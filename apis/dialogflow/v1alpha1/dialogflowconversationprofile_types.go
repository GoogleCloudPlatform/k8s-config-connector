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
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowConversationProfileGVK = GroupVersion.WithKind("DialogflowConversationProfile")

// DialogflowConversationProfileSpec defines the desired state of DialogflowConversationProfile
// +kcc:spec:proto=google.cloud.dialogflow.v2.ConversationProfile
type DialogflowConversationProfileSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location,omitempty"`

	// The DialogflowConversationProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Human readable name for this profile. Max length 1024 bytes.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.display_name
	DisplayName *string `json:"displayName"`

	// Configuration for an automated agent to use with this profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.automated_agent_config
	AutomatedAgentConfig *AutomatedAgentConfig `json:"automatedAgentConfig,omitempty"`

	// Configuration for agent assistance to use with this profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.human_agent_assistant_config
	HumanAgentAssistantConfig *HumanAgentAssistantConfig `json:"humanAgentAssistantConfig,omitempty"`

	// Configuration for connecting to a live agent.
	// Currently, this feature is not general available, please contact Google to get access.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.human_agent_handoff_config
	HumanAgentHandoffConfig *HumanAgentHandoffConfig `json:"humanAgentHandoffConfig,omitempty"`

	// Configuration for publishing conversation lifecycle events.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.notification_config
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`

	// Configuration for logging conversation lifecycle events.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Configuration for publishing new message events. Event will be sent in format of ConversationEvent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.new_message_event_notification_config
	NewMessageEventNotificationConfig *NotificationConfig `json:"newMessageEventNotificationConfig,omitempty"`

	// Optional. Configuration for publishing transcription intermediate results.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.new_recognition_result_notification_config
	NewRecognitionResultNotificationConfig *NotificationConfig `json:"newRecognitionResultNotificationConfig,omitempty"`

	// Settings for speech transcription.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.stt_config
	SttConfig *SpeechToTextConfig `json:"sttConfig,omitempty"`

	// Language code for the conversation profile. If not specified, the language is en-US.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The time zone of this conversational profile from the time zone database, e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Name of the CX SecuritySettings reference for the agent.
	// Format: `projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationProfile.security_settings
	SecuritySettingsRef *DialogflowSecuritySettingsRef `json:"securitySettingsRef,omitempty"`

	// Configuration for Text-to-Speech synthesization.
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
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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

// +kcc:proto=google.cloud.dialogflow.v2.NotificationConfig
type NotificationConfig struct {
	// Name of the Pub/Sub topic to publish conversation events like
	// CONVERSATION_STARTED as serialized ConversationEvent protos.
	// Format: `projects/<Project ID>/locations/<Location ID>/topics/<Topic ID>` or `projects/<Project ID>/topics/<Topic ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.NotificationConfig.topic
	TopicRef *pubsubv1beta1.PubSubTopicRef `json:"topicRef,omitempty"`

	// Format of message.
	// +kcc:proto:field=google.cloud.dialogflow.v2.NotificationConfig.message_format
	MessageFormat *string `json:"messageFormat,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.HumanAgentAssistantConfig.SuggestionQueryConfig.KnowledgeBaseQuerySource
type HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource struct {
	// Required. Knowledge bases to query. Format:
	// `projects/<Project ID>/locations/<Location ID>/knowledgeBases/<Knowledge Base ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.HumanAgentAssistantConfig.SuggestionQueryConfig.KnowledgeBaseQuerySource.knowledge_bases
	KnowledgeBaseRefs []DialogflowKnowledgeBaseRef `json:"knowledgeBaseRefs,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.HumanAgentAssistantConfig.SuggestionConfig
type HumanAgentAssistantConfig_SuggestionConfig struct {
	// Configuration of different suggestion features. One feature can have only
	//  one config.
	// +kcc:proto:field=google.cloud.dialogflow.v2.HumanAgentAssistantConfig.SuggestionConfig.feature_configs
	FeatureConfigs []HumanAgentAssistantConfig_SuggestionFeatureConfig `json:"featureConfigs,omitempty"`

	// If `group_suggestion_responses` is false, and there are multiple
	//  `feature_configs` in `event based suggestion` or
	//  StreamingAnalyzeContent, we will try to deliver suggestions to customers
	//  as soon as we get new suggestion. Different type of suggestions based on
	//  the same context will be in  separate Pub/Sub event or
	//  `StreamingAnalyzeContentResponse`.
	//
	//  If `group_suggestion_responses` set to true. All the suggestions to the
	//  same participant based on the same context will be grouped into a single
	//  Pub/Sub event or StreamingAnalyzeContentResponse.
	// +kcc:proto:field=google.cloud.dialogflow.v2.HumanAgentAssistantConfig.SuggestionConfig.group_suggestion_responses
	GroupSuggestionResponses *bool `json:"groupSuggestionResponses,omitempty"`

	// Optional. List of various generator resource names used in the
	//  conversation profile.
	// +kcc:proto:field=google.cloud.dialogflow.v2.HumanAgentAssistantConfig.SuggestionConfig.generators
	GeneratorRefs []DialogflowGeneratorRef `json:"generatorRefs,omitempty"`

	// Optional. When disable_high_latency_features_sync_delivery is true and
	//  using the AnalyzeContent API, we will not deliver the responses from high
	//  latency features in the API response. The
	//  human_agent_assistant_config.notification_config must be configured and
	//  enable_event_based_suggestion must be set to true to receive the
	//  responses from high latency features in Pub/Sub. High latency feature(s):
	//  KNOWLEDGE_ASSIST
	// +kcc:proto:field=google.cloud.dialogflow.v2.HumanAgentAssistantConfig.SuggestionConfig.disable_high_latency_features_sync_delivery
	DisableHighLatencyFeaturesSyncDelivery *bool `json:"disableHighLatencyFeaturesSyncDelivery,omitempty"`
}
