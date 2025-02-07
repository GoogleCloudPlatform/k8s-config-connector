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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BaremetalsolutionNetworkGVK = GroupVersion.WithKind("BaremetalsolutionNetwork")

// BaremetalsolutionNetworkSpec defines the desired state of BaremetalsolutionNetwork
// +kcc:proto=google.cloud.baremetalsolution.v2.Network
type BaremetalsolutionNetworkSpec struct {
	// The BaremetalsolutionNetwork name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// BaremetalsolutionNetworkStatus defines the config connector machine state of BaremetalsolutionNetwork
type BaremetalsolutionNetworkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BaremetalsolutionNetwork resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BaremetalsolutionNetworkObservedState `json:"observedState,omitempty"`
}

// BaremetalsolutionNetworkObservedState is the state of the BaremetalsolutionNetwork resource as most recently observed in GCP.
// +kcc:proto=google.cloud.baremetalsolution.v2.Network
type BaremetalsolutionNetworkObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpbaremetalsolutionnetwork;gcpbaremetalsolutionnetworks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BaremetalsolutionNetwork is the Schema for the BaremetalsolutionNetwork API
// +k8s:openapi-gen=true
type BaremetalsolutionNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BaremetalsolutionNetworkSpec   `json:"spec,omitempty"`
	Status BaremetalsolutionNetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BaremetalsolutionNetworkList contains a list of BaremetalsolutionNetwork
type BaremetalsolutionNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaremetalsolutionNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaremetalsolutionNetwork{}, &BaremetalsolutionNetworkList{})
}
