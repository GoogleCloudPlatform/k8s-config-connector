// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CertificateManagerDNSAuthorizationGVK = GroupVersion.WithKind("CertificateManagerDNSAuthorization")

// CertificateManagerDNSAuthorizationSpec defines the desired state of CertificateManagerDNSAuthorization
// +kcc:spec:proto=google.cloud.certificatemanager.v1.DnsAuthorization
type CertificateManagerDNSAuthorizationSpec struct {
	/* A human-readable description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Domain field is immutable"
	/* Immutable. A domain which is being authorized. A DnsAuthorization resource covers a
	single domain and its wildcard, e.g. authorization for "example.com" can
	be used to issue certificates for "example.com" and "*.example.com". */
	// +required
	Domain string `json:"domain"`

	/* The project that this resource belongs to. */
	// +required
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. Location represents the geographical location of the DnsAuthorization. If not specified, "global" is used. */
	// +optional
	Location string `json:"location"`

	/* Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// CertificateManagerDNSAuthorizationStatus defines the config connector machine state of CertificateManagerDNSAuthorization
type CertificateManagerDNSAuthorizationStatus struct {
	/* Conditions represent the latest available observations of the
	   CertificateManagerDNSAuthorization's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The structure describing the DNS Resource Record that needs to be added
	to DNS configuration for the authorization to be usable by
	certificate. */
	// +optional
	DnsResourceRecord []DnsAuthorization_DnsResourceRecord `json:"dnsResourceRecord,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CertificateManagerDNSAuthorization resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* NOTYET
	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *CertificateManagerDNSAuthorizationObservedState `json:"observedState,omitempty"`
	*/
}

/* NOTYET
// CertificateManagerDNSAuthorizationSpec defines the desired state of CertificateManagerDNSAuthorization
// +kcc:observedstate:proto=google.cloud.certificatemanager.v1.DnsAuthorization
type CertificateManagerDNSAuthorizationObservedState struct {
}
*/

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcertificatemanagerdnsauthorization;gcpcertificatemanagerdnsauthorizations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable"

// CertificateManagerDNSAuthorization is the Schema for the CertificateManagerDNSAuthorization API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type CertificateManagerDNSAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CertificateManagerDNSAuthorizationSpec   `json:"spec,omitempty"`
	Status CertificateManagerDNSAuthorizationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CertificateManagerDNSAuthorizationList contains a list of CertificateManagerDNSAuthorization
type CertificateManagerDNSAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateManagerDNSAuthorization `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateManagerDNSAuthorization{}, &CertificateManagerDNSAuthorizationList{})
}
