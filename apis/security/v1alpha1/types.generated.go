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


// +kcc:proto=google.cloud.security.privateca.v1.CaPool
type CaPool struct {

	// Required. Immutable. The
	//  [Tier][google.cloud.security.privateca.v1.CaPool.Tier] of this
	//  [CaPool][google.cloud.security.privateca.v1.CaPool].
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.tier
	Tier *string `json:"tier,omitempty"`

	// Optional. The
	//  [IssuancePolicy][google.cloud.security.privateca.v1.CaPool.IssuancePolicy]
	//  to control how
	//  [Certificates][google.cloud.security.privateca.v1.Certificate] will be
	//  issued from this [CaPool][google.cloud.security.privateca.v1.CaPool].
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.issuance_policy
	IssuancePolicy *CaPool_IssuancePolicy `json:"issuancePolicy,omitempty"`

	// Optional. The
	//  [PublishingOptions][google.cloud.security.privateca.v1.CaPool.PublishingOptions]
	//  to follow when issuing
	//  [Certificates][google.cloud.security.privateca.v1.Certificate] from any
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
	//  in this [CaPool][google.cloud.security.privateca.v1.CaPool].
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.publishing_options
	PublishingOptions *CaPool_PublishingOptions `json:"publishingOptions,omitempty"`

	// Optional. Labels with user-defined metadata.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy
type CaPool_IssuancePolicy struct {
	// Optional. If any
	//  [AllowedKeyType][google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType]
	//  is specified, then the certificate request's public key must match one of
	//  the key types listed here. Otherwise, any key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.allowed_key_types
	AllowedKeyTypes []CaPool_IssuancePolicy_AllowedKeyType `json:"allowedKeyTypes,omitempty"`

	// Optional. The maximum lifetime allowed for issued
	//  [Certificates][google.cloud.security.privateca.v1.Certificate]. Note that
	//  if the issuing
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
	//  expires before a
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] resource's
	//  requested maximum_lifetime, the effective lifetime will be explicitly
	//  truncated to match it.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.maximum_lifetime
	MaximumLifetime *string `json:"maximumLifetime,omitempty"`

	// Optional. If specified, then only methods allowed in the
	//  [IssuanceModes][google.cloud.security.privateca.v1.CaPool.IssuancePolicy.IssuanceModes]
	//  may be used to issue
	//  [Certificates][google.cloud.security.privateca.v1.Certificate].
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.allowed_issuance_modes
	AllowedIssuanceModes *CaPool_IssuancePolicy_IssuanceModes `json:"allowedIssuanceModes,omitempty"`

	// Optional. A set of X.509 values that will be applied to all certificates
	//  issued through this [CaPool][google.cloud.security.privateca.v1.CaPool].
	//  If a certificate request includes conflicting values for the same
	//  properties, they will be overwritten by the values defined here. If a
	//  certificate request uses a
	//  [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate]
	//  that defines conflicting
	//  [predefined_values][google.cloud.security.privateca.v1.CertificateTemplate.predefined_values]
	//  for the same properties, the certificate issuance request will fail.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.baseline_values
	BaselineValues *X509Parameters `json:"baselineValues,omitempty"`

	// Optional. Describes constraints on identities that may appear in
	//  [Certificates][google.cloud.security.privateca.v1.Certificate] issued
	//  through this [CaPool][google.cloud.security.privateca.v1.CaPool]. If this
	//  is omitted, then this [CaPool][google.cloud.security.privateca.v1.CaPool]
	//  will not add restrictions on a certificate's identity.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.identity_constraints
	IdentityConstraints *CertificateIdentityConstraints `json:"identityConstraints,omitempty"`

	// Optional. Describes the set of X.509 extensions that may appear in a
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] issued
	//  through this [CaPool][google.cloud.security.privateca.v1.CaPool]. If a
	//  certificate request sets extensions that don't appear in the
	//  [passthrough_extensions][google.cloud.security.privateca.v1.CaPool.IssuancePolicy.passthrough_extensions],
	//  those extensions will be dropped. If a certificate request uses a
	//  [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate]
	//  with
	//  [predefined_values][google.cloud.security.privateca.v1.CertificateTemplate.predefined_values]
	//  that don't appear here, the certificate issuance request will fail. If
	//  this is omitted, then this
	//  [CaPool][google.cloud.security.privateca.v1.CaPool] will not add
	//  restrictions on a certificate's X.509 extensions. These constraints do
	//  not apply to X.509 extensions set in this
	//  [CaPool][google.cloud.security.privateca.v1.CaPool]'s
	//  [baseline_values][google.cloud.security.privateca.v1.CaPool.IssuancePolicy.baseline_values].
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.passthrough_extensions
	PassthroughExtensions *CertificateExtensionConstraints `json:"passthroughExtensions,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType
type CaPool_IssuancePolicy_AllowedKeyType struct {
	// Represents an allowed RSA key type.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.rsa
	Rsa *CaPool_IssuancePolicy_AllowedKeyType_RsaKeyType `json:"rsa,omitempty"`

	// Represents an allowed Elliptic Curve key type.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.elliptic_curve
	EllipticCurve *CaPool_IssuancePolicy_AllowedKeyType_EcKeyType `json:"ellipticCurve,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.EcKeyType
