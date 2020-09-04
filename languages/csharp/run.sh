#!/bin/sh
set -e

printf %s "$1" > program.cs
csc program.cs -nologo >/dev/null && mono program.exe || true
