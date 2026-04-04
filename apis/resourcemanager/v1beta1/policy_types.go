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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ResourceManagerPolicyGVK = GroupVersion.WithKind("ResourceManagerPolicy")

type PolicyAllow struct {
	/* The policy allows or denies all values. */
	// +optional
	All *bool `json:"all,omitempty"`

	/* The policy can define specific values that are allowed or denied. */
	// +optional
	Values []string `json:"values,omitempty"`
}

type PolicyBooleanPolicy struct {
	/* If true, then the Policy is enforced. If false, then any configuration is acceptable. */
	Enforced bool `json:"enforced"`
}

type PolicyDeny struct {
	/* The policy allows or denies all values. */
	// +optional
	All *bool `json:"all,omitempty"`

	/* The policy can define specific values that are allowed or denied. */
	// +optional
	Values []string `json:"values,omitempty"`
}

type PolicyListPolicy struct {
	/* One or the other must be set. */
	// +optional
	Allow *PolicyAllow `json:"allow,omitempty"`

	/* One or the other must be set. */
	// +optional
	Deny *PolicyDeny `json:"deny,omitempty"`

	/* If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy. */
	// +optional
	InheritFromParent *bool `json:"inheritFromParent,omitempty"`

	/* The Google Cloud Console will try to default to a configuration that matches the value specified in this field. */
	// +optional
	SuggestedValue *string `json:"suggestedValue,omitempty"`
}

type PolicyRestorePolicy struct {
	/* May only be set to true. If set, then the default Policy is restored. */
	Default bool `json:"default"`
}

// +kcc:proto=google.cloud.orgpolicy.v2.Policy
type ResourceManagerPolicySpec struct {
	/* A boolean policy is a constraint that is either enforced or not. */
	// +optional
	BooleanPolicy *PolicyBooleanPolicy `json:"booleanPolicy,omitempty"`

	/* Immutable. The name of the Constraint the Policy is configuring, for example, serviceuser.services. */
	Constraint string `json:"constraint"`

	/* The folder on which to configure the constraint. Only one of
	projectRef, folderRef, or organizationRef may be specified. */
	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	/* A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. . */
	// +optional
	ListPolicy *PolicyListPolicy `json:"listPolicy,omitempty"`

	/* The organization on which to configure the constraint. Only one of
	projectRef, folderRef, or organizationRef may be specified. */
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	/* The project on which to configure the constraint. Only one of
	projectRef, folderRef, or organizationRef may be specified. */
	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	/* A restore policy is a constraint to restore the default policy. */
	// +optional
	RestorePolicy *PolicyRestorePolicy `json:"restorePolicy,omitempty"`

	/* Version of the Policy. Default version is 0. */
	// +optional
	Version *int64 `json:"version,omitempty"`

	/* In policies for boolean constraints, the following requirements apply:
	- There must be one and only one policy rule where condition is unset.
	- Boolean policy rules with conditions must set `enforced` to the
	  opposite of the policy rule without a condition.
	- During policy evaluation, policy rules with conditions that are
	  true for a target resource take precedence. */
	// +optional
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.spec.rules
	Rules []PolicySpec_PolicyRule `json:"rules,omitempty"`

	/* Determines the inheritance behavior for this policy.
	If `inherit_from_parent` is true, policy rules set higher up in the
	hierarchy (up to the closest root) are inherited and present in the
	effective policy. If it is false, then no rules are inherited, and this
	policy becomes the new root for evaluation.
	This field can be set only for policies which configure list constraints. */
	// +optional
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.spec.inherit_from_parent
	InheritFromParent *bool `json:"inheritFromParent,omitempty"`

	/* Ignores policies set above this resource and restores the
	`constraint_default` enforcement behavior of the specific constraint at
	this resource.
	This field can be set in policies for either list or boolean
	constraints. If set, `rules` must be empty and `inherit_from_parent`
	must be set to false. */
	// +optional
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.spec.reset
	Reset *bool `json:"reset,omitempty"`

	/* Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced. */
	// +optional
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.dry_run_spec
	DryRunSpec *PolicySpec `json:"dryRunSpec,omitempty"`
}

type ResourceManagerPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   ResourceManagerPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The etag of the organization policy. etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *ResourceManagerPolicyObservedState `json:"observedState,omitempty"`

	/* ExternalRef is a unique specifier for the ResourceManagerPolicy resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v2.Policy
type ResourceManagerPolicyObservedState struct {
	/* Basic information about the Organization Policy. */
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.spec
	Spec *PolicySpecObservedState `json:"spec,omitempty"`

	/* Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced. */
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.dry_run_spec
	DryRunSpec *PolicySpecObservedState `json:"dryRunSpec,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpresourcemanagerpolicy;gcpresourcemanagerpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ResourceManagerPolicy is the Schema for the resourcemanager API
// +k8s:openapi-gen=true
type ResourceManagerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceManagerPolicySpec   `json:"spec,omitempty"`
	Status ResourceManagerPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ResourceManagerPolicyList contains a list of ResourceManagerPolicy
type ResourceManagerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceManagerPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceManagerPolicy{}, &ResourceManagerPolicyList{})
}

// +kcc:proto=google.cloud.orgpolicy.v2.PolicySpec
type PolicySpec struct {
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

// +kcc:proto=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.StringValues
type PolicySpec_PolicyRule_StringValues struct {
	// List of values allowed at this resource.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.StringValues.allowed_values
	AllowedValues []string `json:"allowedValues,omitempty"`

	// List of values denied at this resource.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.StringValues.denied_values
	DeniedValues []string `json:"deniedValues,omitempty"`
}

// +kcc:proto=google.type.Expr
type Expr struct {
	// Textual representation of an expression in Common Expression Language
	//  syntax.
	// +kcc:proto:field=google.type.Expr.expression
	Expression *string `json:"expression,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing
	//  its purpose. This can be used e.g. in UIs which allow to enter the
	//  expression.
	// +kcc:proto:field=google.type.Expr.title
	Title *string `json:"title,omitempty"`

	// Optional. Description of the expression. This is a longer text which
	//  describes the expression, e.g. when hovered over it in a UI.
	// +kcc:proto:field=google.type.Expr.description
	Description *string `json:"description,omitempty"`

	// Optional. String indicating the location of the expression for error
	//  reporting, e.g. a file name and a position in the file.
	// +kcc:proto:field=google.type.Expr.location
	Location *string `json:"location,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.orgpolicy.v2.PolicySpec
type PolicySpecObservedState struct {
	// Output only. The time stamp this was previously updated. This
	//  represents the last time a call to `CreatePolicy` or `UpdatePolicy` was
	//  made for that policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
