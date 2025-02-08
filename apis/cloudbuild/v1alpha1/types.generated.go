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


// +kcc:proto=google.devtools.cloudbuild.v2.Repository
type Repository struct {
	// Immutable. Resource name of the repository, in the format
	//  `projects/*/locations/*/connections/*/repositories/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Repository.name
	Name *string `json:"name,omitempty"`

	// Required. Git Clone HTTPS URI.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Repository.remote_uri
	RemoteURI *string `json:"remoteURI,omitempty"`

	// Allows clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Repository.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Repository.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.Repository
type RepositoryObservedState struct {
	// Output only. Server assigned timestamp for when the connection was created.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Repository.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Server assigned timestamp for when the connection was updated.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Repository.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. External ID of the webhook created for the repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Repository.webhook_id
	WebhookID *string `json:"webhookID,omitempty"`
}
