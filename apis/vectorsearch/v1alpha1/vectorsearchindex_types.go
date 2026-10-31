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

// API sources for VectorsearchIndex, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/vectorsearch/v1/vectorsearch_service.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/docs/vector-search/overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/vector-search-2/reference/rest/v1/projects.locations.collections.indexes

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VectorsearchIndexGVK = GroupVersion.WithKind("VectorsearchIndex")

// VectorsearchIndexSpec defines the desired state of VectorsearchIndex
// +kcc:spec:proto=google.cloud.vectorsearch.v1.Index
// +kcc:required-from-proto
type VectorsearchIndexSpec struct {
	// The VectorSearchCollection this VectorsearchIndex is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *VectorSearchCollectionRef `json:"parentRef"`

	// The VectorsearchIndex name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Dedicated infrastructure for the index.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.dedicated_infrastructure
	DedicatedInfrastructure *DedicatedInfrastructure `json:"dedicatedInfrastructure,omitempty"`

	// Optional. Dense ScaNN index.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.dense_scann
	DenseScann *DenseScannIndex `json:"denseScann,omitempty"`

	// Optional. User-specified display name of the index
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-specified description of the index
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.description
	Description *string `json:"description,omitempty"`

	// Optional. Distance metric used for indexing. If not specified, will default
	//  to DOT_PRODUCT.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.distance_metric
	DistanceMetric *string `json:"distanceMetric,omitempty"`

	// Required. The collection schema field to index.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.index_field
	// +required
	IndexField *string `json:"indexField,omitempty"`

	// Optional. The fields to push into the index to enable fast ANN inline
	//  filtering.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.filter_fields
	FilterFields []string `json:"filterFields,omitempty"`

	// Optional. The fields to push into the index to enable inline data
	//  retrieval.
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.store_fields
	StoreFields []string `json:"storeFields,omitempty"`
}

// VectorsearchIndexStatus defines the config connector machine state of VectorsearchIndex
type VectorsearchIndexStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VectorsearchIndex resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VectorsearchIndexObservedState `json:"observedState,omitempty"`
}

// VectorsearchIndexObservedState is the state of the VectorsearchIndex resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vectorsearch.v1.Index
type VectorsearchIndexObservedState struct {
	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.vectorsearch.v1.Index.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvectorsearchindex;gcpvectorsearchindexes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VectorsearchIndex is the Schema for the VectorsearchIndex API
// +k8s:openapi-gen=true
type VectorsearchIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VectorsearchIndexSpec   `json:"spec,omitempty"`
	Status VectorsearchIndexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VectorsearchIndexList contains a list of VectorsearchIndex
type VectorsearchIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VectorsearchIndex `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VectorsearchIndex{}, &VectorsearchIndexList{})
}
