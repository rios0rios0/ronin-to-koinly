# AGENTS.md

This file provides guidance to AI coding agents (Claude Code, Codex, Cursor, GitHub Copilot) when working with code in this repository.

## Project Overview

`ronin-to-koinly` is a Go CLI that fetches transaction data from the Ronin blockchain wallet API and exports it to a CSV file formatted for [Koinly](https://koinly.io/) (crypto tax tracking). The core pipeline is only partially implemented -- API response unmarshalling is still a placeholder.

## Build and Development Commands

Quality gates run through Makefile targets that include the shared makefiles from `~/Development/github.com/rios0rios0/pipelines`. Use those targets rather than the tool binaries (golangci-lint, Semgrep, Trivy, Gitleaks, ...), because the shared targets load each tool's configuration:

```bash
make build        # Build binary to bin/ronin-to-koinly (stripped)
make debug        # Debug build (no optimizations)
make build-musl   # Static build with musl-gcc
make run          # Run via go run
make install      # Build and copy to ~/.local/bin/

make lint         # Run golangci-lint (via pipelines)
make test         # Run test suite (via pipelines)
make sast         # Run full SAST suite: CodeQL, Semgrep, Trivy, Hadolint, Gitleaks
```

The include path is `SCRIPTS_DIR` (override with `make lint SCRIPTS_DIR=<path>`). The includes are optional, so without that checkout `make lint`, `make test`, and `make sast` fail with "No rule to make target" while the local build targets above still work.

## Architecture

Single-file CLI at `cmd/ronin-to-koinly/main.go`. The pipeline is: resty HTTP client fetches JSON from Ronin API -> unmarshals into `[]Transaction` -> writes to `koinly_transactions.csv` in the current working directory via `encoding/csv`.

The `Transaction` fields and the CSV header in `writeCSV` follow Koinly's column order; dates are written as `2006-01-02 15:04:05`. Changing either changes the import format.

The `internal/` and `test/` directories are scaffolded (empty with `.gitkeep`) for future expansion into domain/infrastructure layers.

## CI/CD

All three workflows (`default.yaml`, `claude-review.yaml`, `claude-mention.yaml`) only call reusable workflows in `rios0rios0/pipelines`; pipeline behavior is changed there, not in this repository.

<!-- chlog:start -->
## Changelog (chlog) — MANDATORY

If the repository you are working in uses chlog (a `.chlog.yaml` or `.chlog.yml`
config file, or a `.changes/` directory, exists at the project root), the
following is binding and ALWAYS applies: whenever you make ANY change, you MUST
create a changelog fragment as part of the same change — automatically, without
being asked, before committing.

- Do NOT edit CHANGELOG.md directly; it is generated from fragments.
- Create the fragment with:
  `chlog new --kind <Kind> --body '<past-tense description>'`
- Write an apostrophe inside the single-quoted body as `'\''`.
- Valid kinds: Added, Changed, Deprecated, Removed, Fixed, Security
- Choose the kind that best matches the change (e.g., new feature → Added,
  bug fix → Fixed, behavior change → Changed, removal → Removed, security fix → Security).
- If the change is backward-INCOMPATIBLE with the public API (a breaking
  change), you MUST add the `--breaking` flag:
  `chlog new --kind <Kind> --breaking --body '<past-tense description>'`.
  This is the ONLY thing that triggers a major version bump — the kind alone
  never does (per SemVer, major = incompatible change). When unsure whether a
  change breaks compatibility, ask the user instead of guessing.
- Fragments are YAML files in `.changes/unreleased/`; stage them with your commit.
- `chlog check` fails the build when a fragment is missing — never skip it.
<!-- chlog:end -->
