#!/usr/bin/env bash

start="$(dirname "$0")"
if [[ ${start:0:1} != '/' ]]
then
  start="$(dirname "$(pwd)/$0")"
fi
set -e

# Ensure everything is built, first, so we're testing the latest code
cd "$(dirname "${start}")"
# GO111MODULE=on go get -v ./...
go get -v ./...
cd libcalends
go build -v -o libcalends.so -buildmode=c-shared .

cd php
which -a zephir || echo "ZEPHIR NOT FOUND!!????"
zephir fullclean
zephir build --dev || (cat compile-errors.log; false)

# Run the actual tests
cd ext
TRAVIS='' NO_INTERACTION=true php run-tests.php -P -d "extension=calends.so" || true
"${start}/failed_test_info.sh"
