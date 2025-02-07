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


// +kcc:proto=google.cloud.aiplatform.v1.MetadataSchema
type MetadataSchema struct {

	// The version of the MetadataSchema. The version's format must match
	//  the following regular expression: `^[0-9]+[.][0-9]+[.][0-9]+$`, which would
	//  allow to order/compare different versions. Example: 1.0.0, 1.0.1, etc.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataSchema.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Required. The raw YAML string representation of the MetadataSchema. The
	//  combination of [MetadataSchema.version] and the schema name given by
	//  `title` in [MetadataSchema.schema] must be unique within a MetadataStore.
	//
	//  The schema is defined as an OpenAPI 3.0.2
	//  [MetadataSchema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataSchema.schema
	Schema *string `json:"schema,omitempty"`

	// The type of the MetadataSchema. This is a property that identifies which
	//  metadata types will use the MetadataSchema.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataSchema.schema_type
	SchemaType *string `json:"schemaType,omitempty"`

	// Description of the Metadata Schema
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataSchema.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.MetadataSchema
type MetadataSchemaObservedState struct {
	// Output only. The resource name of the MetadataSchema.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataSchema.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this MetadataSchema was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MetadataSchema.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
