// Copyright 2025 Google LLC
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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	secretrefv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
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

// +kcc:proto=google.cloud.clouddms.v1.PrivateConnectivity
type PrivateConnectivity struct {
	// Required. The resource name (URI) of the private connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnectivity.private_connection
	// +kcc:ref=CloudDMSPrivateConnection
	PrivateConnectionRef *PrivateConnectionRef `json:"privateConnectionRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword
type AlloyDbSettings_UserPassword struct {
	// The Kubernetes Secret in type "kubernetes.io/basic-auth".
	// * .data.username is the AlloyDB settings user.
	// * .data.password is the AlloyDB settings password.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.user
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.password
	SecretRef *secretrefv1beta1.BasicAuthSecretRef `json:"secretRef,omitempty"`
}
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
	// +kcc:ref=ComputeNetwork
	VPCNetworkRef *computev1beta1.ComputeNetworkRef `json:"vpcNetworkRef,omitempty"`

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

type AlloyDbSettings_EncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//  Each Cloud KMS key is regionalized and has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.EncryptionConfig.kms_key_name
	// +kcc:ref=KMSCryptoKey
	KMSKeyNameRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

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
	// +kcc:ref=CloudDMSConnectionProfile
	SourceRef *CloudDMSConnectionProfileRef `json:"sourceRef,omitempty"`

	// Input only. Initial root password.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password
	RootPasswordSecret *secretrefv1beta1.Legacy `json:"rootPasswordSecret,omitempty"`

	// The Cloud SQL default instance level collation.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.collation
	Collation *string `json:"collation,omitempty"`

	// The KMS key name used for the csql instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.cmek_key_name
	// +kcc:ref=KMSCryptoKey
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

// +kcc:proto=google.cloud.clouddms.v1.MySqlConnectionProfile
type MySQLConnectionProfile struct {
	// Required. The IP or hostname of the source MySQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.host
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source MySQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.port
	Port *int32 `json:"port,omitempty"`

	// The Kubernetes Secret in type "kubernetes.io/basic-auth".
	// * .data.username is the  username that Database Migration Service will use to connect
	//    to the database. The value is encrypted when stored in Database Migration Service.
	// * .data.password is the password for the user that Database Migration
	// 	  Service will be using to connect to the database. This field is not returned on request,
	//    and the value is encrypted when stored in Database Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.username
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.password
	SecretRef *secretrefv1beta1.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// SSL configuration for the destination to connect to the source database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.ssl
	SSL *SSLConfig `json:"ssl,omitempty"`

	// If the source is a Cloud SQL database, use this field to
	//  provide the Cloud SQL instance ID of the source.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.cloud_sql_id
	// +kcc:ref=SQLInstance
	CloudSQLRef *refsv1beta1.SQLInstanceRef `json:"cloudSQLRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.PostgreSqlConnectionProfile
type PostgreSQLConnectionProfile struct {
	// Required. The IP or hostname of the source PostgreSQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.host
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source PostgreSQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.port
	Port *int32 `json:"port,omitempty"`

	// The Kubernetes Secret in type "kubernetes.io/basic-auth".
	// * .data.username is the  username that Database Migration Service will use to connect
	//    to the database. The value is encrypted when stored in Database Migration Service.
	// * .data.password is the password for the user that Database Migration
	// 	  Service will be using to connect to the database. This field is not returned on request,
	//    and the value is encrypted when stored in Database Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.username
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.password
	SecretRef *secretrefv1beta1.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// SSL configuration for the destination to connect to the source database.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.ssl
	SSL *SSLConfig `json:"ssl,omitempty"`

	// If the source is a Cloud SQL database, use this field to
	//  provide the Cloud SQL instance ID of the source.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.cloud_sql_id
	// +kcc:ref=SQLInstance
	CloudSQLRef *refsv1beta1.SQLInstanceRef `json:"cloudSQLRef,omitempty"`

	// Static ip connectivity data (default, no additional details needed).
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.static_ip_connectivity
	StaticIPConnectivity *StaticIPConnectivity `json:"staticIPConnectivity,omitempty"`

	// Private service connect connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.private_service_connect_connectivity
	PrivateServiceConnectConnectivity *PrivateServiceConnectConnectivity `json:"privateServiceConnectConnectivity,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.PrivateServiceConnectConnectivity
type PrivateServiceConnectConnectivity struct {
	// Required. A service attachment that exposes a database, and has the
	//  following format:
	//  projects/{project}/regions/{region}/serviceAttachments/{service_attachment_name}
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateServiceConnectConnectivity.service_attachment
	// +kcc:ref=ComputeServiceAttachment
	ServiceAttachmentRef *refsv1beta1.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`
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
	// +kcc:ref=ComputeNetwork
	PrivateNetworkRef *computev1beta1.ComputeNetworkRef `json:"privateNetworkRef,omitempty"`

	// Optional. The name of the allocated IP address range for the private IP
	//  Cloud SQL instance. This name refers to an already allocated IP range
	//  address. If set, the instance IP address will be created in the allocated
	//  range. Note that this IP address range can't be modified after the instance
	//  is created. If you change the VPC when configuring connectivity settings
	//  for the migration job, this field is not relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.allocated_ip_range
	AllocatedIPRange *string `json:"allocatedIPRange,omitempty"`

	// Whether SSL connections over IP should be enforced or not.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.require_ssl
	RequireSSL *bool `json:"requireSSL,omitempty"`

	// The list of external networks that are allowed to connect to the instance
	//  using the IP. See
	//  https://en.wikipedia.org/wiki/CIDR_notation#CIDR_notation, also known as
	//  'slash' notation (e.g. `192.168.100.0/24`).
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlIpConfig.authorized_networks
	AuthorizedNetworks []SQLAclEntry `json:"authorizedNetworks,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.OracleConnectionProfile
type OracleConnectionProfile struct {
	// Required. The IP or hostname of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.host
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.port
	Port *int32 `json:"port,omitempty"`

	// The Kubernetes Secret in type "kubernetes.io/basic-auth".
	// * .data.username is the  username that Database Migration Service will use to connect
	//    to the database. The value is encrypted when stored in Database Migration Service.
	// * .data.password is the password for the user that Database Migration
	// 	  Service will be using to connect to the database. This field is not returned on request,
	//    and the value is encrypted when stored in Database Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.username
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.password
	SecretRef *secretrefv1beta1.BasicAuthSecretRef `json:"secretRef,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity
type ForwardSSHTunnelConnectivity struct {
	// Required. Hostname for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Port for the SSH tunnel, default value is 22.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.port
	Port *int32 `json:"port,omitempty"`

	// The Kubernetes Secret in type "kubernetes.io/basic-auth".
	// * .data.username is the Username for the SSH tunnel.
	// * .data.password is the SSH password.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.username
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.password
	SecretRef *secretrefv1beta1.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// Input only. SSH private key.
	PrivateKeySecretRef *secretrefv1beta1.Legacy `json:"privateKeySecretRef,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SslConfig
type SSLConfig struct {

	// Input only. The unencrypted PKCS#1 or PKCS#8 PEM-encoded private key
	//  associated with the Client Certificate. If this field is used then the
	//  'client_certificate' field is mandatory.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.client_key
	ClientKey *secretrefv1beta1.Legacy `json:"clientKeySecretRef,omitempty"`

	// Input only. The x509 PEM-encoded certificate that will be used by the
	//  replica to authenticate against the source database server.If this field is
	//  used then the 'client_key' field is mandatory.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.client_certificate
	ClientCertificate *secretrefv1beta1.Legacy `json:"clientCertificateSecretRef,omitempty"`

	// Required. Input only. The x509 PEM-encoded certificate of the CA that
	//  signed the source database server's certificate. The replica will use this
	//  certificate to verify it's connecting to the right host.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.ca_certificate
	CACertificate *secretrefv1beta1.Legacy `json:"caCertificateSecretRef,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbSettings
type AlloyDbSettingsObservedState struct {
	// Required. Input only. Initial user to setup during cluster creation.
	//  Required.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.initial_user
	// InitialUser *AlloyDbSettings_UserPasswordObservedState `json:"initialUser,omitempty"`

	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.primary_instance_settings
	PrimaryInstanceSettings *AlloyDbSettings_PrimaryInstanceSettingsObservedState `json:"primaryInstanceSettings,omitempty"`
}
