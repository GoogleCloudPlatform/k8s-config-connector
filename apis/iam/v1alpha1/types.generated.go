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
// krm.group: iam.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.iam.v2
// resource: IAMDenyPolicy:Policy

package v1alpha1

// +kcc:proto=google.iam.v2.DenyRule
type DenyRule struct {
	// The identities that are prevented from using one or more permissions on
	//  Google Cloud resources. This field can contain the following values:
	//
	//  * `principalSet://goog/public:all`: A special identifier that represents
	//    any principal that is on the internet, even if they do not have a Google
	//    Account or are not logged in.
	//
	//  * `principal://goog/subject/{email_id}`: A specific Google Account.
	//    Includes Gmail, Cloud Identity, and Google Workspace user accounts. For
	//    example, `principal://goog/subject/alice@example.com`.
	//
	//  * `deleted:principal://goog/subject/{email_id}?uid={uid}`: A specific
	//    Google Account that was deleted recently. For example,
	//    `deleted:principal://goog/subject/alice@example.com?uid=1234567890`. If
	//    the Google Account is recovered, this identifier reverts to the standard
	//    identifier for a Google Account.
	//
	//  * `principalSet://goog/group/{group_id}`: A Google group. For example,
	//    `principalSet://goog/group/admins@example.com`.
	//
	//  * `deleted:principalSet://goog/group/{group_id}?uid={uid}`: A Google group
	//    that was deleted recently. For example,
	//    `deleted:principalSet://goog/group/admins@example.com?uid=1234567890`. If
	//    the Google group is restored, this identifier reverts to the standard
	//    identifier for a Google group.
	//
	//  * `principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}`:
	//    A Google Cloud service account. For example,
	//    `principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com`.
	//
	//  * `deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}?uid={uid}`:
	//    A Google Cloud service account that was deleted recently. For example,
	//    `deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com?uid=1234567890`.
	//    If the service account is undeleted, this identifier reverts to the
	//    standard identifier for a service account.
	//
	//  * `principalSet://goog/cloudIdentityCustomerId/{customer_id}`: All of the
	//    principals associated with the specified Google Workspace or Cloud
	//    Identity customer ID. For example,
	//    `principalSet://goog/cloudIdentityCustomerId/C01Abc35`.
	// +kcc:proto:field=google.iam.v2.DenyRule.denied_principals
	DeniedPrincipals []string `json:"deniedPrincipals,omitempty"`

	// The identities that are excluded from the deny rule, even if they are
	//  listed in the `denied_principals`. For example, you could add a Google
	//  group to the `denied_principals`, then exclude specific users who belong to
	//  that group.
	//
	//  This field can contain the same values as the `denied_principals` field,
	//  excluding `principalSet://goog/public:all`, which represents all users on
	//  the internet.
	// +kcc:proto:field=google.iam.v2.DenyRule.exception_principals
	ExceptionPrincipals []string `json:"exceptionPrincipals,omitempty"`

	// The permissions that are explicitly denied by this rule. Each permission
	//  uses the format `{service_fqdn}/{resource}.{verb}`, where `{service_fqdn}`
	//  is the fully qualified domain name for the service. For example,
	//  `iam.googleapis.com/roles.list`.
	// +kcc:proto:field=google.iam.v2.DenyRule.denied_permissions
	DeniedPermissions []string `json:"deniedPermissions,omitempty"`

	// Specifies the permissions that this rule excludes from the set of denied
	//  permissions given by `denied_permissions`. If a permission appears in
	//  `denied_permissions` _and_ in `exception_permissions` then it will _not_ be
	//  denied.
	//
	//  The excluded permissions can be specified using the same syntax as
	//  `denied_permissions`.
	// +kcc:proto:field=google.iam.v2.DenyRule.exception_permissions
	ExceptionPermissions []string `json:"exceptionPermissions,omitempty"`

	// The condition that determines whether this deny rule applies to a request.
	//  If the condition expression evaluates to `true`, then the deny rule is
	//  applied; otherwise, the deny rule is not applied.
	//
	//  Each deny rule is evaluated independently. If this deny rule does not apply
	//  to a request, other deny rules might still apply.
	//
	//  The condition can use CEL functions that evaluate
	//  [resource
	//  tags](https://cloud.google.com/iam/help/conditions/resource-tags). Other
	//  functions and operators are not supported.
	// +kcc:proto:field=google.iam.v2.DenyRule.denial_condition
	DenialCondition *Expr `json:"denialCondition,omitempty"`
}

// +kcc:proto=google.iam.v2.PolicyRule
type PolicyRule struct {
	// A rule for a deny policy.
	// +kcc:proto:field=google.iam.v2.PolicyRule.deny_rule
	DenyRule *DenyRule `json:"denyRule,omitempty"`

	// A user-specified description of the rule. This value can be up to 256
	//  characters.
	// +kcc:proto:field=google.iam.v2.PolicyRule.description
	Description *string `json:"description,omitempty"`
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
