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

// +tool:krm-type-terraform
// proto.message: google.appengine.v1.FirewallRule
// crd.kind: AppEngineFirewallRule
// crd.version: v1alpha1
// terraform.src: github.com/hashicorp/terraform-provider-google-beta/google-beta/services/appengine/resource_app_engine_firewall_rule.go

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AppEngineFirewallRuleGVK = GroupVersion.WithKind("AppEngineFirewallRule")

// AppEngineFirewallRuleSpec defines the desired state of AppEngineFirewallRule
// +kcc:spec:proto=google.appengine.v1.FirewallRule
type AppEngineFirewallRuleSpec struct {
	// The project that this resource belongs to.
	Project *string `json:"project,omitempty"`

	// TODO: Should be projectRef
	// todo+required
	// ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The AppEngineFirewallRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The action to take on matched requests.
	// +kcc:proto:field=google.appengine.v1.FirewallRule.action
	// +required
	Action *string `json:"action,omitempty"`

	// An optional string description of this rule.
	//  This field has a maximum length of 100 characters.
	// +kcc:proto:field=google.appengine.v1.FirewallRule.description
	Description *string `json:"description,omitempty"`

	// NOTYET(terraform)
	// // A positive integer between [1, Int32.MaxValue-1] that defines the order of
	// //  rule evaluation. Rules with the lowest priority are evaluated first.
	// //
	// //  A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic
	// //  when no previous rule matches. Only the action of this rule can be modified
	// //  by the user.
	// // +kcc:proto:field=google.appengine.v1.FirewallRule.priority
	// Priority *int32 `json:"priority,omitempty"`

	// IP address or range, defined using CIDR notation, of requests that this
	//  rule applies to. You can use the wildcard character "*" to match all IPs
	//  equivalent to "0/0" and "::/0" together.
	//  Examples: `192.168.1.1` or `192.168.0.0/16` or `2001:db8::/32`
	//            or `2001:0db8:0000:0042:0000:8a2e:0370:7334`.
	//
	//
	//  <p>Truncation will be silently performed on addresses which are not
	//  properly truncated. For example, `1.2.3.4/24` is accepted as the same
	//  address as `1.2.3.0/24`. Similarly, for IPv6, `2001:db8::1/32` is accepted
	//  as the same address as `2001:db8::/32`.
	// +kcc:proto:field=google.appengine.v1.FirewallRule.source_range
	// +required
	SourceRange *string `json:"sourceRange,omitempty"`
}

// AppEngineFirewallRuleStatus defines the config connector machine state of AppEngineFirewallRule
type AppEngineFirewallRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppEngineFirewallRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// NOTYET(terraform)
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *AppEngineFirewallRuleObservedState `json:"observedState,omitempty"`
}

// AppEngineFirewallRuleObservedState is the state of the AppEngineFirewallRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.appengine.v1.FirewallRule
type AppEngineFirewallRuleObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpappenginefirewallrule;gcpappenginefirewallrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppEngineFirewallRule is the Schema for the AppEngineFirewallRule API
// +k8s:openapi-gen=true
type AppEngineFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppEngineFirewallRuleSpec   `json:"spec,omitempty"`
	Status AppEngineFirewallRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppEngineFirewallRuleList contains a list of AppEngineFirewallRule
type AppEngineFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineFirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppEngineFirewallRule{}, &AppEngineFirewallRuleList{})
}
