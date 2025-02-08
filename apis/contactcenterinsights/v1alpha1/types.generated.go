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


// +kcc:proto=google.cloud.contactcenterinsights.v1.FeedbackLabel
type FeedbackLabel struct {
	// String label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.label
	Label *string `json:"label,omitempty"`

	// QaAnswer label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.qa_answer_label
	QaAnswerLabel *QaAnswer_AnswerValue `json:"qaAnswerLabel,omitempty"`

	// Immutable. Resource name of the FeedbackLabel.
	//  Format:
	//  projects/{project}/locations/{location}/conversations/{conversation}/feedbackLabels/{feedback_label}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.name
	Name *string `json:"name,omitempty"`

	// Resource name of the resource to be labeled.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.labeled_resource
	LabeledResource *string `json:"labeledResource,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue
type QaAnswer_AnswerValue struct {
	// String value.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.str_value
	StrValue *string `json:"strValue,omitempty"`

	// Numerical value.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.num_value
	NumValue *float64 `json:"numValue,omitempty"`

	// Boolean value.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// A value of "Not Applicable (N/A)". Should only ever be `true`.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.na_value
	NaValue *bool `json:"naValue,omitempty"`

	// A short string used as an identifier. Matches the value used in
	//  QaQuestion.AnswerChoice.key.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.key
	Key *string `json:"key,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.FeedbackLabel
type FeedbackLabelObservedState struct {
	// QaAnswer label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.qa_answer_label
	QaAnswerLabel *QaAnswer_AnswerValueObservedState `json:"qaAnswerLabel,omitempty"`

	// Output only. Create time of the label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time of the label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue
type QaAnswer_AnswerValueObservedState struct {
	// Output only. Numerical score of the answer.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.score
	Score *float64 `json:"score,omitempty"`

	// Output only. The maximum potential score of the question.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.potential_score
	PotentialScore *float64 `json:"potentialScore,omitempty"`

	// Output only. Normalized score of the questions. Calculated as score /
	//  potential_score.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerValue.normalized_score
	NormalizedScore *float64 `json:"normalizedScore,omitempty"`
}
