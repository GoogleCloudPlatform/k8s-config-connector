package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanageraccesspolicy;gcpaccesscontextmanageraccesspolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true"

// AccessContextManagerAccessPolicy is the Schema for the accesscontextmanager AccessPolicy resource.
type AccessContextManagerAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerAccessPolicySpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessPolicyStatus `json:"status,omitempty"`
}

// AccessContextManagerAccessPolicySpec defines the desired state of AccessContextManagerAccessPolicy.
type AccessContextManagerAccessPolicySpec struct {
	Parent Parent `json:",inline"`
	// Required. Human readable title. Does not affect behavior.
	// +required
	Title string `json:"title"`
	// The scopes of a policy define which resources an ACM policy can restrict,
	// and where ACM resources can be referenced.
	// For example, a policy with scopes=["folders/123"] has the following
	// behavior:
	// - vpcsc perimeters can only restrict projects within folders/123
	// - access levels can only be referenced by resources within folders/123.
	// If empty, there are no limitations on which resources can be restricted by
	// an ACM policy, and there are no limitations on where ACM resources can be
	// referenced.
	// Only one policy can include a given scope (attempting to create a second
	// policy which includes "folders/123" will result in an error).
	// Currently, scopes cannot be modified after a policy is created.
	// Currently, policies can only have a single scope.
	// Format: list of `folders/{folder_number}` or `projects/{project_number}`
	// +optional
	Scopes []string `json:"scopes,omitempty"`
}

// Parent defines the parent of the AccessContextManagerAccessPolicy.
type Parent struct {
	// Required. The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Currently immutable once created. Format: organizations/{organization_id}
	// +required
	OrganizationRef *refv1beta1.OrganizationRef `json:"organizationRef"`
}

// AccessContextManagerAccessPolicyStatus defines the observed state of AccessContextManagerAccessPolicy.
type AccessContextManagerAccessPolicyStatus struct {
	// Conditions represent the latest available observations of the
	// AccessContextManagerAccessPolicy's current state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Output only. Time the AccessPolicy was created in UTC.
	CreateTime string `json:"createTime,omitempty"`
	// Output only. Time the AccessPolicy was updated in UTC.
	UpdateTime string `json:"updateTime,omitempty"`
	// Output only. An opaque identifier for the current version of the
	// AccessPolicy. This will always be a strongly validated etag, meaning that
	// two Access Polices will be identical if and only if their etags are
	// identical. Clients should not expect this to be in any specific format.
	Etag string `json:"etag,omitempty"`
	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the desired state of the resource.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerAccessPolicyList contains a list of AccessContextManagerAccessPolicy.
type AccessContextManagerAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerAccessPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerAccessPolicy{}, &AccessContextManagerAccessPolicyList{})
}
