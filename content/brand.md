---
title: Brand
type: brand
description: "The tie logo family, the palette and the rules for using them."
---

The family began with anneal's in-app logo: a heavy amber mark on a dark rounded
tile. The site reproduces that geometry directly, so it looks the same at every
size and does not depend on a rendering bug. The full story is in
[The logo that was a bug](/blog/the-logo-that-was-a-bug/).

## Marks

{{< brand-marks >}}

## Palette

{{< brand-palette >}}

| Surface | Dark | Light |
|---|---|---|
| Background | `#17181b` | `#ffffff` |
| Foreground | `#e6e7ea` | `#1f2328` |
| Rail | `#111214` | `#eef0f2` |
| Raised | `#232429` | `#f3f4f6` |
| Border | `#2e3036` | `#e2e4e8` |
| Accent (tie) | `#2dd4bf` | `#0f766e` |
| OK / warn / error | `#4ade80` / `#f59e0b` / `#ef4444` | `#15803d` / `#b45309` / `#b91c1c` |
| anneal amber | `#e6a23c` | `#b45309` |

Radii: 6 for controls, 8 for cards, 999 for pills. Spacing: 4, 8, 12, 16, 24.
The contrast floor is the same one the apps enforce: foreground ≥ 7:1, muted ≥
4.5:1, accent ≥ 3:1.

## Rules

- **Minimum size:** 16 px for the tile mark, 26 px in an activity rail.
- **Clear space:** one eighth of the mark's width on all sides.
- **Do not** stretch, rotate (except the hero's decorative tilt), recolour the
  tile, or place the mark on a background that competes with its hue.
- **Wordmark:** the mark plus the lowercase name in the site's sans. The name is
  set in text, not in the SVG, so it stays sharp and translatable.
- **One colour:** `tie` is teal, `anneal` amber, `tie-fm` sky, `tie-audio` rose,
  `tie-view` violet, `imgview` lime, `tie-browser` fuchsia.

## Typography

The site uses the system sans stack with Noto Sans first, and
`ui-monospace, "DejaVu Sans Mono", monospace` for code. No web fonts are
downloaded. The app screenshots use Fyne's bundled Noto Sans, which is why the
site and the screenshots look related.

## Badges

<div class="t-row" style="margin: 1rem 0 1.5rem">
  <span class="t-badge t-badge--accent">Local</span>
  <span class="t-badge t-badge--muted">Stable</span>
  <span class="t-badge t-badge--warn">Preview</span>
  <span class="t-chip"><i class="t-dot"></i> online</span>
  <span class="t-chip t-chip--plain">BSD-3</span>
</div>

The badge shapes and the teal/amber washes come from anneal's approval card and
the apps' status chips. Use `Preview` for anything whose format can change, and
`Stable` only for the apps that are used every day.
