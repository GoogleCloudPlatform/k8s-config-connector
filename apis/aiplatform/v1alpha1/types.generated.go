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


// +kcc:proto=google.cloud.aiplatform.v1.Annotation
type Annotation struct {

	// Required. Google Cloud Storage URI points to a YAML file describing
	//  [payload][google.cloud.aiplatform.v1.Annotation.payload]. The schema is
	//  defined as an [OpenAPI 3.0.2 Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  The schema files that can be used here are found in
	//  gs://google-cloud-aiplatform/schema/dataset/annotation/, note that the
	//  chosen schema must be consistent with the parent Dataset's
	//  [metadata][google.cloud.aiplatform.v1.Dataset.metadata_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.payload_schema_uri
	PayloadSchemaURI *string `json:"payloadSchemaURI,omitempty"`

	// Required. The schema of the payload can be found in
	//  [payload_schema][google.cloud.aiplatform.v1.Annotation.payload_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.payload
	Payload *Value `json:"payload,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	//  Annotations.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Annotation(System
	//  labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable. Following system labels exist for each Annotation:
	//
	//  * "aiplatform.googleapis.com/annotation_set_name":
	//    optional, name of the UI's annotation set this Annotation belongs to.
	//    If not set, the Annotation is not visible in the UI.
	//
	//  * "aiplatform.googleapis.com/payload_schema":
	//    output only, its value is the
	//    [payload_schema's][google.cloud.aiplatform.v1.Annotation.payload_schema_uri]
	//    title.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.UserActionReference
type UserActionReference struct {
	// For API calls that return a long running operation.
	//  Resource name of the long running operation.
	//  Format:
	//  `projects/{project}/locations/{location}/operations/{operation}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.UserActionReference.operation
	Operation *string `json:"operation,omitempty"`

	// For API calls that start a LabelingJob.
	//  Resource name of the LabelingJob.
	//  Format:
	//  `projects/{project}/locations/{location}/dataLabelingJobs/{data_labeling_job}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.UserActionReference.data_labeling_job
	DataLabelingJob *string `json:"dataLabelingJob,omitempty"`

	// The method name of the API RPC call. For example,
	//  "/google.cloud.aiplatform.{apiVersion}.DatasetService.CreateDataset"
	// +kcc:proto:field=google.cloud.aiplatform.v1.UserActionReference.method
	Method *string `json:"method,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Annotation
type AnnotationObservedState struct {
	// Output only. Resource name of the Annotation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this Annotation was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Annotation was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The source of the Annotation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Annotation.annotation_source
	AnnotationSource *UserActionReference `json:"annotationSource,omitempty"`
}
