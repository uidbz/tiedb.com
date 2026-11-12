---
title: Quickstart
type: docs
weight: 1
---

All five apps are in one repository, [uidbz/tie-gui](https://github.com/uidbz/tie-gui).
They build with Go 1.27+ and CGO, plus the native libraries for video, audio and
devices.

## Build

```bash
git clone --recurse-submodules https://github.com/uidbz/tie-gui
cd tie-gui
go build ./cmd/imgview
go build ./cmd/tie-view
go build ./cmd/tie-fm
go build ./cmd/tie-audio
go build ./cmd/tie-browser
```

`make install` builds and installs the four desktop apps (imgview, tie-view,
tie-fm, tie-audio) into `~/.local/bin`, with a `.desktop` entry and a hicolor
icon. tie-browser is built separately for now.

## Native dependencies

| Library | Needed by | Disable with |
|---|---|---|
| Fyne (X11/Wayland, OpenGL) | every app | — |
| libmpv | imgview, tie-view, tie-fm, tie-audio | `-tags nompv` |
| PipeWire | tie-audio local playback | use the pwplay backend |
| libmtp | tie-fm MTP devices | not buildable out |
| udisks2, gvfs | tie-fm device browsing | Linux only |

`docs/ANDROID.md` covers the Android toolchain, the libmpv cross-compile and the
APK build scripts.

## First run

- **imgview** needs nothing: `imgview ~/Pictures/2014`, or pass an archive.
- **tie-view**, **tie-audio** and **tie-fm** need a tie config. Copy
  `~/.config/tie/config.toml` from your library install, or use the test
  environment. Each app has a connection editor in Settings.
- **tie-browser** needs a collection with scenes; it creates a starter
  `tie:/scenes/home` if none exists.

The shared tie config has per-app collection profiles (`images`, `audio`,
`files`). The apps fall back to `DefaultCollection` if a profile is missing.

## The test environment

The tie repo's `test-env/` starts a local triplestore and filehost on 2161/2162.
tie-audio has an extra layer that starts a pwplay server on 8080:

```bash
cd tie-gui/cmd/tie-audio/test-env
./start.sh
```

Use throwaway credentials (`defaultuser` / `defaultpassword`) and no public
exposure.
