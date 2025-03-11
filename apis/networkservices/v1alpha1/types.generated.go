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


// +kcc:proto=google.cloud.networkservices.v1.ServiceBinding
type ServiceBinding struct {
	// Required. Name of the ServiceBinding resource. It matches pattern
	//  `projects/*/locations/global/serviceBindings/service_binding_name`.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.name
	Name *string `json:"name,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.description
	Description *string `json:"description,omitempty"`

	// Required. The full service directory service name of the format
	//  /projects/*/locations/*/namespaces/*/services/*
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.service
	Service *string `json:"service,omitempty"`

	// Optional. Set of label tags associated with the ServiceBinding resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.ServiceBinding
type ServiceBindingObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
