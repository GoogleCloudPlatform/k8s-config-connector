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
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIDatasetVersionGVK = GroupVersion.WithKind("VertexAIDatasetVersion")

// VertexAIDatasetVersionSpec defines the desired state of VertexAIDatasetVersion
type VertexAIDatasetVersionSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent Dataset of this resource.
	// +required
	DatasetRef *v1beta1.VertexAIDatasetRef `json:"datasetRef"`

	// The VertexAIDatasetVersion name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The user-defined name of the DatasetVersion.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	DisplayName *string `json:"displayName,omitempty"`
}

// VertexAIDatasetVersionStatus defines the config connector machine state of VertexAIDatasetVersion
type VertexAIDatasetVersionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIDatasetVersion resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIDatasetVersionObservedState `json:"observedState,omitempty"`
}

// VertexAIDatasetVersionObservedState is the state of the VertexAIDatasetVersion resource as most recently observed in GCP.
type VertexAIDatasetVersionObservedState struct {
	// Output only. Identifier. The resource name of the DatasetVersion.
	//  Format:
	//  `projects/{project}/locations/{location}/datasets/{dataset}/datasetVersions/{dataset_version}`
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this DatasetVersion was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this DatasetVersion was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Name of the associated BigQuery dataset.
	BigQueryDatasetName *string `json:"bigQueryDatasetName,omitempty"`

	// Required. Output only. Additional information about the DatasetVersion.
	Metadata *Value `json:"metadata,omitempty"`

	// Output only. Reference to the public base model last used by the dataset
	//  version. Only set for prompt dataset versions.
	ModelReference *string `json:"modelReference,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaidatasetversion;gcpvertexaidatasetversions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIDatasetVersion is the Schema for the VertexAIDatasetVersion API
// +k8s:openapi-gen=true
type VertexAIDatasetVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIDatasetVersionSpec   `json:"spec,omitempty"`
	Status VertexAIDatasetVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIDatasetVersionList contains a list of VertexAIDatasetVersion
type VertexAIDatasetVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIDatasetVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIDatasetVersion{}, &VertexAIDatasetVersionList{})
}
