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


// +kcc:proto=google.cloud.migrationcenter.v1.AssetFrame
type AssetFrame struct {
	// Asset information specific for virtual machines.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AssetFrame.machine_details
	MachineDetails *MachineDetails `json:"machineDetails,omitempty"`

	// The time the data was reported.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AssetFrame.report_time
	ReportTime *string `json:"reportTime,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AssetFrame.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Generic asset attributes.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AssetFrame.attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// Asset performance data samples.
	//  Samples that are from more than 40 days ago or after tomorrow are ignored.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AssetFrame.performance_samples
	PerformanceSamples []PerformanceSample `json:"performanceSamples,omitempty"`

	// Optional. Trace token is optionally provided to assist with debugging and
	//  traceability.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AssetFrame.trace_token
	TraceToken *string `json:"traceToken,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.AwsEc2PlatformDetails
type AwsEc2PlatformDetails struct {
	// AWS platform's machine type label.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AwsEc2PlatformDetails.machine_type_label
	MachineTypeLabel *string `json:"machineTypeLabel,omitempty"`

	// The location of the machine in the AWS format.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AwsEc2PlatformDetails.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.AzureVmPlatformDetails
type AzureVmPlatformDetails struct {
	// Azure platform's machine type label.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AzureVmPlatformDetails.machine_type_label
	MachineTypeLabel *string `json:"machineTypeLabel,omitempty"`

	// The location of the machine in the Azure format.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AzureVmPlatformDetails.location
	Location *string `json:"location,omitempty"`

	// Azure platform's provisioning state.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.AzureVmPlatformDetails.provisioning_state
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.BiosDetails
type BiosDetails struct {
	// BIOS name.
	//  This fields is deprecated. Please use the `id` field instead.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.BiosDetails.bios_name
	BiosName *string `json:"biosName,omitempty"`

	// BIOS ID.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.BiosDetails.id
	ID *string `json:"id,omitempty"`

	// BIOS manufacturer.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.BiosDetails.manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`

	// BIOS version.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.BiosDetails.version
	Version *string `json:"version,omitempty"`

	// BIOS release date.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.BiosDetails.release_date
	ReleaseDate *Date `json:"releaseDate,omitempty"`

	// SMBIOS UUID.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.BiosDetails.smbios_uuid
	SmbiosUuid *string `json:"smbiosUuid,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.CpuUsageSample
type CpuUsageSample struct {
	// Percentage of total CPU capacity utilized. Must be in the interval [0,
	//  100]. On most systems can be calculated using 100 - idle percentage.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.CpuUsageSample.utilized_percentage
	UtilizedPercentage *float32 `json:"utilizedPercentage,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.DiskEntry
type DiskEntry struct {
	// Disk capacity.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.capacity_bytes
	CapacityBytes *int64 `json:"capacityBytes,omitempty"`

	// Disk free space.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.free_bytes
	FreeBytes *int64 `json:"freeBytes,omitempty"`

	// Disk label.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.disk_label
	DiskLabel *string `json:"diskLabel,omitempty"`

	// Disk label type (e.g. BIOS/GPT)
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.disk_label_type
	DiskLabelType *string `json:"diskLabelType,omitempty"`

	// Disks interface type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.interface_type
	InterfaceType *string `json:"interfaceType,omitempty"`

	// Partition layout.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.partitions
	Partitions *DiskPartitionList `json:"partitions,omitempty"`

	// Disk hardware address (e.g. 0:1 for SCSI).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.hw_address
	HwAddress *string `json:"hwAddress,omitempty"`

	// VMware disk details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntry.vmware
	Vmware *VmwareDiskConfig `json:"vmware,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.DiskEntryList
type DiskEntryList struct {
	// Disk entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskEntryList.entries
	Entries []DiskEntry `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.DiskPartition
type DiskPartition struct {
	// Partition type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartition.type
	Type *string `json:"type,omitempty"`

	// Partition file system.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartition.file_system
	FileSystem *string `json:"fileSystem,omitempty"`

	// Mount pount (Linux/Windows) or drive letter (Windows).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartition.mount_point
	MountPoint *string `json:"mountPoint,omitempty"`

	// Partition capacity.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartition.capacity_bytes
	CapacityBytes *int64 `json:"capacityBytes,omitempty"`

	// Partition free space.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartition.free_bytes
	FreeBytes *int64 `json:"freeBytes,omitempty"`

	// Partition UUID.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartition.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Sub-partitions.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartition.sub_partitions
	SubPartitions *DiskPartitionList `json:"subPartitions,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.DiskPartitionList
type DiskPartitionList struct {
	// Partition entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskPartitionList.entries
	Entries []DiskPartition `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.DiskUsageSample
type DiskUsageSample struct {
	// Average IOPS sampled over a short window. Must be non-negative.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.DiskUsageSample.average_iops
	AverageIops *float32 `json:"averageIops,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ErrorFrame
type ErrorFrame struct {
}

// +kcc:proto=google.cloud.migrationcenter.v1.FrameViolationEntry
type FrameViolationEntry struct {
	// The field of the original frame where the violation occurred.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FrameViolationEntry.field
	Field *string `json:"field,omitempty"`

	// A message describing the violation.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FrameViolationEntry.violation
	Violation *string `json:"violation,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.FstabEntry
type FstabEntry struct {
	// The block special device or remote filesystem to be mounted.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FstabEntry.spec
	Spec *string `json:"spec,omitempty"`

	// The mount point for the filesystem.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FstabEntry.file
	File *string `json:"file,omitempty"`

	// The type of the filesystem.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FstabEntry.vfstype
	Vfstype *string `json:"vfstype,omitempty"`

	// Mount options associated with the filesystem.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FstabEntry.mntops
	Mntops *string `json:"mntops,omitempty"`

	// Used by dump to determine which filesystems need to be dumped.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FstabEntry.freq
	Freq *int32 `json:"freq,omitempty"`

	// Used by the fsck(8) program to determine the order in which filesystem
	//  checks are done at reboot time.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FstabEntry.passno
	Passno *int32 `json:"passno,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.FstabEntryList
type FstabEntryList struct {
	// Fstab entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FstabEntryList.entries
	Entries []FstabEntry `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.GenericPlatformDetails
type GenericPlatformDetails struct {
	// Free text representation of the machine location.
	//  The format of this field should not be relied on. Different VMs in the same
	//  location may have different string values for this field.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GenericPlatformDetails.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.GuestConfigDetails
type GuestConfigDetails struct {
	// OS issue (typically /etc/issue in Linux).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestConfigDetails.issue
	Issue *string `json:"issue,omitempty"`

	// Mount list (Linux fstab).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestConfigDetails.fstab
	Fstab *FstabEntryList `json:"fstab,omitempty"`

	// Hosts file (/etc/hosts).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestConfigDetails.hosts
	Hosts *HostsEntryList `json:"hosts,omitempty"`

	// NFS exports.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestConfigDetails.nfs_exports
	NfsExports *NfsExportList `json:"nfsExports,omitempty"`

	// Security-Enhanced Linux (SELinux) mode.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestConfigDetails.selinux_mode
	SelinuxMode *string `json:"selinuxMode,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.GuestInstalledApplication
type GuestInstalledApplication struct {
	// Installed application name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestInstalledApplication.application_name
	ApplicationName *string `json:"applicationName,omitempty"`

	// Installed application vendor.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestInstalledApplication.vendor
	Vendor *string `json:"vendor,omitempty"`

	// The time when the application was installed.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestInstalledApplication.install_time
	InstallTime *string `json:"installTime,omitempty"`

	// Source path.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestInstalledApplication.path
	Path *string `json:"path,omitempty"`

	// Installed application version.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestInstalledApplication.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.GuestInstalledApplicationList
type GuestInstalledApplicationList struct {
	// Application entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestInstalledApplicationList.entries
	Entries []GuestInstalledApplication `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.GuestOsDetails
type GuestOsDetails struct {
	// The name of the operating system.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestOsDetails.os_name
	OsName *string `json:"osName,omitempty"`

	// What family the OS belong to, if known.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestOsDetails.family
	Family *string `json:"family,omitempty"`

	// The version of the operating system.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestOsDetails.version
	Version *string `json:"version,omitempty"`

	// OS and app configuration.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestOsDetails.config
	Config *GuestConfigDetails `json:"config,omitempty"`

	// Runtime information.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestOsDetails.runtime
	Runtime *GuestRuntimeDetails `json:"runtime,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.GuestRuntimeDetails
type GuestRuntimeDetails struct {
	// Running background services.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.services
	Services *RunningServiceList `json:"services,omitempty"`

	// Running processes.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.processes
	Processes *RunningProcessList `json:"processes,omitempty"`

	// Runtime network information (connections, ports).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.network
	Network *RuntimeNetworkInfo `json:"network,omitempty"`

	// Last time the OS was booted.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.last_boot_time
	LastBootTime *string `json:"lastBootTime,omitempty"`

	// Domain, e.g. c.stratozone-development.internal.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.domain
	Domain *string `json:"domain,omitempty"`

	// Machine name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.machine_name
	MachineName *string `json:"machineName,omitempty"`

	// Installed applications information.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.installed_apps
	InstalledApps *GuestInstalledApplicationList `json:"installedApps,omitempty"`

	// Open files information.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.GuestRuntimeDetails.open_file_list
	OpenFileList *OpenFileList `json:"openFileList,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.HostsEntry
type HostsEntry struct {
	// IP (raw, IPv4/6 agnostic).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.HostsEntry.ip
	IP *string `json:"ip,omitempty"`

	// List of host names / aliases.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.HostsEntry.host_names
	HostNames []string `json:"hostNames,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.HostsEntryList
type HostsEntryList struct {
	// Hosts entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.HostsEntryList.entries
	Entries []HostsEntry `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.MachineArchitectureDetails
type MachineArchitectureDetails struct {
	// CPU architecture, e.g., "x64-based PC", "x86_64", "i686" etc.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.cpu_architecture
	CpuArchitecture *string `json:"cpuArchitecture,omitempty"`

	// CPU name, e.g., "Intel Xeon E5-2690", "AMD EPYC 7571" etc.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.cpu_name
	CpuName *string `json:"cpuName,omitempty"`

	// Hardware vendor.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.vendor
	Vendor *string `json:"vendor,omitempty"`

	// Number of CPU threads allocated to the machine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.cpu_thread_count
	CpuThreadCount *int32 `json:"cpuThreadCount,omitempty"`

	// Number of processor sockets allocated to the machine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.cpu_socket_count
	CpuSocketCount *int32 `json:"cpuSocketCount,omitempty"`

	// BIOS Details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.bios
	Bios *BiosDetails `json:"bios,omitempty"`

	// Firmware type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.firmware_type
	FirmwareType *string `json:"firmwareType,omitempty"`

	// CPU hyper-threading support.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineArchitectureDetails.hyperthreading
	Hyperthreading *string `json:"hyperthreading,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.MachineDetails
type MachineDetails struct {
	// Machine unique identifier.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Machine name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.machine_name
	MachineName *string `json:"machineName,omitempty"`

	// Machine creation time.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Number of CPU cores in the machine. Must be non-negative.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.core_count
	CoreCount *int32 `json:"coreCount,omitempty"`

	// The amount of memory in the machine. Must be non-negative.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.memory_mb
	MemoryMb *int32 `json:"memoryMb,omitempty"`

	// Power state of the machine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.power_state
	PowerState *string `json:"powerState,omitempty"`

	// Architecture details (vendor, CPU architecture).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.architecture
	Architecture *MachineArchitectureDetails `json:"architecture,omitempty"`

	// Guest OS information.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.guest_os
	GuestOs *GuestOsDetails `json:"guestOs,omitempty"`

	// Network details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.network
	Network *MachineNetworkDetails `json:"network,omitempty"`

	// Disk details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.disks
	Disks *MachineDiskDetails `json:"disks,omitempty"`

	// Platform specific information.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDetails.platform
	Platform *PlatformDetails `json:"platform,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.MachineDiskDetails
type MachineDiskDetails struct {
	// Disk total Capacity.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDiskDetails.total_capacity_bytes
	TotalCapacityBytes *int64 `json:"totalCapacityBytes,omitempty"`

	// Total disk free space.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDiskDetails.total_free_bytes
	TotalFreeBytes *int64 `json:"totalFreeBytes,omitempty"`

	// List of disks.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineDiskDetails.disks
	Disks *DiskEntryList `json:"disks,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.MachineNetworkDetails
type MachineNetworkDetails struct {
	// The primary IP address of the machine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineNetworkDetails.primary_ip_address
	PrimaryIPAddress *string `json:"primaryIPAddress,omitempty"`

	// The public IP address of the machine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineNetworkDetails.public_ip_address
	PublicIPAddress *string `json:"publicIPAddress,omitempty"`

	// MAC address of the machine.
	//  This property is used to uniqly identify the machine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineNetworkDetails.primary_mac_address
	PrimaryMacAddress *string `json:"primaryMacAddress,omitempty"`

	// List of network adapters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineNetworkDetails.adapters
	Adapters *NetworkAdapterList `json:"adapters,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.MemoryUsageSample
type MemoryUsageSample struct {
	// Percentage of system memory utilized. Must be in the interval [0, 100].
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MemoryUsageSample.utilized_percentage
	UtilizedPercentage *float32 `json:"utilizedPercentage,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NetworkAdapterDetails
type NetworkAdapterDetails struct {
	// Network adapter type (e.g. VMXNET3).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAdapterDetails.adapter_type
	AdapterType *string `json:"adapterType,omitempty"`

	// MAC address.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAdapterDetails.mac_address
	MacAddress *string `json:"macAddress,omitempty"`

	// NetworkAddressList
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAdapterDetails.addresses
	Addresses *NetworkAddressList `json:"addresses,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NetworkAdapterList
type NetworkAdapterList struct {
	// Network adapter entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAdapterList.entries
	Entries []NetworkAdapterDetails `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NetworkAddress
type NetworkAddress struct {
	// Assigned or configured IP Address.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAddress.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Subnet mask.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAddress.subnet_mask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// Broadcast address.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAddress.bcast
	Bcast *string `json:"bcast,omitempty"`

	// Fully qualified domain name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAddress.fqdn
	Fqdn *string `json:"fqdn,omitempty"`

	// Whether DHCP is used to assign addresses.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAddress.assignment
	Assignment *string `json:"assignment,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NetworkAddressList
type NetworkAddressList struct {
	// Network address entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkAddressList.entries
	Entries []NetworkAddress `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NetworkConnection
type NetworkConnection struct {
	// Connection protocol (e.g. TCP/UDP).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Local IP address.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.local_ip_address
	LocalIPAddress *string `json:"localIPAddress,omitempty"`

	// Local port.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.local_port
	LocalPort *int32 `json:"localPort,omitempty"`

	// Remote IP address.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.remote_ip_address
	RemoteIPAddress *string `json:"remoteIPAddress,omitempty"`

	// Remote port.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.remote_port
	RemotePort *int32 `json:"remotePort,omitempty"`

	// Network connection state.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.state
	State *string `json:"state,omitempty"`

	// Process ID.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.pid
	Pid *int64 `json:"pid,omitempty"`

	// Process or service name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnection.process_name
	ProcessName *string `json:"processName,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NetworkConnectionList
type NetworkConnectionList struct {
	// Network connection entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkConnectionList.entries
	Entries []NetworkConnection `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NetworkUsageSample
type NetworkUsageSample struct {
	// Average network ingress in B/s sampled over a short window.
	//  Must be non-negative.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkUsageSample.average_ingress_bps
	AverageIngressBps *float32 `json:"averageIngressBps,omitempty"`

	// Average network egress in B/s sampled over a short window.
	//  Must be non-negative.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NetworkUsageSample.average_egress_bps
	AverageEgressBps *float32 `json:"averageEgressBps,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NfsExport
type NfsExport struct {
	// The directory being exported.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NfsExport.export_directory
	ExportDirectory *string `json:"exportDirectory,omitempty"`

	// The hosts or networks to which the export is being shared.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NfsExport.hosts
	Hosts []string `json:"hosts,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.NfsExportList
type NfsExportList struct {
	// NFS export entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.NfsExportList.entries
	Entries []NfsExport `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.OpenFileDetails
type OpenFileDetails struct {
	// Opened file command.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.OpenFileDetails.command
	Command *string `json:"command,omitempty"`

	// Opened file user.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.OpenFileDetails.user
	User *string `json:"user,omitempty"`

	// Opened file file type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.OpenFileDetails.file_type
	FileType *string `json:"fileType,omitempty"`

	// Opened file file path.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.OpenFileDetails.file_path
	FilePath *string `json:"filePath,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.OpenFileList
type OpenFileList struct {
	// Open file details entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.OpenFileList.entries
	Entries []OpenFileDetails `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.PerformanceSample
type PerformanceSample struct {
	// Time the sample was collected.
	//  If omitted, the frame report time will be used.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PerformanceSample.sample_time
	SampleTime *string `json:"sampleTime,omitempty"`

	// Memory usage sample.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PerformanceSample.memory
	Memory *MemoryUsageSample `json:"memory,omitempty"`

	// CPU usage sample.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PerformanceSample.cpu
	Cpu *CpuUsageSample `json:"cpu,omitempty"`

	// Network usage sample.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PerformanceSample.network
	Network *NetworkUsageSample `json:"network,omitempty"`

	// Disk usage sample.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PerformanceSample.disk
	Disk *DiskUsageSample `json:"disk,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.PhysicalPlatformDetails
type PhysicalPlatformDetails struct {
	// Free text representation of the machine location.
	//  The format of this field should not be relied on. Different machines in the
	//  same location may have different string values for this field.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PhysicalPlatformDetails.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.PlatformDetails
type PlatformDetails struct {
	// VMware specific details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PlatformDetails.vmware_details
	VmwareDetails *VmwarePlatformDetails `json:"vmwareDetails,omitempty"`

	// AWS EC2 specific details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PlatformDetails.aws_ec2_details
	AwsEc2Details *AwsEc2PlatformDetails `json:"awsEc2Details,omitempty"`

	// Azure VM specific details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PlatformDetails.azure_vm_details
	AzureVmDetails *AzureVmPlatformDetails `json:"azureVmDetails,omitempty"`

	// Generic platform details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PlatformDetails.generic_details
	GenericDetails *GenericPlatformDetails `json:"genericDetails,omitempty"`

	// Physical machines platform details.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PlatformDetails.physical_details
	PhysicalDetails *PhysicalPlatformDetails `json:"physicalDetails,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.RunningProcess
type RunningProcess struct {
	// Process ID.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningProcess.pid
	Pid *int64 `json:"pid,omitempty"`

	// Process binary path.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningProcess.exe_path
	ExePath *string `json:"exePath,omitempty"`

	// Process full command line.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningProcess.cmdline
	Cmdline *string `json:"cmdline,omitempty"`

	// User running the process.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningProcess.user
	User *string `json:"user,omitempty"`

	// Process extended attributes.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningProcess.attributes
	Attributes map[string]string `json:"attributes,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.RunningProcessList
type RunningProcessList struct {
	// Running process entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningProcessList.entries
	Entries []RunningProcess `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.RunningService
type RunningService struct {
	// Service name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningService.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// Service state (OS-agnostic).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningService.state
	State *string `json:"state,omitempty"`

	// Service start mode (OS-agnostic).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningService.start_mode
	StartMode *string `json:"startMode,omitempty"`

	// Service binary path.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningService.exe_path
	ExePath *string `json:"exePath,omitempty"`

	// Service command line.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningService.cmdline
	Cmdline *string `json:"cmdline,omitempty"`

	// Service pid.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningService.pid
	Pid *int64 `json:"pid,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.RunningServiceList
type RunningServiceList struct {
	// Running service entries.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RunningServiceList.entries
	Entries []RunningService `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.RuntimeNetworkInfo
type RuntimeNetworkInfo struct {
	// Time of the last network scan.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RuntimeNetworkInfo.scan_time
	ScanTime *string `json:"scanTime,omitempty"`

	// Network connections.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RuntimeNetworkInfo.connections
	Connections *NetworkConnectionList `json:"connections,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.VmwareDiskConfig
type VmwareDiskConfig struct {
	// VMDK backing type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareDiskConfig.backing_type
	BackingType *string `json:"backingType,omitempty"`

	// Is VMDK shared with other VMs.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareDiskConfig.shared
	Shared *bool `json:"shared,omitempty"`

	// VMDK disk mode.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareDiskConfig.vmdk_mode
	VmdkMode *string `json:"vmdkMode,omitempty"`

	// RDM compatibility mode.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareDiskConfig.rdm_compatibility
	RdmCompatibility *string `json:"rdmCompatibility,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.VmwarePlatformDetails
type VmwarePlatformDetails struct {
	// vCenter version.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwarePlatformDetails.vcenter_version
	VcenterVersion *string `json:"vcenterVersion,omitempty"`

	// ESX version.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwarePlatformDetails.esx_version
	EsxVersion *string `json:"esxVersion,omitempty"`

	// VMware os enum -
	//  https://vdc-repo.vmware.com/vmwb-repository/dcr-public/da47f910-60ac-438b-8b9b-6122f4d14524/16b7274a-bf8b-4b4c-a05e-746f2aa93c8c/doc/vim.vm.GuestOsDescriptor.GuestOsIdentifier.html.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwarePlatformDetails.osid
	Osid *string `json:"osid,omitempty"`

	// Folder name in vCenter where asset resides.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwarePlatformDetails.vcenter_folder
	VcenterFolder *string `json:"vcenterFolder,omitempty"`

	// vCenter URI used in collection.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwarePlatformDetails.vcenter_uri
	VcenterURI *string `json:"vcenterURI,omitempty"`

	// vCenter VM ID.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwarePlatformDetails.vcenter_vm_id
	VcenterVmID *string `json:"vcenterVmID,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ErrorFrame
type ErrorFrameObservedState struct {
	// Output only. The identifier of the ErrorFrame.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ErrorFrame.name
	Name *string `json:"name,omitempty"`

	// Output only. All the violations that were detected for the frame.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ErrorFrame.violations
	Violations []FrameViolationEntry `json:"violations,omitempty"`

	// Output only. The frame that was originally reported.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ErrorFrame.original_frame
	OriginalFrame *AssetFrame `json:"originalFrame,omitempty"`

	// Output only. Frame ingestion time.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ErrorFrame.ingestion_time
	IngestionTime *string `json:"ingestionTime,omitempty"`
}
