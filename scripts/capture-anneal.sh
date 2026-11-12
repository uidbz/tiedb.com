#!/usr/bin/env bash
# Capture anneal's UI through the headless Fyne harness in cmd/anneal.
set -euo pipefail
cd "$(dirname "$0")/.."
ROOT=$PWD
CACHE=${CACHE:-$ROOT/.cache}
ANNEAL=${ANNEAL_SRC:-$ROOT/../anneal}
OUT=$CACHE/shots/raw/anneal

if [ ! -d "$ANNEAL" ]; then
  echo "capture-anneal: $ANNEAL not found (set ANNEAL_SRC)" >&2
  exit 1
fi
mkdir -p "$OUT"

for theme in dark light; do
  ( cd "$ANNEAL" && \
    ANNEAL_UI_CAPTURE=$OUT \
    ANNEAL_UI_CAPTURE_THEME=$theme \
    go test ./cmd/anneal -run TestUICapture -count=1 )
done
echo "capture-anneal: raw captures in $OUT"
