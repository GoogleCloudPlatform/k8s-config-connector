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

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form
type Form struct {
	// Parameters to collect from the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.parameters
	Parameters []Form_Parameter `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form.Parameter
type Form_Parameter struct {
	// Required. The human-readable name of the parameter, unique within the
	//  form.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates whether the parameter is required. Optional parameters will not
	//  trigger prompts; however, they are filled if the user specifies them.
	//  Required parameters must be filled before form filling concludes.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.required
	Required *bool `json:"required,omitempty"`

	// Required. The entity type of the parameter.
	//  Format:
	//  `projects/-/locations/-/agents/-/entityTypes/<SystemEntityTypeID>` for
	//  system entity types (for example,
	//  `projects/-/locations/-/agents/-/entityTypes/sys.date`), or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/entityTypes/<EntityTypeID>`
	//  for developer entity types.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.is_list
	IsList *bool `json:"isList,omitempty"`

	// Required. Defines fill behavior for the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.fill_behavior
	FillBehavior *Form_Parameter_FillBehavior `json:"fillBehavior,omitempty"`

	// The default value of an optional parameter. If the parameter is required,
	//  the default value will be ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.default_value
	DefaultValue *Value `json:"defaultValue,omitempty"`

	// Indicates whether the parameter content should be redacted in log.  If
	//  redaction is enabled, the parameter content will be replaced by parameter
	//  name during logging.
	//  Note: the parameter content is subject to redaction if either parameter
	//  level redaction or [entity type level
	//  redaction][google.cloud.dialogflow.cx.v3.EntityType.redact] is enabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.redact
	Redact *bool `json:"redact,omitempty"`

	// Hierarchical advanced settings for this parameter. The settings exposed
	//  at the lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form.Parameter.FillBehavior
type Form_Parameter_FillBehavior struct {
	// Required. The fulfillment to provide the initial prompt that the agent
	//  can present to the user in order to fill the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.FillBehavior.initial_prompt_fulfillment
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
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.FillBehavior.reprompt_event_handlers
	RepromptEventHandlers []EventHandler `json:"repromptEventHandlers,omitempty"`
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

// +kcc:proto=google.cloud.dialogflow.cx.v3.Page
type Page struct {
	// The unique identifier of the page.
	//  Required for the
	//  [Pages.UpdatePage][google.cloud.dialogflow.cx.v3.Pages.UpdatePage] method.
	//  [Pages.CreatePage][google.cloud.dialogflow.cx.v3.Pages.CreatePage]
	//  populates the name automatically.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the page, unique within the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the page. The maximum length is 500 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.description
	Description *string `json:"description,omitempty"`

	// The fulfillment to call when the session is entering the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.entry_fulfillment
	EntryFulfillment *Fulfillment `json:"entryFulfillment,omitempty"`

	// The form associated with the page, used for collecting parameters
	//  relevant to the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.form
	Form *Form `json:"form,omitempty"`

	// Ordered list of
	//  [`TransitionRouteGroups`][google.cloud.dialogflow.cx.v3.TransitionRouteGroup]
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
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.transition_route_groups
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
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in flow with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Flow.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in the page with only condition specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with only condition specified.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.transition_routes
	TransitionRoutes []TransitionRoute `json:"transitionRoutes,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors,
	//  no match or no input.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.event_handlers
	EventHandlers []EventHandler `json:"eventHandlers,omitempty"`

	// Hierarchical advanced settings for this page. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`
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

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form
type FormObservedState struct {
	// Parameters to collect from the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.parameters
	Parameters []Form_ParameterObservedState `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form.Parameter
type Form_ParameterObservedState struct {
	// Required. Defines fill behavior for the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.fill_behavior
	FillBehavior *Form_Parameter_FillBehaviorObservedState `json:"fillBehavior,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form.Parameter.FillBehavior
type Form_Parameter_FillBehaviorObservedState struct {
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
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Form.Parameter.FillBehavior.reprompt_event_handlers
	RepromptEventHandlers []EventHandlerObservedState `json:"repromptEventHandlers,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment
type FulfillmentObservedState struct {
	// The list of rich message responses to present to the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.messages
	Messages []ResponseMessageObservedState `json:"messages,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Page
type PageObservedState struct {
	// The fulfillment to call when the session is entering the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.entry_fulfillment
	EntryFulfillment *FulfillmentObservedState `json:"entryFulfillment,omitempty"`

	// The form associated with the page, used for collecting parameters
	//  relevant to the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.form
	Form *FormObservedState `json:"form,omitempty"`

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
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in flow with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Flow.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in the page with only condition specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with only condition specified.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.transition_routes
	TransitionRoutes []TransitionRouteObservedState `json:"transitionRoutes,omitempty"`
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
}
