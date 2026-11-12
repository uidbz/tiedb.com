// Package logo generates the tie logo family: one heavy-stroke mark per app,
// drawn on the same rounded tile as anneal's original in-app logo.
package logo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Spec struct {
	Tile   Tile    `yaml:"tile"`
	Weight float64 `yaml:"weight"`
	Caps   string  `yaml:"caps"`
	Joins  string  `yaml:"joins"`
	Dot    Dot     `yaml:"dot"`
	Marks  []Mark  `yaml:"marks"`
}

type Tile struct {
	Color  string  `yaml:"color"`
	Radius float64 `yaml:"radius"`
}

type Dot struct {
	Radius float64 `yaml:"radius"`
	Color  string  `yaml:"color"`
}

type Mark struct {
	ID     string    `yaml:"id"`
	Name   string    `yaml:"name"`
	Role   string    `yaml:"role"`
	Hue    string    `yaml:"hue"`
	Light  string    `yaml:"light"`
	Paths  []string  `yaml:"paths"`
	Dot    []float64 `yaml:"dot"`
	Weight *float64  `yaml:"weight"`
	Caps   string    `yaml:"caps"`
	Joins  string    `yaml:"joins"`
}

func LoadSpec(path string) (*Spec, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Spec
	if err := yaml.Unmarshal(b, &s); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if s.Weight == 0 {
		s.Weight = 49.2
	}
	if s.Caps == "" {
		s.Caps = "round"
	}
	if s.Joins == "" {
		s.Joins = "round"
	}
	if s.Dot.Radius == 0 {
		s.Dot.Radius = 7
	}
	if s.Dot.Color == "" {
		s.Dot.Color = "#f5f5f5"
	}
	if s.Tile.Color == "" {
		s.Tile.Color = "#181819"
	}
	if s.Tile.Radius == 0 {
		s.Tile.Radius = 28
	}
	for i := range s.Marks {
		if s.Marks[i].Weight == nil {
			w := s.Weight
			s.Marks[i].Weight = &w
		}
		if s.Marks[i].Caps == "" {
			s.Marks[i].Caps = s.Caps
		}
		if s.Marks[i].Joins == "" {
			s.Marks[i].Joins = s.Joins
		}
	}
	return &s, nil
}

func (s *Spec) Mark(id string) (Mark, bool) {
	for _, m := range s.Marks {
		if m.ID == id {
			return m, true
		}
	}
	return Mark{}, false
}

// SVG renders the mark. withTile controls whether the dark rounded tile is
// included. The geometry is drawn with the heavy stroke directly, so browsers
// and every sane rasteriser show the same fat silhouette at any size.
func (s *Spec) SVG(m Mark, withTile bool) string {
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128" role="img" aria-label="%s">`, xmlEscape(m.Name))
	if withTile {
		fmt.Fprintf(&b, `<rect width="128" height="128" rx="%s" fill="%s"/>`, num(s.Tile.Radius), s.Tile.Color)
	}
	fmt.Fprintf(&b, `<g fill="none" stroke="%s" stroke-width="%s" stroke-linecap="%s" stroke-linejoin="%s">`,
		m.Hue, num(*m.Weight), m.Caps, m.Joins)
	for _, p := range m.Paths {
		fmt.Fprintf(&b, `<path d="%s"/>`, p)
	}
	b.WriteString(`</g>`)
	if len(m.Dot) == 2 {
		fmt.Fprintf(&b, `<circle cx="%s" cy="%s" r="%s" fill="%s"/>`, num(m.Dot[0]), num(m.Dot[1]), num(s.Dot.Radius), s.Dot.Color)
	}
	b.WriteString(`</svg>`)
	return b.String()
}

// IconSVG renders a compact inline icon for Hextra's data/icons.yaml.
func (s *Spec) IconSVG(m Mark) string {
	var b strings.Builder
	b.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128" fill="none">`)
	fmt.Fprintf(&b, `<rect width="128" height="128" rx="%s" fill="%s"/>`, num(s.Tile.Radius), s.Tile.Color)
	fmt.Fprintf(&b, `<g stroke="%s" stroke-width="%s" stroke-linecap="%s" stroke-linejoin="%s">`,
		m.Hue, num(*m.Weight), m.Caps, m.Joins)
	for _, p := range m.Paths {
		fmt.Fprintf(&b, `<path d="%s"/>`, p)
	}
	b.WriteString(`</g>`)
	if len(m.Dot) == 2 {
		fmt.Fprintf(&b, `<circle cx="%s" cy="%s" r="%s" fill="%s"/>`, num(m.Dot[0]), num(m.Dot[1]), num(s.Dot.Radius), s.Dot.Color)
	}
	b.WriteString(`</svg>`)
	return b.String()
}

