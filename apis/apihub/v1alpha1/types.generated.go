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

// +kcc:proto=google.cloud.apihub.v1.Issue
type Issue struct {
	// Required. Rule code unique to each rule defined in linter.
	// +kcc:proto:field=google.cloud.apihub.v1.Issue.code
	Code *string `json:"code,omitempty"`

	// Required. An array of strings indicating the location in the analyzed
	//  document where the rule was triggered.
	// +kcc:proto:field=google.cloud.apihub.v1.Issue.path
	Path []string `json:"path,omitempty"`

	// Required. Human-readable message describing the issue found by the linter.
	// +kcc:proto:field=google.cloud.apihub.v1.Issue.message
	Message *string `json:"message,omitempty"`

	// Required. Severity level of the rule violation.
	// +kcc:proto:field=google.cloud.apihub.v1.Issue.severity
	Severity *string `json:"severity,omitempty"`

	// Required. Object describing where in the file the issue was found.
	// +kcc:proto:field=google.cloud.apihub.v1.Issue.range
	Range *Range `json:"range,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.LintResponse
type LintResponse struct {
	// Optional. Array of issues found in the analyzed document.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.issues
	Issues []Issue `json:"issues,omitempty"`

	// Optional. Summary of all issue types and counts for each severity level.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.summary
	Summary []LintResponse_SummaryEntry `json:"summary,omitempty"`

	// Required. Lint state represents success or failure for linting.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.state
	State *string `json:"state,omitempty"`

	// Required. Name of the linting application.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.source
	Source *string `json:"source,omitempty"`

	// Required. Name of the linter used.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.linter
	Linter *string `json:"linter,omitempty"`

	// Required. Timestamp when the linting response was generated.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.LintResponse.SummaryEntry
type LintResponse_SummaryEntry struct {
	// Required. Severity of the issue.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.SummaryEntry.severity
	Severity *string `json:"severity,omitempty"`

	// Required. Count of issues with the given severity.
	// +kcc:proto:field=google.cloud.apihub.v1.LintResponse.SummaryEntry.count
	Count *int32 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.OpenApiSpecDetails
type OpenApiSpecDetails struct {
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

// +kcc:proto=google.cloud.apihub.v1.Point
type Point struct {
	// Required. Line number (zero-indexed).
	// +kcc:proto:field=google.cloud.apihub.v1.Point.line
	Line *int32 `json:"line,omitempty"`

	// Required. Character position within the line (zero-indexed).
	// +kcc:proto:field=google.cloud.apihub.v1.Point.character
	Character *int32 `json:"character,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Range
type Range struct {
	// Required. Start of the issue.
	// +kcc:proto:field=google.cloud.apihub.v1.Range.start
	Start *Point `json:"start,omitempty"`

	// Required. End of the issue.
	// +kcc:proto:field=google.cloud.apihub.v1.Range.end
	End *Point `json:"end,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Spec
type Spec struct {
	// Identifier. The name of the spec.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}`
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the spec.
	//  This can contain the file name of the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The type of spec.
	//  The value should be one of the allowed values defined for
	//  `projects/{project}/locations/{location}/attributes/system-spec-type`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API.
	//
	//  Note, this field is mandatory if content is provided.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.spec_type
	SpecType *AttributeValues `json:"specType,omitempty"`

	// Optional. Input only. The contents of the uploaded spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.contents
	Contents *SpecContents `json:"contents,omitempty"`

	// Optional. The URI of the spec source in case file is uploaded
	//  from an external version control system.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.source_uri
	SourceURI *string `json:"sourceURI,omitempty"`

	// Optional. The lint response for the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.lint_response
	LintResponse *LintResponse `json:"lintResponse,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. The documentation of the spec.
	//  For OpenAPI spec, this will be populated from `externalDocs` in OpenAPI
	//  spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. Input only. Enum specifying the parsing mode for OpenAPI
	//  Specification (OAS) parsing.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.parsing_mode
	ParsingMode *string `json:"parsingMode,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.SpecContents
type SpecContents struct {
	// Required. The contents of the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.SpecContents.contents
	Contents []byte `json:"contents,omitempty"`

	// Required. The mime type of the content for example application/json,
	//  application/yaml, application/wsdl etc.
	// +kcc:proto:field=google.cloud.apihub.v1.SpecContents.mime_type
	MimeType *string `json:"mimeType,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.SpecDetails
type SpecDetails struct {
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValuesObservedState struct {
	// Output only. The name of the attribute.
	//  Format: projects/{project}/locations/{location}/attributes/{attribute}
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.attribute
	Attribute *string `json:"attribute,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.OpenApiSpecDetails
type OpenApiSpecDetailsObservedState struct {
	// Output only. The format of the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.OpenApiSpecDetails.format
	Format *string `json:"format,omitempty"`

	// Output only. The version in the spec.
	//  This maps to `info.version` in OpenAPI spec.
	// +kcc:proto:field=google.cloud.apihub.v1.OpenApiSpecDetails.version
	Version *string `json:"version,omitempty"`

	// Output only. Owner details for the spec.
	//  This maps to `info.contact` in OpenAPI spec.
	// +kcc:proto:field=google.cloud.apihub.v1.OpenApiSpecDetails.owner
	Owner *Owner `json:"owner,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Spec
type SpecObservedState struct {
	// Required. The type of spec.
	//  The value should be one of the allowed values defined for
	//  `projects/{project}/locations/{location}/attributes/system-spec-type`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API.
	//
	//  Note, this field is mandatory if content is provided.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.spec_type
	SpecType *AttributeValuesObservedState `json:"specType,omitempty"`

	// Output only. Details parsed from the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.details
	Details *SpecDetails `json:"details,omitempty"`

	// Output only. The time at which the spec was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the spec was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.SpecDetails
type SpecDetailsObservedState struct {
	// Output only. Additional details apart from `OperationDetails` parsed from
	//  an OpenAPI spec. The OperationDetails parsed from the spec can be
	//  obtained by using
	//  [ListAPIOperations][google.cloud.apihub.v1.ApiHub.ListApiOperations]
	//  method.
	// +kcc:proto:field=google.cloud.apihub.v1.SpecDetails.open_api_spec_details
	OpenApiSpecDetails *OpenApiSpecDetails `json:"openApiSpecDetails,omitempty"`

	// Output only. The description of the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.SpecDetails.description
	Description *string `json:"description,omitempty"`
}
