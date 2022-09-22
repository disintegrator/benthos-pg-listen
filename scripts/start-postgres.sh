#!/bin/sh

set -euf

docker run \
  --name listen-notify-demo \
  -d \
  --rm \
  -v "${PWD}/postgres/docker-entrypoint-initdb.d:/docker-entrypoint-initdb.d" \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -p 5432:5432 \
  postgres:14
