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


// +kcc:proto=google.cloud.securitycenter.v1.ResourceValueConfigMetadata
type ResourceValueConfigMetadata struct {
	// Resource value config name
	// +kcc:proto:field=google.cloud.securitycenter.v1.ResourceValueConfigMetadata.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.ValuedResource
type ValuedResource struct {
	// Valued resource name, for example,
	//   e.g.:
	//   `organizations/123/simulations/456/valuedResources/789`
	// +kcc:proto:field=google.cloud.securitycenter.v1.ValuedResource.name
	Name *string `json:"name,omitempty"`

	// The
	//  [full resource
	//  name](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	//  of the valued resource.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ValuedResource.resource
	Resource *string `json:"resource,omitempty"`

	// The [resource
	//  type](https://cloud.google.com/asset-inventory/docs/supported-asset-types)
	//  of the valued resource.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ValuedResource.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Human-readable name of the valued resource.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ValuedResource.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// How valuable this resource is.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ValuedResource.resource_value
	ResourceValue *string `json:"resourceValue,omitempty"`

	// Exposed score for this valued resource. A value of 0 means no exposure was
	//  detected exposure.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ValuedResource.exposed_score
	ExposedScore *float64 `json:"exposedScore,omitempty"`

	// List of resource value configurations' metadata used to determine the value
	//  of this resource. Maximum of 100.
	// +kcc:proto:field=google.cloud.securitycenter.v1.ValuedResource.resource_value_configs_used
	ResourceValueConfigsUsed []ResourceValueConfigMetadata `json:"resourceValueConfigsUsed,omitempty"`
}
