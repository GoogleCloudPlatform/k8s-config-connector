// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the_license.
// See the_license for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PrivateCACertificateGVK = GroupVersion.WithKind("PrivateCACertificate")

// PrivateCACertificateSpec defines the desired state of PrivateCACertificate
// +kcc:spec:proto=google.cloud.security.privateca.v1.Certificate
type PrivateCACertificateSpec struct {
	// Immutable. Required. The PrivateCACAPool that includes this Certificate.
	CaPoolRef privatecarefs.PrivateCACAPoolRef `json:"caPoolRef"`

	// Immutable. Optional. The PrivateCACertificateAuthority that issued this Certificate.
	// +optional
	CertificateAuthorityRef *PrivateCACertificateAuthorityRef `json:"certificateAuthorityRef,omitempty"`

	// Immutable. Optional. The PrivateCACertificateTemplate used to issue this Certificate.
	// +optional
	CertificateTemplateRef *PrivateCACertificateTemplateRef `json:"certificateTemplateRef,omitempty"`

	// Immutable. Optional. A description of the certificate and key that does not require X.509 or ASN.1.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.config
	// +optional
	Config *Certificate_Config `json:"config,omitempty"`

	// Immutable. Required. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.lifetime
	Lifetime string `json:"lifetime"`

	// Immutable. The location for the resource.
	Location string `json:"location"`

	// Immutable. Optional. A pem-encoded X.509 certificate signing request (CSR).
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.pem_csr
	// +optional
	PemCsr *string `json:"pemCsr,omitempty"`

	// Immutable. Required. The Project that this resource belongs to.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Optional. Specifies how the Certificate's identity fields are to be decided.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.subject_mode
	// +optional
	SubjectMode *string `json:"subjectMode,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateConfig
type Certificate_Config struct {
	// Immutable. Optional. The public key that corresponds to this config.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.public_key
	// +optional
	PublicKey *CertificatePublicKey `json:"publicKey,omitempty"`

	// Immutable. Required. Specifies some of the values in a certificate that are related to the subject.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.subject_config
	// +required
	SubjectConfig *CertificateSubjectConfig `json:"subjectConfig"`

	// Immutable. Required. Describes how some of the technical X.509 fields in a certificate should be populated.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.x509_config
	// +required
	X509Config *CertificateX509Config `json:"x509Config"`
}

// +kcc:proto=google.cloud.security.privateca.v1.PublicKey
type CertificatePublicKey struct {
	// Immutable. Required. The format of the public key. Possible values: KEY_FORMAT_UNSPECIFIED, PEM
	// +kcc:proto:field=google.cloud.security.privateca.v1.PublicKey.format
	Format string `json:"format"`

	// Immutable. Required. A public key. The padding and encoding must match with the `KeyFormat` value specified for the `format` field.
	// +kcc:proto:field=google.cloud.security.privateca.v1.PublicKey.key
	Key string `json:"key"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig
type CertificateSubjectConfig struct {
	// Immutable. Required. Contains distinguished name fields such as the common name, location and organization.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig.subject
	// +required
	Subject *Subject `json:"subject"`

	// Immutable. Optional. The subject alternative name fields.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig.subject_alt_name
	// +optional
	SubjectAltName *CertificateSubjectAltName `json:"subjectAltName,omitempty"`
}

type CertificateSubjectAltName struct {
	// Immutable. Optional. Contains only valid, fully-qualified host names.
	// +optional
	DnsNames []string `json:"dnsNames,omitempty"`

	// Immutable. Optional. Contains only valid RFC 2822 E-mail addresses.
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// Immutable. Optional. Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty"`

	// Immutable. Optional. Contains only valid RFC 3986 URIs.
	// +optional
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Parameters
type CertificateX509Config struct {
	// Immutable. Optional. Describes custom X.509 extensions.
	// +optional
	AdditionalExtensions []CertificateAdditionalExtensions `json:"additionalExtensions,omitempty"`

	// Immutable. Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	// +optional
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	// Immutable. Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
	// +optional
	CAOptions *CertificateCaOptions `json:"caOptions,omitempty"`

	// Immutable. Optional. Indicates the intended use for keys that correspond to a certificate.
	// +optional
	KeyUsage *CertificateKeyUsage `json:"keyUsage,omitempty"`

	// Immutable. Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +optional
	PolicyIds []CertificatePolicyIds `json:"policyIds,omitempty"`
}

type CertificateAdditionalExtensions struct {
	// Immutable. Optional. Indicates whether or not this extension is critical.
	// +optional
	Critical *bool `json:"critical,omitempty"`

	// Immutable. Required. The OID for this X.509 extension.
	ObjectId CertificateObjectId `json:"objectId"`

	// Immutable. Required. The value of this X.509 extension.
	Value string `json:"value"`
}

type CertificateObjectId struct {
	// Immutable. Required. The parts of an OID path. The most significant parts of the path come first.
	// +required
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type CertificateCaOptions struct {
	// Immutable. Optional. When true, the "CA" in Basic Constraints extension will be set to true.
	// +optional
	IsCa *bool `json:"isCa,omitempty"`

	// Immutable. Optional. Refers to the "path length constraint" in Basic Constraints extension.
	// +optional
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength,omitempty"`

	// Immutable. Optional. When true, the "CA" in Basic Constraints extension will be set to false. If both `is_ca` and `non_ca` are unset, the extension will be omitted from the CA certificate.
	// +optional
	NonCa *bool `json:"nonCa,omitempty"`

	// Immutable. Optional. When true, the "path length constraint" in Basic Constraints extension will be set to 0. if both max_issuer_path_length and zero_max_issuer_path_length are unset, the max path length will be omitted from the CA certificate.
	// +optional
	ZeroMaxIssuerPathLength *bool `json:"zeroMaxIssuerPathLength,omitempty"`
}

type CertificateKeyUsage struct {
	// Immutable. Optional. Describes high-level ways in which a key may be used.
	// +optional
	BaseKeyUsage *CertificateBaseKeyUsage `json:"baseKeyUsage,omitempty"`

	// Immutable. Optional. Detailed scenarios in which a key may be used.
	// +optional
	ExtendedKeyUsage *CertificateExtendedKeyUsage `json:"extendedKeyUsage,omitempty"`

	// Immutable. Optional. Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
	// +optional
	UnknownExtendedKeyUsages []CertificateUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateBaseKeyUsage struct {
	// Immutable. Optional. The key may be used to sign certificates.
	// +optional
	CertSign *bool `json:"certSign,omitempty"`

	// Immutable. Optional. The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	// +optional
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	// Immutable. Optional. The key may be used sign certificate revocation lists.
	// +optional
	CrlSign *bool `json:"crlSign,omitempty"`

	// Immutable. Optional. The key may be used to encipher data.
	// +optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	// Immutable. Optional. The key may be used to decipher only.
	// +optional
	DecipherOnly *bool `json:"decipherOnly,omitempty"`

	// Immutable. Optional. The key may be used for digital signatures.
	// +optional
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	// Immutable. Optional. The key may be used to encipher only.
	// +optional
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	// Immutable. Optional. The key may be used in a key agreement protocol.
	// +optional
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	// Immutable. Optional. The key may be used to encipher other keys.
	// +optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`
}

type CertificateExtendedKeyUsage struct {
	// Immutable. Optional. Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	// +optional
	ClientAuth *bool `json:"clientAuth,omitempty"`

	// Immutable. Optional. Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	// +optional
	CodeSigning *bool `json:"codeSigning,omitempty"`

	// Immutable. Optional. Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	// +optional
	EmailProtection *bool `json:"emailProtection,omitempty"`

	// Immutable. Optional. Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	// +optional
	OcspSigning *bool `json:"ocspSigning,omitempty"`

	// Immutable. Optional. Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	// +optional
	ServerAuth *bool `json:"serverAuth,omitempty"`

	// Immutable. Optional. Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	// +optional
	TimeStamping *bool `json:"timeStamping,omitempty"`
}

type CertificateUnknownExtendedKeyUsages struct {
	// Immutable. Required. The parts of an OID path. The most significant parts of the path come first.
	// +required
	ObjectIdPath []int32 `json:"objectIdPath"`
}

type CertificatePolicyIds struct {
	// Immutable. Required. The parts of an OID path. The most significant parts of the path come first.
	// +required
	ObjectIdPath []int64 `json:"objectIdPath"`
}

// PrivateCACertificateStatus defines the config connector machine state of PrivateCACertificate
// +kcc:status:proto=google.cloud.security.privateca.v1.Certificate
type PrivateCACertificateStatus struct {
	/* Conditions represent the latest available observations of the
	   PrivateCACertificate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Output only. A structured description of the issued X.509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.certificate_description
	// +optional
	CertificateDescription *CertificateCertificateDescriptionStatus `json:"certificateDescription,omitempty"`

	// Output only. The time at which this Certificate was created.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.create_time
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The resource name of the issuing CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.issuer_certificate_authority
	// +optional
	IssuerCertificateAuthority *string `json:"issuerCertificateAuthority,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The pem-encoded, signed X.509 certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.pem_certificate
	// +optional
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Output only. The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.pem_certificate_chain
	// +optional
	PemCertificateChain []string `json:"pemCertificateChain,omitempty"`

	// Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.revocation_details
	// +optional
	RevocationDetails *CertificateRevocationDetailsStatus `json:"revocationDetails,omitempty"`

	// Output only. The time at which this Certificate was updated.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.security.privateca.v1.Certificate.update_time
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

type CertificateCertificateDescriptionStatus struct {
	// Optional. Describes lists of issuer CA certificate URLs that appear in the "Authority Information Access" extension in the certificate.
	// +optional
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty"`

	// Optional. Identifies the subject_key_id of the parent certificate, per https://tools.ietf.org/html/rfc5280#section-4.2.1.1.
	// +optional
	AuthorityKeyId *CertificateAuthorityKeyIdStatus `json:"authorityKeyId,omitempty"`

	// Optional. The hash of the x.509 certificate.
	// +optional
	CertFingerprint *CertificateCertFingerprintStatus `json:"certFingerprint,omitempty"`

	// Optional. Describes a list of locations to obtain CRL information, i.e. the DistributionPoint.fullName described by https://tools.ietf.org/html/rfc5280#section-4.2.1.13.
	// +optional
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty"`

	// Optional. The public key that corresponds to an issued certificate.
	// +optional
	PublicKey *CertificatePublicKeyStatus `json:"publicKey,omitempty"`

	// Optional. Describes some of the values in a certificate that are related to the subject and lifetime.
	// +optional
	SubjectDescription *CertificateSubjectDescriptionStatus `json:"subjectDescription,omitempty"`

	// Optional. Provides a means of identifying certificates that contain a particular public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2.
	// +optional
	SubjectKeyId *CertificateSubjectKeyIdStatus `json:"subjectKeyId,omitempty"`

	// Optional. Describes some of the technical X.509 fields in a certificate.
	// +optional
	X509Description *CertificateX509DescriptionStatus `json:"x509Description,omitempty"`
}

type CertificateAuthorityKeyIdStatus struct {
	// Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key.
	// +optional
	KeyId *string `json:"keyId,omitempty"`
}

type CertificateCertFingerprintStatus struct {
	// Optional. The SHA 256 hash, encoded in hexadecimal, of the DER x509 certificate.
	// +optional
	Sha256Hash *string `json:"sha256Hash,omitempty"`
}

type CertificatePublicKeyStatus struct {
	// Optional. The format of the public key. Possible values: KEY_FORMAT_UNSPECIFIED, PEM.
	// +optional
	Format *string `json:"format,omitempty"`

	// Optional. A public key. The padding and encoding must match with the `KeyFormat` value specified for the `format` field.
	// +optional
	Key *string `json:"key,omitempty"`
}

type CertificateSubjectDescriptionStatus struct {
	// Optional. The serial number encoded in lowercase hexadecimal.
	// +optional
	HexSerialNumber *string `json:"hexSerialNumber,omitempty"`

	// Optional. For convenience, the actual lifetime of an issued certificate.
	// +optional
	Lifetime *string `json:"lifetime,omitempty"`

	// Optional. The time after which the certificate is expired. Per RFC 5280, the validity period for a certificate is the period of time from not_before_time through not_after_time, inclusive. Corresponds to 'not_before_time' + 'lifetime' - 1 second.
	// +kubebuilder:validation:Format=date-time
	// +optional
	NotAfterTime *string `json:"notAfterTime,omitempty"`

	// Optional. The time at which the certificate becomes valid.
	// +kubebuilder:validation:Format=date-time
	// +optional
	NotBeforeTime *string `json:"notBeforeTime,omitempty"`

	// Optional. Contains distinguished name fields such as the common name, location and / organization.
	// +optional
	Subject *CertificateSubjectStatus `json:"subject,omitempty"`

	// Optional. The subject alternative name fields.
	// +optional
	SubjectAltName *CertificateSubjectAltNameStatus `json:"subjectAltName,omitempty"`
}

type CertificateSubjectStatus struct {
	// Optional. The "common name" of the subject.
	// +optional
	CommonName *string `json:"commonName,omitempty"`

	// Optional. The country code of the subject.
	// +optional
	CountryCode *string `json:"countryCode,omitempty"`

	// Optional. The locality or city of the subject.
	// +optional
	Locality *string `json:"locality,omitempty"`

	// Optional. The organization of the subject.
	// +optional
	Organization *string `json:"organization,omitempty"`

	// Optional. The organizational_unit of the subject.
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty"`

	// Optional. The postal code of the subject.
	// +optional
	PostalCode *string `json:"postalCode,omitempty"`

	// Optional. The province, territory, or regional state of the subject.
	// +optional
	Province *string `json:"province,omitempty"`

	// Optional. The street address of the subject.
	// +optional
	StreetAddress *string `json:"streetAddress,omitempty"`
}

type CertificateSubjectAltNameStatus struct {
	// Optional. Contains additional subject alternative name values.
	// +optional
	CustomSans []CertificateCustomSansStatus `json:"customSans,omitempty"`

	// Optional. Contains only valid, fully-qualified host names.
	// +optional
	DnsNames []string `json:"dnsNames,omitempty"`

	// Optional. Contains only valid RFC 2822 E-mail addresses.
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// Optional. Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty"`

	// Optional. Contains only valid RFC 3986 URIs.
	// +optional
	Uris []string `json:"uris,omitempty"`
}

type CertificateCustomSansStatus struct {
	// Optional. Indicates whether or not this extension is critical.
	// +optional
	Critical *bool `json:"critical,omitempty"`

	// Optional. The OID for this X.509 extension.
	// +optional
	ObjectId *CertificateObjectIdStatus `json:"objectId,omitempty"`

	// Optional. The value of this X.509 extension.
	// +optional
	Value *string `json:"value,omitempty"`
}

type CertificateObjectIdStatus struct {
	// Optional. The parts of an OID path. The most significant parts of the path come first.
	// +optional
	ObjectIdPath []int64 `json:"objectIdPath,omitempty"`
}

type CertificateX509DescriptionStatus struct {
	// Optional. Describes custom X.509 extensions.
	// +optional
	AdditionalExtensions []CertificateAdditionalExtensionsStatus `json:"additionalExtensions,omitempty"`

	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	// +optional
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	// Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
	// +optional
	CaOptions *CertificateCaOptionsStatus `json:"caOptions,omitempty"`

	// Optional. Indicates the intended use for keys that correspond to a certificate.
	// +optional
	KeyUsage *CertificateKeyUsageStatus `json:"keyUsage,omitempty"`

	// Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +optional
	PolicyIds []CertificatePolicyIdsStatus `json:"policyIds,omitempty"`
}

type CertificateAdditionalExtensionsStatus struct {
	// Optional. Indicates whether or not this extension is critical.
	// +optional
	Critical *bool `json:"critical,omitempty"`

	// Optional. The OID for this X.509 extension.
	// +optional
	ObjectId *CertificateObjectIdStatus `json:"objectId,omitempty"`

	// Optional. The value of this X.509 extension.
	// +optional
	Value *string `json:"value,omitempty"`
}

type CertificateCaOptionsStatus struct {
	// Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate.
	// +optional
	IsCa *bool `json:"isCa,omitempty"`

	// Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate.
	// +optional
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength,omitempty"`
}

type CertificateKeyUsageStatus struct {
	// Optional. Describes high-level ways in which a key may be used.
	// +optional
	BaseKeyUsage *CertificateBaseKeyUsageStatus `json:"baseKeyUsage,omitempty"`

	// Optional. Detailed scenarios in which a key may be used.
	// +optional
	ExtendedKeyUsage *CertificateExtendedKeyUsageStatus `json:"extendedKeyUsage,omitempty"`

	// Optional. Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
	// +optional
	UnknownExtendedKeyUsages []CertificateUnknownExtendedKeyUsagesStatus `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateBaseKeyUsageStatus struct {
	// Optional. The key may be used to sign certificates.
	// +optional
	CertSign *bool `json:"certSign,omitempty"`

	// Optional. The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	// +optional
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	// Optional. The key may be used sign certificate revocation lists.
	// +optional
	CrlSign *bool `json:"crlSign,omitempty"`

	// Optional. The key may be used to encipher data.
	// +optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	// Optional. The key may be used to decipher only.
	// +optional
	DecipherOnly *bool `json:"decipherOnly,omitempty"`

	// Optional. The key may be used for digital signatures.
	// +optional
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	// Optional. The key may be used to encipher only.
	// +optional
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	// Optional. The key may be used in a key agreement protocol.
	// +optional
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	// Optional. The key may be used to encipher other keys.
	// +optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`
}

type CertificateExtendedKeyUsageStatus struct {
	// Optional. Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	// +optional
	ClientAuth *bool `json:"clientAuth,omitempty"`

	// Optional. Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	// +optional
	CodeSigning *bool `json:"codeSigning,omitempty"`

	// Optional. Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	// +optional
	EmailProtection *bool `json:"emailProtection,omitempty"`

	// Optional. Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	// +optional
	OcspSigning *bool `json:"ocspSigning,omitempty"`

	// Optional. Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	// +optional
	ServerAuth *bool `json:"serverAuth,omitempty"`

	// Optional. Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	// +optional
	TimeStamping *bool `json:"timeStamping,omitempty"`
}

type CertificateUnknownExtendedKeyUsagesStatus struct {
	// Optional. The parts of an OID path. The most significant parts of the path come first.
	// +optional
	ObjectIdPath []int64 `json:"objectIdPath,omitempty"`
}

type CertificatePolicyIdsStatus struct {
	// Optional. The parts of an OID path. The most significant parts of the path come first.
	// +optional
	ObjectIdPath []int64 `json:"objectIdPath,omitempty"`
}

type CertificateSubjectKeyIdStatus struct {
	// Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key.
	// +optional
	KeyId *string `json:"keyId,omitempty"`
}

type CertificateRevocationDetailsStatus struct {
	// Optional. Indicates why a Certificate was revoked. Possible values: REVOCATION_REASON_UNSPECIFIED, KEY_COMPROMISE, CERTIFICATE_AUTHORITY_COMPROMISE, AFFILIATION_CHANGED, SUPERSEDED, CESSATION_OF_OPERATION, CERTIFICATE_HOLD, PRIVILEGE_WITHDRAWN, ATTRIBUTE_AUTHORITY_COMPROMISE
	// +optional
	RevocationState *string `json:"revocationState,omitempty"`

	// Optional. The time at which this Certificate was revoked.
	// +kubebuilder:validation:Format=date-time
	// +optional
	RevocationTime *string `json:"revocationTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpprivatecacertificate;gcpprivatecacertificates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PrivateCACertificate is the Schema for the PrivateCACertificate API
// +k8s:openapi-gen=true
type PrivateCACertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PrivateCACertificateSpec   `json:"spec,omitempty"`
	Status PrivateCACertificateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PrivateCACertificateList contains a list of PrivateCACertificate
type PrivateCACertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCACertificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivateCACertificate{}, &PrivateCACertificateList{})
}
