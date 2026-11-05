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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.cloud.dialogflow.cx.v3.SpeechToTextSettings
type AgentSpeechToTextSettings struct {
	// Whether to use speech adaptation for speech recognition.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SpeechToTextSettings.enable_speech_adaptation
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty"`
}

// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Agent
type DialogflowCXAgentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The name of the location this agent is located in.
	//
	// ~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
	//  This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	//  Another options is to use global location so you don't need to manually configure location settings.
	Location string `json:"location"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.avatar_uri
	AvatarUri *string `json:"avatarUri,omitempty"`

	// Immutable. The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be updated after creation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.default_language_code
	// +required
	DefaultLanguageCode string `json:"defaultLanguageCode"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.description
	Description *string `json:"description,omitempty"`

	// The human-readable name of the agent, unique within the location.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.display_name
	// +required
	DisplayName string `json:"displayName"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.enable_spell_correction
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty"`

	// Determines whether this agent should log conversation queries.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.enable_stackdriver_logging
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.security_settings
	SecuritySettings *string `json:"securitySettings,omitempty"`

	// Settings related to speech recognition.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.speech_to_text_settings
	SpeechToTextSettings *AgentSpeechToTextSettings `json:"speechToTextSettings,omitempty"`

	// The list of all languages supported by this agent (except for the default_language_code).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.supported_language_codes
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.time_zone
	// +required
	TimeZone string `json:"timeZone"`
}

type DialogflowCXAgentStatus struct {
	/* Conditions represent the latest available observation of the resource's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowCXAgent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// The unique identifier of the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.name
	Name *string `json:"name,omitempty"`

	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.start_flow
	StartFlow *string `json:"startFlow,omitempty"`
}

// DialogflowCXAgentObservedState is the state of the DialogflowCXAgent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Agent
type DialogflowCXAgentObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxagent;gcpdialogflowcxagents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCXAgent is the Schema for the DialogflowCXAgent API
// +k8s:openapi-gen=true
type DialogflowCXAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCXAgentSpec   `json:"spec,omitempty"`
	Status DialogflowCXAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCXAgentList contains a list of DialogflowCXAgent
type DialogflowCXAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXAgent{}, &DialogflowCXAgentList{})
}
