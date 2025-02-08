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


// +kcc:proto=google.cloud.config.v1.ApplyResults
type ApplyResults struct {
	// Location of a blueprint copy and other manifests in Google Cloud Storage.
	//  Format: `gs://{bucket}/{object}`
	// +kcc:proto:field=google.cloud.config.v1.ApplyResults.content
	Content *string `json:"content,omitempty"`

	// Location of artifacts (e.g. logs) in Google Cloud Storage.
	//  Format: `gs://{bucket}/{object}`
	// +kcc:proto:field=google.cloud.config.v1.ApplyResults.artifacts
	Artifacts *string `json:"artifacts,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.config.v1.GitSource
type GitSource struct {
	// Optional. Repository URL.
	//  Example: 'https://github.com/kubernetes/examples.git'
	// +kcc:proto:field=google.cloud.config.v1.GitSource.repo
	Repo *string `json:"repo,omitempty"`

	// Optional. Subdirectory inside the repository.
	//  Example: 'staging/my-package'
	// +kcc:proto:field=google.cloud.config.v1.GitSource.directory
	Directory *string `json:"directory,omitempty"`

	// Optional. Git reference (e.g. branch or tag).
	// +kcc:proto:field=google.cloud.config.v1.GitSource.ref
	Ref *string `json:"ref,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.Revision
type Revision struct {

	// Revision name. Format:
	//  `projects/{project}/locations/{location}/deployments/{deployment}/
	//  revisions/{revision}`
	// +kcc:proto:field=google.cloud.config.v1.Revision.name
	Name *string `json:"name,omitempty"`

	// Optional. Input to control quota checks for resources in terraform
	//  configuration files. There are limited resources on which quota validation
	//  applies.
	// +kcc:proto:field=google.cloud.config.v1.Revision.quota_validation
	QuotaValidation *string `json:"quotaValidation,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.TerraformBlueprint
type TerraformBlueprint struct {
	// URI of an object in Google Cloud Storage.
	//  Format: `gs://{bucket}/{object}`
	//
	//  URI may also specify an object version for zipped objects.
	//  Format: `gs://{bucket}/{object}#{version}`
	// +kcc:proto:field=google.cloud.config.v1.TerraformBlueprint.gcs_source
	GcsSource *string `json:"gcsSource,omitempty"`

	// URI of a public Git repo.
	// +kcc:proto:field=google.cloud.config.v1.TerraformBlueprint.git_source
	GitSource *GitSource `json:"gitSource,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.config.v1.TerraformError
type TerraformError struct {
	// Address of the resource associated with the error,
	//  e.g. `google_compute_network.vpc_network`.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.resource_address
	ResourceAddress *string `json:"resourceAddress,omitempty"`

	// HTTP response code returned from Google Cloud Platform APIs when Terraform
	//  fails to provision the resource. If unset or 0, no HTTP response code was
	//  returned by Terraform.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.http_response_code
	HTTPResponseCode *int32 `json:"httpResponseCode,omitempty"`

	// A human-readable error description.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.error_description
	ErrorDescription *string `json:"errorDescription,omitempty"`

	// Original error response from underlying Google API, if available.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.error
	Error *Status `json:"error,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.TerraformOutput
type TerraformOutput struct {
	// Identifies whether Terraform has set this output as a potential
	//  sensitive value.
	// +kcc:proto:field=google.cloud.config.v1.TerraformOutput.sensitive
	Sensitive *bool `json:"sensitive,omitempty"`

	// Value of output.
	// +kcc:proto:field=google.cloud.config.v1.TerraformOutput.value
	Value *Value `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.TerraformVariable
type TerraformVariable struct {
	// Input variable value.
	// +kcc:proto:field=google.cloud.config.v1.TerraformVariable.input_value
	InputValue *Value `json:"inputValue,omitempty"`
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

// +kcc:proto=google.cloud.config.v1.Revision
type RevisionObservedState struct {
	// Output only. A blueprint described using Terraform's HashiCorp
	//  Configuration Language as a root module.
	// +kcc:proto:field=google.cloud.config.v1.Revision.terraform_blueprint
	TerraformBlueprint *TerraformBlueprint `json:"terraformBlueprint,omitempty"`

	// Output only. Time when the revision was created.
	// +kcc:proto:field=google.cloud.config.v1.Revision.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the revision was last modified.
	// +kcc:proto:field=google.cloud.config.v1.Revision.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The action which created this revision
	// +kcc:proto:field=google.cloud.config.v1.Revision.action
	Action *string `json:"action,omitempty"`

	// Output only. Current state of the revision.
	// +kcc:proto:field=google.cloud.config.v1.Revision.state
	State *string `json:"state,omitempty"`

	// Output only. Outputs and artifacts from applying a deployment.
	// +kcc:proto:field=google.cloud.config.v1.Revision.apply_results
	ApplyResults *ApplyResults `json:"applyResults,omitempty"`

	// Output only. Additional info regarding the current state.
	// +kcc:proto:field=google.cloud.config.v1.Revision.state_detail
	StateDetail *string `json:"stateDetail,omitempty"`

	// Output only. Code describing any errors that may have occurred.
	// +kcc:proto:field=google.cloud.config.v1.Revision.error_code
	ErrorCode *string `json:"errorCode,omitempty"`

	// Output only. Cloud Build instance UUID associated with this revision.
	// +kcc:proto:field=google.cloud.config.v1.Revision.build
	Build *string `json:"build,omitempty"`

	// Output only. Location of Revision operation logs in
	//  `gs://{bucket}/{object}` format.
	// +kcc:proto:field=google.cloud.config.v1.Revision.logs
	Logs *string `json:"logs,omitempty"`

	// Output only. Errors encountered when creating or updating this deployment.
	//  Errors are truncated to 10 entries, see `delete_results` and `error_logs`
	//  for full details.
	// +kcc:proto:field=google.cloud.config.v1.Revision.tf_errors
	TfErrors []TerraformError `json:"tfErrors,omitempty"`

	// Output only. Location of Terraform error logs in Google Cloud Storage.
	//  Format: `gs://{bucket}/{object}`.
	// +kcc:proto:field=google.cloud.config.v1.Revision.error_logs
	ErrorLogs *string `json:"errorLogs,omitempty"`

	// Output only. User-specified Service Account (SA) to be used as credential
	//  to manage resources. Format:
	//  `projects/{projectID}/serviceAccounts/{serviceAccount}`
	// +kcc:proto:field=google.cloud.config.v1.Revision.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. By default, Infra Manager will return a failure when
	//  Terraform encounters a 409 code (resource conflict error) during actuation.
	//  If this flag is set to true, Infra Manager will instead
	//  attempt to automatically import the resource into the Terraform state (for
	//  supported resource types) and continue actuation.
	//
	//  Not all resource types are supported, refer to documentation.
	// +kcc:proto:field=google.cloud.config.v1.Revision.import_existing_resources
	ImportExistingResources *bool `json:"importExistingResources,omitempty"`

	// Output only. The user-specified Cloud Build worker pool resource in which
	//  the Cloud Build job will execute. Format:
	//  `projects/{project}/locations/{location}/workerPools/{workerPoolId}`.
	//  If this field is unspecified, the default Cloud Build worker pool will be
	//  used.
	// +kcc:proto:field=google.cloud.config.v1.Revision.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

	// Output only. The user-specified Terraform version constraint.
	//  Example: "=1.3.10".
	// +kcc:proto:field=google.cloud.config.v1.Revision.tf_version_constraint
	TfVersionConstraint *string `json:"tfVersionConstraint,omitempty"`

	// Output only. The version of Terraform used to create the Revision.
	//  It is in the format of "Major.Minor.Patch", for example, "1.3.10".
	// +kcc:proto:field=google.cloud.config.v1.Revision.tf_version
	TfVersion *string `json:"tfVersion,omitempty"`

	// Output only. Cloud Storage path containing quota validation results. This
	//  field is set when a user sets Deployment.quota_validation field to ENABLED
	//  or ENFORCED. Format: `gs://{bucket}/{object}`.
	// +kcc:proto:field=google.cloud.config.v1.Revision.quota_validation_results
	QuotaValidationResults *string `json:"quotaValidationResults,omitempty"`
}
