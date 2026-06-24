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
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PrivateCACertificateTemplateGVK = GroupVersion.WithKind("PrivateCACertificateTemplate")

// PrivateCACertificateTemplateSpec defines the desired state of PrivateCACertificateTemplate
// +kcc:spec:proto=google.cloud.security.privateca.v1.CertificateTemplate
type PrivateCACertificateTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The PrivateCACertificateTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of scenarios this template is intended for.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.identity_constraints
	IdentityConstraints *CertificateTemplate_IdentityConstraints `json:"identityConstraints,omitempty"`

	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.passthrough_extensions
	PassthroughExtensions *CertificateTemplate_PassthroughExtensions `json:"passthroughExtensions,omitempty"`

	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.predefined_values
	PredefinedValues *CertificateTemplate_X509Parameters `json:"predefinedValues,omitempty"`
}

// PrivateCACertificateTemplateStatus defines the config connector machine state of PrivateCACertificateTemplate
type PrivateCACertificateTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The time at which this CertificateTemplate was created.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this CertificateTemplate was updated.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

type CertificateTemplate_IdentityConstraints struct {
	// Required. If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded.
	// +required
	AllowSubjectAltNamesPassthrough *bool `json:"allowSubjectAltNamesPassthrough,omitempty"`

	// Required. If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded.
	// +required
	AllowSubjectPassthrough *bool `json:"allowSubjectPassthrough,omitempty"`

	// Optional. A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a certificate is signed. To see the full allowed syntax and some examples, see https://cloud.google.com/certificate-authority-service/docs/using-cel
	CelExpression *CertificateTemplate_Expr `json:"celExpression,omitempty"`
}

type CertificateTemplate_Expr struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `json:"description,omitempty"`
	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `json:"location,omitempty"`
	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title *string `json:"title,omitempty"`
}

type CertificateTemplate_PassthroughExtensions struct {
	// Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
	KnownExtensions []string `json:"knownExtensions,omitempty"`
	// Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.
	AdditionalExtensions []CertificateTemplate_ObjectID `json:"additionalExtensions,omitempty"`
}

type CertificateTemplate_ObjectID struct {
	// Required. The parts of an OID path. The most significant parts of the path come first.
	// +required
	ObjectIDPath []int64 `json:"objectIdPath"`
}

type CertificateTemplate_X509Parameters struct {
	// Optional. Indicates the intended use for keys that correspond to a certificate.
	KeyUsage *CertificateTemplate_KeyUsage `json:"keyUsage,omitempty"`
	// Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
	CAOptions *CertificateTemplate_CAOptions `json:"caOptions,omitempty"`
	// Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	PolicyIds []CertificateTemplate_ObjectID `json:"policyIds,omitempty"`
	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`
	// Optional. Describes custom X.509 extensions.
	AdditionalExtensions []CertificateTemplate_X509Extension `json:"additionalExtensions,omitempty"`
}

type CertificateTemplate_CAOptions struct {
	// Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate.
	IsCA *bool `json:"isCa,omitempty"`
	// Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate.
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength,omitempty"`
}

type CertificateTemplate_X509Extension struct {
	// Required. The OID for this X.509 extension.
	// +required
	ObjectID *CertificateTemplate_ObjectID `json:"objectId,omitempty"`
	// Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
	Critical *bool `json:"critical,omitempty"`
	// Required. The value of this X.509 extension.
	// +required
	Value *string `json:"value,omitempty"`
}

type CertificateTemplate_KeyUsage struct {
	// Describes high-level ways in which a key may be used.
	BaseKeyUsage *CertificateTemplate_KeyUsage_KeyUsageOptions `json:"baseKeyUsage,omitempty"`
	// Detailed scenarios in which a key may be used.
	ExtendedKeyUsage *CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions `json:"extendedKeyUsage,omitempty"`
	// Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
	UnknownExtendedKeyUsages []CertificateTemplate_ObjectID `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateTemplate_KeyUsage_KeyUsageOptions struct {
	// The key may be used to sign certificates.
	CertSign *bool `json:"certSign,omitempty"`
	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	ContentCommitment *bool `json:"contentCommitment,omitempty"`
	// The key may be used sign certificate revocation lists.
	CrlSign *bool `json:"crlSign,omitempty"`
	// The key may be used to encipher data.
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`
	// The key may be used to decipher only.
	DecipherOnly *bool `json:"decipherOnly,omitempty"`
	// The key may be used for digital signatures.
	DigitalSignature *bool `json:"digitalSignature,omitempty"`
	// The key may be used to encipher only.
	EncipherOnly *bool `json:"encipherOnly,omitempty"`
	// The key may be used in a key agreement protocol.
	KeyAgreement *bool `json:"keyAgreement,omitempty"`
	// The key may be used to encipher other keys.
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`
}

type CertificateTemplate_KeyUsage_ExtendedKeyUsageOptions struct {
	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	ClientAuth *bool `json:"clientAuth,omitempty"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	CodeSigning *bool `json:"codeSigning,omitempty"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	EmailProtection *bool `json:"emailProtection,omitempty"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	OcspSigning *bool `json:"ocspSigning,omitempty"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	ServerAuth *bool `json:"serverAuth,omitempty"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	TimeStamping *bool `json:"timeStamping,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpprivatecacertificatetemplate;gcpprivatecacertificatetemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PrivateCACertificateTemplate is the Schema for the PrivateCACertificateTemplate API
// +k8s:openapi-gen=true
type PrivateCACertificateTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PrivateCACertificateTemplateSpec   `json:"spec,omitempty"`
	Status PrivateCACertificateTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PrivateCACertificateTemplateList contains a list of PrivateCACertificateTemplate
type PrivateCACertificateTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCACertificateTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivateCACertificateTemplate{}, &PrivateCACertificateTemplateList{})
}
