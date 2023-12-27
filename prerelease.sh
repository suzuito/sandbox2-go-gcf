#!/bin/sh

set -e

target=$1

NEXT_VERSION=vv$(date +%Y%m%d%H%M%S)
gh release create ${NEXT_VERSION} --title="${NEXT_VERSION}" --generate-notes --target=${target} -p
