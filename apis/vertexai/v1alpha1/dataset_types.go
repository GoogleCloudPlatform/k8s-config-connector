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

var VertexAIDatasetGVK = GroupVersion.WithKind("VertexAIDataset")

// VertexAIDatasetSpec defines the desired state of VertexAIDataset
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.Dataset
type VertexAIDatasetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAIDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user-defined name of the Dataset.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the Dataset.
	Description *string `json:"description,omitempty"`

	// Required. Points to a YAML file stored on Google Cloud Storage describing
	// additional information about the Dataset. The schema is defined as an
	// OpenAPI 3.0.2 Schema Object. The schema files that can be used here are
	// found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	MetadataSchemaURI *string `json:"metadataSchemaURI,omitempty"`

	// Required. Additional information about the Dataset.
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`
	// The labels with user-defined metadata to organize your Datasets.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one Dataset (System
	// labels are excluded).
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	// and are immutable. Following system labels exist for each Dataset:
	//
	// * "aiplatform.googleapis.com/dataset_metadata_schema": output only, its
	//   value is the
	//   [metadata_schema's][google.cloud.aiplatform.v1beta1.Dataset.metadata_schema_uri]
	//   title.
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key spec for a Dataset. If set, this Dataset
	// and all sub-resources of this Dataset will be secured by this key.
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIDatasetStatus defines the config connector machine state of VertexAIDataset
type VertexAIDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIDatasetObservedState `json:"observedState,omitempty"`
}

// VertexAIDatasetObservedState is the state of the VertexAIDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.Dataset
type VertexAIDatasetObservedState struct {
	// Output only. The number of DataItems in this Dataset. Only apply for
	// non-structured Dataset.
	DataItemCount *int64 `json:"dataItemCount,omitempty"`

	// Output only. Timestamp when this Dataset was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Dataset was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// All SavedQueries belong to the Dataset will be returned in List/Get
	// Dataset response. The annotation_specs field
	// will not be populated except for UI cases which will only use
	// [annotation_spec_count][google.cloud.aiplatform.v1beta1.SavedQuery.annotation_spec_count].
	// In CreateDataset request, a SavedQuery is created together if
	// this field is set, up to one SavedQuery can be set in CreateDatasetRequest.
	// The SavedQuery should not contain any AnnotationSpec.
	SavedQueries []SavedQueryObservedState `json:"savedQueries,omitempty"`

	// Output only. The resource name of the Artifact that was created in
	// MetadataStore when creating the Dataset. The Artifact resource name pattern
	// is
	// `projects/{project}/locations/{location}/metadataStores/{metadata_store}/artifacts/{artifact}`.
	MetadataArtifact *string `json:"metadataArtifact,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaidataset;gcpvertexaidatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIDataset is the Schema for the VertexAIDataset API
// +k8s:openapi-gen=true
type VertexAIDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIDatasetSpec   `json:"spec,omitempty"`
	Status VertexAIDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIDatasetList contains a list of VertexAIDataset
type VertexAIDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIDataset{}, &VertexAIDatasetList{})
}
