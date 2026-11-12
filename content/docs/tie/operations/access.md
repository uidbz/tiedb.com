---
title: Access control
type: docs
weight: 21
---

Both HTTP services use Basic auth with three roles:

```
none < read < write
```

Users live in the service config:

```toml
[[Users]]
Username = "you"
Password = "change-me"
Role = "write"
```

`AnonymousAccess` can grant `none`, `read` or `write` to clients without
credentials. That is useful on a loopback-only test environment; it is a bad
idea on a LAN.

## TLS

The servers speak HTTP. Terminate TLS at a reverse proxy (Caddy, nginx, Traefik)
and pass the `Authorization` header through. The `tie` CLI accepts
`--server https://…` and `tie-smb` is a separate protocol.

## Exposure

The intended deployments are:

- loopback on one machine;
- a small trusted LAN, behind the host firewall;
- a WireGuard/Tailscale network for remote access.

There is no rate limiting, no audit log and no per-collection ACL. If you need
those, put tie behind a proxy that provides them, or keep it on the private
network.

## Passwords

Passwords are stored in plaintext in the service config files, which are created
`0600`. Protect the config, or use the test environment's throwaway credentials
for demos. This is documented as a limitation, not a feature.
