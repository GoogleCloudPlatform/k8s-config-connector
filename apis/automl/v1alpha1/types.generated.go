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


// +kcc:proto=google.cloud.automl.v1.ImageClassificationModelMetadata
type ImageClassificationModelMetadata struct {
	// Optional. The ID of the `base` model. If it is specified, the new model
	//  will be created based on the `base` model. Otherwise, the new model will be
	//  created from scratch. The `base` model must be in the same
	//  `project` and `location` as the new model to create, and have the same
	//  `model_type`.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationModelMetadata.base_model_id
	BaseModelID *string `json:"baseModelID,omitempty"`

	// Optional. The train budget of creating this model, expressed in milli node
	//  hours i.e. 1,000 value in this field means 1 node hour. The actual
	//  `train_cost` will be equal or less than this value. If further model
	//  training ceases to provide any improvements, it will stop without using
	//  full budget and the stop_reason will be `MODEL_CONVERGED`.
	//  Note, node_hour  = actual_hour * number_of_nodes_invovled.
	//  For model type `cloud`(default), the train budget must be between 8,000
	//  and 800,000 milli node hours, inclusive. The default value is 192, 000
	//  which represents one day in wall time. For model type
	//  `mobile-low-latency-1`, `mobile-versatile-1`, `mobile-high-accuracy-1`,
	//  `mobile-core-ml-low-latency-1`, `mobile-core-ml-versatile-1`,
	//  `mobile-core-ml-high-accuracy-1`, the train budget must be between 1,000
	//  and 100,000 milli node hours, inclusive. The default value is 24, 000 which
	//  represents one day in wall time.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationModelMetadata.train_budget_milli_node_hours
	TrainBudgetMilliNodeHours *int64 `json:"trainBudgetMilliNodeHours,omitempty"`

	// Optional. Type of the model. The available values are:
	//  *   `cloud` - Model to be used via prediction calls to AutoML API.
	//                This is the default value.
	//  *   `mobile-low-latency-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards. Expected to have low latency, but
	//                may have lower prediction quality than other models.
	//  *   `mobile-versatile-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.
	//  *   `mobile-high-accuracy-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.  Expected to have a higher
	//                latency, but should also have a higher prediction quality
	//                than other models.
	//  *   `mobile-core-ml-low-latency-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile device with Core
	//                ML afterwards. Expected to have low latency, but may have
	//                lower prediction quality than other models.
	//  *   `mobile-core-ml-versatile-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile device with Core
	//                ML afterwards.
	//  *   `mobile-core-ml-high-accuracy-1` - A model that, in addition to
	//                providing prediction via AutoML API, can also be exported
	//                (see [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile device with
	//                Core ML afterwards.  Expected to have a higher latency, but
	//                should also have a higher prediction quality than other
	//                models.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationModelMetadata.model_type
	ModelType *string `json:"modelType,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ImageObjectDetectionModelMetadata
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
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards. Expected to have low latency, but
	//                may have lower prediction quality than other models.
	//  *   `mobile-versatile-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.
	//  *   `mobile-high-accuracy-1` - A model that, in addition to providing
	//                prediction via AutoML API, can also be exported (see
	//                [AutoMl.ExportModel][google.cloud.automl.v1.AutoMl.ExportModel]) and used on a mobile or edge device
	//                with TensorFlow afterwards.  Expected to have a higher
	//                latency, but should also have a higher prediction quality
	//                than other models.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionModelMetadata.model_type
	ModelType *string `json:"modelType,omitempty"`

	// Optional. The train budget of creating this model, expressed in milli node
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
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionModelMetadata.train_budget_milli_node_hours
	TrainBudgetMilliNodeHours *int64 `json:"trainBudgetMilliNodeHours,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.Model
type Model struct {
	// Metadata for translation models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.translation_model_metadata
	TranslationModelMetadata *TranslationModelMetadata `json:"translationModelMetadata,omitempty"`

	// Metadata for image classification models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_classification_model_metadata
	ImageClassificationModelMetadata *ImageClassificationModelMetadata `json:"imageClassificationModelMetadata,omitempty"`

	// Metadata for text classification models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.text_classification_model_metadata
	TextClassificationModelMetadata *TextClassificationModelMetadata `json:"textClassificationModelMetadata,omitempty"`

	// Metadata for image object detection models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_object_detection_model_metadata
	ImageObjectDetectionModelMetadata *ImageObjectDetectionModelMetadata `json:"imageObjectDetectionModelMetadata,omitempty"`

	// Metadata for text extraction models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.text_extraction_model_metadata
	TextExtractionModelMetadata *TextExtractionModelMetadata `json:"textExtractionModelMetadata,omitempty"`

	// Metadata for text sentiment models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.text_sentiment_model_metadata
	TextSentimentModelMetadata *TextSentimentModelMetadata `json:"textSentimentModelMetadata,omitempty"`

	// Output only. Resource name of the model.
	//  Format: `projects/{project_id}/locations/{location_id}/models/{model_id}`
	// +kcc:proto:field=google.cloud.automl.v1.Model.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the model to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores
	//  (_), and ASCII digits 0-9. It must start with a letter.
	// +kcc:proto:field=google.cloud.automl.v1.Model.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The resource ID of the dataset used to create the model. The dataset must
	//  come from the same ancestor project and location.
	// +kcc:proto:field=google.cloud.automl.v1.Model.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// Output only. Timestamp when the model training finished  and can be used for prediction.
	// +kcc:proto:field=google.cloud.automl.v1.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this model was last updated.
	// +kcc:proto:field=google.cloud.automl.v1.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Deployment state of the model. A model can only serve
	//  prediction requests after it gets deployed.
	// +kcc:proto:field=google.cloud.automl.v1.Model.deployment_state
	DeploymentState *string `json:"deploymentState,omitempty"`

	// Used to perform a consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.automl.v1.Model.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your model.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  Label values are optional. Label keys must start with a letter.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	// +kcc:proto:field=google.cloud.automl.v1.Model.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TextClassificationModelMetadata
type TextClassificationModelMetadata struct {
	// Output only. Classification type of the dataset used to train this model.
	// +kcc:proto:field=google.cloud.automl.v1.TextClassificationModelMetadata.classification_type
	ClassificationType *string `json:"classificationType,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TextExtractionModelMetadata
type TextExtractionModelMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1.TextSentimentModelMetadata
type TextSentimentModelMetadata struct {
}

// +kcc:proto=google.cloud.automl.v1.TranslationModelMetadata
type TranslationModelMetadata struct {
	// The resource name of the model to use as a baseline to train the custom
	//  model. If unset, we use the default base model provided by Google
	//  Translate. Format:
	//  `projects/{project_id}/locations/{location_id}/models/{model_id}`
	// +kcc:proto:field=google.cloud.automl.v1.TranslationModelMetadata.base_model
	BaseModel *string `json:"baseModel,omitempty"`

	// Output only. Inferred from the dataset.
	//  The source language (The BCP-47 language code) that is used for training.
	// +kcc:proto:field=google.cloud.automl.v1.TranslationModelMetadata.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// Output only. The target language (The BCP-47 language code) that is used
	//  for training.
	// +kcc:proto:field=google.cloud.automl.v1.TranslationModelMetadata.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ImageClassificationModelMetadata
type ImageClassificationModelMetadataObservedState struct {
	// Output only. The actual train cost of creating this model, expressed in
	//  milli node hours, i.e. 1,000 value in this field means 1 node hour.
	//  Guaranteed to not exceed the train budget.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationModelMetadata.train_cost_milli_node_hours
	TrainCostMilliNodeHours *int64 `json:"trainCostMilliNodeHours,omitempty"`

	// Output only. The reason that this create model operation stopped,
	//  e.g. `BUDGET_REACHED`, `MODEL_CONVERGED`.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationModelMetadata.stop_reason
	StopReason *string `json:"stopReason,omitempty"`

	// Output only. An approximate number of online prediction QPS that can
	//  be supported by this model per each node on which it is deployed.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationModelMetadata.node_qps
	NodeQps *float64 `json:"nodeQps,omitempty"`

	// Output only. The number of nodes this model is deployed on. A node is an
	//  abstraction of a machine resource, which can handle online prediction QPS
	//  as given in the node_qps field.
	// +kcc:proto:field=google.cloud.automl.v1.ImageClassificationModelMetadata.node_count
	NodeCount *int64 `json:"nodeCount,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ImageObjectDetectionModelMetadata
type ImageObjectDetectionModelMetadataObservedState struct {
	// Output only. The number of nodes this model is deployed on. A node is an
	//  abstraction of a machine resource, which can handle online prediction QPS
	//  as given in the qps_per_node field.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionModelMetadata.node_count
	NodeCount *int64 `json:"nodeCount,omitempty"`

	// Output only. An approximate number of online prediction QPS that can
	//  be supported by this model per each node on which it is deployed.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionModelMetadata.node_qps
	NodeQps *float64 `json:"nodeQps,omitempty"`

	// Output only. The reason that this create model operation stopped,
	//  e.g. `BUDGET_REACHED`, `MODEL_CONVERGED`.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionModelMetadata.stop_reason
	StopReason *string `json:"stopReason,omitempty"`

	// Output only. The actual train cost of creating this model, expressed in
	//  milli node hours, i.e. 1,000 value in this field means 1 node hour.
	//  Guaranteed to not exceed the train budget.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionModelMetadata.train_cost_milli_node_hours
	TrainCostMilliNodeHours *int64 `json:"trainCostMilliNodeHours,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.Model
type ModelObservedState struct {
	// Metadata for image classification models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_classification_model_metadata
	ImageClassificationModelMetadata *ImageClassificationModelMetadataObservedState `json:"imageClassificationModelMetadata,omitempty"`

	// Metadata for image object detection models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_object_detection_model_metadata
	ImageObjectDetectionModelMetadata *ImageObjectDetectionModelMetadataObservedState `json:"imageObjectDetectionModelMetadata,omitempty"`
}
