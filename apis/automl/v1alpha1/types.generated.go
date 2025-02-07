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


// +kcc:proto=google.cloud.automl.v1beta1.ArrayStats
type ArrayStats struct {
	// Stats of all the values of all arrays, as if they were a single long
	//  series of data. The type depends on the element type of the array.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ArrayStats.member_stats
	MemberStats *DataStats `json:"memberStats,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.CategoryStats
type CategoryStats struct {
	// The statistics of the top 20 CATEGORY values, ordered by
	//
	//  [count][google.cloud.automl.v1beta1.CategoryStats.SingleCategoryStats.count].
	// +kcc:proto:field=google.cloud.automl.v1beta1.CategoryStats.top_category_stats
	TopCategoryStats []CategoryStats_SingleCategoryStats `json:"topCategoryStats,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.CategoryStats.SingleCategoryStats
type CategoryStats_SingleCategoryStats struct {
	// The CATEGORY value.
	// +kcc:proto:field=google.cloud.automl.v1beta1.CategoryStats.SingleCategoryStats.value
	Value *string `json:"value,omitempty"`

	// The number of occurrences of this value in the series.
	// +kcc:proto:field=google.cloud.automl.v1beta1.CategoryStats.SingleCategoryStats.count
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.ColumnSpec
type ColumnSpec struct {
	// Output only. The resource name of the column specs.
	//  Form:
	//
	//  `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/tableSpecs/{table_spec_id}/columnSpecs/{column_spec_id}`
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.name
	Name *string `json:"name,omitempty"`

	// The data type of elements stored in the column.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.data_type
	DataType *DataType `json:"dataType,omitempty"`

	// Output only. The name of the column to show in the interface. The name can
	//  be up to 100 characters long and can consist only of ASCII Latin letters
	//  A-Z and a-z, ASCII digits 0-9, underscores(_), and forward slashes(/), and
	//  must start with a letter or a digit.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Stats of the series of values in the column.
	//  This field may be stale, see the ancestor's
	//  Dataset.tables_dataset_metadata.stats_update_time field
	//  for the timestamp at which these stats were last updated.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.data_stats
	DataStats *DataStats `json:"dataStats,omitempty"`

	// Deprecated.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.top_correlated_columns
	TopCorrelatedColumns []ColumnSpec_CorrelatedColumn `json:"topCorrelatedColumns,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.ColumnSpec.CorrelatedColumn
type ColumnSpec_CorrelatedColumn struct {
	// The column_spec_id of the correlated column, which belongs to the same
	//  table as the in-context column.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.CorrelatedColumn.column_spec_id
	ColumnSpecID *string `json:"columnSpecID,omitempty"`

	// Correlation between this and the in-context column.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ColumnSpec.CorrelatedColumn.correlation_stats
	CorrelationStats *CorrelationStats `json:"correlationStats,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.CorrelationStats
type CorrelationStats struct {
	// The correlation value using the Cramer's V measure.
	// +kcc:proto:field=google.cloud.automl.v1beta1.CorrelationStats.cramers_v
	CramersV *float64 `json:"cramersV,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.DataStats
type DataStats struct {
	// The statistics for FLOAT64 DataType.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.float64_stats
	Float64Stats *Float64Stats `json:"float64Stats,omitempty"`

	// The statistics for STRING DataType.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.string_stats
	StringStats *StringStats `json:"stringStats,omitempty"`

	// The statistics for TIMESTAMP DataType.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.timestamp_stats
	TimestampStats *TimestampStats `json:"timestampStats,omitempty"`

	// The statistics for ARRAY DataType.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.array_stats
	ArrayStats *ArrayStats `json:"arrayStats,omitempty"`

	// The statistics for STRUCT DataType.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.struct_stats
	StructStats *StructStats `json:"structStats,omitempty"`

	// The statistics for CATEGORY DataType.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.category_stats
	CategoryStats *CategoryStats `json:"categoryStats,omitempty"`

	// The number of distinct values.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.distinct_value_count
	DistinctValueCount *int64 `json:"distinctValueCount,omitempty"`

	// The number of values that are null.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.null_value_count
	NullValueCount *int64 `json:"nullValueCount,omitempty"`

	// The number of values that are valid.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataStats.valid_value_count
	ValidValueCount *int64 `json:"validValueCount,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.DataType
type DataType struct {
	// If [type_code][google.cloud.automl.v1beta1.DataType.type_code] == [ARRAY][google.cloud.automl.v1beta1.TypeCode.ARRAY],
	//  then `list_element_type` is the type of the elements.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataType.list_element_type
	ListElementType *DataType `json:"listElementType,omitempty"`

	// If [type_code][google.cloud.automl.v1beta1.DataType.type_code] == [STRUCT][google.cloud.automl.v1beta1.TypeCode.STRUCT], then `struct_type`
	//  provides type information for the struct's fields.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataType.struct_type
	StructType *StructType `json:"structType,omitempty"`

	// If [type_code][google.cloud.automl.v1beta1.DataType.type_code] == [TIMESTAMP][google.cloud.automl.v1beta1.TypeCode.TIMESTAMP]
	//  then `time_format` provides the format in which that time field is
	//  expressed. The time_format must either be one of:
	//  * `UNIX_SECONDS`
	//  * `UNIX_MILLISECONDS`
	//  * `UNIX_MICROSECONDS`
	//  * `UNIX_NANOSECONDS`
	//  (for respectively number of seconds, milliseconds, microseconds and
	//  nanoseconds since start of the Unix epoch);
	//  or be written in `strftime` syntax. If time_format is not set, then the
	//  default format as described on the type_code is used.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataType.time_format
	TimeFormat *string `json:"timeFormat,omitempty"`

	// Required. The [TypeCode][google.cloud.automl.v1beta1.TypeCode] for this type.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataType.type_code
	TypeCode *string `json:"typeCode,omitempty"`

	// If true, this DataType can also be `NULL`. In .CSV files `NULL` value is
	//  expressed as an empty string.
	// +kcc:proto:field=google.cloud.automl.v1beta1.DataType.nullable
	Nullable *bool `json:"nullable,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.Float64Stats
type Float64Stats struct {
	// The mean of the series.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Float64Stats.mean
	Mean *float64 `json:"mean,omitempty"`

	// The standard deviation of the series.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Float64Stats.standard_deviation
	StandardDeviation *float64 `json:"standardDeviation,omitempty"`

	// Ordered from 0 to k k-quantile values of the data series of n values.
	//  The value at index i is, approximately, the i*n/k-th smallest value in the
	//  series; for i = 0 and i = k these are, respectively, the min and max
	//  values.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Float64Stats.quantiles
	Quantiles []float64 `json:"quantiles,omitempty"`

	// Histogram buckets of the data series. Sorted by the min value of the
	//  bucket, ascendingly, and the number of the buckets is dynamically
	//  generated. The buckets are non-overlapping and completely cover whole
	//  FLOAT64 range with min of first bucket being `"-Infinity"`, and max of
	//  the last one being `"Infinity"`.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Float64Stats.histogram_buckets
	HistogramBuckets []Float64Stats_HistogramBucket `json:"histogramBuckets,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.Float64Stats.HistogramBucket
type Float64Stats_HistogramBucket struct {
	// The minimum value of the bucket, inclusive.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Float64Stats.HistogramBucket.min
	Min *float64 `json:"min,omitempty"`

	// The maximum value of the bucket, exclusive unless max = `"Infinity"`, in
	//  which case it's inclusive.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Float64Stats.HistogramBucket.max
	Max *float64 `json:"max,omitempty"`

	// The number of data values that are in the bucket, i.e. are between
	//  min and max values.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Float64Stats.HistogramBucket.count
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.StringStats
type StringStats struct {
	// The statistics of the top 20 unigrams, ordered by
	//  [count][google.cloud.automl.v1beta1.StringStats.UnigramStats.count].
	// +kcc:proto:field=google.cloud.automl.v1beta1.StringStats.top_unigram_stats
	TopUnigramStats []StringStats_UnigramStats `json:"topUnigramStats,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.StringStats.UnigramStats
type StringStats_UnigramStats struct {
	// The unigram.
	// +kcc:proto:field=google.cloud.automl.v1beta1.StringStats.UnigramStats.value
	Value *string `json:"value,omitempty"`

	// The number of occurrences of this unigram in the series.
	// +kcc:proto:field=google.cloud.automl.v1beta1.StringStats.UnigramStats.count
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.StructStats
type StructStats struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.automl.v1beta1.StructType
type StructType struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.automl.v1beta1.TimestampStats
type TimestampStats struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.automl.v1beta1.TimestampStats.GranularStats
type TimestampStats_GranularStats struct {

	// TODO: unsupported map type with key int32 and value int64

}
