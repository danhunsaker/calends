#!/usr/bin/env bash

set -e

go get -v ./...
go mod vendor

echo "mode: atomic" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -v -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        grep -v '^mode:' profile.out >> coverage.txt || true
        rm profile.out
    fi
done
