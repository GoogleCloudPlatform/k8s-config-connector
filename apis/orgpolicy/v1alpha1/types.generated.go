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

// +generated:types
// krm.group: orgpolicy.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.orgpolicy.v2
// resource: OrgPolicyPolicy:Policy
// resource: OrgPolicyCustomConstraint:CustomConstraint

package v1alpha1

// +kcc:proto=google.cloud.orgpolicy.v2.AlternatePolicySpec
type AlternatePolicySpec struct {
	// Reference to the launch that will be used while audit logging and to
	//  control the launch.
	//  Should be set only in the alternate policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.AlternatePolicySpec.launch
	Launch *string `json:"launch,omitempty"`

	// Specify constraint for configurations of Google Cloud resources.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.AlternatePolicySpec.spec
	Spec *PolicySpec `json:"spec,omitempty"`
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

// +kcc:observedstate:proto=google.cloud.orgpolicy.v2.PolicySpec
type PolicySpecObservedState struct {
	// Output only. The time stamp this was previously updated. This
	//  represents the last time a call to `CreatePolicy` or `UpdatePolicy` was
	//  made for that policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.PolicySpec.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
