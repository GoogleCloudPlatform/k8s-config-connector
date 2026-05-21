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

PROTO_SHA="731d7f2ab6e4e2ea15030c95039e2cb66174d4fb" 
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

go run . generate-types \
    --service google.cloud.kms.v1 \
    --api-version kms.cnrm.cloud.google.com/v1alpha1 \
    --resource KMSAutokeyConfig:AutokeyConfig \
    --resource KMSKeyHandle:KeyHandle \
    --resource KMSEKMConnection:EkmConnection \
    --proto-source-path ${PROTO_OUT}
