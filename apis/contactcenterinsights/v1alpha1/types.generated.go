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


// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueModel
type IssueModel struct {
	// Immutable. The resource name of the issue model.
	//  Format:
	//  projects/{project}/locations/{location}/issueModels/{issue_model}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.name
	Name *string `json:"name,omitempty"`

	// The representative name for the issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Configs for the input data that used to create the issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.input_data_config
	InputDataConfig *IssueModel_InputDataConfig `json:"inputDataConfig,omitempty"`

	// Type of the model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.model_type
	ModelType *string `json:"modelType,omitempty"`

	// Language of the model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueModel.InputDataConfig
type IssueModel_InputDataConfig struct {
	// Medium of conversations used in training data. This field is being
	//  deprecated. To specify the medium to be used in training a new issue
	//  model, set the `medium` field on `filter`.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.InputDataConfig.medium
	Medium *string `json:"medium,omitempty"`

	// A filter to reduce the conversations used for training the model to a
	//  specific subset.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.InputDataConfig.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueModelLabelStats
type IssueModelLabelStats struct {
	// Number of conversations the issue model has analyzed at this point in time.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModelLabelStats.analyzed_conversations_count
	AnalyzedConversationsCount *int64 `json:"analyzedConversationsCount,omitempty"`

	// Number of analyzed conversations for which no issue was applicable at this
	//  point in time.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModelLabelStats.unclassified_conversations_count
	UnclassifiedConversationsCount *int64 `json:"unclassifiedConversationsCount,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueModelLabelStats.IssueStats
type IssueModelLabelStats_IssueStats struct {
	// Issue resource.
	//  Format:
	//  projects/{project}/locations/{location}/issueModels/{issue_model}/issues/{issue}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModelLabelStats.IssueStats.issue
	Issue *string `json:"issue,omitempty"`

	// Number of conversations attached to the issue at this point in time.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModelLabelStats.IssueStats.labeled_conversations_count
	LabeledConversationsCount *int64 `json:"labeledConversationsCount,omitempty"`

	// Display name of the issue.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModelLabelStats.IssueStats.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueModel
type IssueModelObservedState struct {
	// Output only. The time at which this issue model was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the issue model was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Number of issues in this issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.issue_count
	IssueCount *int64 `json:"issueCount,omitempty"`

	// Output only. State of the model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.state
	State *string `json:"state,omitempty"`

	// Configs for the input data that used to create the issue model.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.input_data_config
	InputDataConfig *IssueModel_InputDataConfigObservedState `json:"inputDataConfig,omitempty"`

	// Output only. Immutable. The issue model's label statistics on its training
	//  data.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.training_stats
	TrainingStats *IssueModelLabelStats `json:"trainingStats,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.IssueModel.InputDataConfig
type IssueModel_InputDataConfigObservedState struct {
	// Output only. Number of conversations used in training. Output only.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.IssueModel.InputDataConfig.training_conversations_count
	TrainingConversationsCount *int64 `json:"trainingConversationsCount,omitempty"`
}
