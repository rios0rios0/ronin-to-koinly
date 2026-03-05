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
│   └── workflows/
│       └── default.yaml          # CI/CD pipeline (delegates to shared Go binary workflow)
├── CHANGELOG.md                  # Version history following Keep a Changelog
├── CONTRIBUTING.md               # Development workflow and prerequisites
├── LICENSE                       # MIT license
├── README.md                     # Project description and usage guide
├── go.mod                        # Go module definition (module: github.com/rios0rios0/ronin-to-koinly)
├── go.sum                        # Dependency checksums
└── main.go                       # Single-file entrypoint: API fetch → CSV write
```

## Technology Stack

- **Language**: Go 1.26+
- **HTTP client**: [`github.com/go-resty/resty/v2`](https://github.com/go-resty/resty) v2.17.2
- **Indirect dependency**: `golang.org/x/net`

## Build, Test, and Run Commands

```bash
# Install / tidy dependencies
go mod download

# Build the binary
go build -o ronin-to-koinly .

# Run the binary (requires network access to the Ronin API)
./ronin-to-koinly

# Run all tests
go test ./...

# Format code
go fmt ./...

# Vet code
go vet ./...
```

Build and tests typically complete in a few seconds on any modern machine.

## Architecture and Design Patterns

- **Single-file CLI**: all logic lives in `main.go`; there are no sub-packages yet.
- **Transaction struct** mirrors the Koinly CSV column layout and is tagged with `csv:` field tags.
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

## Development Workflow

1. Fork and clone the repository.
2. Create a feature branch: `git checkout -b feat/my-change`
3. Install dependencies: `go mod download`
4. Make changes to `main.go` (or add new `.go` files as the project grows).
5. Format and vet: `go fmt ./... && go vet ./...`
6. Run tests: `go test ./...`
7. Update `CHANGELOG.md` under `[Unreleased]`.
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
| Cross-compile (e.g., Linux amd64) | `GOOS=linux GOARCH=amd64 go build -o ronin-to-koinly .` |

## Troubleshooting

- **API errors**: The Ronin wallet API endpoint (`https://api.roninchain.com/wallet/transactions`)
  may require authentication headers or a wallet address query parameter. Update the `client.R().Get()`
  call in `main.go` accordingly.
- **Empty CSV**: The JSON → struct unmarshalling is currently a placeholder (commented out). Implement
  `json.Unmarshal(resp.Body(), &transactions)` with the correct struct mapping once the API response
  schema is known.
- **Go version mismatch**: Ensure your local Go version is 1.26 or higher (`go version`).
