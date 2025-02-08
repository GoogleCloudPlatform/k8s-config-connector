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


// +kcc:proto=google.cloud.dialogflow.v2.ConversationDataset
type ConversationDataset struct {

	// Required. The display name of the dataset. Maximum of 64 bytes.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the dataset. Maximum of 10000 bytes.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.ConversationInfo
type ConversationInfo struct {
	// Optional. The language code of the conversation data within this dataset.
	//  See https://cloud.google.com/apis/design/standard_fields for more
	//  information. Supports all UTF-8 languages.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationInfo.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.GcsSources
type GcsSources struct {
	// Required. Google Cloud Storage URIs for the inputs. A URI is of the form:
	//  `gs://bucket/object-prefix-or-name`
	//  Whether a prefix or name is used depends on the use case.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GcsSources.uris
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.InputConfig
type InputConfig struct {
	// The Cloud Storage URI has the form gs://<Google Cloud Storage bucket
	//  name>//agent*.json. Wildcards are allowed and will be expanded into all
	//  matched JSON files, which will be read as one conversation per file.
	// +kcc:proto:field=google.cloud.dialogflow.v2.InputConfig.gcs_source
	GcsSource *GcsSources `json:"gcsSource,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.ConversationDataset
type ConversationDatasetObservedState struct {
	// Output only. ConversationDataset resource name. Format:
	//  `projects/<Project ID>/locations/<Location
	//  ID>/conversationDatasets/<Conversation Dataset ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this dataset.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Input configurations set during conversation data import.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.input_config
	InputConfig *InputConfig `json:"inputConfig,omitempty"`

	// Output only. Metadata set during conversation data import.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.conversation_info
	ConversationInfo *ConversationInfo `json:"conversationInfo,omitempty"`

	// Output only. The number of conversations this conversation dataset
	//  contains.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.conversation_count
	ConversationCount *int64 `json:"conversationCount,omitempty"`

	// Output only. A read only boolean field reflecting Zone Isolation status of
	//  the dataset.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. A read only boolean field reflecting Zone Separation status of
	//  the dataset.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}
