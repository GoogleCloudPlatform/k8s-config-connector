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


// +kcc:proto=google.cloud.translation.v3.Model
type Model struct {
	// The resource name of the model, in form of
	//  `projects/{project-number-or-id}/locations/{location_id}/models/{model_id}`
	// +kcc:proto:field=google.cloud.translation.v3.Model.name
	Name *string `json:"name,omitempty"`

	// The name of the model to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores (_), and ASCII digits 0-9.
	// +kcc:proto:field=google.cloud.translation.v3.Model.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The dataset from which the model is trained, in form of
	//  `projects/{project-number-or-id}/locations/{location_id}/datasets/{dataset_id}`
	// +kcc:proto:field=google.cloud.translation.v3.Model.dataset
	Dataset *string `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.Model
type ModelObservedState struct {
	// Output only. The BCP-47 language code of the source language.
	// +kcc:proto:field=google.cloud.translation.v3.Model.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// Output only. The BCP-47 language code of the target language.
	// +kcc:proto:field=google.cloud.translation.v3.Model.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`

	// Output only. Number of examples (sentence pairs) used to train the model.
	// +kcc:proto:field=google.cloud.translation.v3.Model.train_example_count
	TrainExampleCount *int32 `json:"trainExampleCount,omitempty"`

	// Output only. Number of examples (sentence pairs) used to validate the
	//  model.
	// +kcc:proto:field=google.cloud.translation.v3.Model.validate_example_count
	ValidateExampleCount *int32 `json:"validateExampleCount,omitempty"`

	// Output only. Number of examples (sentence pairs) used to test the model.
	// +kcc:proto:field=google.cloud.translation.v3.Model.test_example_count
	TestExampleCount *int32 `json:"testExampleCount,omitempty"`

	// Output only. Timestamp when the model resource was created, which is also
	//  when the training started.
	// +kcc:proto:field=google.cloud.translation.v3.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this model was last updated.
	// +kcc:proto:field=google.cloud.translation.v3.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
