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


// +kcc:proto=google.cloud.aiplatform.v1.Dataset
type Dataset struct {

	// Required. The user-defined name of the Dataset.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the Dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.description
	Description *string `json:"description,omitempty"`

	// Required. Points to a YAML file stored on Google Cloud Storage describing
	//  additional information about the Dataset. The schema is defined as an
	//  OpenAPI 3.0.2 Schema Object. The schema files that can be used here are
	//  found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.metadata_schema_uri
	MetadataSchemaURI *string `json:"metadataSchemaURI,omitempty"`

	// Required. Additional information about the Dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.metadata
	Metadata *Value `json:"metadata,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Datasets.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Dataset (System
	//  labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable. Following system labels exist for each Dataset:
	//
	//  * "aiplatform.googleapis.com/dataset_metadata_schema": output only, its
	//    value is the
	//    [metadata_schema's][google.cloud.aiplatform.v1.Dataset.metadata_schema_uri]
	//    title.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.labels
	Labels map[string]string `json:"labels,omitempty"`

	// All SavedQueries belong to the Dataset will be returned in List/Get
	//  Dataset response. The annotation_specs field
	//  will not be populated except for UI cases which will only use
	//  [annotation_spec_count][google.cloud.aiplatform.v1.SavedQuery.annotation_spec_count].
	//  In CreateDataset request, a SavedQuery is created together if
	//  this field is set, up to one SavedQuery can be set in CreateDatasetRequest.
	//  The SavedQuery should not contain any AnnotationSpec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.saved_queries
	SavedQueries []SavedQuery `json:"savedQueries,omitempty"`

	// Customer-managed encryption key spec for a Dataset. If set, this Dataset
	//  and all sub-resources of this Dataset will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Reference to the public base model last used by the dataset. Only
	//  set for prompt datasets.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.model_reference
	ModelReference *string `json:"modelReference,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SavedQuery
type SavedQuery struct {

	// Required. The user-defined name of the SavedQuery.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Some additional information about the SavedQuery.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.metadata
	Metadata *Value `json:"metadata,omitempty"`

	// Required. Problem type of the SavedQuery.
	//  Allowed values:
	//
	//  * IMAGE_CLASSIFICATION_SINGLE_LABEL
	//  * IMAGE_CLASSIFICATION_MULTI_LABEL
	//  * IMAGE_BOUNDING_POLY
	//  * IMAGE_BOUNDING_BOX
	//  * TEXT_CLASSIFICATION_SINGLE_LABEL
	//  * TEXT_CLASSIFICATION_MULTI_LABEL
	//  * TEXT_EXTRACTION
	//  * TEXT_SENTIMENT
	//  * VIDEO_CLASSIFICATION
	//  * VIDEO_OBJECT_TRACKING
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.problem_type
	ProblemType *string `json:"problemType,omitempty"`

	// Used to perform a consistent read-modify-write update. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.etag
	Etag *string `json:"etag,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Dataset
type DatasetObservedState struct {
	// Output only. Identifier. The resource name of the Dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.name
	Name *string `json:"name,omitempty"`

	// Output only. The number of DataItems in this Dataset. Only apply for
	//  non-structured Dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.data_item_count
	DataItemCount *int64 `json:"dataItemCount,omitempty"`

	// Output only. Timestamp when this Dataset was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Dataset was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// All SavedQueries belong to the Dataset will be returned in List/Get
	//  Dataset response. The annotation_specs field
	//  will not be populated except for UI cases which will only use
	//  [annotation_spec_count][google.cloud.aiplatform.v1.SavedQuery.annotation_spec_count].
	//  In CreateDataset request, a SavedQuery is created together if
	//  this field is set, up to one SavedQuery can be set in CreateDatasetRequest.
	//  The SavedQuery should not contain any AnnotationSpec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.saved_queries
	SavedQueries []SavedQueryObservedState `json:"savedQueries,omitempty"`

	// Output only. The resource name of the Artifact that was created in
	//  MetadataStore when creating the Dataset. The Artifact resource name pattern
	//  is
	//  `projects/{project}/locations/{location}/metadataStores/{metadata_store}/artifacts/{artifact}`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.metadata_artifact
	MetadataArtifact *string `json:"metadataArtifact,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Dataset.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SavedQuery
type SavedQueryObservedState struct {
	// Output only. Resource name of the SavedQuery.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this SavedQuery was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when SavedQuery was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Filters on the Annotations in the dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.annotation_filter
	AnnotationFilter *string `json:"annotationFilter,omitempty"`

	// Output only. Number of AnnotationSpecs in the context of the SavedQuery.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.annotation_spec_count
	AnnotationSpecCount *int32 `json:"annotationSpecCount,omitempty"`

	// Output only. If the Annotations belonging to the SavedQuery can be used for
	//  AutoML training.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SavedQuery.support_automl_training
	SupportAutomlTraining *bool `json:"supportAutomlTraining,omitempty"`
}
