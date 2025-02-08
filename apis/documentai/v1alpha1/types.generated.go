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


// +kcc:proto=google.cloud.documentai.v1.Evaluation
type Evaluation struct {
	// The resource name of the evaluation.
	//  Format:
	//  `projects/{project}/locations/{location}/processors/{processor}/processorVersions/{processor_version}/evaluations/{evaluation}`
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.name
	Name *string `json:"name,omitempty"`

	// The time that the evaluation was created.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Counters for the documents used in the evaluation.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.document_counters
	DocumentCounters *Evaluation_Counters `json:"documentCounters,omitempty"`

	// Metrics for all the entities in aggregate.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.all_entities_metrics
	AllEntitiesMetrics *Evaluation_MultiConfidenceMetrics `json:"allEntitiesMetrics,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The KMS key name used for encryption.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// The KMS key version with which data is encrypted.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Evaluation.ConfidenceLevelMetrics
type Evaluation_ConfidenceLevelMetrics struct {
	// The confidence level.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.ConfidenceLevelMetrics.confidence_level
	ConfidenceLevel *float32 `json:"confidenceLevel,omitempty"`

	// The metrics at the specific confidence level.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.ConfidenceLevelMetrics.metrics
	Metrics *Evaluation_Metrics `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Evaluation.Counters
type Evaluation_Counters struct {
	// How many documents were sent for evaluation.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Counters.input_documents_count
	InputDocumentsCount *int32 `json:"inputDocumentsCount,omitempty"`

	// How many documents were not included in the evaluation as they didn't
	//  pass validation.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Counters.invalid_documents_count
	InvalidDocumentsCount *int32 `json:"invalidDocumentsCount,omitempty"`

	// How many documents were not included in the evaluation as Document AI
	//  failed to process them.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Counters.failed_documents_count
	FailedDocumentsCount *int32 `json:"failedDocumentsCount,omitempty"`

	// How many documents were used in the evaluation.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Counters.evaluated_documents_count
	EvaluatedDocumentsCount *int32 `json:"evaluatedDocumentsCount,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Evaluation.Metrics
type Evaluation_Metrics struct {
	// The calculated precision.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.precision
	Precision *float32 `json:"precision,omitempty"`

	// The calculated recall.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.recall
	Recall *float32 `json:"recall,omitempty"`

	// The calculated f1 score.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.f1_score
	F1Score *float32 `json:"f1Score,omitempty"`

	// The amount of occurrences in predicted documents.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.predicted_occurrences_count
	PredictedOccurrencesCount *int32 `json:"predictedOccurrencesCount,omitempty"`

	// The amount of occurrences in ground truth documents.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.ground_truth_occurrences_count
	GroundTruthOccurrencesCount *int32 `json:"groundTruthOccurrencesCount,omitempty"`

	// The amount of documents with a predicted occurrence.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.predicted_document_count
	PredictedDocumentCount *int32 `json:"predictedDocumentCount,omitempty"`

	// The amount of documents with a ground truth occurrence.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.ground_truth_document_count
	GroundTruthDocumentCount *int32 `json:"groundTruthDocumentCount,omitempty"`

	// The amount of true positives.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.true_positives_count
	TruePositivesCount *int32 `json:"truePositivesCount,omitempty"`

	// The amount of false positives.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.false_positives_count
	FalsePositivesCount *int32 `json:"falsePositivesCount,omitempty"`

	// The amount of false negatives.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.false_negatives_count
	FalseNegativesCount *int32 `json:"falseNegativesCount,omitempty"`

	// The amount of documents that had an occurrence of this label.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.Metrics.total_documents_count
	TotalDocumentsCount *int32 `json:"totalDocumentsCount,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics
type Evaluation_MultiConfidenceMetrics struct {
	// Metrics across confidence levels with fuzzy matching enabled.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics.confidence_level_metrics
	ConfidenceLevelMetrics []Evaluation_ConfidenceLevelMetrics `json:"confidenceLevelMetrics,omitempty"`

	// Metrics across confidence levels with only exact matching.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics.confidence_level_metrics_exact
	ConfidenceLevelMetricsExact []Evaluation_ConfidenceLevelMetrics `json:"confidenceLevelMetricsExact,omitempty"`

	// The calculated area under the precision recall curve (AUPRC), computed by
	//  integrating over all confidence thresholds.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics.auprc
	Auprc *float32 `json:"auprc,omitempty"`

	// The Estimated Calibration Error (ECE) of the confidence of the predicted
	//  entities.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics.estimated_calibration_error
	EstimatedCalibrationError *float32 `json:"estimatedCalibrationError,omitempty"`

	// The AUPRC for metrics with fuzzy matching disabled, i.e., exact matching
	//  only.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics.auprc_exact
	AuprcExact *float32 `json:"auprcExact,omitempty"`

	// The ECE for the predicted entities with fuzzy matching disabled, i.e.,
	//  exact matching only.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics.estimated_calibration_error_exact
	EstimatedCalibrationErrorExact *float32 `json:"estimatedCalibrationErrorExact,omitempty"`

	// The metrics type for the label.
	// +kcc:proto:field=google.cloud.documentai.v1.Evaluation.MultiConfidenceMetrics.metrics_type
	MetricsType *string `json:"metricsType,omitempty"`
}
