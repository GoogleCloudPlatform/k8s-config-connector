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

// API sources for AIPlatformArtifact, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/artifact.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores.artifacts

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformArtifactGVK = GroupVersion.WithKind("AIPlatformArtifact")

// AIPlatformArtifactSpec defines the desired state of AIPlatformArtifact
// +kcc:spec:proto=google.cloud.aiplatform.v1.Artifact
// +kcc:required-from-proto
type AIPlatformArtifactSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformArtifact name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User provided display name of the Artifact.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The uniform resource identifier of the artifact file.
	//  May be empty if there is no actual artifact file.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.uri
	URI *string `json:"uri,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Artifacts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Artifact (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The state of this Artifact. This is a property of the Artifact, and does
	//  not imply or capture any ongoing process. This property is managed by
	//  clients (such as Vertex AI Pipelines), and the system does not prescribe
	//  or check the validity of state transitions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.state
	State *string `json:"state,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Artifact.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.metadata
	Metadata apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Description of the Artifact
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.description
	Description *string `json:"description,omitempty"`
}

// AIPlatformArtifactStatus defines the config connector machine state of AIPlatformArtifact
type AIPlatformArtifactStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformArtifact resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformArtifactObservedState `json:"observedState,omitempty"`
}

// AIPlatformArtifactObservedState is the state of the AIPlatformArtifact resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Artifact
type AIPlatformArtifactObservedState struct {
	// Output only. Timestamp when this Artifact was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Artifact was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformartifact;gcpaiplatformartifacts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformArtifact is the Schema for the AIPlatformArtifact API
// +k8s:openapi-gen=true
type AIPlatformArtifact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformArtifactSpec   `json:"spec,omitempty"`
	Status AIPlatformArtifactStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformArtifactList contains a list of AIPlatformArtifact
type AIPlatformArtifactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformArtifact `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformArtifact{}, &AIPlatformArtifactList{})
}
