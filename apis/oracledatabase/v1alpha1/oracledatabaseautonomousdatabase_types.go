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
	AdminPassword *string `json:"adminPassword,omitempty"`

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
