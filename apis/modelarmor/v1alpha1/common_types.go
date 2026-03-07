// Copyright 2026 Google LLC
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

// +kcc:proto=google.cloud.modelarmor.v1.SdpAdvancedConfig
type SdpAdvancedConfig struct {
	// Optional. Sensitive Data Protection inspect template resource name
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpAdvancedConfig.inspect_template
	InspectTemplateRef *DLPInspectTemplateRef `json:"inspectTemplateRef,omitempty"`

	// Optional. Optional Sensitive Data Protection Deidentify template resource name.
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpAdvancedConfig.deidentify_template
	DeidentifyTemplateRef *DLPDeidentifyTemplateRef `json:"deidentifyTemplateRef,omitempty"`
}

type DLPInspectTemplateRef struct {
	/* The `name` of a `DLPInspectTemplate` resource. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `DLPInspectTemplate` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `DLPInspectTemplate` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type DLPDeidentifyTemplateRef struct {
	/* The `name` of a `DLPDeidentifyTemplate` resource. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `DLPDeidentifyTemplate` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `DLPDeidentifyTemplate` resource. */
	Namespace string `json:"namespace,omitempty"`
}
