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


// +kcc:proto=google.cloud.apigeeregistry.v1.ApiSpec
type ApiSpec struct {
	// Resource name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.name
	Name *string `json:"name,omitempty"`

	// A possibly-hierarchical name used to refer to the spec from other specs.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.filename
	Filename *string `json:"filename,omitempty"`

	// A detailed description.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.description
	Description *string `json:"description,omitempty"`

	// A style (format) descriptor for this spec that is specified as a Media Type
	//  (https://en.wikipedia.org/wiki/Media_type). Possible values include
	//  `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and
	//  `application/vnd.apigee.graphql`, with possible suffixes representing
	//  compression types. These hypothetical names are defined in the vendor tree
	//  defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final.
	//  Content types can specify compression. Currently only GZip compression is
	//  supported (indicated with "+gzip").
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// The original source URI of the spec (if one exists).
	//  This is an external location that can be used for reference purposes
	//  but which may not be authoritative since this external resource may
	//  change after the spec is retrieved.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.source_uri
	SourceURI *string `json:"sourceURI,omitempty"`

	// Input only. The contents of the spec.
	//  Provided by API callers when specs are created or updated.
	//  To access the contents of a spec, use GetApiSpecContents.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.contents
	Contents []byte `json:"contents,omitempty"`

	// Labels attach identifying metadata to resources. Identifying metadata can
	//  be used to filter list operations.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one resource (System
	//  labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with
	//  `apigeeregistry.googleapis.com/` and cannot be changed.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations attach non-identifying metadata to resources.
	//
	//  Annotation keys and values are less restricted than those of labels, but
	//  should be generally used for small values of broad interest. Larger, topic-
	//  specific metadata should be stored in Artifacts.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.ApiSpec
type ApiSpecObservedState struct {
	// Output only. Immutable. The revision ID of the spec.
	//  A new revision is committed whenever the spec contents are changed.
	//  The format is an 8-character hexadecimal string.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. Creation timestamp; when the spec resource was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Revision creation timestamp; when the represented revision was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. Last update timestamp: when the represented revision was last modified.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.revision_update_time
	RevisionUpdateTime *string `json:"revisionUpdateTime,omitempty"`

	// Output only. The size of the spec file in bytes. If the spec is gzipped, this is the
	//  size of the uncompressed spec.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.size_bytes
	SizeBytes *int32 `json:"sizeBytes,omitempty"`

	// Output only. A SHA-256 hash of the spec's contents. If the spec is gzipped, this is
	//  the hash of the uncompressed spec.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.hash
	Hash *string `json:"hash,omitempty"`
}
