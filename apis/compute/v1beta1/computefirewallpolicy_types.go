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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeFirewallPolicyGVK = GroupVersion.WithKind("ComputeFirewallPolicy")

// ComputeFirewallPolicySpec defines the desired state of ComputeFirewallPolicy
// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicy
type ComputeFirewallPolicySpec struct {
	/* An optional description of this resource. Provide this property when you create the resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.description
	Description *string `json:"description,omitempty"`

	/* Immutable. The Folder that this resource belongs to. Only one of [folderRef, organizationRef] may be specified. */
	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	/* Immutable. The Organization that this resource belongs to. Only one of [folderRef, organizationRef] may be specified. */
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.short_name
	ShortName string `json:"shortName"`
}

// ComputeFirewallPolicyStatus defines the config connector machine state of ComputeFirewallPolicy
type ComputeFirewallPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Fingerprint of the resource. This field is used internally during updates of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.id
	ID *string `json:"id,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.rule_tuple_count
	RuleTupleCount *int64 `json:"ruleTupleCount,omitempty"`

	// Server-defined URL for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Server-defined URL for this resource with the resource id.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicy.self_link_with_id
	SelfLinkWithId *string `json:"selfLinkWithId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputefirewallpolicy;gcpcomputefirewallpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeFirewallPolicy is the Schema for the ComputeFirewallPolicy API
// +k8s:openapi-gen=true
// +kcc:proto=google.cloud.compute.v1.FirewallPolicy
type ComputeFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeFirewallPolicySpec   `json:"spec,omitempty"`
	Status ComputeFirewallPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeFirewallPolicyList contains a list of ComputeFirewallPolicy
type ComputeFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeFirewallPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeFirewallPolicy{}, &ComputeFirewallPolicyList{})
}
