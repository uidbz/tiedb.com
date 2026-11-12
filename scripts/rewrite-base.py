#!/usr/bin/env python3
"""Rewrite root-absolute URLs for a subpath deployment (GitHub project Pages).

Hugo's `relURL` handles links built by templates, but a handful of links in the
content are intentionally root-absolute. When the site is previewed at
`https://<user>.github.io/<repo>/`, those links would escape the project path.
This script prefixes them once, after the build. On a custom domain (path "/")
it does nothing.

Usage: rewrite-base.py PUBLIC_DIR BASE_URL
"""
import os
import re
import sys
from urllib.parse import urlsplit

ATTR = re.compile(r'\b(href|src|poster)=(?:"|\')?(/(?!/)[^"\'\s>]*)')
SRCSET = re.compile(r'\bsrcset=(?:"|\')?([^"\'>]+)')


def rewrite_url(url, base):
    if url == base or url.startswith(base + "/"):
        return url
    return base + url


def rewrite_attr(match, base):
    return f'{match.group(1)}="{rewrite_url(match.group(2), base)}"'


def rewrite_srcset(match, base):
    parts = []
    for item in match.group(1).split(","):
        bits = item.strip().split()
        if bits and bits[0].startswith("/"):
            bits[0] = rewrite_url(bits[0], base)
        parts.append(" ".join(bits))
    return 'srcset="' + ", ".join(parts) + '"'


def main():
    if len(sys.argv) != 3:
        sys.exit("usage: rewrite-base.py PUBLIC_DIR BASE_URL")
    public, base_url = sys.argv[1], sys.argv[2]
    base = urlsplit(base_url).path.rstrip("/")
    if base in ("", "/"):
        print("rewrite-base: root deployment, nothing to do")
        return
    count = 0
    for root, _, files in os.walk(public):
        for name in files:
            if not name.endswith(".html"):
                continue
            path = os.path.join(root, name)
            with open(path, encoding="utf-8") as f:
                text = f.read()
            new = ATTR.sub(lambda m: rewrite_attr(m, base), text)
            new = SRCSET.sub(lambda m: rewrite_srcset(m, base), new)
            if new != text:
                with open(path, "w", encoding="utf-8") as f:
                    f.write(new)
                count += 1
    print(f"rewrite-base: prefixed root URLs with {base} in {count} pages")


if __name__ == "__main__":
    main()
