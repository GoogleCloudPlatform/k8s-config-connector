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

// API sources for DialogflowTransitionRouteGroup, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/transition_route_group.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.transitionRouteGroups

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowTransitionRouteGroupGVK = GroupVersion.WithKind("DialogflowTransitionRouteGroup")

// DialogflowTransitionRouteGroupSpec defines the desired state of DialogflowTransitionRouteGroup
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.TransitionRouteGroup
// +kcc:required-from-proto
type DialogflowTransitionRouteGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DialogflowTransitionRouteGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the transition route group, unique
	//  within the flow. The display name can be no longer than 30 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRouteGroup.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Transition routes associated with the
	//  [TransitionRouteGroup][google.cloud.dialogflow.cx.v3.TransitionRouteGroup].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRouteGroup.transition_routes
	TransitionRoutes []TransitionRoute `json:"transitionRoutes,omitempty"`
}

// DialogflowTransitionRouteGroupStatus defines the config connector machine state of DialogflowTransitionRouteGroup
type DialogflowTransitionRouteGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowTransitionRouteGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowTransitionRouteGroupObservedState `json:"observedState,omitempty"`
}

// DialogflowTransitionRouteGroupObservedState is the state of the DialogflowTransitionRouteGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.TransitionRouteGroup
type DialogflowTransitionRouteGroupObservedState struct {
	// Transition routes associated with the
	//  [TransitionRouteGroup][google.cloud.dialogflow.cx.v3.TransitionRouteGroup].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TransitionRouteGroup.transition_routes
	TransitionRoutes []TransitionRouteObservedState `json:"transitionRoutes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowtransitionroutegroup;gcpdialogflowtransitionroutegroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowTransitionRouteGroup is the Schema for the DialogflowTransitionRouteGroup API
// +k8s:openapi-gen=true
type DialogflowTransitionRouteGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowTransitionRouteGroupSpec   `json:"spec,omitempty"`
	Status DialogflowTransitionRouteGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowTransitionRouteGroupList contains a list of DialogflowTransitionRouteGroup
type DialogflowTransitionRouteGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowTransitionRouteGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowTransitionRouteGroup{}, &DialogflowTransitionRouteGroupList{})
}
