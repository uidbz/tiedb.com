#!/usr/bin/env bash
# Generate the deterministic demo library used by the screenshots.
#
#   .cache/demo/images/   public-domain artworks and NASA photographs
#   .cache/demo/music/    twelve albums of silent FLACs with real tags and covers
#
# Re-running is cheap: existing files are kept unless FORCE=1.
set -euo pipefail
cd "$(dirname "$0")/.."
python3 scripts/demo_media.py
