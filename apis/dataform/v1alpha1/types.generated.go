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

// +kcc:proto=google.cloud.dataform.v1beta1.WorkflowInvocation
type WorkflowInvocation struct {

	// Immutable. The name of the compilation result to use for this invocation.
	//  Must be in the format
	//  `projects/*/locations/*/repositories/*/compilationResults/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowInvocation.compilation_result
	CompilationResult *string `json:"compilationResult,omitempty"`

	// Immutable. The name of the workflow config to invoke. Must be in the
	//  format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowInvocation.workflow_config
	WorkflowConfig *string `json:"workflowConfig,omitempty"`

	// Immutable. If left unset, a default InvocationConfig will be used.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowInvocation.invocation_config
	InvocationConfig *InvocationConfig `json:"invocationConfig,omitempty"`
}

// +kcc:proto=google.type.Interval
type Interval struct {
	// Optional. Inclusive start of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be the same
	//  or after the start.
	// +kcc:proto:field=google.type.Interval.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Exclusive end of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be before the
	//  end.
	// +kcc:proto:field=google.type.Interval.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.WorkflowInvocation
type WorkflowInvocationObservedState struct {
	// Output only. The workflow invocation's name.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowInvocation.name
	Name *string `json:"name,omitempty"`

	// Output only. This workflow invocation's current state.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowInvocation.state
	State *string `json:"state,omitempty"`

	// Output only. This workflow invocation's timing details.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.WorkflowInvocation.invocation_timing
	InvocationTiming *Interval `json:"invocationTiming,omitempty"`
}
