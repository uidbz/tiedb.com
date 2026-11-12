---
title: Verify and repair
type: docs
weight: 15
---

`tie verify` is fsck for a collection. It cross-checks the forward and reverse
indexes, then scans the virtual tree for problems:

```bash
tie verify             # findings as TSV on stdout
tie verify --deep      # also walk every file
tie verify --index     # index checks only
tie verify --check-blobs
```

Exit codes are part of the contract: `0` clean, `1` findings, `2` usage,
`4` operational. That makes it usable in a cron job or a CI check.

## The repair pipeline

Repair is a plan you can read, then an apply you can resume:

```bash
tie verify | tie repair plan > plan.tsv
less plan.tsv
tie repair apply < plan.tsv
tie repair relink
```

`plan` turns findings into operations. `apply` journals each operation, so an
interrupted repair is safe to re-run. `relink` rebuilds containment from the
directory manifests when the tree has lost its `parent` edges.

## What it checks

- Forward and reverse index agreement.
- Orphans (a triple whose subject has no directory or file record).
- Dangling parents, cycles and duplicate paths.
- Missing metadata (filename, size, type) for a file.
- Blob presence and hash, with `--check-blobs`.

The source guide:
[`docs/verify.md`](https://github.com/uidbz/tie/blob/master/docs/verify.md).
