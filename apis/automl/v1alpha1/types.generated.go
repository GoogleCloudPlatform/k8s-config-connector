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


// +kcc:proto=google.cloud.automl.v1beta1.CorrelationStats
type CorrelationStats struct {
	// The correlation value using the Cramer's V measure.
	// +kcc:proto:field=google.cloud.automl.v1beta1.CorrelationStats.cramers_v
	CramersV *float64 `json:"cramersV,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.Dataset
type Dataset struct {
	// Metadata for a dataset used for translation.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.translation_dataset_metadata
	TranslationDatasetMetadata *TranslationDatasetMetadata `json:"translationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for image classification.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.image_classification_dataset_metadata
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `json:"imageClassificationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text classification.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.text_classification_dataset_metadata
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `json:"textClassificationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for image object detection.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.image_object_detection_dataset_metadata
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `json:"imageObjectDetectionDatasetMetadata,omitempty"`

	// Metadata for a dataset used for video classification.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.video_classification_dataset_metadata
	VideoClassificationDatasetMetadata *VideoClassificationDatasetMetadata `json:"videoClassificationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for video object tracking.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.video_object_tracking_dataset_metadata
	VideoObjectTrackingDatasetMetadata *VideoObjectTrackingDatasetMetadata `json:"videoObjectTrackingDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text extraction.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.text_extraction_dataset_metadata
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `json:"textExtractionDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text sentiment.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.text_sentiment_dataset_metadata
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `json:"textSentimentDatasetMetadata,omitempty"`

	// Metadata for a dataset used for Tables.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.tables_dataset_metadata
	TablesDatasetMetadata *TablesDatasetMetadata `json:"tablesDatasetMetadata,omitempty"`

	// Output only. The resource name of the dataset.
	//  Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the dataset to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores
	//  (_), and ASCII digits 0-9.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// User-provided description of the dataset. The description can be up to
	//  25000 characters long.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.description
	Description *string `json:"description,omitempty"`

	// Output only. The number of examples in the dataset.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.example_count
	ExampleCount *int32 `json:"exampleCount,omitempty"`

	// Output only. Timestamp when this dataset was created.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.automl.v1beta1.Dataset.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.ImageClassificationDatasetMetadata
type ImageClassificationDatasetMetadata struct {
	// Required. Type of the classification problem.
	// +kcc:proto:field=google.cloud.automl.v1beta1.ImageClassificationDatasetMetadata.classification_type
	ClassificationType *string `json:"classificationType,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.ImageObjectDetectionDatasetMetadata
type ImageObjectDetectionDatasetMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1beta1.TablesDatasetMetadata
type TablesDatasetMetadata struct {
	// Output only. The table_spec_id of the primary table of this dataset.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesDatasetMetadata.primary_table_spec_id
	PrimaryTableSpecID *string `json:"primaryTableSpecID,omitempty"`

	// column_spec_id of the primary table's column that should be used as the
	//  training & prediction target.
	//  This column must be non-nullable and have one of following data types
	//  (otherwise model creation will error):
	//
	//  * CATEGORY
	//
	//  * FLOAT64
	//
	//  If the type is CATEGORY , only up to
	//  100 unique values may exist in that column across all rows.
	//
	//  NOTE: Updates of this field will instantly affect any other users
	//  concurrently working with the dataset.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesDatasetMetadata.target_column_spec_id
	TargetColumnSpecID *string `json:"targetColumnSpecID,omitempty"`

	// column_spec_id of the primary table's column that should be used as the
	//  weight column, i.e. the higher the value the more important the row will be
	//  during model training.
	//  Required type: FLOAT64.
	//  Allowed values: 0 to 10000, inclusive on both ends; 0 means the row is
	//                  ignored for training.
	//  If not set all rows are assumed to have equal weight of 1.
	//  NOTE: Updates of this field will instantly affect any other users
	//  concurrently working with the dataset.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesDatasetMetadata.weight_column_spec_id
	WeightColumnSpecID *string `json:"weightColumnSpecID,omitempty"`

	// column_spec_id of the primary table column which specifies a possible ML
	//  use of the row, i.e. the column will be used to split the rows into TRAIN,
	//  VALIDATE and TEST sets.
	//  Required type: STRING.
	//  This column, if set, must either have all of `TRAIN`, `VALIDATE`, `TEST`
	//  among its values, or only have `TEST`, `UNASSIGNED` values. In the latter
	//  case the rows with `UNASSIGNED` value will be assigned by AutoML. Note
	//  that if a given ml use distribution makes it impossible to create a "good"
	//  model, that call will error describing the issue.
	//  If both this column_spec_id and primary table's time_column_spec_id are not
	//  set, then all rows are treated as `UNASSIGNED`.
	//  NOTE: Updates of this field will instantly affect any other users
	//  concurrently working with the dataset.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesDatasetMetadata.ml_use_column_spec_id
	MlUseColumnSpecID *string `json:"mlUseColumnSpecID,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. The most recent timestamp when target_column_correlations
	//  field and all descendant ColumnSpec.data_stats and
	//  ColumnSpec.top_correlated_columns fields were last (re-)generated. Any
	//  changes that happened to the dataset afterwards are not reflected in these
	//  fields values. The regeneration happens in the background on a best effort
	//  basis.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TablesDatasetMetadata.stats_update_time
	StatsUpdateTime *string `json:"statsUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TextClassificationDatasetMetadata
type TextClassificationDatasetMetadata struct {
	// Required. Type of the classification problem.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TextClassificationDatasetMetadata.classification_type
	ClassificationType *string `json:"classificationType,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TextExtractionDatasetMetadata
type TextExtractionDatasetMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1beta1.TextSentimentDatasetMetadata
type TextSentimentDatasetMetadata struct {
	// Required. A sentiment is expressed as an integer ordinal, where higher value
	//  means a more positive sentiment. The range of sentiments that will be used
	//  is between 0 and sentiment_max (inclusive on both ends), and all the values
	//  in the range must be represented in the dataset before a model can be
	//  created.
	//  sentiment_max value must be between 1 and 10 (inclusive).
	// +kcc:proto:field=google.cloud.automl.v1beta1.TextSentimentDatasetMetadata.sentiment_max
	SentimentMax *int32 `json:"sentimentMax,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TranslationDatasetMetadata
type TranslationDatasetMetadata struct {
	// Required. The BCP-47 language code of the source language.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TranslationDatasetMetadata.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// Required. The BCP-47 language code of the target language.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TranslationDatasetMetadata.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.VideoClassificationDatasetMetadata
type VideoClassificationDatasetMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1beta1.VideoObjectTrackingDatasetMetadata
type VideoObjectTrackingDatasetMetadata struct {
}
