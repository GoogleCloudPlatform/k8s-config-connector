// Copyright 2025 Google LLC
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

	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
)

var BigtableClusterGVK = GroupVersion.WithKind("BigtableCluster")

type BigtableClusterParent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// +required
	InstanceRef bigtablev1beta1.InstanceRef `json:"instanceRef"`
}

// BigtableClusterSpec defines the desired state of BigtableCluster
// +kcc:proto=google.bigtable.admin.v2.Cluster
type BigtableClusterSpec struct {
	BigtableClusterParent `json:",inline"`

	// The BigtableCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The location where this cluster's nodes and storage reside. For
	//  best performance, clients should be located as close as possible to this
	//  cluster. Currently only zones are supported, so values should be of the
	//  form `projects/{project}/locations/{zone}`.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.location
	Location *string `json:"location,omitempty"`

	// The number of nodes allocated to this cluster. More nodes enable higher
	//  throughput and more consistent performance.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.serve_nodes
	ServeNodes *int32 `json:"serveNodes,omitempty"`

	// Immutable. The node scaling factor of this cluster.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.node_scaling_factor
	NodeScalingFactor *string `json:"nodeScalingFactor,omitempty"`

	// Configuration for this cluster.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.cluster_config
	ClusterConfig *Cluster_ClusterConfig `json:"clusterConfig,omitempty"`

	// Immutable. The type of storage used by this cluster to serve its
	//  parent instance's tables, unless explicitly overridden.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.default_storage_type
	DefaultStorageType *string `json:"defaultStorageType,omitempty"`

	// Immutable. The encryption configuration for CMEK-protected clusters.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.encryption_config
	EncryptionConfig *Cluster_EncryptionConfig `json:"encryptionConfig,omitempty"`
}

// BigtableClusterStatus defines the config connector machine state of BigtableCluster
type BigtableClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigtableClusterObservedState `json:"observedState,omitempty"`
}

// BigtableClusterObservedState is the state of the BigtableCluster resource as most recently observed in GCP.
// +kcc:proto=google.bigtable.admin.v2.Cluster
type BigtableClusterObservedState struct {
	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtablecluster;gcpbigtableclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableCluster is the Schema for the BigtableCluster API
// +k8s:openapi-gen=true
type BigtableCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableClusterSpec   `json:"spec,omitempty"`
	Status BigtableClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableClusterList contains a list of BigtableCluster
type BigtableClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableCluster{}, &BigtableClusterList{})
}
