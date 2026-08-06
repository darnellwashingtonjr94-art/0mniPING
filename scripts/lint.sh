#!/usr/bin/env bash
set -e

echo "Running golangci-lint..."
if command -v golangci-lint &> /dev/null; then
    golangci-lint run ./...
else
    echo "golangci-lint is not installed. Running go vet instead."
    go vet ./...
fi
echo "Linting complete."
