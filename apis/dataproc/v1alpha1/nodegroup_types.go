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
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocNodeGroup is the Schema for the DataprocNodeGroup API
// +k8s:openapi-gen=true
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
