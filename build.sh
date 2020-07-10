#!/bin/bash

if [ "$1" == "" ]; then
  echo ERROR: App name is not provided
  exit 1
fi

VERSION=1.1.0
COMMIT=$(git rev-parse --short HEAD)
TAG=$(git rev-parse --abbrev-ref HEAD)


diff_status=$(git diff-index HEAD)
if [ "$diff_status" != "" ]; then
  COMMIT=$COMMIT-dirty.
fi

PKG=ports_api/internal/config
LD_FLAG="-X ${PKG}.version=$VERSION -X ${PKG}.build=$COMMIT -X ${PKG}.tag=$TAG"


if [ "$2" != "" ]; then
  go build -o "${2}" -ldflags "$LD_FLAG" ./apps/${1}/.
  exit 0
fi

go build -ldflags "$LD_FLAG" ./apps/${1}/.


