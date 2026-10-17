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

// API sources for DialogflowExample, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/example.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v3/projects.locations.agents.playbooks.examples

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowExampleGVK = GroupVersion.WithKind("DialogflowExample")

// DialogflowExampleSpec defines the desired state of DialogflowExample
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Example
// +kcc:required-from-proto
type DialogflowExampleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/agents/{agent}/playbooks/{playbook} this resource belongs to.
	// +kcc:guess
	// PlaybookRef *DialogflowPlaybookRef `json:"playbookRef,omitempty"`

	// The DialogflowExample name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The input to the playbook in the example.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.playbook_input
	PlaybookInput *PlaybookInput `json:"playbookInput,omitempty"`

	// Optional. The output of the playbook in the example.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.playbook_output
	PlaybookOutput *PlaybookOutput `json:"playbookOutput,omitempty"`

	// Required. The ordered list of actions performed by the end user and the
	//  Dialogflow agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.actions
	// +required
	Actions []Action `json:"actions,omitempty"`

	// Required. The display name of the example.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The high level concise description of the example. The max number
	//  of characters is 200.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.description
	Description *string `json:"description,omitempty"`

	// Required. Example's output state.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.conversation_state
	// +required
	ConversationState *string `json:"conversationState,omitempty"`

	// Optional. The language code of the example.
	//  If not specified, the agent's default language is used.
	//  Note: languages must be enabled in the agent before they can be used.
	//  Note: example's language code is not currently used in dialogflow agents.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// DialogflowExampleStatus defines the config connector machine state of DialogflowExample
type DialogflowExampleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowExample resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowExampleObservedState `json:"observedState,omitempty"`
}

// DialogflowExampleObservedState is the state of the DialogflowExample resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Example
type DialogflowExampleObservedState struct {
	// Required. The ordered list of actions performed by the end user and the
	//  Dialogflow agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.actions
	Actions []ActionObservedState `json:"actions,omitempty"`

	// Output only. Estimated number of tokes current example takes when sent to
	//  the LLM.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.token_count
	TokenCount *int64 `json:"tokenCount,omitempty"`

	// Output only. The timestamp of initial example creation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last time the example was updated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Example.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowexample;gcpdialogflowexamples
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowExample is the Schema for the DialogflowExample API
// +k8s:openapi-gen=true
type DialogflowExample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowExampleSpec   `json:"spec,omitempty"`
	Status DialogflowExampleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowExampleList contains a list of DialogflowExample
type DialogflowExampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowExample `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowExample{}, &DialogflowExampleList{})
}
