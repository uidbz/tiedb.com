---
title: "Why triples, and why not just SQL"
type: blog
date: 2026-11-01
description: "An honest look at the least fashionable part of tie, including the sentence where it admits SQLite would have been easier."
---

There is a section in tie's internals note called "Why triples, and why not just
SQL". It contains this sentence:

> We rebuilt a database… SQLite provides all of that for free.

That is the right place to start, because the interesting part of tie was never
the triple store. It is the combination of a **content-addressed blob store** and
a **FUSE tag tree**, joined by the content hash. The triple store exists because
the schema is open and the core query is a reverse tag lookup, but a project that
only wanted a database should use SQLite.

What triples buy in practice:

- **No migrations.** A new relation is just a triple. The apps add `rating`,
  `member-of`, `play-count`, `thumbnail`, `dimensions` without a schema change.
- **Multi-valued by default.** A file has many tags and many paths. In SQL this
  is a join table; here it is the native shape.
- **Reverse queries are the point.** "Which files have this tag?" and "which
  scenes use this photo?" are the same operation, and the opt-in reverse index
  makes it cheap.
- **A path is a query.** The mount turns the index into a directory tree. That is
  not a SQL feature; it is the API the apps and the file manager share.

The cost is also real: a custom storage engine, a custom query path, a custom
fsck. The note documents the trade, including RDF/SPARQL and Datomic as prior
art.

If you are choosing a database for a new project, read the note before choosing
this one. If you want your photos, music and files in one place with tags that
outlive the paths, the combination is worth it.

[`docs/internals.md`](https://github.com/uidbz/tie/blob/master/docs/internals.md)
