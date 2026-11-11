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

// API sources for DialogflowCompanionAgent, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/v2/companion_agent.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v2/projects.locations.companionAgents

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowCompanionAgentGVK = GroupVersion.WithKind("DialogflowCompanionAgent")

// DialogflowCompanionAgentSpec defines the desired state of DialogflowCompanionAgent
// +kcc:spec:proto=google.cloud.dialogflow.v2.CompanionAgent
// +kcc:required-from-proto
type DialogflowCompanionAgentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DialogflowCompanionAgent name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. List of CES toolset specs that the companion agent can choose
	//  from.
	// +kcc:proto:field=google.cloud.dialogflow.v2.CompanionAgent.toolset_tools
	ToolsetTools []ToolsetTool `json:"toolsetTools,omitempty"`

	// Optional. List of CES tool specs that the companion agent can choose from.
	// +kcc:proto:field=google.cloud.dialogflow.v2.CompanionAgent.ces_tool_specs
	CesToolSpecs []CesToolSpec `json:"cesToolSpecs,omitempty"`

	// Optional. Display name for the companion agent.
	//  Character limit is 63.
	// +kcc:proto:field=google.cloud.dialogflow.v2.CompanionAgent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description for the companion agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.CompanionAgent.description
	Description *string `json:"description,omitempty"`

	// Optional. List of skill configs for the companion agent.
	//  Allows at most one instance of each
	//  [SkillConfig][google.cloud.dialogflow.v2.CompanionAgent.SkillConfig] type.
	// +kcc:proto:field=google.cloud.dialogflow.v2.CompanionAgent.skill_configs
	SkillConfigs []CompanionAgent_SkillConfig `json:"skillConfigs,omitempty"`
}

// DialogflowCompanionAgentStatus defines the config connector machine state of DialogflowCompanionAgent
type DialogflowCompanionAgentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowCompanionAgent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowCompanionAgentObservedState `json:"observedState,omitempty"`
}

// DialogflowCompanionAgentObservedState is the state of the DialogflowCompanionAgent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.CompanionAgent
type DialogflowCompanionAgentObservedState struct {
	// Output only. Creation time of this companion agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.CompanionAgent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time of this companion agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.CompanionAgent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcompanionagent;gcpdialogflowcompanionagents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCompanionAgent is the Schema for the DialogflowCompanionAgent API
// +k8s:openapi-gen=true
type DialogflowCompanionAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCompanionAgentSpec   `json:"spec,omitempty"`
	Status DialogflowCompanionAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCompanionAgentList contains a list of DialogflowCompanionAgent
type DialogflowCompanionAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCompanionAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCompanionAgent{}, &DialogflowCompanionAgentList{})
}
