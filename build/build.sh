#!/bin/bash

set -e

git config --global --add safe.directory /build

echo "Running tests..."

# go test go/...

echo "Building amd64..."
GOARCH=amd64 go build -o release/pihole-api_linux_amd64

# echo "Building arm64..."
# GOARCH=arm64 go build -o release/pihole-api_linux_arm64

# echo "Building 386..."
# GOARCH=386 go build -o release/pihole-api_linux_386