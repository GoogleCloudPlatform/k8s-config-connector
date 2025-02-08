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


// +kcc:proto=google.cloud.visionai.v1.Corpus
type Corpus struct {
	// Resource name of the corpus.
	//  Format:
	//  `projects/{project_number}/locations/{location_id}/corpora/{corpus_id}`
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.name
	Name *string `json:"name,omitempty"`

	// Required. The corpus name to shown in the UI. The name can be up to 32
	//  characters long.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the corpus. Can be up to 25000 characters long.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.description
	Description *string `json:"description,omitempty"`

	// Optional. The default TTL value for all assets under the corpus without a
	//  asset level user-defined TTL. For STREAM_VIDEO type corpora, this is
	//  required and the maximum allowed
	//    default_ttl is 10 years.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.default_ttl
	DefaultTtl *string `json:"defaultTtl,omitempty"`

	// Optional. Type of the asset inside corpus.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.type
	Type *string `json:"type,omitempty"`

	// Default search capability setting on corpus level.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.search_capability_setting
	SearchCapabilitySetting *SearchCapabilitySetting `json:"searchCapabilitySetting,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.SearchCapability
type SearchCapability struct {
	// The search capability to enable.
	// +kcc:proto:field=google.cloud.visionai.v1.SearchCapability.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.SearchCapabilitySetting
type SearchCapabilitySetting struct {
	// The metadata of search capability to enable.
	// +kcc:proto:field=google.cloud.visionai.v1.SearchCapabilitySetting.search_capabilities
	SearchCapabilities []SearchCapability `json:"searchCapabilities,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Corpus
type CorpusObservedState struct {
	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Separation enabled via an Org Policy constraint. It is set to true
	//  when the corpus is a valid zone separated corpus and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Isolation enabled via an Org Policy constraint. It is set to true when
	//  the corpus is a valid zone isolated corpus and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
