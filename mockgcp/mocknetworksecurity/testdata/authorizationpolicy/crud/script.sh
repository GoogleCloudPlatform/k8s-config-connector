#!/bin/bash
set -e
envsubst < {{.TestDirectory}}/script.tmpl.yaml | while read -r line; do
  eval $line
done
