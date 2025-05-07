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

// +kcc:proto=google.cloud.asset.v1.IamPolicyAnalysisQuery.AccessSelector
type IAMPolicyAnalysisQuery_AccessSelector struct {
	// Optional. The roles to appear in result.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.AccessSelector.roles
	Roles []string `json:"roles,omitempty"`

	// Optional. The permissions to appear in result.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.AccessSelector.permissions
	Permissions []string `json:"permissions,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.IamPolicyAnalysisQuery.ConditionContext
type IAMPolicyAnalysisQuery_ConditionContext struct {
	// The hypothetical access timestamp to evaluate IAM conditions. Note that
	//  this value must not be earlier than the current time; otherwise, an
	//  INVALID_ARGUMENT error will be returned.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.ConditionContext.access_time
	AccessTime *string `json:"accessTime,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.IamPolicyAnalysisQuery.Options
type IAMPolicyAnalysisQuery_Options struct {
	// Optional. If true, the identities section of the result will expand any
	//  Google groups appearing in an IAM policy binding.
	//
	//  If
	//  [IamPolicyAnalysisQuery.identity_selector][google.cloud.asset.v1.IamPolicyAnalysisQuery.identity_selector]
	//  is specified, the identity in the result will be determined by the
	//  selector, and this flag is not allowed to set.
	//
	//  If true, the default max expansion per group is 1000 for
	//  AssetService.AnalyzeIamPolicy][].
	//
	//  Default is false.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.Options.expand_groups
	ExpandGroups *bool `json:"expandGroups,omitempty"`

	// Optional. If true, the access section of result will expand any roles
	//  appearing in IAM policy bindings to include their permissions.
	//
	//  If
	//  [IamPolicyAnalysisQuery.access_selector][google.cloud.asset.v1.IamPolicyAnalysisQuery.access_selector]
	//  is specified, the access section of the result will be determined by the
	//  selector, and this flag is not allowed to set.
	//
	//  Default is false.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.Options.expand_roles
	ExpandRoles *bool `json:"expandRoles,omitempty"`

	// Optional. If true and
	//  [IamPolicyAnalysisQuery.resource_selector][google.cloud.asset.v1.IamPolicyAnalysisQuery.resource_selector]
	//  is not specified, the resource section of the result will expand any
	//  resource attached to an IAM policy to include resources lower in the
	//  resource hierarchy.
	//
	//  For example, if the request analyzes for which resources user A has
	//  permission P, and the results include an IAM policy with P on a Google
	//  Cloud folder, the results will also include resources in that folder with
	//  permission P.
	//
	//  If true and
	//  [IamPolicyAnalysisQuery.resource_selector][google.cloud.asset.v1.IamPolicyAnalysisQuery.resource_selector]
	//  is specified, the resource section of the result will expand the
	//  specified resource to include resources lower in the resource hierarchy.
	//  Only project or lower resources are supported. Folder and organization
	//  resources cannot be used together with this option.
	//
	//  For example, if the request analyzes for which users have permission P on
	//  a Google Cloud project with this option enabled, the results will include
	//  all users who have permission P on that project or any lower resource.
	//
	//  If true, the default max expansion per resource is 1000 for
	//  AssetService.AnalyzeIamPolicy][] and 100000 for
	//  AssetService.AnalyzeIamPolicyLongrunning][].
	//
	//  Default is false.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.Options.expand_resources
	ExpandResources *bool `json:"expandResources,omitempty"`

	// Optional. If true, the result will output the relevant parent/child
	//  relationships between resources. Default is false.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.Options.output_resource_edges
	OutputResourceEdges *bool `json:"outputResourceEdges,omitempty"`

	// Optional. If true, the result will output the relevant membership
	//  relationships between groups and other groups, and between groups and
	//  principals. Default is false.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.Options.output_group_edges
	OutputGroupEdges *bool `json:"outputGroupEdges,omitempty"`

	// Optional. If true, the response will include access analysis from
	//  identities to resources via service account impersonation. This is a very
	//  expensive operation, because many derived queries will be executed. We
	//  highly recommend you use
	//  [AssetService.AnalyzeIamPolicyLongrunning][google.cloud.asset.v1.AssetService.AnalyzeIamPolicyLongrunning]
	//  RPC instead.
	//
	//  For example, if the request analyzes for which resources user A has
	//  permission P, and there's an IAM policy states user A has
	//  iam.serviceAccounts.getAccessToken permission to a service account SA,
	//  and there's another IAM policy states service account SA has permission P
	//  to a Google Cloud folder F, then user A potentially has access to the
	//  Google Cloud folder F. And those advanced analysis results will be
	//  included in
	//  [AnalyzeIamPolicyResponse.service_account_impersonation_analysis][google.cloud.asset.v1.AnalyzeIamPolicyResponse.service_account_impersonation_analysis].
	//
	//  Another example, if the request analyzes for who has
	//  permission P to a Google Cloud folder F, and there's an IAM policy states
	//  user A has iam.serviceAccounts.actAs permission to a service account SA,
	//  and there's another IAM policy states service account SA has permission P
	//  to the Google Cloud folder F, then user A potentially has access to the
	//  Google Cloud folder F. And those advanced analysis results will be
	//  included in
	//  [AnalyzeIamPolicyResponse.service_account_impersonation_analysis][google.cloud.asset.v1.AnalyzeIamPolicyResponse.service_account_impersonation_analysis].
	//
	//  Only the following permissions are considered in this analysis:
	//
	//  * `iam.serviceAccounts.actAs`
	//  * `iam.serviceAccounts.signBlob`
	//  * `iam.serviceAccounts.signJwt`
	//  * `iam.serviceAccounts.getAccessToken`
	//  * `iam.serviceAccounts.getOpenIdToken`
	//  * `iam.serviceAccounts.implicitDelegation`
	//
	//  Default is false.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.Options.analyze_service_account_impersonation
	AnalyzeServiceAccountImpersonation *bool `json:"analyzeServiceAccountImpersonation,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.SavedQuery.QueryContent
type SavedQuery_QueryContent struct {
	// An IAM Policy Analysis query, which could be used in
	//  the
	//  [AssetService.AnalyzeIamPolicy][google.cloud.asset.v1.AssetService.AnalyzeIamPolicy]
	//  RPC or the
	//  [AssetService.AnalyzeIamPolicyLongrunning][google.cloud.asset.v1.AssetService.AnalyzeIamPolicyLongrunning]
	//  RPC.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.QueryContent.iam_policy_analysis_query
	IAMPolicyAnalysisQuery *IAMPolicyAnalysisQuery `json:"iamPolicyAnalysisQuery,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.FeedOutputConfig
type FeedOutputConfig struct {
	// Destination on Pub/Sub.
	// +kcc:proto:field=google.cloud.asset.v1.FeedOutputConfig.pubsub_destination
	PubsubDestination *PubsubDestination `json:"pubsubDestination,omitempty"`
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
