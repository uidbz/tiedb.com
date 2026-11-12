---
title: Architecture
type: docs
weight: 3
---

anneal is divided into a core that has no UI, a wire protocol, and two front
ends (the Fyne desktop shell and the terminal CLI).

## Packages

| Package | Role |
|---|---|
| `core/app` | Wiring: config, stores, agent, server. |
| `core/agent` | The runner: prompts, tools, hooks, events. |
| `core/events` | A UI-agnostic event bus. |
| `core/server`, `core/client` | The WebSocket protocol and its client. |
| `core/uiapi` | The fynerisor script bridge. |
| `core/model` | The Strata/OpenAI-compatible provider and health monitor. |
| `core/memory`, `skills`, `tieindex`, `review`, `journey` | The learning-loop stores. |
| `core/harness`, `harnesslib`, `improve`, `probe`, `toolset` | The improvement stack. |
| `core/permission`, `script`, `config`, `paths`, `fsutil`, `textutil` | Support. |

The event bus is the seam. The UI subscribes; the agent does not know whether it
is driving a Fyne window, a terminal, or a remote client.

## Threading

The agent runs on its own goroutine and emits events. The Fyne shell applies
them with the `guithread` helper (`fyne.Do` off the UI thread). The Harness pane
runs a Risor VM with its own UI thread guard. Network and tool work never blocks
the UI thread.

## The improvement loop

1. A finished session is reviewed (idle or explicit). Lessons and skills are
   written to bounded files.
2. An improvement run proposes candidates: a skill, a prompt or a guardrail.
3. The probe runs the session's hard parts against each candidate and the
   baseline, and records steps, failures and tokens.
4. Compose keeps the candidates that showed a gain.
5. A successor harness is written as an overlay and added to the library.
6. **Promote** copies the overlay into the active harness. Nothing is promoted
   automatically.

Every step is an event on the journey timeline, which is what the Activity view
displays.

The full diagram and the event table are in
[`docs/ARCHITECTURE.md`](https://github.com/uidbz/anneal/blob/main/docs/ARCHITECTURE.md).
