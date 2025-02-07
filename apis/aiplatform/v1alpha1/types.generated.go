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


// +kcc:proto=google.cloud.aiplatform.v1.DeployedIndexRef
type DeployedIndexRef struct {
	// Immutable. A resource name of the IndexEndpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexRef.index_endpoint
	IndexEndpoint *string `json:"indexEndpoint,omitempty"`

	// Immutable. The ID of the DeployedIndex in the above IndexEndpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexRef.deployed_index_id
	DeployedIndexID *string `json:"deployedIndexID,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.Index
type Index struct {

	// Required. The display name of the Index.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.description
	Description *string `json:"description,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  additional information about the Index, that is specific to it. Unset if
	//  the Index does not have any additional information. The schema is defined
	//  as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.metadata_schema_uri
	MetadataSchemaURI *string `json:"metadataSchemaURI,omitempty"`

	// An additional information about the Index; the schema of the metadata can
	//  be found in
	//  [metadata_schema][google.cloud.aiplatform.v1.Index.metadata_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.metadata
	Metadata *Value `json:"metadata,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Indexes.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The update method to use with this Index. If not set,
	//  BATCH_UPDATE will be used by default.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.index_update_method
	IndexUpdateMethod *string `json:"indexUpdateMethod,omitempty"`

	// Immutable. Customer-managed encryption key spec for an Index. If set, this
	//  Index and all sub-resources of this Index will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IndexStats
type IndexStats struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.DeployedIndexRef
type DeployedIndexRefObservedState struct {
	// Output only. The display name of the DeployedIndex.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexRef.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Index
type IndexObservedState struct {
	// Output only. The resource name of the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.name
	Name *string `json:"name,omitempty"`

	// Output only. The pointers to DeployedIndexes created from this Index.
	//  An Index can be only deleted if all its DeployedIndexes had been undeployed
	//  first.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.deployed_indexes
	DeployedIndexes []DeployedIndexRef `json:"deployedIndexes,omitempty"`

	// Output only. Timestamp when this Index was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Index was most recently updated.
	//  This also includes any update to the contents of the Index.
	//  Note that Operations working on this Index may have their
	//  [Operations.metadata.generic_metadata.update_time]
	//  [google.cloud.aiplatform.v1.GenericOperationMetadata.update_time] a little
	//  after the value of this timestamp, yet that does not mean their results are
	//  not already reflected in the Index. Result of any successfully completed
	//  Operation on the Index is reflected in it.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Stats of the index resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.index_stats
	IndexStats *IndexStats `json:"indexStats,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Index.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IndexStats
type IndexStatsObservedState struct {
	// Output only. The number of dense vectors in the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexStats.vectors_count
	VectorsCount *int64 `json:"vectorsCount,omitempty"`

	// Output only. The number of sparse vectors in the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexStats.sparse_vectors_count
	SparseVectorsCount *int64 `json:"sparseVectorsCount,omitempty"`

	// Output only. The number of shards in the Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexStats.shards_count
	ShardsCount *int32 `json:"shardsCount,omitempty"`
}
