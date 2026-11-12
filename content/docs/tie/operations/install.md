---
title: Install and run
type: docs
weight: 20
---

## From source

```bash
git clone https://github.com/uidbz/tie
cd tie
make build              # binaries in dist/
make install            # client only, rootless, to ~/.local/bin
sudo make install-server
```

`make install-server` runs `contrib/install.sh`: it creates a `tie` service user,
`/etc/tie`, `/var/lib/tie`, and installs systemd or OpenRC units (plus the
Gentoo ebuild metadata). The client config lives in `~/.config/tie/config.toml`
or in the current directory.

`make check` runs `scrub`, `go vet`, `staticcheck`, `govulncheck` and the race
tests. CI runs the same checks.

## Docker

```bash
docker build -t tie .
docker run -d --name tie -p 1161:1161 -p 1162:1162 -v tie-data:/data tie
```

The image contains `tie-triplestore` and `tie-filehost`; the data volume holds
`db/` and `data/`. For the compute plane, add `tie-runner` with its own socket.

## Ports and services

| Service | Default port | Config |
|---|---|---|
| triplestore | 1161 | `tie-triplestore.toml` |
| filehost | 1162 | `tie-filehost.toml` |
| runner | 1163 | `tie-runner.toml` |

The triplestore and filehost run as separate processes on purpose. The
triplestore is the only writer to the `.tie` file; the filehost owns the blob
directory and can be on another machine.

## Gentoo

The `contrib/gentoo` overlay has ebuilds for the releases and a `9999` live
build. `contrib/README.md` covers the service layout, TLS termination and log
locations.

## Python client

```bash
pip install tie-client
```

The Python package is stdlib-only (3.11+), byte-compatible with the Go client,
and does not include the mount. It is a separate distribution with its own
README.
