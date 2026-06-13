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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkEndpointGroupGVK = GroupVersion.WithKind("ComputeNetworkEndpointGroup")

// ComputeNetworkEndpointGroupSpec defines the desired state of ComputeNetworkEndpointGroup
// +kcc:spec:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeNetworkEndpointGroupSpec struct {
	// Immutable. The default port used if the port number is not specified in the network endpoint.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.default_port
	DefaultPort *int32 `json:"defaultPort,omitempty"`

	// Immutable. An optional description of this resource. Provide this property when you create the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.description
	Description *string `json:"description,omitempty"`

	// Location represents the geographical location of the ComputeNetworkEndpointGroup. Specify a zone name.
	// +required
	Location string `json:"location"`

	// Immutable. Type of network endpoints in this network endpoint group.
	// Possible values include: GCE_VM_IP, GCE_VM_IP_PORT, and NON_GCP_PRIVATE_IP_PORT. Default value: "GCE_VM_IP_PORT"
	// +optional
	// +kubebuilder:validation:Enum=GCE_VM_IP;GCE_VM_IP_PORT;NON_GCP_PRIVATE_IP_PORT
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.network_endpoint_type
	NetworkEndpointType *string `json:"networkEndpointType,omitempty"`

	// The network to which all network endpoints in the NEG belong.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.network
	NetworkRef *ComputeNetworkRef `json:"networkRef"`

	// The ComputeNetworkEndpointGroup name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional subnetwork to which all network endpoints in the NEG belong.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.subnetwork
	SubnetworkRef *ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// ComputeNetworkEndpointGroupStatus defines the config connector machine state of ComputeNetworkEndpointGroup
type ComputeNetworkEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNetworkEndpointGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Server-defined URL for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Number of network endpoints in the network endpoint group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.size
	Size *int32 `json:"size,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeNetworkEndpointGroupObservedState `json:"observedState,omitempty"`
}

// ComputeNetworkEndpointGroupObservedState is the state of the ComputeNetworkEndpointGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeNetworkEndpointGroupObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkendpointgroup;gcpcomputenetworkendpointgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkEndpointGroup is the Schema for the ComputeNetworkEndpointGroup API
// +k8s:openapi-gen=true
type ComputeNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkEndpointGroupList contains a list of ComputeNetworkEndpointGroup
type ComputeNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkEndpointGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkEndpointGroup{}, &ComputeNetworkEndpointGroupList{})
}
