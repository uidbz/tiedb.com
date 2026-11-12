---
title: SMB shares
type: docs
weight: 14
---

`tie-smb` serves the live tree over SMB2/3 to Windows Explorer, macOS Finder and
Linux CIFS clients. It is pure Go: no Samba, no `smb.conf`, no exports table.

```toml
# tie-smb.toml
Listen = ":445"
ShareName = "tie"
ReadOnly = false

[[Users]]
Username = "you"
Password = "change-me"
Role = "write"        # none < read < write
```

```bash
tie-smb -config tie-smb.toml
```

It can also serve a `.tie` file directly:

```bash
tie-smb -f ~/.local/share/tie/Main.tie -config tie-smb.toml
```

## Clients

- **Windows**: `\\server\tie` in Explorer.
- **macOS**: Finder → Go → Connect to Server → `smb://server/tie`.
- **Linux**: `mount -t cifs //server/tie /mnt -o user=you`.

The tree is the same `query/`, `files/` and `tags/` view as FUSE. Writes go
through the same store, so a file copied over SMB is visible to the apps
immediately.

## Limits (v1)

- SMB 3.1.1 dialects are not implemented; use SMB 2/3.
- No SMB encryption or signing policy.
- Name escaping follows the SMB rules; the source doc lists the exact cases.
- Authentication is NTLMv2 with the users in the config; there is no domain
  controller integration.

Run it on a trusted LAN, or put it behind WireGuard. The full guide, including
troubleshooting, is
[`docs/smb.md`](https://github.com/uidbz/tie/blob/master/docs/smb.md).
