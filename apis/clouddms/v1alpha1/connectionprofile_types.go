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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSConnectionProfileGVK = GroupVersion.WithKind("CloudDMSConnectionProfile")

// CloudDMSConnectionProfileSpec defines the desired state of CloudDMSConnectionProfile
// +kcc:spec:proto=google.cloud.clouddms.v1.ConnectionProfile
type CloudDMSConnectionProfileSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The CloudDMSConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The connection profile display name.
	// +kcc:proto=display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A MySQL database connection profile.
	// +kcc:proto=mysql
	Mysql *MySQLConnectionProfile `json:"mysql,omitempty"`

	// A PostgreSQL database connection profile.
	// +kcc:proto=postgresql
	Postgresql *PostgreSQLConnectionProfile `json:"postgresql,omitempty"`

	// An Oracle database connection profile.
	// +kcc:proto=oracle
	Oracle *OracleConnectionProfile `json:"oracle,omitempty"`

	// A CloudSQL database connection profile.
	// +kcc:proto=cloudsql
	Cloudsql *CloudSQLConnectionProfile `json:"cloudsql,omitempty"`

	// An AlloyDB cluster connection profile.
	// +kcc:proto=alloydb
	Alloydb *AlloyDbConnectionProfile `json:"alloydb,omitempty"`

	// The database provider.
	// +kcc:proto=provider
	Provider *string `json:"provider,omitempty"`
}

// CloudDMSConnectionProfileStatus defines the config connector machine state of CloudDMSConnectionProfile
type CloudDMSConnectionProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDMSConnectionProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDMSConnectionProfileObservedState `json:"observedState,omitempty"`
}

