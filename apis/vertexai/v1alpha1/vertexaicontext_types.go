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

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var VertexAIContextGVK = GroupVersion.WithKind("VertexAIContext")

// VertexAIContextSpec defines the desired state of VertexAIContext
type VertexAIContextSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent MetadataStore of this resource.
	// +required
	MetadataStoreRef *MetadataStoreRef `json:"metadataStoreRef"`

	// The VertexAIContext name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User provided display name of the Context.
	//  May be up to 128 Unicode characters.
	DisplayName *string `json:"displayName,omitempty"`

	// The labels with user-defined metadata to organize your Contexts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Context (System
	//  labels are excluded).
	Labels map[string]string `json:"labels,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Context.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	Metadata apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Description of the Context
	Description *string `json:"description,omitempty"`
}

// VertexAIContextStatus defines the config connector machine state of VertexAIContext
type VertexAIContextStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIContext resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIContextObservedState `json:"observedState,omitempty"`
}

// VertexAIContextObservedState is the state of the VertexAIContext resource as most recently observed in GCP.
type VertexAIContextObservedState struct {
	// Output only. Timestamp when this Context was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Context was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A list of resource names of Contexts that are parents of this
	//  Context. A Context may have at most 10 parent_contexts.
	ParentContexts []string `json:"parentContexts,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaicontext;gcpvertexaicontexts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIContext is the Schema for the VertexAIContext API
// +k8s:openapi-gen=true
type VertexAIContext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIContextSpec   `json:"spec,omitempty"`
	Status VertexAIContextStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIContextList contains a list of VertexAIContext
type VertexAIContextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIContext `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIContext{}, &VertexAIContextList{})
}
