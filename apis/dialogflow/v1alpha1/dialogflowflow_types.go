// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for DialogflowFlow, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/flow.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows

package v1alpha1

import (
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowFlowGVK = GroupVersion.WithKind("DialogflowFlow")

// DialogflowFlowSpec defines the desired state of DialogflowFlow
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Flow
// +kcc:required-from-proto
type DialogflowFlowSpec struct {
	// The DialogflowCXAgent that this resource belongs to.
	// +required
	AgentRef *DialogflowCXAgentRef `json:"agentRef,omitempty"`

	// The DialogflowFlow name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.display_name
	// +required
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
	//  TransitionRoutes are evaluated in the following order:
	//
	//  *   TransitionRoutes with intent specified.
	//  *   TransitionRoutes with only condition specified.
	//
	//  TransitionRoutes with intent specified are inherited by pages in the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.transition_routes
	TransitionRoutes []DialogflowCX_TransitionRoute `json:"transitionRoutes,omitempty"`

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
	EventHandlers []DialogflowCX_EventHandler `json:"eventHandlers,omitempty"`

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
	TransitionRouteGroupRefs []DialogflowTransitionRouteGroupRef `json:"transitionRouteGroupRefs,omitempty"`

	// NLU related settings of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.nlu_settings
	NluSettings *NluSettings `json:"nluSettings,omitempty"`

	// Hierarchical advanced settings for this flow. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.knowledge_connector_settings
	KnowledgeConnectorSettings *DialogflowCX_KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`

	// Optional. Defined structured input parameters for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.input_parameter_definitions
	InputParameterDefinitions []ParameterDefinition `json:"inputParameterDefinitions,omitempty"`

	// Optional. Defined structured output parameters for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.output_parameter_definitions
	OutputParameterDefinitions []ParameterDefinition `json:"outputParameterDefinitions,omitempty"`

	// Optional. Multi-lingual agent settings for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.multi_language_settings
	MultiLanguageSettings *Flow_MultiLanguageSettings `json:"multiLanguageSettings,omitempty"`

	// Indicates whether the flow is locked for changes. If the flow is locked,
	//  modifications to the flow will be rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.locked
	Locked *bool `json:"locked,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.TransitionRoute
type DialogflowCX_TransitionRoute struct {
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
	IntentRef *DialogflowCXIntentRef `json:"intentRef,omitempty"`

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
	TriggerFulfillment *DialogflowCX_Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.target_page
	TargetPageRef *DialogflowPageRef `json:"targetPageRef,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRoute.target_flow
	TargetFlowRef *DialogflowFlowRef `json:"targetFlowRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.EventHandler
type DialogflowCX_EventHandler struct {
	// Required. The name of the event to handle.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.event
	// +required
	Event *string `json:"event,omitempty"`

	// The fulfillment to call when the event occurs.
	//  Handling webhook errors with a fulfillment enabled with webhook could
	//  cause infinite loop. It is invalid to specify such fulfillment for a
	//  handler handling webhooks.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.trigger_fulfillment
	TriggerFulfillment *DialogflowCX_Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.target_page
	TargetPageRef *DialogflowPageRef `json:"targetPageRef,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.target_flow
	TargetFlowRef *DialogflowFlowRef `json:"targetFlowRef,omitempty"`

	// The target playbook to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EventHandler.target_playbook
	TargetPlaybookRef *DialogflowPlaybookRef `json:"targetPlaybookRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings
type DialogflowCX_KnowledgeConnectorSettings struct {
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
	TriggerFulfillment *DialogflowCX_Fulfillment `json:"triggerFulfillment,omitempty"`

	// The target page to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.target_page
	TargetPageRef *DialogflowPageRef `json:"targetPageRef,omitempty"`

	// The target flow to transition to.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.target_flow
	TargetFlowRef *DialogflowFlowRef `json:"targetFlowRef,omitempty"`

	// Optional. List of related data store connections.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.KnowledgeConnectorSettings.data_store_connections
	DataStoreConnections []DialogflowCX_DataStoreConnection `json:"dataStoreConnections,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.DataStoreConnection
type DialogflowCX_DataStoreConnection struct {
	// The type of the connected data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// The full name of the referenced data store.
	//  Formats:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	//  `projects/{project}/locations/{location}/dataStores/{data_store}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.data_store
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// The document processing mode for the data store connection. Should only be
	//  set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is
	//  considered as DOCUMENTS, as this is the legacy mode.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.document_processing_mode
	DocumentProcessingMode *string `json:"documentProcessingMode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment
type DialogflowCX_Fulfillment struct {
	// The list of rich message responses to present to the user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.messages
	Messages []DialogflowCX_ResponseMessage `json:"messages,omitempty"`

	// The webhook to call.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/webhooks/<WebhookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.webhook
	WebhookRef *DialogflowWebhookRef `json:"webhookRef,omitempty"`

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

	// Optional. The name of the code block function to execute, if this is a code
	//  block fulfillment. The code block itself is implied by the fulfillment's
	//  parent, e.g. a playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.code_block_function
	CodeBlockFunction *string `json:"codeBlockFunction,omitempty"`

	// Set parameter values before executing the webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.set_parameter_actions
	SetParameterActions []Fulfillment_SetParameterAction `json:"setParameterActions,omitempty"`

	// Conditional cases for this fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.conditional_cases
	ConditionalCases []DialogflowCX_Fulfillment_ConditionalCases `json:"conditionalCases,omitempty"`

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

	// A list of Generators to be called during this fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.generators
	Generators []DialogflowCX_Fulfillment_GeneratorSettings `json:"generators,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases
type DialogflowCX_Fulfillment_ConditionalCases struct {
	// A list of cascading if-else conditions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.cases
	Cases []DialogflowCX_Fulfillment_ConditionalCases_Case `json:"cases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case
type DialogflowCX_Fulfillment_ConditionalCases_Case struct {
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
	CaseContent []DialogflowCX_Fulfillment_ConditionalCases_Case_CaseContent `json:"caseContent,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent
type DialogflowCX_Fulfillment_ConditionalCases_Case_CaseContent struct {
	// Returned message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent.message
	Message *DialogflowCX_ResponseMessage `json:"message,omitempty"`

	// Additional cases to be evaluated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent.additional_cases
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	AdditionalCases *DialogflowCX_Fulfillment_ConditionalCases `json:"additionalCases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.GeneratorSettings
type DialogflowCX_Fulfillment_GeneratorSettings struct {
	// Required. The generator to call.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/generators/<GeneratorID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.GeneratorSettings.generator
	// +required
	GeneratorRef *DialogflowGeneratorRef `json:"generatorRef,omitempty"`

	// Map from [placeholder parameter][Generator.Parameter.id] in the
	//  [Generator][google.cloud.dialogflow.cx.v3.Generator] to corresponding
	//  session parameters. By default, Dialogflow uses the session parameter
	//  with the same name to fill in the generator template. e.g. If there is a
	//  placeholder parameter `city` in the Generator, Dialogflow default to fill
	//  in the `$city` with
	//  `$session.params.city`. However, you may choose to fill `$city` with
	//  `$session.params.desination-city`.
	//  - Map key: [parameter ID][Genrator.Parameter.id]
	//  - Map value: session parameter name
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.GeneratorSettings.input_parameters
	InputParameters map[string]string `json:"inputParameters,omitempty"`

	// Required. Output parameter which should contain the generator response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.GeneratorSettings.output_parameter
	// +required
	OutputParameter *string `json:"outputParameter,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage
type DialogflowCX_ResponseMessage struct {
	// Returns a text response.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.text
	Text *ResponseMessage_Text `json:"text,omitempty"`

	// Returns a response containing a custom, platform-specific payload.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.payload
	Payload apiextensionsv1.JSON `json:"payload,omitempty"`

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

	// Returns the definition of a tool call that should be executed by the
	//  client.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResponseMessage.tool_call
	ToolCall *DialogflowCX_ToolCall `json:"toolCall,omitempty"`

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

// +kcc:proto=google.cloud.dialogflow.cx.v3.ToolCall
type DialogflowCX_ToolCall struct {
	// Required. The [tool][google.cloud.dialogflow.cx.v3.Tool] associated with
	//  this call. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/tools/<ToolID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ToolCall.tool
	// +required
	ToolRef *DialogflowToolRef `json:"toolRef,omitempty"`

	// Required. The name of the tool's action associated with this call.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ToolCall.action
	// +required
	Action *string `json:"action,omitempty"`

	// Optional. The action's input parameters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ToolCall.input_parameters
	InputParameters apiextensionsv1.JSON `json:"inputParameters,omitempty"`
}

// DialogflowFlowStatus defines the config connector machine state of DialogflowFlow
type DialogflowFlowStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowFlow resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowFlowObservedState `json:"observedState,omitempty"`
}

// DialogflowFlowObservedState is the state of the DialogflowFlow resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Flow
type DialogflowFlowObservedState struct {
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
	//  TransitionRoutes are evaluated in the following order:
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

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettingsObservedState `json:"knowledgeConnectorSettings,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowflow;gcpdialogflowflows
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowFlow is the Schema for the DialogflowFlow API
// +k8s:openapi-gen=true
type DialogflowFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowFlowSpec   `json:"spec,omitempty"`
	Status DialogflowFlowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowFlowList contains a list of DialogflowFlow
type DialogflowFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowFlow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowFlow{}, &DialogflowFlowList{})
}
