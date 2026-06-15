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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var HypercomputeClusterClusterGVK = GroupVersion.WithKind("HypercomputeClusterCluster")

// FilestoreInstanceRef is a reference to a GCP FilestoreInstance.
type FilestoreInstanceRef struct {
	// A reference to an externally managed FilestoreInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a FilestoreInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a FilestoreInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

// LustreRef is a reference to a GCP Managed Lustre instance.
type LustreRef struct {
	// A reference to an externally managed Managed Lustre instance.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NetworkReference
type NetworkReference struct {
	// Output only. Name of the network, in the format projects/{project}/global/networks/{network}.
	Network *string `json:"network,omitempty"`

	// Output only. Name of the subnetwork, in the format projects/{project}/regions/{region}/subnetworks/{subnetwork}.
	Subnetwork *string `json:"subnetwork,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.FilestoreReference
type FilestoreReference struct {
	// Output only. Name of the Filestore instance, in the format projects/{project}/locations/{location}/instances/{instance}.
	Filestore *string `json:"filestore,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.BucketReference
type BucketReference struct {
	// Output only. Name of the bucket.
	Bucket *string `json:"bucket,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.LustreReference
type LustreReference struct {
	// Output only. Name of the Managed Lustre instance, in the format projects/{project}/locations/{location}/instances/{instance}.
	Lustre *string `json:"lustre,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingNetworkConfig
type ExistingNetworkConfig struct {
	// Required. Immutable. The network to import.
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Required. Immutable. Particular subnetwork to use.
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewNetworkConfig
type NewNetworkConfig struct {
	// Required. Immutable. The network to create.
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Immutable. Description of the network. Maximum of 2048 characters.
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewFilestoreConfig
type NewFilestoreConfig struct {
	// Required. Immutable. Name of the Filestore instance to create.
	FilestoreRef *FilestoreInstanceRef `json:"filestoreRef,omitempty"`

	// Optional. Immutable. Description of the instance. Maximum of 2048 characters.
	Description *string `json:"description,omitempty"`

	// Required. Immutable. File system shares on the instance. Exactly one file share must be specified.
	FileShares []FileShareConfig `json:"fileShares,omitempty"`

	// Required. Immutable. Service tier to use for the instance.
	// +kubebuilder:validation:Enum=ZONAL;REGIONAL
	Tier *string `json:"tier,omitempty"`

	// Optional. Immutable. Access protocol to use for all file shares in the instance. Defaults to NFS V3 if not set.
	// +kubebuilder:validation:Enum=NFSV3;NFSV41
	Protocol *string `json:"protocol,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingFilestoreConfig
type ExistingFilestoreConfig struct {
	// Required. Immutable. Name of the Filestore instance to import.
	FilestoreRef *FilestoreInstanceRef `json:"filestoreRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewBucketConfig
type NewBucketConfig struct {
	// Required. Immutable. Name of the Cloud Storage bucket to create.
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Optional. Immutable. If set, indicates that the bucket should use Autoclass.
	Autoclass *GCSAutoclassConfig `json:"autoclass,omitempty"`

	// Optional. Immutable. If set, uses the provided storage class as the bucket's default storage class.
	// +kubebuilder:validation:Enum=STANDARD;NEARLINE;COLDLINE;ARCHIVE
	StorageClass *string `json:"storageClass,omitempty"`

	// Optional. Immutable. If set, indicates that the bucket should use hierarchical namespaces.
	HierarchicalNamespace *GCSHierarchicalNamespaceConfig `json:"hierarchicalNamespace,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingBucketConfig
type ExistingBucketConfig struct {
	// Required. Immutable. Name of the Cloud Storage bucket to import.
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingLustreConfig
type ExistingLustreConfig struct {
	// Required. Immutable. Name of the Managed Lustre instance to import.
	LustreRef *LustreRef `json:"lustreRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewReservedInstancesConfig
type NewReservedInstancesConfig struct {
	// Optional. Immutable. Name of the reservation from which VM instances should be created.
	ReservationRef *computev1beta1.ComputeReservationRef `json:"reservationRef,omitempty"`
}

// HypercomputeClusterClusterSpec defines the desired state of HypercomputeClusterCluster
// +kcc:spec:proto=google.cloud.hypercomputecluster.v1.Cluster
type HypercomputeClusterClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The HypercomputeClusterCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-provided description of the cluster.
	Description *string `json:"description,omitempty"`

	// Optional. Labels applied to the cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Network resources available to the cluster.
	NetworkResources map[string]NetworkResource `json:"networkResources,omitempty"`

	// Optional. Storage resources available to the cluster.
	StorageResources map[string]StorageResource `json:"storageResources,omitempty"`

	// Optional. Compute resources available to the cluster.
	ComputeResources map[string]ComputeResource `json:"computeResources,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs on the cluster.
	Orchestrator *Orchestrator `json:"orchestrator,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.hypercomputecluster.v1.Cluster
type ClusterObservedState struct {
	// Output only. Time that the cluster was originally created.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time that the cluster was most recently updated.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Indicates whether changes to the cluster are currently in
	//  flight. If this is `true`, then the current state might not match the
	//  cluster's intended state.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs on the cluster.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.orchestrator
	Orchestrator *OrchestratorObservedState `json:"orchestrator,omitempty"`
}

// HypercomputeClusterClusterStatus defines the config connector machine state of HypercomputeClusterCluster
type HypercomputeClusterClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the HypercomputeClusterCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ClusterObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcphypercomputeclustercluster;gcphypercomputeclusterclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// HypercomputeClusterCluster is the Schema for the HypercomputeClusterCluster API
// +k8s:openapi-gen=true
type HypercomputeClusterCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   HypercomputeClusterClusterSpec   `json:"spec,omitempty"`
	Status HypercomputeClusterClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// HypercomputeClusterClusterList contains a list of HypercomputeClusterCluster
type HypercomputeClusterClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HypercomputeClusterCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HypercomputeClusterCluster{}, &HypercomputeClusterClusterList{})
}
