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
// krm.group: kms
// krm.version: v1beta1
// proto.service: google.cloud.kms.v1
// resource: KMSAutokeyConfig:AutokeyConfig
// resource: KMSImportJob:ImportJob
// resource: KMSKeyHandle:KeyHandle

package v1beta1

import refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

// +kcc:proto=google.cloud.kms.v1.KeyOperationAttestation
type KeyOperationAttestation struct {
}

// +kcc:proto=google.cloud.kms.v1.KeyOperationAttestation.CertificateChains
type KeyOperationAttestation_CertificateChains struct {
	// Cavium certificate chain corresponding to the attestation.
	// +kcc:proto:field=google.cloud.kms.v1.KeyOperationAttestation.CertificateChains.cavium_certs
	CaviumCerts []string `json:"caviumCerts,omitempty"`

	// Google card certificate chain corresponding to the attestation.
	// +kcc:proto:field=google.cloud.kms.v1.KeyOperationAttestation.CertificateChains.google_card_certs
	GoogleCardCerts []string `json:"googleCardCerts,omitempty"`

	// Google partition certificate chain corresponding to the attestation.
	// +kcc:proto:field=google.cloud.kms.v1.KeyOperationAttestation.CertificateChains.google_partition_certs
	GooglePartitionCerts []string `json:"googlePartitionCerts,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.kms.v1.KeyOperationAttestation
type KeyOperationAttestationObservedState struct {
	// Output only. The format of the attestation data.
	// +kcc:proto:field=google.cloud.kms.v1.KeyOperationAttestation.format
	Format *string `json:"format,omitempty"`

	// Output only. The attestation data provided by the HSM when the key
	//  operation was performed.
	// +kcc:proto:field=google.cloud.kms.v1.KeyOperationAttestation.content
	Content []byte `json:"content,omitempty"`

	// Output only. The certificate chains needed to validate the attestation
	// +kcc:proto:field=google.cloud.kms.v1.KeyOperationAttestation.cert_chains
	CertChains *KeyOperationAttestation_CertificateChains `json:"certChains,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.AutokeyConfig
type AutokeyConfig struct {
	// Identifier. Name of the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig]
	//  resource, e.g. `folders/{FOLDER_NUMBER}/autokeyConfig`.
	Name *string `json:"name,omitempty"`

	// Optional. Name of the key project, e.g. `projects/{PROJECT_ID}` or
	//  `projects/{PROJECT_NUMBER}`, where Cloud KMS Autokey will provision a new
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] when a
	//  [KeyHandle][google.cloud.kms.v1.KeyHandle] is created. On
	//  [UpdateAutokeyConfig][google.cloud.kms.v1.AutokeyAdmin.UpdateAutokeyConfig],
	//  the caller will require `cloudkms.cryptoKeys.setIamPolicy` permission on
	//  this key project. Once configured, for Cloud KMS Autokey to function
	//  properly, this key project must have the Cloud KMS API activated and the
	//  Cloud KMS Service Agent for this key project must be granted the
	//  `cloudkms.admin` role (or pertinent permissions). A request with an empty
	//  key project field will clear the configuration.
	KeyProject *refs.ProjectRef `json:"keyProject,omitempty"`

	// Output only. The state for the AutokeyConfig.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.KeyHandle
type KeyHandle struct {
	// Identifier. Name of the [KeyHandle][google.cloud.kms.v1.KeyHandle]
	//  resource, e.g.
	//  `projects/{PROJECT_ID}/locations/{LOCATION}/keyHandles/{KEY_HANDLE_ID}`.
	Name *string `json:"name,omitempty"`

	// Output only. Name of a [CryptoKey][google.cloud.kms.v1.CryptoKey] that has
	//  been provisioned for Customer Managed Encryption Key (CMEK) use in the
	//  [KeyHandle][google.cloud.kms.v1.KeyHandle] project and location for the
	//  requested resource type. The [CryptoKey][google.cloud.kms.v1.CryptoKey]
	//  project will reflect the value configured in the
	//  [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] on the resource
	//  project's ancestor folder at the time of the
	//  [KeyHandle][google.cloud.kms.v1.KeyHandle] creation. If more than one
	//  ancestor folder has a configured
	//  [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig], the nearest of these
	//  configurations is used.
	KmsKey *string `json:"kmsKey,omitempty"`

	// Required. Indicates the resource type that the resulting
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] is meant to protect, e.g.
	//  `{SERVICE}.googleapis.com/{TYPE}`. See documentation for supported resource
	//  types.
	ResourceTypeSelector *string `json:"resourceTypeSelector,omitempty"`
}
