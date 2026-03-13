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

var ComputePublicDelegatedPrefixGVK = GroupVersion.WithKind("ComputePublicDelegatedPrefix")

// ComputePublicDelegatedPrefixSpec defines the desired state of ComputePublicDelegatedPrefix
// +kcc:spec:proto=google.cloud.compute.v1.PublicDelegatedPrefix
type ComputePublicDelegatedPrefixSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputePublicDelegatedPrefix name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty"`

	// The IP address range, in CIDR format, represented by this public delegated prefix.
	IPCidrRange *string `json:"ipCidrRange,omitempty"`

	// The URL of parent prefix. Either Create constructed or specified by the user.
	ParentPrefixRef *ComputePublicDelegatedPrefixParentPrefixRef `json:"parentPrefixRef,omitempty"`

	// AllocatablePrefixes: The list of sub-prefixes that are delegated IPv4
	// sub-prefixes. The length of the delegated sub-prefixes must be
	// /64.
	// +optional
	AllocatablePrefixes []string `json:"allocatablePrefixes,omitempty"`
}

type ComputePublicDelegatedPrefixParentPrefixRef struct {
	// The `external` field of a `ComputePublicAdvertisedPrefix` or `ComputePublicDelegatedPrefix` resource.
	External string `json:"external,omitempty"`
	// Kind of the referent. Allowed values: ComputePublicAdvertisedPrefix (default), ComputePublicDelegatedPrefix
	// +optional
	Kind string `json:"kind,omitempty"`
	// Name of the referent.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace of the referent.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ComputePublicDelegatedPrefixStatus defines the config connector machine state of ComputePublicDelegatedPrefix
type ComputePublicDelegatedPrefixStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputePublicDelegatedPrefix resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputePublicDelegatedPrefixObservedState `json:"observedState,omitempty"`
}

// ComputePublicDelegatedPrefixObservedState is the state of the ComputePublicDelegatedPrefix resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.PublicDelegatedPrefix
type ComputePublicDelegatedPrefixObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputepublicdelegatedprefix;gcpcomputepublicdelegatedprefixes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:annotations="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputePublicDelegatedPrefix is the Schema for the ComputePublicDelegatedPrefix API
// +k8s:openapi-gen=true
type ComputePublicDelegatedPrefix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputePublicDelegatedPrefixSpec   `json:"spec,omitempty"`
	Status ComputePublicDelegatedPrefixStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputePublicDelegatedPrefixList contains a list of ComputePublicDelegatedPrefix
type ComputePublicDelegatedPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputePublicDelegatedPrefix `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputePublicDelegatedPrefix{}, &ComputePublicDelegatedPrefixList{})
}
