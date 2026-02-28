# Contributing

Contributions are welcome. By participating, you agree to maintain a respectful and constructive environment.

For coding standards, testing patterns, architecture guidelines, commit conventions, and all
development practices, refer to the **[Development Guide](https://github.com/rios0rios0/guide/wiki)**.

## Prerequisites

- [Go](https://go.dev/dl/) 1.26+
- [Git](https://git-scm.com/downloads) 2.30+

## Development Workflow

1. Fork and clone the repository
2. Create a branch: `git checkout -b feat/my-change`
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Build the project:
   ```bash
   go build -o ronin-to-koinly .
   ```
5. Run the application:
   ```bash
   ./ronin-to-koinly
   ```
6. Run tests:
   ```bash
   go test ./...
   ```
7. Format and vet the code:
   ```bash
   go fmt ./...
   go vet ./...
   ```
8. Update `CHANGELOG.md` under `[Unreleased]`
9. Commit following the [commit conventions](https://github.com/rios0rios0/guide/wiki/Life-Cycle/Git-Flow)
10. Open a pull request against `main`
