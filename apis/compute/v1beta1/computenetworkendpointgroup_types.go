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

var ComputeNetworkEndpointGroupGVK = GroupVersion.WithKind("ComputeNetworkEndpointGroup")

// ComputeNetworkEndpointGroupSpec defines the desired state of ComputeNetworkEndpointGroup
// +kcc:spec:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeNetworkEndpointGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The default port used if the port number is not specified in the network endpoint.
	// +optional
	DefaultPort *int32 `json:"defaultPort,omitempty"`

	// Immutable. An optional description of this resource. Provide this property when you create the resource.
	// +optional
	Description *string `json:"description,omitempty"`

	// Location represents the geographical location of the ComputeNetworkEndpointGroup. Specify a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location"`

	/* Immutable. Type of network endpoints in this network endpoint group.
	NON_GCP_PRIVATE_IP_PORT is used for hybrid connectivity network
	endpoint groups (see https://cloud.google.com/load-balancing/docs/hybrid).
	Note that NON_GCP_PRIVATE_IP_PORT can only be used with Backend Services
	that 1) have the following load balancing schemes: EXTERNAL, EXTERNAL_MANAGED,
	INTERNAL_MANAGED, and INTERNAL_SELF_MANAGED and 2) support the RATE or
	CONNECTION balancing modes.

	Possible values include: GCE_VM_IP, GCE_VM_IP_PORT, and NON_GCP_PRIVATE_IP_PORT. Default value: "GCE_VM_IP_PORT" Possible values: ["GCE_VM_IP", "GCE_VM_IP_PORT", "NON_GCP_PRIVATE_IP_PORT"]. */
	// +optional
	NetworkEndpointType *string `json:"networkEndpointType,omitempty"`

	// The network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	NetworkRef *ComputeNetworkRef `json:"networkRef"`

	// The ComputeNetworkEndpointGroup name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional subnetwork to which all network endpoints in the NEG belong.
	// +optional
	SubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// ComputeNetworkEndpointGroupStatus defines the config connector machine state of ComputeNetworkEndpointGroup
type ComputeNetworkEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNetworkEndpointGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	// Number of network endpoints in the network endpoint group.
	// +optional
	Size *int64 `json:"size,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkendpointgroup;gcpcomputenetworkendpointgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
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
