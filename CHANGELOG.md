# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The Unreleased section on `CHANGELOG.md` gets a version number and date;
3. Open a Pull Request with the bump version changes targeting the `main` branch;
4. When the Pull Request is merged, a new Git tag must be created using [GitHub environment](https://github.com/rios0rios0/ronin-to-koinly/tags).

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

## [0.1.8] - 2026-06-18

### Changed

- changed the Go module dependencies to their latest versions

## [0.1.7] - 2026-06-03

### Changed

- changed the Go version to `1.26.4` and updated all module dependencies
- refreshed `CLAUDE.md` to drop the hard-coded version reference that drifts every release

## [0.1.6] - 2026-05-25

### Changed

- refreshed `CLAUDE.md` to correct the version reference from v0.1.3 to v0.1.5

## [0.1.5] - 2026-05-22

### Changed

- changed the Go module dependencies to their latest versions

## [0.1.4] - 2026-05-19

### Changed

- changed the Go module dependencies to their latest versions
- refreshed `CLAUDE.md` to correct the version reference from v0.1.1 to v0.1.3

## [0.1.3] - 2026-05-08

### Changed

- changed the Go version to `1.26.3` and updated all module dependencies

## [0.1.2] - 2026-04-28

### Changed

- refreshed `.github/copilot-instructions.md` to fix incorrect file paths, add Makefile build targets, correct the repository structure, and remove a false claim about csv struct tags
- refreshed `CLAUDE.md` to correct the version reference from v0.1.0 to v0.1.1

## [0.1.1] - 2026-04-15

### Changed

- changed the Go version to `1.26.2` and updated all module dependencies

## [0.1.0] - 2026-03-12

### Added

- added GitHub Actions workflow for CI/CD pipeline

### Changed

- changed the Go version to `1.26.1` and updated all module dependencies

