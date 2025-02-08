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

package v1alpha1


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

// +kcc:proto=google.cloud.notebooks.v1.EncryptionConfig
type EncryptionConfig struct {
	// The Cloud KMS resource identifier of the customer-managed encryption key
	//  used to protect a resource, such as a disks. It has the following
	//  format:
	//  `projects/{PROJECT_ID}/locations/{REGION}/keyRings/{KEY_RING_NAME}/cryptoKeys/{KEY_NAME}`
	// +kcc:proto:field=google.cloud.notebooks.v1.EncryptionConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.LocalDisk
type LocalDisk struct {

	// Input only. Specifies the parameters for a new disk that will be created
	//  alongside the new instance. Use initialization parameters to create boot
	//  disks or local SSDs attached to the new instance.
	//
	//  This property is mutually exclusive with the source property; you can only
	//  define one or the other, but not both.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.initialize_params
	InitializeParams *LocalDiskInitializeParams `json:"initializeParams,omitempty"`

	// Specifies the disk interface to use for attaching this disk, which is
	//  either SCSI or NVME. The default is SCSI. Persistent disks must always use
	//  SCSI and the request will fail if you attempt to attach a persistent disk
	//  in any other format than SCSI. Local SSDs can use either NVME or SCSI. For
	//  performance characteristics of SCSI over NVMe, see Local SSD performance.
	//  Valid values:
	//
	//  * `NVME`
	//  * `SCSI`
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.interface
	Interface *string `json:"interface,omitempty"`

	// The mode in which to attach this disk, either `READ_WRITE` or `READ_ONLY`.
	//  If not specified, the default is to attach the disk in `READ_WRITE` mode.
	//  Valid values:
	//
	//  * `READ_ONLY`
	//  * `READ_WRITE`
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.mode
	Mode *string `json:"mode,omitempty"`

	// Specifies a valid partial or full URL to an existing Persistent Disk
	//  resource.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.source
	Source *string `json:"source,omitempty"`

	// Specifies the type of the disk, either `SCRATCH` or `PERSISTENT`. If not
	//  specified, the default is `PERSISTENT`.
	//  Valid values:
	//
	//  * `PERSISTENT`
	//  * `SCRATCH`
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.LocalDisk.RuntimeGuestOsFeature
type LocalDisk_RuntimeGuestOsFeature struct {
	// The ID of a supported feature. Read [Enabling guest operating system
	//  features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features)
	//  to see a list of available options.
	//
	//  Valid values:
	//
	//  * `FEATURE_TYPE_UNSPECIFIED`
	//  * `MULTI_IP_SUBNET`
	//  * `SECURE_BOOT`
	//  * `UEFI_COMPATIBLE`
	//  * `VIRTIO_SCSI_MULTIQUEUE`
	//  * `WINDOWS`
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.RuntimeGuestOsFeature.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.LocalDiskInitializeParams
type LocalDiskInitializeParams struct {
	// Optional. Provide this property when creating the disk.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDiskInitializeParams.description
	Description *string `json:"description,omitempty"`

	// Optional. Specifies the disk name. If not specified, the default is to use the name
	//  of the instance. If the disk with the instance name exists already in the
	//  given zone/region, a new name will be automatically generated.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDiskInitializeParams.disk_name
	DiskName *string `json:"diskName,omitempty"`

	// Optional. Specifies the size of the disk in base-2 GB. If not specified, the disk
	//  will be the same size as the image (usually 10GB). If specified, the size
	//  must be equal to or larger than 10GB. Default 100 GB.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDiskInitializeParams.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Input only. The type of the boot disk attached to this instance, defaults to
	//  standard persistent disk (`PD_STANDARD`).
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDiskInitializeParams.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Optional. Labels to apply to this disk. These can be later modified by the
	//  disks.setLabels method. This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDiskInitializeParams.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Runtime
type Runtime struct {

	// Use a Compute Engine VM image to start the managed notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.virtual_machine
	VirtualMachine *VirtualMachine `json:"virtualMachine,omitempty"`

	// The config settings for accessing runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.access_config
	AccessConfig *RuntimeAccessConfig `json:"accessConfig,omitempty"`

	// The config settings for software inside the runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.software_config
	SoftwareConfig *RuntimeSoftwareConfig `json:"softwareConfig,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeAcceleratorConfig
type RuntimeAcceleratorConfig struct {
	// Accelerator model.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeAcceleratorConfig.type
	Type *string `json:"type,omitempty"`

	// Count of cores of this accelerator.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeAcceleratorConfig.core_count
	CoreCount *int64 `json:"coreCount,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeAccessConfig
type RuntimeAccessConfig struct {
	// The type of access mode this instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeAccessConfig.access_type
	AccessType *string `json:"accessType,omitempty"`

	// The owner of this runtime after creation. Format: `alias@example.com`
	//  Currently supports one owner only.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeAccessConfig.runtime_owner
	RuntimeOwner *string `json:"runtimeOwner,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeMetrics
type RuntimeMetrics struct {
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeShieldedInstanceConfig
type RuntimeShieldedInstanceConfig struct {
	// Defines whether the instance has Secure Boot enabled.
	//
	//  Secure Boot helps ensure that the system only runs authentic software by
	//  verifying the digital signature of all boot components, and halting the
	//  boot process if signature verification fails. Disabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Defines whether the instance has the vTPM enabled. Enabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeShieldedInstanceConfig.enable_vtpm
	EnableVTPM *bool `json:"enableVTPM,omitempty"`

	// Defines whether the instance has integrity monitoring enabled.
	//
	//  Enables monitoring and attestation of the boot integrity of the instance.
	//  The attestation is performed against the integrity policy baseline. This
	//  baseline is initially derived from the implicitly trusted boot image when
	//  the instance is created. Enabled by default.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeSoftwareConfig
type RuntimeSoftwareConfig struct {
	// Cron expression in UTC timezone, used to schedule instance auto upgrade.
	//  Please follow the [cron format](https://en.wikipedia.org/wiki/Cron).
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.notebook_upgrade_schedule
	NotebookUpgradeSchedule *string `json:"notebookUpgradeSchedule,omitempty"`

	// Verifies core internal services are running.
	//  Default: True
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.enable_health_monitoring
	EnableHealthMonitoring *bool `json:"enableHealthMonitoring,omitempty"`

	// Runtime will automatically shutdown after idle_shutdown_time.
	//  Default: True
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.idle_shutdown
	IdleShutdown *bool `json:"idleShutdown,omitempty"`

	// Time in minutes to wait before shutting down runtime. Default: 180 minutes
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.idle_shutdown_timeout
	IdleShutdownTimeout *int32 `json:"idleShutdownTimeout,omitempty"`

	// Install Nvidia Driver automatically.
	//  Default: True
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.install_gpu_driver
	InstallGpuDriver *bool `json:"installGpuDriver,omitempty"`

	// Specify a custom Cloud Storage path where the GPU driver is stored.
	//  If not specified, we'll automatically choose from official GPU drivers.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.custom_gpu_driver_path
	CustomGpuDriverPath *string `json:"customGpuDriverPath,omitempty"`

	// Path to a Bash script that automatically runs after a notebook instance
	//  fully boots up. The path must be a URL or
	//  Cloud Storage path (`gs://path-to-file/file-name`).
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.post_startup_script
	PostStartupScript *string `json:"postStartupScript,omitempty"`

	// Optional. Use a list of container images to use as Kernels in the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.kernels
	Kernels []ContainerImage `json:"kernels,omitempty"`

	// Behavior for the post startup script.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.post_startup_script_behavior
	PostStartupScriptBehavior *string `json:"postStartupScriptBehavior,omitempty"`

	// Bool indicating whether JupyterLab terminal will be available or not.
	//  Default: False
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.disable_terminal
	DisableTerminal *bool `json:"disableTerminal,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.VirtualMachine
type VirtualMachine struct {

	// Virtual Machine configuration settings.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachine.virtual_machine_config
	VirtualMachineConfig *VirtualMachineConfig `json:"virtualMachineConfig,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.VirtualMachineConfig
type VirtualMachineConfig struct {

	// Required. The Compute Engine machine type used for runtimes.
	//  Short name is valid. Examples:
	//  * `n1-standard-2`
	//  * `e2-standard-8`
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. Use a list of container images to use as Kernels in the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.container_images
	ContainerImages []ContainerImage `json:"containerImages,omitempty"`

	// Required. Data disk option configuration settings.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.data_disk
	DataDisk *LocalDisk `json:"dataDisk,omitempty"`

	// Optional. Encryption settings for virtual machine data disk.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Optional. Shielded VM Instance configuration settings.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.shielded_instance_config
	ShieldedInstanceConfig *RuntimeShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// Optional. The Compute Engine accelerator configuration for this runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.accelerator_config
	AcceleratorConfig *RuntimeAcceleratorConfig `json:"acceleratorConfig,omitempty"`

	// Optional. The Compute Engine network to be used for machine
	//  communications. Cannot be specified with subnetwork. If neither
	//  `network` nor `subnet` is specified, the "default" network of
	//  the project is used, if it exists.
	//
	//  A full URL or partial URI. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/networks/default`
	//  * `projects/[project_id]/global/networks/default`
	//
	//  Runtimes are managed resources inside Google Infrastructure.
	//  Runtimes support the following network configurations:
	//
	//  * Google Managed Network (Network & subnet are empty)
	//  * Consumer Project VPC (network & subnet are required). Requires
	//  configuring Private Service Access.
	//  * Shared VPC (network & subnet are required). Requires configuring Private
	//  Service Access.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.network
	Network *string `json:"network,omitempty"`

	// Optional. The Compute Engine subnetwork to be used for machine
	//  communications. Cannot be specified with network.
	//
	//  A full URL or partial URI are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0`
	//  * `projects/[project_id]/regions/us-east1/subnetworks/sub0`
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.subnet
	Subnet *string `json:"subnet,omitempty"`

	// Optional. If true, runtime will only have internal IP
	//  addresses. By default, runtimes are not restricted to internal IP
	//  addresses, and will have ephemeral external IP addresses assigned to each
	//  vm. This `internal_ip_only` restriction can only be enabled for
	//  subnetwork enabled networks, and all dependencies must be
	//  configured to be accessible without external IP addresses.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.internal_ip_only
	InternalIPOnly *bool `json:"internalIPOnly,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging
	//  instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.tags
	Tags []string `json:"tags,omitempty"`

	// Optional. The Compute Engine metadata entries to add to virtual machine. (see
	//  [Project and instance
	//  metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. The labels to associate with this runtime.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a cluster.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The type of vNIC to be used on this interface. This may be gVNIC or
	//  VirtioNet.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.nic_type
	NicType *string `json:"nicType,omitempty"`

	// Optional. Reserved IP Range name is used for VPC Peering.
	//  The subnetwork allocation will use the range *name* if it's assigned.
	//
	//  Example: managed-notebooks-range-c
	//
	//      PEERING_RANGE_NAME_3=managed-notebooks-range-c
	//      gcloud compute addresses create $PEERING_RANGE_NAME_3 \
	//        --global \
	//        --prefix-length=24 \
	//        --description="Google Cloud Managed Notebooks Range 24 c" \
	//        --network=$NETWORK \
	//        --addresses=192.168.0.0 \
	//        --purpose=VPC_PEERING
	//
	//  Field value will be: `managed-notebooks-range-c`
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.reserved_ip_range
	ReservedIPRange *string `json:"reservedIPRange,omitempty"`

	// Optional. Boot image metadata used for runtime upgradeability.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.boot_image
	BootImage *VirtualMachineConfig_BootImage `json:"bootImage,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.VirtualMachineConfig.BootImage
type VirtualMachineConfig_BootImage struct {
}

// +kcc:proto=google.cloud.notebooks.v1.LocalDisk
type LocalDiskObservedState struct {
	// Optional. Output only. Specifies whether the disk will be auto-deleted when the
	//  instance is deleted (but not when the disk is detached from the instance).
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.auto_delete
	AutoDelete *bool `json:"autoDelete,omitempty"`

	// Optional. Output only. Indicates that this is a boot disk. The virtual machine
	//  will use the first partition of the disk for its root filesystem.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.boot
	Boot *bool `json:"boot,omitempty"`

	// Optional. Output only. Specifies a unique device name
	//  of your choice that is reflected into the
	//  `/dev/disk/by-id/google-*` tree of a Linux operating system running within
	//  the instance. This name can be used to reference the device for mounting,
	//  resizing, and so on, from within the instance.
	//
	//  If not specified, the server chooses a default device name to apply to this
	//  disk, in the form persistent-disk-x, where x is a number assigned by Google
	//  Compute Engine. This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.device_name
	DeviceName *string `json:"deviceName,omitempty"`

	// Output only. Indicates a list of features to enable on the guest operating system.
	//  Applicable only for bootable images. Read  Enabling guest operating
	//  system features to see a list of available options.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.guest_os_features
	GuestOsFeatures []LocalDisk_RuntimeGuestOsFeature `json:"guestOsFeatures,omitempty"`

	// Output only. A zero-based index to this disk, where 0 is reserved for the
	//  boot disk. If you have many disks attached to an instance, each disk would
	//  have a unique index number.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.index
	Index *int32 `json:"index,omitempty"`

	// Output only. Type of the resource. Always compute#attachedDisk for attached disks.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.kind
	Kind *string `json:"kind,omitempty"`

	// Output only. Any valid publicly visible licenses.
	// +kcc:proto:field=google.cloud.notebooks.v1.LocalDisk.licenses
	Licenses []string `json:"licenses,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Runtime
type RuntimeObservedState struct {
	// Output only. The resource name of the runtime.
	//  Format:
	//  `projects/{project}/locations/{location}/runtimes/{runtimeId}`
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.name
	Name *string `json:"name,omitempty"`

	// Use a Compute Engine VM image to start the managed notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.virtual_machine
	VirtualMachine *VirtualMachineObservedState `json:"virtualMachine,omitempty"`

	// Output only. Runtime state.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.state
	State *string `json:"state,omitempty"`

	// Output only. Runtime health_state.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.health_state
	HealthState *string `json:"healthState,omitempty"`

	// The config settings for accessing runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.access_config
	AccessConfig *RuntimeAccessConfigObservedState `json:"accessConfig,omitempty"`

	// The config settings for software inside the runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.software_config
	SoftwareConfig *RuntimeSoftwareConfigObservedState `json:"softwareConfig,omitempty"`

	// Output only. Contains Runtime daemon metrics such as Service status and JupyterLab
	//  stats.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.metrics
	Metrics *RuntimeMetrics `json:"metrics,omitempty"`

	// Output only. Runtime creation time.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Runtime update time.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeAccessConfig
type RuntimeAccessConfigObservedState struct {
	// Output only. The proxy endpoint that is used to access the runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeAccessConfig.proxy_uri
	ProxyURI *string `json:"proxyURI,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeMetrics
type RuntimeMetricsObservedState struct {
	// Output only. The system metrics.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeMetrics.system_metrics
	SystemMetrics map[string]string `json:"systemMetrics,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.RuntimeSoftwareConfig
type RuntimeSoftwareConfigObservedState struct {
	// Output only. Bool indicating whether an newer image is available in an image family.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.upgradeable
	Upgradeable *bool `json:"upgradeable,omitempty"`

	// Output only. version of boot image such as M100, from release label of the image.
	// +kcc:proto:field=google.cloud.notebooks.v1.RuntimeSoftwareConfig.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.VirtualMachine
type VirtualMachineObservedState struct {
	// Output only. The user-friendly name of the Managed Compute Engine instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachine.instance_name
	InstanceName *string `json:"instanceName,omitempty"`

	// Output only. The unique identifier of the Managed Compute Engine instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachine.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// Virtual Machine configuration settings.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachine.virtual_machine_config
	VirtualMachineConfig *VirtualMachineConfigObservedState `json:"virtualMachineConfig,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.VirtualMachineConfig
type VirtualMachineConfigObservedState struct {
	// Output only. The zone where the virtual machine is located.
	//  If using regional request, the notebooks service will pick a location
	//  in the corresponding runtime region.
	//  On a get request, zone will always be present. Example:
	//  * `us-central1-b`
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.zone
	Zone *string `json:"zone,omitempty"`

	// Required. Data disk option configuration settings.
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.data_disk
	DataDisk *LocalDiskObservedState `json:"dataDisk,omitempty"`

	// Output only. The Compute Engine guest attributes. (see
	//  [Project and instance
	//  guest
	//  attributes](https://cloud.google.com/compute/docs/storing-retrieving-metadata#guest_attributes)).
	// +kcc:proto:field=google.cloud.notebooks.v1.VirtualMachineConfig.guest_attributes
	GuestAttributes map[string]string `json:"guestAttributes,omitempty"`
}
