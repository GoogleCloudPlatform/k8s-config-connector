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

// +kcc:proto=google.cloud.orgpolicy.v2.CustomConstraint
type CustomConstraint struct {
	// Immutable. Name of the constraint. This is unique within the organization.
	//  Format of the name should be
	//
	//  * `organizations/{organization_id}/customConstraints/{custom_constraint_id}`
	//
	//  Example: `organizations/123/customConstraints/custom.createOnlyE2TypeVms`
	//
	//  The max length is 70 characters and the minimum length is 1. Note that the
	//  prefix `organizations/{organization_id}/customConstraints/` is not counted.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.name
	Name *string `json:"name,omitempty"`

	// Immutable. The resource instance type on which this policy applies. Format
	//  will be of the form : `<canonical service name>/<type>` Example:
	//
	//   * `compute.googleapis.com/Instance`.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.resource_types
	ResourceTypes []string `json:"resourceTypes,omitempty"`

	// All the operations being applied for this constraint.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.method_types
	MethodTypes []string `json:"methodTypes,omitempty"`

	// Org policy condition/expression. For example:
	//  `resource.instanceName.matches("[production|test]_.*_(\d)+")` or,
	//  `resource.management.auto_upgrade == true`
	//
	//  The max length of the condition is 1000 characters.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.condition
	Condition *string `json:"condition,omitempty"`

	// Allow or deny type.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.action_type
	ActionType *string `json:"actionType,omitempty"`

	// One line display name for the UI.
	//  The max length of the display_name is 200 characters.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Detailed information about this custom policy constraint.
	//  The max length of the description is 2000 characters.
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v2.CustomConstraint
type CustomConstraintObservedState struct {
	// Output only. The last time this custom constraint was updated. This
	//  represents the last time that the `CreateCustomConstraint` or
	//  `UpdateCustomConstraint` RPC was called
	// +kcc:proto:field=google.cloud.orgpolicy.v2.CustomConstraint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
