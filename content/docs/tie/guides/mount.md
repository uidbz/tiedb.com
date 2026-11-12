---
title: Mounting
type: docs
weight: 10
---

tie can mount two different things, both over FUSE:

```bash
tie mount <dir-hash> ~/snapshot   # one immutable tiedir, read-only
tie mount --db ~/mnt              # the live tag-derived filesystem
```

The live tree has three top-level directories:

```
~/mnt/
  query/     every path is a tag query
  files/     the writable virtual tree
  tags/      tags as text files
```

## query/

A path under `query/` is a tag expression:

```bash
ls "mnt/query/jazz"                 # all jazz
ls "mnt/query/jazz mellow"          # AND
ls "mnt/query/jazz mellow -live"    # exclude
ls "mnt/query/tags"                 # the whole tag vocabulary
```

The space-separated components are intersected. A leading `type:` scopes by
dir-type (`type:audio-dir`). Saved queries from the config's `[Queries]` table
appear as named paths.

## files/

The writable virtual tree. `mv`, `mkdir`, copy, create and edit are reflected in
the triples. Re-importing a changed file creates a new version; `tie versions`
lists them and `tie versions restore` puts one back.

## tags/

Each file has a text file of its tags: one tag per line, editable in place.
Directories expose `.tags` and `.type`, which is how a dir-type is set without
the CLI.

## Caching

Both mounts share a single-flight, LRU blob cache:

```bash
tie mount --db ~/mnt --cache 8        # 8 GB cache
tie mount --db ~/mnt --cache 8 --verify
```

The cache streams to disk, so memory stays bounded regardless of file size.
The FUSE and SMB views of the same store stay in sync because both re-read the
store on every operation.

## Requirements

- Linux with `/dev/fuse` and `fusermount` (or `fusermount3`).
- A writable cache directory; defaults are under the user cache.
- The triplestore and filehost must be reachable, or use `tie -f` for the
  serverless variant.

```bash
tie mount --db ~/mnt &
ls ~/mnt/files/music
fusermount -u ~/mnt
```

See [tie-smb](/docs/tie/guides/smb/) for sharing the same tree with Windows and macOS.
