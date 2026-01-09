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
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1
// resource: ComputeFirewallPolicyRule:FirewallPolicyRule
// resource: ComputeForwardingRule:ForwardingRule
// resource: ComputeInstance:Instance
// resource: ComputeSubnetwork:Subnetwork
// resource: ComputeTargetTcpProxy:TargetTcpProxy

package v1beta1

// +kcc:proto=google.cloud.compute.v1.AcceleratorConfig
type AcceleratorConfig struct {
	// The number of the guest accelerator cards exposed to this instance.
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Full or partial URL of the accelerator type resource to attach to this instance. For example: projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100 If you are creating an instance template, specify only the accelerator name. See GPUs on Compute Engine for a full list of accelerator types.
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AccessConfig
type AccessConfig struct {
	// Applies to ipv6AccessConfigs only. The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. To use a static external IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an external IPv6 address from the instance's subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.external_ipv6
	ExternalIPV6 *string `json:"externalIPV6,omitempty"`

	// Applies to ipv6AccessConfigs only. The prefix length of the external IPv6 range.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.external_ipv6_prefix_length
	ExternalIPV6PrefixLength *int32 `json:"externalIPV6PrefixLength,omitempty"`

	// [Output Only] Type of the resource. Always compute#accessConfig for access configs.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.kind
	Kind *string `json:"kind,omitempty"`

	// The name of this access configuration. In accessConfigs (IPv4), the default and recommended name is External NAT, but you can use any arbitrary string, such as My external IP or Network Access. In ipv6AccessConfigs, the recommend name is External IPv6.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.name
	Name *string `json:"name,omitempty"`

	// Applies to accessConfigs (IPv4) only. An external IP address associated with this instance. Specify an unused static external IP address available to the project or leave this field undefined to use an IP from a shared ephemeral IP address pool. If you specify a static external IP address, it must live in the same region as the zone of the instance.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.nat_i_p
	NATIP *string `json:"natIP,omitempty"`

	// This signifies the networking tier used for configuring this access configuration and can only take the following values: PREMIUM, STANDARD. If an AccessConfig is specified without a valid external IP address, an ephemeral IP will be created with this networkTier. If an AccessConfig with a valid external IP address is specified, it must match that of the networkTier associated with the Address resource owning that IP.
	//  Check the NetworkTier enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.network_tier
	NetworkTier *string `json:"networkTier,omitempty"`

	// The DNS domain name for the public PTR record. You can set this field only if the `setPublicPtr` field is enabled in accessConfig. If this field is unspecified in ipv6AccessConfig, a default PTR record will be created for first IP in associated external IPv6 range.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.public_ptr_domain_name
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`

	// The resource URL for the security policy associated with this access config.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.security_policy
	SecurityPolicy *string `json:"securityPolicy,omitempty"`

	// Specifies whether a public DNS 'PTR' record should be created to map the external IP address of the instance to a DNS domain name. This field is not used in ipv6AccessConfig. A default PTR record will be created if the VM has external IPv6 range associated.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.set_public_ptr
	SetPublicPtr *bool `json:"setPublicPtr,omitempty"`

	// The type of configuration. In accessConfigs (IPv4), the default and only option is ONE_TO_ONE_NAT. In ipv6AccessConfigs, the default and only option is DIRECT_IPV6.
	//  Check the Type enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AdvancedMachineFeatures
type AdvancedMachineFeatures struct {
	// Whether to enable nested virtualization or not (default is false).
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.enable_nested_virtualization
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`

	// Whether to enable UEFI networking for instance creation.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.enable_uefi_networking
	EnableUefiNetworking *bool `json:"enableUefiNetworking,omitempty"`

	// Type of Performance Monitoring Unit requested on instance.
	//  Check the PerformanceMonitoringUnit enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.performance_monitoring_unit
	PerformanceMonitoringUnit *string `json:"performanceMonitoringUnit,omitempty"`

	// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.threads_per_core
	ThreadsPerCore *int32 `json:"threadsPerCore,omitempty"`

	// Turbo frequency mode to use for the instance. Supported modes include: * ALL_CORE_MAX Using empty string or not setting this field will use the platform-specific default turbo mode.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.turbo_mode
	TurboMode *string `json:"turboMode,omitempty"`

	// The number of physical cores to expose to an instance. Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance's nominal CPU count and the underlying platform's SMT width.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.visible_core_count
	VisibleCoreCount *int32 `json:"visibleCoreCount,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AliasIpRange
type AliasIPRange struct {
	// The IP alias ranges to allocate for this interface. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. This range may be a single IP address (such as 10.2.3.4), a netmask (such as /24) or a CIDR-formatted string (such as 10.1.2.0/24).
	// +kcc:proto:field=google.cloud.compute.v1.AliasIpRange.ip_cidr_range
	IPCIDRRange *string `json:"ipCIDRRange,omitempty"`

	// The name of a subnetwork secondary IP range from which to allocate an IP alias range. If not specified, the primary range of the subnetwork is used.
	// +kcc:proto:field=google.cloud.compute.v1.AliasIpRange.subnetwork_range_name
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AttachedDisk
type AttachedDisk struct {
	// [Output Only] The architecture of the attached disk. Valid values are ARM64 or X86_64.
	//  Check the Architecture enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.architecture
	Architecture *string `json:"architecture,omitempty"`

	// Specifies whether the disk will be auto-deleted when the instance is deleted (but not when the disk is detached from the instance).
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.auto_delete
	AutoDelete *bool `json:"autoDelete,omitempty"`

	// Indicates that this is a boot disk. The virtual machine will use the first partition of the disk for its root filesystem.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.boot
	Boot *bool `json:"boot,omitempty"`

	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance. This name can be used to reference the device for mounting, resizing, and so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disk-x, where x is a number assigned by Google Compute Engine. This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.device_name
	DeviceName *string `json:"deviceName,omitempty"`

	// Encrypts or decrypts a disk using a customer-supplied encryption key. If you are creating a new disk, this field encrypts the new disk using an encryption key that you provide. If you are attaching an existing disk that is already encrypted, this field decrypts the disk using the customer-supplied encryption key. If you encrypt a disk using a customer-supplied key, you must provide the same key again when you attempt to use this resource at a later time. For example, you must provide the key when you create a snapshot or an image from the disk or when you attach the disk to a virtual machine instance. If you do not provide an encryption key, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Note: Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt disks in a managed instance group. You cannot create VMs that have disks with customer-supplied keys using the bulk insert method.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.disk_encryption_key
	DiskEncryptionKey *CustomerEncryptionKey `json:"diskEncryptionKey,omitempty"`

	// The size of the disk in GB.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// [Input Only] Whether to force attach the regional disk even if it's currently attached to another instance. If you try to force attach a zonal disk to an instance, you will receive an error.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.force_attach
	ForceAttach *bool `json:"forceAttach,omitempty"`

	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.guest_os_features
	GuestOSFeatures []GuestOSFeature `json:"guestOSFeatures,omitempty"`

	// [Output Only] A zero-based index to this disk, where 0 is reserved for the boot disk. If you have many disks attached to an instance, each disk would have a unique index number.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.index
	Index *int32 `json:"index,omitempty"`

	// [Input Only] Specifies the parameters for a new disk that will be created alongside the new instance. Use initialization parameters to create boot disks or local SSDs attached to the new instance. This property is mutually exclusive with the source property; you can only define one or the other, but not both.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.initialize_params
	InitializeParams *AttachedDiskInitializeParams `json:"initializeParams,omitempty"`

	// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. For most machine types, the default is SCSI. Local SSDs can use either NVME or SCSI. In certain configurations, persistent disks can use NVMe. For more information, see About persistent disks.
	//  Check the Interface enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.interface
	Interface *string `json:"interface,omitempty"`

	// [Output Only] Type of the resource. Always compute#attachedDisk for attached disks.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] Any valid publicly visible licenses.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.licenses
	Licenses []string `json:"licenses,omitempty"`

	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the disk in READ_WRITE mode.
	//  Check the Mode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.mode
	Mode *string `json:"mode,omitempty"`

	// For LocalSSD disks on VM Instances in STOPPED or SUSPENDED state, this field is set to PRESERVED if the LocalSSD data has been saved to a persistent location by customer request. (see the discard_local_ssd option on Stop/Suspend). Read-only in the api.
	//  Check the SavedState enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.saved_state
	SavedState *string `json:"savedState,omitempty"`

	// [Output Only] shielded vm initial state stored on disk
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.shielded_instance_initial_state
	ShieldedInstanceInitialState *InitialStateConfig `json:"shieldedInstanceInitialState,omitempty"`

	// Specifies a valid partial or full URL to an existing Persistent Disk resource. When creating a new instance boot disk, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required. If desired, you can also attach existing non-root persistent disks using this property. This field is only applicable for persistent disks. Note that for InstanceTemplate, specify the disk name for zonal disk, and the URL for regional disk.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.source
	Source *string `json:"source,omitempty"`

	// Specifies the type of the disk, either SCRATCH or PERSISTENT. If not specified, the default is PERSISTENT.
	//  Check the Type enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AttachedDiskInitializeParams
type AttachedDiskInitializeParams struct {
	// The architecture of the attached disk. Valid values are arm64 or x86_64.
	//  Check the Architecture enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.architecture
	Architecture *string `json:"architecture,omitempty"`

	// An optional description. Provide this property when creating the disk.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.description
	Description *string `json:"description,omitempty"`

	// Specifies the disk name. If not specified, the default is to use the name of the instance. If a disk with the same name already exists in the given region, the existing disk is attached to the new instance and the new disk is not created.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.disk_name
	DiskName *string `json:"diskName,omitempty"`

	// Specifies the size of the disk in base-2 GB. The size must be at least 10 GB. If you specify a sourceImage, which is required for boot disks, the default size is the size of the sourceImage. If you do not specify a sourceImage, the default disk size is 500 GB.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Specifies the disk type to use to create the instance. If not specified, the default is pd-standard, specified using the full URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/pd-standard For a full list of acceptable values, see Persistent disk types. If you specify this field when creating a VM, you can provide either the full or partial URL. For example, the following values are valid: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/diskType - projects/project/zones/zone/diskTypes/diskType - zones/zone/diskTypes/diskType If you specify this field when creating or updating an instance template or all-instances configuration, specify the type of the disk, not the URL. For example: pd-standard.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Whether this disk is using confidential compute mode.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.enable_confidential_compute
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`

	// Labels to apply to this disk. These can be later modified by the disks.setLabels method. This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.labels
	Labels map[string]string `json:"labels,omitempty"`

	// A list of publicly visible licenses. Reserved for Google's use.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.licenses
	Licenses []string `json:"licenses,omitempty"`

	// Specifies which action to take on instance update with this disk. Default is to use the existing disk.
	//  Check the OnUpdateAction enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.on_update_action
	OnUpdateAction *string `json:"onUpdateAction,omitempty"`

	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.provisioned_iops
	ProvisionedIops *int64 `json:"provisionedIops,omitempty"`

	// Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle. Values must greater than or equal to 1.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.provisioned_throughput
	ProvisionedThroughput *int64 `json:"provisionedThroughput,omitempty"`

	// Required for each regional disk associated with the instance. Specify the URLs of the zones where the disk should be replicated to. You must provide exactly two replica zones, and one zone must be the same as the instance zone.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.replica_zones
	ReplicaZones []string `json:"replicaZones,omitempty"`

	// Resource manager tags to be bound to the disk. Tag keys and values have the same definition as resource manager tags. Keys must be in the format `tagKeys/{tag_key_id}`, and values are in the format `tagValues/456`. The field is ignored (both PUT & PATCH) when empty.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.resource_manager_tags
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`

	// Resource policies applied to this disk for automatic snapshot creations. Specified using the full or partial URL. For instance template, specify only the resource policy name.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.resource_policies
	ResourcePolicies []string `json:"resourcePolicies,omitempty"`

	// The source image to create this disk. When creating a new instance boot disk, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family If the source image is deleted later, this field will not be set.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.source_image
	SourceImage *string `json:"sourceImage,omitempty"`

	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key. InstanceTemplate and InstancePropertiesPatch do not store customer-supplied encryption keys, so you cannot create disks for instances in a managed instance group if the source images are encrypted with your own keys.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.source_image_encryption_key
	SourceImageEncryptionKey *CustomerEncryptionKey `json:"sourceImageEncryptionKey,omitempty"`

	// The source snapshot to create this disk. When creating a new instance boot disk, one of initializeParams.sourceSnapshot or initializeParams.sourceImage or disks.source is required. To create a disk with a snapshot that you created, specify the snapshot name in the following format: global/snapshots/my-backup If the source snapshot is deleted later, this field will not be set. Note: You cannot create VMs in bulk using a snapshot as the source. Use an image instead when you create VMs using the bulk insert method.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.source_snapshot
	SourceSnapshot *string `json:"sourceSnapshot,omitempty"`

	// The customer-supplied encryption key of the source snapshot.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.source_snapshot_encryption_key
	SourceSnapshotEncryptionKey *CustomerEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty"`

	// The storage pool in which the new disk is created. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /storagePools/storagePool - projects/project/zones/zone/storagePools/storagePool - zones/zone/storagePools/storagePool
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.storage_pool
	StoragePool *string `json:"storagePool,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ConfidentialInstanceConfig
type ConfidentialInstanceConfig struct {
	// Defines the type of technology used by the confidential instance.
	//  Check the ConfidentialInstanceType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.ConfidentialInstanceConfig.confidential_instance_type
	ConfidentialInstanceType *string `json:"confidentialInstanceType,omitempty"`

	// Defines whether the instance should have confidential compute enabled.
	// +kcc:proto:field=google.cloud.compute.v1.ConfidentialInstanceConfig.enable_confidential_compute
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.CustomerEncryptionKey
type CustomerEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS. For example: "kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key The fully-qualifed key name may be returned for resource GET requests. For example: "kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeyVersions/1
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// The service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used. For example: "kmsKeyServiceAccount": "name@project_id.iam.gserviceaccount.com/
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.kms_key_service_account
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource. You can provide either the rawKey or the rsaEncryptedKey. For example: "rawKey": "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0="
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.raw_key
	RawKey *string `json:"rawKey,omitempty"`

	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource. You can provide either the rawKey or the rsaEncryptedKey. For example: "rsaEncryptedKey": "ieCx/NcW06PcT7Ep1X6LUTc/hLvUDYyzSZPPVCVPTVEohpeHASqC8uw5TzyO9U+Fka9JFH z0mBibXUInrC/jEk014kCK/NPjYgEMOyssZ4ZINPKxlUh2zn1bV+MCaTICrdmuSBTWlUUiFoD D6PYznLwh8ZNdaheCeZ8ewEXgFQ8V+sDroLaN3Xs3MDTXQEMMoNUXMCZEIpg9Vtp9x2oe==" The key must meet the following requirements before you can provide it to Compute Engine: 1. The key is wrapped using a RSA public key certificate provided by Google. 2. After being wrapped, the key must be encoded in RFC 4648 base64 encoding. Gets the RSA public key certificate provided by Google at: https://cloud-certs.storage.googleapis.com/google-cloud-csek-ingress.pem
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.rsa_encrypted_key
	RsaEncryptedKey *string `json:"rsaEncryptedKey,omitempty"`

	// [Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	// +kcc:proto:field=google.cloud.compute.v1.CustomerEncryptionKey.sha256
	Sha256 *string `json:"sha256,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.DisplayDevice
type DisplayDevice struct {
	// Defines whether the instance has Display enabled.
	// +kcc:proto:field=google.cloud.compute.v1.DisplayDevice.enable_display
	EnableDisplay *bool `json:"enableDisplay,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Duration
type Duration struct {
	// Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field. Must be from 0 to 999,999,999 inclusive.
	// +kcc:proto:field=google.cloud.compute.v1.Duration.nanos
	Nanos *int32 `json:"nanos,omitempty"`

	// Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive. Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	// +kcc:proto:field=google.cloud.compute.v1.Duration.seconds
	Seconds *int64 `json:"seconds,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FileContentBuffer
type FileContentBuffer struct {
	// The raw content in the secure keys file.
	// +kcc:proto:field=google.cloud.compute.v1.FileContentBuffer.content
	Content *string `json:"content,omitempty"`

	// The file type of source file.
	//  Check the FileType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FileContentBuffer.file_type
	FileType *string `json:"fileType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleSecureTag
type FirewallPolicyRuleSecureTag struct {
	// Name of the secure tag, created with TagManager's TagValue API.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.name
	Name *string `json:"name,omitempty"`

	// [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
	//  Check the State enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.GuestOsFeature
type GuestOSFeature struct {
	// The ID of a supported feature. To add multiple values, use commas to separate values. Set to one or more of the following values: - VIRTIO_SCSI_MULTIQUEUE - WINDOWS - MULTI_IP_SUBNET - UEFI_COMPATIBLE - GVNIC - SEV_CAPABLE - SUSPEND_RESUME_COMPATIBLE - SEV_LIVE_MIGRATABLE_V2 - SEV_SNP_CAPABLE - TDX_CAPABLE - IDPF - SNP_SVSM_CAPABLE For more information, see Enabling guest operating system features.
	//  Check the Type enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.GuestOsFeature.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InitialStateConfig
type InitialStateConfig struct {
	// The Key Database (db).
	// +kcc:proto:field=google.cloud.compute.v1.InitialStateConfig.dbs
	Dbs []FileContentBuffer `json:"dbs,omitempty"`

	// The forbidden key database (dbx).
	// +kcc:proto:field=google.cloud.compute.v1.InitialStateConfig.dbxs
	Dbxs []FileContentBuffer `json:"dbxs,omitempty"`

	// The Key Exchange Key (KEK).
	// +kcc:proto:field=google.cloud.compute.v1.InitialStateConfig.keks
	Keks []FileContentBuffer `json:"keks,omitempty"`

	// The Platform Key (PK).
	// +kcc:proto:field=google.cloud.compute.v1.InitialStateConfig.pk
	Pk *FileContentBuffer `json:"pk,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InstanceParams
type InstanceParams struct {
	// Resource manager tags to be bound to the instance. Tag keys and values have the same definition as resource manager tags. Keys must be in the format `tagKeys/{tag_key_id}`, and values are in the format `tagValues/456`. The field is ignored (both PUT & PATCH) when empty.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceParams.resource_manager_tags
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Items
type Items struct {
	// Key for the metadata entry. Keys must conform to the following regexp: [a-zA-Z0-9-_]+, and be less than 128 bytes in length. This is reflected as part of a URL in the metadata server. Additionally, to avoid ambiguity, keys must not conflict with any other metadata keys for the project.
	// +kcc:proto:field=google.cloud.compute.v1.Items.key
	Key *string `json:"key,omitempty"`

	// Value for the metadata entry. These are free-form strings, and only have meaning as interpreted by the image running in the instance. The only restriction placed on values is that their size must be less than or equal to 262144 bytes (256 KiB).
	// +kcc:proto:field=google.cloud.compute.v1.Items.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Metadata
type Metadata struct {
	// Specifies a fingerprint for this request, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Metadata.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// Array of key/value pairs. The total size of all keys and values must be less than 512 KB.
	// +kcc:proto:field=google.cloud.compute.v1.Metadata.items
	Items []Items `json:"items,omitempty"`

	// [Output Only] Type of the resource. Always compute#metadata for metadata.
	// +kcc:proto:field=google.cloud.compute.v1.Metadata.kind
	Kind *string `json:"kind,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkInterface
type NetworkInterface struct {
	// An array of configurations for this interface. Currently, only one access config, ONE_TO_ONE_NAT, is supported. If there are no accessConfigs specified, then this instance will have no external internet access.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.access_configs
	AccessConfigs []AccessConfig `json:"accessConfigs,omitempty"`

	// An array of alias IP ranges for this network interface. You can only specify this field for network interfaces in VPC networks.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.alias_ip_ranges
	AliasIPRanges []AliasIPRange `json:"aliasIPRanges,omitempty"`

	// Fingerprint hash of contents stored in this network interface. This field will be ignored when inserting an Instance or adding a NetworkInterface. An up-to-date fingerprint must be provided in order to update the NetworkInterface. The request will fail with error 400 Bad Request if the fingerprint is not provided, or 412 Precondition Failed if the fingerprint is out of date.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// The prefix length of the primary internal IPv6 range.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.internal_ipv6_prefix_length
	InternalIPV6PrefixLength *int32 `json:"internalIPV6PrefixLength,omitempty"`

	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.ipv6_access_configs
	IPV6AccessConfigs []AccessConfig `json:"ipv6AccessConfigs,omitempty"`

	// [Output Only] One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork. Valid only if stackType is IPV4_IPV6.
	//  Check the Ipv6AccessType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.ipv6_access_type
	IPV6AccessType *string `json:"ipv6AccessType,omitempty"`

	// An IPv6 internal network address for this network interface. To use a static internal IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.ipv6_address
	IPV6Address *string `json:"ipv6Address,omitempty"`

	// [Output Only] Type of the resource. Always compute#networkInterface for network interfaces.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] The name of the network interface, which is generated by the server. For a VM, the network interface uses the nicN naming format. Where N is a value between 0 and 7. The default interface value is nic0.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.name
	Name *string `json:"name,omitempty"`

	// URL of the VPC network resource for this instance. When creating an instance, if neither the network nor the subnetwork is specified, the default network global/networks/default is used. If the selected project doesn't have the default network, you must specify a network or subnet. If the network is not specified but the subnetwork is specified, the network is inferred. If you specify this property, you can specify the network as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/global/networks/ network - projects/project/global/networks/network - global/networks/default
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.network
	Network *string `json:"network,omitempty"`

	// The URL of the network attachment that this interface should connect to in the following format: projects/{project_number}/regions/{region_name}/networkAttachments/{network_attachment_name}.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.network_attachment
	NetworkAttachment *string `json:"networkAttachment,omitempty"`

	// An IPv4 internal IP address to assign to the instance for this network interface. If not specified by the user, an unused internal IP is assigned by the system.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.network_i_p
	NetworkIP *string `json:"networkIP,omitempty"`

	// The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
	//  Check the NicType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.nic_type
	NicType *string `json:"nicType,omitempty"`

	// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It'll be empty if not specified by the users.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.queue_count
	QueueCount *int32 `json:"queueCount,omitempty"`

	// The stack type for this network interface. To assign only IPv4 addresses, use IPV4_ONLY. To assign both IPv4 and IPv6 addresses, use IPV4_IPV6. If not specified, IPV4_ONLY is used. This field can be both set at instance creation and update network interface operations.
	//  Check the StackType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.stack_type
	StackType *string `json:"stackType,omitempty"`

	// The URL of the Subnetwork resource for this instance. If the network resource is in legacy mode, do not specify this field. If the network is in auto subnet mode, specifying the subnetwork is optional. If the network is in custom subnet mode, specifying the subnetwork is required. If you specify this field, you can specify the subnetwork as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/regions/region /subnetworks/subnetwork - regions/region/subnetworks/subnetwork
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkPerformanceConfig
type NetworkPerformanceConfig struct {
	// Check the TotalEgressBandwidthTier enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkPerformanceConfig.total_egress_bandwidth_tier
	TotalEgressBandwidthTier *string `json:"totalEgressBandwidthTier,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ReservationAffinity
type ReservationAffinity struct {
	// Specifies the type of reservation from which this instance can consume resources: ANY_RESERVATION (default), SPECIFIC_RESERVATION, or NO_RESERVATION. See Consuming reserved instances for examples.
	//  Check the ConsumeReservationType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.ReservationAffinity.consume_reservation_type
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify googleapis.com/reservation-name as the key and specify the name of your reservation as its value.
	// +kcc:proto:field=google.cloud.compute.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Corresponds to the label values of a reservation resource. This can be either a name to a reservation in the same project or "projects/different-project/reservations/some-reservation-name" to target a shared reservation in the same zone but in a different project.
	// +kcc:proto:field=google.cloud.compute.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourceStatus
type ResourceStatus struct {
	// [Output Only] The precise location of your instance within the zone's data center, including the block, sub-block, and host. The field is formatted as follows: blockId/subBlockId/hostId.
	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatus.physical_host
	PhysicalHost *string `json:"physicalHost,omitempty"`

	// [Output Only] A series of fields containing the global name of the Compute Engine cluster, as well as the ID of the block, sub-block, and host on which the running instance is located.
	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatus.physical_host_topology
	PhysicalHostTopology *ResourceStatusPhysicalHostTopology `json:"physicalHostTopology,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatus.scheduling
	Scheduling *ResourceStatusScheduling `json:"scheduling,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatus.upcoming_maintenance
	UpcomingMaintenance *UpcomingMaintenance `json:"upcomingMaintenance,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourceStatusPhysicalHostTopology
type ResourceStatusPhysicalHostTopology struct {
	// [Output Only] The ID of the block in which the running instance is located. Instances within the same block experience low network latency.
	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatusPhysicalHostTopology.block
	Block *string `json:"block,omitempty"`

	// [Output Only] The global name of the Compute Engine cluster where the running instance is located.
	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatusPhysicalHostTopology.cluster
	Cluster *string `json:"cluster,omitempty"`

	// [Output Only] The ID of the host on which the running instance is located. Instances on the same host experience the lowest possible network latency.
	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatusPhysicalHostTopology.host
	Host *string `json:"host,omitempty"`

	// [Output Only] The ID of the sub-block in which the running instance is located. Instances in the same sub-block experience lower network latency than instances in the same block.
	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatusPhysicalHostTopology.subblock
	Subblock *string `json:"subblock,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourceStatusScheduling
type ResourceStatusScheduling struct {
	// Specifies the availability domain to place the instance in. The value must be a number between 1 and the number of availability domains specified in the spread placement policy attached to the instance.
	// +kcc:proto:field=google.cloud.compute.v1.ResourceStatusScheduling.availability_domain
	AvailabilityDomain *int32 `json:"availabilityDomain,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Scheduling
type Scheduling struct {
	// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). You can only set the automatic restart option for standard instances. Preemptible instances cannot be automatically restarted. By default, this is set to true so an instance is automatically restarted if it is terminated by Compute Engine.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.automatic_restart
	AutomaticRestart *bool `json:"automaticRestart,omitempty"`

	// Specifies the availability domain to place the instance in. The value must be a number between 1 and the number of availability domains specified in the spread placement policy attached to the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.availability_domain
	AvailabilityDomain *int32 `json:"availabilityDomain,omitempty"`

	// Specify the time in seconds for host error detection, the value must be within the range of [90, 330] with the increment of 30, if unset, the default behavior of host error recovery will be used.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.host_error_timeout_seconds
	HostErrorTimeoutSeconds *int32 `json:"hostErrorTimeoutSeconds,omitempty"`

	// Specifies the termination action for the instance.
	//  Check the InstanceTerminationAction enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.instance_termination_action
	InstanceTerminationAction *string `json:"instanceTerminationAction,omitempty"`

	// Specifies the maximum amount of time a Local Ssd Vm should wait while recovery of the Local Ssd state is attempted. Its value should be in between 0 and 168 hours with hour granularity and the default value being 1 hour.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.local_ssd_recovery_timeout
	LocalSsdRecoveryTimeout *Duration `json:"localSsdRecoveryTimeout,omitempty"`

	// An opaque location hint used to place the instance close to other resources. This field is for use by internal tools that use the public API.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.location_hint
	LocationHint *string `json:"locationHint,omitempty"`

	// Specifies the max run duration for the given instance. If specified, the instance termination action will be performed at the end of the run duration.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.max_run_duration
	MaxRunDuration *Duration `json:"maxRunDuration,omitempty"`

	// The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.min_node_cpus
	MinNodeCpus *int32 `json:"minNodeCpus,omitempty"`

	// A set of node affinity and anti-affinity configurations. Refer to Configuring node affinity for more information. Overrides reservationAffinity.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.node_affinities
	NodeAffinities []SchedulingNodeAffinity `json:"nodeAffinities,omitempty"`

	// Defines the maintenance behavior for this instance. For standard instances, the default behavior is MIGRATE. For preemptible instances, the default and only possible behavior is TERMINATE. For more information, see Set VM host maintenance policy.
	//  Check the OnHostMaintenance enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.on_host_maintenance
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.on_instance_stop_action
	OnInstanceStopAction *SchedulingOnInstanceStopAction `json:"onInstanceStopAction,omitempty"`

	// Defines whether the instance is preemptible. This can only be set during instance creation or while the instance is stopped and therefore, in a `TERMINATED` state. See Instance Life Cycle for more information on the possible instance states.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Specifies the provisioning model of the instance.
	//  Check the ProvisioningModel enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.provisioning_model
	ProvisioningModel *string `json:"provisioningModel,omitempty"`

	// Specifies the timestamp, when the instance will be terminated, in RFC3339 text format. If specified, the instance termination action will be performed at the termination time.
	// +kcc:proto:field=google.cloud.compute.v1.Scheduling.termination_time
	TerminationTime *string `json:"terminationTime,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SchedulingNodeAffinity
type SchedulingNodeAffinity struct {
	// Corresponds to the label key of Node resource.
	// +kcc:proto:field=google.cloud.compute.v1.SchedulingNodeAffinity.key
	Key *string `json:"key,omitempty"`

	// Defines the operation of node selection. Valid operators are IN for affinity and NOT_IN for anti-affinity.
	//  Check the Operator enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SchedulingNodeAffinity.operator
	Operator *string `json:"operator,omitempty"`

	// Corresponds to the label values of Node resource.
	// +kcc:proto:field=google.cloud.compute.v1.SchedulingNodeAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SchedulingOnInstanceStopAction
type SchedulingOnInstanceStopAction struct {
	// If true, the contents of any attached Local SSD disks will be discarded else, the Local SSD data will be preserved when the instance is stopped at the end of the run duration/termination time.
	// +kcc:proto:field=google.cloud.compute.v1.SchedulingOnInstanceStopAction.discard_local_ssd
	DiscardLocalSsd *bool `json:"discardLocalSsd,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ServiceAccount
type ServiceAccount struct {
	// Email address of the service account.
	// +kcc:proto:field=google.cloud.compute.v1.ServiceAccount.email
	Email *string `json:"email,omitempty"`

	// The list of scopes to be made available for this service account.
	// +kcc:proto:field=google.cloud.compute.v1.ServiceAccount.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled. Enabled by default.
	// +kcc:proto:field=google.cloud.compute.v1.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	// Defines whether the instance has Secure Boot enabled. Disabled by default.
	// +kcc:proto:field=google.cloud.compute.v1.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Defines whether the instance has the vTPM enabled. Enabled by default.
	// +kcc:proto:field=google.cloud.compute.v1.ShieldedInstanceConfig.enable_vtpm
	EnableVTPM *bool `json:"enableVTPM,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ShieldedInstanceIntegrityPolicy
type ShieldedInstanceIntegrityPolicy struct {
	// Updates the integrity policy baseline using the measurements from the VM instance's most recent boot.
	// +kcc:proto:field=google.cloud.compute.v1.ShieldedInstanceIntegrityPolicy.update_auto_learn_policy
	UpdateAutoLearnPolicy *bool `json:"updateAutoLearnPolicy,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SubnetworkParams
type SubnetworkParams struct {
	// Tag keys/values directly bound to this resource. Tag keys and values have the same definition as resource manager tags. The field is allowed for INSERT only. The keys/values to set on the resource should be specified in either ID { : } or Namespaced format { : }. For example the following are valid inputs: * {"tagKeys/333" : "tagValues/444", "tagKeys/123" : "tagValues/456"} * {"123/environment" : "production", "345/abc" : "xyz"} Note: * Invalid combinations of ID & namespaced format is not supported. For instance: {"123/environment" : "tagValues/444"} is invalid.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkParams.resource_manager_tags
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Tags
type Tags struct {
	// Specifies a fingerprint for this request, which is essentially a hash of the tags' contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update tags. You must always provide an up-to-date fingerprint hash in order to update or change tags. To see the latest fingerprint, make get() request to the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Tags.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// An array of tags. Each tag must be 1-63 characters long, and comply with RFC1035.
	// +kcc:proto:field=google.cloud.compute.v1.Tags.items
	Items []string `json:"items,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.UpcomingMaintenance
type UpcomingMaintenance struct {
	// Indicates if the maintenance can be customer triggered.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.can_reschedule
	CanReschedule *bool `json:"canReschedule,omitempty"`

	// The latest time for the planned maintenance window to start. This timestamp value is in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.latest_window_start_time
	LatestWindowStartTime *string `json:"latestWindowStartTime,omitempty"`

	// Indicates whether the UpcomingMaintenance will be triggered on VM shutdown.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.maintenance_on_shutdown
	MaintenanceOnShutdown *bool `json:"maintenanceOnShutdown,omitempty"`

	// The reasons for the maintenance. Only valid for vms.
	//  Check the MaintenanceReasons enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.maintenance_reasons
	MaintenanceReasons []string `json:"maintenanceReasons,omitempty"`

	// Check the MaintenanceStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.maintenance_status
	MaintenanceStatus *string `json:"maintenanceStatus,omitempty"`

	// Defines the type of maintenance.
	//  Check the Type enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.type
	Type *string `json:"type,omitempty"`

	// The time by which the maintenance disruption will be completed. This timestamp value is in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.window_end_time
	WindowEndTime *string `json:"windowEndTime,omitempty"`

	// The current start time of the maintenance window. This timestamp value is in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.UpcomingMaintenance.window_start_time
	WindowStartTime *string `json:"windowStartTime,omitempty"`
}
