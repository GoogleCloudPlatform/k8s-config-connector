#!/bin/bash
# Copyright 2025 Google LLC
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

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/spanner-1-instance/instance.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "SpannerInstance"
if ! grep -q "kind: SpannerInstance" "${FILE_PATH}"; then
  echo "File does not have kind: SpannerInstance"
  exit 1
fi

# Check for the correct settings
if ! grep -q "config: regional-us-central1" "${FILE_PATH}"; then
  echo "config is not set to regional-us-central1"
  exit 1
fi

if ! grep -q "displayName: My Spanner Instance" "${FILE_PATH}"; then
  echo "displayName is not set to My Spanner Instance"
  exit 1
fi

echo "Validation successful!"
exit 0
