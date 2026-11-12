---
title: anneal
type: showcase
mark: anneal
hue: "#e6a23c"
eyebrow: "The agent"
tagline: "The coding agent that anneals."
summary: >-
  A local-first coding agent that works in your project, learns from finished
  sessions, and proposes better harnesses — with evidence, and only if you
  promote them.
badges:
  - "Linux desktop"
  - "Headless CLI"
  - "Container"
  - "dev"
primary:
  label: "Read the guide"
  url: "/docs/anneal/using/"
secondary:
  label: "Architecture"
  url: "/docs/anneal/architecture/"
hero_shot:
  app: anneal
  scene: chat
  themes: dark
  meta: "chat · 1280×800"
  alt: "anneal's chat view: a tool-using session with an approval card, a plan, a context meter and a task list"
  caption: "A permission request, a tool trail and a plan. Every action is visible, and every approval is yours."
docs:
  - title: "Using anneal"
    url: "/docs/anneal/using/"
    desc: "First run, the GUI tour, slash commands and the learning loop."
    icon: "play"
  - title: "Configuration"
    url: "/docs/anneal/config/"
    desc: "Every key in config.json, the env overrides, and where state lives."
    icon: "wrench"
  - title: "Architecture"
    url: "/docs/anneal/architecture/"
    desc: "The event bus, the backend seam, permissions and the harness."
    icon: "layers"
---

## The idea

Most coding agents are a model in a chat box. anneal is a **harness** around the
model that you can inspect and improve: prompts, tools, guardrails, skills and
review rules. When a session ends, anneal can look at what happened, propose a
change to its own harness, test the change on the session it came from, and keep
it only if the evidence is there.

The name is the loop. You heat the work up with a task, then anneal it slowly:
review, memory, a better harness. Nothing is adopted just because a model said it
was a good idea.

<div class="t-grid t-grid--3 t-gap-lg" style="margin: 1.8rem 0 0.5rem">
  <div class="t-principle">
    <div class="pj">01 · Work</div>
    <div class="pd">Chat with tools, approvals, a queue, plan/build modes and a turn budget. The transcript is the work.</div>
  </div>
  <div class="t-principle">
    <div class="pj">02 · Learn</div>
    <div class="pd">A finished session becomes lessons, memory and skills. Cross-session recall lives in tie.</div>
  </div>
  <div class="t-principle">
    <div class="pj">03 · Improve</div>
    <div class="pd">Propose, probe, compose. A successor harness is tested, then waits for an explicit promote.</div>
  </div>
</div>

## Work

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Tools</span>
    <h2 class="t-h2">A real working set, not a chat toy.</h2>
    <p class="t-sub">
      Read, write, edit, bash, glob, grep, fetch and todos. Streaming responses
      with live tool calls, a transcript that keeps the tool trail, and a composer
      that stays usable while a task runs.
    </p>
    <ul class="t-list">
      <li><strong>Approvals</strong> in the transcript: allow once, always, this session, or reject. Saved project grants, OpenCode-style rules, a read-only bash allowlist.</li>
      <li><strong>Plan mode</strong> is read-only. <strong>Build mode</strong> asks. Switch with one control.</li>
      <li><strong>A queue</strong> lets you type the next thing while the current turn runs, with visible delivery state.</li>
    </ul>
  </div>
  <div>
    {{< shot app="anneal" scene="harness" meta="harness · 1280×800" title="anneal" alt="anneal's Harness view: editable system prompt and guardrails, component counts and an improvement library" caption="The Harness view is a fynerisor pane: the system prompt, guardrails and improvement library are editable in place." >}}
  </div>
</div>

<div class="t-feature t-feature--flip">
  <div>
    <span class="t-eyebrow">Budget</span>
    <h2 class="t-h2">A turn has a shape.</h2>
    <p class="t-sub">
      An active-time budget warns at 80% and, at 100%, writes a hand-off summary
      and stops. Continue with one key. Context is compacted at 82% of the
      window, and the context meter in the inspector tells you where you are.
    </p>
    <ul class="t-list">
      <li><strong>Tasks</strong> with a staleness warning when the plan and the actual work drift apart.</li>
      <li><strong>Subagents</strong>: explore, plan and general-purpose, each with its own context.</li>
      <li><strong>MCP</strong> over stdio or HTTP, with optional gortex code intelligence.</li>
    </ul>
  </div>
  <div>
    {{< shot app="anneal" scene="settings" meta="settings · 1280×800" title="anneal" alt="anneal's Settings view with sections for appearance, model, agent, learning and context" caption="Settings is a form over config.json with live changes — and an Edit JSON escape hatch." >}}
  </div>
</div>

