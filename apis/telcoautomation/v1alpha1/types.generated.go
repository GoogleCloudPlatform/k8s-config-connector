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

// +kcc:proto=google.cloud.telcoautomation.v1.HydratedDeployment
type HydratedDeployment struct {

	// Optional. File contents of a hydrated deployment.
	//  When invoking UpdateHydratedBlueprint API, only the modified files should
	//  be included in this. Files that are not included in the update of a
	//  hydrated deployment will not be changed.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.HydratedDeployment.files
	Files []File `json:"files,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.HydratedDeployment
type HydratedDeploymentObservedState struct {
	// Output only. The name of the hydrated deployment.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.HydratedDeployment.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the hydrated deployment (DRAFT, APPLIED).
	// +kcc:proto:field=google.cloud.telcoautomation.v1.HydratedDeployment.state
	State *string `json:"state,omitempty"`

	// Output only. WorkloadCluster identifies which workload cluster will the
	//  hydrated deployment will be deployed on.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.HydratedDeployment.workload_cluster
	WorkloadCluster *string `json:"workloadCluster,omitempty"`
}
