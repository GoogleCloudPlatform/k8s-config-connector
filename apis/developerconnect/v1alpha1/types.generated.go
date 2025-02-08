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


// +kcc:proto=google.cloud.developerconnect.v1.GitRepositoryLink
type GitRepositoryLink struct {
	// Identifier. Resource name of the repository, in the format
	//  `projects/*/locations/*/connections/*/gitRepositoryLinks/*`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.name
	Name *string `json:"name,omitempty"`

	// Required. Git Clone URI.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.clone_uri
	CloneURI *string `json:"cloneURI,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Allows clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitRepositoryLink
type GitRepositoryLinkObservedState struct {
	// Output only. [Output only] Create timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. [Output only] Delete timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. Set to true when the connection is being set up or updated in
	//  the background.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. A system-assigned unique identifier for a the
	//  GitRepositoryLink.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. External ID of the webhook created for the repository.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitRepositoryLink.webhook_id
	WebhookID *string `json:"webhookID,omitempty"`
}
