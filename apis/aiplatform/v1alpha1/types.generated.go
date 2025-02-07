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


// +kcc:proto=google.cloud.aiplatform.v1.FeatureViewSync
type FeatureViewSync struct {
	// Identifier. Name of the FeatureViewSync. Format:
	//  `projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}/featureViewSyncs/{feature_view_sync}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureViewSync.SyncSummary
type FeatureViewSync_SyncSummary struct {

	// Lower bound of the system time watermark for the sync job. This is only
	//  set for continuously syncing feature views.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.SyncSummary.system_watermark_time
	SystemWatermarkTime *string `json:"systemWatermarkTime,omitempty"`
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

// +kcc:proto=google.type.Interval
type Interval struct {
	// Optional. Inclusive start of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be the same
	//  or after the start.
	// +kcc:proto:field=google.type.Interval.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Exclusive end of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be before the
	//  end.
	// +kcc:proto:field=google.type.Interval.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureViewSync
type FeatureViewSyncObservedState struct {
	// Output only. Time when this FeatureViewSync is created. Creation of a
	//  FeatureViewSync means that the job is pending / waiting for sufficient
	//  resources but may not have started the actual data transfer yet.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this FeatureViewSync is finished.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.run_time
	RunTime *Interval `json:"runTime,omitempty"`

	// Output only. Final status of the FeatureViewSync.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.final_status
	FinalStatus *Status `json:"finalStatus,omitempty"`

	// Output only. Summary of the sync job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.sync_summary
	SyncSummary *FeatureViewSync_SyncSummary `json:"syncSummary,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureViewSync.SyncSummary
type FeatureViewSync_SyncSummaryObservedState struct {
	// Output only. Total number of rows synced.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.SyncSummary.row_synced
	RowSynced *int64 `json:"rowSynced,omitempty"`

	// Output only. BigQuery slot milliseconds consumed for the sync job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureViewSync.SyncSummary.total_slot
	TotalSlot *int64 `json:"totalSlot,omitempty"`
}
