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

// API sources for NetappVolume, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/netapp/v1/volume.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.volumes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetappVolumeGVK = GroupVersion.WithKind("NetappVolume")

// NetappVolumeSpec defines the desired state of NetappVolume
// +kcc:spec:proto=google.cloud.netapp.v1.Volume
// +kcc:required-from-proto
type NetappVolumeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetappVolume name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Share name of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.share_name
	// +required
	ShareName *string `json:"shareName,omitempty"`

	// Required. StoragePool name of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.storage_pool
	// +required
	StoragePool *string `json:"storagePool,omitempty"`

	// Required. Capacity in GIB of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.capacity_gib
	// +required
	CapacityGib *int64 `json:"capacityGib,omitempty"`

	// Optional. Export policy of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.export_policy
	ExportPolicy *ExportPolicy `json:"exportPolicy,omitempty"`

	// Required. Protocols required for the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.protocols
	// +required
	Protocols []string `json:"protocols,omitempty"`

	// Optional. SMB share settings for the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.smb_settings
	SmbSettings []string `json:"smbSettings,omitempty"`

	// Optional. Default unix style permission (e.g. 777) the mount point will be
	//  created with. Applicable for NFS protocol types only.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.unix_permissions
	UnixPermissions *string `json:"unixPermissions,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.description
	Description *string `json:"description,omitempty"`

	// Optional. SnapshotPolicy for a volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.snapshot_policy
	SnapshotPolicy *SnapshotPolicy `json:"snapshotPolicy,omitempty"`

	// Optional. Snap_reserve specifies percentage of volume storage reserved for
	//  snapshot storage. Default is 0 percent.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.snap_reserve
	SnapReserve *float64 `json:"snapReserve,omitempty"`

	// Optional. Snapshot_directory if enabled (true) the volume will contain a
	//  read-only .snapshot directory which provides access to each of the volume's
	//  snapshots.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.snapshot_directory
	SnapshotDirectory *bool `json:"snapshotDirectory,omitempty"`

	// Optional. Security Style of the Volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.security_style
	SecurityStyle *string `json:"securityStyle,omitempty"`

	// Optional. Flag indicating if the volume is a kerberos volume or not, export
	//  policy rules control kerberos security modes (krb5, krb5i, krb5p).
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.kerberos_enabled
	KerberosEnabled *bool `json:"kerberosEnabled,omitempty"`

	// Optional. Specifies the source of the volume to be created from.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.restore_parameters
	RestoreParameters *RestoreParameters `json:"restoreParameters,omitempty"`

	// BackupConfig of the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.backup_config
	BackupConfig *BackupConfig `json:"backupConfig,omitempty"`

	// Optional. List of actions that are restricted on this volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.restricted_actions
	RestrictedActions []string `json:"restrictedActions,omitempty"`

	// Optional. Flag indicating if the volume will be a large capacity volume or
	//  a regular volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.large_capacity
	LargeCapacity *bool `json:"largeCapacity,omitempty"`

	// Optional. Flag indicating if the volume will have an IP address per node
	//  for volumes supporting multiple IP endpoints. Only the volume with
	//  large_capacity will be allowed to have multiple endpoints.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.multiple_endpoints
	MultipleEndpoints *bool `json:"multipleEndpoints,omitempty"`

	// Tiering policy for the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.tiering_policy
	TieringPolicy *TieringPolicy `json:"tieringPolicy,omitempty"`

	// Optional. The Hybrid Replication parameters for the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.hybrid_replication_parameters
	HybridReplicationParameters *HybridReplicationParameters `json:"hybridReplicationParameters,omitempty"`
}

// NetappVolumeStatus defines the config connector machine state of NetappVolume
type NetappVolumeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetappVolume resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetappVolumeObservedState `json:"observedState,omitempty"`
}

// NetappVolumeObservedState is the state of the NetappVolume resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.Volume
type NetappVolumeObservedState struct {
	// Output only. State of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Create time of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. This field is not implemented. The values provided in this
	//  field are ignored.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.psa_range
	PsaRange *string `json:"psaRange,omitempty"`

	// Output only. VPC Network name.
	//  Format: projects/{project}/global/networks/{network}
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.network
	Network *string `json:"network,omitempty"`

	// Output only. Service level of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.service_level
	ServiceLevel *string `json:"serviceLevel,omitempty"`

	// Output only. Mount options of this volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.mount_options
	MountOptions []MountOptionObservedState `json:"mountOptions,omitempty"`

	// Output only. Used capacity in GIB of the volume. This is computed
	//  periodically and it does not represent the realtime usage.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.used_gib
	UsedGib *int64 `json:"usedGib,omitempty"`

	// Output only. Flag indicating if the volume is NFS LDAP enabled or not.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.ldap_enabled
	LdapEnabled *bool `json:"ldapEnabled,omitempty"`

	// Output only. Specifies the ActiveDirectory name of a SMB volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.active_directory
	ActiveDirectory *string `json:"activeDirectory,omitempty"`

	// Output only. Specifies the KMS config to be used for volume encryption.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.kms_config
	KMSConfig *string `json:"kmsConfig,omitempty"`

	// Output only. Specified the current volume encryption key source.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Output only. Indicates whether the volume is part of a replication
	//  relationship.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.has_replication
	HasReplication *bool `json:"hasReplication,omitempty"`

	// BackupConfig of the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.backup_config
	BackupConfig *BackupConfigObservedState `json:"backupConfig,omitempty"`

	// Output only. Specifies the replica zone for regional volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.replica_zone
	ReplicaZone *string `json:"replicaZone,omitempty"`

	// Output only. Specifies the active zone for regional volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.zone
	Zone *string `json:"zone,omitempty"`

	// Output only. Size of the volume cold tier data in GiB.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.cold_tier_size_gib
	ColdTierSizeGib *int64 `json:"coldTierSizeGib,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappvolume;gcpnetappvolumes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetappVolume is the Schema for the NetappVolume API
// +k8s:openapi-gen=true
type NetappVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetappVolumeSpec   `json:"spec,omitempty"`
	Status NetappVolumeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetappVolumeList contains a list of NetappVolume
type NetappVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappVolume `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetappVolume{}, &NetappVolumeList{})
}
