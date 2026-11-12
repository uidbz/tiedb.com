#!/usr/bin/env bash
# Capture the tie-gui apps through their env-gated headless tests.
#
# The tests live in the app repositories (cmd/<app>/site_capture_test.go for
# imgview, cmd/tie-audio/internal/ui/site_capture_test.go for tie-audio). They
# are skipped unless the SITE_CAPTURE_DIR gate is set, so normal test runs and
# CI are unaffected.
set -euo pipefail
cd "$(dirname "$0")/.."
ROOT=$PWD
CACHE=${CACHE:-$ROOT/.cache}
GUI=${TIE_GUI_SRC:-$ROOT/../tie-gui}
OUT=$CACHE/shots/raw
mkdir -p "$OUT"

if [ ! -d "$GUI" ]; then
  echo "capture-tie-gui: $GUI not found (set TIE_GUI_SRC)" >&2
  exit 1
fi

have_test() { [ -f "$1" ]; }

capture_imgview() {
  local testfile=$GUI/cmd/imgview/site_capture_test.go
  if ! have_test "$testfile"; then
    echo "capture-tie-gui: skipping imgview (site_capture_test.go missing)" >&2
    return
  fi
  for theme in dark light; do
    ( cd "$GUI" && \
      XDG_CONFIG_HOME=$(mktemp -d) \
      SITE_DEMO_IMAGES=$CACHE/demo/images \
      SITE_CAPTURE_DIR=$OUT/imgview \
      TIE_THEME=$theme \
      go test ./cmd/imgview -run TestSiteCapture -count=1 )
  done
}

capture_audio() {
  local testfile=$GUI/cmd/tie-audio/internal/ui/site_capture_test.go
  if ! have_test "$testfile"; then
    echo "capture-tie-gui: skipping tie-audio (site_capture_test.go missing)" >&2
    return
  fi
  local cfg=$CACHE/tie-env/site-tie.toml
  if [ ! -f "$cfg" ]; then
    echo "capture-tie-gui: $cfg missing; run ./scripts/tie-env.sh start" >&2
    return
  fi
  for theme in dark light; do
    ( cd "$GUI" && \
      FILESDIR=$(mktemp -d) \
      SITE_TIE_CONFIG=$cfg \
      SITE_CAPTURE_DIR=$OUT/tie-audio \
      TIE_THEME=$theme \
      go test ./cmd/tie-audio/internal/ui -run TestSiteCapture -count=1 )
  done
}

capture_imgview
capture_audio
echo "capture-tie-gui: raw captures in $OUT"
