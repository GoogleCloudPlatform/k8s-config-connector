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
  --resource VertexAISpecialistPool:SpecialistPool \
  --resource AIPlatformModel:Model \
  --resource VertexAIPersistentResource:PersistentResource

cd ${REPO_ROOT}

# Inject the necessary apiextensionsv1 import and ListValue definition into types.generated.go
python3 -c '
import os

filepath = "apis/aiplatform/v1alpha1/types.generated.go"
with open(filepath, "r") as f:
    content = f.read()

# Add apiextensionsv1 import if not present
if "import apiextensionsv1" not in content and "apiextensions/v1" not in content:
    content = content.replace(
        "package v1alpha1",
        "package v1alpha1\n\nimport (\n\tapiextensionsv1 \"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1\"\n)"
    )

# Add ListValue struct definition before Value struct if not present
if "type ListValue struct" not in content:
    list_value_def = """// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
// +kcc:proto=google.protobuf.Value"""
    content = content.replace("// +kcc:proto=google.protobuf.Value", list_value_def)

# Comment out recursive ListValue field inside Value struct to break OpenAPI generation loop
content = content.replace(
    "\tListValue *ListValue `json:\"listValue,omitempty\"`\\n}",
    "\t// ListValue *ListValue `json:\"listValue,omitempty\"`\\n}"
)
content = content.replace(
    "\tListValue *ListValue `json:\"listValue,omitempty\"`\n}",
    "\t// ListValue *ListValue `json:\"listValue,omitempty\"`\n}"
)

with open(filepath, "w") as f:
    f.write(content)
'

cd ${REPO_ROOT}/dev/tools/controllerbuilder

go run . generate-mapper \
  --service google.cloud.aiplatform.v1 \
  --api-version aiplatform.cnrm.cloud.google.com/v1alpha1 \
  --api-dir ${REPO_ROOT}/apis/aiplatform/v1alpha1 \
  --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1

cd ${REPO_ROOT}

dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/aiplatform/
