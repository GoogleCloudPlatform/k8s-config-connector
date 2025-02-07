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


// +kcc:proto=google.cloud.automl.v1.BoundingBoxMetricsEntry
type BoundingBoxMetricsEntry struct {
	// Output only. The intersection-over-union threshold value used to compute
	//  this metrics entry.
	// +kcc:proto:field=google.cloud.automl.v1.BoundingBoxMetricsEntry.iou_threshold
	IouThreshold *float32 `json:"iouThreshold,omitempty"`

	// Output only. The mean average precision, most often close to au_prc.
	// +kcc:proto:field=google.cloud.automl.v1.BoundingBoxMetricsEntry.mean_average_precision
	MeanAveragePrecision *float32 `json:"meanAveragePrecision,omitempty"`

	// Output only. Metrics for each label-match confidence_threshold from
	//  0.05,0.10,...,0.95,0.96,0.97,0.98,0.99. Precision-recall curve is
	//  derived from them.
	// +kcc:proto:field=google.cloud.automl.v1.BoundingBoxMetricsEntry.confidence_metrics_entries
	ConfidenceMetricsEntries []BoundingBoxMetricsEntry_ConfidenceMetricsEntry `json:"confidenceMetricsEntries,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry
type BoundingBoxMetricsEntry_ConfidenceMetricsEntry struct {
	// Output only. The confidence threshold value used to compute the metrics.
	// +kcc:proto:field=google.cloud.automl.v1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry.confidence_threshold
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Output only. Recall under the given confidence threshold.
	// +kcc:proto:field=google.cloud.automl.v1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry.recall
	Recall *float32 `json:"recall,omitempty"`

	// Output only. Precision under the given confidence threshold.
	// +kcc:proto:field=google.cloud.automl.v1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry.precision
	Precision *float32 `json:"precision,omitempty"`

	// Output only. The harmonic mean of recall and precision.
	// +kcc:proto:field=google.cloud.automl.v1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry.f1_score
	F1Score *float32 `json:"f1Score,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ClassificationEvaluationMetrics
type ClassificationEvaluationMetrics struct {
	// Output only. The Area Under Precision-Recall Curve metric. Micro-averaged
	//  for the overall evaluation.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.au_prc
	AuPrc *float32 `json:"auPrc,omitempty"`

	// Output only. The Area Under Receiver Operating Characteristic curve metric.
	//  Micro-averaged for the overall evaluation.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.au_roc
	AuRoc *float32 `json:"auRoc,omitempty"`

	// Output only. The Log Loss metric.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.log_loss
	LogLoss *float32 `json:"logLoss,omitempty"`

	// Output only. Metrics for each confidence_threshold in
	//  0.00,0.05,0.10,...,0.95,0.96,0.97,0.98,0.99 and
	//  position_threshold = INT32_MAX_VALUE.
	//  ROC and precision-recall curves, and other aggregated metrics are derived
	//  from them. The confidence metrics entries may also be supplied for
	//  additional values of position_threshold, but from these no aggregated
	//  metrics are computed.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.confidence_metrics_entry
	ConfidenceMetricsEntry []ClassificationEvaluationMetrics_ConfidenceMetricsEntry `json:"confidenceMetricsEntry,omitempty"`

	// Output only. Confusion matrix of the evaluation.
	//  Only set for MULTICLASS classification problems where number
	//  of labels is no more than 10.
	//  Only set for model level evaluation, not for evaluation per label.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.confusion_matrix
	ConfusionMatrix *ClassificationEvaluationMetrics_ConfusionMatrix `json:"confusionMatrix,omitempty"`

	// Output only. The annotation spec ids used for this evaluation.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.annotation_spec_id
	AnnotationSpecID []string `json:"annotationSpecID,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry
type ClassificationEvaluationMetrics_ConfidenceMetricsEntry struct {
	// Output only. Metrics are computed with an assumption that the model
	//  never returns predictions with score lower than this value.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.confidence_threshold
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Output only. Metrics are computed with an assumption that the model
	//  always returns at most this many predictions (ordered by their score,
	//  descendingly), but they all still need to meet the confidence_threshold.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.position_threshold
	PositionThreshold *int32 `json:"positionThreshold,omitempty"`

	// Output only. Recall (True Positive Rate) for the given confidence
	//  threshold.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.recall
	Recall *float32 `json:"recall,omitempty"`

	// Output only. Precision for the given confidence threshold.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.precision
	Precision *float32 `json:"precision,omitempty"`

	// Output only. False Positive Rate for the given confidence threshold.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.false_positive_rate
	FalsePositiveRate *float32 `json:"falsePositiveRate,omitempty"`

	// Output only. The harmonic mean of recall and precision.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.f1_score
	F1Score *float32 `json:"f1Score,omitempty"`

	// Output only. The Recall (True Positive Rate) when only considering the
	//  label that has the highest prediction score and not below the confidence
	//  threshold for each example.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.recall_at1
	RecallAt1 *float32 `json:"recallAt1,omitempty"`

	// Output only. The precision when only considering the label that has the
	//  highest prediction score and not below the confidence threshold for each
	//  example.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.precision_at1
	PrecisionAt1 *float32 `json:"precisionAt1,omitempty"`

	// Output only. The False Positive Rate when only considering the label that
	//  has the highest prediction score and not below the confidence threshold
	//  for each example.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.false_positive_rate_at1
	FalsePositiveRateAt1 *float32 `json:"falsePositiveRateAt1,omitempty"`

	// Output only. The harmonic mean of [recall_at1][google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.recall_at1] and [precision_at1][google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.precision_at1].
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.f1_score_at1
	F1ScoreAt1 *float32 `json:"f1ScoreAt1,omitempty"`

	// Output only. The number of model created labels that match a ground truth
	//  label.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.true_positive_count
	TruePositiveCount *int64 `json:"truePositiveCount,omitempty"`

	// Output only. The number of model created labels that do not match a
	//  ground truth label.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.false_positive_count
	FalsePositiveCount *int64 `json:"falsePositiveCount,omitempty"`

	// Output only. The number of ground truth labels that are not matched
	//  by a model created label.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.false_negative_count
	FalseNegativeCount *int64 `json:"falseNegativeCount,omitempty"`

	// Output only. The number of labels that were not created by the model,
	//  but if they would, they would not match a ground truth label.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.true_negative_count
	TrueNegativeCount *int64 `json:"trueNegativeCount,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix
type ClassificationEvaluationMetrics_ConfusionMatrix struct {
	// Output only. IDs of the annotation specs used in the confusion matrix.
	//  For Tables CLASSIFICATION
	//  [prediction_type][google.cloud.automl.v1p1beta.TablesModelMetadata.prediction_type]
	//  only list of [annotation_spec_display_name-s][] is populated.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.annotation_spec_id
	AnnotationSpecID []string `json:"annotationSpecID,omitempty"`

	// Output only. Display name of the annotation specs used in the confusion
	//  matrix, as they were at the moment of the evaluation. For Tables
	//  CLASSIFICATION
	//  [prediction_type-s][google.cloud.automl.v1p1beta.TablesModelMetadata.prediction_type],
	//  distinct values of the target column at the moment of the model
	//  evaluation are populated here.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.display_name
	DisplayName []string `json:"displayName,omitempty"`

	// Output only. Rows in the confusion matrix. The number of rows is equal to
	//  the size of `annotation_spec_id`.
	//  `row[i].example_count[j]` is the number of examples that have ground
	//  truth of the `annotation_spec_id[i]` and are predicted as
	//  `annotation_spec_id[j]` by the model being evaluated.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.row
	Row []ClassificationEvaluationMetrics_ConfusionMatrix_Row `json:"row,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.Row
type ClassificationEvaluationMetrics_ConfusionMatrix_Row struct {
	// Output only. Value of the specific cell in the confusion matrix.
	//  The number of values each row has (i.e. the length of the row) is equal
	//  to the length of the `annotation_spec_id` field or, if that one is not
	//  populated, length of the [display_name][google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.display_name] field.
	// +kcc:proto:field=google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.Row.example_count
	ExampleCount []int32 `json:"exampleCount,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ImageObjectDetectionEvaluationMetrics
type ImageObjectDetectionEvaluationMetrics struct {
	// Output only. The total number of bounding boxes (i.e. summed over all
	//  images) the ground truth used to create this evaluation had.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionEvaluationMetrics.evaluated_bounding_box_count
	EvaluatedBoundingBoxCount *int32 `json:"evaluatedBoundingBoxCount,omitempty"`

	// Output only. The bounding boxes match metrics for each
	//  Intersection-over-union threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	//  and each label confidence threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	//  pair.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionEvaluationMetrics.bounding_box_metrics_entries
	BoundingBoxMetricsEntries []BoundingBoxMetricsEntry `json:"boundingBoxMetricsEntries,omitempty"`

	// Output only. The single metric for bounding boxes evaluation:
	//  the mean_average_precision averaged over all bounding_box_metrics_entries.
	// +kcc:proto:field=google.cloud.automl.v1.ImageObjectDetectionEvaluationMetrics.bounding_box_mean_average_precision
	BoundingBoxMeanAveragePrecision *float32 `json:"boundingBoxMeanAveragePrecision,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.ModelEvaluation
type ModelEvaluation struct {
	// Model evaluation metrics for image, text, video and tables
	//  classification.
	//  Tables problem is considered a classification when the target column
	//  is CATEGORY DataType.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.classification_evaluation_metrics
	ClassificationEvaluationMetrics *ClassificationEvaluationMetrics `json:"classificationEvaluationMetrics,omitempty"`

	// Model evaluation metrics for translation.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.translation_evaluation_metrics
	TranslationEvaluationMetrics *TranslationEvaluationMetrics `json:"translationEvaluationMetrics,omitempty"`

	// Model evaluation metrics for image object detection.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.image_object_detection_evaluation_metrics
	ImageObjectDetectionEvaluationMetrics *ImageObjectDetectionEvaluationMetrics `json:"imageObjectDetectionEvaluationMetrics,omitempty"`

	// Evaluation metrics for text sentiment models.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.text_sentiment_evaluation_metrics
	TextSentimentEvaluationMetrics *TextSentimentEvaluationMetrics `json:"textSentimentEvaluationMetrics,omitempty"`

	// Evaluation metrics for text extraction models.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.text_extraction_evaluation_metrics
	TextExtractionEvaluationMetrics *TextExtractionEvaluationMetrics `json:"textExtractionEvaluationMetrics,omitempty"`

	// Output only. Resource name of the model evaluation.
	//  Format:
	//  `projects/{project_id}/locations/{location_id}/models/{model_id}/modelEvaluations/{model_evaluation_id}`
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.name
	Name *string `json:"name,omitempty"`

	// Output only. The ID of the annotation spec that the model evaluation applies to. The
	//  The ID is empty for the overall model evaluation.
	//  For Tables annotation specs in the dataset do not exist and this ID is
	//  always not set, but for CLASSIFICATION
	//  [prediction_type-s][google.cloud.automl.v1.TablesModelMetadata.prediction_type]
	//  the
	//  [display_name][google.cloud.automl.v1.ModelEvaluation.display_name]
	//  field is used.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.annotation_spec_id
	AnnotationSpecID *string `json:"annotationSpecID,omitempty"`

	// Output only. The value of
	//  [display_name][google.cloud.automl.v1.AnnotationSpec.display_name]
	//  at the moment when the model was trained. Because this field returns a
	//  value at model training time, for different models trained from the same
	//  dataset, the values may differ, since display names could had been changed
	//  between the two model's trainings. For Tables CLASSIFICATION
	//  [prediction_type-s][google.cloud.automl.v1.TablesModelMetadata.prediction_type]
	//  distinct values of the target column at the moment of the model evaluation
	//  are populated here.
	//  The display_name is empty for the overall model evaluation.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Timestamp when this model evaluation was created.
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The number of examples used for model evaluation, i.e. for
	//  which ground truth from time of model creation is compared against the
	//  predicted annotations created by the model.
	//  For overall ModelEvaluation (i.e. with annotation_spec_id not set) this is
	//  the total number of all examples used for evaluation.
	//  Otherwise, this is the count of examples that according to the ground
	//  truth were annotated by the
	//  [annotation_spec_id][google.cloud.automl.v1.ModelEvaluation.annotation_spec_id].
	// +kcc:proto:field=google.cloud.automl.v1.ModelEvaluation.evaluated_example_count
	EvaluatedExampleCount *int32 `json:"evaluatedExampleCount,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TextExtractionEvaluationMetrics
type TextExtractionEvaluationMetrics struct {
	// Output only. The Area under precision recall curve metric.
	// +kcc:proto:field=google.cloud.automl.v1.TextExtractionEvaluationMetrics.au_prc
	AuPrc *float32 `json:"auPrc,omitempty"`

	// Output only. Metrics that have confidence thresholds.
	//  Precision-recall curve can be derived from it.
	// +kcc:proto:field=google.cloud.automl.v1.TextExtractionEvaluationMetrics.confidence_metrics_entries
	ConfidenceMetricsEntries []TextExtractionEvaluationMetrics_ConfidenceMetricsEntry `json:"confidenceMetricsEntries,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TextExtractionEvaluationMetrics.ConfidenceMetricsEntry
type TextExtractionEvaluationMetrics_ConfidenceMetricsEntry struct {
	// Output only. The confidence threshold value used to compute the metrics.
	//  Only annotations with score of at least this threshold are considered to
	//  be ones the model would return.
	// +kcc:proto:field=google.cloud.automl.v1.TextExtractionEvaluationMetrics.ConfidenceMetricsEntry.confidence_threshold
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Output only. Recall under the given confidence threshold.
	// +kcc:proto:field=google.cloud.automl.v1.TextExtractionEvaluationMetrics.ConfidenceMetricsEntry.recall
	Recall *float32 `json:"recall,omitempty"`

	// Output only. Precision under the given confidence threshold.
	// +kcc:proto:field=google.cloud.automl.v1.TextExtractionEvaluationMetrics.ConfidenceMetricsEntry.precision
	Precision *float32 `json:"precision,omitempty"`

	// Output only. The harmonic mean of recall and precision.
	// +kcc:proto:field=google.cloud.automl.v1.TextExtractionEvaluationMetrics.ConfidenceMetricsEntry.f1_score
	F1Score *float32 `json:"f1Score,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TextSentimentEvaluationMetrics
type TextSentimentEvaluationMetrics struct {
	// Output only. Precision.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.precision
	Precision *float32 `json:"precision,omitempty"`

	// Output only. Recall.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.recall
	Recall *float32 `json:"recall,omitempty"`

	// Output only. The harmonic mean of recall and precision.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.f1_score
	F1Score *float32 `json:"f1Score,omitempty"`

	// Output only. Mean absolute error. Only set for the overall model
	//  evaluation, not for evaluation of a single annotation spec.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.mean_absolute_error
	MeanAbsoluteError *float32 `json:"meanAbsoluteError,omitempty"`

	// Output only. Mean squared error. Only set for the overall model
	//  evaluation, not for evaluation of a single annotation spec.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.mean_squared_error
	MeanSquaredError *float32 `json:"meanSquaredError,omitempty"`

	// Output only. Linear weighted kappa. Only set for the overall model
	//  evaluation, not for evaluation of a single annotation spec.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.linear_kappa
	LinearKappa *float32 `json:"linearKappa,omitempty"`

	// Output only. Quadratic weighted kappa. Only set for the overall model
	//  evaluation, not for evaluation of a single annotation spec.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.quadratic_kappa
	QuadraticKappa *float32 `json:"quadraticKappa,omitempty"`

	// Output only. Confusion matrix of the evaluation.
	//  Only set for the overall model evaluation, not for evaluation of a single
	//  annotation spec.
	// +kcc:proto:field=google.cloud.automl.v1.TextSentimentEvaluationMetrics.confusion_matrix
	ConfusionMatrix *ClassificationEvaluationMetrics_ConfusionMatrix `json:"confusionMatrix,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1.TranslationEvaluationMetrics
type TranslationEvaluationMetrics struct {
	// Output only. BLEU score.
	// +kcc:proto:field=google.cloud.automl.v1.TranslationEvaluationMetrics.bleu_score
	BleuScore *float64 `json:"bleuScore,omitempty"`

	// Output only. BLEU score for base model.
	// +kcc:proto:field=google.cloud.automl.v1.TranslationEvaluationMetrics.base_bleu_score
	BaseBleuScore *float64 `json:"baseBleuScore,omitempty"`
}