// Generate writes every logo asset. root is the repository root.
func (s *Spec) Generate(root string) error {
	logos := filepath.Join(root, "assets", "logos")
	if err := os.MkdirAll(logos, 0o755); err != nil {
		return err
	}
	if _, err := exec.LookPath("rsvg-convert"); err != nil {
		return fmt.Errorf("rsvg-convert is required to rasterise logos: %w", err)
	}
	for _, m := range s.Marks {
		if err := os.WriteFile(filepath.Join(logos, m.ID+".svg"), []byte(s.SVG(m, true)), 0o644); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(logos, m.ID+"-glyph.svg"), []byte(s.SVG(m, false)), 0o644); err != nil {
			return err
		}
	}
	if err := s.generateFavicons(root); err != nil {
		return err
	}
	if err := s.generateIconsData(root); err != nil {
		return err
	}
	if err := s.generateBrand(root); err != nil {
		return err
	}
	if err := s.generateOpenGraph(root); err != nil {
		return err
	}
	if err := s.generateSheet(root); err != nil {
		return err
	}
	return nil
}

func (s *Spec) umbrella() Mark {
	m, _ := s.Mark("tie")
	return m
}

func (s *Spec) generateFavicons(root string) error {
	static := filepath.Join(root, "static")
	if err := os.MkdirAll(static, 0o755); err != nil {
		return err
	}
	u := s.umbrella()
	if err := os.WriteFile(filepath.Join(static, "favicon.svg"), []byte(s.SVG(u, true)), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(static, "favicon-dark.svg"), []byte(s.SVG(u, true)), 0o644); err != nil {
		return err
	}
	sizes := map[int]string{
		16:  "favicon-16x16.png",
		32:  "favicon-32x32.png",
		64:  "favicon.png",
		180: "apple-touch-icon.png",
		192: "android-chrome-192x192.png",
		512: "android-chrome-512x512.png",
	}
	var icoPNGs [][]byte
	for _, size := range []int{16, 32, 48} {
		out := filepath.Join(static, fmt.Sprintf(".tmp-ico-%d.png", size))
		if err := renderSVG(s.SVG(u, true), out, size); err != nil {
			return err
		}
		b, err := os.ReadFile(out)
		if err != nil {
			return err
		}
		icoPNGs = append(icoPNGs, b)
		os.Remove(out)
	}
	for size, name := range sizes {
		out := filepath.Join(static, fmt.Sprintf(".tmp-%d.png", size))
		if err := renderSVG(s.SVG(u, true), out, size); err != nil {
			return err
		}
		if err := os.Rename(out, filepath.Join(static, name)); err != nil {
			return err
		}
	}
	if err := writeICO(filepath.Join(static, "favicon.ico"), icoPNGs, []int{16, 32, 48}); err != nil {
		return err
	}
	manifest := `{
  "name": "tie",
  "short_name": "tie",
  "description": "A local-first content-addressed library, its native apps, and the anneal coding agent.",
  "start_url": "/",
  "display": "standalone",
  "background_color": "#17181b",
  "theme_color": "#17181b",
  "icons": [
    { "src": "/android-chrome-192x192.png", "sizes": "192x192", "type": "image/png" },
    { "src": "/android-chrome-512x512.png", "sizes": "512x512", "type": "image/png" },
    { "src": "/android-chrome-512x512.png", "sizes": "512x512", "type": "image/png", "purpose": "maskable" }
  ]
}
`
	return os.WriteFile(filepath.Join(static, "site.webmanifest"), []byte(manifest), 0o644)
}

