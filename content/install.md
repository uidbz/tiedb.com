---
title: Install
type: default
eyebrow: "Install"
description: "Build tie, the tie apps and anneal from source, or run the backend in Docker."
---

Everything here is built from source. There are no tagged binary releases for the
apps or for anneal yet, and tie's releases are source tags. That is deliberate:
the projects are moving, and a binary built from the commit you are reading is
more honest than a stale download.

## Requirements

| Project | Needs |
|---|---|
| tie | Go 1.27.2+, Linux for FUSE, `fusermount3`, optionally Docker |
| the apps | Go 1.27.2+, CGO, X11/Wayland + OpenGL, libmpv, PipeWire (tie-audio), libmtp (tie-fm, Linux) |
| anneal | Go 1.27.2+, Fyne's C dependencies for the GUI; the headless CLI and server need no GUI libs |

On a Debian/Ubuntu system the Fyne and media packages are roughly:

```bash
sudo apt install build-essential pkg-config libgl1-mesa-dev \
  xorg-dev libxkbcommon-dev libwayland-dev \
  libmpv-dev libpipewire-0.3-dev libmtp-dev
```

The exact package names differ by distribution; the Fyne getting-started page and
each project's README have the current list.

{{< tabs >}}
{{< tab name="tie" >}}
```bash
git clone https://github.com/uidbz/tie
cd tie
make build          # dist/
make install        # client to ~/.local/bin
sudo make install-server   # services, /etc/tie, /var/lib/tie
```

Docker:

```bash
docker build -t tie .
docker run -d --name tie -p 1161:1161 -p 1162:1162 -v tie-data:/data tie
```

Gentoo: the overlay and ebuilds are in `contrib/gentoo`.

{{< /tab >}}
{{< tab name="the apps" >}}
```bash
git clone --recurse-submodules https://github.com/uidbz/tie-gui
cd tie-gui
make install        # imgview, tie-view, tie-fm, tie-audio
go build ./cmd/tie-browser
```

The `third_party/fyne` submodule is required; `go install …@latest` will not work
because of the `replace` directive.

Android: `docs/ANDROID.md` and the `build-android.sh` /
`build-install-android.sh` scripts. The APKs are debug-signed.

{{< /tab >}}
{{< tab name="anneal" >}}
```bash
git clone https://github.com/uidbz/anneal
cd anneal
make build          # bin/anneal, bin/anneal-headless, bin/anneald
make install        # ~/.local/bin, .desktop entry, icon
make docker         # anneal:dev
```

Or run the backend in a container:

```bash
docker run -d --name anneal -v ~/src:/work -p 7777:7777 \
  -e ANNEAL_TOKEN=change-me anneal:dev
```

{{< /tab >}}
{{< /tabs >}}

## Verify the install

```bash
tie version
tie conf show
anneal-headless -h
```

`tie version` prints the client and the versions of the configured triplestore
and filehosts, which is the quickest way to catch a config pointing at the wrong
place.

## Platforms

| | Linux | Windows | macOS | Android |
|---|---|---|---|---|
| tie (FUSE mount) | ✓ | SMB client | SMB client | — |
| tie-fm | ✓ | ✓ | ✓ | — |
| tie-audio | ✓ | ✓ | ✓ | ✓ |
| tie-view | ✓ | ✓ | ✓ | ✓ |
| imgview | ✓ | ✓ | ✓ | ✓ |
| tie-browser | ✓ | ✓ | ✓ | — |
| anneal | ✓ (GUI) | headless | headless | — |

The FUSE mount is Linux-only. The same tree is available on Windows and macOS
through `tie-smb`. tie-fm's device browsing (USB, MTP) is Linux-only.

## A first run in 60 seconds

The quickest end-to-end test is tie's test environment:

```bash
cd tie/test-env
./build.sh && ./start.sh && ./seed.sh
./mount-db.sh
ls mnt/query/vacation
```

Then point an app at the same server (ports 2161/2162, `defaultuser` /
`defaultpassword`). The apps' quickstart has the config.
