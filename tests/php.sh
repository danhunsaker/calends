#!/usr/bin/env bash

php=${1}
ext=${2:-so}

start="$(dirname "$0")"
if [[ ${start:0:1} != '/' ]]
then
  start="$(dirname "$(pwd)/$0")"
fi
set -ex

# Ensure everything is built, first, so we're testing the latest code
cd "$(dirname "${start}")"
go get -v ./...
cd libcalends
go build -v -o libcalends.${ext} -buildmode=c-shared .

# Run the actual tests
cd php

wget "https://www.php.net/distributions/php-${php}.tar.gz" -O - | tar xzv --strip-components 1 php-${php}/run-tests.php

echo "#define FFI_SCOPE \"CALENDS\"" > Calends.h
echo "#define FFI_LIB \"$(dirname $(pwd))/libcalends.${ext}\"" >> Calends.h
cpp -P -D"__attribute__(ARGS)=" ../libcalends.h | egrep -v '_Complex|^static inline .*\{$|return |^\}$' >> Calends.h

composer install

NO_INTERACTION=true php run-tests.php -P -d "ffi.preload=Calends.h" || true

# rm run-tests.php Calends.h

"${start}/failed_test_info.sh"
