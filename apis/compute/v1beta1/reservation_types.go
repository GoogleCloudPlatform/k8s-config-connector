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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeReservationGVK = GroupVersion.WithKind("ComputeReservation")

// ComputeReservationSpec defines the desired state of ComputeReservation
// +kcc:spec:proto=google.cloud.compute.v1.Reservation
type ComputeReservationSpec struct {
	// Immutable. An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.description
	Description *string `json:"description,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Reservation for instances with specific machine shapes.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.specific_reservation
	SpecificReservation *ReservationSpecificReservation `json:"specificReservation"`

	// Immutable. When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it can be consumed by VMs with affinity for any reservation. Defaults to false.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.specific_reservation_required
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty"`

	// Immutable. The share setting for reservations and sole tenancy node groups.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.share_settings
	ShareSettings *ShareSettings `json:"shareSettings,omitempty"`

	// Immutable. The zone where the reservation is made.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.zone
	Zone *string `json:"zone"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationSpecificSKUReservation
type ReservationSpecificReservation struct {
	// The number of resources that are allocated.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUReservation.count
	Count *int32 `json:"count"`

	// How many instances are in use.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUReservation.in_use_count
	InUseCount *int32 `json:"inUseCount,omitempty"`

	// Immutable. The instance properties for the reservation.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUReservation.instance_properties
	InstanceProperties *ReservationInstanceProperties `json:"instanceProperties"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties
type ReservationInstanceProperties struct {
	// Immutable. Guest accelerator type and count.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.guest_accelerators
	GuestAccelerators []ReservationGuestAccelerators `json:"guestAccelerators,omitempty"`

	// Immutable. The amount of local ssd to reserve with each instance. This reserves disks of type 'local-ssd'.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.local_ssds
	LocalSsds []ReservationLocalSsds `json:"localSsds,omitempty"`

	// Immutable. The name of the machine type to reserve.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.machine_type
	MachineType *string `json:"machineType"`

	// Immutable. The minimum CPU platform for the reservation. For example, '"Intel Skylake"'. See the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones) for information on available CPU platforms.
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationReservedInstanceProperties.min_cpu_platform
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AcceleratorConfig
type ReservationGuestAccelerators struct {
	// Immutable. The number of the guest accelerator cards exposed to this instance.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount"`

	// Immutable. The full or partial URL of the accelerator type to attach to this instance. For example: 'projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100' If you are creating an instance template, specify only the accelerator name.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType"`
}

// +kcc:proto=google.cloud.compute.v1.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk
type ReservationLocalSsds struct {
	// Immutable. The size of the disk in base-2 GB.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk.disk_size_gb
	DiskSizeGb *int32 `json:"diskSizeGb"`

	// Immutable. The disk interface to use for attaching this disk. Default value: "SCSI" Possible values: ["SCSI", "NVME"].
	// +kcc:proto:field=google.cloud.compute.v1.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk.interface
	Interface *string `json:"interface,omitempty"`
}

// ComputeReservationStatus defines the config connector machine state of ComputeReservation
type ComputeReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeReservation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.commitment
	Commitment *string `json:"commitment,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// The status of the reservation.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.status
	Status *string `json:"status,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *ComputeReservationObservedState `json:"observedState,omitempty"`
}

// ComputeReservationObservedState is the state of the ComputeReservation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Reservation
/*
type ComputeReservationObservedState struct {
	// [Output Only] A unique identifier for this future reservation. The server defines this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.id
	ID *int64 `json:"id,omitempty"`

	// [Output Only] Allocation Properties of this reservation.
	// +kcc:proto:field=google.cloud.compute.v1.Reservation.resource_status
	ResourceStatus *AllocationResourceStatus `json:"resourceStatus,omitempty"`
}
*/

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputereservation;gcpcomputereservations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeReservation is the Schema for the ComputeReservation API
// +k8s:openapi-gen=true
type ComputeReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeReservationSpec   `json:"spec,omitempty"`
	Status ComputeReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeReservationList contains a list of ComputeReservation
type ComputeReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeReservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeReservation{}, &ComputeReservationList{})
}

// +kcc:proto=google.cloud.compute.v1.ShareSettings
type ShareSettings struct {
	// A map of key(i.e. project or other shared resources) and associated project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	// +kcc:proto:field=google.cloud.compute.v1.ShareSettingsProjectConfig.project_map
	ProjectMap []ShareSettingsProjectMap `json:"projectMap,omitempty"`

	// Type of sharing for this shared-reservation
	//  Check the ShareType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.ShareSettings.share_type
	ShareType *string `json:"shareType,omitempty"`
}

type ShareSettingsProjectMap struct {
	// +required
	KeyRef *refsv1beta1.ExtendedProjectRef `json:"keyRef"`

	Value *ShareSettingsProjectConfig `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ShareSettingsProjectConfig
type ShareSettingsProjectConfig struct {
	// The project ID, should be same as the key of this project config in the
	//  parent map.
	// +kcc:proto:field=google.cloud.compute.v1.ShareSettingsProjectConfig.project_id
	ProjectIDRef *refsv1beta1.ProjectRef `json:"projectIDRef,omitempty"`
}
