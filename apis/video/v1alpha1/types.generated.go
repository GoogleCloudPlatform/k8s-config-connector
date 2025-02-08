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


// +kcc:proto=google.cloud.video.stitcher.v1.AdStitchDetail
type AdStitchDetail struct {
	// Required. The ad break ID of the processed ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AdStitchDetail.ad_break_id
	AdBreakID *string `json:"adBreakID,omitempty"`

	// Required. The ad ID of the processed ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AdStitchDetail.ad_id
	AdID *string `json:"adID,omitempty"`

	// Required. The time offset of the processed ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AdStitchDetail.ad_time_offset
	AdTimeOffset *string `json:"adTimeOffset,omitempty"`

	// Optional. Indicates the reason why the ad has been skipped.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AdStitchDetail.skip_reason
	SkipReason *string `json:"skipReason,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.video.stitcher.v1.VodStitchDetail
type VodStitchDetail struct {
	// The name of the stitch detail in the specified VOD session, in the form of
	//  `projects/{project}/locations/{location}/vodSessions/{vod_session_id}/vodStitchDetails/{id}`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodStitchDetail.name
	Name *string `json:"name,omitempty"`

	// A list of ad processing details for the fetched ad playlist.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodStitchDetail.ad_stitch_details
	AdStitchDetails []AdStitchDetail `json:"adStitchDetails,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}
