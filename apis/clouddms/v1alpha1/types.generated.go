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


// +kcc:proto=google.cloud.clouddms.v1.ApplyHash
type ApplyHash struct {
	// Optional. Generate UUID from the data's byte array
	// +kcc:proto:field=google.cloud.clouddms.v1.ApplyHash.uuid_from_bytes
	UuidFromBytes *Empty `json:"uuidFromBytes,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AssignSpecificValue
type AssignSpecificValue struct {
	// Required. Specific value to be assigned
	// +kcc:proto:field=google.cloud.clouddms.v1.AssignSpecificValue.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ConditionalColumnSetValue
type ConditionalColumnSetValue struct {
	// Optional. Optional filter on source column length. Used for text based
	//  data types like varchar.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.source_text_filter
	SourceTextFilter *SourceTextFilter `json:"sourceTextFilter,omitempty"`

	// Optional. Optional filter on source column precision and scale. Used for
	//  fixed point numbers such as NUMERIC/NUMBER data types.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.source_numeric_filter
	SourceNumericFilter *SourceNumericFilter `json:"sourceNumericFilter,omitempty"`

	// Required. Description of data transformation during migration.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.value_transformation
	ValueTransformation *ValueTransformation `json:"valueTransformation,omitempty"`

	// Optional. Custom engine specific features.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.custom_features
	CustomFeatures map[string]string `json:"customFeatures,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ConvertRowIdToColumn
type ConvertRowIdToColumn struct {
	// Required. Only work on tables without primary key defined
	// +kcc:proto:field=google.cloud.clouddms.v1.ConvertRowIdToColumn.only_if_no_primary_key
	OnlyIfNoPrimaryKey *bool `json:"onlyIfNoPrimaryKey,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.DoubleComparisonFilter
type DoubleComparisonFilter struct {
	// Required. Relation between source value and compare value
	// +kcc:proto:field=google.cloud.clouddms.v1.DoubleComparisonFilter.value_comparison
	ValueComparison *string `json:"valueComparison,omitempty"`

	// Required. Double compare value to be used
	// +kcc:proto:field=google.cloud.clouddms.v1.DoubleComparisonFilter.value
	Value *float64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.EntityMove
type EntityMove struct {
	// Required. The new schema
	// +kcc:proto:field=google.cloud.clouddms.v1.EntityMove.new_schema
	NewSchema *string `json:"newSchema,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.FilterTableColumns
type FilterTableColumns struct {
	// Optional. List of columns to be included for a particular table.
	// +kcc:proto:field=google.cloud.clouddms.v1.FilterTableColumns.include_columns
	IncludeColumns []string `json:"includeColumns,omitempty"`

	// Optional. List of columns to be excluded for a particular table.
	// +kcc:proto:field=google.cloud.clouddms.v1.FilterTableColumns.exclude_columns
	ExcludeColumns []string `json:"excludeColumns,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.IntComparisonFilter
type IntComparisonFilter struct {
	// Required. Relation between source value and compare value
	// +kcc:proto:field=google.cloud.clouddms.v1.IntComparisonFilter.value_comparison
	ValueComparison *string `json:"valueComparison,omitempty"`

	// Required. Integer compare value to be used
	// +kcc:proto:field=google.cloud.clouddms.v1.IntComparisonFilter.value
	Value *int64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MappingRule
type MappingRule struct {
	// Full name of the mapping rule resource, in the form of:
	//  projects/{project}/locations/{location}/conversionWorkspaces/{set}/mappingRule/{rule}.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.name
	Name *string `json:"name,omitempty"`

	// Optional. A human readable name
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The mapping rule state
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.state
	State *string `json:"state,omitempty"`

	// Required. The rule scope
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.rule_scope
	RuleScope *string `json:"ruleScope,omitempty"`

	// Required. The rule filter
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.filter
	Filter *MappingRuleFilter `json:"filter,omitempty"`

	// Required. The order in which the rule is applied. Lower order rules are
	//  applied before higher value rules so they may end up being overridden.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.rule_order
	RuleOrder *int64 `json:"ruleOrder,omitempty"`

	// Optional. Rule to specify how a single entity should be renamed.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.single_entity_rename
	SingleEntityRename *SingleEntityRename `json:"singleEntityRename,omitempty"`

	// Optional. Rule to specify how multiple entities should be renamed.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.multi_entity_rename
	MultiEntityRename *MultiEntityRename `json:"multiEntityRename,omitempty"`

	// Optional. Rule to specify how multiple entities should be relocated into
	//  a different schema.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.entity_move
	EntityMove *EntityMove `json:"entityMove,omitempty"`

	// Optional. Rule to specify how a single column is converted.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.single_column_change
	SingleColumnChange *SingleColumnChange `json:"singleColumnChange,omitempty"`

	// Optional. Rule to specify how multiple columns should be converted to a
	//  different data type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.multi_column_data_type_change
	MultiColumnDataTypeChange *MultiColumnDatatypeChange `json:"multiColumnDataTypeChange,omitempty"`

	// Optional. Rule to specify how the data contained in a column should be
	//  transformed (such as trimmed, rounded, etc) provided that the data meets
	//  certain criteria.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.conditional_column_set_value
	ConditionalColumnSetValue *ConditionalColumnSetValue `json:"conditionalColumnSetValue,omitempty"`

	// Optional. Rule to specify how multiple tables should be converted with an
	//  additional rowid column.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.convert_rowid_column
	ConvertRowidColumn *ConvertRowIdToColumn `json:"convertRowidColumn,omitempty"`

	// Optional. Rule to specify the primary key for a table
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.set_table_primary_key
	SetTablePrimaryKey *SetTablePrimaryKey `json:"setTablePrimaryKey,omitempty"`

	// Optional. Rule to specify how a single package is converted.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.single_package_change
	SinglePackageChange *SinglePackageChange `json:"singlePackageChange,omitempty"`

	// Optional. Rule to change the sql code for an entity, for example,
	//  function, procedure.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.source_sql_change
	SourceSqlChange *SourceSqlChange `json:"sourceSqlChange,omitempty"`

	// Optional. Rule to specify the list of columns to include or exclude from
	//  a table.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.filter_table_columns
	FilterTableColumns *FilterTableColumns `json:"filterTableColumns,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MappingRuleFilter
type MappingRuleFilter struct {
	// Optional. The rule should be applied to entities whose parent entity
	//  (fully qualified name) matches the given value.
	//  For example, if the rule applies to a table entity, the expected value
	//  should be a schema (schema). If the rule applies to a column or index
	//  entity, the expected value can be either a schema (schema) or a table
	//  (schema.table)
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRuleFilter.parent_entity
	ParentEntity *string `json:"parentEntity,omitempty"`

	// Optional. The rule should be applied to entities whose non-qualified name
	//  starts with the given prefix.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRuleFilter.entity_name_prefix
	EntityNamePrefix *string `json:"entityNamePrefix,omitempty"`

	// Optional. The rule should be applied to entities whose non-qualified name
	//  ends with the given suffix.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRuleFilter.entity_name_suffix
	EntityNameSuffix *string `json:"entityNameSuffix,omitempty"`

	// Optional. The rule should be applied to entities whose non-qualified name
	//  contains the given string.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRuleFilter.entity_name_contains
	EntityNameContains *string `json:"entityNameContains,omitempty"`

	// Optional. The rule should be applied to specific entities defined by their
	//  fully qualified names.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRuleFilter.entities
	Entities []string `json:"entities,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MultiColumnDatatypeChange
type MultiColumnDatatypeChange struct {
	// Required. Filter on source data type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.source_data_type_filter
	SourceDataTypeFilter *string `json:"sourceDataTypeFilter,omitempty"`

	// Optional. Filter for text-based data types like varchar.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.source_text_filter
	SourceTextFilter *SourceTextFilter `json:"sourceTextFilter,omitempty"`

	// Optional. Filter for fixed point number data types such as
	//  NUMERIC/NUMBER.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.source_numeric_filter
	SourceNumericFilter *SourceNumericFilter `json:"sourceNumericFilter,omitempty"`

	// Required. New data type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.new_data_type
	NewDataType *string `json:"newDataType,omitempty"`

	// Optional. Column length - e.g. varchar (50) - if not specified and relevant
	//  uses the source column length.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_length
	OverrideLength *int64 `json:"overrideLength,omitempty"`

	// Optional. Column scale - when relevant - if not specified and relevant
	//  uses the source column scale.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_scale
	OverrideScale *int32 `json:"overrideScale,omitempty"`

	// Optional. Column precision - when relevant - if not specified and relevant
	//  uses the source column precision.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_precision
	OverridePrecision *int32 `json:"overridePrecision,omitempty"`

	// Optional. Column fractional seconds precision - used only for timestamp
	//  based datatypes - if not specified and relevant uses the source column
	//  fractional seconds precision.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_fractional_seconds_precision
	OverrideFractionalSecondsPrecision *int32 `json:"overrideFractionalSecondsPrecision,omitempty"`

	// Optional. Custom engine specific features.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.custom_features
	CustomFeatures map[string]string `json:"customFeatures,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MultiEntityRename
type MultiEntityRename struct {
	// Optional. The pattern used to generate the new entity's name. This pattern
	//  must include the characters '{name}', which will be replaced with the name
	//  of the original entity. For example, the pattern 't_{name}' for an entity
	//  name jobs would be converted to 't_jobs'.
	//
	//  If unspecified, the default value for this field is '{name}'
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiEntityRename.new_name_pattern
	NewNamePattern *string `json:"newNamePattern,omitempty"`

	// Optional. Additional transformation that can be done on the source entity
	//  name before it is being used by the new_name_pattern, for example lower
	//  case. If no transformation is desired, use NO_TRANSFORMATION
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiEntityRename.source_name_transformation
	SourceNameTransformation *string `json:"sourceNameTransformation,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.RoundToScale
type RoundToScale struct {
	// Required. Scale value to be used
	// +kcc:proto:field=google.cloud.clouddms.v1.RoundToScale.scale
	Scale *int32 `json:"scale,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SetTablePrimaryKey
type SetTablePrimaryKey struct {
	// Required. List of column names for the primary key
	// +kcc:proto:field=google.cloud.clouddms.v1.SetTablePrimaryKey.primary_key_columns
	PrimaryKeyColumns []string `json:"primaryKeyColumns,omitempty"`

	// Optional. Name for the primary key
	// +kcc:proto:field=google.cloud.clouddms.v1.SetTablePrimaryKey.primary_key
	PrimaryKey *string `json:"primaryKey,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SingleColumnChange
type SingleColumnChange struct {
	// Optional. Column data type name.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.data_type
	DataType *string `json:"dataType,omitempty"`

	// Optional. Charset override - instead of table level charset.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.charset
	Charset *string `json:"charset,omitempty"`

	// Optional. Collation override - instead of table level collation.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.collation
	Collation *string `json:"collation,omitempty"`

	// Optional. Column length - e.g. 50 as in varchar (50) - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.length
	Length *int64 `json:"length,omitempty"`

	// Optional. Column precision - e.g. 8 as in double (8,2) - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.precision
	Precision *int32 `json:"precision,omitempty"`

	// Optional. Column scale - e.g. 2 as in double (8,2) - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.scale
	Scale *int32 `json:"scale,omitempty"`

	// Optional. Column fractional seconds precision - e.g. 2 as in timestamp (2)
	//  - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.fractional_seconds_precision
	FractionalSecondsPrecision *int32 `json:"fractionalSecondsPrecision,omitempty"`

	// Optional. Is the column of array type.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.array
	Array *bool `json:"array,omitempty"`

	// Optional. The length of the array, only relevant if the column type is an
	//  array.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.array_length
	ArrayLength *int32 `json:"arrayLength,omitempty"`

	// Optional. Is the column nullable.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// Optional. Is the column auto-generated/identity.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.auto_generated
	AutoGenerated *bool `json:"autoGenerated,omitempty"`

	// Optional. Is the column a UDT (User-defined Type).
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.udt
	Udt *bool `json:"udt,omitempty"`

	// Optional. Custom engine specific features.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.custom_features
	CustomFeatures map[string]string `json:"customFeatures,omitempty"`

	// Optional. Specifies the list of values allowed in the column.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.set_values
	SetValues []string `json:"setValues,omitempty"`

	// Optional. Comment associated with the column.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.comment
	Comment *string `json:"comment,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SingleEntityRename
type SingleEntityRename struct {
	// Required. The new name of the destination entity
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleEntityRename.new_name
	NewName *string `json:"newName,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SinglePackageChange
type SinglePackageChange struct {
	// Optional. Sql code for package description
	// +kcc:proto:field=google.cloud.clouddms.v1.SinglePackageChange.package_description
	PackageDescription *string `json:"packageDescription,omitempty"`

	// Optional. Sql code for package body
	// +kcc:proto:field=google.cloud.clouddms.v1.SinglePackageChange.package_body
	PackageBody *string `json:"packageBody,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SourceNumericFilter
type SourceNumericFilter struct {
	// Optional. The filter will match columns with scale greater than or equal to
	//  this number.
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceNumericFilter.source_min_scale_filter
	SourceMinScaleFilter *int32 `json:"sourceMinScaleFilter,omitempty"`

	// Optional. The filter will match columns with scale smaller than or equal to
	//  this number.
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceNumericFilter.source_max_scale_filter
	SourceMaxScaleFilter *int32 `json:"sourceMaxScaleFilter,omitempty"`

	// Optional. The filter will match columns with precision greater than or
	//  equal to this number.
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceNumericFilter.source_min_precision_filter
	SourceMinPrecisionFilter *int32 `json:"sourceMinPrecisionFilter,omitempty"`

	// Optional. The filter will match columns with precision smaller than or
	//  equal to this number.
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceNumericFilter.source_max_precision_filter
	SourceMaxPrecisionFilter *int32 `json:"sourceMaxPrecisionFilter,omitempty"`

	// Required. Enum to set the option defining the datatypes numeric filter has
	//  to be applied to
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceNumericFilter.numeric_filter_option
	NumericFilterOption *string `json:"numericFilterOption,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SourceSqlChange
type SourceSqlChange struct {
	// Required. Sql code for source (stored procedure, function, trigger or view)
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceSqlChange.sql_code
	SqlCode *string `json:"sqlCode,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SourceTextFilter
type SourceTextFilter struct {
	// Optional. The filter will match columns with length greater than or equal
	//  to this number.
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceTextFilter.source_min_length_filter
	SourceMinLengthFilter *int64 `json:"sourceMinLengthFilter,omitempty"`

	// Optional. The filter will match columns with length smaller than or equal
	//  to this number.
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceTextFilter.source_max_length_filter
	SourceMaxLengthFilter *int64 `json:"sourceMaxLengthFilter,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ValueListFilter
type ValueListFilter struct {
	// Required. Indicates whether the filter matches rows with values that are
	//  present in the list or those with values not present in it.
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueListFilter.value_present_list
	ValuePresentList *string `json:"valuePresentList,omitempty"`

	// Required. The list to be used to filter by
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueListFilter.values
	Values []string `json:"values,omitempty"`

	// Required. Whether to ignore case when filtering by values. Defaults to
	//  false
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueListFilter.ignore_case
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ValueTransformation
type ValueTransformation struct {
	// Optional. Value is null
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.is_null
	IsNull *Empty `json:"isNull,omitempty"`

	// Optional. Value is found in the specified list.
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.value_list
	ValueList *ValueListFilter `json:"valueList,omitempty"`

	// Optional. Filter on relation between source value and compare value of
	//  type integer.
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.int_comparison
	IntComparison *IntComparisonFilter `json:"intComparison,omitempty"`

	// Optional. Filter on relation between source value and compare value of
	//  type double.
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.double_comparison
	DoubleComparison *DoubleComparisonFilter `json:"doubleComparison,omitempty"`

	// Optional. Set to null
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.assign_null
	AssignNull *Empty `json:"assignNull,omitempty"`

	// Optional. Set to a specific value (value is converted to fit the target
	//  data type)
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.assign_specific_value
	AssignSpecificValue *AssignSpecificValue `json:"assignSpecificValue,omitempty"`

	// Optional. Set to min_value - if integer or numeric, will use
	//  int.minvalue, etc
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.assign_min_value
	AssignMinValue *Empty `json:"assignMinValue,omitempty"`

	// Optional. Set to max_value - if integer or numeric, will use
	//  int.maxvalue, etc
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.assign_max_value
	AssignMaxValue *Empty `json:"assignMaxValue,omitempty"`

	// Optional. Allows the data to change scale
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.round_scale
	RoundScale *RoundToScale `json:"roundScale,omitempty"`

	// Optional. Applies a hash function on the data
	// +kcc:proto:field=google.cloud.clouddms.v1.ValueTransformation.apply_hash
	ApplyHash *ApplyHash `json:"applyHash,omitempty"`
}

// +kcc:proto=google.protobuf.Empty
type Empty struct {
}

// +kcc:proto=google.cloud.clouddms.v1.MappingRule
type MappingRuleObservedState struct {
	// Output only. The revision ID of the mapping rule.
	//  A new revision is committed whenever the mapping rule is changed in any
	//  way. The format is an 8-character hexadecimal string.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`
}
