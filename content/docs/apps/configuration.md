---
title: Configuration
type: docs
weight: 2
---

The apps share one idea: a **tie config** describes the servers and the
collections, and each app has a small config of its own for layout, shortcuts and
theme.

## The tie config

`~/.config/tie/config.toml` (or a file passed with `-C`) is the connection
description. Each app chooses a collection profile from it:

| App | Default profile |
|---|---|
| tie-view | `images` |
| tie-audio | `audio` |
| tie-fm | `files` |
| tie-browser | `DefaultCollection` |

```toml
Username = "you"
Password = "change-me"
Namespace = "Collections"
Collection = "Main"
DefaultCollection = "Main"
TripleStoreURL = "http://localhost:1161"
DefaultFileHosts = ["default"]

[FileHosts.default]
URL = "http://localhost:1162"

[Collections.audio]
Collection = "Music"
# ...
```

The connection editors in tie-view and tie-audio write the same file, so a
change made in one app is visible to the others.

## App configs

| App | Path |
|---|---|
| imgview / tie-view | `~/.config/imgview/config.toml` |
| tie-fm | `~/.config/tie-fm/config.toml` |
| tie-audio | `~/.config/tie-audio/config.toml` |
| tie-browser | `~/.config/tie-browser/config.toml` |

Environment overrides that matter for packaging and tests:

- `TIE_THEME=light|dark` forces the theme variant.
- `FILESDIR` relocates the app config (used by the tests and the Android build).
- `TIE_CONFIG` is the CLI-equivalent config path.

## Theme

All five apps call `tiestyle.Apply` and share the same surfaces: `#17181b` /
`#ffffff` backgrounds, the teal accent (`#2dd4bf` dark, `#0f766e` light), the
same radii and the same warn/ok/error colours. The contrast floor is enforced by
`tiestyle`'s tests: foreground ≥ 7:1, muted ≥ 4.5:1, accent ≥ 3:1.

## Windows and desktop entries

`make install` writes `.desktop` files and a 256px icon per app. On Android, the
manifests and the media session bridge are in the app source; `docs/ANDROID.md`
has the build steps.

## Hotkeys

Every app has a config-driven hotkey table:

- imgview and tie-view: `[Image]` and `[Gallery]` sections in the imgview
  config; one key can run `CmdA` on the current file.
- tie-audio: `[Hotkeys]` with transport, seek and volume bindings.
- tie-fm: a small fixed set (`Ctrl+I`, `Ctrl+T`, `Ctrl+W`, `Ctrl+Tab`).
- tie-browser: `Ctrl+E`/`F4` for the designer, `Ctrl+S` to save.

The manual pages for each app list the defaults.
