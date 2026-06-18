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

# Generate v1beta1 types
go run . generate-types \
    --service google.cloud.aiplatform.v1beta1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAIFeaturestore:Featurestore \
    --resource VertexAIMetadataStore:MetadataStore \
    --resource VertexAIDeploymentResourcePool:DeploymentResourcePool \
    --resource VertexAIExampleStore:ExampleStore \
    --resource VertexAIFeatureGroup:FeatureGroup

mv ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go ${REPO_ROOT}/apis/vertexai/v1alpha1/v1beta1_types.generated.go

# Generate v1 types
go run . generate-types \
    --service google.cloud.aiplatform.v1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAIDataLabelingJob:DataLabelingJob \
    --resource VertexAICachedContent:CachedContent

mv ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go ${REPO_ROOT}/apis/vertexai/v1alpha1/v1_types.generated.go

# Deduplicate overlapping generated types and handwritten reference types
python3 ${REPO_ROOT}/apis/vertexai/v1alpha1/deduplicate_types.py

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/vertexai/
