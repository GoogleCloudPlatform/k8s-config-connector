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

// API sources for MetastoreMetadataImport, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/metastore/v1/metastore.proto
// +kcc:source:service-docs=https://cloud.google.com/dataproc-metastore/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataproc-metastore/docs/reference/rest/v1/projects.locations.services.metadataImports

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MetastoreMetadataImportGVK = GroupVersion.WithKind("MetastoreMetadataImport")

// MetastoreMetadataImportSpec defines the desired state of MetastoreMetadataImport
// +kcc:spec:proto=google.cloud.metastore.v1.MetadataImport
// +kcc:required-from-proto
type MetastoreMetadataImportSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/services/{service} this resource belongs to.
	// +kcc:guess
	// ServiceRef *ServiceRef `json:"serviceRef,omitempty"`

	// The MetastoreMetadataImport name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. A database dump from a pre-existing metastore's database.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.database_dump
	DatabaseDump *MetadataImport_DatabaseDump `json:"databaseDump,omitempty"`

	// The description of the metadata import.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.description
	Description *string `json:"description,omitempty"`
}

// MetastoreMetadataImportStatus defines the config connector machine state of MetastoreMetadataImport
type MetastoreMetadataImportStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MetastoreMetadataImport resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MetastoreMetadataImportObservedState `json:"observedState,omitempty"`
}

// MetastoreMetadataImportObservedState is the state of the MetastoreMetadataImport resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.metastore.v1.MetadataImport
type MetastoreMetadataImportObservedState struct {
	// Output only. The time when the metadata import was started.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the metadata import was last updated.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time when the metadata import finished.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the metadata import.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataImport.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmetastoremetadataimport;gcpmetastoremetadataimports
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MetastoreMetadataImport is the Schema for the MetastoreMetadataImport API
// +k8s:openapi-gen=true
type MetastoreMetadataImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MetastoreMetadataImportSpec   `json:"spec,omitempty"`
	Status MetastoreMetadataImportStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MetastoreMetadataImportList contains a list of MetastoreMetadataImport
type MetastoreMetadataImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetastoreMetadataImport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetastoreMetadataImport{}, &MetastoreMetadataImportList{})
}
