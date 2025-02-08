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


// +kcc:proto=google.cloud.contentwarehouse.v1.DateTimeTypeOptions
type DateTimeTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DocumentSchema
type DocumentSchema struct {
	// The resource name of the document schema.
	//  Format:
	//  projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}.
	//
	//  The name is ignored when creating a document schema.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.name
	Name *string `json:"name,omitempty"`

	// Required. Name of the schema given by the user. Must be unique per project.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Document details.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.property_definitions
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`

	// Document Type, true refers the document is a folder, otherwise it is
	//  a typical document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.document_is_folder
	DocumentIsFolder *bool `json:"documentIsFolder,omitempty"`

	// Schema description.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.EnumTypeOptions
type EnumTypeOptions struct {
	// Required. List of possible enum values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.EnumTypeOptions.possible_values
	PossibleValues []string `json:"possibleValues,omitempty"`

	// Make sure the Enum property value provided in the document is in the
	//  possile value list during document creation. The validation check runs by
	//  default.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.EnumTypeOptions.validation_check_disabled
	ValidationCheckDisabled *bool `json:"validationCheckDisabled,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.FloatTypeOptions
type FloatTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.IntegerTypeOptions
type IntegerTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.MapTypeOptions
type MapTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyDefinition
type PropertyDefinition struct {
	// Required. The name of the metadata property.
	//  Must be unique within a document schema and is case insensitive.
	//  Names must be non-blank, start with a letter, and can contain alphanumeric
	//  characters and: /, :, -, _, and .
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.name
	Name *string `json:"name,omitempty"`

	// The display-name for the property, used for front-end.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Whether the property can have multiple values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_repeatable
	IsRepeatable *bool `json:"isRepeatable,omitempty"`

	// Whether the property can be filtered. If this is a sub-property, all the
	//  parent properties must be marked filterable.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_filterable
	IsFilterable *bool `json:"isFilterable,omitempty"`

	// Indicates that the property should be included in a global search.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_searchable
	IsSearchable *bool `json:"isSearchable,omitempty"`

	// Whether the property is user supplied metadata.
	//  This out-of-the box placeholder setting can be used to tag derived
	//  properties. Its value and interpretation logic should be implemented by API
	//  user.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_metadata
	IsMetadata *bool `json:"isMetadata,omitempty"`

	// Whether the property is mandatory.
	//  Default is 'false', i.e. populating property value can be skipped.
	//  If 'true' then user must populate the value for this property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_required
	IsRequired *bool `json:"isRequired,omitempty"`

	// The retrieval importance of the property during search.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.retrieval_importance
	RetrievalImportance *string `json:"retrievalImportance,omitempty"`

	// Integer property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.integer_type_options
	IntegerTypeOptions *IntegerTypeOptions `json:"integerTypeOptions,omitempty"`

	// Float property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.float_type_options
	FloatTypeOptions *FloatTypeOptions `json:"floatTypeOptions,omitempty"`

	// Text/string property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.text_type_options
	TextTypeOptions *TextTypeOptions `json:"textTypeOptions,omitempty"`

	// Nested structured data property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.property_type_options
	PropertyTypeOptions *PropertyTypeOptions `json:"propertyTypeOptions,omitempty"`

	// Enum/categorical property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.enum_type_options
	EnumTypeOptions *EnumTypeOptions `json:"enumTypeOptions,omitempty"`

	// Date time property.
	//  It is not supported by CMEK compliant deployment.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.date_time_type_options
	DateTimeTypeOptions *DateTimeTypeOptions `json:"dateTimeTypeOptions,omitempty"`

	// Map property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.map_type_options
	MapTypeOptions *MapTypeOptions `json:"mapTypeOptions,omitempty"`

	// Timestamp property.
	//  It is not supported by CMEK compliant deployment.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.timestamp_type_options
	TimestampTypeOptions *TimestampTypeOptions `json:"timestampTypeOptions,omitempty"`

	// The mapping information between this property to another schema source.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.schema_sources
	SchemaSources []PropertyDefinition_SchemaSource `json:"schemaSources,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyDefinition.SchemaSource
type PropertyDefinition_SchemaSource struct {
	// The schema name in the source.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.SchemaSource.name
	Name *string `json:"name,omitempty"`

	// The Doc AI processor type name.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.SchemaSource.processor_type
	ProcessorType *string `json:"processorType,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyTypeOptions
type PropertyTypeOptions struct {
	// Required. List of property definitions.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyTypeOptions.property_definitions
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.TextTypeOptions
type TextTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.TimestampTypeOptions
type TimestampTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DocumentSchema
type DocumentSchemaObservedState struct {
	// Output only. The time when the document schema is last updated.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time when the document schema is created.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
