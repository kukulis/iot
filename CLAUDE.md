# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go module (`darbelis.eu/iot`) for IoT GPS/telemetry data processing. Early-stage project being ported from PHP — some files may contain incomplete conversions. No external dependencies.

## Build & Run Commands

```bash
go build ./...                          # Build all packages
go run ./cmd/consumer                   # Run the consumer entry point
go test ./...                           # Run all tests
go test ./pkg/iot -run TestName         # Run a single test
go vet ./...                            # Static analysis
```

No Makefile, linter config, or CI pipeline exists yet.

## Architecture

- **`cmd/consumer/`** — Main entry point (currently a stub)
- **`pkg/iot/`** — Core domain: `MainRecord` (GPS data with hex serialization) and `InnerRecord` (flexible key-value container with size-categorized maps)
- **`pkg/util/`** — Low-level helpers: `Dechex`, `StrPadLeft`, `BoolToInt`
- **`internal/di/`** — Prepared for dependency injection (empty)

### Key Patterns

- **Hex encoding pipeline**: Domain records serialize to hex strings via `CreateHex()` methods using `Dechex` + `StrPadLeft`
- **Method chaining**: `InnerRecord` setters return `*InnerRecord` for fluent API
- **Constructor functions**: `NewMainRecord()`, `NewInnerRecord()`
