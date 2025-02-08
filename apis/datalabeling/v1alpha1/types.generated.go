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


// +kcc:proto=google.cloud.datalabeling.v1beta1.DataItem
type DataItem struct {
	// The image payload, a container of the image bytes/uri.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.DataItem.image_payload
	ImagePayload *ImagePayload `json:"imagePayload,omitempty"`

	// The text payload, a container of text content.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.DataItem.text_payload
	TextPayload *TextPayload `json:"textPayload,omitempty"`

	// The video payload, a container of the video uri.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.DataItem.video_payload
	VideoPayload *VideoPayload `json:"videoPayload,omitempty"`

	// Output only. Name of the data item, in format of:
	//  projects/{project_id}/datasets/{dataset_id}/dataItems/{data_item_id}
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.DataItem.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ImagePayload
type ImagePayload struct {
	// Image format.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// A byte string of a thumbnail image.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.image_thumbnail
	ImageThumbnail []byte `json:"imageThumbnail,omitempty"`

	// Image uri from the user bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Signed uri of the image file in the service bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.signed_uri
	SignedURI *string `json:"signedURI,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextPayload
type TextPayload struct {
	// Text content.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextPayload.text_content
	TextContent *string `json:"textContent,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoPayload
type VideoPayload struct {
	// Video format.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Video uri from the user bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.video_uri
	VideoURI *string `json:"videoURI,omitempty"`

	// The list of video thumbnails.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.video_thumbnails
	VideoThumbnails []VideoThumbnail `json:"videoThumbnails,omitempty"`

	// FPS of the video.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.frame_rate
	FrameRate *float32 `json:"frameRate,omitempty"`

	// Signed uri of the video file in the service bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.signed_uri
	SignedURI *string `json:"signedURI,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoThumbnail
type VideoThumbnail struct {
	// A byte string of the video frame.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoThumbnail.thumbnail
	Thumbnail []byte `json:"thumbnail,omitempty"`

	// Time offset relative to the beginning of the video, corresponding to the
	//  video frame where the thumbnail has been extracted from.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoThumbnail.time_offset
	TimeOffset *string `json:"timeOffset,omitempty"`
}
