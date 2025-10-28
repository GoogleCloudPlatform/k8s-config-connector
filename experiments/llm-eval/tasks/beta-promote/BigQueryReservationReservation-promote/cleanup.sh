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

set -o nounset
set -o errexit
set -o pipefail

REPO_DIR="$1"
DEST_PATH="$2"

if [ -d "${REPO_DIR}" ]; then
  echo "Moving ${REPO_DIR} to ${DEST_PATH}"
  mv "${REPO_DIR}" "${DEST_PATH}"
else
  echo "Directory ${REPO_DIR} not found, skipping cleanup."
fi

exit 0
