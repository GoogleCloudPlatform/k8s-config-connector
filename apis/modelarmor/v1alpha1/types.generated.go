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

// +generated:types
// krm.group: modelarmor.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.modelarmor.v1
// resource: ModelArmorTemplate:Template

package v1alpha1

// +kcc:proto=google.cloud.modelarmor.v1.FilterConfig
type FilterConfig struct {
	// Optional. Responsible AI settings.
	// +kcc:proto:field=google.cloud.modelarmor.v1.FilterConfig.rai_settings
	RaiSettings *RaiFilterSettings `json:"raiSettings,omitempty"`

	// Optional. Sensitive Data Protection settings.
	// +kcc:proto:field=google.cloud.modelarmor.v1.FilterConfig.sdp_settings
	SdpSettings *SdpFilterSettings `json:"sdpSettings,omitempty"`

	// Optional. Prompt injection and Jailbreak filter settings.
	// +kcc:proto:field=google.cloud.modelarmor.v1.FilterConfig.pi_and_jailbreak_filter_settings
	PiAndJailbreakFilterSettings *PiAndJailbreakFilterSettings `json:"piAndJailbreakFilterSettings,omitempty"`

	// Optional. Malicious URI filter settings.
	// +kcc:proto:field=google.cloud.modelarmor.v1.FilterConfig.malicious_uri_filter_settings
	MaliciousURIFilterSettings *MaliciousURIFilterSettings `json:"maliciousURIFilterSettings,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.MaliciousUriFilterSettings
type MaliciousURIFilterSettings struct {
	// Optional. Tells whether the Malicious URI filter is enabled or disabled.
	// +kcc:proto:field=google.cloud.modelarmor.v1.MaliciousUriFilterSettings.filter_enforcement
	FilterEnforcement *string `json:"filterEnforcement,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.PiAndJailbreakFilterSettings
type PiAndJailbreakFilterSettings struct {
	// Optional. Tells whether Prompt injection and Jailbreak filter is enabled or
	//  disabled.
	// +kcc:proto:field=google.cloud.modelarmor.v1.PiAndJailbreakFilterSettings.filter_enforcement
	FilterEnforcement *string `json:"filterEnforcement,omitempty"`

	// Optional. Confidence level for this filter.
	//  Confidence level is used to determine the threshold for the filter. If
	//  detection confidence is equal to or greater than the specified level, a
	//  positive match is reported. Confidence level will only be used if the
	//  filter is enabled.
	// +kcc:proto:field=google.cloud.modelarmor.v1.PiAndJailbreakFilterSettings.confidence_level
	ConfidenceLevel *string `json:"confidenceLevel,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.RaiFilterSettings
type RaiFilterSettings struct {
	// Required. List of Responsible AI filters enabled for template.
	// +kcc:proto:field=google.cloud.modelarmor.v1.RaiFilterSettings.rai_filters
	RaiFilters []RaiFilterSettings_RaiFilter `json:"raiFilters,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.RaiFilterSettings.RaiFilter
type RaiFilterSettings_RaiFilter struct {
	// Required. Type of responsible AI filter.
	// +kcc:proto:field=google.cloud.modelarmor.v1.RaiFilterSettings.RaiFilter.filter_type
	FilterType *string `json:"filterType,omitempty"`

	// Optional. Confidence level for this RAI filter.
	//  During data sanitization, if data is classified under this filter with a
	//  confidence level equal to or greater than the specified level, a positive
	//  match is reported. If the confidence level is unspecified (i.e., 0), the
	//  system will use a reasonable default level based on the `filter_type`.
	// +kcc:proto:field=google.cloud.modelarmor.v1.RaiFilterSettings.RaiFilter.confidence_level
	ConfidenceLevel *string `json:"confidenceLevel,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.SdpAdvancedConfig
type SdpAdvancedConfig struct {
	// Optional. Sensitive Data Protection inspect template resource name
	//
	//  If only inspect template is provided (de-identify template not provided),
	//  then Sensitive Data Protection InspectContent action is performed during
	//  Sanitization. All Sensitive Data Protection findings identified during
	//  inspection will be returned as SdpFinding in SdpInsepctionResult e.g.
	//  `organizations/{organization}/inspectTemplates/{inspect_template}`,
	//  `projects/{project}/inspectTemplates/{inspect_template}`
	//  `organizations/{organization}/locations/{location}/inspectTemplates/{inspect_template}`
	//  `projects/{project}/locations/{location}/inspectTemplates/{inspect_template}`
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpAdvancedConfig.inspect_template
	InspectTemplate *string `json:"inspectTemplate,omitempty"`

	// Optional. Optional Sensitive Data Protection Deidentify template resource
	//  name.
	//
	//  If provided then DeidentifyContent action is performed during Sanitization
	//  using this template and inspect template. The De-identified data will
	//  be returned in SdpDeidentifyResult.
	//  Note that all info-types present in the deidentify template must be present
	//  in inspect template.
	//
	//  e.g.
	//  `organizations/{organization}/deidentifyTemplates/{deidentify_template}`,
	//  `projects/{project}/deidentifyTemplates/{deidentify_template}`
	//  `organizations/{organization}/locations/{location}/deidentifyTemplates/{deidentify_template}`
	//  `projects/{project}/locations/{location}/deidentifyTemplates/{deidentify_template}`
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpAdvancedConfig.deidentify_template
	DeidentifyTemplate *string `json:"deidentifyTemplate,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.SdpBasicConfig
type SdpBasicConfig struct {
	// Optional. Tells whether the Sensitive Data Protection basic config is
	//  enabled or disabled.
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpBasicConfig.filter_enforcement
	FilterEnforcement *string `json:"filterEnforcement,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.SdpFilterSettings
type SdpFilterSettings struct {
	// Optional. Basic Sensitive Data Protection configuration inspects the
	//  content for sensitive data using a fixed set of six info-types. Sensitive
	//  Data Protection templates cannot be used with basic configuration. Only
	//  Sensitive Data Protection inspection operation is supported with basic
	//  configuration.
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpFilterSettings.basic_config
	BasicConfig *SdpBasicConfig `json:"basicConfig,omitempty"`

	// Optional. Advanced Sensitive Data Protection configuration which enables
	//  use of Sensitive Data Protection templates. Supports both Sensitive Data
	//  Protection inspection and de-identification operations.
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpFilterSettings.advanced_config
	AdvancedConfig *SdpAdvancedConfig `json:"advancedConfig,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.Template.TemplateMetadata
type Template_TemplateMetadata struct {
	// Optional. If true, partial detector failures should be ignored.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.ignore_partial_invocation_failures
	IgnorePartialInvocationFailures *bool `json:"ignorePartialInvocationFailures,omitempty"`

	// Optional. Indicates the custom error code set by the user to be returned
	//  to the end user by the service extension if the prompt trips Model Armor
	//  filters.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.custom_prompt_safety_error_code
	CustomPromptSafetyErrorCode *int32 `json:"customPromptSafetyErrorCode,omitempty"`

	// Optional. Indicates the custom error message set by the user to be
	//  returned to the end user if the prompt trips Model Armor filters.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.custom_prompt_safety_error_message
	CustomPromptSafetyErrorMessage *string `json:"customPromptSafetyErrorMessage,omitempty"`

	// Optional. Indicates the custom error code set by the user to be returned
	//  to the end user if the LLM response trips Model Armor filters.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.custom_llm_response_safety_error_code
	CustomLlmResponseSafetyErrorCode *int32 `json:"customLlmResponseSafetyErrorCode,omitempty"`

	// Optional. Indicates the custom error message set by the user to be
	//  returned to the end user if the LLM response trips Model Armor filters.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.custom_llm_response_safety_error_message
	CustomLlmResponseSafetyErrorMessage *string `json:"customLlmResponseSafetyErrorMessage,omitempty"`

	// Optional. If true, log template crud operations.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.log_template_operations
	LogTemplateOperations *bool `json:"logTemplateOperations,omitempty"`

	// Optional. If true, log sanitize operations.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.log_sanitize_operations
	LogSanitizeOperations *bool `json:"logSanitizeOperations,omitempty"`
}
