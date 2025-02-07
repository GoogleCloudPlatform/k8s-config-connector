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


// +kcc:proto=google.cloud.aiplatform.v1.Context
type Context struct {
	// Immutable. The resource name of the Context.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.name
	Name *string `json:"name,omitempty"`

	// User provided display name of the Context.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Contexts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Context (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Context.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Description of the Context
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Context
type ContextObservedState struct {
	// Output only. Timestamp when this Context was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Context was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A list of resource names of Contexts that are parents of this
	//  Context. A Context may have at most 10 parent_contexts.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.parent_contexts
	ParentContexts []string `json:"parentContexts,omitempty"`
}