func (s *Spec) generateIconsData(root string) error {
	data := filepath.Join(root, "data")
	if err := os.MkdirAll(data, 0o755); err != nil {
		return err
	}
	var b strings.Builder
	b.WriteString("# Generated by `make logos` from tools/specs/marks.yaml. Do not edit.\n")
	for _, m := range s.Marks {
		fmt.Fprintf(&b, "mark-%s: |\n", m.ID)
		for _, line := range strings.Split(s.IconSVG(m), "\n") {
			fmt.Fprintf(&b, "  %s\n", line)
		}
	}
	return os.WriteFile(filepath.Join(data, "icons.yaml"), []byte(b.String()), 0o644)
}

func (s *Spec) generateBrand(root string) error {
	brand := filepath.Join(root, "static", "brand")
	if err := os.MkdirAll(brand, 0o755); err != nil {
		return err
	}
	for _, m := range s.Marks {
		svg := []byte(s.SVG(m, true))
		if err := os.WriteFile(filepath.Join(brand, m.ID+".svg"), svg, 0o644); err != nil {
			return err
		}
		for _, size := range []int{128, 256, 512, 1024} {
			if err := renderSVG(string(svg), filepath.Join(brand, fmt.Sprintf("%s-%d.png", m.ID, size)), size); err != nil {
				return err
			}
		}
	}
	return nil
}

// generateOpenGraph builds 1200x630 social-card bases. Hugo overlays the page
// title with images.Text at build time.
func (s *Spec) generateOpenGraph(root string) error {
	dir := filepath.Join(root, "assets", "og")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, m := range s.Marks {
		tmp, err := os.CreateTemp("", "mark-*.png")
		if err != nil {
			return err
		}
		tmp.Close()
		defer os.Remove(tmp.Name())
		if err := renderSVG(s.SVG(m, true), tmp.Name(), 320); err != nil {
			return err
		}
		mark, err := readPNG(tmp.Name())
		if err != nil {
			return err
		}
		img := cardBase(hexColor(m.Hue), mark)
		if err := writePNG(filepath.Join(dir, m.ID+".png"), img); err != nil {
			return err
		}
	}
	return nil
}

func cardBase(hue color.RGBA, mark image.Image) image.Image {
	const w, h = 1200, 630
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	bg := hexColor("#17181b")
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// A soft glow in the top-right corner, in the app's hue.
			dx := float64(x-980) / 620
			dy := float64(y-40) / 520
			g := math.Exp(-(dx*dx + dy*dy))
			a := g * 0.30
			// A faint teal counter-glow bottom-left.
			dx2 := float64(x-120) / 520
			dy2 := float64(y-620) / 420
			a2 := math.Exp(-(dx2*dx2+dy2*dy2)) * 0.10
			r := mix(bg.R, hue.R, a)
			gg := mix(bg.G, hue.G, a)
			b := mix(bg.B, hue.B, a)
			r = mix(r, 45, a2)
			gg = mix(gg, 212, a2)
			b = mix(b, 191, a2)
			dst.SetRGBA(x, y, color.RGBA{r, gg, b, 255})
		}
	}
	// A hairline frame and a bottom accent bar.
	line := color.RGBA{255, 255, 255, 18}
	for x := 60; x < w-60; x++ {
		dst.SetRGBA(x, 60, line)
	}
	accent := color.RGBA{hue.R, hue.G, hue.B, 255}
	for x := 80; x < 220; x++ {
		for y := 520; y < 524; y++ {
			dst.SetRGBA(x, y, accent)
		}
	}
	draw.Draw(dst, image.Rect(80, 155, 80+mark.Bounds().Dx(), 155+mark.Bounds().Dy()), mark, image.Point{}, draw.Over)
	return dst
}

