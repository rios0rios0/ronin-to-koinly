# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`ronin-to-koinly` is a Go CLI that fetches transaction data from the Ronin blockchain wallet API and exports it to a CSV file formatted for [Koinly](https://koinly.io/) (crypto tax tracking). The core pipeline is only partially implemented -- API response unmarshalling is still a placeholder.

## Build and Development Commands

All quality gate commands use the shared Makefile from `~/Development/github.com/rios0rios0/pipelines`. **Never call tool binaries directly** -- always use Makefile targets:

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

## Architecture

Single-file CLI at `cmd/ronin-to-koinly/main.go`. The pipeline is: resty HTTP client fetches JSON from Ronin API -> unmarshals into `[]Transaction` -> writes to `koinly_transactions.csv` via `encoding/csv`.

The `internal/` and `test/` directories are scaffolded (empty with `.gitkeep`) for future expansion into domain/infrastructure layers.

## Key Types

- **`Transaction`** (in `main.go`): maps to Koinly CSV columns -- Date, SentAmount/Currency, ReceivedAmount/Currency, FeeAmount/Currency, Tag.

## Dependencies

- `github.com/go-resty/resty/v2` -- HTTP client for Ronin API calls
- Go 1.27+ required

## CI/CD

GitHub Actions workflow (`.github/workflows/default.yaml`) delegates entirely to the shared reusable workflow at `rios0rios0/pipelines/.github/workflows/go-binary.yaml@main`. Triggers on pushes to main, tag pushes, PRs targeting main, and manual dispatch.

## Current State

The API response unmarshalling is commented out as a placeholder -- the actual Ronin API response schema needs to be mapped. The output CSV is always named `koinly_transactions.csv` in the current working directory.

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
