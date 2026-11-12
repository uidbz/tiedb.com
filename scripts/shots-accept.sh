#!/usr/bin/env bash
# Copy raw capture PNGs into the site as WebP masters.
#
# Raw:   .cache/shots/raw/<app>/<scene>-<theme>.png
# Final: assets/shots/<app>/<scene>-<theme>.webp
#
# WebP q90 keeps screenshots visually lossless at a fraction of the size. The
# site then resizes them further through Hugo's image pipeline.
set -euo pipefail
cd "$(dirname "$0")/.."
RAW=${RAW:-.cache/shots/raw}
DEST=assets/shots

if ! command -v cwebp >/dev/null; then
  echo "shots-accept: cwebp (libwebp) is required" >&2
  exit 1
fi
if [ ! -d "$RAW" ]; then
  echo "shots-accept: no raw captures at $RAW; run make shots" >&2
  exit 1
fi

count=0
while IFS= read -r -d '' src; do
  rel=${src#"$RAW"/}
  app=${rel%%/*}
  name=$(basename "$rel" .png)
  mkdir -p "$DEST/$app"
  out=$DEST/$app/$name.webp
  cwebp -quiet -q 90 -m 6 "$src" -o "$out"
  count=$((count + 1))
done < <(find "$RAW" -type f -name '*.png' -print0 | sort -z)

echo "shots-accept: wrote $count WebP masters to $DEST"
git status --short "$DEST" | head -40 || true
