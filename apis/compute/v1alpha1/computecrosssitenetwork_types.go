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

// API sources for ComputeCrossSiteNetwork, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/compute/docs/reference/rest

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeCrossSiteNetworkGVK = GroupVersion.WithKind("ComputeCrossSiteNetwork")

// ComputeCrossSiteNetworkSpec defines the desired state of ComputeCrossSiteNetwork
// +kcc:spec:proto=google.cloud.compute.v1.CrossSiteNetwork
// +kcc:required-from-proto
type ComputeCrossSiteNetworkSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ComputeCrossSiteNetwork name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// An optional description of the cross-site network.
	// +kcc:proto:field=google.cloud.compute.v1.CrossSiteNetwork.description
	Description *string `json:"description,omitempty"`
}

// ComputeCrossSiteNetworkStatus defines the config connector machine state of ComputeCrossSiteNetwork
type ComputeCrossSiteNetworkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeCrossSiteNetwork resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeCrossSiteNetworkObservedState `json:"observedState,omitempty"`
}

// ComputeCrossSiteNetworkObservedState is the state of the ComputeCrossSiteNetwork resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.CrossSiteNetwork
type ComputeCrossSiteNetworkObservedState struct {
	// Output only. [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CrossSiteNetwork.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Output only. [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CrossSiteNetwork.id
	ID *uint64 `json:"id,omitempty"`

	// Output only. [Output Only] Type of the resource. Always compute#crossSiteNetwork for cross-site networks.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CrossSiteNetwork.kind
	Kind *string `json:"kind,omitempty"`

	// Output only. [Output Only] Server-defined URL for the resource.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CrossSiteNetwork.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputecrosssitenetwork;gcpcomputecrosssitenetworks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeCrossSiteNetwork is the Schema for the ComputeCrossSiteNetwork API
// +k8s:openapi-gen=true
type ComputeCrossSiteNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeCrossSiteNetworkSpec   `json:"spec,omitempty"`
	Status ComputeCrossSiteNetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeCrossSiteNetworkList contains a list of ComputeCrossSiteNetwork
type ComputeCrossSiteNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeCrossSiteNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeCrossSiteNetwork{}, &ComputeCrossSiteNetworkList{})
}
