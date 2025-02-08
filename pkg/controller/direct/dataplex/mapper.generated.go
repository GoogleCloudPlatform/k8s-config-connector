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

package dataplex

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
)
func DataplexEntityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entity) *krm.DataplexEntityObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexEntityObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ID
	// MISSING: Etag
	// MISSING: Type
	// MISSING: Asset
	// MISSING: DataPath
	// MISSING: DataPathPattern
	// MISSING: CatalogEntry
	// MISSING: System
	// MISSING: Format
	// MISSING: Compatibility
	// MISSING: Access
	// MISSING: Uid
	// MISSING: Schema
	return out
}
func DataplexEntityObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexEntityObservedState) *pb.Entity {
	if in == nil {
		return nil
	}
	out := &pb.Entity{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ID
	// MISSING: Etag
	// MISSING: Type
	// MISSING: Asset
	// MISSING: DataPath
	// MISSING: DataPathPattern
	// MISSING: CatalogEntry
	// MISSING: System
	// MISSING: Format
	// MISSING: Compatibility
	// MISSING: Access
	// MISSING: Uid
	// MISSING: Schema
	return out
}
func DataplexEntitySpec_FromProto(mapCtx *direct.MapContext, in *pb.Entity) *krm.DataplexEntitySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexEntitySpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ID
	// MISSING: Etag
	// MISSING: Type
	// MISSING: Asset
	// MISSING: DataPath
	// MISSING: DataPathPattern
	// MISSING: CatalogEntry
	// MISSING: System
	// MISSING: Format
	// MISSING: Compatibility
	// MISSING: Access
	// MISSING: Uid
	// MISSING: Schema
	return out
}
func DataplexEntitySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexEntitySpec) *pb.Entity {
	if in == nil {
		return nil
	}
	out := &pb.Entity{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ID
	// MISSING: Etag
	// MISSING: Type
	// MISSING: Asset
	// MISSING: DataPath
	// MISSING: DataPathPattern
	// MISSING: CatalogEntry
	// MISSING: System
	// MISSING: Format
	// MISSING: Compatibility
	// MISSING: Access
	// MISSING: Uid
	// MISSING: Schema
	return out
}
func Entity_FromProto(mapCtx *direct.MapContext, in *pb.Entity) *krm.Entity {
	if in == nil {
		return nil
	}
	out := &krm.Entity{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ID = direct.LazyPtr(in.GetId())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Asset = direct.LazyPtr(in.GetAsset())
	out.DataPath = direct.LazyPtr(in.GetDataPath())
	out.DataPathPattern = direct.LazyPtr(in.GetDataPathPattern())
	// MISSING: CatalogEntry
	out.System = direct.Enum_FromProto(mapCtx, in.GetSystem())
	out.Format = StorageFormat_FromProto(mapCtx, in.GetFormat())
	// MISSING: Compatibility
	// MISSING: Access
	// MISSING: Uid
	out.Schema = Schema_FromProto(mapCtx, in.GetSchema())
	return out
}
func Entity_ToProto(mapCtx *direct.MapContext, in *krm.Entity) *pb.Entity {
	if in == nil {
		return nil
	}
	out := &pb.Entity{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Id = direct.ValueOf(in.ID)
	out.Etag = direct.ValueOf(in.Etag)
	out.Type = direct.Enum_ToProto[pb.Entity_Type](mapCtx, in.Type)
	out.Asset = direct.ValueOf(in.Asset)
	out.DataPath = direct.ValueOf(in.DataPath)
	out.DataPathPattern = direct.ValueOf(in.DataPathPattern)
	// MISSING: CatalogEntry
	out.System = direct.Enum_ToProto[pb.StorageSystem](mapCtx, in.System)
	out.Format = StorageFormat_ToProto(mapCtx, in.Format)
	// MISSING: Compatibility
	// MISSING: Access
	// MISSING: Uid
	out.Schema = Schema_ToProto(mapCtx, in.Schema)
	return out
}
func EntityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entity) *krm.EntityObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntityObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ID
	// MISSING: Etag
	// MISSING: Type
	// MISSING: Asset
	// MISSING: DataPath
	// MISSING: DataPathPattern
	out.CatalogEntry = direct.LazyPtr(in.GetCatalogEntry())
	// MISSING: System
	out.Format = StorageFormatObservedState_FromProto(mapCtx, in.GetFormat())
	out.Compatibility = Entity_CompatibilityStatus_FromProto(mapCtx, in.GetCompatibility())
	out.Access = StorageAccess_FromProto(mapCtx, in.GetAccess())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Schema
	return out
}
func EntityObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntityObservedState) *pb.Entity {
	if in == nil {
		return nil
	}
	out := &pb.Entity{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ID
	// MISSING: Etag
	// MISSING: Type
	// MISSING: Asset
	// MISSING: DataPath
	// MISSING: DataPathPattern
	out.CatalogEntry = direct.ValueOf(in.CatalogEntry)
	// MISSING: System
	out.Format = StorageFormatObservedState_ToProto(mapCtx, in.Format)
	out.Compatibility = Entity_CompatibilityStatus_ToProto(mapCtx, in.Compatibility)
	out.Access = StorageAccess_ToProto(mapCtx, in.Access)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Schema
	return out
}
func Entity_CompatibilityStatus_FromProto(mapCtx *direct.MapContext, in *pb.Entity_CompatibilityStatus) *krm.Entity_CompatibilityStatus {
	if in == nil {
		return nil
	}
	out := &krm.Entity_CompatibilityStatus{}
	// MISSING: HiveMetastore
	// MISSING: Bigquery
	return out
}
func Entity_CompatibilityStatus_ToProto(mapCtx *direct.MapContext, in *krm.Entity_CompatibilityStatus) *pb.Entity_CompatibilityStatus {
	if in == nil {
		return nil
	}
	out := &pb.Entity_CompatibilityStatus{}
	// MISSING: HiveMetastore
	// MISSING: Bigquery
	return out
}
func Entity_CompatibilityStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entity_CompatibilityStatus) *krm.Entity_CompatibilityStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Entity_CompatibilityStatusObservedState{}
	out.HiveMetastore = Entity_CompatibilityStatus_Compatibility_FromProto(mapCtx, in.GetHiveMetastore())
	out.Bigquery = Entity_CompatibilityStatus_Compatibility_FromProto(mapCtx, in.GetBigquery())
	return out
}
func Entity_CompatibilityStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Entity_CompatibilityStatusObservedState) *pb.Entity_CompatibilityStatus {
	if in == nil {
		return nil
	}
	out := &pb.Entity_CompatibilityStatus{}
	out.HiveMetastore = Entity_CompatibilityStatus_Compatibility_ToProto(mapCtx, in.HiveMetastore)
	out.Bigquery = Entity_CompatibilityStatus_Compatibility_ToProto(mapCtx, in.Bigquery)
	return out
}
func Entity_CompatibilityStatus_Compatibility_FromProto(mapCtx *direct.MapContext, in *pb.Entity_CompatibilityStatus_Compatibility) *krm.Entity_CompatibilityStatus_Compatibility {
	if in == nil {
		return nil
	}
	out := &krm.Entity_CompatibilityStatus_Compatibility{}
	// MISSING: Compatible
	// MISSING: Reason
	return out
}
func Entity_CompatibilityStatus_Compatibility_ToProto(mapCtx *direct.MapContext, in *krm.Entity_CompatibilityStatus_Compatibility) *pb.Entity_CompatibilityStatus_Compatibility {
	if in == nil {
		return nil
	}
	out := &pb.Entity_CompatibilityStatus_Compatibility{}
	// MISSING: Compatible
	// MISSING: Reason
	return out
}
func Entity_CompatibilityStatus_CompatibilityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entity_CompatibilityStatus_Compatibility) *krm.Entity_CompatibilityStatus_CompatibilityObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Entity_CompatibilityStatus_CompatibilityObservedState{}
	out.Compatible = direct.LazyPtr(in.GetCompatible())
	out.Reason = direct.LazyPtr(in.GetReason())
	return out
}
func Entity_CompatibilityStatus_CompatibilityObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Entity_CompatibilityStatus_CompatibilityObservedState) *pb.Entity_CompatibilityStatus_Compatibility {
	if in == nil {
		return nil
	}
	out := &pb.Entity_CompatibilityStatus_Compatibility{}
	out.Compatible = direct.ValueOf(in.Compatible)
	out.Reason = direct.ValueOf(in.Reason)
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.UserManaged = direct.LazyPtr(in.GetUserManaged())
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Schema_SchemaField_FromProto)
	out.PartitionFields = direct.Slice_FromProto(mapCtx, in.PartitionFields, Schema_PartitionField_FromProto)
	out.PartitionStyle = direct.Enum_FromProto(mapCtx, in.GetPartitionStyle())
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	out.UserManaged = direct.ValueOf(in.UserManaged)
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Schema_SchemaField_ToProto)
	out.PartitionFields = direct.Slice_ToProto(mapCtx, in.PartitionFields, Schema_PartitionField_ToProto)
	out.PartitionStyle = direct.Enum_ToProto[pb.Schema_PartitionStyle](mapCtx, in.PartitionStyle)
	return out
}
func Schema_PartitionField_FromProto(mapCtx *direct.MapContext, in *pb.Schema_PartitionField) *krm.Schema_PartitionField {
	if in == nil {
		return nil
	}
	out := &krm.Schema_PartitionField{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Schema_PartitionField_ToProto(mapCtx *direct.MapContext, in *krm.Schema_PartitionField) *pb.Schema_PartitionField {
	if in == nil {
		return nil
	}
	out := &pb.Schema_PartitionField{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.Enum_ToProto[pb.Schema_Type](mapCtx, in.Type)
	return out
}
func Schema_SchemaField_FromProto(mapCtx *direct.MapContext, in *pb.Schema_SchemaField) *krm.Schema_SchemaField {
	if in == nil {
		return nil
	}
	out := &krm.Schema_SchemaField{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Schema_SchemaField_FromProto)
	return out
}
func Schema_SchemaField_ToProto(mapCtx *direct.MapContext, in *krm.Schema_SchemaField) *pb.Schema_SchemaField {
	if in == nil {
		return nil
	}
	out := &pb.Schema_SchemaField{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Type = direct.Enum_ToProto[pb.Schema_Type](mapCtx, in.Type)
	out.Mode = direct.Enum_ToProto[pb.Schema_Mode](mapCtx, in.Mode)
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Schema_SchemaField_ToProto)
	return out
}
func StorageAccess_FromProto(mapCtx *direct.MapContext, in *pb.StorageAccess) *krm.StorageAccess {
	if in == nil {
		return nil
	}
	out := &krm.StorageAccess{}
	// MISSING: Read
	return out
}
func StorageAccess_ToProto(mapCtx *direct.MapContext, in *krm.StorageAccess) *pb.StorageAccess {
	if in == nil {
		return nil
	}
	out := &pb.StorageAccess{}
	// MISSING: Read
	return out
}
func StorageAccessObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StorageAccess) *krm.StorageAccessObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageAccessObservedState{}
	out.Read = direct.Enum_FromProto(mapCtx, in.GetRead())
	return out
}
func StorageAccessObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageAccessObservedState) *pb.StorageAccess {
	if in == nil {
		return nil
	}
	out := &pb.StorageAccess{}
	out.Read = direct.Enum_ToProto[pb.StorageAccess_AccessMode](mapCtx, in.Read)
	return out
}
func StorageFormat_FromProto(mapCtx *direct.MapContext, in *pb.StorageFormat) *krm.StorageFormat {
	if in == nil {
		return nil
	}
	out := &krm.StorageFormat{}
	// MISSING: Format
	out.CompressionFormat = direct.Enum_FromProto(mapCtx, in.GetCompressionFormat())
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.Csv = StorageFormat_CsvOptions_FromProto(mapCtx, in.GetCsv())
	out.Json = StorageFormat_JsonOptions_FromProto(mapCtx, in.GetJson())
	out.Iceberg = StorageFormat_IcebergOptions_FromProto(mapCtx, in.GetIceberg())
	return out
}
func StorageFormat_ToProto(mapCtx *direct.MapContext, in *krm.StorageFormat) *pb.StorageFormat {
	if in == nil {
		return nil
	}
	out := &pb.StorageFormat{}
	// MISSING: Format
	out.CompressionFormat = direct.Enum_ToProto[pb.StorageFormat_CompressionFormat](mapCtx, in.CompressionFormat)
	out.MimeType = direct.ValueOf(in.MimeType)
	if oneof := StorageFormat_CsvOptions_ToProto(mapCtx, in.Csv); oneof != nil {
		out.Options = &pb.StorageFormat_Csv{Csv: oneof}
	}
	if oneof := StorageFormat_JsonOptions_ToProto(mapCtx, in.Json); oneof != nil {
		out.Options = &pb.StorageFormat_Json{Json: oneof}
	}
	if oneof := StorageFormat_IcebergOptions_ToProto(mapCtx, in.Iceberg); oneof != nil {
		out.Options = &pb.StorageFormat_Iceberg{Iceberg: oneof}
	}
	return out
}
func StorageFormatObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StorageFormat) *krm.StorageFormatObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageFormatObservedState{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	// MISSING: CompressionFormat
	// MISSING: MimeType
	// MISSING: Csv
	// MISSING: Json
	// MISSING: Iceberg
	return out
}
func StorageFormatObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageFormatObservedState) *pb.StorageFormat {
	if in == nil {
		return nil
	}
	out := &pb.StorageFormat{}
	out.Format = direct.Enum_ToProto[pb.StorageFormat_Format](mapCtx, in.Format)
	// MISSING: CompressionFormat
	// MISSING: MimeType
	// MISSING: Csv
	// MISSING: Json
	// MISSING: Iceberg
	return out
}
func StorageFormat_CsvOptions_FromProto(mapCtx *direct.MapContext, in *pb.StorageFormat_CsvOptions) *krm.StorageFormat_CsvOptions {
	if in == nil {
		return nil
	}
	out := &krm.StorageFormat_CsvOptions{}
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.HeaderRows = direct.LazyPtr(in.GetHeaderRows())
	out.Delimiter = direct.LazyPtr(in.GetDelimiter())
	out.Quote = direct.LazyPtr(in.GetQuote())
	return out
}
func StorageFormat_CsvOptions_ToProto(mapCtx *direct.MapContext, in *krm.StorageFormat_CsvOptions) *pb.StorageFormat_CsvOptions {
	if in == nil {
		return nil
	}
	out := &pb.StorageFormat_CsvOptions{}
	out.Encoding = direct.ValueOf(in.Encoding)
	out.HeaderRows = direct.ValueOf(in.HeaderRows)
	out.Delimiter = direct.ValueOf(in.Delimiter)
	out.Quote = direct.ValueOf(in.Quote)
	return out
}
func StorageFormat_IcebergOptions_FromProto(mapCtx *direct.MapContext, in *pb.StorageFormat_IcebergOptions) *krm.StorageFormat_IcebergOptions {
	if in == nil {
		return nil
	}
	out := &krm.StorageFormat_IcebergOptions{}
	out.MetadataLocation = direct.LazyPtr(in.GetMetadataLocation())
	return out
}
func StorageFormat_IcebergOptions_ToProto(mapCtx *direct.MapContext, in *krm.StorageFormat_IcebergOptions) *pb.StorageFormat_IcebergOptions {
	if in == nil {
		return nil
	}
	out := &pb.StorageFormat_IcebergOptions{}
	out.MetadataLocation = direct.ValueOf(in.MetadataLocation)
	return out
}
func StorageFormat_JsonOptions_FromProto(mapCtx *direct.MapContext, in *pb.StorageFormat_JsonOptions) *krm.StorageFormat_JsonOptions {
	if in == nil {
		return nil
	}
	out := &krm.StorageFormat_JsonOptions{}
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	return out
}
func StorageFormat_JsonOptions_ToProto(mapCtx *direct.MapContext, in *krm.StorageFormat_JsonOptions) *pb.StorageFormat_JsonOptions {
	if in == nil {
		return nil
	}
	out := &pb.StorageFormat_JsonOptions{}
	out.Encoding = direct.ValueOf(in.Encoding)
	return out
}
