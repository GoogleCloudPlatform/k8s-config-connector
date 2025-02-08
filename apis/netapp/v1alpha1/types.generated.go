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


// +kcc:proto=google.cloud.netapp.v1.StoragePool
type StoragePool struct {
	// Identifier. Name of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.name
	Name *string `json:"name,omitempty"`

	// Required. Service level of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.service_level
	ServiceLevel *string `json:"serviceLevel,omitempty"`

	// Required. Capacity in GIB of the pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.capacity_gib
	CapacityGib *int64 `json:"capacityGib,omitempty"`

	// Optional. Description of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. VPC Network name.
	//  Format: projects/{project}/global/networks/{network}
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.network
	Network *string `json:"network,omitempty"`

	// Optional. Specifies the Active Directory to be used for creating a SMB
	//  volume.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.active_directory
	ActiveDirectory *string `json:"activeDirectory,omitempty"`

	// Optional. Specifies the KMS config to be used for volume encryption.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.kms_config
	KMSConfig *string `json:"kmsConfig,omitempty"`

	// Optional. Flag indicating if the pool is NFS LDAP enabled or not.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.ldap_enabled
	LdapEnabled *bool `json:"ldapEnabled,omitempty"`

	// Optional. This field is not implemented. The values provided in this field
	//  are ignored.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.psa_range
	PsaRange *string `json:"psaRange,omitempty"`

	// Deprecated. Used to allow SO pool to access AD or DNS server from other
	//  regions.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.global_access_allowed
	GlobalAccessAllowed *bool `json:"globalAccessAllowed,omitempty"`

	// Optional. True if the storage pool supports Auto Tiering enabled volumes.
	//  Default is false. Auto-tiering can be enabled after storage pool creation
	//  but it can't be disabled once enabled.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.allow_auto_tiering
	AllowAutoTiering *bool `json:"allowAutoTiering,omitempty"`

	// Optional. Specifies the replica zone for regional storagePool.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.replica_zone
	ReplicaZone *string `json:"replicaZone,omitempty"`

	// Optional. Specifies the active zone for regional storagePool.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.StoragePool
type StoragePoolObservedState struct {
	// Output only. Allocated size of all volumes in GIB in the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.volume_capacity_gib
	VolumeCapacityGib *int64 `json:"volumeCapacityGib,omitempty"`

	// Output only. Volume count of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.volume_count
	VolumeCount *int32 `json:"volumeCount,omitempty"`

	// Output only. State of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Create time of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Specifies the current pool encryption key source.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`
}
