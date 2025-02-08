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


// +kcc:proto=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule
type EventThreatDetectionCustomModule struct {
	// Identifier. The resource name of the Event Threat Detection custom module,
	//  in one of the following formats:
	//
	//  * `organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{custom_module}`
	//  * `folders/{folder}/locations/{location}/eventThreatDetectionCustomModules/{custom_module}`
	//  * `projects/{project}/locations/{location}/eventThreatDetectionCustomModules/{custom_module}`
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.name
	Name *string `json:"name,omitempty"`

	// Optional. Configuration for the module. For the resident module, its
	//  configuration value is defined at this level. For the inherited module, its
	//  configuration value is inherited from the ancestor module.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.config
	Config map[string]string `json:"config,omitempty"`

	// Optional. The state of enablement for the module at the given level of the
	//  hierarchy.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.enablement_state
	EnablementState *string `json:"enablementState,omitempty"`

	// Optional. Type for the module. For example, `CONFIGURABLE_BAD_IP`.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.type
	Type *string `json:"type,omitempty"`

	// Optional. The human-readable name of the module.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A description of the module.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule
type EventThreatDetectionCustomModuleObservedState struct {
	// Output only. The closest ancestor module that this module inherits the
	//  enablement state from. If empty, indicates that the custom module was
	//  created in the requesting parent organization, folder, or project. The
	//  format is the same as the custom module's resource name.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.ancestor_module
	AncestorModule *string `json:"ancestorModule,omitempty"`

	// Output only. The time the module was last updated.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The editor the module was last updated by.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule.last_editor
	LastEditor *string `json:"lastEditor,omitempty"`
}
