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


// +kcc:proto=google.cloud.automl.v1.Dataset
type Dataset struct {
	// Metadata for a dataset used for translation.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.translation_dataset_metadata
	TranslationDatasetMetadata *TranslationDatasetMetadata `json:"translationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for image classification.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.image_classification_dataset_metadata
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `json:"imageClassificationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text classification.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.text_classification_dataset_metadata
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `json:"textClassificationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for image object detection.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.image_object_detection_dataset_metadata
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `json:"imageObjectDetectionDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text extraction.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.text_extraction_dataset_metadata
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `json:"textExtractionDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text sentiment.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.text_sentiment_dataset_metadata
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `json:"textSentimentDatasetMetadata,omitempty"`

	// Output only. The resource name of the dataset.
	//  Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the dataset to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores
	//  (_), and ASCII digits 0-9.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// User-provided description of the dataset. The description can be up to
	//  25000 characters long.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.description
	Description *string `json:"description,omitempty"`

	// Output only. The number of examples in the dataset.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.example_count
	ExampleCount *int32 `json:"exampleCount,omitempty"`

	// Output only. Timestamp when this dataset was created.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your dataset.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  Label values are optional. Label keys must start with a letter.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	// +kcc:proto:field=google.cloud.automl.v1.Dataset.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ImageClassificationDatasetMetadata
type ImageClassificationDatasetMetadata struct {
	// Required. Type of the classification problem.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationDatasetMetadata.classification_type
	ClassificationType *string `json:"classificationType,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ImageObjectDetectionDatasetMetadata
type ImageObjectDetectionDatasetMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1.TextClassificationDatasetMetadata
type TextClassificationDatasetMetadata struct {
	// Required. Type of the classification problem.
	// +kcc:proto:field=google.cloud.automl.v1.TextClassificationDatasetMetadata.classification_type
	ClassificationType *string `json:"classificationType,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TextExtractionDatasetMetadata
type TextExtractionDatasetMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1.TextSentimentDatasetMetadata
type TextSentimentDatasetMetadata struct {
	// Required. A sentiment is expressed as an integer ordinal, where higher value
	//  means a more positive sentiment. The range of sentiments that will be used
	//  is between 0 and sentiment_max (inclusive on both ends), and all the values
	//  in the range must be represented in the dataset before a model can be
	//  created.
	//  sentiment_max value must be between 1 and 10 (inclusive).
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentDatasetMetadata.sentiment_max
	SentimentMax *int32 `json:"sentimentMax,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TranslationDatasetMetadata
type TranslationDatasetMetadata struct {
	// Required. The BCP-47 language code of the source language.
	// +kcc:proto:field=google.cloud.automl.v1.TranslationDatasetMetadata.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// Required. The BCP-47 language code of the target language.
	// +kcc:proto:field=google.cloud.automl.v1.TranslationDatasetMetadata.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}
