# tiedb.com — build and maintenance
#
# The site itself needs only Hugo. The tools need Go and the sibling repos.

HUGO   ?= hugo
SRC    ?= ../..
DOMAIN ?= tiedb.com

.PHONY: all dev build clean logos facts docs docs-check demo shots shots-review shots-accept check update publish

all: build

## Serve locally with live reload.
dev:
	$(HUGO) server --disableFastRender --port 1313

## A clean, minified production build.
build:
	$(HUGO) -e production --gc --minify --cleanDestinationDir

clean:
	rm -rf public resources .cache site.tar.gz

## Logo family, favicons, brand assets and social-card bases.
logos:
	cd tools && go run ./cmd/sitegen logos -root ..

## Refresh data/facts.yaml from the sibling checkouts.
facts:
	cd tools && go run ./cmd/sitegen facts -root .. -src $(SRC)

## Sync the long-form source notes into /docs/reference/.
docs:
	cd tools && go run ./cmd/sitegen docs -root .. -src $(SRC)

docs-check:
	cd tools && go run ./cmd/sitegen docs -root .. -src $(SRC) --check

## Deterministic demo library: public-domain photos and twelve fictional albums.
demo:
	./scripts/demo-media.sh

## Capture every app (headless Fyne), then accept the shots into assets/shots.
shots:
	./scripts/capture-all.sh

shots-review:
	./scripts/shots-review.sh

shots-accept:
	./scripts/shots-accept.sh

## The quality gates: build, links, alt text, budgets, privacy, CSP rules.
check: build
	cd tools && go run ./cmd/sitegen check -root ..

## Release-day refresh. Review the diff before committing.
update: facts docs logos demo shots check

# GitHub Pages deploys on push to main via .github/workflows/pages.yml.
# `publish` is the older SourceHut Pages path and is kept as a fallback.
publish:
	./scripts/publish.sh $(DOMAIN)
