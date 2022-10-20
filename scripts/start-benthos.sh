#!/bin/sh

set -e

trap "exit 0" EXIT

project="projects/config.yml"

if command -v gum > /dev/null 2>&1; then
  choices="$(ls projects/*.yml)"
  project=$(gum choose --selected ${project} ${choices})
else
  echo "\`gum\` not installed. Using default project: ${project}…"
fi

exec ./bin/benthos --env-file .env -c "${project}"
