#!/bin/sh
set -e

deno eval "$1" || true