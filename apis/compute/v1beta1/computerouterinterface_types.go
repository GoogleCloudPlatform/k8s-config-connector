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

var ComputeRouterInterfaceGVK = GroupVersion.WithKind("ComputeRouterInterface")

type ComputeVPNTunnelRef struct {
	/* Allowed value: The `selfLink` field of a `ComputeVPNTunnel` resource. */
	// +optional
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	// +optional
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ComputeRouterInterfaceSpec defines the desired state of ComputeRouterInterface
// +kcc:spec:proto=google.cloud.compute.v1.RouterInterface
type ComputeRouterInterfaceSpec struct {
	/* The InterconnectAttachment this interface belongs to. */
	// +optional
	InterconnectAttachmentRef *ComputeInterconnectAttachmentRef `json:"interconnectAttachmentRef,omitempty"`

	/* Immutable. The IP address and range of the interface.
	The IP range must be in the RFC3927 link-local IP space. Changing
	this forces a new interface to be created. */
	// +optional
	IPRange *string `json:"ipRange,omitempty"`

	/* The private IP address assigned to this interface. */
	// +optional
	PrivateIPAddressRef *ComputeAddressRef `json:"privateIpAddressRef,omitempty"`

	/* The interface the BGP peer is associated with. */
	// +optional
	RedundantInterfaceRef *ComputeRouterInterfaceRef `json:"redundantInterfaceRef,omitempty"`

	/* Immutable. The region this interface's router sits in.
	If not specified, the project region will be used. Changing this
	forces a new interface to be created. */
	// +required
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The router this interface belongs to. */
	// +required
	RouterRef ComputeRouterRef `json:"routerRef"`

	/* The subnetwork this interface belongs to. */
	// +optional
	SubnetworkRef *ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	/* The VPNTunnel this interface belongs to. */
	// +optional
	VPNTunnelRef *ComputeVPNTunnelRef `json:"vpnTunnelRef,omitempty"`
}

// ComputeRouterInterfaceStatus defines the config connector machine state of ComputeRouterInterface
type ComputeRouterInterfaceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputerouterinterface;gcpcomputerouterinterfaces
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRouterInterface is the Schema for the ComputeRouterInterface API
// +k8s:openapi-gen=true
type ComputeRouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRouterInterfaceSpec   `json:"spec,omitempty"`
	Status ComputeRouterInterfaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRouterInterfaceList contains a list of ComputeRouterInterface
type ComputeRouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRouterInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouterInterface{}, &ComputeRouterInterfaceList{})
}
