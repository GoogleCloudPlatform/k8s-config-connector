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

var ComputeRouterPeerGVK = GroupVersion.WithKind("ComputeRouterPeer")

type RouterpeerAdvertisedIpRanges struct {
	// +optional
	Description *string `json:"description,omitempty"`

	Range string `json:"range"`
}

type RouterpeerBfd struct {
	// +optional
	MinReceiveInterval *int64 `json:"minReceiveInterval,omitempty"`

	// +optional
	MinTransmitInterval *int64 `json:"minTransmitInterval,omitempty"`

	// +optional
	Multiplier *int64 `json:"multiplier,omitempty"`

	SessionInitializationMode string `json:"sessionInitializationMode"`
}

type RouterpeerIpAddress struct {
	// +optional
	External *string `json:"external,omitempty"`
}

// ComputeRouterPeerSpec defines the desired state of ComputeRouterPeer
// +kcc:spec:proto=google.cloud.compute.v1.RouterBgpPeer
type ComputeRouterPeerSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. Region where the router and BgpPeer reside.
	Region string `json:"region"`

	// The ComputeRouterPeer name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	AdvertiseMode *string `json:"advertiseMode,omitempty"`

	// +optional
	AdvertisedGroups []string `json:"advertisedGroups,omitempty"`

	// +optional
	AdvertisedIpRanges []RouterpeerAdvertisedIpRanges `json:"advertisedIpRanges,omitempty"`

	// +optional
	AdvertisedRoutePriority *int64 `json:"advertisedRoutePriority,omitempty"`

	// +optional
	Bfd *RouterpeerBfd `json:"bfd,omitempty"`

	// +optional
	Enable *bool `json:"enable,omitempty"`

	// +optional
	EnableIpv6 *bool `json:"enableIpv6,omitempty"`

	// +optional
	IpAddress *RouterpeerIpAddress `json:"ipAddress,omitempty"`

	// +optional
	Ipv6NexthopAddress *string `json:"ipv6NexthopAddress,omitempty"`

	PeerAsn int64 `json:"peerAsn"`

	// +optional
	PeerIpAddress *string `json:"peerIpAddress,omitempty"`

	// +optional
	PeerIpv6NexthopAddress *string `json:"peerIpv6NexthopAddress,omitempty"`

	// +optional
	RouterApplianceInstanceRef *InstanceRef `json:"routerApplianceInstanceRef,omitempty"`

	RouterInterfaceRef ComputeRouterInterfaceRef `json:"routerInterfaceRef"`

	RouterRef ComputeRouterRef `json:"routerRef"`
}

// ComputeRouterPeerStatus defines the config connector machine state of ComputeRouterPeer
type ComputeRouterPeerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRouterPeer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// +optional
	ManagementType *string `json:"managementType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputerouterpeer;gcpcomputerouterpeers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRouterPeer is the Schema for the ComputeRouterPeer API
// +k8s:openapi-gen=true
type ComputeRouterPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRouterPeerSpec   `json:"spec,omitempty"`
	Status ComputeRouterPeerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRouterPeerList contains a list of ComputeRouterPeer
type ComputeRouterPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRouterPeer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouterPeer{}, &ComputeRouterPeerList{})
}
