#!/usr/bin/env bash
# Build a side-by-side review sheet for the latest raw captures.
# Open .cache/shots/review.html in a browser, then run `make shots-accept`.
set -euo pipefail
cd "$(dirname "$0")/.."
RAW=.cache/shots/raw
OUT=.cache/shots/review.html
mkdir -p .cache/shots

if [ ! -d "$RAW" ]; then
  echo "shots-review: no raw captures at $RAW" >&2
  exit 1
fi

{
  echo '<!doctype html><meta charset="utf-8"><title>capture review</title>'
  echo '<style>body{margin:0;padding:32px;background:#17181b;color:#e6e7ea;font:14px system-ui;background:#17181b}'
  echo 'h1{font-size:22px;margin:0 0 24px} h2{font-size:15px;margin:32px 0 10px;color:#9a9ca3;text-transform:uppercase;letter-spacing:.12em}'
  echo '.row{display:grid;grid-template-columns:1fr 1fr;gap:16px;margin-bottom:16px;align-items:start}'
  echo '.cell{border:1px solid #2e3036;border-radius:10px;overflow:hidden;background:#1c1d21}'
  echo '.cell img{display:block;width:100%;height:auto} .label{padding:6px 10px;font-size:11px;color:#9a9ca3;border-bottom:1px solid #2e3036}'
  echo '</style><h1>Capture review</h1><p style="color:#9a9ca3">Left: the current site asset. Right: the new raw capture.</p>'
  for app in "$RAW"/*; do
    [ -d "$app" ] || continue
    a=$(basename "$app")
    echo "<h2>$a</h2>"
    for img in "$app"/*.png; do
      [ -f "$img" ] || continue
      n=$(basename "$img" .png)
      old="../../assets/shots/$a/$n.webp"
      [ -f "assets/shots/$a/$n.webp" ] || old=$img
      echo "<div class='row'><div class='cell'><div class='label'>current · $n</div><img src=\"$old\" alt=\"\"></div>"
      echo "<div class='cell'><div class='label'>new · $n</div><img src=\"../../$img\" alt=\"\"></div></div>"
    done
  done
} > "$OUT"

echo "shots-review: wrote $OUT"
echo "open: file://$PWD/$OUT"
