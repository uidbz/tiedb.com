// Package docsync copies long-form source notes into the site's reference
// section, adding front matter and rewriting relative links to GitHub.
package docsync

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type Entry struct {
	Repo   string `yaml:"repo"`
	Ref    string `yaml:"ref"`
	Src    string `yaml:"src"`
	Dst    string `yaml:"dst"`
	Title  string `yaml:"title"`
	Weight int    `yaml:"weight"`
}

var (
	h1RE      = regexp.MustCompile(`(?m)^#\s+(.+?)\s*$`)
	linkRE    = regexp.MustCompile(`\]\(([^)\s]+)(\s+"[^"]*")?\)`)
)

func LoadManifest(path string) ([]Entry, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var entries []Entry
	if err := yaml.Unmarshal(b, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

// Sync writes every entry under root/content. In check mode it compares
// instead, returning the list of pages that are out of date.
func Sync(root, srcRoot, manifest string, check bool) ([]string, error) {
	entries, err := LoadManifest(manifest)
	if err != nil {
		return nil, err
	}
	var stale []string
	for _, e := range entries {
		body, err := render(srcRoot, e)
		if err != nil {
			return stale, fmt.Errorf("%s/%s: %w", e.Repo, e.Src, err)
		}
		dst := filepath.Join(root, "content", e.Dst)
		old, _ := os.ReadFile(dst)
		if check {
			if string(old) != body {
				stale = append(stale, e.Dst)
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return stale, err
		}
		if err := os.WriteFile(dst, []byte(body), 0o644); err != nil {
			return stale, err
		}
	}
	return stale, nil
}

func render(srcRoot string, e Entry) (string, error) {
	repoDir := filepath.Join(srcRoot, e.Repo)
	path := filepath.Join(repoDir, filepath.FromSlash(e.Src))
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := stripFrontMatter(string(raw))
	title := e.Title
	if m := h1RE.FindStringSubmatch(text); len(m) == 2 {
		if title == "" {
			title = strings.TrimSpace(m[1])
		}
		text = h1RE.ReplaceAllString(text, "")
	}
	if title == "" {
		title = strings.TrimSuffix(filepath.Base(e.Src), filepath.Ext(e.Src))
	}
	text = rewriteLinks(text, e)
	text = strings.NewReplacer("/home/johan", "/home/you", "/tmp/opencode", "/tmp", "/root/.cache", "/home/you/.cache").Replace(text)
	text = strings.TrimSpace(text)

	last := git(repoDir, "log", "-1", "--format=%cs", "--", e.Src)
	repoURL := repoHTTPS(repoDir)
	blob := fmt.Sprintf("%s/blob/%s/%s", repoURL, e.Ref, e.Src)

	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: %q\n", title)
	b.WriteString("type: docs\n")
	if e.Weight != 0 {
		fmt.Fprintf(&b, "weight: %d\n", e.Weight)
	}
	if last != "" {
		fmt.Fprintf(&b, "lastmod: %q\n", last)
	}
	b.WriteString("generated: true\n")
	b.WriteString("---\n\n")
	b.WriteString("> [!NOTE]\n")
	fmt.Fprintf(&b, "> Synced from [`%s/%s`](%s) on %s. Edit the source, not this copy.\n\n", e.Repo, e.Src, blob, last)
	b.WriteString(text)
	b.WriteString("\n")
	return b.String(), nil
}

func stripFrontMatter(s string) string {
	s = strings.TrimPrefix(s, "\ufeff")
	if !strings.HasPrefix(s, "---\n") {
		return s
	}
	if i := strings.Index(s[4:], "\n---\n"); i >= 0 {
		return s[4+i+5:]
	}
	return s
}

func rewriteLinks(text string, e Entry) string {
	dir := filepath.Dir(e.Src)
	return linkRE.ReplaceAllStringFunc(text, func(m string) string {
		parts := linkRE.FindStringSubmatch(m)
		if len(parts) < 2 {
			return m
		}
		href := parts[1]
		if strings.HasPrefix(href, "http") || strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "mailto:") || strings.HasPrefix(href, "/") {
			return m
		}
		anchor := ""
		if i := strings.IndexByte(href, '#'); i >= 0 {
			anchor = href[i:]
			href = href[:i]
		}
		target := filepath.ToSlash(filepath.Clean(filepath.Join(dir, href)))
		url := fmt.Sprintf("https://github.com/uidbz/%s/blob/%s/%s%s", e.Repo, e.Ref, target, anchor)
		return "](" + url + ")" + parts[2]
	})
}

func repoHTTPS(dir string) string {
	remote := git(dir, "remote", "get-url", "origin")
	remote = strings.TrimSuffix(remote, ".git")
	if strings.HasPrefix(remote, "git@github.com:") {
		return "https://github.com/" + strings.TrimPrefix(remote, "git@github.com:")
	}
	return remote
}

func git(dir string, args ...string) string {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	b, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}
