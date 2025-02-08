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


// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.AffinityClusteringConfig
type AffinityClusteringConfig struct {
	// Number of iterations to perform. Default value is 1.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.AffinityClusteringConfig.compression_round_count
	CompressionRoundCount *int64 `json:"compressionRoundCount,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.BigQueryInputConfig
type BigQueryInputConfig struct {
	// Required. Format is `projects/*/datasets/*/tables/*`.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.BigQueryInputConfig.bigquery_table
	BigqueryTable *string `json:"bigqueryTable,omitempty"`

	// Required. Schema mapping file
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.BigQueryInputConfig.gcs_uri
	GcsURI *string `json:"gcsURI,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ConnectedComponentsConfig
type ConnectedComponentsConfig struct {
	// Threshold used for connected components. Default value is 0.85.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ConnectedComponentsConfig.weight_threshold
	WeightThreshold *float32 `json:"weightThreshold,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob
type EntityReconciliationJob struct {

	// Required. Information about the input BigQuery tables.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.input_config
	InputConfig *InputConfig `json:"inputConfig,omitempty"`

	// Required. The desired output location.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.output_config
	OutputConfig *OutputConfig `json:"outputConfig,omitempty"`

	// Optional. Recon configs to adjust the clustering behavior.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.recon_config
	ReconConfig *ReconConfig `json:"reconConfig,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.InputConfig
type InputConfig struct {
	// Set of input BigQuery tables.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.InputConfig.bigquery_input_configs
	BigqueryInputConfigs []BigQueryInputConfig `json:"bigqueryInputConfigs,omitempty"`

	// Entity type
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.InputConfig.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// Optional. Provide the bigquery table containing the previous results if
	//  cluster ID stability is desired. Format is
	//  `projects/*/datasets/*/tables/*`.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.InputConfig.previous_result_bigquery_table
	PreviousResultBigqueryTable *string `json:"previousResultBigqueryTable,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.OutputConfig
type OutputConfig struct {
	// Format is “projects/*/datasets/*”.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.OutputConfig.bigquery_dataset
	BigqueryDataset *string `json:"bigqueryDataset,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ReconConfig
type ReconConfig struct {
	// Configs for connected components.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.connected_components_config
	ConnectedComponentsConfig *ConnectedComponentsConfig `json:"connectedComponentsConfig,omitempty"`

	// Configs for affinity clustering.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.affinity_clustering_config
	AffinityClusteringConfig *AffinityClusteringConfig `json:"affinityClusteringConfig,omitempty"`

	// Extra options that affect entity clustering behavior.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.options
	Options *ReconConfig_Options `json:"options,omitempty"`

	// Model Configs
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.model_config
	ModelConfig *ReconConfig_ModelConfig `json:"modelConfig,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.ModelConfig
type ReconConfig_ModelConfig struct {
	// Model name. Refer to external documentation for valid names.
	//  If unspecified, it defaults to the one mentioned in the documentation.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.ModelConfig.model_name
	ModelName *string `json:"modelName,omitempty"`

	// Model version tag. Refer to external documentation for valid tags.
	//  If unspecified, it defaults to the one mentioned in the documentation.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.ModelConfig.version_tag
	VersionTag *string `json:"versionTag,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.Options
type ReconConfig_Options struct {
	// If true, separate clusters by their geographic region (from geocoding).
	//  Uses the following entity features:
	//
	//  - schema.org/addressLocality
	//  - schema.org/addressRegion
	//  - schema.org/addressCountry
	//  Warning: processing will no longer be regionalized!
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.ReconConfig.Options.enable_geocoding_separation
	EnableGeocodingSeparation *bool `json:"enableGeocodingSeparation,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob
type EntityReconciliationJobObservedState struct {
	// Output only. Resource name of the EntityReconciliationJob.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when the job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. Time when the EntityReconciliationJob was created.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the EntityReconciliationJob entered any of the
	//  following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`,
	//  `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the EntityReconciliationJob was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.enterpriseknowledgegraph.v1.EntityReconciliationJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
