#!/bin/bash
# Copyright 2026 Google LLC
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

# This script finds the version of golang.org/x/tools in go.mod and sets it to GOLANG_X_TOOLS_VERSION.

REPO_ROOT="$(git rev-parse --show-toplevel)"
export GOLANG_X_TOOLS_VERSION=$(awk '/^[[:space:]]*golang\.org\/x\/tools[[:space:]]+/{print $2}' "${REPO_ROOT}/go.mod")
if [[ -z "${GOLANG_X_TOOLS_VERSION}" ]]; then
  echo "ERROR: Could not determine GOLANG_X_TOOLS_VERSION from ${REPO_ROOT}/go.mod" >&2
  exit 1
fi
