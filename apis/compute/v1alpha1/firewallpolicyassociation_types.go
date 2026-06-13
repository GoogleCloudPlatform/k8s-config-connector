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

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkFirewallPolicyAssociationGVK = GroupVersion.WithKind("ComputeNetworkFirewallPolicyAssociation")

// ComputeNetworkFirewallPolicyAssociationSpec defines the desired state of ComputeNetworkFirewallPolicyAssociation
// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicyAssociation
type ComputeNetworkFirewallPolicyAssociationSpec struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +required
	Location string `json:"location"`

	// The target that the firewall policy is attached to.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyAssociation.attachment_target
	AttachmentTargetRef *computev1beta1.ComputeNetworkRef `json:"attachmentTargetRef"`

	// The firewall policy ID of the association.
	// +required
	FirewallPolicyRef *refsv1beta1.ComputeFirewallPolicyRef `json:"firewallPolicyRef"`

	// The ComputeNetworkFirewallPolicyAssociation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeNetworkFirewallPolicyAssociationStatus defines the config connector machine state of ComputeNetworkFirewallPolicyAssociation
type ComputeNetworkFirewallPolicyAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNetworkFirewallPolicyAssociation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeNetworkFirewallPolicyAssociationObservedState `json:"observedState,omitempty"`
}

// ComputeNetworkFirewallPolicyAssociationObservedState is the state of the ComputeNetworkFirewallPolicyAssociation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.FirewallPolicyAssociation
type ComputeNetworkFirewallPolicyAssociationObservedState struct {
	// Deprecated, please use short name instead. The display name of the firewall policy of the association.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyAssociation.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The firewall policy ID of the association.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyAssociation.firewall_policy_id
	FirewallPolicyID *string `json:"firewallPolicyID,omitempty"`

	// The short name of the firewall policy of the association.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyAssociation.short_name
	ShortName *string `json:"shortName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkfirewallpolicyassociation;gcpcomputenetworkfirewallpolicyassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
