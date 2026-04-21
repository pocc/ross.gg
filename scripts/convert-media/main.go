// convert-media recursively converts JPEG → WebP and MOV → {WebM (AV1) | MP4 (HEVC)}
// under a target folder. Tuned for Apple M-series (especially M4 Max).
//
// Design notes:
//   - Images are small and I/O-ish; encode them with libvips (faster + multi-threaded
//     internally). We run `img-workers` processes, each pinned to VIPS_CONCURRENCY=1
//     so parallel processes don't fight over threads.
//   - Videos in SVT-AV1 software mode already use many cores per encode; running
//     `vid-workers` parallel encodes each limited to `ncpu / vid-workers` logical
//     processors gives cleaner scaling than letting each encode grab all cores.
//   - Videos in HEVC hardware mode go through VideoToolbox's media engine, which
//     is a fixed-function unit; 2–3 concurrent encodes saturate it.
//   - Files are staged through a local scratch dir so pCloud's virtual filesystem
//     isn't in the hot loop during encoding. Writes go to local disk, then one
//     big move back at the end.
//   - Output is atomic via rename-on-success; partial files are cleaned up.
//   - Skips if destination already exists → safe to re-run.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type options struct {
	target     string
	deleteOrig bool
	imgWorkers int
	vidWorkers int
	videoMode  string // "av1" or "hevc"
	stageDir   string
	noStage    bool
	webpQ        int
	webpLossless bool
	webpEffort   int
	av1CRF       int
	av1Preset    int
	av1Extra     string
	hevcQ        int
	tenBit       bool
	audioKbps    int
	retries          int
	retryWaitSec     int
	failureLog       string
	cloudConcurrency int
	verbose          bool
}

type stats struct {
	imgDone, imgSkip, imgFail atomic.Int64
	vidDone, vidSkip, vidFail atomic.Int64
	bytesIn, bytesOut         atomic.Int64
}

var (
	opt        options
	st         stats
	imgEncoder string // "vips" or "cwebp"

	// Health-check / outage reporting. Single-shot messaging so workers don't
	// spam the same "waiting" line N times during one pCloud outage.
	healthMu       sync.Mutex
	outageActive   bool
	outageAttempts atomic.Int64 // total retry attempts across workers (for stats)

	// Failure log
	failLogMu sync.Mutex
	failLogF  *os.File

	// Cloud-I/O semaphore — caps concurrent reads/writes against the source
	// filesystem. Set to a buffered channel sized to opt.cloudConcurrency;
	// nil means unlimited.
	cloudIOSem chan struct{}
)

