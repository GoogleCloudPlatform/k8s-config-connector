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


// +kcc:proto=google.cloud.dialogflow.v2beta1.AutomatedAgentConfig
type AutomatedAgentConfig struct {
	// Required. ID of the Dialogflow agent environment to use.
	//
	//  This project needs to either be the same project as the conversation or you
	//  need to grant `service-<Conversation Project
	//  Number>@gcp-sa-dialogflow.iam.gserviceaccount.com` the `Dialogflow API
	//  Service Agent` role in this project.
	//
	//  - For ES agents, use format: `projects/<Project ID>/locations/<Location
	//  ID>/agent/environments/<Environment ID or '-'>`. If environment is not
	//  specified, the default `draft` environment is used. Refer to
	//  [DetectIntentRequest](/dialogflow/docs/reference/rpc/google.cloud.dialogflow.v2beta1#google.cloud.dialogflow.v2beta1.DetectIntentRequest)
	//  for more details.
	//
	//  - For CX agents, use format `projects/<Project ID>/locations/<Location
	//  ID>/agents/<Agent ID>/environments/<Environment ID
	//  or '-'>`. If environment is not specified, the default `draft` environment
	//  is used.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.AutomatedAgentConfig.agent
	Agent *string `json:"agent,omitempty"`

	// Optional. Configure lifetime of the Dialogflow session.
	//  By default, a Dialogflow CX session remains active and its data is stored
	//  for 30 minutes after the last request is sent for the session.
	//  This value should be no longer than 1 day.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.AutomatedAgentConfig.session_ttl
	SessionTtl *string `json:"sessionTtl,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.ConversationProfile
type ConversationProfile struct {
	// The unique identifier of this conversation profile.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/conversationProfiles/<Conversation Profile ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.name
	Name *string `json:"name,omitempty"`

	// Required. Human readable name for this profile. Max length 1024 bytes.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration for an automated agent to use with this profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.automated_agent_config
	AutomatedAgentConfig *AutomatedAgentConfig `json:"automatedAgentConfig,omitempty"`

	// Configuration for agent assistance to use with this profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.human_agent_assistant_config
	HumanAgentAssistantConfig *HumanAgentAssistantConfig `json:"humanAgentAssistantConfig,omitempty"`

	// Configuration for connecting to a live agent.
	//
	//  Currently, this feature is not general available, please contact Google
	//  to get access.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.human_agent_handoff_config
	HumanAgentHandoffConfig *HumanAgentHandoffConfig `json:"humanAgentHandoffConfig,omitempty"`

	// Configuration for publishing conversation lifecycle events.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.notification_config
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`

	// Configuration for logging conversation lifecycle events.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Configuration for publishing new message events. Event will be sent in
	//  format of
	//  [ConversationEvent][google.cloud.dialogflow.v2beta1.ConversationEvent]
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.new_message_event_notification_config
	NewMessageEventNotificationConfig *NotificationConfig `json:"newMessageEventNotificationConfig,omitempty"`

	// Optional. Configuration for publishing transcription intermediate results.
	//  Event will be sent in format of
	//  [ConversationEvent][google.cloud.dialogflow.v2beta1.ConversationEvent]. If
	//  configured, the following information will be populated as
	//  [ConversationEvent][google.cloud.dialogflow.v2beta1.ConversationEvent]
	//  Pub/Sub message attributes:
	//  - "participant_id"
	//  - "participant_role"
	//  - "message_id"
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.new_recognition_result_notification_config
	NewRecognitionResultNotificationConfig *NotificationConfig `json:"newRecognitionResultNotificationConfig,omitempty"`

	// Settings for speech transcription.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.stt_config
	SttConfig *SpeechToTextConfig `json:"sttConfig,omitempty"`

	// Language code for the conversation profile. If not specified, the language
	//  is en-US. Language at ConversationProfile should be set for all non en-us
	//  languages.
	//  This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt)
	//  language tag. Example: "en-US".
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The time zone of this conversational profile from the
	//  [time zone database](https://www.iana.org/time-zones), e.g.,
	//  America/New_York, Europe/Paris. Defaults to America/New_York.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Name of the CX SecuritySettings reference for the agent.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/securitySettings/<Security Settings ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.security_settings
	SecuritySettings *string `json:"securitySettings,omitempty"`

	// Configuration for Text-to-Speech synthesization.
	//
	//  Used by Phone Gateway to specify synthesization options. If agent defines
	//  synthesization options as well, agent settings overrides the option here.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.tts_config
	TtsConfig *SynthesizeSpeechConfig `json:"ttsConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig
type HumanAgentAssistantConfig struct {
	// Pub/Sub topic on which to publish new agent assistant events.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.notification_config
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`

	// Configuration for agent assistance of human agent participant.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.human_agent_suggestion_config
	HumanAgentSuggestionConfig *HumanAgentAssistantConfig_SuggestionConfig `json:"humanAgentSuggestionConfig,omitempty"`

	// Configuration for agent assistance of end user participant.
	//
	//  Currently, this feature is not general available, please contact Google
	//  to get access.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.end_user_suggestion_config
	EndUserSuggestionConfig *HumanAgentAssistantConfig_SuggestionConfig `json:"endUserSuggestionConfig,omitempty"`

	// Configuration for message analysis.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.message_analysis_config
	MessageAnalysisConfig *HumanAgentAssistantConfig_MessageAnalysisConfig `json:"messageAnalysisConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.ConversationModelConfig
type HumanAgentAssistantConfig_ConversationModelConfig struct {
	// Conversation model resource name. Format: `projects/<Project
	//  ID>/conversationModels/<Model ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.ConversationModelConfig.model
	Model *string `json:"model,omitempty"`

	// Version of current baseline model. It will be ignored if
	//  [model][google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.ConversationModelConfig.model]
	//  is set. Valid versions are:
	//    Article Suggestion baseline model:
	//      - 0.9
	//      - 1.0 (default)
	//    Summarization baseline model:
	//      - 1.0
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.ConversationModelConfig.baseline_model_version
	BaselineModelVersion *string `json:"baselineModelVersion,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.ConversationProcessConfig
type HumanAgentAssistantConfig_ConversationProcessConfig struct {
	// Number of recent non-small-talk sentences to use as context for article
	//  and FAQ suggestion
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.ConversationProcessConfig.recent_sentences_count
	RecentSentencesCount *int32 `json:"recentSentencesCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.MessageAnalysisConfig
type HumanAgentAssistantConfig_MessageAnalysisConfig struct {
	// Enable entity extraction in conversation messages on [agent assist
	//  stage](https://cloud.google.com/dialogflow/priv/docs/contact-center/basics#stages).
	//  If unspecified, defaults to false.
	//
	//  Currently, this feature is not general available, please contact Google
	//  to get access.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.MessageAnalysisConfig.enable_entity_extraction
	EnableEntityExtraction *bool `json:"enableEntityExtraction,omitempty"`

	// Enable sentiment analysis in conversation messages on [agent assist
	//  stage](https://cloud.google.com/dialogflow/priv/docs/contact-center/basics#stages).
	//  If unspecified, defaults to false. Sentiment analysis inspects user input
	//  and identifies the prevailing subjective opinion, especially to determine
	//  a user's attitude as positive, negative, or neutral:
	//  https://cloud.google.com/natural-language/docs/basics#sentiment_analysis
	//  For
	//  [Participants.StreamingAnalyzeContent][google.cloud.dialogflow.v2beta1.Participants.StreamingAnalyzeContent]
	//  method, result will be in
	//  [StreamingAnalyzeContentResponse.message.SentimentAnalysisResult][google.cloud.dialogflow.v2beta1.StreamingAnalyzeContentResponse.message].
	//  For
	//  [Participants.AnalyzeContent][google.cloud.dialogflow.v2beta1.Participants.AnalyzeContent]
	//  method, result will be in
	//  [AnalyzeContentResponse.message.SentimentAnalysisResult][google.cloud.dialogflow.v2beta1.AnalyzeContentResponse.message]
	//  For
	//  [Conversations.ListMessages][google.cloud.dialogflow.v2beta1.Conversations.ListMessages]
	//  method, result will be in
	//  [ListMessagesResponse.messages.SentimentAnalysisResult][google.cloud.dialogflow.v2beta1.ListMessagesResponse.messages]
	//  If Pub/Sub notification is configured, result will be in
	//  [ConversationEvent.new_message_payload.SentimentAnalysisResult][google.cloud.dialogflow.v2beta1.ConversationEvent.new_message_payload].
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.MessageAnalysisConfig.enable_sentiment_analysis
	EnableSentimentAnalysis *bool `json:"enableSentimentAnalysis,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionConfig
type HumanAgentAssistantConfig_SuggestionConfig struct {
	// Configuration of different suggestion features. One feature can have only
	//  one config.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionConfig.feature_configs
	FeatureConfigs []HumanAgentAssistantConfig_SuggestionFeatureConfig `json:"featureConfigs,omitempty"`

	// If `group_suggestion_responses` is false, and there are multiple
	//  `feature_configs` in `event based suggestion` or
	//  StreamingAnalyzeContent, we will try to deliver suggestions to customers
	//  as soon as we get new suggestion. Different type of suggestions based on
	//  the same context will be in  separate Pub/Sub event or
	//  `StreamingAnalyzeContentResponse`.
	//
	//  If `group_suggestion_responses` set to true. All the suggestions to the
	//  same participant based on the same context will be grouped into a single
	//  Pub/Sub event or StreamingAnalyzeContentResponse.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionConfig.group_suggestion_responses
	GroupSuggestionResponses *bool `json:"groupSuggestionResponses,omitempty"`

	// Optional. List of various generator resource names used in the
	//  conversation profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionConfig.generators
	Generators []string `json:"generators,omitempty"`

	// Optional. When disable_high_latency_features_sync_delivery is true and
	//  using the AnalyzeContent API, we will not deliver the responses from high
	//  latency features in the API response. The
	//  human_agent_assistant_config.notification_config must be configured and
	//  enable_event_based_suggestion must be set to true to receive the
	//  responses from high latency features in Pub/Sub. High latency feature(s):
	//  KNOWLEDGE_ASSIST
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionConfig.disable_high_latency_features_sync_delivery
	DisableHighLatencyFeaturesSyncDelivery *bool `json:"disableHighLatencyFeaturesSyncDelivery,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig
type HumanAgentAssistantConfig_SuggestionFeatureConfig struct {
	// The suggestion feature.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.suggestion_feature
	SuggestionFeature *SuggestionFeature `json:"suggestionFeature,omitempty"`

	// Automatically iterates all participants and tries to compile
	//  suggestions.
	//
	//  Supported features: ARTICLE_SUGGESTION, FAQ, DIALOGFLOW_ASSIST,
	//  ENTITY_EXTRACTION, KNOWLEDGE_ASSIST.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.enable_event_based_suggestion
	EnableEventBasedSuggestion *bool `json:"enableEventBasedSuggestion,omitempty"`

	// Optional. Disable the logging of search queries sent by human agents. It
	//  can prevent those queries from being stored at answer records.
	//
	//  Supported features: KNOWLEDGE_SEARCH.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.disable_agent_query_logging
	DisableAgentQueryLogging *bool `json:"disableAgentQueryLogging,omitempty"`

	// Optional. Enable query suggestion even if we can't find its answer.
	//  By default, queries are suggested only if we find its answer.
	//  Supported features: KNOWLEDGE_ASSIST
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.enable_query_suggestion_when_no_answer
	EnableQuerySuggestionWhenNoAnswer *bool `json:"enableQuerySuggestionWhenNoAnswer,omitempty"`

	// Optional. Enable including conversation context during query answer
	//  generation. Supported features: KNOWLEDGE_SEARCH.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.enable_conversation_augmented_query
	EnableConversationAugmentedQuery *bool `json:"enableConversationAugmentedQuery,omitempty"`

	// Optional. Enable query suggestion only.
	//  Supported features: KNOWLEDGE_ASSIST
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.enable_query_suggestion_only
	EnableQuerySuggestionOnly *bool `json:"enableQuerySuggestionOnly,omitempty"`

	// Settings of suggestion trigger.
	//
	//  Currently, only ARTICLE_SUGGESTION, FAQ, and DIALOGFLOW_ASSIST will use
	//  this field.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.suggestion_trigger_settings
	SuggestionTriggerSettings *HumanAgentAssistantConfig_SuggestionTriggerSettings `json:"suggestionTriggerSettings,omitempty"`

	// Configs of query.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.query_config
	QueryConfig *HumanAgentAssistantConfig_SuggestionQueryConfig `json:"queryConfig,omitempty"`

	// Configs of custom conversation model.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.conversation_model_config
	ConversationModelConfig *HumanAgentAssistantConfig_ConversationModelConfig `json:"conversationModelConfig,omitempty"`

	// Configs for processing conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionFeatureConfig.conversation_process_config
	ConversationProcessConfig *HumanAgentAssistantConfig_ConversationProcessConfig `json:"conversationProcessConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig
type HumanAgentAssistantConfig_SuggestionQueryConfig struct {
	// Query from knowledgebase. It is used by:
	//  ARTICLE_SUGGESTION, FAQ.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.knowledge_base_query_source
	KnowledgeBaseQuerySource *HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource `json:"knowledgeBaseQuerySource,omitempty"`

	// Query from knowledge base document. It is used by:
	//  SMART_REPLY, SMART_COMPOSE.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.document_query_source
	DocumentQuerySource *HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource `json:"documentQuerySource,omitempty"`

	// Query from Dialogflow agent. It is used by DIALOGFLOW_ASSIST,
	//  ENTITY_EXTRACTION.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.dialogflow_query_source
	DialogflowQuerySource *HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource `json:"dialogflowQuerySource,omitempty"`

	// Maximum number of results to return. Currently, if unset, defaults to 10.
	//  And the max number is 20.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.max_results
	MaxResults *int32 `json:"maxResults,omitempty"`

	// Confidence threshold of query result.
	//
	//  Agent Assist gives each suggestion a score in the range [0.0, 1.0], based
	//  on the relevance between the suggestion and the current conversation
	//  context. A score of 0.0 has no relevance, while a score of 1.0 has high
	//  relevance. Only suggestions with a score greater than or equal to the
	//  value of this field are included in the results.
	//
	//  For a baseline model (the default), the recommended value is in the range
	//  [0.05, 0.1].
	//
	//  For a custom model, there is no recommended value. Tune this value by
	//  starting from a very low value and slowly increasing until you have
	//  desired results.
	//
	//  If this field is not set, it is default to 0.0, which means that all
	//  suggestions are returned.
	//
	//  Supported features: ARTICLE_SUGGESTION, FAQ, SMART_REPLY, SMART_COMPOSE,
	//  KNOWLEDGE_SEARCH, KNOWLEDGE_ASSIST, ENTITY_EXTRACTION.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.confidence_threshold
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Determines how recent conversation context is filtered when generating
	//  suggestions. If unspecified, no messages will be dropped.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.context_filter_settings
	ContextFilterSettings *HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings `json:"contextFilterSettings,omitempty"`

	// Optional. The customized sections chosen to return when requesting a
	//  summary of a conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.sections
	Sections *HumanAgentAssistantConfig_SuggestionQueryConfig_Sections `json:"sections,omitempty"`

	// Optional. The number of recent messages to include in the context.
	//  Supported features: KNOWLEDGE_ASSIST.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.context_size
	ContextSize *int32 `json:"contextSize,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.ContextFilterSettings
type HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings struct {
	// If set to true, the last message from virtual agent (hand off message)
	//  and the message before it (trigger message of hand off) are dropped.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.ContextFilterSettings.drop_handoff_messages
	DropHandoffMessages *bool `json:"dropHandoffMessages,omitempty"`

	// If set to true, all messages from virtual agent are dropped.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.ContextFilterSettings.drop_virtual_agent_messages
	DropVirtualAgentMessages *bool `json:"dropVirtualAgentMessages,omitempty"`

	// If set to true, all messages from ivr stage are dropped.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.ContextFilterSettings.drop_ivr_messages
	DropIvrMessages *bool `json:"dropIvrMessages,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.DialogflowQuerySource
type HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource struct {
	// Required. The name of a dialogflow virtual agent used for end user side
	//  intent detection and suggestion. Format: `projects/<Project
	//  ID>/locations/<Location ID>/agent`. When multiple agents are allowed in
	//  the same Dialogflow project.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.DialogflowQuerySource.agent
	Agent *string `json:"agent,omitempty"`

	// The Dialogflow assist configuration for human agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.DialogflowQuerySource.human_agent_side_config
	HumanAgentSideConfig *HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig `json:"humanAgentSideConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.DialogflowQuerySource.HumanAgentSideConfig
type HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig struct {
	// Optional. The name of a dialogflow virtual agent used for intent
	//  detection and suggestion triggered by human agent.
	//  Format: `projects/<Project ID>/locations/<Location ID>/agent`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.DialogflowQuerySource.HumanAgentSideConfig.agent
	Agent *string `json:"agent,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.DocumentQuerySource
type HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource struct {
	// Required. Knowledge documents to query from. Format:
	//  `projects/<Project ID>/locations/<Location
	//  ID>/knowledgeBases/<KnowledgeBase ID>/documents/<Document ID>`.
	//  Currently, only one document is supported.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.DocumentQuerySource.documents
	Documents []string `json:"documents,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.KnowledgeBaseQuerySource
type HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource struct {
	// Required. Knowledge bases to query. Format:
	//  `projects/<Project ID>/locations/<Location
	//  ID>/knowledgeBases/<Knowledge Base ID>`. Currently, only one knowledge
	//  base is supported.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.KnowledgeBaseQuerySource.knowledge_bases
	KnowledgeBases []string `json:"knowledgeBases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.Sections
type HumanAgentAssistantConfig_SuggestionQueryConfig_Sections struct {
	// The selected sections chosen to return when requesting a summary of a
	//  conversation. A duplicate selected section will be treated as a single
	//  selected section. If section types are not provided, the default will
	//  be {SITUATION, ACTION, RESULT}.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionQueryConfig.Sections.section_types
	SectionTypes []string `json:"sectionTypes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionTriggerSettings
type HumanAgentAssistantConfig_SuggestionTriggerSettings struct {
	// Do not trigger if last utterance is small talk.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionTriggerSettings.no_small_talk
	NoSmallTalk *bool `json:"noSmallTalk,omitempty"`

	// Only trigger suggestion if participant role of last utterance is
	//  END_USER.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentAssistantConfig.SuggestionTriggerSettings.only_end_user
	OnlyEndUser *bool `json:"onlyEndUser,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig
type HumanAgentHandoffConfig struct {
	// Uses [LivePerson](https://www.liveperson.com).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.live_person_config
	LivePersonConfig *HumanAgentHandoffConfig_LivePersonConfig `json:"livePersonConfig,omitempty"`

	// Uses Salesforce Live Agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.salesforce_live_agent_config
	SalesforceLiveAgentConfig *HumanAgentHandoffConfig_SalesforceLiveAgentConfig `json:"salesforceLiveAgentConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.LivePersonConfig
type HumanAgentHandoffConfig_LivePersonConfig struct {
	// Required. Account number of the LivePerson account to connect. This is
	//  the account number you input at the login page.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.LivePersonConfig.account_number
	AccountNumber *string `json:"accountNumber,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.SalesforceLiveAgentConfig
type HumanAgentHandoffConfig_SalesforceLiveAgentConfig struct {
	// Required. The organization ID of the Salesforce account.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.SalesforceLiveAgentConfig.organization_id
	OrganizationID *string `json:"organizationID,omitempty"`

	// Required. Live Agent deployment ID.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.SalesforceLiveAgentConfig.deployment_id
	DeploymentID *string `json:"deploymentID,omitempty"`

	// Required. Live Agent chat button ID.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.SalesforceLiveAgentConfig.button_id
	ButtonID *string `json:"buttonID,omitempty"`

	// Required. Domain of the Live Agent endpoint for this agent. You can find
	//  the endpoint URL in the `Live Agent settings` page. For example if URL
	//  has the form https://d.la4-c2-phx.salesforceliveagent.com/...,
	//  you should fill in d.la4-c2-phx.salesforceliveagent.com.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.HumanAgentHandoffConfig.SalesforceLiveAgentConfig.endpoint_domain
	EndpointDomain *string `json:"endpointDomain,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.LoggingConfig
type LoggingConfig struct {
	// Whether to log conversation events like
	//  [CONVERSATION_STARTED][google.cloud.dialogflow.v2beta1.ConversationEvent.Type.CONVERSATION_STARTED]
	//  to Stackdriver in the conversation project as JSON format
	//  [ConversationEvent][google.cloud.dialogflow.v2beta1.ConversationEvent]
	//  protos.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.LoggingConfig.enable_stackdriver_logging
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.NotificationConfig
type NotificationConfig struct {
	// Name of the Pub/Sub topic to publish conversation
	//  events like
	//  [CONVERSATION_STARTED][google.cloud.dialogflow.v2beta1.ConversationEvent.Type.CONVERSATION_STARTED]
	//  as serialized
	//  [ConversationEvent][google.cloud.dialogflow.v2beta1.ConversationEvent]
	//  protos.
	//
	//  For telephony integration to receive notification, make sure either this
	//  topic is in the same project as the conversation or you grant
	//  `service-<Conversation Project
	//  Number>@gcp-sa-dialogflow.iam.gserviceaccount.com` the `Dialogflow Service
	//  Agent` role in the topic project.
	//
	//  For chat integration to receive notification, make sure API caller has been
	//  granted the `Dialogflow Service Agent` role for the topic.
	//
	//  Format: `projects/<Project ID>/locations/<Location ID>/topics/<Topic ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.NotificationConfig.topic
	Topic *string `json:"topic,omitempty"`

	// Format of message.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.NotificationConfig.message_format
	MessageFormat *string `json:"messageFormat,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SpeechToTextConfig
type SpeechToTextConfig struct {
	// The speech model used in speech to text.
	//  `SPEECH_MODEL_VARIANT_UNSPECIFIED`, `USE_BEST_AVAILABLE` will be treated as
	//  `USE_ENHANCED`. It can be overridden in
	//  [AnalyzeContentRequest][google.cloud.dialogflow.v2beta1.AnalyzeContentRequest]
	//  and
	//  [StreamingAnalyzeContentRequest][google.cloud.dialogflow.v2beta1.StreamingAnalyzeContentRequest]
	//  request. If enhanced model variant is specified and an enhanced version of
	//  the specified model for the language does not exist, then it would emit an
	//  error.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.speech_model_variant
	SpeechModelVariant *string `json:"speechModelVariant,omitempty"`

	// Which Speech model to select. Select the
	//  model best suited to your domain to get best results. If a model is not
	//  explicitly specified, then Dialogflow auto-selects a model based on other
	//  parameters in the SpeechToTextConfig and Agent settings.
	//  If enhanced speech model is enabled for the agent and an enhanced
	//  version of the specified model for the language does not exist, then the
	//  speech is recognized using the standard version of the specified model.
	//  Refer to
	//  [Cloud Speech API
	//  documentation](https://cloud.google.com/speech-to-text/docs/basics#select-model)
	//  for more details.
	//  If you specify a model, the following models typically have the best
	//  performance:
	//
	//  - phone_call (best for Agent Assist and telephony)
	//  - latest_short (best for Dialogflow non-telephony)
	//  - command_and_search
	//
	//  Leave this field unspecified to use
	//  [Agent Speech
	//  settings](https://cloud.google.com/dialogflow/cx/docs/concept/agent#settings-speech)
	//  for model selection.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.model
	Model *string `json:"model,omitempty"`

	// List of names of Cloud Speech phrase sets that are used for transcription.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.phrase_sets
	PhraseSets []string `json:"phraseSets,omitempty"`

	// Audio encoding of the audio content to process.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.audio_encoding
	AudioEncoding *string `json:"audioEncoding,omitempty"`

	// Sample rate (in Hertz) of the audio content sent in the query.
	//  Refer to
	//  [Cloud Speech API
	//  documentation](https://cloud.google.com/speech-to-text/docs/basics) for
	//  more details.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.sample_rate_hertz
	SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"`

	// The language of the supplied audio. Dialogflow does not do  translations.
	//  See [Language
	//  Support](https://cloud.google.com/dialogflow/docs/reference/language)
	//  for a list of the currently supported language codes. Note that queries in
	//  the same session do not necessarily need to specify the same language.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// If `true`, Dialogflow returns
	//  [SpeechWordInfo][google.cloud.dialogflow.v2beta1.SpeechWordInfo] in
	//  [StreamingRecognitionResult][google.cloud.dialogflow.v2beta1.StreamingRecognitionResult]
	//  with information about the recognized speech words, e.g. start and end time
	//  offsets. If false or unspecified, Speech doesn't return any word-level
	//  information.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.enable_word_info
	EnableWordInfo *bool `json:"enableWordInfo,omitempty"`

	// Use timeout based endpointing, interpreting endpointer sensitivy as
	//  seconds of timeout value.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SpeechToTextConfig.use_timeout_based_endpointing
	UseTimeoutBasedEndpointing *bool `json:"useTimeoutBasedEndpointing,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SuggestionFeature
type SuggestionFeature struct {
	// Type of Human Agent Assistant API feature to request.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SuggestionFeature.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig
type SynthesizeSpeechConfig struct {
	// Optional. Speaking rate/speed, in the range [0.25, 4.0]. 1.0 is the normal
	//  native speed supported by the specific voice. 2.0 is twice as fast, and 0.5
	//  is half as fast. If unset(0.0), defaults to the native 1.0 speed. Any other
	//  values < 0.25 or > 4.0 will return an error.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.speaking_rate
	SpeakingRate *float64 `json:"speakingRate,omitempty"`

	// Optional. Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20
	//  semitones from the original pitch. -20 means decrease 20 semitones from the
	//  original pitch.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.pitch
	Pitch *float64 `json:"pitch,omitempty"`

	// Optional. Volume gain (in dB) of the normal native volume supported by the
	//  specific voice, in the range [-96.0, 16.0]. If unset, or set to a value of
	//  0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB)
	//  will play at approximately half the amplitude of the normal native signal
	//  amplitude. A value of +6.0 (dB) will play at approximately twice the
	//  amplitude of the normal native signal amplitude. We strongly recommend not
	//  to exceed +10 (dB) as there's usually no effective increase in loudness for
	//  any value greater than that.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.volume_gain_db
	VolumeGainDb *float64 `json:"volumeGainDb,omitempty"`

	// Optional. An identifier which selects 'audio effects' profiles that are
	//  applied on (post synthesized) text to speech. Effects are applied on top of
	//  each other in the order they are given.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.effects_profile_id
	EffectsProfileID []string `json:"effectsProfileID,omitempty"`

	// Optional. The desired voice of the synthesized audio.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.voice
	Voice *VoiceSelectionParams `json:"voice,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.VoiceSelectionParams
type VoiceSelectionParams struct {
	// Optional. The name of the voice. If not set, the service will choose a
	//  voice based on the other parameters such as language_code and
	//  [ssml_gender][google.cloud.dialogflow.v2beta1.VoiceSelectionParams.ssml_gender].
	//
	//  For the list of available voices, please refer to [Supported voices and
	//  languages](https://cloud.google.com/text-to-speech/docs/voices).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.VoiceSelectionParams.name
	Name *string `json:"name,omitempty"`

	// Optional. The preferred gender of the voice. If not set, the service will
	//  choose a voice based on the other parameters such as language_code and
	//  [name][google.cloud.dialogflow.v2beta1.VoiceSelectionParams.name]. Note
	//  that this is only a preference, not requirement. If a voice of the
	//  appropriate gender is not available, the synthesizer should substitute a
	//  voice with a different gender rather than failing the request.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.VoiceSelectionParams.ssml_gender
	SsmlGender *string `json:"ssmlGender,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.ConversationProfile
type ConversationProfileObservedState struct {
	// Output only. Create time of the conversation profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time of the conversation profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
