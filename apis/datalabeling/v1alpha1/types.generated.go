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


// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotatedDataset
type AnnotatedDataset struct {
	// Output only. AnnotatedDataset resource name in format of:
	//  projects/{project_id}/datasets/{dataset_id}/annotatedDatasets/
	//  {annotated_dataset_id}
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.name
	Name *string `json:"name,omitempty"`

	// Output only. The display name of the AnnotatedDataset. It is specified in
	//  HumanAnnotationConfig when user starts a labeling task. Maximum of 64
	//  characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The description of the AnnotatedDataset. It is specified in
	//  HumanAnnotationConfig when user starts a labeling task. Maximum of 10000
	//  characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.description
	Description *string `json:"description,omitempty"`

	// Output only. Source of the annotation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.annotation_source
	AnnotationSource *string `json:"annotationSource,omitempty"`

	// Output only. Type of the annotation. It is specified when starting labeling
	//  task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.annotation_type
	AnnotationType *string `json:"annotationType,omitempty"`

	// Output only. Number of examples in the annotated dataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.example_count
	ExampleCount *int64 `json:"exampleCount,omitempty"`

	// Output only. Number of examples that have annotation in the annotated
	//  dataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.completed_example_count
	CompletedExampleCount *int64 `json:"completedExampleCount,omitempty"`

	// Output only. Per label statistics.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.label_stats
	LabelStats *LabelStats `json:"labelStats,omitempty"`

	// Output only. Time the AnnotatedDataset was created.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Additional information about AnnotatedDataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.metadata
	Metadata *AnnotatedDatasetMetadata `json:"metadata,omitempty"`

	// Output only. The names of any related resources that are blocking changes
	//  to the annotated dataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDataset.blocking_resources
	BlockingResources []string `json:"blockingResources,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata
type AnnotatedDatasetMetadata struct {
	// Configuration for image classification task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.image_classification_config
	ImageClassificationConfig *ImageClassificationConfig `json:"imageClassificationConfig,omitempty"`

	// Configuration for image bounding box and bounding poly task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.bounding_poly_config
	BoundingPolyConfig *BoundingPolyConfig `json:"boundingPolyConfig,omitempty"`

	// Configuration for image polyline task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.polyline_config
	PolylineConfig *PolylineConfig `json:"polylineConfig,omitempty"`

	// Configuration for image segmentation task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.segmentation_config
	SegmentationConfig *SegmentationConfig `json:"segmentationConfig,omitempty"`

	// Configuration for video classification task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.video_classification_config
	VideoClassificationConfig *VideoClassificationConfig `json:"videoClassificationConfig,omitempty"`

	// Configuration for video object detection task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.object_detection_config
	ObjectDetectionConfig *ObjectDetectionConfig `json:"objectDetectionConfig,omitempty"`

	// Configuration for video object tracking task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.object_tracking_config
	ObjectTrackingConfig *ObjectTrackingConfig `json:"objectTrackingConfig,omitempty"`

	// Configuration for video event labeling task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.event_config
	EventConfig *EventConfig `json:"eventConfig,omitempty"`

	// Configuration for text classification task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.text_classification_config
	TextClassificationConfig *TextClassificationConfig `json:"textClassificationConfig,omitempty"`

	// Configuration for text entity extraction task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.text_entity_extraction_config
	TextEntityExtractionConfig *TextEntityExtractionConfig `json:"textEntityExtractionConfig,omitempty"`

	// HumanAnnotationConfig used when requesting the human labeling task for this
	//  AnnotatedDataset.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotatedDatasetMetadata.human_annotation_config
	HumanAnnotationConfig *HumanAnnotationConfig `json:"humanAnnotationConfig,omitempty"`
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

// +kcc:proto=google.cloud.datalabeling.v1beta1.EventConfig
type EventConfig struct {
	// Required. The list of annotation spec set resource name. Similar to video
	//  classification, we support selecting event from multiple AnnotationSpecSet
	//  at the same time.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EventConfig.annotation_spec_sets
	AnnotationSpecSets []string `json:"annotationSpecSets,omitempty"`
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

// +kcc:proto=google.cloud.datalabeling.v1beta1.LabelStats
type LabelStats struct {
	// Map of each annotation spec's example count. Key is the annotation spec
	//  name and value is the number of examples for that annotation spec.
	//  If the annotated dataset does not have annotation spec, the map will return
	//  a pair where the key is empty string and value is the total number of
	//  annotations.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.LabelStats.example_count
	ExampleCount map[string]int64 `json:"exampleCount,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ObjectDetectionConfig
type ObjectDetectionConfig struct {
	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ObjectDetectionConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Required. Number of frames per second to be extracted from the video.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ObjectDetectionConfig.extraction_frame_rate
	ExtractionFrameRate *float64 `json:"extractionFrameRate,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ObjectTrackingConfig
type ObjectTrackingConfig struct {
	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ObjectTrackingConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.PolylineConfig
type PolylineConfig struct {
	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PolylineConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Optional. Instruction message showed on contributors UI.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PolylineConfig.instruction_message
	InstructionMessage *string `json:"instructionMessage,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.SegmentationConfig
type SegmentationConfig struct {
	// Required. Annotation spec set resource name. format:
	//  projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.SegmentationConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Instruction message showed on labelers UI.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.SegmentationConfig.instruction_message
	InstructionMessage *string `json:"instructionMessage,omitempty"`
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

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextEntityExtractionConfig
type TextEntityExtractionConfig struct {
	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextEntityExtractionConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoClassificationConfig
type VideoClassificationConfig struct {
	// Required. The list of annotation spec set configs.
	//  Since watching a video clip takes much longer time than an image, we
	//  support label with multiple AnnotationSpecSet at the same time. Labels
	//  in each AnnotationSpecSet will be shown in a group to contributors.
	//  Contributors can select one or more (depending on whether to allow multi
	//  label) from each group.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoClassificationConfig.annotation_spec_set_configs
	AnnotationSpecSetConfigs []VideoClassificationConfig_AnnotationSpecSetConfig `json:"annotationSpecSetConfigs,omitempty"`

	// Optional. Option to apply shot detection on the video.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoClassificationConfig.apply_shot_detection
	ApplyShotDetection *bool `json:"applyShotDetection,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoClassificationConfig.AnnotationSpecSetConfig
type VideoClassificationConfig_AnnotationSpecSetConfig struct {
	// Required. Annotation spec set resource name.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoClassificationConfig.AnnotationSpecSetConfig.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Optional. If allow_multi_label is true, contributors are able to
	//  choose multiple labels from one annotation spec set.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoClassificationConfig.AnnotationSpecSetConfig.allow_multi_label
	AllowMultiLabel *bool `json:"allowMultiLabel,omitempty"`
}
