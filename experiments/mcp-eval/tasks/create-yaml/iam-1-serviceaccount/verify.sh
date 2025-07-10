#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/iam-1-serviceaccount/sa.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "IAMServiceAccount"
if ! grep -q "kind: IAMServiceAccount" "${FILE_PATH}"; then
  echo "File does not have kind: IAMServiceAccount"
  exit 1
fi

# Check for the correct settings
if ! grep -q "displayName: My Service Account" "${FILE_PATH}"; then
  echo "displayName is not set to My Service Account"
  exit 1
fi

echo "Validation successful!"
exit 0
