# dr-demo-telemetry-collector

Ingests service telemetry and forwards to the metrics store.

## Overview

`dr-demo-telemetry-collector` is a Go service in the observability domain. It runs as an internal worker
with a small HTTP control surface.

## Build

```bash
go build ./...
go test ./...
```

## Layout

- `main.go` — HTTP control surface
- `internal/store/store.go` — database access
- `internal/ops/ops.go` — operational helpers
- `internal/ops/ops_test.go` — fixtures
