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

// API sources for DialogflowPlaybook, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/playbook.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.playbooks

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowPlaybookGVK = GroupVersion.WithKind("DialogflowPlaybook")

// DialogflowPlaybookSpec defines the desired state of DialogflowPlaybook
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Playbook
// +kcc:required-from-proto
type DialogflowPlaybookSpec struct {

	// A reference to the projects/{project}/locations/{location}/agents/{agent} this resource belongs to.
	// +required
	AgentRef *DialogflowCXAgentRef `json:"agentRef"`

	// The DialogflowPlaybook name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the playbook, unique within an agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. High level description of the goal the playbook intend to
	//  accomplish. A goal should be concise since it's visible to other playbooks
	//  that may reference this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.goal
	// +required
	Goal *string `json:"goal,omitempty"`

	// Optional. Defined structured input parameters for this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.input_parameter_definitions
	InputParameterDefinitions []ParameterDefinition `json:"inputParameterDefinitions,omitempty"`

	// Optional. Defined structured output parameters for this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.output_parameter_definitions
	OutputParameterDefinitions []ParameterDefinition `json:"outputParameterDefinitions,omitempty"`

	// Instruction to accomplish target goal.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.instruction
	Instruction *Playbook_Instruction `json:"instruction,omitempty"`

	// Optional. The resource name of tools referenced by the current playbook in
	//  the instructions. If not provided explicitly, they are will
	//  be implied using the tool being referenced in goal and steps.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.referenced_tools
	ReferencedTools []string `json:"referencedTools,omitempty"`

	// Optional. The playbook's scoped code block, which may implement handlers
	//  and actions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.code_block
	CodeBlock *CodeBlock `json:"codeBlock,omitempty"`

	// Optional. Llm model settings for the playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.llm_model_settings
	LlmModelSettings *LlmModelSettings `json:"llmModelSettings,omitempty"`

	// Optional. A list of registered handlers to execuate based on the specified
	//  triggers.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.handlers
	Handlers []Handler `json:"handlers,omitempty"`

	// Optional. Type of the playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.playbook_type
	PlaybookType *string `json:"playbookType,omitempty"`
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

	// A list of Generators to be called during this fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.generators
	Generators []Fulfillment_GeneratorSettings `json:"generators,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.GeneratorSettings
type Fulfillment_GeneratorSettings struct {
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

// DialogflowPlaybookStatus defines the config connector machine state of DialogflowPlaybook
type DialogflowPlaybookStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowPlaybook resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowPlaybookObservedState `json:"observedState,omitempty"`
}

// DialogflowPlaybookObservedState is the state of the DialogflowPlaybook resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Playbook
type DialogflowPlaybookObservedState struct {
	// Output only. Estimated number of tokes current playbook takes when sent to
	//  the LLM.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.token_count
	TokenCount *int64 `json:"tokenCount,omitempty"`

	// Output only. The timestamp of initial playbook creation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last time the playbook version was updated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The resource name of other playbooks referenced by the current
	//  playbook in the instructions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.referenced_playbooks
	ReferencedPlaybooks []string `json:"referencedPlaybooks,omitempty"`

	// Output only. The resource name of flows referenced by the current playbook
	//  in the instructions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.referenced_flows
	ReferencedFlows []string `json:"referencedFlows,omitempty"`

	// Optional. Output only. Names of inline actions scoped to this playbook.
	//  These actions are in addition to those belonging to referenced tools, child
	//  playbooks, and flows, e.g. actions that are defined in the playbook's code
	//  block.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.inline_actions
	InlineActions []string `json:"inlineActions,omitempty"`

	// Optional. A list of registered handlers to execuate based on the specified
	//  triggers.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.handlers
	Handlers []HandlerObservedState `json:"handlers,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowplaybook;gcpdialogflowplaybooks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowPlaybook is the Schema for the DialogflowPlaybook API
// +k8s:openapi-gen=true
type DialogflowPlaybook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowPlaybookSpec   `json:"spec,omitempty"`
	Status DialogflowPlaybookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowPlaybookList contains a list of DialogflowPlaybook
type DialogflowPlaybookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowPlaybook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowPlaybook{}, &DialogflowPlaybookList{})
}
