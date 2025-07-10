#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/secretmanager-1-from-k8s/secret.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "SecretManagerSecret"
if ! grep -q "kind: SecretManagerSecret" "${FILE_PATH}"; then
  echo "File does not have kind: SecretManagerSecret"
  exit 1
fi

# Check if the secret data is sourced from the correct k8s secret
if ! grep -q "name: my-k8s-secret" "${FILE_PATH}"; then
  echo "Secret is not sourced from 'my-k8s-secret'"
  exit 1
fi

if ! grep -q "key: secret-data" "${FILE_PATH}"; then
  echo "Secret is not sourced from key 'secret-data'"
  exit 1
fi

echo "Validation successful!"
exit 0
