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

CONTROLLERBUILDER="${CONTROLLERBUILDER:-}"
if [[ -z "${CONTROLLERBUILDER}" ]]; then
  if [[ -x "${REPO_ROOT}/bin/controllerbuilder" ]]; then
    CONTROLLERBUILDER="${REPO_ROOT}/bin/controllerbuilder"
  else
    CONTROLLERBUILDER="go run ${REPO_ROOT}/dev/tools/controllerbuilder"
  fi
fi
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder
./generate-proto.sh

# --- v1alpha1 ---
${CONTROLLERBUILDER} generate-types \
    --service google.cloud.aiplatform.v1beta1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAIFeaturestore:Featurestore \
    --resource VertexAIMetadataStore:MetadataStore \
    --resource VertexAIDeploymentResourcePool:DeploymentResourcePool \
    --resource VertexAIExampleStore:ExampleStore \
    --resource VertexAIFeatureGroup:FeatureGroup \
    --resource VertexAIDataLabelingJob:DataLabelingJob \
    --resource VertexAICustomJob:CustomJob \
    --resource VertexAITensorboard:Tensorboard \
    --resource VertexAITensorboardExperiment:TensorboardExperiment

# --- v1beta1 ---
${CONTROLLERBUILDER} generate-types \
    --service google.cloud.aiplatform.v1beta1 \
    --api-version vertexai.cnrm.cloud.google.com/v1beta1 \
    --include-skipped-output \
    --resource VertexAIMetadataStore:MetadataStore \
    --resource VertexAIDataset:Dataset

${CONTROLLERBUILDER} generate-mapper \
    --service google.cloud.aiplatform.v1beta1 \
    --api-version vertexai.cnrm.cloud.google.com/v1beta1 \
    --include-skipped-output \
    --multiversion

cd ${REPO_ROOT}
sed -i 's/apiextensionsv1.JSON_v1alpha1_/JSON_v1alpha1_/g' "${REPO_ROOT}/pkg/controller/direct/vertexai/mapper.generated.go"

dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/vertexai" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/vertexai/
fi
