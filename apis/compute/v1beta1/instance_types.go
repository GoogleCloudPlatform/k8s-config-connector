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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.cloud.compute.v1.AttachedDisk
type ComputeInstanceAttachedDisk struct {
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance. This name can be used to reference the device for mounting, resizing, and so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disk-x, where x is a number assigned by Google Compute Engine. This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.device_name
	DeviceName *string `json:"deviceName,omitempty"`

	// Encrypts or decrypts a disk using a customer-supplied encryption key. If you are creating a new disk, this field encrypts the new disk using an encryption key that you provide. If you are attaching an existing disk that is already encrypted, this field decrypts the disk using the customer-supplied encryption key. If you encrypt a disk using a customer-supplied key, you must provide the same key again when you attempt to use this resource at a later time. For example, you must provide the key when you create a snapshot or an image from the disk or when you attach the disk to a virtual machine instance. If you do not provide an encryption key, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Note: Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt disks in a managed instance group. You cannot create VMs that have disks with customer-supplied keys using the bulk insert method.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.disk_encryption_key
	DiskEncryptionKey *ComputeInstanceDiskEncryptionKey `json:"diskEncryptionKeyRaw,omitempty"`

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.disk_encryption_key
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the disk in READ_WRITE mode.
	//  Check the Mode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.mode
	Mode *string `json:"mode,omitempty"`

	// Specifies a valid partial or full URL to an existing Persistent Disk resource. When creating a new instance boot disk, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required. If desired, you can also attach existing non-root persistent disks using this property. This field is only applicable for persistent disks. Note that for InstanceTemplate, specify the disk name for zonal disk, and the URL for regional disk.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.source
	SourceDiskRef *ComputeDiskRef `json:"sourceDiskRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AttachedDisk
type ComputeInstanceBootDisk struct {
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance. This name can be used to reference the device for mounting, resizing, and so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disk-x, where x is a number assigned by Google Compute Engine. This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.device_name
	DeviceName *string `json:"deviceName,omitempty"`

	// Encrypts or decrypts a disk using a customer-supplied encryption key. If you are creating a new disk, this field encrypts the new disk using an encryption key that you provide. If you are attaching an existing disk that is already encrypted, this field decrypts the disk using the customer-supplied encryption key. If you encrypt a disk using a customer-supplied key, you must provide the same key again when you attempt to use this resource at a later time. For example, you must provide the key when you create a snapshot or an image from the disk or when you attach the disk to a virtual machine instance. If you do not provide an encryption key, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Note: Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt disks in a managed instance group. You cannot create VMs that have disks with customer-supplied keys using the bulk insert method.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.disk_encryption_key
	DiskEncryptionKey *ComputeInstanceDiskEncryptionKey `json:"diskEncryptionKeyRaw,omitempty"`

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty"`

	// [Input Only] Specifies the parameters for a new disk that will be created alongside the new instance. Use initialization parameters to create boot disks or local SSDs attached to the new instance. This property is mutually exclusive with the source property; you can only define one or the other, but not both.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.initialize_params
	InitializeParams *ComputeInstanceAttachedDiskInitializeParams `json:"initializeParams,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.disk_encryption_key
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the disk in READ_WRITE mode.
	//  Check the Mode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.mode
	Mode *string `json:"mode,omitempty"`

	// Specifies a valid partial or full URL to an existing Persistent Disk resource. When creating a new instance boot disk, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required. If desired, you can also attach existing non-root persistent disks using this property. This field is only applicable for persistent disks. Note that for InstanceTemplate, specify the disk name for zonal disk, and the URL for regional disk.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDisk.source
	SourceDiskRef *ComputeDiskRef `json:"sourceDiskRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.CustomerEncryptionKey
type ComputeInstanceDiskEncryptionKey struct {
	// Value of the field. Cannot be used if 'valueFrom' is specified.
	Value *string `json:"value,omitempty"`

	// Source for the field's value. Cannot be used if 'value' is specified.
	ValueFrom *ComputeInstanceDiskEncryptionKeyValueFrom `json:"valueFrom,omitempty"`
}

type ComputeInstanceDiskEncryptionKeyValueFrom struct {
	// Reference to a value with the given key in the given Secret in the resource's namespace.
	SecretKeyRef *SecretKeySelector `json:"secretKeyRef,omitempty"`
}

type SecretKeySelector struct {
	// Key that identifies the value to be extracted.
	Key string `json:"key"`
	// Name of the Secret to extract a value from.
	Name string `json:"name"`
}

// +kcc:proto=google.cloud.compute.v1.AdvancedMachineFeatures
type AdvancedMachineFeatures struct {
	// Whether to enable nested virtualization or not (default is false).
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.enable_nested_virtualization
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`

	// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.threads_per_core
	ThreadsPerCore *int32 `json:"threadsPerCore,omitempty"`

	// The number of physical cores to expose to an instance. Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance's nominal CPU count and the underlying platform's SMT width.
	// +kcc:proto:field=google.cloud.compute.v1.AdvancedMachineFeatures.visible_core_count
	VisibleCoreCount *int32 `json:"visibleCoreCount,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AttachedDiskInitializeParams
type ComputeInstanceAttachedDiskInitializeParams struct {
	// An optional description. Provide this property when creating the disk.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.description
	Description *string `json:"description,omitempty"`

	// Specifies the disk name. If not specified, the default is to use the name of the instance. If a disk with the same name already exists in the given region, the existing disk is attached to the new instance and the new disk is not created.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.disk_name
	DiskName *string `json:"diskName,omitempty"`

	// Specifies the size of the disk in base-2 GB. The size must be at least 10 GB. If you specify a sourceImage, which is required for boot disks, the default size is the size of the sourceImage. If you do not specify a sourceImage, the default disk size is 500 GB.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.disk_size_gb
	DiskSizeGB *int64 `json:"size,omitempty"`

	// Specifies the disk type to use to create the instance. If not specified, the default is pd-standard, specified using the full URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/pd-standard For a full list of acceptable values, see Persistent disk types. If you specify this field when creating a VM, you can provide either the full or partial URL. For example, the following values are valid: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/diskType - projects/project/zones/zone/diskTypes/diskType - zones/zone/diskTypes/diskType If you specify this field when creating or updating an instance template or all-instances configuration, specify the type of the disk, not the URL. For example: pd-standard.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.disk_type
	DiskType *string `json:"type,omitempty"`

	// Labels to apply to this disk. These can be later modified by the disks.setLabels method. This field is only applicable for persistent disks.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Resource manager tags to be bound to the disk. Tag keys and values have the same definition as resource manager tags. Keys must be in the format `tagKeys/{tag_key_id}`, and values are in the format `tagValues/456`. The field is ignored (both PUT & PATCH) when empty.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.resource_manager_tags
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`

	// The source image to create this disk. When creating a new instance boot disk, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family If the source image is deleted later, this field will not be set.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.source_image
	SourceImageRef *ComputeImageRef `json:"sourceImageRef,omitempty"`

	// The source snapshot to create this disk. When creating a new instance boot disk, one of initializeParams.sourceSnapshot or initializeParams.sourceImage or disks.source is required. To create a disk with a snapshot that you created, specify the snapshot name in the following format: global/snapshots/my-backup If the source snapshot is deleted later, this field will not be set. Note: You cannot create VMs in bulk using a snapshot as the source. Use an image instead when you create VMs using the bulk insert method.
	// +kcc:proto:field=google.cloud.compute.v1.AttachedDiskInitializeParams.source_snapshot
	SourceSnapshotRef *ComputeSnapshotRef `json:"sourceSnapshotRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkInterface
type ComputeInstanceNetworkInterface struct {
	// An array of configurations for this interface. Currently, only one access config, ONE_TO_ONE_NAT, is supported. If there are no accessConfigs specified, then this instance will have no external internet access.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.access_configs
	AccessConfigs []ComputeInstanceAccessConfig `json:"accessConfig,omitempty"`

	// An array of alias IP ranges for this network interface. You can only specify this field for network interfaces in VPC networks.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.alias_ip_ranges
	AliasIPRanges []AliasIPRange `json:"aliasIpRange,omitempty"`

	// The prefix length of the primary internal IPv6 range.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.internal_ipv6_prefix_length
	InternalIPV6PrefixLength *int32 `json:"internalIpv6PrefixLength,omitempty"`

	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.ipv6_access_configs
	IPV6AccessConfigs []ComputeInstanceAccessConfig `json:"ipv6AccessConfig,omitempty"`

	// [Output Only] One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork. Valid only if stackType is IPV4_IPV6.
	//  Check the Ipv6AccessType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.ipv6_access_type
	IPV6AccessType *string `json:"ipv6AccessType,omitempty"`

	// An IPv6 internal network address for this network interface. To use a static internal IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.ipv6_address
	IPV6Address *string `json:"ipv6Address,omitempty"`

	// [Output Only] The name of the network interface, which is generated by the server. For a VM, the network interface uses the nicN naming format. Where N is a value between 0 and 7. The default interface value is nic0.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.name
	Name *string `json:"name,omitempty"`

	// URL of the VPC network resource for this instance. When creating an instance, if neither the network nor the subnetwork is specified, the default network global/networks/default is used. If the selected project doesn't have the default network, you must specify a network or subnet. If the network is not specified but the subnetwork is specified, the network is inferred. If you specify this property, you can specify the network as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/global/networks/ network - projects/project/global/networks/network - global/networks/default
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.network
	NetworkRef *ComputeNetworkRef `json:"networkRef,omitempty"`

	// An IPv4 internal IP address to assign to the instance for this network interface. If not specified by the user, an unused internal IP is assigned by the system.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.network_i_p
	NetworkIPRef *refs.ComputeAddressRef `json:"networkIpRef,omitempty"`

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

	SubnetworkProject *string `json:"subnetworkProject,omitempty"`

	// The URL of the Subnetwork resource for this instance. If the network resource is in legacy mode, do not specify this field. If the network is in auto subnet mode, specifying the subnetwork is optional. If the network is in custom subnet mode, specifying the subnetwork is required. If you specify this field, you can specify the subnetwork as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/regions/region /subnetworks/subnetwork - regions/region/subnetworks/subnetwork
	// +kcc:proto:field=google.cloud.compute.v1.NetworkInterface.subnetwork
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AccessConfig
type ComputeInstanceAccessConfig struct {
	// Applies to ipv6AccessConfigs only. The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. To use a static external IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an external IPv6 address from the instance's subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.external_ipv6
	ExternalIPV6 *string `json:"externalIpv6,omitempty"`

	// Applies to ipv6AccessConfigs only. The prefix length of the external IPv6 range.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.external_ipv6_prefix_length
	ExternalIPV6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty"`

	// The name of this access configuration. In accessConfigs (IPv4), the default and recommended name is External NAT, but you can use any arbitrary string, such as My external IP or Network Access. In ipv6AccessConfigs, the recommend name is External IPv6.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.name
	Name *string `json:"name,omitempty"`

	// Applies to accessConfigs (IPv4) only. An external IP address associated with this instance. Specify an unused static external IP address available to the project or leave this field undefined to use an IP from a shared ephemeral IP address pool. If you specify a static external IP address, it must live in the same region as the zone of the instance.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.nat_i_p
	NatIPRef *refs.ComputeAddressRef `json:"natIpRef,omitempty"`

	// This signifies the networking tier used for configuring this access configuration and can only take the following values: PREMIUM, STANDARD. If an AccessConfig is specified without a valid external IP address, an ephemeral IP will be created with this networkTier. If an AccessConfig with a valid external IP address is specified, it must match that of the networkTier associated with the Address resource owning that IP.
	//  Check the NetworkTier enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.network_tier
	NetworkTier *string `json:"networkTier,omitempty"`

	// The DNS domain name for the public PTR record. You can set this field only if the `setPublicPtr` field is enabled in accessConfig. If this field is unspecified in ipv6AccessConfig, a default PTR record will be created for first IP in associated external IPv6 range.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.public_ptr_domain_name
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`

	// The type of configuration. In accessConfigs (IPv4), the default and only option is ONE_TO_ONE_NAT. In ipv6AccessConfigs, the default and only option is DIRECT_IPV6.
	//  Check the Type enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AccessConfig.type
	// Type *string `json:"type,omitempty"` // Removed because it seems not to be in KRM spec for accessConfig
}

type ComputeInstanceParams struct {
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`
}

type ComputeInstanceReservationAffinity struct {
	SpecificReservation *ComputeInstanceSpecificReservation `json:"specificReservation,omitempty"`
	Type                *string                             `json:"type,omitempty"`
}

type ComputeInstanceSpecificReservation struct {
	Key    *string  `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}

type ComputeInstanceServiceAccount struct {
	Scopes            []string                   `json:"scopes,omitempty"`
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

type ComputeInstanceScratchDisk struct {
	Interface *string `json:"interface,omitempty"`
	Size      *int64  `json:"size,omitempty"`
}

type ComputeInstanceShieldedInstanceConfig struct {
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
	EnableSecureBoot          *bool `json:"enableSecureBoot,omitempty"`
	EnableVtpm                *bool `json:"enableVtpm,omitempty"`
}

// ComputeInstanceSpec defines the desired state of ComputeInstance
// +kcc:spec:proto=google.cloud.compute.v1.Instance
type ComputeInstanceSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// Controls for advanced machine-related behavior features.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.advanced_machine_features
	AdvancedMachineFeatures *AdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	// List of disks attached to the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.disks
	AttachedDisk []ComputeInstanceAttachedDisk `json:"attachedDisk,omitempty"`

	// Immutable. The boot disk for the instance.
	BootDisk *ComputeInstanceBootDisk `json:"bootDisk,omitempty"`

	// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
	// +kcc:proto:field=google.cloud.compute.v1.Instance.can_ip_forward
	CanIPForward *bool `json:"canIpForward,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.confidential_instance_config
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`

	// Whether the resource should be protected against deletion.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.deletion_protection
	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.description
	Description *string `json:"description,omitempty"`

	DesiredStatus *string `json:"desiredStatus,omitempty"`

	// Enables display device for the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.display_device
	// DisplayDevice *DisplayDevice `json:"displayDevice,omitempty"`
	// KRM has enableDisplay boolean instead of DisplayDevice struct.
	EnableDisplay *bool `json:"enableDisplay,omitempty"`

	// A list of the type and count of accelerator cards attached to the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.guest_accelerators
	// GuestAccelerators []AcceleratorConfig `json:"guestAccelerators,omitempty"`
	// KRM has guestAccelerator (singular) and structure is different (count, type) vs (accelerator_count, accelerator_type)
	GuestAccelerator []ComputeInstanceGuestAccelerator `json:"guestAccelerator,omitempty"`

	// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.hostname
	Hostname *string `json:"hostname,omitempty"`

	InstanceTemplateRef *ComputeInstanceTemplateRef `json:"instanceTemplateRef,omitempty"`

	// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// The metadata key/value pairs assigned to this instance. This includes metadata keys that were explicitly defined for the instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.metadata
	// Metadata *Metadata `json:"metadata,omitempty"`
	// KRM has metadata as list of key/value items.
	Metadata []ComputeInstanceMetadataItem `json:"metadata,omitempty"`

	MetadataStartupScript *string `json:"metadataStartupScript,omitempty"`

	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
	// +kcc:proto:field=google.cloud.compute.v1.Instance.min_cpu_platform
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty"`

	// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.network_interfaces
	NetworkInterface []ComputeInstanceNetworkInterface `json:"networkInterface,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.network_performance_config
	NetworkPerformanceConfig *NetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty"`

	// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.params
	Params *ComputeInstanceParams `json:"params,omitempty"`

	// Specifies the reservations that this instance can consume from.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.reservation_affinity
	ReservationAffinity *ComputeInstanceReservationAffinity `json:"reservationAffinity,omitempty"`

	// The ComputeInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Resource policies applied to this instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.resource_policies
	ResourcePolicies []ComputeResourcePolicyRef `json:"resourcePolicies,omitempty"`

	// Sets the scheduling options for this instance.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.scheduling
	Scheduling *Scheduling `json:"scheduling,omitempty"`

	ScratchDisk []ComputeInstanceScratchDisk `json:"scratchDisk,omitempty"`

	// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.service_accounts
	// ServiceAccounts []ServiceAccount `json:"serviceAccounts,omitempty"`
	// KRM has serviceAccount (singular)
	ServiceAccount *ComputeInstanceServiceAccount `json:"serviceAccount,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.shielded_instance_config
	ShieldedInstanceConfig *ComputeInstanceShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.tags
	// Tags *Tags `json:"tags,omitempty"`
	// KRM has tags as list of strings
	Tags []string `json:"tags,omitempty"`

	// [Output Only] URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.zone
	Zone *string `json:"zone,omitempty"`
}

type ComputeInstanceGuestAccelerator struct {
	Count *int32  `json:"count,omitempty"`
	Type  *string `json:"type,omitempty"`
}

type ComputeInstanceMetadataItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// ComputeInstanceStatus defines the config connector machine state of ComputeInstance
type ComputeInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInstanceObservedState `json:"observedState,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.cpu_platform
	CPUPlatform *string `json:"cpuPlatform,omitempty"`

	// [Output Only] An optional, human-readable explanation of the status.
	// +kcc:proto:field=google.cloud.compute.v1.Instance.status_message
	CurrentStatus *string `json:"currentStatus,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.id
	InstanceID *string `json:"instanceId,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.metadata.fingerprint
	MetadataFingerprint *string `json:"metadataFingerprint,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.Instance.tags.fingerprint
	TagsFingerprint *string `json:"tagsFingerprint,omitempty"`
}

// ComputeInstanceObservedState is the state of the ComputeInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Instance
type ComputeInstanceObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstance;gcpcomputeinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstance is the Schema for the ComputeInstance API
// +k8s:openapi-gen=true
type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceSpec   `json:"spec,omitempty"`
	Status ComputeInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceList contains a list of ComputeInstance
type ComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstance{}, &ComputeInstanceList{})
}

type ComputeDiskRef struct {
	/* The ComputeDisk selflink in the form "projects/{{project}}/zones/{{zone}}/disks/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeDisk` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeDisk` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeSnapshotRef struct {
	/* The ComputeSnapshot selflink in the form "projects/{{project}}/global/snapshots/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeSnapshot` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeSnapshot` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeInstanceTemplateRef struct {
	/* The ComputeInstanceTemplate selflink in the form "projects/{{project}}/global/instanceTemplates/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeInstanceTemplate` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeInstanceTemplate` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeResourcePolicyRef struct {
	/* The ComputeResourcePolicy selflink in the form "projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeResourcePolicy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeResourcePolicy` resource. */
	Namespace string `json:"namespace,omitempty"`
}
