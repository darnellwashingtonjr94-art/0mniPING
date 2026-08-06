#!/usr/bin/env bash
set -e

echo "Building omniping binary..."
mkdir -p bin/
go build -o bin/omniping cmd/omniping/main.go
echo "Build complete: bin/omniping"
