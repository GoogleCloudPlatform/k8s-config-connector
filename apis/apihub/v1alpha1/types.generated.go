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

// +kcc:proto=google.cloud.apihub.v1.Api
type Api struct {
	// Identifier. The name of the API resource in the API Hub.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/apis/{api}`
	// +kcc:proto:field=google.cloud.apihub.v1.Api.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.description
	Description *string `json:"description,omitempty"`

	// Optional. The documentation for the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. Owner details for the API resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.owner
	Owner *Owner `json:"owner,omitempty"`

	// Optional. The target users for the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-target-user`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.target_user
	TargetUser *AttributeValues `json:"targetUser,omitempty"`

	// Optional. The team owning the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-team`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.team
	Team *AttributeValues `json:"team,omitempty"`

	// Optional. The business unit owning the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-business-unit`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.business_unit
	BusinessUnit *AttributeValues `json:"businessUnit,omitempty"`

	// Optional. The maturity level of the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-maturity-level`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.maturity_level
	MaturityLevel *AttributeValues `json:"maturityLevel,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. The style of the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-api-style`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.api_style
	ApiStyle *AttributeValues `json:"apiStyle,omitempty"`

	// Optional. The selected version for an API resource.
	//  This can be used when special handling is needed on client side for
	//  particular version of the API. Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}`
	// +kcc:proto:field=google.cloud.apihub.v1.Api.selected_version
	SelectedVersion *string `json:"selectedVersion,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Attribute.AllowedValue
type Attribute_AllowedValue struct {
	// Required. The ID of the allowed value.
	//  * If provided, the same will be used. The service will throw an error if
	//  the specified id is already used by another allowed value in the same
	//  attribute resource.
	//  * If not provided, a system generated id derived from the display name
	//  will be used. In this case, the service will handle conflict resolution
	//  by adding a system generated suffix in case of duplicates.
	//
	//  This value should be 4-63 characters, and valid characters
	//  are /[a-z][0-9]-/.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.id
	ID *string `json:"id,omitempty"`

	// Required. The display name of the allowed value.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The detailed description of the allowed value.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.description
	Description *string `json:"description,omitempty"`

	// Optional. When set to true, the allowed value cannot be updated or
	//  deleted by the user. It can only be true for System defined attributes.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.immutable
	Immutable *bool `json:"immutable,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValues struct {
	// The attribute values associated with a resource in case attribute data
	//  type is enum.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.enum_values
	EnumValues *AttributeValues_EnumAttributeValues `json:"enumValues,omitempty"`

	// The attribute values associated with a resource in case attribute data
	//  type is string.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.string_values
	StringValues *AttributeValues_StringAttributeValues `json:"stringValues,omitempty"`

	// The attribute values associated with a resource in case attribute data
	//  type is JSON.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.json_values
	JsonValues *AttributeValues_StringAttributeValues `json:"jsonValues,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues.EnumAttributeValues
type AttributeValues_EnumAttributeValues struct {
	// Required. The attribute values in case attribute data type is enum.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.EnumAttributeValues.values
	Values []Attribute_AllowedValue `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues.StringAttributeValues
type AttributeValues_StringAttributeValues struct {
	// Required. The attribute values in case attribute data type is string or
	//  JSON.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.StringAttributeValues.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Documentation
type Documentation struct {
	// Optional. The uri of the externally hosted documentation.
	// +kcc:proto:field=google.cloud.apihub.v1.Documentation.external_uri
	ExternalURI *string `json:"externalURI,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Owner
type Owner struct {
	// Optional. The name of the owner.
	// +kcc:proto:field=google.cloud.apihub.v1.Owner.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The email of the owner.
	// +kcc:proto:field=google.cloud.apihub.v1.Owner.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Api
type ApiObservedState struct {
	// Output only. The list of versions present in an API resource.
	//  Note: An API resource can be associated with more than 1 version.
	//  Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}`
	// +kcc:proto:field=google.cloud.apihub.v1.Api.versions
	Versions []string `json:"versions,omitempty"`

	// Output only. The time at which the API resource was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the API resource was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The target users for the API.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-target-user`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Api.target_user
	TargetUser *AttributeValuesObservedState `json:"targetUser,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValuesObservedState struct {
	// Output only. The name of the attribute.
	//  Format: projects/{project}/locations/{location}/attributes/{attribute}
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.attribute
	Attribute *string `json:"attribute,omitempty"`
}
