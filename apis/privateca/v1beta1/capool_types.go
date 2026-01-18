// Copyright 2025 Google LLC
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
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PrivateCACAPoolGVK = GroupVersion.WithKind("PrivateCACAPool")

// PrivateCACAPoolSpec defines the desired state of PrivateCACAPool
// +kcc:spec:proto=google.cloud.security.privateca.v1.CaPool
type PrivateCACAPoolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The PrivateCACAPool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The Tier of this CaPool.
	// Possible values: TIER_UNSPECIFIED, ENTERPRISE, DEVOPS
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.tier
	Tier *string `json:"tier,omitempty"`

	// Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.issuance_policy
	IssuancePolicy *CAPool_IssuancePolicy `json:"issuancePolicy,omitempty"`

	// Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.publishing_options
	PublishingOptions *CAPool_PublishingOptions `json:"publishingOptions,omitempty"`
}

// PrivateCACAPoolStatus defines the config connector machine state of PrivateCACAPool
type PrivateCACAPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the PrivateCACAPool resource in GCP.
	// ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *PrivateCACAPoolObservedState `json:"observedState,omitempty"`
}

// PrivateCACAPoolObservedState is the state of the PrivateCACAPool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.security.privateca.v1.CaPool
type PrivateCACAPoolObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpprivatecacapool;gcpprivatecacapools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PrivateCACAPool is the Schema for the PrivateCACAPool API
// +k8s:openapi-gen=true
type PrivateCACAPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PrivateCACAPoolSpec   `json:"spec,omitempty"`
	Status PrivateCACAPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PrivateCACAPoolList contains a list of PrivateCACAPool
type PrivateCACAPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCACAPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivateCACAPool{}, &PrivateCACAPoolList{})
}

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.IssuancePolicy
type CAPool_IssuancePolicy struct {
	// Optional. If any
	//  [AllowedKeyType][google.cloud.security.privateca.v1.CaPool.IssuancePolicy.AllowedKeyType]
	//  is specified, then the certificate request's public key must match one of
	//  the key types listed here. Otherwise, any key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.allowed_key_types
	AllowedKeyTypes []CAPool_IssuancePolicy_AllowedKeyType `json:"allowedKeyTypes,omitempty"`

	// Optional. The duration to backdate all certificates issued from this
	//  [CaPool][google.cloud.security.privateca.v1.CaPool]. If not set, the
	//  certificates will be issued with a not_before_time of the issuance time
	//  (i.e. the current time). If set, the certificates will be issued with a
	//  not_before_time of the issuance time minus the backdate_duration. The
	//  not_after_time will be adjusted to preserve the requested lifetime. The
	//  backdate_duration must be less than or equal to 48 hours.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.IssuancePolicy.backdate_duration
	// BackdateDuration *string `json:"backdateDuration,omitempty"`

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
	AllowedIssuanceModes *CAPool_IssuancePolicy_IssuanceModes `json:"allowedIssuanceModes,omitempty"`

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

// +kcc:proto=google.cloud.security.privateca.v1.CaPool.PublishingOptions
type CAPool_PublishingOptions struct {
	// Optional. When true, publishes each
	//  [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]'s
	//  CA certificate and includes its URL in the "Authority Information Access"
	//  X.509 extension in all issued
	//  [Certificates][google.cloud.security.privateca.v1.Certificate]. If this
	//  is false, the CA certificate will not be published and the corresponding
	//  X.509 extension will not be written in issued certificates.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CaPool.PublishingOptions.publish_ca_cert
	PublishCACert *bool `json:"publishCaCert,omitempty"`

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
	// EncodingFormat *string `json:"encodingFormat,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Extension
type X509Extension struct {
	// Required. The OID for this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Extension.object_id
	ObjectID *ObjectID `json:"objectId,omitempty"`

	// Optional. Indicates whether or not this extension is critical (i.e., if the
	//  client does not know how to handle this extension, the client should
	//  consider this to be an error).
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Extension.critical
	Critical *bool `json:"critical,omitempty"`

	// Required. The value of this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Extension.value
	Value string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.ObjectId
type ObjectID struct {
	// Required. The parts of an OID path. The most significant parts of the path
	//  come first.
	// +kcc:proto:field=google.cloud.security.privateca.v1.ObjectId.object_id_path
	ObjectIDPath []int32 `json:"objectIdPath,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Parameters
type X509Parameters struct {
	// Optional. Indicates the intended use for keys that correspond to a
	//  certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.key_usage
	KeyUsage *KeyUsage `json:"keyUsage,omitempty"`

	// Optional. Describes options in this
	//  [X509Parameters][google.cloud.security.privateca.v1.X509Parameters] that
	//  are relevant in a CA certificate. If not specified, a default basic
	//  constraints extension with `is_ca=false` will be added for leaf
	//  certificates.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.ca_options
	CAOptions *X509Parameters_CAOptions `json:"caOptions,omitempty"`

	// Optional. Describes the X.509 certificate policy object identifiers, per
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.policy_ids
	PolicyIds []ObjectID `json:"policyIds,omitempty"`

	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint
	//  addresses that appear in the "Authority Information Access" extension in
	//  the certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.aia_ocsp_servers
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	// Optional. Describes the X.509 name constraints extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.name_constraints
	// NameConstraints *X509Parameters_NameConstraints `json:"nameConstraints,omitempty"`

	// Optional. Describes custom X.509 extensions.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.additional_extensions
	AdditionalExtensions []X509Extension `json:"additionalExtensions,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1.X509Parameters.CaOptions
type X509Parameters_CAOptions struct {
	// Optional. Refers to the "CA" boolean field in the X.509 extension.
	//  When this value is missing, the basic constraints extension will be
	//  omitted from the certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.CaOptions.is_ca
	IsCA *bool `json:"isCa,omitempty"`

	// Optional. Refers to the path length constraint field in the X.509
	//  extension. For a CA certificate, this value describes the depth of
	//  subordinate CA certificates that are allowed. If this value is less than
	//  0, the request will fail. If this value is missing, the max path length
	//  will be omitted from the certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1.X509Parameters.CaOptions.max_issuer_path_length
	MaxIssuerPathLength *int32 `json:"maxIssuerPathLength,omitempty"`

	// Optional. When true, the "path length constraint" in Basic Constraints extension will be set to 0. if both max_issuer_path_length and zero_max_issuer_path_length are unset, the max path length will be omitted from the CA certificate.
	ZeroMaxIssuerPathLength *bool `json:"zeroMaxIssuerPathLength,omitempty"`
}
