---
title: tie-fm
type: docs
weight: 10
---

A twin-panel file manager for local disks and tie collections.

## Panes and tabs

Each pane has its own tabs, history, breadcrumb (editable in place) and
bookmarks. `Ctrl+Tab` switches panes, `Ctrl+T` opens a tab, `Ctrl+W` closes it,
and Ctrl+click on a folder opens it in a new tab. `Ctrl+I` filters the current
page.

Bookmarks remember both a path and a collection, so a bookmark can point inside
a tie library.

## Copying into tie

The **Copy/Move into tie as ▸** menu offers the eight built-in dir-types plus
custom ones. Re-importing a changed directory creates a version; the old tree
remains addressable.

The operations panel supports pause, resume and stop. Conflicts offer skip,
rename, overwrite and apply-to-all, in a Dolphin-style dialog.

## Album import

**Import as albums…** discovers albums from a music root and places each one with
a template:

```
/{albumartist}/{year} - {album}/cd{disc}
```

The dry run lists every discovered album and destination. MusicBrainz matching
can use a local Postgres mirror (`TIE_FM_MB_DSN`) or the web service; **Fill
missing covers…** writes the resolved art back into the album directory or the
collection, depending on the mode. **Use as album cover** is available on any
image in the pane.

## Tags

With no selection, the tag panel filters the pane. With rows selected, it edits
their tags. Favourite tags are shared with tie-view and tie-audio. Directories
are taggable; the dir-type dialog sets the type without the CLI.

## Preview and media

Preview mode turns either pane into a thumbnail grid with an in-image viewer and
a video player. It uses the same `gallery/` engine as imgview. Archive extraction
is on the context menu.

On Linux, USB drives are discovered through udisks2 and MTP phones through
libmtp, with a gvfs fallback. Other platforms build without the device manager.

## File associations

Per-extension associations can open a file with an external app, stream it from
tie, or hand a `tie:` URL to tie-view or tie-audio. The context menu's **Open
with…** is the same table.

## Delete

See [What Delete means](/docs/apps/delete/). The short version: local rows delete
the file; tie rows remove the reference and keep the content-addressed blob.
