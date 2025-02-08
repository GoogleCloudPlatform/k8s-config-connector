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


// +kcc:proto=google.cloud.oracledatabase.v1.DbServer
type DbServer struct {
	// Identifier. The name of the database server resource with the format:
	//  projects/{project}/locations/{location}/cloudExadataInfrastructures/{cloud_exadata_infrastructure}/dbServers/{db_server}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServer.name
	Name *string `json:"name,omitempty"`

	// Optional. User friendly name for this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServer.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Various properties of the database server.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServer.properties
	Properties *DbServerProperties `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DbServerProperties
type DbServerProperties struct {

	// Optional. OCPU count per database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.ocpu_count
	OcpuCount *int32 `json:"ocpuCount,omitempty"`

	// Optional. Maximum OCPU count per database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.max_ocpu_count
	MaxOcpuCount *int32 `json:"maxOcpuCount,omitempty"`

	// Optional. Memory allocated in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.memory_size_gb
	MemorySizeGB *int32 `json:"memorySizeGB,omitempty"`

	// Optional. Maximum memory allocated in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.max_memory_size_gb
	MaxMemorySizeGB *int32 `json:"maxMemorySizeGB,omitempty"`

	// Optional. Local storage per VM.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.db_node_storage_size_gb
	DbNodeStorageSizeGB *int32 `json:"dbNodeStorageSizeGB,omitempty"`

	// Optional. Maximum local storage per VM.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.max_db_node_storage_size_gb
	MaxDbNodeStorageSizeGB *int32 `json:"maxDbNodeStorageSizeGB,omitempty"`

	// Optional. Vm count per database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.vm_count
	VmCount *int32 `json:"vmCount,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DbServer
type DbServerObservedState struct {
	// Optional. Various properties of the database server.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServer.properties
	Properties *DbServerPropertiesObservedState `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DbServerProperties
type DbServerPropertiesObservedState struct {
	// Output only. OCID of database server.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.ocid
	Ocid *string `json:"ocid,omitempty"`

	// Output only. State of the database server.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.state
	State *string `json:"state,omitempty"`

	// Output only. OCID of database nodes associated with the database server.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbServerProperties.db_node_ids
	DbNodeIds []string `json:"dbNodeIds,omitempty"`
}
