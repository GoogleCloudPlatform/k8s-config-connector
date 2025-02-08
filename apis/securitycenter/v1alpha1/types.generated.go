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


// +kcc:proto=google.cloud.securitycenter.v2.ResourceValueConfigMetadata
type ResourceValueConfigMetadata struct {
	// Resource value config name
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfigMetadata.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.Simulation
type Simulation struct {
	// Full resource name of the Simulation:
	//  `organizations/123/simulations/456`
	// +kcc:proto:field=google.cloud.securitycenter.v2.Simulation.name
	Name *string `json:"name,omitempty"`

	// Resource value configurations' metadata used in this simulation. Maximum of
	//  100.
	// +kcc:proto:field=google.cloud.securitycenter.v2.Simulation.resource_value_configs_metadata
	ResourceValueConfigsMetadata []ResourceValueConfigMetadata `json:"resourceValueConfigsMetadata,omitempty"`

	// Indicates which cloud provider was used in this simulation.
	// +kcc:proto:field=google.cloud.securitycenter.v2.Simulation.cloud_provider
	CloudProvider *string `json:"cloudProvider,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.Simulation
type SimulationObservedState struct {
	// Output only. Time simulation was created
	// +kcc:proto:field=google.cloud.securitycenter.v2.Simulation.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
