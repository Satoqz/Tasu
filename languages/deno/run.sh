#!/bin/sh
set -e

printf %s "$1" > eval.ts
deno run -A eval.ts || true
