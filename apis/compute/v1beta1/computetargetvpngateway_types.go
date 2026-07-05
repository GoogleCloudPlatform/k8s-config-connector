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

var ComputeTargetVPNGatewayGVK = GroupVersion.WithKind("ComputeTargetVPNGateway")

// ComputeTargetVPNGatewaySpec defines the desired state of ComputeTargetVPNGateway
// +kcc:spec:proto=google.cloud.compute.v1.TargetVpnGateway
type ComputeTargetVPNGatewaySpec struct {
	/* Immutable. An optional description of this resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.description
	Description *string `json:"description,omitempty"`

	/* The network this VPN gateway is accepting traffic for. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.network
	NetworkRef *ComputeNetworkRef `json:"networkRef"`

	/* Immutable. The region this gateway should sit in. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.region
	Region *string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for
	   creation and acquisition. When unset, the value of `metadata.name`
	   is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeTargetVPNGatewayStatus defines the config connector machine state of ComputeTargetVPNGateway
// +kcc:status:proto=google.cloud.compute.v1.TargetVpnGateway
type ComputeTargetVPNGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* The unique identifier for the resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.id
	ID *uint64 `json:"gatewayId,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetVpnGateway.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetvpngateway;gcpcomputetargetvpngateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
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
