---
title: CLI reference
type: docs
weight: 30
---

Generated from `tie --help` at **v0.9.0**. The binary is the source of truth; this
page is a convenience.

```
tie [global options] [command [command options]]
```

## Global options

| Flag | What it does |
|---|---|
| `-C, --config NAME` | Config file. Precedence: `-C`, `$TIE_CONFIG`, `./config.toml`, user config, system config. |
| `-c, --collection NAME` | Collection from `[Collections]`; empty uses `DefaultCollection`. |
| `-f, --file PATH` | Open a `.tie` file directly, with no triplestore server. |
| `--read-only` | With `--file`: lockless snapshot, writes refused. |
| `--reverse-relations LIST` | With `--file`: the relations to index in reverse. |
| `-q, --quiet` / `-V, --verbose` | Less or more diagnostics. |
| `--format tsv\|json\|table` | Listing output. Defaults to table on a terminal, TSV otherwise. |
| `-v, --version` | Client and server versions. |

Stdout is data. Diagnostics go to stderr. Exit codes: `0` clean, `1` findings,
`2` usage, `4` operational.

## Commands

| Command | Aliases | What it does |
|---|---|---|
| `add` | `a` | Add a triple. |
| `del` | `d` | Delete a triple, a relation, or every triple for a key. |
| `get` | `g` | Query triples; multiple keys intersect, `-` reads keys from stdin. |
| `search` | | Free-text search over named relations. |
| `stat` | | Info about a path or hash. |
| `upload` / `download` | | Move files to and from a filehost. |
| `import` | `i` | Upload a tree and tag it, by dir-type. |
| `tag` | | `list`, `add`, `del`, `rename`, `files`, `untagged`, `show`, `set`. |
| `versions` | | `list`, `restore` file version history. |
| `mount` | | A dir hash, or `--db` for the live tree. |
| `dump` / `restore` | | TSV export and additive, idempotent import. |
| `verify` | | fsck: index and tree findings as TSV. |
| `repair` | | `plan`, `apply`, `relink`. |
| `media` | | `covers`, `index`, `backfill`, `rescan`, `restat`, `schema`. |
| `history` | | Recorded plays, newest first. |
| `playlist` | | `list`, `show`, `create`, `delete`, `migrate`. |
| `conf` | | `show`, `create`. |
| `completion` | | bash, zsh, fish, pwsh. |

## get

```bash
tie get pizza                     # everything about pizza
tie get pizza --not cheese        # exclude a value
tie get pizza +topping            # expand an associated key
tie get -r tag travel             # reverse: find subjects for a relation
tie get -f 'title=*Wave*'         # filter
tie get -l 50 -o 100 -s title     # paging and sort
tie get - < keys.txt              # keys from stdin
tie get pizza --table             # the table format
```

## import

```bash
tie import image-dir DIR --tags travel --tags 2014 --dest /photos/2014
tie import audio-dir DIR --albums --by dir --dry-run -y ~/Music
```

Flags: `-t/--tags` (repeatable), `--dest` (virtual path or template),
`--albums`, `--by auto|dir|tags`, `--dry-run`, `-y`, `--host`, `--collection`.

The subcommand is the dir-type: `audio-dir`, `image-dir`, `video-dir`,
`document-dir`, their `*-archive` variants, or a custom type from `[ImportDest]`.

## verify and repair

```bash
tie verify [--deep] [--index] [--check-blobs]
tie verify | tie repair plan > plan.tsv
tie repair apply < plan.tsv
tie repair relink
```

`verify` returns `1` when it finds something, so a cron job can tell the
difference between clean and findings.

## mount

```bash
tie mount HASH ~/snapshot
tie mount --db ~/mnt [--cache 8] [--verify]
```

## media and playlists

```bash
tie media index
tie media covers
tie media backfill
tie playlist create "Sunday morning" --tracks HASH,HASH
tie history --since 30d
```

## Tag

```bash
tie tag list
tie tag add travel HASH
tie tag files travel
tie tag untagged
tie tag rename trip travel
tie tag set HASH travel 2014 reykjavik
```

The full command list and flags come from the binary itself; run
`tie <command> --help`. The source README has longer examples.
