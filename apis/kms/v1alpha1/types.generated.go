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


// +kcc:proto=google.cloud.kms.v1.Certificate
type Certificate struct {
	// Required. The raw certificate bytes in DER format.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.raw_der
	RawDer []byte `json:"rawDer,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection
type EkmConnection struct {

	// Optional. A list of
	//  [ServiceResolvers][google.cloud.kms.v1.EkmConnection.ServiceResolver] where
	//  the EKM can be reached. There should be one ServiceResolver per EKM
	//  replica. Currently, only a single
	//  [ServiceResolver][google.cloud.kms.v1.EkmConnection.ServiceResolver] is
	//  supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.service_resolvers
	ServiceResolvers []EkmConnection_ServiceResolver `json:"serviceResolvers,omitempty"`

	// Optional. Etag of the currently stored
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection].
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Describes who can perform control plane operations on the EKM. If
	//  unset, this defaults to
	//  [MANUAL][google.cloud.kms.v1.EkmConnection.KeyManagementMode.MANUAL].
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.key_management_mode
	KeyManagementMode *string `json:"keyManagementMode,omitempty"`

	// Optional. Identifies the EKM Crypto Space that this
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] maps to. Note: This
	//  field is required if
	//  [KeyManagementMode][google.cloud.kms.v1.EkmConnection.KeyManagementMode] is
	//  [CLOUD_KMS][google.cloud.kms.v1.EkmConnection.KeyManagementMode.CLOUD_KMS].
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.crypto_space_path
	CryptoSpacePath *string `json:"cryptoSpacePath,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection.ServiceResolver
type EkmConnection_ServiceResolver struct {
	// Required. The resource name of the Service Directory service pointing to
	//  an EKM replica, in the format
	//  `projects/*/locations/*/namespaces/*/services/*`.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.service_directory_service
	ServiceDirectoryService *string `json:"serviceDirectoryService,omitempty"`

	// Optional. The filter applied to the endpoints of the resolved service. If
	//  no filter is specified, all endpoints will be considered. An endpoint
	//  will be chosen arbitrarily from the filtered list for each request.
	//
	//  For endpoint filter syntax and examples, see
	//  https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.endpoint_filter
	EndpointFilter *string `json:"endpointFilter,omitempty"`

	// Required. The hostname of the EKM replica used at TLS and HTTP layers.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Required. A list of leaf server certificates used to authenticate HTTPS
	//  connections to the EKM replica. Currently, a maximum of 10
	//  [Certificate][google.cloud.kms.v1.Certificate] is supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.server_certificates
	ServerCertificates []Certificate `json:"serverCertificates,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.Certificate
type CertificateObservedState struct {
	// Output only. True if the certificate was parsed successfully.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.parsed
	Parsed *bool `json:"parsed,omitempty"`

	// Output only. The issuer distinguished name in RFC 2253 format. Only present
	//  if [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.issuer
	Issuer *string `json:"issuer,omitempty"`

	// Output only. The subject distinguished name in RFC 2253 format. Only
	//  present if [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.subject
	Subject *string `json:"subject,omitempty"`

	// Output only. The subject Alternative DNS names. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.subject_alternative_dns_names
	SubjectAlternativeDnsNames []string `json:"subjectAlternativeDnsNames,omitempty"`

	// Output only. The certificate is not valid before this time. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.not_before_time
	NotBeforeTime *string `json:"notBeforeTime,omitempty"`

	// Output only. The certificate is not valid after this time. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.not_after_time
	NotAfterTime *string `json:"notAfterTime,omitempty"`

	// Output only. The certificate serial number as a hex string. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.serial_number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Output only. The SHA-256 certificate fingerprint as a hex string. Only
	//  present if [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.sha256_fingerprint
	Sha256Fingerprint *string `json:"sha256Fingerprint,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection
type EkmConnectionObservedState struct {
	// Output only. The resource name for the
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] in the format
	//  `projects/*/locations/*/ekmConnections/*`.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which the
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] was created.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. A list of
	//  [ServiceResolvers][google.cloud.kms.v1.EkmConnection.ServiceResolver] where
	//  the EKM can be reached. There should be one ServiceResolver per EKM
	//  replica. Currently, only a single
	//  [ServiceResolver][google.cloud.kms.v1.EkmConnection.ServiceResolver] is
	//  supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.service_resolvers
	ServiceResolvers []EkmConnection_ServiceResolverObservedState `json:"serviceResolvers,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection.ServiceResolver
type EkmConnection_ServiceResolverObservedState struct {
	// Required. A list of leaf server certificates used to authenticate HTTPS
	//  connections to the EKM replica. Currently, a maximum of 10
	//  [Certificate][google.cloud.kms.v1.Certificate] is supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.server_certificates
	ServerCertificates []CertificateObservedState `json:"serverCertificates,omitempty"`
}
