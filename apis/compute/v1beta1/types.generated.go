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
// proto.service: google.cloud.compute.v1beta
// resource: ComputeFutureReservation:FutureReservation

package v1beta1

// +kcc:proto=google.cloud.compute.v1beta.AcceleratorConfig
type AcceleratorConfig struct {
	// The number of the guest accelerator cards exposed to this instance.
	// +kcc:proto:field=google.cloud.compute.v1beta.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Full or partial URL of the accelerator type resource to attach to this instance. For example: projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100 If you are creating an instance template, specify only the accelerator name. See GPUs on Compute Engine for a full list of accelerator types.
	// +kcc:proto:field=google.cloud.compute.v1beta.AcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.AllocationAggregateReservation
type AllocationAggregateReservation struct {
	// [Output only] List of resources currently in use.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationAggregateReservation.in_use_resources
	InUseResources []AllocationAggregateReservationReservedResourceInfo `json:"inUseResources,omitempty"`

	// List of reserved resources (CPUs, memory, accelerators).
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationAggregateReservation.reserved_resources
	ReservedResources []AllocationAggregateReservationReservedResourceInfo `json:"reservedResources,omitempty"`

	// The VM family that all instances scheduled against this reservation must belong to.
	//  Check the VmFamily enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationAggregateReservation.vm_family
	VMFamily *string `json:"vmFamily,omitempty"`

	// The workload type of the instances that will target this reservation.
	//  Check the WorkloadType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationAggregateReservation.workload_type
	WorkloadType *string `json:"workloadType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.AllocationAggregateReservationReservedResourceInfo
type AllocationAggregateReservationReservedResourceInfo struct {
	// Properties of accelerator resources in this reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationAggregateReservationReservedResourceInfo.accelerator
	Accelerator *AllocationAggregateReservationReservedResourceInfoAccelerator `json:"accelerator,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.AllocationAggregateReservationReservedResourceInfoAccelerator
type AllocationAggregateReservationReservedResourceInfoAccelerator struct {
	// Number of accelerators of specified type.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationAggregateReservationReservedResourceInfoAccelerator.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Full or partial URL to accelerator type. e.g. "projects/{PROJECT}/zones/{ZONE}/acceleratorTypes/ct4l"
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationAggregateReservationReservedResourceInfoAccelerator.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk
type AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk struct {
	// Specifies the size of the disk in base-2 GB.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. For performance characteristics of SCSI over NVMe, see Local SSD performance.
	//  Check the Interface enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk.interface
	Interface *string `json:"interface,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties
type AllocationSpecificSkuAllocationReservedInstanceProperties struct {
	// Specifies accelerator type and count.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties.guest_accelerators
	GuestAccelerators []AcceleratorConfig `json:"guestAccelerators,omitempty"`

	// Specifies amount of local ssd to reserve with each instance. The type of disk is local-ssd.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties.local_ssds
	LocalSsds []AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk `json:"localSsds,omitempty"`

	// An opaque location hint used to place the allocation close to other resources. This field is for use by internal tools that use the public API.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties.location_hint
	LocationHint *string `json:"locationHint,omitempty"`

	// Specifies type of machine (name only) which has fixed number of vCPUs and fixed amount of memory. This also includes specifying custom machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY pattern.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Specifies the number of hours after reservation creation where instances using the reservation won't be scheduled for maintenance.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties.maintenance_freeze_duration_hours
	MaintenanceFreezeDurationHours *int32 `json:"maintenanceFreezeDurationHours,omitempty"`

	// Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
	//  Check the MaintenanceInterval enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties.maintenance_interval
	MaintenanceInterval *string `json:"maintenanceInterval,omitempty"`

	// Minimum cpu platform the reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.AllocationSpecificSKUAllocationReservedInstanceProperties.min_cpu_platform
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.Duration
type Duration struct {
	// Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field. Must be from 0 to 999,999,999 inclusive.
	// +kcc:proto:field=google.cloud.compute.v1beta.Duration.nanos
	Nanos *int32 `json:"nanos,omitempty"`

	// Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive. Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	// +kcc:proto:field=google.cloud.compute.v1beta.Duration.seconds
	Seconds *int64 `json:"seconds,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.FutureReservationCommitmentInfo
type FutureReservationCommitmentInfo struct {
	// name of the commitment where capacity is being delivered to.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationCommitmentInfo.commitment_name
	CommitmentName *string `json:"commitmentName,omitempty"`

	// Indicates if a Commitment needs to be created as part of FR delivery. If this field is not present, then no commitment needs to be created.
	//  Check the CommitmentPlan enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationCommitmentInfo.commitment_plan
	CommitmentPlan *string `json:"commitmentPlan,omitempty"`

	// Only applicable if FR is delivering to the same reservation. If set, all parent commitments will be extended to match the end date of the plan for this commitment.
	//  Check the PreviousCommitmentTerms enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationCommitmentInfo.previous_commitment_terms
	PreviousCommitmentTerms *string `json:"previousCommitmentTerms,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.FutureReservationSpecificSKUProperties
type FutureReservationSpecificSkuProperties struct {
	// Properties of the SKU instances being reserved.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationSpecificSKUProperties.instance_properties
	InstanceProperties *AllocationSpecificSkuAllocationReservedInstanceProperties `json:"instanceProperties,omitempty"`

	// The instance template that will be used to populate the ReservedInstanceProperties of the future reservation
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationSpecificSKUProperties.source_instance_template
	SourceInstanceTemplate *string `json:"sourceInstanceTemplate,omitempty"`

	// Total number of instances for which capacity assurance is requested at a future time period.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationSpecificSKUProperties.total_count
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.FutureReservationStatusExistingMatchingUsageInfo
type FutureReservationStatusExistingMatchingUsageInfo struct {
	// Count to represent min(FR total_count, matching_reserved_capacity+matching_unreserved_instances)
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusExistingMatchingUsageInfo.count
	Count *int64 `json:"count,omitempty"`

	// Timestamp when the matching usage was calculated
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusExistingMatchingUsageInfo.timestamp
	Timestamp *string `json:"timestamp,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodState
type FutureReservationStatusLastKnownGoodState struct {
	// [Output Only] The description of the FutureReservation before an amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodState.description
	Description *string `json:"description,omitempty"`

	// [Output Only] Represents the matching usage for the future reservation before an amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodState.existing_matching_usage_info
	ExistingMatchingUsageInfo *FutureReservationStatusExistingMatchingUsageInfo `json:"existingMatchingUsageInfo,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodState.future_reservation_specs
	FutureReservationSpecs *FutureReservationStatusLastKnownGoodStateFutureReservationSpecs `json:"futureReservationSpecs,omitempty"`

	// [Output Only] The lock time of the FutureReservation before an amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodState.lock_time
	LockTime *string `json:"lockTime,omitempty"`

	// [Output Only] The name prefix of the Future Reservation before an amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodState.name_prefix
	NamePrefix *string `json:"namePrefix,omitempty"`

	// [Output Only] The status of the last known good state for the Future Reservation.
	//  Check the ProcurementStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodState.procurement_status
	ProcurementStatus *string `json:"procurementStatus,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs
type FutureReservationStatusLastKnownGoodStateFutureReservationSpecs struct {
	// [Output Only] The previous share settings of the Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs.share_settings
	ShareSettings *ShareSettings `json:"shareSettings,omitempty"`

	// [Output Only] The previous instance related properties of the Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs.specific_sku_properties
	SpecificSkuProperties *FutureReservationSpecificSkuProperties `json:"specificSkuProperties,omitempty"`

	// [Output Only] The previous time window of the Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs.time_window
	TimeWindow *FutureReservationTimeWindow `json:"timeWindow,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.FutureReservationStatusSpecificSKUProperties
type FutureReservationStatusSpecificSkuProperties struct {
	// ID of the instance template used to populate the Future Reservation properties.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatusSpecificSKUProperties.source_instance_template_id
	SourceInstanceTemplateID *string `json:"sourceInstanceTemplateID,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.FutureReservationTimeWindow
type FutureReservationTimeWindow struct {
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationTimeWindow.duration
	Duration *Duration `json:"duration,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationTimeWindow.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Start time of the Future Reservation. The start_time is an RFC3339 string.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationTimeWindow.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.ShareSettings
type ShareSettings struct {

	// TODO: unsupported map type with key string and value message

	// A List of Project names to specify consumer projects for this shared-reservation. This is only valid when share_type's value is SPECIFIC_PROJECTS.
	// +kcc:proto:field=google.cloud.compute.v1beta.ShareSettings.projects
	Projects []string `json:"projects,omitempty"`

	// Type of sharing for this shared-reservation
	//  Check the ShareType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.ShareSettings.share_type
	ShareType *string `json:"shareType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1beta.ShareSettingsProjectConfig
type ShareSettingsProjectConfig struct {
	// The project ID, should be same as the key of this project config in the parent map.
	// +kcc:proto:field=google.cloud.compute.v1beta.ShareSettingsProjectConfig.project_id
	ProjectID *string `json:"projectID,omitempty"`
}
