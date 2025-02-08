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


// +kcc:proto=google.cloud.discoveryengine.v1.CustomTuningModel
type CustomTuningModel struct {
	// Required. The fully qualified resource name of the model.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/customTuningModels/{custom_tuning_model}`.
	//
	//  Model must be an alpha-numerical string with limit of 40 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CustomTuningModel.name
	Name *string `json:"name,omitempty"`

	// The display name of the model.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CustomTuningModel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The version of the model.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CustomTuningModel.model_version
	ModelVersion *int64 `json:"modelVersion,omitempty"`

	// The state that the model is in (e.g.`TRAINING` or `TRAINING_FAILED`).
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CustomTuningModel.model_state
	ModelState *string `json:"modelState,omitempty"`

	// Deprecated: Timestamp the Model was created at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CustomTuningModel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Timestamp the model training was initiated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CustomTuningModel.training_start_time
	TrainingStartTime *string `json:"trainingStartTime,omitempty"`

	// TODO: unsupported map type with key string and value double


	// Currently this is only populated if the model state is
	//  `INPUT_VALIDATION_FAILED`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CustomTuningModel.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}
