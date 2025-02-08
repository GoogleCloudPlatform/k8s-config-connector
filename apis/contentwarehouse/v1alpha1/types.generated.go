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


// +kcc:proto=google.cloud.contentwarehouse.v1.SynonymSet
type SynonymSet struct {
	// The resource name of the SynonymSet
	//  This is mandatory for google.api.resource.
	//  Format:
	//  projects/{project_number}/locations/{location}/synonymSets/{context}.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.SynonymSet.name
	Name *string `json:"name,omitempty"`

	// This is a freeform field. Example contexts can be "sales," "engineering,"
	//  "real estate," "accounting," etc.
	//  The context can be supplied during search requests.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.SynonymSet.context
	Context *string `json:"context,omitempty"`

	// List of Synonyms for the context.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.SynonymSet.synonyms
	Synonyms []SynonymSet_Synonym `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.SynonymSet.Synonym
type SynonymSet_Synonym struct {
	// For example: sale, invoice, bill, order
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.SynonymSet.Synonym.words
	Words []string `json:"words,omitempty"`
}
