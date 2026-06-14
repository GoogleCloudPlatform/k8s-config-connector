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

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails
type CloudControlDetailsObservedState struct {
	// Required. The name of the CloudControl in the format:
	//  “organizations/{organization}/locations/{location}/
	//  cloudControls/{cloud-control}”
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.name
	CloudControl *string `json:"cloudControl,omitempty"`

	// Required. Major revision of cloudcontrol
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.major_revision_id
	MajorRevisionID *int64 `json:"majorRevisionID,omitempty"`

	// Optional. Parameters is a key-value pair that is required by the
	//  CloudControl. Eg: { "name": "location","value": "us-west-1"}.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.parameters
	Parameters []Parameter `json:"parameters,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlMetadata
type CloudControlMetadataObservedState struct {
	// Required. CloudControlReference, Deployment mode and parameters for the
	//  cloud_control
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlMetadata.cloud_control_details
	CloudControlDetails *CloudControlDetailsObservedState `json:"cloudControlDetails,omitempty"`

	// Required. Enforcement mode of the cloud control
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlMetadata.enforcement_mode
	EnforcementMode *string `json:"enforcementMode,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment
type CloudControlDeploymentObservedState struct {
	// Identifier. CloudControlDeployment name in either of the following formats:
	//  organizations/{organization}/locations/{location}/cloudControlDeployments/{cloud_control_deployment_id}
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.name
	Name *string `json:"name,omitempty"`

	// Required. target_resource_config referencing either an already existing
	//  target_resource or contains config for a target_resource to be created
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.target_resource_config
	TargetResourceConfig *TargetResourceConfigObservedState `json:"targetResourceConfig,omitempty"`

	// Output only. The resource on which the CloudControl is deployed based on
	//  the provided TargetResourceConfig.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.target_resource
	TargetResource *string `json:"targetResource,omitempty"`

	// Required. CloudControlReference, Deployment mode and parameters for the
	//  cloud_control
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.cloud_control_metadata
	CloudControlMetadata *CloudControlMetadataObservedState `json:"cloudControlMetadata,omitempty"`

	// Optional. User provided description of the deployment
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.description
	Description *string `json:"description,omitempty"`

	// Output only. State of the deployment
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.deployment_state
	DeploymentState *string `json:"deploymentState,omitempty"`

	// Output only. The time at which the resource was created.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the resource last updated.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDeployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlGroupDeployment
type CloudControlGroupDeploymentObservedState struct {
	// Required. Cloud control group
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroupDeployment.cloud_control_group
	CloudControlGroup *CloudControlGroupObservedState `json:"cloudControlGroup,omitempty"`

	// Required. Cloud control deployments in the group
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroupDeployment.cc_deployments
	CcDeployments []CloudControlDeploymentObservedState `json:"ccDeployments,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.TargetResourceConfig
type TargetResourceConfigObservedState struct {
	// Optional. CRM node in format organizations/{organization},
	//  folders/{folder}, projects/{project} or
	//  projects/{project}/locations/{location}/applications/{application}.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceConfig.existing_target_resource
	ExistingTargetResource *string `json:"existingTargetResource,omitempty"`

	// Optional. Config to create a new resource and use that as the
	//  target_resource for deployment
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceConfig.target_resource_creation_config
	TargetResourceCreationConfig *TargetResourceCreationConfigObservedState `json:"targetResourceCreationConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.TargetResourceCreationConfig
type TargetResourceCreationConfigObservedState struct {
	// Optional. Folder creation config
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceCreationConfig.folder_creation_config
	FolderCreationConfig *FolderCreationConfigObservedState `json:"folderCreationConfig,omitempty"`

	// Optional. Project creation config
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.TargetResourceCreationConfig.project_creation_config
	ProjectCreationConfig *ProjectCreationConfigObservedState `json:"projectCreationConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.FolderCreationConfig
type FolderCreationConfigObservedState struct {
	// Required. organizations/{org} or folders/{folder}
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FolderCreationConfig.parent
	Parent *string `json:"parent,omitempty"`

	// Required. Display name of the folder to be created
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FolderCreationConfig.folder_display_name
	FolderDisplayName *string `json:"folderDisplayName,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig
type ProjectCreationConfigObservedState struct {
	// Required. organizations/{org} or folders/{folder}
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig.parent
	Parent *string `json:"parent,omitempty"`

	// Required. Display name of the project to be created
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig.project_display_name
	ProjectDisplayName *string `json:"projectDisplayName,omitempty"`

	// Required. Billing account id to be used for the project
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ProjectCreationConfig.billing_account_id
	BillingAccountID *string `json:"billingAccountID,omitempty"`
}
