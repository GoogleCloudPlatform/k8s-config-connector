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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityGatewaySecurityPolicyRuleGVK = GroupVersion.WithKind("NetworkSecurityGatewaySecurityPolicyRule")

// NetworkSecurityGatewaySecurityPolicyRuleSpec defines the desired state of NetworkSecurityGatewaySecurityPolicyRule
type NetworkSecurityGatewaySecurityPolicyRuleSpec struct {
	// The GatewaySecurityPolicyRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The gateway security policy that this rule belongs to.
	GatewaySecurityPolicyRef refs.NetworkSecurityGatewaySecurityPolicyRef `json:"gatewaySecurityPolicyRef"`

	// Required. The location of the gateway security policy rule.
	Location string `json:"location"`

	// Required. Profile which tells what the primitive action should be. Possible values are: ALLOW, DENY.
	BasicProfile string `json:"basicProfile"`

	// Required. Whether the rule is enforced.
	Enabled bool `json:"enabled"`

	// Required. Priority of the rule. Lower number corresponds to higher precedence.
	Priority int64 `json:"priority"`

	// Required. CEL expression for matching on session criteria.
	SessionMatcher string `json:"sessionMatcher"`

	// Optional. CEL expression for matching on L7/application level criteria.
	ApplicationMatcher *string `json:"applicationMatcher,omitempty"`

	// Optional. Free-text description of the resource.
	Description *string `json:"description,omitempty"`

	// Optional. Flag to enable TLS inspection of traffic matching on, can only be true if the parent GatewaySecurityPolicy references a TLSInspectionConfig.
	TlsInspectionEnabled *bool `json:"tlsInspectionEnabled,omitempty"`
}

// NetworkSecurityGatewaySecurityPolicyRuleStatus defines the config connector machine state of NetworkSecurityGatewaySecurityPolicyRule
type NetworkSecurityGatewaySecurityPolicyRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. Time when the resource was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the resource was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// A unique specifier for the NetworkSecurityGatewaySecurityPolicyRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritygatewaysecuritypolicyrule;gcpnetworksecuritygatewaysecuritypolicyrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityGatewaySecurityPolicyRule is the Schema for the networksecurity API
// +k8s:openapi-gen=true
type NetworkSecurityGatewaySecurityPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

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
