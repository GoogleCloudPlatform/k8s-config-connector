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

var ComputeRouteGVK = GroupVersion.WithKind("ComputeRoute")

// ComputeRouteSpec defines the desired state of ComputeRoute
// +kcc:spec:proto=google.cloud.compute.v1.Route
type ComputeRouteSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The ComputeRoute name. If not given, the metadata.name will be used.
	// +kcc:proto:field=google.cloud.compute.v1.Route.name
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. An optional description of this resource. Provide this property
	// when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Route.description
	Description *string `json:"description,omitempty"`

	// Immutable. The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Route.dest_range
	DestRange *string `json:"destRange,omitempty"`

	// The network that this route applies to.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Route.network
	NetworkRef *ComputeNetworkRef `json:"networkRef"`

	// Immutable. URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
	// * 'projects/project/global/gateways/default-internet-gateway'
	// * 'global/gateways/default-internet-gateway'
	// * The string 'default-internet-gateway'.
	// +kcc:proto:field=google.cloud.compute.v1.Route.next_hop_gateway
	NextHopGateway *string `json:"nextHopGateway,omitempty"`

	// A forwarding rule of type loadBalancingScheme=INTERNAL that should
	// handle matching packets. Note that this can only be used when the
	// destinationRange is a public (non-RFC 1918) IP CIDR range.
	// +kcc:proto:field=google.cloud.compute.v1.Route.next_hop_ilb
	NextHopILBRef *ForwardingRuleRef `json:"nextHopILBRef,omitempty"`

	// Instance that should handle matching packets.
	// +kcc:proto:field=google.cloud.compute.v1.Route.next_hop_instance
	NextHopInstanceRef *InstanceRef `json:"nextHopInstanceRef,omitempty"`

	// Immutable. Network IP address of an instance that should handle matching packets.
	// +kcc:proto:field=google.cloud.compute.v1.Route.next_hop_ip
	NextHopIP *string `json:"nextHopIp,omitempty"`

	// The ComputeVPNTunnel that should handle matching packets.
	// +kcc:proto:field=google.cloud.compute.v1.Route.next_hop_vpn_tunnel
	NextHopVPNTunnelRef *ComputeVPNTunnelRef `json:"nextHopVPNTunnelRef,omitempty"`

	// Immutable. The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	// +kcc:proto:field=google.cloud.compute.v1.Route.priority
	Priority *uint32 `json:"priority,omitempty"`

	// Immutable. A list of instance tags to which this route applies.
	// +kcc:proto:field=google.cloud.compute.v1.Route.tags
	Tags []string `json:"tags,omitempty"`
}

// ComputeRouteStatus defines the config connector machine state of ComputeRoute
type ComputeRouteStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRoute resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRouteObservedState `json:"observedState,omitempty"`

	// URL to a Network that should handle matching packets.
	// +kcc:proto:field=google.cloud.compute.v1.Route.next_hop_network
	NextHopNetwork *string `json:"nextHopNetwork,omitempty"`

	// Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Route.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeRouteObservedState is the state of the ComputeRoute resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Route
type ComputeRouteObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeroute;gcpcomputeroutes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRoute is the Schema for the ComputeRoute API
// +k8s:openapi-gen=true
type ComputeRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRouteSpec   `json:"spec,omitempty"`
	Status ComputeRouteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRouteList contains a list of ComputeRoute
type ComputeRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRoute{}, &ComputeRouteList{})
}
