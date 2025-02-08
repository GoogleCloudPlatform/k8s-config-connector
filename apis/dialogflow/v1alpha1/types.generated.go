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


// +kcc:proto=google.cloud.dialogflow.cx.v3.Generator
type Generator struct {
	// The unique identifier of the generator.
	//  Must be set for the
	//  [Generators.UpdateGenerator][google.cloud.dialogflow.cx.v3.Generators.UpdateGenerator]
	//  method. [Generators.CreateGenerate][] populates the name automatically.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/generators/<GeneratorID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the generator, unique within the
	//  agent. The prompt contains pre-defined parameters such as $conversation,
	//  $last-user-utterance, etc. populated by Dialogflow. It can also contain
	//  custom placeholders which will be resolved during fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Prompt for the LLM model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.prompt_text
	PromptText *Phrase `json:"promptText,omitempty"`

	// Optional. List of custom placeholders in the prompt text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.placeholders
	Placeholders []Generator_Placeholder `json:"placeholders,omitempty"`

	// Parameters passed to the LLM to configure its behavior.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.model_parameter
	ModelParameter *Generator_ModelParameter `json:"modelParameter,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Generator.ModelParameter
type Generator_ModelParameter struct {
	// The temperature used for sampling. Temperature sampling occurs after both
	//  topP and topK have been applied.
	//  Valid range: [0.0, 1.0]
	//  Low temperature = less random. High temperature = more random.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.ModelParameter.temperature
	Temperature *float32 `json:"temperature,omitempty"`

	// The maximum number of tokens to generate.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.ModelParameter.max_decode_steps
	MaxDecodeSteps *int32 `json:"maxDecodeSteps,omitempty"`

	// If set, only the tokens comprising the top top_p probability mass are
	//  considered. If both top_p and top_k are
	//  set, top_p will be used for further refining candidates selected with
	//  top_k.
	//  Valid range: (0.0, 1.0].
	//  Small topP = less random. Large topP = more random.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.ModelParameter.top_p
	TopP *float32 `json:"topP,omitempty"`

	// If set, the sampling process in each step is limited to the top_k tokens
	//  with highest probabilities.
	//  Valid range: [1, 40] or 1000+.
	//  Small topK = less random. Large topK = more random.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.ModelParameter.top_k
	TopK *int32 `json:"topK,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Generator.Placeholder
type Generator_Placeholder struct {
	// Unique ID used to map custom placeholder to parameters in fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.Placeholder.id
	ID *string `json:"id,omitempty"`

	// Custom placeholder value in the prompt text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.Placeholder.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Phrase
type Phrase struct {
	// Required. Text input which can be used for prompt or banned phrases.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Phrase.text
	Text *string `json:"text,omitempty"`
}
