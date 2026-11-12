---
title: Containers
type: docs
weight: 5
---

The backend can run in a container with one bind mount and an idle shutdown.

```bash
docker build -t anneal:dev .
docker run -d --name anneal \
  -v ~/src:/work \
  -e ANNEAL_TOKEN=change-me \
  -p 7777:7777 \
  anneal:dev
```

- One work directory is mounted (`-work /work`).
- The server listens on `:7777`; `/ws` is the protocol, `/healthz` is liveness.
- Idle auto-shutdown stops the container when no client is connected and no run
  is active.
- Multi-user deployments use one instance per project (see
  [Projects](/docs/anneal/projects/)).
- There is a no-port exec transport for environments where a shell is the only
  channel.

The image is the backend, not a desktop in a box: the Fyne GUI needs a display
server. Use the headless CLI or a remote desktop app with `-server`.

[`docs/CONTAINERS.md`](https://github.com/uidbz/anneal/blob/main/docs/CONTAINERS.md)
has the details and the tenancy recipes.
