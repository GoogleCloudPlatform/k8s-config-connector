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
// resource: DatastreamRoute:Route

package v1alpha1

// +kcc:proto=google.cloud.datastream.v1.BigQueryProfile
type BigQueryProfile struct {
}

// +kcc:proto=google.cloud.datastream.v1.HostAddress
type HostAddress struct {
	// Required. Hostname for the connection.
	// +kcc:proto:field=google.cloud.datastream.v1.HostAddress.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Optional. Port for the connection.
	// +kcc:proto:field=google.cloud.datastream.v1.HostAddress.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MongodbProfile
type MongodbProfile struct {
	// Required. List of host addresses for a MongoDB cluster.
	//  For SRV connection format, this list must contain exactly one DNS host
	//  without a port. For Standard connection format, this list must contain all
	//  the required hosts in the cluster with their respective ports.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.host_addresses
	HostAddresses []HostAddress `json:"hostAddresses,omitempty"`

	// Optional. Name of the replica set. Only needed for self hosted replica set
	//  type MongoDB cluster. For SRV connection format, this field must be empty.
	//  For Standard connection format, this field must be specified.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.replica_set
	ReplicaSet *string `json:"replicaSet,omitempty"`

	// Required. Username for the MongoDB connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.username
	Username *string `json:"username,omitempty"`

	// Optional. Password for the MongoDB connection. Mutually exclusive with the
	//  `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.password
	Password *string `json:"password,omitempty"`

	// Optional. A reference to a Secret Manager resource name storing the
	//  SQLServer connection password. Mutually exclusive with the `password`
	//  field.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.secret_manager_stored_password
	SecretManagerStoredPassword *string `json:"secretManagerStoredPassword,omitempty"`

	// Optional. SSL configuration for the MongoDB connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.ssl_config
	SSLConfig *MongodbSSLConfig `json:"sslConfig,omitempty"`

	// Srv connection format.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.srv_connection_format
	SrvConnectionFormat *SrvConnectionFormat `json:"srvConnectionFormat,omitempty"`

	// Standard connection format.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.standard_connection_format
	StandardConnectionFormat *StandardConnectionFormat `json:"standardConnectionFormat,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MongodbSslConfig
type MongodbSSLConfig struct {
	// Optional. Input only. PEM-encoded private key associated with the Client
	//  Certificate. If this field is used then the 'client_certificate' and the
	//  'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbSslConfig.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	// Optional. Input only. PEM-encoded certificate that will be used by the
	//  replica to authenticate against the source database server. If this field
	//  is used then the 'client_key' and the 'ca_certificate' fields are
	//  mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbSslConfig.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Optional. Input only. PEM-encoded certificate of the CA that signed the
	//  source database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`

	// Optional. Input only. A reference to a Secret Manager resource name storing
	//  the PEM-encoded private key associated with the Client Certificate. If this
	//  field is used then the 'client_certificate' and the 'ca_certificate' fields
	//  are mandatory. Mutually exclusive with the `client_key` field.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbSslConfig.secret_manager_stored_client_key
	SecretManagerStoredClientKey *string `json:"secretManagerStoredClientKey,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSslConfig
type MysqlSSLConfig struct {
	// Optional. Input only. PEM-encoded private key associated with the Client
	//  Certificate. If this field is used then the 'client_certificate' and the
	//  'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	// Optional. Input only. PEM-encoded certificate that will be used by the
	//  replica to authenticate against the source database server. If this field
	//  is used then the 'client_key' and the 'ca_certificate' fields are
	//  mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfig struct {
	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`

	// Optional. The distinguished name (DN) mentioned in the server
	//  certificate. This corresponds to SSL_SERVER_CERT_DN sqlnet parameter.
	//  Refer
	//  https://docs.oracle.com/en/database/oracle/oracle-database/19/netrf/local-naming-parameters-in-tns-ora-file.html#GUID-70AB0695-A9AA-4A94-B141-4C605236EEB7
	//  If this field is not provided, the DN matching is not enforced.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.server_certificate_distinguished_name
	ServerCertificateDistinguishedName *string `json:"serverCertificateDistinguishedName,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlSslConfig
type PostgresqlSSLConfig struct {
	// If this field is set, the communication will be encrypted with TLS
	//   encryption and the server identity will be authenticated.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.server_verification
	ServerVerification *PostgresqlSSLConfig_ServerVerification `json:"serverVerification,omitempty"`

	// If this field is set, the communication will be encrypted with TLS
	//  encryption and both the server identity and the client identity will be
	//  authenticated.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.server_and_client_verification
	ServerAndClientVerification *PostgresqlSSLConfig_ServerAndClientVerification `json:"serverAndClientVerification,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlSslConfig.ServerAndClientVerification
type PostgresqlSSLConfig_ServerAndClientVerification struct {
	// Required. Input only. PEM-encoded certificate used by the source database
	//  to authenticate the client identity (i.e., the Datastream's identity).
	//  This certificate is signed by either a root certificate trusted by the
	//  server or one or more intermediate certificates (which is stored with the
	//  leaf certificate) to link the this certificate to the trusted root
	//  certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.ServerAndClientVerification.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Optional. Input only. PEM-encoded private key associated with the client
	//  certificate. This value will be used during the SSL/TLS handshake,
	//  allowing the PostgreSQL server to authenticate the client's identity,
	//  i.e. identity of the Datastream.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.ServerAndClientVerification.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	// Required. Input only. PEM-encoded server root CA certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.ServerAndClientVerification.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`

	// Optional. The hostname mentioned in the Subject or SAN extension of the
	//  server certificate. If this field is not provided, the hostname in the
	//  server certificate is not validated.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.ServerAndClientVerification.server_certificate_hostname
	ServerCertificateHostname *string `json:"serverCertificateHostname,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlSslConfig.ServerVerification
type PostgresqlSSLConfig_ServerVerification struct {
	// Required. Input only. PEM-encoded server root CA certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.ServerVerification.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`

	// Optional. The hostname mentioned in the Subject or SAN extension of the
	//  server certificate. If this field is not provided, the hostname in the
	//  server certificate is not validated.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSslConfig.ServerVerification.server_certificate_hostname
	ServerCertificateHostname *string `json:"serverCertificateHostname,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PscInterfaceConfig
type PSCInterfaceConfig struct {
	// Required. Fully qualified name of the Network Attachment that Datastream
	//  will connect to. Format:
	//  `projects/{project}/regions/{region}/networkAttachments/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.PscInterfaceConfig.network_attachment
	NetworkAttachment *string `json:"networkAttachment,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SalesforceProfile
type SalesforceProfile struct {
	// Required. Domain endpoint for the Salesforce connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.domain
	Domain *string `json:"domain,omitempty"`

	// User-password authentication.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.user_credentials
	UserCredentials *SalesforceProfile_UserCredentials `json:"userCredentials,omitempty"`

	// Connected app authentication.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.oauth2_client_credentials
	OAUTH2ClientCredentials *SalesforceProfile_OAUTH2ClientCredentials `json:"oauth2ClientCredentials,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SalesforceProfile.Oauth2ClientCredentials
type SalesforceProfile_OAUTH2ClientCredentials struct {
	// Required. Client ID for Salesforce OAuth2 Client Credentials.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.Oauth2ClientCredentials.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Optional. Client secret for Salesforce OAuth2 Client Credentials.
	//  Mutually exclusive with the `secret_manager_stored_client_secret` field.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.Oauth2ClientCredentials.client_secret
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Optional. A reference to a Secret Manager resource name storing the
	//  Salesforce OAuth2 client_secret. Mutually exclusive with the
	//  `client_secret` field.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.Oauth2ClientCredentials.secret_manager_stored_client_secret
	SecretManagerStoredClientSecret *string `json:"secretManagerStoredClientSecret,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SalesforceProfile.UserCredentials
type SalesforceProfile_UserCredentials struct {
	// Required. Username for the Salesforce connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.UserCredentials.username
	Username *string `json:"username,omitempty"`

	// Optional. Password for the Salesforce connection.
	//  Mutually exclusive with the `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.UserCredentials.password
	Password *string `json:"password,omitempty"`

	// Optional. Security token for the Salesforce connection.
	//  Mutually exclusive with the `secret_manager_stored_security_token` field.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.UserCredentials.security_token
	SecurityToken *string `json:"securityToken,omitempty"`

	// Optional. A reference to a Secret Manager resource name storing the
	//  Salesforce connection's password. Mutually exclusive with the `password`
	//  field.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.UserCredentials.secret_manager_stored_password
	SecretManagerStoredPassword *string `json:"secretManagerStoredPassword,omitempty"`

	// Optional. A reference to a Secret Manager resource name storing the
	//  Salesforce connection's security token. Mutually exclusive with the
	//  `security_token` field.
	// +kcc:proto:field=google.cloud.datastream.v1.SalesforceProfile.UserCredentials.secret_manager_stored_security_token
	SecretManagerStoredSecurityToken *string `json:"secretManagerStoredSecurityToken,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SrvConnectionFormat
type SrvConnectionFormat struct {
}

// +kcc:proto=google.cloud.datastream.v1.StandardConnectionFormat
type StandardConnectionFormat struct {
	// Optional. Specifies whether the client connects directly to the host[:port]
	//  in the connection URI.
	// +kcc:proto:field=google.cloud.datastream.v1.StandardConnectionFormat.direct_connection
	DirectConnection *bool `json:"directConnection,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.StaticServiceIpConnectivity
type StaticServiceIPConnectivity struct {
}

// +kcc:observedstate:proto=google.cloud.datastream.v1.MongodbProfile
type MongodbProfileObservedState struct {
	// Optional. SSL configuration for the MongoDB connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbProfile.ssl_config
	SSLConfig *MongodbSSLConfigObservedState `json:"sslConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.datastream.v1.MongodbSslConfig
type MongodbSSLConfigObservedState struct {
	// Output only. Indicates whether the client_key field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbSslConfig.client_key_set
	ClientKeySet *bool `json:"clientKeySet,omitempty"`

	// Output only. Indicates whether the client_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbSslConfig.client_certificate_set
	ClientCertificateSet *bool `json:"clientCertificateSet,omitempty"`

	// Output only. Indicates whether the ca_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MongodbSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.datastream.v1.MysqlSslConfig
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

// +kcc:observedstate:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfigObservedState struct {
	// Output only. Indicates whether the ca_certificate field has been set for
	//  this Connection-Profile.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}
