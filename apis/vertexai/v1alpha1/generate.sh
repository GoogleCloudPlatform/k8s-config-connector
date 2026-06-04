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
    --service google.cloud.aiplatform.v1beta1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAIFeaturestore:Featurestore \
    --resource VertexAIMetadataStore:MetadataStore \
    --resource VertexAIDeploymentResourcePool:DeploymentResourcePool \
    --resource VertexAIExampleStore:ExampleStore \
    --resource VertexAIModelMonitor:ModelMonitor

mv ${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go ${REPO_ROOT}/apis/vertexai/v1alpha1/types_v1beta1.generated.go

# go run . generate-mapper     --service google.cloud.aiplatform.v1beta1     --api-version vertexai.cnrm.cloud.google.com/v1alpha1

go run . generate-types \
    --service google.cloud.aiplatform.v1 \
    --api-version vertexai.cnrm.cloud.google.com/v1alpha1 \
    --resource VertexAIDataLabelingJob:DataLabelingJob

# Comment out duplicate ListValue and Value types in types.generated.go to avoid redeclaration issues with types_v1beta1.generated.go
python3 -c "
import sys
path = '${REPO_ROOT}/apis/vertexai/v1alpha1/types.generated.go'
with open(path, 'r') as f:
    content = f.read()
old_block = '''// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value \`json:\"values,omitempty\"\`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string \`json:\"nullValue,omitempty\"\`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 \`json:\"numberValue,omitempty\"\`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string \`json:\"stringValue,omitempty\"\`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool \`json:\"boolValue,omitempty\"\`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue apiextensionsv1.JSON \`json:\"structValue,omitempty\"\`

	// Represents a repeated \`Value\`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue \`json:\"listValue,omitempty\"\`
}'''
new_block = '''/* unreachable type ListValue
// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value \`json:\"values,omitempty\"\`
}
*/

/* unreachable type Value
// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string \`json:\"nullValue,omitempty\"\`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 \`json:\"numberValue,omitempty\"\`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string \`json:\"stringValue,omitempty\"\`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool \`json:\"boolValue,omitempty\"\`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue apiextensionsv1.JSON \`json:\"structValue,omitempty\"\`

	// Represents a repeated \`Value\`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue \`json:\"listValue,omitempty\"\`
}
*/'''
if old_block in content:
    content = content.replace(old_block, new_block)
    with open(path, 'w') as f:
        f.write(content)
        print('Successfully commented out duplicate ListValue and Value types')
else:
    print('Duplicate ListValue and Value block not found in types.generated.go')
"

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/vertexai/