func main() {
	flag.StringVar(&opt.target, "target", "/Users/rj/pCloud Drive/Automatic Upload", "root folder to scan recursively")
	flag.BoolVar(&opt.deleteOrig, "delete", false, "delete originals after successful conversion")
	flag.IntVar(&opt.imgWorkers, "img-workers", 0, "parallel image workers (0 = NumCPU)")
	flag.IntVar(&opt.vidWorkers, "vid-workers", 2, "parallel video workers")
	flag.StringVar(&opt.videoMode, "video", "av1", "video codec: 'av1' (webm, software SVT-AV1) or 'hevc' (mp4, VideoToolbox hardware)")
	flag.StringVar(&opt.stageDir, "stage-dir", "", "local scratch dir (default $TMPDIR)")
	flag.BoolVar(&opt.noStage, "no-stage", false, "don't stage through local disk")
	flag.IntVar(&opt.webpQ, "webp-q", 92, "WebP quality (0–100, ignored if -webp-lossless)")
	flag.BoolVar(&opt.webpLossless, "webp-lossless", false, "lossless WebP (preserves source pixels exactly)")
	flag.IntVar(&opt.webpEffort, "webp-effort", 6, "WebP compression effort (0=fastest ... 6=smallest)")
	flag.IntVar(&opt.av1CRF, "av1-crf", 28, "SVT-AV1 CRF (lower = better; 18 visually lossless, 28 near-lossless)")
	flag.IntVar(&opt.av1Preset, "av1-preset", 6, "SVT-AV1 preset (0=slowest/best ... 13=fastest)")
	flag.StringVar(&opt.av1Extra, "av1-extra", "", "extra svtav1-params (appended to internal defaults)")
	flag.IntVar(&opt.hevcQ, "hevc-q", 55, "HEVC VideoToolbox quality (0–100, higher = better)")
	flag.BoolVar(&opt.tenBit, "10bit", false, "encode video at 10-bit (yuv420p10le) — less banding, slightly slower")
	flag.IntVar(&opt.audioKbps, "audio-kbps", 192, "audio bitrate in kbps (Opus: 96–510; AAC: 96–320)")
	flag.IntVar(&opt.retries, "retries", 8, "retries per file on transient errors (pCloud dropouts)")
	flag.IntVar(&opt.retryWaitSec, "retry-wait", 2, "initial backoff seconds between retries (doubles each attempt, max 30s)")
	flag.StringVar(&opt.failureLog, "failure-log", "convert-media-failures.log", "append failed files + reasons to this path")
	flag.IntVar(&opt.cloudConcurrency, "cloud-concurrency", -1, "max concurrent pCloud I/O ops (stage-in + stage-out). -1=auto (1 for cloud paths, unlimited otherwise). Set 1 to fully serialize.")
	flag.BoolVar(&opt.verbose, "v", false, "verbose output")
	flag.Parse()

	cloudTarget := isCloudPath(opt.target)

	if opt.imgWorkers == 0 {
		// macOS FileProvider-backed mounts (pCloud/iCloud/Dropbox/GDrive/OneDrive)
		// deadlock under heavy parallel reads. We now also serialize the
		// pCloud-facing copy/rename calls (see cloudIOSem), so workers can be
		// as high as we want — they'll queue at the cloud-IO gate.
		if cloudTarget {
			opt.imgWorkers = 8
		} else {
			opt.imgWorkers = runtime.NumCPU()
		}
	}

	// Configure cloud-IO semaphore. Default to fully-serial (1) for cloud
	// paths — a single reader / single writer against FileProvider avoids
	// materialization deadlocks that hit around ~250 parallel reads.
	if opt.cloudConcurrency < 0 {
		if cloudTarget {
			opt.cloudConcurrency = 1
		} else {
			opt.cloudConcurrency = 0 // unlimited
		}
	}
	if opt.cloudConcurrency > 0 {
		cloudIOSem = make(chan struct{}, opt.cloudConcurrency)
	}
	if opt.stageDir == "" {
		opt.stageDir = os.TempDir()
	}
	if opt.videoMode != "av1" && opt.videoMode != "hevc" {
		die("--video must be 'av1' or 'hevc'")
	}

	mustHave("ffmpeg")
	switch {
	case hasBinary("vips"):
		imgEncoder = "vips"
	case hasBinary("cwebp"):
		imgEncoder = "cwebp"
	default:
		die("need either `vips` or `cwebp` installed — try `brew install webp` (fast) or `brew install vips` (faster)")
	}

	if opt.videoMode == "hevc" && !hasEncoder("hevc_videotoolbox") {
		die("hevc_videotoolbox encoder not available in your ffmpeg build")
	}
	if opt.videoMode == "av1" && !hasEncoder("libsvtav1") {
		die("libsvtav1 encoder not available in your ffmpeg build")
	}

	if _, err := os.Stat(opt.target); err != nil {
		die("target not accessible: %v (grant terminal Full Disk Access for pCloud)", err)
	}

	if opt.failureLog != "" {
		f, err := os.OpenFile(opt.failureLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			die("open failure log %s: %v", opt.failureLog, err)
		}
		failLogF = f
		defer f.Close()
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	fmt.Println("Scanning...")
	imgs, vids, err := scan(opt.target)
	if err != nil {
		die("scan: %v", err)
	}

	ext, codec := ".webm", "SVT-AV1 (software, preset "+itoa(opt.av1Preset)+")"
	if opt.videoMode == "hevc" {
		ext, codec = ".mp4", "HEVC (VideoToolbox hardware)"
	}

	fmt.Printf("Target:         %s\n", opt.target)
	fmt.Printf("Images found:   %d\n", len(imgs))
	fmt.Printf("Videos found:   %d\n", len(vids))
	imgDesc := fmt.Sprintf("%s (Q=%d)", imgEncoder, opt.webpQ)
	if opt.webpLossless {
		imgDesc = fmt.Sprintf("%s (lossless, effort=%d)", imgEncoder, opt.webpEffort)
	}
	fmt.Printf("Image encoder:  %s\n", imgDesc)
	fmt.Printf("Video encoder:  %s → *%s\n", codec, ext)
	fmt.Printf("Image workers:  %d\n", opt.imgWorkers)
	fmt.Printf("Video workers:  %d\n", opt.vidWorkers)
	cloudDesc := "unlimited"
	if opt.cloudConcurrency > 0 {
		cloudDesc = fmt.Sprintf("%d concurrent (cloud target detected)", opt.cloudConcurrency)
	}
	fmt.Printf("Cloud I/O:      %s\n", cloudDesc)
	fmt.Printf("Staging:        %s\n", stagingDesc())
	fmt.Printf("Delete originals: %v\n", opt.deleteOrig)
	fmt.Println()

	start := time.Now()
	totalImg, totalVid := int64(len(imgs)), int64(len(vids))
	stopProgress := make(chan struct{})
	go progressLoop(totalImg, totalVid, stopProgress)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		processBatch(ctx, imgs, opt.imgWorkers, processImage)
	}()
	go func() {
		defer wg.Done()
		processBatch(ctx, vids, opt.vidWorkers, processVideo)
	}()
	wg.Wait()
	close(stopProgress)

	dur := time.Since(start)
	fmt.Printf("\n\nDone in %s\n", dur.Round(time.Second))
	fmt.Printf("  Images: %d converted, %d skipped, %d failed\n",
		st.imgDone.Load(), st.imgSkip.Load(), st.imgFail.Load())
	fmt.Printf("  Videos: %d converted, %d skipped, %d failed\n",
		st.vidDone.Load(), st.vidSkip.Load(), st.vidFail.Load())
	if in, out := st.bytesIn.Load(), st.bytesOut.Load(); in > 0 {
		fmt.Printf("  Size:   %s → %s (%.1f%%)\n",
			humanBytes(in), humanBytes(out), 100*float64(out)/float64(in))
	}
}

