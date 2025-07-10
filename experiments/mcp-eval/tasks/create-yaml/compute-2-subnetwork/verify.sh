#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/compute-2-subnetwork/subnetwork.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "ComputeSubnetwork"
if ! grep -q "kind: ComputeSubnetwork" "${FILE_PATH}"; then
  echo "File does not have kind: ComputeSubnetwork"
  exit 1
fi

# Check for the correct settings
if ! grep -q "ipCidrRange: 10.2.0.0/16" "${FILE_PATH}"; then
  echo "ipCidrRange is not set to 10.2.0.0/16"
  exit 1
fi

if ! grep -q "region: us-central1" "${FILE_PATH}"; then
  echo "region is not set to us-central1"
  exit 1
fi

if ! grep -q "name: my-network" "${FILE_PATH}"; then
  echo "networkRef name is not set to my-network"
  exit 1
fi

echo "Validation successful!"
exit 0
