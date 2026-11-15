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

// API sources for DataplexPartition, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/28ba5d15234eff44241492ca72b135e1171fc537/google/cloud/dataplex/v1/metadata.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.lakes.zones.entities.partitions

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexPartitionGVK = GroupVersion.WithKind("DataplexPartition")

// DataplexPartitionSpec defines the desired state of DataplexPartition
// +kcc:spec:proto=google.cloud.dataplex.v1.Partition
// +kcc:required-from-proto
type DataplexPartitionSpec struct {
	// The DataplexEntity that this partition belongs to.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *DataplexEntityRef `json:"parentRef"`

	// The DataplexPartition name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The set of values representing the partition, which
	//  correspond to the partition schema defined in the parent entity.
	// +kcc:proto:field=google.cloud.dataplex.v1.Partition.values
	// +required
	Values []string `json:"values,omitempty"`

	// Required. Immutable. The location of the entity data within the partition,
	//  for example, `gs://bucket/path/to/entity/key1=value1/key2=value2`. Or
	//  `projects/<project_id>/datasets/<dataset_id>/tables/<table_id>`
	// +kcc:proto:field=google.cloud.dataplex.v1.Partition.location
	// +required
	Location *string `json:"location,omitempty"`
}

// DataplexPartitionStatus defines the config connector machine state of DataplexPartition
type DataplexPartitionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexPartition resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DataplexPartitionObservedState `json:"observedState,omitempty"`
}

// DataplexPartitionObservedState is the state of the DataplexPartition resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.Partition
// type DataplexPartitionObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexpartition;gcpdataplexpartitions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexPartition is the Schema for the DataplexPartition API
// +k8s:openapi-gen=true
type DataplexPartition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexPartitionSpec   `json:"spec,omitempty"`
	Status DataplexPartitionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexPartitionList contains a list of DataplexPartition
type DataplexPartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexPartition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexPartition{}, &DataplexPartitionList{})
}
