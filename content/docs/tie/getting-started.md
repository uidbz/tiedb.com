---
title: Getting started
type: docs
weight: 1
---

The fastest way to see tie is the local test environment in the source repo. It
builds the binaries, starts a triplestore and a filehost on **2161** and
**2162**, seeds a few files with tags and mounts the live tree.

```bash
git clone https://github.com/uidbz/tie
cd tie/test-env
./build.sh        # build tie, tie-triplestore and tie-filehost into ./bin
./start.sh        # start both services on localhost
./seed.sh         # upload and tag a few sample files
./mount-db.sh     # mount the live tag-derived filesystem at ./mnt
```

Open another terminal:

```bash
ls mnt/query/tags                  # finance  outdoors  vacation
ls mnt/query/vacation/             # beach.txt  holiday-pics/
ls "mnt/query/vacation outdoors"   # AND query
cat mnt/query/vacation/beach.txt
```

`./demo.sh` runs the whole thing non-interactively, which is a good way to see
what the commands actually do.

## A first library of your own

Create a config with `tie conf create`, or copy
`test-env/config.toml` and change the endpoints. The important fields:

```toml
Username = "defaultuser"
Password = "defaultpassword"
Namespace = "Collections"
Collection = "Main"
DefaultCollection = "Main"
TripleStoreURL = "http://localhost:2161"
DefaultFileHosts = ["default"]

[FileHosts.default]
URL = "http://localhost:2162"
```

Upload one file and tag it:

```bash
tie upload ~/Pictures/reykjavik.jpg
# 9a1f3c0b7e…4d2ac03e

tie add 9a1f3c0b7e…4d2ac03e tag travel
tie add 9a1f3c0b7e…4d2ac03e tag 2014
tie add 9a1f3c0b7e…4d2ac03e title "Reykjavík"
tie add tags all travel
tie add tags all 2014
```

Then query it:

```bash
tie get -r tag travel
tie get 9a1f3c0b7e…4d2ac03e
tie stat 9a1f3c0b7e…4d2ac03e
```

## Import a directory instead

`tie import` uploads a tree, creates the virtual directories and applies tags in
one pass. The subcommand is the dir-type:

```bash
tie import image-dir ~/Pictures/2014 --tags travel --tags 2014 --dest /photos/2014
tie import audio-dir ~/Music/"Slow Cooling" --tags ambient --dest "/music/The Annealers/Slow Cooling"
```

The built-in dir-types are `audio-dir`, `image-dir`, `video-dir`,
`document-dir` and their four `*-archive` variants. Custom types live in the
config's `[ImportDest]` table.

## Mount it

```bash
tie mount --db ~/mnt          # live tree: query/, files/, tags/
# or a single immutable directory by hash
tie mount 4d2ac03e… ~/snapshot
```

See [Mounting](/docs/tie/guides/mount/) for the tree layout, writable paths and the
blob cache.

## No server at all

```bash
tie -f ~/.local/share/tie/Main.tie import image-dir ~/Pictures/2014 --tags travel
tie -f ~/.local/share/tie/Main.tie mount ~/mnt
```

Direct mode opens the `.tie` file in-process; the filehost is still used for
blobs, but there is no triplestore daemon. See
[Direct mode](/docs/tie/guides/direct-mode/).

## Next

- [Concepts](/docs/tie/concepts/) — what a triple is, and why content addressing matters.
- [The media library guide](/docs/tie/guides/media-library/) — albums, tags, covers and playlists.
- [CLI reference](/docs/tie/cli/) — every command.
