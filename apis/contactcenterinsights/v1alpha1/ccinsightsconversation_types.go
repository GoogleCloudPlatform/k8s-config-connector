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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CCInsightsConversationGVK = GroupVersion.WithKind("CCInsightsConversation")

// CCInsightsConversationSpec defines the desired state of CCInsightsConversation
// +kcc:spec:proto=google.cloud.contactcenterinsights.v1.Conversation
type CCInsightsConversationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The CCInsightsConversation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Call-specific metadata.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.call_metadata
	// +kubebuilder:validation:Optional
	CallMetadata *Conversation_CallMetadata `json:"callMetadata,omitempty"`

	// The time at which this conversation should expire. After this time, the
	//  conversation data and any associated analyses will be deleted.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.expire_time
	// +kubebuilder:validation:Optional
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The TTL for this resource. If specified, then this TTL will
	//  be used to calculate the expire time.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.ttl
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty"`

	// The source of the audio and transcription for the conversation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.data_source
	// +kubebuilder:validation:Optional
	DataSource *ConversationDataSource `json:"dataSource,omitempty"`

	// The time at which the conversation started.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.start_time
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty"`

	// A user-specified language code for the conversation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.language_code
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty"`

	// An opaque, user-specified string representing the human agent who handled
	//  the conversation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.agent_id
	// +kubebuilder:validation:Optional
	AgentID *string `json:"agentID,omitempty"`

	// A map for the user to specify any custom fields. A maximum of 100 labels
	//  per conversation is allowed, with a maximum of 256 characters per entry.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.labels
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Conversation metadata related to quality management.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.quality_metadata
	// +kubebuilder:validation:Optional
	QualityMetadata *Conversation_QualityMetadata `json:"qualityMetadata,omitempty"`

	// Input only. JSON metadata encoded as a string.
	//  This field is primarily used by Insights integrations with various telephony
	//  systems and must be in one of Insight's supported formats.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.metadata_json
	// +kubebuilder:validation:Optional
	MetadataJson *string `json:"metadataJson,omitempty"`

	// Immutable. The conversation medium, if unspecified will default to
	//  PHONE_CALL.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.medium
	// +kubebuilder:validation:Enum=MEDIUM_UNSPECIFIED;PHONE_CALL;CHAT
	// +kubebuilder:validation:Optional
	Medium *string `json:"medium,omitempty"`

	// Obfuscated user ID which the customer sent to us.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.obfuscated_user_id
	// +kubebuilder:validation:Optional
	ObfuscatedUserID *string `json:"obfuscatedUserID,omitempty"`
}

// CCInsightsConversationStatus defines the config connector machine state of CCInsightsConversation
type CCInsightsConversationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CCInsightsConversation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CCInsightsConversationObservedState `json:"observedState,omitempty"`
}

// CCInsightsConversationObservedState is the state of the CCInsightsConversation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.Conversation
type CCInsightsConversationObservedState struct {
	// Output only. The time at which the conversation was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the conversation was updated.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The conversation transcript.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.transcript
	// +kubebuilder:validation:Optional
	Transcript *Conversation_Transcript `json:"transcript,omitempty"`

	// Output only. The duration of the conversation.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty"`

	// Output only. The number of turns in the conversation.
	// +kubebuilder:validation:Optional
	TurnCount *int32 `json:"turnCount,omitempty"`

	// Output only. The conversation's latest analysis, if one exists.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.latest_analysis
	// +kubebuilder:validation:Optional
	LatestAnalysis *Analysis `json:"latestAnalysis,omitempty"`

	// Output only. Latest summary of the conversation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.latest_summary
	// +kubebuilder:validation:Optional
	LatestSummary *ConversationSummarizationSuggestionData `json:"latestSummary,omitempty"`

	// Output only. The annotations that were generated during the customer and
	//  agent interaction.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Conversation.runtime_annotations
	// +kubebuilder:validation:Optional
	RuntimeAnnotations []RuntimeAnnotation `json:"runtimeAnnotations,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpccinsightsconversation;gcpccinsightsconversations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CCInsightsConversation is the Schema for the CCInsightsConversation API
// +k8s:openapi-gen=true
type CCInsightsConversation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CCInsightsConversationSpec   `json:"spec,omitempty"`
	Status CCInsightsConversationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CCInsightsConversationList contains a list of CCInsightsConversation
type CCInsightsConversationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CCInsightsConversation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CCInsightsConversation{}, &CCInsightsConversationList{})
}
