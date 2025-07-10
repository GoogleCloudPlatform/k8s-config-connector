#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/bigquery-1-dataset/dataset.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "BigQueryDataset"
if ! grep -q "kind: BigQueryDataset" "${FILE_PATH}"; then
  echo "File does not have kind: BigQueryDataset"
  exit 1
fi

# Check for the correct settings
if ! grep -q "defaultTableExpirationMs: 3600000" "${FILE_PATH}"; then
  echo "defaultTableExpirationMs is not set to 3600000"
  exit 1
fi

if ! grep -q "location: US" "${FILE_PATH}"; then
  echo "location is not set to US"
  exit 1
fi

echo "Validation successful!"
exit 0
