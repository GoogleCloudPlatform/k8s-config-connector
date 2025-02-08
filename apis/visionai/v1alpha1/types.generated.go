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


// +kcc:proto=google.cloud.visionai.v1.SearchHypernym
type SearchHypernym struct {
	// Resource name of the SearchHypernym.
	//  Format:
	//  `projects/{project_number}/locations/{location}/corpora/{corpus}/searchHypernyms/{search_hypernym}`
	// +kcc:proto:field=google.cloud.visionai.v1.SearchHypernym.name
	Name *string `json:"name,omitempty"`

	// Optional. The hypernym.
	// +kcc:proto:field=google.cloud.visionai.v1.SearchHypernym.hypernym
	Hypernym *string `json:"hypernym,omitempty"`

	// Optional. Hyponyms that the hypernym is mapped to.
	// +kcc:proto:field=google.cloud.visionai.v1.SearchHypernym.hyponyms
	Hyponyms []string `json:"hyponyms,omitempty"`
}
