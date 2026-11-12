---
title: tie-view
type: docs
weight: 12
---

A tag-driven image viewer for a tie collection, and a good local image viewer on
its own.

## Open a collection, a tag or a URL

```bash
tie-view
tie-view -config ~/.config/tie/config.toml
tie-view -tag favorite
tie-view "tie:9a1f3c0b7e…4d2ac03e"
tie-view ~/Pictures/2014
```

`-tag` sets the startup tag; a `tie:` URL selects one file; a local path opens
the folder like imgview.

## Tag sidebar

The Tags tab is a query builder:

- click a tag to include it, alt/right-click to exclude;
- **Related tags** refines by the co-tags of the current result;
- the star pins a favourite tag to the top;
- **Untagged** shows the files with no tags.

Sort by name, rating, newest or oldest; filter the rating (Any, Unrated, N, N+).

## Ratings and quick tags

- **1–5 stars** write `rating` on the content hash.
- **Quick tags** are pills over the image edge: one tap or one key per tag, plus
  stars and the favourite heart. Sizes xs–xl; configured in
  `quicktags.toml` or Settings → Quick tags. See
  [Quick tags](/docs/apps/quick-tags/).
- On Android, swipe up for the tagger and use the same pills with a larger touch
  target.

## Files tab

The virtual tree from tie, read-only, for when you need a path rather than a
tag. Clicking a row selects it in the gallery.

## Thumbnails and video

Thumbnails are requested from the filehost's thumbnail service and cached; the
first view of an image may generate and upload a thumbnail, then later views are
instant. Video plays through libmpv; `-tags nompv` removes it.

## Hand-off

tie-fm can open a row here with the collection and file already selected. The
mechanism is the `tie:` URL scheme, so a bookmark or a script can do the same.

## Settings

The Settings tab edits the connection (collection and profile), quick tags and
the startup page (favourites, latest, a tag, or blank). Startup preferences live
in the Fyne app preferences; connection settings live in the tie config.
