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


// +kcc:proto=google.cloud.visionai.v1.Collection
type Collection struct {

	// Optional. The collection name for displaying.
	//  The name can be up to 256 characters long.
	// +kcc:proto:field=google.cloud.visionai.v1.Collection.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the collection. Can be up to 25000 characters
	//  long.
	// +kcc:proto:field=google.cloud.visionai.v1.Collection.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Collection
type CollectionObservedState struct {
	// Output only. Resource name of the collection. Format:
	//  `projects/{project_number}/locations/{location}/corpora/{corpus}/collections/{collection}`
	// +kcc:proto:field=google.cloud.visionai.v1.Collection.name
	Name *string `json:"name,omitempty"`
}
