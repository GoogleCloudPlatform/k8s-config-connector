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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkGVK = GroupVersion.WithKind("ComputeNetwork")

// ComputeNetworkSpec defines the desired state of ComputeNetwork
// +kcc:spec:proto=google.cloud.compute.v1.Network
type ComputeNetworkSpec struct {
	// Immutable. When set to 'true', the network is created in "auto subnet mode" and
	// it will create a subnet for each region automatically across the
	// '10.128.0.0/9' address range.
	//
	// When set to 'false', the network is created in "custom subnet mode" so
	// the user can explicitly connect subnetwork resources.
	// +kcc:proto:field=google.cloud.compute.v1.Network.auto_create_subnetworks
	AutoCreateSubnetworks *bool `json:"autoCreateSubnetworks,omitempty"`

	// If set to 'true', default routes ('0.0.0.0/0') will be deleted
	// immediately after network creation. Defaults to 'false'.
	DeleteDefaultRoutesOnCreate *bool `json:"deleteDefaultRoutesOnCreate,omitempty"`

	// Immutable. An optional description of this resource. The resource must be
	// recreated to modify this field.
	// +kcc:proto:field=google.cloud.compute.v1.Network.description
	Description *string `json:"description,omitempty"`

	// Enable ULA internal ipv6 on this network. Enabling this feature will assign
	// a /48 from google defined ULA prefix fd20::/20.
	// +kcc:proto:field=google.cloud.compute.v1.Network.enable_ula_internal_ipv6
	EnableUlaInternalIpv6 *bool `json:"enableUlaInternalIpv6,omitempty"`

	// Immutable. When enabling ula internal ipv6, caller optionally can specify the /48 range
	// they want from the google defined ULA prefix fd20::/20. The input must be a
	// valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will
	// fail if the speficied /48 is already in used by another resource.
	// If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field.
	// +kcc:proto:field=google.cloud.compute.v1.Network.internal_ipv6_range
	InternalIpv6Range *string `json:"internalIpv6Range,omitempty"`

	// Immutable. Maximum Transmission Unit in bytes. The default value is 1460 bytes.
	// The minimum value for this field is 1300 and the maximum value is 8896 bytes (jumbo frames).
	// Note that packets larger than 1500 bytes (standard Ethernet) can be subject to TCP-MSS clamping or dropped
	// with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or other VPCs
	// with varying MTUs.
	// +kcc:proto:field=google.cloud.compute.v1.Network.mtu
	Mtu *int32 `json:"mtu,omitempty"`

	// Set the order that Firewall Rules and Firewall Policies
	// are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible
	// values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"].
	// +kcc:proto:field=google.cloud.compute.v1.Network.network_firewall_policy_enforcement_order
	NetworkFirewallPolicyEnforcementOrder *string `json:"networkFirewallPolicyEnforcementOrder,omitempty"`

	// Immutable. Optional. The name of the resource. Used for
	// creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// The network-wide routing mode to use. If set to 'REGIONAL', this
	// network's cloud routers will only advertise routes with subnetworks
	// of this network in the same region as the router. If set to 'GLOBAL',
	// this network's cloud routers will advertise routes with all
	// subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"].
	// +kcc:proto:field=google.cloud.compute.v1.Network.routing_config.routing_mode
	RoutingMode *string `json:"routingMode,omitempty"`
}

// ComputeNetworkStatus defines the config connector machine state of ComputeNetwork
// +kcc:status:proto=google.cloud.compute.v1.Network
type ComputeNetworkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The gateway address for default routing out of the network. This value
	// is selected by GCP.
	// +kcc:proto:field=google.cloud.compute.v1.Network.gateway_i_pv4
	GatewayIpv4 *string `json:"gatewayIpv4,omitempty"`

	// A unique specifier for the ComputeNetwork resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Network.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetwork;gcpcomputenetworks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetwork is the Schema for the ComputeNetwork API
// +k8s:openapi-gen=true
type ComputeNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkSpec   `json:"spec,omitempty"`
	Status ComputeNetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkList contains a list of ComputeNetwork
type ComputeNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetwork{}, &ComputeNetworkList{})
}
