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


// +kcc:proto=google.cloud.kms.v1.CryptoKey
type CryptoKey struct {

	// Immutable. The immutable purpose of this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.purpose
	Purpose *string `json:"purpose,omitempty"`

	// At [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time],
	//  the Key Management Service will automatically:
	//
	//  1. Create a new version of this [CryptoKey][google.cloud.kms.v1.CryptoKey].
	//  2. Mark the new version as primary.
	//
	//  Key rotations performed manually via
	//  [CreateCryptoKeyVersion][google.cloud.kms.v1.KeyManagementService.CreateCryptoKeyVersion]
	//  and
	//  [UpdateCryptoKeyPrimaryVersion][google.cloud.kms.v1.KeyManagementService.UpdateCryptoKeyPrimaryVersion]
	//  do not affect
	//  [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time].
	//
	//  Keys with [purpose][google.cloud.kms.v1.CryptoKey.purpose]
	//  [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT]
	//  support automatic rotation. For other keys, this field must be omitted.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.next_rotation_time
	NextRotationTime *string `json:"nextRotationTime,omitempty"`

	// [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time]
	//  will be advanced by this period when the service automatically rotates a
	//  key. Must be at least 24 hours and at most 876,000 hours.
	//
	//  If [rotation_period][google.cloud.kms.v1.CryptoKey.rotation_period] is
	//  set,
	//  [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time]
	//  must also be set.
	//
	//  Keys with [purpose][google.cloud.kms.v1.CryptoKey.purpose]
	//  [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT]
	//  support automatic rotation. For other keys, this field must be omitted.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.rotation_period
	RotationPeriod *string `json:"rotationPeriod,omitempty"`

	// A template describing settings for new
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] instances. The
	//  properties of new [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]
	//  instances created by either
	//  [CreateCryptoKeyVersion][google.cloud.kms.v1.KeyManagementService.CreateCryptoKeyVersion]
	//  or auto-rotation are controlled by this template.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.version_template
	VersionTemplate *CryptoKeyVersionTemplate `json:"versionTemplate,omitempty"`

	// Labels with user-defined metadata. For more information, see
	//  [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. Whether this key may contain imported versions only.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.import_only
	ImportOnly *bool `json:"importOnly,omitempty"`

	// Immutable. The period of time that versions of this key spend in the
	//  [DESTROY_SCHEDULED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROY_SCHEDULED]
	//  state before transitioning to
	//  [DESTROYED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROYED].
	//  If not specified at creation time, the default duration is 30 days.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.destroy_scheduled_duration
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty"`

	// Immutable. The resource name of the backend environment where the key
	//  material for all [CryptoKeyVersions][google.cloud.kms.v1.CryptoKeyVersion]
	//  associated with this [CryptoKey][google.cloud.kms.v1.CryptoKey] reside and
	//  where all related cryptographic operations are performed. Only applicable
	//  if [CryptoKeyVersions][google.cloud.kms.v1.CryptoKeyVersion] have a
	//  [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] of
	//  [EXTERNAL_VPC][google.cloud.kms.v1.ProtectionLevel.EXTERNAL_VPC], with the
	//  resource name in the format `projects/*/locations/*/ekmConnections/*`.
	//  Note, this list is non-exhaustive and may apply to additional
	//  [ProtectionLevels][google.cloud.kms.v1.ProtectionLevel] in the future.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.crypto_key_backend
	CryptoKeyBackend *string `json:"cryptoKeyBackend,omitempty"`

	// Optional. The policy used for Key Access Justifications Policy Enforcement.
	//  If this field is present and this key is enrolled in Key Access
	//  Justifications Policy Enforcement, the policy will be evaluated in encrypt,
	//  decrypt, and sign operations, and the operation will fail if rejected by
	//  the policy. The policy is defined by specifying zero or more allowed
	//  justification codes.
	//  https://cloud.google.com/assured-workloads/key-access-justifications/docs/justification-codes
	//  By default, this field is absent, and all justification codes are allowed.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.key_access_justifications_policy
	KeyAccessJustificationsPolicy *KeyAccessJustificationsPolicy `json:"keyAccessJustificationsPolicy,omitempty"`
}

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

// +kcc:proto=google.cloud.kms.v1.CryptoKeyVersionTemplate
type CryptoKeyVersionTemplate struct {
	// [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] to use when creating
	//  a [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] based on this
	//  template. Immutable. Defaults to
	//  [SOFTWARE][google.cloud.kms.v1.ProtectionLevel.SOFTWARE].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level
	ProtectionLevel *string `json:"protectionLevel,omitempty"`

	// Required.
	//  [Algorithm][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionAlgorithm]
	//  to use when creating a
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] based on this
	//  template.
	//
	//  For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION is implied if both
	//  this field is omitted and
	//  [CryptoKey.purpose][google.cloud.kms.v1.CryptoKey.purpose] is
	//  [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT].
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm
	Algorithm *string `json:"algorithm,omitempty"`
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

// +kcc:proto=google.cloud.kms.v1.KeyAccessJustificationsPolicy
type KeyAccessJustificationsPolicy struct {
	// The list of allowed reasons for access to a
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey]. Zero allowed access reasons
	//  means all encrypt, decrypt, and sign operations for the
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] associated with this policy will
	//  fail.
	// +kcc:proto:field=google.cloud.kms.v1.KeyAccessJustificationsPolicy.allowed_access_reasons
	AllowedAccessReasons []string `json:"allowedAccessReasons,omitempty"`
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

// +kcc:proto=google.cloud.kms.v1.CryptoKey
type CryptoKeyObservedState struct {
	// Output only. The resource name for this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] in the format
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.name
	Name *string `json:"name,omitempty"`

	// Output only. A copy of the "primary"
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] that will be used
	//  by [Encrypt][google.cloud.kms.v1.KeyManagementService.Encrypt] when this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] is given in
	//  [EncryptRequest.name][google.cloud.kms.v1.EncryptRequest.name].
	//
	//  The [CryptoKey][google.cloud.kms.v1.CryptoKey]'s primary version can be
	//  updated via
	//  [UpdateCryptoKeyPrimaryVersion][google.cloud.kms.v1.KeyManagementService.UpdateCryptoKeyPrimaryVersion].
	//
	//  Keys with [purpose][google.cloud.kms.v1.CryptoKey.purpose]
	//  [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT]
	//  may have a primary. For other keys, this field will be omitted.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.primary
	Primary *CryptoKeyVersion `json:"primary,omitempty"`

	// Output only. The time at which this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] was created.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.create_time
	CreateTime *string `json:"createTime,omitempty"`
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
