---
title: Projects
type: docs
weight: 6
---

anneal is one instance per project. A project has its own session list, memory
files, skills, tie index and harness; the model configuration is shared.

## Several side by side

- Desktop: one window per project (`anneal -C ~/src/one`), or one GUI attached to
  separate backends.
- Headless: `anneal-headless -C ~/src/one -p "..."` in a script or a container.
- Server: `anneald -work /projects/one -listen :7777`, one port per project (or
  one reverse proxy route per project).

## Cross-project hand-offs

The session index in tie is per project, so recall does not leak between
projects. If you want a shared library of lessons, copy a skill or a memory file
between project state directories deliberately; anneal does not do it for you.

## The work tree

A remote backend can transfer a work tree to the client (and back) over the
protocol. That is how a thin UI can work on a containerised project without a
shared filesystem.

The recipes and the tenancy notes are in
[`docs/PROJECTS.md`](https://github.com/uidbz/anneal/blob/main/docs/PROJECTS.md).
