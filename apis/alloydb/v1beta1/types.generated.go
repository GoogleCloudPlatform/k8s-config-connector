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

package v1beta1

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

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.PrimaryConfig
type Cluster_PrimaryConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.PscConfig
type Cluster_PSCConfig struct {
	// Optional. Create an instance that allows connections from Private Service
	//  Connect endpoints to the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.PscConfig.psc_enabled
	PSCEnabled *bool `json:"pscEnabled,omitempty"`
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

// +kcc:proto=google.cloud.alloydb.v1beta.EncryptionInfo
type EncryptionInfo struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.GeminiClusterConfig
type GeminiClusterConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.GeminiInstanceConfig
type GeminiInstanceConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.ClientConnectionConfig
type Instance_ClientConnectionConfig struct {
	// Optional. Configuration to enforce connectors only (ex: AuthProxy)
	//  connections to the database.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ClientConnectionConfig.require_connectors
	RequireConnectors *bool `json:"requireConnectors,omitempty"`

	// Optional. SSL configuration option for this instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ClientConnectionConfig.ssl_config
	SSLConfig *SSLConfig `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig.AuthorizedNetwork
type Instance_InstanceNetworkConfig_AuthorizedNetwork struct {
	// CIDR range for one authorzied network of the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig.AuthorizedNetwork.cidr_range
	CIDRRange *string `json:"cidrRange,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.MachineConfig
type Instance_MachineConfig struct {
	// The number of CPU's in the VM instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.MachineConfig.cpu_count
	CPUCount *int32 `json:"cpuCount,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.Node
type Instance_Node struct {
	// The Compute Engine zone of the VM e.g. "us-central1-b".
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.Node.zone_id
	ZoneID *string `json:"zoneID,omitempty"`

	// The identifier of the VM e.g. "test-read-0601-407e52be-ms3l".
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.Node.id
	ID *string `json:"id,omitempty"`

	// The private IP address of the VM e.g. "10.57.0.34".
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.Node.ip
	IP *string `json:"ip,omitempty"`

	// Determined by state of the compute VM and postgres-service health.
	//  Compute VM state can have values listed in
	//  https://cloud.google.com/compute/docs/instances/instance-life-cycle and
	//  postgres-service health can have values: HEALTHY and UNHEALTHY.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.Node.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig
type Instance_ObservabilityInstanceConfig struct {
	// Observability feature status for an instance.
	//  This flag is turned "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Preserve comments in query string for an instance.
	//  This flag is turned "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.preserve_comments
	PreserveComments *bool `json:"preserveComments,omitempty"`

	// Track wait events during query execution for an instance.
	//  This flag is turned "on" by default but tracking is enabled only after
	//  observability enabled flag is also turned on.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.track_wait_events
	TrackWaitEvents *bool `json:"trackWaitEvents,omitempty"`

	// Query string length. The default value is 10k.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.max_query_string_length
	MaxQueryStringLength *int32 `json:"maxQueryStringLength,omitempty"`

	// Record application tags for an instance.
	//  This flag is turned "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.record_application_tags
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty"`

	// Number of query execution plans captured by Insights per minute
	//  for all queries combined. The default value is 200.
	//  Any integer between 0 to 200 is considered valid.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.query_plans_per_minute
	QueryPlansPerMinute *int32 `json:"queryPlansPerMinute,omitempty"`

	// Track actively running queries on the instance.
	//  If not set, this flag is "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.track_active_queries
	TrackActiveQueries *bool `json:"trackActiveQueries,omitempty"`

	// Track client address for an instance.
	//  If not set, default value is "off".
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.track_client_address
	TrackClientAddress *bool `json:"trackClientAddress,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.PscInstanceConfig
type Instance_PSCInstanceConfig struct {

	// Optional. List of consumer projects that are allowed to create
	//  PSC endpoints to service-attachments to this instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.PscInstanceConfig.allowed_consumer_projects
	AllowedConsumerProjects []string `json:"allowedConsumerProjects,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.QueryInsightsInstanceConfig
type Instance_QueryInsightsInstanceConfig struct {
	// Record application tags for an instance.
	//  This flag is turned "on" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.QueryInsightsInstanceConfig.record_application_tags
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty"`

	// Record client address for an instance. Client address is PII information.
	//  This flag is turned "on" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.QueryInsightsInstanceConfig.record_client_address
	RecordClientAddress *bool `json:"recordClientAddress,omitempty"`

	// Query string length. The default value is 1024.
	//  Any integer between 256 and 4500 is considered valid.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.QueryInsightsInstanceConfig.query_string_length
	QueryStringLength *uint32 `json:"queryStringLength,omitempty"`

	// Number of query execution plans captured by Insights per minute
	//  for all queries combined. The default value is 5.
	//  Any integer between 0 and 20 is considered valid.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.QueryInsightsInstanceConfig.query_plans_per_minute
	QueryPlansPerMinute *uint32 `json:"queryPlansPerMinute,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.ReadPoolConfig
type Instance_ReadPoolConfig struct {
	// Read capacity, i.e. number of nodes in a read pool instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ReadPoolConfig.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.UpdatePolicy
type Instance_UpdatePolicy struct {
	// Mode for updating the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.UpdatePolicy.mode
	Mode *string `json:"mode,omitempty"`
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

// +kcc:proto=google.cloud.alloydb.v1beta.SslConfig
type SSLConfig struct {
	// Optional. SSL mode. Specifies client-server SSL/TLS connection behavior.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.SslConfig.ssl_mode
	SSLMode *string `json:"sslMode,omitempty"`

	// Optional. Certificate Authority (CA) source. Only CA_SOURCE_MANAGED is
	//  supported currently, and is the default value.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.SslConfig.ca_source
	CASource *string `json:"caSource,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.PrimaryConfig
type Cluster_PrimaryConfigObservedState struct {
	// Output only. Names of the clusters that are replicating from this
	//  cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.PrimaryConfig.secondary_cluster_names
	SecondaryClusterNames []string `json:"secondaryClusterNames,omitempty"`
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

// +kcc:proto=google.cloud.alloydb.v1beta.GeminiInstanceConfig
type GeminiInstanceConfigObservedState struct {
	// Output only. Whether the Gemini in Databases add-on is enabled for the
	//  instance. It will be true only if the add-on has been enabled for the
	//  billing account corresponding to the instance. Its status is toggled from
	//  the Admin Control Center (ACC) and cannot be toggled using AlloyDB's APIs.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.GeminiInstanceConfig.entitled
	Entitled *bool `json:"entitled,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig
type Instance_ObservabilityInstanceConfigObservedState struct {
	// Output only. Track wait event types during query execution for an
	//  instance. This flag is turned "on" by default but tracking is enabled
	//  only after observability enabled flag is also turned on. This is
	//  read-only flag and only modifiable by producer API.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig.track_wait_event_types
	TrackWaitEventTypes *bool `json:"trackWaitEventTypes,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.PscInstanceConfig
type Instance_PSCInstanceConfigObservedState struct {
	// Output only. The service attachment created when Private
	//  Service Connect (PSC) is enabled for the instance.
	//  The name of the resource will be in the format of
	//  `projects/<alloydb-tenant-project-number>/regions/<region-name>/serviceAttachments/<service-attachment-name>`
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.PscInstanceConfig.service_attachment_link
	ServiceAttachmentLink *string `json:"serviceAttachmentLink,omitempty"`

	// Output only. The DNS name of the instance for PSC connectivity.
	//  Name convention: <uid>.<uid>.<region>.alloydb-psc.goog
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.PscInstanceConfig.psc_dns_name
	PSCDNSName *string `json:"pscDNSName,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.MaintenanceSchedule
type MaintenanceScheduleObservedState struct {
	// Output only. The scheduled start time for the maintenance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceSchedule.start_time
	StartTime *string `json:"startTime,omitempty"`
}
