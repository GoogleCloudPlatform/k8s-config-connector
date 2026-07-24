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

# Unset SKIP_GENERATE_PROTOS so this specific script fetches the newer proto
OLD_SKIP_GENERATE_PROTOS="${SKIP_GENERATE_PROTOS:-}"
unset SKIP_GENERATE_PROTOS

./generate-proto.sh 0fcabfc28371e7bab8107402eb06ad58134ee383 "${REPO_ROOT}/.build/googleapis-mapmanagement.pb"

# Restore SKIP_GENERATE_PROTOS
if [[ -n "${OLD_SKIP_GENERATE_PROTOS}" ]]; then
  export SKIP_GENERATE_PROTOS="${OLD_SKIP_GENERATE_PROTOS}"
fi

go run . generate-types \
  --proto-source-path "${REPO_ROOT}/.build/googleapis-mapmanagement.pb" \
  --service google.maps.mapmanagement.v2beta \
  --api-version mapmanagement.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource MapManagementMapConfig:MapConfig

go run . generate-mapper \
  --proto-source-path "${REPO_ROOT}/.build/googleapis-mapmanagement.pb" \
  --service google.maps.mapmanagement.v2beta \
  --api-version mapmanagement.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/mapmanagement/
