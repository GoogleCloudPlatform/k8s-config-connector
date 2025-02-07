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


// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig
type CertificateIssuanceConfig struct {
	// A user-defined name of the certificate issuance config.
	//  CertificateIssuanceConfig names must be unique globally and match pattern
	//  `projects/*/locations/*/certificateIssuanceConfigs/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.name
	Name *string `json:"name,omitempty"`

	// Set of labels associated with a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.description
	Description *string `json:"description,omitempty"`

	// Required. The CA that issues the workload certificate. It includes the CA
	//  address, type, authentication to CA service, etc.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.certificate_authority_config
	CertificateAuthorityConfig *CertificateIssuanceConfig_CertificateAuthorityConfig `json:"certificateAuthorityConfig,omitempty"`

	// Required. Workload certificate lifetime requested.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.lifetime
	Lifetime *string `json:"lifetime,omitempty"`

	// Required. Specifies the percentage of elapsed time of the certificate
	//  lifetime to wait before renewing the certificate. Must be a number between
	//  1-99, inclusive.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.rotation_window_percentage
	RotationWindowPercentage *int32 `json:"rotationWindowPercentage,omitempty"`

	// Required. The key algorithm to use when generating the private key.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.key_algorithm
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig
type CertificateIssuanceConfig_CertificateAuthorityConfig struct {
	// Defines a CertificateAuthorityServiceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig.certificate_authority_service_config
	CertificateAuthorityServiceConfig *CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig `json:"certificateAuthorityServiceConfig,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig.CertificateAuthorityServiceConfig
type CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig struct {
	// Required. A CA pool resource used to issue a certificate.
	//  The CA pool string has a relative resource path following the form
	//  "projects/{project}/locations/{location}/caPools/{ca_pool}".
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig.CertificateAuthorityServiceConfig.ca_pool
	CaPool *string `json:"caPool,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig
type CertificateIssuanceConfigObservedState struct {
	// Output only. The creation timestamp of a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
