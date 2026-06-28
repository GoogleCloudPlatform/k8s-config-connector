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

var DialogflowGeneratorGVK = GroupVersion.WithKind("DialogflowGenerator")

// DialogflowGeneratorSpec defines the desired state of DialogflowGenerator
// +kcc:spec:proto=google.cloud.dialogflow.v2.Generator
type DialogflowGeneratorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location,omitempty"`

	// The DialogflowGenerator name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Human readable description of the generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.description
	Description *string `json:"description,omitempty"`

	// Input of free from generator to LLM.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.free_form_context
	FreeFormContext *FreeFormContext `json:"freeFormContext,omitempty"`

	// Input of prebuilt Summarization feature.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.summarization_context
	SummarizationContext *SummarizationContext `json:"summarizationContext,omitempty"`

	// Optional. Inference parameters for this generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.inference_parameter
	InferenceParameter *InferenceParameter `json:"inferenceParameter,omitempty"`

	// Optional. The trigger event of the generator. It defines when the generator
	// is triggered in a conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.trigger_event
	TriggerEvent *string `json:"triggerEvent,omitempty"`

	// Optional. The published Large Language Model name.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.published_model
	PublishedModel *string `json:"publishedModel,omitempty"`
}

// DialogflowGeneratorStatus defines the config connector machine state of DialogflowGenerator
type DialogflowGeneratorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowGenerator resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowGeneratorObservedState `json:"observedState,omitempty"`
}

// DialogflowGeneratorObservedState is the state of the DialogflowGenerator resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.Generator
type DialogflowGeneratorObservedState struct {
	// Output only. Creation time of this generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time of this generator.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Generator.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowgenerator;gcpdialogflowgenerators
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowGenerator is the Schema for the DialogflowGenerator API
// +k8s:openapi-gen=true
type DialogflowGenerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowGeneratorSpec   `json:"spec,omitempty"`
	Status DialogflowGeneratorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowGeneratorList contains a list of DialogflowGenerator
type DialogflowGeneratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowGenerator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowGenerator{}, &DialogflowGeneratorList{})
}
