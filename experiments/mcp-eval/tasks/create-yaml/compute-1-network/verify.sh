#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/compute-1-network/network.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "ComputeNetwork"
if ! grep -q "kind: ComputeNetwork" "${FILE_PATH}"; then
  echo "File does not have kind: ComputeNetwork"
  exit 1
fi

# Check for the correct settings
if ! grep -q "autoCreateSubnetworks: false" "${FILE_PATH}"; then
  echo "autoCreateSubnetworks is not set to false"
  exit 1
fi

echo "Validation successful!"
exit 0
