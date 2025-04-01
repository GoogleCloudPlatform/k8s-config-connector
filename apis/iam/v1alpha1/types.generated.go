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
