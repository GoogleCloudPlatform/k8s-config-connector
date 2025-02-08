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


// +kcc:proto=google.cloud.dataform.v1beta1.CodeCompilationConfig
type CodeCompilationConfig struct {
	// Optional. The default database (Google Cloud project ID).
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.default_database
	DefaultDatabase *string `json:"defaultDatabase,omitempty"`

	// Optional. The default schema (BigQuery dataset ID).
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.default_schema
	DefaultSchema *string `json:"defaultSchema,omitempty"`

	// Optional. The default BigQuery location to use. Defaults to "US".
	//  See the BigQuery docs for a full list of locations:
	//  https://cloud.google.com/bigquery/docs/locations.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.default_location
	DefaultLocation *string `json:"defaultLocation,omitempty"`

	// Optional. The default schema (BigQuery dataset ID) for assertions.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.assertion_schema
	AssertionSchema *string `json:"assertionSchema,omitempty"`

	// Optional. User-defined variables that are made available to project code
	//  during compilation.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.vars
	Vars map[string]string `json:"vars,omitempty"`

	// Optional. The suffix that should be appended to all database (Google Cloud
	//  project ID) names.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.database_suffix
	DatabaseSuffix *string `json:"databaseSuffix,omitempty"`

	// Optional. The suffix that should be appended to all schema (BigQuery
	//  dataset ID) names.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.schema_suffix
	SchemaSuffix *string `json:"schemaSuffix,omitempty"`

	// Optional. The prefix that should be prepended to all table names.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CodeCompilationConfig.table_prefix
	TablePrefix *string `json:"tablePrefix,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.ReleaseConfig
type ReleaseConfig struct {

	// Required. Git commit/tag/branch name at which the repository should be
	//  compiled. Must exist in the remote repository. Examples:
	//  - a commit SHA: `12ade345`
	//  - a tag: `tag1`
	//  - a branch name: `branch1`
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.git_commitish
	GitCommitish *string `json:"gitCommitish,omitempty"`

	// Optional. If set, fields of `code_compilation_config` override the default
	//  compilation settings that are specified in dataform.json.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.code_compilation_config
	CodeCompilationConfig *CodeCompilationConfig `json:"codeCompilationConfig,omitempty"`

	// Optional. Optional schedule (in cron format) for automatic creation of
	//  compilation results.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.cron_schedule
	CronSchedule *string `json:"cronSchedule,omitempty"`

	// Optional. Specifies the time zone to be used when interpreting
	//  cron_schedule. Must be a time zone name from the time zone database
	//  (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left
	//  unspecified, the default is UTC.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. The name of the currently released compilation result for this
	//  release config. This value is updated when a compilation result is created
	//  from this release config, or when this resource is updated by API call
	//  (perhaps to roll back to an earlier release). The compilation result must
	//  have been created using this release config. Must be in the format
	//  `projects/*/locations/*/repositories/*/compilationResults/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.release_compilation_result
	ReleaseCompilationResult *string `json:"releaseCompilationResult,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.ReleaseConfig.ScheduledReleaseRecord
type ReleaseConfig_ScheduledReleaseRecord struct {
	// The timestamp of this release attempt.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.ScheduledReleaseRecord.release_time
	ReleaseTime *string `json:"releaseTime,omitempty"`

	// The name of the created compilation result, if one was successfully
	//  created. Must be in the format
	//  `projects/*/locations/*/repositories/*/compilationResults/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.ScheduledReleaseRecord.compilation_result
	CompilationResult *string `json:"compilationResult,omitempty"`

	// The error status encountered upon this attempt to create the
	//  compilation result, if the attempt was unsuccessful.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.ScheduledReleaseRecord.error_status
	ErrorStatus *Status `json:"errorStatus,omitempty"`
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

// +kcc:proto=google.cloud.dataform.v1beta1.ReleaseConfig
type ReleaseConfigObservedState struct {
	// Output only. The release config's name.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.name
	Name *string `json:"name,omitempty"`

	// Output only. Records of the 10 most recent scheduled release attempts,
	//  ordered in in descending order of `release_time`. Updated whenever
	//  automatic creation of a compilation result is triggered by cron_schedule.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.ReleaseConfig.recent_scheduled_release_records
	RecentScheduledReleaseRecords []ReleaseConfig_ScheduledReleaseRecord `json:"recentScheduledReleaseRecords,omitempty"`
}
