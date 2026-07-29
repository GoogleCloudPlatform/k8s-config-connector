// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	modelarmorv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineAssistantGVK = GroupVersion.WithKind("DiscoveryEngineAssistant")

// DiscoveryEngineAssistant is the Schema for the DiscoveryEngineAssistant API
// Public API documentation: https://cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1beta/projects.locations.collections.engines.assistants
// Protobuf definition: https://github.com/googleapis/googleapis/blob/master/google/cloud/discoveryengine/v1beta/assistant.proto

// DiscoveryEngineAssistantSpec defines the desired state of DiscoveryEngineAssistant
// +kcc:spec:proto=google.cloud.discoveryengine.v1beta.Assistant
type DiscoveryEngineAssistantSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// Immutable. The Engine this assistant belongs to.
	// +required
	EngineRef *DiscoveryEngineEngineRef `json:"engineRef"`

	// The DiscoveryEngineAssistant name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The assistant display name. It must be a UTF-8 encoded string with a length limit of 128 characters.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description for additional information. Expected to be shown on the configuration UI, not to the users of the assistant.
	Description *string `json:"description,omitempty"`

	// Optional. Configuration for the generation of the assistant response.
	GenerationConfig *Assistant_GenerationConfig `json:"generationConfig,omitempty"`

	// Optional. The type of web grounding to use.
	// +kubebuilder:validation:Enum=WEB_GROUNDING_TYPE_UNSPECIFIED;WEB_GROUNDING_TYPE_DISABLED;WEB_GROUNDING_TYPE_GOOGLE_SEARCH;WEB_GROUNDING_TYPE_ENTERPRISE_WEB_SEARCH
	WebGroundingType *string `json:"webGroundingType,omitempty"`

	// Optional. This field controls the default web grounding toggle for end users.
	DefaultWebGroundingToggleOff *bool `json:"defaultWebGroundingToggleOff,omitempty"`

	// Optional. The enabled tools on this assistant.
	EnabledTools map[string]Assistant_ToolList `json:"enabledTools,omitempty"`

	// Optional. Customer policy for the assistant.
	CustomerPolicy *Assistant_CustomerPolicy `json:"customerPolicy,omitempty"`
}

// DiscoveryEngineAssistantStatus defines the config connector machine state of DiscoveryEngineAssistant
type DiscoveryEngineAssistantStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +kubebuilder:default=0
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineAssistant resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineAssistantObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineAssistantObservedState is the state of the DiscoveryEngineAssistant resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1beta.Assistant
type DiscoveryEngineAssistantObservedState struct {
	// Output only. Timestamp the assistant was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryengineassistant;gcpdiscoveryengineassistants
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineAssistant is the Schema for the DiscoveryEngineAssistant API
// +k8s:openapi-gen=true
type DiscoveryEngineAssistant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineAssistantSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineAssistantStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineAssistantList contains a list of DiscoveryEngineAssistant
type DiscoveryEngineAssistantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineAssistant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineAssistant{}, &DiscoveryEngineAssistantList{})
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Assistant.CustomerPolicy.ModelArmorConfig
type Assistant_CustomerPolicy_ModelArmorConfig struct {
	// Optional. The resource name of the Model Armor template for sanitizing
	//  user prompts. Format:
	//  `projects/{project}/locations/{location}/templates/{template_id}`
	//
	//  If not specified, no sanitization will be applied to the user prompt.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Assistant.CustomerPolicy.ModelArmorConfig.user_prompt_template
	UserPromptTemplateRef *modelarmorv1alpha1.ModelArmorTemplateRef `json:"userPromptTemplateRef,omitempty"`

	// Optional. The resource name of the Model Armor template for sanitizing
	//  assistant responses. Format:
	//  `projects/{project}/locations/{location}/templates/{template_id}`
	//
	//  If not specified, no sanitization will be applied to the assistant
	//  response.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Assistant.CustomerPolicy.ModelArmorConfig.response_template
	ResponseTemplateRef *modelarmorv1alpha1.ModelArmorTemplateRef `json:"responseTemplateRef,omitempty"`

	// Optional. Defines the failure mode for Model Armor sanitization.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Assistant.CustomerPolicy.ModelArmorConfig.failure_mode
	FailureMode *string `json:"failureMode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Assistant.GenerationConfig
type Assistant_GenerationConfig struct {
	// Optional. The default model to use for assistant.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Assistant.GenerationConfig.default_model_id
	DefaultModelID *string `json:"defaultModelID,omitempty"`

	// Optional. The list of models that are allowed to be used for assistant.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Assistant.GenerationConfig.allowed_model_ids
	AllowedModelIDs []string `json:"allowedModelIDs,omitempty"`

	// System instruction, also known as the prompt preamble for LLM calls.
	//  See also
	//  https://cloud.google.com/vertex-ai/generative-ai/docs/learn/prompts/system-instructions
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Assistant.GenerationConfig.system_instruction
	SystemInstruction *Assistant_GenerationConfig_SystemInstruction `json:"systemInstruction,omitempty"`

	// The default language to use for the generation of the assistant
	//  response.
	//  Use an ISO 639-1 language code such as `en`.
	//  If not specified, the language will be automatically detected.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Assistant.GenerationConfig.default_language
	DefaultLanguage *string `json:"defaultLanguage,omitempty"`
}
