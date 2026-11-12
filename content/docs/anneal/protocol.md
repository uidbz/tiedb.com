---
title: Protocol
type: docs
weight: 4
---

`anneald` speaks a versioned WebSocket protocol on `/ws` (and a `/healthz`
endpoint). The desktop app, the headless CLI and any future client use the same
protocol.

## Framing

- A handshake with the protocol version (`proto: "v1"`), a token and the client
  identity.
- JSON frames with a request id; responses and events share the stream.
- Methods cover sessions, prompts, approvals, the queue, harness operations,
  memory, skills, the journey, improve runs and the work tree.
- The server can transfer a work tree to the client for remote projects.

The normative description (frame types, auth, versioning, error codes) is
[`docs/PROTOCOL.md`](https://github.com/uidbz/anneal/blob/main/docs/PROTOCOL.md).

## Why a protocol

The GUI, the CLI and the container backend need the same behaviour. A protocol
also makes it possible to run the agent on a workstation and the UI on a laptop,
without either one owning the state.

## Compatibility

The protocol version is independent of the app version. Additive methods and
fields are allowed within v1; a breaking change bumps the version and the server
rejects a mismatch with a clear error.
