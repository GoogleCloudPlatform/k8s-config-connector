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

// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation
type Reservation struct {
	// The resource name of the reservation, e.g.,
	//  `projects/*/locations/*/reservations/team1-prod`.
	//  The reservation_id must only contain lower case alphanumeric characters or
	//  dashes. It must start with a letter and must not end with a dash. Its
	//  maximum length is 64 characters.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.name
	Name *string `json:"name,omitempty"`

	// Baseline slots available to this reservation. A slot is a unit of
	//  computational power in BigQuery, and serves as the unit of parallelism.
	//
	//  Queries using this reservation might use more slots during runtime if
	//  ignore_idle_slots is set to false, or autoscaling is enabled.
	//
	//  If edition is EDITION_UNSPECIFIED and total slot_capacity of the
	//  reservation and its siblings exceeds the total slot_count of all capacity
	//  commitments, the request will fail with
	//  `google.rpc.Code.RESOURCE_EXHAUSTED`.
	//
	//  If edition is any value but EDITION_UNSPECIFIED, then the above requirement
	//  is not needed. The total slot_capacity of the reservation and its siblings
	//  may exceed the total slot_count of capacity commitments. In that case, the
	//  exceeding slots will be charged with the autoscale SKU. You can increase
	//  the number of baseline slots in a reservation every few minutes. If you
	//  want to decrease your baseline slots, you are limited to once an hour if
	//  you have recently changed your baseline slot capacity and your baseline
	//  slots exceed your committed slots. Otherwise, you can decrease your
	//  baseline slots every few minutes.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.slot_capacity
	SlotCapacity *int64 `json:"slotCapacity,omitempty"`

	// If false, any query or pipeline job using this reservation will use idle
	//  slots from other reservations within the same admin project. If true, a
	//  query or pipeline job using this reservation will execute with the slot
	//  capacity specified in the slot_capacity field at most.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.ignore_idle_slots
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty"`

	// The configuration parameters for the auto scaling feature.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.autoscale
	Autoscale *Reservation_Autoscale `json:"autoscale,omitempty"`

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

	// Applicable only for reservations located within one of the BigQuery
	//  multi-regions (US or EU).
	//
	//  If set to true, this reservation is placed in the organization's
	//  secondary region which is designated for disaster recovery purposes.
	//  If false, this reservation is placed in the organization's default region.
	//
	//  NOTE: this is a preview feature. Project must be allow-listed in order to
	//  set this field.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.multi_region_auxiliary
	MultiRegionAuxiliary *bool `json:"multiRegionAuxiliary,omitempty"`

	// Edition of the reservation.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.edition
	Edition *string `json:"edition,omitempty"`

	// Optional. The current location of the reservation's primary replica. This
	//  field is only set for reservations using the managed disaster recovery
	//  feature.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.primary_location
	PrimaryLocation *string `json:"primaryLocation,omitempty"`

	// Optional. The current location of the reservation's secondary replica. This
	//  field is only set for reservations using the managed disaster recovery
	//  feature. Users can set this in create reservation calls
	//  to create a failover reservation or in update reservation calls to convert
	//  a non-failover reservation to a failover reservation(or vice versa).
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.secondary_location
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`

	// Optional. The location where the reservation was originally created. This
	//  is set only during the failover reservation's creation. All billing charges
	//  for the failover reservation will be applied to this location.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.original_primary_location
	OriginalPrimaryLocation *string `json:"originalPrimaryLocation,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation.Autoscale
type Reservation_Autoscale struct {

	// Number of slots to be scaled when needed.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.Autoscale.max_slots
	MaxSlots *int64 `json:"maxSlots,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation
type ReservationObservedState struct {
	// The configuration parameters for the auto scaling feature.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.autoscale
	Autoscale *Reservation_AutoscaleObservedState `json:"autoscale,omitempty"`

	// Output only. Creation time of the reservation.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.creation_time
	CreationTime *string `json:"creationTime,omitempty"`

	// Output only. Last update time of the reservation.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation.Autoscale
type Reservation_AutoscaleObservedState struct {
	// Output only. The slot capacity added to this reservation when autoscale
	//  happens. Will be between [0, max_slots]. Note: after users reduce
	//  max_slots, it may take a while before it can be propagated, so
	//  current_slots may stay in the original value and could be larger than
	//  max_slots for that brief period (less than one minute)
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.Autoscale.current_slots
	CurrentSlots *int64 `json:"currentSlots,omitempty"`
}
