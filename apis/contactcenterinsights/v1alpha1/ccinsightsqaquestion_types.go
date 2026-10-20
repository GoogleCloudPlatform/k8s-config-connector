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

// API sources for CCInsightsQaQuestion, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/contactcenterinsights/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/contact-center/insights/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/contact-center/insights/docs/reference/rest/v1/projects.locations.qaScorecards.revisions.qaQuestions

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CCInsightsQaQuestionGVK = GroupVersion.WithKind("CCInsightsQaQuestion")

// CCInsightsQaQuestionSpec defines the desired state of CCInsightsQaQuestion
// +kcc:spec:proto=google.cloud.contactcenterinsights.v1.QaQuestion
// +kcc:required-from-proto
type CCInsightsQaQuestionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CCInsightsQaQuestion name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Short, descriptive string, used in the UI where it's not practical
	//  to display the full question body. E.g., "Greeting".
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.abbreviation
	Abbreviation *string `json:"abbreviation,omitempty"`

	// Question text. E.g., "Did the agent greet the customer?"
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.question_body
	QuestionBody *string `json:"questionBody,omitempty"`

	// Instructions describing how to determine the answer.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.answer_instructions
	AnswerInstructions *string `json:"answerInstructions,omitempty"`

	// A list of valid answers to the question, which the LLM must choose from.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.answer_choices
	AnswerChoices []QaQuestion_AnswerChoice `json:"answerChoices,omitempty"`

	// User-defined list of arbitrary tags for the question. Used for
	//  grouping/organization and for weighting the score of each question.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.tags
	Tags []string `json:"tags,omitempty"`

	// Defines the order of the question within its parent scorecard revision.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.order
	Order *int32 `json:"order,omitempty"`

	// Metadata about the tuning operation for the question.This field will only
	//  be populated if and only if the question is part of a scorecard revision
	//  that has been tuned.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.tuning_metadata
	TuningMetadata *QaQuestion_TuningMetadata `json:"tuningMetadata,omitempty"`
}

// CCInsightsQaQuestionStatus defines the config connector machine state of CCInsightsQaQuestion
type CCInsightsQaQuestionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CCInsightsQaQuestion resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CCInsightsQaQuestionObservedState `json:"observedState,omitempty"`
}

// CCInsightsQaQuestionObservedState is the state of the CCInsightsQaQuestion resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.QaQuestion
type CCInsightsQaQuestionObservedState struct {
	// Output only. The time at which this question was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the question was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Metrics of the underlying tuned LLM over a holdout/test set while fine
	//  tuning the underlying LLM for the given question. This field will only be
	//  populated if and only if the question is part of a scorecard revision that
	//  has been tuned.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaQuestion.metrics
	Metrics *QaQuestion_MetricsObservedState `json:"metrics,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpccinsightsqaquestion;gcpccinsightsqaquestions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CCInsightsQaQuestion is the Schema for the CCInsightsQaQuestion API
// +k8s:openapi-gen=true
type CCInsightsQaQuestion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CCInsightsQaQuestionSpec   `json:"spec,omitempty"`
	Status CCInsightsQaQuestionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CCInsightsQaQuestionList contains a list of CCInsightsQaQuestion
type CCInsightsQaQuestionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CCInsightsQaQuestion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CCInsightsQaQuestion{}, &CCInsightsQaQuestionList{})
}
