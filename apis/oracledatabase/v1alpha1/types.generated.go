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


// +kcc:proto=google.cloud.oracledatabase.v1.DbSystemShape
type DbSystemShape struct {
	// Identifier. The name of the Database System Shape resource with the format:
	//  projects/{project}/locations/{region}/dbSystemShapes/{db_system_shape}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.name
	Name *string `json:"name,omitempty"`

	// Optional. shape
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.shape
	Shape *string `json:"shape,omitempty"`

	// Optional. Minimum number of database servers.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// Optional. Maximum number of database servers.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	// Optional. Minimum number of storage servers.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.min_storage_count
	MinStorageCount *int32 `json:"minStorageCount,omitempty"`

	// Optional. Maximum number of storage servers.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.max_storage_count
	MaxStorageCount *int32 `json:"maxStorageCount,omitempty"`

	// Optional. Number of cores per node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.available_core_count_per_node
	AvailableCoreCountPerNode *int32 `json:"availableCoreCountPerNode,omitempty"`

	// Optional. Memory per database server node in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.available_memory_per_node_gb
	AvailableMemoryPerNodeGB *int32 `json:"availableMemoryPerNodeGB,omitempty"`

	// Optional. Storage per storage server in terabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.available_data_storage_tb
	AvailableDataStorageTb *int32 `json:"availableDataStorageTb,omitempty"`

	// Optional. Minimum core count per node.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.min_core_count_per_node
	MinCoreCountPerNode *int32 `json:"minCoreCountPerNode,omitempty"`

	// Optional. Minimum memory per node in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.min_memory_per_node_gb
	MinMemoryPerNodeGB *int32 `json:"minMemoryPerNodeGB,omitempty"`

	// Optional. Minimum node storage per database server in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DbSystemShape.min_db_node_storage_per_node_gb
	MinDbNodeStoragePerNodeGB *int32 `json:"minDbNodeStoragePerNodeGB,omitempty"`
}
