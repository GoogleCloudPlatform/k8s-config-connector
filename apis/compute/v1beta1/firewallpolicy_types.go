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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkFirewallPolicyGVK = GroupVersion.WithKind("ComputeNetworkFirewallPolicy")

// ComputeNetworkFirewallPolicySpec defines the desired state of ComputeNetworkFirewallPolicy
// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicy
type ComputeNetworkFirewallPolicySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputeNetworkFirewallPolicy name. If not given, the metadata.name will be used.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.name
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.description
	Description *string `json:"description,omitempty"`
}

// ComputeNetworkFirewallPolicyStatus defines the config connector machine state of ComputeNetworkFirewallPolicy
type ComputeNetworkFirewallPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNetworkFirewallPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeNetworkFirewallPolicyObservedState `json:"observedState,omitempty"`
}

// ComputeNetworkFirewallPolicyObservedState is the state of the ComputeNetworkFirewallPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.FirewallPolicy
type ComputeNetworkFirewallPolicyObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.rule_tuple_count
	RuleTupleCount *int32 `json:"ruleTupleCount,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource with the resource id.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkfirewallpolicy;gcpcomputenetworkfirewallpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkFirewallPolicy is the Schema for the ComputeNetworkFirewallPolicy API
// +k8s:openapi-gen=true
type ComputeNetworkFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkFirewallPolicySpec   `json:"spec,omitempty"`
	Status ComputeNetworkFirewallPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkFirewallPolicyList contains a list of ComputeNetworkFirewallPolicy
type ComputeNetworkFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkFirewallPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkFirewallPolicy{}, &ComputeNetworkFirewallPolicyList{})
}
