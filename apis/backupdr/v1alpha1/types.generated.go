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


// +kcc:proto=google.cloud.backupdr.v1.AcceleratorConfig
type AcceleratorConfig struct {
	// Optional. Full or partial URL of the accelerator type resource to attach to
	//  this instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.AcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Optional. The number of the guest accelerator cards exposed to this
	//  instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.AccessConfig
type AccessConfig struct {
	// Optional. In accessConfigs (IPv4), the
	//   default and only option is ONE_TO_ONE_NAT. In
	//   ipv6AccessConfigs, the default and only option is
	//   DIRECT_IPV6.
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.type
	Type *string `json:"type,omitempty"`

	// Optional. The name of this access configuration.
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.name
	Name *string `json:"name,omitempty"`

	// Optional. The external IP address of this access configuration.
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`

	// Optional. The external IPv6 address of this access configuration.
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.external_ipv6
	ExternalIpv6 *string `json:"externalIpv6,omitempty"`

	// Optional. The prefix length of the external IPv6 range.
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.external_ipv6_prefix_length
	ExternalIpv6PrefixLength *int32 `json:"externalIpv6PrefixLength,omitempty"`

	// Optional. Specifies whether a public DNS 'PTR' record should be created to
	//  map the external IP address of the instance to a DNS domain name.
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.set_public_ptr
	SetPublicPtr *bool `json:"setPublicPtr,omitempty"`

	// Optional. The DNS domain name for the public PTR record.
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.public_ptr_domain_name
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`

	// Optional. This signifies the networking tier used for configuring this
	//  access
	// +kcc:proto:field=google.cloud.backupdr.v1.AccessConfig.network_tier
	NetworkTier *string `json:"networkTier,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.AliasIpRange
type AliasIpRange struct {
	// Optional. The IP alias ranges to allocate for this interface.
	// +kcc:proto:field=google.cloud.backupdr.v1.AliasIpRange.ip_cidr_range
	IPCidrRange *string `json:"ipCidrRange,omitempty"`

	// Optional. The name of a subnetwork secondary IP range from which to
	//  allocate an IP alias range. If not specified, the primary range of the
	//  subnetwork is used.
	// +kcc:proto:field=google.cloud.backupdr.v1.AliasIpRange.subnetwork_range_name
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.AttachedDisk
type AttachedDisk struct {
	// Optional. Specifies the parameters to initialize this disk.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.initialize_params
	InitializeParams *AttachedDisk_InitializeParams `json:"initializeParams,omitempty"`

	// Optional. This is used as an identifier for the disks. This is the unique
	//  name has to provided to modify disk parameters like disk_name and
	//  replica_zones (in case of RePDs)
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.device_name
	DeviceName *string `json:"deviceName,omitempty"`

	// Optional. Type of the resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.kind
	Kind *string `json:"kind,omitempty"`

	// Specifies the type of the disk.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.disk_type_deprecated
	DiskTypeDeprecated *string `json:"diskTypeDeprecated,omitempty"`

	// Optional. The mode in which to attach this disk.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Specifies a valid partial or full URL to an existing Persistent
	//  Disk resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.source
	Source *string `json:"source,omitempty"`

	// Optional. A zero-based index to this disk, where 0 is reserved for the
	//  boot disk.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.index
	Index *int64 `json:"index,omitempty"`

	// Optional. Indicates that this is a boot disk. The virtual machine will use
	//  the first partition of the disk for its root filesystem.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.boot
	Boot *bool `json:"boot,omitempty"`

	// Optional. Specifies whether the disk will be auto-deleted when the instance
	//  is deleted (but not when the disk is detached from the instance).
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.auto_delete
	AutoDelete *bool `json:"autoDelete,omitempty"`

	// Optional. Any valid publicly visible licenses.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.license
	License []string `json:"license,omitempty"`

	// Optional. Specifies the disk interface to use for attaching this disk.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.disk_interface
	DiskInterface *string `json:"diskInterface,omitempty"`

	// Optional. A list of features to enable on the guest operating system.
	//  Applicable only for bootable images.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.guest_os_feature
	GuestOsFeature []GuestOsFeature `json:"guestOsFeature,omitempty"`

	// Optional. Encrypts or decrypts a disk using a customer-supplied
	//  encryption key.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.disk_encryption_key
	DiskEncryptionKey *CustomerEncryptionKey `json:"diskEncryptionKey,omitempty"`

	// Optional. The size of the disk in GB.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Optional. Specifies the type of the disk.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.AttachedDisk.InitializeParams
type AttachedDisk_InitializeParams struct {
	// Optional. Specifies the disk name. If not specified, the default is to
	//  use the name of the instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.InitializeParams.disk_name
	DiskName *string `json:"diskName,omitempty"`

	// Optional. URL of the zone where the disk should be created.
	//  Required for each regional disk associated with the instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.InitializeParams.replica_zones
	ReplicaZones []string `json:"replicaZones,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.Backup
type Backup struct {

	// Optional. Resource labels to represent user provided metadata.
	//  No labels currently defined.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The backup can not be deleted before this time.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.enforced_retention_end_time
	EnforcedRetentionEndTime *string `json:"enforcedRetentionEndTime,omitempty"`

	// Optional. When this backup is automatically expired.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Optional. Server specified ETag to prevent updates from overwriting each
	//  other.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The list of BackupLocks taken by the accessor Backup Appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.backup_appliance_locks
	BackupApplianceLocks []BackupLock `json:"backupApplianceLocks,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.Backup.GCPBackupPlanInfo
type Backup_GCPBackupPlanInfo struct {
	// Resource name of backup plan by which workload is protected at the time
	//  of the backup.
	//  Format:
	//  projects/{project}/locations/{location}/backupPlans/{backupPlanId}
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.GCPBackupPlanInfo.backup_plan
	BackupPlan *string `json:"backupPlan,omitempty"`

	// The rule id of the backup plan which triggered this backup in case of
	//  scheduled backup or used for
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.GCPBackupPlanInfo.backup_plan_rule_id
	BackupPlanRuleID *string `json:"backupPlanRuleID,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupApplianceBackupProperties
type BackupApplianceBackupProperties struct {

	// Optional. The earliest timestamp of data available in this Backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupProperties.recovery_range_start_time
	RecoveryRangeStartTime *string `json:"recoveryRangeStartTime,omitempty"`

	// Optional. The latest timestamp of data available in this Backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupProperties.recovery_range_end_time
	RecoveryRangeEndTime *string `json:"recoveryRangeEndTime,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupApplianceLockInfo
type BackupApplianceLockInfo struct {
	// Required. The ID of the backup/recovery appliance that created this lock.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceLockInfo.backup_appliance_id
	BackupApplianceID *int64 `json:"backupApplianceID,omitempty"`

	// Required. The name of the backup/recovery appliance that created this lock.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceLockInfo.backup_appliance_name
	BackupApplianceName *string `json:"backupApplianceName,omitempty"`

	// Required. The reason for the lock: e.g. MOUNT/RESTORE/BACKUP/etc.  The
	//  value of this string is only meaningful to the client and it is not
	//  interpreted by the BackupVault service.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceLockInfo.lock_reason
	LockReason *string `json:"lockReason,omitempty"`

	// The job name on the backup/recovery appliance that created this lock.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceLockInfo.job_name
	JobName *string `json:"jobName,omitempty"`

	// The image name that depends on this Backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceLockInfo.backup_image
	BackupImage *string `json:"backupImage,omitempty"`

	// The SLA on the backup/recovery appliance that owns the lock.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceLockInfo.sla_id
	SlaID *int64 `json:"slaID,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupLock
type BackupLock struct {
	// Required. The time after which this lock is not considered valid and will
	//  no longer protect the Backup from deletion.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupLock.lock_until_time
	LockUntilTime *string `json:"lockUntilTime,omitempty"`

	// If the client is a backup and recovery appliance, this
	//  contains metadata about why the lock exists.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupLock.backup_appliance_lock_info
	BackupApplianceLockInfo *BackupApplianceLockInfo `json:"backupApplianceLockInfo,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ComputeInstanceBackupProperties
type ComputeInstanceBackupProperties struct {
	// An optional text description for the instances that are created from these
	//  properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.description
	Description *string `json:"description,omitempty"`

	// A list of tags to apply to the instances that are created from these
	//  properties. The tags identify valid sources or targets for network
	//  firewalls. The setTags method can modify this list of tags. Each tag within
	//  the list must comply with RFC1035 (https://www.ietf.org/rfc/rfc1035.txt).
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.tags
	Tags *Tags `json:"tags,omitempty"`

	// The machine type to use for instances that are created from these
	//  properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Enables instances created based on these properties to send packets with
	//  source IP addresses other than their own and receive packets with
	//  destination IP addresses other than their own. If these instances will be
	//  used as an IP gateway or it will be set as the next-hop in a Route
	//  resource, specify `true`. If unsure, leave this set to `false`. See the
	//  https://cloud.google.com/vpc/docs/using-routes#canipforward
	//  documentation for more information.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.can_ip_forward
	CanIPForward *bool `json:"canIPForward,omitempty"`

	// An array of network access configurations for this interface.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.network_interface
	NetworkInterface []NetworkInterface `json:"networkInterface,omitempty"`

	// An array of disks that are associated with the instances that are created
	//  from these properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.disk
	Disk []AttachedDisk `json:"disk,omitempty"`

	// The metadata key/value pairs to assign to instances that are created from
	//  these properties. These pairs can consist of custom metadata or predefined
	//  keys. See https://cloud.google.com/compute/docs/metadata/overview for more
	//  information.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.metadata
	Metadata *Metadata `json:"metadata,omitempty"`

	// A list of service accounts with specified scopes. Access tokens for these
	//  service accounts are available to the instances that are created from
	//  these properties. Use metadata queries to obtain the access tokens for
	//  these instances.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.service_account
	ServiceAccount []ServiceAccount `json:"serviceAccount,omitempty"`

	// Specifies the scheduling options for the instances that are created from
	//  these properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.scheduling
	Scheduling *Scheduling `json:"scheduling,omitempty"`

	// A list of guest accelerator cards' type and count to use for instances
	//  created from these properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.guest_accelerator
	GuestAccelerator []AcceleratorConfig `json:"guestAccelerator,omitempty"`

	// Minimum cpu/platform to be used by instances. The instance may be
	//  scheduled on the specified or newer cpu/platform. Applicable values are the
	//  friendly names of CPU platforms, such as
	//  `minCpuPlatform: Intel Haswell` or `minCpuPlatform: Intel Sandy Bridge`.
	//  For more information, read
	//  https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.min_cpu_platform
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	// KeyRevocationActionType of the instance. Supported options are "STOP" and
	//  "NONE". The default value is "NONE" if it is not specified.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.key_revocation_action_type
	KeyRevocationActionType *string `json:"keyRevocationActionType,omitempty"`

	// The source instance used to create this backup. This can be a partial or
	//  full URL to the resource. For example, the following are valid values:
	//    -https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance
	//    -projects/project/zones/zone/instances/instance
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.source_instance
	SourceInstance *string `json:"sourceInstance,omitempty"`

	// Labels to apply to instances that are created from these properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.CustomerEncryptionKey
type CustomerEncryptionKey struct {
	// Optional. Specifies a 256-bit customer-supplied
	//  encryption key.
	// +kcc:proto:field=google.cloud.backupdr.v1.CustomerEncryptionKey.raw_key
	RawKey *string `json:"rawKey,omitempty"`

	// Optional. RSA-wrapped 2048-bit
	//  customer-supplied encryption key to either encrypt or decrypt this
	//  resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.CustomerEncryptionKey.rsa_encrypted_key
	RsaEncryptedKey *string `json:"rsaEncryptedKey,omitempty"`

	// Optional. The name of the encryption key that is stored in Google Cloud
	//  KMS.
	// +kcc:proto:field=google.cloud.backupdr.v1.CustomerEncryptionKey.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Optional. The service account being used for the encryption request for the
	//  given KMS key. If absent, the Compute Engine default service account is
	//  used.
	// +kcc:proto:field=google.cloud.backupdr.v1.CustomerEncryptionKey.kms_key_service_account
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.Entry
type Entry struct {
	// Optional. Key for the metadata entry.
	// +kcc:proto:field=google.cloud.backupdr.v1.Entry.key
	Key *string `json:"key,omitempty"`

	// Optional. Value for the metadata entry. These are free-form strings, and
	//  only have meaning as interpreted by the image running in the instance. The
	//  only restriction placed on values is that their size must be less than
	//  or equal to 262144 bytes (256 KiB).
	// +kcc:proto:field=google.cloud.backupdr.v1.Entry.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.GuestOsFeature
type GuestOsFeature struct {
	// The ID of a supported feature.
	// +kcc:proto:field=google.cloud.backupdr.v1.GuestOsFeature.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.Metadata
type Metadata struct {
	// Optional. Array of key/value pairs. The total size of all keys and values
	//  must be less than 512 KB.
	// +kcc:proto:field=google.cloud.backupdr.v1.Metadata.items
	Items []Entry `json:"items,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.NetworkInterface
type NetworkInterface struct {
	// Optional. URL of the VPC network resource for this instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.network
	Network *string `json:"network,omitempty"`

	// Optional. The URL of the Subnetwork resource for this instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Optional. An IPv4 internal IP address to assign to the instance for this
	//  network interface. If not specified by the user, an unused internal IP is
	//  assigned by the system.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Optional. An IPv6 internal network address for this network interface. To
	//  use a static internal IP address, it must be unused and in the same region
	//  as the instance's zone. If not specified, Google Cloud will automatically
	//  assign an internal IPv6 address from the instance's subnetwork.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.ipv6_address
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	// Optional. The prefix length of the primary internal IPv6 range.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.internal_ipv6_prefix_length
	InternalIpv6PrefixLength *int32 `json:"internalIpv6PrefixLength,omitempty"`

	// Optional. An array of configurations for this interface. Currently, only
	//  one access config,ONE_TO_ONE_NAT is supported. If there are no
	//  accessConfigs specified, then this instance will have
	//  no external internet access.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.access_configs
	AccessConfigs []AccessConfig `json:"accessConfigs,omitempty"`

	// Optional. An array of IPv6 access configurations for this interface.
	//  Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there
	//  is no ipv6AccessConfig specified, then this instance will
	//  have no external IPv6 Internet access.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.ipv6_access_configs
	Ipv6AccessConfigs []AccessConfig `json:"ipv6AccessConfigs,omitempty"`

	// Optional. An array of alias IP ranges for this network interface.
	//  You can only specify this field for network interfaces in VPC networks.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.alias_ip_ranges
	AliasIPRanges []AliasIpRange `json:"aliasIPRanges,omitempty"`

	// The stack type for this network interface.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.stack_type
	StackType *string `json:"stackType,omitempty"`

	// Optional. [Output Only] One of EXTERNAL, INTERNAL to indicate whether the
	//  IP can be accessed from the Internet. This field is always inherited from
	//  its subnetwork.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.ipv6_access_type
	Ipv6AccessType *string `json:"ipv6AccessType,omitempty"`

	// Optional. The networking queue count that's specified by users for the
	//  network interface. Both Rx and Tx queues will be set to this number. It'll
	//  be empty if not specified by the users.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.queue_count
	QueueCount *int32 `json:"queueCount,omitempty"`

	// Optional. The type of vNIC to be used on this interface. This may be gVNIC
	//  or VirtioNet.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.nic_type
	NicType *string `json:"nicType,omitempty"`

	// Optional. The URL of the network attachment that this interface should
	//  connect to in the following format:
	//  projects/{project_number}/regions/{region_name}/networkAttachments/{network_attachment_name}.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.network_attachment
	NetworkAttachment *string `json:"networkAttachment,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.Scheduling
type Scheduling struct {
	// Optional. Defines the maintenance behavior for this instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.on_host_maintenance
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty"`

	// Optional. Specifies whether the instance should be automatically restarted
	//  if it is terminated by Compute Engine (not terminated by a user).
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.automatic_restart
	AutomaticRestart *bool `json:"automaticRestart,omitempty"`

	// Optional. Defines whether the instance is preemptible.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Optional. A set of node affinity and anti-affinity configurations.
	//  Overrides reservationAffinity.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.node_affinities
	NodeAffinities []Scheduling_NodeAffinity `json:"nodeAffinities,omitempty"`

	// Optional. The minimum number of virtual CPUs this instance will consume
	//  when running on a sole-tenant node.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.min_node_cpus
	MinNodeCpus *int32 `json:"minNodeCpus,omitempty"`

	// Optional. Specifies the provisioning model of the instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.provisioning_model
	ProvisioningModel *string `json:"provisioningModel,omitempty"`

	// Optional. Specifies the termination action for the instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.instance_termination_action
	InstanceTerminationAction *string `json:"instanceTerminationAction,omitempty"`

	// Optional. Specifies the maximum amount of time a Local Ssd Vm should wait
	//  while recovery of the Local Ssd state is attempted. Its value should be in
	//  between 0 and 168 hours with hour granularity and the default value being 1
	//  hour.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.local_ssd_recovery_timeout
	LocalSsdRecoveryTimeout *SchedulingDuration `json:"localSsdRecoveryTimeout,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.Scheduling.NodeAffinity
type Scheduling_NodeAffinity struct {
	// Optional. Corresponds to the label key of Node resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.NodeAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Defines the operation of node selection.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.NodeAffinity.operator
	Operator *string `json:"operator,omitempty"`

	// Optional. Corresponds to the label values of Node resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.NodeAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.SchedulingDuration
type SchedulingDuration struct {
	// Optional. Span of time at a resolution of a second.
	// +kcc:proto:field=google.cloud.backupdr.v1.SchedulingDuration.seconds
	Seconds *int64 `json:"seconds,omitempty"`

	// Optional. Span of time that's a fraction of a second at nanosecond
	//  resolution.
	// +kcc:proto:field=google.cloud.backupdr.v1.SchedulingDuration.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ServiceAccount
type ServiceAccount struct {
	// Optional. Email address of the service account.
	// +kcc:proto:field=google.cloud.backupdr.v1.ServiceAccount.email
	Email *string `json:"email,omitempty"`

	// Optional. The list of scopes to be made available for this service account.
	// +kcc:proto:field=google.cloud.backupdr.v1.ServiceAccount.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ServiceLockInfo
type ServiceLockInfo struct {
}

// +kcc:proto=google.cloud.backupdr.v1.Tags
type Tags struct {
	// Optional. An array of tags. Each tag must be 1-63 characters long, and
	//  comply with RFC1035.
	// +kcc:proto:field=google.cloud.backupdr.v1.Tags.items
	Items []string `json:"items,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.AttachedDisk
type AttachedDiskObservedState struct {
	// Optional. Output only. The state of the disk.
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.saved_state
	SavedState *string `json:"savedState,omitempty"`

	// Optional. Output only. The URI of the disk type resource. For example:
	//  projects/project/zones/zone/diskTypes/pd-standard or pd-ssd
	// +kcc:proto:field=google.cloud.backupdr.v1.AttachedDisk.disk_type
	DiskType *string `json:"diskType,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.Backup
type BackupObservedState struct {
	// Output only. Identifier. Name of the backup to create. It must have the
	//  format`"projects/<project>/locations/<location>/backupVaults/<backupvault>/dataSources/{datasource}/backups/{backup}"`.
	//  `{backup}` cannot be changed after creation. It must be between 3-63
	//  characters long and must be unique within the datasource.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.name
	Name *string `json:"name,omitempty"`

	// Output only. The description of the Backup instance (2048 characters or
	//  less).
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.description
	Description *string `json:"description,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The point in time when this backup was captured from the
	//  source.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.consistency_time
	ConsistencyTime *string `json:"consistencyTime,omitempty"`

	// Output only. The Backup resource instance state.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The list of BackupLocks taken by the service to prevent the
	//  deletion of the backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.service_locks
	ServiceLocks []BackupLock `json:"serviceLocks,omitempty"`

	// Output only. Compute Engine specific backup properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.compute_instance_backup_properties
	ComputeInstanceBackupProperties *ComputeInstanceBackupProperties `json:"computeInstanceBackupProperties,omitempty"`

	// Output only. Backup Appliance specific backup properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.backup_appliance_backup_properties
	BackupApplianceBackupProperties *BackupApplianceBackupProperties `json:"backupApplianceBackupProperties,omitempty"`

	// Output only. Type of the backup, unspecified, scheduled or ondemand.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// Output only. Configuration for a Google Cloud resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.gcp_backup_plan_info
	GcpBackupPlanInfo *Backup_GCPBackupPlanInfo `json:"gcpBackupPlanInfo,omitempty"`

	// Output only. source resource size in bytes at the time of the backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.resource_size_bytes
	ResourceSizeBytes *int64 `json:"resourceSizeBytes,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupApplianceBackupProperties
type BackupApplianceBackupPropertiesObservedState struct {
	// Output only. The numeric generation ID of the backup (monotonically
	//  increasing).
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupProperties.generation_id
	GenerationID *int32 `json:"generationID,omitempty"`

	// Output only. The time when this backup object was finalized (if none,
	//  backup is not finalized).
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupProperties.finalize_time
	FinalizeTime *string `json:"finalizeTime,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupLock
type BackupLockObservedState struct {
	// Output only. Contains metadata about the lock exist for Google Cloud
	//  native backups.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupLock.service_lock_info
	ServiceLockInfo *ServiceLockInfo `json:"serviceLockInfo,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ComputeInstanceBackupProperties
type ComputeInstanceBackupPropertiesObservedState struct {
	// An array of network access configurations for this interface.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.network_interface
	NetworkInterface []NetworkInterfaceObservedState `json:"networkInterface,omitempty"`

	// An array of disks that are associated with the instances that are created
	//  from these properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceBackupProperties.disk
	Disk []AttachedDiskObservedState `json:"disk,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.NetworkInterface
type NetworkInterfaceObservedState struct {
	// Output only. [Output Only] The name of the network interface, which is
	//  generated by the server.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkInterface.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ServiceLockInfo
type ServiceLockInfoObservedState struct {
	// Output only. The name of the operation that created this lock.
	//  The lock will automatically be released when the operation completes.
	// +kcc:proto:field=google.cloud.backupdr.v1.ServiceLockInfo.operation
	Operation *string `json:"operation,omitempty"`
}
