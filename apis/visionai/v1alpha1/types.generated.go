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


// +kcc:proto=google.cloud.visionai.v1.DeployedIndex
type DeployedIndex struct {
	// Required. Name of the deployed Index.
	//  Format:
	//  `projects/{project_number}/locations/{location_id}/corpora/{corpus_id}/indexes/{index_id}`
	// +kcc:proto:field=google.cloud.visionai.v1.DeployedIndex.index
	Index *string `json:"index,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.IndexEndpoint
type IndexEndpoint struct {

	// Optional. Display name of the IndexEndpoint. Can be up to 32 characters
	//  long.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the IndexEndpoint. Can be up to 25000 characters
	//  long.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.description
	Description *string `json:"description,omitempty"`

	// Optional. The labels applied to a resource must meet the following
	//  requirements:
	//
	//  * Each resource can have multiple labels, up to a maximum of 64.
	//  * Each label must be a key-value pair.
	//  * Keys have a minimum length of 1 character and a maximum length of 63
	//    characters and cannot be empty. Values can be empty and have a maximum
	//    length of 63 characters.
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//    underscores, and dashes. All characters must use UTF-8 encoding, and
	//    international characters are allowed.
	//  * The key portion of a label must be unique. However, you can use the same
	//    key with multiple resources.
	//  * Keys must start with a lowercase letter or international character.
	//
	//  See [Google Cloud
	//  Document](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements)
	//  for more details.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.IndexEndpoint
type IndexEndpointObservedState struct {
	// Output only. Resource name of the IndexEndpoint.
	//  Format:
	//  `projects/{project}/locations/{location}/indexEndpoints/{index_endpoint_id}`
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.name
	Name *string `json:"name,omitempty"`

	// Output only. The Index deployed in this IndexEndpoint.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.deployed_index
	DeployedIndex *DeployedIndex `json:"deployedIndex,omitempty"`

	// Output only. IndexEndpoint state.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.state
	State *string `json:"state,omitempty"`

	// Output only. Create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Separation enabled via an Org Policy constraint. It is set to true
	//  when the index endpoint is a valid zone separated index endpoint and false
	//  if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Isolation enabled via an Org Policy constraint. It is set to true when
	//  the index endpoint is a valid zone isolated index endpoint and false if it
	//  isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
