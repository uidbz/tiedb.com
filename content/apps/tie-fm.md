---
title: tie-fm
type: app
app: tie-fm
docs:
  - title: "tie-fm manual"
    url: "/docs/apps/tie-fm/"
    desc: "Panes, tabs, bookmarks, file associations and MusicBrainz."
    icon: "book"
  - title: "Album import"
    url: "/docs/apps/album-import/"
    desc: "Destination templates, disc placement and the album planner."
    icon: "music"
  - title: "Delete and versions"
    url: "/docs/apps/delete/"
    desc: "What Delete does to a tie row, and where the bytes go."
    icon: "shield"
---

<div class="t-panes" style="margin-top: 1rem">
  <div class="t-pane">
    <div class="t-pane-bar">
      <img class="pane-mark" src="/brand/tie-fm.svg" alt="" />
      local
      <span class="t-spacer"></span>
      <span>/home/you/Music</span>
    </div>
    <div class="t-pane-body">
      <div class="t-pane-row"><span class="tag">dir</span> The Annealers</div>
      <div class="t-pane-row"><span class="tag">dir</span> Ostrander</div>
      <div class="t-pane-row"><span class="tag">dir</span> Rook &amp; Rowan</div>
      <div class="t-pane-row"><span class="tag">flac</span> 01 - First Frost.flac</div>
      <div class="t-pane-row"><span class="tag">jpg</span> cover.jpg</div>
      <div class="t-pane-row"><span class="tag">zip</span> artwork.zip</div>
    </div>
  </div>
  <div class="t-pane">
    <div class="t-pane-bar">
      <img class="pane-mark" src="/brand/tie.svg" alt="" />
      tie
      <span class="t-spacer"></span>
      <span>tie:/music</span>
    </div>
    <div class="t-pane-body">
      <div class="t-pane-row"><span class="tag">audio-dir</span> Slow Cooling <span class="dim">4 tracks</span></div>
      <div class="t-pane-row"><span class="tag">audio-dir</span> Night Ferry <span class="dim">4 tracks</span></div>
      <div class="t-pane-row"><span class="tag">audio-dir</span> The Long Way Home <span class="dim">4 tracks</span></div>
      <div class="t-pane-row"><span class="dim">tags:</span> <span class="tag">jazz</span><span class="tag">demo</span></div>
      <div class="t-pane-row"><span class="dim">48 tracks · 4.2 MiB · deduplicated</span></div>
    </div>
  </div>
</div>

## A file manager for people with two kinds of files

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Panes</span>
    <h2 class="t-h2">Both worlds, side by side.</h2>
    <p class="t-sub">
      A twin-panel file manager with tabs per pane, independent history, a
      breadcrumb you can type into and bookmarks that remember which tie
      collection they belong to. The left pane can be a local disk; the right pane
      can be the library; drag between them.
    </p>
    <ul class="t-list">
      <li><strong>Copy into tie as ▸</strong> with the eight built-in dir-types, or one of your own.</li>
      <li><strong>Pause, resume, stop</strong> a copy. Conflicts offer skip, rename, overwrite and apply-to-all.</li>
      <li><strong>Preview mode</strong> turns either pane into a thumbnail grid with an in-pane image viewer and video player.</li>
      <li><strong>Hand-off</strong>: open a tie image in tie-view, or a track in tie-audio, from the row's own menu.</li>
    </ul>
  </div>
  <div>
    <span class="t-eyebrow">Import</span>
    <h3 class="t-h3" style="margin-top: 0.9rem">Albums, not folders of tracks.</h3>
    <p class="t-sub" style="font-size: 0.95rem">
      "Import as albums…" discovers albums from a music root, matches them against
      MusicBrainz (a local Postgres mirror or the web service), fills missing
      covers, and places discs with templates like
      <code>/{albumartist}/{year} - {album}/cd{disc}</code>. The dry run shows you
      every destination before a byte moves.
    </p>
    <ul class="t-list">
      <li><strong>By directory, by tags, or auto.</strong></li>
      <li><strong>Use as album cover</strong> from any image in the pane.</li>
      <li><strong>Re-import is a version</strong>, not a duplicate: contents are content-addressed.</li>
    </ul>
  </div>
</div>

<div class="t-feature t-feature--flip">
  <div>
    <span class="t-eyebrow">Tags</span>
    <h2 class="t-h2">The tag panel does both jobs.</h2>
    <p class="t-sub">
      With nothing selected, the tag panel filters the pane. With rows selected,
      the same panel edits their tags. Directories are taggable too, and favourite
      tags are shared with tie-view and tie-audio.
    </p>
    <ul class="t-list">
      <li><strong>Free-text search</strong> in tie panes, ranked, with field selection.</li>
      <li><strong>Breeze file-type icons</strong> in light and dark, with per-extension associations.</li>
      <li><strong>Devices</strong> (Linux): USB drives through udisks2 and MTP phones through libmtp, with a gvfs fallback.</li>
    </ul>
  </div>
  <div class="t-term">
    <div class="t-term-bar">
      <img src="/brand/tie-fm.svg" alt="" />
      <span>keys</span>
    </div>
    <pre class="t-term-body"><code><span class="t-a">Ctrl+I</span>   filter this page     <span class="t-a">Ctrl+T</span>  new tab
<span class="t-a">Ctrl+Tab</span> switch pane          <span class="t-a">Ctrl+W</span>  close tab
<span class="t-a">Ctrl+click</span> open in new tab   <span class="t-a">Esc</span>     leave preview
<span class="t-a">Shift/Ctrl</span> multi-select       <span class="t-a">⌘</span>       the same on macOS</code></pre>
  </div>
</div>

## What Delete means

tie-fm is careful about the word "delete". On a local disk it is the file
manager delete you expect. On a `tie:` row it removes the reference from the
tree, keeps the blob (other paths may point at the same content), and the library
knows what became unreferenced. The manual has the exact semantics and the
roadmap for blob garbage collection.
