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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkFirewallPolicyAssociationGVK = GroupVersion.WithKind("ComputeNetworkFirewallPolicyAssociation")

type ComputeNetworkFirewallPolicyRef struct {
	// A reference to an externally managed ComputeNetworkFirewallPolicy resource.
	// Should be in the format `projects/{{projectID}}/global/firewallPolicies/{{firewallPolicyID}}`.
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeNetworkFirewallPolicy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeNetworkFirewallPolicy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

// ComputeNetworkFirewallPolicyAssociationSpec defines the desired state of ComputeNetworkFirewallPolicyAssociation
// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicyAssociation
type ComputeNetworkFirewallPolicyAssociationSpec struct {
	// The target that the firewall policy is attached to.
	AttachmentTargetRef *ComputeNetworkRef `json:"attachmentTargetRef"`

	// The firewall policy ID of the association.
	FirewallPolicyRef *ComputeNetworkFirewallPolicyRef `json:"firewallPolicyRef"`

	// The project that this resource belongs to.
	ProjectRef *apirefs.ProjectRef `json:"projectRef"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeNetworkFirewallPolicyAssociationStatus defines the config connector machine state of ComputeNetworkFirewallPolicyAssociation
type ComputeNetworkFirewallPolicyAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The short name of the firewall policy of the association.
	// +optional
	ShortName *string `json:"shortName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkfirewallpolicyassociation;gcpcomputenetworkfirewallpolicyassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkFirewallPolicyAssociation is the Schema for the ComputeNetworkFirewallPolicyAssociation API
// +k8s:openapi-gen=true
type ComputeNetworkFirewallPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkFirewallPolicyAssociationSpec   `json:"spec,omitempty"`
	Status ComputeNetworkFirewallPolicyAssociationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkFirewallPolicyAssociationList contains a list of ComputeNetworkFirewallPolicyAssociation
type ComputeNetworkFirewallPolicyAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkFirewallPolicyAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkFirewallPolicyAssociation{}, &ComputeNetworkFirewallPolicyAssociationList{})
}
