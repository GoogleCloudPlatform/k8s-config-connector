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


// +kcc:proto=google.cloud.securityposture.v1.PostureDeployment
type PostureDeployment struct {
	// Required. The name of this PostureDeployment resource, in the format of
	//  organizations/{organization}/locations/{location_id}/postureDeployments/{postureDeployment}.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.name
	Name *string `json:"name,omitempty"`

	// Required. Target resource where the Posture will be deployed. Currently
	//  supported resources are of types: projects/projectNumber,
	//  folders/folderNumber, organizations/organizationNumber.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.target_resource
	TargetResource *string `json:"targetResource,omitempty"`

	// Required. Posture that needs to be deployed.
	//  Format:
	//  organizations/{org_id}/locations/{location_id}/postures/<posture>
	//  Example:
	//  organizations/99/locations/global/postures/les-miserables.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.posture_id
	PostureID *string `json:"postureID,omitempty"`

	// Required. Revision_id of the Posture that is to be deployed.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.posture_revision_id
	PostureRevisionID *string `json:"postureRevisionID,omitempty"`

	// Optional. User provided description of the PostureDeployment.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.description
	Description *string `json:"description,omitempty"`

	// Optional. An opaque tag indicating the current version of the
	//  PostureDeployment, used for concurrency control. When the
	//  `PostureDeployment` is returned from either a `GetPostureDeployment` or a
	//  `ListPostureDeployments` request, this `etag` indicates the version of the
	//  current `PostureDeployment` to use when executing a read-modify-write loop.
	//
	//  When the `PostureDeployment` is used in a `UpdatePostureDeployment` method,
	//  use the `etag` value that was returned from a `GetPostureDeployment`
	//  request as part of a read-modify-write loop for concurrency control. Not
	//  setting the `etag` in a `UpdatePostureDeployment` request will result in an
	//  unconditional write of the `PostureDeployment`.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Google Security Postures.
	//  .
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.securityposture.v1.PostureDeployment
type PostureDeploymentObservedState struct {
	// Output only. State of PostureDeployment resource.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.state
	State *string `json:"state,omitempty"`

	// Output only. The timestamp that the PostureDeployment was created.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp that the PostureDeployment was updated.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Whether or not this Posture is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. This is a output only optional field which will be filled in
	//  case where PostureDeployment state is UPDATE_FAILED or CREATE_FAILED or
	//  DELETE_FAILED. It denotes the desired Posture.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.desired_posture_id
	DesiredPostureID *string `json:"desiredPostureID,omitempty"`

	// Output only. Output only optional field which provides revision_id of the
	//  desired_posture_id.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.desired_posture_revision_id
	DesiredPostureRevisionID *string `json:"desiredPostureRevisionID,omitempty"`

	// Output only. This is a output only optional field which will be filled in
	//  case where PostureDeployment enters a failure state like UPDATE_FAILED or
	//  CREATE_FAILED or DELETE_FAILED.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.failure_message
	FailureMessage *string `json:"failureMessage,omitempty"`
}
