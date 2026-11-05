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

// API sources for DialogflowPage, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/page.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.pages

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowPageGVK = GroupVersion.WithKind("DialogflowPage")

// DialogflowPageSpec defines the desired state of DialogflowPage
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Page
// +kcc:required-from-proto
type DialogflowPageSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DialogflowPage name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the page, unique within the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the page. The maximum length is 500 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.description
	Description *string `json:"description,omitempty"`

	// The fulfillment to call when the session is entering the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.entry_fulfillment
	EntryFulfillment *Fulfillment `json:"entryFulfillment,omitempty"`

	// The form associated with the page, used for collecting parameters
	//  relevant to the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.form
	Form *Form `json:"form,omitempty"`

	// Ordered list of
	//  [`TransitionRouteGroups`][google.cloud.dialogflow.cx.v3.TransitionRouteGroup]
	//  added to the page. Transition route groups must be unique within a page. If
	//  the page links both flow-level transition route groups and agent-level
	//  transition route groups, the flow-level ones will have higher priority and
	//  will be put before the agent-level ones.
	//
	//  *   If multiple transition routes within a page scope refer to the same
	//      intent, then the precedence order is: page's transition route -> page's
	//      transition route group -> flow's transition routes.
	//
	//  *   If multiple transition route groups within a page contain the same
	//      intent, then the first group in the ordered list takes precedence.
	//
	//  Format:`projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  or
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/transitionRouteGroups/<TransitionRouteGroupID>`
	//  for agent-level groups.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.transition_route_groups
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	// A list of transitions for the transition rules of this page.
	//  They route the conversation to another page in the same flow, or another
	//  flow.
	//
	//  When we are in a certain page, the TransitionRoutes are evaluated in the
	//  following order:
	//
	//  *   TransitionRoutes defined in the page with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in flow with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Flow.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in the page with only condition specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with only condition specified.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.transition_routes
	TransitionRoutes []TransitionRoute `json:"transitionRoutes,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors,
	//  no match or no input.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.event_handlers
	EventHandlers []EventHandler `json:"eventHandlers,omitempty"`

	// Hierarchical advanced settings for this page. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettings `json:"knowledgeConnectorSettings,omitempty"`
}

// DialogflowPageStatus defines the config connector machine state of DialogflowPage
type DialogflowPageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowPage resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowPageObservedState `json:"observedState,omitempty"`
}

// DialogflowPageObservedState is the state of the DialogflowPage resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Page
type DialogflowPageObservedState struct {
	// The fulfillment to call when the session is entering the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.entry_fulfillment
	EntryFulfillment *FulfillmentObservedState `json:"entryFulfillment,omitempty"`

	// The form associated with the page, used for collecting parameters
	//  relevant to the page.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.form
	Form *FormObservedState `json:"form,omitempty"`

	// A list of transitions for the transition rules of this page.
	//  They route the conversation to another page in the same flow, or another
	//  flow.
	//
	//  When we are in a certain page, the TransitionRoutes are evaluated in the
	//  following order:
	//
	//  *   TransitionRoutes defined in the page with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in flow with intent specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Flow.transition_route_groups]
	//      with intent specified.
	//  *   TransitionRoutes defined in the page with only condition specified.
	//  *   TransitionRoutes defined in the
	//      [transition route
	//      groups][google.cloud.dialogflow.cx.v3.Page.transition_route_groups]
	//      with only condition specified.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.transition_routes
	TransitionRoutes []TransitionRouteObservedState `json:"transitionRoutes,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors,
	//  no match or no input.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.event_handlers
	EventHandlers []EventHandlerObservedState `json:"eventHandlers,omitempty"`

	// Optional. Knowledge connector configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Page.knowledge_connector_settings
	KnowledgeConnectorSettings *KnowledgeConnectorSettingsObservedState `json:"knowledgeConnectorSettings,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowpage;gcpdialogflowpages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowPage is the Schema for the DialogflowPage API
// +k8s:openapi-gen=true
type DialogflowPage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowPageSpec   `json:"spec,omitempty"`
	Status DialogflowPageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowPageList contains a list of DialogflowPage
type DialogflowPageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowPage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowPage{}, &DialogflowPageList{})
}
