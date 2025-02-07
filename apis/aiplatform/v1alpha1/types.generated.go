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


// +kcc:proto=google.cloud.aiplatform.v1.Artifact
type Artifact struct {

	// User provided display name of the Artifact.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The uniform resource identifier of the artifact file.
	//  May be empty if there is no actual artifact file.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.uri
	URI *string `json:"uri,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Artifacts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Artifact (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The state of this Artifact. This is a property of the Artifact, and does
	//  not imply or capture any ongoing process. This property is managed by
	//  clients (such as Vertex AI Pipelines), and the system does not prescribe
	//  or check the validity of state transitions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.state
	State *string `json:"state,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Artifact.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Description of the Artifact
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Artifact
type ArtifactObservedState struct {
	// Output only. The resource name of the Artifact.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this Artifact was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Artifact was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Artifact.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
