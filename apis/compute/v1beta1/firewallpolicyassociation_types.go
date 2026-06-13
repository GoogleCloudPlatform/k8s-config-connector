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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeFirewallPolicyAssociationGVK = GroupVersion.WithKind("ComputeFirewallPolicyAssociation")

// ComputeFirewallPolicyAssociationSpec defines the desired state of ComputeFirewallPolicyAssociation
// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicyAssociation
type ComputeFirewallPolicyAssociationSpec struct {
	/* Immutable. The target that the firewall policy is attached to. */
	AttachmentTargetRef k8sv1alpha1.ResourceRef `json:"attachmentTargetRef"`

	/* Immutable. The firewall policy of the association. */
	FirewallPolicyRef refsv1beta1.ComputeFirewallPolicyRef `json:"firewallPolicyRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeFirewallPolicyAssociationStatus defines the config connector machine state of ComputeFirewallPolicyAssociation
type ComputeFirewallPolicyAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeFirewallPolicyAssociation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeFirewallPolicyAssociationObservedState `json:"observedState,omitempty"`
}

// ComputeFirewallPolicyAssociationObservedState is the state of the ComputeFirewallPolicyAssociation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.FirewallPolicyAssociation
type ComputeFirewallPolicyAssociationObservedState struct {
	// Deprecated, please use short name instead. The display name of the firewall policy of the association.
	DisplayName *string `json:"displayName,omitempty"`

	// The firewall policy ID of the association.
	FirewallPolicyID *string `json:"firewallPolicyID,omitempty"`

	// The short name of the firewall policy of the association.
	ShortName *string `json:"shortName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputefirewallpolicyassociation;gcpcomputefirewallpolicyassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeFirewallPolicyAssociation is the Schema for the ComputeFirewallPolicyAssociation API
// +k8s:openapi-gen=true
type ComputeFirewallPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeFirewallPolicyAssociationSpec   `json:"spec,omitempty"`
	Status ComputeFirewallPolicyAssociationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeFirewallPolicyAssociationList contains a list of ComputeFirewallPolicyAssociation
type ComputeFirewallPolicyAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeFirewallPolicyAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeFirewallPolicyAssociation{}, &ComputeFirewallPolicyAssociationList{})
}
