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

// +kcc:proto=google.cloud.automl.v1beta1.ImageClassificationModelMetadata
type ImageClassificationModelMetadata struct {
	// Optional. The ID of the `base` model. If it is specified, the new model
	//  will be created based on the `base` model. Otherwise, the new model will be
	//  created from scratch. The `base` model must be in the same
	//  `project` and `location` as the new model to create, and have the same
	//  `model_type`.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationModelMetadata.base_model_id
	BaseModelID *string `json:"baseModelID,omitempty"`

	// Required. The train budget of creating this model, expressed in hours. The
	//  actual `train_cost` will be equal or less than this value.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationModelMetadata.train_budget
	TrainBudget *int64 `json:"trainBudget,omitempty"`

	// Output only. The actual train cost of creating this model, expressed in
	//  hours. If this model is created from a `base` model, the train cost used
	//  to create the `base` model are not included.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationModelMetadata.train_cost
	TrainCost *int64 `json:"trainCost,omitempty"`

	// Output only. The reason that this create model operation stopped,
	//  e.g. `BUDGET_REACHED`, `MODEL_CONVERGED`.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationModelMetadata.stop_reason
	StopReason *string `json:"stopReason,omitempty"`

	// Optional. Type of the model. The available values are:
	//  *   `cloud` - Model to be used via prediction calls to AutoML API.
	//                This is the default value.
	//  *   `mobile-low-latency-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards. Expected to have low latency, but
	//                may have lower prediction quality than other models.
	//  *   `mobile-versatile-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.
	//  *   `mobile-high-accuracy-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.  Expected to have a higher
	//                latency, but should also have a higher prediction quality
	//                than other models.
	//  *   `mobile-core-ml-low-latency-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile device with Core
	//                ML afterwards. Expected to have low latency, but may have
	//                lower prediction quality than other models.
	//  *   `mobile-core-ml-versatile-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile device with Core
	//                ML afterwards.
	//  *   `mobile-core-ml-high-accuracy-1` - A model that, in addition to
	//                providing prediction via AutoML API, can also be exported
	//                (see [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile device with
	//                Core ML afterwards.  Expected to have a higher latency, but
	//                should also have a higher prediction quality than other
	//                models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationModelMetadata.model_type
	ModelType *string `json:"modelType,omitempty"`

	// Output only. An approximate number of online prediction QPS that can
	//  be supported by this model per each node on which it is deployed.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationModelMetadata.node_qps
	NodeQps *float64 `json:"nodeQps,omitempty"`

	// Output only. The number of nodes this model is deployed on. A node is an
	//  abstraction of a machine resource, which can handle online prediction QPS
	//  as given in the node_qps field.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationModelMetadata.node_count
	NodeCount *int64 `json:"nodeCount,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata
type ImageObjectDetectionModelMetadata struct {
	// Optional. Type of the model. The available values are:
	//  *   `cloud-high-accuracy-1` - (default) A model to be used via prediction
	//                calls to AutoML API. Expected to have a higher latency, but
	//                should also have a higher prediction quality than other
	//                models.
	//  *   `cloud-low-latency-1` -  A model to be used via prediction
	//                calls to AutoML API. Expected to have low latency, but may
	//                have lower prediction quality than other models.
	//  *   `mobile-low-latency-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards. Expected to have low latency, but
	//                may have lower prediction quality than other models.
	//  *   `mobile-versatile-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.
	//  *   `mobile-high-accuracy-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1beta1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.  Expected to have a higher
	//                latency, but should also have a higher prediction quality
	//                than other models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata.model_type
	ModelType *string `json:"modelType,omitempty"`

	// Output only. The number of nodes this model is deployed on. A node is an
	//  abstraction of a machine resource, which can handle online prediction QPS
	//  as given in the qps_per_node field.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata.node_count
	NodeCount *int64 `json:"nodeCount,omitempty"`

	// Output only. An approximate number of online prediction QPS that can
	//  be supported by this model per each node on which it is deployed.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata.node_qps
	NodeQps *float64 `json:"nodeQps,omitempty"`

	// Output only. The reason that this create model operation stopped,
	//  e.g. `BUDGET_REACHED`, `MODEL_CONVERGED`.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata.stop_reason
	StopReason *string `json:"stopReason,omitempty"`

	// The train budget of creating this model, expressed in milli node
	//  hours i.e. 1,000 value in this field means 1 node hour. The actual
	//  `train_cost` will be equal or less than this value. If further model
	//  training ceases to provide any improvements, it will stop without using
	//  full budget and the stop_reason will be `MODEL_CONVERGED`.
	//  Note, node_hour  = actual_hour * number_of_nodes_invovled.
	//  For model type `cloud-high-accuracy-1`(default) and `cloud-low-latency-1`,
	//  the train budget must be between 20,000 and 900,000 milli node hours,
	//  inclusive. The default value is 216, 000 which represents one day in
	//  wall time.
	//  For model type `mobile-low-latency-1`, `mobile-versatile-1`,
	//  `mobile-high-accuracy-1`, `mobile-core-ml-low-latency-1`,
	//  `mobile-core-ml-versatile-1`, `mobile-core-ml-high-accuracy-1`, the train
	//  budget must be between 1,000 and 100,000 milli node hours, inclusive.
	//  The default value is 24, 000 which represents one day in wall time.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata.train_budget_milli_node_hours
	TrainBudgetMilliNodeHours *int64 `json:"trainBudgetMilliNodeHours,omitempty"`

	// Output only. The actual train cost of creating this model, expressed in
	//  milli node hours, i.e. 1,000 value in this field means 1 node hour.
	//  Guaranteed to not exceed the train budget.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata.train_cost_milli_node_hours
	TrainCostMilliNodeHours *int64 `json:"trainCostMilliNodeHours,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.Model
type Model struct {
	// Metadata for translation models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.translation_model_metadata
	TranslationModelMetadata *TranslationModelMetadata `json:"translationModelMetadata,omitempty"`

	// Metadata for image classification models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.image_classification_model_metadata
	ImageClassificationModelMetadata *ImageClassificationModelMetadata `json:"imageClassificationModelMetadata,omitempty"`

	// Metadata for text classification models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.text_classification_model_metadata
	TextClassificationModelMetadata *TextClassificationModelMetadata `json:"textClassificationModelMetadata,omitempty"`

	// Metadata for image object detection models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.image_object_detection_model_metadata
	ImageObjectDetectionModelMetadata *ImageObjectDetectionModelMetadata `json:"imageObjectDetectionModelMetadata,omitempty"`

	// Metadata for video classification models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.video_classification_model_metadata
	VideoClassificationModelMetadata *VideoClassificationModelMetadata `json:"videoClassificationModelMetadata,omitempty"`

	// Metadata for video object tracking models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.video_object_tracking_model_metadata
	VideoObjectTrackingModelMetadata *VideoObjectTrackingModelMetadata `json:"videoObjectTrackingModelMetadata,omitempty"`

	// Metadata for text extraction models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.text_extraction_model_metadata
	TextExtractionModelMetadata *TextExtractionModelMetadata `json:"textExtractionModelMetadata,omitempty"`

	// Metadata for Tables models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.tables_model_metadata
	TablesModelMetadata *TablesModelMetadata `json:"tablesModelMetadata,omitempty"`

	// Metadata for text sentiment models.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.text_sentiment_model_metadata
	TextSentimentModelMetadata *TextSentimentModelMetadata `json:"textSentimentModelMetadata,omitempty"`

	// Output only. Resource name of the model.
	//  Format: `projects/{project_id}/locations/{location_id}/models/{model_id}`
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the model to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores
	//  (_), and ASCII digits 0-9. It must start with a letter.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The resource ID of the dataset used to create the model. The dataset must
	//  come from the same ancestor project and location.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// Output only. Timestamp when the model training finished  and can be used for prediction.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this model was last updated.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Deployment state of the model. A model can only serve
	//  prediction requests after it gets deployed.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Model.deployment_state
	DeploymentState *string `json:"deploymentState,omitempty"`
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

// +kcc:proto=google.cloud.automl.v1beta1.TablesModelColumnInfo
type TablesModelColumnInfo struct {
	// Output only. The name of the ColumnSpec describing the column. Not
	//  populated when this proto is outputted to BigQuery.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelColumnInfo.column_spec_name
	ColumnSpecName *string `json:"columnSpecName,omitempty"`

	// Output only. The display name of the column (same as the display_name of
	//  its ColumnSpec).
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelColumnInfo.column_display_name
	ColumnDisplayName *string `json:"columnDisplayName,omitempty"`

	// Output only. When given as part of a Model (always populated):
	//  Measurement of how much model predictions correctness on the TEST data
	//  depend on values in this column. A value between 0 and 1, higher means
	//  higher influence. These values are normalized - for all input feature
	//  columns of a given model they add to 1.
	//
	//  When given back by Predict (populated iff
	//  [feature_importance
	//  param][google.cloud.automl.v1beta1.PredictRequest.params] is set) or Batch
	//  Predict (populated iff
	//  [feature_importance][google.cloud.automl.v1beta1.PredictRequest.params]
	//  param is set):
	//  Measurement of how impactful for the prediction returned for the given row
	//  the value in this column was. Specifically, the feature importance
	//  specifies the marginal contribution that the feature made to the prediction
	//  score compared to the baseline score. These values are computed using the
	//  Sampled Shapley method.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelColumnInfo.feature_importance
	FeatureImportance *float32 `json:"featureImportance,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TablesModelMetadata
type TablesModelMetadata struct {
	// Required when optimization_objective is "MAXIMIZE_PRECISION_AT_RECALL".
	//  Must be between 0 and 1, inclusive.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.optimization_objective_recall_value
	OptimizationObjectiveRecallValue *float32 `json:"optimizationObjectiveRecallValue,omitempty"`

	// Required when optimization_objective is "MAXIMIZE_RECALL_AT_PRECISION".
	//  Must be between 0 and 1, inclusive.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.optimization_objective_precision_value
	OptimizationObjectivePrecisionValue *float32 `json:"optimizationObjectivePrecisionValue,omitempty"`

	// Column spec of the dataset's primary table's column the model is
	//  predicting. Snapshotted when model creation started.
	//  Only 3 fields are used:
	//  name - May be set on CreateModel, if it's not then the ColumnSpec
	//         corresponding to the current target_column_spec_id of the dataset
	//         the model is trained from is used.
	//         If neither is set, CreateModel will error.
	//  display_name - Output only.
	//  data_type - Output only.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.target_column_spec
	TargetColumnSpec *ColumnSpec `json:"targetColumnSpec,omitempty"`

	// Column specs of the dataset's primary table's columns, on which
	//  the model is trained and which are used as the input for predictions.
	//  The
	//
	//  [target_column][google.cloud.automl.v1beta1.TablesModelMetadata.target_column_spec]
	//  as well as, according to dataset's state upon model creation,
	//
	//  [weight_column][google.cloud.automl.v1beta1.TablesDatasetMetadata.weight_column_spec_id],
	//  and
	//
	//  [ml_use_column][google.cloud.automl.v1beta1.TablesDatasetMetadata.ml_use_column_spec_id]
	//  must never be included here.
	//
	//  Only 3 fields are used:
	//
	//  * name - May be set on CreateModel, if set only the columns specified are
	//    used, otherwise all primary table's columns (except the ones listed
	//    above) are used for the training and prediction input.
	//
	//  * display_name - Output only.
	//
	//  * data_type - Output only.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.input_feature_column_specs
	InputFeatureColumnSpecs []ColumnSpec `json:"inputFeatureColumnSpecs,omitempty"`

	// Objective function the model is optimizing towards. The training process
	//  creates a model that maximizes/minimizes the value of the objective
	//  function over the validation set.
	//
	//  The supported optimization objectives depend on the prediction type.
	//  If the field is not set, a default objective function is used.
	//
	//  CLASSIFICATION_BINARY:
	//    "MAXIMIZE_AU_ROC" (default) - Maximize the area under the receiver
	//                                  operating characteristic (ROC) curve.
	//    "MINIMIZE_LOG_LOSS" - Minimize log loss.
	//    "MAXIMIZE_AU_PRC" - Maximize the area under the precision-recall curve.
	//    "MAXIMIZE_PRECISION_AT_RECALL" - Maximize precision for a specified
	//                                    recall value.
	//    "MAXIMIZE_RECALL_AT_PRECISION" - Maximize recall for a specified
	//                                     precision value.
	//
	//  CLASSIFICATION_MULTI_CLASS :
	//    "MINIMIZE_LOG_LOSS" (default) - Minimize log loss.
	//
	//
	//  REGRESSION:
	//    "MINIMIZE_RMSE" (default) - Minimize root-mean-squared error (RMSE).
	//    "MINIMIZE_MAE" - Minimize mean-absolute error (MAE).
	//    "MINIMIZE_RMSLE" - Minimize root-mean-squared log error (RMSLE).
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.optimization_objective
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// Output only. Auxiliary information for each of the
	//  input_feature_column_specs with respect to this particular model.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.tables_model_column_info
	TablesModelColumnInfo []TablesModelColumnInfo `json:"tablesModelColumnInfo,omitempty"`

	// Required. The train budget of creating this model, expressed in milli node
	//  hours i.e. 1,000 value in this field means 1 node hour.
	//
	//  The training cost of the model will not exceed this budget. The final cost
	//  will be attempted to be close to the budget, though may end up being (even)
	//  noticeably smaller - at the backend's discretion. This especially may
	//  happen when further model training ceases to provide any improvements.
	//
	//  If the budget is set to a value known to be insufficient to train a
	//  model for the given dataset, the training won't be attempted and
	//  will error.
	//
	//  The train budget must be between 1,000 and 72,000 milli node hours,
	//  inclusive.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.train_budget_milli_node_hours
	TrainBudgetMilliNodeHours *int64 `json:"trainBudgetMilliNodeHours,omitempty"`

	// Output only. The actual training cost of the model, expressed in milli
	//  node hours, i.e. 1,000 value in this field means 1 node hour. Guaranteed
	//  to not exceed the train budget.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.train_cost_milli_node_hours
	TrainCostMilliNodeHours *int64 `json:"trainCostMilliNodeHours,omitempty"`

	// Use the entire training budget. This disables the early stopping feature.
	//  By default, the early stopping feature is enabled, which means that AutoML
	//  Tables might stop training before the entire training budget has been used.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesModelMetadata.disable_early_stopping
	DisableEarlyStopping *bool `json:"disableEarlyStopping,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TextClassificationModelMetadata
type TextClassificationModelMetadata struct {
	// Output only. Classification type of the dataset used to train this model.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TextClassificationModelMetadata.classification_type
	ClassificationType *string `json:"classificationType,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TextExtractionModelMetadata
type TextExtractionModelMetadata struct {
	// Indicates the scope of model use case.
	//
	//  * `default`: Use to train a general text extraction model. Default value.
	//
	//  * `health_care`: Use to train a text extraction model that is tuned for
	//    healthcare applications.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TextExtractionModelMetadata.model_hint
	ModelHint *string `json:"modelHint,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TextSentimentModelMetadata
type TextSentimentModelMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1beta1.TimestampStats
type TimestampStats struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.automl.v1beta1.TimestampStats.GranularStats
type TimestampStats_GranularStats struct {

	// TODO: unsupported map type with key int32 and value int64

}

// +kcc:proto=google.cloud.automl.v1beta1.TranslationModelMetadata
type TranslationModelMetadata struct {
	// The resource name of the model to use as a baseline to train the custom
	//  model. If unset, we use the default base model provided by Google
	//  Translate. Format:
	//  `projects/{project_id}/locations/{location_id}/models/{model_id}`
	// +kcc:proto:field=google.cloud.automl.v1beta1.TranslationModelMetadata.base_model
	BaseModel *string `json:"baseModel,omitempty"`

	// Output only. Inferred from the dataset.
	//  The source languge (The BCP-47 language code) that is used for training.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TranslationModelMetadata.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// Output only. The target languge (The BCP-47 language code) that is used for
	//  training.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TranslationModelMetadata.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.VideoClassificationModelMetadata
type VideoClassificationModelMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1beta1.VideoObjectTrackingModelMetadata
type VideoObjectTrackingModelMetadata struct {
}
