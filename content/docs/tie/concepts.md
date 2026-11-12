---
title: Concepts
type: docs
weight: 2
---

tie has a small vocabulary. Once these six ideas click, the CLI and the apps stop
feeling like a database with a filesystem bolted on.

## Triples

Everything is a `(key, value1, value2)` row, read as subject, relation, object.

```
pizza        topping   cheese
pizza        topping   basil
spaghetti    topping   tomato
```

The schema is open and multi-valued: adding a relation needs no migration, and a
key can have many values. This is the 2023 `tiedb` recipe example, still the
clearest one.

```
tie add pizza topping cheese
tie add pizza topping basil
tie get pizza
```

### Triples 101, with pasta

The original tiedb page used two recipes:

```sh
tie add "$pasta" title "$pasta"
tie add "$pasta" tag pasta
tie add "$pasta" tag tomato
tie add "$pasta" tag basil

tie get basil chili     # rows associated with both keys
tie get zucchini +url +title
```

The "+relation" form expands an associated key. That is the whole query
language: intersect keys, exclude with `--not` or `!value`, and expand.

## Content addressing

A file's bytes are stored under their 64-character HighwayHash in a sharded
layout:

```
data/9a/1f/9a1f3c0b7e…4d2ac03e
```

The hash is the file's identity. Upload the same bytes twice and they exist once.
A re-import of a changed file is a new blob and a new version of the file, with
the old version still addressable. Deduplication is a property of the model, not
a cleanup job.

The filehost treats its input as untrusted: downloads verify the hash, manifest
names and hashes are validated, and recursion depth is capped.

## tiedir

A directory is an immutable manifest: a blob containing `(hash, filename, size,
head)` entries. Entries are basenames only, so a directory is a tree of content
hashes — the same idea as a git tree. Renaming a directory is a new manifest, not
a rewrite of every child.

`tie stat` shows a directory's hash, children and total size. `tie versions`
lists the manifests a file has belonged to and can restore one.

## Tags and dir-types

A tag is a triple from a content hash to a tag:

```
9a1f…c03e   tag   travel
tags        all   travel
```

Tags attach to the hash, so they survive renames and re-imports. A **dir-type**
tells the library what a directory is: `audio-dir`, `image-dir`, `video-dir`,
`document-dir`, or one of the custom types in your config. The dir-type is how
the apps know to treat a directory as an album, a picture set or a document
folder.

## Collections

The word has two meanings, and both are useful:

- **A storage collection** is a namespace plus a name: one `.tie` file, selected
  with `-c` or `-C`. `Main` is the default; a collection's `_prev` file holds
  its version history.
- **A media collection** is a group of files that belong together: an album, a
  series, a picture set, a playlist. It follows a contract: `member-of`,
  `added-at`, `member-count`, an album summary and a `cover`.

The first is where the triples live. The second is what the apps show you.

## The virtual tree

DirUID nodes carry a `path` (`tie:/…`), a `parent` and `tie-type directory`.
Files hang off them by `parent` edges. This tree is what `tie mount` presents as
`files/`, and what `tie verify` checks for orphans, cycles and duplicate paths.
The mount is a view; the triples are the truth.
