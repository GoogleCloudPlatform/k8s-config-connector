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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FutureReservationGVK = GroupVersion.WithKind("FutureReservation")

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +required
	Location string `json:"location"`
}

// FutureReservationSpec defines the desired state of FutureReservation
// +kcc:spec:proto=google.cloud.compute.v1beta.FutureReservation
type FutureReservationSpec struct {
	Parent `json:",inline"`

	// The FutureReservation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Aggregate reservation details for the future reservation.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.aggregate_reservation
	AggregateReservation *AllocationAggregateReservation `json:"aggregateReservation,omitempty"`

	// Future timestamp when the FR auto-created reservations will be deleted by Compute Engine. Format of this field must be a valid href="https://www.ietf.org/rfc/rfc3339.txt">RFC3339 value.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.auto_created_reservations_delete_time
	AutoCreatedReservationsDeleteTime *string `json:"autoCreatedReservationsDeleteTime,omitempty"`

	// Specifies the duration of auto-created reservations. It represents relative time to future reservation start_time when auto-created reservations will be automatically deleted by Compute Engine. Duration time unit is represented as a count of seconds and fractions of seconds at nanosecond resolution.
	// +kcc:proto:field=google.cloud.compute.v1beta.FutureReservation.auto_created_reservations_duration
	AutoCreatedReservationsDuration *Duration `json:"autoCreatedReservationsDuration,omitempty"`

	// Setting for enabling or disabling automatic deletion for auto-created reservation. If set to true, auto-created reservations will be deleted at Future Reservation's end time (default) or at user's defined timestamp if any of the [auto_created_reservations_delete_time, auto_created_reservations_duration] values is specified. For keeping auto-created reservation indefinitely, this value should be set to false.
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

	// Planning state before being submitted for evaluation
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

// FutureReservationStatus defines the config connector machine state of FutureReservation
type FutureReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FutureReservation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FutureReservationObservedState `json:"observedState,omitempty"`
}

// FutureReservationObservedState is the state of the FutureReservation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1beta.FutureReservation
type FutureReservationObservedState struct {
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
// +kubebuilder:resource:categories=gcp,shortName=gcpfuturereservation;gcpfuturereservations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FutureReservation is the Schema for the FutureReservation API
// +k8s:openapi-gen=true
type FutureReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FutureReservationSpec   `json:"spec,omitempty"`
	Status FutureReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FutureReservationList contains a list of FutureReservation
type FutureReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FutureReservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FutureReservation{}, &FutureReservationList{})
}
