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


// +kcc:proto=google.cloud.dialogflow.v2.ArticleSuggestionModelMetadata
type ArticleSuggestionModelMetadata struct {
	// Optional. Type of the article suggestion model. If not provided, model_type
	//  is used.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ArticleSuggestionModelMetadata.training_model_type
	TrainingModelType *string `json:"trainingModelType,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.ConversationModel
type ConversationModel struct {
	// ConversationModel resource name. Format:
	//  `projects/<Project ID>/conversationModels/<Conversation Model ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the model. At most 64 bytes long.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Datasets used to create model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.datasets
	Datasets []InputDataset `json:"datasets,omitempty"`

	// Language code for the conversation model. If not specified, the language
	//  is en-US. Language at ConversationModel should be set for all non en-us
	//  languages.
	//  This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt)
	//  language tag. Example: "en-US".
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Metadata for article suggestion models.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.article_suggestion_model_metadata
	ArticleSuggestionModelMetadata *ArticleSuggestionModelMetadata `json:"articleSuggestionModelMetadata,omitempty"`

	// Metadata for smart reply models.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.smart_reply_model_metadata
	SmartReplyModelMetadata *SmartReplyModelMetadata `json:"smartReplyModelMetadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.InputDataset
type InputDataset struct {
	// Required. ConversationDataset resource name. Format:
	//  `projects/<Project ID>/locations/<Location
	//  ID>/conversationDatasets/<Conversation Dataset ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2.InputDataset.dataset
	Dataset *string `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SmartReplyModelMetadata
type SmartReplyModelMetadata struct {
	// Optional. Type of the smart reply model. If not provided, model_type is
	//  used.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SmartReplyModelMetadata.training_model_type
	TrainingModelType *string `json:"trainingModelType,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.ConversationModel
type ConversationModelObservedState struct {
	// Output only. Creation time of this model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. State of the model. A model can only serve prediction requests
	//  after it gets deployed.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.state
	State *string `json:"state,omitempty"`

	// Output only. A read only boolean field reflecting Zone Separation
	//  status of the model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. A read only boolean field reflecting Zone Isolation status
	//  of the model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
