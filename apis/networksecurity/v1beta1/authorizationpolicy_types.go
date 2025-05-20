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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityAuthorizationPolicyGVK = GroupVersion.WithKind("NetworkSecurityAuthorizationPolicy")

type Parent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// NetworkSecurityAuthorizationPolicySpec defines the desired state of NetworkSecurityAuthorizationPolicy
// +kcc:spec:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy
type NetworkSecurityAuthorizationPolicySpec struct {
	// The NetworkSecurityAuthorizationPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Required. Name of the AuthorizationPolicy resource. It matches pattern
	//  `projects/{project}/locations/{location}/authorizationPolicies/<authorization_policy>`.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.name
	// NOTYET: this field serves the same purpose as identity
	// Name *string `json:"name,omitempty"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of label tags associated with the AuthorizationPolicy
	//  resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.labels
	// NOTEYET: this field is not in the DCL-based beta resource
	// Labels map[string]string `json:"labels,omitempty"`

	// Required. The action to take when a rule match is found. Possible values
	//  are "ALLOW" or "DENY".
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.action
	// +required
	Action *string `json:"action,omitempty"`

	// Optional. List of rules to match. Note that at least one of the rules must
	//  match in order for the action specified in the 'action' field to be taken.
	//  A rule is a match if there is a matching source and destination. If left
	//  blank, the action specified in the `action` field will be applied on every
	//  request.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.rules
	Rules []AuthorizationPolicy_Rule `json:"rules,omitempty"`
}

// NetworkSecurityAuthorizationPolicyStatus defines the config connector machine state of NetworkSecurityAuthorizationPolicy
// +kcc:status:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy
type NetworkSecurityAuthorizationPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityAuthorizationPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// NOTYET: output only fields are moved directly under "status" due to backward compatibility
	// ObservedState *NetworkSecurityAuthorizationPolicyObservedState `json:"observedState,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// NetworkSecurityAuthorizationPolicyObservedState is the state of the NetworkSecurityAuthorizationPolicy resource as most recently observed in GCP.
// +kcc:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy
//type NetworkSecurityAuthorizationPolicyObservedState struct {
//}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityauthorizationpolicy;gcpnetworksecurityauthorizationpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityAuthorizationPolicy is the Schema for the NetworkSecurityAuthorizationPolicy API
// +k8s:openapi-gen=true
type NetworkSecurityAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityAuthorizationPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityAuthorizationPolicyList contains a list of NetworkSecurityAuthorizationPolicy
type NetworkSecurityAuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityAuthorizationPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityAuthorizationPolicy{}, &NetworkSecurityAuthorizationPolicyList{})
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination.HttpHeaderMatch
type AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch struct {
	// Required. The value of the header must match the regular expression
	//  specified in regexMatch. For regular expression grammar,
	//  please see: en.cppreference.com/w/cpp/regex/ecmascript
	//  For matching against a port specified in the HTTP
	//  request, use a headerMatch with headerName set to Host
	//  and a regular expression that satisfies the RFC2616 Host
	//  header's port specifier.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination.HttpHeaderMatch.regex_match
	// +required
	RegexMatch *string `json:"regexMatch,omitempty"`

	// Required. The name of the HTTP header to match. For matching
	//  against the HTTP request's authority, use a headerMatch
	//  with the header name ":authority". For matching a
	//  request's method, use the headerName ":method".
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination.HttpHeaderMatch.header_name
	// +required
	HeaderName *string `json:"headerName,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination
type AuthorizationPolicy_Rule_Destination struct {
	// Required. List of host names to match. Matched against the ":authority"
	//  header in http requests. At least one host should match. Each host can
	//  be an exact match, or a prefix match (example "mydomain.*") or a suffix
	//  match (example "*.myorg.com") or a presence (any) match "*".
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination.hosts
	// +required
	Hosts []string `json:"hosts,omitempty"`

	// Required. List of destination ports to match. At least one port should
	//  match.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination.ports
	// +required
	Ports []uint32 `json:"ports,omitempty"`

	// Optional. A list of HTTP methods to match. At least one method should
	//  match. Should not be set for gRPC services.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination.methods
	Methods []string `json:"methods,omitempty"`

	// Optional. Match against key:value pair in http header. Provides a
	//  flexible match based on HTTP headers, for potentially advanced use
	//  cases. At least one header should match. Avoid using header matches to
	//  make authorization decisions unless there is a strong guarantee that
	//  requests arrive through a trusted client or proxy.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Destination.http_header_match
	HTTPHeaderMatch *AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch `json:"httpHeaderMatch,omitempty"`
}
