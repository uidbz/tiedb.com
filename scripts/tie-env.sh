#!/usr/bin/env bash
# Start a throwaway tie test environment for the screenshot pipeline.
#
#   ./scripts/tie-env.sh start [--reset]   build + start + import the demo albums
#   ./scripts/tie-env.sh stop
#   ./scripts/tie-env.sh status
#
# Everything lives under .cache/tie-env; the sibling ../tie test-env is only
# copied from, never modified. The result is a config at
# .cache/tie-env/site-tie.toml for the capture tests.
set -euo pipefail
cd "$(dirname "$0")/.."

ROOT=$PWD
CACHE=${CACHE:-$ROOT/.cache}
ENV=$CACHE/tie-env
DEMO=$CACHE/demo
TIE_SRC=${TIE_SRC:-$ROOT/../tie}
TIE_SRC=$(cd "$TIE_SRC" && pwd)

usage() { echo "usage: $0 start [--reset] | stop | status" >&2; exit 2; }

start() {
  local reset=${1:-}
  if [ ! -d "$TIE_SRC/test-env" ]; then
    echo "tie-env: $TIE_SRC/test-env not found (set TIE_SRC)" >&2
    exit 1
  fi
  if [ ! -d "$DEMO/music" ]; then
    echo "tie-env: run ./scripts/demo-media.sh first" >&2
    exit 1
  fi
  if [ "$reset" = "--reset" ] && [ -d "$ENV" ]; then
    ( cd "$ENV" && ./stop.sh >/dev/null 2>&1 || true )
    rm -rf "$ENV"
  fi
  if [ ! -x "$ENV/bin/tie-triplestore" ]; then
    mkdir -p "$ENV"
    cp "$TIE_SRC/test-env"/{build.sh,start.sh,stop.sh,env.sh,config.toml} "$ENV/"
    ( cd "$ENV" && TIE_SRC="$TIE_SRC" ./build.sh )
  fi
  ( cd "$ENV" && TIE_SRC="$TIE_SRC" ./start.sh )

  cat > "$ENV/site-tie.toml" <<'EOF'
Username = "defaultuser"
Password = "defaultpassword"
Namespace = "Collections"
Collection = "Main"
DefaultCollection = "Main"
TripleStoreURL = "http://localhost:2161"
DefaultFileHosts = ["default"]
[FileHosts.default]
URL = "http://localhost:2162"
EOF

  # Import once: a tag list containing "demo" means the library is seeded.
  if ! ( cd "$ENV" && ./bin/tie -C config.toml tag list 2>/dev/null | grep -qx demo ); then
    echo "tie-env: importing the demo albums"
    for dir in "$DEMO"/music/*/*; do
      [ -d "$dir" ] || continue
      artist=$(basename "$(dirname "$dir")")
      album=$(basename "$dir")
      flac=$(find "$dir" -maxdepth 1 -name '*.flac' | sort | head -1)
      genre=$(ffprobe -v quiet -show_entries format_tags=genre -of default=nw=1:nk=1 "$flac" 2>/dev/null || true)
      [ -n "$genre" ] || genre=Jazz
      ( cd "$ENV" && ./bin/tie -C config.toml import audio-dir "$dir" \
          --tags demo --tags "$genre" --dest "/music/$artist/$album" >/dev/null )
    done
  fi
  echo "tie-env: ready at http://localhost:2161 (config: $ENV/site-tie.toml)"
}

stop() {
  if [ -d "$ENV" ]; then
    ( cd "$ENV" && ./stop.sh ) || true
  fi
}

status() {
  local code
  code=$(curl -s -o /dev/null -w '%{http_code}' "http://localhost:2161/-/version" || true)
  case "$code" in
    200|401) echo "tie-env: running (HTTP $code)" ;;
    *) echo "tie-env: stopped" ;;
  esac
}

case "${1:-}" in
  start) start "${2:-}" ;;
  stop) stop ;;
  status) status ;;
  *) usage ;;
esac
