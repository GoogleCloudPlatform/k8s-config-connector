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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ComputeVPNTunnelGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "ComputeVPNTunnel",
}

type VPNTunnelSharedSecret struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *VPNTunnelValueFrom `json:"valueFrom,omitempty"`
}

type VPNTunnelValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *k8sv1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

// ComputeVPNTunnelSpec defines the desired state of ComputeVPNTunnel
// +kcc:spec:proto=google.cloud.compute.v1.VpnTunnel
type ComputeVPNTunnelSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputeVPNTunnel name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. An optional description of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.description
	Description *string `json:"description,omitempty"`

	// Immutable. IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2. Default version is 2.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.ike_version
	IkeVersion *int32 `json:"ikeVersion,omitempty"`

	// Immutable. Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.local_traffic_selector
	LocalTrafficSelector []string `json:"localTrafficSelector,omitempty"`

	// Immutable. The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.peer_external_gateway_interface
	PeerExternalGatewayInterface *int32 `json:"peerExternalGatewayInterface,omitempty"`

	// The peer side external VPN gateway to which this VPN tunnel is connected.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.peer_external_gateway
	PeerExternalGatewayRef *refsv1beta1.ComputeExternalVPNGatewayRef `json:"peerExternalGatewayRef,omitempty"`

	// The peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will automatically use the same VPN gateway interface ID in the peer GCP VPN gateway.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.peer_gcp_gateway
	PeerGCPGatewayRef *refsv1beta1.ComputeVPNGatewayRef `json:"peerGCPGatewayRef,omitempty"`

	// Immutable. IP address of the peer VPN gateway. Only IPv4 is supported.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.peer_ip
	PeerIp *string `json:"peerIp,omitempty"`

	// Immutable. Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.remote_traffic_selector
	RemoteTrafficSelector []string `json:"remoteTrafficSelector,omitempty"`

	// The router to be used for dynamic routing.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.router
	RouterRef *refsv1beta1.ComputeRouterRef `json:"routerRef,omitempty"`

	// Immutable. Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.shared_secret
	SharedSecret VPNTunnelSharedSecret `json:"sharedSecret"`

	// The ComputeTargetVPNGateway with which this VPN tunnel is associated.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.target_vpn_gateway
	TargetVPNGatewayRef *refsv1beta1.ComputeTargetVPNGatewayRef `json:"targetVPNGatewayRef,omitempty"`

	// Immutable. The interface ID of the VPN gateway with which this VPN tunnel is associated.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.vpn_gateway_interface
	VpnGatewayInterface *int32 `json:"vpnGatewayInterface,omitempty"`

	// The ComputeVPNGateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway resource is created.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.vpn_gateway
	VpnGatewayRef *refsv1beta1.ComputeVPNGatewayRef `json:"vpnGatewayRef,omitempty"`
}

// ComputeVPNTunnelStatus defines the config connector machine state of ComputeVPNTunnel
// +kcc:status:proto=google.cloud.compute.v1.VpnTunnel
type ComputeVPNTunnelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeVPNTunnel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeVPNTunnelObservedState `json:"observedState,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Detailed status message for the VPN tunnel.
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.detailed_status
	DetailedStatus *string `json:"detailedStatus,omitempty"`

	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// The SelfLink for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Hash of the shared secret.
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.shared_secret_hash
	SharedSecretHash *string `json:"sharedSecretHash,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.VpnTunnel.id
	TunnelId *string `json:"tunnelId,omitempty"`
}

// ComputeVPNTunnelObservedState is the state of the ComputeVPNTunnel resource as most recently observed in GCP.
type ComputeVPNTunnelObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputevpntunnel;gcpcomputevpntunnels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
