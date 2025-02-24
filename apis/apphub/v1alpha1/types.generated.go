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

// +kcc:proto=google.cloud.apphub.v1.ServiceProjectAttachment
type ServiceProjectAttachment struct {
	// Identifier. The resource name of a ServiceProjectAttachment. Format:
	//  "projects/{host-project-id}/locations/global/serviceProjectAttachments/{service-project-id}."
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProjectAttachment.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. Service project name in the format: "projects/abc" or
	//  "projects/123". As input, project name with either project id or number are
	//  accepted. As output, this field will contain project number.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProjectAttachment.service_project
	ServiceProject *string `json:"serviceProject,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.ServiceProjectAttachment
type ServiceProjectAttachmentObservedState struct {
	// Output only. Create time.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProjectAttachment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. A globally unique identifier (in UUID4 format) for the
	//  `ServiceProjectAttachment`.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProjectAttachment.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. ServiceProjectAttachment state.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProjectAttachment.state
	State *string `json:"state,omitempty"`
}
