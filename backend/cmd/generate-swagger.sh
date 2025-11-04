#!/usr/bin/env bash
set -euo pipefail

echo "=== Swagger docs generation (Bash) ==="

# Resolve repo root (backend/cmd -> repo root two levels up)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
cd "$REPO_ROOT"

# Ensure Go is available
if ! command -v go >/dev/null 2>&1; then
  echo "Go is not installed or not on PATH. Please install Go and try again." >&2
  exit 1
fi

# Ensure GOPATH/bin is on PATH so the installed swag binary can be found
GOPATH_DIR="$(go env GOPATH)"
if [ -n "$GOPATH_DIR" ] && [ -d "$GOPATH_DIR/bin" ]; then
  case ":$PATH:" in
    *":$GOPATH_DIR/bin:"*) ;;
    *) export PATH="$GOPATH_DIR/bin:$PATH" ;;
  esac
fi

# Install swag if missing
if ! command -v swag >/dev/null 2>&1; then
  echo "Installing swag..."
  go install github.com/swaggo/swag@v1.8.12
fi

ENTRY="backend/cmd/api/main.go"
OUT_DIR="backend/docs"

echo "Generating Swagger docs from $ENTRY -> $OUT_DIR"
mkdir -p "$OUT_DIR"

swag init -g "$ENTRY" -o "$OUT_DIR" --parseDependency --parseInternal

echo "Swagger docs generated successfully at $OUT_DIR"


