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


// +kcc:proto=google.cloud.support.v2.Actor
type Actor struct {
	// The name to display for the actor. If not provided, it is inferred from
	//  credentials supplied during case creation. When an email is provided, a
	//  display name must also be provided. This will be obfuscated if the user
	//  is a Google Support agent.
	// +kcc:proto:field=google.cloud.support.v2.Actor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The email address of the actor. If not provided, it is inferred from
	//  credentials supplied during case creation. If the authenticated principal
	//  does not have an email address, one must be provided. When a name is
	//  provided, an email must also be provided. This will be obfuscated if the
	//  user is a Google Support agent.
	// +kcc:proto:field=google.cloud.support.v2.Actor.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.support.v2.Attachment
type Attachment struct {

	// The filename of the attachment (e.g. `"graph.jpg"`).
	// +kcc:proto:field=google.cloud.support.v2.Attachment.filename
	Filename *string `json:"filename,omitempty"`
}

// +kcc:proto=google.cloud.support.v2.Actor
type ActorObservedState struct {
	// Output only. Whether the actor is a Google support actor.
	// +kcc:proto:field=google.cloud.support.v2.Actor.google_support
	GoogleSupport *bool `json:"googleSupport,omitempty"`
}

// +kcc:proto=google.cloud.support.v2.Attachment
type AttachmentObservedState struct {
	// Output only. The resource name of the attachment.
	// +kcc:proto:field=google.cloud.support.v2.Attachment.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which the attachment was created.
	// +kcc:proto:field=google.cloud.support.v2.Attachment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The user who uploaded the attachment. Note, the name and email
	//  will be obfuscated if the attachment was uploaded by Google support.
	// +kcc:proto:field=google.cloud.support.v2.Attachment.creator
	Creator *Actor `json:"creator,omitempty"`

	// Output only. The MIME type of the attachment (e.g. text/plain).
	// +kcc:proto:field=google.cloud.support.v2.Attachment.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Output only. The size of the attachment in bytes.
	// +kcc:proto:field=google.cloud.support.v2.Attachment.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
}
