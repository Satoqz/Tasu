#!/bin/sh
set -e

printf %s "$1" > eval.ts
deno run eval.ts || true
