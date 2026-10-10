// Copyright 2025 Google LLC
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

// API sources for CCInsightsFeedbackLabel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/contactcenterinsights/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/contact-center/insights/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/contact-center/insights/docs/reference/rest/v1/projects.locations.conversations.feedbackLabels

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CCInsightsFeedbackLabelGVK = GroupVersion.WithKind("CCInsightsFeedbackLabel")

// CCInsightsFeedbackLabelSpec defines the desired state of CCInsightsFeedbackLabel
// +kcc:spec:proto=google.cloud.contactcenterinsights.v1.FeedbackLabel
// +kcc:required-from-proto
type CCInsightsFeedbackLabelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/conversations/{conversation} this resource belongs to.
	// +kcc:guess
	// ConversationRef *CCInsightsConversationRef `json:"conversationRef,omitempty"`

	// The CCInsightsFeedbackLabel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// String label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.label
	Label *string `json:"label,omitempty"`

	// QaAnswer label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.qa_answer_label
	QaAnswerLabel *QaAnswer_AnswerValue `json:"qaAnswerLabel,omitempty"`

	// Resource name of the resource to be labeled.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.labeled_resource
	LabeledResource *string `json:"labeledResource,omitempty"`
}

// CCInsightsFeedbackLabelStatus defines the config connector machine state of CCInsightsFeedbackLabel
type CCInsightsFeedbackLabelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CCInsightsFeedbackLabel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CCInsightsFeedbackLabelObservedState `json:"observedState,omitempty"`
}

// CCInsightsFeedbackLabelObservedState is the state of the CCInsightsFeedbackLabel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.FeedbackLabel
type CCInsightsFeedbackLabelObservedState struct {
	// QaAnswer label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.qa_answer_label
	QaAnswerLabel *QaAnswer_AnswerValueObservedState `json:"qaAnswerLabel,omitempty"`

	// Output only. Create time of the label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time of the label.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.FeedbackLabel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpccinsightsfeedbacklabel;gcpccinsightsfeedbacklabels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CCInsightsFeedbackLabel is the Schema for the CCInsightsFeedbackLabel API
// +k8s:openapi-gen=true
type CCInsightsFeedbackLabel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CCInsightsFeedbackLabelSpec   `json:"spec,omitempty"`
	Status CCInsightsFeedbackLabelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CCInsightsFeedbackLabelList contains a list of CCInsightsFeedbackLabel
type CCInsightsFeedbackLabelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CCInsightsFeedbackLabel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CCInsightsFeedbackLabel{}, &CCInsightsFeedbackLabelList{})
}
