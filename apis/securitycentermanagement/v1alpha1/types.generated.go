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


// +kcc:proto=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule
type EffectiveEventThreatDetectionCustomModule struct {
	// Identifier. The resource name of the Event Threat Detection custom module,
	//  in one of the following formats:
	//
	//  * `organizations/{organization}/locations/{location}/effectiveEventThreatDetectionCustomModules/{custom_module}`
	//  * `folders/{folder}/locations/{location}/effectiveEventThreatDetectionCustomModules/{custom_module}`
	//  * `projects/{project}/locations/{location}/effectiveEventThreatDetectionCustomModules/{custom_module}`
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule
type EffectiveEventThreatDetectionCustomModuleObservedState struct {
	// Output only. Configuration for the effective module.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule.config
	Config map[string]string `json:"config,omitempty"`

	// Output only. The effective state of enablement for the module at the given
	//  level of the hierarchy.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule.enablement_state
	EnablementState *string `json:"enablementState,omitempty"`

	// Output only. Type for the module (for example, `CONFIGURABLE_BAD_IP`).
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule.type
	Type *string `json:"type,omitempty"`

	// Output only. The human-readable name of the module.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. A description of the module.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EffectiveEventThreatDetectionCustomModule.description
	Description *string `json:"description,omitempty"`
}
