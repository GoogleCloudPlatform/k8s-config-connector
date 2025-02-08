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


// +kcc:proto=google.cloud.dialogflow.v2beta1.ConversationContext
type ConversationContext struct {
	// Optional. List of message transcripts in the conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationContext.message_entries
	MessageEntries []MessageEntry `json:"messageEntries,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.FewShotExample
type FewShotExample struct {
	// Optional. Conversation transcripts.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.FewShotExample.conversation_context
	ConversationContext *ConversationContext `json:"conversationContext,omitempty"`

	// Optional. Key is the placeholder field name in input, value is the value of
	//  the placeholder. E.g. instruction contains "@price", and ingested data has
	//  <"price", "10">
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.FewShotExample.extra_info
	ExtraInfo map[string]string `json:"extraInfo,omitempty"`

	// Summarization sections.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.FewShotExample.summarization_section_list
	SummarizationSectionList *SummarizationSectionList `json:"summarizationSectionList,omitempty"`

	// Required. Example output of the model.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.FewShotExample.output
	Output *GeneratorSuggestion `json:"output,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Generator
type Generator struct {

	// Optional. Human readable description of the generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Generator.description
	Description *string `json:"description,omitempty"`

	// Input of Summarization feature.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Generator.summarization_context
	SummarizationContext *SummarizationContext `json:"summarizationContext,omitempty"`

	// Optional. Inference parameters for this generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Generator.inference_parameter
	InferenceParameter *InferenceParameter `json:"inferenceParameter,omitempty"`

	// Optional. The trigger event of the generator. It defines when the generator
	//  is triggered in a conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Generator.trigger_event
	TriggerEvent *string `json:"triggerEvent,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.GeneratorSuggestion
type GeneratorSuggestion struct {
	// Optional. Suggested summary.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.GeneratorSuggestion.summary_suggestion
	SummarySuggestion *SummarySuggestion `json:"summarySuggestion,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.InferenceParameter
type InferenceParameter struct {
	// Optional. Maximum number of the output tokens for the generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.InferenceParameter.max_output_tokens
	MaxOutputTokens *int32 `json:"maxOutputTokens,omitempty"`

	// Optional. Controls the randomness of LLM predictions.
	//  Low temperature = less random. High temperature = more random.
	//  If unset (or 0), uses a default value of 0.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.InferenceParameter.temperature
	Temperature *float64 `json:"temperature,omitempty"`

	// Optional. Top-k changes how the model selects tokens for output. A top-k of
	//  1 means the selected token is the most probable among all tokens in the
	//  model's vocabulary (also called greedy decoding), while a top-k of 3 means
	//  that the next token is selected from among the 3 most probable tokens
	//  (using temperature). For each token selection step, the top K tokens with
	//  the highest probabilities are sampled. Then tokens are further filtered
	//  based on topP with the final token selected using temperature sampling.
	//  Specify a lower value for less random responses and a higher value for more
	//  random responses. Acceptable value is [1, 40], default to 40.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.InferenceParameter.top_k
	TopK *int32 `json:"topK,omitempty"`

	// Optional. Top-p changes how the model selects tokens for output. Tokens are
	//  selected from most K (see topK parameter) probable to least until the sum
	//  of their probabilities equals the top-p value. For example, if tokens A, B,
	//  and C have a probability of 0.3, 0.2, and 0.1 and the top-p value is 0.5,
	//  then the model will select either A or B as the next token (using
	//  temperature) and doesn't consider C. The default top-p value is 0.95.
	//  Specify a lower value for less random responses and a higher value for more
	//  random responses. Acceptable value is [0.0, 1.0], default to 0.95.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.InferenceParameter.top_p
	TopP *float64 `json:"topP,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.MessageEntry
type MessageEntry struct {
	// Optional. Participant role of the message.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.MessageEntry.role
	Role *string `json:"role,omitempty"`

	// Optional. Transcript content of the message.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.MessageEntry.text
	Text *string `json:"text,omitempty"`

	// Optional. The language of the text. See [Language
	//  Support](https://cloud.google.com/dialogflow/docs/reference/language) for a
	//  list of the currently supported language codes.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.MessageEntry.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Optional. Create time of the message entry.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.MessageEntry.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SummarizationContext
type SummarizationContext struct {
	// Optional. List of sections. Note it contains both predefined section sand
	//  customer defined sections.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationContext.summarization_sections
	SummarizationSections []SummarizationSection `json:"summarizationSections,omitempty"`

	// Optional. List of few shot examples.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationContext.few_shot_examples
	FewShotExamples []FewShotExample `json:"fewShotExamples,omitempty"`

	// Optional. Version of the feature. If not set, default to latest version.
	//  Current candidates are ["1.0"].
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationContext.version
	Version *string `json:"version,omitempty"`

	// Optional. The target language of the generated summary. The language code
	//  for conversation will be used if this field is empty. Supported 2.0 and
	//  later versions.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationContext.output_language_code
	OutputLanguageCode *string `json:"outputLanguageCode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SummarizationSection
type SummarizationSection struct {
	// Optional. Name of the section, for example, "situation".
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationSection.key
	Key *string `json:"key,omitempty"`

	// Optional. Definition of the section, for example, "what the customer needs
	//  help with or has question about."
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationSection.definition
	Definition *string `json:"definition,omitempty"`

	// Optional. Type of the summarization section.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationSection.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SummarizationSectionList
type SummarizationSectionList struct {
	// Optional. Summarization sections.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarizationSectionList.summarization_sections
	SummarizationSections []SummarizationSection `json:"summarizationSections,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SummarySuggestion
type SummarySuggestion struct {
	// Required. All the parts of generated summary.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarySuggestion.summary_sections
	SummarySections []SummarySuggestion_SummarySection `json:"summarySections,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SummarySuggestion.SummarySection
type SummarySuggestion_SummarySection struct {
	// Required. Name of the section.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarySuggestion.SummarySection.section
	Section *string `json:"section,omitempty"`

	// Required. Summary text for the section.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SummarySuggestion.SummarySection.summary
	Summary *string `json:"summary,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Generator
type GeneratorObservedState struct {
	// Output only. Identifier. The resource name of the generator. Format:
	//  `projects/<Project ID>/locations/<Location ID>/generators/<Generator ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Generator.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Generator.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time of this generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Generator.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
