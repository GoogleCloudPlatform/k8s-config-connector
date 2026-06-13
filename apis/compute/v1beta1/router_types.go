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

var ComputeRouterGVK = GroupVersion.WithKind("ComputeRouter")

// ComputeRouterSpec defines the desired state of ComputeRouter
// +kcc:spec:proto=google.cloud.compute.v1.Router
type ComputeRouterSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputeRouter name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	/* BGP information specific to this router. */
	// +kcc:proto:field=google.cloud.compute.v1.Router.bgp
	BGP *RouterBGP `json:"bgp,omitempty"`

	/* An optional description of this resource. */
	// +kcc:proto:field=google.cloud.compute.v1.Router.description
	Description *string `json:"description,omitempty"`

	/* Immutable. Indicates if a router is dedicated for use with encrypted VLAN
	   attachments (interconnectAttachments). */
	// +kcc:proto:field=google.cloud.compute.v1.Router.encrypted_interconnect_router
	EncryptedInterconnectRouter *bool `json:"encryptedInterconnectRouter,omitempty"`

	/* A reference to the network to which this router belongs. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Router.network
	NetworkRef *ComputeNetworkRef `json:"networkRef"`
}

// +kcc:proto=google.cloud.compute.v1.RouterBgp
type RouterBGP struct {
	// User-specified flag to indicate which mode to use for advertisement. The options are DEFAULT or CUSTOM.
	// +kubebuilder:validation:Enum=DEFAULT;CUSTOM
	// +kcc:proto:field=google.cloud.compute.v1.RouterBgp.advertise_mode
	AdvertiseMode *string `json:"advertiseMode,omitempty"`

	// User-specified list of prefix groups to advertise in custom mode. This field can only be populated if advertise_mode is CUSTOM and is advertised to all peers of the router. These groups will be advertised in addition to any specified prefixes. Leave this field blank to advertise no custom groups.
	// +kcc:proto:field=google.cloud.compute.v1.RouterBgp.advertised_groups
	AdvertisedGroups []string `json:"advertisedGroups,omitempty"`

	// User-specified list of individual IP ranges to advertise in custom mode. This field can only be populated if advertise_mode is CUSTOM and is advertised to all peers of the router. These IP ranges will be advertised in addition to any specified groups. Leave this field blank to advertise no custom IP ranges.
	// +kcc:proto:field=google.cloud.compute.v1.RouterBgp.advertised_ip_ranges
	AdvertisedIPRanges []RouterAdvertisedIPRange `json:"advertisedIPRanges,omitempty"`

	// Local BGP Autonomous System Number (ASN). Must be an RFC6996 private ASN, either 16-bit or 32-bit. The value will be fixed for this router resource. All VPN tunnels that link to this router will have the same local ASN.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterBgp.asn
	Asn *uint32 `json:"asn"`

	// The interval in seconds between BGP keepalive messages that are sent to the peer. Hold time is three times the interval at which keepalive messages are sent, and the hold time is the maximum number of seconds allowed to elapse between successive keepalive messages that BGP receives from a peer. BGP will use the smaller of either the local hold time value or the peer's hold time value as the hold time for the BGP connection between the two peers. If set, this value must be between 20 and 60. The default is 20.
	// +kcc:proto:field=google.cloud.compute.v1.RouterBgp.keepalive_interval
	KeepaliveInterval *uint32 `json:"keepaliveInterval,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.RouterAdvertisedIpRange
type RouterAdvertisedIPRange struct {
	// User-specified description for the IP range.
	// +kcc:proto:field=google.cloud.compute.v1.RouterAdvertisedIpRange.description
	Description *string `json:"description,omitempty"`

	// The IP range to advertise. The value must be a CIDR-formatted string.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RouterAdvertisedIpRange.range
	Range *string `json:"range"`
}

// ComputeRouterStatus defines the config connector machine state of ComputeRouter
type ComputeRouterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRouter resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRouterObservedState `json:"observedState,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.Router.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Router.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeRouterObservedState is the state of the ComputeRouter resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Router
type ComputeRouterObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputerouter;gcpcomputerouters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRouter is the Schema for the ComputeRouter API
// +k8s:openapi-gen=true
type ComputeRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRouterSpec   `json:"spec,omitempty"`
	Status ComputeRouterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRouterList contains a list of ComputeRouter
type ComputeRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRouter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouter{}, &ComputeRouterList{})
}
