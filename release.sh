#!/bin/sh

set -e

LATEST_VERSION=$(gh release list --limit=1 --exclude-drafts --exclude-pre-releases | grep Latest | awk '{print $3}')
NEXT_VERSION=v`bash increment_tag_version.sh ${LATEST_VERSION}`
gh release create ${NEXT_VERSION} --title="${NEXT_VERSION}" --generate-notes --target=main