func stagingDesc() string {
	if opt.noStage {
		return "disabled (encode directly in source filesystem)"
	}
	return opt.stageDir
}

func mustHave(bin string) {
	if !hasBinary(bin) {
		die("missing binary: %s — install it first (e.g. `brew install %s`)", bin, bin)
	}
}

func hasBinary(bin string) bool {
	_, err := exec.LookPath(bin)
	return err == nil
}

// isCloudPath heuristically identifies paths served by macOS FileProvider
// (pCloud Drive, iCloud, Dropbox, Google Drive, OneDrive). These paths need
// low read concurrency to avoid EDEADLK materialization deadlocks.
func isCloudPath(p string) bool {
	for _, pat := range []string{
		"/pCloud Drive/", "/pCloudDrive/",
		"/Library/CloudStorage/",
		"/iCloud Drive/", "/Mobile Documents/",
		"/Dropbox/", "/Google Drive/", "/OneDrive/", "/Box/",
	} {
		if strings.Contains(p, pat) {
			return true
		}
	}
	return false
}

// checkTargetHealth stats the target root. On FileProvider dropouts this
// returns an error quickly (permission denied / no such file).
func checkTargetHealth() bool {
	_, err := os.Stat(opt.target)
	return err == nil
}

// waitForTarget blocks until the target directory is accessible again.
// Prints exactly one "paused" message per outage regardless of how many
// workers are waiting.
func waitForTarget(ctx context.Context) error {
	if checkTargetHealth() {
		return nil
	}
	healthMu.Lock()
	if !outageActive {
		outageActive = true
		fmt.Fprintf(os.Stderr, "\n[%s] target unreachable — pCloud likely dropped, pausing workers...\n",
			time.Now().Format("15:04:05"))
	}
	healthMu.Unlock()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
		}
		if checkTargetHealth() {
			healthMu.Lock()
			if outageActive {
				outageActive = false
				fmt.Fprintf(os.Stderr, "[%s] target back online, resuming\n",
					time.Now().Format("15:04:05"))
			}
			healthMu.Unlock()
			return nil
		}
	}
}

