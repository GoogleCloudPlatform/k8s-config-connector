// Copyright 2025 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ServiceNetworkingConnectionGVK = GroupVersion.WithKind("ServiceNetworkingConnection")

// ServiceNetworkingConnectionSpec defines the desired state of ServiceNetworkingConnection
// +kcc:spec:proto=mockgcp.cloud.servicenetworking.v1.Connection
type ServiceNetworkingConnectionSpec struct {
	/* NOTYET: Terraform
	// The ServiceNetworkingConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	*/

	// Required. The service consumer's VPC network that's connected with service producer network.
	// +kcc:proto:field=mockgcp.cloud.servicenetworking.v1.Connection.network
	// +required
	NetworkRef *v1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// One or more allocated IP address ranges for this service producer of type `PEERING`.
	// Note that invoking CreateConnection method with a different range when connection is already established will not modify already provisioned service producer subnetworks.
	// If CreateConnection method is invoked repeatedly to reconnect when peering connection had been disconnected on the consumer side, leaving this field empty will restore previously allocated IP ranges.
	// +kcc:proto:field=mockgcp.cloud.servicenetworking.v1.Connection.reserved_peering_ranges
	// +required
	ReservedPeeringRanges []*v1beta1.ComputeAddressRef `json:"reservedPeeringRanges,omitempty"`

	/* Immutable. Provider peering service that is managing peering connectivity for a service provider organization. For Google services that support this functionality it is 'servicenetworking.googleapis.com'. */
	// +required
	Service string `json:"service,omitempty"`
}

// ServiceNetworkingConnectionStatus defines the config connector machine state of ServiceNetworkingConnection
type ServiceNetworkingConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET: Terraform
	// A unique specifier for the ServiceNetworkingConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	/*NOTYET: Terraform
	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ServiceNetworkingConnectionObservedState `json:"observedState,omitempty"`
	*/

	// The name of the VPC Network Peering connection that was created by the service producer.
	// +kcc:proto:field=mockgcp.cloud.servicenetworking.v1.Connection.peering
	Peering *string `json:"peering,omitempty"`
}

// ServiceNetworkingConnectionObservedState is the state of the ServiceNetworkingConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=mockgcp.cloud.servicenetworking.v1.Connection
type ServiceNetworkingConnectionObservedState struct {
	/* NOTYET: Terraform
	// Output only. The name of the VPC Network Peering connection that was created by the service producer.
	// +kcc:proto:field=mockgcp.cloud.servicenetworking.v1.Connection.peering
	Peering *string `json:"peering,omitempty"`

	// Output only. The name of the peering service that's associated with this connection, in the following format: `services/{service name}`.
	// +kcc:proto:field=mockgcp.cloud.servicenetworking.v1.Connection.service
	Service *string `json:"service,omitempty"`
	*/
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpservicenetworkingconnection;gcpservicenetworkingconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1beta1"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ServiceNetworkingConnection is the Schema for the ServiceNetworkingConnection API
// +k8s:openapi-gen=true
type ServiceNetworkingConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ServiceNetworkingConnectionSpec   `json:"spec,omitempty"`
	Status ServiceNetworkingConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ServiceNetworkingConnectionList contains a list of ServiceNetworkingConnection
type ServiceNetworkingConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkingConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceNetworkingConnection{}, &ServiceNetworkingConnectionList{})
}
