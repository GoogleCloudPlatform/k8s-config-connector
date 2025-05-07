// Copyright 2024 Google LLC
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

package v1alpha1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryReservationReservationGVK = GroupVersion.WithKind("BigQueryReservationReservation")

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// You can configure spec.secondaryLocation to enable the reservation fail-over to a secondary location,
	// in which case the primary location could be different from the spec.location.
	// +required
	Location string `json:"location"`
}

// BigQueryReservationReservationSpec defines the desired state of BigQueryReservationReservation
// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation
type BigQueryReservationReservationSpec struct {
	Parent `json:",inline"`

	// Immutable. Optional.
	// The BigQuery Reservation ID used for resource creation or acquisition.
	// It must only contain lower case alphanumeric
	// characters or dashes. It must start with a letter and must not end
	// with a dash. Its maximum length is 64 characters.
	// For creation: If specified, this value is used as the Reservation ID. If not provided, a UUID will be generated and assigned as the Reservation ID.
	// For acquisition: This field must be provided to identify the Reservation resource to acquire.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Baseline slots available to this reservation. A slot is a unit of
	//  computational power in BigQuery, and serves as the unit of parallelism.
	//
	//  Queries using this reservation might use more slots during runtime if
	//  ignore_idle_slots is set to false, or autoscaling is enabled.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.slot_capacity
	SlotCapacity *int64 `json:"slotCapacity,omitempty"`

	// If false, any query or pipeline job using this reservation will use idle
	//  slots from other reservations within the same admin project. If true, a
	//  query or pipeline job using this reservation will execute with the slot
	//  capacity specified in the slot_capacity field at most.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.ignore_idle_slots
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty"`

	// Optional. The configuration parameters for the auto scaling feature.
	Autoscale *AutoscaleSpec `json:"autoscale,omitempty"`

	// Job concurrency target which sets a soft upper bound on the number of jobs
	//  that can run concurrently in this reservation. This is a soft target due to
	//  asynchronous nature of the system and various optimizations for small
	//  queries.
	//  Default value is 0 which means that concurrency target will be
	//  automatically computed by the system.
	//  NOTE: this field is exposed as target job concurrency in the Information
	//  Schema, DDL and BigQuery CLI.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.concurrency
	Concurrency *int64 `json:"concurrency,omitempty"`

	// Immutable. Optional.
	// Edition of the reservation. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.edition
	Edition *string `json:"edition,omitempty"`

	// Optional. This field is only set for reservations using the managed disaster recovery
	//  feature. Users can set this to create a failover reservation.
	FailOver *FailoverSpec `json:"failover,omitempty"`
}

// BigQueryReservationReservationStatus defines the config connector machine state of BigQueryReservationReservation
type BigQueryReservationReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryReservationReservation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryReservationReservationObservedState `json:"observedState,omitempty"`
}

// BigQueryReservationReservationSpec defines the desired state of BigQueryReservationReservation
// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation
// BigQueryReservationReservationObservedState is the state of the BigQueryReservationReservation resource as most recently observed in GCP.
type BigQueryReservationReservationObservedState struct {
	FailOver  *FailoverObservedState  `json:"failover,omitempty"`
	Autoscale *AutoscaleObservedState `json:"autoscale,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryreservationreservation;gcpbigqueryreservationreservations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryReservationReservation is the Schema for the BigQueryReservationReservation API
// +k8s:openapi-gen=true
type BigQueryReservationReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryReservationReservationSpec   `json:"spec,omitempty"`
	Status BigQueryReservationReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryReservationReservationList contains a list of BigQueryReservationReservation
type BigQueryReservationReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryReservationReservation `json:"items"`
}

// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.autoscale
type AutoscaleSpec struct {
	// Number of slots to be scaled when needed.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.Autoscale.max_slots
	MaxSlots *int64 `json:"maxSlots,omitempty"`
}

type FailoverSpec struct {
	// Users can update this field to convert a non-failover reservation to a
	// failover reservation (by setting a specific region value) or convert a
	// failover reservation to a non-failover reservation (by removing spec.failover).
	// However, changes from one region to another region will be ignored by the
	// controller. Additionally, if the value of this field changes during manual failover
	// by the API, the controller will not attempt to revert these changes.
	//
	// Note: This field is only available for ENTERPRISE_PLUS edition reservations.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="secondaryLocation field is immutable"
	// Immutable.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.secondary_location
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
}

type FailoverObservedState struct {
	// The current location of the reservation's primary replica. This
	//  field is only set for reservations using the managed disaster recovery
	//  feature.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.primary_location
	PrimaryLocation *string `json:"primaryLocation,omitempty"`
	// The current location of the reservation's secondary replica. This
	//  field is only set for reservations using the managed disaster recovery
	//  feature. Users can set this in create reservation calls
	//  to create a failover reservation or in update reservation calls to convert
	//  a non-failover reservation to a failover reservation(or vice versa).
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.secondary_location
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
	// The location where the reservation was originally created. This
	//  is set only during the failover reservation's creation. All billing charges
	//  for the failover reservation will be applied to this location.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.original_primary_location
	OriginalPrimaryLocation *string `json:"originalPrimaryLocation,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation.Autoscale
type AutoscaleObservedState struct {
	// The slot capacity added to this reservation when autoscale
	//  happens. Will be between [0, max_slots]. Note: after users reduce
	//  max_slots, it may take a while before it can be propagated, so
	//  current_slots may stay in the original value and could be larger than
	//  max_slots for that brief period (less than one minute)
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.Autoscale.current_slots
	CurrentSlots *int64 `json:"currentSlots,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BigQueryReservationReservation{}, &BigQueryReservationReservationList{})
}