// transientPatterns covers errors we see under macOS FileProvider deadlocks
// and routine flaky-cloud-drive behavior.
var transientPatterns = []string{
	"resource deadlock avoided", // EDEADLK — the canonical FileProvider symptom
	"operation not permitted",   // EPERM — pCloud disabled/re-enabling
	"permission denied",         // EACCES
	"no such file or directory", // ENOENT — transient during materialization
	"input/output error",        // EIO
	"i/o error",
	"resource temporarily unavailable", // EAGAIN
	"stale file handle",
	"broken pipe",
	"transport endpoint is not connected",
	"bad file descriptor",
	"device or resource busy",
}

func isTransient(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, syscall.EDEADLK) || errors.Is(err, syscall.EAGAIN) ||
		errors.Is(err, syscall.EACCES) || errors.Is(err, syscall.EPERM) ||
		errors.Is(err, syscall.ENOENT) || errors.Is(err, syscall.EIO) ||
		errors.Is(err, syscall.EBUSY) {
		return true
	}
	s := strings.ToLower(err.Error())
	for _, p := range transientPatterns {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}

// retry executes fn, retrying transient errors with exponential backoff
// and pausing for target-directory availability between attempts.
func retry(ctx context.Context, fn func() error) error {
	backoff := time.Duration(opt.retryWaitSec) * time.Second
	const maxBackoff = 30 * time.Second
	var lastErr error
	for attempt := 0; attempt <= opt.retries; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}
		lastErr = err
		if !isTransient(err) {
			return err
		}
		if attempt == opt.retries {
			break
		}
		outageAttempts.Add(1)
		if werr := waitForTarget(ctx); werr != nil {
			return werr
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
	return fmt.Errorf("after %d retries: %w", opt.retries, lastErr)
}

// acquireCloudIO blocks until we're allowed to touch the source filesystem.
// Respects ctx cancellation. Pairs with releaseCloudIO via defer.
// cloudStat stats a path under the cloud-IO gate. Returns (exists, error);
// treats all transient errors as "not exists" since the file genuinely might
// not be there.
func cloudStat(ctx context.Context, path string) (bool, error) {
	if err := acquireCloudIO(ctx); err != nil {
		return false, err
	}
	defer releaseCloudIO()
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func acquireCloudIO(ctx context.Context) error {
	if cloudIOSem == nil {
		return nil
	}
	select {
	case cloudIOSem <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func releaseCloudIO() {
	if cloudIOSem == nil {
		return
	}
	<-cloudIOSem
}

func logFailure(src string, err error) {
	failLogMu.Lock()
	defer failLogMu.Unlock()
	if failLogF == nil {
		return
	}
	fmt.Fprintf(failLogF, "%s\t%s\t%v\n",
		time.Now().Format(time.RFC3339), src, strings.ReplaceAll(err.Error(), "\n", " "))
	failLogF.Sync()
}

func hasEncoder(name string) bool {
	out, err := exec.Command("ffmpeg", "-hide_banner", "-encoders").CombinedOutput()
	if err != nil {
		return false
	}
	return strings.Contains(string(out), name)
}

func die(f string, a ...any) {
	fmt.Fprintf(os.Stderr, "error: "+f+"\n", a...)
	os.Exit(1)
}

func scan(root string) (imgs, vids []string, err error) {
	err = filepath.WalkDir(root, func(p string, d fs.DirEntry, werr error) error {
		if werr != nil {
			return werr
		}
		if d.IsDir() {
			return nil
		}
		switch strings.ToLower(filepath.Ext(p)) {
		case ".jpg", ".jpeg":
			imgs = append(imgs, p)
		case ".mov":
			vids = append(vids, p)
		}
		return nil
	})
	return
}

func processBatch(ctx context.Context, files []string, workers int, fn func(context.Context, string) error) {
	if len(files) == 0 || workers <= 0 {
		return
	}
	jobs := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for f := range jobs {
				if ctx.Err() != nil {
					return
				}
				if err := fn(ctx, f); err != nil && opt.verbose {
					fmt.Fprintf(os.Stderr, "\n  fail: %s — %v\n", filepath.Base(f), err)
				}
			}
		}(i)
	}
	for _, f := range files {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return
		case jobs <- f:
		}
	}
	close(jobs)
	wg.Wait()
}

