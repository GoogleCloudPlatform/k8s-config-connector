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


// +kcc:proto=google.cloud.dialogflow.v2.EntityType.Entity
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
	// +kcc:proto:field=google.cloud.dialogflow.v2.EntityType.Entity.value
	Value *string `json:"value,omitempty"`

	// Required. A collection of value synonyms. For example, if the entity type
	//  is *vegetable*, and `value` is *scallions*, a synonym could be *green
	//  onions*.
	//
	//  For `KIND_LIST` entity types:
	//
	//  *   This collection must contain exactly one synonym equal to `value`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EntityType.Entity.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SessionEntityType
type SessionEntityType struct {
	// Required. The unique identifier of this session entity type. Format:
	//  `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	//  Display Name>`, or `projects/<Project ID>/agent/environments/<Environment
	//  ID>/users/<User ID>/sessions/<Session ID>/entityTypes/<Entity Type Display
	//  Name>`.
	//  If `Environment ID` is not specified, we assume default 'draft'
	//  environment. If `User ID` is not specified, we assume default '-' user.
	//
	//  `<Entity Type Display Name>` must be the display name of an existing entity
	//  type in the same agent that will be overridden or supplemented.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SessionEntityType.name
	Name *string `json:"name,omitempty"`

	// Required. Indicates whether the additional data should override or
	//  supplement the custom entity type definition.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SessionEntityType.entity_override_mode
	EntityOverrideMode *string `json:"entityOverrideMode,omitempty"`

	// Required. The collection of entities associated with this session entity
	//  type.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SessionEntityType.entities
	Entities []EntityType_Entity `json:"entities,omitempty"`
}
