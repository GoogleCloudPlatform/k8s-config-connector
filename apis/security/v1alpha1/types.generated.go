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


// +kcc:proto=google.cloud.security.privateca.v1.Certificate
type Certificate struct {

	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.pem_csr
	PemCsr *string `json:"pemCsr,omitempty"`

	// Immutable. A description of the certificate and key that does not require
	//  X.509 or ASN.1.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.config
	Config *CertificateConfig `json:"config,omitempty"`

	// Required. Immutable. The desired lifetime of a certificate. Used to create
	//  the "not_before_time" and "not_after_time" fields inside an X.509
	//  certificate. Note that the lifetime may be truncated if it would extend
	//  past the life of any certificate authority in the issuing chain.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.lifetime
	Lifetime *string `json:"lifetime,omitempty"`

	// Immutable. The resource name for a
	//  [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate]
	//  used to issue this certificate, in the format
	//  `projects/*/locations/*/certificateTemplates/*`.
	//  If this is specified, the caller must have the necessary permission to
	//  use this template. If this is omitted, no template will be used.
	//  This template must be in the same location as the
	//  [Certificate][google.cloud.security.privateca.v1.Certificate].
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.certificate_template
	CertificateTemplate *string `json:"certificateTemplate,omitempty"`

	// Immutable. Specifies how the
	//  [Certificate][google.cloud.security.privateca.v1.Certificate]'s identity
	//  fields are to be decided. If this is omitted, the `DEFAULT` subject mode
	//  will be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.subject_mode
	SubjectMode *string `json:"subjectMode,omitempty"`

	// Optional. Labels with user-defined metadata.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.Certificate.RevocationDetails
type Certificate_RevocationDetails struct {
	// Indicates why a
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] was
	//  revoked.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.RevocationDetails.revocation_state
	RevocationState *string `json:"revocationState,omitempty"`

	// The time at which this
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] was
	//  revoked.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.RevocationDetails.revocation_time
	RevocationTime *string `json:"revocationTime,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateConfig
type CertificateConfig struct {
	// Required. Specifies some of the values in a certificate that are related to
	//  the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.subject_config
	SubjectConfig *CertificateConfig_SubjectConfig `json:"subjectConfig,omitempty"`

	// Required. Describes how some of the technical X.509 fields in a certificate
	//  should be populated.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.x509_config
	X509Config *X509Parameters `json:"x509Config,omitempty"`

	// Optional. The public key that corresponds to this config. This is, for
	//  example, used when issuing
	//  [Certificates][google.cloud.security.privateca.v1.Certificate], but not
	//  when creating a self-signed
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
	//  or
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
	//  CSR.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.public_key
	PublicKey *PublicKey `json:"publicKey,omitempty"`

	// Optional. When specified this provides a custom SKI to be used in the
	//  certificate. This should only be used to maintain a SKI of an existing CA
	//  originally created outside CA service, which was not generated using method
	//  (1) described in RFC 5280 section 4.2.1.2.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.subject_key_id
	SubjectKeyID *CertificateConfig_KeyId `json:"subjectKeyID,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateConfig.KeyId
