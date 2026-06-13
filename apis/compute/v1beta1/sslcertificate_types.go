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
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ComputeManagedSSLCertificateGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "ComputeManagedSSLCertificate",
}

// +kcc:proto=google.cloud.compute.v1.SslCertificateManagedSslCertificate
type ComputeManagedSSLCertificateManaged struct {
	// Domains for which a managed SSL certificate will be valid. Currently,
	// there can be up to 100 domains in this list.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificateManagedSslCertificate.domains
	Domains []string `json:"domains"`
}

// ComputeManagedSSLCertificateSpec defines the desired state of ComputeManagedSSLCertificate
// +kcc:spec:proto=google.cloud.compute.v1.SslCertificate
type ComputeManagedSSLCertificateSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +optional
	Location *string `json:"location,omitempty"`

	// The ComputeManagedSSLCertificate name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. An optional description of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.description
	Description *string `json:"description,omitempty"`

	// Immutable. Properties relevant to a managed certificate. These will be used if the
	// certificate is managed (as indicated by a value of 'MANAGED' in 'type').
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.managed
	Managed *ComputeManagedSSLCertificateManaged `json:"managed,omitempty"`

	// Immutable. Enum field whose value is always 'MANAGED' - used to signal to the API
	// which type this is. Default value: "MANAGED" Possible values: ["MANAGED"].
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SslCertificateManagedSslCertificate
type ComputeManagedSSLCertificateObservedManaged struct {
	// [Output only] Detailed statuses of the domains specified for managed certificate resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificateManagedSslCertificate.domain_status
	DomainStatus map[string]string `json:"domainStatus,omitempty"`

	// [Output only] Status of the managed certificate resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificateManagedSslCertificate.status
	Status *string `json:"status,omitempty"`
}

// ComputeManagedSSLCertificateObservedState is the state of the ComputeManagedSSLCertificate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.SslCertificate
type ComputeManagedSSLCertificateObservedState struct {
	// [Output Only] The unique identifier for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.id
	CertificateID *int64 `json:"certificateID,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] Expire time of the certificate. RFC3339
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// [Output only] Server-defined URL for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Domains associated with the certificate via Subject Alternative Name.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.subject_alternative_names
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`

	// [Output Only] Configuration and status of a managed SSL certificate.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SslCertificate.managed
	Managed *ComputeManagedSSLCertificateObservedManaged `json:"managed,omitempty"`
}

// ComputeManagedSSLCertificateStatus defines the config connector machine state of ComputeManagedSSLCertificate
type ComputeManagedSSLCertificateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeManagedSSLCertificate resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *ComputeManagedSSLCertificateObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputemanagedsslcertificate;gcpcomputemanagedsslcertificates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeManagedSSLCertificate is the Schema for the ComputeManagedSSLCertificate API
// +k8s:openapi-gen=true
type ComputeManagedSSLCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeManagedSSLCertificateSpec   `json:"spec,omitempty"`
	Status ComputeManagedSSLCertificateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeManagedSSLCertificateList contains a list of ComputeManagedSSLCertificate
type ComputeManagedSSLCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeManagedSSLCertificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeManagedSSLCertificate{}, &ComputeManagedSSLCertificateList{})
}
