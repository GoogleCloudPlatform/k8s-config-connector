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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	filestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var HypercomputeClusterClusterGVK = GroupVersion.WithKind("HypercomputeClusterCluster")

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
	ObservedState *HypercomputeClusterClusterObservedState `json:"observedState,omitempty"`
}

// HypercomputeClusterClusterObservedState is the state of the HypercomputeClusterCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.hypercomputecluster.v1.Cluster
type HypercomputeClusterClusterObservedState struct {
	// Output only. Time that the cluster was originally created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time that the cluster was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Indicates whether changes to the cluster are currently in
	//  flight. If this is `true`, then the current state might not match the
	//  cluster's intended state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Optional. Network resources available to the cluster.
	NetworkResources map[string]NetworkResourceObservedState `json:"networkResources,omitempty"`

	// Optional. Storage resources available to the cluster.
	StorageResources map[string]StorageResourceObservedState `json:"storageResources,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs
	//  on the cluster.
	Orchestrator *OrchestratorObservedState `json:"orchestrator,omitempty"`
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

// +kcc:proto=google.cloud.hypercomputecluster.v1.NetworkResource
type NetworkResource struct {
	// Immutable. Configuration for this network resource, which describes how it
	//  should be created or imported.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NetworkResource.config
	Config *NetworkResourceConfig `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NetworkResource
type NetworkResourceObservedState struct {
	// Reference to a network in Google Compute Engine.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NetworkResource.network
	Network *NetworkReferenceObservedState `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NetworkReference
type NetworkReferenceObservedState struct {
	// Output only. Name of the network.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NetworkReference.network
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Output only. Name of the particular subnetwork being used by the cluster.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NetworkReference.subnetwork
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NetworkResourceConfig
type NetworkResourceConfig struct {
	// Optional. Immutable. If set, indicates that a new network should be
	//  created.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NetworkResourceConfig.new_network
	NewNetwork *NewNetworkConfig `json:"newNetwork,omitempty"`

	// Optional. Immutable. If set, indicates that an existing network should be
	//  imported.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NetworkResourceConfig.existing_network
	ExistingNetwork *ExistingNetworkConfig `json:"existingNetwork,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewNetworkConfig
type NewNetworkConfig struct {
	// Required. Immutable. Name of the network to create.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewNetworkConfig.network
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Immutable. Description of the network. Maximum of 2048
	//  characters.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewNetworkConfig.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingNetworkConfig
type ExistingNetworkConfig struct {
	// Required. Immutable. Name of the network to import.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ExistingNetworkConfig.network
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Required. Immutable. Particular subnetwork to use.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ExistingNetworkConfig.subnetwork
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.StorageResource
type StorageResource struct {
	// Required. Immutable. Configuration for this storage resource, which
	// describes how it should be created or imported.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResource.config
	Config *StorageResourceConfig `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.StorageResource
type StorageResourceObservedState struct {
	// Reference to a Filestore instance.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResource.filestore
	Filestore *FilestoreReferenceObservedState `json:"filestore,omitempty"`

	// Reference to a Google Cloud Storage bucket.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResource.bucket
	Bucket *BucketReferenceObservedState `json:"bucket,omitempty"`

	// Reference to a Managed Lustre instance.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResource.lustre
	Lustre *LustreReferenceObservedState `json:"lustre,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.FilestoreReference
type FilestoreReferenceObservedState struct {
	// Output only. Name of the Filestore instance.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.FilestoreReference.filestore
	FilestoreRef *filestorev1beta1.FilestoreInstanceRef `json:"filestoreRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.BucketReference
type BucketReferenceObservedState struct {
	// Output only. Name of the bucket.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.BucketReference.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.LustreReference
type LustreReferenceObservedState struct {
	// Output only. Name of the Managed Lustre instance.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.LustreReference.lustre
	Lustre *string `json:"lustre,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.StorageResourceConfig
type StorageResourceConfig struct {
	// Optional. Immutable. If set, indicates that a new Filestore instance
	//  should be created.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResourceConfig.new_filestore
	NewFilestore *NewFilestoreConfig `json:"newFilestore,omitempty"`

	// Optional. Immutable. If set, indicates that an existing Filestore
	//  instance should be imported.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResourceConfig.existing_filestore
	ExistingFilestore *ExistingFilestoreConfig `json:"existingFilestore,omitempty"`

	// Optional. Immutable. If set, indicates that a new Cloud Storage bucket
	//  should be created.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResourceConfig.new_bucket
	NewBucket *NewBucketConfig `json:"newBucket,omitempty"`

	// Optional. Immutable. If set, indicates that an existing Cloud Storage
	//  bucket should be imported.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResourceConfig.existing_bucket
	ExistingBucket *ExistingBucketConfig `json:"existingBucket,omitempty"`

	// Optional. Immutable. If set, indicates that a new Managed Lustre instance
	//  should be created.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResourceConfig.new_lustre
	NewLustre *NewLustreConfig `json:"newLustre,omitempty"`

	// Optional. Immutable. If set, indicates that an existing Managed Lustre
	//  instance should be imported.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageResourceConfig.existing_lustre
	ExistingLustre *ExistingLustreConfig `json:"existingLustre,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewFilestoreConfig
type NewFilestoreConfig struct {
	// Required. Immutable. Name of the Filestore instance to create.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFilestoreConfig.filestore
	FilestoreRef *filestorev1beta1.FilestoreInstanceRef `json:"filestoreRef,omitempty"`

	// Optional. Immutable. Description of the instance. Maximum of 2048
	//  characters.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFilestoreConfig.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. File system shares on the instance. Exactly one file
	//  share must be specified.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFilestoreConfig.file_shares
	FileShares []FileShareConfig `json:"fileShares,omitempty"`

	// Required. Immutable. Service tier to use for the instance.
	// +kubebuilder:validation:Enum=ZONAL;REGIONAL
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFilestoreConfig.tier
	Tier *string `json:"tier,omitempty"`

	// Optional. Immutable. Access protocol to use for all file shares in the
	//  instance. Defaults to NFS V3 if not set.
	// +kubebuilder:validation:Enum=NFSV3;NFSV41
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFilestoreConfig.protocol
	Protocol *string `json:"protocol,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingFilestoreConfig
type ExistingFilestoreConfig struct {
	// Required. Immutable. Name of the Filestore instance to import.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ExistingFilestoreConfig.filestore
	FilestoreRef *filestorev1beta1.FilestoreInstanceRef `json:"filestoreRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewBucketConfig
type NewBucketConfig struct {
	// Optional. Immutable. If set, indicates that the bucket should use
	//  [Autoclass](https://cloud.google.com/storage/docs/autoclass).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewBucketConfig.autoclass
	Autoclass *GCSAutoclassConfig `json:"autoclass,omitempty"`

	// Optional. Immutable. If set, uses the provided storage class as the
	//  bucket's default storage class.
	// +kubebuilder:validation:Enum=STANDARD;NEARLINE;COLDLINE;ARCHIVE
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewBucketConfig.storage_class
	StorageClass *string `json:"storageClass,omitempty"`

	// Required. Immutable. Name of the Cloud Storage bucket to create.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewBucketConfig.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Optional. Immutable. If set, indicates that the bucket should use
	//  [hierarchical
	//  namespaces](https://cloud.google.com/storage/docs/hns-overview).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewBucketConfig.hierarchical_namespace
	HierarchicalNamespace *GCSHierarchicalNamespaceConfig `json:"hierarchicalNamespace,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingBucketConfig
type ExistingBucketConfig struct {
	// Required. Immutable. Name of the Cloud Storage bucket to import.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ExistingBucketConfig.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewLustreConfig
type NewLustreConfig struct {
	// Required. Immutable. Name of the Managed Lustre instance to create.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewLustreConfig.lustre
	Lustre *string `json:"lustre,omitempty"`

	// Optional. Immutable. Description of the Managed Lustre instance. Maximum of
	//  2048 characters.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewLustreConfig.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. Filesystem name for this instance. This name is used
	//  by client-side tools, including when mounting the instance. Must be 8
	//  characters or less and can only contain letters and numbers.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewLustreConfig.filesystem
	Filesystem *string `json:"filesystem,omitempty"`

	// Required. Immutable. Storage capacity of the instance in gibibytes (GiB).
	//  Allowed values are between 18000 and 7632000.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewLustreConfig.capacity_gb
	CapacityGB *int64 `json:"capacityGB,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingLustreConfig
type ExistingLustreConfig struct {
	// Required. Immutable. Name of the Managed Lustre instance to import.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ExistingLustreConfig.lustre
	Lustre *string `json:"lustre,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ComputeResource
type ComputeResource struct {
	// Required. Immutable. Configuration for this compute resource, which
	//  describes how it should be created at runtime.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeResource.config
	Config *ComputeResourceConfig `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ComputeResourceConfig
type ComputeResourceConfig struct {
	// Optional. Immutable. If set, indicates that this resource should use
	//  on-demand VMs.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeResourceConfig.new_on_demand_instances
	NewOnDemandInstances *NewOnDemandInstancesConfig `json:"newOnDemandInstances,omitempty"`

	// Optional. Immutable. If set, indicates that this resource should use spot
	//  VMs.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeResourceConfig.new_spot_instances
	NewSpotInstances *NewSpotInstancesConfig `json:"newSpotInstances,omitempty"`

	// Optional. Immutable. If set, indicates that this resource should use
	//  reserved VMs.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeResourceConfig.new_reserved_instances
	NewReservedInstances *NewReservedInstancesConfig `json:"newReservedInstances,omitempty"`

	// Optional. Immutable. If set, indicates that this resource should use
	//  flex-start VMs.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeResourceConfig.new_flex_start_instances
	NewFlexStartInstances *NewFlexStartInstancesConfig `json:"newFlexStartInstances,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewOnDemandInstancesConfig
type NewOnDemandInstancesConfig struct {
	// Required. Immutable. Name of the zone in which VM instances should run,
	//  e.g., `us-central1-a`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewOnDemandInstancesConfig.zone
	Zone *string `json:"zone,omitempty"`

	// Required. Immutable. Name of the Compute Engine machine type to use.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewOnDemandInstancesConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewSpotInstancesConfig
type NewSpotInstancesConfig struct {
	// Required. Immutable. Name of the zone in which VM instances should run,
	//  e.g., `us-central1-a`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewSpotInstancesConfig.zone
	Zone *string `json:"zone,omitempty"`

	// Required. Immutable. Name of the Compute Engine machine type to use.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewSpotInstancesConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. Termination action for the instance. If not specified, Compute
	//  Engine sets the termination action to DELETE.
	// +kubebuilder:validation:Enum=STOP;DELETE
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewSpotInstancesConfig.termination_action
	TerminationAction *string `json:"terminationAction,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewReservedInstancesConfig
type NewReservedInstancesConfig struct {
	// Optional. Immutable. Name of the reservation from which VM instances
	//  should be created.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewReservedInstancesConfig.reservation
	ReservationRef *computev1beta1.ComputeReservationRef `json:"reservationRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewFlexStartInstancesConfig
type NewFlexStartInstancesConfig struct {
	// Required. Immutable. Name of the zone in which VM instances should run,
	//  e.g., `us-central1-a`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFlexStartInstancesConfig.zone
	Zone *string `json:"zone,omitempty"`

	// Required. Immutable. Name of the Compute Engine machine type to use.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFlexStartInstancesConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Required. Immutable. Specifies the time limit for created instances.
	//  Instances will be terminated at the end of this duration.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.NewFlexStartInstancesConfig.max_duration
	MaxDuration *string `json:"maxDuration,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.Orchestrator
type Orchestrator struct {
	// Optional. If set, indicates that the cluster should use Slurm as the
	//  orchestrator.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Orchestrator.slurm
	Slurm *SlurmOrchestrator `json:"slurm,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.Orchestrator
type OrchestratorObservedState struct {
	// Optional. If set, indicates that the cluster should use Slurm as the
	//  orchestrator.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Orchestrator.slurm
	Slurm *SlurmOrchestratorObservedState `json:"slurm,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmOrchestrator
type SlurmOrchestrator struct {
	// Required. Configuration for login nodes, which allow users to access the
	//  cluster over SSH.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmOrchestrator.login_nodes
	LoginNodes *SlurmLoginNodes `json:"loginNodes,omitempty"`

	// Optional. Compute resource configuration for the Slurm nodesets in your
	//  cluster. If not specified, the cluster won't create any nodes.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmOrchestrator.node_sets
	NodeSets []SlurmNodeSet `json:"nodeSets,omitempty"`

	// Optional. Configuration for the Slurm partitions in your cluster. Each
	//  partition can contain one or more nodesets, and you can submit separate
	//  jobs on each partition.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmOrchestrator.partitions
	Partitions []SlurmPartition `json:"partitions,omitempty"`

	// Optional. Default partition to use for submitted jobs that do not
	//  explicitly specify a partition. Required if and only if there is more than
	//  one partition, in which case it must match the id of one of the partitions.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmOrchestrator.default_partition
	DefaultPartition *string `json:"defaultPartition,omitempty"`

	// Optional. Slurm prolog scripts, which will be executed by compute nodes
	//  before a node begins running a new job. Values must not be empty.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmOrchestrator.prolog_bash_scripts
	PrologBashScripts []string `json:"prologBashScripts,omitempty"`

	// Optional. Slurm epilog scripts, which will be executed by compute nodes
	//  whenever a node finishes running a job. Values must not be empty.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmOrchestrator.epilog_bash_scripts
	EpilogBashScripts []string `json:"epilogBashScripts,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmOrchestrator
type SlurmOrchestratorObservedState struct {
	// Required. Configuration for login nodes, which allow users to access the
	//  cluster over SSH.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmOrchestrator.login_nodes
	LoginNodes *SlurmLoginNodesObservedState `json:"loginNodes,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmLoginNodes
type SlurmLoginNodes struct {
	// Required. Number of login node instances to create.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.count
	Count *int64 `json:"count,omitempty"`

	// Required. Name of the zone in which login nodes should run, e.g.,
	//  `us-central1-a`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.zone
	Zone *string `json:"zone,omitempty"`

	// Required. Name of the Compute Engine machine type to use for login nodes.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. Startup script to be run on each login node instance. Max 256KB.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.startup_script
	StartupScript *string `json:"startupScript,omitempty"`

	// Optional. Whether OS Login should be enabled on login node instances.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.enable_os_login
	EnableOSLogin *bool `json:"enableOSLogin,omitempty"`

	// Optional. Whether login node instances should be assigned external IP
	//  addresses.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.enable_public_ips
	EnablePublicIPs *bool `json:"enablePublicIPs,omitempty"`

	// Optional. Labels that should be applied to each login node instance.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. How storage resources should be mounted on each login node.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.storage_configs
	StorageConfigs []StorageConfig `json:"storageConfigs,omitempty"`

	// Optional. Boot disk for the login node.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.boot_disk
	BootDisk *BootDisk `json:"bootDisk,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmLoginNodes
type SlurmLoginNodesObservedState struct {
	// Output only. Information about the login node instances that were created
	//  in Compute Engine.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.instances
	Instances []ComputeInstanceObservedState `json:"instances,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmNodeSet
type SlurmNodeSet struct {
	// Optional. If set, indicates that the nodeset should be backed by Compute
	//  Engine instances.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmNodeSet.compute_instance
	ComputeInstance *ComputeInstanceSlurmNodeSet `json:"computeInstance,omitempty"`

	// Required. Identifier for the nodeset, which allows it to be referenced by
	//  partitions.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmNodeSet.id
	ID *string `json:"id,omitempty"`

	// Optional. ID of the compute resource on which this nodeset will run.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmNodeSet.compute_id
	ComputeID *string `json:"computeID,omitempty"`

	// Optional. How storage resources should be mounted on each compute node.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmNodeSet.storage_configs
	StorageConfigs []StorageConfig `json:"storageConfigs,omitempty"`

	// Optional. Number of nodes to be statically created for this nodeset.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmNodeSet.static_node_count
	StaticNodeCount *int64 `json:"staticNodeCount,omitempty"`

	// Optional. Controls how many additional nodes a cluster can bring online to
	//  handle workloads.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmNodeSet.max_dynamic_node_count
	MaxDynamicNodeCount *int64 `json:"maxDynamicNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ComputeInstanceSlurmNodeSet
type ComputeInstanceSlurmNodeSet struct {
	// Optional. Startup script to be run on each VM instance in the nodeset. Max
	//  256KB.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeInstanceSlurmNodeSet.startup_script
	StartupScript *string `json:"startupScript,omitempty"`

	// Optional. Labels that should be applied to each VM instance in the
	//  nodeset.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeInstanceSlurmNodeSet.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Boot disk for the compute instance
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeInstanceSlurmNodeSet.boot_disk
	BootDisk *BootDisk `json:"bootDisk,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.BootDisk
type BootDisk struct {
	// Required. Immutable. Persistent disk type.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.BootDisk.type
	Type *string `json:"type,omitempty"`

	// Required. Immutable. Size of the disk in gigabytes. Must be at least 10GB.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.BootDisk.size_gb
	SizeGB *int64 `json:"sizeGB,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.StorageConfig
type StorageConfig struct {
	// Required. ID of the storage resource to mount.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageConfig.id
	ID *string `json:"id,omitempty"`

	// Required. A directory inside the VM instance's file system where the
	//  storage resource should be mounted (e.g., `/mnt/share`).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.StorageConfig.local_mount
	LocalMount *string `json:"localMount,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmPartition
type SlurmPartition struct {
	// Required. ID of the partition, which is how users will identify it.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmPartition.id
	ID *string `json:"id,omitempty"`

	// Required. IDs of the nodesets that make up this partition.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmPartition.node_set_ids
	NodeSetIDs []string `json:"nodeSetIDs,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.FileShareConfig
type FileShareConfig struct {
	// Required. Size of the filestore in GB.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.FileShareConfig.capacity_gb
	CapacityGB *int64 `json:"capacityGB,omitempty"`

	// Required. Filestore share location
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.FileShareConfig.file_share
	FileShare *string `json:"fileShare,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.GcsAutoclassConfig
type GCSAutoclassConfig struct {
	// Required. Enables Auto-class feature.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.GcsAutoclassConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. Terminal storage class of the autoclass bucket
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.GcsAutoclassConfig.terminal_storage_class
	TerminalStorageClass *string `json:"terminalStorageClass,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.GcsHierarchicalNamespaceConfig
type GCSHierarchicalNamespaceConfig struct {
	// Required. Enables hierarchical namespace setup for the bucket.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.GcsHierarchicalNamespaceConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ComputeInstance
type ComputeInstanceObservedState struct {
	// Output only. Name of the VM instance.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeInstance.instance
	InstanceRef *computev1beta1.InstanceRef `json:"instanceRef,omitempty"`
}