type CertificateConfig_KeyId struct {
	// Required. The value of this KeyId encoded in lowercase hexadecimal. This
	//  is most likely the 160 bit SHA-1 hash of the public key.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.KeyId.key_id
	KeyID *string `json:"keyID,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig
type CertificateConfig_SubjectConfig struct {
	// Optional. Contains distinguished name fields such as the common name,
	//  location and organization.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig.subject
	Subject *Subject `json:"subject,omitempty"`

	// Optional. The subject alternative name fields.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig.subject_alt_name
	SubjectAltName *SubjectAltNames `json:"subjectAltName,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription
type CertificateDescription struct {
	// Describes some of the values in a certificate that are related to the
	//  subject and lifetime.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.subject_description
	SubjectDescription *CertificateDescription_SubjectDescription `json:"subjectDescription,omitempty"`

	// Describes some of the technical X.509 fields in a certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.x509_description
	X509Description *X509Parameters `json:"x509Description,omitempty"`

	// The public key that corresponds to an issued certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.public_key
	PublicKey *PublicKey `json:"publicKey,omitempty"`

	// Provides a means of identifiying certificates that contain a particular
	//  public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.subject_key_id
	SubjectKeyID *CertificateDescription_KeyId `json:"subjectKeyID,omitempty"`

	// Identifies the subject_key_id of the parent certificate, per
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.1
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.authority_key_id
	AuthorityKeyID *CertificateDescription_KeyId `json:"authorityKeyID,omitempty"`

	// Describes a list of locations to obtain CRL information, i.e.
	//  the DistributionPoint.fullName described by
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.13
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.crl_distribution_points
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty"`

	// Describes lists of issuer CA certificate URLs that appear in the
	//  "Authority Information Access" extension in the certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.aia_issuing_certificate_urls
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty"`

	// The hash of the x.509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.cert_fingerprint
	CertFingerprint *CertificateDescription_CertificateFingerprint `json:"certFingerprint,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription.CertificateFingerprint
type CertificateDescription_CertificateFingerprint struct {
	// The SHA 256 hash, encoded in hexadecimal, of the DER x509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.CertificateFingerprint.sha256_hash
	Sha256Hash *string `json:"sha256Hash,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription.KeyId
type CertificateDescription_KeyId struct {
	// Optional. The value of this KeyId encoded in lowercase hexadecimal. This
	//  is most likely the 160 bit SHA-1 hash of the public key.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.KeyId.key_id
	KeyID *string `json:"keyID,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription
type CertificateDescription_SubjectDescription struct {
	// Contains distinguished name fields such as the common name, location and
	//  / organization.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription.subject
	Subject *Subject `json:"subject,omitempty"`

	// The subject alternative name fields.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription.subject_alt_name
	SubjectAltName *SubjectAltNames `json:"subjectAltName,omitempty"`

	// The serial number encoded in lowercase hexadecimal.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription.hex_serial_number
	HexSerialNumber *string `json:"hexSerialNumber,omitempty"`

	// For convenience, the actual lifetime of an issued certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription.lifetime
	Lifetime *string `json:"lifetime,omitempty"`

	// The time at which the certificate becomes valid.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription.not_before_time
	NotBeforeTime *string `json:"notBeforeTime,omitempty"`

	// The time after which the certificate is expired.
	//  Per RFC 5280, the validity period for a certificate is the period of time
	//  from not_before_time through not_after_time, inclusive.
	//  Corresponds to 'not_before_time' + 'lifetime' - 1 second.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription.not_after_time
	NotAfterTime *string `json:"notAfterTime,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.KeyUsage
type KeyUsage struct {
	// Describes high-level ways in which a key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.base_key_usage
	BaseKeyUsage *KeyUsage_KeyUsageOptions `json:"baseKeyUsage,omitempty"`

	// Detailed scenarios in which a key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.extended_key_usage
	ExtendedKeyUsage *KeyUsage_ExtendedKeyUsageOptions `json:"extendedKeyUsage,omitempty"`

	// Used to describe extended key usages that are not listed in the
	//  [KeyUsage.ExtendedKeyUsageOptions][google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions]
	//  message.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.unknown_extended_key_usages
	UnknownExtendedKeyUsages []ObjectId `json:"unknownExtendedKeyUsages,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions
type KeyUsage_ExtendedKeyUsageOptions struct {
	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW
	//  server authentication", though regularly used for non-WWW TLS.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions.server_auth
	ServerAuth *bool `json:"serverAuth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW
	//  client authentication", though regularly used for non-WWW TLS.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions.client_auth
	ClientAuth *bool `json:"clientAuth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of
	//  downloadable executable code client authentication".
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions.code_signing
	CodeSigning *bool `json:"codeSigning,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email
	//  protection".
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions.email_protection
	EmailProtection *bool `json:"emailProtection,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding
	//  the hash of an object to a time".
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions.time_stamping
	TimeStamping *bool `json:"timeStamping,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing
	//  OCSP responses".
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.ExtendedKeyUsageOptions.ocsp_signing
	OcspSigning *bool `json:"ocspSigning,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions
type KeyUsage_KeyUsageOptions struct {
	// The key may be used for digital signatures.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.digital_signature
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	// The key may be used for cryptographic commitments. Note that this may
	//  also be referred to as "non-repudiation".
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.content_commitment
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	// The key may be used to encipher other keys.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.key_encipherment
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`

	// The key may be used to encipher data.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.data_encipherment
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	// The key may be used in a key agreement protocol.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.key_agreement
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	// The key may be used to sign certificates.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.cert_sign
	CertSign *bool `json:"certSign,omitempty"`

	// The key may be used sign certificate revocation lists.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.crl_sign
	CrlSign *bool `json:"crlSign,omitempty"`

	// The key may be used to encipher only.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.encipher_only
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	// The key may be used to decipher only.
	// +kcc:proto:field=google.cloud.security.privateca.v1.KeyUsage.KeyUsageOptions.decipher_only
	DecipherOnly *bool `json:"decipherOnly,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.ObjectId
type ObjectId struct {
	// Required. The parts of an OID path. The most significant parts of the path
	//  come first.
	// +kcc:proto:field=google.cloud.security.privateca.v1.ObjectId.object_id_path
	ObjectIDPath []int32 `json:"objectIDPath,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.PublicKey
type PublicKey struct {
	// Required. A public key. The padding and encoding
	//  must match with the `KeyFormat` value specified for the `format` field.
	// +kcc:proto:field=google.cloud.security.privateca.v1.PublicKey.key
	Key []byte `json:"key,omitempty"`

	// Required. The format of the public key.
	// +kcc:proto:field=google.cloud.security.privateca.v1.PublicKey.format
	Format *string `json:"format,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.Subject
type Subject struct {
	// The "common name" of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.common_name
	CommonName *string `json:"commonName,omitempty"`

	// The country code of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.country_code
	CountryCode *string `json:"countryCode,omitempty"`

	// The organization of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.organization
	Organization *string `json:"organization,omitempty"`

	// The organizational_unit of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.organizational_unit
	OrganizationalUnit *string `json:"organizationalUnit,omitempty"`

	// The locality or city of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.locality
	Locality *string `json:"locality,omitempty"`

	// The province, territory, or regional state of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.province
	Province *string `json:"province,omitempty"`

	// The street address of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.street_address
	StreetAddress *string `json:"streetAddress,omitempty"`

	// The postal code of the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.postal_code
	PostalCode *string `json:"postalCode,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.SubjectAltNames
type SubjectAltNames struct {
	// Contains only valid, fully-qualified host names.
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.dns_names
	DnsNames []string `json:"dnsNames,omitempty"`

	// Contains only valid RFC 3986 URIs.
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.uris
	Uris []string `json:"uris,omitempty"`

	// Contains only valid RFC 2822 E-mail addresses.
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.email_addresses
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.ip_addresses
	IPAddresses []string `json:"ipAddresses,omitempty"`

	// Contains additional subject alternative name values.
	//  For each custom_san, the `value` field must contain an ASN.1 encoded
	//  UTF8String.
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.custom_sans
	CustomSans []X509Extension `json:"customSans,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Extension
type X509Extension struct {
	// Required. The OID for this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Extension.object_id
	ObjectID *ObjectId `json:"objectID,omitempty"`

	// Optional. Indicates whether or not this extension is critical (i.e., if the
	//  client does not know how to handle this extension, the client should
	//  consider this to be an error).
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Extension.critical
	Critical *bool `json:"critical,omitempty"`

	// Required. The value of this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Extension.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Parameters
type X509Parameters struct {
	// Optional. Indicates the intended use for keys that correspond to a
	//  certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.key_usage
	KeyUsage *KeyUsage `json:"keyUsage,omitempty"`

	// Optional. Describes options in this
	//  [X509Parameters][google.cloud.security.privateca.v1.X509Parameters] that
	//  are relevant in a CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.ca_options
	CaOptions *X509Parameters_CaOptions `json:"caOptions,omitempty"`

	// Optional. Describes the X.509 certificate policy object identifiers, per
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.policy_ids
	PolicyIds []ObjectId `json:"policyIds,omitempty"`

	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint
	//  addresses that appear in the "Authority Information Access" extension in
	//  the certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.aia_ocsp_servers
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	// Optional. Describes the X.509 name constraints extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.name_constraints
	NameConstraints *X509Parameters_NameConstraints `json:"nameConstraints,omitempty"`

	// Optional. Describes custom X.509 extensions.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.additional_extensions
	AdditionalExtensions []X509Extension `json:"additionalExtensions,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Parameters.CaOptions
type X509Parameters_CaOptions struct {
	// Optional. Refers to the "CA" X.509 extension, which is a boolean value.
	//  When this value is missing, the extension will be omitted from the CA
	//  certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.CaOptions.is_ca
	IsCa *bool `json:"isCa,omitempty"`

	// Optional. Refers to the path length restriction X.509 extension. For a CA
	//  certificate, this value describes the depth of subordinate CA
	//  certificates that are allowed.
	//  If this value is less than 0, the request will fail.
	//  If this value is missing, the max path length will be omitted from the
	//  CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.CaOptions.max_issuer_path_length
	MaxIssuerPathLength *int32 `json:"maxIssuerPathLength,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Parameters.NameConstraints
type X509Parameters_NameConstraints struct {
	// Indicates whether or not the name constraints are marked critical.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.critical
	Critical *bool `json:"critical,omitempty"`

	// Contains permitted DNS names. Any DNS name that can be
	//  constructed by simply adding zero or more labels to
	//  the left-hand side of the name satisfies the name constraint.
	//  For example, `example.com`, `www.example.com`, `www.sub.example.com`
	//  would satisfy `example.com` while `example1.com` does not.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.permitted_dns_names
	PermittedDnsNames []string `json:"permittedDnsNames,omitempty"`

	// Contains excluded DNS names. Any DNS name that can be
	//  constructed by simply adding zero or more labels to
	//  the left-hand side of the name satisfies the name constraint.
	//  For example, `example.com`, `www.example.com`, `www.sub.example.com`
	//  would satisfy `example.com` while `example1.com` does not.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.excluded_dns_names
	ExcludedDnsNames []string `json:"excludedDnsNames,omitempty"`

	// Contains the permitted IP ranges. For IPv4 addresses, the ranges
	//  are expressed using CIDR notation as specified in RFC 4632.
	//  For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
	//  addresses.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.permitted_ip_ranges
	PermittedIPRanges []string `json:"permittedIPRanges,omitempty"`

	// Contains the excluded IP ranges. For IPv4 addresses, the ranges
	//  are expressed using CIDR notation as specified in RFC 4632.
	//  For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
	//  addresses.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.excluded_ip_ranges
	ExcludedIPRanges []string `json:"excludedIPRanges,omitempty"`

	// Contains the permitted email addresses. The value can be a particular
	//  email address, a hostname to indicate all email addresses on that host or
	//  a domain with a leading period (e.g. `.example.com`) to indicate
	//  all email addresses in that domain.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.permitted_email_addresses
	PermittedEmailAddresses []string `json:"permittedEmailAddresses,omitempty"`

	// Contains the excluded email addresses. The value can be a particular
	//  email address, a hostname to indicate all email addresses on that host or
	//  a domain with a leading period (e.g. `.example.com`) to indicate
	//  all email addresses in that domain.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.excluded_email_addresses
	ExcludedEmailAddresses []string `json:"excludedEmailAddresses,omitempty"`

	// Contains the permitted URIs that apply to the host part of the name.
	//  The value can be a hostname or a domain with a
	//  leading period (like `.example.com`)
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.permitted_uris
	PermittedUris []string `json:"permittedUris,omitempty"`

	// Contains the excluded URIs that apply to the host part of the name.
	//  The value can be a hostname or a domain with a
	//  leading period (like `.example.com`)
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.excluded_uris
	ExcludedUris []string `json:"excludedUris,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.Certificate
type CertificateObservedState struct {
	// Output only. The resource name for this
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] in the format
	//  `projects/*/locations/*/caPools/*/certificates/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.name
	Name *string `json:"name,omitempty"`

	// Output only. The resource name of the issuing
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
	//  in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.issuer_certificate_authority
	IssuerCertificateAuthority *string `json:"issuerCertificateAuthority,omitempty"`

	// Output only. Details regarding the revocation of this
	//  [Certificate][google.cloud.security.privateca.v1.Certificate]. This
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] is considered
	//  revoked if and only if this field is present.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.revocation_details
	RevocationDetails *Certificate_RevocationDetails `json:"revocationDetails,omitempty"`

	// Output only. The pem-encoded, signed X.509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.pem_certificate
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Output only. A structured description of the issued X.509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.certificate_description
	CertificateDescription *CertificateDescription `json:"certificateDescription,omitempty"`

	// Output only. The chain that may be used to verify the X.509 certificate.
	//  Expected to be in issuer-to-root order according to RFC 5246.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.pem_certificate_chain
	PemCertificateChain []string `json:"pemCertificateChain,omitempty"`

	// Output only. The time at which this
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] was created.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] was updated.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
