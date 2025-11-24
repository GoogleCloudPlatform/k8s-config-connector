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
// krm.group: certificatemanager.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.certificatemanager.v1
// resource: CertificateManagerCertificateIssuanceConfig:CertificateIssuanceConfig

package v1alpha1

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig
type CertificateIssuanceConfig_CertificateAuthorityConfig struct {
	// Defines a CertificateAuthorityServiceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig.certificate_authority_service_config
	CertificateAuthorityServiceConfig *CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig `json:"certificateAuthorityServiceConfig,omitempty"`
}
