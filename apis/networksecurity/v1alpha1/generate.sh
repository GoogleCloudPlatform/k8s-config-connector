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
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# We need a newer googleapis to get BackendAuthenticationConfig
PROTO_SHA="cdc919ff596e263f2cc55a9780d2f74633da1ced" 
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

# Unset SKIP_GENERATE_PROTOS so this specific script fetches the newer proto
OLD_SKIP_GENERATE_PROTOS="${SKIP_GENERATE_PROTOS:-}"
unset SKIP_GENERATE_PROTOS

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

# Restore SKIP_GENERATE_PROTOS
if [[ -n "${OLD_SKIP_GENERATE_PROTOS}" ]]; then
  export SKIP_GENERATE_PROTOS="${OLD_SKIP_GENERATE_PROTOS}"
fi

go run . generate-types \
  --service google.cloud.networksecurity.v1 \
  --api-version networksecurity.cnrm.cloud.google.com/v1alpha1 \
  --resource NetworkSecurityBackendAuthenticationConfig:BackendAuthenticationConfig \
  --resource NetworkSecurityInterceptDeployment:InterceptDeployment \
  --resource NetworkSecurityInterceptEndpointGroup:InterceptEndpointGroup \
  --resource NetworkSecurityMirroringEndpointGroup:MirroringEndpointGroup \
  --resource NetworkSecuritySecurityProfile:SecurityProfile \
  --proto-source-path ${PROTO_OUT}

cd ${REPO_ROOT}