func processImage(ctx context.Context, src string) error {
	dst := replaceExt(src, ".webp")
	if exists, _ := cloudStat(ctx, dst); exists {
		st.imgSkip.Add(1)
		return nil
	}

	// Stage source to local disk — reduces pCloud pressure, makes the encode
	// I/O all-local, and lets us retry the one cloud-read cleanly.
	input, cleanupIn, err := stageIn(ctx, src)
	if err != nil {
		logFailure(src, fmt.Errorf("stage in: %w", err))
		st.imgFail.Add(1)
		return err
	}
	defer cleanupIn()

	output, finalize, err := stageOut(dst)
	if err != nil {
		logFailure(src, fmt.Errorf("stage out: %w", err))
		st.imgFail.Add(1)
		return err
	}
	tmp := output

	var cmd *exec.Cmd
	if imgEncoder == "vips" {
		// libvips webpsave (CLI form: `vips webpsave in out --flag value ...`):
		//   --Q <n>              lossy quality (ignored when --lossless)
		//   --lossless           exact pixel preservation
		//   --smart-subsample    equivalent of cwebp -sharp_yuv; keeps saturated edges clean
		//   --effort <0-6>       compression effort; for lossless this affects size only
		// libvips 8.15+ preserves EXIF/ICC/XMP by default.
		vipsArgs := []string{"webpsave", input, tmp, "--effort", itoa(opt.webpEffort)}
		if opt.webpLossless {
			vipsArgs = append(vipsArgs, "--lossless")
		} else {
			vipsArgs = append(vipsArgs, "--Q", itoa(opt.webpQ), "--smart-subsample")
		}
		cmd = exec.CommandContext(ctx, "vips", vipsArgs...)
		cmd.Env = append(os.Environ(), "VIPS_CONCURRENCY=1")
	} else {
		// cwebp fallback:
		//   -lossless         exact pixel preservation
		//   -z <0-9>          lossless compression effort (9 = smallest, slowest)
		//   -q <0-100>        lossy quality; in lossless mode controls speed/size tradeoff
		//   -m 6              lossy encoder effort (0–6, 6 = slowest/smallest)
		//   -sharp_yuv        better RGB→YUV; fixes chroma bleeding on saturated edges
		//   -metadata all     preserve EXIF/ICC/XMP
		//   -mt               multithreaded encode (left on; parallel cwebps share cores fine)
		args := []string{"-quiet", "-metadata", "all", "-mt"}
		if opt.webpLossless {
			z := opt.webpEffort
			if z > 9 {
				z = 9
			}
			args = append(args, "-lossless", "-z", itoa(z))
		} else {
			m := opt.webpEffort
			if m > 6 {
				m = 6
			}
			args = append(args, "-q", itoa(opt.webpQ), "-m", itoa(m), "-sharp_yuv", "-pre", "4")
		}
		args = append(args, input, "-o", tmp)
		cmd = exec.CommandContext(ctx, "cwebp", args...)
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		os.Remove(tmp)
		_ = finalize(false)
		logFailure(src, fmt.Errorf("%s encode: %v: %s", imgEncoder, err, strings.TrimSpace(string(out))))
		st.imgFail.Add(1)
		return fmt.Errorf("%s: %v: %s", imgEncoder, err, strings.TrimSpace(string(out)))
	}
	if err := finalize(true); err != nil {
		logFailure(src, fmt.Errorf("finalize: %w", err))
		st.imgFail.Add(1)
		return err
	}
	track(src, dst)
	st.imgDone.Add(1)
	if opt.deleteOrig {
		_ = retry(ctx, func() error {
			if aerr := acquireCloudIO(ctx); aerr != nil {
				return aerr
			}
			defer releaseCloudIO()
			return os.Remove(src)
		})
	}
	return nil
}

