// Copyright 2026 Google LLC
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
// krm.group: privateca.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.security.privateca.v1
// resource: PrivateCACAPool:CaPool

package v1beta1

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType
type CAPool_IssuancePolicy_AllowedKeyType struct {
	// Represents an allowed RSA key type.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.rsa
	Rsa *CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType `json:"rsa,omitempty"`

	// Represents an allowed Elliptic Curve key type.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.elliptic_curve
	EllipticCurve *CAPool_IssuancePolicy_AllowedKeyType_EcKeyType `json:"ellipticCurve,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.EcKeyType
type CAPool_IssuancePolicy_AllowedKeyType_EcKeyType struct {
	// Optional. A signature algorithm that must be used. If this is
	//  omitted, any EC-based signature algorithm will be allowed.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.EcKeyType.signature_algorithm
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType.RsaKeyType
type CAPool_IssuancePolicy_AllowedKeyType_RsaKeyType struct {
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
type CAPool_IssuancePolicy_IssuanceModes struct {
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
	AdditionalExtensions []ObjectID `json:"additionalExtensions,omitempty"`
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
	UnknownExtendedKeyUsages []ObjectID `json:"unknownExtendedKeyUsages,omitempty"`
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
	PermittedDNSNames []string `json:"permittedDNSNames,omitempty"`

	// Contains excluded DNS names. Any DNS name that can be
	//  constructed by simply adding zero or more labels to
	//  the left-hand side of the name satisfies the name constraint.
	//  For example, `example.com`, `www.example.com`, `www.sub.example.com`
	//  would satisfy `example.com` while `example1.com` does not.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.NameConstraints.excluded_dns_names
	ExcludedDNSNames []string `json:"excludedDNSNames,omitempty"`

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
