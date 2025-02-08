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


// +kcc:proto=google.cloud.retail.v2.BigQuerySource
type BigQuerySource struct {
	// BigQuery time partitioned table's _PARTITIONDATE in YYYY-MM-DD format.
	//
	//  Only supported in
	//  [ImportProductsRequest][google.cloud.retail.v2.ImportProductsRequest].
	// +kcc:proto:field=google.cloud.retail.v2.BigQuerySource.partition_date
	PartitionDate *Date `json:"partitionDate,omitempty"`

	// The project ID (can be project # or ID) that the BigQuery source is in with
	//  a length limit of 128 characters. If not specified, inherits the project
	//  ID from the parent request.
	// +kcc:proto:field=google.cloud.retail.v2.BigQuerySource.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The BigQuery data set to copy the data from with a length limit
	//  of 1,024 characters.
	// +kcc:proto:field=google.cloud.retail.v2.BigQuerySource.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// Required. The BigQuery table to copy the data from with a length limit of
	//  1,024 characters.
	// +kcc:proto:field=google.cloud.retail.v2.BigQuerySource.table_id
	TableID *string `json:"tableID,omitempty"`

	// Intermediate Cloud Storage directory used for the import with a length
	//  limit of 2,000 characters. Can be specified if one wants to have the
	//  BigQuery export to a specific Cloud Storage directory.
	// +kcc:proto:field=google.cloud.retail.v2.BigQuerySource.gcs_staging_dir
	GcsStagingDir *string `json:"gcsStagingDir,omitempty"`

	// The schema to use when parsing the data from the source.
	//
	//  Supported values for product imports:
	//
	//  * `product` (default): One JSON [Product][google.cloud.retail.v2.Product]
	//  per line. Each product must
	//    have a valid [Product.id][google.cloud.retail.v2.Product.id].
	//  * `product_merchant_center`: See [Importing catalog data from Merchant
	//    Center](https://cloud.google.com/retail/recommendations-ai/docs/upload-catalog#mc).
	//
	//  Supported values for user events imports:
	//
	//  * `user_event` (default): One JSON
	//  [UserEvent][google.cloud.retail.v2.UserEvent] per line.
	//  * `user_event_ga360`:
	//    The schema is available here:
	//    https://support.google.com/analytics/answer/3437719.
	//  * `user_event_ga4`:
	//    The schema is available here:
	//    https://support.google.com/analytics/answer/7029846.
	//
	//  Supported values for autocomplete imports:
	//
	//  * `suggestions` (default): One JSON completion suggestion per line.
	//  * `denylist`:  One JSON deny suggestion per line.
	//  * `allowlist`:  One JSON allow suggestion per line.
	// +kcc:proto:field=google.cloud.retail.v2.BigQuerySource.data_schema
	DataSchema *string `json:"dataSchema,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.CompletionConfig
type CompletionConfig struct {
	// Required. Immutable. Fully qualified name
	//  `projects/*/locations/*/catalogs/*/completionConfig`
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.name
	Name *string `json:"name,omitempty"`

	// Specifies the matching order for autocomplete suggestions, e.g., a query
	//  consisting of 'sh' with 'out-of-order' specified would suggest "women's
	//  shoes", whereas a query of 'red s' with 'exact-prefix' specified would
	//  suggest "red shoes". Currently supported values:
	//
	//  * 'out-of-order'
	//  * 'exact-prefix'
	//
	//  Default value: 'exact-prefix'.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.matching_order
	MatchingOrder *string `json:"matchingOrder,omitempty"`

	// The maximum number of autocomplete suggestions returned per term. Default
	//  value is 20. If left unset or set to 0, then will fallback to default
	//  value.
	//
	//  Value range is 1 to 20.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.max_suggestions
	MaxSuggestions *int32 `json:"maxSuggestions,omitempty"`

	// The minimum number of characters needed to be typed in order to get
	//  suggestions. Default value is 2. If left unset or set to 0, then will
	//  fallback to default value.
	//
	//  Value range is 1 to 20.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.min_prefix_length
	MinPrefixLength *int32 `json:"minPrefixLength,omitempty"`

	// If set to true, the auto learning function is enabled. Auto learning uses
	//  user data to generate suggestions using ML techniques. Default value is
	//  false. Only after enabling auto learning can users use `cloud-retail`
	//  data in
	//  [CompleteQueryRequest][google.cloud.retail.v2.CompleteQueryRequest].
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.auto_learning
	AutoLearning *bool `json:"autoLearning,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.CompletionDataInputConfig
type CompletionDataInputConfig struct {
	// Required. BigQuery input source.
	//
	//  Add the IAM permission "BigQuery Data Viewer" for
	//  cloud-retail-customer-data-access@system.gserviceaccount.com before
	//  using this feature otherwise an error is thrown.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionDataInputConfig.big_query_source
	BigQuerySource *BigQuerySource `json:"bigQuerySource,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.CompletionConfig
type CompletionConfigObservedState struct {
	// Output only. The source data for the latest import of the autocomplete
	//  suggestion phrases.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.suggestions_input_config
	SuggestionsInputConfig *CompletionDataInputConfig `json:"suggestionsInputConfig,omitempty"`

	// Output only. Name of the LRO corresponding to the latest suggestion terms
	//  list import.
	//
	//  Can use [GetOperation][google.longrunning.Operations.GetOperation] API
	//  method to retrieve the latest state of the Long Running Operation.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.last_suggestions_import_operation
	LastSuggestionsImportOperation *string `json:"lastSuggestionsImportOperation,omitempty"`

	// Output only. The source data for the latest import of the autocomplete
	//  denylist phrases.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.denylist_input_config
	DenylistInputConfig *CompletionDataInputConfig `json:"denylistInputConfig,omitempty"`

	// Output only. Name of the LRO corresponding to the latest denylist import.
	//
	//  Can use [GetOperation][google.longrunning.Operations.GetOperation] API to
	//  retrieve the latest state of the Long Running Operation.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.last_denylist_import_operation
	LastDenylistImportOperation *string `json:"lastDenylistImportOperation,omitempty"`

	// Output only. The source data for the latest import of the autocomplete
	//  allowlist phrases.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.allowlist_input_config
	AllowlistInputConfig *CompletionDataInputConfig `json:"allowlistInputConfig,omitempty"`

	// Output only. Name of the LRO corresponding to the latest allowlist import.
	//
	//  Can use [GetOperation][google.longrunning.Operations.GetOperation] API to
	//  retrieve the latest state of the Long Running Operation.
	// +kcc:proto:field=google.cloud.retail.v2.CompletionConfig.last_allowlist_import_operation
	LastAllowlistImportOperation *string `json:"lastAllowlistImportOperation,omitempty"`
}
