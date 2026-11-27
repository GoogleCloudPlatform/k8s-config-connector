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

// API sources for NetAppVolume, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e09e85d32ca349e1b205514a412817a7692595c6/google/cloud/netapp/v1/volume.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.volumes

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetAppVolumeGVK = GroupVersion.WithKind("NetAppVolume")

// NetAppVolumeSpec defines the desired state of NetAppVolume
// +kcc:spec:proto=google.cloud.netapp.v1.Volume
// +kcc:required-from-proto
type NetAppVolumeSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The NetAppVolume name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Share name of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.share_name
	// +required
	ShareName *string `json:"shareName,omitempty"`

	// Required. StoragePool name of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.storage_pool
	// +required
	StoragePoolRef *NetAppStoragePoolRef `json:"storagePoolRef,omitempty"`

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
	//  a regular volume. This field is used for legacy FILE pools. For Unified
	//  pools, use the `large_capacity_config` field instead. This field and
	//  `large_capacity_config` are mutually exclusive.
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

	// Optional. Throughput of the volume (in MiB/s)
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.throughput_mibps
	ThroughputMibps *float64 `json:"throughputMibps,omitempty"`

	// Optional. Cache parameters for the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.cache_parameters
	CacheParameters *CacheParameters `json:"cacheParameters,omitempty"`

	// Optional. Block devices for the volume.
	//  Currently, only one block device is permitted per Volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.block_devices
	BlockDevices []BlockDevice `json:"blockDevices,omitempty"`

	// Optional. Large capacity config for the volume.
	//  Enables and configures large capacity for volumes in Unified pools with
	//  File protocols. Not applicable for Block protocols in Unified pools.
	//  This field and the legacy `large_capacity` boolean field
	//  are mutually exclusive.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.large_capacity_config
	LargeCapacityConfig *LargeCapacityConfig `json:"largeCapacityConfig,omitempty"`
}

// NetAppVolumeStatus defines the config connector machine state of NetAppVolume
type NetAppVolumeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetAppVolume resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetAppVolumeObservedState `json:"observedState,omitempty"`
}

