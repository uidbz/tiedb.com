---
title: Changelog
type: docs
weight: 40
---

tie is pre-1.0. The full, canonical changelog is
[`CHANGELOG.md`](https://github.com/uidbz/tie/blob/master/CHANGELOG.md) in the
source repo. The highlights below are the ones that changed how the library is
used.

## v0.9.0 — cover artwork resolver

- `tie media covers`: a deterministic resolver for album covers, with the source
  of each decision.
- Cover artwork for albums, series and picture sets, including archive and
  embedded art.
- The media schema command and the matching docs.

## v0.8.x — search, direct mode, the compute plane

- Free-text search with words, substring, prefix and fuzzy modes.
- `tie -f`: the whole client against a local `.tie` file, including mounts and
  SMB.
- `flock` protection against double writers and `--read-only` snapshots.
- `tie-runner`, `tie-run`, `tie-exec` and the fleet admin client.
- A performance pass and the memory optimization note.

## v0.7.x — SMB, and a code review

- `tie-smb`: SMB2/3 from pure Go, with NTLMv2 users and roles.
- The October 2026 review: 81 findings, one fix per commit, each with a
  regression test that failed before the fix. The
  [story is here](/blog/81-findings/).
- Breaking changes around the media contract and the config keys.

## v0.6.0 — the media collection contract

- `member-of`, `added-at`, `member-count`, covers and playlists.
- The album model the apps build on.

## Older

The v0.4/v0.5 line added the tag tree, the FUSE mount and `verify`. The project
began in 2023 as **tiedb**, a key-value-value store; the 2023 page is preserved in
[the concepts page](/docs/tie/concepts/#triples-101).
