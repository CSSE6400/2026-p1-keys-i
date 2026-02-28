#!/bin/bash
#
# Check that the health endpoint is returning 200

# Start go app
go mod download
go build -o /tmp/csse6400_api ./cmd/api
/tmp/csse6400_api -p 6400 >/dev/null 2>&1 &
error=$?
pid=$!
if [[ $error -ne 0 ]]; then
    echo "Failed to start flask app"
    exit 1
fi

# Wait for go app to start
sleep 5

# Check that the health endpoint is returning 200
curl -s -o /dev/null -w "%{http_code}" http://localhost:6400/api/v1/health | grep 200
error=$?
if [[ $error -ne 0 ]]; then
    echo "Failed to get 200 from health endpoint"
    exit 1
fi

# Kill go app
kill $pid 2>/dev/null || true
