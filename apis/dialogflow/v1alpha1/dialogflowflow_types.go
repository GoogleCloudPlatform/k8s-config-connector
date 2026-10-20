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

// API sources for DialogflowFlow, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/flow.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowFlowGVK = GroupVersion.WithKind("DialogflowFlow")

// DialogflowFlowSpec defines the desired state of DialogflowFlow
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Flow
// +kcc:required-from-proto
type DialogflowFlowSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/agents/{agent} this resource belongs to.
	// +kcc:guess
	// AgentRef *DialogflowCXAgentRef `json:"agentRef,omitempty"`

	// The DialogflowFlow name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the flow. The maximum length is 500 characters. If
	//  exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.description
	Description *string `json:"description,omitempty"`

	// A flow's transition routes serve two purposes:
	//
	//  *   They are responsible for matching the user's first utterances in the
	//  flow.
	//  *   They are inherited by every page's [transition
	//  routes][Page.transition_routes] and can support use cases such as the user
	//  saying "help" or "can I talk to a human?", which can be handled in a common
	//  way regardless of the current page. Transition routes defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  TransitionRoutes are evaluated in the following order:
	//
	//  *   TransitionRoutes with intent specified.
	//  *   TransitionRoutes with only condition specified.
	//
	//  TransitionRoutes with intent specified are inherited by pages in the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.transition_routes
	TransitionRoutes []TransitionRoute `json:"transitionRoutes,omitempty"`

	// A flow's event handlers serve two purposes:
	//
	//  *   They are responsible for handling events (e.g. no match,
	//  webhook errors) in the flow.
	//  *   They are inherited by every page's [event
	//  handlers][Page.event_handlers], which can be used to handle common events
	//  regardless of the current page. Event handlers defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  Unlike
	//  [transition_routes][google.cloud.dialogflow.cx.v3.Flow.transition_routes],
	//  these handlers are evaluated on a first-match basis. The first one that
	//  matches the event get executed, with the rest being ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.event_handlers
	EventHandlers []EventHandler `json:"eventHandlers,omitempty"`

	// A flow's transition route group serve two purposes:
	//
	//  *   They are responsible for matching the user's first utterances in the
	//  flow.
	//  *   They are inherited by every page's [transition
	//  route groups][Page.transition_route_groups]. Transition route groups
	//  defined in the page have higher priority than those defined in the flow.
	//
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  for agent-level groups.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.transition_route_groups
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	// NLU related settings of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.nlu_settings
	NluSettings *NluSettings `json:"nluSettings,omitempty"`

	// Hierarchical advanced settings for this flow. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`

	// Optional. Defined structured input parameters for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.input_parameter_definitions
	InputParameterDefinitions []ParameterDefinition `json:"inputParameterDefinitions,omitempty"`

	// Optional. Defined structured output parameters for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.output_parameter_definitions
	OutputParameterDefinitions []ParameterDefinition `json:"outputParameterDefinitions,omitempty"`

	// Optional. Multi-lingual agent settings for this flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.multi_language_settings
	MultiLanguageSettings *Flow_MultiLanguageSettings `json:"multiLanguageSettings,omitempty"`

	// Indicates whether the flow is locked for changes. If the flow is locked,
	//  modifications to the flow will be rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.locked
	Locked *bool `json:"locked,omitempty"`
}

// DialogflowFlowStatus defines the config connector machine state of DialogflowFlow
type DialogflowFlowStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowFlow resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowFlowObservedState `json:"observedState,omitempty"`
}

// DialogflowFlowObservedState is the state of the DialogflowFlow resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Flow
type DialogflowFlowObservedState struct {
	// A flow's transition routes serve two purposes:
	//
	//  *   They are responsible for matching the user's first utterances in the
	//  flow.
	//  *   They are inherited by every page's [transition
	//  routes][Page.transition_routes] and can support use cases such as the user
	//  saying "help" or "can I talk to a human?", which can be handled in a common
	//  way regardless of the current page. Transition routes defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  TransitionRoutes are evaluated in the following order:
	//
	//  *   TransitionRoutes with intent specified.
	//  *   TransitionRoutes with only condition specified.
	//
	//  TransitionRoutes with intent specified are inherited by pages in the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.transition_routes
	TransitionRoutes []TransitionRouteObservedState `json:"transitionRoutes,omitempty"`

	// A flow's event handlers serve two purposes:
	//
	//  *   They are responsible for handling events (e.g. no match,
	//  webhook errors) in the flow.
	//  *   They are inherited by every page's [event
	//  handlers][Page.event_handlers], which can be used to handle common events
	//  regardless of the current page. Event handlers defined in the page
	//  have higher priority than those defined in the flow.
	//
	//  Unlike
	//  [transition_routes][google.cloud.dialogflow.cx.v3.Flow.transition_routes],
	//  these handlers are evaluated on a first-match basis. The first one that
	//  matches the event get executed, with the rest being ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.event_handlers
	EventHandlers []EventHandlerObservedState `json:"eventHandlers,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Flow.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettingsObservedState `json:"knowledgeConnectorSettings,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowflow;gcpdialogflowflows
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowFlow is the Schema for the DialogflowFlow API
// +k8s:openapi-gen=true
type DialogflowFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowFlowSpec   `json:"spec,omitempty"`
	Status DialogflowFlowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowFlowList contains a list of DialogflowFlow
type DialogflowFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowFlow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowFlow{}, &DialogflowFlowList{})
}
