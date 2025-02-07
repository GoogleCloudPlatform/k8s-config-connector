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


// +kcc:proto=google.cloud.bigquery.reservation.v1.CapacityCommitment
type CapacityCommitment struct {

	// Number of slots in this commitment.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.slot_count
	SlotCount *int64 `json:"slotCount,omitempty"`

	// Capacity commitment commitment plan.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.plan
	Plan *string `json:"plan,omitempty"`

	// The plan this capacity commitment is converted to after commitment_end_time
	//  passes. Once the plan is changed, committed period is extended according to
	//  commitment plan. Only applicable for ANNUAL and TRIAL commitments.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.renewal_plan
	RenewalPlan *string `json:"renewalPlan,omitempty"`

	// Applicable only for commitments located within one of the BigQuery
	//  multi-regions (US or EU).
	//
	//  If set to true, this commitment is placed in the organization's
	//  secondary region which is designated for disaster recovery purposes.
	//  If false, this commitment is placed in the organization's default region.
	//
	//  NOTE: this is a preview feature. Project must be allow-listed in order to
	//  set this field.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.multi_region_auxiliary
	MultiRegionAuxiliary *bool `json:"multiRegionAuxiliary,omitempty"`

	// Edition of the capacity commitment.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.edition
	Edition *string `json:"edition,omitempty"`
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

// +kcc:proto=google.cloud.bigquery.reservation.v1.CapacityCommitment
type CapacityCommitmentObservedState struct {
	// Output only. The resource name of the capacity commitment, e.g.,
	//  `projects/myproject/locations/US/capacityCommitments/123`
	//  The commitment_id must only contain lower case alphanumeric characters or
	//  dashes. It must start with a letter and must not end
	//  with a dash. Its maximum length is 64 characters.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the commitment.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.state
	State *string `json:"state,omitempty"`

	// Output only. The start of the current commitment period. It is applicable
	//  only for ACTIVE capacity commitments. Note after the commitment is renewed,
	//  commitment_start_time won't be changed. It refers to the start time of the
	//  original commitment.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.commitment_start_time
	CommitmentStartTime *string `json:"commitmentStartTime,omitempty"`

	// Output only. The end of the current commitment period. It is applicable
	//  only for ACTIVE capacity commitments. Note after renewal,
	//  commitment_end_time is the time the renewed commitment expires. So it would
	//  be at a time after commitment_start_time + committed period, because we
	//  don't change commitment_start_time ,
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.commitment_end_time
	CommitmentEndTime *string `json:"commitmentEndTime,omitempty"`

	// Output only. For FAILED commitment plan, provides the reason of failure.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.failure_status
	FailureStatus *Status `json:"failureStatus,omitempty"`

	// Output only. If true, the commitment is a flat-rate commitment, otherwise,
	//  it's an edition commitment.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.is_flat_rate
	IsFlatRate *bool `json:"isFlatRate,omitempty"`
}
