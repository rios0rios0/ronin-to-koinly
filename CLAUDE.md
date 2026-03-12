# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`ronin-to-koinly` is a Go CLI that fetches transaction data from the Ronin blockchain wallet API and exports it to a CSV file formatted for [Koinly](https://koinly.io/) (crypto tax tracking). Currently at v0.1.0 with the core pipeline partially implemented (API response unmarshalling is a placeholder).

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
- Go 1.26+ required

## CI/CD

GitHub Actions workflow (`.github/workflows/default.yaml`) delegates entirely to the shared reusable workflow at `rios0rios0/pipelines/.github/workflows/go-binary.yaml@main`. Triggers on pushes to main, tag pushes, PRs targeting main, and manual dispatch.

## Current State

The API response unmarshalling is commented out as a placeholder -- the actual Ronin API response schema needs to be mapped. The output CSV is always named `koinly_transactions.csv` in the current working directory.
