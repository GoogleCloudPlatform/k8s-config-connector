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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.EntityType
type EntityType struct {
	// The unique identifier of the entity type.
	//  Required for
	//  [EntityTypes.UpdateEntityType][google.cloud.dialogflow.cx.v3beta1.EntityTypes.UpdateEntityType].
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/entityTypes/<EntityTypeID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the entity type, unique within the
	//  agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Indicates the kind of entity type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.kind
	Kind *string `json:"kind,omitempty"`

	// Indicates whether the entity type can be automatically expanded.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.auto_expansion_mode
	AutoExpansionMode *string `json:"autoExpansionMode,omitempty"`

	// The collection of entity entries associated with the entity type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.entities
	Entities []EntityType_Entity `json:"entities,omitempty"`

	// Collection of exceptional words and phrases that shouldn't be matched.
	//  For example, if you have a size entity type with entry `giant`(an
	//  adjective), you might consider adding `giants`(a noun) as an exclusion.
	//  If the kind of entity type is `KIND_MAP`, then the phrases specified by
	//  entities and excluded phrases should be mutually exclusive.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.excluded_phrases
	ExcludedPhrases []EntityType_ExcludedPhrase `json:"excludedPhrases,omitempty"`

	// Enables fuzzy entity extraction during classification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.enable_fuzzy_extraction
	EnableFuzzyExtraction *bool `json:"enableFuzzyExtraction,omitempty"`

	// Indicates whether parameters of the entity type should be redacted in log.
	//  If redaction is enabled, page parameters and intent parameters referring to
	//  the entity type will be replaced by parameter name during logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.redact
	Redact *bool `json:"redact,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.EntityType.Entity
type EntityType_Entity struct {
	// Required. The primary value associated with this entity entry.
	//  For example, if the entity type is *vegetable*, the value could be
	//  *scallions*.
	//
	//  For `KIND_MAP` entity types:
	//
	//  *   A canonical value to be used in place of synonyms.
	//
	//  For `KIND_LIST` entity types:
	//
	//  *   A string that can contain references to other entity types (with or
	//      without aliases).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.Entity.value
	Value *string `json:"value,omitempty"`

	// Required. A collection of value synonyms. For example, if the entity type
	//  is *vegetable*, and `value` is *scallions*, a synonym could be *green
	//  onions*.
	//
	//  For `KIND_LIST` entity types:
	//
	//  *   This collection must contain exactly one synonym equal to `value`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.Entity.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.EntityType.ExcludedPhrase
type EntityType_ExcludedPhrase struct {
	// Required. The word or phrase to be excluded.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.EntityType.ExcludedPhrase.value
	Value *string `json:"value,omitempty"`
}
