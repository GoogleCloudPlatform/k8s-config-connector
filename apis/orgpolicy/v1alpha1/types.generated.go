// Copyright 2025 Google LLC
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

// +kcc:proto=google.cloud.orgpolicy.v2.Policy
type Policy struct {
	// Immutable. The resource name of the policy. Must be one of the following
	//  forms, where `constraint_name` is the name of the constraint which this
	//  policy configures:
	//
	//  * `projects/{project_number}/policies/{constraint_name}`
	//  * `folders/{folder_id}/policies/{constraint_name}`
	//  * `organizations/{organization_id}/policies/{constraint_name}`
	//
	//  For example, `projects/123/policies/compute.disableSerialPortAccess`.
	//
	//  Note: `projects/{project_id}/policies/{constraint_name}` is also an
	//  acceptable name for API requests, but responses will return the name using
	//  the equivalent project number.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.name
	Name *string `json:"name,omitempty"`

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

// +kcc:proto=google.cloud.orgpolicy.v2.PolicySpec
type PolicySpec struct {
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
	Etag *string `json:"etag,omitempty"`

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
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.values
	Values *PolicySpec_PolicyRule_StringValues `json:"values,omitempty"`

	// Setting this to true means that all values are allowed. This field can
	//  be set only in policies for list constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.allow_all
	AllowAll *bool `json:"allowAll,omitempty"`

	// Setting this to true means that all values are denied. This field can
	//  be set only in policies for list constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.deny_all
	DenyAll *bool `json:"denyAll,omitempty"`

	// If `true`, then the policy is enforced. If `false`, then any
	//  configuration is acceptable.
	//  This field can be set only in policies for boolean constraints.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.enforce
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
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.PolicyRule.condition
	Condition *Expr `json:"condition,omitempty"`
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

// +kcc:proto=google.cloud.orgpolicy.v2.Policy
type PolicyObservedState struct {
	// Basic information about the Organization Policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Policy.spec
	Spec *PolicySpecObservedState `json:"spec,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v2.PolicySpec
type PolicySpecObservedState struct {
	// Output only. The time stamp this was previously updated. This
	//  represents the last time a call to `CreatePolicy` or `UpdatePolicy` was
	//  made for that policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
