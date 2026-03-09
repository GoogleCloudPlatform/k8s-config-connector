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

var ComputePublicAdvertisedPrefixGVK = GroupVersion.WithKind("ComputePublicAdvertisedPrefix")

// ComputePublicAdvertisedPrefixSpec defines the desired state of ComputePublicAdvertisedPrefix
// +kcc:spec:proto=google.cloud.compute.v1.PublicAdvertisedPrefix
type ComputePublicAdvertisedPrefixSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty"`

	// The IPv4 address to be used for reverse DNS verification.
	// +required
	DNSVerificationIP *string `json:"dnsVerificationIp"`

	// The IPv4 address range, in CIDR format, represented by this public advertised prefix.
	// +required
	IPCidrRange *string `json:"ipCidrRange"`

	// The ComputePublicAdvertisedPrefix name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputePublicAdvertisedPrefixStatus defines the config connector machine state of ComputePublicAdvertisedPrefix
type ComputePublicAdvertisedPrefixStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputePublicAdvertisedPrefix resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputePublicAdvertisedPrefixObservedState `json:"observedState,omitempty"`
}

// ComputePublicAdvertisedPrefixObservedState is the state of the ComputePublicAdvertisedPrefix resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.PublicAdvertisedPrefix
type ComputePublicAdvertisedPrefixObservedState struct {
	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	Fingerprint *string `json:"fingerprint,omitempty"`

	// The shared secret to be used for reverse DNS verification.
	SharedSecret *string `json:"sharedSecret,omitempty"`

	// The status of the public advertised prefix. Possible values include:
	// "INITIAL", "PTR_CONFIGURED", "VALIDATED", "REVERSE_DNS_LOOKUP_FAILED",
	// "PREFIX_CONFIGURATION_IN_PROGRESS",
	// "PREFIX_CONFIGURATION_COMPLETE", "PREFIX_REMOVAL_IN_PROGRESS".
	Status *string `json:"status,omitempty"`

	// Server-defined URL for the resource.
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputepublicadvertisedprefix;gcpcomputepublicadvertisedprefixes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:annotations="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputePublicAdvertisedPrefix is the Schema for the ComputePublicAdvertisedPrefix API
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ComputePublicAdvertisedPrefix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputePublicAdvertisedPrefixSpec   `json:"spec,omitempty"`
	Status ComputePublicAdvertisedPrefixStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputePublicAdvertisedPrefixList contains a list of ComputePublicAdvertisedPrefix
type ComputePublicAdvertisedPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputePublicAdvertisedPrefix `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputePublicAdvertisedPrefix{}, &ComputePublicAdvertisedPrefixList{})
}
