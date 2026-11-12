---
title: tie-view
type: app
app: tie-view
docs:
  - title: "tie-view manual"
    url: "/docs/apps/tie-view/"
    desc: "Quick tags, ratings, the tag sidebar and startup pages."
    icon: "book"
  - title: "Quick tags"
    url: "/docs/apps/quick-tags/"
    desc: "The pills, their sizes, and the keyboard mapping in quicktags.toml."
    icon: "tag"
  - title: "The media model"
    url: "/docs/tie/guides/media-library/"
    desc: "How tags, ratings and collections fit together in tie."
    icon: "layers"
---

<div class="t-tagflow" style="margin-top: 1rem">
  <div class="t-tagflow-row">
    <span class="lbl">Image</span>
    <div class="t-tagflow-img" style="flex:1">
      <span class="thumb"></span>
      <div class="meta">
        <span style="color:var(--t-fg); font-weight:600">great-wave.jpg</span>
        <span>9a1f3c…c03e · 1920×1324</span>
        <span>★★★★☆</span>
      </div>
    </div>
  </div>
  <div class="t-tagflow-row">
    <span class="lbl">Tags</span>
    <span class="t-chip">travel</span><span class="t-chip">japan</span><span class="t-chip">ukiyo-e</span><span class="t-chip">favorite</span>
  </div>
  <div class="t-tagflow-row">
    <span class="lbl">Quick</span>
    <span class="t-chip" style="border-color:#e6a23c;color:#b45309">★</span>
    <span class="t-chip">family</span><span class="t-chip">wallpaper</span><span class="t-chip">print</span>
    <span style="color:var(--t-soft);font-size:.72rem">one tap, over the image edge</span>
  </div>
</div>

## Tag at the speed of sight

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Tagging</span>
    <h2 class="t-h2">The image is the interface.</h2>
    <p class="t-sub">
      Open an image and the tag panel is one tap away. Quick-tag pills sit over the
      edge of the picture: one tap, or one key, per tag. Rate with the number row,
      favourite with the heart, swipe up for the full tagger on a phone.
    </p>
    <ul class="t-list">
      <li><strong>Pills sized xs–xl</strong>, with an icon or just text, configured in <code>quicktags.toml</code>.</li>
      <li><strong>Ratings 1–5</strong> and a favourite tag that tie-audio and tie-fm also understand.</li>
      <li><strong>No modal dialogs</strong> in the flow: tag, rate, next.</li>
    </ul>
  </div>
  <div>
    <span class="t-eyebrow">Browsing</span>
    <h3 class="t-h3" style="margin-top: 0.9rem">Include, exclude, refine.</h3>
    <p class="t-sub" style="font-size: 0.95rem">
      The tag sidebar works like a query builder: include tags to narrow, exclude
      tags to remove, and "related tags" to refine by what the current result set
      actually contains. Favourite tags stay at the top.
    </p>
    <ul class="t-list">
      <li><strong>Sort</strong> by name, rating, newest or oldest.</li>
      <li><strong>Untagged</strong> is one button, because a library always has some.</li>
      <li><strong>A Files tab</strong> shows the virtual tree when you need a path instead of a tag.</li>
      <li><strong>Server-cached thumbnails</strong> and archive covers are generated once, then shared.</li>
    </ul>
  </div>
</div>

<div class="t-feature t-feature--flip">
  <div>
    <span class="t-eyebrow">Together</span>
    <h2 class="t-h2">The hand-off from tie-fm.</h2>
    <p class="t-sub">
      tie-fm can open a tile in tie-view with the right collection and the right
      file already selected. The mechanism is a <code>tie:&lt;hash&gt;</code> URL,
      not a temporary export. The same URL works from a terminal, a script or a
      bookmark.
    </p>
    <ul class="t-list">
      <li><strong>Desktop and Android</strong> from the same codebase; the drawer and swipe gestures adapt.</li>
      <li><strong>Connection errors</strong> show a Retry, not a dead window.</li>
      <li><strong>De-import</strong> from a tile's context menu, with the same semantics as tie-fm.</li>
    </ul>
  </div>
  <div class="t-term">
    <div class="t-term-bar">
      <img src="/brand/tie-view.svg" alt="" />
      <span>one URL, one image</span>
    </div>
    <pre class="t-term-body"><code><span class="t-p">$</span> <span class="t-cmd">tie-view "tie:9a1f3c0b7e…4d2ac03e"</span>

<span class="t-o"># or from tie-fm</span>
<span class="t-o">#   right-click → Open in tie-view</span>

<span class="t-o"># or pin a query as a startup page</span>
<span class="t-p">$</span> <span class="t-cmd">tie-view -tag favorite</span></code></pre>
  </div>
</div>

## Shortcuts

| View | Keys | Action |
|---|---|---|
| Gallery | `J` / `K`, arrows | Scroll |
| Gallery | `Backspace` | Parent folder |
| Image | `←` / `→` | Previous / next |
| Image | `T` | Quick tagging |
| Image | `I` | Info overlay |
| Image | `U` | Save image |
| Image | `Esc` | Back to the gallery |
| Any | `Q` | Quit |

Every shortcut is rebindable in the imgview config, which tie-view shares. The
Quick tags settings tab edits the pill layout and key mapping in place.
