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

// API sources for CesAgent, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1526e545e9d26f23b9c5d0f04af17297def8d045/google/cloud/ces/v1/agent.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/reference/rest/v1/projects.locations.apps.agents

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CesAgentGVK = GroupVersion.WithKind("CesAgent")

// CesAgentSpec defines the desired state of CesAgent
// +kcc:spec:proto=google.cloud.ces.v1.Agent
// +kcc:required-from-proto
type CesAgentSpec struct {

	// A reference to the projects/{project}/locations/{location}/apps/{app} this resource belongs to.
	// +required
	AppRef *CesAppRef `json:"appRef"`

	// The CesAgent name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The default agent type.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.llm_agent
	LlmAgent *Agent_LlmAgent `json:"llmAgent,omitempty"`

	// Optional. The remote
	//  [Dialogflow](https://cloud.google.com/dialogflow/cx/docs/concept/console-conversational-agents)
	//  agent to be used for the agent execution. If this field is set, all other
	//  agent level properties will be ignored.
	//
	//  Note: If the Dialogflow agent is in a different project from the app, you
	//  should grant `roles/dialogflow.client` to the CES service agent
	//  `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.remote_dialogflow_agent
	RemoteDialogflowAgent *Agent_RemoteDialogflowAgent `json:"remoteDialogflowAgent,omitempty"`

	// Required. Display name of the agent.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Human-readable description of the agent.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.description
	Description *string `json:"description,omitempty"`

	// Optional. Configurations for the LLM model.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.model_settings
	ModelSettings *ModelSettings `json:"modelSettings,omitempty"`

	// Optional. Instructions for the LLM model to guide the agent's behavior.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.instruction
	Instruction *string `json:"instruction,omitempty"`

	// Optional. List of available tools for the agent.
	//  Format: `projects/{project}/locations/{location}/apps/{app}/tools/{tool}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.tools
	ToolRefs []CesToolRef `json:"toolRefs,omitempty"`

	// Optional. List of child agents in the agent tree.
	//  Format: `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.child_agents
	ChildAgentRefs []CesAgentRef `json:"childAgentRefs,omitempty"`

	// Optional. The callbacks to execute before the agent is called.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.before_agent_callbacks
	BeforeAgentCallbacks []Callback `json:"beforeAgentCallbacks,omitempty"`

	// Optional. The callbacks to execute after the agent is called.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.after_agent_callbacks
	AfterAgentCallbacks []Callback `json:"afterAgentCallbacks,omitempty"`

	// Optional. The callbacks to execute before the model is called. If there are
	//  multiple calls to the model, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.before_model_callbacks
	BeforeModelCallbacks []Callback `json:"beforeModelCallbacks,omitempty"`

	// Optional. The callbacks to execute after the model is called. If there are
	//  multiple calls to the model, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.after_model_callbacks
	AfterModelCallbacks []Callback `json:"afterModelCallbacks,omitempty"`

	// Optional. The callbacks to execute before the tool is invoked. If there are
	//  multiple tool invocations, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.before_tool_callbacks
	BeforeToolCallbacks []Callback `json:"beforeToolCallbacks,omitempty"`

	// Optional. The callbacks to execute after the tool is invoked. If there are
	//  multiple tool invocations, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.after_tool_callbacks
	AfterToolCallbacks []Callback `json:"afterToolCallbacks,omitempty"`

	// Optional. List of guardrails for the agent.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.guardrails
	GuardrailRefs []CesGuardrailRef `json:"guardrailRefs,omitempty"`

	// Optional. List of toolsets for the agent.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.toolsets
	Toolsets []Agent_AgentToolset `json:"toolsets,omitempty"`

	// Optional. Agent transfer rules.
	//  If multiple rules match, the first one in the list will be used.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.transfer_rules
	TransferRules []TransferRule `json:"transferRules,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.Agent.AgentToolset
type Agent_AgentToolset struct {
	// Required. The resource name of the toolset.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.AgentToolset.toolset
	// +required
	ToolsetRef *CesToolsetRef `json:"toolsetRef,omitempty"`

	// Optional. The tools IDs to filter the toolset.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.AgentToolset.tool_ids
	ToolIDs []string `json:"toolIDs,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.TransferRule
type TransferRule struct {
	// Optional. A rule that immediately transfers to the target agent when the
	//  condition is met.
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.deterministic_transfer
	DeterministicTransfer *TransferRule_DeterministicTransfer `json:"deterministicTransfer,omitempty"`

	// Optional. Rule that prevents the planner from transferring to the target
	//  agent.
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.disable_planner_transfer
	DisablePlannerTransfer *TransferRule_DisablePlannerTransfer `json:"disablePlannerTransfer,omitempty"`

	// Required. The resource name of the child agent the rule applies to.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.child_agent
	// +required
	ChildAgentRef *CesAgentRef `json:"childAgentRef,omitempty"`

	// Required. The direction of the transfer.
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.direction
	// +required
	Direction *string `json:"direction,omitempty"`
}

// CesAgentStatus defines the config connector machine state of CesAgent
type CesAgentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CesAgent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CesAgentObservedState `json:"observedState,omitempty"`
}

// CesAgentObservedState is the state of the CesAgent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1.Agent
type CesAgentObservedState struct {
	// Output only. Timestamp when the agent was created.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the agent was last updated.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Etag used to ensure the object hasn't changed during a read-modify-write
	//  operation. If the etag is empty, the update will overwrite any concurrent
	//  changes.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. If the agent is generated by the LLM assistant, this field
	//  contains a descriptive summary of the generation.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.generated_summary
	GeneratedSummary *string `json:"generatedSummary,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesagent;gcpcesagents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CesAgent is the Schema for the CesAgent API
// +k8s:openapi-gen=true
type CesAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CesAgentSpec   `json:"spec,omitempty"`
	Status CesAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CesAgentList contains a list of CesAgent
type CesAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CesAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CesAgent{}, &CesAgentList{})
}
