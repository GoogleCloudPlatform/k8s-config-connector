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
// +kcc:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy
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
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The action to take when a rule match is found. Possible values
	//  are "ALLOW" or "DENY".
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.action
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
type NetworkSecurityAuthorizationPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityAuthorizationPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityAuthorizationPolicyObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityAuthorizationPolicyObservedState is the state of the NetworkSecurityAuthorizationPolicy resource as most recently observed in GCP.
// +kcc:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy
type NetworkSecurityAuthorizationPolicyObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityauthorizationpolicy;gcpnetworksecurityauthorizationpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
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
