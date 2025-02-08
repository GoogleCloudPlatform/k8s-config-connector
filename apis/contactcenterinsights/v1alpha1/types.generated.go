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


// +kcc:proto=google.cloud.contactcenterinsights.v1.QaQuestion
type QaQuestion struct {
	// Identifier. The resource name of the question.
	//  Format:
	//  projects/{project}/locations/{location}/qaScorecards/{qa_scorecard}/revisions/{revision}/qaQuestions/{qa_question}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.name
	Name *string `json:"name,omitempty"`

	// Short, descriptive string, used in the UI where it's not practical
	//  to display the full question body. E.g., "Greeting".
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.abbreviation
	Abbreviation *string `json:"abbreviation,omitempty"`

	// Question text. E.g., "Did the agent greet the customer?"
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.question_body
	QuestionBody *string `json:"questionBody,omitempty"`

	// Instructions describing how to determine the answer.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.answer_instructions
	AnswerInstructions *string `json:"answerInstructions,omitempty"`

	// A list of valid answers to the question, which the LLM must choose from.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.answer_choices
	AnswerChoices []QaQuestion_AnswerChoice `json:"answerChoices,omitempty"`

	// User-defined list of arbitrary tags for the question. Used for
	//  grouping/organization and for weighting the score of each question.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.tags
	Tags []string `json:"tags,omitempty"`

	// Defines the order of the question within its parent scorecard revision.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.order
	Order *int32 `json:"order,omitempty"`

	// Metrics of the underlying tuned LLM over a holdout/test set while fine
	//  tuning the underlying LLM for the given question. This field will only be
	//  populated if and only if the question is part of a scorecard revision that
	//  has been tuned.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.metrics
	Metrics *QaQuestion_Metrics `json:"metrics,omitempty"`

	// Metadata about the tuning operation for the question.This field will only
	//  be populated if and only if the question is part of a scorecard revision
	//  that has been tuned.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.tuning_metadata
	TuningMetadata *QaQuestion_TuningMetadata `json:"tuningMetadata,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaQuestion.AnswerChoice
type QaQuestion_AnswerChoice struct {
	// String value.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.AnswerChoice.str_value
	StrValue *string `json:"strValue,omitempty"`

	// Numerical value.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.AnswerChoice.num_value
	NumValue *float64 `json:"numValue,omitempty"`

	// Boolean value.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.AnswerChoice.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// A value of "Not Applicable (N/A)". If provided, this field may only
	//  be set to `true`. If a question receives this answer, it will be
	//  excluded from any score calculations.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.AnswerChoice.na_value
	NaValue *bool `json:"naValue,omitempty"`

	// A short string used as an identifier.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.AnswerChoice.key
	Key *string `json:"key,omitempty"`

	// Numerical score of the answer, used for generating the overall score of
	//  a QaScorecardResult. If the answer uses na_value, this field is unused.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.AnswerChoice.score
	Score *float64 `json:"score,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaQuestion.Metrics
type QaQuestion_Metrics struct {
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaQuestion.TuningMetadata
type QaQuestion_TuningMetadata struct {
	// Total number of valid labels provided for the question at the time of
	//  tuining.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.TuningMetadata.total_valid_label_count
	TotalValidLabelCount *int64 `json:"totalValidLabelCount,omitempty"`

	// A list of any applicable data validation warnings about the question's
	//  feedback labels.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.TuningMetadata.dataset_validation_warnings
	DatasetValidationWarnings []string `json:"datasetValidationWarnings,omitempty"`

	// Error status of the tuning operation for the question. Will only be set
	//  if the tuning operation failed.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.TuningMetadata.tuning_error
	TuningError *string `json:"tuningError,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaQuestion
type QaQuestionObservedState struct {
	// Output only. The time at which this question was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the question was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Metrics of the underlying tuned LLM over a holdout/test set while fine
	//  tuning the underlying LLM for the given question. This field will only be
	//  populated if and only if the question is part of a scorecard revision that
	//  has been tuned.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.metrics
	Metrics *QaQuestion_MetricsObservedState `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaQuestion.Metrics
type QaQuestion_MetricsObservedState struct {
	// Output only. Accuracy of the model. Measures the percentage of correct
	//  answers the model gave on the test set.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.Metrics.accuracy
	Accuracy *float64 `json:"accuracy,omitempty"`
}
