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
// proto.service: google.iam.v3
// resource: IAMPolicyBinding:PolicyBinding

package v1alpha1

// +kcc:proto=google.iam.v3.PolicyBinding
type PolicyBinding struct {
	// Identifier. The name of the policy binding, in the format
	//  `{binding_parent/locations/{location}/policyBindings/{policy_binding_id}`.
	//  The binding parent is the closest Resource Manager resource (i.e., Project,
	//  Folder or Organization) to the binding target.
	//
	//  Format:
	//
	//  * `projects/{project_id}/locations/{location}/policyBindings/{policy_binding_id}`
	//  * `projects/{project_number}/locations/{location}/policyBindings/{policy_binding_id}`
	//  * `folders/{folder_id}/locations/{location}/policyBindings/{policy_binding_id}`
	//  * `organizations/{organization_id}/locations/{location}/policyBindings/{policy_binding_id}`
	// +kcc:proto:field=google.iam.v3.PolicyBinding.name
	Name *string `json:"name,omitempty"`

	// Optional. The etag for the policy binding.
	//  If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The description of the policy binding. Must be less than or equal
	//  to 63 characters.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined annotations. See
	//  https://google.aip.dev/148#annotations for more details such as format and
	//  size limitations
	// +kcc:proto:field=google.iam.v3.PolicyBinding.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. Immutable. Target is the full resource name of the resource to
	//  which the policy will be bound. Immutable once set.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.target
	Target *PolicyBinding_Target `json:"target,omitempty"`

	// Immutable. The kind of the policy to attach in this binding. This field
	//  must be one of the following:
	//
	//  - Left empty (will be automatically set to the policy kind)
	//  - The input policy kind
	// +kcc:proto:field=google.iam.v3.PolicyBinding.policy_kind
	PolicyKind *string `json:"policyKind,omitempty"`

	// Required. Immutable. The resource name of the policy to be bound. The
	//  binding parent and policy must belong to the same Organization (or
	//  Project).
	// +kcc:proto:field=google.iam.v3.PolicyBinding.policy
	Policy *string `json:"policy,omitempty"`

	// Optional. Condition can either be a principal condition or a resource
	//  condition. It depends on the type of target, the policy it is attached to,
	//  and/or the expression itself. When set, the `expression` field in the
	//  `Expr` must include from 1 to 10 subexpressions, joined by the "||"(Logical
	//  OR),
	//  "&&"(Logical AND) or "!"(Logical NOT) operators and cannot contain more
	//  than 250 characters.
	//  Allowed operations for principal.subject:
	//
	//  - `principal.subject == <principal subject string>`
	//  - `principal.subject != <principal subject string>`
	//  - `principal.subject in [<list of principal subjects>]`
	//  - `principal.subject.startsWith(<string>)`
	//  - `principal.subject.endsWith(<string>)`
	//
	//  Allowed operations for principal.type:
	//
	//  - `principal.type == <principal type string>`
	//  - `principal.type != <principal type string>`
	//  - `principal.type in [<list of principal types>]`
	//
	//  Supported principal types are Workspace, Workforce Pool, Workload Pool and
	//  Service Account. Allowed string must be one of:
	//
	//  - iam.googleapis.com/WorkspaceIdentity
	//  - iam.googleapis.com/WorkforcePoolIdentity
	//  - iam.googleapis.com/WorkloadPoolIdentity
	//  - iam.googleapis.com/ServiceAccount
	//
	//  When the bound policy is a principal access boundary policy, the only
	//  supported attributes in any subexpression are `principal.type` and
	//  `principal.subject`. An example expression is: "principal.type ==
	//  'iam.googleapis.com/ServiceAccount'" or "principal.subject ==
	//  'bob@example.com'".
	// +kcc:proto:field=google.iam.v3.PolicyBinding.condition
	Condition *Expr `json:"condition,omitempty"`
}

// +kcc:proto=google.iam.v3.PolicyBinding.Target
type PolicyBinding_Target struct {
	// Immutable. Full Resource Name used for principal access boundary policy
	//  bindings Examples:
	//
	//  * Organization:
	//  `//cloudresourcemanager.googleapis.com/organizations/ORGANIZATION_ID`
	//  * Folder: `//cloudresourcemanager.googleapis.com/folders/FOLDER_ID`
	//  * Project:
	//      * `//cloudresourcemanager.googleapis.com/projects/PROJECT_NUMBER`
	//      * `//cloudresourcemanager.googleapis.com/projects/PROJECT_ID`
	//  * Workload Identity Pool:
	//  `//iam.googleapis.com/projects/PROJECT_NUMBER/locations/LOCATION/workloadIdentityPools/WORKLOAD_POOL_ID`
	//  * Workforce Identity:
	//  `//iam.googleapis.com/locations/global/workforcePools/WORKFORCE_POOL_ID`
	//  * Workspace Identity:
	//  `//iam.googleapis.com/locations/global/workspace/WORKSPACE_ID`
	// +kcc:proto:field=google.iam.v3.PolicyBinding.Target.principal_set
	PrincipalSet *string `json:"principalSet,omitempty"`
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

// +kcc:proto=google.iam.v3.PolicyBinding
type PolicyBindingObservedState struct {
	// Output only. The globally unique ID of the policy binding. Assigned when
	//  the policy binding is created.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The globally unique ID of the policy to be bound.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.policy_uid
	PolicyUid *string `json:"policyUid,omitempty"`

	// Output only. The time when the policy binding was created.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the policy binding was most recently updated.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
