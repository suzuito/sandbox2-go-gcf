#!/bin/sh

set -e

FILENAME_ZIP="$1.zip"

go mod download && go mod vendor
zip -r $FILENAME_ZIP ./