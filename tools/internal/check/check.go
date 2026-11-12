// Package check runs the site's quality gates: internal links, alt text,
// budgets, privacy tokens and the hosting CSP rules.
package check

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Result struct {
	Errors   []string
	Warnings []string
	Info     []string
}

var (
	hrefRE    = regexp.MustCompile(`(?i)\bhref\s*=\s*"([^"]*)"`)
	srcRE     = regexp.MustCompile(`(?i)\bsrc\s*=\s*"([^"]*)"`)
	imgRE     = regexp.MustCompile(`(?is)<img\b[^>]*>`)
	blankRE   = regexp.MustCompile(`(?i)\btarget\s*=\s*"_blank"`)
	altRE     = regexp.MustCompile(`(?i)\balt(\s*=|\s|>)`)
	privateRE = regexp.MustCompile(`/home/johan|/tmp/opencode|/root/\.cache`)
)

// Run checks the built site under root/public.
func Run(root string) (Result, error) {
	public := filepath.Join(root, "public")
	if _, err := os.Stat(public); err != nil {
		return Result{}, fmt.Errorf("build the site first: %w", err)
	}
	var r Result
	var largestImg struct {
		size int64
		path string
	}
	var totalHTML int64
	var pages int

	err := filepath.Walk(public, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		ext := strings.ToLower(filepath.Ext(path))
		switch ext {
		case ".html":
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			pages++
			totalHTML += int64(len(b))
			rel, _ := filepath.Rel(public, path)
			page := "/" + filepath.ToSlash(rel)
			text := string(b)
			checkLinks(&r, public, page, text)
			checkAlts(&r, page, text)
			if blankRE.MatchString(text) {
				r.Errors = append(r.Errors, fmt.Sprintf("%s: target=\"_blank\" is blocked by the hosting CSP", page))
			}
			if privateRE.MatchString(text) {
				r.Errors = append(r.Errors, fmt.Sprintf("%s: contains a private path", page))
			}
			if strings.Contains(text, "src=\"http://") || strings.Contains(text, "src=\"https://") {
				// Remote scripts are blocked; remote images are allowed by the CSP
				// but we do not use them.
				r.Warnings = append(r.Warnings, fmt.Sprintf("%s: references a remote resource", page))
			}
		case ".webp", ".png", ".jpg", ".jpeg":
			if info.Size() > largestImg.size {
				largestImg.size = info.Size()
				largestImg.path = path
			}
			if info.Size() > 800*1024 {
				rel, _ := filepath.Rel(public, path)
				r.Errors = append(r.Errors, fmt.Sprintf("%s: image is %d KB (budget 800 KB)", rel, info.Size()/1024))
			}
		}
		return nil
	})
	if err != nil {
		return r, err
	}
	if pages == 0 {
		r.Errors = append(r.Errors, "no HTML pages found in public/")
	}
	home := filepath.Join(public, "index.html")
	if info, err := os.Stat(home); err == nil {
		if info.Size() > 250*1024 {
			r.Errors = append(r.Errors, fmt.Sprintf("home page HTML is %d KB (budget 250 KB)", info.Size()/1024))
		}
		r.Info = append(r.Info, fmt.Sprintf("home HTML: %d KB", info.Size()/1024))
	}
	if largestImg.path != "" {
		rel, _ := filepath.Rel(public, largestImg.path)
		r.Info = append(r.Info, fmt.Sprintf("largest image: %s (%d KB)", rel, largestImg.size/1024))
	}
	r.Info = append(r.Info, fmt.Sprintf("%d pages, %.1f KB average HTML", pages, float64(totalHTML)/float64(pages)/1024))
	sort.Strings(r.Errors)
	sort.Strings(r.Warnings)
	return r, nil
}

func checkLinks(r *Result, public, page, text string) {
	seen := map[string]bool{}
	for _, m := range append(hrefRE.FindAllStringSubmatch(text, -1), srcRE.FindAllStringSubmatch(text, -1)...) {
		href := m[1]
		if href == "" || strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") ||
			strings.HasPrefix(href, "mailto:") || strings.HasPrefix(href, "tel:") ||
			strings.HasPrefix(href, "data:") {
			continue
		}
		if i := strings.IndexAny(href, "?#"); i >= 0 {
			href = href[:i]
		}
		if !strings.HasPrefix(href, "/") {
			continue
		}
		if seen[href] {
			continue
		}
		seen[href] = true
		target := filepath.Join(public, filepath.FromSlash(strings.TrimPrefix(href, "/")))
		if strings.HasSuffix(href, "/") {
			target = filepath.Join(target, "index.html")
		}
		info, err := os.Stat(target)
		if err != nil {
			// A pretty URL may exist as a directory with index.html.
			if info2, err2 := os.Stat(filepath.Join(target, "index.html")); err2 == nil && info2.Mode().IsRegular() {
				continue
			}
			r.Errors = append(r.Errors, fmt.Sprintf("%s: broken link %s", page, href))
			continue
		}
		if info.IsDir() {
			if _, err := os.Stat(filepath.Join(target, "index.html")); err != nil {
				r.Errors = append(r.Errors, fmt.Sprintf("%s: broken link %s", page, href))
			}
		}
	}
}

func checkAlts(r *Result, page, text string) {
	for _, img := range imgRE.FindAllString(text, -1) {
		if !altRE.MatchString(img) {
			short := img
			if len(short) > 120 {
				short = short[:120] + "…"
			}
			r.Errors = append(r.Errors, fmt.Sprintf("%s: img without alt: %s", page, short))
		}
	}
}
