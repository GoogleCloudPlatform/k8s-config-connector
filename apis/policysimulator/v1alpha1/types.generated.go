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


// +kcc:proto=google.cloud.policysimulator.v1.AccessStateDiff
type AccessStateDiff struct {
	// The results of evaluating the access tuple under the current (baseline)
	//  policies.
	//
	//  If the [AccessState][google.cloud.policysimulator.v1.AccessState] couldn't
	//  be fully evaluated, this field explains why.
	// +kcc:proto:field=google.cloud.policysimulator.v1.AccessStateDiff.baseline
	Baseline *ExplainedAccess `json:"baseline,omitempty"`

	// The results of evaluating the access tuple under the proposed (simulated)
	//  policies.
	//
	//  If the AccessState couldn't be fully evaluated, this field explains why.
	// +kcc:proto:field=google.cloud.policysimulator.v1.AccessStateDiff.simulated
	Simulated *ExplainedAccess `json:"simulated,omitempty"`

	// How the principal's access, specified in the AccessState field, changed
	//  between the current (baseline) policies and proposed (simulated) policies.
	// +kcc:proto:field=google.cloud.policysimulator.v1.AccessStateDiff.access_change
	AccessChange *string `json:"accessChange,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.AccessTuple
type AccessTuple struct {
	// Required. The principal whose access you want to check, in the form of
	//  the email address that represents that principal. For example,
	//  `alice@example.com` or
	//  `my-service-account@my-project.iam.gserviceaccount.com`.
	//
	//  The principal must be a Google Account or a service account. Other types of
	//  principals are not supported.
	// +kcc:proto:field=google.cloud.policysimulator.v1.AccessTuple.principal
	Principal *string `json:"principal,omitempty"`

	// Required. The full resource name that identifies the resource. For example,
	//  `//compute.googleapis.com/projects/my-project/zones/us-central1-a/instances/my-instance`.
	//
	//  For examples of full resource names for Google Cloud services, see
	//  https://cloud.google.com/iam/help/troubleshooter/full-resource-names.
	// +kcc:proto:field=google.cloud.policysimulator.v1.AccessTuple.full_resource_name
	FullResourceName *string `json:"fullResourceName,omitempty"`

	// Required. The IAM permission to check for the specified principal and
	//  resource.
	//
	//  For a complete list of IAM permissions, see
	//  https://cloud.google.com/iam/help/permissions/reference.
	//
	//  For a complete list of predefined IAM roles and the permissions in each
	//  role, see https://cloud.google.com/iam/help/roles/reference.
	// +kcc:proto:field=google.cloud.policysimulator.v1.AccessTuple.permission
	Permission *string `json:"permission,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.BindingExplanation
type BindingExplanation struct {
	// Required. Indicates whether _this binding_ provides the specified
	//  permission to the specified principal for the specified resource.
	//
	//  This field does _not_ indicate whether the principal actually has the
	//  permission for the resource. There might be another binding that overrides
	//  this binding. To determine whether the principal actually has the
	//  permission, use the `access` field in the
	//  [TroubleshootIamPolicyResponse][IamChecker.TroubleshootIamPolicyResponse].
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.access
	Access *string `json:"access,omitempty"`

	// The role that this binding grants. For example,
	//  `roles/compute.serviceAgent`.
	//
	//  For a complete list of predefined IAM roles, as well as the permissions in
	//  each role, see https://cloud.google.com/iam/help/roles/reference.
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.role
	Role *string `json:"role,omitempty"`

	// Indicates whether the role granted by this binding contains the specified
	//  permission.
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.role_permission
	RolePermission *string `json:"rolePermission,omitempty"`

	// The relevance of the permission's existence, or nonexistence, in the role
	//  to the overall determination for the entire policy.
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.role_permission_relevance
	RolePermissionRelevance *string `json:"rolePermissionRelevance,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The relevance of this binding to the overall determination for the entire
	//  policy.
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.relevance
	Relevance *string `json:"relevance,omitempty"`

	// A condition expression that prevents this binding from granting access
	//  unless the expression evaluates to `true`.
	//
	//  To learn about IAM Conditions, see
	//  https://cloud.google.com/iam/docs/conditions-overview.
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.condition
	Condition *Expr `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.BindingExplanation.AnnotatedMembership
type BindingExplanation_AnnotatedMembership struct {
	// Indicates whether the binding includes the principal.
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.AnnotatedMembership.membership
	Membership *string `json:"membership,omitempty"`

	// The relevance of the principal's status to the overall determination for
	//  the binding.
	// +kcc:proto:field=google.cloud.policysimulator.v1.BindingExplanation.AnnotatedMembership.relevance
	Relevance *string `json:"relevance,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.ExplainedAccess
type ExplainedAccess struct {
	// Whether the principal in the access tuple has permission to access the
	//  resource in the access tuple under the given policies.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedAccess.access_state
	AccessState *string `json:"accessState,omitempty"`

	// If the [AccessState][google.cloud.policysimulator.v1.AccessState] is
	//  `UNKNOWN`, this field contains the policies that led to that result.
	//
	//  If the `AccessState` is `GRANTED` or `NOT_GRANTED`, this field is
	//  omitted.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedAccess.policies
	Policies []ExplainedPolicy `json:"policies,omitempty"`

	// If the [AccessState][google.cloud.policysimulator.v1.AccessState] is
	//  `UNKNOWN`, this field contains a list of errors explaining why the result
	//  is `UNKNOWN`.
	//
	//  If the `AccessState` is `GRANTED` or `NOT_GRANTED`, this field is
	//  omitted.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedAccess.errors
	Errors []Status `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.ExplainedPolicy
type ExplainedPolicy struct {
	// Indicates whether _this policy_ provides the specified permission to the
	//  specified principal for the specified resource.
	//
	//  This field does _not_ indicate whether the principal actually has the
	//  permission for the resource. There might be another policy that overrides
	//  this policy. To determine whether the principal actually has the
	//  permission, use the `access` field in the
	//  [TroubleshootIamPolicyResponse][IamChecker.TroubleshootIamPolicyResponse].
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedPolicy.access
	Access *string `json:"access,omitempty"`

	// The full resource name that identifies the resource. For example,
	//  `//compute.googleapis.com/projects/my-project/zones/us-central1-a/instances/my-instance`.
	//
	//  If the user who created the
	//  [Replay][google.cloud.policysimulator.v1.Replay] does not have
	//  access to the policy, this field is omitted.
	//
	//  For examples of full resource names for Google Cloud services, see
	//  https://cloud.google.com/iam/help/troubleshooter/full-resource-names.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedPolicy.full_resource_name
	FullResourceName *string `json:"fullResourceName,omitempty"`

	// The IAM policy attached to the resource.
	//
	//  If the user who created the
	//  [Replay][google.cloud.policysimulator.v1.Replay] does not have
	//  access to the policy, this field is empty.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedPolicy.policy
	Policy *Policy `json:"policy,omitempty"`

	// Details about how each binding in the policy affects the principal's
	//  ability, or inability, to use the permission for the resource.
	//
	//  If the user who created the
	//  [Replay][google.cloud.policysimulator.v1.Replay] does not have
	//  access to the policy, this field is omitted.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedPolicy.binding_explanations
	BindingExplanations []BindingExplanation `json:"bindingExplanations,omitempty"`

	// The relevance of this policy to the overall determination in the
	//  [TroubleshootIamPolicyResponse][IamChecker.TroubleshootIamPolicyResponse].
	//
	//  If the user who created the
	//  [Replay][google.cloud.policysimulator.v1.Replay] does not have
	//  access to the policy, this field is omitted.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ExplainedPolicy.relevance
	Relevance *string `json:"relevance,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.ReplayDiff
type ReplayDiff struct {
	// A summary and comparison of the principal's access under the current
	//  (baseline) policies and the proposed (simulated) policies for a single
	//  access tuple.
	//
	//  The evaluation of the principal's access is reported in the
	//  [AccessState][google.cloud.policysimulator.v1.AccessState] field.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayDiff.access_diff
	AccessDiff *AccessStateDiff `json:"accessDiff,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.ReplayResult
type ReplayResult struct {
	// The difference between the principal's access under the current
	//  (baseline) policies and the principal's access under the proposed
	//  (simulated) policies.
	//
	//  This field is only included for access tuples that were successfully
	//  replayed and had different results under the current policies and the
	//  proposed policies.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayResult.diff
	Diff *ReplayDiff `json:"diff,omitempty"`

	// The error that caused the access tuple replay to fail.
	//
	//  This field is only included for access tuples that were not replayed
	//  successfully.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayResult.error
	Error *Status `json:"error,omitempty"`

	// The resource name of the `ReplayResult`, in the following format:
	//
	//  `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}/results/{replay-result-id}`,
	//  where `{resource-id}` is the ID of the project, folder, or organization
	//  that owns the [Replay][google.cloud.policysimulator.v1.Replay].
	//
	//  Example:
	//  `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36/results/1234`
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayResult.name
	Name *string `json:"name,omitempty"`

	// The [Replay][google.cloud.policysimulator.v1.Replay] that the access tuple
	//  was included in.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayResult.parent
	Parent *string `json:"parent,omitempty"`

	// The access tuple that was replayed. This field includes information about
	//  the principal, resource, and permission that were involved in the access
	//  attempt.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayResult.access_tuple
	AccessTuple *AccessTuple `json:"accessTuple,omitempty"`

	// The latest date this access tuple was seen in the logs.
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayResult.last_seen_date
	LastSeenDate *Date `json:"lastSeenDate,omitempty"`
}

// +kcc:proto=google.iam.v1.AuditConfig
type AuditConfig struct {
	// Specifies a service that will be enabled for audit logging.
	//  For example, `storage.googleapis.com`, `cloudsql.googleapis.com`.
	//  `allServices` is a special value that covers all services.
	// +kcc:proto:field=google.iam.v1.AuditConfig.service
	Service *string `json:"service,omitempty"`

	// The configuration for logging of each type of permission.
	// +kcc:proto:field=google.iam.v1.AuditConfig.audit_log_configs
	AuditLogConfigs []AuditLogConfig `json:"auditLogConfigs,omitempty"`
}

// +kcc:proto=google.iam.v1.AuditLogConfig
type AuditLogConfig struct {
	// The log type that this config enables.
	// +kcc:proto:field=google.iam.v1.AuditLogConfig.log_type
	LogType *string `json:"logType,omitempty"`

	// Specifies the identities that do not cause logging for this type of
	//  permission.
	//  Follows the same format of
	//  [Binding.members][google.iam.v1.Binding.members].
	// +kcc:proto:field=google.iam.v1.AuditLogConfig.exempted_members
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`
}

// +kcc:proto=google.iam.v1.Binding
type Binding struct {
	// Role that is assigned to the list of `members`, or principals.
	//  For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	// +kcc:proto:field=google.iam.v1.Binding.role
	Role *string `json:"role,omitempty"`

	// Specifies the principals requesting access for a Google Cloud resource.
	//  `members` can have the following values:
	//
	//  * `allUsers`: A special identifier that represents anyone who is
	//     on the internet; with or without a Google account.
	//
	//  * `allAuthenticatedUsers`: A special identifier that represents anyone
	//     who is authenticated with a Google account or a service account.
	//
	//  * `user:{emailid}`: An email address that represents a specific Google
	//     account. For example, `alice@example.com` .
	//
	//
	//  * `serviceAccount:{emailid}`: An email address that represents a service
	//     account. For example, `my-other-app@appspot.gserviceaccount.com`.
	//
	//  * `group:{emailid}`: An email address that represents a Google group.
	//     For example, `admins@example.com`.
	//
	//  * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique
	//     identifier) representing a user that has been recently deleted. For
	//     example, `alice@example.com?uid=123456789012345678901`. If the user is
	//     recovered, this value reverts to `user:{emailid}` and the recovered user
	//     retains the role in the binding.
	//
	//  * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus
	//     unique identifier) representing a service account that has been recently
	//     deleted. For example,
	//     `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`.
	//     If the service account is undeleted, this value reverts to
	//     `serviceAccount:{emailid}` and the undeleted service account retains the
	//     role in the binding.
	//
	//  * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique
	//     identifier) representing a Google group that has been recently
	//     deleted. For example, `admins@example.com?uid=123456789012345678901`. If
	//     the group is recovered, this value reverts to `group:{emailid}` and the
	//     recovered group retains the role in the binding.
	//
	//
	//  * `domain:{domain}`: The G Suite domain (primary) that represents all the
	//     users of that domain. For example, `google.com` or `example.com`.
	// +kcc:proto:field=google.iam.v1.Binding.members
	Members []string `json:"members,omitempty"`

	// The condition that is associated with this binding.
	//
	//  If the condition evaluates to `true`, then this binding applies to the
	//  current request.
	//
	//  If the condition evaluates to `false`, then this binding does not apply to
	//  the current request. However, a different role binding might grant the same
	//  role to one or more of the principals in this binding.
	//
	//  To learn which resources support conditions in their IAM policies, see the
	//  [IAM
	//  documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	// +kcc:proto:field=google.iam.v1.Binding.condition
	Condition *Expr `json:"condition,omitempty"`
}

// +kcc:proto=google.iam.v1.Policy
type Policy struct {
	// Specifies the format of the policy.
	//
	//  Valid values are `0`, `1`, and `3`. Requests that specify an invalid value
	//  are rejected.
	//
	//  Any operation that affects conditional role bindings must specify version
	//  `3`. This requirement applies to the following operations:
	//
	//  * Getting a policy that includes a conditional role binding
	//  * Adding a conditional role binding to a policy
	//  * Changing a conditional role binding in a policy
	//  * Removing any role binding, with or without a condition, from a policy
	//    that includes conditions
	//
	//  **Important:** If you use IAM Conditions, you must include the `etag` field
	//  whenever you call `setIamPolicy`. If you omit this field, then IAM allows
	//  you to overwrite a version `3` policy with a version `1` policy, and all of
	//  the conditions in the version `3` policy are lost.
	//
	//  If a policy does not include any conditions, operations on that policy may
	//  specify any valid version or leave the field unset.
	//
	//  To learn which resources support conditions in their IAM policies, see the
	//  [IAM
	//  documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	// +kcc:proto:field=google.iam.v1.Policy.version
	Version *int32 `json:"version,omitempty"`

	// Associates a list of `members`, or principals, with a `role`. Optionally,
	//  may specify a `condition` that determines how and when the `bindings` are
	//  applied. Each of the `bindings` must contain at least one principal.
	//
	//  The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250
	//  of these principals can be Google groups. Each occurrence of a principal
	//  counts towards these limits. For example, if the `bindings` grant 50
	//  different roles to `user:alice@example.com`, and not to any other
	//  principal, then you can add another 1,450 principals to the `bindings` in
	//  the `Policy`.
	// +kcc:proto:field=google.iam.v1.Policy.bindings
	Bindings []Binding `json:"bindings,omitempty"`

	// Specifies cloud audit logging configuration for this policy.
	// +kcc:proto:field=google.iam.v1.Policy.audit_configs
	AuditConfigs []AuditConfig `json:"auditConfigs,omitempty"`

	// `etag` is used for optimistic concurrency control as a way to help
	//  prevent simultaneous updates of a policy from overwriting each other.
	//  It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform policy updates in order to avoid race
	//  conditions: An `etag` is returned in the response to `getIamPolicy`, and
	//  systems are expected to put that etag in the request to `setIamPolicy` to
	//  ensure that their change will be applied to the same version of the policy.
	//
	//  **Important:** If you use IAM Conditions, you must include the `etag` field
	//  whenever you call `setIamPolicy`. If you omit this field, then IAM allows
	//  you to overwrite a version `3` policy with a version `1` policy, and all of
	//  the conditions in the version `3` policy are lost.
	// +kcc:proto:field=google.iam.v1.Policy.etag
	Etag []byte `json:"etag,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
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
