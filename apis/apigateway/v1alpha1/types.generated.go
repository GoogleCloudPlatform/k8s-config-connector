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

// +kcc:proto=google.cloud.apigateway.v1.Api
type Api struct {

	// Optional. Resource labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Display name.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Immutable. The name of a Google Managed Service (
	//  https://cloud.google.com/service-infrastructure/docs/glossary#managed). If
	//  not specified, a new Service will automatically be created in the same
	//  project as this API.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.managed_service
	ManagedService *string `json:"managedService,omitempty"`
}

// +kcc:proto=google.cloud.apigateway.v1.Api
type ApiObservedState struct {
	// Output only. Resource name of the API.
	//  Format: projects/{project}/locations/global/apis/{api}
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.name
	Name *string `json:"name,omitempty"`

	// Output only. Created time.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Updated time.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the API.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.state
	State *string `json:"state,omitempty"`
}
