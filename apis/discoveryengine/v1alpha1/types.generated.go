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


// +kcc:proto=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig
type Engine_ChatEngineConfig struct {
	// The configurationt generate the Dialogflow agent that is associated to
	//  this Engine.
	//
	//  Note that these configurations are one-time consumed by
	//  and passed to Dialogflow service. It means they cannot be retrieved using
	//  [EngineService.GetEngine][google.cloud.discoveryengine.v1.EngineService.GetEngine]
	//  or
	//  [EngineService.ListEngines][google.cloud.discoveryengine.v1.EngineService.ListEngines]
	//  API after engine creation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.agent_creation_config
	AgentCreationConfig *Engine_ChatEngineConfig_AgentCreationConfig `json:"agentCreationConfig,omitempty"`

	// The resource name of an exist Dialogflow agent to link to this Chat
	//  Engine. Customers can either provide `agent_creation_config` to create
	//  agent or provide an agent name that links the agent with the Chat engine.
	//
	//  Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	//  ID>`.
	//
	//  Note that the `dialogflow_agent_to_link` are one-time consumed by and
	//  passed to Dialogflow service. It means they cannot be retrieved using
	//  [EngineService.GetEngine][google.cloud.discoveryengine.v1.EngineService.GetEngine]
	//  or
	//  [EngineService.ListEngines][google.cloud.discoveryengine.v1.EngineService.ListEngines]
	//  API after engine creation. Use
	//  [ChatEngineMetadata.dialogflow_agent][google.cloud.discoveryengine.v1.Engine.ChatEngineMetadata.dialogflow_agent]
	//  for actual agent association after Engine is created.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.dialogflow_agent_to_link
	DialogflowAgentToLink *string `json:"dialogflowAgentToLink,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig
type Engine_ChatEngineConfig_AgentCreationConfig struct {
	// Name of the company, organization or other entity that the agent
	//  represents. Used for knowledge connector LLM prompt and for knowledge
	//  search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.business
	Business *string `json:"business,omitempty"`

	// Required. The default language of the agent as a language tag.
	//  See [Language
	//  Support](https://cloud.google.com/dialogflow/docs/reference/language)
	//  for a list of the currently supported language codes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.default_language_code
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty"`

	// Required. The time zone of the agent from the [time zone
	//  database](https://www.iana.org/time-zones), e.g., America/New_York,
	//  Europe/Paris.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Agent location for Agent creation, supported values: global/us/eu.
	//  If not provided, us Engine will create Agent using us-central-1 by
	//  default; eu Engine will create Agent using eu-west-1 by default.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.ChatEngineMetadata
type Engine_ChatEngineMetadata struct {
	// The resource name of a Dialogflow agent, that this Chat Engine refers
	//  to.
	//
	//  Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	//  ID>`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineMetadata.dialogflow_agent
	DialogflowAgent *string `json:"dialogflowAgent,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.CommonConfig
type Engine_CommonConfig struct {
	// The name of the company, business or entity that is associated with the
	//  engine. Setting this may help improve LLM related features.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.CommonConfig.company_name
	CompanyName *string `json:"companyName,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.SearchEngineConfig
type Engine_SearchEngineConfig struct {
	// The search feature tier of this engine.
	//
	//  Different tiers might have different
	//  pricing. To learn more, check the pricing documentation.
	//
	//  Defaults to
	//  [SearchTier.SEARCH_TIER_STANDARD][google.cloud.discoveryengine.v1.SearchTier.SEARCH_TIER_STANDARD]
	//  if not specified.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.SearchEngineConfig.search_tier
	SearchTier *string `json:"searchTier,omitempty"`

	// The add-on that this search engine enables.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.SearchEngineConfig.search_add_ons
	SearchAddOns []string `json:"searchAddOns,omitempty"`
}
