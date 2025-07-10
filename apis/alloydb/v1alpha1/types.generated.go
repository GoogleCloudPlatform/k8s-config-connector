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
// krm.group: alloydb.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.alloydb.v1alpha
// resource: AlloyDBCluster:Cluster
// resource: AlloyDBInstance:Instance

package v1alpha1

// +kcc:proto=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy
type AutomatedBackupPolicy struct {
	// Weekly schedule for the Backup.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.weekly_schedule
	WeeklySchedule *AutomatedBackupPolicy_WeeklySchedule `json:"weeklySchedule,omitempty"`

	// Time-based Backup retention policy.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.time_based_retention
	TimeBasedRetention *AutomatedBackupPolicy_TimeBasedRetention `json:"timeBasedRetention,omitempty"`

	// Quantity-based Backup retention policy to retain recent backups.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.quantity_based_retention
	QuantityBasedRetention *AutomatedBackupPolicy_QuantityBasedRetention `json:"quantityBasedRetention,omitempty"`

	// Whether automated automated backups are enabled. If not set, defaults to
	//  true.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The length of the time window during which a backup can be
	//  taken. If a backup does not succeed within this time window, it will be
	//  canceled and considered failed.
	//
	//  The backup window must be at least 5 minutes long. There is no upper bound
	//  on the window. If not set, it defaults to 1 hour.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.backup_window
	BackupWindow *string `json:"backupWindow,omitempty"`

	// Optional. The encryption config can be specified to encrypt the
	//  backups with a customer-managed encryption key (CMEK). When this field is
	//  not specified, the backup will use the cluster's encryption config.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// The location where the backup will be stored. Currently, the only supported
	//  option is to store the backup in the same region as the cluster.
	//
	//  If empty, defaults to the region of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.location
	Location *string `json:"location,omitempty"`

	// Labels to apply to backups created using this configuration.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.QuantityBasedRetention
type AutomatedBackupPolicy_QuantityBasedRetention struct {
	// The number of backups to retain.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.QuantityBasedRetention.count
	Count *int32 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.TimeBasedRetention
type AutomatedBackupPolicy_TimeBasedRetention struct {
	// The retention period.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.AutomatedBackupPolicy.TimeBasedRetention.retention_period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.CloudSQLBackupRunSource
type CloudSQLBackupRunSource struct {
	// The project ID of the source CloudSQL instance. This should be the same as
	//  the AlloyDB cluster's project.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.CloudSQLBackupRunSource.project
	Project *string `json:"project,omitempty"`

	// Required. The CloudSQL instance ID.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.CloudSQLBackupRunSource.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// Required. The CloudSQL backup run ID.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.CloudSQLBackupRunSource.backup_run_id
	BackupRunID *int64 `json:"backupRunID,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Cluster
type Cluster struct {

	// User-settable and human-readable display name for the Cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The database engine major version. This is an optional field and
	//  it is populated at the Cluster creation time. If a database version is not
	//  supplied at cluster creation time, then a default database version will
	//  be used.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.network_config
	NetworkConfig *Cluster_NetworkConfig `json:"networkConfig,omitempty"`

	// Required. The resource link for the VPC network in which cluster resources
	//  are created and from which they are accessible via Private IP. The network
	//  must belong to the same project as the cluster. It is specified in the
	//  form: `projects/{project}/global/networks/{network_id}`. This is required
	//  to create a cluster. Deprecated, use network_config.network instead.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.network
	Network *string `json:"network,omitempty"`

	// For Resource freshness validation (https://google.aip.dev/154)
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.etag
	Etag *string `json:"etag,omitempty"`

	// Annotations to allow client tools to store small amount of arbitrary data.
	//  This is distinct from labels.
	//  https://google.aip.dev/128
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Input only. Initial user to setup during cluster creation. Required.
	//  If used in `RestoreCluster` this is ignored.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.initial_user
	InitialUser *UserPassword `json:"initialUser,omitempty"`

	// The automated backup policy for this cluster.
	//
	//  If no policy is provided then the default policy will be used. If backups
	//  are supported for the cluster, the default policy takes one backup a day,
	//  has a backup window of 1 hour, and retains backups for 14 days.
	//  For more information on the defaults, consult the
	//  documentation for the message type.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.automated_backup_policy
	AutomatedBackupPolicy *AutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`

	// SSL configuration for this AlloyDB cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.ssl_config
	SSLConfig *SSLConfig `json:"sslConfig,omitempty"`

	// Optional. The encryption config can be specified to encrypt the data disks
	//  and other persistent data resources of a cluster with a
	//  customer-managed encryption key (CMEK). When this field is not
	//  specified, the cluster will then use default encryption scheme to
	//  protect the user data.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Optional. Continuous backup configuration for this cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.continuous_backup_config
	ContinuousBackupConfig *ContinuousBackupConfig `json:"continuousBackupConfig,omitempty"`

	// Cross Region replication config specific to SECONDARY cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.secondary_config
	SecondaryConfig *Cluster_SecondaryConfig `json:"secondaryConfig,omitempty"`

	// Optional. The configuration for Private Service Connect (PSC) for the
	//  cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.psc_config
	PSCConfig *Cluster_PSCConfig `json:"pscConfig,omitempty"`

	// Optional. The maintenance update policy determines when to allow or deny
	//  updates.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.maintenance_update_policy
	MaintenanceUpdatePolicy *MaintenanceUpdatePolicy `json:"maintenanceUpdatePolicy,omitempty"`

	// Optional. Deprecated and unused. This field will be removed in the near
	//  future.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.gemini_config
	GeminiConfig *GeminiClusterConfig `json:"geminiConfig,omitempty"`

	// Optional. Subscription type of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.subscription_type
	SubscriptionType *string `json:"subscriptionType,omitempty"`

	// Optional. Input only. Immutable. Tag keys/values directly bound to this
	//  resource. For example:
	//  ```
	//  "123/environment": "production",
	//  "123/costCenter": "marketing"
	//  ```
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.tags
	Tags map[string]string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Cluster.PrimaryConfig
type Cluster_PrimaryConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Cluster.PscConfig
type Cluster_PSCConfig struct {
	// Optional. Create an instance that allows connections from Private Service
	//  Connect endpoints to the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.PscConfig.psc_enabled
	PSCEnabled *bool `json:"pscEnabled,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Cluster.TrialMetadata
type Cluster_TrialMetadata struct {
	// start time of the trial cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.TrialMetadata.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End time of the trial cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.TrialMetadata.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Upgrade time of trial cluster to Standard cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.TrialMetadata.upgrade_time
	UpgradeTime *string `json:"upgradeTime,omitempty"`

	// grace end time of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.TrialMetadata.grace_end_time
	GraceEndTime *string `json:"graceEndTime,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.ContinuousBackupConfig
type ContinuousBackupConfig struct {
	// Whether ContinuousBackup is enabled.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.ContinuousBackupConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The number of days that are eligible to restore from using PITR. To support
	//  the entire recovery window, backups and logs are retained for one day more
	//  than the recovery window. If not set, defaults to 14 days.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.ContinuousBackupConfig.recovery_window_days
	RecoveryWindowDays *int32 `json:"recoveryWindowDays,omitempty"`

	// The encryption config can be specified to encrypt the
	//  backups with a customer-managed encryption key (CMEK). When this field is
	//  not specified, the backup will use the cluster's encryption config.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.ContinuousBackupConfig.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.ContinuousBackupInfo
type ContinuousBackupInfo struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.EncryptionInfo
type EncryptionInfo struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.GCAInstanceConfig
type GcaInstanceConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.GeminiClusterConfig
type GeminiClusterConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.GeminiInstanceConfig
type GeminiInstanceConfig struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance
type Instance struct {

	// User-settable and human-readable display name for the Instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The type of the instance. Specified at creation time.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.instance_type
	InstanceType *string `json:"instanceType,omitempty"`

	// Configurations for the machines that host the underlying
	//  database engine.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.machine_config
	MachineConfig *Instance_MachineConfig `json:"machineConfig,omitempty"`

	// Availability type of an Instance.
	//  If empty, defaults to REGIONAL for primary instances.
	//  For read pools, availability_type is always UNSPECIFIED. Instances in the
	//  read pools are evenly distributed across available zones within the region
	//  (i.e. read pools with more than one node will have a node in at
	//  least two zones).
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.availability_type
	AvailabilityType *string `json:"availabilityType,omitempty"`

	// The Compute Engine zone that the instance should serve from, per
	//  https://cloud.google.com/compute/docs/regions-zones
	//  This can ONLY be specified for ZONAL instances.
	//  If present for a REGIONAL instance, an error will be thrown.
	//  If this is absent for a ZONAL instance, instance is created in a random
	//  zone with available capacity.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.gce_zone
	GCEZone *string `json:"gceZone,omitempty"`

	// Database flags. Set at the instance level.
	//  They are copied from the primary instance on secondary instance creation.
	//  Flags that have restrictions default to the value at primary
	//  instance on read instances during creation. Read instances can set new
	//  flags or override existing flags that are relevant for reads, for example,
	//  for enabling columnar cache on a read instance. Flags set on read instance
	//  might or might not be present on the primary instance.
	//
	//
	//  This is a list of "key": "value" pairs.
	//  "key": The name of the flag. These flags are passed at instance setup time,
	//  so include both server options and system variables for Postgres. Flags are
	//  specified with underscores, not hyphens.
	//  "value": The value of the flag. Booleans are set to **on** for true
	//  and **off** for false. This field must be omitted if the flag
	//  doesn't take a value.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.database_flags
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty"`

	// Configuration for query insights.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.query_insights_config
	QueryInsightsConfig *Instance_QueryInsightsInstanceConfig `json:"queryInsightsConfig,omitempty"`

	// Configuration for observability.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.observability_config
	ObservabilityConfig *Instance_ObservabilityInstanceConfig `json:"observabilityConfig,omitempty"`

	// Read pool instance configuration.
	//  This is required if the value of instanceType is READ_POOL.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.read_pool_config
	ReadPoolConfig *Instance_ReadPoolConfig `json:"readPoolConfig,omitempty"`

	// For Resource freshness validation (https://google.aip.dev/154)
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.etag
	Etag *string `json:"etag,omitempty"`

	// Annotations to allow client tools to store small amount of arbitrary data.
	//  This is distinct from labels.
	//  https://google.aip.dev/128
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Update policy that will be applied during instance update.
	//  This field is not persisted when you update the instance.
	//  To use a non-default update policy, you must
	//  specify explicitly specify the value in each update request.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.update_policy
	UpdatePolicy *Instance_UpdatePolicy `json:"updatePolicy,omitempty"`

	// Optional. Client connection specific configurations
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.client_connection_config
	ClientConnectionConfig *Instance_ClientConnectionConfig `json:"clientConnectionConfig,omitempty"`

	// Optional. The configuration for Private Service Connect (PSC) for the
	//  instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.psc_instance_config
	PSCInstanceConfig *Instance_PSCInstanceConfig `json:"pscInstanceConfig,omitempty"`

	// Optional. Instance-level network configuration.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.network_config
	NetworkConfig *Instance_InstanceNetworkConfig `json:"networkConfig,omitempty"`

	// Optional. Deprecated and unused. This field will be removed in the near
	//  future.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.gemini_config
	GeminiConfig *GeminiInstanceConfig `json:"geminiConfig,omitempty"`

	// Optional. Specifies whether an instance needs to spin up. Once the instance
	//  is active, the activation policy can be updated to the `NEVER` to stop the
	//  instance. Likewise, the activation policy can be updated to `ALWAYS` to
	//  start the instance.
	//  There are restrictions around when an instance can/cannot be activated (for
	//  example, a read pool instance should be stopped before stopping primary
	//  etc.). Please refer to the API documentation for more details.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.activation_policy
	ActivationPolicy *string `json:"activationPolicy,omitempty"`

	// Optional. The configuration for Managed Connection Pool (MCP).
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.connection_pool_config
	ConnectionPoolConfig *Instance_ConnectionPoolConfig `json:"connectionPoolConfig,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.ClientConnectionConfig
type Instance_ClientConnectionConfig struct {
	// Optional. Configuration to enforce connectors only (ex: AuthProxy)
	//  connections to the database.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ClientConnectionConfig.require_connectors
	RequireConnectors *bool `json:"requireConnectors,omitempty"`

	// Optional. SSL configuration option for this instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ClientConnectionConfig.ssl_config
	SSLConfig *SSLConfig `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.ConnectionPoolConfig
type Instance_ConnectionPoolConfig struct {
	// Optional. Whether to enable Managed Connection Pool (MCP).
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ConnectionPoolConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.InstanceNetworkConfig.AuthorizedNetwork
type Instance_InstanceNetworkConfig_AuthorizedNetwork struct {
	// CIDR range for one authorzied network of the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.InstanceNetworkConfig.AuthorizedNetwork.cidr_range
	CIDRRange *string `json:"cidrRange,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.MachineConfig
type Instance_MachineConfig struct {
	// The number of CPU's in the VM instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.MachineConfig.cpu_count
	CPUCount *int32 `json:"cpuCount,omitempty"`

	// Machine type of the VM instance. E.g. "n2-highmem-4",
	//  "n2-highmem-8", "c4a-highmem-4-lssd".
	//  cpu_count must match the number of vCPUs in the machine type.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.MachineConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.Node
type Instance_Node struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig
type Instance_ObservabilityInstanceConfig struct {
	// Observability feature status for an instance.
	//  This flag is turned "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Preserve comments in query string for an instance.
	//  This flag is turned "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.preserve_comments
	PreserveComments *bool `json:"preserveComments,omitempty"`

	// Track wait events during query execution for an instance.
	//  This flag is turned "on" by default but tracking is enabled only after
	//  observability enabled flag is also turned on.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.track_wait_events
	TrackWaitEvents *bool `json:"trackWaitEvents,omitempty"`

	// Query string length. The default value is 10k.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.max_query_string_length
	MaxQueryStringLength *int32 `json:"maxQueryStringLength,omitempty"`

	// Record application tags for an instance.
	//  This flag is turned "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.record_application_tags
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty"`

	// Number of query execution plans captured by Insights per minute
	//  for all queries combined. The default value is 200.
	//  Any integer between 0 to 200 is considered valid.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.query_plans_per_minute
	QueryPlansPerMinute *int32 `json:"queryPlansPerMinute,omitempty"`

	// Track actively running queries on the instance.
	//  If not set, this flag is "off" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.track_active_queries
	TrackActiveQueries *bool `json:"trackActiveQueries,omitempty"`

	// Track client address for an instance.
	//  If not set, default value is "off".
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.track_client_address
	TrackClientAddress *bool `json:"trackClientAddress,omitempty"`

	// Whether assistive experiences are enabled for this AlloyDB instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.assistive_experiences_enabled
	AssistiveExperiencesEnabled *bool `json:"assistiveExperiencesEnabled,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.PscAutoConnectionConfig
type Instance_PSCAutoConnectionConfig struct {
	// The consumer project to which the PSC service automation endpoint will
	//  be created.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscAutoConnectionConfig.consumer_project
	ConsumerProject *string `json:"consumerProject,omitempty"`

	// The consumer network for the PSC service automation, example:
	//  "projects/vpc-host-project/global/networks/default".
	//  The consumer network might be hosted a different project than the
	//  consumer project.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscAutoConnectionConfig.consumer_network
	ConsumerNetwork *string `json:"consumerNetwork,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig
type Instance_PSCInstanceConfig struct {

	// Optional. List of consumer projects that are allowed to create
	//  PSC endpoints to service-attachments to this instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig.allowed_consumer_projects
	AllowedConsumerProjects []string `json:"allowedConsumerProjects,omitempty"`

	// Optional. Configurations for setting up PSC interfaces attached to the
	//  instance which are used for outbound connectivity. Only primary instances
	//  can have PSC interface attached. Currently we only support 0 or 1 PSC
	//  interface.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig.psc_interface_configs
	PSCInterfaceConfigs []Instance_PSCInterfaceConfig `json:"pscInterfaceConfigs,omitempty"`

	// Optional. Configurations for setting up PSC service automation.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig.psc_auto_connections
	PSCAutoConnections []Instance_PSCAutoConnectionConfig `json:"pscAutoConnections,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.PscInterfaceConfig
type Instance_PSCInterfaceConfig struct {
	// The network attachment resource created in the consumer network to which
	//  the PSC interface will be linked. This is of the format:
	//  "projects/${CONSUMER_PROJECT}/regions/${REGION}/networkAttachments/${NETWORK_ATTACHMENT_NAME}".
	//  The network attachment must be in the same region as the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscInterfaceConfig.network_attachment_resource
	NetworkAttachmentResource *string `json:"networkAttachmentResource,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.QueryInsightsInstanceConfig
type Instance_QueryInsightsInstanceConfig struct {
	// Record application tags for an instance.
	//  This flag is turned "on" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.QueryInsightsInstanceConfig.record_application_tags
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty"`

	// Record client address for an instance. Client address is PII information.
	//  This flag is turned "on" by default.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.QueryInsightsInstanceConfig.record_client_address
	RecordClientAddress *bool `json:"recordClientAddress,omitempty"`

	// Query string length. The default value is 1024.
	//  Any integer between 256 and 4500 is considered valid.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.QueryInsightsInstanceConfig.query_string_length
	QueryStringLength *uint32 `json:"queryStringLength,omitempty"`

	// Number of query execution plans captured by Insights per minute
	//  for all queries combined. The default value is 5.
	//  Any integer between 0 and 20 is considered valid.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.QueryInsightsInstanceConfig.query_plans_per_minute
	QueryPlansPerMinute *uint32 `json:"queryPlansPerMinute,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.ReadPoolConfig
type Instance_ReadPoolConfig struct {
	// Read capacity, i.e. number of nodes in a read pool instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ReadPoolConfig.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.Instance.UpdatePolicy
type Instance_UpdatePolicy struct {
	// Mode for updating the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.UpdatePolicy.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.MaintenanceSchedule
type MaintenanceSchedule struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.MaintenanceUpdatePolicy
type MaintenanceUpdatePolicy struct {
	// Preferred windows to perform maintenance. Currently limited to 1.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.MaintenanceUpdatePolicy.maintenance_windows
	MaintenanceWindows []MaintenanceUpdatePolicy_MaintenanceWindow `json:"maintenanceWindows,omitempty"`

	// Periods to deny maintenance. Currently limited to 1.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.MaintenanceUpdatePolicy.deny_maintenance_periods
	DenyMaintenancePeriods []MaintenanceUpdatePolicy_DenyMaintenancePeriod `json:"denyMaintenancePeriods,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.MaintenanceUpdatePolicy.DenyMaintenancePeriod
type MaintenanceUpdatePolicy_DenyMaintenancePeriod struct {
	// Deny period start date.
	//  This can be:
	//  * A full date, with non-zero year, month and day values OR
	//  * A month and day value, with a zero year for recurring
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.MaintenanceUpdatePolicy.DenyMaintenancePeriod.start_date
	StartDate *Date `json:"startDate,omitempty"`

	// Deny period end date.
	//  This can be:
	//  * A full date, with non-zero year, month and day values OR
	//  * A month and day value, with a zero year for recurring
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.MaintenanceUpdatePolicy.DenyMaintenancePeriod.end_date
	EndDate *Date `json:"endDate,omitempty"`

	// Time in UTC when the deny period starts on start_date and ends on
	//  end_date. This can be:
	//  * Full time OR
	//  * All zeros for 00:00:00 UTC
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.MaintenanceUpdatePolicy.DenyMaintenancePeriod.time
	Time *TimeOfDay `json:"time,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1alpha.MigrationSource
type MigrationSource struct {
}

// +kcc:proto=google.cloud.alloydb.v1alpha.SslConfig
type SSLConfig struct {
	// Optional. SSL mode. Specifies client-server SSL/TLS connection behavior.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.SslConfig.ssl_mode
	SSLMode *string `json:"sslMode,omitempty"`

	// Optional. Certificate Authority (CA) source. Only CA_SOURCE_MANAGED is
	//  supported currently, and is the default value.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.SslConfig.ca_source
	CASource *string `json:"caSource,omitempty"`
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

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Cluster.PrimaryConfig
type Cluster_PrimaryConfigObservedState struct {
	// Output only. Names of the clusters that are replicating from this
	//  cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.PrimaryConfig.secondary_cluster_names
	SecondaryClusterNames []string `json:"secondaryClusterNames,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Cluster.PscConfig
type Cluster_PSCConfigObservedState struct {
	// Output only. The project number that needs to be allowlisted on the
	//  network attachment to enable outbound connectivity.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Cluster.PscConfig.service_owned_project_number
	ServiceOwnedProjectNumber *int64 `json:"serviceOwnedProjectNumber,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.EncryptionInfo
type EncryptionInfoObservedState struct {
	// Output only. Type of encryption.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.EncryptionInfo.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Output only. Cloud KMS key versions that are being used to protect the
	//  database or the backup.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.EncryptionInfo.kms_key_versions
	KMSKeyVersions []string `json:"kmsKeyVersions,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.GCAInstanceConfig
type GcaInstanceConfigObservedState struct {
	// Output only. Represents the GCA entitlement state of the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.GCAInstanceConfig.gca_entitlement
	GcaEntitlement *string `json:"gcaEntitlement,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.GeminiClusterConfig
type GeminiClusterConfigObservedState struct {
	// Output only. Deprecated and unused. This field will be removed in the near
	//  future.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.GeminiClusterConfig.entitled
	Entitled *bool `json:"entitled,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.GeminiInstanceConfig
type GeminiInstanceConfigObservedState struct {
	// Output only. Deprecated and unused. This field will be removed in the near
	//  future.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.GeminiInstanceConfig.entitled
	Entitled *bool `json:"entitled,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Instance
type InstanceObservedState struct {
	// Output only. The name of the instance resource with the format:
	//   * projects/{project}/locations/{region}/clusters/{cluster_id}/instances/{instance_id}
	//  where the cluster and instance ID segments should satisfy the regex
	//  expression `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`, e.g. 1-63 characters of
	//  lowercase letters, numbers, and dashes, starting with a letter, and ending
	//  with a letter or number. For more details see https://google.aip.dev/122.
	//  The prefix of the instance resource name is the name of the parent
	//  resource:
	//   * projects/{project}/locations/{region}/clusters/{cluster_id}
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.name
	Name *string `json:"name,omitempty"`

	// Output only. The system-generated UID of the resource. The UID is assigned
	//  when the resource is created, and it is retained until it is deleted.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Create time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Delete time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The current serving state of the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. This is set for the read-write VM of the PRIMARY instance
	//  only.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.writable_node
	WritableNode *Instance_Node `json:"writableNode,omitempty"`

	// Output only. List of available read-only VMs in this instance, including
	//  the standby for a PRIMARY instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.nodes
	Nodes []Instance_Node `json:"nodes,omitempty"`

	// Configuration for observability.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.observability_config
	ObservabilityConfig *Instance_ObservabilityInstanceConfigObservedState `json:"observabilityConfig,omitempty"`

	// Output only. The IP address for the Instance.
	//  This is the connection endpoint for an end-user application.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Output only. The public IP addresses for the Instance. This is available
	//  ONLY when enable_public_ip is set. This is the connection endpoint for an
	//  end-user application.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.public_ip_address
	PublicIPAddress *string `json:"publicIPAddress,omitempty"`

	// Output only. Reconciling (https://google.aip.dev/128#reconciliation).
	//  Set to true if the current state of Instance does not match the user's
	//  intended state, and the service is actively updating the resource to
	//  reconcile them. This can happen due to user-triggered updates or
	//  system actions like failover or maintenance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Optional. The configuration for Private Service Connect (PSC) for the
	//  instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.psc_instance_config
	PSCInstanceConfig *Instance_PSCInstanceConfigObservedState `json:"pscInstanceConfig,omitempty"`

	// Optional. Instance-level network configuration.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.network_config
	NetworkConfig *Instance_InstanceNetworkConfigObservedState `json:"networkConfig,omitempty"`

	// Optional. Deprecated and unused. This field will be removed in the near
	//  future.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.gemini_config
	GeminiConfig *GeminiInstanceConfigObservedState `json:"geminiConfig,omitempty"`

	// Output only. All outbound public IP addresses configured for the instance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.outbound_public_ip_addresses
	OutboundPublicIPAddresses []string `json:"outboundPublicIPAddresses,omitempty"`

	// Output only. Configuration parameters related to Gemini Cloud Assist.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.gca_config
	GcaConfig *GcaInstanceConfig `json:"gcaConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Instance.InstanceNetworkConfig
type Instance_InstanceNetworkConfigObservedState struct {
	// Output only. The resource link for the VPC network in which instance
	//  resources are created and from which they are accessible via Private IP.
	//  This will be the same value as the parent cluster's network. It is
	//  specified in the form: //
	//  `projects/{project_number}/global/networks/{network_id}`.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.InstanceNetworkConfig.network
	Network *string `json:"network,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Instance.Node
type Instance_NodeObservedState struct {
	// Output only. The Compute Engine zone of the VM e.g. "us-central1-b".
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.Node.zone_id
	ZoneID *string `json:"zoneID,omitempty"`

	// Output only. The identifier of the VM e.g.
	//  "test-read-0601-407e52be-ms3l".
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.Node.id
	ID *string `json:"id,omitempty"`

	// Output only. The private IP address of the VM e.g. "10.57.0.34".
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.Node.ip
	IP *string `json:"ip,omitempty"`

	// Output only. Determined by state of the compute VM and postgres-service
	//  health. Compute VM state can have values listed in
	//  https://cloud.google.com/compute/docs/instances/instance-life-cycle and
	//  postgres-service health can have values: HEALTHY and UNHEALTHY.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.Node.state
	State *string `json:"state,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig
type Instance_ObservabilityInstanceConfigObservedState struct {
	// Output only. Track wait event types during query execution for an
	//  instance. This flag is turned "on" by default but tracking is enabled
	//  only after observability enabled flag is also turned on. This is
	//  read-only flag and only modifiable by internal API.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.ObservabilityInstanceConfig.track_wait_event_types
	TrackWaitEventTypes *bool `json:"trackWaitEventTypes,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Instance.PscAutoConnectionConfig
type Instance_PSCAutoConnectionConfigObservedState struct {
	// Output only. The IP address of the PSC service automation endpoint.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscAutoConnectionConfig.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Output only. The status of the PSC service automation connection.
	//  Possible values:
	//    "STATE_UNSPECIFIED" - An invalid state as the default case.
	//    "ACTIVE" - The connection has been created successfully.
	//    "FAILED" - The connection is not functional since some resources on the
	//  connection fail to be created.
	//    "CREATING" - The connection is being created.
	//    "DELETING" - The connection is being deleted.
	//    "CREATE_REPAIRING" - The connection is being repaired to complete
	//  creation.
	//    "DELETE_REPAIRING" - The connection is being repaired to complete
	//  deletion.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscAutoConnectionConfig.status
	Status *string `json:"status,omitempty"`

	// Output only. The status of the service connection policy.
	//  Possible values:
	//    "STATE_UNSPECIFIED" - Default state, when Connection Map is created
	//  initially.
	//    "VALID" - Set when policy and map configuration is valid, and their
	//  matching can lead to allowing creation of PSC Connections subject to
	//  other constraints like connections limit.
	//    "CONNECTION_POLICY_MISSING" - No Service Connection Policy found for
	//  this network and Service Class
	//    "POLICY_LIMIT_REACHED" - Service Connection Policy limit reached for
	//    this network and Service Class
	//    "CONSUMER_INSTANCE_PROJECT_NOT_ALLOWLISTED" - The consumer instance
	//  project is not in AllowedGoogleProducersResourceHierarchyLevels of the
	//  matching ServiceConnectionPolicy.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscAutoConnectionConfig.consumer_network_status
	ConsumerNetworkStatus *string `json:"consumerNetworkStatus,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig
type Instance_PSCInstanceConfigObservedState struct {
	// Output only. The service attachment created when Private
	//  Service Connect (PSC) is enabled for the instance.
	//  The name of the resource will be in the format of
	//  `projects/<alloydb-tenant-project-number>/regions/<region-name>/serviceAttachments/<service-attachment-name>`
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig.service_attachment_link
	ServiceAttachmentLink *string `json:"serviceAttachmentLink,omitempty"`

	// Output only. The DNS name of the instance for PSC connectivity.
	//  Name convention: <uid>.<uid>.<region>.alloydb-psc.goog
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig.psc_dns_name
	PSCDNSName *string `json:"pscDNSName,omitempty"`

	// Optional. Configurations for setting up PSC service automation.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.Instance.PscInstanceConfig.psc_auto_connections
	PSCAutoConnections []Instance_PSCAutoConnectionConfigObservedState `json:"pscAutoConnections,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.alloydb.v1alpha.MaintenanceSchedule
type MaintenanceScheduleObservedState struct {
	// Output only. The scheduled start time for the maintenance.
	// +kcc:proto:field=google.cloud.alloydb.v1alpha.MaintenanceSchedule.start_time
	StartTime *string `json:"startTime,omitempty"`
}
