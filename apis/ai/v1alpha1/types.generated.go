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


// +kcc:proto=google.ai.generativelanguage.v1beta.Chunk
type Chunk struct {
	// Immutable. Identifier. The `Chunk` resource name. The ID (name excluding
	//  the "corpora/*/documents/*/chunks/" prefix) can contain up to 40 characters
	//  that are lowercase alphanumeric or dashes (-). The ID cannot start or end
	//  with a dash. If the name is empty on create, a random 12-character unique
	//  ID will be generated.
	//  Example: `corpora/{corpus_id}/documents/{document_id}/chunks/123a456b789c`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Chunk.name
	Name *string `json:"name,omitempty"`

	// Required. The content for the `Chunk`, such as the text string.
	//  The maximum number of tokens per chunk is 2043.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Chunk.data
	Data *ChunkData `json:"data,omitempty"`

	// Optional. User provided custom metadata stored as key-value pairs.
	//  The maximum number of `CustomMetadata` per chunk is 20.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Chunk.custom_metadata
	CustomMetadata []CustomMetadata `json:"customMetadata,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.ChunkData
type ChunkData struct {
	// The `Chunk` content as a string.
	//  The maximum number of tokens per chunk is 2043.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.ChunkData.string_value
	StringValue *string `json:"stringValue,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.CustomMetadata
type CustomMetadata struct {
	// The string value of the metadata to store.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CustomMetadata.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// The StringList value of the metadata to store.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CustomMetadata.string_list_value
	StringListValue *StringList `json:"stringListValue,omitempty"`

	// The numeric value of the metadata to store.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CustomMetadata.numeric_value
	NumericValue *float32 `json:"numericValue,omitempty"`

	// Required. The key of the metadata to store.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.CustomMetadata.key
	Key *string `json:"key,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.StringList
type StringList struct {
	// The string values of the metadata to store.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.StringList.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Chunk
type ChunkObservedState struct {
	// Output only. The Timestamp of when the `Chunk` was created.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Chunk.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The Timestamp of when the `Chunk` was last updated.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Chunk.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the `Chunk`.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Chunk.state
	State *string `json:"state,omitempty"`
}
