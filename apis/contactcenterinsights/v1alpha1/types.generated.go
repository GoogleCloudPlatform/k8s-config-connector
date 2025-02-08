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


// +kcc:proto=google.cloud.contactcenterinsights.v1.Analysis
type Analysis struct {
	// Immutable. The resource name of the analysis.
	//  Format:
	//  projects/{project}/locations/{location}/conversations/{conversation}/analyses/{analysis}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Analysis.name
	Name *string `json:"name,omitempty"`

	// To select the annotators to run and the phrase matchers to use
	//  (if any). If not specified, all annotators will be run.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Analysis.annotator_selector
	AnnotatorSelector *AnnotatorSelector `json:"annotatorSelector,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnalysisResult
type AnalysisResult struct {
	// Call-specific metadata created by the analysis.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.call_analysis_metadata
	CallAnalysisMetadata *AnalysisResult_CallAnalysisMetadata `json:"callAnalysisMetadata,omitempty"`

	// The time at which the analysis ended.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata
type AnalysisResult_CallAnalysisMetadata struct {
	// A list of call annotations that apply to this call.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata.annotations
	Annotations []CallAnnotation `json:"annotations,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Overall conversation-level sentiment for each channel of the call.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata.sentiments
	Sentiments []ConversationLevelSentiment `json:"sentiments,omitempty"`

	// Overall conversation-level silence during the call.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata.silence
	Silence *ConversationLevelSilence `json:"silence,omitempty"`

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// Overall conversation-level issue modeling result.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata.issue_model_result
	IssueModelResult *IssueModelResult `json:"issueModelResult,omitempty"`

	// Results of scoring QaScorecards.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata.qa_scorecard_results
	QaScorecardResults []QaScorecardResult `json:"qaScorecardResults,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotationBoundary
type AnnotationBoundary struct {
	// The word index of this boundary with respect to the first word in the
	//  transcript piece. This index starts at zero.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotationBoundary.word_index
	WordIndex *int32 `json:"wordIndex,omitempty"`

	// The index in the sequence of transcribed pieces of the conversation where
	//  the boundary is located. This index starts at zero.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotationBoundary.transcript_index
	TranscriptIndex *int32 `json:"transcriptIndex,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotatorSelector
type AnnotatorSelector struct {
	// Whether to run the interruption annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_interruption_annotator
	RunInterruptionAnnotator *bool `json:"runInterruptionAnnotator,omitempty"`

	// Whether to run the silence annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_silence_annotator
	RunSilenceAnnotator *bool `json:"runSilenceAnnotator,omitempty"`

	// Whether to run the active phrase matcher annotator(s).
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_phrase_matcher_annotator
	RunPhraseMatcherAnnotator *bool `json:"runPhraseMatcherAnnotator,omitempty"`

	// The list of phrase matchers to run. If not provided, all active phrase
	//  matchers will be used. If inactive phrase matchers are provided, they will
	//  not be used. Phrase matchers will be run only if
	//  run_phrase_matcher_annotator is set to true. Format:
	//  projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.phrase_matchers
	PhraseMatchers []string `json:"phraseMatchers,omitempty"`

	// Whether to run the sentiment annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_sentiment_annotator
	RunSentimentAnnotator *bool `json:"runSentimentAnnotator,omitempty"`

	// Whether to run the entity annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_entity_annotator
	RunEntityAnnotator *bool `json:"runEntityAnnotator,omitempty"`

	// Whether to run the intent annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_intent_annotator
	RunIntentAnnotator *bool `json:"runIntentAnnotator,omitempty"`

	// Whether to run the issue model annotator. A model should have already been
	//  deployed for this to take effect.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_issue_model_annotator
	RunIssueModelAnnotator *bool `json:"runIssueModelAnnotator,omitempty"`

	// The issue model to run. If not provided, the most recently deployed topic
	//  model will be used. The provided issue model will only be used for
	//  inference if the issue model is deployed and if run_issue_model_annotator
	//  is set to true. If more than one issue model is provided, only the first
	//  provided issue model will be used for inference.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.issue_models
	IssueModels []string `json:"issueModels,omitempty"`

	// Whether to run the summarization annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_summarization_annotator
	RunSummarizationAnnotator *bool `json:"runSummarizationAnnotator,omitempty"`

	// Configuration for the summarization annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.summarization_config
	SummarizationConfig *AnnotatorSelector_SummarizationConfig `json:"summarizationConfig,omitempty"`

	// Whether to run the QA annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_qa_annotator
	RunQaAnnotator *bool `json:"runQaAnnotator,omitempty"`

	// Configuration for the QA annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.qa_config
	QaConfig *AnnotatorSelector_QaConfig `json:"qaConfig,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotatorSelector.QaConfig
type AnnotatorSelector_QaConfig struct {
	// A manual list of scorecards to score.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.QaConfig.scorecard_list
	ScorecardList *AnnotatorSelector_QaConfig_ScorecardList `json:"scorecardList,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotatorSelector.QaConfig.ScorecardList
type AnnotatorSelector_QaConfig_ScorecardList struct {
	// List of QaScorecardRevisions.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.QaConfig.ScorecardList.qa_scorecard_revisions
	QaScorecardRevisions []string `json:"qaScorecardRevisions,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotatorSelector.SummarizationConfig
type AnnotatorSelector_SummarizationConfig struct {
	// Resource name of the Dialogflow conversation profile.
	//  Format:
	//  projects/{project}/locations/{location}/conversationProfiles/{conversation_profile}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.SummarizationConfig.conversation_profile
	ConversationProfile *string `json:"conversationProfile,omitempty"`

	// Default summarization model to be used.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.SummarizationConfig.summarization_model
	SummarizationModel *string `json:"summarizationModel,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.CallAnnotation
type CallAnnotation struct {
	// Data specifying an interruption.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.interruption_data
	InterruptionData *InterruptionData `json:"interruptionData,omitempty"`

	// Data specifying sentiment.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.sentiment_data
	SentimentData *SentimentData `json:"sentimentData,omitempty"`

	// Data specifying silence.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.silence_data
	SilenceData *SilenceData `json:"silenceData,omitempty"`

	// Data specifying a hold.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.hold_data
	HoldData *HoldData `json:"holdData,omitempty"`

	// Data specifying an entity mention.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.entity_mention_data
	EntityMentionData *EntityMentionData `json:"entityMentionData,omitempty"`

	// Data specifying an intent match.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.intent_match_data
	IntentMatchData *IntentMatchData `json:"intentMatchData,omitempty"`

	// Data specifying a phrase match.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.phrase_match_data
	PhraseMatchData *PhraseMatchData `json:"phraseMatchData,omitempty"`

	// Data specifying an issue match.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.issue_match_data
	IssueMatchData *IssueMatchData `json:"issueMatchData,omitempty"`

	// The channel of the audio where the annotation occurs. For single-channel
	//  audio, this field is not populated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.channel_tag
	ChannelTag *int32 `json:"channelTag,omitempty"`

	// The boundary in the conversation where the annotation starts, inclusive.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.annotation_start_boundary
	AnnotationStartBoundary *AnnotationBoundary `json:"annotationStartBoundary,omitempty"`

	// The boundary in the conversation where the annotation ends, inclusive.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.CallAnnotation.annotation_end_boundary
	AnnotationEndBoundary *AnnotationBoundary `json:"annotationEndBoundary,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.ConversationLevelSentiment
type ConversationLevelSentiment struct {
	// The channel of the audio that the data applies to.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.ConversationLevelSentiment.channel_tag
	ChannelTag *int32 `json:"channelTag,omitempty"`

	// Data specifying sentiment.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.ConversationLevelSentiment.sentiment_data
	SentimentData *SentimentData `json:"sentimentData,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.ConversationLevelSilence
type ConversationLevelSilence struct {
	// Amount of time calculated to be in silence.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.ConversationLevelSilence.silence_duration
	SilenceDuration *string `json:"silenceDuration,omitempty"`

	// Percentage of the total conversation spent in silence.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.ConversationLevelSilence.silence_percentage
	SilencePercentage *float32 `json:"silencePercentage,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.Entity
type Entity struct {
	// The representative name for the entity.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Entity.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The entity type.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Entity.type
	Type *string `json:"type,omitempty"`

	// Metadata associated with the entity.
	//
	//  For most entity types, the metadata is a Wikipedia URL (`wikipedia_url`)
	//  and Knowledge Graph MID (`mid`), if they are available. For the metadata
	//  associated with other entity types, see the Type table below.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Entity.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// The salience score associated with the entity in the [0, 1.0] range.
	//
	//  The salience score for an entity provides information about the
	//  importance or centrality of that entity to the entire document text.
	//  Scores closer to 0 are less salient, while scores closer to 1.0 are highly
	//  salient.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Entity.salience
	Salience *float32 `json:"salience,omitempty"`

	// The aggregate sentiment expressed for this entity in the conversation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Entity.sentiment
	Sentiment *SentimentData `json:"sentiment,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.EntityMentionData
type EntityMentionData struct {
	// The key of this entity in conversation entities.
	//  Can be used to retrieve the exact `Entity` this mention is attached to.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.EntityMentionData.entity_unique_id
	EntityUniqueID *string `json:"entityUniqueID,omitempty"`

	// The type of the entity mention.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.EntityMentionData.type
	Type *string `json:"type,omitempty"`

	// Sentiment expressed for this mention of the entity.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.EntityMentionData.sentiment
	Sentiment *SentimentData `json:"sentiment,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.HoldData
type HoldData struct {
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.Intent
type Intent struct {
	// The unique identifier of the intent.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Intent.id
	ID *string `json:"id,omitempty"`

	// The human-readable name of the intent.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Intent.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IntentMatchData
type IntentMatchData struct {
	// The id of the matched intent.
	//  Can be used to retrieve the corresponding intent information.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IntentMatchData.intent_unique_id
	IntentUniqueID *string `json:"intentUniqueID,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.InterruptionData
type InterruptionData struct {
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueAssignment
type IssueAssignment struct {
	// Resource name of the assigned issue.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueAssignment.issue
	Issue *string `json:"issue,omitempty"`

	// Score indicating the likelihood of the issue assignment.
	//  currently bounded on [0,1].
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueAssignment.score
	Score *float64 `json:"score,omitempty"`

	// Immutable. Display name of the assigned issue. This field is set at time of
	//  analyis and immutable since then.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueAssignment.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueMatchData
type IssueMatchData struct {
	// Information about the issue's assignment.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueMatchData.issue_assignment
	IssueAssignment *IssueAssignment `json:"issueAssignment,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueModelResult
type IssueModelResult struct {
	// Issue model that generates the result.
	//  Format: projects/{project}/locations/{location}/issueModels/{issue_model}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModelResult.issue_model
	IssueModel *string `json:"issueModel,omitempty"`

	// All the matched issues.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModelResult.issues
	Issues []IssueAssignment `json:"issues,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.PhraseMatchData
type PhraseMatchData struct {
	// The unique identifier (the resource name) of the phrase matcher.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchData.phrase_matcher
	PhraseMatcher *string `json:"phraseMatcher,omitempty"`

	// The human-readable name of the phrase matcher.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchData.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

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

// +kcc:proto=google.cloud.contactcenterinsights.v1.SentimentData
type SentimentData struct {
	// A non-negative number from 0 to infinity which represents the abolute
	//  magnitude of sentiment regardless of score.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.SentimentData.magnitude
	Magnitude *float32 `json:"magnitude,omitempty"`

	// The sentiment score between -1.0 (negative) and 1.0 (positive).
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.SentimentData.score
	Score *float32 `json:"score,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.SilenceData
type SilenceData struct {
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.Analysis
type AnalysisObservedState struct {
	// Output only. The time at which the analysis was requested.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Analysis.request_time
	RequestTime *string `json:"requestTime,omitempty"`

	// Output only. The time at which the analysis was created, which occurs when
	//  the long-running operation completes.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Analysis.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The result of the analysis, which is populated when the
	//  analysis finishes.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Analysis.analysis_result
	AnalysisResult *AnalysisResult `json:"analysisResult,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnalysisResult
type AnalysisResultObservedState struct {
	// Call-specific metadata created by the analysis.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.call_analysis_metadata
	CallAnalysisMetadata *AnalysisResult_CallAnalysisMetadataObservedState `json:"callAnalysisMetadata,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata
type AnalysisResult_CallAnalysisMetadataObservedState struct {
	// Results of scoring QaScorecards.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisResult.CallAnalysisMetadata.qa_scorecard_results
	QaScorecardResults []QaScorecardResultObservedState `json:"qaScorecardResults,omitempty"`
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
