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

// API sources for DialogflowConversationModel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dialogflow/v2/conversation_model.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v2/projects.locations.conversationModels

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowConversationModelGVK = GroupVersion.WithKind("DialogflowConversationModel")

// DialogflowConversationModelSpec defines the desired state of DialogflowConversationModel
// +kcc:spec:proto=google.cloud.dialogflow.v2.ConversationModel
// +kcc:required-from-proto
type DialogflowConversationModelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DialogflowConversationModel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the model. At most 64 bytes long.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Datasets used to create model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.datasets
	// +required
	Datasets []InputDataset `json:"datasets,omitempty"`

	// Language code for the conversation model. If not specified, the language
	//  is en-US. Language at ConversationModel should be set for all non en-us
	//  languages.
	//  This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt)
	//  language tag. Example: "en-US".
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Metadata for article suggestion models.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.article_suggestion_model_metadata
	ArticleSuggestionModelMetadata *ArticleSuggestionModelMetadata `json:"articleSuggestionModelMetadata,omitempty"`

	// Metadata for smart reply models.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.smart_reply_model_metadata
	SmartReplyModelMetadata *SmartReplyModelMetadata `json:"smartReplyModelMetadata,omitempty"`
}

// DialogflowConversationModelStatus defines the config connector machine state of DialogflowConversationModel
type DialogflowConversationModelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowConversationModel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowConversationModelObservedState `json:"observedState,omitempty"`
}

// DialogflowConversationModelObservedState is the state of the DialogflowConversationModel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.ConversationModel
type DialogflowConversationModelObservedState struct {
	// Output only. Creation time of this model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. State of the model. A model can only serve prediction requests
	//  after it gets deployed.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.state
	State *string `json:"state,omitempty"`

	// Output only. A read only boolean field reflecting Zone Separation
	//  status of the model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. A read only boolean field reflecting Zone Isolation status
	//  of the model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModel.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowconversationmodel;gcpdialogflowconversationmodels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowConversationModel is the Schema for the DialogflowConversationModel API
// +k8s:openapi-gen=true
type DialogflowConversationModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowConversationModelSpec   `json:"spec,omitempty"`
	Status DialogflowConversationModelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowConversationModelList contains a list of DialogflowConversationModel
type DialogflowConversationModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowConversationModel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowConversationModel{}, &DialogflowConversationModelList{})
}
