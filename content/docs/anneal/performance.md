---
title: Performance
type: docs
weight: 7
---

anneal has a per-run timing tracer, exposed as `/perf` and written into the
session. The source note,
[`docs/PERFORMANCE.md`](https://github.com/uidbz/anneal/blob/main/docs/PERFORMANCE.md),
includes the baseline and the comparison harness against opencode.

## What is measured

- Time to first token and total run time.
- Tool call latency, by tool.
- Compaction and summary costs.
- Prompt assembly and memory injection.
- The improvement loop: propose, probe and compose per candidate.

The probe records steps, failures and token counts, because "better" has to mean
something measurable. A candidate with no measurable gain is skipped and the
reason is recorded in the run.

## Reproducing

The harness in `testing/perf/` runs the same task against anneal and opencode,
with a local model, and compares the traces. It is intentionally boring: a
fixed task, a fixed model, a fixed machine, and the raw numbers saved.

```bash
cd testing/perf
./ab.sh
```

Do not trust a marketing claim that is not in this note. If the number changes,
the note changes first.
