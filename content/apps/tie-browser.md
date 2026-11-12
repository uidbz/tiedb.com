---
title: tie-browser
type: app
app: tie-browser
badges: ["preview"]
docs:
  - title: "tie-browser manual"
    url: "/docs/apps/tie-browser/"
    desc: "Scenes, props, actions, scripting and the designer."
    icon: "book"
  - title: "Scenes and fynerisor"
    url: "/ecosystem/"
    desc: "The Risor bindings and the scene renderer behind the pages."
    icon: "code"
  - title: "The media model"
    url: "/docs/tie/guides/media-library/"
    desc: "Where the images, audio and tags on a page come from."
    icon: "layers"
---

<div class="t-feature" style="border-top:0; padding-top:1rem">
  <div>
    <span class="t-eyebrow">The idea</span>
    <h2 class="t-h2">Pages made of your library.</h2>
    <p class="t-sub">
      A scene is a page stored in tie: a layout, some hotspots, and components
      that read your library live. Point one at a photo, another at a tag query,
      another at an audio file. "Which scenes use this photo?" is then a tie query,
      not a grep through HTML.
    </p>
    <ul class="t-list">
      <li><strong>Scenes are folders</strong>: <code>tie:/scenes/house/living-room/</code> holds <code>scene.json</code>, a thumbnail and an optional script.</li>
      <li><strong>Hotspots</strong> navigate, open or play. Containers lay out the rest.</li>
      <li><strong>Versioned saves</strong> with conflict handling: a scene is a directory in a content-addressed store.</li>
    </ul>
  </div>
  <div class="t-term">
    <div class="t-term-bar">
      <img src="/brand/tie-browser.svg" alt="" />
      <span>scene.json</span>
    </div>
    <pre class="t-term-body"><code>{
  <span class="t-b">"format"</span>: <span class="t-g">"fynerisor.scene/1"</span>,
  <span class="t-b">"title"</span>: <span class="t-g">"Living room"</span>,
  <span class="t-b">"root"</span>: {
    <span class="t-b">"type"</span>: <span class="t-g">"free"</span>,
    <span class="t-b">"props"</span>: { <span class="t-b">"designW"</span>: 1600, <span class="t-b">"background"</span>: <span class="t-g">"tie:9a1f…"</span> },
    <span class="t-b">"children"</span>: [
      { <span class="t-b">"type"</span>: <span class="t-g">"hotspot"</span>, <span class="t-b">"name"</span>: <span class="t-g">"Door"</span>,
        <span class="t-b">"props"</span>: { <span class="t-b">"label"</span>: <span class="t-g">"Kitchen"</span> },
        <span class="t-b">"on"</span>: { <span class="t-b">"tapped"</span>: [{ <span class="t-b">"do"</span>: <span class="t-g">"navigate"</span>, <span class="t-b">"to"</span>: <span class="t-g">"../kitchen"</span> }] } },
      { <span class="t-b">"type"</span>: <span class="t-g">"tie.gallery"</span>,
        <span class="t-b">"props"</span>: { <span class="t-b">"tags"</span>: [<span class="t-g">"family"</span>, <span class="t-g">"2014"</span>] } }
    ]
  }
}</code></pre>
  </div>
</div>

## The components

| Component | What it renders |
|---|---|
| `tie.image` | One image by hash or path. |
| `tie.gallery` | A live query: images tagged with all of these tags. |
| `tie.collection` | An album, series or picture set. |
| `tie.tags` | A tag picker that reflects the library. |
| `tie.related` | The co-tags of the current selection. |
| `tie.document` | A PDF or document from the blob store. |
| `tie.audio` | Playback through pwplay, with the page's own transport. |
| `tie.video` | Video by hash, with a poster frame. |

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Designer</span>
    <h2 class="t-h2">Edit the page in place.</h2>
    <p class="t-sub">
      <span class="t-kbd">Ctrl</span> + <span class="t-kbd">E</span> turns the page
      into a designer: move boxes, set props, pick media from the library, and
      save. A persistent indicator tells you whether the scene is saved, saving,
      or failed to save. <span class="t-kbd">Ctrl</span> + <span class="t-kbd">S</span>
      saves a new version.
    </p>
    <ul class="t-list">
      <li><strong>Media and scene pickers</strong> browse the real library.</li>
      <li><strong>A story map</strong> shows how scenes link to each other.</li>
      <li><strong>Auto-generated thumbnails</strong> render offscreen, so the library grid has a preview.</li>
    </ul>
  </div>
  <div>
    <span class="t-eyebrow">Script</span>
    <h3 class="t-h3" style="margin-top: 0.9rem">A page can have behaviour.</h3>
    <p class="t-sub" style="font-size: 0.95rem">
      Add a <code>page.risor</code> sidecar when a hotspot needs more than
      navigate/open/play. It is ordinary Risor running in the fynerisor runtime,
      with the same theme and the same window — edit it in place from the menu.
    </p>
    <div class="t-badge t-badge--warn" style="margin-top: 1rem">Preview</div>
    <p class="t-sub" style="font-size: 0.9rem; margin-top: 0.6rem">
      tie-browser is the newest app in the family. The scene format can still
      change between releases, and the roadmap includes native video surfaces and
      richer tag pickers. Everything else on this site is built to be stable
      enough to rely on; this one is built to be explored.
    </p>
  </div>
</div>
