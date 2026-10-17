// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for NetappStoragePool, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/netapp/v1/storage_pool.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.storagePools

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetappStoragePoolGVK = GroupVersion.WithKind("NetappStoragePool")

// NetappStoragePoolSpec defines the desired state of NetappStoragePool
// +kcc:spec:proto=google.cloud.netapp.v1.StoragePool
// +kcc:required-from-proto
type NetappStoragePoolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetappStoragePool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Service level of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.service_level
	// +required
	ServiceLevel *string `json:"serviceLevel,omitempty"`

	// Required. Capacity in GIB of the pool
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.capacity_gib
	// +required
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
	// +required
	Network *string `json:"network,omitempty"`

	// Optional. Specifies the Active Directory to be used for creating a SMB
	//  volume.
	// +kcc:guess=possible-reference target=NetAppActiveDirectory
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.active_directory
	ActiveDirectory *string `json:"activeDirectory,omitempty"`

	// Optional. Specifies the KMS config to be used for volume encryption.
	// +kcc:guess=possible-reference target=NetappKmsConfig
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.kms_config
	KMSConfig *string `json:"kmsConfig,omitempty"`

	// Optional. Flag indicating if the pool is NFS LDAP enabled or not.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.ldap_enabled
	LdapEnabled *bool `json:"ldapEnabled,omitempty"`

	// Optional. This field is not implemented. The values provided in this field
	//  are ignored.
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.psa_range
	PsaRange *string `json:"psaRange,omitempty"`

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

	// Optional. True if using Independent Scaling of capacity and performance
	//  (Hyperdisk) By default set to false
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.custom_performance_enabled
	CustomPerformanceEnabled *bool `json:"customPerformanceEnabled,omitempty"`

	// Optional. Custom Performance Total Throughput of the pool (in MiB/s)
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.total_throughput_mibps
	TotalThroughputMibps *int64 `json:"totalThroughputMibps,omitempty"`

	// Optional. Custom Performance Total IOPS of the pool
	//  If not provided, it will be calculated based on the total_throughput_mibps
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.total_iops
	TotalIops *int64 `json:"totalIops,omitempty"`
}

// NetappStoragePoolStatus defines the config connector machine state of NetappStoragePool
type NetappStoragePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetappStoragePool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetappStoragePoolObservedState `json:"observedState,omitempty"`
}

// NetappStoragePoolObservedState is the state of the NetappStoragePool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.StoragePool
type NetappStoragePoolObservedState struct {
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

	// Output only. Reserved for future use
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use
	// +kcc:proto:field=google.cloud.netapp.v1.StoragePool.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappstoragepool;gcpnetappstoragepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetappStoragePool is the Schema for the NetappStoragePool API
// +k8s:openapi-gen=true
type NetappStoragePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetappStoragePoolSpec   `json:"spec,omitempty"`
	Status NetappStoragePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetappStoragePoolList contains a list of NetappStoragePool
type NetappStoragePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappStoragePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetappStoragePool{}, &NetappStoragePoolList{})
}
