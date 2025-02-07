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

// +kcc:proto=google.ai.generativelanguage.v1beta.Document
type Document struct {
	// Immutable. Identifier. The `Document` resource name. The ID (name excluding
	//  the "corpora/*/documents/" prefix) can contain up to 40 characters that are
	//  lowercase alphanumeric or dashes (-). The ID cannot start or end with a
	//  dash. If the name is empty on create, a unique name will be derived from
	//  `display_name` along with a 12 character random suffix.
	//  Example: `corpora/{corpus_id}/documents/my-awesome-doc-123a456b789c`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Document.name
	Name *string `json:"name,omitempty"`

	// Optional. The human-readable display name for the `Document`. The display
	//  name must be no more than 512 characters in length, including spaces.
	//  Example: "Semantic Retriever Documentation"
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Document.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User provided custom metadata stored as key-value pairs used for
	//  querying. A `Document` can have a maximum of 20 `CustomMetadata`.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Document.custom_metadata
	CustomMetadata []CustomMetadata `json:"customMetadata,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.StringList
type StringList struct {
	// The string values of the metadata to store.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.StringList.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Document
type DocumentObservedState struct {
	// Output only. The Timestamp of when the `Document` was last updated.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Document.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Timestamp of when the `Document` was created.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Document.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
