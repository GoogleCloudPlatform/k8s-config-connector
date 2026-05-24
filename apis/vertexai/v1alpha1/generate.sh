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

./generate-proto.sh

# Backup original types.generated.go and zz_generated.deepcopy.go
cp ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go.bak
cp ${REPO_ROOT}/apis/vertexai/v1alpha1/zz_generated.deepcopy.go ${REPO_ROOT}/apis/vertexai/v1alpha1/zz_generated.deepcopy.go.bak

# Generate only VertexAIIndexEndpoint types
go run . generate-types \
    --service google.cloud.aiplatform.v1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAIIndexEndpoint:IndexEndpoint

mv ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go ${REPO_ROOT}/apis/vertexai/v1alpha1/vertexaiindexendpoint_types.generated.go

# Restore the original files from backup
mv ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go.bak ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go
mv ${REPO_ROOT}/apis/vertexai/v1alpha1/zz_generated.deepcopy.go.bak ${REPO_ROOT}/apis/vertexai/v1alpha1/zz_generated.deepcopy.go

# Run deduplication to strip any types already declared in the restored types.generated.go
python3 ${REPO_ROOT}/dev/tasks/dedup_vertexai_types.py

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/vertexai/
