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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeRouterInterfaceGVK = GroupVersion.WithKind("ComputeRouterInterface")

// ComputeRouterInterfaceSpec defines the desired state of ComputeRouterInterface
// +kcc:spec:proto=google.cloud.compute.v1.RouterInterface
type ComputeRouterInterfaceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource (region).
	Location string `json:"location"`

	// The ComputeRouterInterface name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The ComputeRouter that this interface belongs to.
	RouterRef *ComputeRouterRef `json:"routerRef"`

	// IP address and range of the interface.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.ip_range
	IpRange *string `json:"ipRange,omitempty"`

	// IP version of this interface.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.ip_version
	IpVersion *string `json:"ipVersion,omitempty"`

	// The regional private internal IP address.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.private_ip_address
	PrivateIpAddressRef *refsv1beta1.ComputeAddressRef `json:"privateIpAddressRef,omitempty"`

	// Name of the interface that will be redundant with the current interface.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.redundant_interface
	RedundantInterfaceRef *ComputeRouterInterfaceRef `json:"redundantInterfaceRef,omitempty"`

	// The subnetwork resource that this interface belongs to.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.subnetwork
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// The linked VPN tunnel.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.linked_vpn_tunnel
	VpnTunnelRef *ComputeVPNTunnelRef `json:"vpnTunnelRef,omitempty"`

	// The linked Interconnect attachment.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.linked_interconnect_attachment
	InterconnectAttachmentRef *ComputeInterconnectAttachmentRef `json:"interconnectAttachmentRef,omitempty"`
}

// ComputeRouterInterfaceStatus defines the config connector machine state of ComputeRouterInterface
type ComputeRouterInterfaceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRouterInterface resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRouterInterfaceObservedState `json:"observedState,omitempty"`
}

// ComputeRouterInterfaceObservedState is the state of the ComputeRouterInterface resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.RouterInterface
type ComputeRouterInterfaceObservedState struct {
	// The resource that configures and manages this interface.
	// +kcc:proto:field=google.cloud.compute.v1.RouterInterface.management_type
	ManagementType *string `json:"managementType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputerouterinterface;gcpcomputerouterinterfaces
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
