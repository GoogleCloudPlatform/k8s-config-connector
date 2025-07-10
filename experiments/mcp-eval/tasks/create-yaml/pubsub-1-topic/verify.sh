#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/pubsub-1-topic/topic.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "PubSubTopic"
if ! grep -q "kind: PubSubTopic" "${FILE_PATH}"; then
  echo "File does not have kind: PubSubTopic"
  exit 1
fi

echo "Validation successful!"
exit 0
