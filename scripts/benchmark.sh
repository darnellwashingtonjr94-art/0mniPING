#!/usr/bin/env bash
set -e

echo "Running performance benchmarks..."
go test -bench=. -benchmem ./tests/benchmarks/...
echo "Benchmarks complete."
