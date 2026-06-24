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

go run . generate-types \
  --service google.cloud.aiplatform.v1 \
  --api-version aiplatform.cnrm.cloud.google.com/v1alpha1 \
  --resource AIPlatformModel:Model \
  --resource VertexAIHyperparameterTuningJob:HyperparameterTuningJob \
  --resource VertexAISpecialistPool:SpecialistPool

cd ${REPO_ROOT}

# Inline pruning and acronym correction using a quoted heredoc (EOF) to avoid shell substitution
python3 << 'EOF'
import re

with open('apis/aiplatform/v1alpha1/types.generated.go', 'r') as f:
    code = f.read()

# Replace StructValue apiextensionsv1.JSON with map[string]string
code = re.sub(r'StructValue\s+apiextensionsv1\.JSON\s+`json:"structValue,omitempty"`', r'StructValue map[string]string `json:"structValue,omitempty"`', code)

# Comment out ListValue field in Value
code = re.sub(r'(// \+kcc:proto:field=google\.protobuf\.Value\.list_value\n\s+)(ListValue\s+\*ListValue\s+`json:"listValue,omitempty"`\n)', r'\1// ListValue refers to Value struct and vice versa, causing it to be recursive.\n\t// The recursive structure is causing CRD generation failure\n\t// \2', code)

# Comment out OutputIndices in ExplanationParameters
code = re.sub(r'(// \+kcc:proto:field=google\.cloud\.aiplatform\.v1\.ExplanationParameters\.output_indices\n\s+)(OutputIndices\s+\*ListValue\s+`json:"outputIndices,omitempty"`\n)', r'\1// ListValue struct refers to Value struct and vice versa, causing it to be recursive.\n\t// The recursive structure is causing CRD generation failure\n\t// \2', code)

# Comment out ParameterSpec in StudySpec_ParameterSpec_ConditionalParameterSpec
code = re.sub(r'(// \+kcc:proto:field=google\.cloud\.aiplatform\.v1\.StudySpec\.ParameterSpec\.ConditionalParameterSpec\.parameter_spec\n\s+)(ParameterSpec\s+\*StudySpec_ParameterSpec\s+`json:"parameterSpec,omitempty"`\n)', r'\1// ParameterSpec is a recursive type (StudySpec_ParameterSpec contains StudySpec_ParameterSpec_ConditionalParameterSpec which contains StudySpec_ParameterSpec).\n\t// controller-gen doesn\'t support recursive types and omits the type, causing fielddesc panic.\n\t// \2', code)

# Acronym corrections for WorkerPoolSpec -> pythonPackageSpec -> packageUris
code = code.replace('packageUris,omitempty', 'packageURIs,omitempty')
code = code.replace('PackageUris ', 'PackageURIs ')

# Acronym corrections for TrialObservedState -> webAccessUris
code = code.replace('webAccessUris,omitempty', 'webAccessURIs,omitempty')
code = code.replace('WebAccessUris ', 'WebAccessURIs ')

with open('apis/aiplatform/v1alpha1/types.generated.go', 'w') as f:
    f.write(code)
EOF

cd ${REPO_ROOT}/dev/tools/controllerbuilder
go run . generate-mapper \
    --service google.cloud.aiplatform.v1 \
    --api-version aiplatform.cnrm.cloud.google.com/v1alpha1 \
    --include-skipped-output

cd ${REPO_ROOT}

dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/aiplatform/
