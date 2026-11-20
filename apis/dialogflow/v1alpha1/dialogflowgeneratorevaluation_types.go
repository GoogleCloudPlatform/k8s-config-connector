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

// API sources for DialogflowGeneratorEvaluation, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/v2/generator_evaluation.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v2/projects.locations.generators.evaluations

package v1alpha1

import (
	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowGeneratorEvaluationGVK = GroupVersion.WithKind("DialogflowGeneratorEvaluation")

// DialogflowGeneratorEvaluationSpec defines the desired state of DialogflowGeneratorEvaluation
// +kcc:spec:proto=google.cloud.dialogflow.v2.GeneratorEvaluation
// +kcc:required-from-proto
type DialogflowGeneratorEvaluationSpec struct {
	// The DialogflowGenerator that this resource belongs to.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *DialogflowGeneratorRef `json:"parentRef"`

	// The server-assigned evaluation ID of the DialogflowGeneratorEvaluation.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The display name of the generator evaluation. At most 64 bytes
	//  long.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The configuration of the evaluation task.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.generator_evaluation_config
	// +required
	GeneratorEvaluationConfig *GeneratorEvaluationConfig `json:"generatorEvaluationConfig,omitempty"`

	// Required. The initial generator that was used when creating this
	//  evaluation. This is a copy of the generator read from storage when creating
	//  the evaluation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.initial_generator
	// +required
	InitialGenerator *DialogflowGeneratorEvaluation_Generator `json:"initialGenerator,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Generator
type DialogflowGeneratorEvaluation_Generator struct {
	// Optional. Human readable description of the generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.description
	Description *string `json:"description,omitempty"`

	// Input of free from generator to LLM.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.free_form_context
	FreeFormContext *FreeFormContext `json:"freeFormContext,omitempty"`

	// Input of summarization generator to LLM.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.summarization_context
	SummarizationContext *SummarizationContextWithRequired `json:"summarizationContext,omitempty"`

	// Optional. Inference parameters for this generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.inference_parameter
	InferenceParameter *InferenceParameter `json:"inferenceParameter,omitempty"`

	// Optional. The trigger event of the generator. It defines when the generator
	//  is triggered in a conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.trigger_event
	TriggerEvent *string `json:"triggerEvent,omitempty"`

	// Optional. The published Large Language Model name.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.published_model
	PublishedModel *string `json:"publishedModel,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SummarizationContext
type SummarizationContextWithRequired struct {
	// Optional. List of few shot examples.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SummarizationContext.few_shot_examples
	FewShotExamples []FewShotExampleWithRequired `json:"fewShotExamples,omitempty"`

	// Optional. The language of the text. See [Language
	//  Support](https://cloud.google.com/dialogflow/docs/reference/language) for a
	//  list of the currently supported language codes.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SummarizationContext.output_language_code
	OutputLanguageCode *string `json:"outputLanguageCode,omitempty"`

	// Optional. List of sections. Note it contains both predefined section sand
	//  customer defined sections.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SummarizationContext.summarization_sections
	SummarizationSections []SummarizationSection `json:"summarizationSections,omitempty"`

	// Optional. Version of the feature. If not set, default to latest version.
	//  Current candidates are ["1.0"].
	// +kcc:proto:field=google.cloud.dialogflow.v2.SummarizationContext.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.FewShotExample
type FewShotExampleWithRequired struct {
	// Optional. Conversation transcripts.
	// +kcc:proto:field=google.cloud.dialogflow.v2.FewShotExample.conversation_context
	ConversationContext *ConversationContext `json:"conversationContext,omitempty"`

	// Optional. Key is the placeholder field name in input, value is the value of
	//  the placeholder. E.g. instruction contains "@price", and ingested data has
	//  <"price", "10">
	// +kcc:proto:field=google.cloud.dialogflow.v2.FewShotExample.extra_info
	ExtraInfo map[string]string `json:"extraInfo,omitempty"`

	// Required. Example output of the model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.FewShotExample.output
	// +required
	Output *GeneratorSuggestionWithRequired `json:"output,omitempty"`

	// Summarization sections.
	// +kcc:proto:field=google.cloud.dialogflow.v2.FewShotExample.summarization_section_list
	SummarizationSectionList *SummarizationSectionList `json:"summarizationSectionList,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.GeneratorSuggestion
type GeneratorSuggestionWithRequired struct {
	// Optional. Free form suggestion.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorSuggestion.free_form_suggestion
	FreeFormSuggestion *FreeFormSuggestionWithRequired `json:"freeFormSuggestion,omitempty"`

	// Optional. Suggested summary.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorSuggestion.summary_suggestion
	SummarySuggestion *SummarySuggestionWithRequired `json:"summarySuggestion,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.FreeFormSuggestion
type FreeFormSuggestionWithRequired struct {
	// Required. Free form suggestion.
	// +kcc:proto:field=google.cloud.dialogflow.v2.FreeFormSuggestion.response
	// +required
	Response *string `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SummarySuggestion
type SummarySuggestionWithRequired struct {
	// Required. All the parts of generated summary.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SummarySuggestion.summary_sections
	// +required
	SummarySections []SummarySuggestion_SummarySectionWithRequired `json:"summarySections,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SummarySuggestion.SummarySection
type SummarySuggestion_SummarySectionWithRequired struct {
	// Required. Name of the section.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SummarySuggestion.SummarySection.section
	// +required
	Section *string `json:"section,omitempty"`

	// Required. Summary text for the section.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SummarySuggestion.SummarySection.summary
	// +required
	Summary *string `json:"summary,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.InputDataConfig
type GeneratorEvaluationConfig_InputDataConfig struct {
	// Optional. Desired number of conversation-summary pairs to be evaluated.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.InputDataConfig.sample_size
	SampleSize *int32 `json:"sampleSize,omitempty"`

	// Optional. Option to control whether summaries are generated during
	//  evaluation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.InputDataConfig.summary_generation_option
	SummaryGenerationOption *string `json:"summaryGenerationOption,omitempty"`

	// The distinctive configs for Agent Assist conversations as the
	//  conversation source.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.InputDataConfig.agent_assist_input_data_config
	AgentAssistInputDataConfig *GeneratorEvaluationConfig_AgentAssistInputDataConfig `json:"agentAssistInputDataConfig,omitempty"`

	// The distinctive configs for dataset as the conversation source.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.InputDataConfig.dataset_input_data_config
	DatasetInputDataConfig *GeneratorEvaluationConfig_DatasetInputDataConfig `json:"datasetInputDataConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.DatasetInputDataConfig
type GeneratorEvaluationConfig_DatasetInputDataConfig struct {
	// Required. The VertexAIDataset to be evaluated.
	// Format: `projects/<ProjectId>/locations/<LocationID>/datasets/<DatasetID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.DatasetInputDataConfig.dataset
	// +required
	DatasetRef *vertexaiv1beta1.VertexAIDatasetRef `json:"datasetRef"`
}

// +kcc:proto=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.SummarizationConfig
type GeneratorEvaluationConfig_SummarizationConfig struct {
	// Optional. Version for summarization evaluation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluationConfig.SummarizationConfig.evaluator_version
	EvaluatorVersion *string `json:"evaluatorVersion,omitempty"`
}

// DialogflowGeneratorEvaluationStatus defines the config connector machine state of DialogflowGeneratorEvaluation
type DialogflowGeneratorEvaluationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowGeneratorEvaluation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowGeneratorEvaluationObservedState `json:"observedState,omitempty"`
}

// DialogflowGeneratorEvaluationObservedState is the state of the DialogflowGeneratorEvaluation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.GeneratorEvaluation
type DialogflowGeneratorEvaluationObservedState struct {
	// Output only. Creation time of this generator evaluation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Completion time of this generator evaluation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Required. The initial generator that was used when creating this
	//  evaluation. This is a copy of the generator read from storage when creating
	//  the evaluation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.initial_generator
	InitialGenerator *DialogflowGeneratorObservedState `json:"initialGenerator,omitempty"`

	// Output only. Only available when the summarization generator is provided.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.summarization_metrics
	SummarizationMetrics *SummarizationEvaluationMetricsObservedState `json:"summarizationMetrics,omitempty"`

	// Output only. The result status of the evaluation pipeline. Provides the
	//  status information including if the evaluation is still in progress,
	//  completed or failed with certain error and user actionable message.
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.evaluation_status
	EvaluationStatus *EvaluationStatusObservedState `json:"evaluationStatus,omitempty"`

	// Output only. A read only boolean field reflecting Zone Separation
	//  status of the model. The field is an aggregated
	//  value of ZS status of its underlying dependencies. See more details in
	//  go/zicy-resource-placement#resource-status
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. A read only boolean field reflecting Zone Isolation status
	//  of the model. The field is an aggregated value of
	//  ZI status of its underlying dependencies. See more details in
	//  go/zicy-resource-placement#resource-status
	// +kcc:proto:field=google.cloud.dialogflow.v2.GeneratorEvaluation.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowgeneratorevaluation;gcpdialogflowgeneratorevaluations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowGeneratorEvaluation is the Schema for the DialogflowGeneratorEvaluation API
// +k8s:openapi-gen=true
type DialogflowGeneratorEvaluation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowGeneratorEvaluationSpec   `json:"spec,omitempty"`
	Status DialogflowGeneratorEvaluationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowGeneratorEvaluationList contains a list of DialogflowGeneratorEvaluation
type DialogflowGeneratorEvaluationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowGeneratorEvaluation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowGeneratorEvaluation{}, &DialogflowGeneratorEvaluationList{})
}
