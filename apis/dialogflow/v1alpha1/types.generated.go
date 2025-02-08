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


// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings
type AdvancedSettings struct {
	// If present, incoming audio is exported by Dialogflow to the configured
	//  Google Cloud Storage destination.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.audio_export_gcs_destination
	AudioExportGcsDestination *GcsDestination `json:"audioExportGcsDestination,omitempty"`

	// Settings for speech to text detection.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	//  - Page level
	//  - Parameter level
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.speech_settings
	SpeechSettings *AdvancedSettings_SpeechSettings `json:"speechSettings,omitempty"`

	// Settings for DTMF.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	//  - Page level
	//  - Parameter level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.dtmf_settings
	DtmfSettings *AdvancedSettings_DtmfSettings `json:"dtmfSettings,omitempty"`

	// Settings for logging.
	//  Settings for Dialogflow History, Contact Center messages, StackDriver logs,
	//  and speech logging.
	//  Exposed at the following levels:
	//  - Agent level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.logging_settings
	LoggingSettings *AdvancedSettings_LoggingSettings `json:"loggingSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings
type AdvancedSettings_DtmfSettings struct {
	// If true, incoming audio is processed for DTMF (dual tone multi frequency)
	//  events. For example, if the caller presses a button on their telephone
	//  keypad and DTMF processing is enabled, Dialogflow will detect the
	//  event (e.g. a "3" was pressed) in the incoming audio and pass the event
	//  to the bot to drive business logic (e.g. when 3 is pressed, return the
	//  account balance).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Max length of DTMF digits.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.max_digits
	MaxDigits *int32 `json:"maxDigits,omitempty"`

	// The digit that terminates a DTMF digit sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.finish_digit
	FinishDigit *string `json:"finishDigit,omitempty"`

	// Interdigit timeout setting for matching dtmf input to regex.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.interdigit_timeout_duration
	InterdigitTimeoutDuration *string `json:"interdigitTimeoutDuration,omitempty"`

	// Endpoint timeout setting for matching dtmf input to regex.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.endpointing_timeout_duration
	EndpointingTimeoutDuration *string `json:"endpointingTimeoutDuration,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings
type AdvancedSettings_LoggingSettings struct {
	// Enables Google Cloud Logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings.enable_stackdriver_logging
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	// Enables DF Interaction logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings.enable_interaction_logging
	EnableInteractionLogging *bool `json:"enableInteractionLogging,omitempty"`

	// Enables consent-based end-user input redaction, if true, a pre-defined
	//  session parameter `$session.params.conversation-redaction` will be
	//  used to determine if the utterance should be redacted.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings.enable_consent_based_redaction
	EnableConsentBasedRedaction *bool `json:"enableConsentBasedRedaction,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings
type AdvancedSettings_SpeechSettings struct {
	// Sensitivity of the speech model that detects the end of speech.
	//  Scale from 0 to 100.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.endpointer_sensitivity
	EndpointerSensitivity *int32 `json:"endpointerSensitivity,omitempty"`

	// Timeout before detecting no speech.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.no_speech_timeout
	NoSpeechTimeout *string `json:"noSpeechTimeout,omitempty"`

	// Use timeout based endpointing, interpreting endpointer sensitivy as
	//  seconds of timeout value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.use_timeout_based_endpointing
	UseTimeoutBasedEndpointing *bool `json:"useTimeoutBasedEndpointing,omitempty"`

	// Mapping from language to Speech-to-Text model. The mapped Speech-to-Text
	//  model will be selected for requests from its corresponding language.
	//  For more information, see
	//  [Speech
	//  models](https://cloud.google.com/dialogflow/cx/docs/concept/speech-models).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.models
	Models map[string]string `json:"models,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.DataStoreConnection
type DataStoreConnection struct {
	// The type of the connected data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// The full name of the referenced data store.
	//  Formats:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	//  `projects/{project}/locations/{location}/dataStores/{data_store}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.data_store
	DataStore *string `json:"dataStore,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.EventHandler
type EventHandler struct {

	// Required. The name of the event to handle.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.event
	Event *string `json:"event,omitempty"`

	// The fulfillment to call when the event occurs.
	//  Handling webhook errors with a fulfillment enabled with webhook could
	//  cause infinite loop. It is invalid to specify such fulfillment for a
	//  handler handling webhooks.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.trigger_fulfillment
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.target_page
	TargetPage *string `json:"targetPage,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.target_flow
	TargetFlow *string `json:"targetFlow,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Flow
type Flow struct {
	// The unique identifier of the flow.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the flow. The maximum length is 500 characters. If
	//  exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.description
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
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.transition_routes
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
	//  [transition_routes][google.cloud.dialogflow.cx.v3.Flow.transition_routes],
	//  these handlers are evaluated on a first-match basis. The first one that
	//  matches the event get executed, with the rest being ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.event_handlers
	EventHandlers []EventHandler `json:"eventHandlers,omitempty"`

	// A flow's transition route group serve two purposes:
	//
	//  *   They are responsible for matching the user's first utterances in the
	//  flow.
	//  *   They are inherited by every page's [transition
	//  route groups][Page.transition_route_groups]. Transition route groups
	//  defined in the page have higher priority than those defined in the flow.
	//
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  for agent-level groups.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.transition_route_groups
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	// NLU related settings of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.nlu_settings
	NluSettings *NluSettings `json:"nluSettings,omitempty"`

	// Hierarchical advanced settings for this flow. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`

	// Optional. Multi-lingual agent settings for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.multi_language_settings
	MultiLanguageSettings *Flow_MultiLanguageSettings `json:"multiLanguageSettings,omitempty"`

	// Indicates whether the flow is locked for changes. If the flow is locked,
	//  modifications to the flow will be rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.locked
	Locked *bool `json:"locked,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Flow.MultiLanguageSettings
type Flow_MultiLanguageSettings struct {
	// Optional. Enable multi-language detection for this flow. This can be set
	//  only if [agent level multi language
	//  setting][Agent.enable_multi_language_training] is enabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.MultiLanguageSettings.enable_multi_language_detection
	EnableMultiLanguageDetection *bool `json:"enableMultiLanguageDetection,omitempty"`

	// Optional. Agent will respond in the detected language if the detected
	//  language code is in the supported resolved languages for this flow. This
	//  will be used only if multi-language training is enabled in the
	//  [agent][google.cloud.dialogflow.cx.v3.Agent.enable_multi_language_training]
	//  and multi-language detection is enabled in the
	//  [flow][google.cloud.dialogflow.cx.v3.Flow.MultiLanguageSettings.enable_multi_language_detection].
	//  The supported languages must be a subset of the languages supported by
	//  the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.MultiLanguageSettings.supported_response_language_codes
	SupportedResponseLanguageCodes []string `json:"supportedResponseLanguageCodes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment
type Fulfillment struct {
	// The list of rich message responses to present to the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.messages
	Messages []ResponseMessage `json:"messages,omitempty"`

	// The webhook to call.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/webhooks/<WebhookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.webhook
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
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.return_partial_responses
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty"`

	// The value of this field will be populated in the
	//  [WebhookRequest][google.cloud.dialogflow.cx.v3.WebhookRequest]
	//  `fulfillmentInfo.tag` field by Dialogflow when the associated webhook is
	//  called.
	//  The tag is typically used by the webhook service to identify which
	//  fulfillment is being called, but it could be used for other purposes.
	//  This field is required if `webhook` is specified.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.tag
	Tag *string `json:"tag,omitempty"`

	// Set parameter values before executing the webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.set_parameter_actions
	SetParameterActions []Fulfillment_SetParameterAction `json:"setParameterActions,omitempty"`

	// Conditional cases for this fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.conditional_cases
	ConditionalCases []Fulfillment_ConditionalCases `json:"conditionalCases,omitempty"`

	// Hierarchical advanced settings for this fulfillment. The settings exposed
	//  at the lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// If the flag is true, the agent will utilize LLM to generate a text
	//  response. If LLM generation fails, the defined
	//  [responses][google.cloud.dialogflow.cx.v3.Fulfillment.messages] in the
	//  fulfillment will be respected. This flag is only useful for fulfillments
	//  associated with no-match event handlers.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.enable_generative_fallback
	EnableGenerativeFallback *bool `json:"enableGenerativeFallback,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases
type Fulfillment_ConditionalCases struct {
	// A list of cascading if-else conditions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.cases
	Cases []Fulfillment_ConditionalCases_Case `json:"cases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case
type Fulfillment_ConditionalCases_Case struct {
	// The condition to activate and select this case. Empty means the
	//  condition is always true. The condition is evaluated against [form
	//  parameters][Form.parameters] or [session
	//  parameters][SessionInfo.parameters].
	//
	//  See the [conditions
	//  reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.condition
	Condition *string `json:"condition,omitempty"`

	// A list of case content.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.case_content
	CaseContent []Fulfillment_ConditionalCases_Case_CaseContent `json:"caseContent,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent
type Fulfillment_ConditionalCases_Case_CaseContent struct {
	// Returned message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent.message
	Message *ResponseMessage `json:"message,omitempty"`

	// Additional cases to be evaluated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent.additional_cases
	AdditionalCases *Fulfillment_ConditionalCases `json:"additionalCases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.SetParameterAction
type Fulfillment_SetParameterAction struct {
	// Display name of the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.SetParameterAction.parameter
	Parameter *string `json:"parameter,omitempty"`

	// The new value of the parameter. A null value clears the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.SetParameterAction.value
	Value *Value `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.GcsDestination
type GcsDestination struct {
	// Required. The Google Cloud Storage URI for the exported objects. A URI is
	//  of the form: `gs://bucket/object-name-or-prefix` Whether a full object
	//  name, or just a prefix, its usage depends on the Dialogflow operation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GcsDestination.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings
type KnowledgeConnectorSettings struct {
	// Whether Knowledge Connector is enabled or not.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The fulfillment to be triggered.
	//
	//  When the answers from the Knowledge Connector are selected by Dialogflow,
	//  you can utitlize the request scoped parameter `$request.knowledge.answers`
	//  (contains up to the 5 highest confidence answers) and
	//  `$request.knowledge.questions` (contains the corresponding questions) to
	//  construct the fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.trigger_fulfillment
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.target_page
	TargetPage *string `json:"targetPage,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.target_flow
	TargetFlow *string `json:"targetFlow,omitempty"`

	// Optional. List of related data store connections.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.data_store_connections
	DataStoreConnections []DataStoreConnection `json:"dataStoreConnections,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.NluSettings
type NluSettings struct {
	// Indicates the type of NLU model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.NluSettings.model_type
	ModelType *string `json:"modelType,omitempty"`

	// To filter out false positive results and still get variety in matched
	//  natural language inputs for your agent, you can tune the machine learning
	//  classification threshold. If the returned score value is less than the
	//  threshold value, then a no-match event will be triggered. The score values
	//  range from 0.0 (completely uncertain) to 1.0 (completely certain). If set
	//  to 0.0, the default of 0.3 is used. You can set a separate classification
	//  threshold for the flow in each language enabled for the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.NluSettings.classification_threshold
	ClassificationThreshold *float32 `json:"classificationThreshold,omitempty"`

	// Indicates NLU model training mode.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.NluSettings.model_training_mode
	ModelTrainingMode *string `json:"modelTrainingMode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage
type ResponseMessage struct {
	// Returns a text response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.text
	Text *ResponseMessage_Text `json:"text,omitempty"`

	// Returns a response containing a custom, platform-specific payload.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.payload
	Payload map[string]string `json:"payload,omitempty"`

	// Indicates that the conversation succeeded.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.conversation_success
	ConversationSuccess *ResponseMessage_ConversationSuccess `json:"conversationSuccess,omitempty"`

	// A text or ssml response that is preferentially used for TTS output audio
	//  synthesis, as described in the comment on the ResponseMessage message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.output_audio_text
	OutputAudioText *ResponseMessage_OutputAudioText `json:"outputAudioText,omitempty"`

	// Hands off conversation to a human agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.live_agent_handoff
	LiveAgentHandoff *ResponseMessage_LiveAgentHandoff `json:"liveAgentHandoff,omitempty"`

	// Signal that the client should play an audio clip hosted at a
	//  client-specific URI. Dialogflow uses this to construct
	//  [mixed_audio][google.cloud.dialogflow.cx.v3.ResponseMessage.mixed_audio].
	//  However, Dialogflow itself does not try to read or process the URI in any
	//  way.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.play_audio
	PlayAudio *ResponseMessage_PlayAudio `json:"playAudio,omitempty"`

	// A signal that the client should transfer the phone call connected to
	//  this agent to a third-party endpoint.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.telephony_transfer_call
	TelephonyTransferCall *ResponseMessage_TelephonyTransferCall `json:"telephonyTransferCall,omitempty"`

	// Represents info card for knowledge answers, to be better rendered in
	//  Dialogflow Messenger.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.knowledge_info_card
	KnowledgeInfoCard *ResponseMessage_KnowledgeInfoCard `json:"knowledgeInfoCard,omitempty"`

	// Response type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.response_type
	ResponseType *string `json:"responseType,omitempty"`

	// The channel which the response is associated with. Clients can specify the
	//  channel via
	//  [QueryParameters.channel][google.cloud.dialogflow.cx.v3.QueryParameters.channel],
	//  and only associated channel response will be returned.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.channel
	Channel *string `json:"channel,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.ConversationSuccess
type ResponseMessage_ConversationSuccess struct {
	// Custom metadata. Dialogflow doesn't impose any structure on this.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.ConversationSuccess.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.EndInteraction
type ResponseMessage_EndInteraction struct {
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.KnowledgeInfoCard
type ResponseMessage_KnowledgeInfoCard struct {
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.LiveAgentHandoff
type ResponseMessage_LiveAgentHandoff struct {
	// Custom metadata for your handoff procedure. Dialogflow doesn't impose
	//  any structure on this.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.LiveAgentHandoff.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio
type ResponseMessage_MixedAudio struct {
	// Segments this audio response is composed of.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio.segments
	Segments []ResponseMessage_MixedAudio_Segment `json:"segments,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio.Segment
type ResponseMessage_MixedAudio_Segment struct {
	// Raw audio synthesized from the Dialogflow agent's response using
	//  the output config specified in the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio.Segment.audio
	Audio []byte `json:"audio,omitempty"`

	// Client-specific URI that points to an audio clip accessible to the
	//  client. Dialogflow does not impose any validation on it.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio.Segment.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.OutputAudioText
type ResponseMessage_OutputAudioText struct {
	// The raw text to be synthesized.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.OutputAudioText.text
	Text *string `json:"text,omitempty"`

	// The SSML text to be synthesized. For more information, see
	//  [SSML](/speech/text-to-speech/docs/ssml).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.OutputAudioText.ssml
	Ssml *string `json:"ssml,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.PlayAudio
type ResponseMessage_PlayAudio struct {
	// Required. URI of the audio clip. Dialogflow does not impose any
	//  validation on this value. It is specific to the client that reads it.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.PlayAudio.audio_uri
	AudioURI *string `json:"audioURI,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.TelephonyTransferCall
type ResponseMessage_TelephonyTransferCall struct {
	// Transfer the call to a phone number
	//  in [E.164 format](https://en.wikipedia.org/wiki/E.164).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.TelephonyTransferCall.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.Text
type ResponseMessage_Text struct {
	// Required. A collection of text response variants. If multiple variants
	//  are defined, only one text response variant is returned at runtime.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.Text.text
	Text []string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.TransitionRoute
type TransitionRoute struct {

	// Optional. The description of the transition route. The maximum length is
	//  500 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.description
	Description *string `json:"description,omitempty"`

	// The unique identifier of an [Intent][google.cloud.dialogflow.cx.v3.Intent].
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/intents/<IntentID>`.
	//  Indicates that the transition can only happen when the given intent is
	//  matched.
	//  At least one of `intent` or `condition` must be specified. When both
	//  `intent` and `condition` are specified, the transition can only happen
	//  when both are fulfilled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.intent
	Intent *string `json:"intent,omitempty"`

	// The condition to evaluate against [form
	//  parameters][google.cloud.dialogflow.cx.v3.Form.parameters] or [session
	//  parameters][google.cloud.dialogflow.cx.v3.SessionInfo.parameters].
	//
	//  See the [conditions
	//  reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	//  At least one of `intent` or `condition` must be specified. When both
	//  `intent` and `condition` are specified, the transition can only happen
	//  when both are fulfilled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.condition
	Condition *string `json:"condition,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of
	//  `trigger_fulfillment` and `target` must be specified. When both are
	//  defined, `trigger_fulfillment` is executed first.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.trigger_fulfillment
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.target_page
	TargetPage *string `json:"targetPage,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.target_flow
	TargetFlow *string `json:"targetFlow,omitempty"`
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

// +kcc:proto=google.cloud.dialogflow.cx.v3.EventHandler
type EventHandlerObservedState struct {
	// Output only. The unique identifier of this event handler.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Flow
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
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.transition_routes
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
	//  [transition_routes][google.cloud.dialogflow.cx.v3.Flow.transition_routes],
	//  these handlers are evaluated on a first-match basis. The first one that
	//  matches the event get executed, with the rest being ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.event_handlers
	EventHandlers []EventHandlerObservedState `json:"eventHandlers,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment
type FulfillmentObservedState struct {
	// The list of rich message responses to present to the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.messages
	Messages []ResponseMessageObservedState `json:"messages,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage
type ResponseMessageObservedState struct {
	// Returns a text response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.text
	Text *ResponseMessage_TextObservedState `json:"text,omitempty"`

	// A text or ssml response that is preferentially used for TTS output audio
	//  synthesis, as described in the comment on the ResponseMessage message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.output_audio_text
	OutputAudioText *ResponseMessage_OutputAudioTextObservedState `json:"outputAudioText,omitempty"`

	// Output only. A signal that indicates the interaction with the Dialogflow
	//  agent has ended. This message is generated by Dialogflow only when the
	//  conversation reaches `END_SESSION` page. It is not supposed to be defined
	//  by the user.
	//
	//  It's guaranteed that there is at most one such message in each response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.end_interaction
	EndInteraction *ResponseMessage_EndInteraction `json:"endInteraction,omitempty"`

	// Signal that the client should play an audio clip hosted at a
	//  client-specific URI. Dialogflow uses this to construct
	//  [mixed_audio][google.cloud.dialogflow.cx.v3.ResponseMessage.mixed_audio].
	//  However, Dialogflow itself does not try to read or process the URI in any
	//  way.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.play_audio
	PlayAudio *ResponseMessage_PlayAudioObservedState `json:"playAudio,omitempty"`

	// Output only. An audio response message composed of both the synthesized
	//  Dialogflow agent responses and responses defined via
	//  [play_audio][google.cloud.dialogflow.cx.v3.ResponseMessage.play_audio].
	//  This message is generated by Dialogflow only and not supposed to be
	//  defined by the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.mixed_audio
	MixedAudio *ResponseMessage_MixedAudio `json:"mixedAudio,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio
type ResponseMessage_MixedAudioObservedState struct {
	// Segments this audio response is composed of.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio.segments
	Segments []ResponseMessage_MixedAudio_SegmentObservedState `json:"segments,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio.Segment
type ResponseMessage_MixedAudio_SegmentObservedState struct {
	// Output only. Whether the playback of this segment can be interrupted by
	//  the end user's speech and the client should then start the next
	//  Dialogflow request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.MixedAudio.Segment.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.OutputAudioText
type ResponseMessage_OutputAudioTextObservedState struct {
	// Output only. Whether the playback of this message can be interrupted by
	//  the end user's speech and the client can then starts the next Dialogflow
	//  request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.OutputAudioText.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.PlayAudio
type ResponseMessage_PlayAudioObservedState struct {
	// Output only. Whether the playback of this message can be interrupted by
	//  the end user's speech and the client can then starts the next Dialogflow
	//  request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.PlayAudio.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.Text
type ResponseMessage_TextObservedState struct {
	// Output only. Whether the playback of this message can be interrupted by
	//  the end user's speech and the client can then starts the next Dialogflow
	//  request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.Text.allow_playback_interruption
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.TransitionRoute
type TransitionRouteObservedState struct {
	// Output only. The unique identifier of this transition route.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.name
	Name *string `json:"name,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of
	//  `trigger_fulfillment` and `target` must be specified. When both are
	//  defined, `trigger_fulfillment` is executed first.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.trigger_fulfillment
	TriggerFulfillment *FulfillmentObservedState `json:"triggerFulfillment,omitempty"`
}
