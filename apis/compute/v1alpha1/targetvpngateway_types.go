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

var ComputeTargetVPNGatewayGVK = GroupVersion.WithKind("ComputeTargetVPNGateway")

// ComputeTargetVPNGatewaySpec defines the desired state of ComputeTargetVPNGateway
// +kcc:spec:proto=google.cloud.compute.v1.TargetVpnGateway
type ComputeTargetVPNGatewaySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputeTargetVPNGateway name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.description
	Description *string `json:"description,omitempty"`

	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef"`

	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// ComputeTargetVPNGatewayStatus defines the config connector machine state of ComputeTargetVPNGateway
type ComputeTargetVPNGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetVPNGateway resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *ComputeTargetVPNGatewayObservedState `json:"observedState,omitempty"`
}

// ComputeTargetVPNGatewayObservedState is the state of the ComputeTargetVPNGateway resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.TargetVpnGateway
type ComputeTargetVPNGatewayObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] A list of URLs to the ForwardingRule resources. ForwardingRules are created using compute.forwardingRules.insert and associated with a VPN gateway.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.forwarding_rules
	ForwardingRules []string `json:"forwardingRules,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of resource. Always compute#targetVpnGateway for target VPN gateways.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the target VPN gateway resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] The status of the VPN gateway, which can be one of the following: CREATING, READY, FAILED, or DELETING.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.status
	Status *string `json:"status,omitempty"`

	// [Output Only] A list of URLs to VpnTunnel resources. VpnTunnels are created using the compute.vpntunnels.insert method and associated with a VPN gateway.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.tunnels
	Tunnels []string `json:"tunnels,omitempty"`

	// A fingerprint for the labels being applied to this TargetVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a TargetVpnGateway.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetvpngateway;gcpcomputetargetvpngateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetVPNGateway is the Schema for the ComputeTargetVPNGateway API
// +k8s:openapi-gen=true
type ComputeTargetVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetVPNGatewaySpec   `json:"spec,omitempty"`
	Status ComputeTargetVPNGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetVPNGatewayList contains a list of ComputeTargetVPNGateway
type ComputeTargetVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetVPNGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetVPNGateway{}, &ComputeTargetVPNGatewayList{})
}
