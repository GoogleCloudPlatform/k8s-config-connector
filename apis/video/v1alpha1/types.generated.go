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


// +kcc:proto=google.cloud.video.livestream.v1.Event
type Event struct {
	// The resource name of the event, in the form of:
	//  `projects/{project}/locations/{location}/channels/{channelId}/events/{eventId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.name
	Name *string `json:"name,omitempty"`

	// User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Switches to another input stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.input_switch
	InputSwitch *Event_InputSwitchTask `json:"inputSwitch,omitempty"`

	// Inserts a new ad opportunity.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.ad_break
	AdBreak *Event_AdBreakTask `json:"adBreak,omitempty"`

	// Stops any running ad break.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.return_to_program
	ReturnToProgram *Event_ReturnToProgramTask `json:"returnToProgram,omitempty"`

	// Inserts a slate.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.slate
	Slate *Event_SlateTask `json:"slate,omitempty"`

	// Mutes the stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.mute
	Mute *Event_MuteTask `json:"mute,omitempty"`

	// Unmutes the stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.unmute
	Unmute *Event_UnmuteTask `json:"unmute,omitempty"`

	// When this field is set to true, the event will be executed at the earliest
	//  time that the server can schedule the event and
	//  [execution_time][google.cloud.video.livestream.v1.Event.execution_time]
	//  will be populated with the time that the server actually schedules the
	//  event.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.execute_now
	ExecuteNow *bool `json:"executeNow,omitempty"`

	// The time to execute the event. If you set
	//  [execute_now][google.cloud.video.livestream.v1.Event.execute_now] to
	//  `true`, then do not set this field in the `CreateEvent` request. In
	//  this case, the server schedules the event and populates this field. If you
	//  set [execute_now][google.cloud.video.livestream.v1.Event.execute_now] to
	//  `false`, then you must set this field to at least 10 seconds in the future
	//  or else the event can't be created.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.execution_time
	ExecutionTime *string `json:"executionTime,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Event.AdBreakTask
type Event_AdBreakTask struct {
	// Duration of an ad opportunity. Must be greater than 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.AdBreakTask.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Event.InputSwitchTask
type Event_InputSwitchTask struct {
	// The
	//  [InputAttachment.key][google.cloud.video.livestream.v1.InputAttachment.key]
	//  of the input to switch to.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.InputSwitchTask.input_key
	InputKey *string `json:"inputKey,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Event.MuteTask
type Event_MuteTask struct {
	// Duration for which the stream should be muted. If omitted, the stream
	//  will be muted until an UnmuteTask event is sent.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.MuteTask.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Event.ReturnToProgramTask
type Event_ReturnToProgramTask struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Event.SlateTask
type Event_SlateTask struct {
	// Optional. Duration of the slate. Must be greater than 0 if specified.
	//  Omit this field for a long running slate.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.SlateTask.duration
	Duration *string `json:"duration,omitempty"`

	// Slate asset to use for the duration. If its duration is less than the
	//  duration of the SlateTask, then the slate loops. The slate must be
	//  represented in the form of:
	//  `projects/{project}/locations/{location}/assets/{assetId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.SlateTask.asset
	Asset *string `json:"asset,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Event.UnmuteTask
type Event_UnmuteTask struct {
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

// +kcc:proto=google.cloud.video.livestream.v1.Event
type EventObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the event.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.state
	State *string `json:"state,omitempty"`

	// Output only. An error object that describes the reason for the failure.
	//  This property is always present when `state` is `FAILED`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Event.error
	Error *Status `json:"error,omitempty"`
}
