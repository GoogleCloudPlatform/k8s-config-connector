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

var DialogflowConversationDatasetGVK = GroupVersion.WithKind("DialogflowConversationDataset")

// DialogflowConversationDatasetSpec defines the desired state of DialogflowConversationDataset
// +kcc:spec:proto=google.cloud.dialogflow.v2.ConversationDataset
type DialogflowConversationDatasetSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The DialogflowConversationDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the dataset. Maximum of 64 bytes.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.display_name
	DisplayName *string `json:"displayName"`

	// Optional. The description of the dataset. Maximum of 10000 bytes.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.description
	Description *string `json:"description,omitempty"`
}

// DialogflowConversationDatasetStatus defines the config connector machine state of DialogflowConversationDataset
type DialogflowConversationDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowConversationDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowConversationDatasetObservedState `json:"observedState,omitempty"`
}

// DialogflowConversationDatasetObservedState is the state of the DialogflowConversationDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.ConversationDataset
type DialogflowConversationDatasetObservedState struct {
	// Output only. Creation time of this dataset.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Input configurations set during conversation data import.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.input_config
	InputConfig *InputConfig `json:"inputConfig,omitempty"`

	// Output only. Metadata set during conversation data import.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.conversation_info
	ConversationInfo *ConversationInfo `json:"conversationInfo,omitempty"`

	// Output only. The number of conversations this conversation dataset
	//  contains.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.conversation_count
	ConversationCount *int64 `json:"conversationCount,omitempty"`

	// Output only. A read only boolean field reflecting Zone Isolation status of
	//  the dataset.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. A read only boolean field reflecting Zone Separation status of
	//  the dataset.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationDataset.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowconversationdataset;gcpdialogflowconversationdatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowConversationDataset is the Schema for the DialogflowConversationDataset API
// +k8s:openapi-gen=true
type DialogflowConversationDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowConversationDatasetSpec   `json:"spec,omitempty"`
	Status DialogflowConversationDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowConversationDatasetList contains a list of DialogflowConversationDataset
type DialogflowConversationDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowConversationDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowConversationDataset{}, &DialogflowConversationDatasetList{})
}
