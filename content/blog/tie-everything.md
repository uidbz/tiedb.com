---
title: "Tie it all together"
type: blog
date: 2026-11-12
description: "From a 2023 key-value-value page to a library, five native apps and a coding agent. What this site is for."
---

tiedb.com used to be a single page about a key-value-value store and a recipe
example. The project is now three things that belong together:

- **tie**, a content-addressed library where tags are a filesystem;
- **the tie apps**, five native tools that use it;
- **anneal**, a coding agent that stores its session memory in it.

This site is the front door. It is built with Hugo and Hextra, it has no
analytics, no CDNs and no web fonts, and its screenshots are generated from the
real applications by scripts in the repository. If a page and the app disagree,
the app is right and the page is a bug.

The plan is simple: a landing page that explains the idea, a page per app that
shows the actual UI, documentation that is useful, and stories that say what was
measured and what was reverted. There is a brand page with the logo family, a
glossary, and a 404 page that is a tag query.

The old page is not lost. The pasta-and-basil example still explains a triple
better than any diagram, so it lives on in
[Triples 101](/docs/tie/concepts/#triples-101). The name moved on; the idea
didn't.
