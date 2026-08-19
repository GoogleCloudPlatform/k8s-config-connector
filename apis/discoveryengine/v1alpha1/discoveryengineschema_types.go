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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineSchemaGVK = GroupVersion.WithKind("DiscoveryEngineSchema")

// DiscoveryEngineSchemaSpec defines the desired state of DiscoveryEngineSchema
// +kcc:spec:proto=google.cloud.discoveryengine.v1.Schema
type DiscoveryEngineSchemaSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// Immutable. The DataStore this schema belongs to.
	// +required
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef"`

	// The DiscoveryEngineSchema name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The structured representation of the schema.
	StructSchema *apiextensionsv1.JSON `json:"structSchema,omitempty"`

	// The JSON representation of the schema.
	JsonSchema *string `json:"jsonSchema,omitempty"`
}

// DiscoveryEngineSchemaStatus defines the config connector machine state of DiscoveryEngineSchema
type DiscoveryEngineSchemaStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineSchema resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineSchemaObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineSchemaObservedState is the state of the DiscoveryEngineSchema resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.Schema
type DiscoveryEngineSchemaObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryengineschema;gcpdiscoveryengineschemas
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineSchema is the Schema for the DiscoveryEngineSchema API
// +k8s:openapi-gen=true
type DiscoveryEngineSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineSchemaSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineSchemaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineSchemaList contains a list of DiscoveryEngineSchema
type DiscoveryEngineSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineSchema `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineSchema{}, &DiscoveryEngineSchemaList{})
}
