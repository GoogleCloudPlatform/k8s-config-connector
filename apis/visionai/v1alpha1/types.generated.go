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


// +kcc:proto=google.cloud.visionai.v1.Stream
type Stream struct {
	// Name of the resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations to allow clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// The display name for the stream resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Whether to enable the HLS playback service on this stream.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.enable_hls_playback
	EnableHlsPlayback *bool `json:"enableHlsPlayback,omitempty"`

	// The name of the media warehouse asset for long term storage of stream data.
	//  Format: projects/${p_id}/locations/${l_id}/corpora/${c_id}/assets/${a_id}
	//  Remain empty if the media warehouse storage is not needed for the stream.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.media_warehouse_asset
	MediaWarehouseAsset *string `json:"mediaWarehouseAsset,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Stream
type StreamObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
