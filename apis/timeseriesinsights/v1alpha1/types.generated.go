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


// +kcc:proto=google.cloud.timeseriesinsights.v1.BigqueryMapping
type BigqueryMapping struct {
	// The column which should be used as the event timestamps. If not specified
	//  'Timestamp' is used by default. The column may have TIMESTAMP or INT64
	//  type (the latter is interpreted as microseconds since the Unix epoch).
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.BigqueryMapping.timestamp_column
	TimestampColumn *string `json:"timestampColumn,omitempty"`

	// The column which should be used as the group ID (grouping events into
	//  sessions). If not specified 'GroupId' is used by default, if the input
	//  table does not have such a column, random unique group IDs are
	//  generated automatically (different group ID per input row).
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.BigqueryMapping.group_id_column
	GroupIDColumn *string `json:"groupIDColumn,omitempty"`

	// The list of columns that should be translated to dimensions. If empty,
	//  all columns are translated to dimensions. The timestamp and group_id
	//  columns should not be listed here again. Columns are expected to have
	//  primitive types (STRING, INT64, FLOAT64 or NUMERIC).
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.BigqueryMapping.dimension_column
	DimensionColumn []string `json:"dimensionColumn,omitempty"`
}

// +kcc:proto=google.cloud.timeseriesinsights.v1.DataSet
type DataSet struct {
	// The dataset name, which will be used for querying, status and unload
	//  requests. This must be unique within a project.
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSet.name
	Name *string `json:"name,omitempty"`

	// [Data dimension names][google.cloud.timeseriesinsights.v1.EventDimension.name] allowed for this `DataSet`.
	//
	//  If left empty, all dimension names are included. This field works as a
	//  filter to avoid regenerating the data.
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSet.data_names
	DataNames []string `json:"dataNames,omitempty"`

	// Input data.
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSet.data_sources
	DataSources []DataSource `json:"dataSources,omitempty"`

	// Dataset state in the system.
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSet.state
	State *string `json:"state,omitempty"`

	// Dataset processing status.
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSet.status
	Status *Status `json:"status,omitempty"`

	// Periodically we discard dataset [Event][google.cloud.timeseriesinsights.v1.Event] objects that have
	//  timestamps older than 'ttl'.  Omitting this field or a zero value means no
	//  events are discarded.
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSet.ttl
	Ttl *string `json:"ttl,omitempty"`
}

// +kcc:proto=google.cloud.timeseriesinsights.v1.DataSource
type DataSource struct {
	// Data source URI.
	//
	//  1) Google Cloud Storage files (JSON) are defined in the following form.
	//  `gs://bucket_name/object_name`. For more information on Cloud Storage URIs,
	//  please see https://cloud.google.com/storage/docs/reference-uris.
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSource.uri
	URI *string `json:"uri,omitempty"`

	// For BigQuery inputs defines the columns that should be used for dimensions
	//  (including time and group ID).
	// +kcc:proto:field=google.cloud.timeseriesinsights.v1.DataSource.bq_mapping
	BqMapping *BigqueryMapping `json:"bqMapping,omitempty"`
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
