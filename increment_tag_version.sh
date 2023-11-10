#!/bin/sh

set -e

TAG=${1}

REGEXP=^v[0-9]+$

if ! [[ $TAG =~ ${REGEXP} ]]; then
  >&2 echo "Invalid tag ${TAG}. Must be ${REGEXP}"
  exit 1
fi
N=`echo $TAG | cut -d v -f2`
N=$((N+1))

echo $N