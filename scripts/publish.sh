#!/usr/bin/env bash
# Build and publish to SourceHut Pages.
#
#   ./scripts/publish.sh [domain]
#
# Requires hut (https://sr.ht) and a pages site for the domain. The build is
# clean: nothing stale from a previous run can leak into the tarball.
set -euo pipefail
cd "$(dirname "$0")/.."
DOMAIN=${1:-tiedb.com}

if ! command -v hut >/dev/null; then
  echo "publish: hut is required (go install git.sr.ht/~sircmpwn/hut@latest)" >&2
  exit 1
fi

hugo -e production --gc --minify --cleanDestinationDir
tar -C public -czf site.tar.gz .
hut pages publish -d "$DOMAIN" site.tar.gz

echo "publish: checking https://$DOMAIN/"
curl -fsS -o /dev/null -w '  HTTP %{http_code}\n' "https://$DOMAIN/"
