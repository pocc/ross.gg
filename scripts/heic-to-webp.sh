#!/usr/bin/env bash
# Copy HEIC files from a source folder into a destination, then convert each
# to WebP (1200px long edge, q78). Strips EXIF/GPS/XMP/IPTC but preserves the
# ICC color profile so iPhone Display P3 photos render correctly.
# Usage: ./heic-to-webp.sh [--source PATH] [--dest PATH]

set -euo pipefail

SOURCE="/Users/rj/pCloud Drive/Automatic Upload/Rocinante/2026/04"
DEST="/private/tmp/photos"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --source) SOURCE="$2"; shift 2 ;;
    --dest)   DEST="$2";   shift 2 ;;
    -h|--help) sed -n '2,5p' "$0"; exit 0 ;;
    *) echo "Unknown arg: $1" >&2; exit 1 ;;
  esac
done

[[ -d "$SOURCE" ]] || { echo "Error: source not found: $SOURCE" >&2; exit 1; }
mkdir -p "$DEST"

command -v magick >/dev/null || { echo "magick not installed (brew install imagemagick)" >&2; exit 1; }

JOBS=$(( $(sysctl -n hw.ncpu) - 1 ))
(( JOBS < 1 )) && JOBS=1

convert_one() {
  local src="$1"
  local dest="$2"
  local base; base="$(basename "$src")"
  local heic_copy="$dest/$base"
  local webp="$dest/${base%.*}.webp"

  # Copy HEIC (idempotent)
  if [[ ! -f "$heic_copy" ]]; then
    cp -p "$src" "$heic_copy"
  fi

  # Convert + tonemap + strip non-ICC metadata in one magick pass.
  # Tonemap compensates for the HDR gain map that Apple's renderers apply at
  # display time but which no CLI decoder preserves — without it, the WebP
  # looks notably darker than the HEIC in Photos/Preview.
  #   -resize 1200x>              cap width at 1200, shrink-only
  #   -colorspace sRGB            avoid browser P3/v4 ICC rendering inconsistency
  #   -sigmoidal-contrast 5x45%   S-curve contrast for highlight/shadow separation
  #   -modulate 115,110           +15% brightness, +10% saturation to approximate HDR pop
  #   -quality 78                 photo sweet spot
  #   -define webp:method=6       slowest/best compression (fine for one-shots)
  #   +profile exif/xmp/iptc/8bim drop privacy-sensitive metadata, keep ICC
  if magick "$heic_copy" \
       -resize '1200x>' \
       -colorspace sRGB \
       -sigmoidal-contrast 5x45% \
       -modulate 115,110 \
       -quality 78 \
       -define webp:method=6 \
       +profile exif \
       +profile xmp \
       +profile iptc \
       +profile 8bim \
       "$webp"; then
    echo "✓ $base → $(basename "$webp")"
  else
    echo "✗ FAILED: $base" >&2
    rm -f "$webp"
  fi
}

export -f convert_one

echo "Source: $SOURCE"
echo "Dest:   $DEST"
echo "Jobs:   $JOBS"
echo

find "$SOURCE" -type f -iname '*.heic' -print0 \
  | xargs -0 -n1 -P "$JOBS" -I{} bash -c 'convert_one "$@"' _ {} "$DEST"

echo
echo "Done."
