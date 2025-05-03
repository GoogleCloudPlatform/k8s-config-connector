// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +generated:types
// krm.group: notebooks.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.notebooks.v1
// resource: NotebookInstance:Instance

package v1beta1

// +kcc:proto=google.cloud.notebooks.v1.ContainerImage
type ContainerImage struct {
	// Required. The path to the container image repository. For example:
	//  `gcr.io/{project_id}/{image_name}`
	// +kcc:proto:field=google.cloud.notebooks.v1.ContainerImage.repository
	Repository *string `json:"repository,omitempty"`

	// The tag of the container image. If not specified, this defaults
	//  to the latest tag.
	// +kcc:proto:field=google.cloud.notebooks.v1.ContainerImage.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Instance.AcceleratorConfig
type Instance_AcceleratorConfig struct {
	// Type of this accelerator.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.AcceleratorConfig.type
	Type *string `json:"type,omitempty"`

	// Count of cores of this accelerator.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.AcceleratorConfig.core_count
	CoreCount *int64 `json:"coreCount,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Instance.Disk
type Instance_Disk struct {
	// Indicates whether the disk will be auto-deleted when the instance is
	//  deleted (but not when the disk is detached from the instance).
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.auto_delete
	AutoDelete *bool `json:"autoDelete,omitempty"`

	// Indicates that this is a boot disk. The virtual machine will use the
	//  first partition of the disk for its root filesystem.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.boot
	Boot *bool `json:"boot,omitempty"`

	// Indicates a unique device name of your choice that is reflected into the
	//  `/dev/disk/by-id/google-*` tree of a Linux operating system running
	//  within the instance. This name can be used to reference the device for
	//  mounting, resizing, and so on, from within the instance.
	//
	//  If not specified, the server chooses a default device name to apply to
	//  this disk, in the form persistent-disk-x, where x is a number assigned by
	//  Google Compute Engine.This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.device_name
	DeviceName *string `json:"deviceName,omitempty"`

	// Indicates the size of the disk in base-2 GB.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Indicates a list of features to enable on the guest operating system.
	//  Applicable only for bootable images. Read  Enabling guest operating
	//  system features to see a list of available options.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.guest_os_features
	GuestOSFeatures []Instance_Disk_GuestOSFeature `json:"guestOSFeatures,omitempty"`

	// A zero-based index to this disk, where 0 is reserved for the
	//  boot disk. If you have many disks attached to an instance, each disk
	//  would have a unique index number.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.index
	Index *int64 `json:"index,omitempty"`

	// Indicates the disk interface to use for attaching this disk, which is
	//  either SCSI or NVME. The default is SCSI. Persistent disks must always
	//  use SCSI and the request will fail if you attempt to attach a persistent
	//  disk in any other format than SCSI. Local SSDs can use either NVME or
	//  SCSI. For performance characteristics of SCSI over NVMe, see Local SSD
	//  performance.
	//  Valid values:
	//
	//  * `NVME`
	//  * `SCSI`
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.interface
	Interface *string `json:"interface,omitempty"`

	// Type of the resource. Always compute#attachedDisk for attached
	//  disks.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.kind
	Kind *string `json:"kind,omitempty"`

	// A list of publicly visible licenses. Reserved for Google's use.
	//  A License represents billing and aggregate usage data for public
	//  and marketplace images.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.licenses
	Licenses []string `json:"licenses,omitempty"`

	// The mode in which to attach this disk, either `READ_WRITE` or
	//  `READ_ONLY`. If not specified, the default is to attach the disk in
	//  `READ_WRITE` mode. Valid values:
	//
	//  * `READ_ONLY`
	//  * `READ_WRITE`
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.mode
	Mode *string `json:"mode,omitempty"`

	// Indicates a valid partial or full URL to an existing Persistent Disk
	//  resource.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.source
	Source *string `json:"source,omitempty"`

	// Indicates the type of the disk, either `SCRATCH` or `PERSISTENT`.
	//  Valid values:
	//
	//  * `PERSISTENT`
	//  * `SCRATCH`
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Instance.Disk.GuestOsFeature
type Instance_Disk_GuestOSFeature struct {
	// The ID of a supported feature. Read  Enabling guest operating system
	//  features to see a list of available options.
	//  Valid values:
	//
	//  * `FEATURE_TYPE_UNSPECIFIED`
	//  * `MULTI_IP_SUBNET`
	//  * `SECURE_BOOT`
	//  * `UEFI_COMPATIBLE`
	//  * `VIRTIO_SCSI_MULTIQUEUE`
	//  * `WINDOWS`
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.Disk.GuestOsFeature.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Instance.ShieldedInstanceConfig
type Instance_ShieldedInstanceConfig struct {
	// Defines whether the instance has Secure Boot enabled.
	//
	//  Secure Boot helps ensure that the system only runs authentic software by
	//  verifying the digital signature of all boot components, and halting the
	//  boot process if signature verification fails. Disabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Defines whether the instance has the vTPM enabled. Enabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.ShieldedInstanceConfig.enable_vtpm
	EnableVTPM *bool `json:"enableVTPM,omitempty"`

	// Defines whether the instance has integrity monitoring enabled.
	//
	//  Enables monitoring and attestation of the boot integrity of the instance.
	//  The attestation is performed against the integrity policy baseline. This
	//  baseline is initially derived from the implicitly trusted boot image when
	//  the instance is created. Enabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry
type Instance_UpgradeHistoryEntry struct {
	// The snapshot of the boot disk of this notebook instance before upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.snapshot
	Snapshot *string `json:"snapshot,omitempty"`

	// The VM image before this instance upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.vm_image
	VMImage *string `json:"vmImage,omitempty"`

	// The container image before this instance upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.container_image
	ContainerImage *string `json:"containerImage,omitempty"`

	// The framework of this notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.framework
	Framework *string `json:"framework,omitempty"`

	// The version of the notebook instance before this upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.version
	Version *string `json:"version,omitempty"`

	// The state of this instance upgrade history entry.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.state
	State *string `json:"state,omitempty"`

	// The time that this instance upgrade history entry is created.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Target VM Image. Format: `ainotebooks-vm/project/image-name/name`.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.target_image
	TargetImage *string `json:"targetImage,omitempty"`

	// Action. Rolloback or Upgrade.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.action
	Action *string `json:"action,omitempty"`

	// Target VM Version, like m63.
	// +kcc:proto:field=google.cloud.notebooks.v1.Instance.UpgradeHistoryEntry.target_version
	TargetVersion *string `json:"targetVersion,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.ReservationAffinity
type ReservationAffinity struct {
	// Optional. Type of reservation to consume
	// +kcc:proto:field=google.cloud.notebooks.v1.ReservationAffinity.consume_reservation_type
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	// Optional. Corresponds to the label key of reservation resource.
	// +kcc:proto:field=google.cloud.notebooks.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of reservation resource.
	// +kcc:proto:field=google.cloud.notebooks.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.VmImage
type VMImage struct {
	// Required. The name of the Google Cloud project that this VM image belongs to.
	//  Format: `{project_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1.VmImage.project
	Project *string `json:"project,omitempty"`

	// Use VM image name to find the image.
	// +kcc:proto:field=google.cloud.notebooks.v1.VmImage.image_name
	ImageName *string `json:"imageName,omitempty"`

	// Use this VM image family to find the image; the newest image in this
	//  family will be used.
	// +kcc:proto:field=google.cloud.notebooks.v1.VmImage.image_family
	ImageFamily *string `json:"imageFamily,omitempty"`
}
