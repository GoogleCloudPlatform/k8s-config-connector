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


// +kcc:proto=google.firestore.admin.v1.Index
type Index struct {
	// Output only. A server defined name for this index.
	//  The form of this name for composite indexes will be:
	//  `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}`
	//  For single field indexes, this field will be empty.
	// +kcc:proto:field=google.firestore.admin.v1.Index.name
	Name *string `json:"name,omitempty"`

	// Indexes with a collection query scope specified allow queries
	//  against a collection that is the child of a specific document, specified at
	//  query time, and that has the same collection ID.
	//
	//  Indexes with a collection group query scope specified allow queries against
	//  all collections descended from a specific document, specified at query
	//  time, and that have the same collection ID as this index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.query_scope
	QueryScope *string `json:"queryScope,omitempty"`

	// The API scope supported by this index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.api_scope
	ApiScope *string `json:"apiScope,omitempty"`

	// The fields supported by this index.
	//
	//  For composite indexes, this requires a minimum of 2 and a maximum of 100
	//  fields. The last field entry is always for the field path `__name__`. If,
	//  on creation, `__name__` was not specified as the last field, it will be
	//  added automatically with the same direction as that of the last field
	//  defined. If the final field in a composite index is not directional, the
	//  `__name__` will be ordered ASCENDING (unless explicitly specified).
	//
	//  For single field indexes, this will always be exactly one entry with a
	//  field path equal to the field path of the associated field.
	// +kcc:proto:field=google.firestore.admin.v1.Index.fields
	Fields []Index_IndexField `json:"fields,omitempty"`

	// Output only. The serving state of the index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index.IndexField
type Index_IndexField struct {
	// Can be __name__.
	//  For single field indexes, this must match the name of the field or may
	//  be omitted.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.field_path
	FieldPath *string `json:"fieldPath,omitempty"`

	// Indicates that this field supports ordering by the specified order or
	//  comparing using =, !=, <, <=, >, >=.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.order
	Order *string `json:"order,omitempty"`

	// Indicates that this field supports operations on `array_value`s.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.array_config
	ArrayConfig *string `json:"arrayConfig,omitempty"`

	// Indicates that this field supports nearest neighbor and distance
	//  operations on vector.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.vector_config
	VectorConfig *Index_IndexField_VectorConfig `json:"vectorConfig,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index.IndexField.VectorConfig
type Index_IndexField_VectorConfig struct {
	// Required. The vector dimension this configuration applies to.
	//
	//  The resulting index will only include vectors of this dimension, and
	//  can be used for vector search with the same dimension.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.VectorConfig.dimension
	Dimension *int32 `json:"dimension,omitempty"`

	// Indicates the vector index is a flat index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.VectorConfig.flat
	Flat *Index_IndexField_VectorConfig_FlatIndex `json:"flat,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index.IndexField.VectorConfig.FlatIndex
type Index_IndexField_VectorConfig_FlatIndex struct {
}
