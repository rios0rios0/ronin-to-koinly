# Copilot Instructions

## Project Overview

`ronin-to-koinly` is a Go CLI application that fetches transaction data from the Ronin blockchain
wallet API and exports it to a CSV file formatted for import into [Koinly](https://koinly.io/), a
crypto tax and portfolio tracking tool.

## Repository Structure

```
.
├── .github/
│   ├── copilot-instructions.md   # AI assistant context (this file)
│   ├── skills/
│   │   └── code-review/SKILL.md  # Copilot code-review skill for this repo
│   └── workflows/
│       ├── default.yaml          # CI/CD pipeline (delegates to shared Go binary workflow)
│       ├── claude-review.yaml    # Automated PR review (shared reusable workflow)
│       └── claude-mention.yaml   # @claude mention responder (shared reusable workflow)
├── cmd/
│   └── ronin-to-koinly/
│       └── main.go               # CLI entrypoint: API fetch → CSV write
├── internal/                     # Scaffolded for future domain/infrastructure layers
├── test/                         # Scaffolded for future tests
├── CHANGELOG.md                  # Version history following Keep a Changelog
├── CONTRIBUTING.md               # Development workflow and prerequisites
├── LICENSE                       # MIT license
├── Makefile                      # Build targets (wraps shared pipelines Makefile)
├── README.md                     # Project description and usage guide
├── go.mod                        # Go module definition (module: github.com/rios0rios0/ronin-to-koinly)
└── go.sum                        # Dependency checksums
```

## Technology Stack

- **Language**: Go 1.27+
- **HTTP client**: [`github.com/go-resty/resty/v2`](https://github.com/go-resty/resty) v2.17.2
- **Indirect dependency**: `golang.org/x/net`

## Build, Test, and Run Commands

All quality gate commands use the shared Makefile from `~/Development/github.com/rios0rios0/pipelines`. Always use Makefile targets instead of calling tool binaries directly:

```bash
make build        # Build binary to bin/ronin-to-koinly (stripped)
make debug        # Debug build (no optimizations)
make run          # Run via go run
make install      # Build and copy to ~/.local/bin/

make lint         # Run golangci-lint (via pipelines)
make test         # Run test suite (via pipelines)
make sast         # Run full SAST suite: CodeQL, Semgrep, Trivy, Hadolint, Gitleaks

go mod download   # Install dependencies
go mod tidy       # Remove unused dependencies
```

## Architecture and Design Patterns

- **Single-file CLI**: all logic lives in `cmd/ronin-to-koinly/main.go`; there are no sub-packages yet.
- **Transaction struct** mirrors the Koinly CSV column layout (Date, SentAmount/Currency, ReceivedAmount/Currency, FeeAmount/Currency, Tag).
- **HTTP → CSV pipeline**: `resty` client fetches JSON from the Ronin API → data is unmarshalled
  into `[]Transaction` → written row-by-row to `koinly_transactions.csv` via the standard
  `encoding/csv` writer.
- Date values are formatted as `2006-01-02 15:04:05` (Go reference time) to match Koinly's
  expected input.

## CI/CD Pipeline

The pipeline is declared in `.github/workflows/default.yaml` and delegates entirely to the shared
reusable workflow at `rios0rios0/pipelines/.github/workflows/go-binary.yaml@main`, passing
`binary_name: ronin-to-koinly`. It runs on:

- Pushes to `main`
- Any tag push
- Pull requests targeting `main`
- Manual `workflow_dispatch`

The shared workflow handles linting, testing, building, and releasing the binary automatically.

Two further workflows drive AI automation via shared reusable workflows: `claude-review.yaml`
(automated pull-request review) and `claude-mention.yaml` (`@claude` mention responder), both
authenticating with the `CLAUDE_CODE_OAUTH_TOKEN` secret.

## Development Workflow

1. Fork and clone the repository.
2. Create a feature branch: `git checkout -b feat/my-change`
3. Install dependencies: `go mod download`
4. Make changes to `cmd/ronin-to-koinly/main.go` (or add new `.go` files as the project grows).
5. Format and vet: `go fmt ./... && go vet ./...`
6. Run tests: `go test ./...`
7. Add a changelog fragment: `chlog new --kind <Kind> --body "..."`, committed from `.changes/unreleased/`. Never edit `CHANGELOG.md` — it is generated from the fragments.
8. Commit following [conventional commits](https://github.com/rios0rios0/guide/wiki/Life-Cycle/Git-Flow).
9. Open a pull request against `main`.

## Coding Conventions

- Follow standard Go idioms and the conventions described in the
  [Development Guide](https://github.com/rios0rios0/guide/wiki).
- Prefer the standard library; only add third-party dependencies when strictly necessary.
- All exported types and functions must have Go doc comments.
- Use `log.Fatalf` for fatal errors at the CLI entry point; propagate errors up via return values in
  library-style functions.
- Output CSV file is always named `koinly_transactions.csv` and written to the current working
  directory.

## Common Tasks

| Task | Command |
|------|---------|
| Add a dependency | `go get <module>@<version>` |
| Remove unused dependencies | `go mod tidy` |
| Check for outdated deps | `go list -u -m all` |
| Static musl build | `make build-musl` |

## Troubleshooting

- **API errors**: The Ronin wallet API endpoint (`https://api.roninchain.com/wallet/transactions`)
  may require authentication headers or a wallet address query parameter. Update the `client.R().Get()`
  call in `cmd/ronin-to-koinly/main.go` accordingly.
- **Empty CSV**: The JSON → struct unmarshalling is currently a placeholder (commented out). Implement
  `json.Unmarshal(resp.Body(), &transactions)` with the correct struct mapping once the API response
  schema is known.
- **Go version mismatch**: Ensure your local Go version is 1.27 or higher (`go version`).

<!-- chlog:start -->
## Changelog (chlog) — MANDATORY

If the repository you are working in uses chlog (a `.chlog.yaml` or `.chlog.yml`
config file, or a `.changes/` directory, exists at the project root), the
following is binding and ALWAYS applies: whenever you make ANY change, you MUST
create a changelog fragment as part of the same change — automatically, without
being asked, before committing.

- Do NOT edit CHANGELOG.md directly; it is generated from fragments.
- Create the fragment with:
  `chlog new --kind <Kind> --body "<imperative description>"`
- Valid kinds: Added, Changed, Deprecated, Removed, Fixed, Security
- Choose the kind that best matches the change (e.g., new feature → Added,
  bug fix → Fixed, behavior change → Changed, removal → Removed, security fix → Security).
- If the change is backward-INCOMPATIBLE with the public API (a breaking
  change), you MUST add the `--breaking` flag:
  `chlog new --kind <Kind> --breaking --body "<description>"`.
  This is the ONLY thing that triggers a major version bump — the kind alone
  never does (per SemVer, major = incompatible change). When unsure whether a
  change breaks compatibility, ask the user instead of guessing.
- Fragments are YAML files in `.changes/unreleased/`; stage them with your commit.
- `chlog check` fails the build when a fragment is missing — never skip it.
<!-- chlog:end -->
