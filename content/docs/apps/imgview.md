---
title: imgview
type: docs
weight: 13
---

A local image and archive viewer that needs no library.

```bash
imgview                       # the current directory
imgview ~/Pictures/2014
imgview page-001.cbz
imgview -config my.toml ~/Pictures
```

## Gallery

- Justified rows that fill the width; pagination at 500 images per page.
- Thumbnails are cached on disk under `ThumbnailDir`.
- Folder tiles show a preview of their contents; swipe (or navigate) through up
  to ten images/video frames.
- `N` toggles filenames, `Backspace` goes to the parent directory.

## Single image

- `←`/`→` navigate, `S` shows the original size, `X` fits the window.
- `H`/`L` rotate, `B` toggles filtering, `F` fullscreen.
- `I` opens the info overlay (EXIF, dimensions, path).
- `E` runs the command in `CmdA` on the current file.

## Archives and video

zip, rar, tar, 7z, cbr and cbz open like directories. Video is played with
libmpv and thumbnailed with frame extraction; build with `-tags nompv` to remove
both.

## Android

The APK build is in `docs/ANDROID.md`. On first run the empty state offers the
system folder picker; the gallery, image view and info overlay are the same.

## Config

`~/.config/imgview/config.toml` holds the image and gallery hotkeys, `CmdA`,
`SaveDir` and the general settings (tile width, gap, workers, images per page,
default window size). The bundled defaults are in `gallery/config.toml`. The
same file is read by tie-view, which adds its own keys.
