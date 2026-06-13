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

var ComputeVPNGatewayGVK = GroupVersion.WithKind("ComputeVPNGateway")

// +kcc:proto=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface
type VPNGatewayVPNGatewayInterfaceSpec struct {
	// Immutable. Numeric identifier for this VPN interface associated with the VPN gateway.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface.id
	ID *uint32 `json:"id,omitempty"`

	// URL of the VLAN attachment (interconnectAttachment) resource for this VPN gateway interface.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface.interconnect_attachment
	InterconnectAttachmentRef *refsv1beta1.ComputeInterconnectAttachmentRef `json:"interconnectAttachmentRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface
type VPNGatewayVPNGatewayInterfaceStatus struct {
	// [Output Only] Numeric identifier for this VPN interface associated with the VPN gateway.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface.id
	ID *uint32 `json:"id,omitempty"`

	// URL of the VLAN attachment (interconnectAttachment) resource for this VPN gateway interface.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface.interconnect_attachment
	InterconnectAttachmentRef *refsv1beta1.ComputeInterconnectAttachmentRef `json:"interconnectAttachmentRef,omitempty"`

	// [Output Only] IP address for this VPN interface associated with the VPN gateway.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// [Output Only] IPv6 address for this VPN interface associated with the VPN gateway.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGatewayVpnGatewayInterface.ipv6_address
	IPV6Address *string `json:"ipv6Address,omitempty"`
}

// ComputeVPNGatewaySpec defines the desired state of ComputeVPNGateway
// +kcc:spec:proto=google.cloud.compute.v1.VpnGateway
type ComputeVPNGatewaySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeVPNGateway name. If not given, the metadata.name will be used.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.name
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.description
	Description *string `json:"description,omitempty"`

	// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.gateway_ip_version
	GatewayIPVersion *string `json:"gatewayIPVersion,omitempty"`

	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.network
	NetworkRef *ComputeNetworkRef `json:"networkRef"`

	// The stack type for this VPN gateway to identify the IP protocols that are enabled. Possible values are: IPV4_ONLY, IPV4_IPV6, IPV6_ONLY.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.stack_type
	StackType *string `json:"stackType,omitempty"`

	// The list of VPN interfaces associated with this VPN gateway.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.vpn_interfaces
	VPNInterfaces []VPNGatewayVPNGatewayInterfaceSpec `json:"vpnInterfaces,omitempty"`
}

// ComputeVPNGatewayStatus defines the config connector machine state of ComputeVPNGateway
type ComputeVPNGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeVPNGateway resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeVPNGatewayObservedState `json:"observedState,omitempty"`
}

// ComputeVPNGatewayObservedState is the state of the ComputeVPNGateway resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.VpnGateway
type ComputeVPNGatewayObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of resource. Always compute#vpnGateway for VPN gateways.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the VPN gateway resides.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// The list of VPN interfaces associated with this VPN gateway.
	// +kcc:proto:field=google.cloud.compute.v1.VpnGateway.vpn_interfaces
	VPNInterfaces []VPNGatewayVPNGatewayInterfaceStatus `json:"vpnInterfaces,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputevpngateway;gcpcomputevpngateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeVPNGateway is the Schema for the ComputeVPNGateway API
// +k8s:openapi-gen=true
type ComputeVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeVPNGatewaySpec   `json:"spec,omitempty"`
	Status ComputeVPNGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeVPNGatewayList contains a list of ComputeVPNGateway
type ComputeVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeVPNGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeVPNGateway{}, &ComputeVPNGatewayList{})
}
