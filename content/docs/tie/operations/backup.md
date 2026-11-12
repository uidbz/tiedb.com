---
title: Backup and storage
type: docs
weight: 22
---

A tie deployment has two data sets:

- **the triplestore** — one `.tie` file per collection (plus `_prev` history);
- **the filehost** — the content-addressed blob directory.

They can be backed up independently, but restore them together.

## The simplest backup

```bash
# stop writes, or use a snapshot
tie dump --file triples.tsv
rsync -a ~/.local/share/tie/ backup/tie-db/
rsync -a /var/lib/tie/data/ backup/tie-data/
```

`tie dump` is a TSV of every triple; `tie restore` is additive and idempotent.
The blob store is content-addressed, so a plain copy is valid: a blob is either
present or not, and its name is its hash.

## Filesystem snapshots

On Btrfs, a read-only snapshot of the triplestore directory and the blob
directory is consistent enough for `btrfs send`:

```bash
btrfs subvolume snapshot -r /var/lib/tie /var/lib/tie-snap
btrfs send /var/lib/tie-snap | ssh backup btrfs receive /backup
```

`contrib/backup/` has `btrfs-send-backup.sh`, `rsync-backup.sh` and
`triplestore-db-backup.sh`.

## The reaper

The filehost can reap blobs that are no longer referenced. `ReapInterval = "0"`
disables it (the default in the test environment). If you enable it, make sure
the triplestore backup you keep still references what the blobs need.

## Restore order

1. Stop the services.
2. Restore the blob directory.
3. Restore the `.tie` file(s).
4. Start the triplestore, then the filehost.
5. `tie verify --deep` and `tie verify --check-blobs`.

The source note is
[`docs/filehost-storage-and-backup.md`](https://github.com/uidbz/tie/blob/master/docs/filehost-storage-and-backup.md).
