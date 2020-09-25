#!/bin/sh
set -e

printf %s "$1" > program.nim
nim build c --run ./program.nim || true