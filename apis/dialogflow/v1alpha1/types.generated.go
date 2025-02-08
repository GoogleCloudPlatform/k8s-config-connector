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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Action
type Action struct {
	// Optional. Agent obtained a message from the customer.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Action.user_utterance
	UserUtterance *UserUtterance `json:"userUtterance,omitempty"`

	// Optional. Action performed by the agent as a message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Action.agent_utterance
	AgentUtterance *AgentUtterance `json:"agentUtterance,omitempty"`

	// Optional. Action performed on behalf of the agent by calling a plugin
	//  tool.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Action.tool_use
	ToolUse *ToolUse `json:"toolUse,omitempty"`

	// Optional. Action performed on behalf of the agent by invoking a child
	//  playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Action.playbook_invocation
	PlaybookInvocation *PlaybookInvocation `json:"playbookInvocation,omitempty"`

	// Optional. Action performed on behalf of the agent by invoking a CX flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Action.flow_invocation
	FlowInvocation *FlowInvocation `json:"flowInvocation,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings
type AdvancedSettings struct {
	// If present, incoming audio is exported by Dialogflow to the configured
	//  Google Cloud Storage destination.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.audio_export_gcs_destination
	AudioExportGcsDestination *GcsDestination `json:"audioExportGcsDestination,omitempty"`

	// Settings for speech to text detection.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	//  - Page level
	//  - Parameter level
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.speech_settings
	SpeechSettings *AdvancedSettings_SpeechSettings `json:"speechSettings,omitempty"`

	// Settings for DTMF.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	//  - Page level
	//  - Parameter level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.dtmf_settings
	DtmfSettings *AdvancedSettings_DtmfSettings `json:"dtmfSettings,omitempty"`

	// Settings for logging.
	//  Settings for Dialogflow History, Contact Center messages, StackDriver logs,
	//  and speech logging.
	//  Exposed at the following levels:
	//  - Agent level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.logging_settings
	LoggingSettings *AdvancedSettings_LoggingSettings `json:"loggingSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings
type AdvancedSettings_DtmfSettings struct {
	// If true, incoming audio is processed for DTMF (dual tone multi frequency)
	//  events. For example, if the caller presses a button on their telephone
	//  keypad and DTMF processing is enabled, Dialogflow will detect the
	//  event (e.g. a "3" was pressed) in the incoming audio and pass the event
	//  to the bot to drive business logic (e.g. when 3 is pressed, return the
	//  account balance).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Max length of DTMF digits.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings.max_digits
	MaxDigits *int32 `json:"maxDigits,omitempty"`

	// The digit that terminates a DTMF digit sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings.finish_digit
	FinishDigit *string `json:"finishDigit,omitempty"`

	// Interdigit timeout setting for matching dtmf input to regex.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings.interdigit_timeout_duration
	InterdigitTimeoutDuration *string `json:"interdigitTimeoutDuration,omitempty"`

	// Endpoint timeout setting for matching dtmf input to regex.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings.endpointing_timeout_duration
	EndpointingTimeoutDuration *string `json:"endpointingTimeoutDuration,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings
type AdvancedSettings_LoggingSettings struct {
	// Enables Google Cloud Logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings.enable_stackdriver_logging
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	// Enables DF Interaction logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings.enable_interaction_logging
	EnableInteractionLogging *bool `json:"enableInteractionLogging,omitempty"`

	// Enables consent-based end-user input redaction, if true, a pre-defined
	//  session parameter `$session.params.conversation-redaction` will be
	//  used to determine if the utterance should be redacted.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings.enable_consent_based_redaction
	EnableConsentBasedRedaction *bool `json:"enableConsentBasedRedaction,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings
type AdvancedSettings_SpeechSettings struct {
	// Sensitivity of the speech model that detects the end of speech.
	//  Scale from 0 to 100.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.endpointer_sensitivity
	EndpointerSensitivity *int32 `json:"endpointerSensitivity,omitempty"`

	// Timeout before detecting no speech.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.no_speech_timeout
	NoSpeechTimeout *string `json:"noSpeechTimeout,omitempty"`

	// Use timeout based endpointing, interpreting endpointer sensitivy as
	//  seconds of timeout value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.use_timeout_based_endpointing
	UseTimeoutBasedEndpointing *bool `json:"useTimeoutBasedEndpointing,omitempty"`

	// Mapping from language to Speech-to-Text model. The mapped Speech-to-Text
	//  model will be selected for requests from its corresponding language.
	//  For more information, see
	//  [Speech
	//  models](https://cloud.google.com/dialogflow/cx/docs/concept/speech-models).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.models
	Models map[string]string `json:"models,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.AgentUtterance
type AgentUtterance struct {
	// Required. Message content in text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AgentUtterance.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.AudioInput
type AudioInput struct {
	// Required. Instructs the speech recognizer how to process the speech audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AudioInput.config
	Config *InputAudioConfig `json:"config,omitempty"`

	// The natural language speech audio to be processed.
	//  A single request can contain up to 2 minutes of speech audio data.
	//  The [transcribed
	//  text][google.cloud.dialogflow.cx.v3beta1.QueryResult.transcript] cannot
	//  contain more than 256 bytes.
	//
	//  For non-streaming audio detect intent, both `config` and `audio` must be
	//  provided.
	//  For streaming audio detect intent, `config` must be provided in
	//  the first request and `audio` must be provided in all following requests.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.AudioInput.audio
	Audio []byte `json:"audio,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.BargeInConfig
type BargeInConfig struct {
	// Duration that is not eligible for barge-in at the beginning of the input
	//  audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BargeInConfig.no_barge_in_duration
	NoBargeInDuration *string `json:"noBargeInDuration,omitempty"`

	// Total duration for the playback at the beginning of the input audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BargeInConfig.total_duration
	TotalDuration *string `json:"totalDuration,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.BoostSpec
type BoostSpec struct {
	// Optional. Condition boost specifications. If a document matches multiple
	//  conditions in the specifictions, boost scores from these specifications are
	//  all applied and combined in a non-linear way. Maximum number of
	//  specifications is 20.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.condition_boost_specs
	ConditionBoostSpecs []BoostSpec_ConditionBoostSpec `json:"conditionBoostSpecs,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec
type BoostSpec_ConditionBoostSpec struct {
	// Optional. An expression which specifies a boost condition. The syntax and
	//  supported fields are the same as a filter expression.
	//  Examples:
	//
	//  * To boost documents with document ID "doc_1" or "doc_2", and
	//  color
	//    "Red" or "Blue":
	//      * (id: ANY("doc_1", "doc_2")) AND (color: ANY("Red","Blue"))
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.condition
	Condition *string `json:"condition,omitempty"`

	// Optional. Strength of the condition boost, which should be in [-1, 1].
	//  Negative boost means demotion. Default is 0.0.
	//
	//  Setting to 1.0 gives the document a big promotion. However, it does not
	//  necessarily mean that the boosted document will be the top result at
	//  all times, nor that other documents will be excluded. Results could
	//  still be shown even when none of them matches the condition. And
	//  results that are significantly more relevant to the search query can
	//  still trump your heavily favored but irrelevant documents.
	//
	//  Setting to -1.0 gives the document a big demotion. However, results
	//  that are deeply relevant might still be shown. The document will have
	//  an upstream battle to get a fairly high ranking, but it is not blocked
	//  out completely.
	//
	//  Setting to 0.0 means no boost applied. The boosting condition is
	//  ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.boost
	Boost *float32 `json:"boost,omitempty"`

	// Optional. Complex specification for custom ranking based on customer
	//  defined attribute value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.boost_control_spec
	BoostControlSpec *BoostSpec_ConditionBoostSpec_BoostControlSpec `json:"boostControlSpec,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec
type BoostSpec_ConditionBoostSpec_BoostControlSpec struct {
	// Optional. The name of the field whose value will be used to determine
	//  the boost amount.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec.field_name
	FieldName *string `json:"fieldName,omitempty"`

	// Optional. The attribute type to be used to determine the boost amount.
	//  The attribute value can be derived from the field value of the
	//  specified field_name. In the case of numerical it is straightforward
	//  i.e. attribute_value = numerical_field_value. In the case of freshness
	//  however, attribute_value = (time.now() - datetime_field_value).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec.attribute_type
	AttributeType *string `json:"attributeType,omitempty"`

	// Optional. The interpolation type to be applied to connect the control
	//  points listed below.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec.interpolation_type
	InterpolationType *string `json:"interpolationType,omitempty"`

	// Optional. The control points used to define the curve. The monotonic
	//  function (defined through the interpolation_type above) passes through
	//  the control points listed here.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec.control_points
	ControlPoints []BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint `json:"controlPoints,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec.ControlPoint
type BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint struct {
	// Optional. Can be one of:
	//  1. The numerical field value.
	//  2. The duration spec for freshness:
	//  The value must be formatted as an XSD `dayTimeDuration` value (a
	//  restricted subset of an ISO 8601 duration value). The pattern for
	//  this is: `[nD][T[nH][nM][nS]]`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec.ControlPoint.attribute_value
	AttributeValue *string `json:"attributeValue,omitempty"`

	// Optional. The value between -1 to 1 by which to boost the score if
	//  the attribute_value evaluates to the value specified above.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpec.ConditionBoostSpec.BoostControlSpec.ControlPoint.boost_amount
	BoostAmount *float32 `json:"boostAmount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.BoostSpecs
type BoostSpecs struct {
	// Optional. Data Stores where the boosting configuration is applied. The full
	//  names of the referenced data stores. Formats:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	//  `projects/{project}/locations/{location}/dataStores/{data_store}
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpecs.data_stores
	DataStores []string `json:"dataStores,omitempty"`

	// Optional. A list of boosting specifications.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.BoostSpecs.spec
	Spec []BoostSpec `json:"spec,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation
type Conversation struct {
	// Identifier. The identifier of the conversation.
	//  If conversation ID is reused, interactions happened later than 48 hours of
	//  the conversation's create time will be ignored. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/conversations/<ConversationID>`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.name
	Name *string `json:"name,omitempty"`

	// The type of the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.type
	Type *string `json:"type,omitempty"`

	// The language of the conversation, which is the language of the first
	//  request in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Start time of the conversation, which is the time of the first request of
	//  the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Duration of the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.duration
	Duration *string `json:"duration,omitempty"`

	// Conversation metrics.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.metrics
	Metrics *Conversation_Metrics `json:"metrics,omitempty"`

	// All the matched [Intent][google.cloud.dialogflow.cx.v3beta1.Intent] in the
	//  conversation. Only `name` and `display_name` are filled in this message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.intents
	Intents []Intent `json:"intents,omitempty"`

	// All the [Flow][google.cloud.dialogflow.cx.v3beta1.Flow] the conversation
	//  has went through. Only `name` and `display_name` are filled in this
	//  message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.flows
	Flows []Flow `json:"flows,omitempty"`

	// All the [Page][google.cloud.dialogflow.cx.v3beta1.Page] the conversation
	//  has went through. Only `name` and `display_name` are filled in this
	//  message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.pages
	Pages []Page `json:"pages,omitempty"`

	// Interactions of the conversation.
	//  Only populated for `GetConversation` and empty for `ListConversations`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.interactions
	Interactions []Conversation_Interaction `json:"interactions,omitempty"`

	// Environment of the conversation.
	//  Only `name` and `display_name` are filled in this message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.environment
	Environment *Environment `json:"environment,omitempty"`

	// Flow versions used in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.flow_versions
	FlowVersions map[string]int64 `json:"flowVersions,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction
type Conversation_Interaction struct {
	// The request of the interaction.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.request
	Request *DetectIntentRequest `json:"request,omitempty"`

	// The final response of the interaction.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.response
	Response *DetectIntentResponse `json:"response,omitempty"`

	// The partial responses of the interaction. Empty if there is no partial
	//  response in the interaction.
	//  See the
	//  [partial response
	//  documentation][https://cloud.google.com/dialogflow/cx/docs/concept/fulfillment#queue].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.partial_responses
	PartialResponses []DetectIntentResponse `json:"partialResponses,omitempty"`

	// The input text or the transcript of the input audio in the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.request_utterances
	RequestUtterances *string `json:"requestUtterances,omitempty"`

	// The output text or the transcript of the output audio in the responses.
	//  If multiple output messages are returned, they will be concatenated into
	//  one.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.response_utterances
	ResponseUtterances *string `json:"responseUtterances,omitempty"`

	// The time that the interaction was created.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Missing transition predicted for the interaction. This field is set only
	//  if the interaction match type was no-match.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.missing_transition
	MissingTransition *Conversation_Interaction_MissingTransition `json:"missingTransition,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.MissingTransition
type Conversation_Interaction_MissingTransition struct {
	// Name of the intent that could have triggered.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.MissingTransition.intent_display_name
	IntentDisplayName *string `json:"intentDisplayName,omitempty"`

	// Score of the above intent. The higher it is the more likely a
	//  transition was missed on a given page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.MissingTransition.score
	Score *float32 `json:"score,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics
type Conversation_Metrics struct {
	// The number of interactions in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.interaction_count
	InteractionCount *int32 `json:"interactionCount,omitempty"`

	// Duration of all the input's audio in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.input_audio_duration
	InputAudioDuration *string `json:"inputAudioDuration,omitempty"`

	// Duration of all the output's audio in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.output_audio_duration
	OutputAudioDuration *string `json:"outputAudioDuration,omitempty"`

	// Maximum latency of the
	//  [Webhook][google.cloud.dialogflow.cx.v3beta1.Webhook] calls in the
	//  conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.max_webhook_latency
	MaxWebhookLatency *string `json:"maxWebhookLatency,omitempty"`

	// A signal that indicates the interaction with the Dialogflow agent has
	//  ended.
	//  If any response has the
	//  [ResponseMessage.end_interaction][google.cloud.dialogflow.cx.v3beta1.ResponseMessage.end_interaction]
	//  signal, this is set to true.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.has_end_interaction
	HasEndInteraction *bool `json:"hasEndInteraction,omitempty"`

	// Hands off conversation to a human agent.
	//  If any response has the
	//  [ResponseMessage.live_agent_handoff][google.cloud.dialogflow.cx.v3beta1.ResponseMessage.live_agent_handoff]signal,
	//  this is set to true.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.has_live_agent_handoff
	HasLiveAgentHandoff *bool `json:"hasLiveAgentHandoff,omitempty"`

	// The average confidence all of the
	//  [Match][google.cloud.dialogflow.cx.v3beta1.Match] in the conversation.
	//  Values range from 0.0 (completely uncertain) to 1.0 (completely certain).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.average_match_confidence
	AverageMatchConfidence *float32 `json:"averageMatchConfidence,omitempty"`

	// Query input counts.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.query_input_count
	QueryInputCount *Conversation_Metrics_QueryInputCount `json:"queryInputCount,omitempty"`

	// Match type counts.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.match_type_count
	MatchTypeCount *Conversation_Metrics_MatchTypeCount `json:"matchTypeCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount
type Conversation_Metrics_MatchTypeCount struct {
	// The number of matches with type
	//  [Match.MatchType.MATCH_TYPE_UNSPECIFIED][google.cloud.dialogflow.cx.v3beta1.Match.MatchType.MATCH_TYPE_UNSPECIFIED].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount.unspecified_count
	UnspecifiedCount *int32 `json:"unspecifiedCount,omitempty"`

	// The number of matches with type
	//  [Match.MatchType.INTENT][google.cloud.dialogflow.cx.v3beta1.Match.MatchType.INTENT].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount.intent_count
	IntentCount *int32 `json:"intentCount,omitempty"`

	// The number of matches with type
	//  [Match.MatchType.DIRECT_INTENT][google.cloud.dialogflow.cx.v3beta1.Match.MatchType.DIRECT_INTENT].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount.direct_intent_count
	DirectIntentCount *int32 `json:"directIntentCount,omitempty"`

	// The number of matches with type
	//  [Match.MatchType.PARAMETER_FILLING][google.cloud.dialogflow.cx.v3beta1.Match.MatchType.PARAMETER_FILLING].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount.parameter_filling_count
	ParameterFillingCount *int32 `json:"parameterFillingCount,omitempty"`

	// The number of matches with type
	//  [Match.MatchType.NO_MATCH][google.cloud.dialogflow.cx.v3beta1.Match.MatchType.NO_MATCH].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount.no_match_count
	NoMatchCount *int32 `json:"noMatchCount,omitempty"`

	// The number of matches with type
	//  [Match.MatchType.NO_INPUT][google.cloud.dialogflow.cx.v3beta1.Match.MatchType.NO_INPUT].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount.no_input_count
	NoInputCount *int32 `json:"noInputCount,omitempty"`

	// The number of matches with type
	//  [Match.MatchType.EVENT][google.cloud.dialogflow.cx.v3beta1.Match.MatchType.EVENT].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.MatchTypeCount.event_count
	EventCount *int32 `json:"eventCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.QueryInputCount
type Conversation_Metrics_QueryInputCount struct {
	// The number of [TextInput][google.cloud.dialogflow.cx.v3beta1.TextInput]
	//  in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.QueryInputCount.text_count
	TextCount *int32 `json:"textCount,omitempty"`

	// The number of
	//  [IntentInput][google.cloud.dialogflow.cx.v3beta1.IntentInput] in the
	//  conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.QueryInputCount.intent_count
	IntentCount *int32 `json:"intentCount,omitempty"`

	// The number of
	//  [AudioInput][google.cloud.dialogflow.cx.v3beta1.AudioInput] in the
	//  conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.QueryInputCount.audio_count
	AudioCount *int32 `json:"audioCount,omitempty"`

	// The number of
	//  [EventInput][google.cloud.dialogflow.cx.v3beta1.EventInput] in the
	//  conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.QueryInputCount.event_count
	EventCount *int32 `json:"eventCount,omitempty"`

	// The number of [DtmfInput][google.cloud.dialogflow.cx.v3beta1.DtmfInput]
	//  in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Metrics.QueryInputCount.dtmf_count
	DtmfCount *int32 `json:"dtmfCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection
type DataStoreConnection struct {
	// The type of the connected data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// The full name of the referenced data store.
	//  Formats:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	//  `projects/{project}/locations/{location}/dataStores/{data_store}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store
	DataStore *string `json:"dataStore,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals
type DataStoreConnectionSignals struct {
	// Optional. Diagnostic info related to the rewriter model call.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.rewriter_model_call_signals
	RewriterModelCallSignals *DataStoreConnectionSignals_RewriterModelCallSignals `json:"rewriterModelCallSignals,omitempty"`

	// Optional. Rewritten string query used for search.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.rewritten_query
	RewrittenQuery *string `json:"rewrittenQuery,omitempty"`

	// Optional. Search snippets included in the answer generation prompt.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.search_snippets
	SearchSnippets []DataStoreConnectionSignals_SearchSnippet `json:"searchSnippets,omitempty"`

	// Optional. Diagnostic info related to the answer generation model call.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.answer_generation_model_call_signals
	AnswerGenerationModelCallSignals *DataStoreConnectionSignals_AnswerGenerationModelCallSignals `json:"answerGenerationModelCallSignals,omitempty"`

	// Optional. The final compiled answer.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.answer
	Answer *string `json:"answer,omitempty"`

	// Optional. Answer parts with relevant citations.
	//  Concatenation of texts should add up the `answer` (not counting
	//  whitespaces).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.answer_parts
	AnswerParts []DataStoreConnectionSignals_AnswerPart `json:"answerParts,omitempty"`

	// Optional. Snippets cited by the answer generation model from the most to
	//  least relevant.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.cited_snippets
	CitedSnippets []DataStoreConnectionSignals_CitedSnippet `json:"citedSnippets,omitempty"`

	// Optional. Grounding signals.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.grounding_signals
	GroundingSignals *DataStoreConnectionSignals_GroundingSignals `json:"groundingSignals,omitempty"`

	// Optional. Safety check result.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.safety_signals
	SafetySignals *DataStoreConnectionSignals_SafetySignals `json:"safetySignals,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.AnswerGenerationModelCallSignals
type DataStoreConnectionSignals_AnswerGenerationModelCallSignals struct {
	// Prompt as sent to the model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.AnswerGenerationModelCallSignals.rendered_prompt
	RenderedPrompt *string `json:"renderedPrompt,omitempty"`

	// Output of the generative model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.AnswerGenerationModelCallSignals.model_output
	ModelOutput *string `json:"modelOutput,omitempty"`

	// Name of the generative model. For example, "gemini-ultra", "gemini-pro",
	//  "gemini-1.5-flash" etc. Defaults to "Other" if the model is unknown.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.AnswerGenerationModelCallSignals.model
	Model *string `json:"model,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.AnswerPart
type DataStoreConnectionSignals_AnswerPart struct {
	// Substring of the answer.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.AnswerPart.text
	Text *string `json:"text,omitempty"`

	// Citations for this answer part. Indices of `search_snippets`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.AnswerPart.supporting_indices
	SupportingIndices []int32 `json:"supportingIndices,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.CitedSnippet
type DataStoreConnectionSignals_CitedSnippet struct {
	// Details of the snippet.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.CitedSnippet.search_snippet
	SearchSnippet *DataStoreConnectionSignals_SearchSnippet `json:"searchSnippet,omitempty"`

	// Index of the snippet in `search_snippets` field.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.CitedSnippet.snippet_index
	SnippetIndex *int32 `json:"snippetIndex,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.GroundingSignals
type DataStoreConnectionSignals_GroundingSignals struct {
	// Represents the decision of the grounding check.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.GroundingSignals.decision
	Decision *string `json:"decision,omitempty"`

	// Grounding score bucket setting.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.GroundingSignals.score
	Score *string `json:"score,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.RewriterModelCallSignals
type DataStoreConnectionSignals_RewriterModelCallSignals struct {
	// Prompt as sent to the model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.RewriterModelCallSignals.rendered_prompt
	RenderedPrompt *string `json:"renderedPrompt,omitempty"`

	// Output of the generative model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.RewriterModelCallSignals.model_output
	ModelOutput *string `json:"modelOutput,omitempty"`

	// Name of the generative model. For example, "gemini-ultra", "gemini-pro",
	//  "gemini-1.5-flash" etc. Defaults to "Other" if the model is unknown.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.RewriterModelCallSignals.model
	Model *string `json:"model,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SafetySignals
type DataStoreConnectionSignals_SafetySignals struct {
	// Safety decision.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SafetySignals.decision
	Decision *string `json:"decision,omitempty"`

	// Specifies banned phrase match subject.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SafetySignals.banned_phrase_match
	BannedPhraseMatch *string `json:"bannedPhraseMatch,omitempty"`

	// The matched banned phrase if there was a match.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SafetySignals.matched_banned_phrase
	MatchedBannedPhrase *string `json:"matchedBannedPhrase,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SearchSnippet
type DataStoreConnectionSignals_SearchSnippet struct {
	// Title of the enclosing document.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SearchSnippet.document_title
	DocumentTitle *string `json:"documentTitle,omitempty"`

	// Uri for the document. Present if specified for the document.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SearchSnippet.document_uri
	DocumentURI *string `json:"documentURI,omitempty"`

	// Text included in the prompt.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnectionSignals.SearchSnippet.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DetectIntentRequest
type DetectIntentRequest struct {
	// Required. The name of the session this query is sent to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/sessions/<SessionID>`
	//  or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/environments/<EnvironmentID>/sessions/<SessionID>`.
	//  If `Environment ID` is not specified, we assume default 'draft'
	//  environment. It's up to the API caller to choose an appropriate `Session
	//  ID`. It can be a random number or some type of session identifiers
	//  (preferably hashed). The length of the `Session ID` must not exceed 36
	//  characters.
	//
	//  For more information, see the [sessions
	//  guide](https://cloud.google.com/dialogflow/cx/docs/concept/session).
	//
	//  Note: Always use agent versions for production traffic.
	//  See [Versions and
	//  environments](https://cloud.google.com/dialogflow/cx/docs/concept/version).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentRequest.session
	Session *string `json:"session,omitempty"`

	// The parameters of this query.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentRequest.query_params
	QueryParams *QueryParameters `json:"queryParams,omitempty"`

	// Required. The input specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentRequest.query_input
	QueryInput *QueryInput `json:"queryInput,omitempty"`

	// Instructs the speech synthesizer how to generate the output audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentRequest.output_audio_config
	OutputAudioConfig *OutputAudioConfig `json:"outputAudioConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse
type DetectIntentResponse struct {
	// Output only. The unique identifier of the response. It can be used to
	//  locate a response in the training example set or for reporting issues.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse.response_id
	ResponseID *string `json:"responseID,omitempty"`

	// The result of the conversational query.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse.query_result
	QueryResult *QueryResult `json:"queryResult,omitempty"`

	// The audio data bytes encoded as specified in the request.
	//  Note: The output audio is generated based on the values of default platform
	//  text responses found in the
	//  [`query_result.response_messages`][google.cloud.dialogflow.cx.v3beta1.QueryResult.response_messages]
	//  field. If multiple default text responses exist, they will be concatenated
	//  when generating audio. If no default platform text responses exist, the
	//  generated audio content will be empty.
	//
	//  In some scenarios, multiple output audio fields may be present in the
	//  response structure. In these cases, only the top-most-level audio output
	//  has content.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse.output_audio
	OutputAudio []byte `json:"outputAudio,omitempty"`

	// The config used by the speech synthesizer to generate the output audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse.output_audio_config
	OutputAudioConfig *OutputAudioConfig `json:"outputAudioConfig,omitempty"`

	// Response type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse.response_type
	ResponseType *string `json:"responseType,omitempty"`

	// Indicates whether the partial response can be cancelled when a later
	//  response arrives. e.g. if the agent specified some music as partial
	//  response, it can be cancelled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse.allow_cancellation
	AllowCancellation *bool `json:"allowCancellation,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DtmfInput
type DtmfInput struct {
	// The dtmf digits.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DtmfInput.digits
	Digits *string `json:"digits,omitempty"`

	// The finish digit (if any).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DtmfInput.finish_digit
	FinishDigit *string `json:"finishDigit,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.EntityType.Entity
type EntityType_Entity struct {
	// Required. The primary value associated with this entity entry.
	//  For example, if the entity type is *vegetable*, the value could be
	//  *scallions*.
	//
	//  For `KIND_MAP` entity types:
	//
	//  *   A canonical value to be used in place of synonyms.
	//
	//  For `KIND_LIST` entity types:
	//
	//  *   A string that can contain references to other entity types (with or
	//      without aliases).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.Entity.value
	Value *string `json:"value,omitempty"`

	// Required. A collection of value synonyms. For example, if the entity type
	//  is *vegetable*, and `value` is *scallions*, a synonym could be *green
	//  onions*.
	//
	//  For `KIND_LIST` entity types:
	//
	//  *   This collection must contain exactly one synonym equal to `value`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.Entity.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Environment
type Environment struct {
	// The name of the environment.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/environments/<EnvironmentID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the environment (unique in an agent).
	//  Limit of 64 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The human-readable description of the environment. The maximum length is
	//  500 characters. If exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.description
	Description *string `json:"description,omitempty"`

	// A list of configurations for flow versions. You should include version
	//  configs for all flows that are reachable from [`Start
	//  Flow`][Agent.start_flow] in the agent. Otherwise, an error will be
	//  returned.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.version_configs
	VersionConfigs []Environment_VersionConfig `json:"versionConfigs,omitempty"`

	// The test cases config for continuous tests of this environment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.test_cases_config
	TestCasesConfig *Environment_TestCasesConfig `json:"testCasesConfig,omitempty"`

	// The webhook configuration for this environment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.webhook_config
	WebhookConfig *Environment_WebhookConfig `json:"webhookConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Environment.TestCasesConfig
type Environment_TestCasesConfig struct {
	// A list of test case names to run. They should be under the same agent.
	//  Format of each test case name:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/testCases/<TestCaseID>`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.TestCasesConfig.test_cases
	TestCases []string `json:"testCases,omitempty"`

	// Whether to run test cases in
	//  [TestCasesConfig.test_cases][google.cloud.dialogflow.cx.v3beta1.Environment.TestCasesConfig.test_cases]
	//  periodically. Default false. If set to true, run once a day.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.TestCasesConfig.enable_continuous_run
	EnableContinuousRun *bool `json:"enableContinuousRun,omitempty"`

	// Whether to run test cases in
	//  [TestCasesConfig.test_cases][google.cloud.dialogflow.cx.v3beta1.Environment.TestCasesConfig.test_cases]
	//  before deploying a flow version to the environment. Default false.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.TestCasesConfig.enable_predeployment_run
	EnablePredeploymentRun *bool `json:"enablePredeploymentRun,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Environment.VersionConfig
type Environment_VersionConfig struct {
	// Required. Both flow and playbook versions are supported.
	//  Format for flow version:
	//  projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/versions/<VersionID>.
	//  Format for playbook version:
	//  projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>/versions/<VersionID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.VersionConfig.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Environment.WebhookConfig
type Environment_WebhookConfig struct {
	// The list of webhooks to override for the agent environment. The webhook
	//  must exist in the agent. You can override fields in
	//  [`generic_web_service`][google.cloud.dialogflow.cx.v3beta1.Webhook.generic_web_service]
	//  and
	//  [`service_directory`][google.cloud.dialogflow.cx.v3beta1.Webhook.service_directory].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.WebhookConfig.webhook_overrides
	WebhookOverrides []Webhook `json:"webhookOverrides,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.EventHandler
type EventHandler struct {

	// Required. The name of the event to handle.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EventHandler.event
	Event *string `json:"event,omitempty"`

	// The fulfillment to call when the event occurs.
	//  Handling webhook errors with a fulfillment enabled with webhook could
	//  cause infinite loop. It is invalid to specify such fulfillment for a
	//  handler handling webhooks.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EventHandler.trigger_fulfillment
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EventHandler.target_page
	TargetPage *string `json:"targetPage,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EventHandler.target_flow
	TargetFlow *string `json:"targetFlow,omitempty"`

	// The target playbook to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EventHandler.target_playbook
	TargetPlaybook *string `json:"targetPlaybook,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.EventInput
type EventInput struct {
	// Name of the event.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EventInput.event
	Event *string `json:"event,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Example
type Example struct {
	// The unique identifier of the playbook example.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>/examples/<ExampleID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.name
	Name *string `json:"name,omitempty"`

	// Optional. The input to the playbook in the example.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.playbook_input
	PlaybookInput *PlaybookInput `json:"playbookInput,omitempty"`

	// Optional. The output of the playbook in the example.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.playbook_output
	PlaybookOutput *PlaybookOutput `json:"playbookOutput,omitempty"`

	// Required. The ordered list of actions performed by the end user and the
	//  Dialogflow agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.actions
	Actions []Action `json:"actions,omitempty"`

	// Required. The display name of the example.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The high level concise description of the example. The max number
	//  of characters is 200.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.description
	Description *string `json:"description,omitempty"`

	// Required. Example's output state.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.conversation_state
	ConversationState *string `json:"conversationState,omitempty"`

	// Optional. The language code of the example.
	//  If not specified, the agent's default language is used.
	//  Note: languages must be enabled in the agent before they can be used.
	//  Note: example's language code is not currently used in dialogflow agents.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.FilterSpecs
type FilterSpecs struct {
	// Optional. Data Stores where the boosting configuration is applied. The full
	//  names of the referenced data stores. Formats:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	//  `projects/{project}/locations/{location}/dataStores/{data_store}
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.FilterSpecs.data_stores
	DataStores []string `json:"dataStores,omitempty"`

	// Optional. The filter expression to be applied.
	//  Expression syntax is documented at
	//  https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata#filter-expression-syntax
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.FilterSpecs.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Flow
type Flow struct {
	// The unique identifier of the flow.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the flow. The maximum length is 500 characters. If
	//  exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.description
	Description *string `json:"description,omitempty"`

	// A flow's transition routes serve two purposes:
	//
	//  *   They are responsible for matching the user's first utterances in the
	//  flow.
	//  *   They are inherited by every page's [transition
	//  routes][Page.transition_routes] and can support use cases such as the user
	//  saying "help" or "can I talk to a human?", which can be handled in a common
	//  way regardless of the current page. Transition routes defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  TransitionRoutes are evalauted in the following order:
	//
	//  *   TransitionRoutes with intent specified.
	//  *   TransitionRoutes with only condition specified.
	//
	//  TransitionRoutes with intent specified are inherited by pages in the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.transition_routes
	TransitionRoutes []TransitionRoute `json:"transitionRoutes,omitempty"`

	// A flow's event handlers serve two purposes:
	//
	//  *   They are responsible for handling events (e.g. no match,
	//  webhook errors) in the flow.
	//  *   They are inherited by every page's [event
	//  handlers][Page.event_handlers], which can be used to handle common events
	//  regardless of the current page. Event handlers defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  Unlike
	//  [transition_routes][google.cloud.dialogflow.cx.v3beta1.Flow.transition_routes],
	//  these handlers are evaluated on a first-match basis. The first one that
	//  matches the event get executed, with the rest being ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.event_handlers
	EventHandlers []EventHandler `json:"eventHandlers,omitempty"`

	// A flow's transition route group serve two purposes:
	//
	//  *   They are responsible for matching the user's first utterances in the
	//  flow.
	//  *   They are inherited by every page's [transition
	//  route groups][Page.transition_route_groups]. Transition route groups
	//  defined in the page have higher priority than those defined in the flow.
	//
	//  Format:`projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  for agent-level groups.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.transition_route_groups
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	// NLU related settings of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.nlu_settings
	NluSettings *NluSettings `json:"nluSettings,omitempty"`

	// Hierarchical advanced settings for this flow. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`

	// Optional. Multi-lingual agent settings for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.multi_language_settings
	MultiLanguageSettings *Flow_MultiLanguageSettings `json:"multiLanguageSettings,omitempty"`

	// Indicates whether the flow is locked for changes. If the flow is locked,
	//  modifications to the flow will be rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.locked
	Locked *bool `json:"locked,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Flow.MultiLanguageSettings
type Flow_MultiLanguageSettings struct {
	// Optional. Enable multi-language detection for this flow. This can be set
	//  only if [agent level multi language
	//  setting][Agent.enable_multi_language_training] is enabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.MultiLanguageSettings.enable_multi_language_detection
	EnableMultiLanguageDetection *bool `json:"enableMultiLanguageDetection,omitempty"`

	// Optional. Agent will respond in the detected language if the detected
	//  language code is in the supported resolved languages for this flow. This
	//  will be used only if multi-language training is enabled in the
	//  [agent][google.cloud.dialogflow.cx.v3beta1.Agent.enable_multi_language_training]
	//  and multi-language detection is enabled in the
	//  [flow][google.cloud.dialogflow.cx.v3beta1.Flow.MultiLanguageSettings.enable_multi_language_detection].
	//  The supported languages must be a subset of the languages supported by
	//  the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.MultiLanguageSettings.supported_response_language_codes
	SupportedResponseLanguageCodes []string `json:"supportedResponseLanguageCodes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.FlowInvocation
type FlowInvocation struct {
	// Required. The unique identifier of the flow.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.FlowInvocation.flow
	Flow *string `json:"flow,omitempty"`

	// Optional. A list of input parameters for the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.FlowInvocation.input_action_parameters
	InputActionParameters map[string]string `json:"inputActionParameters,omitempty"`

	// Optional. A list of output parameters generated by the flow invocation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.FlowInvocation.output_action_parameters
	OutputActionParameters map[string]string `json:"outputActionParameters,omitempty"`

	// Required. Flow invocation's output state.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.FlowInvocation.flow_state
	FlowState *string `json:"flowState,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Form
type Form struct {
	// Parameters to collect from the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.parameters
	Parameters []Form_Parameter `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Form.Parameter
type Form_Parameter struct {
	// Required. The human-readable name of the parameter, unique within the
	//  form.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates whether the parameter is required. Optional parameters will not
	//  trigger prompts; however, they are filled if the user specifies them.
	//  Required parameters must be filled before form filling concludes.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.required
	Required *bool `json:"required,omitempty"`

	// Required. The entity type of the parameter.
	//  Format:
	//  `projects/-/locations/-/agents/-/entityTypes/<SystemEntityTypeID>` for
	//  system entity types (for example,
	//  `projects/-/locations/-/agents/-/entityTypes/sys.date`), or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/entityTypes/<EntityTypeID>`
	//  for developer entity types.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.is_list
	IsList *bool `json:"isList,omitempty"`

	// Required. Defines fill behavior for the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.fill_behavior
	FillBehavior *Form_Parameter_FillBehavior `json:"fillBehavior,omitempty"`

	// The default value of an optional parameter. If the parameter is required,
	//  the default value will be ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.default_value
	DefaultValue *Value `json:"defaultValue,omitempty"`

	// Indicates whether the parameter content should be redacted in log.  If
	//  redaction is enabled, the parameter content will be replaced by parameter
	//  name during logging.
	//  Note: the parameter content is subject to redaction if either parameter
	//  level redaction or [entity type level
	//  redaction][google.cloud.dialogflow.cx.v3beta1.EntityType.redact] is
	//  enabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.redact
	Redact *bool `json:"redact,omitempty"`

	// Hierarchical advanced settings for this parameter. The settings exposed
	//  at the lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.FillBehavior
type Form_Parameter_FillBehavior struct {
	// Required. The fulfillment to provide the initial prompt that the agent
	//  can present to the user in order to fill the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.FillBehavior.initial_prompt_fulfillment
	InitialPromptFulfillment *Fulfillment `json:"initialPromptFulfillment,omitempty"`

	// The handlers for parameter-level events, used to provide reprompt for
	//  the parameter or transition to a different page/flow. The supported
	//  events are:
	//  *   `sys.no-match-<N>`, where N can be from 1 to 6
	//  *   `sys.no-match-default`
	//  *   `sys.no-input-<N>`, where N can be from 1 to 6
	//  *   `sys.no-input-default`
	//  *   `sys.invalid-parameter`
	//
	//  `initial_prompt_fulfillment` provides the first prompt for the
	//  parameter.
	//
	//  If the user's response does not fill the parameter, a
	//  no-match/no-input event will be triggered, and the fulfillment
	//  associated with the `sys.no-match-1`/`sys.no-input-1` handler (if
	//  defined) will be called to provide a prompt. The
	//  `sys.no-match-2`/`sys.no-input-2` handler (if defined) will respond to
	//  the next no-match/no-input event, and so on.
	//
	//  A `sys.no-match-default` or `sys.no-input-default` handler will be used
	//  to handle all following no-match/no-input events after all numbered
	//  no-match/no-input handlers for the parameter are consumed.
	//
	//  A `sys.invalid-parameter` handler can be defined to handle the case
	//  where the parameter values have been `invalidated` by webhook. For
	//  example, if the user's response fill the parameter, however the
	//  parameter was invalidated by webhook, the fulfillment associated with
	//  the `sys.invalid-parameter` handler (if defined) will be called to
	//  provide a prompt.
	//
	//  If the event handler for the corresponding event can't be found on the
	//  parameter, `initial_prompt_fulfillment` will be re-prompted.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Form.Parameter.FillBehavior.reprompt_event_handlers
	RepromptEventHandlers []EventHandler `json:"repromptEventHandlers,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Fulfillment
type Fulfillment struct {
	// The list of rich message responses to present to the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.messages
	Messages []ResponseMessage `json:"messages,omitempty"`

	// The webhook to call.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/webhooks/<WebhookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.webhook
	Webhook *string `json:"webhook,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response
	//  messages in streaming APIs. If a webhook is specified, it happens before
	//  Dialogflow invokes webhook.
	//  Warning:
	//  1) This flag only affects streaming API. Responses are still queued
	//  and returned once in non-streaming API.
	//  2) The flag can be enabled in any fulfillment but only the first 3 partial
	//  responses will be returned. You may only want to apply it to fulfillments
	//  that have slow webhooks.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.return_partial_responses
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty"`

	// The value of this field will be populated in the
	//  [WebhookRequest][google.cloud.dialogflow.cx.v3beta1.WebhookRequest]
	//  `fulfillmentInfo.tag` field by Dialogflow when the associated webhook is
	//  called.
	//  The tag is typically used by the webhook service to identify which
	//  fulfillment is being called, but it could be used for other purposes.
	//  This field is required if `webhook` is specified.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.tag
	Tag *string `json:"tag,omitempty"`

	// Set parameter values before executing the webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.set_parameter_actions
	SetParameterActions []Fulfillment_SetParameterAction `json:"setParameterActions,omitempty"`

	// Conditional cases for this fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.conditional_cases
	ConditionalCases []Fulfillment_ConditionalCases `json:"conditionalCases,omitempty"`

	// Hierarchical advanced settings for this fulfillment. The settings exposed
	//  at the lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// If the flag is true, the agent will utilize LLM to generate a text
	//  response. If LLM generation fails, the defined
	//  [responses][google.cloud.dialogflow.cx.v3beta1.Fulfillment.messages] in the
	//  fulfillment will be respected. This flag is only useful for fulfillments
	//  associated with no-match event handlers.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.enable_generative_fallback
	EnableGenerativeFallback *bool `json:"enableGenerativeFallback,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases
type Fulfillment_ConditionalCases struct {
	// A list of cascading if-else conditions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.cases
	Cases []Fulfillment_ConditionalCases_Case `json:"cases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case
type Fulfillment_ConditionalCases_Case struct {
	// The condition to activate and select this case. Empty means the
	//  condition is always true. The condition is evaluated against [form
	//  parameters][Form.parameters] or [session
	//  parameters][SessionInfo.parameters].
	//
	//  See the [conditions
	//  reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.condition
	Condition *string `json:"condition,omitempty"`

	// A list of case content.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.case_content
	CaseContent []Fulfillment_ConditionalCases_Case_CaseContent `json:"caseContent,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.CaseContent
type Fulfillment_ConditionalCases_Case_CaseContent struct {
	// Returned message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.CaseContent.message
	Message *ResponseMessage `json:"message,omitempty"`

	// Additional cases to be evaluated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.CaseContent.additional_cases
	AdditionalCases *Fulfillment_ConditionalCases `json:"additionalCases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Fulfillment.SetParameterAction
type Fulfillment_SetParameterAction struct {
	// Display name of the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.SetParameterAction.parameter
	Parameter *string `json:"parameter,omitempty"`

	// The new value of the parameter. A null value clears the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.SetParameterAction.value
	Value *Value `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.GcsDestination
type GcsDestination struct {
	// Required. The Google Cloud Storage URI for the exported objects. A URI is
	//  of the form: `gs://bucket/object-name-or-prefix` Whether a full object
	//  name, or just a prefix, its usage depends on the Dialogflow operation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.GcsDestination.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.GenerativeInfo
type GenerativeInfo struct {
	// The stack of [playbooks][google.cloud.dialogflow.cx.v3beta1.Playbook] that
	//  the conversation has currently entered, with the most recent one on the
	//  top.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.GenerativeInfo.current_playbooks
	CurrentPlaybooks []string `json:"currentPlaybooks,omitempty"`

	// The actions performed by the generative playbook for the current agent
	//  response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.GenerativeInfo.action_tracing_info
	ActionTracingInfo *Example `json:"actionTracingInfo,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig
type InputAudioConfig struct {
	// Required. Audio encoding of the audio content to process.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.audio_encoding
	AudioEncoding *string `json:"audioEncoding,omitempty"`

	// Sample rate (in Hertz) of the audio content sent in the query.
	//  Refer to
	//  [Cloud Speech API
	//  documentation](https://cloud.google.com/speech-to-text/docs/basics) for
	//  more details.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.sample_rate_hertz
	SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"`

	// Optional. If `true`, Dialogflow returns
	//  [SpeechWordInfo][google.cloud.dialogflow.cx.v3beta1.SpeechWordInfo] in
	//  [StreamingRecognitionResult][google.cloud.dialogflow.cx.v3beta1.StreamingRecognitionResult]
	//  with information about the recognized speech words, e.g. start and end time
	//  offsets. If false or unspecified, Speech doesn't return any word-level
	//  information.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.enable_word_info
	EnableWordInfo *bool `json:"enableWordInfo,omitempty"`

	// Optional. A list of strings containing words and phrases that the speech
	//  recognizer should recognize with higher likelihood.
	//
	//  See [the Cloud Speech
	//  documentation](https://cloud.google.com/speech-to-text/docs/basics#phrase-hints)
	//  for more details.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.phrase_hints
	PhraseHints []string `json:"phraseHints,omitempty"`

	// Optional. Which Speech model to select for the given request.
	//  For more information, see
	//  [Speech
	//  models](https://cloud.google.com/dialogflow/cx/docs/concept/speech-models).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.model
	Model *string `json:"model,omitempty"`

	// Optional. Which variant of the [Speech
	//  model][google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.model] to use.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.model_variant
	ModelVariant *string `json:"modelVariant,omitempty"`

	// Optional. If `false` (default), recognition does not cease until the
	//  client closes the stream.
	//  If `true`, the recognizer will detect a single spoken utterance in input
	//  audio. Recognition ceases when it detects the audio's voice has
	//  stopped or paused. In this case, once a detected intent is received, the
	//  client should close the stream and start a new request with a new stream as
	//  needed.
	//  Note: This setting is relevant only for streaming methods.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.single_utterance
	SingleUtterance *bool `json:"singleUtterance,omitempty"`

	// Configuration of barge-in behavior during the streaming of input audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.barge_in_config
	BargeInConfig *BargeInConfig `json:"bargeInConfig,omitempty"`

	// If `true`, the request will opt out for STT conformer model migration.
	//  This field will be deprecated once force migration takes place in June
	//  2024. Please refer to [Dialogflow CX Speech model
	//  migration](https://cloud.google.com/dialogflow/cx/docs/concept/speech-model-migration).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.InputAudioConfig.opt_out_conformer_model_migration
	OptOutConformerModelMigration *bool `json:"optOutConformerModelMigration,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent
type Intent struct {
	// The unique identifier of the intent.
	//  Required for the
	//  [Intents.UpdateIntent][google.cloud.dialogflow.cx.v3beta1.Intents.UpdateIntent]
	//  method.
	//  [Intents.CreateIntent][google.cloud.dialogflow.cx.v3beta1.Intents.CreateIntent]
	//  populates the name automatically.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/intents/<IntentID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the intent, unique within the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The collection of training phrases the agent is trained on to identify the
	//  intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.training_phrases
	TrainingPhrases []Intent_TrainingPhrase `json:"trainingPhrases,omitempty"`

	// The collection of parameters associated with the intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.parameters
	Parameters []Intent_Parameter `json:"parameters,omitempty"`

	// The priority of this intent. Higher numbers represent higher
	//  priorities.
	//
	//  - If the supplied value is unspecified or 0, the service
	//    translates the value to 500,000, which corresponds to the
	//    `Normal` priority in the console.
	//  - If the supplied value is negative, the intent is ignored
	//    in runtime detect intent requests.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.priority
	Priority *int32 `json:"priority,omitempty"`

	// Indicates whether this is a fallback intent. Currently only default
	//  fallback intent is allowed in the agent, which is added upon agent
	//  creation.
	//  Adding training phrases to fallback intent is useful in the case of
	//  requests that are mistakenly matched, since training phrases assigned to
	//  fallback intents act as negative examples that triggers no-match event.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.is_fallback
	IsFallback *bool `json:"isFallback,omitempty"`

	// The key/value metadata to label an intent. Labels can contain
	//  lowercase letters, digits and the symbols '-' and '_'. International
	//  characters are allowed, including letters from unicase alphabets. Keys must
	//  start with a letter. Keys and values can be no longer than 63 characters
	//  and no more than 128 bytes.
	//
	//  Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed
	//  Dialogflow defined labels include:
	//  * sys-head
	//  * sys-contextual
	//  The above labels do not require value. "sys-head" means the intent is a
	//  head intent. "sys-contextual" means the intent is a contextual intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Human readable description for better understanding an intent like its
	//  scope, content, result etc. Maximum character limit: 140 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter
type Intent_Parameter struct {
	// Required. The unique identifier of the parameter. This field
	//  is used by [training
	//  phrases][google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase] to
	//  annotate their
	//  [parts][google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.id
	ID *string `json:"id,omitempty"`

	// Required. The entity type of the parameter.
	//  Format:
	//  `projects/-/locations/-/agents/-/entityTypes/<SystemEntityTypeID>` for
	//  system entity types (for example,
	//  `projects/-/locations/-/agents/-/entityTypes/sys.date`), or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/entityTypes/<EntityTypeID>`
	//  for developer entity types.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.is_list
	IsList *bool `json:"isList,omitempty"`

	// Indicates whether the parameter content should be redacted in log. If
	//  redaction is enabled, the parameter content will be replaced by parameter
	//  name during logging.
	//  Note: the parameter content is subject to redaction if either parameter
	//  level redaction or [entity type level
	//  redaction][google.cloud.dialogflow.cx.v3beta1.EntityType.redact] is
	//  enabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.redact
	Redact *bool `json:"redact,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase
type Intent_TrainingPhrase struct {
	// Output only. The unique identifier of the training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.id
	ID *string `json:"id,omitempty"`

	// Required. The ordered list of training phrase parts.
	//  The parts are concatenated in order to form the training phrase.
	//
	//  Note: The API does not automatically annotate training phrases like the
	//  Dialogflow Console does.
	//
	//  Note: Do not forget to include whitespace at part boundaries, so the
	//  training phrase is well formatted when the parts are concatenated.
	//
	//  If the training phrase does not need to be annotated with parameters,
	//  you just need a single part with only the
	//  [Part.text][google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part.text]
	//  field set.
	//
	//  If you want to annotate the training phrase, you must create multiple
	//  parts, where the fields of each part are populated in one of two ways:
	//
	//  -   `Part.text` is set to a part of the phrase that has no parameters.
	//  -   `Part.text` is set to a part of the phrase that you want to annotate,
	//      and the `parameter_id` field is set.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.parts
	Parts []Intent_TrainingPhrase_Part `json:"parts,omitempty"`

	// Indicates how many times this example was added to the intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.repeat_count
	RepeatCount *int32 `json:"repeatCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part
type Intent_TrainingPhrase_Part struct {
	// Required. The text for this part.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part.text
	Text *string `json:"text,omitempty"`

	// The [parameter][google.cloud.dialogflow.cx.v3beta1.Intent.Parameter]
	//  used to annotate this part of the training phrase. This field is
	//  required for annotated parts of the training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part.parameter_id
	ParameterID *string `json:"parameterID,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.IntentInput
type IntentInput struct {
	// Required. The unique identifier of the intent.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/intents/<IntentID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.IntentInput.intent
	Intent *string `json:"intent,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.KnowledgeConnectorSettings
type KnowledgeConnectorSettings struct {
	// Whether Knowledge Connector is enabled or not.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.KnowledgeConnectorSettings.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The fulfillment to be triggered.
	//
	//  When the answers from the Knowledge Connector are selected by Dialogflow,
	//  you can utitlize the request scoped parameter `$request.knowledge.answers`
	//  (contains up to the 5 highest confidence answers) and
	//  `$request.knowledge.questions` (contains the corresponding questions) to
	//  construct the fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.KnowledgeConnectorSettings.trigger_fulfillment
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.KnowledgeConnectorSettings.target_page
	TargetPage *string `json:"targetPage,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.KnowledgeConnectorSettings.target_flow
	TargetFlow *string `json:"targetFlow,omitempty"`

	// Optional. List of related data store connections.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.KnowledgeConnectorSettings.data_store_connections
	DataStoreConnections []DataStoreConnection `json:"dataStoreConnections,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.LlmModelSettings
type LlmModelSettings struct {
	// The selected LLM model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.LlmModelSettings.model
	Model *string `json:"model,omitempty"`

	// The custom prompt to use.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.LlmModelSettings.prompt_text
	PromptText *string `json:"promptText,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Match
type Match struct {
	// The [Intent][google.cloud.dialogflow.cx.v3beta1.Intent] that matched the
	//  query. Some, not all fields are filled in this message, including but not
	//  limited to: `name` and `display_name`. Only filled for
	//  [`INTENT`][google.cloud.dialogflow.cx.v3beta1.Match.MatchType] match type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Match.intent
	Intent *Intent `json:"intent,omitempty"`

	// The event that matched the query. Filled for
	//  [`EVENT`][google.cloud.dialogflow.cx.v3beta1.Match.MatchType],
	//  [`NO_MATCH`][google.cloud.dialogflow.cx.v3beta1.Match.MatchType] and
	//  [`NO_INPUT`][google.cloud.dialogflow.cx.v3beta1.Match.MatchType] match
	//  types.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Match.event
	Event *string `json:"event,omitempty"`

	// The collection of parameters extracted from the query.
	//
	//  Depending on your protocol or client library language, this is a
	//  map, associative array, symbol table, dictionary, or JSON object
	//  composed of a collection of (MapKey, MapValue) pairs:
	//
	//  * MapKey type: string
	//  * MapKey value: parameter name
	//  * MapValue type: If parameter's entity type is a composite entity then use
	//  map, otherwise, depending on the parameter value type, it could be one of
	//  string, number, boolean, null, list or map.
	//  * MapValue value: If parameter's entity type is a composite entity then use
	//  map from composite entity property names to property values, otherwise,
	//  use parameter value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Match.parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// Final text input which was matched during MatchIntent. This value can be
	//  different from original input sent in request because of spelling
	//  correction or other processing.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Match.resolved_input
	ResolvedInput *string `json:"resolvedInput,omitempty"`

	// Type of this [Match][google.cloud.dialogflow.cx.v3beta1.Match].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Match.match_type
	MatchType *string `json:"matchType,omitempty"`

	// The confidence of this match. Values range from 0.0 (completely uncertain)
	//  to 1.0 (completely certain).
	//  This value is for informational purpose only and is only used to help match
	//  the best intent within the classification threshold. This value may change
	//  for the same end-user expression at any time due to a model retraining or
	//  change in implementation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Match.confidence
	Confidence *float32 `json:"confidence,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.NluSettings
type NluSettings struct {
	// Indicates the type of NLU model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.NluSettings.model_type
	ModelType *string `json:"modelType,omitempty"`

	// To filter out false positive results and still get variety in matched
	//  natural language inputs for your agent, you can tune the machine learning
	//  classification threshold. If the returned score value is less than the
	//  threshold value, then a no-match event will be triggered. The score values
	//  range from 0.0 (completely uncertain) to 1.0 (completely certain). If set
	//  to 0.0, the default of 0.3 is used. You can set a separate classification
	//  threshold for the flow in each language enabled for the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.NluSettings.classification_threshold
	ClassificationThreshold *float32 `json:"classificationThreshold,omitempty"`

	// Indicates NLU model training mode.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.NluSettings.model_training_mode
	ModelTrainingMode *string `json:"modelTrainingMode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.OutputAudioConfig
type OutputAudioConfig struct {
	// Required. Audio encoding of the synthesized audio content.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.OutputAudioConfig.audio_encoding
	AudioEncoding *string `json:"audioEncoding,omitempty"`

	// Optional. The synthesis sample rate (in hertz) for this audio. If not
	//  provided, then the synthesizer will use the default sample rate based on
	//  the audio encoding. If this is different from the voice's natural sample
	//  rate, then the synthesizer will honor this request by converting to the
	//  desired sample rate (which might result in worse audio quality).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.OutputAudioConfig.sample_rate_hertz
	SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"`

	// Optional. Configuration of how speech should be synthesized.
	//  If not specified,
	//  [Agent.text_to_speech_settings][google.cloud.dialogflow.cx.v3beta1.Agent.text_to_speech_settings]
	//  is applied.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.OutputAudioConfig.synthesize_speech_config
	SynthesizeSpeechConfig *SynthesizeSpeechConfig `json:"synthesizeSpeechConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Page
type Page struct {
	// The unique identifier of the page.
	//  Required for the
	//  [Pages.UpdatePage][google.cloud.dialogflow.cx.v3beta1.Pages.UpdatePage]
	//  method.
	//  [Pages.CreatePage][google.cloud.dialogflow.cx.v3beta1.Pages.CreatePage]
	//  populates the name automatically.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the page, unique within the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the page. The maximum length is 500 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.description
	Description *string `json:"description,omitempty"`

	// The fulfillment to call when the session is entering the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.entry_fulfillment
	EntryFulfillment *Fulfillment `json:"entryFulfillment,omitempty"`

	// The form associated with the page, used for collecting parameters
	//  relevant to the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.form
	Form *Form `json:"form,omitempty"`

	// Ordered list of
	//  [`TransitionRouteGroups`][google.cloud.dialogflow.cx.v3beta1.TransitionRouteGroup]
	//  added to the page. Transition route groups must be unique within a page. If
	//  the page links both flow-level transition route groups and agent-level
	//  transition route groups, the flow-level ones will have higher priority and
	//  will be put before the agent-level ones.
	//
	//  *   If multiple transition routes within a page scope refer to the same
	//      intent, then the precedence order is: page's transition route -> page's
	//      transition route group -> flow's transition routes.
	//
	//  *   If multiple transition route groups within a page contain the same
	//      intent, then the first group in the ordered list takes precedence.
	//
	//  Format:`projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  for agent-level groups.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.transition_route_groups
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	// A list of transitions for the transition rules of this page.
	//  They route the conversation to another page in the same flow, or another
	//  flow.
	//
	//  When we are in a certain page, the TransitionRoutes are evalauted in the
	//  following order:
	//
	//  *   TransitionRoutes defined in the page with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3beta1.Page.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in flow with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3beta1.Flow.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in the page with only condition specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3beta1.Page.transition_route_groups]
	//      with only condition specified.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.transition_routes
	TransitionRoutes []TransitionRoute `json:"transitionRoutes,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors,
	//  no match or no input.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.event_handlers
	EventHandlers []EventHandler `json:"eventHandlers,omitempty"`

	// Hierarchical advanced settings for this page. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Page.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.PlaybookInput
type PlaybookInput struct {
	// Optional. Summary string of the preceding conversation for the child
	//  playbook invocation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookInput.preceding_conversation_summary
	PrecedingConversationSummary *string `json:"precedingConversationSummary,omitempty"`

	// Optional. A list of input parameters for the action.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookInput.action_parameters
	ActionParameters map[string]string `json:"actionParameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.PlaybookInvocation
type PlaybookInvocation struct {
	// Required. The unique identifier of the playbook.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookInvocation.playbook
	Playbook *string `json:"playbook,omitempty"`

	// Optional. Input of the child playbook invocation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookInvocation.playbook_input
	PlaybookInput *PlaybookInput `json:"playbookInput,omitempty"`

	// Optional. Output of the child playbook invocation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookInvocation.playbook_output
	PlaybookOutput *PlaybookOutput `json:"playbookOutput,omitempty"`

	// Required. Playbook invocation's output state.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookInvocation.playbook_state
	PlaybookState *string `json:"playbookState,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.PlaybookOutput
type PlaybookOutput struct {
	// Optional. Summary string of the execution result of the child playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookOutput.execution_summary
	ExecutionSummary *string `json:"executionSummary,omitempty"`

	// Optional. A Struct object of output parameters for the action.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.PlaybookOutput.action_parameters
	ActionParameters map[string]string `json:"actionParameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.QueryInput
type QueryInput struct {
	// The natural language text to be processed.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryInput.text
	Text *TextInput `json:"text,omitempty"`

	// The intent to be triggered.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryInput.intent
	Intent *IntentInput `json:"intent,omitempty"`

	// The natural language speech audio to be processed.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryInput.audio
	Audio *AudioInput `json:"audio,omitempty"`

	// The event to be triggered.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryInput.event
	Event *EventInput `json:"event,omitempty"`

	// The DTMF event to be handled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryInput.dtmf
	Dtmf *DtmfInput `json:"dtmf,omitempty"`

	// The results of a tool executed by the client.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryInput.tool_call_result
	ToolCallResult *ToolCallResult `json:"toolCallResult,omitempty"`

	// Required. The language of the input. See [Language
	//  Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	//  for a list of the currently supported language codes. Note that queries in
	//  the same session do not necessarily need to specify the same language.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryInput.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.QueryParameters
type QueryParameters struct {
	// The time zone of this conversational query from the [time zone
	//  database](https://www.iana.org/time-zones), e.g., America/New_York,
	//  Europe/Paris. If not provided, the time zone specified in the agent is
	//  used.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// The geo location of this conversational query.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.geo_location
	GeoLocation *LatLng `json:"geoLocation,omitempty"`

	// Additional session entity types to replace or extend developer entity types
	//  with. The entity synonyms apply to all languages and persist for the
	//  session of this query.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.session_entity_types
	SessionEntityTypes []SessionEntityType `json:"sessionEntityTypes,omitempty"`

	// This field can be used to pass custom data into the webhook associated with
	//  the agent. Arbitrary JSON objects are supported.
	//  Some integrations that query a Dialogflow agent may provide additional
	//  information in the payload.
	//  In particular, for the Dialogflow Phone Gateway integration, this field has
	//  the form:
	//  ```
	//  {
	//   "telephony": {
	//     "caller_id": "+18558363987"
	//   }
	//  }
	//  ```
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.payload
	Payload map[string]string `json:"payload,omitempty"`

	// Additional parameters to be put into [session
	//  parameters][SessionInfo.parameters]. To remove a
	//  parameter from the session, clients should explicitly set the parameter
	//  value to null.
	//
	//  You can reference the session parameters in the agent with the following
	//  format: $session.params.parameter-id.
	//
	//  Depending on your protocol or client library language, this is a
	//  map, associative array, symbol table, dictionary, or JSON object
	//  composed of a collection of (MapKey, MapValue) pairs:
	//
	//  * MapKey type: string
	//  * MapKey value: parameter name
	//  * MapValue type: If parameter's entity type is a composite entity then use
	//  map, otherwise, depending on the parameter value type, it could be one of
	//  string, number, boolean, null, list or map.
	//  * MapValue value: If parameter's entity type is a composite entity then use
	//  map from composite entity property names to property values, otherwise,
	//  use parameter value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// The unique identifier of the
	//  [page][google.cloud.dialogflow.cx.v3beta1.Page] to override the [current
	//  page][QueryResult.current_page] in the session.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	//
	//   If `current_page` is specified, the previous state of the session will be
	//   ignored by Dialogflow, including the [previous
	//   page][QueryResult.current_page] and the [previous session
	//   parameters][QueryResult.parameters].
	//   In most cases,
	//   [current_page][google.cloud.dialogflow.cx.v3beta1.QueryParameters.current_page]
	//   and
	//   [parameters][google.cloud.dialogflow.cx.v3beta1.QueryParameters.parameters]
	//   should be configured together to direct a session to a specific state.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.current_page
	CurrentPage *string `json:"currentPage,omitempty"`

	// Whether to disable webhook calls for this request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.disable_webhook
	DisableWebhook *bool `json:"disableWebhook,omitempty"`

	// Configures whether sentiment analysis should be performed. If not
	//  provided, sentiment analysis is not performed.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.analyze_query_text_sentiment
	AnalyzeQueryTextSentiment *bool `json:"analyzeQueryTextSentiment,omitempty"`

	// This field can be used to pass HTTP headers for a webhook
	//  call. These headers will be sent to webhook along with the headers that
	//  have been configured through Dialogflow web console. The headers defined
	//  within this field will overwrite the headers configured through Dialogflow
	//  console if there is a conflict. Header names are case-insensitive.
	//  Google's specified headers are not allowed. Including: "Host",
	//  "Content-Length", "Connection", "From", "User-Agent", "Accept-Encoding",
	//  "If-Modified-Since", "If-None-Match", "X-Forwarded-For", etc.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.webhook_headers
	WebhookHeaders map[string]string `json:"webhookHeaders,omitempty"`

	// A list of flow versions to override for the request.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/versions/<VersionID>`.
	//
	//   If version 1 of flow X is included in this list, the traffic of
	//   flow X will go through version 1 regardless of the version configuration
	//   in the environment. Each flow can have at most one version specified in
	//   this list.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.flow_versions
	FlowVersions []string `json:"flowVersions,omitempty"`

	// Optional. Start the session with the specified
	//  [playbook][google.cloud.dialogflow.cx.v3beta1.Playbook]. You can only
	//  specify the playbook at the beginning of the session. Otherwise, an error
	//  will be thrown.
	//
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.current_playbook
	CurrentPlaybook *string `json:"currentPlaybook,omitempty"`

	// Optional. Use the specified LLM model settings for processing the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.llm_model_settings
	LlmModelSettings *LlmModelSettings `json:"llmModelSettings,omitempty"`

	// The channel which this query is for.
	//
	//  If specified, only the
	//  [ResponseMessage][google.cloud.dialogflow.cx.v3beta1.ResponseMessage]
	//  associated with the channel will be returned. If no
	//  [ResponseMessage][google.cloud.dialogflow.cx.v3beta1.ResponseMessage] is
	//  associated with the channel, it falls back to the
	//  [ResponseMessage][google.cloud.dialogflow.cx.v3beta1.ResponseMessage] with
	//  unspecified channel.
	//
	//  If unspecified, the
	//  [ResponseMessage][google.cloud.dialogflow.cx.v3beta1.ResponseMessage] with
	//  unspecified channel will be returned.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.channel
	Channel *string `json:"channel,omitempty"`

	// Optional. Configure lifetime of the Dialogflow session.
	//  By default, a Dialogflow session remains active and its data is stored for
	//  30 minutes after the last request is sent for the session.
	//  This value should be no longer than 1 day.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.session_ttl
	SessionTtl *string `json:"sessionTtl,omitempty"`

	// Optional. Information about the end-user to improve the relevance and
	//  accuracy of generative answers.
	//
	//  This will be interpreted and used by a language model, so, for good
	//  results, the data should be self-descriptive, and in a simple structure.
	//
	//  Example:
	//
	//  ```json
	//  {
	//    "subscription plan": "Business Premium Plus",
	//    "devices owned": [
	//      {"model": "Google Pixel 7"},
	//      {"model": "Google Pixel Tablet"}
	//    ]
	//  }
	//  ```
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.end_user_metadata
	EndUserMetadata map[string]string `json:"endUserMetadata,omitempty"`

	// Optional. Search configuration for UCS search queries.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.search_config
	SearchConfig *SearchConfig `json:"searchConfig,omitempty"`

	// Optional. If set to true and data stores are involved in serving the
	//  request then
	//  DetectIntentResponse.query_result.data_store_connection_signals
	//  will be filled with data that can help evaluations.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryParameters.populate_data_store_connection_signals
	PopulateDataStoreConnectionSignals *bool `json:"populateDataStoreConnectionSignals,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.QueryResult
type QueryResult struct {
	// If [natural language text][google.cloud.dialogflow.cx.v3beta1.TextInput]
	//  was provided as input, this field will contain a copy of the text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.text
	Text *string `json:"text,omitempty"`

	// If an [intent][google.cloud.dialogflow.cx.v3beta1.IntentInput] was
	//  provided as input, this field will contain a copy of the intent
	//  identifier. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/intents/<IntentID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.trigger_intent
	TriggerIntent *string `json:"triggerIntent,omitempty"`

	// If [natural language speech
	//  audio][google.cloud.dialogflow.cx.v3beta1.AudioInput] was provided as
	//  input, this field will contain the transcript for the audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.transcript
	Transcript *string `json:"transcript,omitempty"`

	// If an [event][google.cloud.dialogflow.cx.v3beta1.EventInput] was provided
	//  as input, this field will contain the name of the event.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.trigger_event
	TriggerEvent *string `json:"triggerEvent,omitempty"`

	// If a [DTMF][google.cloud.dialogflow.cx.v3beta1.DtmfInput] was provided as
	//  input, this field will contain a copy of the
	//  [DtmfInput][google.cloud.dialogflow.cx.v3beta1.DtmfInput].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.dtmf
	Dtmf *DtmfInput `json:"dtmf,omitempty"`

	// The language that was triggered during intent detection.
	//  See [Language
	//  Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	//  for a list of the currently supported language codes.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The collected [session
	//  parameters][google.cloud.dialogflow.cx.v3beta1.SessionInfo.parameters].
	//
	//  Depending on your protocol or client library language, this is a
	//  map, associative array, symbol table, dictionary, or JSON object
	//  composed of a collection of (MapKey, MapValue) pairs:
	//
	//  * MapKey type: string
	//  * MapKey value: parameter name
	//  * MapValue type: If parameter's entity type is a composite entity then use
	//  map, otherwise, depending on the parameter value type, it could be one of
	//  string, number, boolean, null, list or map.
	//  * MapValue value: If parameter's entity type is a composite entity then use
	//  map from composite entity property names to property values, otherwise,
	//  use parameter value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// The list of rich messages returned to the client. Responses vary from
	//  simple text messages to more sophisticated, structured payloads used
	//  to drive complex logic.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.response_messages
	ResponseMessages []ResponseMessage `json:"responseMessages,omitempty"`

	// The list of webhook ids in the order of call sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.webhook_ids
	WebhookIds []string `json:"webhookIds,omitempty"`

	// The list of webhook display names in the order of call sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.webhook_display_names
	WebhookDisplayNames []string `json:"webhookDisplayNames,omitempty"`

	// The list of webhook latencies in the order of call sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.webhook_latencies
	WebhookLatencies []string `json:"webhookLatencies,omitempty"`

	// The list of webhook tags in the order of call sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.webhook_tags
	WebhookTags []string `json:"webhookTags,omitempty"`

	// The list of webhook call status in the order of call sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.webhook_statuses
	WebhookStatuses []Status `json:"webhookStatuses,omitempty"`

	// The list of webhook payload in
	//  [WebhookResponse.payload][google.cloud.dialogflow.cx.v3beta1.WebhookResponse.payload],
	//  in the order of call sequence. If some webhook call fails or doesn't return
	//  any payload, an empty `Struct` would be used instead.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.webhook_payloads
	WebhookPayloads []map[string]string `json:"webhookPayloads,omitempty"`

	// The current [Page][google.cloud.dialogflow.cx.v3beta1.Page]. Some, not all
	//  fields are filled in this message, including but not limited to `name` and
	//  `display_name`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.current_page
	CurrentPage *Page `json:"currentPage,omitempty"`

	// The current [Flow][google.cloud.dialogflow.cx.v3beta1.Flow]. Some, not all
	//  fields are filled in this message, including but not limited to `name` and
	//  `display_name`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.current_flow
	CurrentFlow *Flow `json:"currentFlow,omitempty"`

	// The [Intent][google.cloud.dialogflow.cx.v3beta1.Intent] that matched the
	//  conversational query. Some, not all fields are filled in this message,
	//  including but not limited to: `name` and `display_name`. This field is
	//  deprecated, please use
	//  [QueryResult.match][google.cloud.dialogflow.cx.v3beta1.QueryResult.match]
	//  instead.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.intent
	Intent *Intent `json:"intent,omitempty"`

	// The intent detection confidence. Values range from 0.0 (completely
	//  uncertain) to 1.0 (completely certain).
	//  This value is for informational purpose only and is only used to
	//  help match the best intent within the classification threshold.
	//  This value may change for the same end-user expression at any time due to a
	//  model retraining or change in implementation.
	//  This field is deprecated, please use
	//  [QueryResult.match][google.cloud.dialogflow.cx.v3beta1.QueryResult.match]
	//  instead.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.intent_detection_confidence
	IntentDetectionConfidence *float32 `json:"intentDetectionConfidence,omitempty"`

	// Intent match result, could be an intent or an event.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.match
	Match *Match `json:"match,omitempty"`

	// The free-form diagnostic info. For example, this field could contain
	//  webhook call latency. The fields of this data can change without notice,
	//  so you should not write code that depends on its structure.
	//
	//  One of the fields is called "Alternative Matched Intents", which may
	//  aid with debugging. The following describes these intent results:
	//
	//  - The list is empty if no intent was matched to end-user input.
	//  - Only intents that are referenced in the currently active flow are
	//    included.
	//  - The matched intent is included.
	//  - Other intents that could have matched end-user input, but did not match
	//    because they are referenced by intent routes that are out of
	//    [scope](https://cloud.google.com/dialogflow/cx/docs/concept/handler#scope),
	//    are included.
	//  - Other intents referenced by intent routes in scope that matched end-user
	//    input, but had a lower confidence score.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.diagnostic_info
	DiagnosticInfo map[string]string `json:"diagnosticInfo,omitempty"`

	// The information of a query if handled by generative agent resources.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.generative_info
	GenerativeInfo *GenerativeInfo `json:"generativeInfo,omitempty"`

	// The sentiment analyss result, which depends on
	//  [`analyze_query_text_sentiment`]
	//  [google.cloud.dialogflow.cx.v3beta1.QueryParameters.analyze_query_text_sentiment],
	//  specified in the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.sentiment_analysis_result
	SentimentAnalysisResult *SentimentAnalysisResult `json:"sentimentAnalysisResult,omitempty"`

	// Returns the current advanced settings including IVR settings. Even though
	//  the operations configured by these settings are performed by Dialogflow,
	//  the client may need to perform special logic at the moment. For example, if
	//  Dialogflow exports audio to Google Cloud Storage, then the client may need
	//  to wait for the resulting object to appear in the bucket before proceeding.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Indicates whether the Thumbs up/Thumbs down rating controls are need to be
	//  shown for the response in the Dialogflow Messenger widget.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.allow_answer_feedback
	AllowAnswerFeedback *bool `json:"allowAnswerFeedback,omitempty"`

	// Optional. Data store connection feature output signals.
	//  Filled only when data stores are involved in serving the query and
	//  DetectIntentRequest.populate_data_store_connection_signals is set to true
	//  in the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.data_store_connection_signals
	DataStoreConnectionSignals *DataStoreConnectionSignals `json:"dataStoreConnectionSignals,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage
type ResponseMessage struct {
	// Returns a text response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.text
	Text *ResponseMessage_Text `json:"text,omitempty"`

	// Returns a response containing a custom, platform-specific payload.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.payload
	Payload map[string]string `json:"payload,omitempty"`

	// Indicates that the conversation succeeded.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.conversation_success
	ConversationSuccess *ResponseMessage_ConversationSuccess `json:"conversationSuccess,omitempty"`

	// A text or ssml response that is preferentially used for TTS output audio
	//  synthesis, as described in the comment on the ResponseMessage message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.output_audio_text
	OutputAudioText *ResponseMessage_OutputAudioText `json:"outputAudioText,omitempty"`

	// Hands off conversation to a human agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.live_agent_handoff
	LiveAgentHandoff *ResponseMessage_LiveAgentHandoff `json:"liveAgentHandoff,omitempty"`

	// Signal that the client should play an audio clip hosted at a
	//  client-specific URI. Dialogflow uses this to construct
	//  [mixed_audio][google.cloud.dialogflow.cx.v3beta1.ResponseMessage.mixed_audio].
	//  However, Dialogflow itself does not try to read or process the URI in any
	//  way.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.play_audio
	PlayAudio *ResponseMessage_PlayAudio `json:"playAudio,omitempty"`

	// A signal that the client should transfer the phone call connected to
	//  this agent to a third-party endpoint.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.telephony_transfer_call
	TelephonyTransferCall *ResponseMessage_TelephonyTransferCall `json:"telephonyTransferCall,omitempty"`

	// Represents info card for knowledge answers, to be better rendered in
	//  Dialogflow Messenger.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.knowledge_info_card
	KnowledgeInfoCard *ResponseMessage_KnowledgeInfoCard `json:"knowledgeInfoCard,omitempty"`

	// Returns the definition of a tool call that should be executed by the
	//  client.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.tool_call
	ToolCall *ToolCall `json:"toolCall,omitempty"`

	// The channel which the response is associated with. Clients can specify the
	//  channel via
	//  [QueryParameters.channel][google.cloud.dialogflow.cx.v3beta1.QueryParameters.channel],
	//  and only associated channel response will be returned.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.channel
	Channel *string `json:"channel,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.ConversationSuccess
type ResponseMessage_ConversationSuccess struct {
	// Custom metadata. Dialogflow doesn't impose any structure on this.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.ConversationSuccess.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.EndInteraction
type ResponseMessage_EndInteraction struct {
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.KnowledgeInfoCard
type ResponseMessage_KnowledgeInfoCard struct {
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.LiveAgentHandoff
type ResponseMessage_LiveAgentHandoff struct {
	// Custom metadata for your handoff procedure. Dialogflow doesn't impose
	//  any structure on this.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.LiveAgentHandoff.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio
type ResponseMessage_MixedAudio struct {
	// Segments this audio response is composed of.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio.segments
	Segments []ResponseMessage_MixedAudio_Segment `json:"segments,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio.Segment
type ResponseMessage_MixedAudio_Segment struct {
	// Raw audio synthesized from the Dialogflow agent's response using
	//  the output config specified in the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio.Segment.audio
	Audio []byte `json:"audio,omitempty"`

	// Client-specific URI that points to an audio clip accessible to the
	//  client. Dialogflow does not impose any validation on it.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio.Segment.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.OutputAudioText
type ResponseMessage_OutputAudioText struct {
	// The raw text to be synthesized.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.OutputAudioText.text
	Text *string `json:"text,omitempty"`

	// The SSML text to be synthesized. For more information, see
	//  [SSML](/speech/text-to-speech/docs/ssml).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.OutputAudioText.ssml
	Ssml *string `json:"ssml,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.PlayAudio
type ResponseMessage_PlayAudio struct {
	// Required. URI of the audio clip. Dialogflow does not impose any
	//  validation on this value. It is specific to the client that reads it.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.PlayAudio.audio_uri
	AudioURI *string `json:"audioURI,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.TelephonyTransferCall
type ResponseMessage_TelephonyTransferCall struct {
	// Transfer the call to a phone number
	//  in [E.164 format](https://en.wikipedia.org/wiki/E.164).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.TelephonyTransferCall.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.Text
type ResponseMessage_Text struct {
	// Required. A collection of text response variants. If multiple variants
	//  are defined, only one text response variant is returned at runtime.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.Text.text
	Text []string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.SearchConfig
type SearchConfig struct {
	// Optional. Boosting configuration for the datastores.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SearchConfig.boost_specs
	BoostSpecs []BoostSpecs `json:"boostSpecs,omitempty"`

	// Optional. Filter configuration for the datastores.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SearchConfig.filter_specs
	FilterSpecs []FilterSpecs `json:"filterSpecs,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.SentimentAnalysisResult
type SentimentAnalysisResult struct {
	// Sentiment score between -1.0 (negative sentiment) and 1.0 (positive
	//  sentiment).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SentimentAnalysisResult.score
	Score *float32 `json:"score,omitempty"`

	// A non-negative number in the [0, +inf) range, which represents the absolute
	//  magnitude of sentiment, regardless of score (positive or negative).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SentimentAnalysisResult.magnitude
	Magnitude *float32 `json:"magnitude,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.SessionEntityType
type SessionEntityType struct {
	// Required. The unique identifier of the session entity type.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/sessions/<SessionID>/entityTypes/<EntityTypeID>`
	//  or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/environments/<EnvironmentID>/sessions/<SessionID>/entityTypes/<EntityTypeID>`.
	//  If `Environment ID` is not specified, we assume default 'draft'
	//  environment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SessionEntityType.name
	Name *string `json:"name,omitempty"`

	// Required. Indicates whether the additional data should override or
	//  supplement the custom entity type definition.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SessionEntityType.entity_override_mode
	EntityOverrideMode *string `json:"entityOverrideMode,omitempty"`

	// Required. The collection of entities to override or supplement the custom
	//  entity type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SessionEntityType.entities
	Entities []EntityType_Entity `json:"entities,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.SynthesizeSpeechConfig
type SynthesizeSpeechConfig struct {
	// Optional. Speaking rate/speed, in the range [0.25, 4.0]. 1.0 is the normal
	//  native speed supported by the specific voice. 2.0 is twice as fast, and
	//  0.5 is half as fast. If unset(0.0), defaults to the native 1.0 speed. Any
	//  other values < 0.25 or > 4.0 will return an error.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SynthesizeSpeechConfig.speaking_rate
	SpeakingRate *float64 `json:"speakingRate,omitempty"`

	// Optional. Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20
	//  semitones from the original pitch. -20 means decrease 20 semitones from the
	//  original pitch.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SynthesizeSpeechConfig.pitch
	Pitch *float64 `json:"pitch,omitempty"`

	// Optional. Volume gain (in dB) of the normal native volume supported by the
	//  specific voice, in the range [-96.0, 16.0]. If unset, or set to a value of
	//  0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB)
	//  will play at approximately half the amplitude of the normal native signal
	//  amplitude. A value of +6.0 (dB) will play at approximately twice the
	//  amplitude of the normal native signal amplitude. We strongly recommend not
	//  to exceed +10 (dB) as there's usually no effective increase in loudness for
	//  any value greater than that.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SynthesizeSpeechConfig.volume_gain_db
	VolumeGainDb *float64 `json:"volumeGainDb,omitempty"`

	// Optional. An identifier which selects 'audio effects' profiles that are
	//  applied on (post synthesized) text to speech. Effects are applied on top of
	//  each other in the order they are given.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SynthesizeSpeechConfig.effects_profile_id
	EffectsProfileID []string `json:"effectsProfileID,omitempty"`

	// Optional. The desired voice of the synthesized audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.SynthesizeSpeechConfig.voice
	Voice *VoiceSelectionParams `json:"voice,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.TextInput
type TextInput struct {
	// Required. The UTF-8 encoded natural language text to be processed.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TextInput.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ToolCall
type ToolCall struct {
	// The [tool][Tool] associated with this call.
	//  Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	//  ID>/tools/<Tool ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCall.tool
	Tool *string `json:"tool,omitempty"`

	// The name of the tool's action associated with this call.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCall.action
	Action *string `json:"action,omitempty"`

	// The action's input parameters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCall.input_parameters
	InputParameters map[string]string `json:"inputParameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ToolCallResult
type ToolCallResult struct {
	// The [tool][Tool] associated with this call.
	//  Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	//  ID>/tools/<Tool ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCallResult.tool
	Tool *string `json:"tool,omitempty"`

	// The name of the tool's action associated with this call.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCallResult.action
	Action *string `json:"action,omitempty"`

	// The tool call's error.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCallResult.error
	Error *ToolCallResult_Error `json:"error,omitempty"`

	// The tool call's output parameters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCallResult.output_parameters
	OutputParameters map[string]string `json:"outputParameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ToolCallResult.Error
type ToolCallResult_Error struct {
	// The error message of the function.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolCallResult.Error.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ToolUse
type ToolUse struct {
	// Required. The [tool][google.cloud.dialogflow.cx.v3beta1.Tool] that should
	//  be used. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/tools/<ToolID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolUse.tool
	Tool *string `json:"tool,omitempty"`

	// Optional. Name of the action to be called during the tool use.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolUse.action
	Action *string `json:"action,omitempty"`

	// Optional. A list of input parameters for the action.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolUse.input_action_parameters
	InputActionParameters map[string]string `json:"inputActionParameters,omitempty"`

	// Optional. A list of output parameters generated by the action.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ToolUse.output_action_parameters
	OutputActionParameters map[string]string `json:"outputActionParameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.TransitionRoute
type TransitionRoute struct {

	// Optional. The description of the transition route. The maximum length is
	//  500 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.description
	Description *string `json:"description,omitempty"`

	// The unique identifier of an
	//  [Intent][google.cloud.dialogflow.cx.v3beta1.Intent]. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/intents/<IntentID>`.
	//  Indicates that the transition can only happen when the given intent is
	//  matched.
	//  At least one of `intent` or `condition` must be specified. When both
	//  `intent` and `condition` are specified, the transition can only happen
	//  when both are fulfilled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.intent
	Intent *string `json:"intent,omitempty"`

	// The condition to evaluate against [form
	//  parameters][google.cloud.dialogflow.cx.v3beta1.Form.parameters] or [session
	//  parameters][google.cloud.dialogflow.cx.v3beta1.SessionInfo.parameters].
	//
	//  See the [conditions
	//  reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	//  At least one of `intent` or `condition` must be specified. When both
	//  `intent` and `condition` are specified, the transition can only happen
	//  when both are fulfilled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.condition
	Condition *string `json:"condition,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of
	//  `trigger_fulfillment` and `target` must be specified. When both are
	//  defined, `trigger_fulfillment` is executed first.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.trigger_fulfillment
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.target_page
	TargetPage *string `json:"targetPage,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.target_flow
	TargetFlow *string `json:"targetFlow,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.UserUtterance
type UserUtterance struct {
	// Required. Message content in text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.UserUtterance.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.VoiceSelectionParams
type VoiceSelectionParams struct {
	// Optional. The name of the voice. If not set, the service will choose a
	//  voice based on the other parameters such as language_code and
	//  [ssml_gender][google.cloud.dialogflow.cx.v3beta1.VoiceSelectionParams.ssml_gender].
	//
	//  For the list of available voices, please refer to [Supported voices and
	//  languages](https://cloud.google.com/text-to-speech/docs/voices).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VoiceSelectionParams.name
	Name *string `json:"name,omitempty"`

	// Optional. The preferred gender of the voice. If not set, the service will
	//  choose a voice based on the other parameters such as language_code and
	//  [name][google.cloud.dialogflow.cx.v3beta1.VoiceSelectionParams.name]. Note
	//  that this is only a preference, not requirement. If a voice of the
	//  appropriate gender is not available, the synthesizer should substitute a
	//  voice with a different gender rather than failing the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VoiceSelectionParams.ssml_gender
	SsmlGender *string `json:"ssmlGender,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook
type Webhook struct {
	// The unique identifier of the webhook.
	//  Required for the
	//  [Webhooks.UpdateWebhook][google.cloud.dialogflow.cx.v3beta1.Webhooks.UpdateWebhook]
	//  method.
	//  [Webhooks.CreateWebhook][google.cloud.dialogflow.cx.v3beta1.Webhooks.CreateWebhook]
	//  populates the name automatically. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/webhooks/<WebhookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the webhook, unique within the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration for a generic web service.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.generic_web_service
	GenericWebService *Webhook_GenericWebService `json:"genericWebService,omitempty"`

	// Configuration for a [Service
	//  Directory](https://cloud.google.com/service-directory) service.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.service_directory
	ServiceDirectory *Webhook_ServiceDirectoryConfig `json:"serviceDirectory,omitempty"`

	// Webhook execution timeout. Execution is considered failed if Dialogflow
	//  doesn't receive a response from webhook at the end of the timeout period.
	//  Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Indicates whether the webhook is disabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService
type Webhook_GenericWebService struct {
	// Required. The webhook URI for receiving POST requests. It must use https
	//  protocol.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.uri
	URI *string `json:"uri,omitempty"`

	// The user name for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.username
	Username *string `json:"username,omitempty"`

	// The password for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.password
	Password *string `json:"password,omitempty"`

	// The HTTP request headers to send together with webhook
	//  requests.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.request_headers
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// Optional. Specifies a list of allowed custom CA certificates (in DER
	//  format) for HTTPS verification. This overrides the default SSL trust
	//  store. If this is empty or unspecified, Dialogflow will use Google's
	//  default trust store to verify certificates. N.B. Make sure the HTTPS
	//  server certificates are signed with "subject alt name". For instance a
	//  certificate can be self-signed using the following command,
	//  ```
	//     openssl x509 -req -days 200 -in example.com.csr \
	//       -signkey example.com.key \
	//       -out example.com.crt \
	//       -extfile <(printf "\nsubjectAltName='DNS:www.example.com'")
	//  ```
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.allowed_ca_certs
	AllowedCaCerts [][]byte `json:"allowedCaCerts,omitempty"`

	// Optional. The OAuth configuration of the webhook. If specified,
	//  Dialogflow will initiate the OAuth client credential flow to exchange an
	//  access token from the 3rd party platform and put it in the auth header.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.oauth_config
	OauthConfig *Webhook_GenericWebService_OAuthConfig `json:"oauthConfig,omitempty"`

	// Optional. Indicate the auth token type generated from the [Diglogflow
	//  service
	//  agent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent).
	//  The generated token is sent in the Authorization header.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.service_agent_auth
	ServiceAgentAuth *string `json:"serviceAgentAuth,omitempty"`

	// Optional. Type of the webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.webhook_type
	WebhookType *string `json:"webhookType,omitempty"`

	// Optional. HTTP method for the flexible webhook calls. Standard webhook
	//  always uses POST.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.http_method
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// Optional. Defines a custom JSON object as request body to send to
	//  flexible webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.request_body
	RequestBody *string `json:"requestBody,omitempty"`

	// Optional. Maps the values extracted from specific fields of the flexible
	//  webhook response into session parameters.
	//  - Key: session parameter name
	//  - Value: field path in the webhook response
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.parameter_mapping
	ParameterMapping map[string]string `json:"parameterMapping,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig
type Webhook_GenericWebService_OAuthConfig struct {
	// Required. The client ID provided by the 3rd party platform.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Required. The client secret provided by the 3rd party platform.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.client_secret
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Required. The token endpoint provided by the 3rd party platform to
	//  exchange an access token.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.token_endpoint
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`

	// Optional. The OAuth scopes to grant.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook.ServiceDirectoryConfig
type Webhook_ServiceDirectoryConfig struct {
	// Required. The name of [Service
	//  Directory](https://cloud.google.com/service-directory) service.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/namespaces/<NamespaceID>/services/<ServiceID>`.
	//  `Location ID` of the service directory must be the same as the location
	//  of the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.ServiceDirectoryConfig.service
	Service *string `json:"service,omitempty"`

	// Generic Service configuration of this webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.ServiceDirectoryConfig.generic_web_service
	GenericWebService *Webhook_GenericWebService `json:"genericWebService,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.type.LatLng
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	// +kcc:proto:field=google.type.LatLng.latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	// +kcc:proto:field=google.type.LatLng.longitude
	Longitude *float64 `json:"longitude,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation
type ConversationObservedState struct {
	// All the [Flow][google.cloud.dialogflow.cx.v3beta1.Flow] the conversation
	//  has went through. Only `name` and `display_name` are filled in this
	//  message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.flows
	Flows []FlowObservedState `json:"flows,omitempty"`

	// Interactions of the conversation.
	//  Only populated for `GetConversation` and empty for `ListConversations`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.interactions
	Interactions []Conversation_InteractionObservedState `json:"interactions,omitempty"`

	// Environment of the conversation.
	//  Only `name` and `display_name` are filled in this message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.environment
	Environment *EnvironmentObservedState `json:"environment,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction
type Conversation_InteractionObservedState struct {
	// The final response of the interaction.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Conversation.Interaction.response
	Response *DetectIntentResponseObservedState `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse
type DetectIntentResponseObservedState struct {
	// The result of the conversational query.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse.query_result
	QueryResult *QueryResultObservedState `json:"queryResult,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Environment
type EnvironmentObservedState struct {
	// Output only. Update time of this environment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Environment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.EventHandler
type EventHandlerObservedState struct {
	// Output only. The unique identifier of this event handler.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EventHandler.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Example
type ExampleObservedState struct {
	// Output only. Estimated number of tokes current example takes when sent to
	//  the LLM.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.token_count
	TokenCount *int64 `json:"tokenCount,omitempty"`

	// Output only. The timestamp of initial example creation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last time the example was updated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Example.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Flow
type FlowObservedState struct {
	// A flow's transition routes serve two purposes:
	//
	//  *   They are responsible for matching the user's first utterances in the
	//  flow.
	//  *   They are inherited by every page's [transition
	//  routes][Page.transition_routes] and can support use cases such as the user
	//  saying "help" or "can I talk to a human?", which can be handled in a common
	//  way regardless of the current page. Transition routes defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  TransitionRoutes are evalauted in the following order:
	//
	//  *   TransitionRoutes with intent specified.
	//  *   TransitionRoutes with only condition specified.
	//
	//  TransitionRoutes with intent specified are inherited by pages in the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.transition_routes
	TransitionRoutes []TransitionRouteObservedState `json:"transitionRoutes,omitempty"`

	// A flow's event handlers serve two purposes:
	//
	//  *   They are responsible for handling events (e.g. no match,
	//  webhook errors) in the flow.
	//  *   They are inherited by every page's [event
	//  handlers][Page.event_handlers], which can be used to handle common events
	//  regardless of the current page. Event handlers defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  Unlike
	//  [transition_routes][google.cloud.dialogflow.cx.v3beta1.Flow.transition_routes],
	//  these handlers are evaluated on a first-match basis. The first one that
	//  matches the event get executed, with the rest being ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Flow.event_handlers
	EventHandlers []EventHandlerObservedState `json:"eventHandlers,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Fulfillment
type FulfillmentObservedState struct {
	// The list of rich message responses to present to the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Fulfillment.messages
	Messages []ResponseMessageObservedState `json:"messages,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.GenerativeInfo
type GenerativeInfoObservedState struct {
	// The actions performed by the generative playbook for the current agent
	//  response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.GenerativeInfo.action_tracing_info
	ActionTracingInfo *ExampleObservedState `json:"actionTracingInfo,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.QueryResult
type QueryResultObservedState struct {
	// The information of a query if handled by generative agent resources.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.QueryResult.generative_info
	GenerativeInfo *GenerativeInfoObservedState `json:"generativeInfo,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage
type ResponseMessageObservedState struct {
	// Returns a text response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.text
	Text *ResponseMessage_TextObservedState `json:"text,omitempty"`

	// A text or ssml response that is preferentially used for TTS output audio
	//  synthesis, as described in the comment on the ResponseMessage message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.output_audio_text
	OutputAudioText *ResponseMessage_OutputAudioTextObservedState `json:"outputAudioText,omitempty"`

	// Output only. A signal that indicates the interaction with the Dialogflow
	//  agent has ended. This message is generated by Dialogflow only when the
	//  conversation reaches `END_SESSION` page. It is not supposed to be defined
	//  by the user.
	//
	//  It's guaranteed that there is at most one such message in each response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.end_interaction
	EndInteraction *ResponseMessage_EndInteraction `json:"endInteraction,omitempty"`

	// Signal that the client should play an audio clip hosted at a
	//  client-specific URI. Dialogflow uses this to construct
	//  [mixed_audio][google.cloud.dialogflow.cx.v3beta1.ResponseMessage.mixed_audio].
	//  However, Dialogflow itself does not try to read or process the URI in any
	//  way.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.play_audio
	PlayAudio *ResponseMessage_PlayAudioObservedState `json:"playAudio,omitempty"`

	// Output only. An audio response message composed of both the synthesized
	//  Dialogflow agent responses and responses defined via
	//  [play_audio][google.cloud.dialogflow.cx.v3beta1.ResponseMessage.play_audio].
	//  This message is generated by Dialogflow only and not supposed to be
	//  defined by the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.mixed_audio
	MixedAudio *ResponseMessage_MixedAudio `json:"mixedAudio,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio
type ResponseMessage_MixedAudioObservedState struct {
	// Segments this audio response is composed of.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio.segments
	Segments []ResponseMessage_MixedAudio_SegmentObservedState `json:"segments,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio.Segment
type ResponseMessage_MixedAudio_SegmentObservedState struct {
	// Output only. Whether the playback of this segment can be interrupted by
	//  the end user's speech and the client should then start the next
	//  Dialogflow request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.MixedAudio.Segment.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.OutputAudioText
type ResponseMessage_OutputAudioTextObservedState struct {
	// Output only. Whether the playback of this message can be interrupted by
	//  the end user's speech and the client can then starts the next Dialogflow
	//  request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.OutputAudioText.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.PlayAudio
type ResponseMessage_PlayAudioObservedState struct {
	// Output only. Whether the playback of this message can be interrupted by
	//  the end user's speech and the client can then starts the next Dialogflow
	//  request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.PlayAudio.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.Text
type ResponseMessage_TextObservedState struct {
	// Output only. Whether the playback of this message can be interrupted by
	//  the end user's speech and the client can then starts the next Dialogflow
	//  request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ResponseMessage.Text.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.TransitionRoute
type TransitionRouteObservedState struct {
	// Output only. The unique identifier of this transition route.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.name
	Name *string `json:"name,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of
	//  `trigger_fulfillment` and `target` must be specified. When both are
	//  defined, `trigger_fulfillment` is executed first.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.TransitionRoute.trigger_fulfillment
	TriggerFulfillment *FulfillmentObservedState `json:"triggerFulfillment,omitempty"`
}
