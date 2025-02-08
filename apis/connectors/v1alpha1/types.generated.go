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


// +kcc:proto=google.cloud.connectors.v1.AuthConfig
type AuthConfig struct {
	// The type of authentication configured.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.auth_type
	AuthType *string `json:"authType,omitempty"`

	// UserPassword.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.user_password
	UserPassword *AuthConfig_UserPassword `json:"userPassword,omitempty"`

	// Oauth2JwtBearer.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.oauth2_jwt_bearer
	Oauth2JwtBearer *AuthConfig_Oauth2JwtBearer `json:"oauth2JwtBearer,omitempty"`

	// Oauth2ClientCredentials.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.oauth2_client_credentials
	Oauth2ClientCredentials *AuthConfig_Oauth2ClientCredentials `json:"oauth2ClientCredentials,omitempty"`

	// SSH Public Key.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.ssh_public_key
	SSHPublicKey *AuthConfig_SshPublicKey `json:"sshPublicKey,omitempty"`

	// List containing additional auth configs.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.additional_variables
	AdditionalVariables []ConfigVariable `json:"additionalVariables,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.Oauth2ClientCredentials
type AuthConfig_Oauth2ClientCredentials struct {
	// The client identifier.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2ClientCredentials.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Secret version reference containing the client secret.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2ClientCredentials.client_secret
	ClientSecret *Secret `json:"clientSecret,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer
type AuthConfig_Oauth2JwtBearer struct {
	// Secret version reference containing a PKCS#8 PEM-encoded private
	//  key associated with the Client Certificate. This private key will be
	//  used to sign JWTs used for the jwt-bearer authorization grant.
	//  Specified in the form as: `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.client_key
	ClientKey *Secret `json:"clientKey,omitempty"`

	// JwtClaims providers fields to generate the token.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.jwt_claims
	JwtClaims *AuthConfig_Oauth2JwtBearer_JwtClaims `json:"jwtClaims,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.JwtClaims
type AuthConfig_Oauth2JwtBearer_JwtClaims struct {
	// Value for the "iss" claim.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.JwtClaims.issuer
	Issuer *string `json:"issuer,omitempty"`

	// Value for the "sub" claim.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.JwtClaims.subject
	Subject *string `json:"subject,omitempty"`

	// Value for the "aud" claim.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.JwtClaims.audience
	Audience *string `json:"audience,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.SshPublicKey
type AuthConfig_SshPublicKey struct {
	// The user account used to authenticate.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.username
	Username *string `json:"username,omitempty"`

	// SSH Client Cert. It should contain both public and private key.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.ssh_client_cert
	SSHClientCert *Secret `json:"sshClientCert,omitempty"`

	// Format of SSH Client cert.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.cert_type
	CertType *string `json:"certType,omitempty"`

	// Password (passphrase) for ssh client certificate if it has one.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.ssh_client_cert_pass
	SSHClientCertPass *Secret `json:"sshClientCertPass,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.UserPassword
type AuthConfig_UserPassword struct {
	// Username.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.UserPassword.username
	Username *string `json:"username,omitempty"`

	// Secret version reference containing the password.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.UserPassword.password
	Password *Secret `json:"password,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ConfigVariable
type ConfigVariable struct {
	// Key of the config variable.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariable.key
	Key *string `json:"key,omitempty"`

	// Value is an integer
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariable.int_value
	IntValue *int64 `json:"intValue,omitempty"`

	// Value is a bool.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariable.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Value is a string.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariable.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Value is a secret.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariable.secret_value
	SecretValue *Secret `json:"secretValue,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.Connection
type Connection struct {

	// Optional. Resource labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the resource.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.description
	Description *string `json:"description,omitempty"`

	// Required. Connector version on which the connection is created.
	//  The format is:
	//  projects/*/locations/*/providers/*/connectors/*/versions/*
	//  Only global location is supported for ConnectorVersion resource.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.connector_version
	ConnectorVersion *string `json:"connectorVersion,omitempty"`

	// Optional. Configuration for configuring the connection with an external
	//  system.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.config_variables
	ConfigVariables []ConfigVariable `json:"configVariables,omitempty"`

	// Optional. Configuration for establishing the connection's authentication
	//  with an external system.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.auth_config
	AuthConfig *AuthConfig `json:"authConfig,omitempty"`

	// Optional. Configuration that indicates whether or not the Connection can be
	//  edited.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.lock_config
	LockConfig *LockConfig `json:"lockConfig,omitempty"`

	// Optional. Configuration of the Connector's destination. Only accepted for
	//  Connectors that accepts user defined destination(s).
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.destination_configs
	DestinationConfigs []DestinationConfig `json:"destinationConfigs,omitempty"`

	// Optional. Service account needed for runtime plane to access GCP resources.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Suspended indicates if a user has suspended a connection or not.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Optional. Node configuration for the connection.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.node_config
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// Optional. Ssl config of a connection
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.ssl_config
	SslConfig *SslConfig `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ConnectionStatus
type ConnectionStatus struct {
	// State.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionStatus.state
	State *string `json:"state,omitempty"`

	// Description.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionStatus.description
	Description *string `json:"description,omitempty"`

	// Status provides detailed information for the state.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionStatus.status
	Status *string `json:"status,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.Destination
type Destination struct {
	// PSC service attachments.
	//  Format: projects/*/regions/*/serviceAttachments/*
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`

	// For publicly routable host.
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.host
	Host *string `json:"host,omitempty"`

	// The port is the target port number that is accepted by the destination.
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.DestinationConfig
type DestinationConfig struct {
	// The key is the destination identifier that is supported by the Connector.
	// +kcc:proto:field=google.cloud.connectors.v1.DestinationConfig.key
	Key *string `json:"key,omitempty"`

	// The destinations for the key.
	// +kcc:proto:field=google.cloud.connectors.v1.DestinationConfig.destinations
	Destinations []Destination `json:"destinations,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.LockConfig
type LockConfig struct {
	// Indicates whether or not the connection is locked.
	// +kcc:proto:field=google.cloud.connectors.v1.LockConfig.locked
	Locked *bool `json:"locked,omitempty"`

	// Describes why a connection is locked.
	// +kcc:proto:field=google.cloud.connectors.v1.LockConfig.reason
	Reason *string `json:"reason,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.NodeConfig
type NodeConfig struct {
	// Minimum number of nodes in the runtime nodes.
	// +kcc:proto:field=google.cloud.connectors.v1.NodeConfig.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// Maximum number of nodes in the runtime nodes.
	// +kcc:proto:field=google.cloud.connectors.v1.NodeConfig.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.Secret
type Secret struct {
	// The resource name of the secret version in the format,
	//  format as: `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.cloud.connectors.v1.Secret.secret_version
	SecretVersion *string `json:"secretVersion,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.SslConfig
type SslConfig struct {
	// Controls the ssl type for the given connector version.
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.type
	Type *string `json:"type,omitempty"`

	// Trust Model of the SSL connection
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.trust_model
	TrustModel *string `json:"trustModel,omitempty"`

	// Private Server Certificate. Needs to be specified if trust model is
	//  `PRIVATE`.
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.private_server_certificate
	PrivateServerCertificate *Secret `json:"privateServerCertificate,omitempty"`

	// Client Certificate
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_certificate
	ClientCertificate *Secret `json:"clientCertificate,omitempty"`

	// Client Private Key
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_private_key
	ClientPrivateKey *Secret `json:"clientPrivateKey,omitempty"`

	// Secret containing the passphrase protecting the Client Private Key
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_private_key_pass
	ClientPrivateKeyPass *Secret `json:"clientPrivateKeyPass,omitempty"`

	// Type of Server Cert (PEM/JKS/.. etc.)
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.server_cert_type
	ServerCertType *string `json:"serverCertType,omitempty"`

	// Type of Client Cert (PEM/JKS/.. etc.)
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_cert_type
	ClientCertType *string `json:"clientCertType,omitempty"`

	// Bool for enabling SSL
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.use_ssl
	UseSsl *bool `json:"useSsl,omitempty"`

	// Additional SSL related field values
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.additional_variables
	AdditionalVariables []ConfigVariable `json:"additionalVariables,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.Connection
type ConnectionObservedState struct {
	// Output only. Resource name of the Connection.
	//  Format: projects/{project}/locations/{location}/connections/{connection}
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.name
	Name *string `json:"name,omitempty"`

	// Output only. Created time.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Updated time.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current status of the connection.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.status
	Status *ConnectionStatus `json:"status,omitempty"`

	// Output only. GCR location where the runtime image is stored.
	//  formatted like: gcr.io/{bucketName}/{imageName}
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.image_location
	ImageLocation *string `json:"imageLocation,omitempty"`

	// Output only. The name of the Service Directory service name. Used for
	//  Private Harpoon to resolve the ILB address.
	//  e.g.
	//  "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.service_directory
	ServiceDirectory *string `json:"serviceDirectory,omitempty"`

	// Output only. GCR location where the envoy image is stored.
	//  formatted like: gcr.io/{bucketName}/{imageName}
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.envoy_image_location
	EnvoyImageLocation *string `json:"envoyImageLocation,omitempty"`
}
