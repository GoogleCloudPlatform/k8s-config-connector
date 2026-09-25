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

package v1alpha1

import (
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OracleDatabaseAutonomousDatabaseGVK = GroupVersion.WithKind("OracleDatabaseAutonomousDatabase")

// OracleDatabaseAutonomousDatabaseSpec defines the desired state of OracleDatabaseAutonomousDatabase
// +kcc:spec:proto=google.cloud.oracledatabase.v1.AutonomousDatabase
type OracleDatabaseAutonomousDatabaseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The OracleDatabaseAutonomousDatabase name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Immutable. The name of the Autonomous Database. The database name
	//  must be unique in the project. The name must begin with a letter and can
	//  contain a maximum of 30 alphanumeric characters.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.database
	Database *string `json:"database,omitempty"`

	// Optional. Immutable. The display name for the Autonomous Database. The name
	//  does not have to be unique within your project.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Immutable. The password for the default ADMIN user.
	//  Note: Only one of `admin_password_secret_version` or `admin_password` can
	//  be populated.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.admin_password
	AdminPassword *refsv1beta1secret.Legacy `json:"adminPassword,omitempty"`

	// Optional. Immutable. The resource name of a secret version in Secret
	//  Manager which contains the database admin user's password.
	AdminPasswordSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"adminPasswordSecretVersionRef,omitempty"`

	// Optional. The properties of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.properties
	Properties *AutonomousDatabaseProperties `json:"properties,omitempty"`

	// Optional. The labels or tags associated with the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Immutable. The ComputeNetwork associated with the Autonomous Database.
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Immutable. The subnet CIDR range for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.cidr
	CIDR *string `json:"cidr,omitempty"`

	// Optional. Immutable. The OdbNetwork associated with the Autonomous Database.
	OdbNetworkRef *OracleDatabaseODBNetworkRef `json:"odbNetworkRef,omitempty"`

	// Optional. Immutable. The OdbSubnet associated with the Autonomous Database.
	OdbSubnetRef *OracleDatabaseODBSubnetRef `json:"odbSubnetRef,omitempty"`

	// Optional. Immutable. The source Autonomous Database configuration for the
	//  standby Autonomous Database. The source Autonomous Database is configured
	//  while creating the Peer Autonomous Database and can't be updated after
	//  creation.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.source_config
	SourceConfig *SourceConfig `json:"sourceConfig,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.SourceConfig
type SourceConfig struct {
	// Optional. The primary Autonomous Database that is used to
	//  create a Peer Autonomous Database from a source.
	AutonomousDatabaseRef *OracleDatabaseAutonomousDatabaseRef `json:"autonomousDatabaseRef,omitempty"`

	// Optional. This field specifies if the replication of automatic backups is
	//  enabled when creating a Data Guard.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.automatic_backups_replication_enabled
	AutomaticBackupsReplicationEnabled *bool `json:"automaticBackupsReplicationEnabled,omitempty"`

	// Optional. The source type of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// Optional. The clone type of the Autonomous Database. This field is only
	//  applicable in case of cloning
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.clone_type
	CloneType *string `json:"cloneType,omitempty"`

	// Optional. The refresh mode of the clone.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.refreshable_mode
	RefreshableMode *string `json:"refreshableMode,omitempty"`

	// Optional. The frequency in seconds a refreshable clone is refreshed after
	//  auto-refresh is enabled.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.auto_refresh_frequency_seconds
	AutoRefreshFrequencySeconds *int32 `json:"autoRefreshFrequencySeconds,omitempty"`

	// Optional. The time, in seconds, the data of the automatic refreshable clone
	//  lags the primary database at the point of refresh.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.auto_refresh_point_lag_seconds
	AutoRefreshPointLagSeconds *int32 `json:"autoRefreshPointLagSeconds,omitempty"`

	// Optional. The date and time that auto-refreshing will begin for an
	//  Autonomous Database refreshable clone. This value controls only the start
	//  time for the first refresh operation.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.auto_refresh_start_time
	AutoRefreshStartTime *string `json:"autoRefreshStartTime,omitempty"`

	// Optional. The Autonomous Database Backup resource.
	//  Required when source_type is BACKUP_FROM_ID.
	AutonomousDatabaseBackupRef *OracleDatabaseAutonomousDatabaseBackupRef `json:"autonomousDatabaseBackupRef,omitempty"`

	// Optional. The timestamp specified for the point-in-time clone of the source
	//  Autonomous Database. This field is only applicable
	//  in case of BACKUP_FROM_TIMESTAMP source type and when
	//  use_latest_available_backup is false.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.backup_time
	BackupTime *string `json:"backupTime,omitempty"`

	// Optional. Clone from latest available backup timestamp. This field is only
	//  applicable in case of BACKUP_FROM_TIMESTAMP source type.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.SourceConfig.use_latest_available_backup
	UseLatestAvailableBackup *bool `json:"useLatestAvailableBackup,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.EncryptionKey
type EncryptionKey struct {
	// Optional. The provider of the encryption key.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.EncryptionKey.provider
	Provider *string `json:"provider,omitempty"`

	// Optional. The KMS key used to encrypt the Autonomous Database.
	//  This field is required if the provider is GOOGLE_MANAGED.
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.EncryptionKey
type EncryptionKeyObservedState struct {
	// Optional. The provider of the encryption key.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.EncryptionKey.provider
	Provider *string `json:"provider,omitempty"`

	// Optional. The KMS key used to encrypt the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.EncryptionKey.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.EncryptionKeyHistoryEntry
type EncryptionKeyHistoryEntryObservedState struct {
	// Output only. The encryption key used to encrypt the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.EncryptionKeyHistoryEntry.encryption_key
	EncryptionKey *EncryptionKeyObservedState `json:"encryptionKey,omitempty"`

	// Output only. The date and time when the encryption key was activated on the
	//  Autonomous Database..
	// +kcc:proto:field=google.cloud.oracledatabase.v1.EncryptionKeyHistoryEntry.activation_time
	ActivationTime *string `json:"activationTime,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties
type AutonomousDatabaseProperties struct {
	// Optional. Immutable. The number of compute servers for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.compute_count
	ComputeCount *float32 `json:"computeCount,omitempty"`

	// Optional. Immutable. The number of CPU cores to be made available to the
	//  database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.cpu_core_count
	CPUCoreCount *int32 `json:"cpuCoreCount,omitempty"`

	// Optional. Immutable. The size of the data stored in the database, in
	//  terabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.data_storage_size_tb
	DataStorageSizeTb *int32 `json:"dataStorageSizeTb,omitempty"`

	// Optional. Immutable. The size of the data stored in the database, in
	//  gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.data_storage_size_gb
	DataStorageSizeGB *int32 `json:"dataStorageSizeGB,omitempty"`

	// Required. Immutable. The workload type of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.db_workload
	DbWorkload *string `json:"dbWorkload,omitempty"`

	// Optional. Immutable. The edition of the Autonomous Databases.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.db_edition
	DbEdition *string `json:"dbEdition,omitempty"`

	// Optional. Immutable. The character set for the Autonomous Database. The
	//  default is AL32UTF8.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.character_set
	CharacterSet *string `json:"characterSet,omitempty"`

	// Optional. Immutable. The national character set for the Autonomous
	//  Database. The default is AL16UTF16.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.n_character_set
	NCharacterSet *string `json:"nCharacterSet,omitempty"`

	// Optional. Immutable. The private endpoint IP address for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.private_endpoint_ip
	PrivateEndpointIP *string `json:"privateEndpointIP,omitempty"`

	// Optional. Immutable. The private endpoint label for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.private_endpoint_label
	PrivateEndpointLabel *string `json:"privateEndpointLabel,omitempty"`

	// Optional. Immutable. The Oracle Database version for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.db_version
	DbVersion *string `json:"dbVersion,omitempty"`

	// Optional. Immutable. This field indicates if auto scaling is enabled for
	//  the Autonomous Database CPU core count.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.is_auto_scaling_enabled
	IsAutoScalingEnabled *bool `json:"isAutoScalingEnabled,omitempty"`

	// Optional. Immutable. This field indicates if auto scaling is enabled for
	//  the Autonomous Database storage.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.is_storage_auto_scaling_enabled
	IsStorageAutoScalingEnabled *bool `json:"isStorageAutoScalingEnabled,omitempty"`

	// Required. Immutable. The license type used for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.license_type
	LicenseType *string `json:"licenseType,omitempty"`

	// Optional. Immutable. The list of customer contacts.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.customer_contacts
	CustomerContacts []CustomerContact `json:"customerContacts,omitempty"`

	// Optional. Immutable. The ID of the Oracle Cloud Infrastructure vault
	//  secret.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.secret_id
	SecretID *string `json:"secretID,omitempty"`

	// Optional. Immutable. The ID of the Oracle Cloud Infrastructure vault.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.vault_id
	VaultID *string `json:"vaultID,omitempty"`

	// Optional. Immutable. The maintenance schedule of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.maintenance_schedule_type
	MaintenanceScheduleType *string `json:"maintenanceScheduleType,omitempty"`

	// Optional. Immutable. This field specifies if the Autonomous Database
	//  requires mTLS connections.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.mtls_connection_required
	MtlsConnectionRequired *bool `json:"mtlsConnectionRequired,omitempty"`

	// Optional. Immutable. The retention period for the Autonomous Database. This
	//  field is specified in days, can range from 1 day to 60 days, and has a
	//  default value of 60 days.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.backup_retention_period_days
	BackupRetentionPeriodDays *int32 `json:"backupRetentionPeriodDays,omitempty"`

	// Optional. Immutable. The list of allowlisted IP addresses for the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.allowlisted_ips
	AllowlistedIPs []string `json:"allowlistedIPs,omitempty"`

	// Optional. The encryption key used to encrypt the Autonomous Database.
	//  Updating this field will add a new entry in the
	//  `encryption_key_history_entries` field with the former version.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.encryption_key
	EncryptionKey *EncryptionKey `json:"encryptionKey,omitempty"`

	// Optional. Indicates whether the Autonomous Database has a local (in-region)
	//  standby database. Not applicable to cross-region Data Guard or dedicated
	//  Exadata infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.local_data_guard_enabled
	LocalDataGuardEnabled *bool `json:"localDataGuardEnabled,omitempty"`

	// Optional. This field indicates the maximum data loss limit for an
	//  Autonomous Database, in seconds.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.local_adg_auto_failover_max_data_loss_limit_duration
	LocalAdgAutoFailoverMaxDataLossLimitDuration *int32 `json:"localAdgAutoFailoverMaxDataLossLimitDuration,omitempty"`

	// Optional. Indicates if the Autonomous Database is a refreshable clone. This
	//  field is used in update flow to connect / disconnect a refreshable clone
	//  from its source database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.refreshable_clone
	RefreshableClone *bool `json:"refreshableClone,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls
type AutonomousDatabaseConnectionURLsObservedState struct {
	// Output only. Oracle Application Express (APEX) URL.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.apex_uri
	ApexURI *string `json:"apexURI,omitempty"`

	// Output only. The URL of the Database Transforms for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.database_transforms_uri
	DatabaseTransformsURI *string `json:"databaseTransformsURI,omitempty"`

	// Output only. The URL of the Graph Studio for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.graph_studio_uri
	GraphStudioURI *string `json:"graphStudioURI,omitempty"`

	// Output only. The URL of the Oracle Machine Learning (OML) Notebook for the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.machine_learning_notebook_uri
	MachineLearningNotebookURI *string `json:"machineLearningNotebookURI,omitempty"`

	// Output only. The URL of Machine Learning user management the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.machine_learning_user_management_uri
	MachineLearningUserManagementURI *string `json:"machineLearningUserManagementURI,omitempty"`

	// Output only. The URL of the MongoDB API for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.mongo_db_uri
	MongoDbURI *string `json:"mongoDbURI,omitempty"`

	// Output only. The Oracle REST Data Services (ORDS) URL of the Web Access for
	//  the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.ords_uri
	OrdsURI *string `json:"ordsURI,omitempty"`

	// Output only. The URL of the Oracle SQL Developer Web for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls.sql_dev_web_uri
	SQLDevWebURI *string `json:"sqlDevWebURI,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties
type AutonomousDatabasePropertiesObservedState struct {
	// Output only. OCID of the Autonomous Database.
	//  https://docs.oracle.com/en-us/iaas/Content/General/Concepts/identifiers.htm#Oracle
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.ocid
	Ocid *string `json:"ocid,omitempty"`

	// Output only. The amount of storage currently being used for user and system
	//  data, in terabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.actual_used_data_storage_size_tb
	ActualUsedDataStorageSizeTb *float64 `json:"actualUsedDataStorageSizeTb,omitempty"`

	// Output only. The amount of storage currently allocated for the database
	//  tables and billed for, rounded up in terabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.allocated_storage_size_tb
	AllocatedStorageSizeTb *float64 `json:"allocatedStorageSizeTb,omitempty"`

	// Output only. The details for the Oracle APEX Application Development.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.apex_details
	ApexDetails *AutonomousDatabaseApexObservedState `json:"apexDetails,omitempty"`

	// Output only. This field indicates the status of Data Guard and Access
	//  control for the Autonomous Database. The field's value is null if Data
	//  Guard is disabled or Access Control is disabled. The field's value is TRUE
	//  if both Data Guard and Access Control are enabled, and the Autonomous
	//  Database is using primary IP access control list (ACL) for standby. The
	//  field's value is FALSE if both Data Guard and Access Control are enabled,
	//  and the Autonomous Database is using a different IP access control list
	//  (ACL) for standby compared to primary.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.are_primary_allowlisted_ips_used
	ArePrimaryAllowlistedIPsUsed *bool `json:"arePrimaryAllowlistedIPsUsed,omitempty"`

	// Output only. The details of the current lifestyle state of the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.lifecycle_details
	LifecycleDetails *string `json:"lifecycleDetails,omitempty"`

	// Output only. The current lifecycle state of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.state
	State *string `json:"state,omitempty"`

	// Output only. The Autonomous Container Database OCID.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.autonomous_container_database_id
	AutonomousContainerDatabaseID *string `json:"autonomousContainerDatabaseID,omitempty"`

	// Output only. The list of available Oracle Database upgrade versions for an
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.available_upgrade_versions
	AvailableUpgradeVersions []string `json:"availableUpgradeVersions,omitempty"`

	// Output only. The connection strings used to connect to an Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.connection_strings
	ConnectionStrings *AutonomousDatabaseConnectionStringsObservedState `json:"connectionStrings,omitempty"`

	// Output only. The Oracle Connection URLs for an Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.connection_urls
	ConnectionURLs *AutonomousDatabaseConnectionURLsObservedState `json:"connectionURLs,omitempty"`

	// Output only. This field indicates the number of seconds of data loss during
	//  a Data Guard failover.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.failed_data_recovery_duration
	FailedDataRecoveryDuration *string `json:"failedDataRecoveryDuration,omitempty"`

	// Output only. The memory assigned to in-memory tables in an Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.memory_table_gbs
	MemoryTableGBs *int32 `json:"memoryTableGBs,omitempty"`

	// Output only. Deprecated: Please use `local_data_guard_enabled` instead.
	//  This field indicates whether the Autonomous Database has local (in-region)
	//  Data Guard enabled.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.is_local_data_guard_enabled
	IsLocalDataGuardEnabled *bool `json:"isLocalDataGuardEnabled,omitempty"`

	// Output only. Deprecated: Please use
	//  `local_adg_auto_failover_max_data_loss_limit_duration` instead.
	//  This field indicates the maximum data loss limit for an Autonomous
	//  Database, in seconds.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.local_adg_auto_failover_max_data_loss_limit
	LocalAdgAutoFailoverMaxDataLossLimit *int32 `json:"localAdgAutoFailoverMaxDataLossLimit,omitempty"`

	// Output only. The details of the Autonomous Data Guard standby database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.local_standby_db
	LocalStandbyDb *AutonomousDatabaseStandbySummaryObservedState `json:"localStandbyDb,omitempty"`

	// Output only. The amount of memory enabled per ECPU, in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.memory_per_oracle_compute_unit_gbs
	MemoryPerOracleComputeUnitGBs *int32 `json:"memoryPerOracleComputeUnitGBs,omitempty"`

	// Output only. This field indicates the local disaster recovery (DR) type of
	//  an Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.local_disaster_recovery_type
	LocalDisasterRecoveryType *string `json:"localDisasterRecoveryType,omitempty"`

	// Output only. The current state of the Data Safe registration for the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.data_safe_state
	DataSafeState *string `json:"dataSafeState,omitempty"`

	// Output only. The current state of database management for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.database_management_state
	DatabaseManagementState *string `json:"databaseManagementState,omitempty"`

	// Output only. This field indicates the current mode of the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.open_mode
	OpenMode *string `json:"openMode,omitempty"`

	// Output only. This field indicates the state of Operations Insights for the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.operations_insights_state
	OperationsInsightsState *string `json:"operationsInsightsState,omitempty"`

	// Output only. The list of OCIDs of standby databases located in Autonomous
	//  Data Guard remote regions that are associated with the source database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.peer_db_ids
	PeerDbIDs []string `json:"peerDbIDs,omitempty"`

	// Output only. The permission level of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.permission_level
	PermissionLevel *string `json:"permissionLevel,omitempty"`

	// Output only. The private endpoint for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.private_endpoint
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`

	// Output only. The refresh mode of the cloned Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.refreshable_mode
	RefreshableMode *string `json:"refreshableMode,omitempty"`

	// Output only. The refresh State of the clone.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.refreshable_state
	RefreshableState *string `json:"refreshableState,omitempty"`

	// Output only. The Data Guard role of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.role
	Role *string `json:"role,omitempty"`

	// Output only. The list and details of the scheduled operations of the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.scheduled_operation_details
	ScheduledOperationDetails []ScheduledOperationDetailsObservedState `json:"scheduledOperationDetails,omitempty"`

	// Output only. The SQL Web Developer URL for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.sql_web_developer_url
	SQLWebDeveloperURL *string `json:"sqlWebDeveloperURL,omitempty"`

	// Output only. The list of available regions that can be used to create a
	//  clone for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.supported_clone_regions
	SupportedCloneRegions []string `json:"supportedCloneRegions,omitempty"`

	// Output only. The storage space used by Autonomous Database, in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.used_data_storage_size_tbs
	UsedDataStorageSizeTbs *int32 `json:"usedDataStorageSizeTbs,omitempty"`

	// Output only. The Oracle Cloud Infrastructure link for the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.oci_url
	OciURL *string `json:"ociURL,omitempty"`

	// Output only. The storage space used by automatic backups of Autonomous
	//  Database, in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.total_auto_backup_storage_size_gbs
	TotalAutoBackupStorageSizeGBs *float32 `json:"totalAutoBackupStorageSizeGBs,omitempty"`

	// Output only. The long term backup schedule of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.next_long_term_backup_time
	NextLongTermBackupTime *string `json:"nextLongTermBackupTime,omitempty"`

	// Output only. The date and time the Autonomous Data Guard role was changed
	//  for the standby Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.data_guard_role_changed_time
	DataGuardRoleChangedTime *string `json:"dataGuardRoleChangedTime,omitempty"`

	// Output only. The date and time the Disaster Recovery role was changed for
	//  the standby Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.disaster_recovery_role_changed_time
	DisasterRecoveryRoleChangedTime *string `json:"disasterRecoveryRoleChangedTime,omitempty"`

	// Output only. The date and time when maintenance will begin.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.maintenance_begin_time
	MaintenanceBeginTime *string `json:"maintenanceBeginTime,omitempty"`

	// Output only. The date and time when maintenance will end.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.maintenance_end_time
	MaintenanceEndTime *string `json:"maintenanceEndTime,omitempty"`

	// Output only. The history of the encryption keys used to encrypt the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.encryption_key_history_entries
	EncryptionKeyHistoryEntries []EncryptionKeyHistoryEntryObservedState `json:"encryptionKeyHistoryEntries,omitempty"`

	// Output only. An Oracle-managed Google Cloud service account on which
	//  customers can grant roles to access resources in the customer project.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.service_agent_email
	ServiceAgentEmail *string `json:"serviceAgentEmail,omitempty"`
}

// OracleDatabaseAutonomousDatabaseStatus defines the config connector machine state of OracleDatabaseAutonomousDatabase
type OracleDatabaseAutonomousDatabaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the OracleDatabaseAutonomousDatabase resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *OracleDatabaseAutonomousDatabaseObservedState `json:"observedState,omitempty"`
}

// OracleDatabaseAutonomousDatabaseObservedState is the state of the OracleDatabaseAutonomousDatabase resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.AutonomousDatabase
type OracleDatabaseAutonomousDatabaseObservedState struct {
	// Output only. The ID of the subscription entitlement associated with the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.entitlement_id
	EntitlementID *string `json:"entitlementID,omitempty"`

	// Optional. The properties of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.properties
	Properties *AutonomousDatabasePropertiesObservedState `json:"properties,omitempty"`

	// Output only. The peer Autonomous Database names of the given Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.peer_autonomous_databases
	PeerAutonomousDatabases []string `json:"peerAutonomousDatabases,omitempty"`

	// Output only. The date and time that the Autonomous Database was created.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. List of supported GCP region to clone the Autonomous Database
	//  for disaster recovery. Format: `project/{project}/locations/{location}`.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.disaster_recovery_supported_locations
	DisasterRecoverySupportedLocations []string `json:"disasterRecoverySupportedLocations,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcporacledatabaseautonomousdatabase;gcporacledatabaseautonomousdatabases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OracleDatabaseAutonomousDatabase is the Schema for the OracleDatabaseAutonomousDatabase API
// +k8s:openapi-gen=true
type OracleDatabaseAutonomousDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OracleDatabaseAutonomousDatabaseSpec   `json:"spec,omitempty"`
	Status OracleDatabaseAutonomousDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OracleDatabaseAutonomousDatabaseList contains a list of OracleDatabaseAutonomousDatabase
type OracleDatabaseAutonomousDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OracleDatabaseAutonomousDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OracleDatabaseAutonomousDatabase{}, &OracleDatabaseAutonomousDatabaseList{})
}
