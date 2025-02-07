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


// +kcc:proto=google.ai.generativelanguage.v1beta.Corpus
type Corpus struct {
	// Immutable. Identifier. The `Corpus` resource name. The ID (name excluding
	//  the "corpora/" prefix) can contain up to 40 characters that are lowercase
	//  alphanumeric or dashes
	//  (-). The ID cannot start or end with a dash. If the name is empty on
	//  create, a unique name will be derived from `display_name` along with a 12
	//  character random suffix.
	//  Example: `corpora/my-awesome-corpora-123a456b789c`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Corpus.name
	Name *string `json:"name,omitempty"`

	// Optional. The human-readable display name for the `Corpus`. The display
	//  name must be no more than 512 characters in length, including spaces.
	//  Example: "Docs on Semantic Retriever"
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Corpus.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.Corpus
type CorpusObservedState struct {
	// Output only. The Timestamp of when the `Corpus` was created.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Corpus.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The Timestamp of when the `Corpus` was last updated.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.Corpus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
