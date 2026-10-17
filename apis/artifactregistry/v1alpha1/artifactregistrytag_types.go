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

// API sources for ArtifactRegistryTag, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/devtools/artifactregistry/v1/tag.proto
// +kcc:source:service-docs=https://cloud.google.com/artifacts/docs/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/artifacts/docs/reference/rest/v1/projects.locations.repositories.packages.tags

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ArtifactRegistryTagGVK = GroupVersion.WithKind("ArtifactRegistryTag")

// ArtifactRegistryTagSpec defines the desired state of ArtifactRegistryTag
// +kcc:spec:proto=google.devtools.artifactregistry.v1.Tag
// +kcc:required-from-proto
type ArtifactRegistryTagSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ArtifactRegistryTag name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The name of the version the tag refers to, for example:
	//  `projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811`
	//  If the package or version ID parts contain slashes, the slashes are
	//  escaped.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Tag.version
	Version *string `json:"version,omitempty"`
}

// ArtifactRegistryTagStatus defines the config connector machine state of ArtifactRegistryTag
type ArtifactRegistryTagStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ArtifactRegistryTag resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *ArtifactRegistryTagObservedState `json:"observedState,omitempty"`
}

// ArtifactRegistryTagObservedState is the state of the ArtifactRegistryTag resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.artifactregistry.v1.Tag
// type ArtifactRegistryTagObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpartifactregistrytag;gcpartifactregistrytags
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ArtifactRegistryTag is the Schema for the ArtifactRegistryTag API
// +k8s:openapi-gen=true
type ArtifactRegistryTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ArtifactRegistryTagSpec   `json:"spec,omitempty"`
	Status ArtifactRegistryTagStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ArtifactRegistryTagList contains a list of ArtifactRegistryTag
type ArtifactRegistryTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArtifactRegistryTag `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArtifactRegistryTag{}, &ArtifactRegistryTagList{})
}