// NetAppVolumeObservedState is the state of the NetAppVolume resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.Volume
type NetAppVolumeObservedState struct {
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

	// Output only. Size of the volume cold tier data rounded down to the nearest
	//  GiB.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.cold_tier_size_gib
	ColdTierSizeGib *int64 `json:"coldTierSizeGib,omitempty"`

	// Optional. Cache parameters for the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.cache_parameters
	CacheParameters *CacheParametersObservedState `json:"cacheParameters,omitempty"`

	// Output only. Total hot tier data rounded down to the nearest GiB used by
	//  the Volume. This field is only used for flex Service Level
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.hot_tier_size_used_gib
	HotTierSizeUsedGib *int64 `json:"hotTierSizeUsedGib,omitempty"`

	// Optional. Block devices for the volume.
	//  Currently, only one block device is permitted per Volume.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.block_devices
	BlockDevices []BlockDeviceObservedState `json:"blockDevices,omitempty"`

	// Output only. If this volume is a clone, this field contains details about
	//  the clone.
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.clone_details
	CloneDetails *Volume_CloneDetailsObservedState `json:"cloneDetails,omitempty"`
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

// NetAppVolume is the Schema for the NetAppVolume API
// +k8s:openapi-gen=true
type NetAppVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetAppVolumeSpec   `json:"spec,omitempty"`
	Status NetAppVolumeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetAppVolumeList contains a list of NetAppVolume
type NetAppVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetAppVolume `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetAppVolume{}, &NetAppVolumeList{})
}

// +kcc:proto=google.cloud.netapp.v1.BackupConfig
type BackupConfig struct {
	// Optional. When specified, schedule backups will be created based on the
	//  policy configuration.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.backup_policies
	BackupPolicyRefs []BackupPolicyRef `json:"backupPolicyRefs,omitempty"`

	// Optional. Name of backup vault.
	//  Format:
	//  projects/{project_id}/locations/{location}/backupVaults/{backup_vault_id}
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.backup_vault
	BackupVaultRef *BackupVaultRef `json:"backupVaultRef,omitempty"`

	// Optional. When set to true, scheduled backup is enabled on the volume.
	//  This field should be nil when there's no backup policy attached.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.scheduled_backup_enabled
	ScheduledBackupEnabled *bool `json:"scheduledBackupEnabled,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.netapp.v1.BackupConfig
type BackupConfigObservedState struct {
	// Output only. Total size of all backups in a chain in bytes = baseline
	//  backup size + sum(incremental backup size).
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.backup_chain_bytes
	BackupChainBytes *int64 `json:"backupChainBytes,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.BlockDevice
type BlockDevice struct {
	// Optional. User-defined name for the block device, unique within the volume.
	//  In case no user input is provided, name will be auto-generated in the
	//  backend. The name must meet the following requirements:
	//  *   Be between 1 and 255 characters long.
	//  *   Contain only uppercase or lowercase letters (A-Z, a-z), numbers (0-9),
	//      and the following special characters: "-", "_", "}", "{", ".".
	//  *   Spaces are not allowed.
	// +kcc:proto:field=google.cloud.netapp.v1.BlockDevice.name
	Name *string `json:"name,omitempty"`

	// Optional. A list of host groups that identify hosts that can mount the
	//  block volume. Format:
	//  `projects/{project_id}/locations/{location}/hostGroups/{host_group_id}`
	//  This field can be updated after the block device is created.
	// +kcc:proto:field=google.cloud.netapp.v1.BlockDevice.host_groups
	HostGroupRefs []NetAppHostGroupRef `json:"hostGroupRefs,omitempty"`

	// Optional. The size of the block device in GiB.
	//  Any value provided for the `size_gib` field during volume creation is
	//  ignored. The block device's size is system-managed and will be set to match
	//  the parent Volume's `capacity_gib`.
	// +kcc:proto:field=google.cloud.netapp.v1.BlockDevice.size_gib
	SizeGib *int64 `json:"sizeGib,omitempty"`

	// Required. Immutable. The OS type of the volume.
	//  This field can't be changed after the block device is created.
	// +kcc:proto:field=google.cloud.netapp.v1.BlockDevice.os_type
	// +required
	OSType *string `json:"osType,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.netapp.v1.BlockDevice
type BlockDeviceObservedState struct {
	// Output only. Device identifier of the block volume. This represents
	//  `lun_serial_number` for iSCSI volumes.
	// +kcc:proto:field=google.cloud.netapp.v1.BlockDevice.identifier
	Identifier *string `json:"identifier,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.HybridReplicationParameters
type HybridReplicationParameters struct {
	// Required. Desired name for the replication of this volume.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.replication
	// +required
	ReplicationRef *NetappReplicationRef `json:"replicationRef,omitempty"`

	// Required. Name of the user's local source volume to be peered with the
	//  destination volume.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_volume_name
	// +required
	PeerVolumeName *string `json:"peerVolumeName,omitempty"`

	// Required. Name of the user's local source cluster to be peered with the
	//  destination cluster.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_cluster_name
	// +required
	PeerClusterName *string `json:"peerClusterName,omitempty"`

	// Required. Name of the user's local source vserver svm to be peered with the
	//  destination vserver svm.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_svm_name
	// +required
	PeerSvmName *string `json:"peerSvmName,omitempty"`

	// Required. List of node ip addresses to be peered with.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_ip_addresses
	// +required
	PeerIPAddresses []string `json:"peerIPAddresses,omitempty"`

	// Optional. Name of source cluster location associated with the Hybrid
	//  replication. This is a free-form field for the display purpose only.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.cluster_location
	ClusterLocation *string `json:"clusterLocation,omitempty"`

	// Optional. Description of the replication.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels to be added to the replication as the key value pairs.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Replication Schedule for the replication created.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.replication_schedule
	ReplicationSchedule *string `json:"replicationSchedule,omitempty"`

	// Optional. Type of the hybrid replication.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.hybrid_replication_type
	HybridReplicationType *string `json:"hybridReplicationType,omitempty"`

	// Optional. Constituent volume count for large volume.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.large_volume_constituent_count
	LargeVolumeConstituentCount *int32 `json:"largeVolumeConstituentCount,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.RestoreParameters
type RestoreParameters struct {
	// Full name of the snapshot resource.
	//  Format:
	//  projects/{project}/locations/{location}/volumes/{volume}/snapshots/{snapshot}
	// +kcc:proto:field=google.cloud.netapp.v1.RestoreParameters.source_snapshot
	SourceSnapshotRef *NetAppSnapshotRef `json:"sourceSnapshotRef,omitempty"`

	// Full name of the backup resource.
	//  Format for standard backup:
	//  projects/{project}/locations/{location}/backupVaults/{backup_vault_id}/backups/{backup_id}
	//  Format for BackupDR backup:
	//  projects/{project}/locations/{location}/backupVaults/{backup_vault}/dataSources/{data_source}/backups/{backup}
	// +kcc:proto:field=google.cloud.netapp.v1.RestoreParameters.source_backup
	SourceBackupRef *NetAppBackupRef `json:"sourceBackupRef,omitempty"`
}
