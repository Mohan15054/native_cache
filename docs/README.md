# GoCache

A modular, scalable, and extensible native caching system in Go.

## Features

- Pluggable storage engines
- Configurable eviction policies (LRU, LFU, etc.)
- TTL and max memory support
- Metrics and monitoring
- CLI and REST API interfaces

## Project Structure

- `cmd/` - CLI entry points
- `internal/` - Core logic (storage, eviction, config, metrics)
- `api/` - REST API server
- `configs/` - Configuration files
- `docs/` - Documentation
- `test/` - Integration and system tests

## Getting Started

1. `go mod tidy`
2. `go run main.go`
