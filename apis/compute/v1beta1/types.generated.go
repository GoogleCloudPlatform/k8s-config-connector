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
// resource: ComputeSubnetwork:Subnetwork
// resource: ComputeTargetTcpProxy:TargetTcpProxy
// resource: ComputeInstance:Instance

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
