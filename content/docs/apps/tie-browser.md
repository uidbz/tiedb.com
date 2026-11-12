---
title: tie-browser
type: docs
weight: 14
---

A scene browser and visual designer. **Preview:** the scene format can still
change.

## A scene is a directory

```
tie:/scenes/house/living-room/
  scene.json      the page
  page.risor      optional behaviour
  thumbnail.png   optional preview
```

`scene.json` has a `format` (`fynerisor.scene/1`), a `title`, an optional
`script`, and a `root` object with `place` (x, y, w, h), `props` and `children`.

## Run it

```bash
tie-browser
tie-browser -config ~/.config/tie/config.toml -collection Main
tie-browser -home tie:/scenes/home
```

If the home scene is missing, tie-browser creates a starter one. `-pwplay` sets
the audio server URL (`http://localhost:8080` by default).

## Components

| Type | Props |
|---|---|
| `tie.image` | `hash` or `path`, fit/fill |
| `tie.gallery` | `tags` (a live query), layout |
| `tie.collection` | album/series id |
| `tie.tags` | the tag picker |
| `tie.related` | co-tags of the current subject |
| `tie.document` | a PDF/document hash |
| `tie.audio` | an audio hash, transport |
| `tie.video` | a video hash, poster |

Containers (`vbox`, `hbox`, `grid`, `free`) and hotspots (`rect`, `circle`) are
built in. Actions are `navigate`, `open`, `play` and a small set of scene
operations.

## Design

`Ctrl+E` (or `F4`) opens the designer. Move and resize boxes, edit props, pick
media from the library, and save. The save indicator is persistent: Unsaved,
Saving…, Saved, Save failed. `Ctrl+S` saves a version; conflicts are handled by
the store.

## Thumbnails and indexes

`SaveThumbnail` renders the scene offscreen at 640×400 through the software
canvas. The store writes `scene-doc`, `member-of` and `links-to` triples, which is
what makes "which scenes use this photo?" a query.

## Roadmap

Native mpv video surfaces and richer tag/chip pickers. The app is the newest in
the family; expect the format to move.
