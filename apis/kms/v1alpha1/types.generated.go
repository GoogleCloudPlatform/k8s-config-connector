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


// +kcc:proto=google.cloud.kms.v1.CryptoKeyVersion
type CryptoKeyVersion struct {

	// The current state of the
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.state
	State *string `json:"state,omitempty"`

	// ExternalProtectionLevelOptions stores a group of additional fields for
	//  configuring a [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] that
	//  are specific to the
	//  [EXTERNAL][google.cloud.kms.v1.ProtectionLevel.EXTERNAL] protection level
	//  and [EXTERNAL_VPC][google.cloud.kms.v1.ProtectionLevel.EXTERNAL_VPC]
	//  protection levels.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.external_protection_level_options
	ExternalProtectionLevelOptions *ExternalProtectionLevelOptions `json:"externalProtectionLevelOptions,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.ExternalProtectionLevelOptions
type ExternalProtectionLevelOptions struct {
	// The URI for an external resource that this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] represents.
	// +kcc:proto:field=google.cloud.kms.v1.ExternalProtectionLevelOptions.external_key_uri
	ExternalKeyURI *string `json:"externalKeyURI,omitempty"`

	// The path to the external key material on the EKM when using
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] e.g., "v0/my/key". Set
	//  this field instead of external_key_uri when using an
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection].
	// +kcc:proto:field=google.cloud.kms.v1.ExternalProtectionLevelOptions.ekm_connection_key_path
	EkmConnectionKeyPath *string `json:"ekmConnectionKeyPath,omitempty"`
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

// +kcc:proto=google.cloud.kms.v1.CryptoKeyVersion
type CryptoKeyVersionObservedState struct {
	// Output only. The resource name for this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] in the format
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.name
	Name *string `json:"name,omitempty"`

	// Output only. The [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel]
	//  describing how crypto operations are performed with this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.protection_level
	ProtectionLevel *string `json:"protectionLevel,omitempty"`

	// Output only. The
	//  [CryptoKeyVersionAlgorithm][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionAlgorithm]
	//  that this [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]
	//  supports.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.algorithm
	Algorithm *string `json:"algorithm,omitempty"`

	// Output only. Statement that was generated and signed by the HSM at key
	//  creation time. Use this statement to verify attributes of the key as stored
	//  on the HSM, independently of Google. Only provided for key versions with
	//  [protection_level][google.cloud.kms.v1.CryptoKeyVersion.protection_level]
	//  [HSM][google.cloud.kms.v1.ProtectionLevel.HSM].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.attestation
	Attestation *KeyOperationAttestation `json:"attestation,omitempty"`

	// Output only. The time at which this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] was created.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]'s key material was
	//  generated.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.generate_time
	GenerateTime *string `json:"generateTime,omitempty"`

	// Output only. The time this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]'s key material is
	//  scheduled for destruction. Only present if
	//  [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [DESTROY_SCHEDULED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROY_SCHEDULED].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.destroy_time
	DestroyTime *string `json:"destroyTime,omitempty"`

	// Output only. The time this CryptoKeyVersion's key material was
	//  destroyed. Only present if
	//  [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [DESTROYED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROYED].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.destroy_event_time
	DestroyEventTime *string `json:"destroyEventTime,omitempty"`

	// Output only. The name of the [ImportJob][google.cloud.kms.v1.ImportJob]
	//  used in the most recent import of this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]. Only present if
	//  the underlying key material was imported.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.import_job
	ImportJob *string `json:"importJob,omitempty"`

	// Output only. The time at which this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]'s key material was
	//  most recently imported.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.import_time
	ImportTime *string `json:"importTime,omitempty"`

	// Output only. The root cause of the most recent import failure. Only present
	//  if [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [IMPORT_FAILED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.IMPORT_FAILED].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.import_failure_reason
	ImportFailureReason *string `json:"importFailureReason,omitempty"`

	// Output only. The root cause of the most recent generation failure. Only
	//  present if [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [GENERATION_FAILED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.GENERATION_FAILED].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.generation_failure_reason
	GenerationFailureReason *string `json:"generationFailureReason,omitempty"`

	// Output only. The root cause of the most recent external destruction
	//  failure. Only present if
	//  [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [EXTERNAL_DESTRUCTION_FAILED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.EXTERNAL_DESTRUCTION_FAILED].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.external_destruction_failure_reason
	ExternalDestructionFailureReason *string `json:"externalDestructionFailureReason,omitempty"`

	// Output only. Whether or not this key version is eligible for reimport, by
	//  being specified as a target in
	//  [ImportCryptoKeyVersionRequest.crypto_key_version][google.cloud.kms.v1.ImportCryptoKeyVersionRequest.crypto_key_version].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.reimport_eligible
	ReimportEligible *bool `json:"reimportEligible,omitempty"`
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
