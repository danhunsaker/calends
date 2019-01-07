#!/usr/bin/env bash

start="$(dirname "$0")"
if [[ ${start:0:1} != '/' ]]
then
  start="$(dirname "$(pwd)/$0")"
fi
set -e

# Ensure everything is built, first, so we're testing the latest code
cd "$(dirname "${start}")"
go get -v -t ./...
cd libcalends
go build -v -o libcalends.so -buildmode=c-shared .

cd php
zephir fullclean
zephir build --dev -v || (cat compile-errors.log; false)

# Run the actual tests
cd ext
TRAVIS='' NO_INTERACTION=true php run-tests.php -P -d "extension=calends.so" || true
"${start}/failed_test_info.sh"
