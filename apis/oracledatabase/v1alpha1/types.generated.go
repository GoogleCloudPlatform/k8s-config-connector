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


// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup
type AutonomousDatabaseBackup struct {
	// Identifier. The name of the Autonomous Database Backup resource with the
	//  format:
	//  projects/{project}/locations/{region}/autonomousDatabaseBackups/{autonomous_database_backup}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the Autonomous Database resource for which the backup
	//  is being created. Format:
	//  projects/{project}/locations/{region}/autonomousDatabases/{autonomous_database}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup.autonomous_database
	AutonomousDatabase *string `json:"autonomousDatabase,omitempty"`

	// Optional. User friendly name for the Backup. The name does not have to be
	//  unique.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Various properties of the backup.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup.properties
	Properties *AutonomousDatabaseBackupProperties `json:"properties,omitempty"`

	// Optional. labels or tags associated with the resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties
type AutonomousDatabaseBackupProperties struct {

	// Optional. Retention period in days for the backup.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.retention_period_days
	RetentionPeriodDays *int32 `json:"retentionPeriodDays,omitempty"`

	// Optional. The OCID of the key store of Oracle Vault.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.key_store_id
	KeyStoreID *string `json:"keyStoreID,omitempty"`

	// Optional. The wallet name for Oracle Key Vault.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.key_store_wallet
	KeyStoreWallet *string `json:"keyStoreWallet,omitempty"`

	// Optional. The OCID of the key container that is used as the master
	//  encryption key in database transparent data encryption (TDE) operations.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.kms_key_id
	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	// Optional. The OCID of the key container version that is used in database
	//  transparent data encryption (TDE) operations KMS Key can have multiple key
	//  versions. If none is specified, the current key version (latest) of the Key
	//  Id is used for the operation. Autonomous Database Serverless does not use
	//  key versions, hence is not applicable for Autonomous Database Serverless
	//  instances.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.kms_key_version_id
	KMSKeyVersionID *string `json:"kmsKeyVersionID,omitempty"`

	// Optional. The OCID of the vault.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.vault_id
	VaultID *string `json:"vaultID,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup
type AutonomousDatabaseBackupObservedState struct {
	// Optional. Various properties of the backup.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackup.properties
	Properties *AutonomousDatabaseBackupPropertiesObservedState `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties
type AutonomousDatabaseBackupPropertiesObservedState struct {
	// Output only. OCID of the Autonomous Database backup.
	//  https://docs.oracle.com/en-us/iaas/Content/General/Concepts/identifiers.htm#Oracle
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.ocid
	Ocid *string `json:"ocid,omitempty"`

	// Output only. The OCID of the compartment.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.compartment_id
	CompartmentID *string `json:"compartmentID,omitempty"`

	// Output only. The quantity of data in the database, in terabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.database_size_tb
	DatabaseSizeTb *float32 `json:"databaseSizeTb,omitempty"`

	// Output only. A valid Oracle Database version for Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.db_version
	DbVersion *string `json:"dbVersion,omitempty"`

	// Output only. Indicates if the backup is long term backup.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.is_long_term_backup
	IsLongTermBackup *bool `json:"isLongTermBackup,omitempty"`

	// Output only. Indicates if the backup is automatic or user initiated.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.is_automatic_backup
	IsAutomaticBackup *bool `json:"isAutomaticBackup,omitempty"`

	// Output only. Indicates if the backup can be used to restore the Autonomous
	//  Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.is_restorable
	IsRestorable *bool `json:"isRestorable,omitempty"`

	// Output only. Additional information about the current lifecycle state.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.lifecycle_details
	LifecycleDetails *string `json:"lifecycleDetails,omitempty"`

	// Output only. The lifecycle state of the backup.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.lifecycle_state
	LifecycleState *string `json:"lifecycleState,omitempty"`

	// Output only. The backup size in terabytes.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.size_tb
	SizeTb *float32 `json:"sizeTb,omitempty"`

	// Output only. Timestamp until when the backup will be available.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.available_till_time
	AvailableTillTime *string `json:"availableTillTime,omitempty"`

	// Output only. The date and time the backup completed.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The date and time the backup started.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The type of the backup.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseBackupProperties.type
	Type *string `json:"type,omitempty"`
}
