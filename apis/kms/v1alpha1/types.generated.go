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
// krm.group: kms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.kms.v1
// resource: KMSImportJob:ImportJob

package v1alpha1

// +kcc:proto=google.cloud.kms.v1.ImportJob.WrappingPublicKey
type ImportJob_WrappingPublicKey struct {
	// The public key, encoded in PEM format. For more information, see the [RFC
	//  7468](https://tools.ietf.org/html/rfc7468) sections for [General
	//  Considerations](https://tools.ietf.org/html/rfc7468#section-2) and
	//  [Textual Encoding of Subject Public Key Info]
	//  (https://tools.ietf.org/html/rfc7468#section-13).
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.WrappingPublicKey.pem
	Pem *string `json:"pem,omitempty"`
}

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

// +kcc:proto=google.cloud.kms.v1.KeyOperationAttestation
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
