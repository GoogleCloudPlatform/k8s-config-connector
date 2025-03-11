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


// +kcc:proto=google.cloud.edgecontainer.v1.Machine
type Machine struct {
	// Required. The resource name of the machine.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Canonical resource name of the node that this machine is responsible for
	//  hosting e.g.
	//  projects/{project}/locations/{location}/clusters/{cluster_id}/nodePools/{pool_id}/{node},
	//  Or empty if the machine is not assigned to assume the role of a node.
	//
	//  For control plane nodes hosted on edge machines, this will return
	//  the following format:
	//    "projects/{project}/locations/{location}/clusters/{cluster_id}/controlPlaneNodes/{node}".
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.hosted_node
	HostedNode *string `json:"hostedNode,omitempty"`

	// The Google Distributed Cloud Edge zone of this machine.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Machine
type MachineObservedState struct {
	// Output only. The time when the node pool was created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the node pool was last updated.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The software version of the machine.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.version
	Version *string `json:"version,omitempty"`

	// Output only. Whether the machine is disabled. If disabled, the machine is
	//  unable to enter service.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.disabled
	Disabled *bool `json:"disabled,omitempty"`
}
