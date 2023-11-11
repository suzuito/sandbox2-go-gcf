#!/bin/sh

set -e

env=$1
region=$2
version=$3

curl -L \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer ${GH_TOKEN}" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/suzuito/terraform/dispatches \
  -d "{\"event_type\":\"update_crawler_version\",\"client_payload\":{"env":\"${env}\","region":\"${region}\","version":\"${version}\"}}"