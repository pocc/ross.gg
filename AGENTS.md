# AGENTS.md

Notes for working on this repo. Keep this file focused: facts that aren't obvious from the code, not narration.

## Image pipeline: iPhone HEIC → WebP for the web

Travel and blog photos come from iPhone as HEIC. For the web we convert to WebP (AVIF optional, see below). The non-obvious part is the **tonemap**, not the encoder settings.

### The gain-map problem

iPhone HEICs since ~iPhone 12 store an SDR base image plus an HDR gain map. Photos.app, Preview, Finder QuickLook, and Safari 26+ apply the gain map at display time — that's why the HEIC looks bright and punchy in Preview. **No CLI tool in 2026 applies the gain map on export** (verified: ImageMagick, `sips`, `heif-convert`, `qlmanage`, and Swift via `kCGImageSourceDecodeToHDR` into 16-bit TIFF all return the SDR base). The converted WebP comes out noticeably darker than the HEIC looks in Preview.

To compensate, apply a gentle **gamma + soft S-curve + saturation lift** during conversion. Anything more aggressive (e.g. `-modulate 115,110`) clips white roofs and skies — 10%+ of pixels lost to pure white in testing.

### Default command

```bash
magick input.heic \
  -resize '1200x>' \
  -colorspace sRGB \
  -gamma 1.08 \
  -sigmoidal-contrast 2x50% \
  -modulate 100,108 \
  -quality 82 -define webp:method=6 \
  +profile exif +profile xmp +profile iptc +profile 8bim \
  output.webp
```

- `-resize '1200x>'` — cap width at 1200 (retina-ready for blog column), shrink-only, preserves aspect
- `-colorspace sRGB` — avoid browser ICC v4 / Display P3 rendering inconsistency
- `-gamma 1.08` — lift midtones without clipping highlights (asymptotic at white)
- `-sigmoidal-contrast 2x50%` — soft S-curve for perceived contrast
- `-modulate 100,108` — +8% saturation, no brightness change (gamma handles that)
- `-quality 82` — photo sweet spot; q78 is 15% smaller with minimal quality loss
- `+profile exif/xmp/iptc/8bim` — strip GPS/camera metadata; **keep ICC** (Display P3 stays intact for wide-gamut browsers)
- No `-strip` — that also removes ICC, causing muted reds/greens on wide-gamut Apple displays

### Script

`scripts/heic-to-webp.sh` runs this in parallel over a source folder. Usage:

```bash
post_name="{post_name}"
./scripts/heic-to-webp.sh --source "/path/2/source" --dest "./src/assets/images/${post_name}"
```

### AVIF: not worth it at blog resolution

Tested extensively: at 1200px, q75–82, AVIF is **10–25% *larger* than WebP** for the same visual quality. AVIF's compression advantage only shows at >2000px or aggressive quality (q50–65). For a standard blog post, stick with WebP.

If you want broader-format `<picture>` markup anyway:

```html
<picture>
  <source srcset="/img/photo.avif" type="image/avif">
  <img src="/img/photo.webp" alt="..." loading="lazy" width="1200" height="800">
</picture>
```

### HDR preservation — not solved

Goal: make Safari 26+ users see HDR gain-map rendering. Paths investigated and rejected:

- **`<picture>` with HEIC source**: `sips` resize strips the `tmap` auxiliary image (gain map). Shipping originals at 1.5+ MB/photo is too heavy.
- **HDR AVIF (HLG/PQ)**: requires extracting HDR pixels from HEIC, which magick/libvips/sips don't do. A Swift helper using `kCGImageSourceDecodeToHDR` into float TIFF would work but is non-trivial.
- **Ultra HDR JPG via `toGainMapHDR`**: requires Xcode build of a third-party tool.

Conclusion: ship SDR WebP with the gamma tonemap. Most readers are on SDR displays anyway, and on HDR displays the WebP is only slightly flatter than Preview's HEIC render — not worth the complexity budget.

### Lazy loading in Astro

Images in `public/` are served verbatim; use native `loading="lazy"`:

```astro
<img
  src="/img/trip/photo.webp"
  alt="..."
  width="1200" height="800"
  loading={index === 0 ? 'eager' : 'lazy'}
  decoding="async"
/>
```

The hero/first image gets `eager` + `fetchpriority="high"` — lazy-loading it hurts LCP. Always set `width`/`height` to prevent layout shift.

Images in `src/assets/` can go through Astro's `<Image>` component (lazy by default, auto srcset), but `astro:assets` doesn't process HEIC — always feed it WebP.
