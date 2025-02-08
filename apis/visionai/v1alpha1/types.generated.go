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


// +kcc:proto=google.cloud.visionai.v1.Annotation
type Annotation struct {
	// Resource name of the annotation.
	//  Format:
	//  `projects/{project_number}/locations/{location}/corpora/{corpus}/assets/{asset}/annotations/{annotation}`
	// +kcc:proto:field=google.cloud.visionai.v1.Annotation.name
	Name *string `json:"name,omitempty"`

	// User provided annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.Annotation.user_specified_annotation
	UserSpecifiedAnnotation *UserSpecifiedAnnotation `json:"userSpecifiedAnnotation,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AnnotationCustomizedStruct
type AnnotationCustomizedStruct struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.visionai.v1.AnnotationList
type AnnotationList struct {
	// The values of `LIST` data type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationList.values
	Values []AnnotationValue `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AnnotationValue
type AnnotationValue struct {
	// Value of int type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.int_value
	IntValue *int64 `json:"intValue,omitempty"`

	// Value of float type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.float_value
	FloatValue *float32 `json:"floatValue,omitempty"`

	// Value of string type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.str_value
	StrValue *string `json:"strValue,omitempty"`

	// Value of date time type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.datetime_value
	DatetimeValue *string `json:"datetimeValue,omitempty"`

	// Value of geo coordinate type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.geo_coordinate
	GeoCoordinate *GeoCoordinate `json:"geoCoordinate,omitempty"`

	// Value of any proto value.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.proto_any_value
	ProtoAnyValue *Any `json:"protoAnyValue,omitempty"`

	// Value of boolean type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Value of customized struct annotation. This field does not have effects.
	//  Use customized_struct_value instead for customized struct annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.customized_struct_data_value
	CustomizedStructDataValue map[string]string `json:"customizedStructDataValue,omitempty"`

	// Value of list type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.list_value
	ListValue *AnnotationList `json:"listValue,omitempty"`

	// Value of custom struct type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.customized_struct_value
	CustomizedStructValue *AnnotationCustomizedStruct `json:"customizedStructValue,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.GeoCoordinate
type GeoCoordinate struct {
	// Latitude Coordinate. Degrees [-90 .. 90]
	// +kcc:proto:field=google.cloud.visionai.v1.GeoCoordinate.latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude Coordinate. Degrees [-180 .. 180]
	// +kcc:proto:field=google.cloud.visionai.v1.GeoCoordinate.longitude
	Longitude *float64 `json:"longitude,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Partition
type Partition struct {
	// Partition of asset in time.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.temporal_partition
	TemporalPartition *Partition_TemporalPartition `json:"temporalPartition,omitempty"`

	// Partition of asset in space.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.spatial_partition
	SpatialPartition *Partition_SpatialPartition `json:"spatialPartition,omitempty"`

	// Partition of asset in time.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.relative_temporal_partition
	RelativeTemporalPartition *Partition_RelativeTemporalPartition `json:"relativeTemporalPartition,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Partition.RelativeTemporalPartition
type Partition_RelativeTemporalPartition struct {
	// Start time offset of the partition.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.RelativeTemporalPartition.start_offset
	StartOffset *string `json:"startOffset,omitempty"`

	// End time offset of the partition.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.RelativeTemporalPartition.end_offset
	EndOffset *string `json:"endOffset,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Partition.SpatialPartition
type Partition_SpatialPartition struct {
	// The minimum x coordinate value.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.SpatialPartition.x_min
	XMin *int64 `json:"xMin,omitempty"`

	// The minimum y coordinate value.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.SpatialPartition.y_min
	YMin *int64 `json:"yMin,omitempty"`

	// The maximum x coordinate value.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.SpatialPartition.x_max
	XMax *int64 `json:"xMax,omitempty"`

	// The maximum y coordinate value.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.SpatialPartition.y_max
	YMax *int64 `json:"yMax,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Partition.TemporalPartition
type Partition_TemporalPartition struct {
	// Start time of the partition.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.TemporalPartition.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End time of the partition.
	// +kcc:proto:field=google.cloud.visionai.v1.Partition.TemporalPartition.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.UserSpecifiedAnnotation
type UserSpecifiedAnnotation struct {
	// Required. Key of the annotation. The key must be set with type by
	//  CreateDataSchema.
	// +kcc:proto:field=google.cloud.visionai.v1.UserSpecifiedAnnotation.key
	Key *string `json:"key,omitempty"`

	// Value of the annotation. The value must be able to convert
	//  to the type according to the data schema.
	// +kcc:proto:field=google.cloud.visionai.v1.UserSpecifiedAnnotation.value
	Value *AnnotationValue `json:"value,omitempty"`

	// Partition information in time and space for the sub-asset level annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.UserSpecifiedAnnotation.partition
	Partition *Partition `json:"partition,omitempty"`
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
