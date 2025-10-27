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

// +generated:types
// krm.group: firestore.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.firestore.admin.v1
// resource: FirestoreDocument:google.firestore.v1.Document
// resource: FirestoreField:Field
// resource: FirestoreBackupSchedule:BackupSchedule

package v1alpha1

// +kcc:proto=google.firestore.admin.v1.DailyRecurrence
type DailyRecurrence struct {
}

// +kcc:proto=google.firestore.admin.v1.Field.TtlConfig
type Field_TTLConfig struct {
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

// +kcc:proto=google.firestore.admin.v1.WeeklyRecurrence
type WeeklyRecurrence struct {
	// The day of week to run.
	//
	//  DAY_OF_WEEK_UNSPECIFIED is not allowed.
	// +kcc:proto:field=google.firestore.admin.v1.WeeklyRecurrence.day
	Day *string `json:"day,omitempty"`
}

// +kcc:proto=google.firestore.v1.ArrayValue
type ArrayValue struct {
	// Values in the array.
	// +kcc:proto:field=google.firestore.v1.ArrayValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.firestore.v1.MapValue
type MapValue struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.firestore.v1.Value
type Value struct {
	// A null value.
	// +kcc:proto:field=google.firestore.v1.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// A boolean value.
	// +kcc:proto:field=google.firestore.v1.Value.boolean_value
	BooleanValue *bool `json:"booleanValue,omitempty"`

	// An integer value.
	// +kcc:proto:field=google.firestore.v1.Value.integer_value
	IntegerValue *int64 `json:"integerValue,omitempty"`

	// A double value.
	// +kcc:proto:field=google.firestore.v1.Value.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// A timestamp value.
	//
	//  Precise only to microseconds. When stored, any additional precision is
	//  rounded down.
	// +kcc:proto:field=google.firestore.v1.Value.timestamp_value
	TimestampValue *string `json:"timestampValue,omitempty"`

	// A string value.
	//
	//  The string, represented as UTF-8, must not exceed 1 MiB - 89 bytes.
	//  Only the first 1,500 bytes of the UTF-8 representation are considered by
	//  queries.
	// +kcc:proto:field=google.firestore.v1.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// A bytes value.
	//
	//  Must not exceed 1 MiB - 89 bytes.
	//  Only the first 1,500 bytes are considered by queries.
	// +kcc:proto:field=google.firestore.v1.Value.bytes_value
	BytesValue []byte `json:"bytesValue,omitempty"`

	// A reference to a document. For example:
	//  `projects/{project_id}/databases/{database_id}/documents/{document_path}`.
	// +kcc:proto:field=google.firestore.v1.Value.reference_value
	ReferenceValue *string `json:"referenceValue,omitempty"`

	// A geo point value representing a point on the surface of Earth.
	// +kcc:proto:field=google.firestore.v1.Value.geo_point_value
	GeoPointValue *LatLng `json:"geoPointValue,omitempty"`

	// An array value.
	//
	//  Cannot directly contain another array value, though can contain a
	//  map which contains another array.
	// +kcc:proto:field=google.firestore.v1.Value.array_value
	ArrayValue *ArrayValue `json:"arrayValue,omitempty"`

	// A map value.
	// +kcc:proto:field=google.firestore.v1.Value.map_value
	MapValue *MapValue `json:"mapValue,omitempty"`
}

// +kcc:proto=google.type.LatLng
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	// +kcc:proto:field=google.type.LatLng.latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	// +kcc:proto:field=google.type.LatLng.longitude
	Longitude *float64 `json:"longitude,omitempty"`
}

// +kcc:observedstate:proto=google.firestore.admin.v1.Field.TtlConfig
type Field_TTLConfigObservedState struct {
	// Output only. The state of the TTL configuration.
	// +kcc:proto:field=google.firestore.admin.v1.Field.TtlConfig.state
	State *string `json:"state,omitempty"`
}
