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

// +kcc:proto=google.cloud.vmmigration.v1.PostProcessingStep
type PostProcessingStep struct {
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