func processVideo(ctx context.Context, src string) error {
	var ext string
	if opt.videoMode == "hevc" {
		ext = ".mp4"
	} else {
		ext = ".webm"
	}
	dst := replaceExt(src, ext)
	if exists, _ := cloudStat(ctx, dst); exists {
		st.vidSkip.Add(1)
		return nil
	}

	input, cleanupIn, err := stageIn(ctx, src)
	if err != nil {
		logFailure(src, fmt.Errorf("stage in: %w", err))
		st.vidFail.Add(1)
		return fmt.Errorf("stage in: %w", err)
	}
	defer cleanupIn()

	output, finalize, err := stageOut(dst)
	if err != nil {
		logFailure(src, fmt.Errorf("stage out: %w", err))
		st.vidFail.Add(1)
		return fmt.Errorf("stage out: %w", err)
	}

	args := []string{
		"-hide_banner", "-loglevel", "error", "-y",
		"-nostdin",
		"-i", input,
	}

	pixFmt := "yuv420p"
	if opt.tenBit {
		pixFmt = "yuv420p10le"
	}
	audioBR := fmt.Sprintf("%dk", opt.audioKbps)

	if opt.videoMode == "hevc" {
		// Hardware HEVC via VideoToolbox — minimal quality drop for ~20–50× speedup
		// on Apple Silicon's media engine. -tag:v hvc1 makes the file play in
		// QuickTime/Safari/Finder (vs hev1 which is more standards-correct but
		// less compatible with Apple tooling).
		// 10-bit HEVC uses yuv420p10le + Main 10 profile, automatically selected.
		args = append(args,
			"-c:v", "hevc_videotoolbox",
			"-q:v", itoa(opt.hevcQ),
			"-tag:v", "hvc1",
			"-pix_fmt", pixFmt,
			"-c:a", "aac", "-b:a", audioBR,
			"-movflags", "+faststart",
			"-f", "mp4",
		)
	} else {
		// Software AV1 via SVT-AV1. Preset sweet spots:
		//   preset 2   near-best quality, very slow
		//   preset 4   quality-focused encodes; ~1% quality loss vs preset 2, ~3–4× faster
		//   preset 6   default balanced; ~3–5× faster than libvpx-vp9 at equivalent quality
		//   preset 8+  real-time territory
		// -svtav1-params lp=<N>: logical processors per encode, so parallel encodes
		//   don't fight over cores.
		lp := runtime.NumCPU() / opt.vidWorkers
		if lp < 1 {
			lp = 1
		}
		svtParams := fmt.Sprintf("tune=0:enable-overlays=1:lp=%d", lp)
		if opt.av1Extra != "" {
			svtParams += ":" + opt.av1Extra
		}
		args = append(args,
			"-c:v", "libsvtav1",
			"-crf", itoa(opt.av1CRF),
			"-preset", itoa(opt.av1Preset),
			"-svtav1-params", svtParams,
			"-g", "240",
			"-pix_fmt", pixFmt,
			"-c:a", "libopus", "-b:a", audioBR,
			"-f", "webm",
		)
	}
	args = append(args, output)

	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		os.Remove(output)
		_ = finalize(false)
		logFailure(src, fmt.Errorf("ffmpeg: %v: %s", err, strings.TrimSpace(string(out))))
		st.vidFail.Add(1)
		return fmt.Errorf("ffmpeg: %v: %s", err, strings.TrimSpace(string(out)))
	}
	if err := finalize(true); err != nil {
		logFailure(src, fmt.Errorf("finalize: %w", err))
		st.vidFail.Add(1)
		return fmt.Errorf("finalize: %w", err)
	}
	track(src, dst)
	st.vidDone.Add(1)
	if opt.deleteOrig {
		_ = retry(ctx, func() error {
			if aerr := acquireCloudIO(ctx); aerr != nil {
				return aerr
			}
			defer releaseCloudIO()
			return os.Remove(src)
		})
	}
	return nil
}