func (s *Spec) generateSheet(root string) error {
	dir := filepath.Join(root, ".cache", "logos")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	var b strings.Builder
	b.WriteString(`<!doctype html><meta charset="utf-8"><title>tie logo family</title>
<style>
  body{margin:0;padding:48px;background:#0d0e10;color:#e6e7ea;font:15px/1.5 system-ui,sans-serif}
  h1{font-size:26px;letter-spacing:-.02em;margin:0 0 4px}
  p{color:#9a9ca3;margin:0 0 36px}
  table{border-collapse:collapse;width:100%}
  td,th{padding:16px;border-top:1px solid #2e3036;text-align:left;vertical-align:middle}
  th{font-size:11px;text-transform:uppercase;letter-spacing:.14em;color:#9a9ca3;font-weight:600}
  img{display:block;image-rendering:auto}
  .row{display:flex;gap:18px;align-items:center}
  .dark{background:#17181b;padding:14px;border-radius:10px}
  .light{background:#eef0f2;padding:14px;border-radius:10px}
  .hue{width:54px;height:54px;border-radius:12px;display:inline-block;vertical-align:middle;margin-right:10px}
  .name{font-weight:700;font-size:16px}
  .role{color:#9a9ca3;font-size:12px}
  .note{margin-top:36px;color:#9a9ca3;font-size:13px}
</style>
<h1>tie logo family</h1>
<p>Heavy-stroke marks on the anneal tile. Sizes 16, 26, 48, 128 and 512 px, on dark and light.</p>
<table><thead><tr><th>mark</th><th>dark</th><th>light</th><th>hue</th></tr></thead><tbody>
`)
	for _, m := range s.Marks {
		svg := "../assets/logos/" + m.ID + ".svg"
		fmt.Fprintf(&b, `<tr><td><div class="name">%s</div><div class="role">%s</div></td>`, m.Name, m.Role)
		b.WriteString(`<td><div class="row dark">`)
		for _, size := range []int{16, 26, 48, 128} {
			fmt.Fprintf(&b, `<img src="%s" width="%d" height="%d" alt="">`, svg, size, size)
		}
		b.WriteString(`</div></td><td><div class="row light">`)
		for _, size := range []int{16, 26, 48, 128} {
			fmt.Fprintf(&b, `<img src="%s" width="%d" height="%d" alt="">`, svg, size, size)
		}
		fmt.Fprintf(&b, `</div></td><td><span class="hue" style="background:%s"></span><code>%s</code><div class="role">%s</div></td></tr>`, m.Hue, m.Hue, m.Light)
	}
	b.WriteString(`</tbody></table><p class="note">Generated by <code>make logos</code>. Edit tools/specs/marks.yaml and re-run.</p>`)
	return os.WriteFile(filepath.Join(dir, "index.html"), []byte(b.String()), 0o644)
}

func renderSVG(svg, out string, size int) error {
	f, err := os.CreateTemp("", "mark-*.svg")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err := f.WriteString(svg); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	cmd := exec.Command("rsvg-convert", "-w", fmt.Sprint(size), "-h", fmt.Sprint(size), "-o", out, f.Name())
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("rsvg-convert: %w: %s", err, stderr.String())
	}
	return nil
}

func writeICO(path string, pngs [][]byte, sizes []int) error {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uint16(0))
	binary.Write(&buf, binary.LittleEndian, uint16(1))
	binary.Write(&buf, binary.LittleEndian, uint16(len(pngs)))
	offset := 6 + 16*len(pngs)
	for i, p := range pngs {
		w := byte(sizes[i])
		if sizes[i] >= 256 {
			w = 0
		}
		buf.WriteByte(w)
		buf.WriteByte(w)
		buf.WriteByte(0)
		buf.WriteByte(0)
		binary.Write(&buf, binary.LittleEndian, uint16(1))
		binary.Write(&buf, binary.LittleEndian, uint16(32))
		binary.Write(&buf, binary.LittleEndian, uint32(len(p)))
		binary.Write(&buf, binary.LittleEndian, uint32(offset))
		offset += len(p)
	}
	for _, p := range pngs {
		buf.Write(p)
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

func readPNG(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func writePNG(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func hexColor(s string) color.RGBA {
	s = strings.TrimPrefix(s, "#")
	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}
	var r, g, b uint8
	fmt.Sscanf(s, "%02x%02x%02x", &r, &g, &b)
	return color.RGBA{r, g, b, 255}
}

func mix(a, b uint8, t float64) uint8 {
	if t <= 0 {
		return a
	}
	if t >= 1 {
		return b
	}
	return uint8(float64(a)*(1-t) + float64(b)*t)
}

func num(f float64) string {
	s := fmt.Sprintf("%.2f", f)
	s = strings.TrimRight(s, "0")
	return strings.TrimRight(s, ".")
}

func xmlEscape(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&quot;")
	return r.Replace(s)
}

// OrderedIDs returns the mark IDs in spec order.
func (s *Spec) OrderedIDs() []string {
	ids := make([]string, 0, len(s.Marks))
	for _, m := range s.Marks {
		ids = append(ids, m.ID)
	}
	sort.Strings(ids)
	return ids
}
