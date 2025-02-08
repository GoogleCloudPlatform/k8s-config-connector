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


// +kcc:proto=google.cloud.contactcenterinsights.v1.QaAnswer
type QaAnswer struct {
	// The QaQuestion answered by this answer.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.qa_question
	QaQuestion *string `json:"qaQuestion,omitempty"`

	// The conversation the answer applies to.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.conversation
	Conversation *string `json:"conversation,omitempty"`

	// Question text. E.g., "Did the agent greet the customer?"
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.question_body
	QuestionBody *string `json:"questionBody,omitempty"`

	// The main answer value, incorporating any manual edits if they exist.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.answer_value
	AnswerValue *QaAnswer_AnswerValue `json:"answerValue,omitempty"`

	// User-defined list of arbitrary tags. Matches the value from
	//  QaScorecard.ScorecardQuestion.tags. Used for grouping/organization and
	//  for weighting the score of each answer.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.tags
	Tags []string `json:"tags,omitempty"`

	// List of all individual answers given to the question.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.answer_sources
	AnswerSources []QaAnswer_AnswerSource `json:"answerSources,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerSource
type QaAnswer_AnswerSource struct {
	// What created the answer.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerSource.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// The answer value from this source.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.AnswerSource.answer_value
	AnswerValue *QaAnswer_AnswerValue `json:"answerValue,omitempty"`
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

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaScorecardResult
type QaScorecardResult struct {
	// Identifier. The name of the scorecard result.
	//  Format:
	//  projects/{project}/locations/{location}/qaScorecardResults/{qa_scorecard_result}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.name
	Name *string `json:"name,omitempty"`

	// The QaScorecardRevision scored by this result.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.qa_scorecard_revision
	QaScorecardRevision *string `json:"qaScorecardRevision,omitempty"`

	// The conversation scored by this result.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.conversation
	Conversation *string `json:"conversation,omitempty"`

	// ID of the agent that handled the conversation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.agent_id
	AgentID *string `json:"agentID,omitempty"`

	// Set of QaAnswers represented in the result.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.qa_answers
	QaAnswers []QaAnswer `json:"qaAnswers,omitempty"`

	// The overall numerical score of the result, incorporating any manual edits
	//  if they exist.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.score
	Score *float64 `json:"score,omitempty"`

	// The maximum potential overall score of the scorecard. Any questions
	//  answered using `na_value` are excluded from this calculation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.potential_score
	PotentialScore *float64 `json:"potentialScore,omitempty"`

	// The normalized score, which is the score divided by the potential score.
	//  Any manual edits are included if they exist.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.normalized_score
	NormalizedScore *float64 `json:"normalizedScore,omitempty"`

	// Collection of tags and their scores.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.qa_tag_results
	QaTagResults []QaScorecardResult_QaTagResult `json:"qaTagResults,omitempty"`

	// List of all individual score sets.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.score_sources
	ScoreSources []QaScorecardResult_ScoreSource `json:"scoreSources,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaScorecardResult.QaTagResult
type QaScorecardResult_QaTagResult struct {
	// The tag the score applies to.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.QaTagResult.tag
	Tag *string `json:"tag,omitempty"`

	// The score the tag applies to.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.QaTagResult.score
	Score *float64 `json:"score,omitempty"`

	// The potential score the tag applies to.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.QaTagResult.potential_score
	PotentialScore *float64 `json:"potentialScore,omitempty"`

	// The normalized score the tag applies to.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.QaTagResult.normalized_score
	NormalizedScore *float64 `json:"normalizedScore,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaScorecardResult.ScoreSource
type QaScorecardResult_ScoreSource struct {
	// What created the score.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.ScoreSource.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// The overall numerical score of the result.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.ScoreSource.score
	Score *float64 `json:"score,omitempty"`

	// The maximum potential overall score of the scorecard. Any questions
	//  answered using `na_value` are excluded from this calculation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.ScoreSource.potential_score
	PotentialScore *float64 `json:"potentialScore,omitempty"`

	// The normalized score, which is the score divided by the potential score.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.ScoreSource.normalized_score
	NormalizedScore *float64 `json:"normalizedScore,omitempty"`

	// Collection of tags and their scores.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.ScoreSource.qa_tag_results
	QaTagResults []QaScorecardResult_QaTagResult `json:"qaTagResults,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaAnswer
type QaAnswerObservedState struct {
	// The main answer value, incorporating any manual edits if they exist.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaAnswer.answer_value
	AnswerValue *QaAnswer_AnswerValueObservedState `json:"answerValue,omitempty"`
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

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaScorecardResult
type QaScorecardResultObservedState struct {
	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Set of QaAnswers represented in the result.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardResult.qa_answers
	QaAnswers []QaAnswerObservedState `json:"qaAnswers,omitempty"`
}
