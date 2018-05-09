#!/usr/bin/env bash

for diff in $(find tests/ -iname '*.diff' | sort -h)
do
  FAILED=true
  f=${diff/%.diff/}
  echo -e "\n--------\n${f//tests\//Test }\n========"
  cat "${f}.php"
  echo "--RESULT--"
  cat "${f}.diff"
  echo
done
[[ "$FAILED" != "true" ]]
