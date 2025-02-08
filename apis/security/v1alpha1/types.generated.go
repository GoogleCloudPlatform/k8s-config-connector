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


// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority
type CertificateAuthority struct {

	// Required. Immutable. The [Type][google.cloud.security.privateca.v1beta1.CertificateAuthority.Type] of this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.type
	Type *string `json:"type,omitempty"`

	// Required. Immutable. The [Tier][google.cloud.security.privateca.v1beta1.CertificateAuthority.Tier] of this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.tier
	Tier *string `json:"tier,omitempty"`

	// Required. Immutable. The config used to create a self-signed X.509 certificate or CSR.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.config
	Config *CertificateConfig `json:"config,omitempty"`

	// Required. The desired lifetime of the CA certificate. Used to create the
	//  "not_before_time" and "not_after_time" fields inside an X.509
	//  certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.lifetime
	Lifetime *string `json:"lifetime,omitempty"`

	// Required. Immutable. Used when issuing certificates for this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]. If this
	//  [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] is a self-signed CertificateAuthority, this key
	//  is also used to sign the self-signed CA certificate. Otherwise, it
	//  is used to sign a CSR.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.key_spec
	KeySpec *CertificateAuthority_KeyVersionSpec `json:"keySpec,omitempty"`

	// Optional. The [CertificateAuthorityPolicy][google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy] to enforce when issuing
	//  [Certificates][google.cloud.security.privateca.v1beta1.Certificate] from this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.certificate_policy
	CertificatePolicy *CertificateAuthority_CertificateAuthorityPolicy `json:"certificatePolicy,omitempty"`

	// Optional. The [IssuingOptions][google.cloud.security.privateca.v1beta1.CertificateAuthority.IssuingOptions] to follow when issuing [Certificates][google.cloud.security.privateca.v1beta1.Certificate]
	//  from this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.issuing_options
	IssuingOptions *CertificateAuthority_IssuingOptions `json:"issuingOptions,omitempty"`

	// Optional. If this is a subordinate [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority], this field will be set
	//  with the subordinate configuration, which describes its issuers. This may
	//  be updated, but this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] must continue to validate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.subordinate_config
	SubordinateConfig *SubordinateConfig `json:"subordinateConfig,omitempty"`

	// Immutable. The name of a Cloud Storage bucket where this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] will
	//  publish content, such as the CA certificate and CRLs. This must be a bucket
	//  name, without any prefixes (such as `gs://`) or suffixes (such as
	//  `.googleapis.com`). For example, to use a bucket named `my-bucket`, you
	//  would simply specify `my-bucket`. If not specified, a managed bucket will
	//  be created.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.gcs_bucket
	GcsBucket *string `json:"gcsBucket,omitempty"`

	// Optional. Labels with user-defined metadata.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority.AccessUrls
type CertificateAuthority_AccessUrls struct {
	// The URL where this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]'s CA certificate is
	//  published. This will only be set for CAs that have been activated.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.AccessUrls.ca_certificate_access_url
	CaCertificateAccessURL *string `json:"caCertificateAccessURL,omitempty"`

	// The URL where this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]'s CRLs are published. This
	//  will only be set for CAs that have been activated.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.AccessUrls.crl_access_url
	CrlAccessURL *string `json:"crlAccessURL,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy
type CertificateAuthority_CertificateAuthorityPolicy struct {
	// Optional. All [Certificates][google.cloud.security.privateca.v1beta1.Certificate] issued by the [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]
	//  must match at least one listed [ReusableConfigWrapper][google.cloud.security.privateca.v1beta1.ReusableConfigWrapper] in the list.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.allowed_config_list
	AllowedConfigList *CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList `json:"allowedConfigList,omitempty"`

	// Optional. All [Certificates][google.cloud.security.privateca.v1beta1.Certificate] issued by the [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]
	//  will use the provided configuration values, overwriting any requested
	//  configuration values.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.overwrite_config_values
	OverwriteConfigValues *ReusableConfigWrapper `json:"overwriteConfigValues,omitempty"`

	// Optional. If any [Subject][google.cloud.security.privateca.v1beta1.Subject] is specified here, then all
	//  [Certificates][google.cloud.security.privateca.v1beta1.Certificate] issued by the [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] must
	//  match at least one listed [Subject][google.cloud.security.privateca.v1beta1.Subject]. If a [Subject][google.cloud.security.privateca.v1beta1.Subject] has an empty
	//  field, any value will be allowed for that field.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.allowed_locations_and_organizations
	AllowedLocationsAndOrganizations []Subject `json:"allowedLocationsAndOrganizations,omitempty"`

	// Optional. If any value is specified here, then all
	//  [Certificates][google.cloud.security.privateca.v1beta1.Certificate] issued by the [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] must
	//  match at least one listed value. If no value is specified, all values
	//  will be allowed for this fied. Glob patterns are also supported.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.allowed_common_names
	AllowedCommonNames []string `json:"allowedCommonNames,omitempty"`

	// Optional. If a [AllowedSubjectAltNames][google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames] is specified here, then all
	//  [Certificates][google.cloud.security.privateca.v1beta1.Certificate] issued by the [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] must
	//  match [AllowedSubjectAltNames][google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames]. If no value or an empty value
	//  is specified, any value will be allowed for the [SubjectAltNames][google.cloud.security.privateca.v1beta1.SubjectAltNames]
	//  field.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.allowed_sans
	AllowedSans *CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames `json:"allowedSans,omitempty"`

	// Optional. The maximum lifetime allowed by the [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]. Note that
	//  if the any part if the issuing chain expires before a [Certificate][google.cloud.security.privateca.v1beta1.Certificate]'s
	//  requested maximum_lifetime, the effective lifetime will be explicitly
	//  truncated.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.maximum_lifetime
	MaximumLifetime *string `json:"maximumLifetime,omitempty"`

	// Optional. If specified, then only methods allowed in the [IssuanceModes][google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.IssuanceModes] may be
	//  used to issue [Certificates][google.cloud.security.privateca.v1beta1.Certificate].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.allowed_issuance_modes
	AllowedIssuanceModes *CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes `json:"allowedIssuanceModes,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedConfigList
type CertificateAuthority_CertificateAuthorityPolicy_AllowedConfigList struct {
	// Required. All [Certificates][google.cloud.security.privateca.v1beta1.Certificate] issued by the [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]
	//  must match at least one listed [ReusableConfigWrapper][google.cloud.security.privateca.v1beta1.ReusableConfigWrapper]. If a
	//  [ReusableConfigWrapper][google.cloud.security.privateca.v1beta1.ReusableConfigWrapper] has an empty field, any value will be
	//  allowed for that field.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedConfigList.allowed_config_values
	AllowedConfigValues []ReusableConfigWrapper `json:"allowedConfigValues,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames
type CertificateAuthority_CertificateAuthorityPolicy_AllowedSubjectAltNames struct {
	// Optional. Contains valid, fully-qualified host names. Glob patterns are also
	//  supported. To allow an explicit wildcard certificate, escape with
	//  backlash (i.e. `\*`).
	//  E.g. for globbed entries: `*bar.com` will allow `foo.bar.com`, but not
	//  `*.bar.com`, unless the [allow_globbing_dns_wildcards][google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allow_globbing_dns_wildcards] field is set.
	//  E.g. for wildcard entries: `\*.bar.com` will allow `*.bar.com`, but not
	//  `foo.bar.com`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allowed_dns_names
	AllowedDnsNames []string `json:"allowedDnsNames,omitempty"`

	// Optional. Contains valid RFC 3986 URIs. Glob patterns are also supported. To
	//  match across path seperators (i.e. '/') use the double star glob
	//  pattern (i.e. '**').
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allowed_uris
	AllowedUris []string `json:"allowedUris,omitempty"`

	// Optional. Contains valid RFC 2822 E-mail addresses. Glob patterns are also
	//  supported.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allowed_email_addresses
	AllowedEmailAddresses []string `json:"allowedEmailAddresses,omitempty"`

	// Optional. Contains valid 32-bit IPv4 addresses and subnet ranges or RFC 4291 IPv6
	//  addresses and subnet ranges. Subnet ranges are specified using the
	//  '/' notation (e.g. 10.0.0.0/8, 2001:700:300:1800::/64). Glob patterns
	//  are supported only for ip address entries (i.e. not for subnet ranges).
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allowed_ips
	AllowedIps []string `json:"allowedIps,omitempty"`

	// Optional. Specifies if glob patterns used for [allowed_dns_names][google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allowed_dns_names] allows
	//  wildcard certificates.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allow_globbing_dns_wildcards
	AllowGlobbingDnsWildcards *bool `json:"allowGlobbingDnsWildcards,omitempty"`

	// Optional. Specifies if to allow custom X509Extension values.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.AllowedSubjectAltNames.allow_custom_sans
	AllowCustomSans *bool `json:"allowCustomSans,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.IssuanceModes
type CertificateAuthority_CertificateAuthorityPolicy_IssuanceModes struct {
	// Required. When true, allows callers to create [Certificates][google.cloud.security.privateca.v1beta1.Certificate] by
	//  specifying a CSR.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.IssuanceModes.allow_csr_based_issuance
	AllowCsrBasedIssuance *bool `json:"allowCsrBasedIssuance,omitempty"`

	// Required. When true, allows callers to create [Certificates][google.cloud.security.privateca.v1beta1.Certificate] by
	//  specifying a [CertificateConfig][google.cloud.security.privateca.v1beta1.CertificateConfig].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.CertificateAuthorityPolicy.IssuanceModes.allow_config_based_issuance
	AllowConfigBasedIssuance *bool `json:"allowConfigBasedIssuance,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority.IssuingOptions
type CertificateAuthority_IssuingOptions struct {
	// Required. When true, includes a URL to the issuing CA certificate in the
	//  "authority information access" X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.IssuingOptions.include_ca_cert_url
	IncludeCaCertURL *bool `json:"includeCaCertURL,omitempty"`

	// Required. When true, includes a URL to the CRL corresponding to certificates
	//  issued from a [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority].
	//  CRLs will expire 7 days from their creation. However, we will rebuild
	//  daily. CRLs are also rebuilt shortly after a certificate is revoked.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.IssuingOptions.include_crl_access_url
	IncludeCrlAccessURL *bool `json:"includeCrlAccessURL,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority.KeyVersionSpec
type CertificateAuthority_KeyVersionSpec struct {
	// Required. The resource name for an existing Cloud KMS CryptoKeyVersion in the
	//  format
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
	//  This option enables full flexibility in the key's capabilities and
	//  properties.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.KeyVersionSpec.cloud_kms_key_version
	CloudKMSKeyVersion *string `json:"cloudKMSKeyVersion,omitempty"`

	// Required. The algorithm to use for creating a managed Cloud KMS key for a for a
	//  simplified experience. All managed keys will be have their
	//  [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] as `HSM`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.KeyVersionSpec.algorithm
	Algorithm *string `json:"algorithm,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateConfig
type CertificateConfig struct {
	// Required. Specifies some of the values in a certificate that are related to the
	//  subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateConfig.subject_config
	SubjectConfig *CertificateConfig_SubjectConfig `json:"subjectConfig,omitempty"`

	// Required. Describes how some of the technical fields in a certificate should be
	//  populated.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateConfig.reusable_config
	ReusableConfig *ReusableConfigWrapper `json:"reusableConfig,omitempty"`

	// Optional. The public key that corresponds to this config. This is, for example, used
	//  when issuing [Certificates][google.cloud.security.privateca.v1beta1.Certificate], but not when creating a
	//  self-signed [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] or [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] CSR.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateConfig.public_key
	PublicKey *PublicKey `json:"publicKey,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateConfig.SubjectConfig
type CertificateConfig_SubjectConfig struct {
	// Required. Contains distinguished name fields such as the location and organization.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateConfig.SubjectConfig.subject
	Subject *Subject `json:"subject,omitempty"`

	// Optional. The "common name" of the distinguished name.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateConfig.SubjectConfig.common_name
	CommonName *string `json:"commonName,omitempty"`

	// Optional. The subject alternative name fields.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateConfig.SubjectConfig.subject_alt_name
	SubjectAltName *SubjectAltNames `json:"subjectAltName,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateDescription
type CertificateDescription struct {
	// Describes some of the values in a certificate that are related to the
	//  subject and lifetime.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.subject_description
	SubjectDescription *CertificateDescription_SubjectDescription `json:"subjectDescription,omitempty"`

	// Describes some of the technical fields in a certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.config_values
	ConfigValues *ReusableConfigValues `json:"configValues,omitempty"`

	// The public key that corresponds to an issued certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.public_key
	PublicKey *PublicKey `json:"publicKey,omitempty"`

	// Provides a means of identifiying certificates that contain a particular
	//  public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.subject_key_id
	SubjectKeyID *CertificateDescription_KeyId `json:"subjectKeyID,omitempty"`

	// Identifies the subject_key_id of the parent certificate, per
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.1
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.authority_key_id
	AuthorityKeyID *CertificateDescription_KeyId `json:"authorityKeyID,omitempty"`

	// Describes a list of locations to obtain CRL information, i.e.
	//  the DistributionPoint.fullName described by
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.13
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.crl_distribution_points
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty"`

	// Describes lists of issuer CA certificate URLs that appear in the
	//  "Authority Information Access" extension in the certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.aia_issuing_certificate_urls
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty"`

	// The hash of the x.509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.cert_fingerprint
	CertFingerprint *CertificateDescription_CertificateFingerprint `json:"certFingerprint,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateDescription.CertificateFingerprint
type CertificateDescription_CertificateFingerprint struct {
	// The SHA 256 hash, encoded in hexadecimal, of the DER x509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.CertificateFingerprint.sha256_hash
	Sha256Hash *string `json:"sha256Hash,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateDescription.KeyId
type CertificateDescription_KeyId struct {
	// Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most
	//  likely the 160 bit SHA-1 hash of the public key.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.KeyId.key_id
	KeyID *string `json:"keyID,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription
type CertificateDescription_SubjectDescription struct {
	// Contains distinguished name fields such as the location and organization.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription.subject
	Subject *Subject `json:"subject,omitempty"`

	// The "common name" of the distinguished name.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription.common_name
	CommonName *string `json:"commonName,omitempty"`

	// The subject alternative name fields.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription.subject_alt_name
	SubjectAltName *SubjectAltNames `json:"subjectAltName,omitempty"`

	// The serial number encoded in lowercase hexadecimal.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription.hex_serial_number
	HexSerialNumber *string `json:"hexSerialNumber,omitempty"`

	// For convenience, the actual lifetime of an issued certificate.
	//  Corresponds to 'not_after_time' - 'not_before_time'.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription.lifetime
	Lifetime *string `json:"lifetime,omitempty"`

	// The time at which the certificate becomes valid.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription.not_before_time
	NotBeforeTime *string `json:"notBeforeTime,omitempty"`

	// The time at which the certificate expires.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateDescription.SubjectDescription.not_after_time
	NotAfterTime *string `json:"notAfterTime,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.KeyUsage
type KeyUsage struct {
	// Describes high-level ways in which a key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.base_key_usage
	BaseKeyUsage *KeyUsage_KeyUsageOptions `json:"baseKeyUsage,omitempty"`

	// Detailed scenarios in which a key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.extended_key_usage
	ExtendedKeyUsage *KeyUsage_ExtendedKeyUsageOptions `json:"extendedKeyUsage,omitempty"`

	// Used to describe extended key usages that are not listed in the
	//  [KeyUsage.ExtendedKeyUsageOptions][google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions] message.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.unknown_extended_key_usages
	UnknownExtendedKeyUsages []ObjectId `json:"unknownExtendedKeyUsages,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions
type KeyUsage_ExtendedKeyUsageOptions struct {
	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW
	//  server authentication", though regularly used for non-WWW TLS.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.server_auth
	ServerAuth *bool `json:"serverAuth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW
	//  client authentication", though regularly used for non-WWW TLS.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.client_auth
	ClientAuth *bool `json:"clientAuth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of
	//  downloadable executable code client authentication".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.code_signing
	CodeSigning *bool `json:"codeSigning,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email
	//  protection".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.email_protection
	EmailProtection *bool `json:"emailProtection,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding
	//  the hash of an object to a time".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.time_stamping
	TimeStamping *bool `json:"timeStamping,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing
	//  OCSP responses".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.ocsp_signing
	OcspSigning *bool `json:"ocspSigning,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions
type KeyUsage_KeyUsageOptions struct {
	// The key may be used for digital signatures.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.digital_signature
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	// The key may be used for cryptographic commitments. Note that this may
	//  also be referred to as "non-repudiation".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.content_commitment
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	// The key may be used to encipher other keys.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.key_encipherment
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`

	// The key may be used to encipher data.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.data_encipherment
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	// The key may be used in a key agreement protocol.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.key_agreement
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	// The key may be used to sign certificates.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.cert_sign
	CertSign *bool `json:"certSign,omitempty"`

	// The key may be used sign certificate revocation lists.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.crl_sign
	CrlSign *bool `json:"crlSign,omitempty"`

	// The key may be used to encipher only.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.encipher_only
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	// The key may be used to decipher only.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.decipher_only
	DecipherOnly *bool `json:"decipherOnly,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ObjectId
type ObjectId struct {
	// Required. The parts of an OID path. The most significant parts of the path come
	//  first.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ObjectId.object_id_path
	ObjectIDPath []int32 `json:"objectIDPath,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.PublicKey
type PublicKey struct {
	// Required. The type of public key.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.PublicKey.type
	Type *string `json:"type,omitempty"`

	// Required. A public key. Padding and encoding varies by 'KeyType' and is described
	//  along with the KeyType values.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.PublicKey.key
	Key []byte `json:"key,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ReusableConfigValues
type ReusableConfigValues struct {
	// Optional. Indicates the intended use for keys that correspond to a certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.key_usage
	KeyUsage *KeyUsage `json:"keyUsage,omitempty"`

	// Optional. Describes options in this [ReusableConfigValues][google.cloud.security.privateca.v1beta1.ReusableConfigValues] that are
	//  relevant in a CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.ca_options
	CaOptions *ReusableConfigValues_CaOptions `json:"caOptions,omitempty"`

	// Optional. Describes the X.509 certificate policy object identifiers, per
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.policy_ids
	PolicyIds []ObjectId `json:"policyIds,omitempty"`

	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses
	//  that appear in the "Authority Information Access" extension in the
	//  certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.aia_ocsp_servers
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	// Optional. Describes custom X.509 extensions.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.additional_extensions
	AdditionalExtensions []X509Extension `json:"additionalExtensions,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ReusableConfigValues.CaOptions
type ReusableConfigValues_CaOptions struct {
	// Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this
	//  value is missing, the extension will be omitted from the CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.CaOptions.is_ca
	IsCa *bool `json:"isCa,omitempty"`

	// Optional. Refers to the path length restriction X.509 extension. For a CA
	//  certificate, this value describes the depth of subordinate CA
	//  certificates that are allowed.
	//  If this value is less than 0, the request will fail.
	//  If this value is missing, the max path length will be omitted from the
	//  CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.CaOptions.max_issuer_path_length
	MaxIssuerPathLength *Int32Value `json:"maxIssuerPathLength,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ReusableConfigWrapper
type ReusableConfigWrapper struct {
	// Required. A resource path to a [ReusableConfig][google.cloud.security.privateca.v1beta1.ReusableConfig] in the format
	//  `projects/*/locations/*/reusableConfigs/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigWrapper.reusable_config
	ReusableConfig *string `json:"reusableConfig,omitempty"`

	// Required. A user-specified inline [ReusableConfigValues][google.cloud.security.privateca.v1beta1.ReusableConfigValues].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigWrapper.reusable_config_values
	ReusableConfigValues *ReusableConfigValues `json:"reusableConfigValues,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.Subject
type Subject struct {
	// The country code of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.Subject.country_code
	CountryCode *string `json:"countryCode,omitempty"`

	// The organization of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.Subject.organization
	Organization *string `json:"organization,omitempty"`

	// The organizational_unit of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.Subject.organizational_unit
	OrganizationalUnit *string `json:"organizationalUnit,omitempty"`

	// The locality or city of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.Subject.locality
	Locality *string `json:"locality,omitempty"`

	// The province, territory, or regional state of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.Subject.province
	Province *string `json:"province,omitempty"`

	// The street address of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.Subject.street_address
	StreetAddress *string `json:"streetAddress,omitempty"`

	// The postal code of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.Subject.postal_code
	PostalCode *string `json:"postalCode,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.SubjectAltNames
type SubjectAltNames struct {
	// Contains only valid, fully-qualified host names.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubjectAltNames.dns_names
	DnsNames []string `json:"dnsNames,omitempty"`

	// Contains only valid RFC 3986 URIs.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubjectAltNames.uris
	Uris []string `json:"uris,omitempty"`

	// Contains only valid RFC 2822 E-mail addresses.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubjectAltNames.email_addresses
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubjectAltNames.ip_addresses
	IPAddresses []string `json:"ipAddresses,omitempty"`

	// Contains additional subject alternative name values.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubjectAltNames.custom_sans
	CustomSans []X509Extension `json:"customSans,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.SubordinateConfig
type SubordinateConfig struct {
	// Required. This can refer to a [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] in the same project that
	//  was used to create a subordinate [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]. This field
	//  is used for information and usability purposes only. The resource name
	//  is in the format `projects/*/locations/*/certificateAuthorities/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubordinateConfig.certificate_authority
	CertificateAuthority *string `json:"certificateAuthority,omitempty"`

	// Required. Contains the PEM certificate chain for the issuers of this
	//  [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority], but not pem certificate for this CA itself.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubordinateConfig.pem_issuer_chain
	PemIssuerChain *SubordinateConfig_SubordinateConfigChain `json:"pemIssuerChain,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.SubordinateConfig.SubordinateConfigChain
type SubordinateConfig_SubordinateConfigChain struct {
	// Required. Expected to be in leaf-to-root order according to RFC 5246.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.SubordinateConfig.SubordinateConfigChain.pem_certificates
	PemCertificates []string `json:"pemCertificates,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.X509Extension
type X509Extension struct {
	// Required. The OID for this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.X509Extension.object_id
	ObjectID *ObjectId `json:"objectID,omitempty"`

	// Required. Indicates whether or not this extension is critical (i.e., if the client
	//  does not know how to handle this extension, the client should consider this
	//  to be an error).
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.X509Extension.critical
	Critical *bool `json:"critical,omitempty"`

	// Required. The value of this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.X509Extension.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.Int32Value
type Int32Value struct {
	// The int32 value.
	// +kcc:proto:field=google.protobuf.Int32Value.value
	Value *int32 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateAuthority
type CertificateAuthorityObservedState struct {
	// Output only. The resource name for this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] in the
	//  format `projects/*/locations/*/certificateAuthorities/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.name
	Name *string `json:"name,omitempty"`

	// Output only. The [State][google.cloud.security.privateca.v1beta1.CertificateAuthority.State] for this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.state
	State *string `json:"state,omitempty"`

	// Output only. This [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]'s certificate chain, including the current
	//  [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]'s certificate. Ordered such that the root issuer
	//  is the final element (consistent with RFC 5246). For a self-signed CA, this
	//  will only list the current [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]'s certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.pem_ca_certificates
	PemCaCertificates []string `json:"pemCaCertificates,omitempty"`

	// Output only. A structured description of this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority]'s CA certificate
	//  and its issuers. Ordered as self-to-root.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.ca_certificate_descriptions
	CaCertificateDescriptions []CertificateDescription `json:"caCertificateDescriptions,omitempty"`

	// Output only. URLs for accessing content published by this CA, such as the CA certificate
	//  and CRLs.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.access_urls
	AccessUrls *CertificateAuthority_AccessUrls `json:"accessUrls,omitempty"`

	// Output only. The time at which this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] was created.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] was updated.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this [CertificateAuthority][google.cloud.security.privateca.v1beta1.CertificateAuthority] will be deleted, if
	//  scheduled for deletion.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateAuthority.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`
}
