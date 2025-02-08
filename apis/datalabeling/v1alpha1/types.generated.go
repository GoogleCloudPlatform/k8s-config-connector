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


// +kcc:proto=google.cloud.datalabeling.v1beta1.Attempt
type Attempt struct {
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Attempt.attempt_time
	AttemptTime *string `json:"attemptTime,omitempty"`

	// Details of errors that occurred.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Attempt.partial_failures
	PartialFailures []Status `json:"partialFailures,omitempty"`
}

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

// +kcc:proto=google.cloud.datalabeling.v1beta1.BoundingBoxEvaluationOptions
type BoundingBoxEvaluationOptions struct {
	// Minimum
	//  [intersection-over-union
	//
	//  (IOU)](/vision/automl/object-detection/docs/evaluate#intersection-over-union)
	//  required for 2 bounding boxes to be considered a match. This must be a
	//  number between 0 and 1.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.BoundingBoxEvaluationOptions.iou_threshold
	IouThreshold *float32 `json:"iouThreshold,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.BoundingPolyConfig
type BoundingPolyConfig struct {
	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.BoundingPolyConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Optional. Instruction message showed on contributors UI.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.BoundingPolyConfig.instruction_message
	InstructionMessage *string `json:"instructionMessage,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ClassificationMetadata
type ClassificationMetadata struct {
	// Whether the classification task is multi-label or not.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ClassificationMetadata.is_multi_label
	IsMultiLabel *bool `json:"isMultiLabel,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.EvaluationConfig
type EvaluationConfig struct {
	// Only specify this field if the related model performs image object
	//  detection (`IMAGE_BOUNDING_BOX_ANNOTATION`). Describes how to evaluate
	//  bounding boxes.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationConfig.bounding_box_evaluation_options
	BoundingBoxEvaluationOptions *BoundingBoxEvaluationOptions `json:"boundingBoxEvaluationOptions,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.EvaluationJob
type EvaluationJob struct {
	// Output only. After you create a job, Data Labeling Service assigns a name
	//  to the job with the following format:
	//
	//  "projects/<var>{project_id}</var>/evaluationJobs/<var>{evaluation_job_id}</var>"
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.name
	Name *string `json:"name,omitempty"`

	// Required. Description of the job. The description can be up to 25,000
	//  characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.description
	Description *string `json:"description,omitempty"`

	// Output only. Describes the current state of the job.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.state
	State *string `json:"state,omitempty"`

	// Required. Describes the interval at which the job runs. This interval must
	//  be at least 1 day, and it is rounded to the nearest day. For example, if
	//  you specify a 50-hour interval, the job runs every 2 days.
	//
	//  You can provide the schedule in
	//  [crontab format](/scheduler/docs/configuring/cron-job-schedules) or in an
	//  [English-like
	//  format](/appengine/docs/standard/python/config/cronref#schedule_format).
	//
	//  Regardless of what you specify, the job will run at 10:00 AM UTC. Only the
	//  interval from this schedule is used, not the specific time of day.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Required. The [AI Platform Prediction model
	//  version](/ml-engine/docs/prediction-overview) to be evaluated. Prediction
	//  input and output is sampled from this model version. When creating an
	//  evaluation job, specify the model version in the following format:
	//
	//  "projects/<var>{project_id}</var>/models/<var>{model_name}</var>/versions/<var>{version_name}</var>"
	//
	//  There can only be one evaluation job per model version.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.model_version
	ModelVersion *string `json:"modelVersion,omitempty"`

	// Required. Configuration details for the evaluation job.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.evaluation_job_config
	EvaluationJobConfig *EvaluationJobConfig `json:"evaluationJobConfig,omitempty"`

	// Required. Name of the [AnnotationSpecSet][google.cloud.datalabeling.v1beta1.AnnotationSpecSet] describing all the
	//  labels that your machine learning model outputs. You must create this
	//  resource before you create an evaluation job and provide its name in the
	//  following format:
	//
	//  "projects/<var>{project_id}</var>/annotationSpecSets/<var>{annotation_spec_set_id}</var>"
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Required. Whether you want Data Labeling Service to provide ground truth
	//  labels for prediction input. If you want the service to assign human
	//  labelers to annotate your data, set this to `true`. If you want to provide
	//  your own ground truth labels in the evaluation job's BigQuery table, set
	//  this to `false`.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.label_missing_ground_truth
	LabelMissingGroundTruth *bool `json:"labelMissingGroundTruth,omitempty"`

	// Output only. Every time the evaluation job runs and an error occurs, the
	//  failed attempt is appended to this array.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.attempts
	Attempts []Attempt `json:"attempts,omitempty"`

	// Output only. Timestamp of when this evaluation job was created.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.EvaluationJobAlertConfig
type EvaluationJobAlertConfig struct {
	// Required. An email address to send alerts to.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobAlertConfig.email
	Email *string `json:"email,omitempty"`

	// Required. A number between 0 and 1 that describes a minimum mean average
	//  precision threshold. When the evaluation job runs, if it calculates that
	//  your model version's predictions from the recent interval have
	//  [meanAveragePrecision][google.cloud.datalabeling.v1beta1.PrCurve.mean_average_precision] below this
	//  threshold, then it sends an alert to your specified email.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobAlertConfig.min_acceptable_mean_average_precision
	MinAcceptableMeanAveragePrecision *float64 `json:"minAcceptableMeanAveragePrecision,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.EvaluationJobConfig
type EvaluationJobConfig struct {
	// Specify this field if your model version performs image classification or
	//  general classification.
	//
	//  `annotationSpecSet` in this configuration must match
	//  [EvaluationJob.annotationSpecSet][google.cloud.datalabeling.v1beta1.EvaluationJob.annotation_spec_set].
	//  `allowMultiLabel` in this configuration must match
	//  `classificationMetadata.isMultiLabel` in [input_config][google.cloud.datalabeling.v1beta1.EvaluationJobConfig.input_config].
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.image_classification_config
	ImageClassificationConfig *ImageClassificationConfig `json:"imageClassificationConfig,omitempty"`

	// Specify this field if your model version performs image object detection
	//  (bounding box detection).
	//
	//  `annotationSpecSet` in this configuration must match
	//  [EvaluationJob.annotationSpecSet][google.cloud.datalabeling.v1beta1.EvaluationJob.annotation_spec_set].
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.bounding_poly_config
	BoundingPolyConfig *BoundingPolyConfig `json:"boundingPolyConfig,omitempty"`

	// Specify this field if your model version performs text classification.
	//
	//  `annotationSpecSet` in this configuration must match
	//  [EvaluationJob.annotationSpecSet][google.cloud.datalabeling.v1beta1.EvaluationJob.annotation_spec_set].
	//  `allowMultiLabel` in this configuration must match
	//  `classificationMetadata.isMultiLabel` in [input_config][google.cloud.datalabeling.v1beta1.EvaluationJobConfig.input_config].
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.text_classification_config
	TextClassificationConfig *TextClassificationConfig `json:"textClassificationConfig,omitempty"`

	// Rquired. Details for the sampled prediction input. Within this
	//  configuration, there are requirements for several fields:
	//
	//  * `dataType` must be one of `IMAGE`, `TEXT`, or `GENERAL_DATA`.
	//  * `annotationType` must be one of `IMAGE_CLASSIFICATION_ANNOTATION`,
	//    `TEXT_CLASSIFICATION_ANNOTATION`, `GENERAL_CLASSIFICATION_ANNOTATION`,
	//    or `IMAGE_BOUNDING_BOX_ANNOTATION` (image object detection).
	//  * If your machine learning model performs classification, you must specify
	//    `classificationMetadata.isMultiLabel`.
	//  * You must specify `bigquerySource` (not `gcsSource`).
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.input_config
	InputConfig *InputConfig `json:"inputConfig,omitempty"`

	// Required. Details for calculating evaluation metrics and creating
	//  [Evaulations][google.cloud.datalabeling.v1beta1.Evaluation]. If your model version performs image object
	//  detection, you must specify the `boundingBoxEvaluationOptions` field within
	//  this configuration. Otherwise, provide an empty object for this
	//  configuration.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.evaluation_config
	EvaluationConfig *EvaluationConfig `json:"evaluationConfig,omitempty"`

	// Optional. Details for human annotation of your data. If you set
	//  [labelMissingGroundTruth][google.cloud.datalabeling.v1beta1.EvaluationJob.label_missing_ground_truth] to
	//  `true` for this evaluation job, then you must specify this field. If you
	//  plan to provide your own ground truth labels, then omit this field.
	//
	//  Note that you must create an [Instruction][google.cloud.datalabeling.v1beta1.Instruction] resource before you can
	//  specify this field. Provide the name of the instruction resource in the
	//  `instruction` field within this configuration.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.human_annotation_config
	HumanAnnotationConfig *HumanAnnotationConfig `json:"humanAnnotationConfig,omitempty"`

	// Required. Prediction keys that tell Data Labeling Service where to find the
	//  data for evaluation in your BigQuery table. When the service samples
	//  prediction input and output from your model version and saves it to
	//  BigQuery, the data gets stored as JSON strings in the BigQuery table. These
	//  keys tell Data Labeling Service how to parse the JSON.
	//
	//  You can provide the following entries in this field:
	//
	//  * `data_json_key`: the data key for prediction input. You must provide
	//    either this key or `reference_json_key`.
	//  * `reference_json_key`: the data reference key for prediction input. You
	//    must provide either this key or `data_json_key`.
	//  * `label_json_key`: the label key for prediction output. Required.
	//  * `label_score_json_key`: the score key for prediction output. Required.
	//  * `bounding_box_json_key`: the bounding box key for prediction output.
	//    Required if your model version perform image object detection.
	//
	//  Learn [how to configure prediction
	//  keys](/ml-engine/docs/continuous-evaluation/create-job#prediction-keys).
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.bigquery_import_keys
	BigqueryImportKeys map[string]string `json:"bigqueryImportKeys,omitempty"`

	// Required. The maximum number of predictions to sample and save to BigQuery
	//  during each [evaluation interval][google.cloud.datalabeling.v1beta1.EvaluationJob.schedule]. This limit
	//  overrides `example_sample_percentage`: even if the service has not sampled
	//  enough predictions to fulfill `example_sample_perecentage` during an
	//  interval, it stops sampling predictions when it meets this limit.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.example_count
	ExampleCount *int32 `json:"exampleCount,omitempty"`

	// Required. Fraction of predictions to sample and save to BigQuery during
	//  each [evaluation interval][google.cloud.datalabeling.v1beta1.EvaluationJob.schedule]. For example, 0.1 means
	//  10% of predictions served by your model version get saved to BigQuery.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.example_sample_percentage
	ExampleSamplePercentage *float64 `json:"exampleSamplePercentage,omitempty"`

	// Optional. Configuration details for evaluation job alerts. Specify this
	//  field if you want to receive email alerts if the evaluation job finds that
	//  your predictions have low mean average precision during a run.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJobConfig.evaluation_job_alert_config
	EvaluationJobAlertConfig *EvaluationJobAlertConfig `json:"evaluationJobAlertConfig,omitempty"`
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

// +kcc:proto=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig
type HumanAnnotationConfig struct {
	// Required. Instruction resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.instruction
	Instruction *string `json:"instruction,omitempty"`

	// Required. A human-readable name for AnnotatedDataset defined by
	//  users. Maximum of 64 characters
	//  .
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.annotated_dataset_display_name
	AnnotatedDatasetDisplayName *string `json:"annotatedDatasetDisplayName,omitempty"`

	// Optional. A human-readable description for AnnotatedDataset.
	//  The description can be up to 10000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.annotated_dataset_description
	AnnotatedDatasetDescription *string `json:"annotatedDatasetDescription,omitempty"`

	// Optional. A human-readable label used to logically group labeling tasks.
	//  This string must match the regular expression `[a-zA-Z\\d_-]{0,128}`.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.label_group
	LabelGroup *string `json:"labelGroup,omitempty"`

	// Optional. The Language of this question, as a
	//  [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	//  Default value is en-US.
	//  Only need to set this when task is language related. For example, French
	//  text classification.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Optional. Replication of questions. Each question will be sent to up to
	//  this number of contributors to label. Aggregated answers will be returned.
	//  Default is set to 1.
	//  For image related labeling, valid values are 1, 3, 5.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.replica_count
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Optional. Maximum duration for contributors to answer a question. Maximum
	//  is 3600 seconds. Default is 3600 seconds.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.question_duration
	QuestionDuration *string `json:"questionDuration,omitempty"`

	// Optional. If you want your own labeling contributors to manage and work on
	//  this labeling request, you can set these contributors here. We will give
	//  them access to the question types in crowdcompute. Note that these
	//  emails must be registered in crowdcompute worker UI:
	//  https://crowd-compute.appspot.com/
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.contributor_emails
	ContributorEmails []string `json:"contributorEmails,omitempty"`

	// Email of the user who started the labeling task and should be notified by
	//  email. If empty no notification will be sent.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.HumanAnnotationConfig.user_email_address
	UserEmailAddress *string `json:"userEmailAddress,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ImageClassificationConfig
type ImageClassificationConfig struct {
	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageClassificationConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Optional. If allow_multi_label is true, contributors are able to choose
	//  multiple labels for one image.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageClassificationConfig.allow_multi_label
	AllowMultiLabel *bool `json:"allowMultiLabel,omitempty"`

	// Optional. The type of how to aggregate answers.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageClassificationConfig.answer_aggregation_type
	AnswerAggregationType *string `json:"answerAggregationType,omitempty"`
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

// +kcc:proto=google.cloud.datalabeling.v1beta1.SentimentConfig
type SentimentConfig struct {
	// If set to true, contributors will have the option to select sentiment of
	//  the label they selected, to mark it as negative or positive label. Default
	//  is false.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.SentimentConfig.enable_label_sentiment_selection
	EnableLabelSentimentSelection *bool `json:"enableLabelSentimentSelection,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextClassificationConfig
type TextClassificationConfig struct {
	// Optional. If allow_multi_label is true, contributors are able to choose
	//  multiple labels for one text segment.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextClassificationConfig.allow_multi_label
	AllowMultiLabel *bool `json:"allowMultiLabel,omitempty"`

	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextClassificationConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Optional. Configs for sentiment selection.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextClassificationConfig.sentiment_config
	SentimentConfig *SentimentConfig `json:"sentimentConfig,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextMetadata
type TextMetadata struct {
	// The language of this text, as a
	//  [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	//  Default value is en-US.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextMetadata.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}
