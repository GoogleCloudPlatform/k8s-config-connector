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


// +kcc:proto=google.cloud.translation.v3.Dataset
type Dataset struct {
	// The resource name of the dataset, in form of
	//  `projects/{project-number-or-id}/locations/{location_id}/datasets/{dataset_id}`
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.name
	Name *string `json:"name,omitempty"`

	// The name of the dataset to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores (_), and ASCII digits 0-9.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The BCP-47 language code of the source language.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// The BCP-47 language code of the target language.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.Dataset
type DatasetObservedState struct {
	// Output only. The number of examples in the dataset.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.example_count
	ExampleCount *int32 `json:"exampleCount,omitempty"`

	// Output only. Number of training examples (sentence pairs).
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.train_example_count
	TrainExampleCount *int32 `json:"trainExampleCount,omitempty"`

	// Output only. Number of validation examples (sentence pairs).
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.validate_example_count
	ValidateExampleCount *int32 `json:"validateExampleCount,omitempty"`

	// Output only. Number of test examples (sentence pairs).
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.test_example_count
	TestExampleCount *int32 `json:"testExampleCount,omitempty"`

	// Output only. Timestamp when this dataset was created.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this dataset was last updated.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
