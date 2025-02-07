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


// +kcc:proto=google.cloud.alloydb.v1beta.AutomatedBackupPolicy
type AutomatedBackupPolicy struct {
	// Weekly schedule for the Backup.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.weekly_schedule
	WeeklySchedule *AutomatedBackupPolicy_WeeklySchedule `json:"weeklySchedule,omitempty"`

	// Time-based Backup retention policy.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.time_based_retention
	TimeBasedRetention *AutomatedBackupPolicy_TimeBasedRetention `json:"timeBasedRetention,omitempty"`

	// Quantity-based Backup retention policy to retain recent backups.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.quantity_based_retention
	QuantityBasedRetention *AutomatedBackupPolicy_QuantityBasedRetention `json:"quantityBasedRetention,omitempty"`

	// Whether automated automated backups are enabled. If not set, defaults to
	//  true.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The length of the time window during which a backup can be
	//  taken. If a backup does not succeed within this time window, it will be
	//  canceled and considered failed.
	//
	//  The backup window must be at least 5 minutes long. There is no upper bound
	//  on the window. If not set, it defaults to 1 hour.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.backup_window
	BackupWindow *string `json:"backupWindow,omitempty"`

	// Optional. The encryption config can be specified to encrypt the
	//  backups with a customer-managed encryption key (CMEK). When this field is
	//  not specified, the backup will then use default encryption scheme to
	//  protect the user data.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// The location where the backup will be stored. Currently, the only supported
	//  option is to store the backup in the same region as the cluster.
	//
	//  If empty, defaults to the region of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.location
	Location *string `json:"location,omitempty"`

	// Labels to apply to backups created using this configuration.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.QuantityBasedRetention
type AutomatedBackupPolicy_QuantityBasedRetention struct {
	// The number of backups to retain.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.QuantityBasedRetention.count
	Count *int32 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.TimeBasedRetention
type AutomatedBackupPolicy_TimeBasedRetention struct {
	// The retention period.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.TimeBasedRetention.retention_period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.WeeklySchedule
type AutomatedBackupPolicy_WeeklySchedule struct {
	// The times during the day to start a backup. The start times are assumed
	//  to be in UTC and to be an exact hour (e.g., 04:00:00).
	//
	//  If no start times are provided, a single fixed start time is chosen
	//  arbitrarily.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.WeeklySchedule.start_times
	StartTimes []TimeOfDay `json:"startTimes,omitempty"`

	// The days of the week to perform a backup.
	//
	//  If this field is left empty, the default of every day of the week is
	//  used.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.WeeklySchedule.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.BackupSource
type BackupSource struct {

	// Required. The name of the backup resource with the format:
	//   * projects/{project}/locations/{region}/backups/{backup_id}
	// +kcc:proto:field=google.cloud.alloydb.v1beta.BackupSource.backup_name
	BackupName *string `json:"backupName,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.CloudSQLBackupRunSource
type CloudSQLBackupRunSource struct {
	// The project ID of the source CloudSQL instance. This should be the same as
	//  the AlloyDB cluster's project.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.CloudSQLBackupRunSource.project
	Project *string `json:"project,omitempty"`

	// Required. The CloudSQL instance ID.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.CloudSQLBackupRunSource.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// Required. The CloudSQL backup run ID.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.CloudSQLBackupRunSource.backup_run_id
	BackupRunID *int64 `json:"backupRunID,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster
type Cluster struct {

	// User-settable and human-readable display name for the Cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The database engine major version. This is an optional field and
	//  it is populated at the Cluster creation time. If a database version is not
	//  supplied at cluster creation time, then a default database version will
	//  be used.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.network_config
	NetworkConfig *Cluster_NetworkConfig `json:"networkConfig,omitempty"`

	// Required. The resource link for the VPC network in which cluster resources
	//  are created and from which they are accessible via Private IP. The network
	//  must belong to the same project as the cluster. It is specified in the
	//  form: `projects/{project}/global/networks/{network_id}`. This is required
	//  to create a cluster. Deprecated, use network_config.network instead.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.network
	Network *string `json:"network,omitempty"`

	// For Resource freshness validation (https://google.aip.dev/154)
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.etag
	Etag *string `json:"etag,omitempty"`

	// Annotations to allow client tools to store small amount of arbitrary data.
	//  This is distinct from labels.
	//  https://google.aip.dev/128
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Input only. Initial user to setup during cluster creation. Required.
	//  If used in `RestoreCluster` this is ignored.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.initial_user
	InitialUser *UserPassword `json:"initialUser,omitempty"`

	// The automated backup policy for this cluster.
	//
	//  If no policy is provided then the default policy will be used. If backups
	//  are supported for the cluster, the default policy takes one backup a day,
	//  has a backup window of 1 hour, and retains backups for 14 days.
	//  For more information on the defaults, consult the
	//  documentation for the message type.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.automated_backup_policy
	AutomatedBackupPolicy *AutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`

	// SSL configuration for this AlloyDB cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.ssl_config
	SslConfig *SslConfig `json:"sslConfig,omitempty"`

	// Optional. The encryption config can be specified to encrypt the data disks
	//  and other persistent data resources of a cluster with a
	//  customer-managed encryption key (CMEK). When this field is not
	//  specified, the cluster will then use default encryption scheme to
	//  protect the user data.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Optional. Continuous backup configuration for this cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.continuous_backup_config
	ContinuousBackupConfig *ContinuousBackupConfig `json:"continuousBackupConfig,omitempty"`

	// Cross Region replication config specific to SECONDARY cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.secondary_config
	SecondaryConfig *Cluster_SecondaryConfig `json:"secondaryConfig,omitempty"`

	// Optional. The configuration for Private Service Connect (PSC) for the
	//  cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.psc_config
	PscConfig *Cluster_PscConfig `json:"pscConfig,omitempty"`

	// Optional. The maintenance update policy determines when to allow or deny
	//  updates.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.maintenance_update_policy
	MaintenanceUpdatePolicy *MaintenanceUpdatePolicy `json:"maintenanceUpdatePolicy,omitempty"`

	// Optional. Configuration parameters related to the Gemini in Databases
	//  add-on.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.gemini_config
	GeminiConfig *GeminiClusterConfig `json:"geminiConfig,omitempty"`

	// Optional. Subscription type of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.subscription_type
	SubscriptionType *string `json:"subscriptionType,omitempty"`

	// Optional. Input only. Immutable. Tag keys/values directly bound to this
	//  resource. For example:
	//  ```
	//  "123/environment": "production",
	//  "123/costCenter": "marketing"
	//  ```
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.tags
	Tags map[string]string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.NetworkConfig
type Cluster_NetworkConfig struct {
	// Optional. The resource link for the VPC network in which cluster
	//  resources are created and from which they are accessible via Private IP.
	//  The network must belong to the same project as the cluster. It is
	//  specified in the form:
	//  `projects/{project_number}/global/networks/{network_id}`. This is
	//  required to create a cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.NetworkConfig.network
	Network *string `json:"network,omitempty"`

	// Optional. Name of the allocated IP range for the private IP AlloyDB
	//  cluster, for example: "google-managed-services-default". If set, the
	//  instance IPs for this cluster will be created in the allocated range. The
	//  range name must comply with RFC 1035. Specifically, the name must be 1-63
	//  characters long and match the regular expression
	//  `[a-z]([-a-z0-9]*[a-z0-9])?`.
	//  Field name is intended to be consistent with Cloud SQL.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.NetworkConfig.allocated_ip_range
	AllocatedIPRange *string `json:"allocatedIPRange,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.PrimaryConfig
type Cluster_PrimaryConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.PscConfig
type Cluster_PscConfig struct {
	// Optional. Create an instance that allows connections from Private Service
	//  Connect endpoints to the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.PscConfig.psc_enabled
	PscEnabled *bool `json:"pscEnabled,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.SecondaryConfig
type Cluster_SecondaryConfig struct {
	// The name of the primary cluster name with the format:
	//  * projects/{project}/locations/{region}/clusters/{cluster_id}
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.SecondaryConfig.primary_cluster_name
	PrimaryClusterName *string `json:"primaryClusterName,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.TrialMetadata
type Cluster_TrialMetadata struct {
	// start time of the trial cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.TrialMetadata.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End time of the trial cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.TrialMetadata.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Upgrade time of trial cluster to Standard cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.TrialMetadata.upgrade_time
	UpgradeTime *string `json:"upgradeTime,omitempty"`

	// grace end time of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.TrialMetadata.grace_end_time
	GraceEndTime *string `json:"graceEndTime,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.ContinuousBackupConfig
type ContinuousBackupConfig struct {
	// Whether ContinuousBackup is enabled.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The number of days that are eligible to restore from using PITR. To support
	//  the entire recovery window, backups and logs are retained for one day more
	//  than the recovery window. If not set, defaults to 14 days.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupConfig.recovery_window_days
	RecoveryWindowDays *int32 `json:"recoveryWindowDays,omitempty"`

	// The encryption config can be specified to encrypt the
	//  backups with a customer-managed encryption key (CMEK). When this field is
	//  not specified, the backup will then use default encryption scheme to
	//  protect the user data.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupConfig.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.ContinuousBackupInfo
type ContinuousBackupInfo struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.EncryptionConfig
type EncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//  Each Cloud KMS key is regionalized and has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.alloydb.v1beta.EncryptionConfig.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.EncryptionInfo
type EncryptionInfo struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.GeminiClusterConfig
type GeminiClusterConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.MaintenanceSchedule
type MaintenanceSchedule struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy
type MaintenanceUpdatePolicy struct {
	// Preferred windows to perform maintenance. Currently limited to 1.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.maintenance_windows
	MaintenanceWindows []MaintenanceUpdatePolicy_MaintenanceWindow `json:"maintenanceWindows,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.MaintenanceWindow
type MaintenanceUpdatePolicy_MaintenanceWindow struct {
	// Preferred day of the week for maintenance, e.g. MONDAY, TUESDAY, etc.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.MaintenanceWindow.day
	Day *string `json:"day,omitempty"`

	// Preferred time to start the maintenance operation on the specified day.
	//  Maintenance will start within 1 hour of this time.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.MaintenanceWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.MigrationSource
type MigrationSource struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.SslConfig
type SslConfig struct {
	// Optional. SSL mode. Specifies client-server SSL/TLS connection behavior.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.SslConfig.ssl_mode
	SslMode *string `json:"sslMode,omitempty"`

	// Optional. Certificate Authority (CA) source. Only CA_SOURCE_MANAGED is
	//  supported currently, and is the default value.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.SslConfig.ca_source
	CaSource *string `json:"caSource,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.UserPassword
type UserPassword struct {
	// The database username.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.UserPassword.user
	User *string `json:"user,omitempty"`

	// The initial password for the user.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.UserPassword.password
	Password *string `json:"password,omitempty"`
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

// +kcc:proto=google.cloud.alloydb.v1beta.BackupSource
type BackupSourceObservedState struct {
	// Output only. The system-generated UID of the backup which was used to
	//  create this resource. The UID is generated when the backup is created, and
	//  it is retained until the backup is deleted.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.BackupSource.backup_uid
	BackupUid *string `json:"backupUid,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster
type ClusterObservedState struct {
	// Output only. Cluster created from backup.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.backup_source
	BackupSource *BackupSource `json:"backupSource,omitempty"`

	// Output only. Cluster created via DMS migration.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.migration_source
	MigrationSource *MigrationSource `json:"migrationSource,omitempty"`

	// Output only. Cluster created from CloudSQL snapshot.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.cloudsql_backup_run_source
	CloudsqlBackupRunSource *CloudSQLBackupRunSource `json:"cloudsqlBackupRunSource,omitempty"`

	// Output only. The name of the cluster resource with the format:
	//   * projects/{project}/locations/{region}/clusters/{cluster_id}
	//  where the cluster ID segment should satisfy the regex expression
	//  `[a-z0-9-]+`. For more details see https://google.aip.dev/122.
	//  The prefix of the cluster resource name is the name of the parent resource:
	//   * projects/{project}/locations/{region}
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.name
	Name *string `json:"name,omitempty"`

	// Output only. The system-generated UID of the resource. The UID is assigned
	//  when the resource is created, and it is retained until it is deleted.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Create time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Delete time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The current serving state of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.state
	State *string `json:"state,omitempty"`

	// Output only. The type of the cluster. This is an output-only field and it's
	//  populated at the Cluster creation time or the Cluster promotion
	//  time. The cluster type is determined by which RPC was used to create
	//  the cluster (i.e. `CreateCluster` vs. `CreateSecondaryCluster`
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.cluster_type
	ClusterType *string `json:"clusterType,omitempty"`

	// Output only. Reconciling (https://google.aip.dev/128#reconciliation).
	//  Set to true if the current state of Cluster does not match the user's
	//  intended state, and the service is actively updating the resource to
	//  reconcile them. This can happen due to user-triggered updates or
	//  system actions like failover or maintenance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The encryption information for the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.encryption_info
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`

	// Output only. Continuous backup properties for this cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.continuous_backup_info
	ContinuousBackupInfo *ContinuousBackupInfo `json:"continuousBackupInfo,omitempty"`

	// Output only. Cross Region replication config specific to PRIMARY cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.primary_config
	PrimaryConfig *Cluster_PrimaryConfig `json:"primaryConfig,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. The maintenance schedule for the cluster, generated for a
	//  specific rollout if a maintenance window is set.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.maintenance_schedule
	MaintenanceSchedule *MaintenanceSchedule `json:"maintenanceSchedule,omitempty"`

	// Optional. Configuration parameters related to the Gemini in Databases
	//  add-on.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.gemini_config
	GeminiConfig *GeminiClusterConfigObservedState `json:"geminiConfig,omitempty"`

	// Output only. Metadata for free trial clusters
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.trial_metadata
	TrialMetadata *Cluster_TrialMetadata `json:"trialMetadata,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.PrimaryConfig
type Cluster_PrimaryConfigObservedState struct {
	// Output only. Names of the clusters that are replicating from this
	//  cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.PrimaryConfig.secondary_cluster_names
	SecondaryClusterNames []string `json:"secondaryClusterNames,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.ContinuousBackupInfo
type ContinuousBackupInfoObservedState struct {
	// Output only. The encryption information for the WALs and backups required
	//  for ContinuousBackup.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.encryption_info
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`

	// Output only. When ContinuousBackup was most recently enabled. Set to null
	//  if ContinuousBackup is not enabled.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.enabled_time
	EnabledTime *string `json:"enabledTime,omitempty"`

	// Output only. Days of the week on which a continuous backup is taken. Output
	//  only field. Ignored if passed into the request.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.schedule
	Schedule []string `json:"schedule,omitempty"`

	// Output only. The earliest restorable time that can be restored to. Output
	//  only field.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.earliest_restorable_time
	EarliestRestorableTime *string `json:"earliestRestorableTime,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.EncryptionInfo
type EncryptionInfoObservedState struct {
	// Output only. Type of encryption.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.EncryptionInfo.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Output only. Cloud KMS key versions that are being used to protect the
	//  database or the backup.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.EncryptionInfo.kms_key_versions
	KMSKeyVersions []string `json:"kmsKeyVersions,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.GeminiClusterConfig
type GeminiClusterConfigObservedState struct {
	// Output only. Whether the Gemini in Databases add-on is enabled for the
	//  cluster. It will be true only if the add-on has been enabled for the
	//  billing account corresponding to the cluster. Its status is toggled from
	//  the Admin Control Center (ACC) and cannot be toggled using AlloyDB's APIs.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.GeminiClusterConfig.entitled
	Entitled *bool `json:"entitled,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.MaintenanceSchedule
type MaintenanceScheduleObservedState struct {
	// Output only. The scheduled start time for the maintenance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceSchedule.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.MigrationSource
type MigrationSourceObservedState struct {
	// Output only. The host and port of the on-premises instance in host:port
	//  format
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MigrationSource.host_port
	HostPort *string `json:"hostPort,omitempty"`

	// Output only. Place holder for the external source identifier(e.g DMS job
	//  name) that created the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MigrationSource.reference_id
	ReferenceID *string `json:"referenceID,omitempty"`

	// Output only. Type of migration source.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MigrationSource.source_type
	SourceType *string `json:"sourceType,omitempty"`
}
