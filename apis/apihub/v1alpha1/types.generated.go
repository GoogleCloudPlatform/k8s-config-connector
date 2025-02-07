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

// +kcc:proto=google.cloud.apihub.v1.Deployment
type Deployment struct {
	// Identifier. The name of the deployment.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/deployments/{deployment}`
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the deployment.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the deployment.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.description
	Description *string `json:"description,omitempty"`

	// Optional. The documentation of the deployment.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Required. The type of deployment.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-deployment-type`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.deployment_type
	DeploymentType *AttributeValues `json:"deploymentType,omitempty"`

	// Required. A URI to the runtime resource. This URI can be used to manage the
	//  resource. For example, if the runtime resource is of type APIGEE_PROXY,
	//  then this field will contain the URI to the management UI of the proxy.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// Required. The endpoints at which this deployment resource is listening for
	//  API requests. This could be a list of complete URIs, hostnames or an IP
	//  addresses.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.endpoints
	Endpoints []string `json:"endpoints,omitempty"`

	// Optional. The SLO for this deployment.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-slo`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.slo
	Slo *AttributeValues `json:"slo,omitempty"`

	// Optional. The environment mapping to this deployment.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-environment`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.environment
	Environment *AttributeValues `json:"environment,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.apihub.v1.Documentation
type Documentation struct {
	// Optional. The uri of the externally hosted documentation.
	// +kcc:proto:field=google.cloud.apihub.v1.Documentation.external_uri
	ExternalURI *string `json:"externalURI,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValuesObservedState struct {
	// Output only. The name of the attribute.
	//  Format: projects/{project}/locations/{location}/attributes/{attribute}
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.attribute
	Attribute *string `json:"attribute,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Deployment
type DeploymentObservedState struct {
	// Required. The type of deployment.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-deployment-type`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.deployment_type
	DeploymentType *AttributeValuesObservedState `json:"deploymentType,omitempty"`

	// Output only. The API versions linked to this deployment.
	//  Note: A particular deployment could be linked to multiple different API
	//  versions (of same or different APIs).
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.api_versions
	ApiVersions []string `json:"apiVersions,omitempty"`

	// Output only. The time at which the deployment was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the deployment was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Deployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
