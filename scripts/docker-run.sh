#!/usr/bin/env bash
set -e

echo "Building Docker image for omniping..."
docker build -t omniping:latest .

echo "Running omniping container against 1.1.1.1..."
docker run --rm omniping:latest --target 1.1.1.1
