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


// +kcc:proto=google.cloud.workflows.v1.Workflow
type Workflow struct {
	// The resource name of the workflow.
	//  Format: projects/{project}/locations/{location}/workflows/{workflow}
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.name
	Name *string `json:"name,omitempty"`

	// Description of the workflow provided by the user.
	//  Must be at most 1000 unicode characters long.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.description
	Description *string `json:"description,omitempty"`

	// Labels associated with this workflow.
	//  Labels can contain at most 64 entries. Keys and values can be no longer
	//  than 63 characters and can only contain lowercase letters, numeric
	//  characters, underscores, and dashes. Label keys must start with a letter.
	//  International characters are allowed.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The service account associated with the latest workflow version.
	//  This service account represents the identity of the workflow and determines
	//  what permissions the workflow has.
	//  Format: projects/{project}/serviceAccounts/{account} or {account}
	//
	//  Using `-` as a wildcard for the `{project}` or not providing one at all
	//  will infer the project from the account. The `{account}` value can be the
	//  `email` address or the `unique_id` of the service account.
	//
	//  If not provided, workflow will use the project's default service account.
	//  Modifying this field for an existing workflow results in a new workflow
	//  revision.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Workflow code to be executed. The size limit is 128KB.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.source_contents
	SourceContents *string `json:"sourceContents,omitempty"`

	// Optional. The resource name of a KMS crypto key used to encrypt or decrypt
	//  the data associated with the workflow.
	//
	//  Format:
	//  projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	//
	//  Using `-` as a wildcard for the `{project}` or not providing one at all
	//  will infer the project from the account.
	//
	//  If not provided, data associated with the workflow will not be
	//  CMEK-encrypted.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.crypto_key_name
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`

	// Optional. Describes the level of platform logging to apply to calls and
	//  call responses during executions of this workflow. If both the workflow and
	//  the execution specify a logging level, the execution level takes
	//  precedence.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.call_log_level
	CallLogLevel *string `json:"callLogLevel,omitempty"`

	// Optional. User-defined environment variables associated with this workflow
	//  revision. This map has a maximum length of 20. Each string can take up to
	//  40KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or
	//  “WORKFLOWS".
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.user_env_vars
	UserEnvVars map[string]string `json:"userEnvVars,omitempty"`
}

// +kcc:proto=google.cloud.workflows.v1.Workflow.StateError
type Workflow_StateError struct {
	// Provides specifics about the error.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.StateError.details
	Details *string `json:"details,omitempty"`

	// The type of this state error.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.StateError.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.workflows.v1.Workflow
type WorkflowObservedState struct {
	// Output only. State of the workflow deployment.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.state
	State *string `json:"state,omitempty"`

	// Output only. The revision of the workflow.
	//  A new revision of a workflow is created as a result of updating the
	//  following properties of a workflow:
	//
	//  - [Service account][google.cloud.workflows.v1.Workflow.service_account]
	//  - [Workflow code to be
	//  executed][google.cloud.workflows.v1.Workflow.source_contents]
	//
	//  The format is "000001-a4d", where the first six characters define
	//  the zero-padded revision ordinal number. They are followed by a hyphen and
	//  three hexadecimal random characters.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp for when the workflow was created.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp for when the workflow was last updated.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The timestamp for the latest revision of the workflow's
	//  creation.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. Error regarding the state of the workflow. For example, this
	//  field will have error details if the execution data is unavailable due to
	//  revoked KMS key permissions.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.state_error
	StateError *Workflow_StateError `json:"stateError,omitempty"`
}