// CloudDMSConnectionProfileObservedState is the state of the CloudDMSConnectionProfile resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.ConnectionProfile
type CloudDMSConnectionProfileObservedState struct {
	// The name of this connection profile resource in the form of
	//  projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
	// +kcc:proto=name
	Name *string `json:"name,omitempty"`

	// The current connection profile state (e.g. DRAFT, READY, or FAILED).
	// +kcc:proto=state
	State *string `json:"state,omitempty"`

	// Output only. The timestamp when the resource was created.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto=create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto=update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The error details in case of state FAILED.
	// +kcc:proto=error
	Error *Status `json:"error,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsconnectionprofile;gcpclouddmsconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDMSConnectionProfile is the Schema for the CloudDMSConnectionProfile API
// +k8s:openapi-gen=true
type CloudDMSConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDMSConnectionProfileSpec   `json:"spec,omitempty"`
	Status CloudDMSConnectionProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDMSConnectionProfileList contains a list of CloudDMSConnectionProfile
type CloudDMSConnectionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDMSConnectionProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDMSConnectionProfile{}, &CloudDMSConnectionProfileList{})
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

	// The settings for IP Management. This allows to enable or disable the
	//  instance IP and manage which external networks can connect to the instance.
	//  The IPv4 address cannot be disabled.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.ip_config
	IPConfig *SQLIPConfig `json:"ipConfig,omitempty"`

	// [default: ON] If you enable this setting, Cloud SQL checks your available
	//  storage every 30 seconds. If the available storage falls below a threshold
	//  size, Cloud SQL automatically adds additional storage capacity. If the
	//  available storage repeatedly falls below the threshold size, Cloud SQL
	//  continues to add storage until it reaches the maximum of 30 TB.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.auto_storage_increase
	AutoStorageIncrease *bool `json:"autoStorageIncrease,omitempty"`

	// The database flags passed to the Cloud SQL instance at startup.
	//  An object containing a list of "key": value pairs.
	//  Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.database_flags
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty"`

	// The type of storage: `PD_SSD` (default) or `PD_HDD`.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.data_disk_type
	DataDiskType *string `json:"dataDiskType,omitempty"`

	// The storage capacity available to the database, in GB.
	//  The minimum (and default) size is 10GB.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.data_disk_size_gb
	DataDiskSizeGB *int64 `json:"dataDiskSizeGB,omitempty"`

	// The Google Cloud Platform zone where your Cloud SQL database instance is
	//  located.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.zone
	Zone *string `json:"zone,omitempty"`

	// Optional. The Google Cloud Platform zone where the failover Cloud SQL
	//  database instance is located. Used when the Cloud SQL database availability
	//  type is REGIONAL (i.e. multiple zones / highly available).
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.secondary_zone
	SecondaryZone *string `json:"secondaryZone,omitempty"`

	// The Database Migration Service source connection profile ID,
	//  in the format:
	//  `projects/my_project_name/locations/us-central1/connectionProfiles/connection_profile_ID`
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.source_id
	SourceID *string `json:"sourceID,omitempty"`

	// NOTYET
	// Input only. Initial root password.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password
	// RootPassword *string `json:"rootPassword,omitempty"`

	// The Cloud SQL default instance level collation.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.collation
	Collation *string `json:"collation,omitempty"`

	// The KMS key name used for the csql instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.cmek_key_name
	CmekKeyName *string `json:"cmekKeyName,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity
type ForwardSSHTunnelConnectivity struct {
	// Required. Hostname for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Required. Username for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.username
	Username *string `json:"username,omitempty"`

	// Port for the SSH tunnel, default value is 22.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.port
	Port *int32 `json:"port,omitempty"`

	// NOTYET
	// Input only. SSH password.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.password
	// Password *string `json:"password,omitempty"`

	// Input only. SSH private key.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.private_key
	PrivateKey *string `json:"privateKey,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MySqlConnectionProfile
type MySQLConnectionProfile struct {
	// Required. The IP or hostname of the source MySQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.host
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source MySQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. The username that Database Migration Service will use to connect
	//  to the database. The value is encrypted when stored in Database Migration
	//  Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.username
	Username *string `json:"username,omitempty"`

	// NOTYET
	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.password
	// Password *string `json:"password,omitempty"`

	// SSL configuration for the destination to connect to the source database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.ssl
	SSL *SSLConfig `json:"ssl,omitempty"`

	// If the source is a Cloud SQL database, use this field to
	//  provide the Cloud SQL instance ID of the source.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.cloud_sql_id
	CloudSQLID *string `json:"cloudSQLID,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.OracleConnectionProfile
type OracleConnectionProfile struct {
	// Required. The IP or hostname of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.host
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. The username that Database Migration Service will use to connect
	//  to the database. The value is encrypted when stored in Database Migration
	//  Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.username
	Username *string `json:"username,omitempty"`

	// NOTYET
	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.password
	// Password *string `json:"password,omitempty"`

	// Required. Database service for the Oracle connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.database_service
	DatabaseService *string `json:"databaseService,omitempty"`

	// SSL configuration for the connection to the source Oracle database.
	//
	//   * Only `SERVER_ONLY` configuration is supported for Oracle SSL.
	//   * SSL is supported for Oracle versions 12 and above.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.ssl
	SSL *SSLConfig `json:"ssl,omitempty"`

	// Static Service IP connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.static_service_ip_connectivity
	StaticServiceIPConnectivity *StaticServiceIPConnectivity `json:"staticServiceIPConnectivity,omitempty"`

	// Forward SSH tunnel connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.forward_ssh_connectivity
	ForwardSSHConnectivity *ForwardSSHTunnelConnectivity `json:"forwardSSHConnectivity,omitempty"`

	// Private connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.private_connectivity
	PrivateConnectivity *PrivateConnectivity `json:"privateConnectivity,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.PostgreSqlConnectionProfile
type PostgreSQLConnectionProfile struct {
	// Required. The IP or hostname of the source PostgreSQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.host
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source PostgreSQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. The username that Database Migration Service will use to connect
	//  to the database. The value is encrypted when stored in Database Migration
	//  Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.username
	Username *string `json:"username,omitempty"`

	// NOTYET
	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.password
	// Password *string `json:"password,omitempty"`

	// SSL configuration for the destination to connect to the source database.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.ssl
	SSL *SSLConfig `json:"ssl,omitempty"`

	// If the source is a Cloud SQL database, use this field to
	//  provide the Cloud SQL instance ID of the source.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.cloud_sql_id
	CloudSQLID *string `json:"cloudSQLID,omitempty"`

	// Static ip connectivity data (default, no additional details needed).
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.static_ip_connectivity
	StaticIPConnectivity *StaticIPConnectivity `json:"staticIPConnectivity,omitempty"`

	// Private service connect connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.private_service_connect_connectivity
	PrivateServiceConnectConnectivity *PrivateServiceConnectConnectivity `json:"privateServiceConnectConnectivity,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.CloudSqlSettings
type CloudSQLSettingsObservedState struct {
	// Output only. Indicates If this connection profile root password is stored.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password_set
	RootPasswordSet *bool `json:"rootPasswordSet,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword
type AlloyDbSettings_UserPassword struct {
	// The database username.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.user
	User *string `json:"user,omitempty"`

	// NOTYET
	// The initial password for the user.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.password
	// Password *string `json:"password,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword
type AlloyDbSettings_UserPasswordObservedState struct {
	// Output only. Indicates if the initial_user.password field has been set.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.password_set
	PasswordSet *bool `json:"passwordSet,omitempty"`
}
