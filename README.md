# Wii/DS Game Stats Server

OpenPak's Wii/DS Game Stats Server.

| | |
| --- | --- |
| Network | Wii, DS |
| Role | GameSpy gamestats TCP and HTTP statistics protocol service |
| Depends on | Existing WFC database and shared `WFC_TOKEN_KEY_FILE`; no remote WFC service |
| Image | `ghcr.io/openpak/wfc-gamestats` |
| Based on | WiiLink WFC (AGPL-3.0) |
| Licence | AGPL-3.0-only |

## What it serves

| Endpoint | What |
| --- | --- |
| GameSpy gamestats | TCP statistics protocol |
| Internal HTTP | `WFC_HTTP_ADDR` (default `:20205`); `GET /healthz` is unauthenticated |

## Build

```sh
docker build -t wfc-gamestats .
```

## Docs

- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)
- [Notes](docs/notes.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
