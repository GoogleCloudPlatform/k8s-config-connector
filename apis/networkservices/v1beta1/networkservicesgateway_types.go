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
	networksecurityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesGatewayGVK = GroupVersion.WithKind("NetworkServicesGateway")

// NetworkServicesGatewaySpec defines the desired state of NetworkServicesGateway
// +kcc:spec:proto=google.cloud.networkservices.v1.Gateway
type NetworkServicesGatewaySpec struct {
	// Optional. Zero or one IPv4 or IPv6 address on which the Gateway will
	//  receive the traffic. When no address is provided, an IP from the subnetwork
	//  is allocated
	//
	//  This field only applies to gateways of type 'SECURE_WEB_GATEWAY'.
	//  Gateways of type 'OPEN_MESH' listen on 0.0.0.0 for IPv4 and :: for IPv6.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.addresses
	Addresses []string `json:"addresses,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024 characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.description
	Description *string `json:"description,omitempty"`

	// Immutable. The location for the resource.
	Location string `json:"location"`

	// Required. One or more ports that the Gateway must receive traffic on. The proxy binds to the ports specified. Gateway listen on 0.0.0.0 on the ports specified below.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.ports
	Ports []int64 `json:"ports"`

	// Immutable. The Project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Required. Immutable. Scope determines how configuration across multiple Gateway instances are merged. The configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer. Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.scope
	Scope string `json:"scope"`

	// Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.server_tls_policy
	ServerTlsPolicyRef *networksecurityv1beta1.NetworkSecurityServerTLSPolicyRef `json:"serverTlsPolicyRef,omitempty"`

	// Immutable. Immutable. The type of the customer managed gateway. Possible values: TYPE_UNSPECIFIED, OPEN_MESH, SECURE_WEB_GATEWAY
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.type
	Type *string `json:"type,omitempty"`
}

// NetworkServicesGatewayStatus defines the config connector machine state of NetworkServicesGateway
type NetworkServicesGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. Server-defined URL of this resource
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.create_time
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.update_time
	// +kubebuilder:validation:Format=date-time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesgateway;gcpnetworkservicesgateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesGateway is the Schema for the NetworkServicesGateway API
// +k8s:openapi-gen=true
type NetworkServicesGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesGatewaySpec   `json:"spec,omitempty"`
	Status NetworkServicesGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesGatewayList contains a list of NetworkServicesGateway
type NetworkServicesGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesGateway{}, &NetworkServicesGatewayList{})
}
