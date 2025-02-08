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


// +kcc:proto=google.cloud.discoveryengine.v1.Document
type Document struct {
	// The structured JSON data for the document. It should conform to the
	//  registered [Schema][google.cloud.discoveryengine.v1.Schema] or an
	//  `INVALID_ARGUMENT` error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.struct_data
	StructData map[string]string `json:"structData,omitempty"`

	// The JSON string representation of the document. It should conform to the
	//  registered [Schema][google.cloud.discoveryengine.v1.Schema] or an
	//  `INVALID_ARGUMENT` error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.json_data
	JsonData *string `json:"jsonData,omitempty"`

	// Immutable. The full resource name of the document.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.name
	Name *string `json:"name,omitempty"`

	// Immutable. The identifier of the document.
	//
	//  Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034)
	//  standard with a length limit of 63 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.id
	ID *string `json:"id,omitempty"`

	// The identifier of the schema located in the same data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.schema_id
	SchemaID *string `json:"schemaID,omitempty"`

	// The unstructured data linked to this document. Content must be set if this
	//  document is under a
	//  `CONTENT_REQUIRED` data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.content
	Content *Document_Content `json:"content,omitempty"`

	// The identifier of the parent document. Currently supports at most two level
	//  document hierarchy.
	//
	//  Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034)
	//  standard with a length limit of 63 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.parent_document_id
	ParentDocumentID *string `json:"parentDocumentID,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Document.Content
type Document_Content struct {
	// The content represented as a stream of bytes. The maximum length is
	//  1,000,000 bytes (1 MB / ~0.95 MiB).
	//
	//  Note: As with all `bytes` fields, this field is represented as pure
	//  binary in Protocol Buffers and base64-encoded string in JSON. For
	//  example, `abc123!?$*&()'-=@~` should be represented as
	//  `YWJjMTIzIT8kKiYoKSctPUB+` in JSON. See
	//  https://developers.google.com/protocol-buffers/docs/proto3#json.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.Content.raw_bytes
	RawBytes []byte `json:"rawBytes,omitempty"`

	// The URI of the content. Only Cloud Storage URIs (e.g.
	//  `gs://bucket-name/path/to/file`) are supported. The maximum file size
	//  is 2.5 MB for text-based formats, 200 MB for other formats.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.Content.uri
	URI *string `json:"uri,omitempty"`

	// The MIME type of the content. Supported types:
	//
	//  * `application/pdf` (PDF, only native PDFs are supported for now)
	//  * `text/html` (HTML)
	//  * `application/vnd.openxmlformats-officedocument.wordprocessingml.document` (DOCX)
	//  * `application/vnd.openxmlformats-officedocument.presentationml.presentation` (PPTX)
	//  * `text/plain` (TXT)
	//
	//  See https://www.iana.org/assignments/media-types/media-types.xhtml.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.Content.mime_type
	MimeType *string `json:"mimeType,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Document.IndexStatus
type Document_IndexStatus struct {
	// The time when the document was indexed.
	//  If this field is populated, it means the document has been indexed.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.IndexStatus.index_time
	IndexTime *string `json:"indexTime,omitempty"`

	// A sample of errors encountered while indexing the document.
	//  If this field is populated, the document is not indexed due to errors.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.IndexStatus.error_samples
	ErrorSamples []Status `json:"errorSamples,omitempty"`
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

// +kcc:proto=google.cloud.discoveryengine.v1.Document
type DocumentObservedState struct {
	// Output only. This field is OUTPUT_ONLY.
	//  It contains derived data that are not in the original input document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.derived_struct_data
	DerivedStructData map[string]string `json:"derivedStructData,omitempty"`

	// Output only. The last time the document was indexed. If this field is set,
	//  the document could be returned in search results.
	//
	//  This field is OUTPUT_ONLY. If this field is not populated, it means the
	//  document has never been indexed.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.index_time
	IndexTime *string `json:"indexTime,omitempty"`

	// Output only. The index status of the document.
	//
	//  * If document is indexed successfully, the index_time field is populated.
	//  * Otherwise, if document is not indexed due to errors, the error_samples
	//    field is populated.
	//  * Otherwise, index_status is unset.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.index_status
	IndexStatus *Document_IndexStatus `json:"indexStatus,omitempty"`
}
