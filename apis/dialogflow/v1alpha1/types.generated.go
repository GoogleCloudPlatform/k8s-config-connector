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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent
type Intent struct {
	// The unique identifier of the intent.
	//  Required for the
	//  [Intents.UpdateIntent][google.cloud.dialogflow.cx.v3beta1.Intents.UpdateIntent]
	//  method.
	//  [Intents.CreateIntent][google.cloud.dialogflow.cx.v3beta1.Intents.CreateIntent]
	//  populates the name automatically.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/intents/<IntentID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the intent, unique within the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The collection of training phrases the agent is trained on to identify the
	//  intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.training_phrases
	TrainingPhrases []Intent_TrainingPhrase `json:"trainingPhrases,omitempty"`

	// The collection of parameters associated with the intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.parameters
	Parameters []Intent_Parameter `json:"parameters,omitempty"`

	// The priority of this intent. Higher numbers represent higher
	//  priorities.
	//
	//  - If the supplied value is unspecified or 0, the service
	//    translates the value to 500,000, which corresponds to the
	//    `Normal` priority in the console.
	//  - If the supplied value is negative, the intent is ignored
	//    in runtime detect intent requests.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.priority
	Priority *int32 `json:"priority,omitempty"`

	// Indicates whether this is a fallback intent. Currently only default
	//  fallback intent is allowed in the agent, which is added upon agent
	//  creation.
	//  Adding training phrases to fallback intent is useful in the case of
	//  requests that are mistakenly matched, since training phrases assigned to
	//  fallback intents act as negative examples that triggers no-match event.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.is_fallback
	IsFallback *bool `json:"isFallback,omitempty"`

	// The key/value metadata to label an intent. Labels can contain
	//  lowercase letters, digits and the symbols '-' and '_'. International
	//  characters are allowed, including letters from unicase alphabets. Keys must
	//  start with a letter. Keys and values can be no longer than 63 characters
	//  and no more than 128 bytes.
	//
	//  Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed
	//  Dialogflow defined labels include:
	//  * sys-head
	//  * sys-contextual
	//  The above labels do not require value. "sys-head" means the intent is a
	//  head intent. "sys-contextual" means the intent is a contextual intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Human readable description for better understanding an intent like its
	//  scope, content, result etc. Maximum character limit: 140 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter
type Intent_Parameter struct {
	// Required. The unique identifier of the parameter. This field
	//  is used by [training
	//  phrases][google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase] to
	//  annotate their
	//  [parts][google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.id
	ID *string `json:"id,omitempty"`

	// Required. The entity type of the parameter.
	//  Format:
	//  `projects/-/locations/-/agents/-/entityTypes/<SystemEntityTypeID>` for
	//  system entity types (for example,
	//  `projects/-/locations/-/agents/-/entityTypes/sys.date`), or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/entityTypes/<EntityTypeID>`
	//  for developer entity types.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.is_list
	IsList *bool `json:"isList,omitempty"`

	// Indicates whether the parameter content should be redacted in log. If
	//  redaction is enabled, the parameter content will be replaced by parameter
	//  name during logging.
	//  Note: the parameter content is subject to redaction if either parameter
	//  level redaction or [entity type level
	//  redaction][google.cloud.dialogflow.cx.v3beta1.EntityType.redact] is
	//  enabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.Parameter.redact
	Redact *bool `json:"redact,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase
type Intent_TrainingPhrase struct {
	// Output only. The unique identifier of the training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.id
	ID *string `json:"id,omitempty"`

	// Required. The ordered list of training phrase parts.
	//  The parts are concatenated in order to form the training phrase.
	//
	//  Note: The API does not automatically annotate training phrases like the
	//  Dialogflow Console does.
	//
	//  Note: Do not forget to include whitespace at part boundaries, so the
	//  training phrase is well formatted when the parts are concatenated.
	//
	//  If the training phrase does not need to be annotated with parameters,
	//  you just need a single part with only the
	//  [Part.text][google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part.text]
	//  field set.
	//
	//  If you want to annotate the training phrase, you must create multiple
	//  parts, where the fields of each part are populated in one of two ways:
	//
	//  -   `Part.text` is set to a part of the phrase that has no parameters.
	//  -   `Part.text` is set to a part of the phrase that you want to annotate,
	//      and the `parameter_id` field is set.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.parts
	Parts []Intent_TrainingPhrase_Part `json:"parts,omitempty"`

	// Indicates how many times this example was added to the intent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.repeat_count
	RepeatCount *int32 `json:"repeatCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part
type Intent_TrainingPhrase_Part struct {
	// Required. The text for this part.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part.text
	Text *string `json:"text,omitempty"`

	// The [parameter][google.cloud.dialogflow.cx.v3beta1.Intent.Parameter]
	//  used to annotate this part of the training phrase. This field is
	//  required for annotated parts of the training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Intent.TrainingPhrase.Part.parameter_id
	ParameterID *string `json:"parameterID,omitempty"`
}
