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

var BeyondCorpAppConnectionGVK = GroupVersion.WithKind("BeyondCorpAppConnection")

// +kcc:proto=google.cloud.beyondcorp.appconnections.v1.AppConnection.ApplicationEndpoint
type AppConnectionApplicationEndpoint struct {
	// Hostname or IP address of the remote application endpoint.
	// +kcc:proto:field=google.cloud.beyondcorp.appconnections.v1.AppConnection.ApplicationEndpoint.host
	Host string `json:"host"`

	// Port of the remote application endpoint.
	// +kcc:proto:field=google.cloud.beyondcorp.appconnections.v1.AppConnection.ApplicationEndpoint.port
	Port int `json:"port"`
}

// +kcc:proto=google.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway
type AppConnectionGateway struct {
	// AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.
	// +kcc:proto:field=google.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway.app_gateway
	AppGateway string `json:"appGateway"`

	// Ingress port reserved on the gateways for this AppConnection, if not specified or zero, the default port is 19443.
	// +kcc:proto:field=google.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway.ingress_port
	IngressPort *int `json:"ingressPort,omitempty"`

	// The type of hosting used by the gateway. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	// for a list of possible values.
	// +kcc:proto:field=google.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway.type
	Type *string `json:"type,omitempty"`

	// Server-defined URI for this resource.
	// +kcc:proto:field=google.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:spec:proto=google.cloud.beyondcorp.appconnections.v1.AppConnection
type BeyondCorpAppConnectionSpec struct {
	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	// +required
	ApplicationEndpoint *AppConnectionApplicationEndpoint `json:"applicationEndpoint"`

	// List of AppConnectors that are authorised to be associated with this AppConnection.
	Connectors []string `json:"connectors,omitempty"`

	// An arbitrary user-provided name for the AppConnection.
	DisplayName *string `json:"displayName,omitempty"`

	// Gateway used by the AppConnection.
	Gateway *AppConnectionGateway `json:"gateway,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The region of the AppConnection.
	Region string `json:"region"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The type of network connectivity used by the AppConnection. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
	// for a list of possible values.
	Type *string `json:"type,omitempty"`
}

type BeyondCorpAppConnectionStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BeyondCorpAppConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// BeyondCorpAppConnectionObservedState is the state of the BeyondCorpAppConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.beyondcorp.appconnections.v1.AppConnection
type BeyondCorpAppConnectionObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbeyondcorpappconnection;gcpbeyondcorpappconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BeyondCorpAppConnection is the Schema for the beyondcorp API
// +k8s:openapi-gen=true
type BeyondCorpAppConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BeyondCorpAppConnectionSpec   `json:"spec,omitempty"`
	Status BeyondCorpAppConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BeyondCorpAppConnectionList contains a list of BeyondCorpAppConnection
type BeyondCorpAppConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BeyondCorpAppConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BeyondCorpAppConnection{}, &BeyondCorpAppConnectionList{})
}