## Learn

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Memory</span>
    <h2 class="t-h2">The project remembers.</h2>
    <p class="t-sub">
      Project memory and user memory are bounded files with a duplicate guard and
      an injection scan. Skills are <code>SKILL.md</code> files the agent can
      write, review and reuse. Neither is a hidden vector store: they are files
      you can read, diff and delete.
    </p>
    <ul class="t-list">
      <li><strong>Sessions</strong> are JSONL under <code>~/.anneal/sessions</code>, replayable and resumable per project.</li>
      <li><strong>Recall</strong> is <code>session_search</code>, a cross-session search over a tie index.</li>
      <li><strong>Review</strong> runs in the background while idle, mining finished sessions for lessons.</li>
    </ul>
  </div>
  <div class="t-term">
    <div class="t-term-bar">
      <img src="/brand/anneal.svg" alt="" />
      <span>the loop in practice</span>
    </div>
    <pre class="t-term-body"><code><span class="t-o">a session ends, the reviewer runs</span>

<span class="t-a">REVIEW</span>       <span class="t-o">saved 3 lessons · 1 skill</span>
<span class="t-a">MEMORY-ADD</span>   <span class="t-o">target: user</span>
<span class="t-a">IMPROVE</span>      <span class="t-o">run 20261030-101500</span>
             <span class="t-g">skills  probed: 2 steps → adopted</span>
             <span class="t-a">prompt  skipped: no measurable gain</span>
             <span class="t-o">decision: evidence-backed, awaiting promote</span></code></pre>
  </div>
</div>

## Improve

The improvement stack is the part that makes anneal different. It is not a
prompt that says "improve yourself". It is a loop with a budget, a probe and a
gate.

<div class="t-grid t-grid--4 t-gap-lg" style="margin: 1.8rem 0">
  <div class="t-principle"><div class="pj">Propose</div><div class="pd">A candidate change to a skill, a prompt or a guardrail, derived from a finished session.</div></div>
  <div class="t-principle"><div class="pj">Probe</div><div class="pd">Run the session's hard parts against the candidate and the baseline. Measure steps, failures and tokens.</div></div>
  <div class="t-principle"><div class="pj">Compose</div><div class="pd">Combine the candidates that showed a gain. No gain means no change.</div></div>
  <div class="t-principle"><div class="pj">Promote</div><div class="pd">A successor harness sits in an overlay until you explicitly promote it.</div></div>
</div>

<div class="t-feature t-feature--flip">
  <div>
    <span class="t-eyebrow">Evidence</span>
    <h2 class="t-h2">The Activity view is the receipt.</h2>
    <p class="t-sub">
      Every review, memory addition and improvement run is a card on the journey
      timeline. You can see which candidate was adopted, which was skipped, what
      the probe cost, and why. k-step recursion is visible, not implied.
    </p>
    <ul class="t-list">
      <li><strong>An improvement library</strong> keeps successors and their evidence.</li>
      <li><strong>Ctrl+I</strong> starts an improvement run; idle time starts one automatically.</li>
      <li><strong>Promote</strong> is the only path from an overlay to the harness you actually run.</li>
    </ul>
  </div>
  <div>
    {{< shot app="anneal" scene="activity" meta="activity · 1280×800" title="anneal" alt="anneal's Activity view: a journey timeline of reviews, memory additions and improvement runs, with probe results" caption="The Activity view: a journey of reviews and improvement runs, with the adopted and skipped candidates shown." >}}
  </div>
</div>

## Run anywhere

| Surface | What it is |
|---|---|
| `anneal` | The Fyne desktop app. Linux is the supported GUI target. |
| `anneal-headless` | The same core without GUI libraries: `-p`, `-i`, `-json`, `-yolo`. |
| `anneald` | A CGO-free backend with a versioned WebSocket protocol, containers and idle shutdown. |

```
$ anneal-headless -p "Fix the parser test in core/textutil" -json
$ anneald -work ~/src -listen :7777
```

One instance per project. The desktop app can embed the core or attach to a
remote `anneald` with `-server`; sessions, memory and the tie index stay per
project. The Docker image is the backend, not a desktop in a box.

## Local models, your choice

The default model is a local Strata server at `127.0.0.1:8081`, with sample
models like `qwen3-coder-30b-a3b` at a 262k context. Any OpenAI-compatible
chat-completions endpoint works. There is no anneal account and no model proxy:
the agent talks to the endpoint you configure.

## Scriptable

The Harness view is a Risor script hosted by fynerisor, sharing the host theme.
The agent also has a sandboxed Risor `script` tool: confined filesystem, no exec,
no network. If you want a pane, a rule or a report, you can write it without
rebuilding anneal.

## Status

> [!WARNING]
> anneal is young: milestones M0–M16 are complete, the version is `dev`, and
> there are no tagged releases yet. Install it from source with `make install`,
> or build the container with `make docker`. The wire protocol is versioned
> separately as v1. Expect the harness format to keep moving.

The manual's [performance note](/docs/anneal/performance/) includes the per-run
tracing and the opencode comparison harness, so you can reproduce the numbers
rather than take them on faith.