type CaPool_IssuancePolicy_AllowedKeyType_EcKeyType struct {
	// Optional. A signature algorithm that must be used. If this is
	//  omitted, any EC-based signature algorithm will be allowed.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.EcKeyType.signature_algorithm
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.RsaKeyType
type CaPool_IssuancePolicy_AllowedKeyType_RsaKeyType struct {
	// Optional. The minimum allowed RSA modulus size (inclusive), in bits.
	//  If this is not set, or if set to zero, the service-level min RSA
	//  modulus size will continue to apply.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.RsaKeyType.min_modulus_size
	MinModulusSize *int64 `json:"minModulusSize,omitempty"`

	// Optional. The maximum allowed RSA modulus size (inclusive), in bits.
	//  If this is not set, or if set to zero, the service will not enforce
	//  an explicit upper bound on RSA modulus sizes.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.RsaKeyType.max_modulus_size
	MaxModulusSize *int64 `json:"maxModulusSize,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.IssuanceModes
type CaPool_IssuancePolicy_IssuanceModes struct {
	// Optional. When true, allows callers to create
	//  [Certificates][google.cloud.security.privateca.v1.Certificate] by
	//  specifying a CSR.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.IssuanceModes.allow_csr_based_issuance
	AllowCsrBasedIssuance *bool `json:"allowCsrBasedIssuance,omitempty"`

	// Optional. When true, allows callers to create
	//  [Certificates][google.cloud.security.privateca.v1.Certificate] by
	//  specifying a
	//  [CertificateConfig][google.cloud.security.privateca.v1.CertificateConfig].
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.IssuanceModes.allow_config_based_issuance
	AllowConfigBasedIssuance *bool `json:"allowConfigBasedIssuance,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.PublishingOptions
type CaPool_PublishingOptions struct {
	// Optional. When true, publishes each
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]'s
	//  CA certificate and includes its URL in the "Authority Information Access"
	//  X.509 extension in all issued
	//  [Certificates][google.cloud.security.privateca.v1.Certificate]. If this
	//  is false, the CA certificate will not be published and the corresponding
	//  X.509 extension will not be written in issued certificates.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.PublishingOptions.publish_ca_cert
	PublishCaCert *bool `json:"publishCaCert,omitempty"`

	// Optional. When true, publishes each
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]'s
	//  CRL and includes its URL in the "CRL Distribution Points" X.509 extension
	//  in all issued
	//  [Certificates][google.cloud.security.privateca.v1.Certificate]. If this
	//  is false, CRLs will not be published and the corresponding X.509
	//  extension will not be written in issued certificates. CRLs will expire 7
	//  days from their creation. However, we will rebuild daily. CRLs are also
	//  rebuilt shortly after a certificate is revoked.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.PublishingOptions.publish_crl
	PublishCrl *bool `json:"publishCrl,omitempty"`

	// Optional. Specifies the encoding format of each
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
	//  resource's CA certificate and CRLs. If this is omitted, CA certificates
	//  and CRLs will be published in PEM.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.PublishingOptions.encoding_format
	EncodingFormat *string `json:"encodingFormat,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateExtensionConstraints
type CertificateExtensionConstraints struct {
	// Optional. A set of named X.509 extensions. Will be combined with
	//  [additional_extensions][google.cloud.security.privateca.v1.CertificateExtensionConstraints.additional_extensions]
	//  to determine the full set of X.509 extensions.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateExtensionConstraints.known_extensions
	KnownExtensions []string `json:"knownExtensions,omitempty"`

	// Optional. A set of [ObjectIds][google.cloud.security.privateca.v1.ObjectId]
	//  identifying custom X.509 extensions. Will be combined with
	//  [known_extensions][google.cloud.security.privateca.v1.CertificateExtensionConstraints.known_extensions]
	//  to determine the full set of X.509 extensions.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateExtensionConstraints.additional_extensions
	AdditionalExtensions []ObjectId `json:"additionalExtensions,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateIdentityConstraints
type CertificateIdentityConstraints struct {
	// Optional. A CEL expression that may be used to validate the resolved X.509
	//  Subject and/or Subject Alternative Name before a certificate is signed. To
	//  see the full allowed syntax and some examples, see
	//  https://cloud.google.com/certificate-authority-service/docs/using-cel
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateIdentityConstraints.cel_expression
	CelExpression *Expr `json:"celExpression,omitempty"`

	// Required. If this is true, the
	//  [Subject][google.cloud.security.privateca.v1.Subject] field may be copied
	//  from a certificate request into the signed certificate. Otherwise, the
	//  requested [Subject][google.cloud.security.privateca.v1.Subject] will be
	//  discarded.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateIdentityConstraints.allow_subject_passthrough
	AllowSubjectPassthrough *bool `json:"allowSubjectPassthrough,omitempty"`

	// Required. If this is true, the
	//  [SubjectAltNames][google.cloud.security.privateca.v1.SubjectAltNames]
	//  extension may be copied from a certificate request into the signed
	//  certificate. Otherwise, the requested
	//  [SubjectAltNames][google.cloud.security.privateca.v1.SubjectAltNames] will
	//  be discarded.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateIdentityConstraints.allow_subject_alt_names_passthrough
	AllowSubjectAltNamesPassthrough *bool `json:"allowSubjectAltNamesPassthrough,omitempty"`
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

// +kcc:proto=google.type.Expr
type Expr struct {
	// Textual representation of an expression in Common Expression Language
	//  syntax.
	// +kcc:proto:field=google.type.Expr.expression
	Expression *string `json:"expression,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing
	//  its purpose. This can be used e.g. in UIs which allow to enter the
	//  expression.
	// +kcc:proto:field=google.type.Expr.title
	Title *string `json:"title,omitempty"`

	// Optional. Description of the expression. This is a longer text which
	//  describes the expression, e.g. when hovered over it in a UI.
	// +kcc:proto:field=google.type.Expr.description
	Description *string `json:"description,omitempty"`

	// Optional. String indicating the location of the expression for error
	//  reporting, e.g. a file name and a position in the file.
	// +kcc:proto:field=google.type.Expr.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool
type CaPoolObservedState struct {
	// Output only. The resource name for this
	//  [CaPool][google.cloud.security.privateca.v1.CaPool] in the format
	//  `projects/*/locations/*/caPools/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.name
	Name *string `json:"name,omitempty"`
}
