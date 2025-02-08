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


// +kcc:proto=google.cloud.vmwareengine.v1.NodeType
type NodeType struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.NodeType
type NodeTypeObservedState struct {
	// Output only. The resource name of this node type.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-proj/locations/us-central1-a/nodeTypes/standard-72`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.name
	Name *string `json:"name,omitempty"`

	// Output only. The canonical identifier of the node type
	//  (corresponds to the `NodeType`). For example: standard-72.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.node_type_id
	NodeTypeID *string `json:"nodeTypeID,omitempty"`

	// Output only. The friendly name for this node type.
	//  For example: ve1-standard-72
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The total number of virtual CPUs in a single node.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.virtual_cpu_count
	VirtualCpuCount *int32 `json:"virtualCpuCount,omitempty"`

	// Output only. The total number of CPU cores in a single node.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.total_core_count
	TotalCoreCount *int32 `json:"totalCoreCount,omitempty"`

	// Output only. The amount of physical memory available, defined in GB.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.memory_gb
	MemoryGB *int32 `json:"memoryGB,omitempty"`

	// Output only. The amount of storage available, defined in GB.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.disk_size_gb
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`

	// Output only. List of possible values of custom core count.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.available_custom_core_counts
	AvailableCustomCoreCounts []int32 `json:"availableCustomCoreCounts,omitempty"`

	// Output only. The type of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.kind
	Kind *string `json:"kind,omitempty"`

	// Output only. Families of the node type.
	//  For node types to be in the same cluster
	//  they must share at least one element in the `families`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.families
	Families []string `json:"families,omitempty"`

	// Output only. Capabilities of this node type.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeType.capabilities
	Capabilities []string `json:"capabilities,omitempty"`
}
