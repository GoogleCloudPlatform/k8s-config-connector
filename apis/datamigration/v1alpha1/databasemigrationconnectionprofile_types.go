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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatabaseMigrationConnectionProfileGVK = GroupVersion.WithKind("DatabaseMigrationConnectionProfile")

// +kcc:proto=google.cloud.clouddms.v1.PrivateConnectivity
type PrivateConnectivity struct {
	// The DatabaseMigrationPrivateConnection that this resource peers with.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnectivity.private_connection
	PrivateConnectionRef *DatabaseMigrationPrivateConnectionRef `json:"privateConnectionRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.CloudSqlSettings
type CloudSQLSettings struct {
	// The database engine type and version.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	// The resource labels for a Cloud SQL instance to use to annotate any related
	//  underlying resources such as Compute Engine VMs.
	//  An object containing a list of "key": "value" pairs.
	//
	//  Example: `{ "name": "wrench", "mass": "18kg", "count": "3" }`.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.user_labels
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// The tier (or machine type) for this instance, for example:
	//  `db-n1-standard-1` (MySQL instances) or
	//  `db-custom-1-3840` (PostgreSQL instances).
	//  For more information, see
	//  [Cloud SQL Instance
	//  Settings](https://cloud.google.com/sql/docs/mysql/instance-settings).
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.tier
	Tier *string `json:"tier,omitempty"`

	// The maximum size to which storage capacity can be automatically increased.
	//  The default value is 0, which specifies that there is no limit.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.storage_auto_resize_limit
	StorageAutoResizeLimit *int64 `json:"storageAutoResizeLimit,omitempty"`

	// The activation policy specifies when the instance is activated; it is
	//  applicable only when the instance state is 'RUNNABLE'. Valid values:
	//
	//  'ALWAYS': The instance is on, and remains so even in
	//  the absence of connection requests.
	//
	//  `NEVER`: The instance is off; it is not activated, even if a
	//  connection request arrives.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.activation_policy
	ActivationPolicy *string `json:"activationPolicy,omitempty"`

	// IP Config settings for the Cloud SQL instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.ip_config
	IPConfig *SQLIPConfig `json:"ipConfig,omitempty"`

	// Storage auto resize enabled/disabled.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.auto_storage_increase
	AutoStorageIncrease *bool `json:"autoStorageIncrease,omitempty"`

	// Database flags of the Cloud SQL instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.database_flags
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty"`

	// The type of disk for data storage.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.data_disk_type
	DataDiskType *string `json:"dataDiskType,omitempty"`

	// The size of data disk in GB.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.data_disk_size_gb
	DataDiskSizeGB *int64 `json:"dataDiskSizeGB,omitempty"`

	// The primary zone of the Cloud SQL instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.zone
	Zone *string `json:"zone,omitempty"`

	// The secondary zone of the Cloud SQL instance (applicable if highly available).
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.secondary_zone
	SecondaryZone *string `json:"secondaryZone,omitempty"`

	// The source connection profile ID.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.source_id
	SourceID *string `json:"sourceID,omitempty"`

	// Input only. Initial root password.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password
	RootPassword *string `json:"rootPassword,omitempty"`

	// The Cloud SQL default instance level collation.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.collation
	Collation *string `json:"collation,omitempty"`

	// The KMS key name used for the csql instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.cmek_key_name
	CmekKeyNameRef *refsv1beta1.KMSCryptoKeyRef `json:"cmekKeyNameRef,omitempty"`

	// Optional. Availability type. Potential values:
	//  *  `ZONAL`: The instance serves data from only one zone. Outages in that
	//  zone affect data availability.
	//  *  `REGIONAL`: The instance can serve data from more than one zone in a
	//  region (it is highly available).
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.availability_type
	AvailabilityType *string `json:"availabilityType,omitempty"`

	// Optional. The edition of the given Cloud SQL instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.edition
	Edition *string `json:"edition,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword
type AlloyDbSettings_UserPassword struct {
	// The database username.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.user
	User *string `json:"user,omitempty"`

	// The initial password for the user.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.password
	Password *string `json:"password,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.MachineConfig
type AlloyDbSettings_PrimaryInstanceSettings_MachineConfig struct {
	// The number of CPU's in the VM instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.MachineConfig.cpu_count
	CPUCount *int32 `json:"cpuCount,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings
type AlloyDbSettings_PrimaryInstanceSettings struct {
	// Required. The ID of the AlloyDB primary instance. The ID must satisfy the
	//  regex expression "[a-z0-9-]+".
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.id
	ID *string `json:"id,omitempty"`

	// Machine config for primary instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.machine_config
	MachineConfig *AlloyDbSettings_PrimaryInstanceSettings_MachineConfig `json:"machineConfig,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.EncryptionConfig
type AlloyDbSettings_EncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//  Each Cloud KMS key is regionalized and has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.EncryptionConfig.kms_key_name
	KmsKeyNameRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyNameRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings
type AlloyDbSettings struct {
	// Required. Input only. Initial user to setup during cluster creation.
	//  Required.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.initial_user
	InitialUser *AlloyDbSettings_UserPassword `json:"initialUser,omitempty"`

	// Required. The resource link for the VPC network in which cluster resources
	//  are created and from which they are accessible via Private IP. The network
	//  must belong to the same project as the cluster. It is specified in the
	//  form: "projects/{project_number}/global/networks/{network_id}". This is
	//  required to create a cluster.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.vpc_network
	VpcNetworkRef *computev1beta1.ComputeNetworkRef `json:"vpcNetworkRef,omitempty"`

	// Labels for the AlloyDB cluster created by DMS. An object containing a list
	//  of 'key', 'value' pairs.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.labels
	Labels map[string]string `json:"labels,omitempty"`

	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.primary_instance_settings
	PrimaryInstanceSettings *AlloyDbSettings_PrimaryInstanceSettings `json:"primaryInstanceSettings,omitempty"`

	// Optional. The encryption config can be specified to encrypt the data disks
	//  and other persistent data resources of a cluster with a
	//  customer-managed encryption key (CMEK). When this field is not
	//  specified, the cluster will then use default encryption scheme to
	//  protect the user data.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.encryption_config
	EncryptionConfig *AlloyDbSettings_EncryptionConfig `json:"encryptionConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbSettings
type AlloyDbSettingsObservedState struct {
	// Required. Input only. Initial user to setup during cluster creation.
	//  Required.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.initial_user
	InitialUser *AlloyDbSettings_UserPasswordObservedState `json:"initialUser,omitempty"`

	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.primary_instance_settings
	PrimaryInstanceSettings *AlloyDbSettings_PrimaryInstanceSettingsObservedState `json:"primaryInstanceSettings,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword
type AlloyDbSettings_UserPasswordObservedState struct {
	// Output only. Indicates if the initial_user.password field has been set.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.password_set
	PasswordSet *bool `json:"passwordSet,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings
type AlloyDbSettings_PrimaryInstanceSettingsObservedState struct {
	// Output only. The private IP address for the Instance.
	//  This is the connection endpoint for an end-user application.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.private_ip
	PrivateIP *string `json:"privateIP,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.CloudSqlSettings
type CloudSQLSettingsObservedState struct {
	// Output only. Indicates If this connection profile root password is stored.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password_set
	RootPasswordSet *bool `json:"rootPasswordSet,omitempty"`
}

// DatabaseMigrationConnectionProfileSpec defines the desired state of DatabaseMigrationConnectionProfile
// +kcc:spec:proto=google.cloud.clouddms.v1.ConnectionProfile
type DatabaseMigrationConnectionProfileSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The DatabaseMigrationConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The resource labels for connection profile to use to annotate any related
	//  underlying resources such as Compute Engine VMs. An object containing a
	//  list of "key": "value" pairs.
	//
	//  Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The connection profile display name.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A MySQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.mysql
	Mysql *MySQLConnectionProfile `json:"mysql,omitempty"`

	// A PostgreSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.postgresql
	Postgresql *PostgreSQLConnectionProfile `json:"postgresql,omitempty"`

	// An Oracle database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.oracle
	Oracle *OracleConnectionProfile `json:"oracle,omitempty"`

	// A CloudSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.cloudsql
	Cloudsql *CloudSQLConnectionProfile `json:"cloudsql,omitempty"`

	// An AlloyDB cluster connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.alloydb
	Alloydb *AlloyDbConnectionProfile `json:"alloydb,omitempty"`

	// The database provider.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.provider
	Provider *string `json:"provider,omitempty"`
}

// DatabaseMigrationConnectionProfileStatus defines the config connector machine state of DatabaseMigrationConnectionProfile
type DatabaseMigrationConnectionProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatabaseMigrationConnectionProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatabaseMigrationConnectionProfileObservedState `json:"observedState,omitempty"`
}

// DatabaseMigrationConnectionProfileObservedState is the state of the DatabaseMigrationConnectionProfile resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.ConnectionProfile
type DatabaseMigrationConnectionProfileObservedState struct {
	// Output only. The timestamp when the resource was created.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The current connection profile state (e.g. DRAFT, READY, or FAILED).
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.state
	State *string `json:"state,omitempty"`

	// A MySQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.mysql
	Mysql *MySQLConnectionProfileObservedState `json:"mysql,omitempty"`

	// A PostgreSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.postgresql
	Postgresql *PostgreSQLConnectionProfileObservedState `json:"postgresql,omitempty"`

	// An Oracle database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.oracle
	Oracle *OracleConnectionProfileObservedState `json:"oracle,omitempty"`

	// A CloudSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.cloudsql
	Cloudsql *CloudSQLConnectionProfileObservedState `json:"cloudsql,omitempty"`

	// An AlloyDB cluster connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.alloydb
	Alloydb *AlloyDbConnectionProfileObservedState `json:"alloydb,omitempty"`

	// Output only. The error details in case of state FAILED.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.error
	Error *common.Status `json:"error,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatabasemigrationconnectionprofile;gcpdatabasemigrationconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatabaseMigrationConnectionProfile is the Schema for the DatabaseMigrationConnectionProfile API
// +k8s:openapi-gen=true
type DatabaseMigrationConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatabaseMigrationConnectionProfileSpec   `json:"spec,omitempty"`
	Status DatabaseMigrationConnectionProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatabaseMigrationConnectionProfileList contains a list of DatabaseMigrationConnectionProfile
type DatabaseMigrationConnectionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigrationConnectionProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseMigrationConnectionProfile{}, &DatabaseMigrationConnectionProfileList{})
}
