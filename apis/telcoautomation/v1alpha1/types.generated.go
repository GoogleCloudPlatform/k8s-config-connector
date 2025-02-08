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


// +kcc:proto=google.cloud.telcoautomation.v1.EdgeSlm
type EdgeSlm struct {
	// Name of the EdgeSlm resource.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.name
	Name *string `json:"name,omitempty"`

	// Immutable. Reference to the orchestration cluster on which templates for
	//  this resources will be applied. This should be of format
	//  projects/{project}/locations/{location}/orchestrationClusters/{orchestration_cluster}.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.orchestration_cluster
	OrchestrationCluster *string `json:"orchestrationCluster,omitempty"`

	// Optional. Labels as key value pairs. The key and value should contain
	//  characters which are UTF-8 compliant and less than 50 characters.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Type of workload cluster for which an EdgeSLM resource is
	//  created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.workload_cluster_type
	WorkloadClusterType *string `json:"workloadClusterType,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.EdgeSlm
type EdgeSlmObservedState struct {
	// Output only. [Output only] Create time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Provides the active TNA version for this resource.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.tna_version
	TnaVersion *string `json:"tnaVersion,omitempty"`

	// Output only. State of the EdgeSlm resource.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.state
	State *string `json:"state,omitempty"`
}
