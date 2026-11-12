---
title: The compute plane
type: docs
weight: 23
---

`tie-runner` is optional. It turns content-addressed trees into a compute
substrate: a job takes an input hash, runs in a container, and produces an output
hash.

```bash
tie-run run ghcr.io/you/tool:1 <input-hash> <output-hash>
tie-run status
tie-run jobs
tie-run logs <job>
tie-run deploy --prod <image>
```

The runner has its own HTTP API (port 1163) with jobs, pipelines, services,
capacity, a queue and cancellation. `tie-exec` is the container entrypoint: it
reads inputs and writes outputs with `#!TIE` markers on stdout. `tie-fleet` is a
thin admin client for fleets of runners (brandable via `fleetcmd.Run`).

The normative protocol is in
[`docs/runner-api.md`](https://github.com/uidbz/tie/blob/master/docs/runner-api.md);
the wire fixtures are in `testdata/runner/`.

This is the most opinionated part of tie and the least necessary: if you want a
library, you never need to start a runner. If you want reproducible derivatives
— thumbnails, transcodes, OCR, ML inference — the hash-in/hash-out contract is
the point.
