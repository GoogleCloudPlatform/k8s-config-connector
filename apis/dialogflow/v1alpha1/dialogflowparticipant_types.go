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

// API sources for DialogflowParticipant, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dialogflow/v2/participant.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v2/projects.conversations.participants

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowParticipantGVK = GroupVersion.WithKind("DialogflowParticipant")

// DialogflowParticipantSpec defines the desired state of DialogflowParticipant
// +kcc:spec:proto=google.cloud.dialogflow.v2.Participant
// +kcc:required-from-proto
type DialogflowParticipantSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DialogflowParticipant name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. The role this participant plays in the conversation. This field
	//  must be set during participant creation and is then immutable.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.role
	Role *string `json:"role,omitempty"`

	// Optional. Label applied to streams representing this participant in SIPREC
	//  XML metadata and SDP. This is used to assign transcriptions from that
	//  media stream to this participant. This field can be updated.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.sip_recording_media_label
	SipRecordingMediaLabel *string `json:"sipRecordingMediaLabel,omitempty"`

	// Optional. Obfuscated user id that should be associated with the created
	//  participant.
	//
	//  You can specify a user id as follows:
	//
	//  1. If you set this field in
	//     [CreateParticipantRequest][google.cloud.dialogflow.v2.CreateParticipantRequest.participant]
	//     or
	//     [UpdateParticipantRequest][google.cloud.dialogflow.v2.UpdateParticipantRequest.participant],
	//     Dialogflow adds the obfuscated user id with the participant.
	//
	//  2. If you set this field in
	//     [AnalyzeContent][google.cloud.dialogflow.v2.AnalyzeContentRequest.participant]
	//     or
	//     [StreamingAnalyzeContent][google.cloud.dialogflow.v2.StreamingAnalyzeContentRequest.participant],
	//     Dialogflow will update
	//     [Participant.obfuscated_external_user_id][google.cloud.dialogflow.v2.Participant.obfuscated_external_user_id].
	//
	//  Dialogflow returns an error if you try to add a user id for a
	//  non-[END_USER][google.cloud.dialogflow.v2.Participant.Role.END_USER]
	//  participant.
	//
	//  Dialogflow uses this user id for billing and measurement purposes. For
	//  example, Dialogflow determines whether a user in one conversation returned
	//  in a later conversation.
	//
	//  Note:
	//
	//  * Please never pass raw user ids to Dialogflow. Always obfuscate your user
	//    id first.
	//  * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a
	//    hash function like SHA-512.
	//  * The length of the user id must be <= 256 characters.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.obfuscated_external_user_id
	ObfuscatedExternalUserID *string `json:"obfuscatedExternalUserID,omitempty"`

	// Optional. Key-value filters on the metadata of documents returned by
	//  article suggestion. If specified, article suggestion only returns suggested
	//  documents that match all filters in their
	//  [Document.metadata][google.cloud.dialogflow.v2.Document.metadata]. Multiple
	//  values for a metadata key should be concatenated by comma. For example,
	//  filters to match all documents that have 'US' or 'CA' in their market
	//  metadata values and 'agent' in their user metadata values will be
	//  ```
	//  documents_metadata_filters {
	//    key: "market"
	//    value: "US,CA"
	//  }
	//  documents_metadata_filters {
	//    key: "user"
	//    value: "agent"
	//  }
	//  ```
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.documents_metadata_filters
	DocumentsMetadataFilters map[string]string `json:"documentsMetadataFilters,omitempty"`
}

// DialogflowParticipantStatus defines the config connector machine state of DialogflowParticipant
type DialogflowParticipantStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowParticipant resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DialogflowParticipantObservedState `json:"observedState,omitempty"`
}

// DialogflowParticipantObservedState is the state of the DialogflowParticipant resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.Participant
// type DialogflowParticipantObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowparticipant;gcpdialogflowparticipants
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowParticipant is the Schema for the DialogflowParticipant API
// +k8s:openapi-gen=true
type DialogflowParticipant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowParticipantSpec   `json:"spec,omitempty"`
	Status DialogflowParticipantStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowParticipantList contains a list of DialogflowParticipant
type DialogflowParticipantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowParticipant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowParticipant{}, &DialogflowParticipantList{})
}
