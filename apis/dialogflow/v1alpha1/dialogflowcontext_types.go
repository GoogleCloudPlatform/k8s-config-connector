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

// API sources for DialogflowContext, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dialogflow/v2/context.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.sessions.contexts

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowContextGVK = GroupVersion.WithKind("DialogflowContext")

// DialogflowContextSpec defines the desired state of DialogflowContext
// +kcc:spec:proto=google.cloud.dialogflow.v2.Context
// +kcc:required-from-proto
type DialogflowContextSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DialogflowContext name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The number of conversational query requests after which the
	//  context expires. The default is `0`. If set to `0`, the context expires
	//  immediately. Contexts expire automatically after 20 minutes if there
	//  are no matching queries.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Context.lifespan_count
	LifespanCount *int32 `json:"lifespanCount,omitempty"`

	// Optional. The collection of parameters associated with this context.
	//
	//  Depending on your protocol or client library language, this is a
	//  map, associative array, symbol table, dictionary, or JSON object
	//  composed of a collection of (MapKey, MapValue) pairs:
	//
	//  * MapKey type: string
	//  * MapKey value: parameter name
	//  * MapValue type: If parameter's entity type is a composite entity then use
	//  map, otherwise, depending on the parameter value type, it could be one of
	//  string, number, boolean, null, list or map.
	//  * MapValue value: If parameter's entity type is a composite entity then use
	//  map from composite entity property names to property values, otherwise,
	//  use parameter value.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Context.parameters
	Parameters apiextensionsv1.JSON `json:"parameters,omitempty"`
}

// DialogflowContextStatus defines the config connector machine state of DialogflowContext
type DialogflowContextStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowContext resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DialogflowContextObservedState `json:"observedState,omitempty"`
}

// DialogflowContextObservedState is the state of the DialogflowContext resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.Context
// type DialogflowContextObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcontext;gcpdialogflowcontexts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowContext is the Schema for the DialogflowContext API
// +k8s:openapi-gen=true
type DialogflowContext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowContextSpec   `json:"spec,omitempty"`
	Status DialogflowContextStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowContextList contains a list of DialogflowContext
type DialogflowContextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowContext `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowContext{}, &DialogflowContextList{})
}
