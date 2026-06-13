// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
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

var ComputeOrganizationSecurityPolicyAssociationGVK = GroupVersion.WithKind("ComputeOrganizationSecurityPolicyAssociation")

// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicyAssociation
type ComputeOrganizationSecurityPolicyAssociationSpec struct {
	/* Immutable. The resource that the security policy is attached to. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyAssociation.attachment_target
	AttachmentTarget string `json:"attachmentId"`

	/* Immutable. The security policy ID of the association. */
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyAssociation.firewall_policy_id
	FirewallPolicyID string `json:"policyId"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:status:proto=google.cloud.compute.v1.FirewallPolicyAssociation
type ComputeOrganizationSecurityPolicyAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeOrganizationSecurityPolicyAssociation's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The display name of the security policy of the association. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyAssociation.display_name
	DisplayName *string `json:"displayName,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeorganizationsecuritypolicyassociation;gcpcomputeorganizationsecuritypolicyassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeOrganizationSecurityPolicyAssociation is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeOrganizationSecurityPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeOrganizationSecurityPolicyAssociationSpec   `json:"spec,omitempty"`
	Status ComputeOrganizationSecurityPolicyAssociationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeOrganizationSecurityPolicyAssociationList contains a list of ComputeOrganizationSecurityPolicyAssociation
type ComputeOrganizationSecurityPolicyAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeOrganizationSecurityPolicyAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeOrganizationSecurityPolicyAssociation{}, &ComputeOrganizationSecurityPolicyAssociationList{})
}
