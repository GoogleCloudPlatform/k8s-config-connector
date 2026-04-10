// Copyright 2026 Google LLC
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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeFutureReservationGVK = GroupVersion.WithKind("ComputeFutureReservation")

type FutureReservationParent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +required
	Zone string `json:"zone"`
}

// ComputeFutureReservationSpec defines the desired state of ComputeFutureReservation
// +kcc:spec:proto=google.cloud.compute.v1beta.FutureReservation
type ComputeFutureReservationSpec struct {
	FutureReservationParent `json:",inline"`

	// Immutable. Optional. The FutureReservation name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Aggregate reservation details for the future reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.aggregate_reservation
	AggregateReservation *AllocationAggregateReservation `json:"aggregateReservation,omitempty"`

	// Future timestamp when the FR auto-created reservations will be deleted by Compute Engine. Format of this field must be a valid href="https://www.ietf.org/rfc/rfc3339.txt">RFC3339 value.
	// GCP computes this fields' value based on other field values [auto_delete_auto_created_reservations, auto_created_reservations_duration].
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.auto_created_reservations_delete_time
	AutoCreatedReservationsDeleteTime *string `json:"autoCreatedReservationsDeleteTime,omitempty"`

	// Specifies the duration of auto-created reservations. It represents relative time to future reservation start_time when auto-created reservations will be automatically deleted by Compute Engine. Duration time unit is represented as a count of seconds and fractions of seconds at nanosecond resolution.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.auto_created_reservations_duration
	AutoCreatedReservationsDuration *Duration `json:"autoCreatedReservationsDuration,omitempty"`

	// Setting for enabling or disabling automatic deletion for auto-created reservation. If set to true, auto-created reservations will be deleted at Future Reservation's end time (default) or at user's defined timestamp if any of the [auto_created_reservations_delete_time, auto_created_reservations_duration] values is specified. For keeping auto-created reservation indefinitely, this value should be set to false.
	// GCP resets this to false in DRAFTING state.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.auto_delete_auto_created_reservations
	AutoDeleteAutoCreatedReservations *bool `json:"autoDeleteAutoCreatedReservations,omitempty"`

	// If not present, then FR will not deliver a new commitment or update an existing commitment.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.commitment_info
	CommitmentInfo *FutureReservationCommitmentInfo `json:"commitmentInfo,omitempty"`

	// Type of the deployment requested as part of future reservation.
	//  Check the DeploymentType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.deployment_type
	DeploymentType *string `json:"deploymentType,omitempty"`

	// An optional description of this resource. Provide this property when you create the future reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.description
	Description *string `json:"description,omitempty"`

	// Indicates if this group of VMs have emergent maintenance enabled.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.enable_emergent_maintenance
	EnableEmergentMaintenance *bool `json:"enableEmergentMaintenance,omitempty"`

	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.name
	Name *string `json:"name,omitempty"`

	// Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.name_prefix
	NamePrefix *string `json:"namePrefix,omitempty"`

	// Planning state before being submitted for evaluation. This field is NOT required to be set, GCP manages lifecycle state transitions and will set it accordingly.
	//  Check the PlanningStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.planning_status
	PlanningStatus *string `json:"planningStatus,omitempty"`

	// The reservation mode which determines reservation-termination behavior and expected pricing.
	//  Check the ReservationMode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.reservation_mode
	ReservationMode *string `json:"reservationMode,omitempty"`

	// Name of reservations where the capacity is provisioned at the time of delivery of future reservations. If the reservation with the given name does not exist already, it is created automatically at the time of Approval with INACTIVE state till specified start-time. Either provide the reservation_name or a name_prefix.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.reservation_name
	ReservationName *string `json:"reservationName,omitempty"`

	// Maintenance information for this reservation
	//  Check the SchedulingType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.scheduling_type
	SchedulingType *string `json:"schedulingType,omitempty"`

	// List of Projects/Folders to share with.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.share_settings
	ShareSettings *ShareSettings `json:"shareSettings,omitempty"`

	// Indicates whether the auto-created reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from the delivered reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.specific_reservation_required
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty"`

	// Future Reservation configuration to indicate instance properties and total count.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.specific_sku_properties
	SpecificSkuProperties *FutureReservationSpecificSkuProperties `json:"specificSkuProperties,omitempty"`

	// Time window for this Future Reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.time_window
	TimeWindow *FutureReservationTimeWindow `json:"timeWindow,omitempty"`
}

// ComputeFutureReservationStatus defines the config connector machine state of ComputeFutureReservation
type ComputeFutureReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeFutureReservation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeFutureReservationObservedState `json:"observedState,omitempty"`
}

// ComputeFutureReservationObservedState is the state of the ComputeFutureReservation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1beta.FutureReservation
type ComputeFutureReservationObservedState struct {
	// [Output Only] The current status of the requested amendment.
	//  Check the AmendmentStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.amendment_status
	AmendmentStatus *string `json:"amendmentStatus,omitempty"`

	// Fully qualified urls of the automatically created reservations at start_time.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.auto_created_reservations
	AutoCreatedReservations []string `json:"autoCreatedReservations,omitempty"`

	// [Output Only] Represents the existing matching usage for the future reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.existing_matching_usage_info
	ExistingMatchingUsageInfo *FutureReservationStatusExistingMatchingUsageInfo `json:"existingMatchingUsageInfo,omitempty"`

	// This count indicates the fulfilled capacity so far. This is set during "PROVISIONING" state. This count also includes capacity delivered as part of existing matching reservations.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.fulfilled_count
	FulfilledCount *int64 `json:"fulfilledCount,omitempty"`

	// [Output Only] This field represents the future reservation before an amendment was requested. If the amendment is declined, the Future Reservation will be reverted to the last known good state. The last known good state is not set when updating a future reservation whose Procurement Status is DRAFTING.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.last_known_good_state
	LastKnownGoodState *FutureReservationStatusLastKnownGoodState `json:"lastKnownGoodState,omitempty"`

	// Time when Future Reservation would become LOCKED, after which no modifications to Future Reservation will be allowed. Applicable only after the Future Reservation is in the APPROVED state. The lock_time is an RFC3339 string. The procurement_status will transition to PROCURING state at this time.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.lock_time
	LockTime *string `json:"lockTime,omitempty"`

	// Current state of this Future Reservation
	//  Check the ProcurementStatus enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.procurement_status
	ProcurementStatus *string `json:"procurementStatus,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservationStatus.specific_sku_properties
	SpecificSkuProperties *FutureReservationStatusSpecificSkuProperties `json:"specificSkuProperties,omitempty"`

	// [Output Only] The creation timestamp for this future reservation in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] A unique identifier for this future reservation. The server defines this identifier.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#futureReservation for future reservations.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] Server-defined fully-qualified URL for this resource.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource with the resource id.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// [Output Only] URL of the Zone where this future reservation resides.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.zone
	Zone *string `json:"zone,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputefuturereservation;gcpcomputefuturereservations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeFutureReservation is the Schema for the ComputeFutureReservation API
// +k8s:openapi-gen=true
type ComputeFutureReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeFutureReservationSpec   `json:"spec,omitempty"`
	Status ComputeFutureReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeFutureReservationList contains a list of ComputeFutureReservation
type ComputeFutureReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeFutureReservation `json:"items"`
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

func init() {
	SchemeBuilder.Register(&ComputeFutureReservation{}, &ComputeFutureReservationList{})
}
