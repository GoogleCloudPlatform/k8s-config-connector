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


${CONTROLLERBUILDER} generate-types \
    --service google.cloud.aiplatform.v1,google.cloud.aiplatform.v1beta1 \
    --api-version aiplatform.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAISpecialistPool:SpecialistPool \
    --resource AIPlatformModel:Model \
    --resource VertexAIFeatureOnlineStore:FeatureOnlineStore \
    --resource VertexAIPipelineJob:PipelineJob \
    --resource VertexAITuningJob:TuningJob \
    --resource VertexAIStudy:Study \
    --resource VertexAITrainingPipeline:TrainingPipeline \
    --resource VertexAISchedule:Schedule \
    --resource AIPlatformModelMonitor:ModelMonitor

# Post-processing to resolve v1/v1beta1 collisions for types that must map to v1 for existing controllers
sed -i 's/google.cloud.aiplatform.v1beta1.ExplanationMetadata/google.cloud.aiplatform.v1.ExplanationMetadata/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.ExplanationSpec/google.cloud.aiplatform.v1.ExplanationSpec/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.ExplanationParameters/google.cloud.aiplatform.v1.ExplanationParameters/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.SampledShapleyAttribution/google.cloud.aiplatform.v1.SampledShapleyAttribution/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.IntegratedGradientsAttribution/google.cloud.aiplatform.v1.IntegratedGradientsAttribution/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.XraiAttribution/google.cloud.aiplatform.v1.XraiAttribution/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.Examples/google.cloud.aiplatform.v1.Examples/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.BlurBaselineConfig/google.cloud.aiplatform.v1.BlurBaselineConfig/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.SmoothGradConfig/google.cloud.aiplatform.v1.SmoothGradConfig/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.FeatureNoiseSigma/google.cloud.aiplatform.v1.FeatureNoiseSigma/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.EncryptionSpec/google.cloud.aiplatform.v1.EncryptionSpec/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.Presets/google.cloud.aiplatform.v1.Presets/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.GcsSource/google.cloud.aiplatform.v1.GcsSource/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.GcsDestination/google.cloud.aiplatform.v1.GcsDestination/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"
sed -i 's/google.cloud.aiplatform.v1beta1.MachineSpec/google.cloud.aiplatform.v1.MachineSpec/g' "${REPO_ROOT}/apis/aiplatform/v1alpha1/types.generated.go"

# Handled recursive self-referential fields by defining ListValue, Value, and ExplanationParameters manually in recursive_types.go

${CONTROLLERBUILDER} generate-mapper \
    --service google.cloud.aiplatform.v1,google.cloud.aiplatform.v1beta1 \
    --api-version aiplatform.cnrm.cloud.google.com/v1alpha1 \
    --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/aiplatform" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/aiplatform/
fi
