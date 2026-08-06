#!/usr/bin/env bash
set -e

echo "Cleaning build artifacts and logs..."
rm -rf bin/ *.log data/
echo "Clean complete."
