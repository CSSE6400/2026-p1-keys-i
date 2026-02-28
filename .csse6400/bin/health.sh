#!/bin/bash
#
# Check that the Go health endpoint is returning 200

set -euo pipefail

PORT=6400
URL="http://localhost:${PORT}/api/v1/health"

go mod download

go run ./cmd/api -p "${PORT}" &
pid=$!

cleanup() {
  # Don't fail cleanup if the process is already gone
  kill "${pid}" 2>/dev/null || true
}
trap cleanup EXIT

# Wait for server to start (retry up to ~10s)
for _ in {1..20}; do
  if curl -s -o /dev/null "${URL}"; then
    break
  fi
  sleep 0.5
done

# Assert HTTP 200
code="$(curl -s -o /dev/null -w "%{http_code}" "${URL}")"
if [[ "${code}" != "200" ]]; then
  echo "Failed to get 200 from health endpoint (got ${code})"
  exit 1
fi

echo "OK: health endpoint returned 200"
