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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataprocNodeGroupGVK = GroupVersion.WithKind("DataprocNodeGroup")

// DataprocNodeGroupSpec defines the desired state of DataprocNodeGroup
// +kcc:spec:proto=google.cloud.dataproc.v1.NodeGroup
type DataprocNodeGroupSpec struct {
	// Required. Node group roles.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.roles
	Roles []string `json:"roles,omitempty"`

	// Optional. The node group instance group configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.node_group_config
	NodeGroupConfig *InstanceGroupConfig `json:"nodeGroupConfig,omitempty"`

	// Optional. Node group labels.
	//
	//  * Label **keys** must consist of from 1 to 63 characters and conform to
	//    [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  * Label **values** can be empty. If specified, they must consist of from
	//    1 to 63 characters and conform to [RFC 1035]
	//    (https://www.ietf.org/rfc/rfc1035.txt).
	//  * The node group must have no more than 32 labels.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The cluster that this node group belongs to.
	ClusterRef *refsv1beta1.DataprocClusterRef `json:"clusterRef,omitempty"`

	*Parent `json:",inline"`

	// The DataprocNodeGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DiskConfig
type DiskConfig struct {
	// Optional. Type of the boot disk (default is "pd-standard").
	//  Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive),
	//  "pd-ssd" (Persistent Disk Solid State Drive),
	//  or "pd-standard" (Persistent Disk Hard Disk Drive).
	//  See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Optional. Size in GB of the boot disk (default is 500GB).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_size_gb
	BootDiskSizeGB *int32 `json:"bootDiskSizeGB,omitempty"`

	// Optional. Number of attached SSDs, from 0 to 8 (default is 0).
	//  If SSDs are not attached, the boot disk is used to store runtime logs and
	//  [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data.
	//  If one or more SSDs are attached, this runtime bulk
	//  data is spread across them, and the boot disk contains only basic
	//  config and installed binaries.
	//
	//  Note: Local SSD options may vary by machine type and number of vCPUs
	//  selected.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.num_local_ssds
	NumLocalSSDs *int32 `json:"numLocalSSDs,omitempty"`

	// Optional. Interface type of local SSDs (default is "scsi").
	//  Valid values: "scsi" (Small Computer System Interface),
	//  "nvme" (Non-Volatile Memory Express).
	//  See [local SSD
	//  performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.local_ssd_interface
	LocalSSDInterface *string `json:"localSSDInterface,omitempty"`

	// Optional. Indicates how many IOPS to provision for the disk. This sets the
	//  number of I/O operations per second that the disk can handle. Note: This
	//  field is only supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_iops
	BootDiskProvisionedIOPs *int64 `json:"bootDiskProvisionedIOPs,omitempty"`

	// Optional. Indicates how much throughput to provision for the disk. This
	//  sets the number of throughput mb per second that the disk can handle.
	//  Values must be greater than or equal to 1. Note: This field is only
	//  supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_throughput
	BootDiskProvisionedThroughput *int64 `json:"bootDiskProvisionedThroughput,omitempty"`
}

// DataprocNodeGroupStatus defines the config connector machine state of DataprocNodeGroup
type DataprocNodeGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocNodeGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocNodeGroupObservedState `json:"observedState,omitempty"`
}

// DataprocNodeGroupObservedState is the state of the DataprocNodeGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataproc.v1.NodeGroup
type DataprocNodeGroupObservedState struct {
	// Optional. The node group instance group configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.node_group_config
	NodeGroupConfig *InstanceGroupConfigObservedState `json:"nodeGroupConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocnodegroup;gcpdataprocnodegroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocNodeGroup is the Schema for the DataprocNodeGroup API
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
type DataprocNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocNodeGroupSpec   `json:"spec,omitempty"`
	Status DataprocNodeGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocNodeGroupList contains a list of DataprocNodeGroup
type DataprocNodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocNodeGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocNodeGroup{}, &DataprocNodeGroupList{})
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicyObservedState struct {
	// Output only. A list of instance selection results in the group.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_results
	InstanceSelectionResults []InstanceFlexibilityPolicy_InstanceSelectionResultObservedState `json:"instanceSelectionResults,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResultObservedState struct {
	// Output only. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Output only. Number of VM provisioned with the machine_type.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.vm_count
	VMCount *int32 `json:"vmCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicy struct {
	// Optional. Defines how the Group selects the provisioning model to ensure
	//  required reliability.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.provisioning_model_mix
	ProvisioningModelMix *InstanceFlexibilityPolicy_ProvisioningModelMix `json:"provisioningModelMix,omitempty"`

	// Optional. List of instance selection options that the group will use when
	//  creating new VMs.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_list
	InstanceSelectionList []InstanceFlexibilityPolicy_InstanceSelection `json:"instanceSelectionList,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection
type InstanceFlexibilityPolicy_InstanceSelection struct {
	// Optional. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.machine_types
	MachineTypes []string `json:"machineTypes,omitempty"`

	// Optional. Preference of this instance selection. Lower number means
	//  higher preference. Dataproc will first try to create a VM based on the
	//  machine-type with priority rank and fallback to next rank based on
	//  availability. Machine types and instance selections with the same
	//  priority have the same preference.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.rank
	Rank *int32 `json:"rank,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResult struct {
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix
type InstanceFlexibilityPolicy_ProvisioningModelMix struct {
	// Optional. The base capacity that will always use Standard VMs to avoid
	//  risk of more preemption than the minimum capacity you need. Dataproc will
	//  create only standard VMs until it reaches standard_capacity_base, then it
	//  will start using standard_capacity_percent_above_base to mix Spot with
	//  Standard VMs. eg. If 15 instances are requested and
	//  standard_capacity_base is 5, Dataproc will create 5 standard VMs and then
	//  start mixing spot and standard VMs for remaining 10 instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_base
	StandardCapacityBase *int32 `json:"standardCapacityBase,omitempty"`

	// Optional. The percentage of target capacity that should use Standard VM.
	//  The remaining percentage will use Spot VMs. The percentage applies only to
	//  the capacity above standard_capacity_base. eg. If 15 instances are
	//  requested and standard_capacity_base is 5 and
	//  standard_capacity_percent_above_base is 30, Dataproc will create 5
	//  standard VMs and then start mixing spot and standard VMs for remaining 10
	//  instances. The mix will be 30% standard and 70% spot, which means 3
	//  standard VMs and 7 spot VMs will be created. The total number of standard
	//  VMs created will be 8 and spot VMs will be 7.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_percent_above_base
	StandardCapacityPercentAboveBase *int32 `json:"standardCapacityPercentAboveBase,omitempty"`
}
