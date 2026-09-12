// Copyright 2026 Google LLC
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

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExampleStore
type ExampleStore struct {
	// Identifier. The resource name of the ExampleStore. This is a unique
	//  identifier. Format:
	//  projects/{project}/locations/{location}/exampleStores/{example_store}
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.name
	Name *string `json:"name,omitempty"`

	// Required. Display name of the ExampleStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the ExampleStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.description
	Description *string `json:"description,omitempty"`

	// Required. Example Store config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.example_store_config
	ExampleStoreConfig *ExampleStoreConfig `json:"exampleStoreConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExampleStoreConfig
type ExampleStoreConfig struct {
	// Required. The embedding model to be used for vector embedding.
	//  Immutable.
	//  Supported models:
	//  * "textembedding-gecko@003"
	//  * "text-embedding-004"
	//  * "text-embedding-005"
	//  * "text-multilingual-embedding-002"
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStoreConfig.vertex_embedding_model
	VertexEmbeddingModel *string `json:"vertexEmbeddingModel,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.ExampleStore
type ExampleStoreObservedState struct {
	// Output only. Timestamp when this ExampleStore was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ExampleStore was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExampleStore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
