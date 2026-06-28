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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PubSubSchemaGVK = GroupVersion.WithKind("PubSubSchema")

// PubSubSchemaSpec defines the desired state of PubSubSchema
// +kcc:spec:proto=google.pubsub.v1.Schema
type PubSubSchemaSpec struct {
	// The project that this resource belongs to.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// The PubSubSchema name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The definition of the schema. This should contain a string representing
	//  the full definition of the schema that is a valid schema definition of
	//  the type specified in `type`.
	// +kcc:proto:field=google.pubsub.v1.Schema.definition
	Definition *string `json:"definition,omitempty"`

	// The type of the schema definition.
	// +kcc:proto:field=google.pubsub.v1.Schema.type
	Type *string `json:"type,omitempty"`
}

// PubSubSchemaStatus defines the config connector machine state of PubSubSchema
type PubSubSchemaStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsubschema;gcppubsubschemas
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubSchema is the Schema for the PubSubSchema API
// +k8s:openapi-gen=true
type PubSubSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PubSubSchemaSpec   `json:"spec,omitempty"`
	Status PubSubSchemaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PubSubSchemaList contains a list of PubSubSchema
type PubSubSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubSchema `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubSchema{}, &PubSubSchemaList{})
}
