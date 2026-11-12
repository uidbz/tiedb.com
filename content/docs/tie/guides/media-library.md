---
title: Build a media library
type: docs
weight: 11
---

tie is a good personal media library because it does not need a media server.
The files stay where they are (or in the blob store), the tags describe them, and
the apps read the tags.

## Import

Import a directory tree as a media type:

```bash
# photos: one directory, tags on the whole tree
tie import image-dir ~/Pictures/2014 \
  --tags travel --tags 2014 --dest /photos/2014

# music: album folders
tie import audio-dir ~/Music/"Slow Cooling" \
  --tags ambient --dest "/music/The Annealers/Slow Cooling"

# a whole library, one album per directory
tie import audio-dir --albums --by dir \
  --dest "/music/{albumartist}/{year} - {album}" \
  -y ~/Music
```

`--albums --by auto` groups by directory and then merges or splits by tags;
`--dry-run` shows the destinations first. The templates understand
`{albumartist}`, `{album}`, `{year}`, `{disc}` and the other media tags.

The built-in types are `audio-dir`, `image-dir`, `video-dir`, `document-dir`,
plus `audio-archive`, `image-archive`, `video-archive`, `document-archive`.
Custom types come from `[ImportDest]` in the config, for example:

```toml
[ImportDest]
podcast-dir = "/podcasts/{artist}/{title}"
```

## Tag

Tags are triples with a content hash on the left. The CLI has a full tag
tool:

```bash
tie tag list
tie tag add travel 9a1f…c03e
tie tag files travel
tie tag untagged
tie tag rename trip travel
tie tag set 9a1f…c03e travel 2014 reykjavik
tie tag show 9a1f…c03e
```

Tags attach to the hash, so they apply everywhere the content appears. The
**registry** (`tags all <tag>`) is what makes the tag list and the `query/`
directory complete.

## Collections

For an album, series or picture set, use the media contract:

- `member-of` from each item to the collection
- `added-at` on each membership
- `member-count` and an album summary on the collection
- a `cover` relation (or a `cover.jpg` / embedded picture)

`tie media index` creates the model for an existing tree, `tie media backfill`
fills gaps, and `tie media covers` shows the resolved cover and its source.
`tie media restat` refreshes counts.

## Playlists and history

```bash
tie playlist create "Sunday morning" --tracks 9a1f…,4d2a…,c03e…
tie playlist list
tie playlist show "Sunday morning"
tie history --since 30d
```

A playlist is a `tie-type playlist` subject with ordered `playlist-track`
entries. A play is an idempotent `play-event` with `play-count` and
`last-played`. tie-audio writes those events, and its shelves are queries over
them. `tie media schema` prints the relations the model uses.

## Look at it

The apps are the front end:

- **tie-fm** for import, cleanup, album matching and tag editing.
- **tie-view** for images: include/exclude tags, ratings, quick tags.
- **tie-audio** for the shelves, the cover wall, playlists and history.

The full source guide has the end-to-end walkthrough with more commands:
[`docs/media-database.md`](https://github.com/uidbz/tie/blob/master/docs/media-database.md).
