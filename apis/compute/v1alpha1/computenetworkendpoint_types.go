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

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkEndpointGVK = GroupVersion.WithKind("ComputeNetworkEndpoint")

// ComputeNetworkEndpointSpec defines the desired state of ComputeNetworkEndpoint
// +kcc:spec:proto=google.cloud.compute.v1.NetworkEndpoint
type ComputeNetworkEndpointSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. Zone where the containing network endpoint group is located.
	// +required
	Zone string `json:"zone"`

	// The network endpoint group to which this network endpoint belongs.
	// +required
	NetworkEndpointGroupRef *computev1beta1.ComputeNetworkEndpointGroupRef `json:"networkEndpointGroupRef"`

	// +optional
	InstanceRef *computev1beta1.InstanceRef `json:"instanceRef,omitempty"`

	// Immutable. IPv4 address of network endpoint. The IP address must belong to a VM in Compute Engine (either the primary IP or as part of an aliased IP range).
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpoint.ip_address
	IpAddress string `json:"ipAddress"`

	// Immutable. Optional. The port of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeNetworkEndpointStatus defines the config connector machine state of ComputeNetworkEndpoint
type ComputeNetworkEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkendpoint;gcpcomputenetworkendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkEndpoint is the Schema for the ComputeNetworkEndpoint API
// +k8s:openapi-gen=true
type ComputeNetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkEndpointSpec   `json:"spec,omitempty"`
	Status ComputeNetworkEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkEndpointList contains a list of ComputeNetworkEndpoint
type ComputeNetworkEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkEndpoint{}, &ComputeNetworkEndpointList{})
}
