// Copyright 2026 Google LLC
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

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowCXFlowGVK = GroupVersion.WithKind("DialogflowCXFlow")

// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Flow
type DialogflowCXFlowSpec struct {
	// Immutable. The agent to create a flow for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `json:"parent,omitempty"`

	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `json:"description,omitempty"`

	// The human-readable name of the flow.
	// +required
	DisplayName string `json:"displayName"`

	// A flow's event handlers serve two purposes:
	// They are responsible for handling events (e.g. no match, webhook errors) in the flow.
	// They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
	// Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers []DialogflowCXFlowEventHandlers `json:"eventHandlers,omitempty"`

	// Immutable. The language of the following fields in flow:
	// Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases
	// Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `json:"languageCode,omitempty"`

	// NLU related settings of the flow.
	NluSettings *DialogflowCXFlowNluSettings `json:"nluSettings,omitempty"`

	// A flow's transition route group serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	// A flow's transition routes serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
	TransitionRoutes []DialogflowCXFlowTransitionRoutes `json:"transitionRoutes,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.EventHandler
type DialogflowCXFlowEventHandlers struct {
	// The name of the event to handle.
	Event *string `json:"event,omitempty"`

	// The unique identifier of this event handler.
	Name *string `json:"name,omitempty"`

	// The target flow to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	TargetFlow *string `json:"targetFlow,omitempty"`

	// The target page to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	TargetPage *string `json:"targetPage,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	TriggerFulfillment *DialogflowCXFlowTriggerFulfillment `json:"triggerFulfillment,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment
type DialogflowCXFlowTriggerFulfillment struct {
	// Conditional cases for this fulfillment.
	ConditionalCases []DialogflowCXFlowConditionalCases `json:"conditionalCases,omitempty"`

	// The list of rich message responses to present to the user.
	Messages []DialogflowCXFlowMessages `json:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty"`

	// Set parameter values before executing the webhook.
	SetParameterActions []DialogflowCXFlowSetParameterActions `json:"setParameterActions,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	Tag *string `json:"tag,omitempty"`

	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	Webhook *string `json:"webhook,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases
type DialogflowCXFlowConditionalCases struct {
	// A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.
	// See Case for the schema.
	Cases *string `json:"cases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage
type DialogflowCXFlowMessages struct {
	// The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.
	Channel *string `json:"channel,omitempty"`

	// Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.
	ConversationSuccess *DialogflowCXFlowConversationSuccess `json:"conversationSuccess,omitempty"`

	// Indicates that the conversation should be handed off to a live agent.
	LiveAgentHandoff *DialogflowCXFlowLiveAgentHandoff `json:"liveAgentHandoff,omitempty"`

	// A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.
	OutputAudioText *DialogflowCXFlowOutputAudioText `json:"outputAudioText,omitempty"`

	// A custom, platform-specific payload.
	Payload *string `json:"payload,omitempty"`

	// Specifies an audio clip to be played by the client as part of the response.
	PlayAudio *DialogflowCXFlowPlayAudio `json:"playAudio,omitempty"`

	// Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.
	TelephonyTransferCall *DialogflowCXFlowTelephonyTransferCall `json:"telephonyTransferCall,omitempty"`

	// The text response message.
	Text *DialogflowCXFlowText `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.ConversationSuccess
type DialogflowCXFlowConversationSuccess struct {
	// Custom metadata. Dialogflow doesn't impose any structure on this.
	Metadata *string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.LiveAgentHandoff
type DialogflowCXFlowLiveAgentHandoff struct {
	// Custom metadata. Dialogflow doesn't impose any structure on this.
	Metadata *string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.OutputAudioText
type DialogflowCXFlowOutputAudioText struct {
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`

	// The SSML text to be synthesized. For more information, see SSML.
	Ssml *string `json:"ssml,omitempty"`

	// The raw text to be synthesized.
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.PlayAudio
type DialogflowCXFlowPlayAudio struct {
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`

	// URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.
	// +required
	AudioUri string `json:"audioUri"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.TelephonyTransferCall
type DialogflowCXFlowTelephonyTransferCall struct {
	// Transfer the call to a phone number in E.164 format.
	// +required
	PhoneNumber string `json:"phoneNumber"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResponseMessage.Text
type DialogflowCXFlowText struct {
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`

	// A collection of text responses.
	Text []string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.SetParameterAction
type DialogflowCXFlowSetParameterActions struct {
	// Display name of the parameter.
	Parameter *string `json:"parameter,omitempty"`

	// The new JSON-encoded value of the parameter. A null value clears the parameter.
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.NluSettings
type DialogflowCXFlowNluSettings struct {
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
	// If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold *float64 `json:"classificationThreshold,omitempty"`

	// Indicates NLU model training mode.
	// * MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
	// * MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train. Possible values: ["MODEL_TRAINING_MODE_AUTOMATIC", "MODEL_TRAINING_MODE_MANUAL"].
	ModelTrainingMode *string `json:"modelTrainingMode,omitempty"`

	// Indicates the type of NLU model.
	// * MODEL_TYPE_STANDARD: Use standard NLU model.
	// * MODEL_TYPE_ADVANCED: Use advanced NLU model. Possible values: ["MODEL_TYPE_STANDARD", "MODEL_TYPE_ADVANCED"].
	ModelType *string `json:"modelType,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.TransitionRoute
type DialogflowCXFlowTransitionRoutes struct {
	// The condition to evaluate against form parameters or session parameters.
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	Condition *string `json:"condition,omitempty"`

	// The unique identifier of an Intent.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	Intent *string `json:"intent,omitempty"`

	// The unique identifier of this transition route.
	Name *string `json:"name,omitempty"`

	// The target flow to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	TargetFlow *string `json:"targetFlow,omitempty"`

	// The target page to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	TargetPage *string `json:"targetPage,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.
	TriggerFulfillment *DialogflowCXFlowTriggerFulfillment `json:"triggerFulfillment,omitempty"`
}

type DialogflowCXFlowStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The unique identifier of the flow.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowCXFlow resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// DialogflowCXFlowObservedState is the state of the DialogflowCXFlow resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Flow
type DialogflowCXFlowObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxflow;gcpdialogflowcxflows
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCXFlow is the Schema for the DialogflowCXFlow API
// +k8s:openapi-gen=true
type DialogflowCXFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCXFlowSpec   `json:"spec,omitempty"`
	Status DialogflowCXFlowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCXFlowList contains a list of DialogflowCXFlow
type DialogflowCXFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXFlow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXFlow{}, &DialogflowCXFlowList{})
}
