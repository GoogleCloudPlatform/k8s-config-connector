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

// +kcc:proto=google.cloud.documentai.v1.Processor
type Processor struct {

	// The processor type, such as: `OCR_PROCESSOR`, `INVOICE_PROCESSOR`.
	//  To get a list of processor types, see
	//  [FetchProcessorTypes][google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes].
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.type
	Type *string `json:"type,omitempty"`

	// The display name of the processor.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The default processor version.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.default_processor_version
	DefaultProcessorVersion *string `json:"defaultProcessorVersion,omitempty"`

	// The time the processor was created.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The [KMS key](https://cloud.google.com/security-key-management) used for
	//  encryption and decryption in CMEK scenarios.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.ProcessorVersionAlias
type ProcessorVersionAlias struct {
	// The alias in the form of `processor_version` resource name.
	// +kcc:proto:field=google.cloud.documentai.v1.ProcessorVersionAlias.alias
	Alias *string `json:"alias,omitempty"`

	// The resource name of aliased processor version.
	// +kcc:proto:field=google.cloud.documentai.v1.ProcessorVersionAlias.processor_version
	ProcessorVersion *string `json:"processorVersion,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Processor
type ProcessorObservedState struct {
	// Output only. Immutable. The resource name of the processor.
	//  Format: `projects/{project}/locations/{location}/processors/{processor}`
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.name
	Name *string `json:"name,omitempty"`

	// Output only. The state of the processor.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.state
	State *string `json:"state,omitempty"`

	// Output only. The processor version aliases.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.processor_version_aliases
	ProcessorVersionAliases []ProcessorVersionAlias `json:"processorVersionAliases,omitempty"`

	// Output only. Immutable. The http endpoint that can be called to invoke
	//  processing.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.process_endpoint
	ProcessEndpoint *string `json:"processEndpoint,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
