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


// +kcc:proto=google.cloud.securitycentermanagement.v1.SecurityCenterService
type SecurityCenterService struct {
	// Identifier. The name of the service, in one of the following formats:
	//
	//  * `organizations/{organization}/locations/{location}/securityCenterServices/{service}`
	//  * `folders/{folder}/locations/{location}/securityCenterServices/{service}`
	//  * `projects/{project}/locations/{location}/securityCenterServices/{service}`
	//
	//  The following values are valid for `{service}`:
	//
	//  * `container-threat-detection`
	//  * `event-threat-detection`
	//  * `security-health-analytics`
	//  * `vm-threat-detection`
	//  * `web-security-scanner`
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SecurityCenterService.name
	Name *string `json:"name,omitempty"`

	// Optional. The intended enablement state for the service at its level of the
	//  resource hierarchy. A `DISABLED` state will override all module enablement
	//  states to `DISABLED`.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SecurityCenterService.intended_enablement_state
	IntendedEnablementState *string `json:"intendedEnablementState,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. Additional service-specific configuration. Not all services will
	//  utilize this field.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SecurityCenterService.service_config
	ServiceConfig map[string]string `json:"serviceConfig,omitempty"`
}

// +kcc:proto=google.cloud.securitycentermanagement.v1.SecurityCenterService.ModuleSettings
type SecurityCenterService_ModuleSettings struct {
	// Optional. The intended enablement state for the module at its level of
	//  the resource hierarchy.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SecurityCenterService.ModuleSettings.intended_enablement_state
	IntendedEnablementState *string `json:"intendedEnablementState,omitempty"`
}

// +kcc:proto=google.cloud.securitycentermanagement.v1.SecurityCenterService
type SecurityCenterServiceObservedState struct {
	// Output only. The effective enablement state for the service at its level of
	//  the resource hierarchy. If the intended state is set to `INHERITED`, the
	//  effective state will be inherited from the enablement state of an ancestor.
	//  This state may differ from the intended enablement state due to billing
	//  eligibility or onboarding status.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SecurityCenterService.effective_enablement_state
	EffectiveEnablementState *string `json:"effectiveEnablementState,omitempty"`

	// Output only. The time the service was last updated. This could be due to an
	//  explicit user update or due to a side effect of another system change, such
	//  as billing subscription expiry.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SecurityCenterService.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
