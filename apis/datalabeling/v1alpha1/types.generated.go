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


// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotationSpec
type AnnotationSpec struct {
	// Required. The display name of the AnnotationSpec. Maximum of 64 characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the annotation specification.
	//  The description can be up to 10,000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpec.description
	Description *string `json:"description,omitempty"`
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

// +kcc:proto=google.cloud.datalabeling.v1beta1.ClassificationMetrics
type ClassificationMetrics struct {
	// Precision-recall curve based on ground truth labels, predicted labels, and
	//  scores for the predicted labels.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ClassificationMetrics.pr_curve
	PrCurve *PrCurve `json:"prCurve,omitempty"`

	// Confusion matrix of predicted labels vs. ground truth labels.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ClassificationMetrics.confusion_matrix
	ConfusionMatrix *ConfusionMatrix `json:"confusionMatrix,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ConfusionMatrix
type ConfusionMatrix struct {
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ConfusionMatrix.row
	Row []ConfusionMatrix_Row `json:"row,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ConfusionMatrix.ConfusionMatrixEntry
type ConfusionMatrix_ConfusionMatrixEntry struct {
	// The annotation spec of a predicted label.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ConfusionMatrix.ConfusionMatrixEntry.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`

	// Number of items predicted to have this label. (The ground truth label for
	//  these items is the `Row.annotationSpec` of this entry's parent.)
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ConfusionMatrix.ConfusionMatrixEntry.item_count
	ItemCount *int32 `json:"itemCount,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ConfusionMatrix.Row
type ConfusionMatrix_Row struct {
	// The annotation spec of the ground truth label for this row.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ConfusionMatrix.Row.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`

	// A list of the confusion matrix entries. One entry for each possible
	//  predicted label.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ConfusionMatrix.Row.entries
	Entries []ConfusionMatrix_ConfusionMatrixEntry `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.Evaluation
type Evaluation struct {
	// Output only. Resource name of an evaluation. The name has the following
	//  format:
	//
	//  "projects/<var>{project_id}</var>/datasets/<var>{dataset_id}</var>/evaluations/<var>{evaluation_id</var>}'
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Evaluation.name
	Name *string `json:"name,omitempty"`

	// Output only. Options used in the evaluation job that created this
	//  evaluation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Evaluation.config
	Config *EvaluationConfig `json:"config,omitempty"`

	// Output only. Timestamp for when the evaluation job that created this
	//  evaluation ran.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Evaluation.evaluation_job_run_time
	EvaluationJobRunTime *string `json:"evaluationJobRunTime,omitempty"`

	// Output only. Timestamp for when this evaluation was created.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Evaluation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Metrics comparing predictions to ground truth labels.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Evaluation.evaluation_metrics
	EvaluationMetrics *EvaluationMetrics `json:"evaluationMetrics,omitempty"`

	// Output only. Type of task that the model version being evaluated performs,
	//  as defined in the
	//
	//  [evaluationJobConfig.inputConfig.annotationType][google.cloud.datalabeling.v1beta1.EvaluationJobConfig.input_config]
	//  field of the evaluation job that created this evaluation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Evaluation.annotation_type
	AnnotationType *string `json:"annotationType,omitempty"`

	// Output only. The number of items in the ground truth dataset that were used
	//  for this evaluation. Only populated when the evaulation is for certain
	//  AnnotationTypes.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Evaluation.evaluated_item_count
	EvaluatedItemCount *int64 `json:"evaluatedItemCount,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.EvaluationConfig
type EvaluationConfig struct {
	// Only specify this field if the related model performs image object
	//  detection (`IMAGE_BOUNDING_BOX_ANNOTATION`). Describes how to evaluate
	//  bounding boxes.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationConfig.bounding_box_evaluation_options
	BoundingBoxEvaluationOptions *BoundingBoxEvaluationOptions `json:"boundingBoxEvaluationOptions,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.EvaluationMetrics
type EvaluationMetrics struct {
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationMetrics.classification_metrics
	ClassificationMetrics *ClassificationMetrics `json:"classificationMetrics,omitempty"`

	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationMetrics.object_detection_metrics
	ObjectDetectionMetrics *ObjectDetectionMetrics `json:"objectDetectionMetrics,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ObjectDetectionMetrics
type ObjectDetectionMetrics struct {
	// Precision-recall curve.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ObjectDetectionMetrics.pr_curve
	PrCurve *PrCurve `json:"prCurve,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.PrCurve
type PrCurve struct {
	// The annotation spec of the label for which the precision-recall curve
	//  calculated. If this field is empty, that means the precision-recall curve
	//  is an aggregate curve for all labels.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`

	// Area under the precision-recall curve. Not to be confused with area under
	//  a receiver operating characteristic (ROC) curve.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.area_under_curve
	AreaUnderCurve *float32 `json:"areaUnderCurve,omitempty"`

	// Entries that make up the precision-recall graph. Each entry is a "point" on
	//  the graph drawn for a different `confidence_threshold`.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.confidence_metrics_entries
	ConfidenceMetricsEntries []PrCurve_ConfidenceMetricsEntry `json:"confidenceMetricsEntries,omitempty"`

	// Mean average prcision of this curve.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.mean_average_precision
	MeanAveragePrecision *float32 `json:"meanAveragePrecision,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry
type PrCurve_ConfidenceMetricsEntry struct {
	// Threshold used for this entry.
	//
	//  For classification tasks, this is a classification threshold: a
	//  predicted label is categorized as positive or negative (in the context of
	//  this point on the PR curve) based on whether the label's score meets this
	//  threshold.
	//
	//  For image object detection (bounding box) tasks, this is the
	//  [intersection-over-union
	//
	//  (IOU)](/vision/automl/object-detection/docs/evaluate#intersection-over-union)
	//  threshold for the context of this point on the PR curve.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.confidence_threshold
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Recall value.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.recall
	Recall *float32 `json:"recall,omitempty"`

	// Precision value.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.precision
	Precision *float32 `json:"precision,omitempty"`

	// Harmonic mean of recall and precision.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.f1_score
	F1Score *float32 `json:"f1Score,omitempty"`

	// Recall value for entries with label that has highest score.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.recall_at1
	RecallAt1 *float32 `json:"recallAt1,omitempty"`

	// Precision value for entries with label that has highest score.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.precision_at1
	PrecisionAt1 *float32 `json:"precisionAt1,omitempty"`

	// The harmonic mean of [recall_at1][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.recall_at1] and [precision_at1][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.precision_at1].
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.f1_score_at1
	F1ScoreAt1 *float32 `json:"f1ScoreAt1,omitempty"`

	// Recall value for entries with label that has highest 5 scores.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.recall_at5
	RecallAt5 *float32 `json:"recallAt5,omitempty"`

	// Precision value for entries with label that has highest 5 scores.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.precision_at5
	PrecisionAt5 *float32 `json:"precisionAt5,omitempty"`

	// The harmonic mean of [recall_at5][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.recall_at5] and [precision_at5][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.precision_at5].
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.f1_score_at5
	F1ScoreAt5 *float32 `json:"f1ScoreAt5,omitempty"`
}
