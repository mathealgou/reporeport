#!/usr/bin/bash
set -euo pipefail

APP_NAME="reporeport"
OUTPUT_DIR="dist"
mkdir -p "${OUTPUT_DIR}"

export GOOS=windows
export GOARCH=amd64

go mod tidy
go build -o "${OUTPUT_DIR}/${APP_NAME}.exe"