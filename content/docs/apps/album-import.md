---
title: Album import
type: docs
weight: 20
---

tie-fm's **Import as albums…** and the CLI's `tie import audio-dir --albums`
share the planner. The difference is the UI: tie-fm shows the form, the
destinations and the MusicBrainz matches before you commit.

## The planner

```bash
tie import audio-dir --albums --by dir \
  --dest "/music/{albumartist}/{year} - {album}/cd{disc}" \
  --dry-run -y ~/Music
```

- `--by dir` — one directory is one album.
- `--by tags` — cluster by album tags.
- `--by auto` — directories first, then merge or split by tags.

Templates: `{albumartist}`, `{artist}`, `{album}`, `{year}`, `{genre}`,
`{disc}`, `{title}` and the rest of the media tags. A missing value becomes an
empty component; the dry run shows the result.

## MusicBrainz

tie-fm can query a local Postgres mirror (fast, offline) or the web service
(rate-limited). Matching fills the album artist, year, genre, track numbers and
disc numbers, and can download the cover. The match is a suggestion, not an
automatic decision; the form shows score and release.

`Fill missing covers…` walks the albums that have no cover and writes the
resolved art back. `Use as album cover` sets the cover from any image in the
pane.

## Discs and multi-disc sets

`cd{disc}` in the destination template creates `cd1/`, `cd2/` (or `disc1/`) and
sets the disc numbers. The album model groups them under one collection.

## Re-imports

Content addressing makes a re-import a version, not a duplicate: unchanged
tracks keep their hashes and tags, changed tracks become new hashes, and the
collection manifest gets a new version. The old version remains restorable with
`tie versions`.

## The source

[`docs/TIE-FM-ALBUM-IMPORT.md`](https://github.com/uidbz/tie-gui/blob/master/docs/TIE-FM-ALBUM-IMPORT.md)
and
[`docs/TIE-FM-MUSICBRAINZ.md`](https://github.com/uidbz/tie-gui/blob/master/docs/TIE-FM-MUSICBRAINZ.md).
