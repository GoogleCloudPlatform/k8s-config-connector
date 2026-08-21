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
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/kmsrefs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PrivateCACertificateAuthorityGVK = GroupVersion.WithKind("PrivateCACertificateAuthority")

// PrivateCACertificateAuthoritySpec defines the desired state of PrivateCACertificateAuthority
// +kcc:spec:proto=google.cloud.security.privateca.v1.CertificateAuthority
type PrivateCACertificateAuthoritySpec struct {
	// The project that this resource belongs to.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The PrivateCACertificateAuthority name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Required. Immutable. The config used to create a self-signed X.509 certificate or CSR. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateAuthority.config
	Config CertificateConfig `json:"config"`

	/* Immutable. Required. Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateAuthority.key_spec
	KeySpec CertificateAuthority_KeyVersionSpec `json:"keySpec"`

	/* Immutable. Required. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateAuthority.lifetime
	Lifetime string `json:"lifetime"`

	/* Immutable. Required. Immutable. The Type of this CertificateAuthority. Possible values: SELF_SIGNED, SUBORDINATE */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateAuthority.type
	Type string `json:"type"`

	/* Immutable. Optional. The StorageBucket where this CertificateAuthority's CA certificate is published. */
	// +optional
	GcsBucketRef *storagev1beta1.StorageBucketRef `json:"gcsBucketRef,omitempty"`

	/* Immutable. Optional. The PrivateCACAPool that includes this CertificateAuthority. */
	CaPoolRef privatecarefs.PrivateCACAPoolRef `json:"caPoolRef"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateConfig
type CertificateConfig struct {
	/* Required. Specifies some of the values in a certificate that are related to the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.subject_config
	SubjectConfig CertificateConfig_SubjectConfig `json:"subjectConfig"`

	/* Required. Describes how some of the technical X.509 fields in a certificate should be populated. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.x509_config
	X509Config CertificateConfig_X509Config `json:"x509Config"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig
type CertificateConfig_SubjectConfig struct {
	/* Required. Contains distinguished name fields such as the common name, location and organization. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig.subject
	Subject Subject `json:"subject"`

	/* Optional. The subject alternative name fields. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateConfig.SubjectConfig.subject_alt_name
	// +optional
	SubjectAltName *SubjectAltNames `json:"subjectAltName,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.Subject
type Subject struct {
	/* The "common name" of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.common_name
	// +optional
	CommonName *string `json:"commonName,omitempty"`

	/* The country code of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.country_code
	// +optional
	CountryCode *string `json:"countryCode,omitempty"`

	/* The locality or city of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.locality
	// +optional
	Locality *string `json:"locality,omitempty"`

	/* The organization of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.organization
	// +optional
	Organization *string `json:"organization,omitempty"`

	/* The organizational_unit of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.organizational_unit
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty"`

	/* The postal code of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.postal_code
	// +optional
	PostalCode *string `json:"postalCode,omitempty"`

	/* The province, territory, or regional state of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.province
	// +optional
	Province *string `json:"province,omitempty"`

	/* The street address of the subject. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.Subject.street_address
	// +optional
	StreetAddress *string `json:"streetAddress,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.SubjectAltNames
type SubjectAltNames struct {
	/* Contains additional subject alternative name values. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.custom_sans
	// +optional
	CustomSans []CertificateAuthority_X509Extension `json:"customSans,omitempty"`

	/* Contains only valid, fully-qualified host names. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.dns_names
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`

	/* Contains only valid RFC 2822 E-mail addresses. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.email_addresses
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	/* Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.ip_addresses
	// +optional
	IPAddresses []string `json:"ipAddresses,omitempty"`

	/* Contains only valid RFC 3986 URIs. */
	// +kcc:proto:field=google.cloud.security.privateca.v1.SubjectAltNames.uris
	// +optional
	Uris []string `json:"uris,omitempty"`
}

type SubjectAltNamesStatus struct {
	/* Contains additional subject alternative name values. */
	// +optional
	CustomSans []CertificateAuthority_X509ExtensionStatus `json:"customSans,omitempty"`

	/* Contains only valid, fully-qualified host names. */
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`

	/* Contains only valid RFC 2822 E-mail addresses. */
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	/* Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses. */
	// +optional
	IPAddresses []string `json:"ipAddresses,omitempty"`

	/* Contains only valid RFC 3986 URIs. */
	// +optional
	Uris []string `json:"uris,omitempty"`
}

type CertificateConfig_X509Config struct {
	/* Optional. Describes custom X.509 extensions. */
	// +optional
	AdditionalExtensions []CertificateAuthority_X509Extension `json:"additionalExtensions,omitempty"`

	/* Optional. Describes options in this X509Parameters that are relevant in a CA certificate. */
	// +optional
	CaOptions *X509Parameters_CAOptions `json:"caOptions,omitempty"`

	/* Optional. Indicates the intended use for keys that correspond to a certificate. */
	// +optional
	KeyUsage *CertificateAuthority_KeyUsage `json:"keyUsage,omitempty"`

	/* Optional. Describes the X.509 certificate policy object identifiers. */
	// +optional
	PolicyIds []CertificateAuthority_ObjectID `json:"policyIds,omitempty"`
}

type CertificateAuthority_X509Extension struct {
	/* Required. The OID for this X.509 extension. */
	ObjectID CertificateAuthority_ObjectID `json:"objectId"`

	/* Optional. Indicates whether or not this extension is critical. */
	// +optional
	Critical *bool `json:"critical,omitempty"`

	/* Required. The value of this X.509 extension. */
	Value string `json:"value"`
}

type CertificateAuthority_X509ExtensionStatus struct {
	/* Optional. The OID for this X.509 extension. */
	// +optional
	ObjectID *CertificateAuthority_ObjectIDStatus `json:"objectId,omitempty"`

	/* Optional. Indicates whether or not this extension is critical. */
	// +optional
	Critical *bool `json:"critical,omitempty"`

	/* Optional. The value of this X.509 extension. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type CertificateAuthority_ObjectID struct {
	/* Required. The parts of an OID path. */
	ObjectIDPath []int64 `json:"objectIdPath"`
}

type CertificateAuthority_ObjectIDStatus struct {
	/* Optional. The parts of an OID path. */
	// +optional
	ObjectIDPath []int64 `json:"objectIdPath,omitempty"`
}

type CertificateAuthority_KeyUsage struct {
	/* Describes high-level ways in which a key may be used. */
	// +optional
	BaseKeyUsage *KeyUsage_KeyUsageOptions `json:"baseKeyUsage,omitempty"`

	/* Detailed scenarios in which a key may be used. */
	// +optional
	ExtendedKeyUsage *KeyUsage_ExtendedKeyUsageOptions `json:"extendedKeyUsage,omitempty"`

	/* Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message. */
	// +optional
	UnknownExtendedKeyUsages []CertificateAuthority_ObjectID `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateAuthority_KeyUsageStatus struct {
	/* Describes high-level ways in which a key may be used. */
	// +optional
	BaseKeyUsage *KeyUsage_KeyUsageOptions `json:"baseKeyUsage,omitempty"`

	/* Detailed scenarios in which a key may be used. */
	// +optional
	ExtendedKeyUsage *KeyUsage_ExtendedKeyUsageOptions `json:"extendedKeyUsage,omitempty"`

	/* Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message. */
	// +optional
	UnknownExtendedKeyUsages []CertificateAuthority_ObjectIDStatus `json:"unknownExtendedKeyUsages,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateAuthority.KeyVersionSpec
type CertificateAuthority_KeyVersionSpec struct {
	/* Optional. The resource name for an existing Cloud KMS CryptoKeyVersion. */
	// +optional
	CloudKmsKeyVersionRef *kmsrefs.KMSCryptoKeyVersionRef `json:"cloudKmsKeyVersionRef,omitempty"`

	/* Optional. The algorithm to use for creating a managed Cloud KMS key. */
	// +optional
	Algorithm *string `json:"algorithm,omitempty"`
}

// PrivateCACertificateAuthorityStatus defines the config connector machine state of PrivateCACertificateAuthority
type PrivateCACertificateAuthorityStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. URLs for accessing content published by this CA, such as the CA certificate and CRLs. */
	// +optional
	AccessUrls *CertificateAuthority_AccessUrls `json:"accessUrls,omitempty"`

	/* Output only. A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
	   Ordered as self-to-root. */
	// +optional
	CaCertificateDescriptions []CertificateDescription `json:"caCertificateDescriptions,omitempty"`

	// +optional
	Config *CertificateAuthority_ConfigStatus `json:"config,omitempty"`

	// Output only. The time at which this CertificateAuthority was created.
	// +kubebuilder:validation:Format=date-time
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this CertificateAuthority was soft deleted, if it is in the DELETED state.
	// +kubebuilder:validation:Format=date-time
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which this CertificateAuthority will be permanently purged, if it is in the DELETED state.
	// +kubebuilder:validation:Format=date-time
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	/* Output only. This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate. */
	// +optional
	PemCaCertificates []string `json:"pemCaCertificates,omitempty"`

	/* Output only. The State for this CertificateAuthority. Possible values: ENABLED, DISABLED, STAGED, AWAITING_USER_ACTIVATION, DELETED */
	// +optional
	State *string `json:"state,omitempty"`

	/* Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate. */
	// +optional
	SubordinateConfig *SubordinateConfig `json:"subordinateConfig,omitempty"`

	/* Output only. The CaPool.Tier of the CaPool that includes this CertificateAuthority. Possible values: ENTERPRISE, DEVOPS */
	// +optional
	Tier *string `json:"tier,omitempty"`

	/* Output only. The time at which this CertificateAuthority was last updated. */
	// +kubebuilder:validation:Format=date-time
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateAuthority.AccessUrls
type CertificateAuthority_AccessUrls struct {
	/* The URL where this CertificateAuthority's CA certificate is published. This will only be set for CAs that have been activated. */
	// +optional
	CaCertificateAccessUrl *string `json:"caCertificateAccessUrl,omitempty"`

	/* The URLs where this CertificateAuthority's CRLs are published. This will only be set for CAs that have been activated. */
	// +optional
	CrlAccessUrls []string `json:"crlAccessUrls,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription
type CertificateDescription struct {
	/* Describes lists of issuer CA certificate URLs that appear in the "Authority Information Access" extension in the certificate. */
	// +optional
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty"`

	/* Identifies the subject_key_id of the parent certificate, per https://tools.ietf.org/html/rfc5280#section-4.2.1.1 */
	// +optional
	AuthorityKeyId *CertificateDescription_KeyID `json:"authorityKeyId,omitempty"`

	/* The hash of the x.509 certificate. */
	// +optional
	CertFingerprint *CertificateDescription_CertificateFingerprint `json:"certFingerprint,omitempty"`

	/* Describes a list of locations to obtain CRL information, i.e. the DistributionPoint.fullName described by https://tools.ietf.org/html/rfc5280#section-4.2.1.13 */
	// +optional
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty"`

	/* The public key that corresponds to an issued certificate. */
	// +optional
	PublicKey *PublicKey `json:"publicKey,omitempty"`

	/* Describes some of the values in a certificate that are related to the subject and lifetime. */
	// +optional
	SubjectDescription *CertificateDescription_SubjectDescription `json:"subjectDescription,omitempty"`

	/* Provides a means of identifiying certificates that contain a particular public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2. */
	// +optional
	SubjectKeyId *CertificateDescription_KeyID `json:"subjectKeyId,omitempty"`

	/* Describes some of the technical X.509 fields in a certificate. */
	// +optional
	X509Description *CertificateDescription_X509Description `json:"x509Description,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription.KeyId
type CertificateDescription_KeyID struct {
	/* Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key. */
	// +optional
	KeyId *string `json:"keyId,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription.CertificateFingerprint
type CertificateDescription_CertificateFingerprint struct {
	/* The SHA 256 hash, encoded in hexadecimal, of the DER x509 certificate. */
	// +optional
	Sha256Hash *string `json:"sha256Hash,omitempty"`
}

type PublicKey struct {
	/* Required. The format of the public key. Possible values: PEM */
	// +optional
	Format *string `json:"format,omitempty"`

	/* Required. A public key. */
	// +optional
	Key *string `json:"key,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateDescription.SubjectDescription
type CertificateDescription_SubjectDescription struct {
	/* The serial number encoded in lowercase hexadecimal. */
	// +optional
	HexSerialNumber *string `json:"hexSerialNumber,omitempty"`

	/* For convenience, the actual lifetime of an issued certificate. */
	// +optional
	Lifetime *string `json:"lifetime,omitempty"`

	/* The time after which the certificate is expired. Per RFC 5280, the validity period for a certificate is the period of time from not_before_time through not_after_time, inclusive. Corresponds to 'not_before_time' + 'lifetime' - 1 second. */
	// +optional
	// +kubebuilder:validation:Format=date-time
	NotAfterTime *string `json:"notAfterTime,omitempty"`

	/* The time at which the certificate becomes valid. */
	// +optional
	// +kubebuilder:validation:Format=date-time
	NotBeforeTime *string `json:"notBeforeTime,omitempty"`

	/* Contains distinguished name fields such as the common name, location and organization. */
	// +optional
	Subject *Subject `json:"subject,omitempty"`

	/* The subject alternative name fields. */
	// +optional
	SubjectAltName *SubjectAltNamesStatus `json:"subjectAltName,omitempty"`
}

type CertificateDescription_X509Description struct {
	/* Optional. Describes custom X.509 extensions. */
	// +optional
	AdditionalExtensions []CertificateAuthority_X509ExtensionStatus `json:"additionalExtensions,omitempty"`

	/* Optional. Describes lists of Ocsps servers in "Authority Information Access" extension. */
	// +optional
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	/* Optional. Describes options in this X509Parameters that are relevant in a CA certificate. */
	// +optional
	CaOptions *CertificateAuthority_CaOptionsStatus `json:"caOptions,omitempty"`

	/* Optional. Indicates the intended use for keys that correspond to a certificate. */
	// +optional
	KeyUsage *CertificateAuthority_KeyUsageStatus `json:"keyUsage,omitempty"`

	/* Optional. Describes the X.509 certificate policy object identifiers. */
	// +optional
	PolicyIds []CertificateAuthority_ObjectIDStatus `json:"policyIds,omitempty"`
}

type CertificateAuthority_CaOptionsStatus struct {
	/* Optional. Refers to the "CA" field in the Basic Constraints extension. If this is true, the certificate is a CA certificate, and can be used to sign other certificates. */
	// +optional
	IsCa *bool `json:"isCa,omitempty"`

	/* Optional. Refers to the path length constraint field in the Basic Constraints extension. This constraint specifies the maximum number of non-self-issued intermediate certificates that may follow this certificate in a valid certification path. */
	// +optional
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.CertificateAuthority.UserDefinedAccessUrls
type CertificateAuthority_UserDefinedAccessUrls struct {
	// +optional
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty"`

	// +optional
	CrlAccessUrls []string `json:"crlAccessUrls,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.SubordinateConfig
type SubordinateConfig struct {
	/* Required. This can refer to a CertificateAuthority that was used to sign the subordinate CA. */
	// +optional
	CertificateAuthority *string `json:"certificateAuthority,omitempty"`

	/* Required. Contains the certificate chain of the subordinate CA's issuer. */
	// +optional
	PemIssuerChain *SubordinateConfig_SubordinateConfigChain `json:"pemIssuerChain,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.SubordinateConfig.SubordinateConfigChain
type SubordinateConfig_SubordinateConfigChain struct {
	/* Required. Expected to be in leaf-to-root order according to RFC 5246. */
	// +optional
	PemCertificates []string `json:"pemCertificates,omitempty"`
}

type CertificateAuthority_ConfigStatus struct {
	/* Optional. The public key that corresponds to this config. */
	// +optional
	PublicKey *PublicKey `json:"publicKey,omitempty"`

	// +optional
	X509Config *CertificateAuthority_X509ConfigStatus `json:"x509Config,omitempty"`
}

type CertificateAuthority_X509ConfigStatus struct {
	/* Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate. */
	// +optional
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpprivatecacertificateauthority;gcpprivatecacertificateauthorities
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PrivateCACertificateAuthority is the Schema for the PrivateCACertificateAuthority API
// +k8s:openapi-gen=true
type PrivateCACertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PrivateCACertificateAuthoritySpec   `json:"spec,omitempty"`
	Status PrivateCACertificateAuthorityStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PrivateCACertificateAuthorityList contains a list of PrivateCACertificateAuthority
type PrivateCACertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCACertificateAuthority `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivateCACertificateAuthority{}, &PrivateCACertificateAuthorityList{})
	kccscheme.RegisterType(PrivateCACertificateAuthorityGVK, &PrivateCACertificateAuthority{})
}
