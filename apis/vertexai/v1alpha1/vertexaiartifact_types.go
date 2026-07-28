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

var VertexAIArtifactGVK = GroupVersion.WithKind("VertexAIArtifact")

// VertexAIArtifactSpec defines the desired state of VertexAIArtifact
type VertexAIArtifactSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent MetadataStore of this resource.
	// +required
	MetadataStoreRef *MetadataStoreRef `json:"metadataStoreRef"`

	// The VertexAIArtifact name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User provided display name of the Artifact.
	//  May be up to 128 Unicode characters.
	DisplayName *string `json:"displayName,omitempty"`

	// The uniform resource identifier of the artifact file.
	//  May be empty if there is no actual artifact file.
	URI *string `json:"uri,omitempty"`

	// The labels with user-defined metadata to organize your Artifacts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Artifact (System
	//  labels are excluded).
	Labels map[string]string `json:"labels,omitempty"`

	// The state of this Artifact. This is a property of the Artifact, and does
	//  not imply or capture any ongoing process. This property is managed by
	//  clients (such as Vertex AI Pipelines), and the system does not prescribe
	//  or check the validity of state transitions.
	State *string `json:"state,omitempty"`

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

	// Properties of the Artifact.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	Metadata apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Description of the Artifact
	Description *string `json:"description,omitempty"`
}

// VertexAIArtifactStatus defines the config connector machine state of VertexAIArtifact
type VertexAIArtifactStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIArtifact resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIArtifactObservedState `json:"observedState,omitempty"`
}

// VertexAIArtifactObservedState is the state of the VertexAIArtifact resource as most recently observed in GCP.
type VertexAIArtifactObservedState struct {
	// Output only. The resource name of the Artifact.
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this Artifact was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Artifact was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiartifact;gcpvertexaiartifacts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIArtifact is the Schema for the VertexAIArtifact API
// +k8s:openapi-gen=true
type VertexAIArtifact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIArtifactSpec   `json:"spec,omitempty"`
	Status VertexAIArtifactStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIArtifactList contains a list of VertexAIArtifact
type VertexAIArtifactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIArtifact `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIArtifact{}, &VertexAIArtifactList{})
}
