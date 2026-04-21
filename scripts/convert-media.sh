#!/usr/bin/env bash
# Convert JPEG → WebP and MOV → WebM recursively under a target folder.
# Originals are preserved. Pass --delete to remove originals after a successful conversion.
# Usage: ./convert-media.sh [--delete] [path]

set -euo pipefail

DELETE=0
TARGET="/Users/rj/pCloud Drive/Automatic Upload"

for arg in "$@"; do
  case "$arg" in
    --delete) DELETE=1 ;;
    -h|--help)
      sed -n '2,5p' "$0"; exit 0 ;;
    *) TARGET="$arg" ;;
  esac
done

if [[ ! -d "$TARGET" ]]; then
  echo "Error: directory not found: $TARGET" >&2
  exit 1
fi

command -v cwebp  >/dev/null || { echo "cwebp not installed (brew install webp)"  >&2; exit 1; }
command -v ffmpeg >/dev/null || { echo "ffmpeg not installed (brew install ffmpeg)" >&2; exit 1; }

# Parallelism — leave one core free.
JOBS=$(( $(sysctl -n hw.ncpu) - 1 ))
(( JOBS < 1 )) && JOBS=1

convert_jpeg() {
  local src="$1"
  local dst="${src%.*}.webp"
  if [[ -f "$dst" ]]; then
    echo "skip (exists): $dst"
    return
  fi
  # Quality-first WebP:
  # -q 92          high quality; visually indistinguishable from source for photos
  # -m 6           slowest/best compression effort
  # -sharp_yuv     better RGB→YUV conversion, preserves fine color detail
  # -pre 4         spatial noise shaping preprocessor for cleaner output
  # -metadata all  preserve EXIF/ICC/XMP
  # -mt            multithreaded encode
  if cwebp -quiet -q 92 -m 6 -sharp_yuv -pre 4 -mt -metadata all "$src" -o "$dst"; then
    echo "jpeg→webp: $src"
    (( DELETE )) && rm -f "$src"
  else
    echo "FAILED jpeg: $src" >&2
    rm -f "$dst"
  fi
}

convert_mov() {
  local src="$1"
  local dst="${src%.*}.webm"
  if [[ -f "$dst" ]]; then
    echo "skip (exists): $dst"
    return
  fi
  # Quality-first VP9, two-pass constant-quality (Google's recommended workflow):
  # -crf 24 -b:v 0   high quality CQ mode (lower = better; 24 ≈ visually lossless-ish for most footage)
  # -deadline good -cpu-used 0   slowest/best analysis (much slower than cpu-used 2, meaningfully better)
  # -row-mt 1 -tile-columns 2 -threads 0   multithread without hurting quality
  # -g 240 -keyint_min 240   2× framerate keyframe interval — standard for VOD
  # -auto-alt-ref 1 -lag-in-frames 25   enables alt-ref frames for higher quality
  # -pix_fmt yuv420p   broadest browser compatibility
  # libopus 192k       audiophile-level audio
  local pass_log
  pass_log="$(mktemp -t vp9pass.XXXXXX)"
  if ffmpeg -hide_banner -loglevel error -y -i "$src" \
       -c:v libvpx-vp9 -pass 1 -passlogfile "$pass_log" \
       -crf 24 -b:v 0 \
       -deadline good -cpu-used 4 \
       -row-mt 1 -tile-columns 2 -threads 0 \
       -g 240 -keyint_min 240 \
       -auto-alt-ref 1 -lag-in-frames 25 \
       -pix_fmt yuv420p -an -f null /dev/null \
  && ffmpeg -hide_banner -loglevel error -y -i "$src" \
       -c:v libvpx-vp9 -pass 2 -passlogfile "$pass_log" \
       -crf 24 -b:v 0 \
       -deadline good -cpu-used 0 \
       -row-mt 1 -tile-columns 2 -threads 0 \
       -g 240 -keyint_min 240 \
       -auto-alt-ref 1 -lag-in-frames 25 \
       -pix_fmt yuv420p \
       -c:a libopus -b:a 192k \
       "$dst"; then
    echo "mov→webm:  $src"
    (( DELETE )) && rm -f "$src"
  else
    echo "FAILED mov: $src" >&2
    rm -f "$dst"
  fi
  rm -f "${pass_log}"*.log 2>/dev/null || true
}

export -f convert_jpeg convert_mov
export DELETE

echo "Target:   $TARGET"
echo "Delete originals: $([[ $DELETE -eq 1 ]] && echo yes || echo no)"
echo "Parallel jobs: $JOBS"
echo

# Images — parallel (cheap, I/O-bound-ish)
find "$TARGET" -type f \( -iname '*.jpg' -o -iname '*.jpeg' \) -print0 \
  | xargs -0 -n1 -P "$JOBS" -I{} bash -c 'convert_jpeg "$@"' _ {}

# Videos — serial (VP9 encode already saturates CPU)
find "$TARGET" -type f -iname '*.mov' -print0 \
  | while IFS= read -r -d '' f; do convert_mov "$f"; done

echo
echo "Done."
