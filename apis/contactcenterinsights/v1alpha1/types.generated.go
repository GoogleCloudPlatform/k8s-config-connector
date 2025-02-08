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


// +kcc:proto=google.cloud.contactcenterinsights.v1.AnalysisRule
type AnalysisRule struct {
	// Identifier. The resource name of the analysis rule.
	//  Format:
	//  projects/{project}/locations/{location}/analysisRules/{analysis_rule}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.name
	Name *string `json:"name,omitempty"`

	// Display Name of the analysis rule.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Filter for the conversations that should apply this analysis
	//  rule. An empty filter means this analysis rule applies to all
	//  conversations.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.conversation_filter
	ConversationFilter *string `json:"conversationFilter,omitempty"`

	// Selector of annotators to run and the phrase matchers to use for
	//  conversations that matches the conversation_filter. If not specified, NO
	//  annotators will be run.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.annotator_selector
	AnnotatorSelector *AnnotatorSelector `json:"annotatorSelector,omitempty"`

	// Percentage of conversations that we should apply this analysis setting
	//  automatically, between [0, 1]. For example, 0.1 means 10%. Conversations
	//  are sampled in a determenestic way. The original runtime_percentage &
	//  upload percentage will be replaced by defining filters on the conversation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.analysis_percentage
	AnalysisPercentage *float64 `json:"analysisPercentage,omitempty"`

	// If true, apply this rule to conversations. Otherwise, this rule is
	//  inactive and saved as a draft.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.active
	Active *bool `json:"active,omitempty"`
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

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnalysisRule
type AnalysisRuleObservedState struct {
	// Output only. The time at which this analysis rule was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which this analysis rule was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
