#!/bin/bash
#
# Validate that the repository has the following structure:
# -- README.md
# -- go.mod
# -- cmd
#    | -- api
#         | -- main.go
# -- internal
#    | -- todo
#         | -- routes.go
#         | -- routes
#              | -- health.go
#              | -- todo.go

set -euo pipefail

failed=0

# Required root files
for file in README.md go.mod; do
  if [[ ! -f "$file" ]]; then
    echo "FAIL: Missing $file"
    failed=1
  fi
done

# Required directories
for dir in cmd cmd/api internal internal/todo internal/todo/routes; do
  if [[ ! -d "$dir" ]]; then
    echo "FAIL: Missing $dir directory"
    failed=1
  fi
done

# Required Go source files
for file in cmd/api/main.go internal/todo/routes.go internal/todo/routes/health.go internal/todo/routes/todo.go; do
  if [[ ! -f "$file" ]]; then
    echo "FAIL: Missing $file"
    failed=1
  fi
done

if [[ $failed -eq 1 ]]; then
  echo "Repository structure is not valid. Please fix the errors above."
  exit 1
fi
