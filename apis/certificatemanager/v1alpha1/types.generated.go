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


// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig
type TrustConfig struct {
	// A user-defined name of the trust config. TrustConfig names must be
	//  unique globally and match pattern
	//  `projects/*/locations/*/trustConfigs/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.name
	Name *string `json:"name,omitempty"`

	// Set of labels associated with a TrustConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// One or more paragraphs of text description of a TrustConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.description
	Description *string `json:"description,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.etag
	Etag *string `json:"etag,omitempty"`

	// Set of trust stores to perform validation against.
	//
	//  This field is supported when TrustConfig is configured with Load Balancers,
	//  currently not supported for SPIFFE certificate validation.
	//
	//  Only one TrustStore specified is currently allowed.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.trust_stores
	TrustStores []TrustConfig_TrustStore `json:"trustStores,omitempty"`
}

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
	IntermediateCas []TrustConfig_IntermediateCA `json:"intermediateCas,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig
type TrustConfigObservedState struct {
	// Output only. The creation timestamp of a TrustConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a TrustConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
