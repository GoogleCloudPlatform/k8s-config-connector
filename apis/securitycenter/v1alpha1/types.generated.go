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


// +kcc:proto=google.cloud.securitycenter.v1beta1.Asset
type Asset struct {
	// The relative resource name of this asset. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  Example:
	//  "organizations/{organization_id}/assets/{asset_id}".
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.name
	Name *string `json:"name,omitempty"`

	// Security Command Center managed properties. These properties are managed by
	//  Security Command Center and cannot be modified by the user.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.security_center_properties
	SecurityCenterProperties *Asset_SecurityCenterProperties `json:"securityCenterProperties,omitempty"`

	// TODO: unsupported map type with key string and value message


	// User specified security marks. These marks are entirely managed by the user
	//  and come from the SecurityMarks resource that belongs to the asset.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.security_marks
	SecurityMarks *SecurityMarks `json:"securityMarks,omitempty"`

	// The time at which the asset was created in Security Command Center.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The time at which the asset was last updated, added, or deleted in Security
	//  Command Center.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1beta1.Asset.SecurityCenterProperties
type Asset_SecurityCenterProperties struct {
	// Immutable. The full resource name of the Google Cloud resource this asset
	//  represents. This field is immutable after create time. See:
	//  https://cloud.google.com/apis/design/resource_names#full_resource_name
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.SecurityCenterProperties.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// The type of the Google Cloud resource. Examples include: APPLICATION,
	//  PROJECT, and ORGANIZATION. This is a case insensitive field defined by
	//  Security Command Center and/or the producer of the resource and is
	//  immutable after create time.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.SecurityCenterProperties.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// The full resource name of the immediate parent of the resource. See:
	//  https://cloud.google.com/apis/design/resource_names#full_resource_name
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.SecurityCenterProperties.resource_parent
	ResourceParent *string `json:"resourceParent,omitempty"`

	// The full resource name of the project the resource belongs to. See:
	//  https://cloud.google.com/apis/design/resource_names#full_resource_name
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.SecurityCenterProperties.resource_project
	ResourceProject *string `json:"resourceProject,omitempty"`

	// Owners of the Google Cloud resource.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Asset.SecurityCenterProperties.resource_owners
	ResourceOwners []string `json:"resourceOwners,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1beta1.SecurityMarks
type SecurityMarks struct {
	// The relative resource name of the SecurityMarks. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  Examples:
	//  "organizations/{organization_id}/assets/{asset_id}/securityMarks"
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}/securityMarks".
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.SecurityMarks.name
	Name *string `json:"name,omitempty"`

	// Mutable user specified security marks belonging to the parent resource.
	//  Constraints are as follows:
	//
	//    * Keys and values are treated as case insensitive
	//    * Keys must be between 1 - 256 characters (inclusive)
	//    * Keys must be letters, numbers, underscores, or dashes
	//    * Values have leading and trailing whitespace trimmed, remaining
	//      characters must be between 1 - 4096 characters (inclusive)
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.SecurityMarks.marks
	Marks map[string]string `json:"marks,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}
