---
title: The apps
type: apps
eyebrow: tie-gui
mark: tie
hue: "#2dd4bf"
tagline: "Five native apps on one library."
summary: >-
  tie-fm, tie-audio, tie-view, imgview and tie-browser are Go and Fyne programs
  that share a theme, a URL scheme and a tag vocabulary. They talk to the same
  triples and the same content-addressed blobs — and imgview will happily open a
  folder with no library at all.
badges:
  - "Linux"
  - "Windows"
  - "macOS"
  - "Android"
  - "BSD-3"
primary:
  label: "Install the apps"
  url: "/install/"
secondary:
  label: "Read the manual"
  url: "/docs/apps/"
docs:
  - title: "Quickstart"
    url: "/docs/apps/quickstart/"
    desc: "Install, connect to a library, and open your first album or folder."
    icon: "play"
  - title: "Configuration"
    url: "/docs/apps/configuration/"
    desc: "Shared tie profiles, per-app collections, themes and shortcuts."
    icon: "wrench"
  - title: "Architecture"
    url: "/docs/apps/architecture/"
    desc: "The shared gallery engine, thumbnails, threading and the Fyne fork."
    icon: "layers"
---

<div class="t-panes" style="margin-top: 1rem">
  <div class="t-pane">
    <div class="t-pane-bar">
      <img class="pane-mark" src="/brand/tie-fm.svg" alt="" />
      your disk
      <span class="t-spacer"></span>
      <span>/home/you/Music</span>
    </div>
    <div class="t-pane-body">
      <div class="t-pane-row"><span class="dim">▸</span> The Annealers</div>
      <div class="t-pane-row"><span class="dim">▸</span> Ostrander</div>
      <div class="t-pane-row"><span class="dim">▸</span> Rook &amp; Rowan</div>
      <div class="t-pane-row"><span class="dim">  </span> notes.md</div>
      <div class="t-pane-row"><span class="dim">  </span> covers.zip</div>
    </div>
  </div>
  <div class="t-pane">
    <div class="t-pane-bar">
      <img class="pane-mark" src="/brand/tie.svg" alt="" />
      the library
      <span class="t-spacer"></span>
      <span>tie:/music</span>
    </div>
    <div class="t-pane-body">
      <div class="t-pane-row"><span class="dim">▸</span> The Annealers — Slow Cooling <span class="tag">2026</span></div>
      <div class="t-pane-row"><span class="dim">▸</span> Ostrander — Night Ferry <span class="tag">jazz</span></div>
      <div class="t-pane-row"><span class="dim">▸</span> Rook &amp; Rowan — The Long Way Home</div>
      <div class="t-pane-row"><span>⤷</span> 48 tracks <span class="dim">· 4.2 MiB</span></div>
      <div class="t-pane-row"><span class="tag">tagged</span><span class="tag">deduplicated</span><span class="tag">versioned</span></div>
    </div>
  </div>
</div>

## One library, five ways in

Each app keeps its own collection profile (`images`, `audio`, `files`) and its own
window, but they share the same `tie:` namespace. That is why tie-fm can hand an
image to tie-view, and why tie-audio can play the album you found in the file
manager without importing it again.

<div class="t-grid t-grid--3 t-gap-lg" style="margin-top:1.5rem">
  <div class="t-principle">
    <div class="pj">A shared gallery engine</div>
    <div class="pd">Justified rows, thumbnails, caching and keyboard navigation are one package (<code>gallery/</code>) used by imgview, tie-view and tie-audio's cover wall.</div>
  </div>
  <div class="t-principle">
    <div class="pj">A shared style</div>
    <div class="pd">One theme package (<code>tiestyle</code>) gives every app the same surfaces, teal accent, light/dark variants and contrast floor.</div>
  </div>
  <div class="t-principle">
    <div class="pj">A shared hand-off</div>
    <div class="pd">Tiles and rows carry <code>tie:&lt;hash&gt;</code> URLs. Other apps open them in place, with the right collection already selected.</div>
  </div>
</div>

## Pick the one you need

<div class="t-grid t-grid--2 t-gap-lg" style="margin-top:1.5rem">
  <div class="t-card">
    <div class="t-h3">If you just want to look at pictures</div>
    <p class="t-app-line">imgview needs no library, no import and no server. Point it at a folder, an archive or a video file.</p>
    <a class="t-more" href="//apps/imgview/">Open imgview →</a>
  </div>
  <div class="t-card">
    <div class="t-h3">If you are building a library</div>
    <p class="t-app-line">Use tie-fm for the import and the cleanup, tie-view for tagging and rating, and tie-audio for listening.</p>
    <a class="t-more" href="//tie/">Start with tie →</a>
  </div>
</div>

## A note on platforms

<div class="t-chip" style="margin: 0.2rem 0">Linux</div>
<div class="t-chip" style="margin: 0.2rem 0">Windows</div>
<div class="t-chip" style="margin: 0.2rem 0">macOS</div>
<div class="t-chip" style="margin: 0.2rem 0">Android</div>

The five apps are Linux-first, and the desktop builds work on Windows and macOS
through Fyne. tie-view and tie-audio also ship Android builds, and imgview has a
folder picker for the same. tie-fm's device browsing (USB drives, MTP phones)
is Linux-only because it uses udisks2 and libmtp directly.
