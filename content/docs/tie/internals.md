---
title: Internals
type: docs
weight: 3
---

## The file

A `.tie` file is a flat log of fixed 50-byte little-endian records:

| Type | Size | Payload |
|---|---|---|
| `TYPE_ENTRY` (11) | 50 | a 24-byte trie chunk |
| `TYPE_ASSOCIATION` (12) | 50 | three IDs: key, relation, value |
| `TYPE_HASH` (13) | 50 | a raw 32-byte hash |
| `TYPE_DELETE` (10) | 50 | a tombstone; the slot is reused |

Strings are interned into a per-level trie of 24-byte chunks. There are forward
and reverse association indexes; the reverse index is opt-in per relation and
defaults to `tag,path,parent,tie-type,version-of,member-of`.

An `AssociationSet` is tiered: inline, then a slice, then a 16-shard map above
256 entries. One writer goroutine fsyncs on a 10-second ticker and drains on
SIGTERM. Load is two-pass and parallel. New files are created `0600`.

The stored and wire formats are the stable contract. The storage layer is
deliberately boring: a log, a trie, two indexes, and a writer.

## Memory

The real benchmark database is 7.4 GiB on disk with 60.7 million entries and
98.2 million associations. It loads in 2m46s with a 19.4 GiB heap on 8 workers.

The optimization note documents the path:

- at 1M triples, the heap went 862 → 786 → 633 → 473 MiB (−45%), and objects per
  triple went from 21.5 to 10.0;
- the opt-in reverse index saves about 29% of heap;
- the `TYPE_HASH` record cut heap by 28% at 100k triples;
- sharding the outer tree cut load time by 39% (4m32s → 2m46s).

Two experiments were reverted because the real database disagreed with the
microbenchmark. The note's discipline is the point: "measure both axes on the
real DB. Guessing that a design 'should' be better is not sufficient."

## Why triples, and why not just SQL

The honest answer is in the source note: SQLite would provide transactions,
indexes and durability for free, and the genuinely distinctive part of tie was
never the triple store. It is the **content-addressed blob store + FUSE tag
tree** combination. The triple store exists because the schema is open and the
reverse tag query is the core operation; RDF/SPARQL and Datomic are prior art.

If you are choosing a database for a new project, read the note before choosing
this one.

## The transport seam

```
TieClient → Doer → { webservice.Client (HTTP)
                   | localTransport → webservice.Dispatch → api.RegisterAll → tiedb }
```

Direct mode and the server use the same handlers. That is why `tie -f` has the
same behavior as the daemon, and why the Go client can be tested against an
in-process server.

The full note:
[`docs/internals.md`](https://github.com/uidbz/tie/blob/master/docs/internals.md).
