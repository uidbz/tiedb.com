# Screenshot and demo pipeline

The point of this directory is that the site's media is reproducible. No
screenshot is hand-edited; each one comes from a real app, a known demo library
and a headless render.

## The flow

```
scripts/demo-media.sh       public-domain photos + 12 fictional albums
scripts/tie-env.sh start    throwaway triplestore/filehost on 2161/2162
scripts/capture-all.sh      runs the app harnesses, then shots-accept
```

`make shots` is the same as the last line.

### Demo library

`.cache/demo/images` holds sixteen public-domain works from Wikimedia Commons,
with `credits.json`. `.cache/demo/music` holds twelve albums whose tracks are
silent FLACs with real tags, durations and generated covers. The names are
fictional (The Annealers, Ostrander, Hana Sato Trio, …), so nothing private
reaches a screenshot.

Re-running is cheap; `FORCE=1` rebuilds everything.

### tie environment

`.cache/tie-env` is a copy of `../tie/test-env`'s scripts, built from the
sibling checkout and run with its data in the cache. The sibling `test-env/db`
and `test-env/data` are never touched. The demo albums are imported as
`audio-dir` with a `demo` tag and their genre; the script writes
`site-tie.toml` for the capture tests.

### Capture tests

Each app owns its harness. They are ordinary Go tests that skip unless
`SITE_CAPTURE_DIR` is set:

| App | File | Scenes |
|---|---|---|
| imgview | `cmd/imgview/site_capture_test.go` | gallery, image, info |
| tie-audio | `cmd/tie-audio/internal/ui/site_capture_test.go` | home, wall, album, compact |
| anneal | `cmd/anneal/capture_test.go` | chat, harness, activity, settings, empty, narrow |

The tests use Fyne's software driver at 2x and apply `tiestyle` so the captures
match the real app. Runtime is about 7 s (imgview), 18 s (tie-audio, both
themes) and a few seconds (anneal).

If a test file is missing, `capture-tie-gui.sh` skips that app with a warning.
The site falls back to a labelled placeholder for a missing capture.

### Accept and review

`assets/shots/<app>/<scene>-<theme>.webp` are the committed masters, produced
from `.cache/shots/raw/**.png` by `shots-accept.sh` (cwebp q90). Hugo resizes
them to 1280/1920 and serves a `srcset`.

`make shots-review` writes `.cache/shots/review.html` with the current asset
next to the new raw capture. Check it before accepting, then commit the WebP
changes.

### Adding an app

1. Add a scene to the app's capture test that writes `<scene>-<theme>.png`.
2. Add `hero_shot` or a `{{< shot >}}` shortcode to the app page.
3. Run `make shots`.

The shot shortcode searches `assets/shots/<app>/<scene>-<theme>.webp` and then
`.png`, so both raw and processed assets work.

## Known capture quirks

- **anneal's session sidebar.** The harness seeds a session with a title, but
  the sidebar row is built from the backend's session list and still says
  "Untitled session". The checked-in `chat-dark.webp` has the row title patched.
  Seed the session list in the harness before regenerating, or expect to patch
  that one label.
- **imgview filenames.** The checked-in gallery was captured before the demo
  image names were cleaned up, so the labels carry `13-`, `14-` prefixes. A
  fresh capture from `scripts/demo_media.py` uses the clean names.
- **The apps still to capture.** tie-fm, tie-view and tie-browser have no
  `site_capture_test.go` yet. Their pages use diagrams and the app chrome rather
  than a placeholder screenshot. The `shot` shortcode will render a labelled
  placeholder as soon as a scene is referenced but not captured.
