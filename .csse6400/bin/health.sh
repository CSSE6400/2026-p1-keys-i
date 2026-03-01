#!/bin/bash
#
# Check that the health endpoint is returning 200

# Start go app
go mod download
go run ./cmd/api -p 6400 &
error=$?
pid=$!
if [[ $error -ne 0 ]]; then
    echo "Failed to start go api"
    exit 1
fi

# Wait for go api to start
sleep 5

# Check that the health endpoint is returning 200
curl -s -o /dev/null -w "%{http_code}" http://localhost:6400/api/v1/health | grep 200
error=$?
if [[ $error -ne 0 ]]; then
    echo "Failed to get 200 from health endpoint"
    exit 1
fi

# Kill go app
kill $pid
