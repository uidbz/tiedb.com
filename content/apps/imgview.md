---
title: imgview
type: app
app: imgview
hero_shot:
  scene: gallery
  meta: "gallery · 1280×800"
  alt: "imgview's gallery: a justified grid of public-domain paintings and NASA photographs, with file names underneath each tile"
  caption: "A folder of mixed artwork and photographs, at 2x. No library, no import, no server."
---

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Zero setup</span>
    <h2 class="t-h2">Point it at a folder.</h2>
    <p class="t-sub">
      imgview is the app you open before you have decided to organise anything.
      Give it a directory, a single image, a video or an archive. It reads the
      folder, builds a gallery and gets out of the way.
    </p>
    <ul class="t-list">
      <li><strong>Justified rows</strong> that fill the window, paginated at 500 images and cached on disk.</li>
      <li><strong>Folder tiles</strong> show a preview and swipe through up to ten images inside.</li>
      <li><strong>No database.</strong> Nothing to configure, nothing to back up, nothing to migrate.</li>
    </ul>
  </div>
  <div>
    {{< shot app="imgview" scene="image" meta="single image · 1280×800" title="imgview" alt="A single photograph filling imgview's window, letterboxed against the dark background" caption="The single-image view: fill the window, or show it at original size." >}}
  </div>
</div>

<div class="t-feature t-feature--flip">
  <div>
    <span class="t-eyebrow">Information</span>
    <h2 class="t-h2">The details, on demand.</h2>
    <p class="t-sub">
      Press <span class="t-kbd">I</span> for the image information overlay: camera,
      exposure, dimensions and the file itself. It loads asynchronously, so large
      files do not block the next image.
    </p>
    <ul class="t-list">
      <li><strong>EXIF and file info</strong> without leaving the viewer.</li>
      <li><strong>Zoom, pan, rotate</strong> and fullscreen, all from the keyboard or the mouse.</li>
      <li><strong>Rebindable keys</strong> for everything, including one key that runs your own command on the file.</li>
    </ul>
  </div>
  <div>
    {{< shot app="imgview" scene="info" meta="info overlay · 1280×800" title="imgview" alt="The image information overlay in imgview, showing file metadata over a photograph" caption="The info overlay: EXIF, dimensions and the path, loaded without blocking the image." >}}
  </div>
</div>

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Files</span>
    <h2 class="t-h2">Archives open like folders.</h2>
    <p class="t-sub">
      zip, rar, tar, 7z, cbr and cbz are all just sources of images. Open one and
      imgview treats it as a directory, with the archive's own cover as the tile
      preview. Video is played by libmpv, with frame thumbnails on the tile and a
      strip of frames for a quick scan. Build with <code>-tags nompv</code> and the
      video code disappears.
    </p>
    <ul class="t-list">
      <li><strong>Comic readers</strong>: cbz/cbr pages in reading order.</li>
      <li><strong>Android</strong> uses the system folder picker; the same gallery, a touch-first layout.</li>
      <li><strong>One binary</strong>, no service, no tie connection required.</li>
    </ul>
  </div>
  <div class="t-term">
    <div class="t-term-bar">
      <img src="/brand/imgview.svg" alt="" />
      <span>keys · defaults</span>
    </div>
    <pre class="t-term-body"><code><span class="t-a">J / ↓</span>  scroll down        <span class="t-a">K / ↑</span>  scroll up
<span class="t-a">← / →</span>  previous / next    <span class="t-a">S</span>      original size
<span class="t-a">H / L</span>  rotate left / right <span class="t-a">X</span>      fit window
<span class="t-a">N</span>      filenames          <span class="t-a">I</span>      image info
<span class="t-a">T</span>      tag bar            <span class="t-a">E</span>      run your command
<span class="t-a">F</span>      fullscreen         <span class="t-a">Esc</span>    back / quit</code></pre>
  </div>
</div>

## The gallery engine

imgview and tie-view share the same `gallery/` package: the justified layout,
the thumbnail pipeline, pagination, the tile updater, the info overlay and the
drawer. tie-fm uses it for its preview pane, and tie-audio for the cover wall.
That is why a fix to scrolling or HiDPI tiles lands in every app at once.
