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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BeyondCorpAppGatewayGVK = GroupVersion.WithKind("BeyondCorpAppGateway")

// +kcc:spec:proto=google.cloud.beyondcorp.appgateways.v1.AppGateway
type BeyondCorpAppGatewaySpec struct {
	// Immutable. An arbitrary user-provided name for the AppGateway.
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The type of hosting used by the AppGateway.
	// Default value: "HOST_TYPE_UNSPECIFIED" Possible values: ["HOST_TYPE_UNSPECIFIED", "GCP_REGIONAL_MIG"].
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.host_type
	HostType *string `json:"hostType,omitempty"`

	// The project that this resource belongs to.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Immutable. The region of the AppGateway.
	Region string `json:"region"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The type of network connectivity used by the AppGateway.
	// Default value: "TYPE_UNSPECIFIED" Possible values: ["TYPE_UNSPECIFIED", "TCP_PROXY"].
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.type
	Type *string `json:"type,omitempty"`
}

// +kcc:status:proto=google.cloud.beyondcorp.appgateways.v1.AppGateway
type BeyondCorpAppGatewayStatus struct {
	// A list of connections allocated for the Gateway.
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.allocated_connections
	AllocatedConnections []AppGatewayAllocatedConnectionStatus `json:"allocatedConnections,omitempty"`

	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BeyondCorpAppGateway resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Represents the different states of an AppGateway.
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.state
	State *string `json:"state,omitempty"`

	// Server-defined URI for this resource.
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.uri
	Uri *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.beyondcorp.appgateways.v1.AppGateway.AllocatedConnection
type AppGatewayAllocatedConnectionStatus struct {
	// The ingress port of an allocated connection.
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.AllocatedConnection.ingress_port
	IngressPort *int32 `json:"ingressPort,omitempty"`

	// The PSC uri of an allocated connection.
	// +kcc:proto:field=google.cloud.beyondcorp.appgateways.v1.AppGateway.AllocatedConnection.psc_uri
	PscUri *string `json:"pscUri,omitempty"`
}

// BeyondCorpAppGatewayObservedState is the state of the BeyondCorpAppGateway resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.beyondcorp.appgateways.v1.AppGateway
type BeyondCorpAppGatewayObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbeyondcorpappgateway;gcpbeyondcorpappgateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BeyondCorpAppGateway is the Schema for the BeyondCorpAppGateway API
// +k8s:openapi-gen=true
type BeyondCorpAppGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BeyondCorpAppGatewaySpec   `json:"spec,omitempty"`
	Status BeyondCorpAppGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BeyondCorpAppGatewayList contains a list of BeyondCorpAppGateway
type BeyondCorpAppGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BeyondCorpAppGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BeyondCorpAppGateway{}, &BeyondCorpAppGatewayList{})
}
