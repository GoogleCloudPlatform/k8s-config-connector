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

var ReCAPTCHAEnterpriseFirewallPolicyGVK = GroupVersion.WithKind("ReCAPTCHAEnterpriseFirewallPolicy")

// ReCAPTCHAEnterpriseFirewallPolicySpec defines the desired state of ReCAPTCHAEnterpriseFirewallPolicy
// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallPolicy
type ReCAPTCHAEnterpriseFirewallPolicySpec struct {
	// The ReCAPTCHAEnterpriseFirewallPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A description of what this policy aims to achieve, for
	//  convenience purposes. The description can at most include 256 UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. The path for which this policy applies, specified as a glob
	//  pattern. For more information on glob, see the [manual
	//  page](https://man7.org/linux/man-pages/man7/glob.7.html).
	//  A path has a max length of 200 characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.path
	Path *string `json:"path,omitempty"`

	// Optional. A CEL (Common Expression Language) conditional expression that
	//  specifies if this policy applies to an incoming user request. If this
	//  condition evaluates to true and the requested path matched the path
	//  pattern, the associated actions should be executed by the caller. The
	//  condition string is checked for CEL syntax correctness on creation. For
	//  more information, see the [CEL spec](https://github.com/google/cel-spec)
	//  and its [language
	//  definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md).
	//  A condition has a max length of 500 characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.condition
	Condition *string `json:"condition,omitempty"`

	// Optional. The actions that the caller should take regarding user access.
	//  There should be at most one terminal action. A terminal action is any
	//  action that forces a response, such as `AllowAction`,
	//  `BlockAction` or `SubstituteAction`.
	//  Zero or more non-terminal actions such as `SetHeader` might be
	//  specified. A single policy can contain up to 16 actions.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.actions
	Actions []FirewallAction `json:"actions,omitempty"`
}
	// The ReCAPTCHAEnterpriseFirewallPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. A description of what this policy aims to achieve, for
	//  convenience purposes. The description can at most include 256 UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. The path for which this policy applies, specified as a glob
	//  pattern. For more information on glob, see the [manual
	//  page](https://man7.org/linux/man-pages/man7/glob.7.html).
	//  A path has a max length of 200 characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.path
	Path *string `json:"path,omitempty"`

	// Optional. A CEL (Common Expression Language) conditional expression that
	//  specifies if this policy applies to an incoming user request. If this
	//  condition evaluates to true and the requested path matched the path
	//  pattern, the associated actions should be executed by the caller. The
	//  condition string is checked for CEL syntax correctness on creation. For
	//  more information, see the [CEL spec](https://github.com/google/cel-spec)
	//  and its [language
	//  definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md).
	//  A condition has a max length of 500 characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.condition
	Condition *string `json:"condition,omitempty"`

	// Optional. The actions that the caller should take regarding user access.
	//  There should be at most one terminal action. A terminal action is any
	//  action that forces a response, such as `AllowAction`,
	//  `BlockAction` or `SubstituteAction`.
	//  Zero or more non-terminal actions such as `SetHeader` might be
	//  specified. A single policy can contain up to 16 actions.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.actions
	Actions []FirewallAction `json:"actions,omitempty"`
}

// ReCAPTCHAEnterpriseFirewallPolicyStatus defines the config connector machine state of ReCAPTCHAEnterpriseFirewallPolicy
type ReCAPTCHAEnterpriseFirewallPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ReCAPTCHAEnterpriseFirewallPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ReCAPTCHAEnterpriseFirewallPolicyObservedState `json:"observedState,omitempty"`
}

// ReCAPTCHAEnterpriseFirewallPolicyObservedState is the state of the ReCAPTCHAEnterpriseFirewallPolicy resource as most recently observed in GCP.
type ReCAPTCHAEnterpriseFirewallPolicyObservedState struct {

}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcprecaptchaenterprisefirewallpolicy;gcprecaptchaenterprisefirewallpolicys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ReCAPTCHAEnterpriseFirewallPolicy is the Schema for the ReCAPTCHAEnterpriseFirewallPolicy API
// +k8s:openapi-gen=true
type ReCAPTCHAEnterpriseFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ReCAPTCHAEnterpriseFirewallPolicySpec   `json:"spec,omitempty"`
	Status ReCAPTCHAEnterpriseFirewallPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ReCAPTCHAEnterpriseFirewallPolicyList contains a list of ReCAPTCHAEnterpriseFirewallPolicy
type ReCAPTCHAEnterpriseFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReCAPTCHAEnterpriseFirewallPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReCAPTCHAEnterpriseFirewallPolicy{}, &ReCAPTCHAEnterpriseFirewallPolicyList{})
}
