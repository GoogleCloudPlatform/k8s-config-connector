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

# Pin a googleapis SHA that contains the google.cloud.hypercomputecluster.v1 service definition
PROTO_SHA="b8486a2f44f15dc578a9dc1e17b144253079d5c1"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

go run . generate-types \
  --service google.cloud.hypercomputecluster.v1 \
  --api-version hypercomputecluster.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource HypercomputeClusterCluster:Cluster \
  --proto-source-path ${PROTO_OUT}

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/hypercomputecluster" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/hypercomputecluster/
fi
