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


// +kcc:proto=google.cloud.discoveryengine.v1beta.Project
type Project struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Project.ServiceTerms
type Project_ServiceTerms struct {
	// The unique identifier of this terms of service.
	//  Available terms:
	//
	//  * `GA_DATA_USE_TERMS`: [Terms for data
	//  use](https://cloud.google.com/retail/data-use-terms). When using this as
	//  `id`, the acceptable
	//  [version][google.cloud.discoveryengine.v1beta.Project.ServiceTerms.version]
	//  to provide is `2022-11-23`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.ServiceTerms.id
	ID *string `json:"id,omitempty"`

	// The version string of the terms of service.
	//  For acceptable values, see the comments for
	//  [id][google.cloud.discoveryengine.v1beta.Project.ServiceTerms.id] above.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.ServiceTerms.version
	Version *string `json:"version,omitempty"`

	// Whether the project has accepted/rejected the service terms or it is
	//  still pending.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.ServiceTerms.state
	State *string `json:"state,omitempty"`

	// The last time when the project agreed to the terms of service.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.ServiceTerms.accept_time
	AcceptTime *string `json:"acceptTime,omitempty"`

	// The last time when the project declined or revoked the agreement to terms
	//  of service.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.ServiceTerms.decline_time
	DeclineTime *string `json:"declineTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Project
type ProjectObservedState struct {
	// Output only. Full resource name of the project, for example
	//  `projects/{project}`.
	//  Note that when making requests, project number and project id are both
	//  acceptable, but the server will always respond in project number.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when this project is created.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this project is successfully provisioned.
	//  Empty value means this project is still provisioning and is not ready for
	//  use.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Project.provision_completion_time
	ProvisionCompletionTime *string `json:"provisionCompletionTime,omitempty"`

	// TODO: unsupported map type with key string and value message

}
