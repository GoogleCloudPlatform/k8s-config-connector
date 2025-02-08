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


// +kcc:proto=google.cloud.connectors.v1.ConnectionSchemaMetadata
type ConnectionSchemaMetadata struct {
}

// +kcc:proto=google.cloud.connectors.v1.ConnectionSchemaMetadata
type ConnectionSchemaMetadataObservedState struct {
	// Output only. List of entity names.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionSchemaMetadata.entities
	Entities []string `json:"entities,omitempty"`

	// Output only. List of actions.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionSchemaMetadata.actions
	Actions []string `json:"actions,omitempty"`

	// Output only. Resource name.
	//  Format:
	//  projects/{project}/locations/{location}/connections/{connection}/connectionSchemaMetadata
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionSchemaMetadata.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when the connection runtime schema was updated.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionSchemaMetadata.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when the connection runtime schema refresh was
	//  triggered.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionSchemaMetadata.refresh_time
	RefreshTime *string `json:"refreshTime,omitempty"`

	// Output only. The current state of runtime schema.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectionSchemaMetadata.state
	State *string `json:"state,omitempty"`
}
