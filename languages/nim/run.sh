#!/bin/sh
set -e

printf %s "$1" > program.nim
nim c -o:program --nimcache:./ --debuginfo:off --verbosity:0 --hints:off ./program.nim
./program || true
