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
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1
// resource: DataCatalogTag:Tag

package v1alpha1

// +kcc:proto=google.cloud.datacatalog.v1.Tag
type Tag struct {
	// Identifier. The resource name of the tag in URL format where tag ID is a
	//  system-generated identifier.
	//
	//  Note: The tag itself might not be stored in the location specified in its
	//  name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.name
	Name *string `json:"name,omitempty"`

	// Required. The resource name of the tag template this tag uses. Example:
	//
	//  `projects/{PROJECT_ID}/locations/{LOCATION}/tagTemplates/{TAG_TEMPLATE_ID}`
	//
	//  This field cannot be modified after creation.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.template
	Template *string `json:"template,omitempty"`

	// Resources like entry can have schemas associated with them. This scope
	//  allows you to attach tags to an individual column based on that schema.
	//
	//  To attach a tag to a nested column, separate column names with a dot
	//  (`.`). Example: `column.nested_column`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.column
	Column *string `json:"column,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.datacatalog.v1.TagField
type TagField struct {

	// The value of a tag field with a double type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagField.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// The value of a tag field with a string type.
	//
	//  The maximum length is 2000 UTF-8 characters.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagField.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// The value of a tag field with a boolean type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagField.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// The value of a tag field with a timestamp type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagField.timestamp_value
	TimestampValue *string `json:"timestampValue,omitempty"`

	// The value of a tag field with an enum type.
	//
	//  This value must be one of the allowed values listed in this enum.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagField.enum_value
	EnumValue *TagField_EnumValue `json:"enumValue,omitempty"`

	// The value of a tag field with a rich text type.
	//
	//  The maximum length is 10 MiB as this value holds HTML descriptions
	//  including encoded images. The maximum length of the text without images
	//  is 100 KiB.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagField.richtext_value
	RichtextValue *string `json:"richtextValue,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.TagField.EnumValue
type TagField_EnumValue struct {
	// The display name of the enum value.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagField.EnumValue.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Tag
type TagObservedState struct {
	// Output only. The display name of the tag template.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.template_display_name
	TemplateDisplayName *string `json:"templateDisplayName,omitempty"`

	// Output only. Denotes the transfer status of the Tag Template.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.dataplex_transfer_status
	DataplexTransferStatus *string `json:"dataplexTransferStatus,omitempty"`
}
