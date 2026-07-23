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

# Update the Discovery API JSON specification from upstream
"${REPO_ROOT}/apis/identityplatform/v1beta1/update.sh"

# Generate KRM Go types from the OpenAPI / Discovery API JSON specification, ignoring all 'kind' fields
go run "${REPO_ROOT}/dev/tools/openapi-to-krm/main.go" \
  --schema-file "${REPO_ROOT}/apis/identityplatform/v1beta1/identitytoolkit-api.json" \
  --api-version "identityplatform.cnrm.cloud.google.com/v1beta1" \
  --resource "IdentityPlatformConfig:GoogleCloudIdentitytoolkitAdminV2Config" \
  --output-file "${REPO_ROOT}/apis/identityplatform/v1beta1/types.generated.go"

cd ${REPO_ROOT}
dev/tasks/generate-crds
