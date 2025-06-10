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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OrgPolicyPolicyGVK = GroupVersion.WithKind("OrgPolicyPolicy")

// OrgPolicyPolicySpec defines the desired state of OrgPolicyPolicy
// +kcc:proto=google.cloud.orgpolicy.v2.Policy
type OrgPolicyPolicySpec struct {
	// Immutable. The Project that this resource belongs to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The Folder that this resource belongs to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	// Immutable. The Organization that this resource belongs to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	// The OrgPolicyPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Basic information about the Organization Policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.spec
	Spec *PolicySpec `json:"spec,omitempty"`

	// Dry-run policy.
	//  Audit-only policy, can be used to monitor how the policy would have
	//  impacted the existing and future resources if it's enforced.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.dry_run_spec
	DryRunSpec *PolicySpec `json:"dryRunSpec,omitempty"`

	// Optional. An opaque tag indicating the current state of the policy, used
	//  for concurrency control. This 'etag' is computed by the server based on the
	//  value of other fields, and may be sent on update and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.etag
	Etag *string `json:"etag,omitempty"`
}

// OrgPolicyPolicyStatus defines the config connector machine state of OrgPolicyPolicy
type OrgPolicyPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the OrgPolicyPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *OrgPolicyPolicyObservedState `json:"observedState,omitempty"`
}

// OrgPolicyPolicyObservedState is the state of the OrgPolicyPolicy resource as most recently observed in GCP.
// +kcc:proto=google.cloud.orgpolicy.v2.Policy
type OrgPolicyPolicyObservedState struct {
	// Basic information about the Organization Policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.spec
	Spec *PolicySpecObservedState `json:"spec,omitempty"`

	// Dry-run policy.
	//  Audit-only policy, can be used to monitor how the policy would have
	//  impacted the existing and future resources if it's enforced.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.dry_run_spec
	DryRunSpec *PolicySpecObservedState `json:"dryRunSpec,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcporgpolicypolicy;gcporgpolicypolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OrgPolicyPolicy is the Schema for the OrgPolicyPolicy API
// +k8s:openapi-gen=true
type OrgPolicyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OrgPolicyPolicySpec   `json:"spec,omitempty"`
	Status OrgPolicyPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OrgPolicyPolicyList contains a list of OrgPolicyPolicy
type OrgPolicyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrgPolicyPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OrgPolicyPolicy{}, &OrgPolicyPolicyList{})
}
