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

var OrgPolicyCustomConstraintGVK = GroupVersion.WithKind("OrgPolicyCustomConstraint")

// OrgPolicyCustomConstraintSpec defines the desired state of OrgPolicyCustomConstraint
// +kcc:proto=google.cloud.orgpolicy.v2.CustomConstraint
type OrgPolicyCustomConstraintSpec struct {
	// The Organization that this resource belongs to.
	// +required
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The OrgPolicyCustomConstraint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Name of the constraint. This is unique within the organization.
	//  Format of the name should be
	//
	//  * `organizations/{organization_id}/customConstraints/{custom_constraint_id}`
	//
	//  Example: `organizations/123/customConstraints/custom.createOnlyE2TypeVms`
	//
	//  The max length is 70 characters and the minimum length is 1. Note that the
	//  prefix `organizations/{organization_id}/customConstraints/` is not counted.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.name
	Name *string `json:"name,omitempty"`

	// Immutable. The resource instance type on which this policy applies. Format
	//  will be of the form : `<canonical service name>/<type>` Example:
	//
	//   * `compute.googleapis.com/Instance`.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.resource_types
	ResourceTypes []string `json:"resourceTypes,omitempty"`

	// All the operations being applied for this constraint.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.method_types
	MethodTypes []string `json:"methodTypes,omitempty"`

	// Org policy condition/expression. For example:
	//  `resource.instanceName.matches("[production|test]_.*_(\d)+")` or,
	//  `resource.management.auto_upgrade == true`
	//
	//  The max length of the condition is 1000 characters.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.condition
	Condition *string `json:"condition,omitempty"`

	// Allow or deny type.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.action_type
	ActionType *string `json:"actionType,omitempty"`

	// One line display name for the UI.
	//  The max length of the display_name is 200 characters.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Detailed information about this custom policy constraint.
	//  The max length of the description is 2000 characters.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.description
	Description *string `json:"description,omitempty"`
}

// OrgPolicyCustomConstraintStatus defines the config connector machine state of OrgPolicyCustomConstraint
type OrgPolicyCustomConstraintStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the OrgPolicyCustomConstraint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *OrgPolicyCustomConstraintObservedState `json:"observedState,omitempty"`
}

// OrgPolicyCustomConstraintObservedState is the state of the OrgPolicyCustomConstraint resource as most recently observed in GCP.
// +kcc:proto=google.cloud.orgpolicy.v2.CustomConstraint
type OrgPolicyCustomConstraintObservedState struct {
	// Output only. The last time this custom constraint was updated. This
	//  represents the last time that the `CreateCustomConstraint` or
	//  `UpdateCustomConstraint` RPC was called
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcporgpolicycustomconstraint;gcporgpolicycustomconstraints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OrgPolicyCustomConstraint is the Schema for the OrgPolicyCustomConstraint API
// +k8s:openapi-gen=true
type OrgPolicyCustomConstraint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OrgPolicyCustomConstraintSpec   `json:"spec,omitempty"`
	Status OrgPolicyCustomConstraintStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OrgPolicyCustomConstraintList contains a list of OrgPolicyCustomConstraint
type OrgPolicyCustomConstraintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrgPolicyCustomConstraint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OrgPolicyCustomConstraint{}, &OrgPolicyCustomConstraintList{})
}
