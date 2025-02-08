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


// +kcc:proto=google.cloud.dialogflow.cx.v3.GenerativeSettings
type GenerativeSettings struct {
	// Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/generativeSettings`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.name
	Name *string `json:"name,omitempty"`

	// Settings for Generative Fallback.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.fallback_settings
	FallbackSettings *GenerativeSettings_FallbackSettings `json:"fallbackSettings,omitempty"`

	// Settings for Generative Safety.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.generative_safety_settings
	GenerativeSafetySettings *SafetySettings `json:"generativeSafetySettings,omitempty"`

	// Settings for knowledge connector.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.knowledge_connector_settings
	KnowledgeConnectorSettings *GenerativeSettings_KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`

	// Language for this settings.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.GenerativeSettings.FallbackSettings
type GenerativeSettings_FallbackSettings struct {
	// Display name of the selected prompt.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.FallbackSettings.selected_prompt
	SelectedPrompt *string `json:"selectedPrompt,omitempty"`

	// Stored prompts that can be selected, for example default templates like
	//  "conservative" or "chatty", or user defined ones.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.FallbackSettings.prompt_templates
	PromptTemplates []GenerativeSettings_FallbackSettings_PromptTemplate `json:"promptTemplates,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.GenerativeSettings.FallbackSettings.PromptTemplate
type GenerativeSettings_FallbackSettings_PromptTemplate struct {
	// Prompt name.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.FallbackSettings.PromptTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Prompt text that is sent to a LLM on no-match default, placeholders are
	//  filled downstream. For example: "Here is a conversation $conversation,
	//  a response is: "
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.FallbackSettings.PromptTemplate.prompt_text
	PromptText *string `json:"promptText,omitempty"`

	// If the flag is true, the prompt is frozen and cannot be modified by
	//  users.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.FallbackSettings.PromptTemplate.frozen
	Frozen *bool `json:"frozen,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.GenerativeSettings.KnowledgeConnectorSettings
type GenerativeSettings_KnowledgeConnectorSettings struct {
	// Name of the company, organization or other entity that the agent
	//  represents. Used for knowledge connector LLM prompt and for knowledge
	//  search.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.KnowledgeConnectorSettings.business
	Business *string `json:"business,omitempty"`

	// Name of the virtual agent. Used for LLM prompt. Can be left empty.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.KnowledgeConnectorSettings.agent
	Agent *string `json:"agent,omitempty"`

	// Identity of the agent, e.g. "virtual agent", "AI assistant".
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.KnowledgeConnectorSettings.agent_identity
	AgentIdentity *string `json:"agentIdentity,omitempty"`

	// Company description, used for LLM prompt, e.g. "a family company selling
	//  freshly roasted coffee beans".
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.KnowledgeConnectorSettings.business_description
	BusinessDescription *string `json:"businessDescription,omitempty"`

	// Agent scope, e.g. "Example company website", "internal Example
	//  company website for employees", "manual of car owner".
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.KnowledgeConnectorSettings.agent_scope
	AgentScope *string `json:"agentScope,omitempty"`

	// Whether to disable fallback to Data Store search results (in case the LLM
	//  couldn't pick a proper answer). Per default the feature is enabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GenerativeSettings.KnowledgeConnectorSettings.disable_data_store_fallback
	DisableDataStoreFallback *bool `json:"disableDataStoreFallback,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.SafetySettings
type SafetySettings struct {
	// Banned phrases for generated text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SafetySettings.banned_phrases
	BannedPhrases []SafetySettings_Phrase `json:"bannedPhrases,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.SafetySettings.Phrase
type SafetySettings_Phrase struct {
	// Required. Text input which can be used for prompt or banned phrases.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SafetySettings.Phrase.text
	Text *string `json:"text,omitempty"`

	// Required. Language code of the phrase.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SafetySettings.Phrase.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}
