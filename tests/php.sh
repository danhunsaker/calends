#!/usr/bin/env bash

start=$(dirname $(readlink -f $0))
set -e

# Ensure everything is built, first, so we're testing the latest code
cd libcalends
go build -v -o libcalends.so -buildmode=c-shared .

cd php
zephir fullclean
zephir builddev -v || (cat compile-errors.log; false)

# Run the actual tests
cd ext
TRAVIS='' NO_INTERACTION=true php run-tests.php -P -d "extension=calends.so" || ${start}/failed_test_info.sh
