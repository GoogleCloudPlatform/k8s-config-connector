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
