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


// +kcc:proto=google.cloud.dialogflow.v2beta1.EntityType
type EntityType struct {
	// The unique identifier of the entity type.
	//  Required for
	//  [EntityTypes.UpdateEntityType][google.cloud.dialogflow.v2beta1.EntityTypes.UpdateEntityType]
	//  and
	//  [EntityTypes.BatchUpdateEntityTypes][google.cloud.dialogflow.v2beta1.EntityTypes.BatchUpdateEntityTypes]
	//  methods. Supported formats:
	//  - `projects/<Project ID>/agent/entityTypes/<Entity Type ID>`
	//  - `projects/<Project ID>/locations/<Location ID>/agent/entityTypes/<Entity
	//    Type ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the entity type.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Indicates the kind of entity type.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.kind
	Kind *string `json:"kind,omitempty"`

	// Optional. Indicates whether the entity type can be automatically
	//  expanded.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.auto_expansion_mode
	AutoExpansionMode *string `json:"autoExpansionMode,omitempty"`

	// Optional. The collection of entity entries associated with the entity type.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.entities
	Entities []EntityType_Entity `json:"entities,omitempty"`

	// Optional. Enables fuzzy entity extraction during classification.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.enable_fuzzy_extraction
	EnableFuzzyExtraction *bool `json:"enableFuzzyExtraction,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.EntityType.Entity
type EntityType_Entity struct {
	// Required. The primary value associated with this entity entry.
	//  For example, if the entity type is *vegetable*, the value could be
	//  *scallions*.
	//
	//  For `KIND_MAP` entity types:
	//
	//  *   A reference value to be used in place of synonyms.
	//
	//  For `KIND_LIST` entity types:
	//
	//  *   A string that can contain references to other entity types (with or
	//      without aliases).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.Entity.value
	Value *string `json:"value,omitempty"`

	// Required. A collection of value synonyms. For example, if the entity type
	//  is *vegetable*, and `value` is *scallions*, a synonym could be *green
	//  onions*.
	//
	//  For `KIND_LIST` entity types:
	//
	//  *   This collection must contain exactly one synonym equal to `value`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EntityType.Entity.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}
