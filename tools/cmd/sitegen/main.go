// Command sitegen keeps the tie website up to date from the sibling projects.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/uidbz/tiedb.com/tools/internal/check"
	"github.com/uidbz/tiedb.com/tools/internal/docsync"
	"github.com/uidbz/tiedb.com/tools/internal/facts"
	"github.com/uidbz/tiedb.com/tools/internal/logo"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	cmd := os.Args[1]
	args := os.Args[2:]
	switch cmd {
	case "logos":
		fs := flag.NewFlagSet("logos", flag.ExitOnError)
		root := fs.String("root", "..", "repository root")
		spec := fs.String("spec", "specs/marks.yaml", "logo spec")
		fs.Parse(args)
		if err := runLogos(*root, *spec); err != nil {
			fatal(err)
		}
	case "facts":
		fs := flag.NewFlagSet("facts", flag.ExitOnError)
		root := fs.String("root", "..", "repository root")
		src := fs.String("src", "../..", "parent directory of the sibling checkouts")
		fs.Parse(args)
		if err := facts.Generate(*root, *src); err != nil {
			fatal(err)
		}
		fmt.Println("facts: wrote data/facts.yaml")
	case "docs":
		fs := flag.NewFlagSet("docs", flag.ExitOnError)
		root := fs.String("root", "..", "repository root")
		src := fs.String("src", "../..", "parent directory of the sibling checkouts")
		spec := fs.String("spec", "specs/docs.yaml", "sync manifest")
		checkOnly := fs.Bool("check", false, "report stale pages without writing")
		fs.Parse(args)
		stale, err := docsync.Sync(*root, *src, *spec, *checkOnly)
		if err != nil {
			fatal(err)
		}
		if *checkOnly {
			if len(stale) > 0 {
				for _, p := range stale {
					fmt.Println("stale:", p)
				}
				fmt.Fprintf(os.Stderr, "docs: %d synced page(s) out of date; run make docs\n", len(stale))
				os.Exit(1)
			}
			fmt.Println("docs: all synced pages are current")
			return
		}
		fmt.Println("docs: synced reference pages")
	case "check":
		fs := flag.NewFlagSet("check", flag.ExitOnError)
		root := fs.String("root", "..", "repository root")
		fs.Parse(args)
		r, err := check.Run(*root)
		if err != nil {
			fatal(err)
		}
		for _, s := range r.Info {
			fmt.Println("info:", s)
		}
		for _, s := range r.Warnings {
			fmt.Println("warn:", s)
		}
		for _, s := range r.Errors {
			fmt.Println("error:", s)
		}
		if len(r.Errors) > 0 {
			fmt.Fprintf(os.Stderr, "check: %d error(s)\n", len(r.Errors))
			os.Exit(1)
		}
		fmt.Println("check: clean")
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", cmd)
		usage()
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `sitegen - keep the tie website current

Usage: sitegen <command> [flags]

Commands:
  logos   regenerate the logo family, favicons, brand assets and social cards
  facts   refresh data/facts.yaml from the sibling repositories
  docs    sync the long-form source notes into the reference docs
  check   check internal links, alt text, budgets and privacy tokens

Run from the tools directory (the Makefile does this for you).
`)
	os.Exit(2)
}

func runLogos(root, specPath string) error {
	if !filepath.IsAbs(specPath) {
		if _, err := os.Stat(specPath); err != nil {
			specPath = filepath.Join("specs", filepath.Base(specPath))
		}
	}
	s, err := logo.LoadSpec(specPath)
	if err != nil {
		return err
	}
	if err := s.Generate(root); err != nil {
		return err
	}
	fmt.Printf("logos: wrote %d marks, favicons, brand assets and social cards\n", len(s.Marks))
	return nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "sitegen:", err)
	os.Exit(1)
}
