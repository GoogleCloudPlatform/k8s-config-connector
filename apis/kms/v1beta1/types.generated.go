// Copyright 2024 Google LLC
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

package v1beta1

import refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

// +kcc:proto=google.cloud.kms.v1.CryptoKey
type CryptoKey struct {
	// Output only. The resource name for this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] in the format
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
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
	Primary *CryptoKeyVersion `json:"primary,omitempty"`

	// Immutable. The immutable purpose of this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey].
	Purpose *string `json:"purpose,omitempty"`

	// Output only. The time at which this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] was created.
	CreateTime *string `json:"createTime,omitempty"`

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
	RotationPeriod *string `json:"rotationPeriod,omitempty"`

	// A template describing settings for new
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] instances. The
	//  properties of new [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]
	//  instances created by either
	//  [CreateCryptoKeyVersion][google.cloud.kms.v1.KeyManagementService.CreateCryptoKeyVersion]
	//  or auto-rotation are controlled by this template.
	VersionTemplate *CryptoKeyVersionTemplate `json:"versionTemplate,omitempty"`

	// Labels with user-defined metadata. For more information, see
	//  [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. Whether this key may contain imported versions only.
	ImportOnly *bool `json:"importOnly,omitempty"`

	// Immutable. The period of time that versions of this key spend in the
	//  [DESTROY_SCHEDULED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROY_SCHEDULED]
	//  state before transitioning to
	//  [DESTROYED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROYED].
	//  If not specified at creation time, the default duration is 24 hours.
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty"`

	// Immutable. The resource name of the backend environment where the key
	//  material for all [CryptoKeyVersions][google.cloud.kms.v1.CryptoKeyVersion]
	//  associated with this [CryptoKey][google.cloud.kms.v1.CryptoKey] reside and
	//  where all related cryptographic operations are performed. Only applicable
	//  if [CryptoKeyVersions][google.cloud.kms.v1.CryptoKeyVersion] have a
	//  [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] of
	//  [EXTERNAL_VPC][CryptoKeyVersion.ProtectionLevel.EXTERNAL_VPC], with the
	//  resource name in the format `projects/*/locations/*/ekmConnections/*`.
	//  Note, this list is non-exhaustive and may apply to additional
	//  [ProtectionLevels][google.cloud.kms.v1.ProtectionLevel] in the future.
	CryptoKeyBackend *string `json:"cryptoKeyBackend,omitempty"`

	// Optional. The policy used for Key Access Justifications Policy Enforcement.
	//  If this field is present and this key is enrolled in Key Access
	//  Justifications Policy Enforcement, the policy will be evaluated in encrypt,
	//  decrypt, and sign operations, and the operation will fail if rejected by
	//  the policy. The policy is defined by specifying zero or more allowed
	//  justification codes.
	//  https://cloud.google.com/assured-workloads/key-access-justifications/docs/justification-codes
	//  By default, this field is absent, and all justification codes are allowed.
	KeyAccessJustificationsPolicy *KeyAccessJustificationsPolicy `json:"keyAccessJustificationsPolicy,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.CryptoKeyVersion
type CryptoKeyVersion struct {
	// Output only. The resource name for this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] in the format
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
	Name *string `json:"name,omitempty"`

	// The current state of the
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion].
	State *string `json:"state,omitempty"`

	// Output only. The [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel]
	//  describing how crypto operations are performed with this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion].
	ProtectionLevel *string `json:"protectionLevel,omitempty"`

	// Output only. The
	//  [CryptoKeyVersionAlgorithm][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionAlgorithm]
	//  that this [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]
	//  supports.
	Algorithm *string `json:"algorithm,omitempty"`

	// Output only. Statement that was generated and signed by the HSM at key
	//  creation time. Use this statement to verify attributes of the key as stored
	//  on the HSM, independently of Google. Only provided for key versions with
	//  [protection_level][google.cloud.kms.v1.CryptoKeyVersion.protection_level]
	//  [HSM][google.cloud.kms.v1.ProtectionLevel.HSM].
	Attestation *KeyOperationAttestation `json:"attestation,omitempty"`

	// Output only. The time at which this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]'s key material was
	//  generated.
	GenerateTime *string `json:"generateTime,omitempty"`

	// Output only. The time this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]'s key material is
	//  scheduled for destruction. Only present if
	//  [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [DESTROY_SCHEDULED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROY_SCHEDULED].
	DestroyTime *string `json:"destroyTime,omitempty"`

	// Output only. The time this CryptoKeyVersion's key material was
	//  destroyed. Only present if
	//  [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [DESTROYED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROYED].
	DestroyEventTime *string `json:"destroyEventTime,omitempty"`

	// Output only. The name of the [ImportJob][google.cloud.kms.v1.ImportJob]
	//  used in the most recent import of this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]. Only present if
	//  the underlying key material was imported.
	ImportJob *string `json:"importJob,omitempty"`

	// Output only. The time at which this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]'s key material was
	//  most recently imported.
	ImportTime *string `json:"importTime,omitempty"`

	// Output only. The root cause of the most recent import failure. Only present
	//  if [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [IMPORT_FAILED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.IMPORT_FAILED].
	ImportFailureReason *string `json:"importFailureReason,omitempty"`

	// Output only. The root cause of the most recent generation failure. Only
	//  present if [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [GENERATION_FAILED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.GENERATION_FAILED].
	GenerationFailureReason *string `json:"generationFailureReason,omitempty"`

	// Output only. The root cause of the most recent external destruction
	//  failure. Only present if
	//  [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	//  [EXTERNAL_DESTRUCTION_FAILED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.EXTERNAL_DESTRUCTION_FAILED].
	ExternalDestructionFailureReason *string `json:"externalDestructionFailureReason,omitempty"`

	// ExternalProtectionLevelOptions stores a group of additional fields for
	//  configuring a [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] that
	//  are specific to the
	//  [EXTERNAL][google.cloud.kms.v1.ProtectionLevel.EXTERNAL] protection level
	//  and [EXTERNAL_VPC][google.cloud.kms.v1.ProtectionLevel.EXTERNAL_VPC]
	//  protection levels.
	ExternalProtectionLevelOptions *ExternalProtectionLevelOptions `json:"externalProtectionLevelOptions,omitempty"`

	// Output only. Whether or not this key version is eligible for reimport, by
	//  being specified as a target in
	//  [ImportCryptoKeyVersionRequest.crypto_key_version][google.cloud.kms.v1.ImportCryptoKeyVersionRequest.crypto_key_version].
	ReimportEligible *bool `json:"reimportEligible,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.CryptoKeyVersionTemplate
