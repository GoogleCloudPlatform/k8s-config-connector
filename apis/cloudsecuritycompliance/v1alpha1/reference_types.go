// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.TargetResourceConfig
type TargetResourceConfig struct {
	// Optional. CRM node in format organizations/{organization},
	//  folders/{folder}, projects/{project} or
	//  projects/{project}/locations/{location}/applications/{application}.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceConfig.existing_target_resource
	ExistingTargetResourceRef *CloudSecurityComplianceTargetResourceRef `json:"existingTargetResourceRef,omitempty"`

	// Optional. Config to create a new resource and use that as the
	//  target_resource for deployment
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceConfig.target_resource_creation_config
	TargetResourceCreationConfig *TargetResourceCreationConfig `json:"targetResourceCreationConfig,omitempty"`
}

type CloudSecurityComplianceTargetResourceRef struct {
	// A reference to an externally managed target resource.
	// Can be in format organizations/{organization}, folders/{folder}, projects/{project} or projects/{project}/locations/{location}/applications/{application}.
	External string `json:"external,omitempty"`
}

type CloudSecurityComplianceParentRef struct {
	// A reference to an Organization or Folder.
	// Can be in format folders/{folder_id} or organizations/{org_id}.
	External string `json:"external,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.FolderCreationConfig
type FolderCreationConfig struct {
	// Required. organizations/{org} or folders/{folder}
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FolderCreationConfig.parent
	ParentRef *CloudSecurityComplianceParentRef `json:"parentRef,omitempty"`

	// Required. Display name of the folder to be created
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FolderCreationConfig.folder_display_name
	FolderDisplayName *string `json:"folderDisplayName,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig
type ProjectCreationConfig struct {
	// Required. organizations/{org} or folders/{folder}
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig.parent
	ParentRef *CloudSecurityComplianceParentRef `json:"parentRef,omitempty"`

	// Required. Display name of the project to be created
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig.project_display_name
	ProjectDisplayName *string `json:"projectDisplayName,omitempty"`

	// Required. Billing account id to be used for the project
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig.billing_account_id
	BillingAccountID *string `json:"billingAccountID,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.FrameworkReference
type FrameworkReference struct {
	// Required. Framework resource reference
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkReference.framework
	FrameworkRef *CloudSecurityComplianceFrameworkRef `json:"frameworkRef,omitempty"`

	// Optional. Major revision id of the framework. If not specified, corresponds
	//  to the latest revision of the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkReference.major_revision_id
	MajorRevisionID *int64 `json:"majorRevisionID,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlMetadata
type CloudControlMetadata struct {
	// Required. Cloud control details
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlMetadata.cloud_control_details
	CloudControlDetails *CloudControlDetails `json:"cloudControlDetails,omitempty"`

	// Required. Enforcement mode of the cloud control
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlMetadata.enforcement_mode
	EnforcementMode *string `json:"enforcementMode,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.TargetResourceCreationConfig
type TargetResourceCreationConfig struct {
	// Optional. Folder creation config
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceCreationConfig.folder_creation_config
	FolderCreationConfig *FolderCreationConfig `json:"folderCreationConfig,omitempty"`

	// Optional. Project creation config
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceCreationConfig.project_creation_config
	ProjectCreationConfig *ProjectCreationConfig `json:"projectCreationConfig,omitempty"`
}
