#!/bin/sh

set -euf

docker run \
  --link listen-notify-demo \
  -it --rm \
  postgres:14 \
  psql -h listen-notify-demo -U postgres -p 5432 -w
