#!/bin/sh

set -e

trap "exit 0" EXIT

dir=$(CDPATH="" cd -- "$(dirname -- "$0")" && pwd)

while true; do
  "${dir}/insert-item.sh"
  sleep 0.1
done
