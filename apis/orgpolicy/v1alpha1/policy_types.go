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

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
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

	// NOTYET: not supported in Config Connector reconciliation
	// Optional. An opaque tag indicating the current state of the policy, used
	//  for concurrency control. This 'etag' is computed by the server based on the
	//  value of other fields, and may be sent on update and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.etag
	// Etag *string `json:"etag,omitempty"`
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

// +kcc:proto=google.cloud.orgpolicy.v2.PolicySpec
type PolicySpec struct {
	// NOTYET: not supported in Config Connector reconciliation
	// An opaque tag indicating the current version of the policySpec, used for
	//  concurrency control.
	//
	//  This field is ignored if used in a `CreatePolicy` request.
	//
	//  When the policy is returned from either a `GetPolicy` or a
	//  `ListPolicies` request, this `etag` indicates the version of the
	//  current policySpec to use when executing a read-modify-write loop.
	//
	//  When the policy is returned from a `GetEffectivePolicy` request, the
	//  `etag` will be unset.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.etag
	// Etag *string `json:"etag,omitempty"`

	// In policies for boolean constraints, the following requirements apply:
	//
	//    - There must be one and only one policy rule where condition is unset.
	//    - Boolean policy rules with conditions must set `enforced` to the
	//      opposite of the policy rule without a condition.
	//    - During policy evaluation, policy rules with conditions that are
	//      true for a target resource take precedence.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.rules
	Rules []PolicySpec_PolicyRule `json:"rules,omitempty"`

	// Determines the inheritance behavior for this policy.
	//
	//  If `inherit_from_parent` is true, policy rules set higher up in the
	//  hierarchy (up to the closest root) are inherited and present in the
	//  effective policy. If it is false, then no rules are inherited, and this
	//  policy becomes the new root for evaluation.
	//  This field can be set only for policies which configure list constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.inherit_from_parent
	InheritFromParent *bool `json:"inheritFromParent,omitempty"`

	// Ignores policies set above this resource and restores the
	//  `constraint_default` enforcement behavior of the specific constraint at
	//  this resource.
	//  This field can be set in policies for either list or boolean
	//  constraints. If set, `rules` must be empty and `inherit_from_parent`
	//  must be set to false.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.reset
	Reset *bool `json:"reset,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.orgpolicy.v2.PolicySpec
type PolicySpecObservedState struct {
	// Output only. The time stamp this was previously updated. This
	//  represents the last time a call to `CreatePolicy` or `UpdatePolicy` was
	//  made for that policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule
type PolicySpec_PolicyRule struct {
	// List of values to be used for this policy rule. This field can be set
	//  only in policies for list constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicyRule.values
	Values *PolicySpec_PolicyRule_StringValues `json:"values,omitempty"`

	// Setting this to true means that all values are allowed. This field can
	//  be set only in policies for list constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicyRule.allow_all
	AllowAll *bool `json:"allowAll,omitempty"`

	// Setting this to true means that all values are denied. This field can
	//  be set only in policies for list constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicyRule.deny_all
	DenyAll *bool `json:"denyAll,omitempty"`

	// If `true`, then the policy is enforced. If `false`, then any
	//  configuration is acceptable.
	//  This field can be set only in policies for boolean constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicyRule.enforce
	Enforce *bool `json:"enforce,omitempty"`

	// A condition which determines whether this rule is used
	//  in the evaluation of the policy. When set, the `expression` field in
	//  the `Expr' must include from 1 to 10 subexpressions, joined by the "||"
	//  or "&&" operators. Each subexpression must be of the form
	//  "resource.matchTag('<ORG_ID>/tag_key_short_name,
	//  'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id',
	//  'tagValues/value_id')". where key_name and value_name are the resource
	//  names for Label Keys and Values. These names are available from the Tag
	//  Manager Service. An example expression is:
	//  "resource.matchTag('123456789/environment,
	//  'prod')". or "resource.matchTagId('tagKeys/123',
	//  'tagValues/456')".
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicyRule.condition
	Condition *Expr `json:"condition,omitempty"`

	// Optional. Required for managed constraints if parameters are defined.
	//  Passes parameter values when policy enforcement is enabled. Ensure that
	//  parameter value types match those defined in the constraint definition.
	//  For example:
	//  {
	//    "allowedLocations" : ["us-east1", "us-west1"],
	//    "allowAll" : true
	//  }
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicyRule.parameters
	Parameters *apiextensionsv1.JSON `json:"parameters,omitempty"`
}
