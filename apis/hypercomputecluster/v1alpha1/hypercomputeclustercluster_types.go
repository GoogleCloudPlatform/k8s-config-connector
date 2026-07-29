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
	NetworkResources map[string]*NetworkResource `json:"networkResources,omitempty"`

	// Optional. Storage resources available to the cluster.
	StorageResources map[string]*StorageResource `json:"storageResources,omitempty"`

	// Optional. Compute resources available to the cluster.
	ComputeResources map[string]*ComputeResource `json:"computeResources,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs on the cluster.
	Orchestrator *Orchestrator `json:"orchestrator,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewNetworkConfig
type NewNetworkConfig struct {
	// Required. Immutable. Reference to the network to create.
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Immutable. Description of the network. Maximum of 2048 characters.
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingNetworkConfig
type ExistingNetworkConfig struct {
	// Required. Immutable. Reference to the network to import.
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Required. Immutable. Particular subnetwork to use.
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewFilestoreConfig
type NewFilestoreConfig struct {
	// Required. Immutable. Reference to the Filestore instance to create.
	FilestoreRef *filestorev1beta1.FilestoreInstanceRef `json:"filestoreRef,omitempty"`

	// Optional. Immutable. Description of the instance. Maximum of 2048 characters.
	Description *string `json:"description,omitempty"`

	// Required. Immutable. File system shares on the instance. Exactly one file share must be specified.
	FileShares []FileShareConfig `json:"fileShares,omitempty"`

	// Required. Immutable. Service tier to use for the instance.
	Tier *string `json:"tier,omitempty"`

	// Optional. Immutable. Access protocol to use for all file shares in the instance. Defaults to NFS V3 if not set.
	Protocol *string `json:"protocol,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingFilestoreConfig
type ExistingFilestoreConfig struct {
	// Required. Immutable. Reference to the Filestore instance to import.
	FilestoreRef *filestorev1beta1.FilestoreInstanceRef `json:"filestoreRef,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NewBucketConfig
type NewBucketConfig struct {
	// Optional. Immutable. If set, indicates that the bucket should use Autoclass.
	Autoclass *GCSAutoclassConfig `json:"autoclass,omitempty"`

	// Optional. Immutable. If set, uses the provided storage class as the bucket's default storage class.
	StorageClass *string `json:"storageClass,omitempty"`

	// Required. Immutable. Reference to the Cloud Storage bucket to create.
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Optional. Immutable. If set, indicates that the bucket should use hierarchical namespaces.
	HierarchicalNamespace *GCSHierarchicalNamespaceConfig `json:"hierarchicalNamespace,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ExistingBucketConfig
type ExistingBucketConfig struct {
	// Required. Immutable. Reference to the Cloud Storage bucket to import.
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`
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
	// flight. If this is `true`, then the current state might not match the
	// cluster's intended state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs
	// on the cluster.
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

// +kcc:proto=google.cloud.hypercomputecluster.v1.BucketReference
type BucketReference struct {
	// Reference to a Google Cloud Storage bucket.
	Bucket *string `json:"bucket,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.FilestoreReference
type FilestoreReference struct {
	// Reference to a Filestore instance.
	Filestore *string `json:"filestore,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NetworkReference
type NetworkReference struct {
	// Reference to a Compute Network.
	Network *string `json:"network,omitempty"`

	// Reference to a Compute Subnetwork.
	Subnetwork *string `json:"subnetwork,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.LustreReference
type LustreReference struct {
	// Reference to a Lustre instance.
	Lustre *string `json:"lustre,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmLoginNodes
type SlurmLoginNodes struct {
	// Required. Number of login node instances to create.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.count
	Count *int64 `json:"count,omitempty"`

	// Required. Name of the zone in which login nodes should run, e.g.,
	//  `us-central1-a`. Must be in the same region as the cluster, and must match
	//  the zone of any other resources specified in the cluster.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.zone
	Zone *string `json:"zone,omitempty"`

	// Required. Name of the Compute Engine machine type to use for login nodes, e.g. `n2-standard-2`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. Startup script to be run on each login node instance. Max 256KB.
	//  The script must complete within the system-defined default timeout of 5
	//  minutes. For tasks that require more time, consider running them in the
	//  background using methods such as `&` or `nohup`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.startup_script
	StartupScript *string `json:"startupScript,omitempty"`

	// Optional. Whether OS Login should be enabled on login node instances.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.enable_os_login
	EnableOSLogin *bool `json:"enableOSLogin,omitempty"`

	// Optional. Whether login node instances should be assigned external IP addresses.
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

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmPartition
type SlurmPartition struct {
	// Required. ID of the partition, which is how users will identify it. Must
	//  conform to RFC-1034 (lower-case, alphanumeric, and at most 63 characters).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmPartition.id
	ID *string `json:"id,omitempty"`

	// Required. IDs of the nodesets that make up this partition. Values must
	//  match SlurmNodeSet.id.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmPartition.node_set_ids
	NodeSetIDs []string `json:"nodeSetIDs,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.hypercomputecluster.v1.SlurmLoginNodes
type SlurmLoginNodesObservedState struct {
	// Output only. Information about the login node instances that were created in Compute Engine.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.instances
	Instances []ComputeInstanceObservedState `json:"instances,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.hypercomputecluster.v1.ComputeInstance
type ComputeInstanceObservedState struct {
	// Output only. Name of the VM instance, in the format projects/{project}/zones/{zone}/instances/{instance}.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeInstance.instance
	Instance *string `json:"instance,omitempty"`
}
