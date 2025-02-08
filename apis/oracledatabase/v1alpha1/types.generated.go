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


// +kcc:proto=google.cloud.oracledatabase.v1.AllConnectionStrings
type AllConnectionStrings struct {
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabase
type AutonomousDatabase struct {
	// Identifier. The name of the Autonomous Database resource in the following
	//  format:
	//  projects/{project}/locations/{region}/autonomousDatabases/{autonomous_database}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.name
	Name *string `json:"name,omitempty"`

	// Optional. The name of the Autonomous Database. The database name must be
	//  unique in the project. The name must begin with a letter and can contain a
	//  maximum of 30 alphanumeric characters.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.database
	Database *string `json:"database,omitempty"`

	// Optional. The display name for the Autonomous Database. The name does not
	//  have to be unique within your project.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The password for the default ADMIN user.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.admin_password
	AdminPassword *string `json:"adminPassword,omitempty"`

	// Optional. The properties of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.properties
	Properties *AutonomousDatabaseProperties `json:"properties,omitempty"`

	// Optional. The labels or tags associated with the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The name of the VPC network used by the Autonomous Database in
	//  the following format: projects/{project}/global/networks/{network}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.network
	Network *string `json:"network,omitempty"`

	// Required. The subnet CIDR range for the Autonmous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.cidr
	Cidr *string `json:"cidr,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseApex
type AutonomousDatabaseApex struct {
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings
type AutonomousDatabaseConnectionStrings struct {
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls
type AutonomousDatabaseConnectionUrls struct {
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties
type AutonomousDatabaseProperties struct {

	// Optional. The number of compute servers for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.compute_count
	ComputeCount *float32 `json:"computeCount,omitempty"`

	// Optional. The number of CPU cores to be made available to the database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.cpu_core_count
	CpuCoreCount *int32 `json:"cpuCoreCount,omitempty"`

	// Optional. The size of the data stored in the database, in terabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.data_storage_size_tb
	DataStorageSizeTb *int32 `json:"dataStorageSizeTb,omitempty"`

	// Optional. The size of the data stored in the database, in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.data_storage_size_gb
	DataStorageSizeGB *int32 `json:"dataStorageSizeGB,omitempty"`

	// Required. The workload type of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.db_workload
	DbWorkload *string `json:"dbWorkload,omitempty"`

	// Optional. The edition of the Autonomous Databases.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.db_edition
	DbEdition *string `json:"dbEdition,omitempty"`

	// Optional. The character set for the Autonomous Database. The default is
	//  AL32UTF8.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.character_set
	CharacterSet *string `json:"characterSet,omitempty"`

	// Optional. The national character set for the Autonomous Database. The
	//  default is AL16UTF16.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.n_character_set
	NCharacterSet *string `json:"nCharacterSet,omitempty"`

	// Optional. The private endpoint IP address for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.private_endpoint_ip
	PrivateEndpointIP *string `json:"privateEndpointIP,omitempty"`

	// Optional. The private endpoint label for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.private_endpoint_label
	PrivateEndpointLabel *string `json:"privateEndpointLabel,omitempty"`

	// Optional. The Oracle Database version for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.db_version
	DbVersion *string `json:"dbVersion,omitempty"`

	// Optional. This field indicates if auto scaling is enabled for the
	//  Autonomous Database CPU core count.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.is_auto_scaling_enabled
	IsAutoScalingEnabled *bool `json:"isAutoScalingEnabled,omitempty"`

	// Optional. This field indicates if auto scaling is enabled for the
	//  Autonomous Database storage.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.is_storage_auto_scaling_enabled
	IsStorageAutoScalingEnabled *bool `json:"isStorageAutoScalingEnabled,omitempty"`

	// Required. The license type used for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.license_type
	LicenseType *string `json:"licenseType,omitempty"`

	// Optional. The list of customer contacts.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.customer_contacts
	CustomerContacts []CustomerContact `json:"customerContacts,omitempty"`

	// Optional. The ID of the Oracle Cloud Infrastructure vault secret.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.secret_id
	SecretID *string `json:"secretID,omitempty"`

	// Optional. The ID of the Oracle Cloud Infrastructure vault.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.vault_id
	VaultID *string `json:"vaultID,omitempty"`

	// Optional. The maintenance schedule of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.maintenance_schedule_type
	MaintenanceScheduleType *string `json:"maintenanceScheduleType,omitempty"`

	// Optional. This field specifies if the Autonomous Database requires mTLS
	//  connections.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.mtls_connection_required
	MtlsConnectionRequired *bool `json:"mtlsConnectionRequired,omitempty"`

	// Optional. The retention period for the Autonomous Database. This field is
	//  specified in days, can range from 1 day to 60 days, and has a default value
	//  of 60 days.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.backup_retention_period_days
	BackupRetentionPeriodDays *int32 `json:"backupRetentionPeriodDays,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseStandbySummary
type AutonomousDatabaseStandbySummary struct {
}

// +kcc:proto=google.cloud.oracledatabase.v1.CustomerContact
type CustomerContact struct {
	// Required. The email address used by Oracle to send notifications regarding
	//  databases and infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CustomerContact.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile
type DatabaseConnectionStringProfile struct {
}

// +kcc:proto=google.cloud.oracledatabase.v1.ScheduledOperationDetails
type ScheduledOperationDetails struct {
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

// +kcc:proto=google.cloud.oracledatabase.v1.AllConnectionStrings
type AllConnectionStringsObservedState struct {
	// Output only. The database service provides the highest level of resources
	//  to each SQL statement.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AllConnectionStrings.high
	High *string `json:"high,omitempty"`

	// Output only. The database service provides the least level of resources to
	//  each SQL statement.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AllConnectionStrings.low
	Low *string `json:"low,omitempty"`

	// Output only. The database service provides a lower level of resources to
	//  each SQL statement.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AllConnectionStrings.medium
	Medium *string `json:"medium,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabase
type AutonomousDatabaseObservedState struct {
	// Output only. The ID of the subscription entitlement associated with the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.entitlement_id
	EntitlementID *string `json:"entitlementID,omitempty"`

	// Optional. The properties of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.properties
	Properties *AutonomousDatabasePropertiesObservedState `json:"properties,omitempty"`

	// Output only. The date and time that the Autonomous Database was created.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseApex
type AutonomousDatabaseApexObservedState struct {
	// Output only. The Oracle APEX Application Development version.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseApex.apex_version
	ApexVersion *string `json:"apexVersion,omitempty"`

	// Output only. The Oracle REST Data Services (ORDS) version.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseApex.ords_version
	OrdsVersion *string `json:"ordsVersion,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings
type AutonomousDatabaseConnectionStringsObservedState struct {
	// Output only. Returns all connection strings that can be used to connect to
	//  the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings.all_connection_strings
	AllConnectionStrings *AllConnectionStrings `json:"allConnectionStrings,omitempty"`

	// Output only. The database service provides the least level of resources to
	//  each SQL statement, but supports the most number of concurrent SQL
	//  statements.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings.dedicated
	Dedicated *string `json:"dedicated,omitempty"`

	// Output only. The database service provides the highest level of resources
	//  to each SQL statement.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings.high
	High *string `json:"high,omitempty"`

	// Output only. The database service provides the least level of resources to
	//  each SQL statement.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings.low
	Low *string `json:"low,omitempty"`

	// Output only. The database service provides a lower level of resources to
	//  each SQL statement.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings.medium
	Medium *string `json:"medium,omitempty"`

	// Output only. A list of connection string profiles to allow clients to
	//  group, filter, and select values based on the structured metadata.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionStrings.profiles
	Profiles []DatabaseConnectionStringProfile `json:"profiles,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseConnectionUrls
type AutonomousDatabaseConnectionUrlsObservedState struct {
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
	SqlDevWebURI *string `json:"sqlDevWebURI,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties
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
	ApexDetails *AutonomousDatabaseApex `json:"apexDetails,omitempty"`

	// Output only. This field indicates the status of Data Guard and Access
	//  control for the Autonomous Database. The field's value is null if Data
	//  Guard is disabled or Access Control is disabled. The field's value is TRUE
	//  if both Data Guard and Access Control are enabled, and the Autonomous
	//  Database is using primary IP access control list (ACL) for standby. The
	//  field's value is FALSE if both Data Guard and Access Control are enabled,
	//  and the Autonomous Database is using a different IP access control list
	//  (ACL) for standby compared to primary.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.are_primary_allowlisted_ips_used
	ArePrimaryAllowlistedIpsUsed *bool `json:"arePrimaryAllowlistedIpsUsed,omitempty"`

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
	ConnectionStrings *AutonomousDatabaseConnectionStrings `json:"connectionStrings,omitempty"`

	// Output only. The Oracle Connection URLs for an Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.connection_urls
	ConnectionUrls *AutonomousDatabaseConnectionUrls `json:"connectionUrls,omitempty"`

	// Output only. This field indicates the number of seconds of data loss during
	//  a Data Guard failover.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.failed_data_recovery_duration
	FailedDataRecoveryDuration *string `json:"failedDataRecoveryDuration,omitempty"`

	// Output only. The memory assigned to in-memory tables in an Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.memory_table_gbs
	MemoryTableGbs *int32 `json:"memoryTableGbs,omitempty"`

	// Output only. This field indicates whether the Autonomous Database has local
	//  (in-region) Data Guard enabled.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.is_local_data_guard_enabled
	IsLocalDataGuardEnabled *bool `json:"isLocalDataGuardEnabled,omitempty"`

	// Output only. This field indicates the maximum data loss limit for an
	//  Autonomous Database, in seconds.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.local_adg_auto_failover_max_data_loss_limit
	LocalAdgAutoFailoverMaxDataLossLimit *int32 `json:"localAdgAutoFailoverMaxDataLossLimit,omitempty"`

	// Output only. The details of the Autonomous Data Guard standby database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.local_standby_db
	LocalStandbyDb *AutonomousDatabaseStandbySummary `json:"localStandbyDb,omitempty"`

	// Output only. The amount of memory enabled per ECPU, in gigabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.memory_per_oracle_compute_unit_gbs
	MemoryPerOracleComputeUnitGbs *int32 `json:"memoryPerOracleComputeUnitGbs,omitempty"`

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
	PeerDbIds []string `json:"peerDbIds,omitempty"`

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
	ScheduledOperationDetails []ScheduledOperationDetails `json:"scheduledOperationDetails,omitempty"`

	// Output only. The SQL Web Developer URL for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.sql_web_developer_url
	SqlWebDeveloperURL *string `json:"sqlWebDeveloperURL,omitempty"`

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
	TotalAutoBackupStorageSizeGbs *float32 `json:"totalAutoBackupStorageSizeGbs,omitempty"`

	// Output only. The long term backup schedule of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.next_long_term_backup_time
	NextLongTermBackupTime *string `json:"nextLongTermBackupTime,omitempty"`

	// Output only. The date and time when maintenance will begin.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.maintenance_begin_time
	MaintenanceBeginTime *string `json:"maintenanceBeginTime,omitempty"`

	// Output only. The date and time when maintenance will end.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseProperties.maintenance_end_time
	MaintenanceEndTime *string `json:"maintenanceEndTime,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseStandbySummary
type AutonomousDatabaseStandbySummaryObservedState struct {
	// Output only. The amount of time, in seconds, that the data of the standby
	//  database lags in comparison to the data of the primary database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseStandbySummary.lag_time_duration
	LagTimeDuration *string `json:"lagTimeDuration,omitempty"`

	// Output only. The additional details about the current lifecycle state of
	//  the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseStandbySummary.lifecycle_details
	LifecycleDetails *string `json:"lifecycleDetails,omitempty"`

	// Output only. The current lifecycle state of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseStandbySummary.state
	State *string `json:"state,omitempty"`

	// Output only. The date and time the Autonomous Data Guard role was switched
	//  for the standby Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseStandbySummary.data_guard_role_changed_time
	DataGuardRoleChangedTime *string `json:"dataGuardRoleChangedTime,omitempty"`

	// Output only. The date and time the Disaster Recovery role was switched for
	//  the standby Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseStandbySummary.disaster_recovery_role_changed_time
	DisasterRecoveryRoleChangedTime *string `json:"disasterRecoveryRoleChangedTime,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile
type DatabaseConnectionStringProfileObservedState struct {
	// Output only. The current consumer group being used by the connection.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.consumer_group
	ConsumerGroup *string `json:"consumerGroup,omitempty"`

	// Output only. The display name for the database connection.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The host name format being currently used in connection
	//  string.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.host_format
	HostFormat *string `json:"hostFormat,omitempty"`

	// Output only. This field indicates if the connection string is regional and
	//  is only applicable for cross-region Data Guard.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.is_regional
	IsRegional *bool `json:"isRegional,omitempty"`

	// Output only. The protocol being used by the connection.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Output only. The current session mode of the connection.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.session_mode
	SessionMode *string `json:"sessionMode,omitempty"`

	// Output only. The syntax of the connection string.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.syntax_format
	SyntaxFormat *string `json:"syntaxFormat,omitempty"`

	// Output only. This field indicates the TLS authentication type of the
	//  connection.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.tls_authentication
	TlsAuthentication *string `json:"tlsAuthentication,omitempty"`

	// Output only. The value of the connection string.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.DatabaseConnectionStringProfile.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.ScheduledOperationDetails
type ScheduledOperationDetailsObservedState struct {
	// Output only. Day of week.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ScheduledOperationDetails.day_of_week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// Output only. Auto start time.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ScheduledOperationDetails.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Output only. Auto stop time.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ScheduledOperationDetails.stop_time
	StopTime *TimeOfDay `json:"stopTime,omitempty"`
}
