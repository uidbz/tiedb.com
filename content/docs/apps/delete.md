---
title: What Delete means
type: docs
weight: 22
---

tie-fm's Delete is deliberately conservative.

## A local row

Delete removes the file from the local filesystem, after the usual confirmation.
There is no trash integration beyond the desktop's own; the app uses the same
semantics as the rest of the pane.

## A tie row

Delete removes the **reference** from the virtual tree:

- the `parent` edge and the `path`/`name` metadata go away;
- the content-addressed blob stays, because other paths, versions or
  collections may point at the same hash;
- the tag triples stay on the hash, so a future reference sees them again.

That is why deleting a file in tie is not the same as destroying the bytes. Blob
garbage collection is the separate operation; the source doc,
[`docs/TIE-FM-DELETE.md`](https://github.com/uidbz/tie-gui/blob/master/docs/TIE-FM-DELETE.md),
describes the current state and the roadmap.

## Versions

If a file was replaced (a new hash under the same path), the old version is still
in the history:

```bash
tie versions list /photos/2014/reykjavik.jpg
tie versions restore /photos/2014/reykjavik.jpg <version>
```

`tie verify` will report a dangling reference if a manifest points at a missing
blob; `tie repair` can relink the tree from the manifests.
