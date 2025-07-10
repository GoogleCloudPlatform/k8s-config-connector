#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/sql-1-instance/instance.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "SQLInstance"
if ! grep -q "kind: SQLInstance" "${FILE_PATH}"; then
  echo "File does not have kind: SQLInstance"
  exit 1
fi

# Check for the correct settings
if ! grep -q "databaseVersion: MYSQL_8_0" "${FILE_PATH}"; then
  echo "databaseVersion is not set to MYSQL_8_0"
  exit 1
fi

if ! grep -q "region: us-central1" "${FILE_PATH}"; then
  echo "region is not set to us-central1"
  exit 1
fi

echo "Validation successful!"
exit 0
