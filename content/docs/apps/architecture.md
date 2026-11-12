---
title: Architecture
type: docs
weight: 30
---

The five apps share more than a look. The repository is one Go module with
shared packages and one `cmd/` directory per app.

## Shared packages

| Package | Used by | What it is |
|---|---|---|
| `gallery/` | imgview, tie-view, tie-fm preview, tie-audio cover wall | Justified layout, tiles, pagination, image view, drawer, filter chips |
| `tagselection/` | tie-view, tie-audio | The tag picker, backed by a trie |
| `mpvplayer/` | imgview, tie-view, tie-fm, tie-audio | libmpv video widget with a stub build tag |
| `tieconfig/` | all | Config path resolution, the connection editor, per-app collection choice |
| `tiethumb/` | tie-view, tie-fm, tie-audio | Filehost-backed thumbnails |
| `tiestyle/` | all | The theme: surfaces, accent, contrast tests |
| `internal/tietest/` | tests | A test config against the local environment |

The apps talk to tie through `github.com/uidbz/tie/client`. The standard config
is `~/.config/tie/config.toml`; the client opens the triplestore over HTTP and
the filehost over HTTP, with the same code that the CLI uses.

## Fyne fork

`third_party/fyne` is a submodule of `uidbz/fyne` on the `imgview` branch. It
adds:

- `canvas.GLVideo` for mpv video surfaces;
- Android system-bar control and a media-session driver;
- text-shaping fixes used by the gallery labels.

Builds use a `replace` directive, which is why `go install …@latest` does not
work for these apps; build from a checkout.

## Concurrency

The Fyne UI is single-threaded. The apps do network, disk and decoding work on
goroutines and hand results back through `fyne.Do` (or the fynerisor UI thread
guard for tie-browser). The gallery has a bounded loader pool and a debounced
tile updater; tie-audio has a UI-thread test driver that models the real
compositor.

## Caching

- Thumbnails: server-side (filehost) plus a local disk cache.
- Covers: decoded and cached in tie-audio, capped at 512 px for the wall.
- Video: frame extraction for tiles, not a full decode.
- Images: the gallery keeps a bounded tile cache and releases off-screen views.

`docs/OPTIMIZATIONS.md` has the measurements and the mobile tuning.

## Build tags

| Tag | Effect |
|---|---|
| `nompv` | Remove libmpv and video surfaces. |
| `nomtp` | Remove libmtp/MTP support from tie-fm. |
| `mobile` | The mobile device profile (used by the Android build). |

The shared packages compile on Linux, Windows, macOS and Android; the device
manager and PipeWire are behind build tags.
