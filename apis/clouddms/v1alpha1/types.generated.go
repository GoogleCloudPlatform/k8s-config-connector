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

// +kcc:proto=google.cloud.clouddms.v1.CloudSqlConnectionProfile
type CloudSQLConnectionProfile struct {

	// Immutable. Metadata used to create the destination Cloud SQL database.
	// +kcc:proto:field=google.cloud.clouddms.v1.CloudSqlConnectionProfile.settings
	Settings *CloudSQLSettings `json:"settings,omitempty"`
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

// +kcc:observedstate:proto=google.cloud.clouddms.v1.SslConfig
type SSLConfigObservedState struct {
	// Output only. The ssl config type according to 'client_key',
	//  'client_certificate' and 'ca_certificate'.
	// +kcc:proto:field=google.cloud.clouddms.v1.SslConfig.type
	Type *string `json:"type,omitempty"`
}
