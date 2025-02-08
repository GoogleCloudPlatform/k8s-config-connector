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


// +kcc:proto=google.cloud.dataform.v1beta1.InvocationConfig
type InvocationConfig struct {
	// Optional. The set of action identifiers to include.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.InvocationConfig.included_targets
	IncludedTargets []Target `json:"includedTargets,omitempty"`

	// Optional. The set of tags to include.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.InvocationConfig.included_tags
	IncludedTags []string `json:"includedTags,omitempty"`

	// Optional. When set to true, transitive dependencies of included actions
	//  will be executed.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.InvocationConfig.transitive_dependencies_included
	TransitiveDependenciesIncluded *bool `json:"transitiveDependenciesIncluded,omitempty"`

	// Optional. When set to true, transitive dependents of included actions will
	//  be executed.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.InvocationConfig.transitive_dependents_included
	TransitiveDependentsIncluded *bool `json:"transitiveDependentsIncluded,omitempty"`

	// Optional. When set to true, any incremental tables will be fully refreshed.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.InvocationConfig.fully_refresh_incremental_tables_enabled
	FullyRefreshIncrementalTablesEnabled *bool `json:"fullyRefreshIncrementalTablesEnabled,omitempty"`

	// Optional. The service account to run workflow invocations under.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.InvocationConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.Target
type Target struct {
	// The action's database (Google Cloud project ID) .
	// +kcc:proto:field=google.cloud.dataform.v1beta1.Target.database
	Database *string `json:"database,omitempty"`

	// The action's schema (BigQuery dataset ID), within `database`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.Target.schema
	Schema *string `json:"schema,omitempty"`

	// The action's name, within `database` and `schema`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.Target.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.WorkflowConfig
type WorkflowConfig struct {

	// Required. The name of the release config whose release_compilation_result
	//  should be executed. Must be in the format
	//  `projects/*/locations/*/repositories/*/releaseConfigs/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.release_config
	ReleaseConfig *string `json:"releaseConfig,omitempty"`

	// Optional. If left unset, a default InvocationConfig will be used.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.invocation_config
	InvocationConfig *InvocationConfig `json:"invocationConfig,omitempty"`

	// Optional. Optional schedule (in cron format) for automatic execution of
	//  this workflow config.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.cron_schedule
	CronSchedule *string `json:"cronSchedule,omitempty"`

	// Optional. Specifies the time zone to be used when interpreting
	//  cron_schedule. Must be a time zone name from the time zone database
	//  (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left
	//  unspecified, the default is UTC.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.WorkflowConfig.ScheduledExecutionRecord
type WorkflowConfig_ScheduledExecutionRecord struct {
	// The timestamp of this execution attempt.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.ScheduledExecutionRecord.execution_time
	ExecutionTime *string `json:"executionTime,omitempty"`

	// The name of the created workflow invocation, if one was successfully
	//  created. Must be in the format
	//  `projects/*/locations/*/repositories/*/workflowInvocations/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.ScheduledExecutionRecord.workflow_invocation
	WorkflowInvocation *string `json:"workflowInvocation,omitempty"`

	// The error status encountered upon this attempt to create the
	//  workflow invocation, if the attempt was unsuccessful.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.ScheduledExecutionRecord.error_status
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

// +kcc:proto=google.cloud.dataform.v1beta1.WorkflowConfig
type WorkflowConfigObservedState struct {
	// Output only. The workflow config's name.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.name
	Name *string `json:"name,omitempty"`

	// Output only. Records of the 10 most recent scheduled execution attempts,
	//  ordered in in descending order of `execution_time`. Updated whenever
	//  automatic creation of a workflow invocation is triggered by cron_schedule.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowConfig.recent_scheduled_execution_records
	RecentScheduledExecutionRecords []WorkflowConfig_ScheduledExecutionRecord `json:"recentScheduledExecutionRecords,omitempty"`
}
