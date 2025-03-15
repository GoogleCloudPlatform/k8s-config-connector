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

// +kcc:proto=google.cloud.gkebackup.v1.ExclusionWindow
type ExclusionWindow struct {
	// Required. Specifies the start time of the window using time of the day in
	//  UTC.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Required. Specifies duration of the window.
	//  Duration must be >= 5 minutes and < (target RPO - 20 minutes).
	//  Additional restrictions based on the recurrence type to allow some time for
	//  backup to happen:
	//  - single_occurrence_date:  no restriction, but UI may warn about this when
	//  duration >= target RPO
	//  - daily window: duration < 24 hours
	//  - weekly window:
	//    - days of week includes all seven days of a week: duration < 24 hours
	//    - all other weekly window: duration < 168 hours (i.e., 24 * 7 hours)
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.duration
	Duration *string `json:"duration,omitempty"`

	// No recurrence. The exclusion window occurs only once and on this
	//  date in UTC.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.single_occurrence_date
	SingleOccurrenceDate *Date `json:"singleOccurrenceDate,omitempty"`

	// The exclusion window occurs every day if set to "True".
	//  Specifying this field to "False" is an error.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.daily
	Daily *bool `json:"daily,omitempty"`

	// The exclusion window occurs on these days of each week in UTC.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.days_of_week
	DaysOfWeek *ExclusionWindow_DayOfWeekList `json:"daysOfWeek,omitempty"`
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

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter
type RestoreConfig_ResourceFilter struct {
	// Optional. (Filtering parameter) Any resource subject to transformation
	//  must be contained within one of the listed Kubernetes Namespace in the
	//  Backup. If this field is not provided, no namespace filtering will be
	//  performed (all resources in all Namespaces, including all cluster-scoped
	//  resources, will be candidates for transformation).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter.namespaces
	Namespaces []string `json:"namespaces,omitempty"`

	// Optional. (Filtering parameter) Any resource subject to transformation
	//  must belong to one of the listed "types". If this field is not provided,
	//  no type filtering will be performed (all resources of all types matching
	//  previous filtering parameters will be candidates for transformation).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter.group_kinds
	GroupKinds []RestoreConfig_GroupKind `json:"groupKinds,omitempty"`

	// Optional. This is a [JSONPath]
	//  (https://github.com/json-path/JsonPath/blob/master/README.md)
	//  expression that matches specific fields of candidate
	//  resources and it operates as a filtering parameter (resources that
	//  are not matched with this expression will not be candidates for
	//  transformation).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter.json_path
	JsonPath *string `json:"jsonPath,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder
type RestoreConfig_RestoreOrder struct {
	// Optional. Contains a list of group kind dependency pairs provided
	//  by the customer, that is used by Backup for GKE to
	//  generate a group kind restore order.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.group_kind_dependencies
	GroupKindDependencies []RestoreConfig_RestoreOrder_GroupKindDependency `json:"groupKindDependencies,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.GroupKindDependency
type RestoreConfig_RestoreOrder_GroupKindDependency struct {
	// Required. The satisfying group kind must be restored first
	//  in order to satisfy the dependency.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.GroupKindDependency.satisfying
	Satisfying *RestoreConfig_GroupKind `json:"satisfying,omitempty"`

	// Required. The requiring group kind requires that the other
	//  group kind be restored first.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.GroupKindDependency.requiring
	Requiring *RestoreConfig_GroupKind `json:"requiring,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule
type RestoreConfig_SubstitutionRule struct {
	// Optional. (Filtering parameter) Any resource subject to substitution must
	//  be contained within one of the listed Kubernetes Namespace in the Backup.
	//  If this field is not provided, no namespace filtering will be performed
	//  (all resources in all Namespaces, including all cluster-scoped resources,
	//  will be candidates for substitution).
	//  To mix cluster-scoped and namespaced resources in the same rule, use an
	//  empty string ("") as one of the target namespaces.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.target_namespaces
	TargetNamespaces []string `json:"targetNamespaces,omitempty"`

	// Optional. (Filtering parameter) Any resource subject to substitution must
	//  belong to one of the listed "types". If this field is not provided, no
	//  type filtering will be performed (all resources of all types matching
	//  previous filtering parameters will be candidates for substitution).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.target_group_kinds
	TargetGroupKinds []RestoreConfig_GroupKind `json:"targetGroupKinds,omitempty"`

	// Required. This is a [JSONPath]
	//  (https://kubernetes.io/docs/reference/kubectl/jsonpath/)
	//  expression that matches specific fields of candidate
	//  resources and it operates as both a filtering parameter (resources that
	//  are not matched with this expression will not be candidates for
	//  substitution) as well as a field identifier (identifies exactly which
	//  fields out of the candidate resources will be modified).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.target_json_path
	TargetJsonPath *string `json:"targetJsonPath,omitempty"`

	// Optional. (Filtering parameter) This is a [regular expression]
	//  (https://en.wikipedia.org/wiki/Regular_expression)
	//  that is compared against the fields matched by the target_json_path
	//  expression (and must also have passed the previous filters).
	//  Substitution will not be performed against fields whose
	//  value does not match this expression. If this field is NOT specified,
	//  then ALL fields matched by the target_json_path expression will undergo
	//  substitution. Note that an empty (e.g., "", rather than unspecified)
	//  value for this field will only match empty fields.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.original_value_pattern
	OriginalValuePattern *string `json:"originalValuePattern,omitempty"`

	// Optional. This is the new value to set for any fields that pass the
	//  filtering and selection criteria. To remove a value from a Kubernetes
	//  resource, either leave this field unspecified, or set it to the empty
	//  string ("").
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.new_value
	NewValue *string `json:"newValue,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule
type RestoreConfig_TransformationRule struct {
	// Required. A list of transformation rule actions to take against candidate
	//  resources. Actions are executed in order defined - this order matters, as
	//  they could potentially interfere with each other and the first operation
	//  could affect the outcome of the second operation.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule.field_actions
	FieldActions []RestoreConfig_TransformationRuleAction `json:"fieldActions,omitempty"`

	// Optional. This field is used to specify a set of fields that should be
	//  used to determine which resources in backup should be acted upon by the
	//  supplied transformation rule actions, and this will ensure that only
	//  specific resources are affected by transformation rule actions.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule.resource_filter
	ResourceFilter *RestoreConfig_ResourceFilter `json:"resourceFilter,omitempty"`

	// Optional. The description is a user specified string description of the
	//  transformation rule.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction
type RestoreConfig_TransformationRuleAction struct {
	// Required. op specifies the operation to perform.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.op
	Op *string `json:"op,omitempty"`

	// Optional. A string containing a JSON Pointer value that references the
	//  location in the target document to move the value from.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.from_path
	FromPath *string `json:"fromPath,omitempty"`

	// Optional. A string containing a JSON-Pointer value that references a
	//  location within the target document where the operation is performed.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.path
	Path *string `json:"path,omitempty"`

	// Optional. A string that specifies the desired value in string format to
	//  use for transformation.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.VolumeDataRestorePolicyBinding
type RestoreConfig_VolumeDataRestorePolicyBinding struct {
	// Required. The VolumeDataRestorePolicy to apply when restoring volumes in
	//  scope.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.VolumeDataRestorePolicyBinding.policy
	Policy *string `json:"policy,omitempty"`

	// The volume type, as determined by the PVC's bound PV,
	//  to apply the policy to.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.VolumeDataRestorePolicyBinding.volume_type
	VolumeType *string `json:"volumeType,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestorePlan
type RestorePlan struct {

	// Optional. User specified descriptive string for this RestorePlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. A reference to the
	//  [BackupPlan][google.cloud.gkebackup.v1.BackupPlan] from which Backups may
	//  be used as the source for Restores created via this RestorePlan. Format:
	//  `projects/*/locations/*/backupPlans/*`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.backup_plan
	BackupPlan *string `json:"backupPlan,omitempty"`

	// Required. Immutable. The target cluster into which Restores created via
	//  this RestorePlan will restore data. NOTE: the cluster's region must be the
	//  same as the RestorePlan. Valid formats:
	//
	//    - `projects/*/locations/*/clusters/*`
	//    - `projects/*/zones/*/clusters/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Required. Configuration of Restores created via this RestorePlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.restore_config
	RestoreConfig *RestoreConfig `json:"restoreConfig,omitempty"`

	// Optional. A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.labels
	Labels map[string]string `json:"labels,omitempty"`
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

// +kcc:proto=google.cloud.gkebackup.v1.RestorePlan
type RestorePlanObservedState struct {
	// Output only. The full name of the RestorePlan resource.
	//  Format: `projects/*/locations/*/restorePlans/*`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.name
	Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The timestamp when this RestorePlan resource was
	//  created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this RestorePlan resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a restore from overwriting each other.
	//  It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform restore updates in order to avoid
	//  race conditions: An `etag` is returned in the response to `GetRestorePlan`,
	//  and systems are expected to put that etag in the request to
	//  `UpdateRestorePlan` or `DeleteRestorePlan` to ensure that their change
	//  will be applied to the same version of the resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. State of the RestorePlan. This State field reflects the
	//  various stages a RestorePlan can be in
	//  during the Create operation.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.state
	State *string `json:"state,omitempty"`

	// Output only. Human-readable description of why RestorePlan is in the
	//  current `state`
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.state_reason
	StateReason *string `json:"stateReason,omitempty"`
}
