---
title: Using anneal
type: docs
weight: 1
---

## Install

```bash
git clone https://github.com/uidbz/anneal
cd anneal
make build          # bin/anneal, bin/anneal-headless, bin/anneald
make install        # ~/.local/bin, .desktop entry, hicolor icon
```

Go 1.27+ and the Fyne C dependencies (X11/Wayland and OpenGL) are required for
the GUI. `anneal-headless` and `anneald` build without them. The container image
is `make docker` (`anneal:dev`).

The state directory is `~/.anneal` (`ANNEAL_HOME`). Sessions are JSONL under
`~/.anneal/sessions`.

## First run

1. Open a project (`-C <dir>` or the project picker).
2. Configure a model: a local Strata server or any OpenAI-compatible
   chat-completions endpoint. The default is `http://127.0.0.1:8081/v1/chat/completions`.
3. Say what you want. `/` opens the slash commands; `Ctrl+1–4` switch Chat,
   Harness, Activity and Settings.

## The GUI

- **Rail** (left): Chat, Harness, Activity, Settings.
- **Sessions** (sidebar): filter, groups, resume.
- **Transcript**: user and assistant turns, tool calls and results, approvals.
- **Composer**: Build/Plan, model and context, queue, Stop.
- **Inspector** (right): tasks, context meter, session ID, working directory.
- **Status bar**: model, index and MCP health.

## Plan and Build

Plan mode is read-only: the agent can inspect and propose, not edit. Build mode
has the tools. The approval card offers allow once, always, this session, or
reject; project grants are saved.

## The queue

Type while a task runs. The tray shows each message's delivery state. The queue
is part of the session, not a separate chat.

## Turn budget

The run has an active-time budget. At 80% you get a warning; at 100% anneal
writes a hand-off summary and stops. Continue resumes with the summary in
context. Compaction happens at 82% of the context window.

## Slash commands

| Command | What it does |
|---|---|
| `/improve` | Start an improvement run (also `Ctrl+I`). |
| `/learn` | Ask the reviewer to mine the current session. |
| `/memory` | Show or edit project/user memory. |
| `/journey` | Open the Activity timeline. |
| `/theme` | Switch light/dark/system. |
| `/help` | The full list. |

## Headless

```bash
anneal-headless -p "Fix the parser test"          # one prompt, print the result
anneal-headless -i                                # interactive terminal
anneal-headless -p "..." -json                    # event stream
anneal-headless -p "..." -yolo                    # skip approvals (careful)
anneal-headless -session <id> -p "Continue"
anneal-headless -improve -k 2
```

## Remote backend

```bash
anneald -work ~/src -listen :7777
anneal -server ws://host:7777 -token "$ANNEAL_TOKEN"
```

The wire protocol is versioned separately (v1); see
[Protocol](/docs/anneal/protocol/). The desktop app can embed the core or attach
to a remote backend.

## Scripting

The Harness view is a fynerisor pane (`assets/ui/harness.risor`). The agent also
has a sandboxed Risor `script` tool: confined filesystem, no exec, no network.
For a custom pane or report, write Risor against the `uiapi` bridge.
