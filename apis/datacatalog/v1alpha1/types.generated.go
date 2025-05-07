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
