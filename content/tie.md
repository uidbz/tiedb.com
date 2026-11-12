---
title: tie
type: showcase
mark: tie
hue: "#2dd4bf"
eyebrow: "The library"
tagline: "Tag it. Find it. Mount it."
summary: >-
  A triple store and a content-addressed blob store, joined by the content hash.
  Files stay files, meaning lives in triples, and the whole library is one thing
  you can back up, verify and mount as a folder.
badges:
  - "v0.9.0"
  - "BSD-3"
  - "Linux-first"
  - "Go & Python clients"
primary:
  label: "Install tie"
  url: "/install/"
secondary:
  label: "Take the 60-second tour"
  url: "/docs/tie/getting-started/"
docs:
  - title: "Getting started"
    url: "/docs/tie/getting-started/"
    desc: "The test environment, one import, one mount, and a tag query."
    icon: "play"
  - title: "Concepts"
    url: "/docs/tie/concepts/"
    desc: "Triples, content addressing, tiedirs, dir-types and collections."
    icon: "layers"
  - title: "CLI reference"
    url: "/docs/tie/cli/"
    desc: "Every command and flag, generated from the binary."
    icon: "terminal"
---

## The model

tie splits what a file **is** from what it **means**. The bytes go into a
content-addressed blob store under their HighwayHash; the meaning goes into
triples — `tag`, `member-of`, `rating`, `path`, `play-count`. The hash is the
join, which gives you some unusual properties for free:

- **Deduplication is structural.** Import the same photo from two folders and
  the bytes exist once. A re-import becomes a *version* of the file, not a
  second copy.
- **Tags follow content.** Rename the file, move it, re-import it: the tags
  still point at the bytes. The filesystem is a view, not the source of truth.
- **The tree is versioned.** A directory is a manifest (`tiedir`) with entries
  and hashes. Every import that changes it produces a new manifest, so you can
  list and restore versions.

<div class="t-term" style="margin: 2rem 0">
  <div class="t-term-bar">
    <img src="/brand/tie.svg" alt="" />
    <span>one hash, both ways</span>
    <span class="t-spacer" style="flex:1"></span>
    <span>tie v0.9.0</span>
  </div>
  <pre class="t-term-body"><code><span class="t-p">$</span> <span class="t-cmd">tie upload ~/Pictures/reykjavik.jpg</span>
<span class="t-g">9a1f3c0b7e…4d2ac03e</span>
<span class="t-p">$</span> <span class="t-cmd">tie add 9a1f3c0b7e…4d2ac03e tag travel</span>
<span class="t-p">$</span> <span class="t-cmd">tie add 9a1f3c0b7e…4d2ac03e rating 4</span>
<span class="t-p">$</span> <span class="t-cmd">tie add 9a1f3c0b7e…4d2ac03e member-of iceland-2014</span>
<span class="t-p">$</span> <span class="t-cmd">tie get -r tag travel</span>
<span class="t-o">9a1f3c…c03e  tag  travel</span></code></pre>
</div>

## A path is a query

Mount the live library and the tag vocabulary becomes a directory tree:

```
~/mnt/
  query/          every path is a tag query
    jazz mellow -live/
    travel 2014 -screenshots/
  files/          the writable virtual tree
    music/
    photos/2014/
  tags/           one text file per file's tags; .tags and .type per directory
```

```
$ ls "~/mnt/query/jazz mellow -live"
baker-street/   kind-of-blue/   slow-cooling/
```

Both views are live. Write through `files/`, and the tags and queries change on
the next read, because both adapters re-read the store on every operation. The
same tree is served over SMB to Windows Explorer and macOS Finder by a pure-Go
SMB2/3 server — no Samba, no exports table, no ACL file.

## The store is a file

A collection is a single `.tie` file: a log of fixed 50-byte records with
interned strings, forward and reverse association indexes, and tombstoned slots
that get reused. One writer goroutine fsyncs on a 10-second ticker and drains on
SIGTERM.

- **No server required.** `tie -f library.tie mount ~/mnt` opens the file
  in-process and mounts it. `tie -f` also serves SMB and runs tag queries.
- **No double writers.** The file is `flock`ed. `--read-only` opens a lockless
  snapshot beside a live writer.
