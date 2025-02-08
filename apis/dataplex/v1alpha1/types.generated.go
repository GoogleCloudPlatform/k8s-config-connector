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


// +kcc:proto=google.cloud.dataplex.v1.DataDiscoveryResult
type DataDiscoveryResult struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoveryResult.BigQueryPublishing
type DataDiscoveryResult_BigQueryPublishing struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoverySpec
type DataDiscoverySpec struct {
	// Optional. Configuration for metadata publishing.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.bigquery_publishing_config
	BigqueryPublishingConfig *DataDiscoverySpec_BigQueryPublishingConfig `json:"bigqueryPublishingConfig,omitempty"`

	// Cloud Storage related configurations.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.storage_config
	StorageConfig *DataDiscoverySpec_StorageConfig `json:"storageConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig
type DataDiscoverySpec_BigQueryPublishingConfig struct {
	// Optional. Determines whether to  publish discovered tables as BigLake
	//  external tables or non-BigLake external tables.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig.table_type
	TableType *string `json:"tableType,omitempty"`

	// Optional. The BigQuery connection used to create BigLake tables.
	//  Must be in the form
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig.connection
	Connection *string `json:"connection,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig
type DataDiscoverySpec_StorageConfig struct {
	// Optional. Defines the data to include during discovery when only a subset
	//  of the data should be considered. Provide a list of patterns that
	//  identify the data to include. For Cloud Storage bucket assets, these
	//  patterns are interpreted as glob patterns used to match object names. For
	//  BigQuery dataset assets, these patterns are interpreted as patterns to
	//  match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.include_patterns
	IncludePatterns []string `json:"includePatterns,omitempty"`

	// Optional. Defines the data to exclude during discovery. Provide a list of
	//  patterns that identify the data to exclude. For Cloud Storage bucket
	//  assets, these patterns are interpreted as glob patterns used to match
	//  object names. For BigQuery dataset assets, these patterns are interpreted
	//  as patterns to match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.exclude_patterns
	ExcludePatterns []string `json:"excludePatterns,omitempty"`

	// Optional. Configuration for CSV data.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.csv_options
	CsvOptions *DataDiscoverySpec_StorageConfig_CsvOptions `json:"csvOptions,omitempty"`

	// Optional. Configuration for JSON data.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.json_options
	JsonOptions *DataDiscoverySpec_StorageConfig_JsonOptions `json:"jsonOptions,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions
type DataDiscoverySpec_StorageConfig_CsvOptions struct {
	// Optional. The number of rows to interpret as header rows that should be
	//  skipped when reading data rows.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions.header_rows
	HeaderRows *int32 `json:"headerRows,omitempty"`

	// Optional. The delimiter that is used to separate values. The default is
	//  `,` (comma).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions.delimiter
	Delimiter *string `json:"delimiter,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data types for CSV data.
	//  If true, all columns are registered as strings.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions.type_inference_disabled
	TypeInferenceDisabled *bool `json:"typeInferenceDisabled,omitempty"`

	// Optional. The character used to quote column values. Accepts `"`
	//  (double quotation mark) or `'` (single quotation mark). If unspecified,
	//  defaults to `"` (double quotation mark).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions.quote
	Quote *string `json:"quote,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.JsonOptions
type DataDiscoverySpec_StorageConfig_JsonOptions struct {
	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.JsonOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data types for JSON data.
	//  If true, all columns are registered as their primitive types
	//  (strings, number, or boolean).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.JsonOptions.type_inference_disabled
	TypeInferenceDisabled *bool `json:"typeInferenceDisabled,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult
type DataProfileResult struct {
	// The count of rows scanned.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.row_count
	RowCount *int64 `json:"rowCount,omitempty"`

	// The profile information per field.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.profile
	Profile *DataProfileResult_Profile `json:"profile,omitempty"`

	// The data scanned for this result.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.scanned_data
	ScannedData *ScannedData `json:"scannedData,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult
type DataProfileResult_PostScanActionsResult struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult.BigQueryExportResult
type DataProfileResult_PostScanActionsResult_BigQueryExportResult struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.Profile
type DataProfileResult_Profile struct {
	// List of fields with structural and profile information for each field.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.fields
	Fields []DataProfileResult_Profile_Field `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.Profile.Field
type DataProfileResult_Profile_Field struct {
	// The name of the field.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.name
	Name *string `json:"name,omitempty"`

	// The data type retrieved from the schema of the data source. For
	//  instance, for a BigQuery native table, it is the [BigQuery Table
	//  Schema](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#tablefieldschema).
	//  For a Dataplex Entity, it is the [Entity
	//  Schema](https://cloud.google.com/dataplex/docs/reference/rpc/google.cloud.dataplex.v1#type_3).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.type
	Type *string `json:"type,omitempty"`

	// The mode of the field. Possible values include:
	//
	//  * REQUIRED, if it is a required field.
	//  * NULLABLE, if it is an optional field.
	//  * REPEATED, if it is a repeated field.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.mode
	Mode *string `json:"mode,omitempty"`

	// Profile information for the corresponding field.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.profile
	Profile *DataProfileResult_Profile_Field_ProfileInfo `json:"profile,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo
type DataProfileResult_Profile_Field_ProfileInfo struct {
	// Ratio of rows with null value against total scanned rows.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.null_ratio
	NullRatio *float64 `json:"nullRatio,omitempty"`

	// Ratio of rows with distinct values against total scanned rows.
	//  Not available for complex non-groupable field type, including RECORD,
	//  ARRAY, GEOGRAPHY, and JSON, as well as fields with REPEATABLE mode.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.distinct_ratio
	DistinctRatio *float64 `json:"distinctRatio,omitempty"`

	// The list of top N non-null values, frequency and ratio with which
	//  they occur in the scanned data. N is 10 or equal to the number of
	//  distinct values in the field, whichever is smaller. Not available for
	//  complex non-groupable field type, including RECORD, ARRAY, GEOGRAPHY,
	//  and JSON, as well as fields with REPEATABLE mode.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.top_n_values
	TopNValues []DataProfileResult_Profile_Field_ProfileInfo_TopNValue `json:"topNValues,omitempty"`

	// String type field information.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.string_profile
	StringProfile *DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo `json:"stringProfile,omitempty"`

	// Integer type field information.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.integer_profile
	IntegerProfile *DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo `json:"integerProfile,omitempty"`

	// Double type field information.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.double_profile
	DoubleProfile *DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo `json:"doubleProfile,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.DoubleFieldInfo
type DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo struct {
	// Average of non-null values in the scanned data. NaN, if the field
	//  has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.DoubleFieldInfo.average
	Average *float64 `json:"average,omitempty"`

	// Standard deviation of non-null values in the scanned data. NaN, if
	//  the field has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.DoubleFieldInfo.standard_deviation
	StandardDeviation *float64 `json:"standardDeviation,omitempty"`

	// Minimum of non-null values in the scanned data. NaN, if the field
	//  has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.DoubleFieldInfo.min
	Min *float64 `json:"min,omitempty"`

	// A quartile divides the number of data points into four parts, or
	//  quarters, of more-or-less equal size. Three main quartiles used
	//  are: The first quartile (Q1) splits off the lowest 25% of data from
	//  the highest 75%. It is also known as the lower or 25th empirical
	//  quartile, as 25% of the data is below this point. The second
	//  quartile (Q2) is the median of a data set. So, 50% of the data lies
	//  below this point. The third quartile (Q3) splits off the highest
	//  25% of data from the lowest 75%. It is known as the upper or 75th
	//  empirical quartile, as 75% of the data lies below this point.
	//  Here, the quartiles is provided as an ordered list of quartile
	//  values for the scanned data, occurring in order Q1, median, Q3.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.DoubleFieldInfo.quartiles
	Quartiles []float64 `json:"quartiles,omitempty"`

	// Maximum of non-null values in the scanned data. NaN, if the field
	//  has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.DoubleFieldInfo.max
	Max *float64 `json:"max,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.IntegerFieldInfo
type DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo struct {
	// Average of non-null values in the scanned data. NaN, if the field
	//  has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.IntegerFieldInfo.average
	Average *float64 `json:"average,omitempty"`

	// Standard deviation of non-null values in the scanned data. NaN, if
	//  the field has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.IntegerFieldInfo.standard_deviation
	StandardDeviation *float64 `json:"standardDeviation,omitempty"`

	// Minimum of non-null values in the scanned data. NaN, if the field
	//  has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.IntegerFieldInfo.min
	Min *int64 `json:"min,omitempty"`

	// A quartile divides the number of data points into four parts, or
	//  quarters, of more-or-less equal size. Three main quartiles used
	//  are: The first quartile (Q1) splits off the lowest 25% of data from
	//  the highest 75%. It is also known as the lower or 25th empirical
	//  quartile, as 25% of the data is below this point. The second
	//  quartile (Q2) is the median of a data set. So, 50% of the data lies
	//  below this point. The third quartile (Q3) splits off the highest
	//  25% of data from the lowest 75%. It is known as the upper or 75th
	//  empirical quartile, as 75% of the data lies below this point.
	//  Here, the quartiles is provided as an ordered list of approximate
	//  quartile values for the scanned data, occurring in order Q1,
	//  median, Q3.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.IntegerFieldInfo.quartiles
	Quartiles []int64 `json:"quartiles,omitempty"`

	// Maximum of non-null values in the scanned data. NaN, if the field
	//  has a NaN.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.IntegerFieldInfo.max
	Max *int64 `json:"max,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.StringFieldInfo
type DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo struct {
	// Minimum length of non-null values in the scanned data.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.StringFieldInfo.min_length
	MinLength *int64 `json:"minLength,omitempty"`

	// Maximum length of non-null values in the scanned data.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.StringFieldInfo.max_length
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Average length of non-null values in the scanned data.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.StringFieldInfo.average_length
	AverageLength *float64 `json:"averageLength,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.TopNValue
type DataProfileResult_Profile_Field_ProfileInfo_TopNValue struct {
	// String value of a top N non-null value.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.TopNValue.value
	Value *string `json:"value,omitempty"`

	// Count of the corresponding value in the scanned data.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.TopNValue.count
	Count *int64 `json:"count,omitempty"`

	// Ratio of the corresponding value in the field against the total
	//  number of rows in the scanned data.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.Profile.Field.ProfileInfo.TopNValue.ratio
	Ratio *float64 `json:"ratio,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileSpec
type DataProfileSpec struct {
	// Optional. The percentage of the records to be selected from the dataset for
	//  DataScan.
	//
	//  * Value can range between 0.0 and 100.0 with up to 3 significant decimal
	//  digits.
	//  * Sampling is not applied if `sampling_percent` is not specified, 0 or
	//  100.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.sampling_percent
	SamplingPercent *float32 `json:"samplingPercent,omitempty"`

	// Optional. A filter applied to all rows in a single DataScan job.
	//  The filter needs to be a valid SQL expression for a WHERE clause in
	//  BigQuery standard SQL syntax.
	//  Example: col1 >= 0 AND col2 < 10
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.row_filter
	RowFilter *string `json:"rowFilter,omitempty"`

	// Optional. Actions to take upon job completion..
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.post_scan_actions
	PostScanActions *DataProfileSpec_PostScanActions `json:"postScanActions,omitempty"`

	// Optional. The fields to include in data profile.
	//
	//  If not specified, all fields at the time of profile scan job execution are
	//  included, except for ones listed in `exclude_fields`.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.include_fields
	IncludeFields *DataProfileSpec_SelectedFields `json:"includeFields,omitempty"`

	// Optional. The fields to exclude from data profile.
	//
	//  If specified, the fields will be excluded from data profile, regardless of
	//  `include_fields` value.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.exclude_fields
	ExcludeFields *DataProfileSpec_SelectedFields `json:"excludeFields,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileSpec.PostScanActions
type DataProfileSpec_PostScanActions struct {
	// Optional. If set, results will be exported to the provided BigQuery
	//  table.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.PostScanActions.bigquery_export
	BigqueryExport *DataProfileSpec_PostScanActions_BigQueryExport `json:"bigqueryExport,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileSpec.PostScanActions.BigQueryExport
type DataProfileSpec_PostScanActions_BigQueryExport struct {
	// Optional. The BigQuery table to export DataProfileScan results to.
	//  Format:
	//  //bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.PostScanActions.BigQueryExport.results_table
	ResultsTable *string `json:"resultsTable,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileSpec.SelectedFields
type DataProfileSpec_SelectedFields struct {
	// Optional. Expected input is a list of fully qualified names of fields as
	//  in the schema.
	//
	//  Only top-level field names for nested fields are supported.
	//  For instance, if 'x' is of nested field type, listing 'x' is supported
	//  but 'x.y.z' is not supported. Here 'y' and 'y.z' are nested fields of
	//  'x'.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileSpec.SelectedFields.field_names
	FieldNames []string `json:"fieldNames,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityColumnResult
type DataQualityColumnResult struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityDimension
type DataQualityDimension struct {
	// The dimension name a rule belongs to. Supported dimensions are
	//  ["COMPLETENESS", "ACCURACY", "CONSISTENCY", "VALIDITY", "UNIQUENESS",
	//  "FRESHNESS", "VOLUME"]
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityDimension.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityDimensionResult
type DataQualityDimensionResult struct {

	// Whether the dimension passed or failed.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityDimensionResult.passed
	Passed *bool `json:"passed,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityResult
type DataQualityResult struct {
	// Overall data quality result -- `true` if all rules passed.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.passed
	Passed *bool `json:"passed,omitempty"`

	// A list of results at the dimension level.
	//
	//  A dimension will have a corresponding `DataQualityDimensionResult` if and
	//  only if there is at least one rule with the 'dimension' field set to it.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.dimensions
	Dimensions []DataQualityDimensionResult `json:"dimensions,omitempty"`

	// A list of all the rules in a job, and their results.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.rules
	Rules []DataQualityRuleResult `json:"rules,omitempty"`

	// The count of rows processed.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.row_count
	RowCount *int64 `json:"rowCount,omitempty"`

	// The data scanned for this result.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.scanned_data
	ScannedData *ScannedData `json:"scannedData,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult
type DataQualityResult_PostScanActionsResult struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult.BigQueryExportResult
type DataQualityResult_PostScanActionsResult_BigQueryExportResult struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule
type DataQualityRule struct {
	// Row-level rule which evaluates whether each column value lies between a
	//  specified range.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.range_expectation
	RangeExpectation *DataQualityRule_RangeExpectation `json:"rangeExpectation,omitempty"`

	// Row-level rule which evaluates whether each column value is null.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.non_null_expectation
	NonNullExpectation *DataQualityRule_NonNullExpectation `json:"nonNullExpectation,omitempty"`

	// Row-level rule which evaluates whether each column value is contained by
	//  a specified set.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.set_expectation
	SetExpectation *DataQualityRule_SetExpectation `json:"setExpectation,omitempty"`

	// Row-level rule which evaluates whether each column value matches a
	//  specified regex.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.regex_expectation
	RegexExpectation *DataQualityRule_RegexExpectation `json:"regexExpectation,omitempty"`

	// Row-level rule which evaluates whether each column value is unique.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.uniqueness_expectation
	UniquenessExpectation *DataQualityRule_UniquenessExpectation `json:"uniquenessExpectation,omitempty"`

	// Aggregate rule which evaluates whether the column aggregate
	//  statistic lies between a specified range.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.statistic_range_expectation
	StatisticRangeExpectation *DataQualityRule_StatisticRangeExpectation `json:"statisticRangeExpectation,omitempty"`

	// Row-level rule which evaluates whether each row in a table passes the
	//  specified condition.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.row_condition_expectation
	RowConditionExpectation *DataQualityRule_RowConditionExpectation `json:"rowConditionExpectation,omitempty"`

	// Aggregate rule which evaluates whether the provided expression is true
	//  for a table.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.table_condition_expectation
	TableConditionExpectation *DataQualityRule_TableConditionExpectation `json:"tableConditionExpectation,omitempty"`

	// Aggregate rule which evaluates the number of rows returned for the
	//  provided statement. If any rows are returned, this rule fails.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.sql_assertion
	SqlAssertion *DataQualityRule_SqlAssertion `json:"sqlAssertion,omitempty"`

	// Optional. The unnested column which this rule is evaluated against.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.column
	Column *string `json:"column,omitempty"`

	// Optional. Rows with `null` values will automatically fail a rule, unless
	//  `ignore_null` is `true`. In that case, such `null` rows are trivially
	//  considered passing.
	//
	//  This field is only valid for the following type of rules:
	//
	//  * RangeExpectation
	//  * RegexExpectation
	//  * SetExpectation
	//  * UniquenessExpectation
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.ignore_null
	IgnoreNull *bool `json:"ignoreNull,omitempty"`

	// Required. The dimension a rule belongs to. Results are also aggregated at
	//  the dimension level. Supported dimensions are **["COMPLETENESS",
	//  "ACCURACY", "CONSISTENCY", "VALIDITY", "UNIQUENESS", "FRESHNESS",
	//  "VOLUME"]**
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.dimension
	Dimension *string `json:"dimension,omitempty"`

	// Optional. The minimum ratio of **passing_rows / total_rows** required to
	//  pass this rule, with a range of [0.0, 1.0].
	//
	//  0 indicates default value (i.e. 1.0).
	//
	//  This field is only valid for row-level type rules.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.threshold
	Threshold *float64 `json:"threshold,omitempty"`

	// Optional. A mutable name for the rule.
	//
	//  * The name must contain only letters (a-z, A-Z), numbers (0-9), or
	//  hyphens (-).
	//  * The maximum length is 63 characters.
	//  * Must start with a letter.
	//  * Must end with a number or a letter.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.name
	Name *string `json:"name,omitempty"`

	// Optional. Description of the rule.
	//
	//  * The maximum length is 1,024 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.description
	Description *string `json:"description,omitempty"`

	// Optional. Whether the Rule is active or suspended.
	//  Default is false.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.suspended
	Suspended *bool `json:"suspended,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.NonNullExpectation
type DataQualityRule_NonNullExpectation struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.RangeExpectation
type DataQualityRule_RangeExpectation struct {
	// Optional. The minimum column value allowed for a row to pass this
	//  validation. At least one of `min_value` and `max_value` need to be
	//  provided.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.RangeExpectation.min_value
	MinValue *string `json:"minValue,omitempty"`

	// Optional. The maximum column value allowed for a row to pass this
	//  validation. At least one of `min_value` and `max_value` need to be
	//  provided.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.RangeExpectation.max_value
	MaxValue *string `json:"maxValue,omitempty"`

	// Optional. Whether each value needs to be strictly greater than ('>') the
	//  minimum, or if equality is allowed.
	//
	//  Only relevant if a `min_value` has been defined. Default = false.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.RangeExpectation.strict_min_enabled
	StrictMinEnabled *bool `json:"strictMinEnabled,omitempty"`

	// Optional. Whether each value needs to be strictly lesser than ('<') the
	//  maximum, or if equality is allowed.
	//
	//  Only relevant if a `max_value` has been defined. Default = false.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.RangeExpectation.strict_max_enabled
	StrictMaxEnabled *bool `json:"strictMaxEnabled,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.RegexExpectation
type DataQualityRule_RegexExpectation struct {
	// Optional. A regular expression the column value is expected to match.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.RegexExpectation.regex
	Regex *string `json:"regex,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.RowConditionExpectation
type DataQualityRule_RowConditionExpectation struct {
	// Optional. The SQL expression.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.RowConditionExpectation.sql_expression
	SqlExpression *string `json:"sqlExpression,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.SetExpectation
type DataQualityRule_SetExpectation struct {
	// Optional. Expected values for the column value.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.SetExpectation.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.SqlAssertion
type DataQualityRule_SqlAssertion struct {
	// Optional. The SQL statement.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.SqlAssertion.sql_statement
	SqlStatement *string `json:"sqlStatement,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.StatisticRangeExpectation
type DataQualityRule_StatisticRangeExpectation struct {
	// Optional. The aggregate metric to evaluate.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.StatisticRangeExpectation.statistic
	Statistic *string `json:"statistic,omitempty"`

	// Optional. The minimum column statistic value allowed for a row to pass
	//  this validation.
	//
	//  At least one of `min_value` and `max_value` need to be provided.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.StatisticRangeExpectation.min_value
	MinValue *string `json:"minValue,omitempty"`

	// Optional. The maximum column statistic value allowed for a row to pass
	//  this validation.
	//
	//  At least one of `min_value` and `max_value` need to be provided.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.StatisticRangeExpectation.max_value
	MaxValue *string `json:"maxValue,omitempty"`

	// Optional. Whether column statistic needs to be strictly greater than
	//  ('>') the minimum, or if equality is allowed.
	//
	//  Only relevant if a `min_value` has been defined. Default = false.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.StatisticRangeExpectation.strict_min_enabled
	StrictMinEnabled *bool `json:"strictMinEnabled,omitempty"`

	// Optional. Whether column statistic needs to be strictly lesser than ('<')
	//  the maximum, or if equality is allowed.
	//
	//  Only relevant if a `max_value` has been defined. Default = false.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.StatisticRangeExpectation.strict_max_enabled
	StrictMaxEnabled *bool `json:"strictMaxEnabled,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.TableConditionExpectation
type DataQualityRule_TableConditionExpectation struct {
	// Optional. The SQL expression.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRule.TableConditionExpectation.sql_expression
	SqlExpression *string `json:"sqlExpression,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.UniquenessExpectation
type DataQualityRule_UniquenessExpectation struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRuleResult
type DataQualityRuleResult struct {
	// The rule specified in the DataQualitySpec, as is.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.rule
	Rule *DataQualityRule `json:"rule,omitempty"`

	// Whether the rule passed or failed.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.passed
	Passed *bool `json:"passed,omitempty"`

	// The number of rows a rule was evaluated against.
	//
	//  This field is only valid for row-level type rules.
	//
	//  Evaluated count can be configured to either
	//
	//  * include all rows (default) - with `null` rows automatically failing rule
	//  evaluation, or
	//  * exclude `null` rows from the `evaluated_count`, by setting
	//  `ignore_nulls = true`.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.evaluated_count
	EvaluatedCount *int64 `json:"evaluatedCount,omitempty"`

	// The number of rows which passed a rule evaluation.
	//
	//  This field is only valid for row-level type rules.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.passed_count
	PassedCount *int64 `json:"passedCount,omitempty"`

	// The number of rows with null values in the specified column.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.null_count
	NullCount *int64 `json:"nullCount,omitempty"`

	// The ratio of **passed_count / evaluated_count**.
	//
	//  This field is only valid for row-level type rules.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.pass_ratio
	PassRatio *float64 `json:"passRatio,omitempty"`

	// The query to find rows that did not pass this rule.
	//
	//  This field is only valid for row-level type rules.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.failing_rows_query
	FailingRowsQuery *string `json:"failingRowsQuery,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec
type DataQualitySpec struct {
	// Required. The list of rules to evaluate against a data source. At least one
	//  rule is required.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.rules
	Rules []DataQualityRule `json:"rules,omitempty"`

	// Optional. The percentage of the records to be selected from the dataset for
	//  DataScan.
	//
	//  * Value can range between 0.0 and 100.0 with up to 3 significant decimal
	//  digits.
	//  * Sampling is not applied if `sampling_percent` is not specified, 0 or
	//  100.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.sampling_percent
	SamplingPercent *float32 `json:"samplingPercent,omitempty"`

	// Optional. A filter applied to all rows in a single DataScan job.
	//  The filter needs to be a valid SQL expression for a WHERE clause in
	//  BigQuery standard SQL syntax.
	//  Example: col1 >= 0 AND col2 < 10
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.row_filter
	RowFilter *string `json:"rowFilter,omitempty"`

	// Optional. Actions to take upon job completion.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.post_scan_actions
	PostScanActions *DataQualitySpec_PostScanActions `json:"postScanActions,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions
type DataQualitySpec_PostScanActions struct {
	// Optional. If set, results will be exported to the provided BigQuery
	//  table.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.bigquery_export
	BigqueryExport *DataQualitySpec_PostScanActions_BigQueryExport `json:"bigqueryExport,omitempty"`

	// Optional. If set, results will be sent to the provided notification
	//  receipts upon triggers.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.notification_report
	NotificationReport *DataQualitySpec_PostScanActions_NotificationReport `json:"notificationReport,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.BigQueryExport
type DataQualitySpec_PostScanActions_BigQueryExport struct {
	// Optional. The BigQuery table to export DataQualityScan results to.
	//  Format:
	//  //bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.BigQueryExport.results_table
	ResultsTable *string `json:"resultsTable,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.JobEndTrigger
type DataQualitySpec_PostScanActions_JobEndTrigger struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.JobFailureTrigger
type DataQualitySpec_PostScanActions_JobFailureTrigger struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.NotificationReport
type DataQualitySpec_PostScanActions_NotificationReport struct {
	// Required. The recipients who will receive the notification report.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.NotificationReport.recipients
	Recipients *DataQualitySpec_PostScanActions_Recipients `json:"recipients,omitempty"`

	// Optional. If set, report will be sent when score threshold is met.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.NotificationReport.score_threshold_trigger
	ScoreThresholdTrigger *DataQualitySpec_PostScanActions_ScoreThresholdTrigger `json:"scoreThresholdTrigger,omitempty"`

	// Optional. If set, report will be sent when a scan job fails.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.NotificationReport.job_failure_trigger
	JobFailureTrigger *DataQualitySpec_PostScanActions_JobFailureTrigger `json:"jobFailureTrigger,omitempty"`

	// Optional. If set, report will be sent when a scan job ends.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.NotificationReport.job_end_trigger
	JobEndTrigger *DataQualitySpec_PostScanActions_JobEndTrigger `json:"jobEndTrigger,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.Recipients
type DataQualitySpec_PostScanActions_Recipients struct {
	// Optional. The email recipients who will receive the DataQualityScan
	//  results report.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.Recipients.emails
	Emails []string `json:"emails,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.ScoreThresholdTrigger
type DataQualitySpec_PostScanActions_ScoreThresholdTrigger struct {
	// Optional. The score range is in [0,100].
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.ScoreThresholdTrigger.score_threshold
	ScoreThreshold *float32 `json:"scoreThreshold,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataScanJob
type DataScanJob struct {
}

// +kcc:proto=google.cloud.dataplex.v1.ScannedData
type ScannedData struct {
	// The range denoted by values of an incremental field
	// +kcc:proto:field=google.cloud.dataplex.v1.ScannedData.incremental_field
	IncrementalField *ScannedData_IncrementalField `json:"incrementalField,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.ScannedData.IncrementalField
type ScannedData_IncrementalField struct {
	// The field that contains values which monotonically increases over time
	//  (e.g. a timestamp column).
	// +kcc:proto:field=google.cloud.dataplex.v1.ScannedData.IncrementalField.field
	Field *string `json:"field,omitempty"`

	// Value that marks the start of the range.
	// +kcc:proto:field=google.cloud.dataplex.v1.ScannedData.IncrementalField.start
	Start *string `json:"start,omitempty"`

	// Value that marks the end of the range.
	// +kcc:proto:field=google.cloud.dataplex.v1.ScannedData.IncrementalField.end
	End *string `json:"end,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoveryResult
type DataDiscoveryResultObservedState struct {
	// Output only. Configuration for metadata publishing.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoveryResult.bigquery_publishing
	BigqueryPublishing *DataDiscoveryResult_BigQueryPublishing `json:"bigqueryPublishing,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoveryResult.BigQueryPublishing
type DataDiscoveryResult_BigQueryPublishingObservedState struct {
	// Output only. The BigQuery dataset to publish to. It takes the form
	//  `projects/{project_id}/datasets/{dataset_id}`.
	//  If not set, the service creates a default publishing dataset.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataDiscoveryResult.BigQueryPublishing.dataset
	Dataset *string `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult
type DataProfileResultObservedState struct {
	// Output only. The result of post scan actions.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.post_scan_actions_result
	PostScanActionsResult *DataProfileResult_PostScanActionsResult `json:"postScanActionsResult,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult
type DataProfileResult_PostScanActionsResultObservedState struct {
	// Output only. The result of BigQuery export post scan action.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult.bigquery_export_result
	BigqueryExportResult *DataProfileResult_PostScanActionsResult_BigQueryExportResult `json:"bigqueryExportResult,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult.BigQueryExportResult
type DataProfileResult_PostScanActionsResult_BigQueryExportResultObservedState struct {
	// Output only. Execution state for the BigQuery exporting.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult.BigQueryExportResult.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the BigQuery exporting.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult.BigQueryExportResult.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityColumnResult
type DataQualityColumnResultObservedState struct {
	// Output only. The column specified in the DataQualityRule.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityColumnResult.column
	Column *string `json:"column,omitempty"`

	// Output only. The column-level data quality score for this data scan job if
	//  and only if the 'column' field is set.
	//
	//  The score ranges between between [0, 100] (up to two decimal
	//  points).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityColumnResult.score
	Score *float32 `json:"score,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityDimensionResult
type DataQualityDimensionResultObservedState struct {
	// Output only. The dimension config specified in the DataQualitySpec, as is.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityDimensionResult.dimension
	Dimension *DataQualityDimension `json:"dimension,omitempty"`

	// Output only. The dimension-level data quality score for this data scan job
	//  if and only if the 'dimension' field is set.
	//
	//  The score ranges between [0, 100] (up to two decimal
	//  points).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityDimensionResult.score
	Score *float32 `json:"score,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityResult
type DataQualityResultObservedState struct {
	// Output only. The overall data quality score.
	//
	//  The score ranges between [0, 100] (up to two decimal points).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.score
	Score *float32 `json:"score,omitempty"`

	// A list of results at the dimension level.
	//
	//  A dimension will have a corresponding `DataQualityDimensionResult` if and
	//  only if there is at least one rule with the 'dimension' field set to it.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.dimensions
	Dimensions []DataQualityDimensionResultObservedState `json:"dimensions,omitempty"`

	// Output only. A list of results at the column level.
	//
	//  A column will have a corresponding `DataQualityColumnResult` if and only if
	//  there is at least one rule with the 'column' field set to it.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.columns
	Columns []DataQualityColumnResult `json:"columns,omitempty"`

	// A list of all the rules in a job, and their results.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.rules
	Rules []DataQualityRuleResultObservedState `json:"rules,omitempty"`

	// Output only. The result of post scan actions.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.post_scan_actions_result
	PostScanActionsResult *DataQualityResult_PostScanActionsResult `json:"postScanActionsResult,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult
type DataQualityResult_PostScanActionsResultObservedState struct {
	// Output only. The result of BigQuery export post scan action.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult.bigquery_export_result
	BigqueryExportResult *DataQualityResult_PostScanActionsResult_BigQueryExportResult `json:"bigqueryExportResult,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult.BigQueryExportResult
type DataQualityResult_PostScanActionsResult_BigQueryExportResultObservedState struct {
	// Output only. Execution state for the BigQuery exporting.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult.BigQueryExportResult.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the BigQuery exporting.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult.BigQueryExportResult.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRuleResult
type DataQualityRuleResultObservedState struct {
	// Output only. The number of rows returned by the SQL statement in a SQL
	//  assertion rule.
	//
	//  This field is only valid for SQL assertion rules.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataQualityRuleResult.assertion_row_count
	AssertionRowCount *int64 `json:"assertionRowCount,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataScanJob
type DataScanJobObservedState struct {
	// Output only. Identifier. The relative resource name of the DataScanJob, of
	//  the form:
	//  `projects/{project}/locations/{location_id}/dataScans/{datascan_id}/jobs/{job_id}`,
	//  where `project` refers to a *project_id* or *project_number* and
	//  `location_id` refers to a GCP region.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the DataScanJob.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the DataScanJob was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the DataScanJob was started.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the DataScanJob ended.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Execution state for the DataScanJob.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current state.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.message
	Message *string `json:"message,omitempty"`

	// Output only. The type of the parent DataScan.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.type
	Type *string `json:"type,omitempty"`

	// Output only. Settings for a data quality scan.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.data_quality_spec
	DataQualitySpec *DataQualitySpec `json:"dataQualitySpec,omitempty"`

	// Output only. Settings for a data profile scan.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.data_profile_spec
	DataProfileSpec *DataProfileSpec `json:"dataProfileSpec,omitempty"`

	// Output only. Settings for a data discovery scan.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.data_discovery_spec
	DataDiscoverySpec *DataDiscoverySpec `json:"dataDiscoverySpec,omitempty"`

	// Output only. The result of a data quality scan.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.data_quality_result
	DataQualityResult *DataQualityResult `json:"dataQualityResult,omitempty"`

	// Output only. The result of a data profile scan.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.data_profile_result
	DataProfileResult *DataProfileResult `json:"dataProfileResult,omitempty"`

	// Output only. The result of a data discovery scan.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataScanJob.data_discovery_result
	DataDiscoveryResult *DataDiscoveryResult `json:"dataDiscoveryResult,omitempty"`
}
