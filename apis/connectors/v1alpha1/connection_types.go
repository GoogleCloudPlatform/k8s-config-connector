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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ConnectorsConnectionGVK = GroupVersion.WithKind("ConnectorsConnection")

// ConnectorsConnectionSpec defines the desired state of ConnectorsConnection
// +kcc:spec:proto=google.cloud.connectors.v1.Connection
type ConnectorsConnectionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ConnectorsConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
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
	//  projects/-*-/locations/-*-/providers/-*-/connectors/-*-/versions/-*
	//  Only global location is supported for ConnectorVersion resource.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.connector_version
	// +kubebuilder:validation:Required
	ConnectorVersionRef *ConnectorsConnectorVersionRef `json:"connectorVersionRef,omitempty"`

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
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. Suspended indicates if a user has suspended a connection or not.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Optional. Node configuration for the connection.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.node_config
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// Optional. Ssl config of a connection
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.ssl_config
	SSLConfig *SSLConfig `json:"sslConfig,omitempty"`
}

// ConnectorsConnectionStatus defines the config connector machine state of ConnectorsConnection
type ConnectorsConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ConnectorsConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ConnectorsConnectionObservedState `json:"observedState,omitempty"`
}

// ConnectorsConnectionObservedState is the state of the ConnectorsConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.connectors.v1.Connection
type ConnectorsConnectionObservedState struct {

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpconnectorsconnection;gcpconnectorsconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ConnectorsConnection is the Schema for the ConnectorsConnection API
// +k8s:openapi-gen=true
type ConnectorsConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ConnectorsConnectionSpec   `json:"spec,omitempty"`
	Status ConnectorsConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ConnectorsConnectionList contains a list of ConnectorsConnection
type ConnectorsConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectorsConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConnectorsConnection{}, &ConnectorsConnectionList{})
}

// +kcc:proto=google.cloud.connectors.v1.Destination
type Destination struct {
	// PSC service attachments.
	//  Format: projects/*/regions/*/serviceAttachments/*
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.service_attachment
	ServiceAttachmentRef *refsv1beta1.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`

	// For publicly routable host.
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.host
	Host *string `json:"host,omitempty"`

	// The port is the target port number that is accepted by the destination.
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.UserPassword
type AuthConfig_UserPassword struct {
	// Username.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.UserPassword.username
	Username *string `json:"username,omitempty"`

	// Secret version reference containing the password.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.UserPassword.password
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.SslConfig
type SSLConfig struct {
	// Controls the ssl type for the given connector version.
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.type
	Type *string `json:"type,omitempty"`

	// Trust Model of the SSL connection
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.trust_model
	TrustModel *string `json:"trustModel,omitempty"`

	// Private Server Certificate. Needs to be specified if trust model is
	//  `PRIVATE`.
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.private_server_certificate
	PrivateServerCertificateRef *secretmanagerv1beta1.SecretRef `json:"privateServerCertificateRef,omitempty"`

	// Client Certificate
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_certificate
	ClientCertificateRef *secretmanagerv1beta1.SecretRef `json:"clientCertificateRef,omitempty"`

	// Client Private Key
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_private_key
	ClientPrivateKeyRef *secretmanagerv1beta1.SecretRef `json:"clientPrivateKeyRef,omitempty"`

	// Secret containing the passphrase protecting the Client Private Key
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_private_key_pass
	ClientPrivateKeyPassRef *secretmanagerv1beta1.SecretRef `json:"clientPrivateKeyPassRef,omitempty"`

	// Type of Server Cert (PEM/JKS/.. etc.)
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.server_cert_type
	ServerCertType *string `json:"serverCertType,omitempty"`

	// Type of Client Cert (PEM/JKS/.. etc.)
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.client_cert_type
	ClientCertType *string `json:"clientCertType,omitempty"`

	// Bool for enabling SSL
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.use_ssl
	UseSSL *bool `json:"useSSL,omitempty"`

	// Additional SSL related field values
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfig.additional_variables
	AdditionalVariables []ConfigVariable `json:"additionalVariables,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.SshPublicKey
type AuthConfig_SSHPublicKey struct {
	// The user account used to authenticate.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.username
	Username *string `json:"username,omitempty"`

	// SSH Client Cert. It should contain both public and private key.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.ssh_client_cert
	SSHClientCertRef *secretmanagerv1beta1.SecretRef `json:"sshClientCertRef,omitempty"`

	// Format of SSH Client cert.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.cert_type
	CertType *string `json:"certType,omitempty"`

	// Password (passphrase) for ssh client certificate if it has one.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.SshPublicKey.ssh_client_cert_pass
	SSHClientCertPassRef *secretmanagerv1beta1.SecretRef `json:"sshClientCertPassRef,omitempty"`
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
	SecretValueRef *secretmanagerv1beta1.SecretRef `json:"secretValueRef,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.Oauth2ClientCredentials
type AuthConfig_OAUTH2ClientCredentials struct {
	// The client identifier.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2ClientCredentials.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Secret version reference containing the client secret.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2ClientCredentials.client_secret
	ClientSecretRef *secretmanagerv1beta1.SecretRef `json:"clientSecretRef,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer
type AuthConfig_OAUTH2JwtBearer struct {
	// Secret version reference containing a PKCS#8 PEM-encoded private
	//  key associated with the Client Certificate. This private key will be
	//  used to sign JWTs used for the jwt-bearer authorization grant.
	//  Specified in the form as: `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.client_key
	ClientKeyRef *secretmanagerv1beta1.SecretRef `json:"clientKeyRef,omitempty"`

	// JwtClaims providers fields to generate the token.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.Oauth2JwtBearer.jwt_claims
	JwtClaims *AuthConfig_OAUTH2JwtBearer_JwtClaims `json:"jwtClaims,omitempty"`
}
