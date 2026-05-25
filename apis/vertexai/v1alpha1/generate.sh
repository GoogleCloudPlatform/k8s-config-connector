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

# Types for v1alpha1
go run . generate-types \
    --service google.cloud.aiplatform.v1beta1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAIFeaturestore:Featurestore \
    --resource VertexAIMetadataStore:MetadataStore \
    --resource VertexAIDeploymentResourcePool:DeploymentResourcePool \
    --resource VertexAIExampleStore:ExampleStore \
    --resource VertexAICustomJob:CustomJob \
    --resource VertexAIDataLabelingJob:DataLabelingJob

# CustomJob Mapper (isolated sub-package)
go run . generate-mapper \
    --service google.cloud.aiplatform.v1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --output-dir ${REPO_ROOT}/pkg/controller/direct/vertexai/customjob/

# Fix up the generated file for customjob
mv ${REPO_ROOT}/pkg/controller/direct/vertexai/customjob/vertexai/mapper.generated.go ${REPO_ROOT}/pkg/controller/direct/vertexai/customjob/mapper.generated.go
rmdir ${REPO_ROOT}/pkg/controller/direct/vertexai/customjob/vertexai
sed -i "s/package vertexai/package customjob/" ${REPO_ROOT}/pkg/controller/direct/vertexai/customjob/mapper.generated.go

# Original mapper run for existing resources
go run . generate-mapper \
    --service google.cloud.aiplatform.v1beta1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1

# Fix up the generated mappers
cd ${REPO_ROOT}
python3 apis/vertexai/v1alpha1/scripts/fix_vertexai_mappers.py pkg/controller/direct/vertexai/mapper.generated.go pkg/controller/direct/vertexai/customjob/mapper.generated.go
python3 apis/vertexai/v1alpha1/scripts/fix_redeclarations.py pkg/controller/direct/vertexai/mapper.generated.go pkg/controller/direct/vertexai/customjob/mapper.generated.go
python3 apis/vertexai/v1alpha1/scripts/fix_customjob_mappers.py
python3 apis/vertexai/v1alpha1/scripts/prune_customjob_mapper.py

dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/vertexai/
