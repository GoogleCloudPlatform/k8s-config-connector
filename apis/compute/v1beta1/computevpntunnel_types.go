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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeVPNTunnelGVK = GroupVersion.WithKind("ComputeVPNTunnel")

// ComputeVPNTunnelSpec defines the desired state of ComputeVPNTunnel
// +kcc:spec:proto=google.cloud.compute.v1.VpnTunnel
type ComputeVPNTunnelSpec struct {
	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. IKE protocol version to use when establishing the VPN tunnel with
	peer VPN gateway.
	Acceptable IKE versions are 1 or 2. Default version is 2. */
	// +optional
	IkeVersion *int64 `json:"ikeVersion,omitempty"`

	/* Immutable. Local traffic selector to use when establishing the VPN tunnel with
	peer VPN gateway. The value should be a CIDR formatted string,
	for example '192.168.0.0/16'. The ranges should be disjoint.
	Only IPv4 is supported. */
	// +optional
	LocalTrafficSelector []string `json:"localTrafficSelector,omitempty"`

	/* Immutable. The interface ID of the external VPN gateway to which this VPN tunnel is connected. */
	// +optional
	PeerExternalGatewayInterface *int64 `json:"peerExternalGatewayInterface,omitempty"`

	/* The peer side external VPN gateway to which this VPN tunnel
	is connected. */
	// +optional
	PeerExternalGatewayRef *ComputeExternalVPNGatewayRef `json:"peerExternalGatewayRef,omitempty"`

	/* The peer side HA GCP VPN gateway to which this VPN tunnel is
	connected. If provided, the VPN tunnel will automatically use the
	same VPN gateway interface ID in the peer GCP VPN gateway. */
	// +optional
	PeerGCPGatewayRef *ComputeVPNGatewayRef `json:"peerGCPGatewayRef,omitempty"`

	/* Immutable. IP address of the peer VPN gateway. Only IPv4 is supported. */
	// +optional
	PeerIp *string `json:"peerIp,omitempty"`

	/* Immutable. The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'. */
	Region string `json:"region"`

	/* Immutable. Remote traffic selector to use when establishing the VPN tunnel with
	peer VPN gateway. The value should be a CIDR formatted string,
	for example '192.168.0.0/16'. The ranges should be disjoint.
	Only IPv4 is supported. */
	// +optional
	RemoteTrafficSelector []string `json:"remoteTrafficSelector,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The router to be used for dynamic routing. */
	// +optional
	RouterRef *ComputeRouterRef `json:"routerRef,omitempty"`

	/* Immutable. Shared secret used to set the secure session between the Cloud VPN
	gateway and the peer VPN gateway. */
	SharedSecret secret.Legacy `json:"sharedSecret"`

	/* The ComputeTargetVPNGateway with which this VPN tunnel is
	associated. */
	// +optional
	TargetVPNGatewayRef *refsv1beta1.ComputeTargetVPNGatewayRef `json:"targetVPNGatewayRef,omitempty"`

	/* Immutable. The interface ID of the VPN gateway with which this VPN tunnel is associated. */
	// +optional
	VpnGatewayInterface *int64 `json:"vpnGatewayInterface,omitempty"`

	/* The ComputeVPNGateway with which this VPN tunnel is associated.
	This must be used if a High Availability VPN gateway resource is
	created. */
	// +optional
	VpnGatewayRef *ComputeVPNGatewayRef `json:"vpnGatewayRef,omitempty"`
}

// ComputeVPNTunnelStatus defines the config connector machine state of ComputeVPNTunnel
type ComputeVPNTunnelStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeVPNTunnel's current state. */
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* Detailed status message for the VPN tunnel. */
	// +optional
	DetailedStatus *string `json:"detailedStatus,omitempty"`

	/* The fingerprint used for optimistic locking of this resource.  Used
	internally during updates. */
	// +optional
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* Hash of the shared secret. */
	// +optional
	SharedSecretHash *string `json:"sharedSecretHash,omitempty"`

	/* The unique identifier for the resource. This identifier is defined by the server. */
	// +optional
	TunnelId *string `json:"tunnelId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputevpntunnel;gcpcomputevpntunnels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeVPNTunnel is the Schema for the ComputeVPNTunnel API
// +k8s:openapi-gen=true
type ComputeVPNTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeVPNTunnelSpec   `json:"spec,omitempty"`
	Status ComputeVPNTunnelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeVPNTunnelList contains a list of ComputeVPNTunnel
type ComputeVPNTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeVPNTunnel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeVPNTunnel{}, &ComputeVPNTunnelList{})
}
