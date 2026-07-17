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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkFirewallPolicyGVK = GroupVersion.WithKind("ComputeNetworkFirewallPolicy")

// ComputeNetworkFirewallPolicySpec defines the desired state of ComputeNetworkFirewallPolicy
// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicy
type ComputeNetworkFirewallPolicySpec struct {
	/* An optional description of this resource. Provide this property when you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	// The project that this resource belongs to.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// The ComputeNetworkFirewallPolicy name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeNetworkFirewallPolicyStatus defines the config connector machine state of ComputeNetworkFirewallPolicy
type ComputeNetworkFirewallPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNetworkFirewallPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* Fingerprint of the resource. This field is used internally during updates of this resource. */
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* The unique identifier for the resource. This identifier is defined by the server. */
	// +optional
	NetworkFirewallPolicyId *string `json:"networkFirewallPolicyId,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples. */
	// +optional
	RuleTupleCount *int64 `json:"ruleTupleCount,omitempty"`

	/* Server-defined URL for the resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* Server-defined URL for this resource with the resource id. */
	// +optional
	SelfLinkWithId *string `json:"selfLinkWithId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkfirewallpolicy;gcpcomputenetworkfirewallpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkFirewallPolicy is the Schema for the ComputeNetworkFirewallPolicy API
// +k8s:openapi-gen=true
// +kcc:proto=google.cloud.compute.v1.FirewallPolicy
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
