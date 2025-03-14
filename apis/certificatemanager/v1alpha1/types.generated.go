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

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig.IntermediateCA
type TrustConfig_IntermediateCA struct {
	// PEM intermediate certificate used for building up paths
	//  for validation.
	//
	//  Each certificate provided in PEM format may occupy up to 5kB.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.IntermediateCA.pem_certificate
	PemCertificate *string `json:"pemCertificate,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig.TrustAnchor
type TrustConfig_TrustAnchor struct {
	// PEM root certificate of the PKI used for validation.
	//
	//  Each certificate provided in PEM format may occupy up to 5kB.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.TrustAnchor.pem_certificate
	PemCertificate *string `json:"pemCertificate,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig.TrustStore
type TrustConfig_TrustStore struct {
	// List of Trust Anchors to be used while performing validation
	//  against a given TrustStore.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.TrustStore.trust_anchors
	TrustAnchors []TrustConfig_TrustAnchor `json:"trustAnchors,omitempty"`

	// Set of intermediate CA certificates used for the path building
	//  phase of chain validation.
	//
	//  The field is currently not supported if TrustConfig is used for the
	//  workload certificate feature.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.TrustStore.intermediate_cas
	IntermediateCas []TrustConfig_IntermediateCA `json:"intermediateCAs,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.DnsAuthorization.DnsResourceRecord
type DnsAuthorization_DnsResourceRecord struct {
	// Output only. Fully qualified name of the DNS Resource Record.
	//  e.g. `_acme-challenge.example.com`
	Name *string `json:"name,omitempty"`

	// Output only. Type of the DNS Resource Record.
	//  Currently always set to "CNAME".
	Type *string `json:"type,omitempty"`

	// Output only. Data of the DNS Resource Record.
	Data *string `json:"data,omitempty"`
}
