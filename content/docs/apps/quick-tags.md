---
title: Quick tags
type: docs
weight: 21
---

Quick tags are the pills over the edge of the image in tie-view. They are
configured in `quicktags.toml` (Settings → Quick tags edits the same file).

```toml
# ~/.config/tieview/quicktags.toml
[general]
enabled = true
size = "md"          # xs | sm | md | lg | xl
position = "bottom"  # top | bottom | left | right

[[tags]]
tag = "favorite"
icon = "heart"
key = "f"

[[tags]]
tag = "family"
key = "1"

[[tags]]
tag = "wallpaper"
label = "Wall"
key = "2"
```

- **Include/exclude**: a quick tag toggles; the state comes from the file's tags.
- **Icons**: a heart, a star, or any of the built-in icon names; without an icon
  the label is used.
- **Keys**: one key per tag, plus the rating keys for 1–5 stars. The keys are
  global while the image view is focused.
- **Sizes**: xs–xl, scaled for touch on Android.

The pill layout is stored per collection profile, so a phone and a desktop can
use different sizes. The defaults are sensible: a favourite heart, five stars and
a few common tags.

The source is `cmd/tie-view/quicktagconfig.go`; the config is documented by
example in `quicktags.toml` in the repo.
