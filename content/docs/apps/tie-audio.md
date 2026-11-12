---
title: tie-audio
type: docs
weight: 11
---

A tag-aware music player for a tie library.

## Start and connect

```bash
tie-audio
tie-audio -config ~/.config/tie/config.toml
```

The first run asks for a tie config and a collection; the Settings tab has the
same connection editor as tie-view. `-host` selects a filehost by name.

## Views

- **Home** — shelves: Jump back in, Recently played, Recently added, Because you
  played, Haven't heard in a while, On this day, Favorites, Never played, Your
  playlists.
- **Albums** — Covers, Table or Now Playing; filter by tags, search, sort.
- **Tags** — the tag sidebar, including favourites.
- **Files** — the virtual tree.
- **Playlist** — the queue, with shuffle, repeat, save, clear, columns and drag
  reorder.
- **Settings** — connection, layout, playback backend, startup page, shortcuts.

The compact layout (phones, narrow windows) uses a bottom nav bar, a mini
transport and a full-screen Now Playing view.

## Playback

Two backends:

- **local** — in-process, PipeWire on Linux, OpenSL ES on Android; gapless.
- **pwplay** — a remote `pwplay-server` over HTTP.

Formats: FLAC, MP3, WAV, OGG and Opus. CUE sheets and zipped albums are albums.
MPRIS 2 exposes the player to media keys and `playerctl`; Android uses the media
session and lock-screen controls.

`tie-audio -control play-pause|play|pause|stop|next|previous` drives a running
instance and exits, which is how the desktop hotkeys and scripts talk to it.

## History

Plays are recorded into tie as `play-event` subjects with `play-count` and
`last-played`; the home shelves are queries over them. The recorder writes a
spool file and a cursor under the app's state directory, then flushes to the
store. It can be disabled in Settings.

## Shortcuts

| Key | Action |
|---|---|
| `Space` | Play / pause |
| `X` | Stop |
| `N` / `P` | Next / previous |
| `→` / `←` | Seek ±10 s |
| `=` / `-` | Volume ±10 % |
| `Esc` | Back |

All of them are rebindable under `[Hotkeys]`.

## Building

PipeWire development packages are required for the local backend. The Android
build follows `docs/ANDROID.md`; the libmpv cross-compile is in
`build-libmpv-android.sh`.
