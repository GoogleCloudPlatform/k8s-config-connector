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
// krm.group: datastream.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datastream.v1
// resource: DatastreamPrivateConnection:PrivateConnection
// resource: DatastreamConnectionProfile:ConnectionProfile

package v1alpha1

// +kcc:proto=google.cloud.datastream.v1.BigQueryProfile
type BigQueryProfile struct {
}

// +kcc:proto=google.cloud.datastream.v1.ConnectionProfile
type ConnectionProfile struct {

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Oracle ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.oracle_profile
	OracleProfile *OracleProfile `json:"oracleProfile,omitempty"`

	// Cloud Storage ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.gcs_profile
	GCSProfile *GCSProfile `json:"gcsProfile,omitempty"`

	// MySQL ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.mysql_profile
	MysqlProfile *MysqlProfile `json:"mysqlProfile,omitempty"`

	// BigQuery Connection Profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.bigquery_profile
	BigqueryProfile *BigQueryProfile `json:"bigqueryProfile,omitempty"`

	// PostgreSQL Connection Profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.postgresql_profile
	PostgresqlProfile *PostgresqlProfile `json:"postgresqlProfile,omitempty"`

	// SQLServer Connection Profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.sql_server_profile
	SQLServerProfile *SQLServerProfile `json:"sqlServerProfile,omitempty"`

	// Static Service IP connectivity.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.static_service_ip_connectivity
	StaticServiceIPConnectivity *StaticServiceIPConnectivity `json:"staticServiceIPConnectivity,omitempty"`

	// Forward SSH tunnel connectivity.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.forward_ssh_connectivity
	ForwardSSHConnectivity *ForwardSSHTunnelConnectivity `json:"forwardSSHConnectivity,omitempty"`

	// Private connectivity.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.private_connectivity
	PrivateConnectivity *PrivateConnectivity `json:"privateConnectivity,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.ForwardSshTunnelConnectivity
type ForwardSSHTunnelConnectivity struct {
	// Required. Hostname for the SSH tunnel.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Required. Username for the SSH tunnel.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.username
	Username *string `json:"username,omitempty"`

	// Port for the SSH tunnel, default value is 22.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.port
	Port *int32 `json:"port,omitempty"`

	// Input only. SSH password.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.password
	Password *string `json:"password,omitempty"`

	// Input only. SSH private key.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.private_key
	PrivateKey *string `json:"privateKey,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.GcsProfile
type GCSProfile struct {
	// Required. The Cloud Storage bucket name.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsProfile.bucket
	Bucket *string `json:"bucket,omitempty"`

	// The root path inside the Cloud Storage bucket.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsProfile.root_path
	RootPath *string `json:"rootPath,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlProfile
type MysqlProfile struct {
	// Required. Hostname for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Port for the MySQL connection, default value is 3306.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.username
	Username *string `json:"username,omitempty"`

	// Optional. Input only. Password for the MySQL connection. Mutually exclusive
	//  with the `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.password
	Password *string `json:"password,omitempty"`

	// SSL configuration for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.ssl_config
	SSLConfig *MysqlSSLConfig `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSslConfig
type MysqlSSLConfig struct {
	// Input only. PEM-encoded private key associated with the Client Certificate.
	//  If this field is used then the 'client_certificate' and the
	//  'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	// Input only. PEM-encoded certificate that will be used by the replica to
	//  authenticate against the source database server. If this field is used
	//  then the 'client_key' and the 'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleAsmConfig
type OracleAsmConfig struct {
	// Required. Hostname for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Required. Port for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.username
	Username *string `json:"username,omitempty"`

	// Required. Password for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.password
	Password *string `json:"password,omitempty"`

	// Required. ASM service name for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.asm_service
	AsmService *string `json:"asmService,omitempty"`

	// Optional. Connection string attributes
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.connection_attributes
	ConnectionAttributes map[string]string `json:"connectionAttributes,omitempty"`

	// Optional. SSL configuration for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.oracle_ssl_config
	OracleSSLConfig *OracleSSLConfig `json:"oracleSSLConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleProfile
type OracleProfile struct {
	// Required. Hostname for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Port for the Oracle connection, default value is 1521.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.username
	Username *string `json:"username,omitempty"`

	// Optional. Password for the Oracle connection. Mutually exclusive with the
	//  `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.password
	Password *string `json:"password,omitempty"`

	// Required. Database for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.database_service
	DatabaseService *string `json:"databaseService,omitempty"`

	// Connection string attributes
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.connection_attributes
	ConnectionAttributes map[string]string `json:"connectionAttributes,omitempty"`

	// Optional. SSL configuration for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.oracle_ssl_config
	OracleSSLConfig *OracleSSLConfig `json:"oracleSSLConfig,omitempty"`

	// Optional. Configuration for Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.oracle_asm_config
	OracleAsmConfig *OracleAsmConfig `json:"oracleAsmConfig,omitempty"`

	// Optional. A reference to a Secret Manager resource name storing the Oracle
	//  connection password. Mutually exclusive with the `password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.secret_manager_stored_password
	SecretManagerStoredPassword *string `json:"secretManagerStoredPassword,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfig struct {
	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlProfile
type PostgresqlProfile struct {
	// Required. Hostname for the PostgreSQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Port for the PostgreSQL connection, default value is 5432.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the PostgreSQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.username
	Username *string `json:"username,omitempty"`

	// Optional. Password for the PostgreSQL connection. Mutually exclusive with
	//  the `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.password
	Password *string `json:"password,omitempty"`

	// Required. Database for the PostgreSQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.database
	Database *string `json:"database,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PrivateConnectivity
type PrivateConnectivity struct {
	// Required. A reference to a private connection resource.
	//  Format: `projects/{project}/locations/{location}/privateConnections/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnectivity.private_connection
	PrivateConnection *string `json:"privateConnection,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerProfile
type SQLServerProfile struct {
	// Required. Hostname for the SQLServer connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Port for the SQLServer connection, default value is 1433.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the SQLServer connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.username
	Username *string `json:"username,omitempty"`

	// Optional. Password for the SQLServer connection. Mutually exclusive with
	//  the `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.password
	Password *string `json:"password,omitempty"`

	// Required. Database for the SQLServer connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.database
	Database *string `json:"database,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.StaticServiceIpConnectivity
type StaticServiceIPConnectivity struct {
}

// +kcc:proto=google.cloud.datastream.v1.ConnectionProfile
type ConnectionProfileObservedState struct {
	// Output only. The resource's name.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Oracle ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.oracle_profile
	OracleProfile *OracleProfileObservedState `json:"oracleProfile,omitempty"`

	// MySQL ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.mysql_profile
	MysqlProfile *MysqlProfileObservedState `json:"mysqlProfile,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlProfile
type MysqlProfileObservedState struct {
	// SSL configuration for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.ssl_config
	SSLConfig *MysqlSSLConfigObservedState `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSslConfig
type MysqlSSLConfigObservedState struct {
	// Output only. Indicates whether the client_key field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_key_set
	ClientKeySet *bool `json:"clientKeySet,omitempty"`

	// Output only. Indicates whether the client_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_certificate_set
	ClientCertificateSet *bool `json:"clientCertificateSet,omitempty"`

	// Output only. Indicates whether the ca_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleProfile
type OracleProfileObservedState struct {
	// Optional. SSL configuration for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.oracle_ssl_config
	OracleSSLConfig *OracleSSLConfigObservedState `json:"oracleSSLConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfigObservedState struct {
	// Output only. Indicates whether the ca_certificate field has been set for
	//  this Connection-Profile.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}
