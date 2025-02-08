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


// +kcc:proto=google.cloud.dialogflow.v2.Agent
type Agent struct {
	// Required. The project of this agent.
	//  Format: `projects/<Project ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.parent
	Parent *string `json:"parent,omitempty"`

	// Required. The name of this agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The default language of the agent as a language tag. See
	//  [Language
	//  Support](https://cloud.google.com/dialogflow/docs/reference/language)
	//  for a list of the currently supported language codes. This field cannot be
	//  set by the `Update` method.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.default_language_code
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty"`

	// Optional. The list of all languages supported by this agent (except for the
	//  `default_language_code`).
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.supported_language_codes
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	// Required. The time zone of this agent from the
	//  [time zone database](https://www.iana.org/time-zones), e.g.,
	//  America/New_York, Europe/Paris.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. The description of this agent.
	//  The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.description
	Description *string `json:"description,omitempty"`

	// Optional. The URI of the agent's avatar.
	//  Avatars are used throughout the Dialogflow console and in the self-hosted
	//  [Web
	//  Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo)
	//  integration.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.avatar_uri
	AvatarURI *string `json:"avatarURI,omitempty"`

	// Optional. Determines whether this agent should log conversation queries.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.enable_logging
	EnableLogging *bool `json:"enableLogging,omitempty"`

	// Optional. Determines how intents are detected from user queries.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.match_mode
	MatchMode *string `json:"matchMode,omitempty"`

	// Optional. To filter out false positive results and still get variety in
	//  matched natural language inputs for your agent, you can tune the machine
	//  learning classification threshold. If the returned score value is less than
	//  the threshold value, then a fallback intent will be triggered or, if there
	//  are no fallback intents defined, no intent will be triggered. The score
	//  values range from 0.0 (completely uncertain) to 1.0 (completely certain).
	//  If set to 0.0, the default of 0.3 is used.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.classification_threshold
	ClassificationThreshold *float32 `json:"classificationThreshold,omitempty"`

	// Optional. API version displayed in Dialogflow console. If not specified,
	//  V2 API is assumed. Clients are free to query different service endpoints
	//  for different API versions. However, bots connectors and webhook calls will
	//  follow the specified API version.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.api_version
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Optional. The agent tier. If not specified, TIER_STANDARD is assumed.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.tier
	Tier *string `json:"tier,omitempty"`
}
