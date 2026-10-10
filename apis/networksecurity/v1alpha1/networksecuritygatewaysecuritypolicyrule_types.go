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

// API sources for NetworkSecurityGatewaySecurityPolicyRule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/networksecurity/v1/gateway_security_policy_rule.proto
// +kcc:source:service-docs=https://cloud.google.com/networking
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/networking/docs/reference/rest/v1/projects.locations.gatewaySecurityPolicies.rules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityGatewaySecurityPolicyRuleGVK = GroupVersion.WithKind("NetworkSecurityGatewaySecurityPolicyRule")

// NetworkSecurityGatewaySecurityPolicyRuleSpec defines the desired state of NetworkSecurityGatewaySecurityPolicyRule
// +kcc:spec:proto=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule
// +kcc:required-from-proto
type NetworkSecurityGatewaySecurityPolicyRuleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy} this resource belongs to.
	// +kcc:guess
	// GatewaySecurityPolicyRef *NetworkSecurityGatewaySecurityPolicyRef `json:"gatewaySecurityPolicyRef,omitempty"`

	// The NetworkSecurityGatewaySecurityPolicyRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Profile which tells what the primitive action should be.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.basic_profile
	// +required
	BasicProfile *string `json:"basicProfile,omitempty"`

	// Required. Whether the rule is enforced.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.enabled
	// +required
	Enabled *bool `json:"enabled,omitempty"`

	// Required. Priority of the rule.
	//  Lower number corresponds to higher precedence.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.priority
	// +required
	Priority *int32 `json:"priority,omitempty"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.description
	Description *string `json:"description,omitempty"`

	// Required. CEL expression for matching on session criteria.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.session_matcher
	// +required
	SessionMatcher *string `json:"sessionMatcher,omitempty"`

	// Optional. CEL expression for matching on L7/application level criteria.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.application_matcher
	ApplicationMatcher *string `json:"applicationMatcher,omitempty"`

	// Optional. Flag to enable TLS inspection of traffic matching on
	//  <session_matcher>, can only be true if the parent GatewaySecurityPolicy
	//  references a TLSInspectionConfig.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.tls_inspection_enabled
	TLSInspectionEnabled *bool `json:"tlsInspectionEnabled,omitempty"`
}

// NetworkSecurityGatewaySecurityPolicyRuleStatus defines the config connector machine state of NetworkSecurityGatewaySecurityPolicyRule
type NetworkSecurityGatewaySecurityPolicyRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityGatewaySecurityPolicyRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityGatewaySecurityPolicyRuleObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityGatewaySecurityPolicyRuleObservedState is the state of the NetworkSecurityGatewaySecurityPolicyRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule
type NetworkSecurityGatewaySecurityPolicyRuleObservedState struct {
	// Output only. Time when the rule was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the rule was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.GatewaySecurityPolicyRule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritygatewaysecuritypolicyrule;gcpnetworksecuritygatewaysecuritypolicyrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityGatewaySecurityPolicyRule is the Schema for the NetworkSecurityGatewaySecurityPolicyRule API
// +k8s:openapi-gen=true
type NetworkSecurityGatewaySecurityPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityGatewaySecurityPolicyRuleSpec   `json:"spec,omitempty"`
	Status NetworkSecurityGatewaySecurityPolicyRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityGatewaySecurityPolicyRuleList contains a list of NetworkSecurityGatewaySecurityPolicyRule
type NetworkSecurityGatewaySecurityPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGatewaySecurityPolicyRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityGatewaySecurityPolicyRule{}, &NetworkSecurityGatewaySecurityPolicyRuleList{})
}
