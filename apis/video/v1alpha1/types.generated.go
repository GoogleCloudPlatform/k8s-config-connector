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


// +kcc:proto=google.cloud.video.livestream.v1.Asset
type Asset struct {
	// The resource name of the asset, in the form of:
	//  `projects/{project}/locations/{location}/assets/{assetId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.name
	Name *string `json:"name,omitempty"`

	// User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.labels
	Labels map[string]string `json:"labels,omitempty"`

	// VideoAsset represents a video.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.video
	Video *Asset_VideoAsset `json:"video,omitempty"`

	// ImageAsset represents an image.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.image
	Image *Asset_ImageAsset `json:"image,omitempty"`

	// Based64-encoded CRC32c checksum of the asset file. For more information,
	//  see the crc32c checksum of the [Cloud Storage Objects
	//  resource](https://cloud.google.com/storage/docs/json_api/v1/objects).
	//  If crc32c is omitted or left empty when the asset is created, this field is
	//  filled by the crc32c checksum of the Cloud Storage object indicated by
	//  [VideoAsset.uri][google.cloud.video.livestream.v1.Asset.VideoAsset.uri] or
	//  [ImageAsset.uri][google.cloud.video.livestream.v1.Asset.ImageAsset.uri]. If
	//  crc32c is set, the asset can't be created if the crc32c value does not
	//  match with the crc32c checksum of the Cloud Storage object indicated by
	//  [VideoAsset.uri][google.cloud.video.livestream.v1.Asset.VideoAsset.uri] or
	//  [ImageAsset.uri][google.cloud.video.livestream.v1.Asset.ImageAsset.uri].
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.crc32c
	Crc32c *string `json:"crc32c,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Asset.ImageAsset
type Asset_ImageAsset struct {
	// Cloud Storage URI of the image. The format is `gs://my-bucket/my-object`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.ImageAsset.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Asset.VideoAsset
type Asset_VideoAsset struct {
	// Cloud Storage URI of the video. The format is `gs://my-bucket/my-object`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.VideoAsset.uri
	URI *string `json:"uri,omitempty"`
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

// +kcc:proto=google.cloud.video.livestream.v1.Asset
type AssetObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the asset resource.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.state
	State *string `json:"state,omitempty"`

	// Output only. Only present when `state` is `ERROR`. The reason for the error
	//  state of the asset.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.error
	Error *Status `json:"error,omitempty"`
}
