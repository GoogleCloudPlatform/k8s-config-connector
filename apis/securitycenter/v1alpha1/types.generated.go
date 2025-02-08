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


// +kcc:proto=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule
type EventThreatDetectionCustomModule struct {
	// Immutable. The resource name of the Event Threat Detection custom module.
	//
	//  Its format is:
	//
	//    * `organizations/{organization}/eventThreatDetectionSettings/customModules/{module}`.
	//    * `folders/{folder}/eventThreatDetectionSettings/customModules/{module}`.
	//    * `projects/{project}/eventThreatDetectionSettings/customModules/{module}`.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.name
	Name *string `json:"name,omitempty"`

	// Config for the module. For the resident module, its config value is defined
	//  at this level. For the inherited module, its config value is inherited from
	//  the ancestor module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.config
	Config map[string]string `json:"config,omitempty"`

	// The state of enablement for the module at the given level of the hierarchy.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.enablement_state
	EnablementState *string `json:"enablementState,omitempty"`

	// Type for the module. e.g. CONFIGURABLE_BAD_IP.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.type
	Type *string `json:"type,omitempty"`

	// The human readable name to be displayed for the module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description for the module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule
type EventThreatDetectionCustomModuleObservedState struct {
	// Output only. The closest ancestor module that this module inherits the
	//  enablement state from. The format is the same as the
	//  EventThreatDetectionCustomModule resource name.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.ancestor_module
	AncestorModule *string `json:"ancestorModule,omitempty"`

	// Output only. The time the module was last updated.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The editor the module was last updated by.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.last_editor
	LastEditor *string `json:"lastEditor,omitempty"`
}
