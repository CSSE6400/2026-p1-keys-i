#!/bin/bash
#
# Copy the tests directory and run Go unit tests
cp -r .csse6400/tests .

go mod download

go test ./... -count=1 -v
