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


// +kcc:proto=google.cloud.aiplatform.v1.BigQuerySource
type BigQuerySource struct {
	// Required. BigQuery URI to a table, up to 2000 characters long.
	//  Accepted forms:
	//
	//  *  BigQuery path. For example: `bq://projectId.bqDatasetId.bqTableId`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BigQuerySource.input_uri
	InputURI *string `json:"inputURI,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureGroup
type FeatureGroup struct {
	// Indicates that features for this group come from BigQuery Table/View.
	//  By default treats the source as a sparse time series source. The BigQuery
	//  source table or view must have at least one entity ID column and a column
	//  named `feature_timestamp`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.big_query
	BigQuery *FeatureGroup_BigQuery `json:"bigQuery,omitempty"`

	// Identifier. Name of the FeatureGroup. Format:
	//  `projects/{project}/locations/{location}/featureGroups/{featureGroup}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.name
	Name *string `json:"name,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	//  FeatureGroup.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//  No more than 64 user labels can be associated with one
	//  FeatureGroup(System labels are excluded)." System reserved label keys
	//  are prefixed with "aiplatform.googleapis.com/" and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the FeatureGroup.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureGroup.BigQuery
type FeatureGroup_BigQuery struct {
	// Required. Immutable. The BigQuery source URI that points to either a
	//  BigQuery Table or View.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.BigQuery.big_query_source
	BigQuerySource *BigQuerySource `json:"bigQuerySource,omitempty"`

	// Optional. Columns to construct entity_id / row keys.
	//  If not provided defaults to `entity_id`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.BigQuery.entity_id_columns
	EntityIDColumns []string `json:"entityIDColumns,omitempty"`

	// Optional. Set if the data source is not a time-series.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.BigQuery.static_data_source
	StaticDataSource *bool `json:"staticDataSource,omitempty"`

	// Optional. If the source is a time-series source, this can be set to
	//  control how downstream sources (ex:
	//  [FeatureView][google.cloud.aiplatform.v1.FeatureView] ) will treat
	//  time-series sources. If not set, will treat the source as a time-series
	//  source with `feature_timestamp` as timestamp column and no scan boundary.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.BigQuery.time_series
	TimeSeries *FeatureGroup_BigQuery_TimeSeries `json:"timeSeries,omitempty"`

	// Optional. If set, all feature values will be fetched
	//  from a single row per unique entityId including nulls.
	//  If not set, will collapse all rows for each unique entityId into a singe
	//  row with any non-null values if present, if no non-null values are
	//  present will sync null.
	//  ex: If source has schema
	//  `(entity_id, feature_timestamp, f0, f1)` and the following rows:
	//  `(e1, 2020-01-01T10:00:00.123Z, 10, 15)`
	//  `(e1, 2020-02-01T10:00:00.123Z, 20, null)`
	//  If dense is set, `(e1, 20, null)` is synced to online stores. If dense is
	//  not set, `(e1, 20, 15)` is synced to online stores.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.BigQuery.dense
	Dense *bool `json:"dense,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureGroup.BigQuery.TimeSeries
type FeatureGroup_BigQuery_TimeSeries struct {
	// Optional. Column hosting timestamp values for a time-series source.
	//  Will be used to determine the latest `feature_values` for each entity.
	//  Optional. If not provided, column named `feature_timestamp` of
	//  type `TIMESTAMP` will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.BigQuery.TimeSeries.timestamp_column
	TimestampColumn *string `json:"timestampColumn,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureGroup
type FeatureGroupObservedState struct {
	// Output only. Timestamp when this FeatureGroup was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this FeatureGroup was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