type CryptoKeyVersionTemplate struct {
	// [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] to use when creating
	//  a [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] based on this
	//  template. Immutable. Defaults to
	//  [SOFTWARE][google.cloud.kms.v1.ProtectionLevel.SOFTWARE].
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
	Algorithm *string `json:"algorithm,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.ExternalProtectionLevelOptions
type ExternalProtectionLevelOptions struct {
	// The URI for an external resource that this
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] represents.
	ExternalKeyUri *string `json:"externalKeyUri,omitempty"`

	// The path to the external key material on the EKM when using
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] e.g., "v0/my/key". Set
	//  this field instead of external_key_uri when using an
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection].
	EkmConnectionKeyPath *string `json:"ekmConnectionKeyPath,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.ImportJob
type ImportJob struct {
	// Output only. The resource name for this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] in the format
	//  `projects/*/locations/*/keyRings/*/importJobs/*`.
	Name *string `json:"name,omitempty"`

	// Required. Immutable. The wrapping method to be used for incoming key
	//  material.
	ImportMethod *string `json:"importMethod,omitempty"`

	// Required. Immutable. The protection level of the
	//  [ImportJob][google.cloud.kms.v1.ImportJob]. This must match the
	//  [protection_level][google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level]
	//  of the [version_template][google.cloud.kms.v1.CryptoKey.version_template]
	//  on the [CryptoKey][google.cloud.kms.v1.CryptoKey] you attempt to import
	//  into.
	ProtectionLevel *string `json:"protectionLevel,omitempty"`

	// Output only. The time at which this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this [ImportJob][google.cloud.kms.v1.ImportJob]'s key
	//  material was generated.
	GenerateTime *string `json:"generateTime,omitempty"`

	// Output only. The time at which this
	//  [ImportJob][google.cloud.kms.v1.ImportJob] is scheduled for expiration and
	//  can no longer be used to import key material.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. The time this [ImportJob][google.cloud.kms.v1.ImportJob]
	//  expired. Only present if [state][google.cloud.kms.v1.ImportJob.state] is
	//  [EXPIRED][google.cloud.kms.v1.ImportJob.ImportJobState.EXPIRED].
	ExpireEventTime *string `json:"expireEventTime,omitempty"`

	// Output only. The current state of the
	//  [ImportJob][google.cloud.kms.v1.ImportJob], indicating if it can be used.
	State *string `json:"state,omitempty"`

	// Output only. The public key with which to wrap key material prior to
	//  import. Only returned if [state][google.cloud.kms.v1.ImportJob.state] is
	//  [ACTIVE][google.cloud.kms.v1.ImportJob.ImportJobState.ACTIVE].
	PublicKey *ImportJob_WrappingPublicKey `json:"publicKey,omitempty"`

	// Output only. Statement that was generated and signed by the key creator
	//  (for example, an HSM) at key creation time. Use this statement to verify
	//  attributes of the key as stored on the HSM, independently of Google.
	//  Only present if the chosen
	//  [ImportMethod][google.cloud.kms.v1.ImportJob.ImportMethod] is one with a
	//  protection level of [HSM][google.cloud.kms.v1.ProtectionLevel.HSM].
	Attestation *KeyOperationAttestation `json:"attestation,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.ImportJob.WrappingPublicKey
type ImportJob_WrappingPublicKey struct {
	// The public key, encoded in PEM format. For more information, see the [RFC
	//  7468](https://tools.ietf.org/html/rfc7468) sections for [General
	//  Considerations](https://tools.ietf.org/html/rfc7468#section-2) and
	//  [Textual Encoding of Subject Public Key Info]
	//  (https://tools.ietf.org/html/rfc7468#section-13).
	Pem *string `json:"pem,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.KeyAccessJustificationsPolicy
type KeyAccessJustificationsPolicy struct {
	// The list of allowed reasons for access to a
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey]. Zero allowed access reasons
	//  means all encrypt, decrypt, and sign operations for the
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] associated with this policy will
	//  fail.
	AllowedAccessReasons []string `json:"allowedAccessReasons,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.KeyOperationAttestation
type KeyOperationAttestation struct {
	// Output only. The format of the attestation data.
	Format *string `json:"format,omitempty"`

	// Output only. The attestation data provided by the HSM when the key
	//  operation was performed.
	Content *[]byte `json:"content,omitempty"`

	// Output only. The certificate chains needed to validate the attestation
	CertChains *KeyOperationAttestation_CertificateChains `json:"certChains,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.KeyOperationAttestation.CertificateChains
type KeyOperationAttestation_CertificateChains struct {
	// Cavium certificate chain corresponding to the attestation.
	CaviumCerts []string `json:"caviumCerts,omitempty"`

	// Google card certificate chain corresponding to the attestation.
	GoogleCardCerts []string `json:"googleCardCerts,omitempty"`

	// Google partition certificate chain corresponding to the attestation.
	GooglePartitionCerts []string `json:"googlePartitionCerts,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.PublicKey
type PublicKey struct {
	// The public key, encoded in PEM format. For more information, see the
	//  [RFC 7468](https://tools.ietf.org/html/rfc7468) sections for
	//  [General Considerations](https://tools.ietf.org/html/rfc7468#section-2) and
	//  [Textual Encoding of Subject Public Key Info]
	//  (https://tools.ietf.org/html/rfc7468#section-13).
	Pem *string `json:"pem,omitempty"`

	// The
	//  [Algorithm][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionAlgorithm]
	//  associated with this key.
	Algorithm *string `json:"algorithm,omitempty"`

	// Integrity verification field. A CRC32C checksum of the returned
	//  [PublicKey.pem][google.cloud.kms.v1.PublicKey.pem]. An integrity check of
	//  [PublicKey.pem][google.cloud.kms.v1.PublicKey.pem] can be performed by
	//  computing the CRC32C checksum of
	//  [PublicKey.pem][google.cloud.kms.v1.PublicKey.pem] and comparing your
	//  results to this field. Discard the response in case of non-matching
	//  checksum values, and perform a limited number of retries. A persistent
	//  mismatch may indicate an issue in your computation of the CRC32C checksum.
	//  Note: This field is defined as int64 for reasons of compatibility across
	//  different languages. However, it is a non-negative integer, which will
	//  never exceed 2^32-1, and can be safely downconverted to uint32 in languages
	//  that support this type.
	//
	//  NOTE: This field is in Beta.
	PemCrc32c *int64 `json:"pemCrc32c,omitempty"`

	// The [name][google.cloud.kms.v1.CryptoKeyVersion.name] of the
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] public key.
	//  Provided here for verification.
	//
	//  NOTE: This field is in Beta.
	Name *string `json:"name,omitempty"`

	// The [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] of the
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] public key.
	ProtectionLevel *string `json:"protectionLevel,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.Certificate
type Certificate struct {
	// Required. The raw certificate bytes in DER format.
	RawDer *[]byte `json:"rawDer,omitempty"`

	// Output only. True if the certificate was parsed successfully.
	Parsed *bool `json:"parsed,omitempty"`

	// Output only. The issuer distinguished name in RFC 2253 format. Only present
	//  if [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	Issuer *string `json:"issuer,omitempty"`

	// Output only. The subject distinguished name in RFC 2253 format. Only
	//  present if [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	Subject *string `json:"subject,omitempty"`

	// Output only. The subject Alternative DNS names. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	SubjectAlternativeDnsNames []string `json:"subjectAlternativeDnsNames,omitempty"`

	// Output only. The certificate is not valid before this time. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	NotBeforeTime *string `json:"notBeforeTime,omitempty"`

	// Output only. The certificate is not valid after this time. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	NotAfterTime *string `json:"notAfterTime,omitempty"`

	// Output only. The certificate serial number as a hex string. Only present if
	//  [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Output only. The SHA-256 certificate fingerprint as a hex string. Only
	//  present if [parsed][google.cloud.kms.v1.Certificate.parsed] is true.
	Sha256Fingerprint *string `json:"sha256Fingerprint,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConfig
type EkmConfig struct {
	// Output only. The resource name for the
	//  [EkmConfig][google.cloud.kms.v1.EkmConfig] in the format
	//  `projects/*/locations/*/ekmConfig`.
	Name *string `json:"name,omitempty"`

	// Optional. Resource name of the default
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection]. Setting this field to
	//  the empty string removes the default.
	DefaultEkmConnection *string `json:"defaultEkmConnection,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection
type EkmConnection struct {
	// Output only. The resource name for the
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] in the format
	//  `projects/*/locations/*/ekmConnections/*`.
	Name *string `json:"name,omitempty"`

	// Output only. The time at which the
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] was created.
	CreateTime *string `json:"createTime,omitempty"`

	// A list of
	//  [ServiceResolvers][google.cloud.kms.v1.EkmConnection.ServiceResolver] where
	//  the EKM can be reached. There should be one ServiceResolver per EKM
	//  replica. Currently, only a single
	//  [ServiceResolver][google.cloud.kms.v1.EkmConnection.ServiceResolver] is
	//  supported.
	ServiceResolvers []EkmConnection_ServiceResolver `json:"serviceResolvers,omitempty"`

	// Optional. Etag of the currently stored
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection].
	Etag *string `json:"etag,omitempty"`

	// Optional. Describes who can perform control plane operations on the EKM. If
	//  unset, this defaults to
	//  [MANUAL][google.cloud.kms.v1.EkmConnection.KeyManagementMode.MANUAL].
	KeyManagementMode *string `json:"keyManagementMode,omitempty"`

	// Optional. Identifies the EKM Crypto Space that this
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] maps to. Note: This
	//  field is required if
	//  [KeyManagementMode][google.cloud.kms.v1.EkmConnection.KeyManagementMode] is
	//  [CLOUD_KMS][google.cloud.kms.v1.EkmConnection.KeyManagementMode.CLOUD_KMS].
	CryptoSpacePath *string `json:"cryptoSpacePath,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection.ServiceResolver
type EkmConnection_ServiceResolver struct {
	// Required. The resource name of the Service Directory service pointing to
	//  an EKM replica, in the format
	//  `projects/*/locations/*/namespaces/*/services/*`.
	ServiceDirectoryService *string `json:"serviceDirectoryService,omitempty"`

	// Optional. The filter applied to the endpoints of the resolved service. If
	//  no filter is specified, all endpoints will be considered. An endpoint
	//  will be chosen arbitrarily from the filtered list for each request.
	//
	//  For endpoint filter syntax and examples, see
	//  https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.
	EndpointFilter *string `json:"endpointFilter,omitempty"`

	// Required. The hostname of the EKM replica used at TLS and HTTP layers.
	Hostname *string `json:"hostname,omitempty"`

	// Required. A list of leaf server certificates used to authenticate HTTPS
	//  connections to the EKM replica. Currently, a maximum of 10
	//  [Certificate][google.cloud.kms.v1.Certificate] is supported.
	ServerCertificates []Certificate `json:"serverCertificates,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.CreateKeyHandleMetadata
type CreateKeyHandleMetadata struct {
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

// +kcc:proto=google.cloud.kms.v1.Digest
type Digest struct {
	// A message digest produced with the SHA-256 algorithm.
	Sha256 *[]byte `json:"sha256,omitempty"`

	// A message digest produced with the SHA-384 algorithm.
	Sha384 *[]byte `json:"sha384,omitempty"`

	// A message digest produced with the SHA-512 algorithm.
	Sha512 *[]byte `json:"sha512,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.LocationMetadata
type LocationMetadata struct {
	// Indicates whether [CryptoKeys][google.cloud.kms.v1.CryptoKey] with
	//  [protection_level][google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level]
	//  [HSM][google.cloud.kms.v1.ProtectionLevel.HSM] can be created in this
	//  location.
	HsmAvailable *bool `json:"hsmAvailable,omitempty"`

	// Indicates whether [CryptoKeys][google.cloud.kms.v1.CryptoKey] with
	//  [protection_level][google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level]
	//  [EXTERNAL][google.cloud.kms.v1.ProtectionLevel.EXTERNAL] can be created in
	//  this location.
	EkmAvailable *bool `json:"ekmAvailable,omitempty"`
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
