#!/bin/bash
#
# Copy the tests directory and run Go benchmarks

set -euo pipefail

rm -rf tests
cp -r .csse6400/tests .

go mod download

go test ./... -run '^$' -bench . -benchmem -count=1
