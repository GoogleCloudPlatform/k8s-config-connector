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

var ComputeExternalVPNGatewayGVK = GroupVersion.WithKind("ComputeExternalVPNGateway")

// ComputeExternalVPNGatewaySpec defines the desired state of ComputeExternalVPNGateway
// +kcc:spec:proto=google.cloud.compute.v1.ExternalVpnGateway
type ComputeExternalVPNGatewaySpec struct {
	// Immutable. An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.ExternalVpnGateway.description
	Description *string `json:"description,omitempty"`

	// Immutable. A list of interfaces on this external VPN gateway.
	// +kcc:proto:field=google.cloud.compute.v1.ExternalVpnGateway.interfaces
	Interfaces []ComputeExternalVPNGatewayInterface `json:"interface,omitempty"`

	// Immutable. Indicates the redundancy type of this external VPN gateway.
	// Possible values: ["FOUR_IPS_REDUNDANCY", "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"].
	// +kcc:proto:field=google.cloud.compute.v1.ExternalVpnGateway.redundancy_type
	RedundancyType *string `json:"redundancyType,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition.
	// When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ExternalVpnGatewayInterface
type ComputeExternalVPNGatewayInterface struct {
	// Immutable. The numeric ID for this interface. Allowed values are based on the redundancy type
	// of this external VPN gateway
	// * '0 - SINGLE_IP_INTERNALLY_REDUNDANT'
	// * '0, 1 - TWO_IPS_REDUNDANCY'
	// * '0, 1, 2, 3 - FOUR_IPS_REDUNDANCY'.
	// +kcc:proto:field=google.cloud.compute.v1.ExternalVpnGatewayInterface.id
	ID *int64 `json:"id,omitempty"`

	// Immutable. IP address of the interface in the external VPN gateway.
	// Only IPv4 is supported. This IP address can be either from
	// your on-premise gateway or another Cloud provider's VPN gateway,
	// it cannot be an IP address from Google Compute Engine.
	// +kcc:proto:field=google.cloud.compute.v1.ExternalVpnGatewayInterface.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`
}

// ComputeExternalVPNGatewayStatus defines the config connector machine state of ComputeExternalVPNGateway
// +kcc:status:proto=google.cloud.compute.v1.ExternalVpnGateway
type ComputeExternalVPNGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	// +kcc:proto:field=google.cloud.compute.v1.ExternalVpnGateway.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The self link of the ComputeExternalVPNGateway.
	// +kcc:proto:field=google.cloud.compute.v1.ExternalVpnGateway.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeexternalvpngateway;gcpcomputeexternalvpngateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeExternalVPNGateway is the Schema for the ComputeExternalVPNGateway API
// +k8s:openapi-gen=true
type ComputeExternalVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeExternalVPNGatewaySpec   `json:"spec,omitempty"`
	Status ComputeExternalVPNGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeExternalVPNGatewayList contains a list of ComputeExternalVPNGateway
type ComputeExternalVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeExternalVPNGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeExternalVPNGateway{}, &ComputeExternalVPNGatewayList{})
}
