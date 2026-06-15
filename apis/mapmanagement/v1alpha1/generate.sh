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

# Explicitly compile the required protos to a private, isolated pb file.
# Unset SKIP_GENERATE_PROTOS so the compilation runs even if skip is set globally.
SKIP_GENERATE_PROTOS= ./generate-proto.sh d23047f4bae1f796aff726e6e3399292d29804ef "${REPO_ROOT}/apis/mapmanagement/v1alpha1/mapmanagement.pb"

go run . generate-types \
    --service google.maps.mapmanagement.v2beta \
    --api-version mapmanagement.cnrm.cloud.google.com/v1alpha1 \
    --resource MapManagementMapConfig:MapConfig \
    --proto-source-path "${REPO_ROOT}/apis/mapmanagement/v1alpha1/mapmanagement.pb"

# Revert types.generated.go if needed, but since this is greenfield, we probably don't need to revert
# unless we already had other types.

go run . generate-mapper \
    --service google.maps.mapmanagement.v2beta \
    --api-version mapmanagement.cnrm.cloud.google.com/v1alpha1 \
    --proto-source-path "${REPO_ROOT}/apis/mapmanagement/v1alpha1/mapmanagement.pb"

cd ${REPO_ROOT}

dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/mapmanagement/
