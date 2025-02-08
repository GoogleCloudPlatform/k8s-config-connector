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


// +kcc:proto=google.cloud.orgpolicy.v2.Constraint
type Constraint struct {
	// Immutable. The resource name of the constraint. Must be in one of
	//  the following forms:
	//
	//  * `projects/{project_number}/constraints/{constraint_name}`
	//  * `folders/{folder_id}/constraints/{constraint_name}`
	//  * `organizations/{organization_id}/constraints/{constraint_name}`
	//
	//  For example, "/projects/123/constraints/compute.disableSerialPortAccess".
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.name
	Name *string `json:"name,omitempty"`

	// The human readable name.
	//
	//  Mutable.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Detailed description of what this constraint controls as well as how and
	//  where it is enforced.
	//
	//  Mutable.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.description
	Description *string `json:"description,omitempty"`

	// The evaluation behavior of this constraint in the absence of a policy.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.constraint_default
	ConstraintDefault *string `json:"constraintDefault,omitempty"`

	// Defines this constraint as being a ListConstraint.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.list_constraint
	ListConstraint *Constraint_ListConstraint `json:"listConstraint,omitempty"`

	// Defines this constraint as being a BooleanConstraint.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.boolean_constraint
	BooleanConstraint *Constraint_BooleanConstraint `json:"booleanConstraint,omitempty"`

	// Shows if dry run is supported for this constraint or not.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.supports_dry_run
	SupportsDryRun *bool `json:"supportsDryRun,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v2.Constraint.BooleanConstraint
type Constraint_BooleanConstraint struct {
}

// +kcc:proto=google.cloud.orgpolicy.v2.Constraint.ListConstraint
type Constraint_ListConstraint struct {
	// Indicates whether values grouped into categories can be used in
	//  `Policy.allowed_values` and `Policy.denied_values`. For example,
	//  `"in:Python"` would match any value in the 'Python' group.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.ListConstraint.supports_in
	SupportsIn *bool `json:"supportsIn,omitempty"`

	// Indicates whether subtrees of the Resource Manager resource hierarchy
	//  can be used in `Policy.allowed_values` and `Policy.denied_values`. For
	//  example, `"under:folders/123"` would match any resource under the
	//  'folders/123' folder.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.Constraint.ListConstraint.supports_under
	SupportsUnder *bool `json:"supportsUnder,omitempty"`
}