- **Backups are boring.** `rsync` the directory, snapshot the Btrfs subvolume, or
  `tie dump` to TSV and `tie restore` into a fresh collection. The dump is
  additive and idempotent; the blob store is content-addressed, so a backup is
  just a copy of the blobs you care about.

> [!TIP]
> The real benchmark library is 7.4 GiB on disk with 60.7 million entries and
> 98.2 million associations. It loads in 2m46s with a 19.4 GiB heap on 8 workers.
> The number is in the [internals note](/docs/tie/internals/) with the design
> decisions and the two experiments that were reverted.

## Built to be checked

The filehost is treated as untrusted: every download verifies the hash, every
manifest filename and hash is validated, and recursion depth is capped.
`tie verify` cross-checks the indexes, then scans the virtual tree for orphans,
dangling parents, cycles, duplicate paths and missing metadata.

```
$ tie verify --deep | tie repair plan > plan.tsv
$ tie repair apply < plan.tsv
```

The plan is a TSV you can read, edit and re-run. The apply step journals every
operation, so an interrupted repair is resumable. `tie repair relink` rebuilds
containment from the directory manifests when the tree has lost its edges.

## Search without a sidecar

Free-text search is an in-memory corpus over named relations, rebuilt on demand
and cached. Accents and case are folded; queries can be words, substrings,
prefixes or fuzzy.

| Workload | Time |
|---|---|
| Cold build + query, 10k subjects | ~18 ms |
| Cold build + query, 100k subjects | ~223 ms |
| Warm query, 10k subjects | ~1.2 ms |
| Warm query, 100k subjects (4 fields each) | ~15 ms |
| Warm fuzzy query, 50k subjects | ~18 ms |

Those numbers are from [the search note](/docs/tie/guides/search/), measured on
a Ryzen 9 9950X with a single run. There is no SQLite FTS5, no Bleve, no sidecar
index — the note explains why, and what would change its mind.

## Media-aware

The library knows about albums, series and picture sets, not just files:

- **Collections** follow a contract: `member-of`, `added-at`, `member-count`,
  album summary, `cover`. A collection can be a directory, an archive, or a
  cluster of tracks.
- **Covers** are resolved by a resolver with a deterministic preference order;
  `tie media covers` shows what it picked and why.
- **Playlists** are ordered `playlist-track` entries.
- **Play history** is an idempotent `play-event` subject with `play-count` and
  `last-played`. That is what tie-audio's shelves are made of.

## The compute plane

`tie-runner` turns content-addressed trees into a compute substrate: a job takes
an input hash, runs in a container, and produces an output hash. Pipelines,
services and fleets are in the same store. It is optional; tie works as a
library without it.

## Developers

The Go client is the reference implementation; the Python client is byte-
compatible and dependency-free. The CLI is built on the same client, so anything
the CLI does is a library call.

```go
tc, _ := client.NewTieClientFor(cfg, "Main")
hash, _ := tc.ImportFile("/photos/reykjavik.jpg", "default", "image-file",
    []string{"travel", "2014"}, "/photos/2014", 0)
rows, _ := tc.Get("tag", "travel")
```

- **Go**: `github.com/uidbz/tie/client` — context variants for every call,
  `ErrNotFound` as the sentinel, search profiles for audio and files.
- **Python**: `tie-client` 0.1.0, stdlib only, Python 3.11+.
- **CLI**: exits 0 clean, 1 findings, 2 usage, 4 operational; stdout is data.

## Limits, honestly

> [!WARNING]
> tie is built for a personal library on one PC or a small trusted LAN. It is
> not a multi-tenant service. Passwords live in the config file, the servers
> speak HTTP by default, and SMB is v2/3 without encryption. Run it behind
> WireGuard or on a trusted network, and read the [access-control notes](/docs/tie/operations/access/) before exposing a port.

It is also pre-1.0. The stored and wire formats have been stable for a while,
but the CLI and the media contract can still change. The
[changelog](/docs/tie/changelog/) is the honest record.

## From tiedb, with love

This domain used to hold a 2023 page for **tiedb**, the key-value-value store
that became tie. The old recipe example — pasta, basil, chili — is still the
clearest way to explain a triple, so it lives on in
[Triples 101](/docs/tie/concepts/#triples-101). The name moved on; the idea
didn't.
