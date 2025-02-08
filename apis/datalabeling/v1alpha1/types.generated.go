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


// +kcc:proto=google.cloud.datalabeling.v1beta1.BigQuerySource
type BigQuerySource struct {
	// Required. BigQuery URI to a table, up to 2,000 characters long. If you
	//  specify the URI of a table that does not exist, Data Labeling Service
	//  creates a table at the URI with the correct schema when you create your
	//  [EvaluationJob][google.cloud.datalabeling.v1beta1.EvaluationJob]. If you specify the URI of a table that already exists,
	//  it must have the
	//  [correct
	//  schema](/ml-engine/docs/continuous-evaluation/create-job#table-schema).
	//
	//  Provide the table URI in the following format:
	//
	//  "bq://<var>{your_project_id}</var>/<var>{your_dataset_name}</var>/<var>{your_table_name}</var>"
	//
	//  [Learn
	//  more](/ml-engine/docs/continuous-evaluation/create-job#table-schema).
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.BigQuerySource.input_uri
	InputURI *string `json:"inputURI,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ClassificationMetadata
type ClassificationMetadata struct {
	// Whether the classification task is multi-label or not.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ClassificationMetadata.is_multi_label
	IsMultiLabel *bool `json:"isMultiLabel,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.Dataset
type Dataset struct {
	// Output only. Dataset resource name, format is:
	//  projects/{project_id}/datasets/{dataset_id}
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Dataset.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the dataset. Maximum of 64 characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Dataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the annotation specification set.
	//  The description can be up to 10000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Dataset.description
	Description *string `json:"description,omitempty"`

	// Output only. Time the dataset is created.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. This is populated with the original input configs
	//  where ImportData is called. It is available only after the clients
	//  import data to this dataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Dataset.input_configs
	InputConfigs []InputConfig `json:"inputConfigs,omitempty"`

	// Output only. The names of any related resources that are blocking changes
	//  to the dataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Dataset.blocking_resources
	BlockingResources []string `json:"blockingResources,omitempty"`

	// Output only. The number of data items in the dataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Dataset.data_item_count
	DataItemCount *int64 `json:"dataItemCount,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.GcsSource
type GcsSource struct {
	// Required. The input URI of source file. This must be a Cloud Storage path
	//  (`gs://...`).
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.GcsSource.input_uri
	InputURI *string `json:"inputURI,omitempty"`

	// Required. The format of the source file. Only "text/csv" is supported.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.GcsSource.mime_type
	MimeType *string `json:"mimeType,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.InputConfig
type InputConfig struct {
	// Required for text import, as language code must be specified.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.InputConfig.text_metadata
	TextMetadata *TextMetadata `json:"textMetadata,omitempty"`

	// Source located in Cloud Storage.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.InputConfig.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`

	// Source located in BigQuery. You must specify this field if you are using
	//  this InputConfig in an [EvaluationJob][google.cloud.datalabeling.v1beta1.EvaluationJob].
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.InputConfig.bigquery_source
	BigquerySource *BigQuerySource `json:"bigquerySource,omitempty"`

	// Required. Data type must be specifed when user tries to import data.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.InputConfig.data_type
	DataType *string `json:"dataType,omitempty"`

	// Optional. The type of annotation to be performed on this data. You must
	//  specify this field if you are using this InputConfig in an
	//  [EvaluationJob][google.cloud.datalabeling.v1beta1.EvaluationJob].
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.InputConfig.annotation_type
	AnnotationType *string `json:"annotationType,omitempty"`

	// Optional. Metadata about annotations for the input. You must specify this
	//  field if you are using this InputConfig in an [EvaluationJob][google.cloud.datalabeling.v1beta1.EvaluationJob] for a
	//  model version that performs classification.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.InputConfig.classification_metadata
	ClassificationMetadata *ClassificationMetadata `json:"classificationMetadata,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextMetadata
type TextMetadata struct {
	// The language of this text, as a
	//  [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	//  Default value is en-US.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextMetadata.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}
