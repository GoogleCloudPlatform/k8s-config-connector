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


// +kcc:proto=google.cloud.policysimulator.v1.Replay
type Replay struct {

	// Required. The configuration used for the `Replay`.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.config
	Config *ReplayConfig `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.Replay.ResultsSummary
type Replay_ResultsSummary struct {
	// The total number of log entries replayed.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.ResultsSummary.log_count
	LogCount *int32 `json:"logCount,omitempty"`

	// The number of replayed log entries with no difference between
	//  baseline and simulated policies.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.ResultsSummary.unchanged_count
	UnchangedCount *int32 `json:"unchangedCount,omitempty"`

	// The number of replayed log entries with a difference between baseline and
	//  simulated policies.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.ResultsSummary.difference_count
	DifferenceCount *int32 `json:"differenceCount,omitempty"`

	// The number of log entries that could not be replayed.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.ResultsSummary.error_count
	ErrorCount *int32 `json:"errorCount,omitempty"`

	// The date of the oldest log entry replayed.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.ResultsSummary.oldest_date
	OldestDate *Date `json:"oldestDate,omitempty"`

	// The date of the newest log entry replayed.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.ResultsSummary.newest_date
	NewestDate *Date `json:"newestDate,omitempty"`
}

// +kcc:proto=google.cloud.policysimulator.v1.ReplayConfig
type ReplayConfig struct {

	// TODO: unsupported map type with key string and value message


	// The logs to use as input for the
	//  [Replay][google.cloud.policysimulator.v1.Replay].
	// +kcc:proto:field=google.cloud.policysimulator.v1.ReplayConfig.log_source
	LogSource *string `json:"logSource,omitempty"`
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

// +kcc:proto=google.cloud.policysimulator.v1.Replay
type ReplayObservedState struct {
	// Output only. The resource name of the `Replay`, which has the following
	//  format:
	//
	//  `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`,
	//  where `{resource-id}` is the ID of the project, folder, or organization
	//  that owns the Replay.
	//
	//  Example:
	//  `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.name
	Name *string `json:"name,omitempty"`

	// Output only. The current state of the `Replay`.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.state
	State *string `json:"state,omitempty"`

	// Output only. Summary statistics about the replayed log entries.
	// +kcc:proto:field=google.cloud.policysimulator.v1.Replay.results_summary
	ResultsSummary *Replay_ResultsSummary `json:"resultsSummary,omitempty"`
}
