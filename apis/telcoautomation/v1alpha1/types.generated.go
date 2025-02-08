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


// +kcc:proto=google.cloud.telcoautomation.v1.PublicBlueprint
type PublicBlueprint struct {
	// Name of the public blueprint.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.PublicBlueprint.name
	Name *string `json:"name,omitempty"`

	// The display name of the public blueprint.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.PublicBlueprint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the public blueprint.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.PublicBlueprint.description
	Description *string `json:"description,omitempty"`

	// DeploymentLevel of a blueprint signifies where the blueprint will be
	//  applied. e.g. [HYDRATION, SINGLE_DEPLOYMENT, MULTI_DEPLOYMENT]
	// +kcc:proto:field=google.cloud.telcoautomation.v1.PublicBlueprint.deployment_level
	DeploymentLevel *string `json:"deploymentLevel,omitempty"`

	// Source provider is the author of a public blueprint. e.g. Google, vendors
	// +kcc:proto:field=google.cloud.telcoautomation.v1.PublicBlueprint.source_provider
	SourceProvider *string `json:"sourceProvider,omitempty"`
}

// +kcc:proto=google.cloud.telcoautomation.v1.PublicBlueprint
type PublicBlueprintObservedState struct {
	// Output only. Indicates if the deployment created from this blueprint can be
	//  rolled back.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.PublicBlueprint.rollback_support
	RollbackSupport *bool `json:"rollbackSupport,omitempty"`
}
