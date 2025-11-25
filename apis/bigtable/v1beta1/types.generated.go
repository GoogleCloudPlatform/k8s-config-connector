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
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.bigtable.admin.v2
// resource: BigtableAppProfile:AppProfile
// resource: BigtableTable:Table

package v1beta1

// +kcc:proto=google.bigtable.admin.v2.AppProfile.DataBoostIsolationReadOnly
type AppProfile_DataBoostIsolationReadOnly struct {
	// The Compute Billing Owner for this Data Boost App Profile.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.DataBoostIsolationReadOnly.compute_billing_owner
	ComputeBillingOwner *string `json:"computeBillingOwner,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny
type AppProfile_MultiClusterRoutingUseAny struct {
	// The set of clusters to route to. The order is ignored; clusters will be
	//  tried in order of distance. If left empty, all clusters are eligible.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny.cluster_ids
	ClusterIds []string `json:"clusterIds,omitempty"`

	// Row affinity sticky routing based on the row key of the request.
	//  Requests that span multiple rows are routed non-deterministically.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny.row_affinity
	RowAffinity *AppProfile_MultiClusterRoutingUseAny_RowAffinity `json:"rowAffinity,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny.RowAffinity
type AppProfile_MultiClusterRoutingUseAny_RowAffinity struct {
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.StandardIsolation
type AppProfile_StandardIsolation struct {
	// The priority of requests sent using this app profile.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.StandardIsolation.priority
	Priority *string `json:"priority,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.BackupInfo
type BackupInfo struct {
}

// +kcc:proto=google.bigtable.admin.v2.ChangeStreamConfig
type ChangeStreamConfig struct {
	// How long the change stream should be retained. Change stream data older
	//  than the retention period will not be returned when reading the change
	//  stream from the table.
	//  Values must be at least 1 day and at most 7 days, and will be truncated to
	//  microsecond granularity.
	// +kcc:proto:field=google.bigtable.admin.v2.ChangeStreamConfig.retention_period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.GcRule
type GcRule struct {
	// Delete all cells in a column except the most recent N.
	// +kcc:proto:field=google.bigtable.admin.v2.GcRule.max_num_versions
	MaxNumVersions *int32 `json:"maxNumVersions,omitempty"`

	// Delete cells in a column older than the given age.
	//  Values must be at least one millisecond, and will be truncated to
	//  microsecond granularity.
	// +kcc:proto:field=google.bigtable.admin.v2.GcRule.max_age
	MaxAge *string `json:"maxAge,omitempty"`

	// Delete cells that would be deleted by every nested rule.
	// +kcc:proto:field=google.bigtable.admin.v2.GcRule.intersection
	Intersection *GcRule_Intersection `json:"intersection,omitempty"`

	// Delete cells that would be deleted by any nested rule.
	// +kcc:proto:field=google.bigtable.admin.v2.GcRule.union
	Union *GcRule_Union `json:"union,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.GcRule.Intersection
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	// +kcc:proto:field=google.bigtable.admin.v2.GcRule.Intersection.rules
	Rules []GcRule `json:"rules,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.GcRule.Union
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	// +kcc:proto:field=google.bigtable.admin.v2.GcRule.Union.rules
	Rules []GcRule `json:"rules,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.RestoreInfo
type RestoreInfo struct {
	// The type of the restore source.
	// +kcc:proto:field=google.bigtable.admin.v2.RestoreInfo.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// Information about the backup used to restore the table. The backup
	//  may no longer exist.
	// +kcc:proto:field=google.bigtable.admin.v2.RestoreInfo.backup_info
	BackupInfo *BackupInfo `json:"backupInfo,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Table.AutomatedBackupPolicy
type Table_AutomatedBackupPolicy struct {
	// Required. How long the automated backups should be retained. The only
	//  supported value at this time is 3 days.
	// +kcc:proto:field=google.bigtable.admin.v2.Table.AutomatedBackupPolicy.retention_period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`

	// Required. How frequently automated backups should occur. The only
	//  supported value at this time is 24 hours.
	// +kcc:proto:field=google.bigtable.admin.v2.Table.AutomatedBackupPolicy.frequency
	Frequency *string `json:"frequency,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Table.ClusterState
type Table_ClusterState struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type
type Type struct {
	// Bytes
	// +kcc:proto:field=google.bigtable.admin.v2.Type.bytes_type
	BytesType *Type_Bytes `json:"bytesType,omitempty"`

	// String
	// +kcc:proto:field=google.bigtable.admin.v2.Type.string_type
	StringType *Type_String `json:"stringType,omitempty"`

	// Int64
	// +kcc:proto:field=google.bigtable.admin.v2.Type.int64_type
	Int64Type *Type_Int64 `json:"int64Type,omitempty"`

	// Float32
	// +kcc:proto:field=google.bigtable.admin.v2.Type.float32_type
	Float32Type *Type_Float32 `json:"float32Type,omitempty"`

	// Float64
	// +kcc:proto:field=google.bigtable.admin.v2.Type.float64_type
	Float64Type *Type_Float64 `json:"float64Type,omitempty"`

	// Bool
	// +kcc:proto:field=google.bigtable.admin.v2.Type.bool_type
	BoolType *Type_Bool `json:"boolType,omitempty"`

	// Timestamp
	// +kcc:proto:field=google.bigtable.admin.v2.Type.timestamp_type
	TimestampType *Type_Timestamp `json:"timestampType,omitempty"`

	// Date
	// +kcc:proto:field=google.bigtable.admin.v2.Type.date_type
	DateType *Type_Date `json:"dateType,omitempty"`

	// Aggregate
	// +kcc:proto:field=google.bigtable.admin.v2.Type.aggregate_type
	AggregateType *Type_Aggregate `json:"aggregateType,omitempty"`

	// Struct
	// +kcc:proto:field=google.bigtable.admin.v2.Type.struct_type
	StructType *Type_Struct `json:"structType,omitempty"`

	// Array
	// +kcc:proto:field=google.bigtable.admin.v2.Type.array_type
	ArrayType *Type_Array `json:"arrayType,omitempty"`

	// Map
	// +kcc:proto:field=google.bigtable.admin.v2.Type.map_type
	MapType *Type_Map `json:"mapType,omitempty"`

	// Proto
	// +kcc:proto:field=google.bigtable.admin.v2.Type.proto_type
	ProtoType *Type_Proto `json:"protoType,omitempty"`

	// Enum
	// +kcc:proto:field=google.bigtable.admin.v2.Type.enum_type
	EnumType *Type_Enum `json:"enumType,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Aggregate
type Type_Aggregate struct {
	// Type of the inputs that are accumulated by this `Aggregate`, which must
	//  specify a full encoding.
	//  Use `AddInput` mutations to accumulate new inputs.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Aggregate.input_type
	InputType *Type `json:"inputType,omitempty"`

	// Sum aggregator.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Aggregate.sum
	Sum *Type_Aggregate_Sum `json:"sum,omitempty"`

	// HyperLogLogPlusPlusUniqueCount aggregator.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Aggregate.hllpp_unique_count
	HllppUniqueCount *Type_Aggregate_HyperLogLogPlusPlusUniqueCount `json:"hllppUniqueCount,omitempty"`

	// Max aggregator.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Aggregate.max
	Max *Type_Aggregate_Max `json:"max,omitempty"`

	// Min aggregator.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Aggregate.min
	Min *Type_Aggregate_Min `json:"min,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Aggregate.HyperLogLogPlusPlusUniqueCount
type Type_Aggregate_HyperLogLogPlusPlusUniqueCount struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Aggregate.Max
type Type_Aggregate_Max struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Aggregate.Min
type Type_Aggregate_Min struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Aggregate.Sum
type Type_Aggregate_Sum struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Array
type Type_Array struct {
	// The type of the elements in the array. This must not be `Array`.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Array.element_type
	ElementType *Type `json:"elementType,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Bool
type Type_Bool struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Bytes
type Type_Bytes struct {
	// The encoding to use when converting to or from lower level types.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Bytes.encoding
	Encoding *Type_Bytes_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Bytes.Encoding
type Type_Bytes_Encoding struct {
	// Use `Raw` encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Bytes.Encoding.raw
	Raw *Type_Bytes_Encoding_Raw `json:"raw,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Bytes.Encoding.Raw
type Type_Bytes_Encoding_Raw struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Date
type Type_Date struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Enum
type Type_Enum struct {
	// The ID of the schema bundle that this enum is defined in.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Enum.schema_bundle_id
	SchemaBundleID *string `json:"schemaBundleID,omitempty"`

	// The fully qualified name of the protobuf enum message, including package.
	//  In the format of "foo.bar.EnumMessage".
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Enum.enum_name
	EnumName *string `json:"enumName,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Float32
type Type_Float32 struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Float64
type Type_Float64 struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Int64
type Type_Int64 struct {
	// The encoding to use when converting to or from lower level types.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Int64.encoding
	Encoding *Type_Int64_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Int64.Encoding
type Type_Int64_Encoding struct {
	// Use `BigEndianBytes` encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Int64.Encoding.big_endian_bytes
	BigEndianBytes *Type_Int64_Encoding_BigEndianBytes `json:"bigEndianBytes,omitempty"`

	// Use `OrderedCodeBytes` encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Int64.Encoding.ordered_code_bytes
	OrderedCodeBytes *Type_Int64_Encoding_OrderedCodeBytes `json:"orderedCodeBytes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Int64.Encoding.BigEndianBytes
type Type_Int64_Encoding_BigEndianBytes struct {
	// Deprecated: ignored if set.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Int64.Encoding.BigEndianBytes.bytes_type
	BytesType *Type_Bytes `json:"bytesType,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Int64.Encoding.OrderedCodeBytes
type Type_Int64_Encoding_OrderedCodeBytes struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Map
type Type_Map struct {
	// The type of a map key.
	//  Only `Bytes`, `String`, and `Int64` are allowed as key types.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Map.key_type
	KeyType *Type `json:"keyType,omitempty"`

	// The type of the values in a map.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Map.value_type
	ValueType *Type `json:"valueType,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Proto
type Type_Proto struct {
	// The ID of the schema bundle that this proto is defined in.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Proto.schema_bundle_id
	SchemaBundleID *string `json:"schemaBundleID,omitempty"`

	// The fully qualified name of the protobuf message, including package. In
	//  the format of "foo.bar.Message".
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Proto.message_name
	MessageName *string `json:"messageName,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.String
type Type_String struct {
	// The encoding to use when converting to or from lower level types.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.String.encoding
	Encoding *Type_String_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.String.Encoding
type Type_String_Encoding struct {
	// Deprecated: if set, converts to an empty `utf8_bytes`.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.String.Encoding.utf8_raw
	Utf8Raw *Type_String_Encoding_Utf8Raw `json:"utf8Raw,omitempty"`

	// Use `Utf8Bytes` encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.String.Encoding.utf8_bytes
	Utf8Bytes *Type_String_Encoding_Utf8Bytes `json:"utf8Bytes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.String.Encoding.Utf8Bytes
type Type_String_Encoding_Utf8Bytes struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.String.Encoding.Utf8Raw
type Type_String_Encoding_Utf8Raw struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Struct
type Type_Struct struct {
	// The names and types of the fields in this struct.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.fields
	Fields []Type_Struct_Field `json:"fields,omitempty"`

	// The encoding to use when converting to or from lower level types.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.encoding
	Encoding *Type_Struct_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Struct.Encoding
type Type_Struct_Encoding struct {
	// Use `Singleton` encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.Encoding.singleton
	Singleton *Type_Struct_Encoding_Singleton `json:"singleton,omitempty"`

	// Use `DelimitedBytes` encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.Encoding.delimited_bytes
	DelimitedBytes *Type_Struct_Encoding_DelimitedBytes `json:"delimitedBytes,omitempty"`

	// User `OrderedCodeBytes` encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.Encoding.ordered_code_bytes
	OrderedCodeBytes *Type_Struct_Encoding_OrderedCodeBytes `json:"orderedCodeBytes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Struct.Encoding.DelimitedBytes
type Type_Struct_Encoding_DelimitedBytes struct {
	// Byte sequence used to delimit concatenated fields. The delimiter must
	//  contain at least 1 character and at most 50 characters.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.Encoding.DelimitedBytes.delimiter
	Delimiter []byte `json:"delimiter,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Struct.Encoding.OrderedCodeBytes
type Type_Struct_Encoding_OrderedCodeBytes struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Struct.Encoding.Singleton
type Type_Struct_Encoding_Singleton struct {
}

// +kcc:proto=google.bigtable.admin.v2.Type.Struct.Field
type Type_Struct_Field struct {
	// The field name (optional). Fields without a `field_name` are considered
	//  anonymous and cannot be referenced by name.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.Field.field_name
	FieldName *string `json:"fieldName,omitempty"`

	// The type of values in this field.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.Field.type
	Type *Type `json:"type,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Timestamp
type Type_Timestamp struct {
	// The encoding to use when converting to or from lower level types.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Timestamp.encoding
	Encoding *Type_Timestamp_Encoding `json:"encoding,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Type.Timestamp.Encoding
type Type_Timestamp_Encoding struct {
	// Encodes the number of microseconds since the Unix epoch using the
	//  given `Int64` encoding. Values must be microsecond-aligned.
	//
	//  Compatible with:
	//
	//   - Java `Instant.truncatedTo()` with `ChronoUnit.MICROS`
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Timestamp.Encoding.unix_micros_int64
	UnixMicrosInt64 *Type_Int64_Encoding `json:"unixMicrosInt64,omitempty"`
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

// +kcc:observedstate:proto=google.bigtable.admin.v2.BackupInfo
type BackupInfoObservedState struct {
	// Output only. Name of the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.BackupInfo.backup
	Backup *string `json:"backup,omitempty"`

	// Output only. The time that the backup was started. Row data in the backup
	//  will be no older than this timestamp.
	// +kcc:proto:field=google.bigtable.admin.v2.BackupInfo.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. This time that the backup was finished. Row data in the
	//  backup will be no newer than this timestamp.
	// +kcc:proto:field=google.bigtable.admin.v2.BackupInfo.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Name of the table the backup was created from.
	// +kcc:proto:field=google.bigtable.admin.v2.BackupInfo.source_table
	SourceTable *string `json:"sourceTable,omitempty"`

	// Output only. Name of the backup from which this backup was copied. If a
	//  backup is not created by copying a backup, this field will be empty. Values
	//  are of the form:
	//  projects/<project>/instances/<instance>/clusters/<cluster>/backups/<backup>
	// +kcc:proto:field=google.bigtable.admin.v2.BackupInfo.source_backup
	SourceBackup *string `json:"sourceBackup,omitempty"`
}

// +kcc:observedstate:proto=google.bigtable.admin.v2.RestoreInfo
type RestoreInfoObservedState struct {
	// Information about the backup used to restore the table. The backup
	//  may no longer exist.
	// +kcc:proto:field=google.bigtable.admin.v2.RestoreInfo.backup_info
	BackupInfo *BackupInfoObservedState `json:"backupInfo,omitempty"`
}

// +kcc:observedstate:proto=google.bigtable.admin.v2.Type
type TypeObservedState struct {
	// Aggregate
	// +kcc:proto:field=google.bigtable.admin.v2.Type.aggregate_type
	AggregateType *Type_AggregateObservedState `json:"aggregateType,omitempty"`
}

// +kcc:observedstate:proto=google.bigtable.admin.v2.Type.Aggregate
type Type_AggregateObservedState struct {
	// Output only. Type that holds the internal accumulator state for the
	//  `Aggregate`. This is a function of the `input_type` and `aggregator`
	//  chosen, and will always specify a full encoding.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Aggregate.state_type
	StateType *Type `json:"stateType,omitempty"`
}

// +kcc:observedstate:proto=google.bigtable.admin.v2.Type.Struct
type Type_StructObservedState struct {
	// The names and types of the fields in this struct.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.fields
	Fields []Type_Struct_FieldObservedState `json:"fields,omitempty"`
}

// +kcc:observedstate:proto=google.bigtable.admin.v2.Type.Struct.Field
type Type_Struct_FieldObservedState struct {
	// The type of values in this field.
	// +kcc:proto:field=google.bigtable.admin.v2.Type.Struct.Field.type
	Type *TypeObservedState `json:"type,omitempty"`
}
