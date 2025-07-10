#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/kms-1-keyring/keyring.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "KMSKeyRing"
if ! grep -q "kind: KMSKeyRing" "${FILE_PATH}"; then
  echo "File does not have kind: KMSKeyRing"
  exit 1
fi

# Check for the correct settings
if ! grep -q "location: global" "${FILE_PATH}"; then
  echo "location is not set to global"
  exit 1
fi

echo "Validation successful!"
exit 0
