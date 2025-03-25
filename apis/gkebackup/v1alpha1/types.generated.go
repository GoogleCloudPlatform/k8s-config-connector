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

// +generated:types
// krm.group: gkebackup.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.gkebackup.v1
// resource: GKEBackupBackupPlan:BackupPlan
// resource: GKEBackupRestorePlan:RestorePlan
// resource: GKEBackupBackup:Backup

package v1alpha1

// +kcc:proto=google.cloud.gkebackup.v1.Backup
type Backup struct {

	// Optional. A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Minimum age for this Backup (in days). If this field is set to a
	//  non-zero value, the Backup will be "locked" against deletion (either manual
	//  or automatic deletion) for the number of days provided (measured from the
	//  creation time of the Backup).  MUST be an integer value between 0-90
	//  (inclusive).
	//
	//  Defaults to parent BackupPlan's
	//  [backup_delete_lock_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days]
	//  setting and may only be increased
	//  (either at creation time or in a subsequent update).
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.delete_lock_days
	DeleteLockDays *int32 `json:"deleteLockDays,omitempty"`

	// Optional. The age (in days) after which this Backup will be automatically
	//  deleted. Must be an integer value >= 0:
	//
	//  - If 0, no automatic deletion will occur for this Backup.
	//  - If not 0, this must be >=
	//  [delete_lock_days][google.cloud.gkebackup.v1.Backup.delete_lock_days] and
	//  <= 365.
	//
	//  Once a Backup is created, this value may only be increased.
	//
	//  Defaults to the parent BackupPlan's
	//  [backup_retain_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.retain_days
	RetainDays *int32 `json:"retainDays,omitempty"`

	// Optional. User specified descriptive string for this Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.Backup.ClusterMetadata
type Backup_ClusterMetadata struct {
}

// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan.BackupConfig
type BackupPlan_BackupConfig struct {
	// If True, include all namespaced resources
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.all_namespaces
	AllNamespaces *bool `json:"allNamespaces,omitempty"`

	// If set, include just the resources in the listed namespaces.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.selected_namespaces
	SelectedNamespaces *Namespaces `json:"selectedNamespaces,omitempty"`

	// If set, include just the resources referenced by the listed
	//  ProtectedApplications.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.selected_applications
	SelectedApplications *NamespacedNames `json:"selectedApplications,omitempty"`

	// Optional. This flag specifies whether volume data should be backed up
	//  when PVCs are included in the scope of a Backup.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_volume_data
	IncludeVolumeData *bool `json:"includeVolumeData,omitempty"`

	// Optional. This flag specifies whether Kubernetes Secret resources should
	//  be included when they fall into the scope of Backups.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_secrets
	IncludeSecrets *bool `json:"includeSecrets,omitempty"`

	// Optional. This defines a customer managed encryption key that will be
	//  used to encrypt the "config" portion (the Kubernetes resources) of
	//  Backups created via this plan.
	//
	//  Default (empty): Config backup artifacts will not be encrypted.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.encryption_key
	EncryptionKey *EncryptionKey `json:"encryptionKey,omitempty"`

	// Optional. If false, Backups will fail when Backup for GKE detects
	//  Kubernetes configuration that is non-standard or
	//  requires additional setup to restore.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.permissive_mode
	PermissiveMode *bool `json:"permissiveMode,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy
type BackupPlan_RetentionPolicy struct {
	// Optional. Minimum age for Backups created via this BackupPlan (in days).
	//  This field MUST be an integer value between 0-90 (inclusive).
	//  A Backup created under this BackupPlan will NOT be deletable until it
	//  reaches Backup's (create_time + backup_delete_lock_days).
	//  Updating this field of a BackupPlan does NOT affect existing Backups
	//  under it. Backups created AFTER a successful update will inherit
	//  the new value.
	//
	//  Default: 0 (no delete blocking)
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days
	BackupDeleteLockDays *int32 `json:"backupDeleteLockDays,omitempty"`

	// Optional. The default maximum age of a Backup created via this
	//  BackupPlan. This field MUST be an integer value >= 0 and <= 365. If
	//  specified, a Backup created under this BackupPlan will be automatically
	//  deleted after its age reaches (create_time + backup_retain_days). If not
	//  specified, Backups created under this BackupPlan will NOT be subject to
	//  automatic deletion. Updating this field does NOT affect existing Backups
	//  under it. Backups created AFTER a successful update will automatically
	//  pick up the new value. NOTE: backup_retain_days must be >=
	//  [backup_delete_lock_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days].
	//  If
	//  [cron_schedule][google.cloud.gkebackup.v1.BackupPlan.Schedule.cron_schedule]
	//  is defined, then this must be
	//  <= 360 * the creation interval. If
	//  [rpo_config][google.cloud.gkebackup.v1.BackupPlan.Schedule.rpo_config] is
	//  defined, then this must be
	//  <= 360 * [target_rpo_minutes][Schedule.rpo_config.target_rpo_minutes] /
	//  (1440minutes/day).
	//
	//  Default: 0 (no automatic deletion)
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days
	BackupRetainDays *int32 `json:"backupRetainDays,omitempty"`

	// Optional. This flag denotes whether the retention policy of this
	//  BackupPlan is locked.  If set to True, no further update is allowed on
	//  this policy, including the `locked` field itself.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.locked
	Locked *bool `json:"locked,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.ExclusionWindow.DayOfWeekList
type ExclusionWindow_DayOfWeekList struct {
	// Optional. A list of days of week.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.DayOfWeekList.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.NamespacedName
type NamespacedName struct {
	// Optional. The Namespace of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.namespace
	Namespace *string `json:"namespace,omitempty"`

	// Optional. The name of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.NamespacedNames
type NamespacedNames struct {
	// Optional. A list of namespaced Kubernetes resources.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedNames.namespaced_names
	NamespacedNames []NamespacedName `json:"namespacedNames,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.Namespaces
type Namespaces struct {
	// Optional. A list of Kubernetes Namespaces
	// +kcc:proto:field=google.cloud.gkebackup.v1.Namespaces.namespaces
	Namespaces []string `json:"namespaces,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig
type RestoreConfig struct {
	// Optional. Specifies the mechanism to be used to restore volume data.
	//  Default: VOLUME_DATA_RESTORE_POLICY_UNSPECIFIED (will be treated as
	//  NO_VOLUME_DATA_RESTORATION).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.volume_data_restore_policy
	VolumeDataRestorePolicy *string `json:"volumeDataRestorePolicy,omitempty"`

	// Optional. Defines the behavior for handling the situation where
	//  cluster-scoped resources being restored already exist in the target
	//  cluster. This MUST be set to a value other than
	//  CLUSTER_RESOURCE_CONFLICT_POLICY_UNSPECIFIED if
	//  [cluster_resource_restore_scope][google.cloud.gkebackup.v1.RestoreConfig.cluster_resource_restore_scope]
	//  is not empty.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.cluster_resource_conflict_policy
	ClusterResourceConflictPolicy *string `json:"clusterResourceConflictPolicy,omitempty"`

	// Optional. Defines the behavior for handling the situation where sets of
	//  namespaced resources being restored already exist in the target cluster.
	//  This MUST be set to a value other than
	//  NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.namespaced_resource_restore_mode
	NamespacedResourceRestoreMode *string `json:"namespacedResourceRestoreMode,omitempty"`

	// Optional. Identifies the cluster-scoped resources to restore from the
	//  Backup. Not specifying it means NO cluster resource will be restored.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.cluster_resource_restore_scope
	ClusterResourceRestoreScope *RestoreConfig_ClusterResourceRestoreScope `json:"clusterResourceRestoreScope,omitempty"`

	// Restore all namespaced resources in the Backup if set to "True".
	//  Specifying this field to "False" is an error.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.all_namespaces
	AllNamespaces *bool `json:"allNamespaces,omitempty"`

	// A list of selected Namespaces to restore from the Backup. The listed
	//  Namespaces and all resources contained in them will be restored.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.selected_namespaces
	SelectedNamespaces *Namespaces `json:"selectedNamespaces,omitempty"`

	// A list of selected ProtectedApplications to restore. The listed
	//  ProtectedApplications and all the resources to which they refer will be
	//  restored.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.selected_applications
	SelectedApplications *NamespacedNames `json:"selectedApplications,omitempty"`

	// Do not restore any namespaced resources if set to "True".
	//  Specifying this field to "False" is not allowed.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.no_namespaces
	NoNamespaces *bool `json:"noNamespaces,omitempty"`

	// A list of selected namespaces excluded from restoration. All
	//  namespaces except those in this list will be restored.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.excluded_namespaces
	ExcludedNamespaces *Namespaces `json:"excludedNamespaces,omitempty"`

	// Optional. A list of transformation rules to be applied against Kubernetes
	//  resources as they are selected for restoration from a Backup. Rules are
	//  executed in order defined - this order matters, as changes made by a rule
	//  may impact the filtering logic of subsequent rules. An empty list means no
	//  substitution will occur.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.substitution_rules
	SubstitutionRules []RestoreConfig_SubstitutionRule `json:"substitutionRules,omitempty"`

	// Optional. A list of transformation rules to be applied against Kubernetes
	//  resources as they are selected for restoration from a Backup. Rules are
	//  executed in order defined - this order matters, as changes made by a rule
	//  may impact the filtering logic of subsequent rules. An empty list means no
	//  transformation will occur.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.transformation_rules
	TransformationRules []RestoreConfig_TransformationRule `json:"transformationRules,omitempty"`

	// Optional. A table that binds volumes by their scope to a restore policy.
	//  Bindings must have a unique scope. Any volumes not scoped in the bindings
	//  are subject to the policy defined in volume_data_restore_policy.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.volume_data_restore_policy_bindings
	VolumeDataRestorePolicyBindings []RestoreConfig_VolumeDataRestorePolicyBinding `json:"volumeDataRestorePolicyBindings,omitempty"`

	// Optional. RestoreOrder contains custom ordering to use on a Restore.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.restore_order
	RestoreOrder *RestoreConfig_RestoreOrder `json:"restoreOrder,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.ClusterResourceRestoreScope
type RestoreConfig_ClusterResourceRestoreScope struct {
	// Optional. A list of cluster-scoped resource group kinds to restore from
	//  the backup. If specified, only the selected resources will be restored.
	//  Mutually exclusive to any other field in the message.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ClusterResourceRestoreScope.selected_group_kinds
	SelectedGroupKinds []RestoreConfig_GroupKind `json:"selectedGroupKinds,omitempty"`

	// Optional. A list of cluster-scoped resource group kinds to NOT restore
	//  from the backup. If specified, all valid cluster-scoped resources will be
	//  restored except for those specified in the list.
	//  Mutually exclusive to any other field in the message.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ClusterResourceRestoreScope.excluded_group_kinds
	ExcludedGroupKinds []RestoreConfig_GroupKind `json:"excludedGroupKinds,omitempty"`

	// Optional. If True, all valid cluster-scoped resources will be restored.
	//  Mutually exclusive to any other field in the message.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ClusterResourceRestoreScope.all_group_kinds
	AllGroupKinds *bool `json:"allGroupKinds,omitempty"`

	// Optional. If True, no cluster-scoped resources will be restored.
	//  This has the same restore scope as if the message is not defined.
	//  Mutually exclusive to any other field in the message.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ClusterResourceRestoreScope.no_group_kinds
	NoGroupKinds *bool `json:"noGroupKinds,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.GroupKind
type RestoreConfig_GroupKind struct {
	// Optional. API group string of a Kubernetes resource, e.g.
	//  "apiextensions.k8s.io", "storage.k8s.io", etc.
	//  Note: use empty string for core API group
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.GroupKind.resource_group
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// Optional. Kind of a Kubernetes resource, must be in UpperCamelCase
	//  (PascalCase) and singular form. E.g. "CustomResourceDefinition",
	//  "StorageClass", etc.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.GroupKind.resource_kind
	ResourceKind *string `json:"resourceKind,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder
type RestoreConfig_RestoreOrder struct {
	// Optional. Contains a list of group kind dependency pairs provided
	//  by the customer, that is used by Backup for GKE to
	//  generate a group kind restore order.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.group_kind_dependencies
	GroupKindDependencies []RestoreConfig_RestoreOrder_GroupKindDependency `json:"groupKindDependencies,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.type.TimeOfDay
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	//  to allow the value "24:00:00" for scenarios like business closing time.
	// +kcc:proto:field=google.type.TimeOfDay.hours
	Hours *int32 `json:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.TimeOfDay.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	//  allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.TimeOfDay.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kcc:proto:field=google.type.TimeOfDay.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.Backup
type BackupObservedState struct {
	// Output only. The fully qualified name of the Backup.
	//  `projects/*/locations/*/backupPlans/*/backups/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.name
	Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID4](https://en.wikipedia.org/wiki/Universally_unique_identifier)
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The timestamp when this Backup resource was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this Backup resource was last updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. This flag indicates whether this Backup resource was created
	//  manually by a user or via a schedule in the BackupPlan. A value of True
	//  means that the Backup was created manually.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.manual
	Manual *bool `json:"manual,omitempty"`

	// Output only. The time at which an existing delete lock will expire for this
	//  backup (calculated from create_time +
	//  [delete_lock_days][google.cloud.gkebackup.v1.Backup.delete_lock_days]).
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.delete_lock_expire_time
	DeleteLockExpireTime *string `json:"deleteLockExpireTime,omitempty"`

	// Output only. The time at which this Backup will be automatically deleted
	//  (calculated from create_time +
	//  [retain_days][google.cloud.gkebackup.v1.Backup.retain_days]).
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.retain_expire_time
	RetainExpireTime *string `json:"retainExpireTime,omitempty"`

	// Output only. The customer managed encryption key that was used to encrypt
	//  the Backup's artifacts.  Inherited from the parent BackupPlan's
	//  [encryption_key][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.encryption_key]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.encryption_key
	EncryptionKey *EncryptionKey `json:"encryptionKey,omitempty"`

	// Output only. If True, all namespaces were included in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.all_namespaces
	AllNamespaces *bool `json:"allNamespaces,omitempty"`

	// Output only. If set, the list of namespaces that were included in the
	//  Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.selected_namespaces
	SelectedNamespaces *Namespaces `json:"selectedNamespaces,omitempty"`

	// Output only. If set, the list of ProtectedApplications whose resources
	//  were included in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.selected_applications
	SelectedApplications *NamespacedNames `json:"selectedApplications,omitempty"`

	// Output only. Whether or not the Backup contains volume data.  Controlled by
	//  the parent BackupPlan's
	//  [include_volume_data][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_volume_data]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.contains_volume_data
	ContainsVolumeData *bool `json:"containsVolumeData,omitempty"`

	// Output only. Whether or not the Backup contains Kubernetes Secrets.
	//  Controlled by the parent BackupPlan's
	//  [include_secrets][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_secrets]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.contains_secrets
	ContainsSecrets *bool `json:"containsSecrets,omitempty"`

	// Output only. Information about the GKE cluster from which this Backup was
	//  created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.cluster_metadata
	ClusterMetadata *Backup_ClusterMetadata `json:"clusterMetadata,omitempty"`

	// Output only. Current state of the Backup
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. Human-readable description of why the backup is in the current
	//  `state`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.state_reason
	StateReason *string `json:"stateReason,omitempty"`

	// Output only. Completion time of the Backup
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. The total number of Kubernetes resources included in the
	//  Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.resource_count
	ResourceCount *int32 `json:"resourceCount,omitempty"`

	// Output only. The total number of volume backups contained in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.volume_count
	VolumeCount *int32 `json:"volumeCount,omitempty"`

	// Output only. The total size of the Backup in bytes = config backup size +
	//  sum(volume backup sizes)
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a backup from overwriting each other.
	//  It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform backup updates in order to avoid
	//  race conditions: An `etag` is returned in the response to `GetBackup`,
	//  and systems are expected to put that etag in the request to
	//  `UpdateBackup` or `DeleteBackup` to ensure that their change will be
	//  applied to the same version of the resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The total number of Kubernetes Pods contained in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.pod_count
	PodCount *int32 `json:"podCount,omitempty"`

	// Output only. The size of the config backup in bytes.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.config_backup_size_bytes
	ConfigBackupSizeBytes *int64 `json:"configBackupSizeBytes,omitempty"`

	// Output only. If false, Backup will fail when Backup for GKE detects
	//  Kubernetes configuration that is non-standard or
	//  requires additional setup to restore.
	//
	//  Inherited from the parent BackupPlan's
	//  [permissive_mode][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.permissive_mode]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.permissive_mode
	PermissiveMode *bool `json:"permissiveMode,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.Backup.ClusterMetadata
type Backup_ClusterMetadataObservedState struct {
	// Output only. The source cluster from which this Backup was created.
	//  Valid formats:
	//
	//    - `projects/*/locations/*/clusters/*`
	//    - `projects/*/zones/*/clusters/*`
	//
	//  This is inherited from the parent BackupPlan's
	//  [cluster][google.cloud.gkebackup.v1.BackupPlan.cluster] field.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Output only. The Kubernetes server version of the source cluster.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.k8s_version
	K8sVersion *string `json:"k8sVersion,omitempty"`

	// Output only. A list of the Backup for GKE CRD versions found in the
	//  cluster.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.backup_crd_versions
	BackupCrdVersions map[string]string `json:"backupCrdVersions,omitempty"`

	// Output only. GKE version
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.gke_version
	GkeVersion *string `json:"gkeVersion,omitempty"`

	// Output only. Anthos version
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.anthos_version
	AnthosVersion *string `json:"anthosVersion,omitempty"`
}
