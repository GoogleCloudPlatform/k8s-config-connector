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


// +kcc:proto=google.cloud.networksecurity.v1.CertificateProvider
type CertificateProvider struct {
	// gRPC specific configuration to access the gRPC server to
	//  obtain the cert and private key.
	// +kcc:proto:field=google.cloud.networksecurity.v1.CertificateProvider.grpc_endpoint
	GrpcEndpoint *GrpcEndpoint `json:"grpcEndpoint,omitempty"`

	// The certificate provider instance specification that will be passed to
	//  the data plane, which will be used to load necessary credential
	//  information.
	// +kcc:proto:field=google.cloud.networksecurity.v1.CertificateProvider.certificate_provider_instance
	CertificateProviderInstance *CertificateProviderInstance `json:"certificateProviderInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.CertificateProviderInstance
type CertificateProviderInstance struct {
	// Required. Plugin instance name, used to locate and load CertificateProvider instance
	//  configuration. Set to "google_cloud_private_spiffe" to use Certificate
	//  Authority Service certificate provider instance.
	// +kcc:proto:field=google.cloud.networksecurity.v1.CertificateProviderInstance.plugin_instance
	PluginInstance *string `json:"pluginInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.ClientTlsPolicy
type ClientTlsPolicy struct {
	// Required. Name of the ClientTlsPolicy resource. It matches the pattern
	//  `projects/*/locations/{location}/clientTlsPolicies/{client_tls_policy}`
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.name
	Name *string `json:"name,omitempty"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of label tags associated with the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Server Name Indication string to present to the server during TLS
	//  handshake. E.g: "secure.example.com".
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.sni
	Sni *string `json:"sni,omitempty"`

	// Optional. Defines a mechanism to provision client identity (public and private keys)
	//  for peer to peer authentication. The presence of this dictates mTLS.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.client_certificate
	ClientCertificate *CertificateProvider `json:"clientCertificate,omitempty"`

	// Optional. Defines the mechanism to obtain the Certificate Authority certificate to
	//  validate the server certificate. If empty, client does not validate the
	//  server certificate.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.server_validation_ca
	ServerValidationCa []ValidationCA `json:"serverValidationCa,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.GrpcEndpoint
type GrpcEndpoint struct {
	// Required. The target URI of the gRPC endpoint. Only UDS path is supported, and
	//  should start with "unix:".
	// +kcc:proto:field=google.cloud.networksecurity.v1.GrpcEndpoint.target_uri
	TargetURI *string `json:"targetURI,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.ValidationCA
type ValidationCA struct {
	// gRPC specific configuration to access the gRPC server to
	//  obtain the CA certificate.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ValidationCA.grpc_endpoint
	GrpcEndpoint *GrpcEndpoint `json:"grpcEndpoint,omitempty"`

	// The certificate provider instance specification that will be passed to
	//  the data plane, which will be used to load necessary credential
	//  information.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ValidationCA.certificate_provider_instance
	CertificateProviderInstance *CertificateProviderInstance `json:"certificateProviderInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.ClientTlsPolicy
type ClientTlsPolicyObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.ClientTlsPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
