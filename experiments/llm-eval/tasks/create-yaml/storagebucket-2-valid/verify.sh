#!/bin/bash
# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -o errexit
set -o nounset
set -o pipefail

# Get the top-level directory of the git repository
TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/storagebucket-2-valid/storagebucket-valid.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "StorageBucket"
if ! grep -q "kind: StorageBucket" "${FILE_PATH}"; then
  echo "File does not have kind: StorageBucket"
  exit 1
fi

# Check if the file has the annotation "cnrm.cloud.google.com/project-id"
if grep -q 'cnrm.cloud.google.com/project-id' "${FILE_PATH}"; then
  echo "Found project-id annotation"
  exit 0
fi

echo "Validation failed: Missing project-id annotation"
exit 1
