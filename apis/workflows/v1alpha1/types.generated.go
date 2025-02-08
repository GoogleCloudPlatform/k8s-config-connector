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


// +kcc:proto=google.cloud.workflows.v1beta.Workflow
type Workflow struct {
	// The resource name of the workflow.
	//  Format: projects/{project}/locations/{location}/workflows/{workflow}
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.name
	Name *string `json:"name,omitempty"`

	// Description of the workflow provided by the user.
	//  Must be at most 1000 unicode characters long.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.description
	Description *string `json:"description,omitempty"`

	// Labels associated with this workflow.
	//  Labels can contain at most 64 entries. Keys and values can be no longer
	//  than 63 characters and can only contain lowercase letters, numeric
	//  characters, underscores and dashes. Label keys must start with a letter.
	//  International characters are allowed.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Name of the service account associated with the latest workflow version.
	//  This service account represents the identity of the workflow and determines
	//  what permissions the workflow has.
	//  Format: projects/{project}/serviceAccounts/{account}
	//
	//  Using `-` as a wildcard for the `{project}` will infer the project from
	//  the account. The `{account}` value can be the `email` address or the
	//  `unique_id` of the service account.
	//
	//  If not provided, workflow will use the project's default service account.
	//  Modifying this field for an existing workflow results in a new workflow
	//  revision.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Workflow code to be executed. The size limit is 32KB.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.source_contents
	SourceContents *string `json:"sourceContents,omitempty"`
}

// +kcc:proto=google.cloud.workflows.v1beta.Workflow
type WorkflowObservedState struct {
	// Output only. State of the workflow deployment.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.state
	State *string `json:"state,omitempty"`

	// Output only. The revision of the workflow.
	//  A new revision of a workflow is created as a result of updating the
	//  following fields of a workflow:
	//  - `source_code`
	//  - `service_account`
	//  The format is "000001-a4d", where the first 6 characters define
	//  the zero-padded revision ordinal number. They are followed by a hyphen and
	//  3 hexadecimal random characters.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp of when the workflow was created.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the workflow.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The timestamp that the latest revision of the workflow
	//  was created.
	// +kcc:proto:field=google.cloud.workflows.v1beta.Workflow.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`
}
