# Working on tiedb.com

- The site is Hugo + Hextra. `make dev` for a preview, `make build` for a clean
  production build.
- Do not hand-edit generated files: `data/facts.yaml`, `data/icons.yaml`,
  `assets/logos/`, `assets/og/`, `assets/shots/`, `content/docs/reference/`.
  Regenerate with `make facts`, `make logos`, `make shots`, `make docs`.
- Version numbers on pages come from `hugo.Data.facts` or the `fact` shortcode.
  Do not type a new one into prose.
- Screenshots are synthetic. Never commit a capture of a real library or a
  path outside `/home/you`. `make check` enforces the obvious cases.
- External links must not use `target="_blank"`; the hosting CSP sandbox blocks
  it and `make check` fails. The site's render-link hook already omits it.
- No CDNs, web fonts or analytics. KaTeX/Mermaid/asciinema are off for the same
  reason.
- New pages: use the `showcase`, `app`, `brand` or `default` type as the closest
  match, and add a description for social cards.
- A new app goes in `data/apps.yaml`; the mark goes in `tools/specs/marks.yaml`
  and is regenerated.
