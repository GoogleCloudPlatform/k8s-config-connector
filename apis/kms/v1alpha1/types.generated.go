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

// +kcc:proto=google.cloud.kms.v1.ImportJob
type ImportJob struct {

	// Required. Immutable. The wrapping method to be used for incoming key
	//  material.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.import_method
	ImportMethod *string `json:"importMethod,omitempty"`

	// Required. Immutable. The protection level of the
	//  [ImportJob][google.cloud.kms.v1.ImportJob]. This must match the
	//  [protection_level][google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level]
	//  of the [version_template][google.cloud.kms.v1.CryptoKey.version_template]
	//  on the [CryptoKey][google.cloud.kms.v1.CryptoKey] you attempt to import
	//  into.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.protection_level
	ProtectionLevel *string `json:"protectionLevel,omitempty"`
}

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

// +kcc:proto=google.cloud.kms.v1.ImportJob
type ImportJobObservedState struct {
	// Output only. The resource name for this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] in the format
	//  `projects/*/locations/*/keyRings/*/importJobs/*`.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] was created.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this [ImportJob][google.cloud.kms.v1.ImportJob]'s key
	//  material was generated.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.generate_time
	GenerateTime *string `json:"generateTime,omitempty"`

	// Output only. The time at which this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] is scheduled for expiration and
	//  can no longer be used to import key material.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. The time this [ImportJob][google.cloud.kms.v1.ImportJob]
	//  expired. Only present if [state][google.cloud.kms.v1.ImportJob.state] is
	//  [EXPIRED][google.cloud.kms.v1.ImportJob.ImportJobState.EXPIRED].
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.expire_event_time
	ExpireEventTime *string `json:"expireEventTime,omitempty"`

	// Output only. The current state of the
	//  [ImportJob][google.cloud.kms.v1.ImportJob], indicating if it can be used.
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.state
	State *string `json:"state,omitempty"`

	// Output only. The public key with which to wrap key material prior to
	//  import. Only returned if [state][google.cloud.kms.v1.ImportJob.state] is
	//  [ACTIVE][google.cloud.kms.v1.ImportJob.ImportJobState.ACTIVE].
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.public_key
	PublicKey *ImportJob_WrappingPublicKey `json:"publicKey,omitempty"`

	// Output only. Statement that was generated and signed by the key creator
	//  (for example, an HSM) at key creation time. Use this statement to verify
	//  attributes of the key as stored on the HSM, independently of Google.
	//  Only present if the chosen
	//  [ImportMethod][google.cloud.kms.v1.ImportJob.ImportMethod] is one with a
	//  protection level of [HSM][google.cloud.kms.v1.ProtectionLevel.HSM].
	// +kcc:proto:field=google.cloud.kms.v1.ImportJob.attestation
	Attestation *KeyOperationAttestation `json:"attestation,omitempty"`
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
