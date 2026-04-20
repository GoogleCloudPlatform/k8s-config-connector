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

// +kcc:proto=google.cloud.compute.v1.AcceleratorConfig
type AcceleratorConfig struct {
	// The number of the guest accelerator cards exposed to this instance.
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Full or partial URL of the accelerator type resource to attach to this
	//  instance. For example:projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100
	//  If you are creating an instance template, specify only the
	//  accelerator name.
	//  See GPUs on Compute Engine
	//  for a full list of accelerator types.
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationAggregateReservation
type AllocationAggregateReservation struct {
	// Output only. [Output only] List of resources currently in use.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationAggregateReservation.in_use_resources
	InUseResources []AllocationAggregateReservationReservedResourceInfo `json:"inUseResources,omitempty"`

	// List of reserved resources (CPUs, memory, accelerators).
	// +kcc:proto:field=google.cloud.compute.v1.AllocationAggregateReservation.reserved_resources
	ReservedResources []AllocationAggregateReservationReservedResourceInfo `json:"reservedResources,omitempty"`

	// The VM family that all instances scheduled against this reservation must
	//  belong to.
	//  Check the VmFamily enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationAggregateReservation.vm_family
	VMFamily *string `json:"vmFamily,omitempty"`

	// The workload type of the instances that will target this reservation.
	//  Check the WorkloadType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationAggregateReservation.workload_type
	WorkloadType *string `json:"workloadType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationAggregateReservationReservedResourceInfo
type AllocationAggregateReservationReservedResourceInfo struct {
	// Properties of accelerator resources in this reservation.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationAggregateReservationReservedResourceInfo.accelerator
	Accelerator *AllocationAggregateReservationReservedResourceInfoAccelerator `json:"accelerator,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationAggregateReservationReservedResourceInfoAccelerator
type AllocationAggregateReservationReservedResourceInfoAccelerator struct {
	// Number of accelerators of specified type.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationAggregateReservationReservedResourceInfoAccelerator.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Full or partial URL to accelerator type. e.g.
	//  "projects/{PROJECT}/zones/{ZONE}/acceleratorTypes/ct4l"
	// +kcc:proto:field=google.cloud.compute.v1.AllocationAggregateReservationReservedResourceInfoAccelerator.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk
type AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk struct {
	// Specifies the size of the disk in base-2 GB.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Specifies the disk interface to use for attaching this disk, which is
	//  either SCSI or NVME. The default isSCSI.
	//  For performance characteristics of SCSI over NVMe, seeLocal SSD performance.
	//  Check the Interface enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk.interface
	Interface *string `json:"interface,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties
type AllocationSpecificSkuAllocationReservedInstanceProperties struct {
	// Specifies accelerator type and count.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.guest_accelerators
	GuestAccelerators []AcceleratorConfig `json:"guestAccelerators,omitempty"`

	// Specifies amount of local ssd to reserve with each instance. The type
	//  of disk is local-ssd.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.local_ssds
	LocalSsds []AllocationSpecificSkuAllocationAllocatedInstancePropertiesReservedDisk `json:"localSsds,omitempty"`

	// An opaque location hint used to place the allocation close to other
	//  resources.
	//  This field is for use by internal tools that use the public API.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.location_hint
	LocationHint *string `json:"locationHint,omitempty"`

	// Specifies type of machine (name only) which has fixed number of vCPUs
	//  and fixed amount of memory. This also includes specifying custom
	//  machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY pattern.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Minimum cpu platform the reservation.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.min_cpu_platform
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Duration
type Duration struct {
	// Span of time that's a fraction of a second at nanosecond resolution.
	//  Durations less than one second are represented with a 0
	//  `seconds` field and a positive `nanos` field. Must be from 0
	//  to 999,999,999 inclusive.
	// +kcc:proto:field=google.cloud.compute.v1.Duration.nanos
	Nanos *int32 `json:"nanos,omitempty"`

	// Span of time at a resolution of a second. Must be from 0
	//  to 315,576,000,000 inclusive. Note: these bounds are computed from:
	//  60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	// +kcc:proto:field=google.cloud.compute.v1.Duration.seconds
	Seconds *int64 `json:"seconds,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservation
type FutureReservation struct {
	// Aggregate reservation details for the future reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.aggregate_reservation
	AggregateReservation *AllocationAggregateReservation `json:"aggregateReservation,omitempty"`

	// Future timestamp when the FR auto-created reservations will be deleted by
	//  Compute Engine. Format of this field must be a valid
	//  href="https://www.ietf.org/rfc/rfc3339.txt">RFC3339 value.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.auto_created_reservations_delete_time
	AutoCreatedReservationsDeleteTime *string `json:"autoCreatedReservationsDeleteTime,omitempty"`

	// Specifies the duration of auto-created reservations.
	//  It represents relative time to future reservation start_time when
	//  auto-created reservations will be automatically deleted by Compute
	//  Engine. Duration time unit is represented as a count of seconds
	//  and fractions of seconds at nanosecond resolution.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.auto_created_reservations_duration
	AutoCreatedReservationsDuration *Duration `json:"autoCreatedReservationsDuration,omitempty"`

	// Setting for enabling or disabling automatic deletion for auto-created
	//  reservation. If set to true, auto-created reservations will be
	//  deleted at Future Reservation's end time (default) or at user's defined
	//  timestamp if any of the
	//  [auto_created_reservations_delete_time, auto_created_reservations_duration]
	//  values is specified.
	//  For keeping auto-created reservation indefinitely, this value should be set
	//  to false.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.auto_delete_auto_created_reservations
	AutoDeleteAutoCreatedReservations *bool `json:"autoDeleteAutoCreatedReservations,omitempty"`

	// If not present, then FR will not deliver a new commitment or update an
	//  existing commitment.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.commitment_info
	CommitmentInfo *FutureReservationCommitmentInfo `json:"commitmentInfo,omitempty"`

	// Output only. [Output Only] The creation timestamp for this future reservation inRFC3339
	//  text format.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Type of the deployment requested as part of future reservation.
	//  Check the DeploymentType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.deployment_type
	DeploymentType *string `json:"deploymentType,omitempty"`

	// An optional description of this resource. Provide this property when you
	//  create the future reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.description
	Description *string `json:"description,omitempty"`

	// Indicates if this group of VMs have emergent maintenance enabled.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.enable_emergent_maintenance
	EnableEmergentMaintenance *bool `json:"enableEmergentMaintenance,omitempty"`

	// Output only. [Output Only] A unique identifier for this future reservation. The server
	//  defines this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.id
	ID *uint64 `json:"id,omitempty"`

	// Output only. [Output Only] Type of the resource. Alwayscompute#futureReservation for future reservations.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.kind
	Kind *string `json:"kind,omitempty"`

	// The name of the resource, provided by the client when initially creating
	//  the resource. The resource name must be 1-63 characters long, and comply
	//  withRFC1035.
	//  Specifically, the name must be 1-63 characters long and match the regular
	//  expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first
	//  character must be a lowercase letter, and all following characters must be
	//  a dash, lowercase letter, or digit, except the last character, which cannot
	//  be a dash.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.name
	Name *string `json:"name,omitempty"`

	// Name prefix for the reservations to be created at the time of
	//  delivery. The name prefix must comply with RFC1035.
	//  Maximum allowed length for name prefix is 20. Automatically created
	//  reservations name format will be -date-####.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.name_prefix
	NamePrefix *string `json:"namePrefix,omitempty"`

	// Planning state before being submitted for evaluation
	//  Check the PlanningStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.planning_status
	PlanningStatus *string `json:"planningStatus,omitempty"`

	// The reservation mode which determines reservation-termination behavior and
	//  expected pricing.
	//  Check the ReservationMode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.reservation_mode
	ReservationMode *string `json:"reservationMode,omitempty"`

	// Name of reservations where the capacity is provisioned at the time of
	//  delivery of  future reservations. If the reservation with the given name
	//  does not exist already, it is created automatically at the time of Approval
	//  with INACTIVE state till specified start-time. Either provide the
	//  reservation_name or a name_prefix.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.reservation_name
	ReservationName *string `json:"reservationName,omitempty"`

	// Maintenance information for this reservation
	//  Check the SchedulingType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.scheduling_type
	SchedulingType *string `json:"schedulingType,omitempty"`

	// Output only. [Output Only] Server-defined fully-qualified URL for this resource.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. [Output Only] Server-defined URL for this resource with the resource id.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// List of Projects/Folders to share with.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.share_settings
	ShareSettings *ShareSettings `json:"shareSettings,omitempty"`

	// Indicates whether the auto-created reservation can be consumed by VMs with
	//  affinity for "any" reservation. If the field is set, then only VMs that
	//  target the reservation by name can consume from the delivered reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.specific_reservation_required
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty"`

	// Future Reservation configuration to indicate instance properties and
	//  total count.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.specific_sku_properties
	SpecificSkuProperties *FutureReservationSpecificSkuProperties `json:"specificSkuProperties,omitempty"`

	// Output only. [Output only] Status of the Future Reservation
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.status
	Status *FutureReservationStatus `json:"status,omitempty"`

	// Time window for this Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.time_window
	TimeWindow *FutureReservationTimeWindow `json:"timeWindow,omitempty"`

	// Output only. [Output Only] URL of the Zone where this future reservation resides.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservation.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationCommitmentInfo
type FutureReservationCommitmentInfo struct {
	// name of the commitment where capacity is being delivered to.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationCommitmentInfo.commitment_name
	CommitmentName *string `json:"commitmentName,omitempty"`

	// Indicates if a Commitment needs to be created as part of FR delivery. If
	//  this field is not present, then no
	//  commitment needs to be created.
	//  Check the CommitmentPlan enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationCommitmentInfo.commitment_plan
	CommitmentPlan *string `json:"commitmentPlan,omitempty"`

	// Only applicable if FR is delivering to the same reservation. If set, all
	//  parent commitments will be extended to match the end date of the plan for
	//  this commitment.
	//  Check the PreviousCommitmentTerms enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationCommitmentInfo.previous_commitment_terms
	PreviousCommitmentTerms *string `json:"previousCommitmentTerms,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationSpecificSKUProperties
type FutureReservationSpecificSkuProperties struct {
	// Properties of the SKU instances being reserved.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationSpecificSKUProperties.instance_properties
	InstanceProperties *AllocationSpecificSkuAllocationReservedInstanceProperties `json:"instanceProperties,omitempty"`

	// The instance template that will be used to populate the
	//  ReservedInstanceProperties of the future reservation
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationSpecificSKUProperties.source_instance_template
	SourceInstanceTemplate *string `json:"sourceInstanceTemplate,omitempty"`

	// Total number of instances for which capacity assurance is requested at a
	//  future time period.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationSpecificSKUProperties.total_count
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationStatus
type FutureReservationStatus struct {
	// Output only. [Output Only] The current status of the requested amendment.
	//  Check the AmendmentStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.amendment_status
	AmendmentStatus *string `json:"amendmentStatus,omitempty"`

	// Output only. Fully qualified urls of the automatically created reservations at
	//  start_time.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.auto_created_reservations
	AutoCreatedReservations []string `json:"autoCreatedReservations,omitempty"`

	// Output only. [Output Only] Represents the existing matching usage for the future
	//  reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.existing_matching_usage_info
	ExistingMatchingUsageInfo *FutureReservationStatusExistingMatchingUsageInfo `json:"existingMatchingUsageInfo,omitempty"`

	// Output only. This count indicates the fulfilled capacity so far. This is set during
	//  "PROVISIONING" state. This count also includes capacity delivered as part
	//  of existing matching reservations.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.fulfilled_count
	FulfilledCount *int64 `json:"fulfilledCount,omitempty"`

	// Output only. [Output Only] This field represents the future reservation before an
	//  amendment was requested. If the amendment is declined, the Future
	//  Reservation will be reverted to the last known good state. The last known
	//  good state is not set when updating a future reservation whose
	//  Procurement Status is DRAFTING.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.last_known_good_state
	LastKnownGoodState *FutureReservationStatusLastKnownGoodState `json:"lastKnownGoodState,omitempty"`

	// Output only. Time when Future Reservation would become LOCKED, after which no
	//  modifications to Future Reservation will be allowed. Applicable only
	//  after the Future Reservation is in the APPROVED state. The lock_time is
	//  an RFC3339 string. The procurement_status will transition to PROCURING
	//  state at this time.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.lock_time
	LockTime *string `json:"lockTime,omitempty"`

	// Output only. Current state of this Future Reservation
	//  Check the ProcurementStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.procurement_status
	ProcurementStatus *string `json:"procurementStatus,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatus.specific_sku_properties
	SpecificSkuProperties *FutureReservationStatusSpecificSkuProperties `json:"specificSkuProperties,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationStatusExistingMatchingUsageInfo
type FutureReservationStatusExistingMatchingUsageInfo struct {
	// Output only. Count to represent min(FR total_count,
	//  matching_reserved_capacity+matching_unreserved_instances)
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusExistingMatchingUsageInfo.count
	Count *int64 `json:"count,omitempty"`

	// Output only. Timestamp when the matching usage was calculated
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusExistingMatchingUsageInfo.timestamp
	Timestamp *string `json:"timestamp,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodState
type FutureReservationStatusLastKnownGoodState struct {
	// Output only. [Output Only] The description of the FutureReservation before an
	//  amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodState.description
	Description *string `json:"description,omitempty"`

	// Output only. [Output Only] Represents the matching usage for the future
	//  reservation before an amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodState.existing_matching_usage_info
	ExistingMatchingUsageInfo *FutureReservationStatusExistingMatchingUsageInfo `json:"existingMatchingUsageInfo,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodState.future_reservation_specs
	FutureReservationSpecs *FutureReservationStatusLastKnownGoodStateFutureReservationSpecs `json:"futureReservationSpecs,omitempty"`

	// Output only. [Output Only] The lock time of the FutureReservation before an
	//  amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodState.lock_time
	LockTime *string `json:"lockTime,omitempty"`

	// Output only. [Output Only] The name prefix of the Future Reservation before an
	//  amendment was requested.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodState.name_prefix
	NamePrefix *string `json:"namePrefix,omitempty"`

	// Output only. [Output Only] The status of the last known good state for the Future
	//  Reservation.
	//  Check the ProcurementStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodState.procurement_status
	ProcurementStatus *string `json:"procurementStatus,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs
type FutureReservationStatusLastKnownGoodStateFutureReservationSpecs struct {
	// Output only. [Output Only] The previous share settings of the Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs.share_settings
	ShareSettings *ShareSettings `json:"shareSettings,omitempty"`

	// Output only. [Output Only] The previous instance related properties of the
	//  Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs.specific_sku_properties
	SpecificSkuProperties *FutureReservationSpecificSkuProperties `json:"specificSkuProperties,omitempty"`

	// Output only. [Output Only] The previous time window of the Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusLastKnownGoodStateFutureReservationSpecs.time_window
	TimeWindow *FutureReservationTimeWindow `json:"timeWindow,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationStatusSpecificSKUProperties
type FutureReservationStatusSpecificSkuProperties struct {
	// ID of the instance template used to populate the Future Reservation
	//  properties.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationStatusSpecificSKUProperties.source_instance_template_id
	SourceInstanceTemplateID *string `json:"sourceInstanceTemplateID,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FutureReservationTimeWindow
type FutureReservationTimeWindow struct {
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationTimeWindow.duration
	Duration *Duration `json:"duration,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationTimeWindow.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Start time of the Future Reservation. The start_time is an RFC3339
	//  string.
	// +kcc:proto:field=google.cloud.compute.v1.FutureReservationTimeWindow.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectCircuitInfo
type InterconnectCircuitInfo struct {
	// Customer-side demarc ID for this circuit.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectCircuitInfo.customer_demarc_id
	CustomerDemarcID *string `json:"customerDemarcID,omitempty"`

	// Google-assigned unique ID for this circuit. Assigned at circuit turn-up.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectCircuitInfo.google_circuit_id
	GoogleCircuitID *string `json:"googleCircuitID,omitempty"`

	// Google-side demarc ID for this circuit. Assigned at circuit turn-up and provided by Google to the customer in the LOA.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectCircuitInfo.google_demarc_id
	GoogleDemarcID *string `json:"googleDemarcID,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectMacsec
type InterconnectMacsec struct {
	// If set to true, the Interconnect connection is configured with a should-secure MACsec security policy, that allows the Google router to fallback to cleartext traffic if the MKA session cannot be established. By default, the Interconnect connection is configured with a must-secure security policy that drops all traffic if the MKA session cannot be established with your router.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsec.fail_open
	FailOpen *bool `json:"failOpen,omitempty"`

	// Required. A keychain placeholder describing a set of named key objects along with their start times. A MACsec CKN/CAK is generated for each key in the key chain. Google router automatically picks the key with the most recent startTime when establishing or re-establishing a MACsec secure link.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsec.pre_shared_keys
	PreSharedKeys []InterconnectMacsecPreSharedKey `json:"preSharedKeys,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectMacsecPreSharedKey
type InterconnectMacsecPreSharedKey struct {
	// Required. A name for this pre-shared key. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsecPreSharedKey.name
	Name *string `json:"name,omitempty"`

	// A RFC3339 timestamp on or after which the key is valid. startTime can be in the future. If the keychain has a single key, startTime can be omitted. If the keychain has multiple keys, startTime is mandatory for each key. The start times of keys must be in increasing order. The start times of two consecutive keys must be at least 6 hours apart.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsecPreSharedKey.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectOutageNotification
type InterconnectOutageNotification struct {
	// If issue_type is IT_PARTIAL_OUTAGE, a list of the Google-side circuit IDs that will be affected.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.affected_circuits
	AffectedCircuits []string `json:"affectedCircuits,omitempty"`

	// A description about the purpose of the outage.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.description
	Description *string `json:"description,omitempty"`

	// Scheduled end time for the outage (milliseconds since Unix epoch).
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.end_time
	EndTime *int64 `json:"endTime,omitempty"`

	// Form this outage is expected to take, which can take one of the following values: - OUTAGE: The Interconnect may be completely out of service for some or all of the specified window. - PARTIAL_OUTAGE: Some circuits comprising the Interconnect as a whole should remain up, but with reduced bandwidth. Note that the versions of this enum prefixed with "IT_" have been deprecated in favor of the unprefixed values.
	//  Check the IssueType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.issue_type
	IssueType *string `json:"issueType,omitempty"`

	// Unique identifier for this outage notification.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.name
	Name *string `json:"name,omitempty"`

	// The party that generated this notification, which can take the following value: - GOOGLE: this notification as generated by Google. Note that the value of NSRC_GOOGLE has been deprecated in favor of GOOGLE.
	//  Check the Source enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.source
	Source *string `json:"source,omitempty"`

	// Scheduled start time for the outage (milliseconds since Unix epoch).
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.start_time
	StartTime *int64 `json:"startTime,omitempty"`

	// State of this notification, which can take one of the following values: - ACTIVE: This outage notification is active. The event could be in the past, present, or future. See start_time and end_time for scheduling. - CANCELLED: The outage associated with this notification was cancelled before the outage was due to start. - COMPLETED: The outage associated with this notification is complete. Note that the versions of this enum prefixed with "NS_" have been deprecated in favor of the unprefixed values.
	//  Check the State enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint
type NetworkAttachmentConnectedEndpoint struct {
	// The IPv4 address assigned to the producer instance network interface. This value will be a range in case of Serverless.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The IPv6 address assigned to the producer instance network interface. This is only assigned when the stack types of both the instance network interface and the consumer subnet are IPv4_IPv6.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.ipv6_address
	IPV6Address *string `json:"ipv6Address,omitempty"`

	// The project id or number of the interface to which the IP was assigned.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.project_id_or_num
	ProjectIDOrNum *string `json:"projectIDOrNum,omitempty"`

	// Alias IP ranges from the same subnetwork.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.secondary_ip_cidr_ranges
	SecondaryIPCIDRRanges []string `json:"secondaryIPCIDRRanges,omitempty"`

	// The status of a connected endpoint to this network attachment.
	//  Check the Status enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.status
	Status *string `json:"status,omitempty"`

	// The subnetwork used to assign the IP to the producer instance network interface.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// [Output Only] The CIDR range of the subnet from which the IPv4 internal IP was allocated from.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkAttachmentConnectedEndpoint.subnetwork_cidr_range
	SubnetworkCIDRRange *string `json:"subnetworkCIDRRange,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ShareSettings
type ShareSettings struct {

	// TODO: unsupported map type with key string and value message

	// Type of sharing for this shared-reservation
	//  Check the ShareType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.ShareSettings.share_type
	ShareType *string `json:"shareType,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ShareSettingsProjectConfig
type ShareSettingsProjectConfig struct {
	// The project ID, should be same as the key of this project config in the
	//  parent map.
	// +kcc:proto:field=google.cloud.compute.v1.ShareSettingsProjectConfig.project_id
	ProjectID *string `json:"projectID,omitempty"`
}
