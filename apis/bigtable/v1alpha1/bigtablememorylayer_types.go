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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableMemoryLayerGVK = GroupVersion.WithKind("BigtableMemoryLayer")

type BigtableMemoryLayerParent struct {
	// Immutable. The BigtableCluster whose memory layer this resource manages.
	// +required
	ClusterRef *ClusterRef `json:"clusterRef,omitempty"`
}

// BigtableMemoryLayerSpec defines the desired state of BigtableMemoryLayer.
//
// A memory layer serves reads from memory without hitting the backing
// persistent data store. It is a singleton of its cluster: the existence of a
// BigtableMemoryLayer object enables the memory layer on the referenced
// cluster, and deleting the object disables it.
// +kcc:spec:proto=google.bigtable.admin.v2.MemoryLayer
type BigtableMemoryLayerSpec struct {
	// NOTE: ResourceID field is not required for MemoryLayer as its ID has the
	// format projects/<projectID>/instances/<instanceID>/clusters/<clusterID>/memoryLayer
	// i.e., it doesnt have any unique ID of its own and relies on parent ID for uniqueness.

	BigtableMemoryLayerParent `json:",inline"`
}

// BigtableMemoryLayerStatus defines the config connector machine state of BigtableMemoryLayer
type BigtableMemoryLayerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableMemoryLayer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigtableMemoryLayerObservedState `json:"observedState,omitempty"`
}

// BigtableMemoryLayerObservedState is the state of the BigtableMemoryLayer resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.bigtable.admin.v2.MemoryLayer
type BigtableMemoryLayerObservedState struct {
	// Output only. The current state of the memory layer.
	// +optional
	// +kcc:proto:field=google.bigtable.admin.v2.MemoryLayer.state
	State *string `json:"state,omitempty"`

	// Output only. The configuration of this memory layer.
	// +optional
	// +kcc:proto:field=google.bigtable.admin.v2.MemoryLayer.memory_config
	MemoryConfig *MemoryLayer_MemoryConfigObservedState `json:"memoryConfig,omitempty"`
}

// MemoryLayer_MemoryConfigObservedState is the observed configuration of a memory layer.
// +kcc:observedstate:proto=google.bigtable.admin.v2.MemoryLayer.MemoryConfig
type MemoryLayer_MemoryConfigObservedState struct {
	// Output only. Reporting the current size of the memory layer in GiB.
	// +optional
	// +kcc:proto:field=google.bigtable.admin.v2.MemoryLayer.MemoryConfig.storage_size_gib
	StorageSizeGib *int32 `json:"storageSizeGib,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtablememorylayer;gcpbigtablememorylayers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableMemoryLayer is the Schema for the BigtableMemoryLayer API
// +k8s:openapi-gen=true
type BigtableMemoryLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableMemoryLayerSpec   `json:"spec,omitempty"`
	Status BigtableMemoryLayerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableMemoryLayerList contains a list of BigtableMemoryLayer
type BigtableMemoryLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableMemoryLayer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableMemoryLayer{}, &BigtableMemoryLayerList{})
}
