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


// +kcc:proto=google.cloud.vmwareengine.v1.Node
type Node struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.Node
type NodeObservedState struct {
	// Output only. The resource name of this node.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  projects/my-project/locations/us-central1-a/privateClouds/my-cloud/clusters/my-cluster/nodes/my-node
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Node.name
	Name *string `json:"name,omitempty"`

	// Output only. Fully qualified domain name of the node.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Node.fqdn
	Fqdn *string `json:"fqdn,omitempty"`

	// Output only. Internal IP address of the node.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Node.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// Output only. The canonical identifier of the node type (corresponds to the
	//  `NodeType`).
	//  For example: standard-72.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Node.node_type_id
	NodeTypeID *string `json:"nodeTypeID,omitempty"`

	// Output only. The version number of the VMware ESXi
	//  management component in this cluster.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Node.version
	Version *string `json:"version,omitempty"`

	// Output only. Customized number of cores
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Node.custom_core_count
	CustomCoreCount *int64 `json:"customCoreCount,omitempty"`

	// Output only. The state of the appliance.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Node.state
	State *string `json:"state,omitempty"`
}
