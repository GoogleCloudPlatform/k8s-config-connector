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


// +kcc:proto=google.cloud.networksecurity.v1beta1.CertificateProvider
type CertificateProvider struct {
	// gRPC specific configuration to access the gRPC server to
	//  obtain the cert and private key.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.CertificateProvider.grpc_endpoint
	GrpcEndpoint *GrpcEndpoint `json:"grpcEndpoint,omitempty"`

	// The certificate provider instance specification that will be passed to
	//  the data plane, which will be used to load necessary credential
	//  information.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.CertificateProvider.certificate_provider_instance
	CertificateProviderInstance *CertificateProviderInstance `json:"certificateProviderInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.CertificateProviderInstance
type CertificateProviderInstance struct {
	// Required. Plugin instance name, used to locate and load CertificateProvider
	//  instance configuration. Set to "google_cloud_private_spiffe" to use
	//  Certificate Authority Service certificate provider instance.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.CertificateProviderInstance.plugin_instance
	PluginInstance *string `json:"pluginInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.GrpcEndpoint
type GrpcEndpoint struct {
	// Required. The target URI of the gRPC endpoint. Only UDS path is supported,
	//  and should start with "unix:".
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.GrpcEndpoint.target_uri
	TargetURI *string `json:"targetURI,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.ServerTlsPolicy
type ServerTlsPolicy struct {
	// Required. Name of the ServerTlsPolicy resource. It matches the pattern
	//  `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.name
	Name *string `json:"name,omitempty"`

	// Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.description
	Description *string `json:"description,omitempty"`

	// Set of label tags associated with the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Determines if server allows plaintext connections. If set to true, server
	//  allows plain text connections. By default, it is set to false. This setting
	//  is not exclusive of other encryption modes. For example, if `allow_open`
	//  and `mtls_policy` are set, server allows both plain text and mTLS
	//  connections. See documentation of other encryption modes to confirm
	//  compatibility.
	//
	//  Consider using it if you wish to upgrade in place your deployment to TLS
	//  while having mixed TLS and non-TLS traffic reaching port :80.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.allow_open
	AllowOpen *bool `json:"allowOpen,omitempty"`

	// Defines a mechanism to provision server identity (public and private keys).
	//  Cannot be combined with `allow_open` as a permissive mode that allows both
	//  plain text and TLS is not supported.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.server_certificate
	ServerCertificate *CertificateProvider `json:"serverCertificate,omitempty"`

	// Defines a mechanism to provision peer validation certificates for peer to
	//  peer authentication (Mutual TLS - mTLS). If not specified, client
	//  certificate will not be requested. The connection is treated as TLS and not
	//  mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain
	//  text and mTLS connections.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.mtls_policy
	MtlsPolicy *ServerTlsPolicy_MTLSPolicy `json:"mtlsPolicy,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.MTLSPolicy
type ServerTlsPolicy_MTLSPolicy struct {
	// Defines the mechanism to obtain the Certificate Authority certificate to
	//  validate the client certificate.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.MTLSPolicy.client_validation_ca
	ClientValidationCa []ValidationCA `json:"clientValidationCa,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.ValidationCA
type ValidationCA struct {
	// gRPC specific configuration to access the gRPC server to
	//  obtain the CA certificate.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ValidationCA.grpc_endpoint
	GrpcEndpoint *GrpcEndpoint `json:"grpcEndpoint,omitempty"`

	// The certificate provider instance specification that will be passed to
	//  the data plane, which will be used to load necessary credential
	//  information.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ValidationCA.certificate_provider_instance
	CertificateProviderInstance *CertificateProviderInstance `json:"certificateProviderInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.ServerTlsPolicy
type ServerTlsPolicyObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ServerTlsPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
