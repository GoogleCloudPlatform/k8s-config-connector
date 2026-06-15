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
)

var ApigeeRegistryArtifactGVK = GroupVersion.WithKind("ApigeeRegistryArtifact")

// ApigeeRegistryArtifactSpec defines the desired state of ApigeeRegistryArtifact
// +kcc:spec:proto=google.cloud.apigeeregistry.v1.Artifact
type ApigeeRegistryArtifactSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ApigeeRegistryArtifact name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A content type specifier for the artifact.
	// Content type specifiers are Media Types
	// (https://en.wikipedia.org/wiki/Media_type) with a possible "schema"
	// parameter that specifies a schema for the stored information.
	// Content types can specify compression. Currently only GZip compression is
	// supported (indicated with "+gzip").
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Input only. The contents of the artifact.
	// Provided by API callers when artifacts are created or replaced.
	// To access the contents of an artifact, use GetArtifactContents.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.contents
	Contents []byte `json:"contents,omitempty"`
}

// ApigeeRegistryArtifactStatus defines the config connector machine state of ApigeeRegistryArtifact
type ApigeeRegistryArtifactStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeRegistryArtifact resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *ApigeeRegistryArtifactObservedState `json:"observedState,omitempty"`
}

// ApigeeRegistryArtifactObservedState is the state of the ApigeeRegistryArtifact resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.apigeeregistry.v1.Artifact
type ApigeeRegistryArtifactObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The size of the artifact in bytes. If the artifact is gzipped, this is
	// the size of the uncompressed artifact.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.size_bytes
	SizeBytes *int32 `json:"sizeBytes,omitempty"`

	// Output only. A SHA-256 hash of the artifact's contents. If the artifact is gzipped,
	// this is the hash of the uncompressed artifact.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.hash
	Hash *string `json:"hash,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeregistryartifact;gcpapigeeregistryartifacts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeRegistryArtifact is the Schema for the ApigeeRegistryArtifact API
// +k8s:openapi-gen=true
type ApigeeRegistryArtifact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeRegistryArtifactSpec   `json:"spec,omitempty"`
	Status ApigeeRegistryArtifactStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeRegistryArtifactList contains a list of ApigeeRegistryArtifact
type ApigeeRegistryArtifactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeRegistryArtifact `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeRegistryArtifact{}, &ApigeeRegistryArtifactList{})
}
