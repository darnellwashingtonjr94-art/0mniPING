#!/usr/bin/env bash
set -e

echo "Running all unit and integration tests with coverage..."
go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
echo "Tests complete."
