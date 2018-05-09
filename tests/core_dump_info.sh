#!/bin/bash

cd ${1:-.}

shopt -s nullglob
export LC_ALL=C
for i in core*; do
	if [ -f "$i" -a "$(file "$i" | grep -o 'core file')" ]; then
    echo -e "\n--------\nExpanding info for $i\n========"
		gdb -q $((phpenv which php || which php)2>/dev/null) "$i" <<EOF
set pagination 0
backtrace full
info registers
x/16i \$pc
thread apply all backtrace
quit
EOF
	fi
done

echo
