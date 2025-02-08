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


// +kcc:proto=google.cloud.eventarc.v1.Enrollment
type Enrollment struct {
	// Identifier. Resource name of the form
	//  projects/{project}/locations/{location}/enrollments/{enrollment}
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.name
	Name *string `json:"name,omitempty"`

	// Optional. Resource labels.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Resource annotations.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Resource display name.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. A CEL expression identifying which messages this enrollment
	//  applies to.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.cel_match
	CelMatch *string `json:"celMatch,omitempty"`

	// Required. Resource name of the message bus identifying the source of the
	//  messages. It matches the form
	//  projects/{project}/locations/{location}/messageBuses/{messageBus}.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.message_bus
	MessageBus *string `json:"messageBus,omitempty"`

	// Required. Destination is the Pipeline that the Enrollment is delivering to.
	//  It must point to the full resource name of a Pipeline. Format:
	//  "projects/{PROJECT_ID}/locations/{region}/pipelines/{PIPELINE_ID)"
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.destination
	Destination *string `json:"destination,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Enrollment
type EnrollmentObservedState struct {
	// Output only. Server assigned unique identifier for the channel. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields, and might be sent only on update and delete requests to
	//  ensure that the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
