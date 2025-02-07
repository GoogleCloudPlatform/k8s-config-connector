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

// +kcc:proto=google.cloud.apihub.v1.Version
type Version struct {
	// Identifier. The name of the version.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the version.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the version.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.description
	Description *string `json:"description,omitempty"`

	// Optional. The documentation of the version.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. The deployments linked to this API version.
	//  Note: A particular API version could be deployed to multiple deployments
	//  (for dev deployment, UAT deployment, etc)
	//  Format is
	//  `projects/{project}/locations/{location}/deployments/{deployment}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.deployments
	Deployments []string `json:"deployments,omitempty"`

	// Optional. The lifecycle of the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-lifecycle`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.lifecycle
	Lifecycle *AttributeValues `json:"lifecycle,omitempty"`

	// Optional. The compliance associated with the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-compliance`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.compliance
	Compliance *AttributeValues `json:"compliance,omitempty"`

	// Optional. The accreditations associated with the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-accreditation`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.accreditation
	Accreditation *AttributeValues `json:"accreditation,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. The selected deployment for a Version resource.
	//  This can be used when special handling is needed on client side for a
	//  particular deployment linked to the version.
	//  Format is
	//  `projects/{project}/locations/{location}/deployments/{deployment}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.selected_deployment
	SelectedDeployment *string `json:"selectedDeployment,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValuesObservedState struct {
	// Output only. The name of the attribute.
	//  Format: projects/{project}/locations/{location}/attributes/{attribute}
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.attribute
	Attribute *string `json:"attribute,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Version
type VersionObservedState struct {
	// Output only. The specs associated with this version.
	//  Note that an API version can be associated with multiple specs.
	//  Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.specs
	Specs []string `json:"specs,omitempty"`

	// Output only. The operations contained in the API version.
	//  These operations will be added to the version when a new spec is
	//  added or when an existing spec is updated. Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/operations/{operation}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.api_operations
	ApiOperations []string `json:"apiOperations,omitempty"`

	// Output only. The definitions contained in the API version.
	//  These definitions will be added to the version when a new spec is
	//  added or when an existing spec is updated. Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/definitions/{definition}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.definitions
	Definitions []string `json:"definitions,omitempty"`

	// Output only. The time at which the version was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the version was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The lifecycle of the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-lifecycle`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.lifecycle
	Lifecycle *AttributeValuesObservedState `json:"lifecycle,omitempty"`
}
