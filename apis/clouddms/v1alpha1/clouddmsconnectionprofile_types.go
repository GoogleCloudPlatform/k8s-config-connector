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

// API sources for CloudDMSConnectionProfile, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/clouddms/v1/clouddms_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/database-migration/
// +kcc:source:resource-docs=https://docs.cloud.google.com/database-migration/docs/reference/rest/v1/projects.locations.connectionProfiles

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSConnectionProfileGVK = GroupVersion.WithKind("CloudDMSConnectionProfile")

// CloudDMSConnectionProfileSpec defines the desired state of CloudDMSConnectionProfile
// +kcc:spec:proto=google.cloud.clouddms.v1.ConnectionProfile
// +kcc:required-from-proto
type CloudDMSConnectionProfileSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The CloudDMSConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.EncryptionConfig
type AlloyDbSettings_EncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//  Each Cloud KMS key is regionalized and has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.EncryptionConfig.kms_key_name
	KMSKeyNameRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword
type AlloyDbSettings_UserPassword struct {
	// Required. The database username.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.user
	User *string `json:"user,omitempty"`

	// The initial password for the user.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.password
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings
type AlloyDbSettings struct {
	// Required. Input only. Initial user to setup during cluster creation.
	//  Required.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.initial_user
	// +required
	InitialUser *AlloyDbSettings_UserPassword `json:"initialUser,omitempty"`

	// Required. The resource link for the VPC network in which cluster resources
	//  are created and from which they are accessible via Private IP. The network
	//  must belong to the same project as the cluster. It is specified in the
	//  form: "projects/{project_number}/global/networks/{network_id}". This is
	//  required to create a cluster.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.vpc_network
	// +required
	VPCNetworkRef *computerefs.ComputeNetworkRef `json:"vpcNetworkRef,omitempty"`

	// Labels for the AlloyDB cluster created by DMS. An object containing a list
	//  of 'key', 'value' pairs.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Settings for the cluster's primary instance
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

// +kcc:proto=google.cloud.clouddms.v1.CloudSqlSettings
type CloudSQLSettings struct {
	// The database engine type and version.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	// The labels for the Cloud SQL database instance. An object containing a list
	//  of 'key', 'value' pairs.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.user_labels
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// The tier of the Cloud SQL database instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.tier
	Tier *string `json:"tier,omitempty"`

	// The storage capacity of the database instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.storage_auto_resize_limit
	StorageAutoResizeLimit *int64 `json:"storageAutoResizeLimit,omitempty"`

	// The activation policy for the database instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.activation_policy
	ActivationPolicy *string `json:"activationPolicy,omitempty"`

	// The settings for IP Management.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.ip_config
	IPConfig *SQLIPConfig `json:"ipConfig,omitempty"`

	// Storage auto increase.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.auto_storage_increase
	AutoStorageIncrease *bool `json:"autoStorageIncrease,omitempty"`

	// Database flags to pass to Cloud SQL database when the database is created.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.database_flags
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty"`

	// The type of storage: `PD_SSD` (default) or `PD_HDD`.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.data_disk_type
	DataDiskType *string `json:"dataDiskType,omitempty"`

	// The storage capacity for the database.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.data_disk_size_gb
	DataDiskSizeGB *int64 `json:"dataDiskSizeGB,omitempty"`

	// The Google Cloud Platform zone where your Cloud SQL database instance is
	//  located.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.zone
	Zone *string `json:"zone,omitempty"`

	// The Google Cloud Platform secondary zone where your Cloud SQL database
	//  instance will be located in case of a failover.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.secondary_zone
	SecondaryZone *string `json:"secondaryZone,omitempty"`

	// The Database Migration Service source connection profile ID,
	//  in the format:
	//  `projects/my_project_name/locations/us-central1/connectionProfiles/connection_profile_ID`
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.source_id
	SourceIDRef *CloudDMSConnectionProfileRef `json:"sourceRef,omitempty"`

	// Input only. Initial root password.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password
	RootPassword *refsv1beta1secret.Legacy `json:"rootPassword,omitempty"`

	// The Cloud SQL default instance level collation.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.collation
	Collation *string `json:"collation,omitempty"`

	// The KMS key name used for the csql instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.cmek_key_name
	CmekKeyNameRef *kmsv1beta1.KMSCryptoKeyRef `json:"cmekKeyNameRef,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.SqlIpConfig
type SQLIPConfig struct {
	// Whether the instance should be assigned an IPv4 address or not.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.enable_ipv4
	EnableIPV4 *bool `json:"enableIPV4,omitempty"`

	// The resource link for the VPC network from which the Cloud SQL instance is
	//  accessible for private IP. For example,
	//  `projects/myProject/global/networks/default`. This setting can
	//  be updated, but it cannot be removed after it is set.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.private_network
	PrivateNetworkRef *computerefs.ComputeNetworkRef `json:"privateNetworkRef,omitempty"`

	// Optional. The name of the allocated IP address range for the private IP
	//  Cloud SQL instance. This name refers to an already allocated IP range
	//  address. If set, the instance IP address will be created in the allocated
	//  range. Note that this IP address range can't be modified after the instance
	//  is created. If you change the VPC when configuring connectivity settings
	//  for the migration job, this field is not relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.allocated_ip_range
	AllocatedIPRange *string `json:"allocatedIPRange,omitempty"`

	// Whether SSL connections over IP are enforced or not.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.require_ssl
	RequireSSL *bool `json:"requireSSL,omitempty"`

	// The list of external networks that are allowed to connect to the instance
	//  using the IP. See
	//  https://en.wikipedia.org/wiki/CIDR_notation#CIDR_address_blocks
	//  for information on block notation.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.authorized_networks
	AuthorizedNetworks []SQLAclEntry `json:"authorizedNetworks,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.PrivateConnectivity
type PrivateConnectivity struct {
	// Required. The resource name (URI) of the private connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnectivity.private_connection
	// +required
	PrivateConnectionRef *PrivateConnectionRef `json:"privateConnectionRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.PrivateServiceConnectConnectivity
type PrivateServiceConnectConnectivity struct {
	// Required. A service attachment that exposes a database, and has the
	//  following format:
	//  projects/{project}/regions/{region}/serviceAttachments/{service_attachment_name}
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateServiceConnectConnectivity.service_attachment
	// +required
	ServiceAttachmentRef *refsv1beta1.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MySqlConnectionProfile
type MySQLConnectionProfile struct {
	// Required. The IP or hostname of the source MySQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.host
	// +required
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source MySQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.port
	// +required
	Port *int32 `json:"port,omitempty"`

	// Required. The username that Database Migration Service will use to connect
	//  to the database. The value is encrypted when stored in Database Migration
	//  Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.username
	// +required
	Username *string `json:"username,omitempty"`

	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.password
	// +required
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`

	// SSL configuration for the destination to connect to the source database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.ssl
	SSL *SSLConfig `json:"ssl,omitempty"`

	// If the source is a Cloud SQL database, use this field to
	//  provide the Cloud SQL instance ID of the source.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.cloud_sql_id
	CloudSQLID *string `json:"cloudSQLID,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity
type ForwardSSHTunnelConnectivity struct {
	// Required. Hostname for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Required. Username for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.username
	// +required
	Username *string `json:"username,omitempty"`

	// Port for the SSH tunnel, default value is 22.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.port
	Port *int32 `json:"port,omitempty"`

	// Input only. SSH password.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.password
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`

	// Input only. SSH private key.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.private_key
	PrivateKey *string `json:"privateKey,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.OracleConnectionProfile
type OracleConnectionProfile struct {
	// Required. The IP or hostname of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.host
	// +required
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.port
	// +required
	Port *int32 `json:"port,omitempty"`

	// Required. The username that Database Migration Service will use to connect
	//  to the database. The value is encrypted when stored in Database Migration
	//  Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.username
	// +required
	Username *string `json:"username,omitempty"`

	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.password
	// +required
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`

	// Required. Database service for the Oracle connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.database_service
	// +required
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
	// +required
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source PostgreSQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.port
	// +required
	Port *int32 `json:"port,omitempty"`

	// Required. The username that Database Migration Service will use to connect
	//  to the database. The value is encrypted when stored in Database Migration
	//  Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.username
	// +required
	Username *string `json:"username,omitempty"`

	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.password
	// +required
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`

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

	// Output only. The current connection profile state (e.g. DRAFT, READY, or FAILED).
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
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsconnectionprofile;gcpclouddmsconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
