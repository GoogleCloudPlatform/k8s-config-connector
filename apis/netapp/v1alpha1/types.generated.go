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


// +kcc:proto=google.cloud.netapp.v1.BackupConfig
type BackupConfig struct {
	// Optional. When specified, schedule backups will be created based on the
	//  policy configuration.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.backup_policies
	BackupPolicies []string `json:"backupPolicies,omitempty"`

	// Optional. Name of backup vault.
	//  Format:
	//  projects/{project_id}/locations/{location}/backupVaults/{backup_vault_id}
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.backup_vault
	BackupVault *string `json:"backupVault,omitempty"`

	// Optional. When set to true, scheduled backup is enabled on the volume.
	//  This field should be nil when there's no backup policy attached.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.scheduled_backup_enabled
	ScheduledBackupEnabled *bool `json:"scheduledBackupEnabled,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.DailySchedule
type DailySchedule struct {
	// The maximum number of Snapshots to keep for the hourly schedule
	// +kcc:proto:field=google.cloud.netapp.v1.DailySchedule.snapshots_to_keep
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty"`

	// Set the minute of the hour to start the snapshot (0-59), defaults to the
	//  top of the hour (0).
	// +kcc:proto:field=google.cloud.netapp.v1.DailySchedule.minute
	Minute *float64 `json:"minute,omitempty"`

	// Set the hour to start the snapshot (0-23), defaults to midnight (0).
	// +kcc:proto:field=google.cloud.netapp.v1.DailySchedule.hour
	Hour *float64 `json:"hour,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.ExportPolicy
type ExportPolicy struct {
	// Required. List of export policy rules
	// +kcc:proto:field=google.cloud.netapp.v1.ExportPolicy.rules
	Rules []SimpleExportPolicyRule `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.HourlySchedule
type HourlySchedule struct {
	// The maximum number of Snapshots to keep for the hourly schedule
	// +kcc:proto:field=google.cloud.netapp.v1.HourlySchedule.snapshots_to_keep
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty"`

	// Set the minute of the hour to start the snapshot (0-59), defaults to the
	//  top of the hour (0).
	// +kcc:proto:field=google.cloud.netapp.v1.HourlySchedule.minute
	Minute *float64 `json:"minute,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.HybridReplicationParameters
type HybridReplicationParameters struct {
	// Required. Desired name for the replication of this volume.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.replication
	Replication *string `json:"replication,omitempty"`

	// Required. Name of the user's local source volume to be peered with the
	//  destination volume.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_volume_name
	PeerVolumeName *string `json:"peerVolumeName,omitempty"`

	// Required. Name of the user's local source cluster to be peered with the
	//  destination cluster.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_cluster_name
	PeerClusterName *string `json:"peerClusterName,omitempty"`

	// Required. Name of the user's local source vserver svm to be peered with the
	//  destination vserver svm.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_svm_name
	PeerSvmName *string `json:"peerSvmName,omitempty"`

	// Required. List of node ip addresses to be peered with.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridReplicationParameters.peer_ip_addresses
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
}

// +kcc:proto=google.cloud.netapp.v1.MonthlySchedule
type MonthlySchedule struct {
	// The maximum number of Snapshots to keep for the hourly schedule
	// +kcc:proto:field=google.cloud.netapp.v1.MonthlySchedule.snapshots_to_keep
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty"`

	// Set the minute of the hour to start the snapshot (0-59), defaults to the
	//  top of the hour (0).
	// +kcc:proto:field=google.cloud.netapp.v1.MonthlySchedule.minute
	Minute *float64 `json:"minute,omitempty"`

	// Set the hour to start the snapshot (0-23), defaults to midnight (0).
	// +kcc:proto:field=google.cloud.netapp.v1.MonthlySchedule.hour
	Hour *float64 `json:"hour,omitempty"`

	// Set the day or days of the month to make a snapshot (1-31). Accepts a
	//  comma separated number of days. Defaults to '1'.
	// +kcc:proto:field=google.cloud.netapp.v1.MonthlySchedule.days_of_month
	DaysOfMonth *string `json:"daysOfMonth,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.MountOption
type MountOption struct {
	// Export string
	// +kcc:proto:field=google.cloud.netapp.v1.MountOption.export
	Export *string `json:"export,omitempty"`

	// Full export string
	// +kcc:proto:field=google.cloud.netapp.v1.MountOption.export_full
	ExportFull *string `json:"exportFull,omitempty"`

	// Protocol to mount with.
	// +kcc:proto:field=google.cloud.netapp.v1.MountOption.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Instructions for mounting
	// +kcc:proto:field=google.cloud.netapp.v1.MountOption.instructions
	Instructions *string `json:"instructions,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.RestoreParameters
type RestoreParameters struct {
	// Full name of the snapshot resource.
	//  Format:
	//  projects/{project}/locations/{location}/volumes/{volume}/snapshots/{snapshot}
	// +kcc:proto:field=google.cloud.netapp.v1.RestoreParameters.source_snapshot
	SourceSnapshot *string `json:"sourceSnapshot,omitempty"`

	// Full name of the backup resource.
	//  Format:
	//  projects/{project}/locations/{location}/backupVaults/{backup_vault_id}/backups/{backup_id}
	// +kcc:proto:field=google.cloud.netapp.v1.RestoreParameters.source_backup
	SourceBackup *string `json:"sourceBackup,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.SimpleExportPolicyRule
type SimpleExportPolicyRule struct {
	// Comma separated list of allowed clients IP addresses
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.allowed_clients
	AllowedClients *string `json:"allowedClients,omitempty"`

	// Whether Unix root access will be granted.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.has_root_access
	HasRootAccess *string `json:"hasRootAccess,omitempty"`

	// Access type (ReadWrite, ReadOnly, None)
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.access_type
	AccessType *string `json:"accessType,omitempty"`

	// NFS V3 protocol.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.nfsv3
	Nfsv3 *bool `json:"nfsv3,omitempty"`

	// NFS V4 protocol.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.nfsv4
	Nfsv4 *bool `json:"nfsv4,omitempty"`

	// If enabled (true) the rule defines a read only access for clients matching
	//  the 'allowedClients' specification. It enables nfs clients to mount using
	//  'authentication' kerberos security mode.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.kerberos_5_read_only
	Kerberos5ReadOnly *bool `json:"kerberos5ReadOnly,omitempty"`

	// If enabled (true) the rule defines read and write access for clients
	//  matching the 'allowedClients' specification. It enables nfs clients to
	//  mount using 'authentication' kerberos security mode. The
	//  'kerberos5ReadOnly' value be ignored if this is enabled.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.kerberos_5_read_write
	Kerberos5ReadWrite *bool `json:"kerberos5ReadWrite,omitempty"`

	// If enabled (true) the rule defines a read only access for clients matching
	//  the 'allowedClients' specification. It enables nfs clients to mount using
	//  'integrity' kerberos security mode.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.kerberos_5i_read_only
	Kerberos5iReadOnly *bool `json:"kerberos5iReadOnly,omitempty"`

	// If enabled (true) the rule defines read and write access for clients
	//  matching the 'allowedClients' specification. It enables nfs clients to
	//  mount using 'integrity' kerberos security mode. The 'kerberos5iReadOnly'
	//  value be ignored if this is enabled.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.kerberos_5i_read_write
	Kerberos5iReadWrite *bool `json:"kerberos5iReadWrite,omitempty"`

	// If enabled (true) the rule defines a read only access for clients matching
	//  the 'allowedClients' specification. It enables nfs clients to mount using
	//  'privacy' kerberos security mode.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.kerberos_5p_read_only
	Kerberos5pReadOnly *bool `json:"kerberos5pReadOnly,omitempty"`

	// If enabled (true) the rule defines read and write access for clients
	//  matching the 'allowedClients' specification. It enables nfs clients to
	//  mount using 'privacy' kerberos security mode. The 'kerberos5pReadOnly'
	//  value be ignored if this is enabled.
	// +kcc:proto:field=google.cloud.netapp.v1.SimpleExportPolicyRule.kerberos_5p_read_write
	Kerberos5pReadWrite *bool `json:"kerberos5pReadWrite,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.SnapshotPolicy
type SnapshotPolicy struct {
	// If enabled, make snapshots automatically according to the schedules.
	//  Default is false.
	// +kcc:proto:field=google.cloud.netapp.v1.SnapshotPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Hourly schedule policy.
	// +kcc:proto:field=google.cloud.netapp.v1.SnapshotPolicy.hourly_schedule
	HourlySchedule *HourlySchedule `json:"hourlySchedule,omitempty"`

	// Daily schedule policy.
	// +kcc:proto:field=google.cloud.netapp.v1.SnapshotPolicy.daily_schedule
	DailySchedule *DailySchedule `json:"dailySchedule,omitempty"`

	// Weekly schedule policy.
	// +kcc:proto:field=google.cloud.netapp.v1.SnapshotPolicy.weekly_schedule
	WeeklySchedule *WeeklySchedule `json:"weeklySchedule,omitempty"`

	// Monthly schedule policy.
	// +kcc:proto:field=google.cloud.netapp.v1.SnapshotPolicy.monthly_schedule
	MonthlySchedule *MonthlySchedule `json:"monthlySchedule,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.TieringPolicy
type TieringPolicy struct {
	// Optional. Flag indicating if the volume has tiering policy enable/pause.
	//  Default is PAUSED.
	// +kcc:proto:field=google.cloud.netapp.v1.TieringPolicy.tier_action
	TierAction *string `json:"tierAction,omitempty"`

	// Optional. Time in days to mark the volume's data block as cold and make it
	//  eligible for tiering, can be range from 7-183. Default is 31.
	// +kcc:proto:field=google.cloud.netapp.v1.TieringPolicy.cooling_threshold_days
	CoolingThresholdDays *int32 `json:"coolingThresholdDays,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.Volume
type Volume struct {
	// Identifier. Name of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.name
	Name *string `json:"name,omitempty"`

	// Required. Share name of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.share_name
	ShareName *string `json:"shareName,omitempty"`

	// Required. StoragePool name of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.storage_pool
	StoragePool *string `json:"storagePool,omitempty"`

	// Required. Capacity in GIB of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.capacity_gib
	CapacityGib *int64 `json:"capacityGib,omitempty"`

	// Optional. Export policy of the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.export_policy
	ExportPolicy *ExportPolicy `json:"exportPolicy,omitempty"`

	// Required. Protocols required for the volume
	// +kcc:proto:field=google.cloud.netapp.v1.Volume.protocols
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

// +kcc:proto=google.cloud.netapp.v1.WeeklySchedule
type WeeklySchedule struct {
	// The maximum number of Snapshots to keep for the hourly schedule
	// +kcc:proto:field=google.cloud.netapp.v1.WeeklySchedule.snapshots_to_keep
	SnapshotsToKeep *float64 `json:"snapshotsToKeep,omitempty"`

	// Set the minute of the hour to start the snapshot (0-59), defaults to the
	//  top of the hour (0).
	// +kcc:proto:field=google.cloud.netapp.v1.WeeklySchedule.minute
	Minute *float64 `json:"minute,omitempty"`

	// Set the hour to start the snapshot (0-23), defaults to midnight (0).
	// +kcc:proto:field=google.cloud.netapp.v1.WeeklySchedule.hour
	Hour *float64 `json:"hour,omitempty"`

	// Set the day or days of the week to make a snapshot. Accepts a comma
	//  separated days of the week. Defaults to 'Sunday'.
	// +kcc:proto:field=google.cloud.netapp.v1.WeeklySchedule.day
	Day *string `json:"day,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.BackupConfig
type BackupConfigObservedState struct {
	// Output only. Total size of all backups in a chain in bytes = baseline
	//  backup size + sum(incremental backup size).
	// +kcc:proto:field=google.cloud.netapp.v1.BackupConfig.backup_chain_bytes
	BackupChainBytes *int64 `json:"backupChainBytes,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.Volume
type VolumeObservedState struct {
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
	MountOptions []MountOption `json:"mountOptions,omitempty"`

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
