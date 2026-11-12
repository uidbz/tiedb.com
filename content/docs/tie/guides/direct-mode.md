---
title: Direct mode
type: docs
weight: 13
---

`tie -f <file.tie>` runs the whole client against a local `.tie` file, with no
triplestore daemon. The filehost is still used for blobs (or point it at a local
one).

```bash
tie -f ~/.local/share/tie/Main.tie get -r tag travel
tie -f ~/.local/share/tie/Main.tie mount --db ~/mnt
tie -f ~/.local/share/tie/Main.tie import image-dir ~/Pictures/2014 --tags travel
```

Direct mode is the same code path as the HTTP client: a local transport
dispatches to the same API handlers. That means the CLI, the mounts, SMB and the
APIs behave the same with or without a server.

## One writer

The file is `flock`ed. A second writer fails fast instead of corrupting the
store. For a read-only view beside a live writer:

```bash
tie -f ~/.local/share/tie/Main.tie --read-only get -r tag travel
```

`--read-only` takes a lockless snapshot and refuses writes.

## Reverse relations

The reverse index is opt-in per relation. The server defaults to
`tag,path,parent,tie-type,version-of,member-of`. With `-f` you can choose:

```bash
tie -f Main.tie --reverse-relations tag,path,parent,tie-type get -r tag travel
```

Indexing fewer relations saves memory; indexing more makes reverse queries
faster. For a large database, match the server's subset.

## Mount and SMB

Both mounts and `tie-smb` can run in direct mode:

```bash
tie -f Main.tie mount --db ~/mnt
tie-smb -f Main.tie -config tie-smb.toml
```

That is enough for a laptop with no service manager, or a single-file library on
a USB disk.
