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

// API sources for DialogflowConversationProfile, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dialogflow/v2/conversation_profile.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v2/projects.conversationProfiles

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowConversationProfileGVK = GroupVersion.WithKind("DialogflowConversationProfile")

// DialogflowConversationProfileSpec defines the desired state of DialogflowConversationProfile
// +kcc:spec:proto=google.cloud.dialogflow.v2.ConversationProfile
// +kcc:required-from-proto
type DialogflowConversationProfileSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DialogflowConversationProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Human readable name for this profile. Max length 1024 bytes.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration for an automated agent to use with this profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.automated_agent_config
	AutomatedAgentConfig *AutomatedAgentConfig `json:"automatedAgentConfig,omitempty"`

	// Configuration for agent assistance to use with this profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.human_agent_assistant_config
	HumanAgentAssistantConfig *HumanAgentAssistantConfig `json:"humanAgentAssistantConfig,omitempty"`

	// Configuration for connecting to a live agent.
	//
	//  Currently, this feature is not general available, please contact Google
	//  to get access.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.human_agent_handoff_config
	HumanAgentHandoffConfig *HumanAgentHandoffConfig `json:"humanAgentHandoffConfig,omitempty"`

	// Configuration for publishing conversation lifecycle events.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.notification_config
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`

	// Configuration for logging conversation lifecycle events.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Configuration for publishing new message events. Event will be sent in
	//  format of [ConversationEvent][google.cloud.dialogflow.v2.ConversationEvent]
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.new_message_event_notification_config
	NewMessageEventNotificationConfig *NotificationConfig `json:"newMessageEventNotificationConfig,omitempty"`

	// Optional. Configuration for publishing transcription intermediate results.
	//  Event will be sent in format of
	//  [ConversationEvent][google.cloud.dialogflow.v2.ConversationEvent]. If
	//  configured, the following information will be populated as
	//  [ConversationEvent][google.cloud.dialogflow.v2.ConversationEvent] Pub/Sub
	//  message attributes:
	//  - "participant_id"
	//  - "participant_role"
	//  - "message_id"
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.new_recognition_result_notification_config
	NewRecognitionResultNotificationConfig *NotificationConfig `json:"newRecognitionResultNotificationConfig,omitempty"`

	// Settings for speech transcription.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.stt_config
	SttConfig *SpeechToTextConfig `json:"sttConfig,omitempty"`

	// Language code for the conversation profile. If not specified, the language
	//  is en-US. Language at ConversationProfile should be set for all non en-US
	//  languages.
	//  This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt)
	//  language tag. Example: "en-US".
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The time zone of this conversational profile from the
	//  [time zone database](https://www.iana.org/time-zones), e.g.,
	//  America/New_York, Europe/Paris. Defaults to America/New_York.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Name of the CX SecuritySettings reference for the agent.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/securitySettings/<Security Settings ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.security_settings
	SecuritySettings *string `json:"securitySettings,omitempty"`

	// Configuration for Text-to-Speech synthesization.
	//
	//  Used by Phone Gateway to specify synthesization options. If agent defines
	//  synthesization options as well, agent settings overrides the option here.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.tts_config
	TtsConfig *SynthesizeSpeechConfig `json:"ttsConfig,omitempty"`
}

// DialogflowConversationProfileStatus defines the config connector machine state of DialogflowConversationProfile
type DialogflowConversationProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowConversationProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowConversationProfileObservedState `json:"observedState,omitempty"`
}

// DialogflowConversationProfileObservedState is the state of the DialogflowConversationProfile resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.ConversationProfile
type DialogflowConversationProfileObservedState struct {
	// Output only. Create time of the conversation profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time of the conversation profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowconversationprofile;gcpdialogflowconversationprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowConversationProfile is the Schema for the DialogflowConversationProfile API
// +k8s:openapi-gen=true
type DialogflowConversationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowConversationProfileSpec   `json:"spec,omitempty"`
	Status DialogflowConversationProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowConversationProfileList contains a list of DialogflowConversationProfile
type DialogflowConversationProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowConversationProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowConversationProfile{}, &DialogflowConversationProfileList{})
}
