---
title: Search
type: docs
weight: 12
---

tie has free-text search over named relations. There is no sidecar index and no
external service: the corpus is built in memory from the triples, cached, and
rebuilt when the collection changes.

```bash
tie search reykjavik
tie search "great wave" --relation title
tie search bach --mode fuzzy --type audio-dir
```

## Modes

| Mode | What it does |
|---|---|
| `words` | Token match, accent- and case-folded. The default. |
| `substring` | Match inside a word. |
| `prefix` | Match the start of a word. |
| `fuzzy` | Edit-distance match for typos. |

`--relation` limits the fields searched; `--type` scopes by dir-type;
`--weight` biases a field. The Go API exposes `Search` and `SearchPage`, and the
profiles `AudioSearchProfile` and `FileSearchProfile` bundle sensible defaults
for the apps.

## Measured

These are from the search note in the source (single run, Ryzen 9 9950X):

| Workload | Time |
|---|---|
| Cold build + query, 1k subjects | ~1.8 ms |
| Cold build + query, 10k subjects | ~18 ms |
| Cold build + query, 100k subjects | ~223 ms |
| Warm query, 10k subjects | ~1.2 ms |
| Warm query, 100k subjects (4 fields each) | ~15 ms |
| Warm fuzzy query, 50k subjects | ~18 ms |

The corpus cache defaults to 64 MiB. The note explains why SQLite FTS5, Bleve
and a sidecar index were rejected, and what would change the decision.

## In the apps

tie-audio's search uses the album/artist/track fields; tie-fm searches tie panes
with a field selector; tie-view filters by tags rather than free text. All of
them call the same client.
