---
title: "Halving the heap"
type: blog
date: 2026-11-05
description: "Three memory optimizations, two reverted experiments and the discipline of measuring on the real database."
---

The real benchmark library is 7.4 GiB on disk with 60.7 million entries and 98.2
million associations. Loading it used to take 19.4 GiB of heap. The optimization
note in the source repo is a record of what actually moved the number.

At 1M triples, the heap went **862 → 786 → 633 → 473 MiB** — a 45% reduction —
and objects per triple went from **21.5 to 10.0**. Three changes did most of the
work:

1. **A `TYPE_HASH` record.** Hashes were being stored as strings in the trie.
   A raw 32-byte record cut the heap by **28%** at 100k triples.
2. **Interning the trie chunks.** Identical chunks in different levels are
   referenced in memory but written to disk, which is the same trade the
   original design made for strings.
3. **Sharding the outer tree.** Load time went from **4m32s to 2m46s** (−39%)
   and the peak heap came down with it.

Two other experiments were reverted. They looked better in a microbenchmark and
were worse on the real database. The note says why, and that is the point:

> Measure both axes on the real DB. Guessing that a design "should" be better is
> not sufficient.

The opt-in reverse index saves about 29% of heap. That is a real trade: reverse
queries become fast, but the index is opt-in per relation and the server's
default subset is deliberately small (`tag,path,parent,tie-type,version-of,member-of`).

The full note, including the two reverted experiments, is
[`docs/memory-optimization.md`](https://github.com/uidbz/tie/blob/master/docs/memory-optimization.md).
