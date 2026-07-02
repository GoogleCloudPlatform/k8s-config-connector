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
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	Location string `json:"location"`

	// Display Name of the analysis rule.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.display_name
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// Filter for the conversations that should apply this analysis
	//  rule. An empty filter means this analysis rule applies to all
	//  conversations.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.conversation_filter
	// +optional
	ConversationFilter *string `json:"conversationFilter,omitempty"`

	// Selector of annotators to run and the phrase matchers to use for
	//  conversations that matches the conversation_filter. If not specified, NO
	//  annotators will be run.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.annotator_selector
	// +optional
	AnnotatorSelector *AnnotatorSelector `json:"annotatorSelector,omitempty"`

	// Percentage of conversations that we should apply this analysis setting
	//  automatically, between [0, 1]. For example, 0.1 means 10%. Conversations
	//  are sampled in a deterministic way.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.analysis_percentage
	// +optional
	AnalysisPercentage *float64 `json:"analysisPercentage,omitempty"`

	// If true, apply this rule to conversations. Otherwise, this rule is
	//  inactive and saved as a draft.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnalysisRule.active
	// +optional
	Active *bool `json:"active,omitempty"`

	// The CCInsightsAnalysisRule name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.AnnotatorSelector
type AnnotatorSelector struct {
	// Whether to run the interruption annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_interruption_annotator
	// +optional
	RunInterruptionAnnotator *bool `json:"runInterruptionAnnotator,omitempty"`

	// Whether to run the silence annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_silence_annotator
	// +optional
	RunSilenceAnnotator *bool `json:"runSilenceAnnotator,omitempty"`

	// Whether to run the active phrase matcher annotator(s).
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_phrase_matcher_annotator
	// +optional
	RunPhraseMatcherAnnotator *bool `json:"runPhraseMatcherAnnotator,omitempty"`

	// The list of phrase matchers to run. If not provided, all active phrase
	//  matchers will be used. If inactive phrase matchers are provided, they will
	//  not be used. Phrase matchers will be run only if
	//  run_phrase_matcher_annotator is set to true. Format:
	//  projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.phrase_matchers
	// +optional
	PhraseMatchers []CCInsightsPhraseMatcherRef `json:"phraseMatchers,omitempty"`

	// Whether to run the sentiment annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_sentiment_annotator
	// +optional
	RunSentimentAnnotator *bool `json:"runSentimentAnnotator,omitempty"`

	// Whether to run the entity annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_entity_annotator
	// +optional
	RunEntityAnnotator *bool `json:"runEntityAnnotator,omitempty"`

	// Whether to run the intent annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_intent_annotator
	// +optional
	RunIntentAnnotator *bool `json:"runIntentAnnotator,omitempty"`

	// Whether to run the issue model annotator. A model should have already been
	//  deployed for this to take effect.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_issue_model_annotator
	// +optional
	RunIssueModelAnnotator *bool `json:"runIssueModelAnnotator,omitempty"`

	// The issue model to run. If not provided, the most recently deployed topic
	//  model will be used. The provided issue model will only be used for
	//  inference if the issue model is deployed and if run_issue_model_annotator
	//  is set to true. If more than one issue model is provided, only the first
	//  provided issue model will be used for inference.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.issue_models
	// +optional
	IssueModels []CCInsightsIssueModelRef `json:"issueModels,omitempty"`

	// Whether to run the summarization annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_summarization_annotator
	// +optional
	RunSummarizationAnnotator *bool `json:"runSummarizationAnnotator,omitempty"`

	// Configuration for the summarization annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.summarization_config
	// +optional
	SummarizationConfig *AnnotatorSelector_SummarizationConfig `json:"summarizationConfig,omitempty"`

	// Whether to run the QA annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.run_qa_annotator
	// +optional
	RunQaAnnotator *bool `json:"runQaAnnotator,omitempty"`

	// Configuration for the QA annotator.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.AnnotatorSelector.qa_config
	// +optional
	QaConfig *AnnotatorSelector_QaConfig `json:"qaConfig,omitempty"`
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
