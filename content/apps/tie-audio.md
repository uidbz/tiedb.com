---
title: tie-audio
type: app
app: tie-audio
hero_shot:
  scene: wall
  meta: "cover wall · 1440×900"
  alt: "tie-audio's cover wall: a grid of album artwork with the tag sidebar on the left and the transport bar at the bottom"
  caption: "The cover wall, with the tag sidebar, the Covers/Table/Now Playing tabs and the transport pinned to the bottom."
---

<div class="t-feature t-feature--top">
  <div>
    <span class="t-eyebrow">Home</span>
    <h2 class="t-h2">A home that remembers.</h2>
    <p class="t-sub">
      Jump back in, recently played, recently added, because you played something
      else, haven't heard in a while, on this day. The shelves are queries over the
      play history that tie-audio records in tie, not a separate recommendation
      service.
    </p>
    <ul class="t-list">
      <li><strong>Play history is data.</strong> Each play is an idempotent event with a count and a timestamp.</li>
      <li><strong>On this day</strong> is a date query, not a playlist you maintain by hand.</li>
      <li><strong>Favourites</strong> come from the same tags tie-view uses.</li>
    </ul>
  </div>
  <div>
    {{< shot app="tie-audio" scene="compact" class="t-shot--phone" meta="compact · 420×860" title="tie-audio" alt="tie-audio's compact phone home: shelves of album covers above a mini transport bar and a bottom navigation bar" caption="The compact layout is the same app: shelves, a mini player and a bottom nav." >}}
  </div>
</div>

<div class="t-feature t-feature--flip">
  <div>
    <span class="t-eyebrow">Albums</span>
    <h2 class="t-h2">The wall, the table, the queue.</h2>
    <p class="t-sub">
      Covers for browsing, a sortable table for finding, and a queue that is
      visible while you are anywhere else. Choose columns, reorder, shuffle, save
      the queue as a playlist — then find it under Playlists tomorrow.
    </p>
    <ul class="t-list">
      <li><strong>Gapless local playback</strong> over PipeWire, or a remote <code>pwplay-server</code>.</li>
      <li><strong>CUE sheets and zipped albums</strong> are first-class albums, not a folder of tracks.</li>
      <li><strong>Search</strong> covers albums, artists and tracks with the same field selector as the rest of the app.</li>
    </ul>
  </div>
  <div>
    {{< shot app="tie-audio" scene="album" meta="album · 1440×900" title="tie-audio" alt="An album page in tie-audio: title, artist, year, Play album and Add to playlist buttons, and a track table with duration" caption="An album is a page, with Play album, Add to playlist, and a track table you can sort." >}}
  </div>
</div>

<div class="t-feature">
  <div>
    <span class="t-eyebrow">Everywhere</span>
    <h2 class="t-h2">Media keys, lock screens, no surprises.</h2>
    <p class="t-sub">
      MPRIS 2 on Linux means your keyboard's play button and <code>playerctl</code>
      just work. On Android, playback sits in the media session with lock-screen
      controls and audio focus. The history recorder is the only thing that writes
      back, and it writes to tie.
    </p>
    <ul class="t-list">
      <li><strong>Flac, MP3, WAV, OGG, Opus</strong> — decoded by your system, not bundled codecs.</li>
      <li><strong>Remote control</strong> via <code>tie-audio -control play-pause</code> for a script or a key binding.</li>
      <li><strong>No account.</strong> No scrobbling, no cloud library, no listening profile.</li>
    </ul>
  </div>
  <div class="t-term">
    <div class="t-term-bar">
      <img src="/brand/tie-audio.svg" alt="" />
      <span>hotkeys</span>
    </div>
    <pre class="t-term-body"><code><span class="t-o">desktop · rebindable in [Hotkeys]</span>
<span class="t-a">Space</span>  play / pause        <span class="t-a">N</span>  next
<span class="t-a">X</span>      stop                <span class="t-a">P</span>  previous
<span class="t-a">→ ←</span>    seek ±10s           <span class="t-a">= -</span> volume
<span class="t-a">Esc</span>    back

<span class="t-o">android · media session + lock screen</span>
<span class="t-a">swipe</span>  between home, albums, tags, playlist</code></pre>
  </div>
</div>

## Under the hood

tie-audio is a tag-aware player, not a music database with a UI. Albums come from
`audio-dir` imports in tie; covers are resolved through the media model; play
events are recorded as `play-event` subjects. If you already have the files, you
already have the library — tie-audio just gives it a front end.

| Piece | What it does |
|---|---|
| `audio-dir` | A directory that is an album: tracks, disc numbers, cover art. |
| `member-of` | The album contract that groups tracks without moving them. |
| `play-event` | An idempotent play with `play-count` and `last-played`. |
| `playlist` | An ordered `playlist-track` list, stored in the same collection. |
| pwplay | The playback engine, in-process or over HTTP. |

The **manual** covers the connection setup, the two playback backends, the
config file and the Android build. The **shortcut reference** is generated from
the app's defaults by the site tooling.
