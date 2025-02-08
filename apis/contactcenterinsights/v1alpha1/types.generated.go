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

// +kcc:proto=google.cloud.contactcenterinsights.v1.RedactionConfig
type RedactionConfig struct {
	// The fully-qualified DLP deidentify template resource name.
	//  Format:
	//  `projects/{project}/deidentifyTemplates/{template}`
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.RedactionConfig.deidentify_template
	DeidentifyTemplate *string `json:"deidentifyTemplate,omitempty"`

	// The fully-qualified DLP inspect template resource name.
	//  Format:
	//  `projects/{project}/locations/{location}/inspectTemplates/{template}`
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.RedactionConfig.inspect_template
	InspectTemplate *string `json:"inspectTemplate,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.Settings
type Settings struct {
	// Immutable. The resource name of the settings resource.
	//  Format:
	//  projects/{project}/locations/{location}/settings
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.name
	Name *string `json:"name,omitempty"`

	// A language code to be applied to each transcript segment unless the segment
	//  already specifies a language code. Language code defaults to "en-US" if it
	//  is neither specified on the segment nor here.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The default TTL for newly-created conversations. If a conversation has a
	//  specified expiration, that value will be used instead. Changing this
	//  value will not change the expiration of existing conversations.
	//  Conversations with no expire time persist until they are deleted.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.conversation_ttl
	ConversationTtl *string `json:"conversationTtl,omitempty"`

	// A map that maps a notification trigger to a Pub/Sub topic. Each time a
	//  specified trigger occurs, Insights will notify the corresponding Pub/Sub
	//  topic.
	//
	//  Keys are notification triggers. Supported keys are:
	//
	//  * "all-triggers": Notify each time any of the supported triggers occurs.
	//  * "create-analysis": Notify each time an analysis is created.
	//  * "create-conversation": Notify each time a conversation is created.
	//  * "export-insights-data": Notify each time an export is complete.
	//  * "ingest-conversations": Notify each time an IngestConversations LRO is
	//  complete.
	//  * "update-conversation": Notify each time a conversation is updated via
	//  UpdateConversation.
	//  * "upload-conversation": Notify when an UploadConversation LRO is complete.
	//
	//  Values are Pub/Sub topics. The format of each Pub/Sub topic is:
	//  projects/{project}/topics/{topic}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.pubsub_notification_settings
	PubsubNotificationSettings map[string]string `json:"pubsubNotificationSettings,omitempty"`

	// Default analysis settings.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.analysis_config
	AnalysisConfig *Settings_AnalysisConfig `json:"analysisConfig,omitempty"`

	// Default DLP redaction resources to be applied while ingesting
	//  conversations. This applies to conversations ingested from the
	//  `UploadConversation` and `IngestConversations` endpoints, including
	//  conversations coming from CCAI Platform.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.redaction_config
	RedactionConfig *RedactionConfig `json:"redactionConfig,omitempty"`

	// Optional. Default Speech-to-Text resources to use while ingesting audio
	//  files. Optional, CCAI Insights will create a default if not provided. This
	//  applies to conversations ingested from the `UploadConversation` and
	//  `IngestConversations` endpoints, including conversations coming from CCAI
	//  Platform.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.speech_config
	SpeechConfig *SpeechConfig `json:"speechConfig,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.Settings.AnalysisConfig
type Settings_AnalysisConfig struct {
	// Percentage of conversations created using Dialogflow runtime integration
	//  to analyze automatically, between [0, 100].
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.AnalysisConfig.runtime_integration_analysis_percentage
	RuntimeIntegrationAnalysisPercentage *float64 `json:"runtimeIntegrationAnalysisPercentage,omitempty"`

	// Percentage of conversations created using the UploadConversation endpoint
	//  to analyze automatically, between [0, 100].
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.AnalysisConfig.upload_conversation_analysis_percentage
	UploadConversationAnalysisPercentage *float64 `json:"uploadConversationAnalysisPercentage,omitempty"`

	// To select the annotators to run and the phrase matchers to use
	//  (if any). If not specified, all annotators will be run.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.AnalysisConfig.annotator_selector
	AnnotatorSelector *AnnotatorSelector `json:"annotatorSelector,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.SpeechConfig
type SpeechConfig struct {
	// The fully-qualified Speech Recognizer resource name.
	//  Format:
	//  `projects/{project_id}/locations/{location}/recognizer/{recognizer}`
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.SpeechConfig.speech_recognizer
	SpeechRecognizer *string `json:"speechRecognizer,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.Settings
type SettingsObservedState struct {
	// Output only. The time at which the settings was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the settings were last updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Settings.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
