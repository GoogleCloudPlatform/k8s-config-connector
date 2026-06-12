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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PrivateCACertificateTemplateGVK = GroupVersion.WithKind("PrivateCACertificateTemplate")

// PrivateCACertificateTemplateSpec defines the desired state of PrivateCACertificateTemplate
// +kcc:spec:proto=google.cloud.security.privateca.v1.CertificateTemplate
type PrivateCACertificateTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The PrivateCACertificateTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of scenarios this template is
	//  intended for.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Describes constraints on identities that may be appear in
	//  [Certificates][google.cloud.security.privateca.v1.Certificate] issued using
	//  this template. If this is omitted, then this template will not add
	//  restrictions on a certificate's identity.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.identity_constraints
	IdentityConstraints *CertificateIdentityConstraints `json:"identityConstraints,omitempty"`

	// Optional. Describes the set of X.509 extensions that may appear in a
	//  [Certificate][google.cloud.security.privateca.v1.Certificate] issued using
	//  this
	//  [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate].
	//  If a certificate request sets extensions that don't appear in the
	//  [passthrough_extensions][google.cloud.security.privateca.v1.CertificateTemplate.passthrough_extensions],
	//  those extensions will be dropped. If the issuing
	//  [CaPool][google.cloud.security.privateca.v1.CaPool]'s
	//  [IssuancePolicy][google.cloud.security.privateca.v1.CaPool.IssuancePolicy]
	//  defines
	//  [baseline_values][google.cloud.security.privateca.v1.CaPool.IssuancePolicy.baseline_values]
	//  that don't appear here, the certificate issuance request will fail. If this
	//  is omitted, then this template will not add restrictions on a certificate's
	//  X.509 extensions. These constraints do not apply to X.509 extensions set in
	//  this
	//  [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate]'s
	//  [predefined_values][google.cloud.security.privateca.v1.CertificateTemplate.predefined_values].
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.passthrough_extensions
	PassthroughExtensions *CertificateExtensionConstraints `json:"passthroughExtensions,omitempty"`

	// Optional. A set of X.509 values that will be applied to all issued
	//  certificates that use this template. If the certificate request includes
	//  conflicting values for the same properties, they will be overwritten by the
	//  values defined here. If the issuing
	//  [CaPool][google.cloud.security.privateca.v1.CaPool]'s
	//  [IssuancePolicy][google.cloud.security.privateca.v1.CaPool.IssuancePolicy]
	//  defines conflicting
	//  [baseline_values][google.cloud.security.privateca.v1.CaPool.IssuancePolicy.baseline_values]
	//  for the same properties, the certificate issuance request will fail.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateTemplate.predefined_values
	PredefinedValues *X509Parameters `json:"predefinedValues,omitempty"`
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
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this CertificateTemplate was updated.
	// +kubebuilder:validation:Format=date-time
	UpdateTime *string `json:"updateTime,omitempty"`
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
	// +required
	AllowSubjectPassthrough *bool `json:"allowSubjectPassthrough"`

	// Required. If this is true, the
	//  [SubjectAltNames][google.cloud.security.privateca.v1.SubjectAltNames]
	//  extension may be copied from a certificate request into the signed
	//  certificate. Otherwise, the requested
	//  [SubjectAltNames][google.cloud.security.privateca.v1.SubjectAltNames] will
	//  be discarded.
	// +kcc:proto:field=google.cloud.security.privateca.v1.CertificateIdentityConstraints.allow_subject_alt_names_passthrough
	// +required
	AllowSubjectAltNamesPassthrough *bool `json:"allowSubjectAltNamesPassthrough"`
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
