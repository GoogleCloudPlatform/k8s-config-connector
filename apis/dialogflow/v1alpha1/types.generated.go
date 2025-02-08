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


// +kcc:proto=google.cloud.dialogflow.v2.KnowledgeBase
type KnowledgeBase struct {
	// The knowledge base resource name.
	//  The name must be empty when creating a knowledge base.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/knowledgeBases/<Knowledge Base ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.KnowledgeBase.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the knowledge base. The name must be 1024
	//  bytes or less; otherwise, the creation request fails.
	// +kcc:proto:field=google.cloud.dialogflow.v2.KnowledgeBase.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Language which represents the KnowledgeBase. When the KnowledgeBase is
	//  created/updated, expect this to be present for non en-us languages. When
	//  unspecified, the default language code en-us applies.
	// +kcc:proto:field=google.cloud.dialogflow.v2.KnowledgeBase.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}
