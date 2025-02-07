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


// +kcc:proto=google.cloud.asset.v1p7beta1.Asset
type Asset struct {
	// The last update timestamp of an asset. update_time is updated when
	//  create/update/delete operation is performed.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The full name of the asset. Example:
	//  `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`
	//
	//  See [Resource
	//  names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	//  for more information.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.name
	Name *string `json:"name,omitempty"`

	// The type of the asset. Example: `compute.googleapis.com/Disk`
	//
	//  See [Supported asset
	//  types](https://cloud.google.com/asset-inventory/docs/supported-asset-types)
	//  for more information.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.asset_type
	AssetType *string `json:"assetType,omitempty"`

	// A representation of the resource.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.resource
	Resource *Resource `json:"resource,omitempty"`

	// A representation of the IAM policy set on a Google Cloud resource.
	//  There can be a maximum of one IAM policy set on any given resource.
	//  In addition, IAM policies inherit their granted access scope from any
	//  policies set on parent resources in the resource hierarchy. Therefore, the
	//  effectively policy is the union of both the policy set on this resource
	//  and each policy set on all of the resource's ancestry resource levels in
	//  the hierarchy. See
	//  [this topic](https://cloud.google.com/iam/help/allow-policies/inheritance)
	//  for more information.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.iam_policy
	IamPolicy *Policy `json:"iamPolicy,omitempty"`

	// A representation of an [organization
	//  policy](https://cloud.google.com/resource-manager/docs/organization-policy/overview#organization_policy).
	//  There can be more than one organization policy with different constraints
	//  set on a given resource.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.org_policy
	OrgPolicy []Policy `json:"orgPolicy,omitempty"`

	// Please also refer to the [access policy user
	//  guide](https://cloud.google.com/access-context-manager/docs/overview#access-policies).
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.access_policy
	AccessPolicy *AccessPolicy `json:"accessPolicy,omitempty"`

	// Please also refer to the [access level user
	//  guide](https://cloud.google.com/access-context-manager/docs/overview#access-levels).
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.access_level
	AccessLevel *AccessLevel `json:"accessLevel,omitempty"`

	// Please also refer to the [service perimeter user
	//  guide](https://cloud.google.com/vpc-service-controls/docs/overview).
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.service_perimeter
	ServicePerimeter *ServicePerimeter `json:"servicePerimeter,omitempty"`

	// The related assets of the asset of one relationship type.
	//  One asset only represents one type of relationship.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.related_assets
	RelatedAssets *RelatedAssets `json:"relatedAssets,omitempty"`

	// The ancestry path of an asset in Google Cloud [resource
	//  hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy),
	//  represented as a list of relative resource names. An ancestry path starts
	//  with the closest ancestor in the hierarchy and ends at root. If the asset
	//  is a project, folder, or organization, the ancestry path starts from the
	//  asset itself.
	//
	//  Example: `["projects/123456789", "folders/5432", "organizations/1234"]`
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Asset.ancestors
	Ancestors []string `json:"ancestors,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1p7beta1.RelatedAsset
type RelatedAsset struct {
	// The full name of the asset. Example:
	//  `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`
	//
	//  See [Resource
	//  names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	//  for more information.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelatedAsset.asset
	Asset *string `json:"asset,omitempty"`

	// The type of the asset. Example: `compute.googleapis.com/Disk`
	//
	//  See [Supported asset
	//  types](https://cloud.google.com/asset-inventory/docs/supported-asset-types)
	//  for more information.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelatedAsset.asset_type
	AssetType *string `json:"assetType,omitempty"`

	// The ancestors of an asset in Google Cloud [resource
	//  hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy),
	//  represented as a list of relative resource names. An ancestry path starts
	//  with the closest ancestor in the hierarchy and ends at root.
	//
	//  Example: `["projects/123456789", "folders/5432", "organizations/1234"]`
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelatedAsset.ancestors
	Ancestors []string `json:"ancestors,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1p7beta1.RelatedAssets
type RelatedAssets struct {
	// The detailed relation attributes.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelatedAssets.relationship_attributes
	RelationshipAttributes *RelationshipAttributes `json:"relationshipAttributes,omitempty"`

	// The peer resources of the relationship.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelatedAssets.assets
	Assets []RelatedAsset `json:"assets,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1p7beta1.RelationshipAttributes
type RelationshipAttributes struct {
	// The unique identifier of the relationship type. Example:
	//  `INSTANCE_TO_INSTANCEGROUP`
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelationshipAttributes.type
	Type *string `json:"type,omitempty"`

	// The source asset type. Example: `compute.googleapis.com/Instance`
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelationshipAttributes.source_resource_type
	SourceResourceType *string `json:"sourceResourceType,omitempty"`

	// The target asset type. Example: `compute.googleapis.com/Disk`
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelationshipAttributes.target_resource_type
	TargetResourceType *string `json:"targetResourceType,omitempty"`

	// The detail of the relationship, e.g. `contains`, `attaches`
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.RelationshipAttributes.action
	Action *string `json:"action,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1p7beta1.Resource
type Resource struct {
	// The API version. Example: `v1`
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Resource.version
	Version *string `json:"version,omitempty"`

	// The URL of the discovery document containing the resource's JSON schema.
	//  Example:
	//  `https://www.googleapis.com/discovery/v1/apis/compute/v1/rest`
	//
	//  This value is unspecified for resources that do not have an API based on a
	//  discovery document, such as Cloud Bigtable.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Resource.discovery_document_uri
	DiscoveryDocumentURI *string `json:"discoveryDocumentURI,omitempty"`

	// The JSON schema name listed in the discovery document. Example:
	//  `Project`
	//
	//  This value is unspecified for resources that do not have an API based on a
	//  discovery document, such as Cloud Bigtable.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Resource.discovery_name
	DiscoveryName *string `json:"discoveryName,omitempty"`

	// The REST URL for accessing the resource. An HTTP `GET` request using this
	//  URL returns the resource itself. Example:
	//  `https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123`
	//
	//  This value is unspecified for resources without a REST API.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Resource.resource_url
	ResourceURL *string `json:"resourceURL,omitempty"`

	// The full name of the immediate parent of this resource. See
	//  [Resource
	//  Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	//  for more information.
	//
	//  For Google Cloud assets, this value is the parent resource defined in the
	//  [IAM policy
	//  hierarchy](https://cloud.google.com/iam/docs/overview#policy_hierarchy).
	//  Example:
	//  `//cloudresourcemanager.googleapis.com/projects/my_project_123`
	//
	//  For third-party assets, this field may be set differently.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Resource.parent
	Parent *string `json:"parent,omitempty"`

	// The content of the resource, in which some sensitive fields are removed
	//  and may not be present.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Resource.data
	Data map[string]string `json:"data,omitempty"`

	// The location of the resource in Google Cloud, such as its zone and region.
	//  For more information, see https://cloud.google.com/about/locations/.
	// +kcc:proto:field=google.cloud.asset.v1p7beta1.Resource.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v1.Policy
type Policy struct {
	// Version of the `Policy`. Default version is 0;
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.version
	Version *int32 `json:"version,omitempty"`

	// The name of the `Constraint` the `Policy` is configuring, for example,
	//  `constraints/serviceuser.services`.
	//
	//  Immutable after creation.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.constraint
	Constraint *string `json:"constraint,omitempty"`

	// An opaque tag indicating the current version of the `Policy`, used for
	//  concurrency control.
	//
	//  When the `Policy` is returned from either a `GetPolicy` or a
	//  `ListOrgPolicy` request, this `etag` indicates the version of the current
	//  `Policy` to use when executing a read-modify-write loop.
	//
	//  When the `Policy` is returned from a `GetEffectivePolicy` request, the
	//  `etag` will be unset.
	//
	//  When the `Policy` is used in a `SetOrgPolicy` method, use the `etag` value
	//  that was returned from a `GetOrgPolicy` request as part of a
	//  read-modify-write loop for concurrency control. Not setting the `etag`in a
	//  `SetOrgPolicy` request will result in an unconditional write of the
	//  `Policy`.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.etag
	Etag []byte `json:"etag,omitempty"`

	// The time stamp the `Policy` was previously updated. This is set by the
	//  server, not specified by the caller, and represents the last time a call to
	//  `SetOrgPolicy` was made for that `Policy`. Any value set by the client will
	//  be ignored.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// List of values either allowed or disallowed.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.list_policy
	ListPolicy *Policy_ListPolicy `json:"listPolicy,omitempty"`

	// For boolean `Constraints`, whether to enforce the `Constraint` or not.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.boolean_policy
	BooleanPolicy *Policy_BooleanPolicy `json:"booleanPolicy,omitempty"`

	// Restores the default behavior of the constraint; independent of
	//  `Constraint` type.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.restore_default
	RestoreDefault *Policy_RestoreDefault `json:"restoreDefault,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v1.Policy.BooleanPolicy
type Policy_BooleanPolicy struct {
	// If `true`, then the `Policy` is enforced. If `false`, then any
	//  configuration is acceptable.
	//
	//  Suppose you have a `Constraint`
	//  `constraints/compute.disableSerialPortAccess` with `constraint_default`
	//  set to `ALLOW`. A `Policy` for that `Constraint` exhibits the following
	//  behavior:
	//    - If the `Policy` at this resource has enforced set to `false`, serial
	//      port connection attempts will be allowed.
	//    - If the `Policy` at this resource has enforced set to `true`, serial
	//      port connection attempts will be refused.
	//    - If the `Policy` at this resource is `RestoreDefault`, serial port
	//      connection attempts will be allowed.
	//    - If no `Policy` is set at this resource or anywhere higher in the
	//      resource hierarchy, serial port connection attempts will be allowed.
	//    - If no `Policy` is set at this resource, but one exists higher in the
	//      resource hierarchy, the behavior is as if the`Policy` were set at
	//      this resource.
	//
	//  The following examples demonstrate the different possible layerings:
	//
	//  Example 1 (nearest `Constraint` wins):
	//    `organizations/foo` has a `Policy` with:
	//      {enforced: false}
	//    `projects/bar` has no `Policy` set.
	//  The constraint at `projects/bar` and `organizations/foo` will not be
	//  enforced.
	//
	//  Example 2 (enforcement gets replaced):
	//    `organizations/foo` has a `Policy` with:
	//      {enforced: false}
	//    `projects/bar` has a `Policy` with:
	//      {enforced: true}
	//  The constraint at `organizations/foo` is not enforced.
	//  The constraint at `projects/bar` is enforced.
	//
	//  Example 3 (RestoreDefault):
	//    `organizations/foo` has a `Policy` with:
	//      {enforced: true}
	//    `projects/bar` has a `Policy` with:
	//      {RestoreDefault: {}}
	//  The constraint at `organizations/foo` is enforced.
	//  The constraint at `projects/bar` is not enforced, because
	//  `constraint_default` for the `Constraint` is `ALLOW`.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.BooleanPolicy.enforced
	Enforced *bool `json:"enforced,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v1.Policy.ListPolicy
type Policy_ListPolicy struct {
	// List of values allowed  at this resource. Can only be set if `all_values`
	//  is set to `ALL_VALUES_UNSPECIFIED`.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.ListPolicy.allowed_values
	AllowedValues []string `json:"allowedValues,omitempty"`

	// List of values denied at this resource. Can only be set if `all_values`
	//  is set to `ALL_VALUES_UNSPECIFIED`.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.ListPolicy.denied_values
	DeniedValues []string `json:"deniedValues,omitempty"`

	// The policy all_values state.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.ListPolicy.all_values
	AllValues *string `json:"allValues,omitempty"`

	// Optional. The Google Cloud Console will try to default to a configuration
	//  that matches the value specified in this `Policy`. If `suggested_value`
	//  is not set, it will inherit the value specified higher in the hierarchy,
	//  unless `inherit_from_parent` is `false`.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.ListPolicy.suggested_value
	SuggestedValue *string `json:"suggestedValue,omitempty"`

	// Determines the inheritance behavior for this `Policy`.
	//
	//  By default, a `ListPolicy` set at a resource supercedes any `Policy` set
	//  anywhere up the resource hierarchy. However, if `inherit_from_parent` is
	//  set to `true`, then the values from the effective `Policy` of the parent
	//  resource are inherited, meaning the values set in this `Policy` are
	//  added to the values inherited up the hierarchy.
	//
	//  Setting `Policy` hierarchies that inherit both allowed values and denied
	//  values isn't recommended in most circumstances to keep the configuration
	//  simple and understandable. However, it is possible to set a `Policy` with
	//  `allowed_values` set that inherits a `Policy` with `denied_values` set.
	//  In this case, the values that are allowed must be in `allowed_values` and
	//  not present in `denied_values`.
	//
	//  For example, suppose you have a `Constraint`
	//  `constraints/serviceuser.services`, which has a `constraint_type` of
	//  `list_constraint`, and with `constraint_default` set to `ALLOW`.
	//  Suppose that at the Organization level, a `Policy` is applied that
	//  restricts the allowed API activations to {`E1`, `E2`}. Then, if a
	//  `Policy` is applied to a project below the Organization that has
	//  `inherit_from_parent` set to `false` and field all_values set to DENY,
	//  then an attempt to activate any API will be denied.
	//
	//  The following examples demonstrate different possible layerings for
	//  `projects/bar` parented by `organizations/foo`:
	//
	//  Example 1 (no inherited values):
	//    `organizations/foo` has a `Policy` with values:
	//      {allowed_values: "E1" allowed_values:"E2"}
	//    `projects/bar` has `inherit_from_parent` `false` and values:
	//      {allowed_values: "E3" allowed_values: "E4"}
	//  The accepted values at `organizations/foo` are `E1`, `E2`.
	//  The accepted values at `projects/bar` are `E3`, and `E4`.
	//
	//  Example 2 (inherited values):
	//    `organizations/foo` has a `Policy` with values:
	//      {allowed_values: "E1" allowed_values:"E2"}
	//    `projects/bar` has a `Policy` with values:
	//      {value: "E3" value: "E4" inherit_from_parent: true}
	//  The accepted values at `organizations/foo` are `E1`, `E2`.
	//  The accepted values at `projects/bar` are `E1`, `E2`, `E3`, and `E4`.
	//
	//  Example 3 (inheriting both allowed and denied values):
	//    `organizations/foo` has a `Policy` with values:
	//      {allowed_values: "E1" allowed_values: "E2"}
	//    `projects/bar` has a `Policy` with:
	//      {denied_values: "E1"}
	//  The accepted values at `organizations/foo` are `E1`, `E2`.
	//  The value accepted at `projects/bar` is `E2`.
	//
	//  Example 4 (RestoreDefault):
	//    `organizations/foo` has a `Policy` with values:
	//      {allowed_values: "E1" allowed_values:"E2"}
	//    `projects/bar` has a `Policy` with values:
	//      {RestoreDefault: {}}
	//  The accepted values at `organizations/foo` are `E1`, `E2`.
	//  The accepted values at `projects/bar` are either all or none depending on
	//  the value of `constraint_default` (if `ALLOW`, all; if
	//  `DENY`, none).
	//
	//  Example 5 (no policy inherits parent policy):
	//    `organizations/foo` has no `Policy` set.
	//    `projects/bar` has no `Policy` set.
	//  The accepted values at both levels are either all or none depending on
	//  the value of `constraint_default` (if `ALLOW`, all; if
	//  `DENY`, none).
	//
	//  Example 6 (ListConstraint allowing all):
	//    `organizations/foo` has a `Policy` with values:
	//      {allowed_values: "E1" allowed_values: "E2"}
	//    `projects/bar` has a `Policy` with:
	//      {all: ALLOW}
	//  The accepted values at `organizations/foo` are `E1`, E2`.
	//  Any value is accepted at `projects/bar`.
	//
	//  Example 7 (ListConstraint allowing none):
	//    `organizations/foo` has a `Policy` with values:
	//      {allowed_values: "E1" allowed_values: "E2"}
	//    `projects/bar` has a `Policy` with:
	//      {all: DENY}
	//  The accepted values at `organizations/foo` are `E1`, E2`.
	//  No value is accepted at `projects/bar`.
	//
	//  Example 10 (allowed and denied subtrees of Resource Manager hierarchy):
	//  Given the following resource hierarchy
	//    O1->{F1, F2}; F1->{P1}; F2->{P2, P3},
	//    `organizations/foo` has a `Policy` with values:
	//      {allowed_values: "under:organizations/O1"}
	//    `projects/bar` has a `Policy` with:
	//      {allowed_values: "under:projects/P3"}
	//      {denied_values: "under:folders/F2"}
	//  The accepted values at `organizations/foo` are `organizations/O1`,
	//    `folders/F1`, `folders/F2`, `projects/P1`, `projects/P2`,
	//    `projects/P3`.
	//  The accepted values at `projects/bar` are `organizations/O1`,
	//    `folders/F1`, `projects/P1`.
	// +kcc:proto:field=google.cloud.orgpolicy.v1.Policy.ListPolicy.inherit_from_parent
	InheritFromParent *bool `json:"inheritFromParent,omitempty"`
}

// +kcc:proto=google.cloud.orgpolicy.v1.Policy.RestoreDefault
type Policy_RestoreDefault struct {
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

// +kcc:proto=google.identity.accesscontextmanager.v1.AccessLevel
type AccessLevel struct {
	// Required. Resource name for the Access Level. The `short_name` component
	//  must begin with a letter and only include alphanumeric and '_'. Format:
	//  `accessPolicies/{access_policy}/accessLevels/{access_level}`. The maximum
	//  length of the `access_level` component is 50 characters.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.name
	Name *string `json:"name,omitempty"`

	// Human readable title. Must be unique within the Policy.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.title
	Title *string `json:"title,omitempty"`

	// Description of the `AccessLevel` and its use. Does not affect behavior.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.description
	Description *string `json:"description,omitempty"`

	// A `BasicLevel` composed of `Conditions`.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.basic
	Basic *BasicLevel `json:"basic,omitempty"`

	// A `CustomLevel` written in the Common Expression Language.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.custom
	Custom *CustomLevel `json:"custom,omitempty"`

	// Output only. Time the `AccessLevel` was created in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the `AccessLevel` was updated in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.AccessPolicy
type AccessPolicy struct {
	// Output only. Resource name of the `AccessPolicy`. Format:
	//  `accessPolicies/{access_policy}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.name
	Name *string `json:"name,omitempty"`

	// Required. The parent of this `AccessPolicy` in the Cloud Resource
	//  Hierarchy. Currently immutable once created. Format:
	//  `organizations/{organization_id}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.parent
	Parent *string `json:"parent,omitempty"`

	// Required. Human readable title. Does not affect behavior.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.title
	Title *string `json:"title,omitempty"`

	// The scopes of a policy define which resources an ACM policy can restrict,
	//  and where ACM resources can be referenced.
	//  For example, a policy with scopes=["folders/123"] has the following
	//  behavior:
	//  - vpcsc perimeters can only restrict projects within folders/123
	//  - access levels can only be referenced by resources within folders/123.
	//  If empty, there are no limitations on which resources can be restricted by
	//  an ACM policy, and there are no limitations on where ACM resources can be
	//  referenced.
	//  Only one policy can include a given scope (attempting to create a second
	//  policy which includes "folders/123" will result in an error).
	//  Currently, scopes cannot be modified after a policy is created.
	//  Currently, policies can only have a single scope.
	//  Format: list of `folders/{folder_number}` or `projects/{project_number}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.scopes
	Scopes []string `json:"scopes,omitempty"`

	// Output only. Time the `AccessPolicy` was created in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the `AccessPolicy` was updated in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. An opaque identifier for the current version of the
	//  `AccessPolicy`. This will always be a strongly validated etag, meaning that
	//  two Access Polices will be identical if and only if their etags are
	//  identical. Clients should not expect this to be in any specific format.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.BasicLevel
type BasicLevel struct {
	// Required. A list of requirements for the `AccessLevel` to be granted.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.BasicLevel.conditions
	Conditions []Condition `json:"conditions,omitempty"`

	// How the `conditions` list should be combined to determine if a request is
	//  granted this `AccessLevel`. If AND is used, each `Condition` in
	//  `conditions` must be satisfied for the `AccessLevel` to be applied. If OR
	//  is used, at least one `Condition` in `conditions` must be satisfied for the
	//  `AccessLevel` to be applied. Default behavior is AND.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.BasicLevel.combining_function
	CombiningFunction *string `json:"combiningFunction,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.Condition
type Condition struct {
	// CIDR block IP subnetwork specification. May be IPv4 or IPv6. Note that for
	//  a CIDR IP address block, the specified IP address portion must be properly
	//  truncated (i.e. all the host bits must be zero) or the input is considered
	//  malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24" is
	//  not. Similarly, for IPv6, "2001:db8::/32" is accepted whereas
	//  "2001:db8::1/32" is not. The originating IP of a request must be in one of
	//  the listed subnets in order for this Condition to be true. If empty, all IP
	//  addresses are allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.ip_subnetworks
	IPSubnetworks []string `json:"ipSubnetworks,omitempty"`

	// Device specific restrictions, all restrictions must hold for the
	//  Condition to be true. If not specified, all devices are allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.device_policy
	DevicePolicy *DevicePolicy `json:"devicePolicy,omitempty"`

	// A list of other access levels defined in the same `Policy`, referenced by
	//  resource name. Referencing an `AccessLevel` which does not exist is an
	//  error. All access levels listed must be granted for the Condition
	//  to be true. Example:
	//  "`accessPolicies/MY_POLICY/accessLevels/LEVEL_NAME"`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.required_access_levels
	RequiredAccessLevels []string `json:"requiredAccessLevels,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes a NAND over
	//  its non-empty fields, each field must be false for the Condition overall to
	//  be satisfied. Defaults to false.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.negate
	Negate *bool `json:"negate,omitempty"`

	// The request must be made by one of the provided user or service
	//  accounts. Groups are not supported.
	//  Syntax:
	//  `user:{emailid}`
	//  `serviceAccount:{emailid}`
	//  If not specified, a request may come from any user.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.members
	Members []string `json:"members,omitempty"`

	// The request must originate from one of the provided countries/regions.
	//  Must be valid ISO 3166-1 alpha-2 codes.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.regions
	Regions []string `json:"regions,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.CustomLevel
type CustomLevel struct {
	// Required. A Cloud CEL expression evaluating to a boolean.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.CustomLevel.expr
	Expr *Expr `json:"expr,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.DevicePolicy
type DevicePolicy struct {
	// Whether or not screenlock is required for the DevicePolicy to be true.
	//  Defaults to `false`.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.DevicePolicy.require_screenlock
	RequireScreenlock *bool `json:"requireScreenlock,omitempty"`

	// Allowed encryptions statuses, an empty list allows all statuses.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.DevicePolicy.allowed_encryption_statuses
	AllowedEncryptionStatuses []string `json:"allowedEncryptionStatuses,omitempty"`

	// Allowed OS versions, an empty list allows all types and all versions.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.DevicePolicy.os_constraints
	OsConstraints []OsConstraint `json:"osConstraints,omitempty"`

	// Allowed device management levels, an empty list allows all management
	//  levels.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.DevicePolicy.allowed_device_management_levels
	AllowedDeviceManagementLevels []string `json:"allowedDeviceManagementLevels,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.DevicePolicy.require_admin_approval
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty"`

	// Whether the device needs to be corp owned.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.DevicePolicy.require_corp_owned
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.OsConstraint
type OsConstraint struct {
	// Required. The allowed OS type.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.OsConstraint.os_type
	OsType *string `json:"osType,omitempty"`

	// The minimum allowed OS version. If not set, any version of this OS
	//  satisfies the constraint. Format: `"major.minor.patch"`.
	//  Examples: `"10.5.301"`, `"9.2.1"`.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.OsConstraint.minimum_version
	MinimumVersion *string `json:"minimumVersion,omitempty"`

	// Only allows requests from devices with a verified Chrome OS.
	//  Verifications includes requirements that the device is enterprise-managed,
	//  conformant to domain policies, and the caller has permission to call
	//  the API targeted by the request.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.OsConstraint.require_verified_chrome_os
	RequireVerifiedChromeOs *bool `json:"requireVerifiedChromeOs,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeter
type ServicePerimeter struct {
	// Required. Resource name for the ServicePerimeter.  The `short_name`
	//  component must begin with a letter and only include alphanumeric and '_'.
	//  Format:
	//  `accessPolicies/{access_policy}/servicePerimeters/{service_perimeter}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.name
	Name *string `json:"name,omitempty"`

	// Human readable title. Must be unique within the Policy.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.title
	Title *string `json:"title,omitempty"`

	// Description of the `ServicePerimeter` and its use. Does not affect
	//  behavior.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.description
	Description *string `json:"description,omitempty"`

	// Output only. Time the `ServicePerimeter` was created in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the `ServicePerimeter` was updated in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Perimeter type indicator. A single project is
	//  allowed to be a member of single regular perimeter, but multiple service
	//  perimeter bridges. A project cannot be a included in a perimeter bridge
	//  without being included in regular perimeter. For perimeter bridges,
	//  the restricted service list as well as access level lists must be
	//  empty.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.perimeter_type
	PerimeterType *string `json:"perimeterType,omitempty"`

	// Current ServicePerimeter configuration. Specifies sets of resources,
	//  restricted services and access levels that determine perimeter
	//  content and boundaries.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.status
	Status *ServicePerimeterConfig `json:"status,omitempty"`

	// Proposed (or dry run) ServicePerimeter configuration. This configuration
	//  allows to specify and test ServicePerimeter configuration without enforcing
	//  actual access restrictions. Only allowed to be set when the
	//  "use_explicit_dry_run_spec" flag is set.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.spec
	Spec *ServicePerimeterConfig `json:"spec,omitempty"`

	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly
	//  exists  for all Service Perimeters, and that spec is identical to the
	//  status for those Service Perimeters. When this flag is set, it inhibits the
	//  generation of the implicit spec, thereby allowing the user to explicitly
	//  provide a configuration ("spec") to use in a dry-run version of the Service
	//  Perimeter. This allows the user to test changes to the enforced config
	//  ("status") without actually enforcing them. This testing is done through
	//  analyzing the differences between currently enforced and suggested
	//  restrictions. use_explicit_dry_run_spec must bet set to True if any of the
	//  fields in the spec are set to non-default values.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.use_explicit_dry_run_spec
	UseExplicitDryRunSpec *bool `json:"useExplicitDryRunSpec,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig
type ServicePerimeterConfig struct {
	// A list of Google Cloud resources that are inside of the service perimeter.
	//  Currently only projects are allowed. Format: `projects/{project_number}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.resources
	Resources []string `json:"resources,omitempty"`

	// A list of `AccessLevel` resource names that allow resources within the
	//  `ServicePerimeter` to be accessed from the internet. `AccessLevels` listed
	//  must be in the same policy as this `ServicePerimeter`. Referencing a
	//  nonexistent `AccessLevel` is a syntax error. If no `AccessLevel` names are
	//  listed, resources within the perimeter can only be accessed via Google
	//  Cloud calls with request origins within the perimeter. Example:
	//  `"accessPolicies/MY_POLICY/accessLevels/MY_LEVEL"`.
	//  For Service Perimeter Bridge, must be empty.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.access_levels
	AccessLevels []string `json:"accessLevels,omitempty"`

	// Google Cloud services that are subject to the Service Perimeter
	//  restrictions. For example, if `storage.googleapis.com` is specified, access
	//  to the storage buckets inside the perimeter must meet the perimeter's
	//  access restrictions.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.restricted_services
	RestrictedServices []string `json:"restrictedServices,omitempty"`

	// Configuration for APIs allowed within Perimeter.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.vpc_accessible_services
	VpcAccessibleServices *ServicePerimeterConfig_VpcAccessibleServices `json:"vpcAccessibleServices,omitempty"`

	// List of [IngressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  to apply to the perimeter. A perimeter may have multiple [IngressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy],
	//  each of which is evaluated separately. Access is granted if any [Ingress
	//  Policy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  grants it. Must be empty for a perimeter bridge.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ingress_policies
	IngressPolicies []ServicePerimeterConfig_IngressPolicy `json:"ingressPolicies,omitempty"`

	// List of [EgressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  to apply to the perimeter. A perimeter may have multiple [EgressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy],
	//  each of which is evaluated separately. Access is granted if any
	//  [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  grants it. Must be empty for a perimeter bridge.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.egress_policies
	EgressPolicies []ServicePerimeterConfig_EgressPolicy `json:"egressPolicies,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation
type ServicePerimeterConfig_ApiOperation struct {
	// The name of the API whose methods or permissions the [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  or [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  want to allow. A single [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  with `service_name` field set to `*` will allow all methods AND
	//  permissions for all services.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// API methods or permissions to allow. Method or permission must belong to
	//  the service specified by `service_name` field. A single [MethodSelector]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector]
	//  entry with `*` specified for the `method` field will allow all methods
	//  AND permissions for the service specified in `service_name`.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation.method_selectors
	MethodSelectors []ServicePerimeterConfig_MethodSelector `json:"methodSelectors,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom
type ServicePerimeterConfig_EgressFrom struct {
	// A list of identities that are allowed access through this [EgressPolicy].
	//  Should be in the format of email address. The email address should
	//  represent individual user or service account only.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom.identities
	Identities []string `json:"identities,omitempty"`

	// Specifies the type of identities that are allowed access to outside the
	//  perimeter. If left unspecified, then members of `identities` field will
	//  be allowed access.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom.identity_type
	IdentityType *string `json:"identityType,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy
type ServicePerimeterConfig_EgressPolicy struct {
	// Defines conditions on the source of a request causing this [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy.egress_from
	EgressFrom *ServicePerimeterConfig_EgressFrom `json:"egressFrom,omitempty"`

	// Defines the conditions on the [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  and destination resources that cause this [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy.egress_to
	EgressTo *ServicePerimeterConfig_EgressTo `json:"egressTo,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo
type ServicePerimeterConfig_EgressTo struct {
	// A list of resources, currently only projects in the form
	//  `projects/<projectnumber>`, that are allowed to be accessed by sources
	//  defined in the corresponding [EgressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom].
	//  A request matches if it contains a resource in this list.  If `*` is
	//  specified for `resources`, then this [EgressTo]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo]
	//  rule will authorize access to all resources outside the perimeter.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo.resources
	Resources []string `json:"resources,omitempty"`

	// A list of [ApiOperations]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  allowed to be performed by the sources specified in the corresponding
	//  [EgressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom].
	//  A request matches if it uses an operation/service in this list.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo.operations
	Operations []ServicePerimeterConfig_ApiOperation `json:"operations,omitempty"`

	// A list of external resources that are allowed to be accessed. Only AWS
	//  and Azure resources are supported. For Amazon S3, the supported format is
	//  s3://BUCKET_NAME. For Azure Storage, the supported format is
	//  azure://myaccount.blob.core.windows.net/CONTAINER_NAME. A request matches
	//  if it contains an external resource in this list (Example:
	//  s3://bucket/path). Currently '*' is not allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo.external_resources
	ExternalResources []string `json:"externalResources,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom
type ServicePerimeterConfig_IngressFrom struct {
	// Sources that this [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  authorizes access from.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom.sources
	Sources []ServicePerimeterConfig_IngressSource `json:"sources,omitempty"`

	// A list of identities that are allowed access through this ingress
	//  policy. Should be in the format of email address. The email address
	//  should represent individual user or service account only.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom.identities
	Identities []string `json:"identities,omitempty"`

	// Specifies the type of identities that are allowed access from outside the
	//  perimeter. If left unspecified, then members of `identities` field will
	//  be allowed access.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom.identity_type
	IdentityType *string `json:"identityType,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy
type ServicePerimeterConfig_IngressPolicy struct {
	// Defines the conditions on the source of a request causing this
	//  [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy.ingress_from
	IngressFrom *ServicePerimeterConfig_IngressFrom `json:"ingressFrom,omitempty"`

	// Defines the conditions on the [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  and request destination that cause this [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy.ingress_to
	IngressTo *ServicePerimeterConfig_IngressTo `json:"ingressTo,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource
type ServicePerimeterConfig_IngressSource struct {
	// An [AccessLevel]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] resource
	//  name that allow resources within the [ServicePerimeters]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter] to be
	//  accessed from the internet. [AccessLevels]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] listed must
	//  be in the same policy as this [ServicePerimeter]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter].
	//  Referencing a nonexistent [AccessLevel]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] will cause
	//  an error. If no [AccessLevel]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] names are
	//  listed, resources within the perimeter can only be accessed via Google
	//  Cloud calls with request origins within the perimeter. Example:
	//  `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL`. If a single `*` is
	//  specified for `access_level`, then all [IngressSources]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource]
	//  will be allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource.access_level
	AccessLevel *string `json:"accessLevel,omitempty"`

	// A Google Cloud resource that is allowed to ingress the perimeter.
	//  Requests from these resources will be allowed to access perimeter data.
	//  Currently only projects are allowed.
	//  Format: `projects/{project_number}`
	//  The project may be in any Google Cloud organization, not just the
	//  organization that the perimeter is defined in. `*` is not allowed, the
	//  case of allowing all Google Cloud resources only is not supported.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource.resource
	Resource *string `json:"resource,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressTo
type ServicePerimeterConfig_IngressTo struct {
	// A list of [ApiOperations]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  allowed to be performed by the sources specified in corresponding
	//  [IngressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom]
	//  in this [ServicePerimeter]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter].
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressTo.operations
	Operations []ServicePerimeterConfig_ApiOperation `json:"operations,omitempty"`

	// A list of resources, currently only projects in the form
	//  `projects/<projectnumber>`, protected by this [ServicePerimeter]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter] that are
	//  allowed to be accessed by sources defined in the corresponding
	//  [IngressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom].
	//  If a single `*` is specified, then access to all resources inside the
	//  perimeter are allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressTo.resources
	Resources []string `json:"resources,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector
type ServicePerimeterConfig_MethodSelector struct {
	// Value for `method` should be a valid method name for the corresponding
	//  `service_name` in [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation].
	//  If `*` used as value for `method`, then ALL methods and permissions are
	//  allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector.method
	Method *string `json:"method,omitempty"`

	// Value for `permission` should be a valid Cloud IAM permission for the
	//  corresponding `service_name` in [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation].
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector.permission
	Permission *string `json:"permission,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices
type ServicePerimeterConfig_VpcAccessibleServices struct {
	// Whether to restrict API calls within the Service Perimeter to the list of
	//  APIs specified in 'allowed_services'.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices.enable_restriction
	EnableRestriction *bool `json:"enableRestriction,omitempty"`

	// The list of APIs usable within the Service Perimeter. Must be empty
	//  unless 'enable_restriction' is True. You can specify a list of individual
	//  services, as well as include the 'RESTRICTED-SERVICES' value, which
	//  automatically includes all of the services protected by the perimeter.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices.allowed_services
	AllowedServices []string `json:"allowedServices,omitempty"`
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
