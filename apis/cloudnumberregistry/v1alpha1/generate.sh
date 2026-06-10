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

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# Build the custom .pb for numberregistry using fe9f668e59b5448d27564d2b89b5aed97b74f8d7 version of googleapis if it doesn't exist
if [ ! -f "${REPO_ROOT}/.build/googleapis-numberregistry.pb" ]; then
  # Unset SKIP_GENERATE_PROTOS so this specific script fetches the newer/custom proto
  OLD_SKIP_GENERATE_PROTOS="${SKIP_GENERATE_PROTOS:-}"
  unset SKIP_GENERATE_PROTOS

  ./generate-proto.sh fe9f668e59b5448d27564d2b89b5aed97b74f8d7 "${REPO_ROOT}/.build/googleapis-numberregistry.pb"
  # Restore the default googleapis.pb and checkout
  ./generate-proto.sh

  # Restore SKIP_GENERATE_PROTOS
  if [[ -n "${OLD_SKIP_GENERATE_PROTOS}" ]]; then
    export SKIP_GENERATE_PROTOS="${OLD_SKIP_GENERATE_PROTOS}"
  fi
fi

go run . generate-types \
  --proto-source-path "${REPO_ROOT}/.build/googleapis-numberregistry.pb" \
  --service google.cloud.numberregistry.v1alpha \
  --api-version cloudnumberregistry.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource CloudNumberRegistryCustomRange:CustomRange

# generate-mapper skipped in Step 1

cd ${REPO_ROOT}
dev/tasks/generate-crds

# goimports on direct/cloudnumberregistry skipped in Step 1
