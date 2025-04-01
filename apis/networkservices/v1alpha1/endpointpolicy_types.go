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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesEndpointPolicyGVK = GroupVersion.WithKind("NetworkServicesEndpointPolicy")

// NetworkServicesEndpointPolicySpec defines the desired state of NetworkServicesEndpointPolicy
// +kcc:proto=google.cloud.networkservices.v1.EndpointPolicy
type NetworkServicesEndpointPolicySpec struct {
	// Optional. Set of label tags associated with the EndpointPolicy resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The type of endpoint policy. This is primarily used to validate
	//  the configuration.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.type
	Type *string `json:"type"` // NOTE: Marked required based on proto comment

	// Optional. This field specifies the URL of AuthorizationPolicy resource that
	//  applies authorization policies to the inbound traffic at the
	//  matched endpoints. Refer to Authorization. If this field is not
	//  specified, authorization is disabled(no authz checks) for this
	//  endpoint.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.authorization_policy
	AuthorizationPolicy *string `json:"authorizationPolicy,omitempty"`

	// Required. A matcher that selects endpoints to which the policies should be
	//  applied.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.endpoint_matcher
	EndpointMatcher *EndpointMatcher `json:"endpointMatcher"` // NOTE: Marked required based on proto comment

	// Optional. Port selector for the (matched) endpoints. If no port selector is
	//  provided, the matched config is applied to all ports.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.traffic_port_selector
	TrafficPortSelector *TrafficPortSelector `json:"trafficPortSelector,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is
	//  used to determine the authentication policy to be applied to terminate the
	//  inbound traffic at the identified backends. If this field is not set,
	//  authentication is disabled(open) for this endpoint.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.server_tls_policy
	ServerTLSPolicy *string `json:"serverTLSPolicy,omitempty"`

	// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy
	//  can be set to specify the authentication for traffic from the proxy to the
	//  actual endpoints. More specifically, it is applied to the outgoing traffic
	//  from the proxy to the endpoint. This is typically used for sidecar model
	//  where the proxy identifies itself as endpoint to the control plane, with
	//  the connection between sidecar and endpoint requiring authentication. If
	//  this field is not set, authentication is disabled(open). Applicable only
	//  when EndpointPolicyType is SIDECAR_PROXY.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.client_tls_policy
	ClientTLSPolicy *string `json:"clientTLSPolicy,omitempty"`

	// The NetworkServicesEndpointPolicy name. If not given, the metadata.name will be used.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.name,resource_name=true
	ResourceID *string `json:"resourceID,omitempty"`
}

// NetworkServicesEndpointPolicyStatus defines the config connector machine state of NetworkServicesEndpointPolicy
type NetworkServicesEndpointPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesEndpointPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesEndpointPolicyObservedState `json:"observedState,omitempty"`
}

// NetworkServicesEndpointPolicyObservedState is the state of the NetworkServicesEndpointPolicy resource as most recently observed in GCP.
// +kcc:proto=google.cloud.networkservices.v1.EndpointPolicy
type NetworkServicesEndpointPolicyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesendpointpolicy;gcpnetworkservicesendpointpolicys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesEndpointPolicy is the Schema for the NetworkServicesEndpointPolicy API
// +k8s:openapi-gen=true
type NetworkServicesEndpointPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesEndpointPolicySpec   `json:"spec,omitempty"`
	Status NetworkServicesEndpointPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesEndpointPolicyList contains a list of NetworkServicesEndpointPolicy
type NetworkServicesEndpointPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesEndpointPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesEndpointPolicy{}, &NetworkServicesEndpointPolicyList{})
}
