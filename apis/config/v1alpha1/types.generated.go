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


// +kcc:proto=google.cloud.config.v1.Resource
type Resource struct {
}

// +kcc:proto=google.cloud.config.v1.ResourceCAIInfo
type ResourceCAIInfo struct {
	// CAI resource name in the format following
	//  https://cloud.google.com/apis/design/resource_names#full_resource_name
	// +kcc:proto:field=google.cloud.config.v1.ResourceCAIInfo.full_resource_name
	FullResourceName *string `json:"fullResourceName,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.ResourceTerraformInfo
type ResourceTerraformInfo struct {
	// TF resource address that uniquely identifies this resource within this
	//  deployment.
	// +kcc:proto:field=google.cloud.config.v1.ResourceTerraformInfo.address
	Address *string `json:"address,omitempty"`

	// TF resource type
	// +kcc:proto:field=google.cloud.config.v1.ResourceTerraformInfo.type
	Type *string `json:"type,omitempty"`

	// ID attribute of the TF resource
	// +kcc:proto:field=google.cloud.config.v1.ResourceTerraformInfo.id
	ID *string `json:"id,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.Resource
type ResourceObservedState struct {
	// Output only. Resource name.
	//  Format:
	//  `projects/{project}/locations/{location}/deployments/{deployment}/revisions/{revision}/resources/{resource}`
	// +kcc:proto:field=google.cloud.config.v1.Resource.name
	Name *string `json:"name,omitempty"`

	// Output only. Terraform-specific info if this resource was created using
	//  Terraform.
	// +kcc:proto:field=google.cloud.config.v1.Resource.terraform_info
	TerraformInfo *ResourceTerraformInfo `json:"terraformInfo,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. Intent of the resource.
	// +kcc:proto:field=google.cloud.config.v1.Resource.intent
	Intent *string `json:"intent,omitempty"`

	// Output only. Current state of the resource.
	// +kcc:proto:field=google.cloud.config.v1.Resource.state
	State *string `json:"state,omitempty"`
}
