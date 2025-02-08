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

// +kcc:proto=google.cloud.dataform.v1beta1.CompilationResult
type CompilationResult struct {

	// Immutable. Git commit/tag/branch name at which the repository should be
	//  compiled. Must exist in the remote repository. Examples:
	//  - a commit SHA: `12ade345`
	//  - a tag: `tag1`
	//  - a branch name: `branch1`
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.git_commitish
	GitCommitish *string `json:"gitCommitish,omitempty"`

	// Immutable. The name of the workspace to compile. Must be in the format
	//  `projects/*/locations/*/repositories/*/workspaces/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.workspace
	Workspace *string `json:"workspace,omitempty"`

	// Immutable. The name of the release config to compile. The release
	//  config's 'current_compilation_result' field will be updated to this
	//  compilation result. Must be in the format
	//  `projects/*/locations/*/repositories/*/releaseConfigs/*`.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.release_config
	ReleaseConfig *string `json:"releaseConfig,omitempty"`

	// Immutable. If set, fields of `code_compilation_config` override the default
	//  compilation settings that are specified in dataform.json.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.code_compilation_config
	CodeCompilationConfig *CodeCompilationConfig `json:"codeCompilationConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.CompilationResult.CompilationError
type CompilationResult_CompilationError struct {
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

// +kcc:proto=google.cloud.dataform.v1beta1.CompilationResult
type CompilationResultObservedState struct {
	// Output only. The compilation result's name.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.name
	Name *string `json:"name,omitempty"`

	// Output only. The fully resolved Git commit SHA of the code that was
	//  compiled. Not set for compilation results whose source is a workspace.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.resolved_git_commit_sha
	ResolvedGitCommitSha *string `json:"resolvedGitCommitSha,omitempty"`

	// Output only. The version of `@dataform/core` that was used for compilation.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.dataform_core_version
	DataformCoreVersion *string `json:"dataformCoreVersion,omitempty"`

	// Output only. Errors encountered during project compilation.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.compilation_errors
	CompilationErrors []CompilationResult_CompilationError `json:"compilationErrors,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.CompilationResult.CompilationError
type CompilationResult_CompilationErrorObservedState struct {
	// Output only. The error's top level message.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.CompilationError.message
	Message *string `json:"message,omitempty"`

	// Output only. The error's full stack trace.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.CompilationError.stack
	Stack *string `json:"stack,omitempty"`

	// Output only. The path of the file where this error occurred, if
	//  available, relative to the project root.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.CompilationError.path
	Path *string `json:"path,omitempty"`

	// Output only. The identifier of the action where this error occurred, if
	//  available.
	// +kcc:proto:field=google.cloud.dataform.v1beta1.CompilationResult.CompilationError.action_target
	ActionTarget *Target `json:"actionTarget,omitempty"`
}
