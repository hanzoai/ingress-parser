# ingress-parser

Fork of Traefik's [`paerser`](https://github.com/traefik/paerser) (`github.com/hanzoai/ingress-parser`): loads configuration from CLI flags, config files (YAML/TOML/JSON), and environment variables, plus a small CLI command system. Used by the Hanzo ingress stack for config decoding.

Packages: `flag`, `file`, `env` (each exposes `Decode`).

- Test: `go test ./...`

Full docs: README.md
