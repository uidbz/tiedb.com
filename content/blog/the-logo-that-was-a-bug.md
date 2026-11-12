---
title: "The logo that was a bug"
type: blog
date: 2026-11-10
description: "Why anneal's logo looks like a solid amber blob, and how a rendering bug became a family of marks."
---

anneal's logo was drawn as a thin amber outline of an **A** on a dark rounded
tile: two legs, a crossbar and a small white dot on the apex. In the app it looks
like a fat, solid amber shape with a pinhole. That is not what the SVG says.

The cause is a rendering bug. Fyne rasterises SVG with oksvg. In
`oksvg/svg_path.go`, `DrawTransformed` applies the transform matrix to the path
points, but calls `SetStroke` with the raw `LineWidth`. The stroke is therefore
always **10 device pixels**, regardless of the scale.

In anneal's 26 px activity rail, a 10 px stroke is 38% of the icon. On the
128-unit viewBox that is a stroke width of 49.2, not 10. The thin outline fills
in: the crossbar disappears, the upper counter closes, a small notch remains
between the legs, and the apex dot becomes a white pinhole. The amber nearly
reaches the tile edges.

It looks good. It is also fragile: at a different display scale the stroke would
be proportionally thinner, and if the bug is ever fixed the logo turns back into
an outline. So the site does not depend on the bug. The marks are drawn with the
heavy geometry directly, as strokes with `stroke-width="49.2"` (rendered
correctly by every browser), and the generator includes a fidelity check that
compares the result with the app's actual 26 px render.

That gave us a family rule:

- a 128-unit tile, `#181819`, radius 28;
- one heavy stroke, round caps and joins;
- one white pinhole on the shape;
- a silhouette that survives being filled in.

The marks are the **A** for anneal, a bow tie for tie, a folder for tie-fm,
headphones for tie-audio, a star for tie-view, two peaks for imgview and a map
pin for tie-browser. The palette is anneal's: teal, amber, and a small set of
saturated hues for the apps.

The [brand page](/brand/) has the downloads and the full family. The generator
is in `tools/`; the contact sheet is one `make logos` away.
