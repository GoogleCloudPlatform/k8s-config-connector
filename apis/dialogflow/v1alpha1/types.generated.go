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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.LlmModelSettings
type LlmModelSettings struct {
	// The selected LLM model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.LlmModelSettings.model
	Model *string `json:"model,omitempty"`

	// The custom prompt to use.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.LlmModelSettings.prompt_text
	PromptText *string `json:"promptText,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.ParameterDefinition
type ParameterDefinition struct {
	// Name of parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ParameterDefinition.name
	Name *string `json:"name,omitempty"`

	// Type of parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ParameterDefinition.type
	Type *string `json:"type,omitempty"`

	// Human-readable description of the parameter. Limited to 300 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.ParameterDefinition.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Playbook
type Playbook struct {
	// The unique identifier of the playbook.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/playbooks/<PlaybookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the playbook, unique within an agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. High level description of the goal the playbook intend to
	//  accomplish.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.goal
	Goal *string `json:"goal,omitempty"`

	// Optional. Defined structured input parameters for this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.input_parameter_definitions
	InputParameterDefinitions []ParameterDefinition `json:"inputParameterDefinitions,omitempty"`

	// Optional. Defined structured output parameters for this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.output_parameter_definitions
	OutputParameterDefinitions []ParameterDefinition `json:"outputParameterDefinitions,omitempty"`

	// Instruction to accomplish target goal.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.instruction
	Instruction *Playbook_Instruction `json:"instruction,omitempty"`

	// Optional. The resource name of tools referenced by the current playbook in
	//  the instructions. If not provided explicitly, they are will
	//  be implied using the tool being referenced in goal and steps.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.referenced_tools
	ReferencedTools []string `json:"referencedTools,omitempty"`

	// Optional. Llm model settings for the playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.llm_model_settings
	LlmModelSettings *LlmModelSettings `json:"llmModelSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Playbook.Instruction
type Playbook_Instruction struct {
	// Ordered list of step by step execution instructions to accomplish
	//  target goal.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.Instruction.steps
	Steps []Playbook_Step `json:"steps,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Playbook.Step
type Playbook_Step struct {
	// Step instruction in text format.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.Step.text
	Text *string `json:"text,omitempty"`

	// Sub-processing needed to execute the current step.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.Step.steps
	Steps []Playbook_Step `json:"steps,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Playbook
type PlaybookObservedState struct {
	// Output only. Estimated number of tokes current playbook takes when sent to
	//  the LLM.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.token_count
	TokenCount *int64 `json:"tokenCount,omitempty"`

	// Output only. The timestamp of initial playbook creation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last time the playbook version was updated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The resource name of other playbooks referenced by the current
	//  playbook in the instructions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.referenced_playbooks
	ReferencedPlaybooks []string `json:"referencedPlaybooks,omitempty"`

	// Output only. The resource name of flows referenced by the current playbook
	//  in the instructions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Playbook.referenced_flows
	ReferencedFlows []string `json:"referencedFlows,omitempty"`
}