// stageIn copies src to a local scratch path so ffmpeg reads from local disk
// instead of the (possibly slow) source filesystem like pCloud Drive.
// Returns the path to use, and a cleanup func.
func stageIn(ctx context.Context, src string) (string, func(), error) {
	if opt.noStage {
		return src, func() {}, nil
	}
	f, err := os.CreateTemp(opt.stageDir, "cm-in-*"+filepath.Ext(src))
	if err != nil {
		return "", nil, err
	}
	path := f.Name()
	f.Close()
	// Retry the cloud-read under the cloud-IO semaphore: this is where
	// FileProvider deadlocks most often.
	err = retry(ctx, func() error {
		if aerr := acquireCloudIO(ctx); aerr != nil {
			return aerr
		}
		defer releaseCloudIO()
		return copyFile(ctx, src, path)
	})
	if err != nil {
		os.Remove(path)
		return "", nil, err
	}
	return path, func() { os.Remove(path) }, nil
}

// stageOut returns a path ffmpeg should write to, and a finalize func that
// either moves the finished file to dst (on success) or cleans it up.
func stageOut(dst string) (string, func(bool) error, error) {
	if opt.noStage {
		tmp := dst + ".part"
		os.Remove(tmp)
		return tmp, func(ok bool) error {
			if !ok {
				os.Remove(tmp)
				return nil
			}
			return retry(context.Background(), func() error {
				if aerr := acquireCloudIO(context.Background()); aerr != nil {
					return aerr
				}
				defer releaseCloudIO()
				return os.Rename(tmp, dst)
			})
		}, nil
	}
	f, err := os.CreateTemp(opt.stageDir, "cm-out-*"+filepath.Ext(dst))
	if err != nil {
		return "", nil, err
	}
	path := f.Name()
	f.Close()
	os.Remove(path) // encoder will recreate
	return path, func(ok bool) error {
		if !ok {
			os.Remove(path)
			return nil
		}
		defer os.Remove(path)
		// Retry the cloud-write under the cloud-IO semaphore.
		return retry(context.Background(), func() error {
			if aerr := acquireCloudIO(context.Background()); aerr != nil {
				return aerr
			}
			defer releaseCloudIO()
			return moveFile(path, dst)
		})
	}, nil
}

func copyFile(ctx context.Context, src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, in)
	if cerr := out.Close(); err == nil {
		err = cerr
	}
	return err
}

func moveFile(src, dst string) error {
	if err := os.Rename(src, dst); err == nil {
		return nil
	}
	// Cross-device rename fails — fall back to copy+delete.
	if err := copyFile(context.Background(), src, dst); err != nil {
		return err
	}
	return os.Remove(src)
}

func track(src, dst string) {
	// Stat both files under the cloud-IO gate; silent if it times out.
	_ = acquireCloudIO(context.Background())
	if si, err := os.Stat(src); err == nil {
		st.bytesIn.Add(si.Size())
	}
	if di, err := os.Stat(dst); err == nil {
		st.bytesOut.Add(di.Size())
	}
	releaseCloudIO()
}

func replaceExt(p, newExt string) string {
	return strings.TrimSuffix(p, filepath.Ext(p)) + newExt
}

func progressLoop(totalImg, totalVid int64, stop <-chan struct{}) {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-stop:
			return
		case <-t.C:
			iDone := st.imgDone.Load() + st.imgSkip.Load()
			vDone := st.vidDone.Load() + st.vidSkip.Load()
			suffix := ""
			if n := outageAttempts.Load(); n > 0 {
				suffix = fmt.Sprintf("  retries %d", n)
			}
			fmt.Printf("\r  img %d/%d  vid %d/%d  fails %d%s   ",
				iDone, totalImg, vDone, totalVid,
				st.imgFail.Load()+st.vidFail.Load(), suffix)
		}
	}
}

func humanBytes(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for v := n / unit; v >= unit; v /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(n)/float64(div), "KMGTPE"[exp])
}

func itoa(n int) string { return fmt.Sprintf("%d", n) }
