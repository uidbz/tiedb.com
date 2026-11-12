---
title: "81 findings, 81 regression tests"
type: blog
date: 2026-10-20
description: "What a whole-repo review found in tie, and why every fix shipped with a test that failed first."
---

In October 2026, tie got a full code review: 81 findings across the storage
engine, the HTTP layer, the mounts, the importers and the tests. The rule for the
fixes was simple and slow:

1. one commit per issue;
2. a regression test that fails before the fix;
3. the fix, and the test passes;
4. no drive-by cleanups in the same commit.

The categories were the usual ones for a young storage project: error paths that
lost the error, lock ordering, partial writes on a full disk, path traversal in
archive extraction, a cache that could return a stale size, and a handful of
places where a context was not threaded through a retry.

None of it is glamorous. All of it is why the library now has a `verify` command
you can trust, a filehost that treats its input as untrusted, and a test suite
that includes the bug that was fixed.

The full review is in the source repo
([`docs/code-review.md`](https://github.com/uidbz/tie/blob/master/docs/code-review.md)).
It is worth reading as a picture of what "done" meant at that point, not as a
list of current bugs.
