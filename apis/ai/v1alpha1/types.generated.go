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


// +kcc:proto=google.ai.generativelanguage.v1beta.File
type File struct {

	// Immutable. Identifier. The `File` resource name. The ID (name excluding the
	//  "files/" prefix) can contain up to 40 characters that are lowercase
	//  alphanumeric or dashes (-). The ID cannot start or end with a dash. If the
	//  name is empty on create, a unique name will be generated. Example:
	//  `files/123-456`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.name
	Name *string `json:"name,omitempty"`

	// Optional. The human-readable display name for the `File`. The display name
	//  must be no more than 512 characters in length, including spaces. Example:
	//  "Welcome Image"
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta.VideoMetadata
type VideoMetadata struct {
	// Duration of the video.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.VideoMetadata.video_duration
	VideoDuration *string `json:"videoDuration,omitempty"`
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

// +kcc:proto=google.ai.generativelanguage.v1beta.File
type FileObservedState struct {
	// Output only. Metadata for a video.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.video_metadata
	VideoMetadata *VideoMetadata `json:"videoMetadata,omitempty"`

	// Output only. MIME type of the file.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Output only. Size of the file in bytes.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The timestamp of when the `File` was created.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of when the `File` was last updated.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The timestamp of when the `File` will be deleted. Only set if
	//  the `File` is scheduled to expire.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.expiration_time
	ExpirationTime *string `json:"expirationTime,omitempty"`

	// Output only. SHA-256 hash of the uploaded bytes.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.sha256_hash
	Sha256Hash []byte `json:"sha256Hash,omitempty"`

	// Output only. The uri of the `File`.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.uri
	URI *string `json:"uri,omitempty"`

	// Output only. Processing state of the File.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.state
	State *string `json:"state,omitempty"`

	// Output only. Error status if File processing failed.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta.File.error
	Error *Status `json:"error,omitempty"`
}
