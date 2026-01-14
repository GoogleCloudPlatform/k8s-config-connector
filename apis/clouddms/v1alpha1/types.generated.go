// Copyright 2026 Google LLC
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
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1
// resource: CloudDMSConversionWorkspace:ConversionWorkspace
// resource: CloudDMSPrivateConnection:PrivateConnection
// resource: CloudDMSMigrationJob:MigrationJob
// resource: CloudDMSMappingRule:MappingRule

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

// +kcc:proto=google.cloud.clouddms.v1.ConversionWorkspaceInfo
type ConversionWorkspaceInfo struct {
	// The resource name (URI) of the conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspaceInfo.name
	Name *string `json:"name,omitempty"`

	// The commit ID of the conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspaceInfo.commit_id
	CommitID *string `json:"commitID,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ConvertRowIdToColumn
type ConvertRowIDToColumn struct {
	// Required. Only work on tables without primary key defined
	// +kcc:proto:field=google.cloud.clouddms.v1.ConvertRowIdToColumn.only_if_no_primary_key
	OnlyIfNoPrimaryKey *bool `json:"onlyIfNoPrimaryKey,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.DatabaseEngineInfo
type DatabaseEngineInfo struct {
	// Required. Engine type.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.engine
	Engine *string `json:"engine,omitempty"`

	// Required. Engine named version, for example 12.c.1.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.DatabaseType
type DatabaseType struct {
	// The database provider.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseType.provider
	Provider *string `json:"provider,omitempty"`

	// The database engine.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseType.engine
	Engine *string `json:"engine,omitempty"`
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

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.DumpFlag
type MigrationJob_DumpFlag struct {
	// The name of the flag
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlag.name
	Name *string `json:"name,omitempty"`

	// The value of the flag.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlag.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.DumpFlags
type MigrationJob_DumpFlags struct {
	// The flags for the initial dump.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlags.dump_flags
	DumpFlags []MigrationJob_DumpFlag `json:"dumpFlags,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.PerformanceConfig
type MigrationJob_PerformanceConfig struct {
	// Initial dump parallelism level.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.PerformanceConfig.dump_parallel_level
	DumpParallelLevel *string `json:"dumpParallelLevel,omitempty"`
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
type SourceSQLChange struct {
	// Required. Sql code for source (stored procedure, function, trigger or view)
	// +kcc:proto:field=google.cloud.clouddms.v1.SourceSqlChange.sql_code
	SQLCode *string `json:"sqlCode,omitempty"`
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

// +kcc:proto=google.cloud.clouddms.v1.StaticIpConnectivity
type StaticIPConnectivity struct {
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
