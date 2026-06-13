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

package v1beta1

import (
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FilestoreInstanceGVK = GroupVersion.WithKind("FilestoreInstance")

// +kcc:proto=google.cloud.filestore.v1.FileShareConfig
type InstanceFileShares struct {
	/* File share capacity in gigabytes (GB). Cloud Filestore defines 1 GB as 1024^3 bytes. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.capacity_gb
	CapacityGb *int64 `json:"capacityGb,omitempty"`

	/* The name of the file share (must be 16 characters or less). */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.name
	Name *string `json:"name,omitempty"`

	/* Nfs Export Options. There is a limit of 10 export options per file share. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.nfs_export_options
	NfsExportOptions []InstanceNfsExportOptions `json:"nfsExportOptions,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.source_backup
	SourceBackupRef *FilestoreBackupRef `json:"sourceBackupRef,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1.NetworkConfig
type InstanceNetworks struct {
	/* Immutable. Output only. IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or IPv6 addresses in the format `{block1}:{block2}:{block3}:{block4}:{block5}:{block6}:{block7}:{block8}`. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.ip_addresses
	IpAddresses []string `json:"ipAddresses,omitempty"`

	/* Immutable. Internet protocol versions for which the instance has IP addresses assigned. For this version, only MODE_IPV4 is supported. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.modes
	Modes []string `json:"modes,omitempty"`

	/* Immutable. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.network
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	/* Immutable. A /29 CIDR block in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/) that identifies the range of IP addresses reserved for this instance. For example, 10.0.0.0/29 or 192.168.0.0/29. The range you specify can't overlap with either existing subnets or assigned IP address ranges for other Cloud Filestore instances in the selected VPC network. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.reserved_ip_range
	ReservedIPRange *string `json:"reservedIPRange,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1.NfsExportOptions
type InstanceNfsExportOptions struct {
	/* Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE. Possible values: ACCESS_MODE_UNSPECIFIED, READ_ONLY, READ_WRITE */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.access_mode
	AccessMode *string `json:"accessMode,omitempty"`

	/* An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.anon_gid
	AnonGid *int64 `json:"anonGid,omitempty"`

	/* An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.anon_uid
	AnonUid *int64 `json:"anonUid,omitempty"`

	/* List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions. */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.ip_ranges
	IpRanges []string `json:"ipRanges,omitempty"`

	/* Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH. Possible values: SQUASH_MODE_UNSPECIFIED, NO_ROOT_SQUASH, ROOT_SQUASH */
	// +optional
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.squash_mode
	SquashMode *string `json:"squashMode,omitempty"`
}

// FilestoreInstanceSpec defines the desired state of FilestoreInstance
// +kcc:spec:proto=google.cloud.filestore.v1.Instance
type FilestoreInstanceSpec struct {
	/* The description of the instance (2048 characters or less). */
	// +optional
	Description *string `json:"description,omitempty"`

	/* File system shares on the instance. For this version, only a single file share is supported. */
	// +optional
	FileShares []InstanceFileShares `json:"fileShares,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. VPC networks to which the instance is connected. For this version, only a single network is supported. */
	// +optional
	Networks []InstanceNetworks `json:"networks,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The service tier of the instance. Possible values: TIER_UNSPECIFIED, STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ENTERPRISE */
	// +optional
	Tier *string `json:"tier,omitempty"`
}

// FilestoreInstanceStatus defines the config connector machine state of FilestoreInstance
type FilestoreInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   FilestoreInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Output only. The time when the instance was created. */
	// +optional
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	/* Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. The instance state. Possible values: STATE_UNSPECIFIED, CREATING, READY, REPAIRING, DELETING, ERROR */
	// +optional
	State *string `json:"state,omitempty"`

	/* Output only. Additional information about the instance state, if available. */
	// +optional
	StatusMessage *string `json:"statusMessage,omitempty"`

	/* A unique specifier for the FilestoreInstance resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfilestoreinstance;gcpfilestoreinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FilestoreInstance is the Schema for the FilestoreInstance API
// +k8s:openapi-gen=true
type FilestoreInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FilestoreInstanceSpec   `json:"spec,omitempty"`
	Status FilestoreInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FilestoreInstanceList contains a list of FilestoreInstance
type FilestoreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FilestoreInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FilestoreInstance{}, &FilestoreInstanceList{})
}
