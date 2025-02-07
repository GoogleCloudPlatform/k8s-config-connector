// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.aiplatform.v1.DatasetVersion
type DatasetVersion struct {

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.etag
	Etag *string `json:"etag,omitempty"`

	// The user-defined name of the DatasetVersion.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DatasetVersion
type DatasetVersionObservedState struct {
	// Output only. Identifier. The resource name of the DatasetVersion.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this DatasetVersion was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this DatasetVersion was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Name of the associated BigQuery dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.big_query_dataset_name
	BigQueryDatasetName *string `json:"bigQueryDatasetName,omitempty"`

	// Required. Output only. Additional information about the DatasetVersion.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.metadata
	Metadata *Value `json:"metadata,omitempty"`

	// Output only. Reference to the public base model last used by the dataset
	//  version. Only set for prompt dataset versions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.model_reference
	ModelReference *string `json:"modelReference,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
