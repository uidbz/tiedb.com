# tiedb.com

The website for **tie**, the tie apps and **anneal**.

Hugo + Hextra, no JavaScript framework, no web fonts, no CDNs. The site builds
with plain `hugo`; the maintenance tools refresh the generated parts.

## Requirements

| Task | Needs |
|---|---|
| `make dev`, `make build` | Hugo 0.146+ extended |
| `make logos` | Go 1.27+, `rsvg-convert` |
| `make facts`, `make docs`, `make check` | Go 1.27+, the sibling checkouts in `..` |
| `make demo` | Python 3.11+, ffmpeg, ImageMagick |
| `make shots` | the above, a CGO Fyne toolchain, and the capture tests in the app repos |
| `make publish` | `hut` with a SourceHut Pages site for the domain |

## Common commands

```bash
make dev        # live-reloading preview on :1313
make build      # clean production build into public/
make logos      # regenerate marks, favicons, brand PNGs and social cards
make facts      # refresh versions/dates/commit counts from ../tie, ../tie-gui, ../anneal
make docs       # sync the long-form source notes into /docs/reference/
make demo       # build the demo photo/music library under .cache/demo
make shots      # capture the apps, then accept the screenshots
make check      # build + links + alt text + budgets + privacy + CSP rules
make publish    # build, tar and publish to tiedb.com
```

`make update` runs facts, docs, logos, demo, shots and check. Review the diff
before committing; the generated content is committed so the site builds
without the sibling repositories.

## Layout

```
content/            pages: the landing copy, showcase pages, docs, stories
data/apps.yaml      the curated app registry (one source of truth)
data/facts.yaml     generated versions and dates
data/icons.yaml     generated mark icons for Hextra
assets/logos/       generated SVG marks
assets/shots/       WebP screenshot masters (generated, committed)
assets/css/custom.css
layouts/            home, showcase, app, brand, docs hub, 404, shortcodes
tools/              the sitegen binary (logos, facts, docs sync, check)
scripts/            demo data, tie test-env, captures, review, publish
```

## Screenshots

The screenshots are real: each app is built on Fyne's software test driver and
captured at 2x. The harness is a `TestSiteCapture` in the app repository, gated
by `SITE_CAPTURE_DIR`, so it is skipped in normal test runs.

Currently captured:

- `imgview` — `cmd/imgview/site_capture_test.go`
- `tie-audio` — `cmd/tie-audio/internal/ui/site_capture_test.go`
- `anneal` — `cmd/anneal/capture_test.go` (already in the repo)

Still to add: `tie-fm`, `tie-view` and `tie-browser`. Their pages use the
designed placeholder until a capture lands. The `shot` shortcode renders the
placeholder automatically when the asset is missing, so the site never has a
broken image.

`scripts/capture-all.sh` builds the demo library, starts a throwaway tie
environment on 2161/2162, imports twelve fictional albums, runs the captures
and converts the PNGs to WebP. See `scripts/README.md`.

## Privacy

All screenshots use synthetic data and the sandbox path `/home/you`. The
`make check` task fails if a private path or a `target="_blank"` link appears in
the output. Do not publish a capture of a real library.

## Deploying to GitHub Pages

`.github/workflows/pages.yml` builds and deploys on every push to `main`.

1. In the GitHub repo: **Settings → Pages → Build and deployment → Source:
   GitHub Actions**.
2. Push to `main`, or run the workflow manually. The site is served at
   `https://uidbz.github.io/tiedb.com/` until a custom domain is set.
3. The workflow builds with the Pages base URL and runs
   `scripts/rewrite-base.py`, so root-absolute content links work under the
   project subpath too.
4. For the custom domain, see below; the next deployment picks up
   `tiedb.com` as the base URL.

`make publish` is the older SourceHut Pages path and is kept as a fallback.

## Custom domain

The canonical domain is `tiedb.com`. At the DNS provider, replace the current
records (they point at SourceHut) with GitHub Pages:

| Type | Name | Value |
|---|---|---|
| A | `@` | `185.199.108.153` |
| A | `@` | `185.199.109.153` |
| A | `@` | `185.199.110.153` |
| A | `@` | `185.199.111.153` |
| AAAA | `@` | `2606:50c0:8000::153` |
| AAAA | `@` | `2606:50c0:8001::153` |
| AAAA | `@` | `2606:50c0:8002::153` |
| AAAA | `@` | `2606:50c0:8003::153` |
| CNAME | `www` | `uidbz.github.io` |

Then in **Settings → Pages → Custom domain**, enter `tiedb.com` and save. Wait
for GitHub's DNS check, then enable **Enforce HTTPS** (certificate issuance can
take a few minutes). If GitHub shows a `_github-pages-challenge-uidbz` TXT
record, add it at the registrar. If the domain is behind Cloudflare, set the
proxy to **DNS only** while the certificate is provisioned.

DNS has to be changed at the registrar; the old SourceHut records
(`46.23.81.157` and the parking address on `www`) must go. Once the domain
resolves to GitHub Pages, `make publish`/`hut` no longer serves this site.

## Domain config

The canonical domain is a config value: `baseURL` in `hugo.yaml`. The CI build
overrides it with the Pages URL (or the custom domain, once set).
