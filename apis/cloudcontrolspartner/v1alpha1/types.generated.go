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


// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Violation
type Violation struct {
	// Identifier. Format:
	//  `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}/violations/{violation}`
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.name
	Name *string `json:"name,omitempty"`

	// The folder_id of the violation
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.folder_id
	FolderID *int64 `json:"folderID,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Violation.Remediation
type Violation_Remediation struct {
	// Required. Remediation instructions to resolve violations
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.instructions
	Instructions *Violation_Remediation_Instructions `json:"instructions,omitempty"`

	// Values that can resolve the violation
	//  For example: for list org policy violations, this will either be the list
	//  of allowed or denied values
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.compliant_values
	CompliantValues []string `json:"compliantValues,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions
type Violation_Remediation_Instructions struct {
	// Remediation instructions to resolve violation via gcloud cli
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.gcloud_instructions
	GcloudInstructions *Violation_Remediation_Instructions_Gcloud `json:"gcloudInstructions,omitempty"`

	// Remediation instructions to resolve violation via cloud console
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.console_instructions
	ConsoleInstructions *Violation_Remediation_Instructions_Console `json:"consoleInstructions,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Console
type Violation_Remediation_Instructions_Console struct {
	// Link to console page where violations can be resolved
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Console.console_uris
	ConsoleUris []string `json:"consoleUris,omitempty"`

	// Steps to resolve violation via cloud console
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Console.steps
	Steps []string `json:"steps,omitempty"`

	// Additional urls for more information about steps
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Console.additional_links
	AdditionalLinks []string `json:"additionalLinks,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Gcloud
type Violation_Remediation_Instructions_Gcloud struct {
	// Gcloud command to resolve violation
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Gcloud.gcloud_commands
	GcloudCommands []string `json:"gcloudCommands,omitempty"`

	// Steps to resolve violation via gcloud cli
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Gcloud.steps
	Steps []string `json:"steps,omitempty"`

	// Additional urls for more information about steps
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.Instructions.Gcloud.additional_links
	AdditionalLinks []string `json:"additionalLinks,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Violation
type ViolationObservedState struct {
	// Output only. Description for the Violation.
	//  e.g. OrgPolicy gcp.resourceLocations has non compliant value.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.description
	Description *string `json:"description,omitempty"`

	// Output only. Time of the event which triggered the Violation.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.begin_time
	BeginTime *string `json:"beginTime,omitempty"`

	// Output only. The last time when the Violation record was updated.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time of the event which fixed the Violation.
	//  If the violation is ACTIVE this will be empty.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.resolve_time
	ResolveTime *string `json:"resolveTime,omitempty"`

	// Output only. Category under which this violation is mapped.
	//  e.g. Location, Service Usage, Access, Encryption, etc.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.category
	Category *string `json:"category,omitempty"`

	// Output only. State of the violation
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.state
	State *string `json:"state,omitempty"`

	// Output only. Immutable. Name of the OrgPolicy which was modified with
	//  non-compliant change and resulted this violation. Format:
	//   `projects/{project_number}/policies/{constraint_name}`
	//   `folders/{folder_id}/policies/{constraint_name}`
	//   `organizations/{organization_id}/policies/{constraint_name}`
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.non_compliant_org_policy
	NonCompliantOrgPolicy *string `json:"nonCompliantOrgPolicy,omitempty"`

	// Output only. Compliance violation remediation
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.remediation
	Remediation *Violation_Remediation `json:"remediation,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Violation.Remediation
type Violation_RemediationObservedState struct {
	// Output only. Remediation type based on the type of org policy values
	//  violated
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Violation.Remediation.remediation_type
	RemediationType *string `json:"remediationType,omitempty"`
}
