# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Beacon is a Go CLI tool that enables EC2 user data scripts to emit custom lifecycle events to AWS EventBridge during instance initialization. It solves the observability gap for EC2 auto-scaling startup processes.

## Commands

This project uses [Task](https://taskfile.dev) (v3.x) as the build runner — not `make`.

```bash
task build        # Format code and build .build/beacon
task test         # Run all unit tests
task fmt          # Format Go code
task check        # Run all security scans (gosec, go vet, govulncheck)
task coverage     # Generate test coverage report
task dist         # Build for all target platforms (linux/darwin/windows, multiple arches)
task clean        # Remove .build/ and .dist/ artifacts
```

Run a single test:
```bash
go test ./... -run TestFunctionName
go test ./cmd/... -run TestChooseStatusMessage   # cmd package tests
go test . -run TestEmit                          # root package tests
```

## Architecture

### Package Layout

The repo has two packages:

- **Root package (`beacon`)** — importable library: event emission, IMDS access, config loading, type definitions with validation
- **`cmd/` package** — Cobra CLI that wires the library together

### Core Data Flow

```
beacon CLI flags
    ↓
LoadConfig()          # /etc/beacon/config.yml or --config path
AWS SDK Config        # standard credential chain
RetrieveInstanceARN() # fetches from EC2 IMDS: account/region/instance-id → ARN
chooseStatusMessage() # priority: --status > --fail > --pass > --info > default
    ↓
Emitter.Emit()
    ↓ validates payload size (<256KB) and ARN format
PutEvents → EventBridge
```

### Key Design Decisions

**Interface-based AWS clients** — `EventBridgeClient` and `IMDSClient` are interfaces defined in `types.go`. The CLI passes real SDK clients; tests inject mocks via struct function fields. New AWS interactions should follow this same pattern.

**Validation on types** — Domain types (`Status`, `DetailType`, `InstanceARN`, `Project`) each carry a `Validate() error` method. Validation is structural (size limits, regex, allowed values) and checked before any AWS call is made.

**Status flag priority** — `cmd/status.go:chooseStatusMessage()` resolves which flag wins when multiple are provided: `--status` (custom) → `--fail` → `--pass` → `--info` → default info. This is intentional UX, not a bug.

**IMDS vs explicit ARN** — The `Emitter` accepts either a pre-populated `InstanceARN` or fetches one at emit time via the `IMDSClient`. This lets the library work outside EC2 (e.g., tests) with an explicit ARN.

### Constants & Limits (constants.go)

AWS-imposed limits enforced locally before making API calls:
- Event payload: 256 KB
- Detail type: 128 characters
- ARN: 2048 characters
- Default config path: `/etc/beacon/config.yml`
- Default timeout: 30 seconds

### Version Injection

The `VERSION` variable in the root package is set at build time via ldflags:
```
-X github.com/scottbrown/beacon.VERSION={{.VERSION}}
```

### CI

GitHub Actions runs `task test` and `task build` on every push and PR (`.github/workflows/test.yml`). Dependabot PRs for `gomod` and `github-actions` are auto-merged weekly via `.github/workflows/automerge-dependabot.yml` (minor/patch only).

## Conventions

- **Constants**: `UPPER_SNAKE_CASE` (e.g., `STATUS_FAIL`, `DEFAULT_CONFIG_PATH`)
- **Error wrapping**: `fmt.Errorf("context: %w", err)` — no custom error types
- **Tests**: table-driven with `[]struct` slices; mock clients use exported function fields (`PutEventsFn`, `GetInstanceIdentityDocumentFn`)
- **IAM**: the tool requires only `events:PutEvents` — keep it that way
