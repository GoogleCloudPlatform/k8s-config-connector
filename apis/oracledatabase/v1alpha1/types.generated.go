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


// +kcc:proto=google.cloud.oracledatabase.v1.DbNode
type DbNode struct {
	// Identifier. The name of the database node resource in the following format:
	//  projects/{project}/locations/{location}/cloudVmClusters/{cloud_vm_cluster}/dbNodes/{db_node}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNode.name
	Name *string `json:"name,omitempty"`

	// Optional. Various properties of the database node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNode.properties
	Properties *DbNodeProperties `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DbNodeProperties
type DbNodeProperties struct {

	// Optional. OCPU count per database node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.ocpu_count
	OcpuCount *int32 `json:"ocpuCount,omitempty"`

	// Memory allocated in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.memory_size_gb
	MemorySizeGB *int32 `json:"memorySizeGB,omitempty"`

	// Optional. Local storage per database node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.db_node_storage_size_gb
	DbNodeStorageSizeGB *int32 `json:"dbNodeStorageSizeGB,omitempty"`

	// Optional. Database server OCID.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.db_server_ocid
	DbServerOcid *string `json:"dbServerOcid,omitempty"`

	// Optional. DNS
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Total CPU core count of the database node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.total_cpu_core_count
	TotalCpuCoreCount *int32 `json:"totalCpuCoreCount,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DbNode
type DbNodeObservedState struct {
	// Optional. Various properties of the database node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNode.properties
	Properties *DbNodePropertiesObservedState `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DbNodeProperties
type DbNodePropertiesObservedState struct {
	// Output only. OCID of database node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.ocid
	Ocid *string `json:"ocid,omitempty"`

	// Output only. State of the database node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbNodeProperties.state
	State *string `json:"state,omitempty"`
}
