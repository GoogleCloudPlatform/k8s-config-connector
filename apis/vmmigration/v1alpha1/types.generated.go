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


// +kcc:proto=google.cloud.vmmigration.v1.AdaptingOSStep
type AdaptingOSStep struct {
}

// +kcc:proto=google.cloud.vmmigration.v1.AppliedLicense
type AppliedLicense struct {
	// The license type that was used in OS adaptation.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AppliedLicense.type
	Type *string `json:"type,omitempty"`

	// The OS license returned from the adaptation module's report.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AppliedLicense.os_license
	OsLicense *string `json:"osLicense,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.AwsSourceVmDetails
type AwsSourceVmDetails struct {
	// The firmware type of the source VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceVmDetails.firmware
	Firmware *string `json:"firmware,omitempty"`

	// The total size of the disks being migrated in bytes.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceVmDetails.committed_storage_bytes
	CommittedStorageBytes *int64 `json:"committedStorageBytes,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.CloneJob
type CloneJob struct {
}

// +kcc:proto=google.cloud.vmmigration.v1.CloneStep
type CloneStep struct {
	// Adapting OS step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneStep.adapting_os
	AdaptingOs *AdaptingOSStep `json:"adaptingOs,omitempty"`

	// Preparing VM disks step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneStep.preparing_vm_disks
	PreparingVmDisks *PreparingVMDisksStep `json:"preparingVmDisks,omitempty"`

	// Instantiating migrated VM step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneStep.instantiating_migrated_vm
	InstantiatingMigratedVm *InstantiatingMigratedVMStep `json:"instantiatingMigratedVm,omitempty"`

	// The time the step has started.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneStep.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The time the step has ended.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneStep.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults
type ComputeEngineTargetDefaults struct {
	// The name of the VM to create.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.vm_name
	VmName *string `json:"vmName,omitempty"`

	// The full path of the resource of type TargetProject which represents the
	//  Compute Engine project in which to create this VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.target_project
	TargetProject *string `json:"targetProject,omitempty"`

	// The zone in which to create the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.zone
	Zone *string `json:"zone,omitempty"`

	// The machine type series to create the VM with.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.machine_type_series
	MachineTypeSeries *string `json:"machineTypeSeries,omitempty"`

	// The machine type to create the VM with.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// A map of network tags to associate with the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`

	// List of NICs connected to this VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.network_interfaces
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces,omitempty"`

	// The service account to associate the VM with.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The disk type to use in the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// A map of labels to associate with the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The license type to use in OS adaptation.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.license_type
	LicenseType *string `json:"licenseType,omitempty"`

	// Compute instance scheduling information (if empty default is used).
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.compute_scheduling
	ComputeScheduling *ComputeScheduling `json:"computeScheduling,omitempty"`

	// Defines whether the instance has Secure Boot enabled.
	//  This can be set to true only if the vm boot option is EFI.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.secure_boot
	SecureBoot *bool `json:"secureBoot,omitempty"`

	// The metadata key/value pairs to assign to the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Additional licenses to assign to the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.additional_licenses
	AdditionalLicenses []string `json:"additionalLicenses,omitempty"`

	// The hostname to assign to the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.hostname
	Hostname *string `json:"hostname,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.ComputeEngineTargetDetails
type ComputeEngineTargetDetails struct {
	// The name of the VM to create.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.vm_name
	VmName *string `json:"vmName,omitempty"`

	// The Google Cloud target project ID or project name.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.project
	Project *string `json:"project,omitempty"`

	// The zone in which to create the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.zone
	Zone *string `json:"zone,omitempty"`

	// The machine type series to create the VM with.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.machine_type_series
	MachineTypeSeries *string `json:"machineTypeSeries,omitempty"`

	// The machine type to create the VM with.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// A map of network tags to associate with the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`

	// List of NICs connected to this VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.network_interfaces
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces,omitempty"`

	// The service account to associate the VM with.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The disk type to use in the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// A map of labels to associate with the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The license type to use in OS adaptation.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.license_type
	LicenseType *string `json:"licenseType,omitempty"`

	// The OS license returned from the adaptation module report.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.applied_license
	AppliedLicense *AppliedLicense `json:"appliedLicense,omitempty"`

	// Compute instance scheduling information (if empty default is used).
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.compute_scheduling
	ComputeScheduling *ComputeScheduling `json:"computeScheduling,omitempty"`

	// Defines whether the instance has Secure Boot enabled.
	//  This can be set to true only if the vm boot option is EFI.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.secure_boot
	SecureBoot *bool `json:"secureBoot,omitempty"`

	// The VM Boot Option, as set in the source vm.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.boot_option
	BootOption *string `json:"bootOption,omitempty"`

	// The metadata key/value pairs to assign to the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Additional licenses to assign to the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.additional_licenses
	AdditionalLicenses []string `json:"additionalLicenses,omitempty"`

	// The hostname to assign to the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDetails.hostname
	Hostname *string `json:"hostname,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.ComputeScheduling
type ComputeScheduling struct {
	// How the instance should behave when the host machine undergoes
	//  maintenance that may temporarily impact instance performance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeScheduling.on_host_maintenance
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty"`

	// Whether the Instance should be automatically restarted whenever it is
	//  terminated by Compute Engine (not terminated by user).
	//  This configuration is identical to `automaticRestart` field in Compute
	//  Engine create instance under scheduling.
	//  It was changed to an enum (instead of a boolean) to match the default
	//  value in Compute Engine which is automatic restart.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeScheduling.restart_type
	RestartType *string `json:"restartType,omitempty"`

	// A set of node affinity and anti-affinity configurations for sole tenant
	//  nodes.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeScheduling.node_affinities
	NodeAffinities []SchedulingNodeAffinity `json:"nodeAffinities,omitempty"`

	// The minimum number of virtual CPUs this instance will consume when
	//  running on a sole-tenant node. Ignored if no node_affinites are
	//  configured.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeScheduling.min_node_cpus
	MinNodeCpus *int32 `json:"minNodeCpus,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.CutoverJob
type CutoverJob struct {
}

// +kcc:proto=google.cloud.vmmigration.v1.CutoverStep
type CutoverStep struct {
	// A replication cycle prior cutover step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverStep.previous_replication_cycle
	PreviousReplicationCycle *ReplicationCycle `json:"previousReplicationCycle,omitempty"`

	// Shutting down VM step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverStep.shutting_down_source_vm
	ShuttingDownSourceVm *ShuttingDownSourceVMStep `json:"shuttingDownSourceVm,omitempty"`

	// Final sync step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverStep.final_sync
	FinalSync *ReplicationCycle `json:"finalSync,omitempty"`

	// Preparing VM disks step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverStep.preparing_vm_disks
	PreparingVmDisks *PreparingVMDisksStep `json:"preparingVmDisks,omitempty"`

	// Instantiating migrated VM step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverStep.instantiating_migrated_vm
	InstantiatingMigratedVm *InstantiatingMigratedVMStep `json:"instantiatingMigratedVm,omitempty"`

	// The time the step has started.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverStep.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The time the step has ended.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverStep.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.CycleStep
type CycleStep struct {
	// Initializing replication step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CycleStep.initializing_replication
	InitializingReplication *InitializingReplicationStep `json:"initializingReplication,omitempty"`

	// Replicating step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CycleStep.replicating
	Replicating *ReplicatingStep `json:"replicating,omitempty"`

	// Post processing step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CycleStep.post_processing
	PostProcessing *PostProcessingStep `json:"postProcessing,omitempty"`

	// The time the cycle step has started.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CycleStep.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The time the cycle step has ended.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CycleStep.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.InitializingReplicationStep
type InitializingReplicationStep struct {
}

// +kcc:proto=google.cloud.vmmigration.v1.InstantiatingMigratedVMStep
type InstantiatingMigratedVMStep struct {
}

// +kcc:proto=google.cloud.vmmigration.v1.MigratingVm
type MigratingVm struct {
	// Details of the target VM in Compute Engine.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.compute_engine_target_defaults
	ComputeEngineTargetDefaults *ComputeEngineTargetDefaults `json:"computeEngineTargetDefaults,omitempty"`

	// The unique ID of the VM in the source.
	//  The VM's name in vSphere can be changed, so this is not the VM's name but
	//  rather its moRef id. This id is of the form vm-<num>.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.source_vm_id
	SourceVmID *string `json:"sourceVmID,omitempty"`

	// The display name attached to the MigratingVm by the user.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description attached to the migrating VM by the user.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.description
	Description *string `json:"description,omitempty"`

	// The replication schedule policy.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.policy
	Policy *SchedulePolicy `json:"policy,omitempty"`

	// The labels of the migrating VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.NetworkInterface
type NetworkInterface struct {
	// The network to connect the NIC to.
	// +kcc:proto:field=google.cloud.vmmigration.v1.NetworkInterface.network
	Network *string `json:"network,omitempty"`

	// The subnetwork to connect the NIC to.
	// +kcc:proto:field=google.cloud.vmmigration.v1.NetworkInterface.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// The internal IP to define in the NIC.
	//  The formats accepted are: `ephemeral` \ ipv4 address \ a named address
	//  resource full path.
	// +kcc:proto:field=google.cloud.vmmigration.v1.NetworkInterface.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// The external IP to define in the NIC.
	// +kcc:proto:field=google.cloud.vmmigration.v1.NetworkInterface.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.PostProcessingStep
type PostProcessingStep struct {
}

// +kcc:proto=google.cloud.vmmigration.v1.PreparingVMDisksStep
type PreparingVMDisksStep struct {
}

// +kcc:proto=google.cloud.vmmigration.v1.ReplicatingStep
type ReplicatingStep struct {
	// Total bytes to be handled in the step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicatingStep.total_bytes
	TotalBytes *int64 `json:"totalBytes,omitempty"`

	// Replicated bytes in the step.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicatingStep.replicated_bytes
	ReplicatedBytes *int64 `json:"replicatedBytes,omitempty"`

	// The source disks replication rate for the last 2 minutes in bytes per
	//  second.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicatingStep.last_two_minutes_average_bytes_per_second
	LastTwoMinutesAverageBytesPerSecond *int64 `json:"lastTwoMinutesAverageBytesPerSecond,omitempty"`

	// The source disks replication rate for the last 30 minutes in bytes per
	//  second.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicatingStep.last_thirty_minutes_average_bytes_per_second
	LastThirtyMinutesAverageBytesPerSecond *int64 `json:"lastThirtyMinutesAverageBytesPerSecond,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.ReplicationCycle
type ReplicationCycle struct {
	// The identifier of the ReplicationCycle.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.name
	Name *string `json:"name,omitempty"`

	// The cycle's ordinal number.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.cycle_number
	CycleNumber *int32 `json:"cycleNumber,omitempty"`

	// The time the replication cycle has started.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The time the replication cycle has ended.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.end_time
	EndTime *string `json:"endTime,omitempty"`

	// The accumulated duration the replication cycle was paused.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.total_pause_duration
	TotalPauseDuration *string `json:"totalPauseDuration,omitempty"`

	// The current progress in percentage of this cycle.
	//  Was replaced by 'steps' field, which breaks down the cycle progression more
	//  accurately.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.progress_percent
	ProgressPercent *int32 `json:"progressPercent,omitempty"`

	// The cycle's steps list representing its progress.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.steps
	Steps []CycleStep `json:"steps,omitempty"`

	// State of the ReplicationCycle.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.state
	State *string `json:"state,omitempty"`

	// Provides details on the state of the cycle in case of an error.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationCycle.error
	Error *Status `json:"error,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.ReplicationSync
type ReplicationSync struct {
	// The most updated snapshot created time in the source that finished
	//  replication.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ReplicationSync.last_sync_time
	LastSyncTime *string `json:"lastSyncTime,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.SchedulePolicy
type SchedulePolicy struct {
	// The idle duration between replication stages.
	// +kcc:proto:field=google.cloud.vmmigration.v1.SchedulePolicy.idle_duration
	IdleDuration *string `json:"idleDuration,omitempty"`

	// A flag to indicate whether to skip OS adaptation during the replication
	//  sync. OS adaptation is a process where the VM's operating system undergoes
	//  changes and adaptations to fully function on Compute Engine.
	// +kcc:proto:field=google.cloud.vmmigration.v1.SchedulePolicy.skip_os_adaptation
	SkipOsAdaptation *bool `json:"skipOsAdaptation,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.SchedulingNodeAffinity
type SchedulingNodeAffinity struct {
	// The label key of Node resource to reference.
	// +kcc:proto:field=google.cloud.vmmigration.v1.SchedulingNodeAffinity.key
	Key *string `json:"key,omitempty"`

	// The operator to use for the node resources specified in the `values`
	//  parameter.
	// +kcc:proto:field=google.cloud.vmmigration.v1.SchedulingNodeAffinity.operator
	Operator *string `json:"operator,omitempty"`

	// Corresponds to the label values of Node resource.
	// +kcc:proto:field=google.cloud.vmmigration.v1.SchedulingNodeAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.ShuttingDownSourceVMStep
type ShuttingDownSourceVMStep struct {
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.CloneJob
type CloneJobObservedState struct {
	// Output only. Details of the target VM in Compute Engine.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.compute_engine_target_details
	ComputeEngineTargetDetails *ComputeEngineTargetDetails `json:"computeEngineTargetDetails,omitempty"`

	// Output only. The time the clone job was created (as an API call, not when
	//  it was actually created in the target).
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the clone job was ended.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The name of the clone.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the clone job.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.state
	State *string `json:"state,omitempty"`

	// Output only. The time the state was last updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. Provides details for the errors that led to the Clone Job's
	//  state.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. The clone steps list representing its progress.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CloneJob.steps
	Steps []CloneStep `json:"steps,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults
type ComputeEngineTargetDefaultsObservedState struct {
	// Output only. The OS license returned from the adaptation module report.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.applied_license
	AppliedLicense *AppliedLicense `json:"appliedLicense,omitempty"`

	// Output only. The VM Boot Option, as set in the source vm.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ComputeEngineTargetDefaults.boot_option
	BootOption *string `json:"bootOption,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.CutoverJob
type CutoverJobObservedState struct {
	// Output only. Details of the target VM in Compute Engine.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.compute_engine_target_details
	ComputeEngineTargetDetails *ComputeEngineTargetDetails `json:"computeEngineTargetDetails,omitempty"`

	// Output only. The time the cutover job was created (as an API call, not when
	//  it was actually created in the target).
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the cutover job had finished.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The name of the cutover job.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the cutover job.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.state
	State *string `json:"state,omitempty"`

	// Output only. The time the state was last updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. The current progress in percentage of the cutover job.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.progress_percent
	ProgressPercent *int32 `json:"progressPercent,omitempty"`

	// Output only. Provides details for the errors that led to the Cutover Job's
	//  state.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. A message providing possible extra details about the current
	//  state.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The cutover steps list representing its progress.
	// +kcc:proto:field=google.cloud.vmmigration.v1.CutoverJob.steps
	Steps []CutoverStep `json:"steps,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.MigratingVm
type MigratingVmObservedState struct {
	// Details of the target VM in Compute Engine.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.compute_engine_target_defaults
	ComputeEngineTargetDefaults *ComputeEngineTargetDefaultsObservedState `json:"computeEngineTargetDefaults,omitempty"`

	// Output only. Details of the VM from an AWS source.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.aws_source_vm_details
	AwsSourceVmDetails *AwsSourceVmDetails `json:"awsSourceVmDetails,omitempty"`

	// Output only. The identifier of the MigratingVm.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.name
	Name *string `json:"name,omitempty"`

	// Output only. The time the migrating VM was created (this refers to this
	//  resource and not to the time it was installed in the source).
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time the migrating VM resource was updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The most updated snapshot created time in the source that
	//  finished replication.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.last_sync
	LastSync *ReplicationSync `json:"lastSync,omitempty"`

	// Output only. State of the MigratingVm.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.state
	State *string `json:"state,omitempty"`

	// Output only. The last time the migrating VM state was updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. The percentage progress of the current running replication
	//  cycle.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.current_sync_info
	CurrentSyncInfo *ReplicationCycle `json:"currentSyncInfo,omitempty"`

	// Output only. The group this migrating vm is included in, if any. The group
	//  is represented by the full path of the appropriate
	//  [Group][google.cloud.vmmigration.v1.Group] resource.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.group
	Group *string `json:"group,omitempty"`

	// Output only. The recent [clone jobs][google.cloud.vmmigration.v1.CloneJob]
	//  performed on the migrating VM. This field holds the vm's last completed
	//  clone job and the vm's running clone job, if one exists.
	//  Note: To have this field populated you need to explicitly request it via
	//  the "view" parameter of the Get/List request.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.recent_clone_jobs
	RecentCloneJobs []CloneJob `json:"recentCloneJobs,omitempty"`

	// Output only. Provides details on the state of the Migrating VM in case of
	//  an error in replication.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.error
	Error *Status `json:"error,omitempty"`

	// Output only. The recent cutover jobs performed on the migrating VM.
	//  This field holds the vm's last completed cutover job and the vm's
	//  running cutover job, if one exists.
	//  Note: To have this field populated you need to explicitly request it via
	//  the "view" parameter of the Get/List request.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.recent_cutover_jobs
	RecentCutoverJobs []CutoverJob `json:"recentCutoverJobs,omitempty"`
}
