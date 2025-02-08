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


// +kcc:proto=google.cloud.dialogflow.v2beta1.Document
type Document struct {
	// Optional. The document resource name.
	//  The name must be empty when creating a document.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/knowledgeBases/<Knowledge Base ID>/documents/<Document ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the document. The name must be 1024 bytes or
	//  less; otherwise, the creation request fails.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The MIME type of this document.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Required. The knowledge type of document content.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.knowledge_types
	KnowledgeTypes []string `json:"knowledgeTypes,omitempty"`

	// The URI where the file content is located.
	//
	//  For documents stored in Google Cloud Storage, these URIs must have
	//  the form `gs://<bucket-name>/<object-name>`.
	//
	//  NOTE: External URLs must correspond to public webpages, i.e., they must
	//  be indexed by Google Search. In particular, URLs for showing documents in
	//  Google Cloud Storage (i.e. the URL in your browser) are not supported.
	//  Instead use the `gs://` format URI described above.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.content_uri
	ContentURI *string `json:"contentURI,omitempty"`

	// The raw content of the document. This field is only permitted for
	//  EXTRACTIVE_QA and FAQ knowledge types.
	//  Note: This field is in the process of being deprecated, please use
	//  raw_content instead.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.content
	Content *string `json:"content,omitempty"`

	// The raw content of the document. This field is only permitted for
	//  EXTRACTIVE_QA and FAQ knowledge types.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.raw_content
	RawContent []byte `json:"rawContent,omitempty"`

	// Optional. If true, we try to automatically reload the document every day
	//  (at a time picked by the system). If false or unspecified, we don't try
	//  to automatically reload the document.
	//
	//  Currently you can only enable automatic reload for documents sourced from
	//  a public url, see `source` field for the source types.
	//
	//  Reload status can be tracked in `latest_reload_status`. If a reload
	//  fails, we will keep the document unchanged.
	//
	//  If a reload fails with internal errors, the system will try to reload the
	//  document on the next day.
	//  If a reload fails with non-retriable errors (e.g. PERMISSION_DENIED), the
	//  system will not try to reload the document anymore. You need to manually
	//  reload the document successfully by calling `ReloadDocument` and clear the
	//  errors.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.enable_auto_reload
	EnableAutoReload *bool `json:"enableAutoReload,omitempty"`

	// Optional. Metadata for the document. The metadata supports arbitrary
	//  key-value pairs. Suggested use cases include storing a document's title,
	//  an external URL distinct from the document's content_uri, etc.
	//  The max size of a `key` or a `value` of the metadata is 1024 bytes.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Document.ReloadStatus
type Document_ReloadStatus struct {
	// Output only. The time of a reload attempt.
	//  This reload may have been triggered automatically or manually and may
	//  not have succeeded.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.ReloadStatus.time
	Time *string `json:"time,omitempty"`

	// Output only. The status of a reload attempt or the initial load.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.ReloadStatus.status
	Status *Status `json:"status,omitempty"`
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

// +kcc:proto=google.cloud.dialogflow.v2beta1.Document
type DocumentObservedState struct {
	// Output only. The time and status of the latest reload.
	//  This reload may have been triggered automatically or manually
	//  and may not have succeeded.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.latest_reload_status
	LatestReloadStatus *Document_ReloadStatus `json:"latestReloadStatus,omitempty"`

	// Output only. The current state of the document.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Document.state
	State *string `json:"state,omitempty"`
}
