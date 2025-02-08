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


// +kcc:proto=google.cloud.telcoautomation.v1.Deployment
type Deployment struct {
	// The name of the deployment.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.name
	Name *string `json:"name,omitempty"`

	// Required. The blueprint revision from which this deployment was created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.source_blueprint_revision
	SourceBlueprintRevision *string `json:"sourceBlueprintRevision,omitempty"`

	// Optional. Human readable name of a Deployment.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Files present in a deployment.
	//  When invoking UpdateDeployment API, only the modified files should be
	//  included in this. Files that are not included in the update of a deployment
	//  will not be changed.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.files
	Files []File `json:"files,omitempty"`

	// Optional. Labels are key-value attributes that can be set on a deployment
	//  resource by the user.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Immutable. The WorkloadCluster on which to create the Deployment.
	//  This field should only be passed when the deployment_level of the source
	//  blueprint specifies deployments on workload clusters e.g.
	//  WORKLOAD_CLUSTER_DEPLOYMENT.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.workload_cluster
	WorkloadCluster *string `json:"workloadCluster,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.File
type File struct {
	// Required. Path of the file in package.
	//  e.g. `gdce/v1/cluster.yaml`
	// +kcc:proto:field=google.cloud.telcoautomation.v1.File.path
	Path *string `json:"path,omitempty"`

	// Optional. The contents of a file in string format.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.File.content
	Content *string `json:"content,omitempty"`

	// Optional. Signifies whether a file is marked for deletion.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.File.deleted
	Deleted *bool `json:"deleted,omitempty"`

	// Optional. Indicates whether changes are allowed to a file. If the field is
	//  not set, the file cannot be edited.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.File.editable
	Editable *bool `json:"editable,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.Deployment
type DeploymentObservedState struct {
	// Output only. Immutable. The revision ID of the deployment.
	//  A new revision is committed whenever a change in deployment is applied.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. State of the deployment (DRAFT, APPLIED, DELETING).
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.state
	State *string `json:"state,omitempty"`

	// Output only. Name of the repository where the deployment package files are
	//  stored.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.repository
	Repository *string `json:"repository,omitempty"`

	// Output only. Deployment creation time.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the deployment was updated.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Source provider is the author of a public blueprint, from
	//  which this deployment is created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.source_provider
	SourceProvider *string `json:"sourceProvider,omitempty"`

	// Output only. Attributes to where the deployment can inflict changes. The
	//  value can only be [SINGLE_DEPLOYMENT, MULTI_DEPLOYMENT].
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.deployment_level
	DeploymentLevel *string `json:"deploymentLevel,omitempty"`

	// Output only. Indicates if the deployment can be rolled back, exported from
	//  public blueprint.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.rollback_support
	RollbackSupport *bool `json:"rollbackSupport,omitempty"`
}
