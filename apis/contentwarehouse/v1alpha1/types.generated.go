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


// +kcc:proto=google.cloud.contentwarehouse.v1.AccessControlAction
type AccessControlAction struct {
	// Identifies the type of operation.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.AccessControlAction.operation_type
	OperationType *string `json:"operationType,omitempty"`

	// Represents the new policy from which bindings are added, removed or
	//  replaced based on the type of the operation. the policy is limited to a few
	//  10s of KB.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.AccessControlAction.policy
	Policy *Policy `json:"policy,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.Action
type Action struct {
	// ID of the action. Managed internally.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.action_id
	ActionID *string `json:"actionID,omitempty"`

	// Action triggering access control operations.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.access_control
	AccessControl *AccessControlAction `json:"accessControl,omitempty"`

	// Action triggering data validation operations.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.data_validation
	DataValidation *DataValidationAction `json:"dataValidation,omitempty"`

	// Action triggering data update operations.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.data_update
	DataUpdate *DataUpdateAction `json:"dataUpdate,omitempty"`

	// Action triggering create document link operation.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.add_to_folder
	AddToFolder *AddToFolderAction `json:"addToFolder,omitempty"`

	// Action publish to Pub/Sub operation.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.publish_to_pub_sub
	PublishToPubSub *PublishAction `json:"publishToPubSub,omitempty"`

	// Action removing a document from a folder.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.remove_from_folder_action
	RemoveFromFolderAction *RemoveFromFolderAction `json:"removeFromFolderAction,omitempty"`

	// Action deleting the document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Action.delete_document_action
	DeleteDocumentAction *DeleteDocumentAction `json:"deleteDocumentAction,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.AddToFolderAction
type AddToFolderAction struct {
	// Names of the folder under which new document is to be added.
	//  Format:
	//  projects/{project_number}/locations/{location}/documents/{document_id}.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.AddToFolderAction.folders
	Folders []string `json:"folders,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DataUpdateAction
type DataUpdateAction struct {
	// Map of (K, V) -> (valid name of the field, new value of the field)
	//  E.g., ("age", "60") entry triggers update of field age with a value of 60.
	//  If the field is not present then new entry is added.
	//  During update action execution, value strings will be casted to
	//  appropriate types.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DataUpdateAction.entries
	Entries map[string]string `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DataValidationAction
type DataValidationAction struct {
	// Map of (K, V) -> (field, string condition to be evaluated on the field)
	//  E.g., ("age", "age > 18  && age < 60") entry triggers validation of field
	//  age with the given condition. Map entries will be ANDed during validation.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DataValidationAction.conditions
	Conditions map[string]string `json:"conditions,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DeleteDocumentAction
type DeleteDocumentAction struct {
	// Boolean field to select between hard vs soft delete options.
	//  Set 'true' for 'hard delete' and 'false' for 'soft delete'.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DeleteDocumentAction.enable_hard_delete
	EnableHardDelete *bool `json:"enableHardDelete,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PublishAction
type PublishAction struct {
	// The topic id in the Pub/Sub service for which messages will be published
	//  to.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PublishAction.topic_id
	TopicID *string `json:"topicID,omitempty"`

	// Messages to be published.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PublishAction.messages
	Messages []string `json:"messages,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.RemoveFromFolderAction
type RemoveFromFolderAction struct {
	// Condition of the action to be executed.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.RemoveFromFolderAction.condition
	Condition *string `json:"condition,omitempty"`

	// Name of the folder under which new document is to be added.
	//  Format:
	//  projects/{project_number}/locations/{location}/documents/{document_id}.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.RemoveFromFolderAction.folder
	Folder *string `json:"folder,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.Rule
type Rule struct {
	// Short description of the rule and its context.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Rule.description
	Description *string `json:"description,omitempty"`

	// ID of the rule. It has to be unique across all the examples.
	//  This is managed internally.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Rule.rule_id
	RuleID *string `json:"ruleID,omitempty"`

	// Identifies the trigger type for running the policy.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Rule.trigger_type
	TriggerType *string `json:"triggerType,omitempty"`

	// Represents the conditional expression to be evaluated.
	//  Expression should evaluate to a boolean result.
	//  When the condition is true actions are executed.
	//  Example: user_role = "hsbc_role_1" AND doc.salary > 20000
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Rule.condition
	Condition *string `json:"condition,omitempty"`

	// List of actions that are executed when the rule is satisfied.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Rule.actions
	Actions []Action `json:"actions,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.RuleSet
type RuleSet struct {
	// The resource name of the rule set. Managed internally.
	//  Format:
	//  projects/{project_number}/locations/{location}/ruleSet/{rule_set_id}.
	//
	//  The name is ignored when creating a rule set.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.RuleSet.name
	Name *string `json:"name,omitempty"`

	// Short description of the rule-set.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.RuleSet.description
	Description *string `json:"description,omitempty"`

	// Source of the rules i.e., customer name.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.RuleSet.source
	Source *string `json:"source,omitempty"`

	// List of rules given by the customer.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.RuleSet.rules
	Rules []Rule `json:"rules,omitempty"`
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
