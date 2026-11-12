---
title: Configuration
type: docs
weight: 2
---

`config.json` lives in `~/.anneal` (or `ANNEAL_HOME`). The Settings view edits it
with live changes; **Edit JSON** opens the raw file.

## Top-level keys

| Key | What it controls |
|---|---|
| `server` | The model endpoint, model name, context window, timeouts. |
| `provider` | Health checks and fallbacks. |
| `agent` | Tools, turn budget, compaction, subagents. |
| `permission` | Allow/ask/deny rules, the read-only bash allowlist, saved grants. |
| `ui` | Theme, text scale, layout, pane sizes. |
| `memory` | Caps, duplicate guard, injection scan. |
| `skills` | Skill directories and authoring. |
| `review` | Idle review, lesson caps, model. |
| `tie` | The session index: file/URL, collection, offline behaviour. |
| `mcp` | stdio and HTTP servers. |
| `harness` | Components, guardrails, the improvement library. |
| `improve` | k-step recursion, timeouts, the promotion policy. |
| `script` | The Risor tool's sandbox roots. |
| `compaction` | The threshold (82%) and the summary prompt. |

## Environment

| Variable | Effect |
|---|---|
| `ANNEAL_HOME` | The state directory. |
| `ANNEAL_THEME` | `light` or `dark`. |
| `ANNEAL_UI_DIR` | A directory of scripted panes for tests. |
| `ANNEAL_MODEL` / provider vars | Override the model endpoint. |

## Example

```json
{
  "server": {
    "url": "http://127.0.0.1:8081/v1/chat/completions",
    "model": "qwen3-coder-30b-a3b",
    "context_window": 262144
  },
  "ui": { "theme": "dark", "text_scale": 1.0 },
  "permission": {
    "bash": { "allow": ["git status", "go test ./...", "ls"] }
  },
  "tie": { "enabled": true, "collection": "anneal" },
  "improve": { "k": 2, "mode": "sequential", "promote": "manual" }
}
```

The full reference, with every field and default, is
[`docs/CONFIG.md`](https://github.com/uidbz/anneal/blob/main/docs/CONFIG.md).
