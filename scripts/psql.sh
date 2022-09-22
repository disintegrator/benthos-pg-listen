#!/bin/sh

set -euf

cmd="psql -h listen-notify-demo -U postgres -p 5432 -w"

docker run \
  --link listen-notify-demo \
  -it --rm \
  postgres:14 \
  "${cmd}"
