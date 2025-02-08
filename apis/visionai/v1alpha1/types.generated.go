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


// +kcc:proto=google.cloud.visionai.v1.DeployedIndexReference
type DeployedIndexReference struct {
	// Immutable. A resource name of the IndexEndpoint.
	// +kcc:proto:field=google.cloud.visionai.v1.DeployedIndexReference.index_endpoint
	IndexEndpoint *string `json:"indexEndpoint,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Index
type Index struct {
	// Include all assets under the corpus.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.entire_corpus
	EntireCorpus *bool `json:"entireCorpus,omitempty"`

	// Optional. Optional user-specified display name of the index.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Optional description of the index.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Index
type IndexObservedState struct {
	// Output only. Resource name of the Index resource.
	//  Format:
	//  `projects/{project_number}/locations/{location_id}/corpora/{corpus_id}/indexes/{index_id}`
	// +kcc:proto:field=google.cloud.visionai.v1.Index.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the index.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.state
	State *string `json:"state,omitempty"`

	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. References to the deployed index instance.
	//  Index of VIDEO_ON_DEMAND corpus can have at most one deployed index.
	//  Index of IMAGE corpus can have multiple deployed indexes.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.deployed_indexes
	DeployedIndexes []DeployedIndexReference `json:"deployedIndexes,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Separation enabled via an Org Policy constraint. It is set to true
	//  when the index is a valid zone separated index and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Isolation enabled via an Org Policy constraint. It is set to true when
	//  the index is a valid zone isolated index and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
