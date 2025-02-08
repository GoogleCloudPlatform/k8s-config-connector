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


// +kcc:proto=google.cloud.translation.v3.Example
type Example struct {

	// Sentence in source language.
	// +kcc:proto:field=google.cloud.translation.v3.Example.source_text
	SourceText *string `json:"sourceText,omitempty"`

	// Sentence in target language.
	// +kcc:proto:field=google.cloud.translation.v3.Example.target_text
	TargetText *string `json:"targetText,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.Example
type ExampleObservedState struct {
	// Output only. The resource name of the example, in form of
	//  `projects/{project-number-or-id}/locations/{location_id}/datasets/{dataset_id}/examples/{example_id}`
	// +kcc:proto:field=google.cloud.translation.v3.Example.name
	Name *string `json:"name,omitempty"`

	// Output only. Usage of the sentence pair. Options are TRAIN|VALIDATION|TEST.
	// +kcc:proto:field=google.cloud.translation.v3.Example.usage
	Usage *string `json:"usage,omitempty"`
}
