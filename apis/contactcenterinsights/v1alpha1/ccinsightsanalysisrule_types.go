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

var CCInsightsAnalysisRuleGVK = GroupVersion.WithKind("CCInsightsAnalysisRule")

// CCInsightsAnalysisRuleSpec defines the desired state of CCInsightsAnalysisRule
// +kcc:spec:proto=google.cloud.contactcenterinsights.v1.AnalysisRule
type CCInsightsAnalysisRuleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	Location *string `json:"location"`

	// The CCInsightsAnalysisRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Display Name of the analysis rule.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// Filter for the conversations that should apply this analysis rule. An empty filter means this analysis rule applies to all conversations.
	// +optional
	ConversationFilter *string `json:"conversationFilter,omitempty"`

	// Selector of annotators to run and the phrase matchers to use for conversations that matches the conversation_filter. If not specified, NO annotators will be run.
	// +optional
	AnnotatorSelector *CCInsightsAnalysisRuleAnnotatorSelector `json:"annotatorSelector,omitempty"`

	// Percentage of conversations that we should apply this analysis setting automatically, between [0, 1]. For example, 0.1 means 10%. Conversations are sampled in a determenestic way.
	// +optional
	AnalysisPercentage *float64 `json:"analysisPercentage,omitempty"`

	// If true, apply this rule to conversations. Otherwise, this rule is inactive and saved as a draft.
	// +optional
	Active *bool `json:"active,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotatorSelector
type CCInsightsAnalysisRuleAnnotatorSelector struct {
	// Whether to run the interruption annotator.
	// +optional
	RunInterruptionAnnotator *bool `json:"runInterruptionAnnotator,omitempty"`

	// Whether to run the silence annotator.
	// +optional
	RunSilenceAnnotator *bool `json:"runSilenceAnnotator,omitempty"`

	// Whether to run the active phrase matcher annotator(s).
	// +optional
	RunPhraseMatcherAnnotator *bool `json:"runPhraseMatcherAnnotator,omitempty"`

	// The list of phrase matchers to run. If not provided, all active phrase matchers will be used.
	// +optional
	PhraseMatchersRefs []CCInsightsPhraseMatcherRef `json:"phraseMatchersRefs,omitempty"`

	// Whether to run the sentiment annotator.
	// +optional
	RunSentimentAnnotator *bool `json:"runSentimentAnnotator,omitempty"`

	// Whether to run the entity annotator.
	// +optional
	RunEntityAnnotator *bool `json:"runEntityAnnotator,omitempty"`

	// Whether to run the intent annotator.
	// +optional
	RunIntentAnnotator *bool `json:"runIntentAnnotator,omitempty"`

	// Whether to run the issue model annotator. A model should have already been deployed for this to take effect.
	// +optional
	RunIssueModelAnnotator *bool `json:"runIssueModelAnnotator,omitempty"`

	// The issue model to run. If not provided, the most recently deployed topic model will be used.
	// +optional
	IssueModelsRefs []CCInsightsIssueModelRef `json:"issueModelsRefs,omitempty"`

	// Whether to run the summarization annotator.
	// +optional
	RunSummarizationAnnotator *bool `json:"runSummarizationAnnotator,omitempty"`

	// Configuration for the summarization annotator.
	// +optional
	SummarizationConfig *CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig `json:"summarizationConfig,omitempty"`

	// Whether to run the QA annotator.
	// +optional
	RunQaAnnotator *bool `json:"runQaAnnotator,omitempty"`

	// Configuration for the QA annotator.
	// +optional
	QaConfig *AnnotatorSelector_QaConfig `json:"qaConfig,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotatorSelector.SummarizationConfig
type CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig struct {
	// Resource name of the Dialogflow conversation profile.
	// +optional
	ConversationProfileRef *DialogflowConversationProfileRef `json:"conversationProfileRef,omitempty"`

	// Default summarization model to be used.
	// +optional
	SummarizationModel *string `json:"summarizationModel,omitempty"`
}

// CCInsightsAnalysisRuleStatus defines the config connector machine state of CCInsightsAnalysisRule
type CCInsightsAnalysisRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CCInsightsAnalysisRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CCInsightsAnalysisRuleObservedState `json:"observedState,omitempty"`
}

// CCInsightsAnalysisRuleObservedState is the state of the CCInsightsAnalysisRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.AnalysisRule
type CCInsightsAnalysisRuleObservedState struct {
	// Output only. The time at which this analysis rule was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which this analysis rule was updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpccinsightsanalysisrule;gcpccinsightsanalysisrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CCInsightsAnalysisRule is the Schema for the CCInsightsAnalysisRule API
// +k8s:openapi-gen=true
type CCInsightsAnalysisRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CCInsightsAnalysisRuleSpec   `json:"spec,omitempty"`
	Status CCInsightsAnalysisRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CCInsightsAnalysisRuleList contains a list of CCInsightsAnalysisRule
type CCInsightsAnalysisRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CCInsightsAnalysisRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CCInsightsAnalysisRule{}, &CCInsightsAnalysisRuleList{})
}
