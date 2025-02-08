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


// +kcc:proto=google.cloud.video.livestream.v1.Clip
type Clip struct {
	// The resource name of the clip, in the following format:
	//  `projects/{project}/locations/{location}/channels/{c}/clips/{clipId}`.
	//  `{clipId}` is a user-specified resource id that conforms to the following
	//  criteria:
	//
	//  1. 1 character minimum, 63 characters maximum
	//  2. Only contains letters, digits, underscores, and hyphens
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.name
	Name *string `json:"name,omitempty"`

	// The labels associated with this resource. Each label is a key-value pair.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Specify the `output_uri` to determine where to place the clip segments and
	//  clip manifest files in Cloud Storage. The manifests specified in
	//  `clip_manifests` fields will be placed under this URI. The exact URI of the
	//  generated manifests will be provided in `clip_manifests.output_uri` for
	//  each manifest.
	//  Example:
	//  "output_uri": "gs://my-bucket/clip-outputs"
	//  "clip_manifests.output_uri": "gs://my-bucket/clip-outputs/main.m3u8"
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.output_uri
	OutputURI *string `json:"outputURI,omitempty"`

	// The specified ranges of segments to generate a clip.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.slices
	Slices []Clip_Slice `json:"slices,omitempty"`

	// Required. A list of clip manifests. Currently only one clip manifest is
	//  allowed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.clip_manifests
	ClipManifests []Clip_ClipManifest `json:"clipManifests,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Clip.ClipManifest
type Clip_ClipManifest struct {
	// Required. A unique key that identifies a manifest config in the parent
	//  channel. This key is the same as `channel.manifests.key` for the selected
	//  manifest.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.ClipManifest.manifest_key
	ManifestKey *string `json:"manifestKey,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Clip.Slice
type Clip_Slice struct {
	// A slice in form of a tuple of Unix epoch time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.Slice.time_slice
	TimeSlice *Clip_TimeSlice `json:"timeSlice,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Clip.TimeSlice
type Clip_TimeSlice struct {
	// The mark-in Unix epoch time in the original live stream manifest.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.TimeSlice.markin_time
	MarkinTime *string `json:"markinTime,omitempty"`

	// The mark-out Unix epoch time in the original live stream manifest.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.TimeSlice.markout_time
	MarkoutTime *string `json:"markoutTime,omitempty"`
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

// +kcc:proto=google.cloud.video.livestream.v1.Clip
type ClipObservedState struct {
	// Output only. The creation timestamp of the clip resource.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the clip request starts to be processed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The update timestamp of the clip resource.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the clip.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.state
	State *string `json:"state,omitempty"`

	// Output only. An error object that describes the reason for the failure.
	//  This property only presents when `state` is `FAILED`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.error
	Error *Status `json:"error,omitempty"`

	// Required. A list of clip manifests. Currently only one clip manifest is
	//  allowed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.clip_manifests
	ClipManifests []Clip_ClipManifestObservedState `json:"clipManifests,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Clip.ClipManifest
type Clip_ClipManifestObservedState struct {
	// Output only. The output URI of the generated clip manifest. This field
	//  will be populated when the CreateClip request is accepted. Current output
	//  format is provided below but may change in the future. Please read this
	//  field to get the uri to the generated clip manifest. Format:
	//  {clip.output_uri}/{channel.manifest.fileName} Example:
	//  gs://my-bucket/clip-outputs/main.m3u8
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.ClipManifest.output_uri
	OutputURI *string `json:"outputURI,omitempty"`
}
