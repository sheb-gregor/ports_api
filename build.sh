#!/bin/bash

VERSION=1.1.0
COMMIT=$(git rev-parse --short HEAD)
TAG=$(git rev-parse --abbrev-ref HEAD)
SERVICE_PATH=ports_api

diff_status=$(git diff-index HEAD)
if [ "$diff_status" != "" ]; then
  COMMIT=$COMMIT-dirty.
fi

if [ "$1" != "" ]; then
  go build -o "${1}" -ldflags "-X $SERVICE_PATH/config.version=$VERSION -X $SERVICE_PATH/config.build=$COMMIT -X $SERVICE_PATH/config.tag=$TAG" .
  exit 0
fi

go build -ldflags "-X $SERVICE_PATH/config.version=$VERSION -X $SERVICE_PATH/config.build=$COMMIT -X $SERVICE_PATH/config.tag=$TAG" .


