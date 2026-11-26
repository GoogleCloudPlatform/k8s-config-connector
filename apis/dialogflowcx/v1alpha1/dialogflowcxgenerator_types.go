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

// API sources for DialogflowCXGenerator, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dialogflow/cx/v3/generator.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.generators

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowCXGeneratorGVK = GroupVersion.WithKind("DialogflowCXGenerator")

// DialogflowCXGeneratorSpec defines the desired state of DialogflowCXGenerator
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Generator
// +kcc:required-from-proto
type DialogflowCXGeneratorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/agents/{agent} this resource belongs to.
	// +kcc:guess
	// AgentRef *DialogflowCXAgentRef `json:"agentRef,omitempty"`

	// The DialogflowCXGenerator name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the generator, unique within the
	//  agent. The prompt contains pre-defined parameters such as $conversation,
	//  $last-user-utterance, etc. populated by Dialogflow. It can also contain
	//  custom placeholders which will be resolved during fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Prompt for the LLM model.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.prompt_text
	// +required
	PromptText *Phrase `json:"promptText,omitempty"`

	// Optional. List of custom placeholders in the prompt text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.placeholders
	Placeholders []Generator_Placeholder `json:"placeholders,omitempty"`

	// Parameters passed to the LLM to configure its behavior.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Generator.model_parameter
	ModelParameter *Generator_ModelParameter `json:"modelParameter,omitempty"`
}

// DialogflowCXGeneratorStatus defines the config connector machine state of DialogflowCXGenerator
type DialogflowCXGeneratorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowCXGenerator resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DialogflowCXGeneratorObservedState `json:"observedState,omitempty"`
}

// DialogflowCXGeneratorObservedState is the state of the DialogflowCXGenerator resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Generator
// type DialogflowCXGeneratorObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxgenerator;gcpdialogflowcxgenerators
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCXGenerator is the Schema for the DialogflowCXGenerator API
// +k8s:openapi-gen=true
type DialogflowCXGenerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCXGeneratorSpec   `json:"spec,omitempty"`
	Status DialogflowCXGeneratorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCXGeneratorList contains a list of DialogflowCXGenerator
type DialogflowCXGeneratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXGenerator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXGenerator{}, &DialogflowCXGeneratorList{})
}
