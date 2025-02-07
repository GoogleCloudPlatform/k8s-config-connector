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


// +kcc:proto=google.cloud.apigeeregistry.v1.Artifact
type Artifact struct {
	// Resource name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.name
	Name *string `json:"name,omitempty"`

	// A content type specifier for the artifact.
	//  Content type specifiers are Media Types
	//  (https://en.wikipedia.org/wiki/Media_type) with a possible "schema"
	//  parameter that specifies a schema for the stored information.
	//  Content types can specify compression. Currently only GZip compression is
	//  supported (indicated with "+gzip").
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Input only. The contents of the artifact.
	//  Provided by API callers when artifacts are created or replaced.
	//  To access the contents of an artifact, use GetArtifactContents.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.contents
	Contents []byte `json:"contents,omitempty"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.Artifact
type ArtifactObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The size of the artifact in bytes. If the artifact is gzipped, this is
	//  the size of the uncompressed artifact.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.size_bytes
	SizeBytes *int32 `json:"sizeBytes,omitempty"`

	// Output only. A SHA-256 hash of the artifact's contents. If the artifact is gzipped,
	//  this is the hash of the uncompressed artifact.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Artifact.hash
	Hash *string `json:"hash,omitempty"`
}
