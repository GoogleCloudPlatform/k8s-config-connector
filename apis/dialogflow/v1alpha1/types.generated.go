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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Changelog
type Changelog struct {
	// The unique identifier of the changelog.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/changelogs/<ChangelogID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.name
	Name *string `json:"name,omitempty"`

	// Email address of the authenticated user.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.user_email
	UserEmail *string `json:"userEmail,omitempty"`

	// The affected resource display name of the change.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The action of the change.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.action
	Action *string `json:"action,omitempty"`

	// The affected resource type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.type
	Type *string `json:"type,omitempty"`

	// The affected resource name of the change.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.resource
	Resource *string `json:"resource,omitempty"`

	// The timestamp of the change.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The affected language code of the change.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Changelog.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}
