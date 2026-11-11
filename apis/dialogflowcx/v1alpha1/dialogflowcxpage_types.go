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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowCXPageGVK = GroupVersion.WithKind("DialogflowCXPage")

// +kcc:proto=google.cloud.dialogflow.cx.v3.Page
type DialogflowCXPageSpec struct {
	/* The DialogflowCXPage name. If not given, the metadata.name will be used. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The human-readable name of the page, unique within the agent.
	// +required
	DisplayName string `json:"displayName"`

	// The fulfillment to call when the session is entering the page.
	// +optional
	EntryFulfillment *Fulfillment `json:"entryFulfillment,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// +optional
	EventHandlers []EventHandler `json:"eventHandlers,omitempty"`

	// The form associated with the page, used for collecting parameters relevant to the page.
	// +optional
	Form *Form `json:"form,omitempty"`

	// Immutable. The language of the following fields in page:
	//
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	// +optional
	LanguageCode *string `json:"languageCode,omitempty"`

	// Immutable. The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +optional
	Parent *string `json:"parent,omitempty"`

	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	// +optional
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// +optional
	TransitionRoutes []TransitionRoute `json:"transitionRoutes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment
type Fulfillment struct {
	// Conditional cases for this fulfillment.
	// +optional
	ConditionalCases []Fulfillment_ConditionalCases `json:"conditionalCases,omitempty"`

	// The list of rich message responses to present to the user.
	// +optional
	Messages []ResponseMessage `json:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty"`

	// Set parameter values before executing the webhook.
	// +optional
	SetParameterActions []Fulfillment_SetParameterAction `json:"setParameterActions,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +optional
	Tag *string `json:"tag,omitempty"`

	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	// +optional
	Webhook *string `json:"webhook,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases
type Fulfillment_ConditionalCases struct {
	// A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.
	// See [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.
	// +optional
	Cases *string `json:"cases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage
type ResponseMessage struct {
	// The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.
	// +optional
	Channel *string `json:"channel,omitempty"`

	// Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.
	// Dialogflow only uses this to determine which conversations should be counted as successful and doesn't process the metadata in this message in any way. Note that Dialogflow also considers conversations that get to the conversation end page as successful even if they don't return ConversationSuccess.
	// You may set this, for example:
	// * In the entryFulfillment of a Page if entering the page indicates that the conversation succeeded.
	// * In a webhook response when you determine that you handled the customer issue.
	// +optional
	ConversationSuccess *ResponseMessage_ConversationSuccess `json:"conversationSuccess,omitempty"`

	// Indicates that the conversation should be handed off to a live agent.
	// Dialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures.
	// You may set this, for example:
	// * In the entryFulfillment of a Page if entering the page indicates something went extremely wrong in the conversation.
	// * In a webhook response when you determine that the customer issue can only be handled by a human.
	// +optional
	LiveAgentHandoff *ResponseMessage_LiveAgentHandoff `json:"liveAgentHandoff,omitempty"`

	// A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.
	// +optional
	OutputAudioText *ResponseMessage_OutputAudioText `json:"outputAudioText,omitempty"`

	// A custom, platform-specific payload.
	// +optional
	Payload *string `json:"payload,omitempty"`

	// Specifies an audio clip to be played by the client as part of the response.
	// +optional
	PlayAudio *ResponseMessage_PlayAudio `json:"playAudio,omitempty"`

	// Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.
	// +optional
	TelephonyTransferCall *ResponseMessage_TelephonyTransferCall `json:"telephonyTransferCall,omitempty"`

	// The text response message.
	// +optional
	Text *ResponseMessage_Text `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.ConversationSuccess
type ResponseMessage_ConversationSuccess struct {
	// Custom metadata. Dialogflow doesn't impose any structure on this.
	// +optional
	Metadata *string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.LiveAgentHandoff
type ResponseMessage_LiveAgentHandoff struct {
	// Custom metadata. Dialogflow doesn't impose any structure on this.
	// +optional
	Metadata *string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.OutputAudioText
type ResponseMessage_OutputAudioText struct {
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	// +optional
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`

	// The SSML text to be synthesized. For more information, see SSML.
	// +optional
	Ssml *string `json:"ssml,omitempty"`

	// The raw text to be synthesized.
	// +optional
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.PlayAudio
type ResponseMessage_PlayAudio struct {
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	// +optional
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`

	// URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.
	// +required
	AudioUri string `json:"audioUri"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.TelephonyTransferCall
type ResponseMessage_TelephonyTransferCall struct {
	// Transfer the call to a phone number in E.164 format.
	// +required
	PhoneNumber string `json:"phoneNumber"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.Text
type ResponseMessage_Text struct {
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	// +optional
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`

	// A collection of text responses.
	// +optional
	Text []string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.SetParameterAction
type Fulfillment_SetParameterAction struct {
	// Display name of the parameter.
	// +optional
	Parameter *string `json:"parameter,omitempty"`

	// The new JSON-encoded value of the parameter. A null value clears the parameter.
	// +optional
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.EventHandler
type EventHandler struct {
	// The name of the event to handle.
	// +optional
	Event *string `json:"event,omitempty"`

	// The unique identifier of this event handler.
	// +optional
	Name *string `json:"name,omitempty"`

	// The target flow to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +optional
	TargetFlow *string `json:"targetFlow,omitempty"`

	// The target page to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	// +optional
	TargetPage *string `json:"targetPage,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// +optional
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form
type Form struct {
	// Parameters to collect from the user.
	// +optional
	Parameters []Form_Parameter `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form.Parameter
type Form_Parameter struct {
	// The default value of the parameter.
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty"`

	// The human-readable name of the parameter, unique within the form.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// The entity type of the parameter.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<EntityType ID>.
	// +optional
	EntityType *string `json:"entityType,omitempty"`

	// Defines fill behavior for the parameter.
	// +optional
	FillBehavior *Form_Parameter_FillBehavior `json:"fillBehavior,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +optional
	IsList *bool `json:"isList,omitempty"`

	// Indicates whether the parameter content should be redacted in log.
	// If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
	// +optional
	Redact *bool `json:"redact,omitempty"`

	// Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them.
	// Required parameters must be filled before form filling concludes.
	// +optional
	Required *bool `json:"required,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Form.Parameter.FillBehavior
type Form_Parameter_FillBehavior struct {
	// The fulfillment to call when the parameter is requested.
	// +optional
	InitialPromptFulfillment *Fulfillment `json:"initialPromptFulfillment,omitempty"`

	// The event handlers for the parameter. If specified, when an event occurs during parameter filling, the handler's trigger_fulfillment will be executed.
	// +optional
	RepromptEventHandlers []EventHandler `json:"repromptEventHandlers,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.TransitionRoute
type TransitionRoute struct {
	// The condition to evaluate against form parameters or session parameters.
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +optional
	Condition *string `json:"condition,omitempty"`

	// The unique identifier of an Intent.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +optional
	Intent *string `json:"intent,omitempty"`

	// The unique identifier of this transition route.
	// +optional
	Name *string `json:"name,omitempty"`

	// The target flow to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +optional
	TargetFlow *string `json:"targetFlow,omitempty"`

	// The target page to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	// +optional
	TargetPage *string `json:"targetPage,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.
	// +optional
	TriggerFulfillment *Fulfillment `json:"triggerFulfillment,omitempty"`
}

type DialogflowCXPageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The unique identifier of the page.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	// +optional
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowCXPage resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// DialogflowCXPageObservedState is the state of the DialogflowCXPage resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Page
type DialogflowCXPageObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxpage;gcpdialogflowcxpages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCXPage is the Schema for the DialogflowCXPage API
// +k8s:openapi-gen=true
type DialogflowCXPage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCXPageSpec   `json:"spec,omitempty"`
	Status DialogflowCXPageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCXPageList contains a list of DialogflowCXPage
type DialogflowCXPageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXPage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXPage{}, &DialogflowCXPageList{})
}
