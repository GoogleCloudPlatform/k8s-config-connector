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
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1
// resource: CloudDMSConversionWorkspace:ConversionWorkspace
// resource: CloudDMSPrivateConnection:PrivateConnection
// resource: CloudDMSMigrationJob:MigrationJob
// resource: CloudDMSConnectionProfile:ConnectionProfile

package v1alpha1

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbConnectionProfile
type AlloyDbConnectionProfile struct {
	// Required. The AlloyDB cluster ID that this connection profile is associated
	//  with.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbConnectionProfile.cluster_id
	ClusterID *string `json:"clusterID,omitempty"`

	// Immutable. Metadata used to create the destination AlloyDB cluster.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbConnectionProfile.settings
	Settings *AlloyDbSettings `json:"settings,omitempty"`
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
	VPCNetwork *string `json:"vpcNetwork,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.EncryptionConfig
type AlloyDbSettings_EncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//  Each Cloud KMS key is regionalized and has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.EncryptionConfig.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings
type AlloyDbSettings_PrimaryInstanceSettings struct {
	// Required. The ID of the AlloyDB primary instance. The ID must satisfy the
	//  regex expression "[a-z0-9-]+".
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.id
	ID *string `json:"id,omitempty"`

	// Configuration for the machines that host the underlying
	//  database engine.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.machine_config
	MachineConfig *AlloyDbSettings_PrimaryInstanceSettings_MachineConfig `json:"machineConfig,omitempty"`

	// Database flags to pass to AlloyDB when DMS is creating the AlloyDB
	//  cluster and instances. See the AlloyDB documentation for how these can be
	//  used.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.database_flags
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty"`

	// Labels for the AlloyDB primary instance created by DMS. An object
	//  containing a list of 'key', 'value' pairs.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.MachineConfig
type AlloyDbSettings_PrimaryInstanceSettings_MachineConfig struct {
	// The number of CPU's in the VM instance.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.MachineConfig.cpu_count
	CPUCount *int32 `json:"cpuCount,omitempty"`
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

// +kcc:proto=google.cloud.clouddms.v1.CloudSqlConnectionProfile
type CloudSQLConnectionProfile struct {

	// Immutable. Metadata used to create the destination Cloud SQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlConnectionProfile.settings
	Settings *CloudSQLSettings `json:"settings,omitempty"`
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

	// Input only. Initial root password.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password
	RootPassword *string `json:"rootPassword,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.ConversionWorkspaceInfo
type ConversionWorkspaceInfo struct {
	// The resource name (URI) of the conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspaceInfo.name
	Name *string `json:"name,omitempty"`

	// The commit ID of the conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspaceInfo.commit_id
	CommitID *string `json:"commitID,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.DatabaseEngineInfo
type DatabaseEngineInfo struct {
	// Required. Engine type.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.engine
	Engine *string `json:"engine,omitempty"`

	// Required. Engine named version, for example 12.c.1.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.DatabaseType
type DatabaseType struct {
	// The database provider.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseType.provider
	Provider *string `json:"provider,omitempty"`

	// The database engine.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseType.engine
	Engine *string `json:"engine,omitempty"`
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

	// Input only. SSH password.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.password
	Password *string `json:"password,omitempty"`

	// Input only. SSH private key.
	// +kcc:proto:field=google.cloud.clouddms.v1.ForwardSshTunnelConnectivity.private_key
	PrivateKey *string `json:"privateKey,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.DumpFlag
type MigrationJob_DumpFlag struct {
	// The name of the flag
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlag.name
	Name *string `json:"name,omitempty"`

	// The value of the flag.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlag.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.DumpFlags
type MigrationJob_DumpFlags struct {
	// The flags for the initial dump.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlags.dump_flags
	DumpFlags []MigrationJob_DumpFlag `json:"dumpFlags,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.PerformanceConfig
type MigrationJob_PerformanceConfig struct {
	// Initial dump parallelism level.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.PerformanceConfig.dump_parallel_level
	DumpParallelLevel *string `json:"dumpParallelLevel,omitempty"`
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

	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.password
	Password *string `json:"password,omitempty"`

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

	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.password
	Password *string `json:"password,omitempty"`

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

	// Required. Input only. The password for the user that Database Migration
	//  Service will be using to connect to the database. This field is not
	//  returned on request, and the value is encrypted when stored in Database
	//  Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.password
	Password *string `json:"password,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.PrivateConnectivity
type PrivateConnectivity struct {
	// Required. The resource name (URI) of the private connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnectivity.private_connection
	PrivateConnection *string `json:"privateConnection,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.PrivateServiceConnectConnectivity
type PrivateServiceConnectConnectivity struct {
	// Required. A service attachment that exposes a database, and has the
	//  following format:
	//  projects/{project}/regions/{region}/serviceAttachments/{service_attachment_name}
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateServiceConnectConnectivity.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SqlAclEntry
type SQLAclEntry struct {
	// The allowlisted value for the access control list.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlAclEntry.value
	Value *string `json:"value,omitempty"`

	// The time when this access control entry expires in
	//  [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example:
	//  `2012-11-15T16:19:00.094Z`.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlAclEntry.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The time-to-leave of this access control entry.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlAclEntry.ttl
	TTL *string `json:"ttl,omitempty"`

	// A label to identify this entry.
	// +kcc:proto:field=google.cloud.clouddms.v1.SqlAclEntry.label
	Label *string `json:"label,omitempty"`
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
	PrivateNetwork *string `json:"privateNetwork,omitempty"`

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

// +kcc:proto=google.cloud.clouddms.v1.SslConfig
type SSLConfig struct {

	// Input only. The unencrypted PKCS#1 or PKCS#8 PEM-encoded private key
	//  associated with the Client Certificate. If this field is used then the
	//  'client_certificate' field is mandatory.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	// Input only. The x509 PEM-encoded certificate that will be used by the
	//  replica to authenticate against the source database server.If this field is
	//  used then the 'client_key' field is mandatory.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Required. Input only. The x509 PEM-encoded certificate of the CA that
	//  signed the source database server's certificate. The replica will use this
	//  certificate to verify it's connecting to the right host.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.StaticIpConnectivity
type StaticIPConnectivity struct {
}

// +kcc:proto=google.cloud.clouddms.v1.StaticServiceIpConnectivity
type StaticServiceIPConnectivity struct {
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbConnectionProfile
type AlloyDbConnectionProfileObservedState struct {
	// Immutable. Metadata used to create the destination AlloyDB cluster.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbConnectionProfile.settings
	Settings *AlloyDbSettingsObservedState `json:"settings,omitempty"`
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

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings
type AlloyDbSettings_PrimaryInstanceSettingsObservedState struct {
	// Output only. The private IP address for the Instance.
	//  This is the connection endpoint for an end-user application.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.PrimaryInstanceSettings.private_ip
	PrivateIP *string `json:"privateIP,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword
type AlloyDbSettings_UserPasswordObservedState struct {
	// Output only. Indicates if the initial_user.password field has been set.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbSettings.UserPassword.password_set
	PasswordSet *bool `json:"passwordSet,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.CloudSqlConnectionProfile
type CloudSQLConnectionProfileObservedState struct {
	// Output only. The Cloud SQL instance ID that this connection profile is
	//  associated with.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlConnectionProfile.cloud_sql_id
	CloudSQLID *string `json:"cloudSQLID,omitempty"`

	// Immutable. Metadata used to create the destination Cloud SQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlConnectionProfile.settings
	Settings *CloudSQLSettingsObservedState `json:"settings,omitempty"`

	// Output only. The Cloud SQL database instance's private IP.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlConnectionProfile.private_ip
	PrivateIP *string `json:"privateIP,omitempty"`

	// Output only. The Cloud SQL database instance's public IP.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlConnectionProfile.public_ip
	PublicIP *string `json:"publicIP,omitempty"`

	// Output only. The Cloud SQL database instance's additional (outgoing) public
	//  IP. Used when the Cloud SQL database availability type is REGIONAL (i.e.
	//  multiple zones / highly available).
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlConnectionProfile.additional_public_ip
	AdditionalPublicIP *string `json:"additionalPublicIP,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.CloudSqlSettings
type CloudSQLSettingsObservedState struct {
	// Output only. Indicates If this connection profile root password is stored.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlSettings.root_password_set
	RootPasswordSet *bool `json:"rootPasswordSet,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.MySqlConnectionProfile
type MySQLConnectionProfileObservedState struct {
	// Output only. Indicates If this connection profile password is stored.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.password_set
	PasswordSet *bool `json:"passwordSet,omitempty"`

	// SSL configuration for the destination to connect to the source database.
	// +kcc:proto:field=google.cloud.clouddms.v1.MySqlConnectionProfile.ssl
	SSL *SSLConfigObservedState `json:"ssl,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.OracleConnectionProfile
type OracleConnectionProfileObservedState struct {
	// Output only. Indicates whether a new password is included in the request.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.password_set
	PasswordSet *bool `json:"passwordSet,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.PostgreSqlConnectionProfile
type PostgreSQLConnectionProfileObservedState struct {
	// Output only. Indicates If this connection profile password is stored.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.password_set
	PasswordSet *bool `json:"passwordSet,omitempty"`

	// Output only. If the source is a Cloud SQL database, this field indicates
	//  the network architecture it's associated with.
	// +kcc:proto:field=google.cloud.clouddms.v1.PostgreSqlConnectionProfile.network_architecture
	NetworkArchitecture *string `json:"networkArchitecture,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.clouddms.v1.SslConfig
type SSLConfigObservedState struct {
	// Output only. The ssl config type according to 'client_key',
	//  'client_certificate' and 'ca_certificate'.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.type
	Type *string `json:"type,omitempty"`
}
