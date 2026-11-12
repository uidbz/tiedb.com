#!/usr/bin/env bash
# Run the whole screenshot pipeline:
#   1. build the demo library (public-domain photos + fictional albums)
#   2. start a throwaway tie environment and seed it
#   3. capture each app through its headless Fyne harness (dark + light)
#   4. accept the raw PNGs as WebP masters under assets/shots
#
# The capture tests are env-gated and live in the app repositories; see
# scripts/README.md. Review the diff before committing.
set -euo pipefail
cd "$(dirname "$0")/.."

./scripts/demo-media.sh
./scripts/tie-env.sh start --reset
trap './scripts/tie-env.sh stop' EXIT

./scripts/capture-tie-gui.sh
./scripts/capture-anneal.sh
./scripts/shots-accept.sh
