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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var HypercomputeClusterClusterGVK = GroupVersion.WithKind("HypercomputeClusterCluster")

// HypercomputeClusterClusterSpec defines the desired state of HypercomputeClusterCluster
// +kcc:spec:proto=google.cloud.hypercomputecluster.v1.Cluster
type HypercomputeClusterClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The HypercomputeClusterCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-provided description of the cluster.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.description
	Description *string `json:"description,omitempty"`

	// Optional.
	//  [Labels](https://cloud.google.com/compute/docs/labeling-resources) applied
	//  to the cluster. Labels can be used to organize clusters and to filter them
	//  in queries.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Network resources available to the cluster. Must contain at most
	//  one value. Keys specify the ID of the network resource by which it can be
	//  referenced elsewhere, and must conform to
	//  [RFC-1034](https://datatracker.ietf.org/doc/html/rfc1034) (lower-case,
	//  alphanumeric, and at most 63 characters).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.network_resources
	NetworkResources map[string]NetworkResource `json:"networkResources,omitempty"`

	// Optional. Storage resources available to the cluster. Keys specify the ID
	//  of the storage resource by which it can be referenced elsewhere, and must
	//  conform to [RFC-1034](https://datatracker.ietf.org/doc/html/rfc1034)
	//  (lower-case, alphanumeric, and at most 63 characters).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.storage_resources
	StorageResources map[string]StorageResource `json:"storageResources,omitempty"`

	// Optional. Compute resources available to the cluster. Keys specify the ID
	//  of the compute resource by which it can be referenced elsewhere, and must
	//  conform to [RFC-1034](https://datatracker.ietf.org/doc/html/rfc1034)
	//  (lower-case, alphanumeric, and at most 63 characters).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.compute_resources
	ComputeResources map[string]ComputeResource `json:"computeResources,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs
	//  on the cluster.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.orchestrator
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

	// Optional. Orchestrator that is responsible for scheduling and running jobs
	//  on the cluster.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.Cluster.orchestrator
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
// +kubebuilder:validation:XPreserveUnknownFields
type BucketReference struct {
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.ComputeInstance
// +kubebuilder:validation:XPreserveUnknownFields
type ComputeInstance struct {
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.FilestoreReference
// +kubebuilder:validation:XPreserveUnknownFields
type FilestoreReference struct {
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.LustreReference
// +kubebuilder:validation:XPreserveUnknownFields
type LustreReference struct {
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.NetworkReference
// +kubebuilder:validation:XPreserveUnknownFields
type NetworkReference struct {
}

// +kcc:observedstate:proto=google.cloud.hypercomputecluster.v1.ComputeInstance
type ComputeInstanceObservedState struct {
	// Output only. Name of the VM instance, in the format
	//  `projects/{project}/zones/{zone}/instances/{instance}`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.ComputeInstance.instance
	Instance *string `json:"instance,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.BootDisk
type BootDisk struct {
	// Required. Immutable. [Persistent disk
	//  type](https://cloud.google.com/compute/docs/disks#disk-types), in the
	//  format `projects/{project}/zones/{zone}/diskTypes/{disk_type}`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.BootDisk.type
	TypeRef *refsv1beta1.ComputeDiskTypeRef `json:"typeRef,omitempty"`

	// Required. Immutable. Size of the disk in gigabytes. Must be at least 10GB.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.BootDisk.size_gb
	SizeGB *int64 `json:"sizeGB,omitempty"`
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

	// Required. Name of the Compute Engine [machine
	//  type](https://cloud.google.com/compute/docs/machine-resource) to use for
	//  login nodes, e.g. `n2-standard-2`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. [Startup
	//  script](https://cloud.google.com/compute/docs/instances/startup-scripts/linux)
	//  to be run on each login node instance. Max 256KB.
	//  The script must complete within the system-defined default timeout of 5
	//  minutes. For tasks that require more time, consider running them in the
	//  background using methods such as `&` or `nohup`.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.startup_script
	StartupScript *string `json:"startupScript,omitempty"`

	// Optional. Whether [OS Login](https://cloud.google.com/compute/docs/oslogin)
	//  should be enabled on login node instances.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.enable_os_login
	EnableOSLogin *bool `json:"enableOSLogin,omitempty"`

	// Optional. Whether login node instances should be assigned [external IP
	//  addresses](https://cloud.google.com/compute/docs/ip-addresses#externaladdresses).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.enable_public_ips
	EnablePublicIPs *bool `json:"enablePublicIPs,omitempty"`

	// Optional.
	//  [Labels](https://cloud.google.com/compute/docs/labeling-resources) that
	//  should be applied to each login node instance.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. How [storage
	//  resources][google.cloud.hypercomputecluster.v1.StorageResource] should be
	//  mounted on each login node.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.storage_configs
	StorageConfigs []StorageConfig `json:"storageConfigs,omitempty"`

	// Optional. Boot disk for the login node.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.boot_disk
	BootDisk *BootDisk `json:"bootDisk,omitempty"`
}

// +kcc:proto=google.cloud.hypercomputecluster.v1.SlurmPartition
type SlurmPartition struct {
	// Required. ID of the partition, which is how users will identify it. Must
	//  conform to [RFC-1034](https://datatracker.ietf.org/doc/html/rfc1034)
	//  (lower-case, alphanumeric, and at most 63 characters).
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmPartition.id
	ID *string `json:"id,omitempty"`

	// Required. IDs of the nodesets that make up this partition. Values must
	//  match
	//  [SlurmNodeSet.id][google.cloud.hypercomputecluster.v1.SlurmNodeSet.id].
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmPartition.node_set_ids
	NodeSetIDs []string `json:"nodeSetIDs,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.hypercomputecluster.v1.SlurmLoginNodes
type SlurmLoginNodesObservedState struct {
	// Output only. Information about the login node instances that were created
	//  in Compute Engine.
	// +kcc:proto:field=google.cloud.hypercomputecluster.v1.SlurmLoginNodes.instances
	Instances []ComputeInstanceObservedState `json:"instances,omitempty"`
}
