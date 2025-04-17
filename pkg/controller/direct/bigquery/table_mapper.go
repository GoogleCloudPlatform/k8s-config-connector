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

// +generated:mapper
// krm.group: bigquery.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.bigquery.v2

package bigquery

import (
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/api/bigquery/v2"
)

func BigQueryTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmv1beta1.BigQueryTableSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigQueryTableSpec{}
	// MISSING: Kind
	// MISSING: Etag
	// MISSING: ID
	// MISSING: SelfLink
	// MISSING: TableReference
	out.FriendlyName = direct.LazyPtr(in.FriendlyName)
	out.Description = direct.LazyPtr(in.Description)
	out.Labels = in.Labels
	out.Schema = Table_Schema_FromProto(mapCtx, in.Schema)
	out.TimePartitioning = TimePartitioning_FromProto(mapCtx, in.TimePartitioning)
	out.RangePartitioning = RangePartitioning_FromProto(mapCtx, in.RangePartitioning)
	if in.Clustering != nil {
		out.Clustering = in.Clustering.Fields
	}
	out.RequirePartitionFilter = direct.LazyPtr(in.RequirePartitionFilter)
	// MISSING: PartitionDefinition
	// MISSING: NumBytes
	// MISSING: NumPhysicalBytes
	// MISSING: NumLongTermBytes
	// MISSING: NumRows
	// MISSING: CreationTime
	out.ExpirationTime = direct.LazyPtr(in.ExpirationTime)
	// MISSING: LastModifiedTime
	// MISSING: Type
	out.View = ViewDefinition_FromProto(mapCtx, in.View)
	out.MaterializedView = MaterializedViewDefinition_FromProto(mapCtx, in.MaterializedView)
	// MISSING: MaterializedViewStatus
	out.ExternalDataConfiguration = ExternalDataConfiguration_FromProto(mapCtx, in.ExternalDataConfiguration)
	// MISSING: BiglakeConfiguration
	// MISSING: Location
	// MISSING: StreamingBuffer
	out.EncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.EncryptionConfiguration)
	// MISSING: SnapshotDefinition
	// MISSING: DefaultCollation
	// MISSING: DefaultRoundingMode
	// MISSING: CloneDefinition
	// MISSING: NumTimeTravelPhysicalBytes
	// MISSING: NumTotalLogicalBytes
	// MISSING: NumActiveLogicalBytes
	// MISSING: NumLongTermLogicalBytes
	// MISSING: NumCurrentPhysicalBytes
	// MISSING: NumTotalPhysicalBytes
	// MISSING: NumActivePhysicalBytes
	// MISSING: NumLongTermPhysicalBytes
	// MISSING: NumPartitions
	out.MaxStaleness = direct.LazyPtr(in.MaxStaleness)
	// MISSING: Restrictions
	out.TableConstraints = TableConstraints_FromProto(mapCtx, in.TableConstraints)
	// MISSING: ResourceTags
	// MISSING: TableReplicationInfo
	// MISSING: Replicas
	// MISSING: ExternalCatalogTableOptions
	return out
}
func BigQueryTableSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigQueryTableSpec) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	// MISSING: Kind
	// MISSING: Etag
	// MISSING: ID
	// MISSING: SelfLink
	// MISSING: TableReference
	out.FriendlyName = direct.ValueOf(in.FriendlyName)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Schema = Table_Schema_ToProto(mapCtx, in.Schema)
	out.TimePartitioning = TimePartitioning_ToProto(mapCtx, in.TimePartitioning)
	out.RangePartitioning = RangePartitioning_ToProto(mapCtx, in.RangePartitioning)
	if len(in.Clustering) > 0 {
		out.Clustering = &pb.Clustering{}
		out.Clustering.Fields = in.Clustering
	}
	out.RequirePartitionFilter = direct.ValueOf(in.RequirePartitionFilter)
	// MISSING: PartitionDefinition
	// MISSING: NumBytes
	// MISSING: NumPhysicalBytes
	// MISSING: NumLongTermBytes
	// MISSING: NumRows
	// MISSING: CreationTime
	out.ExpirationTime = direct.ValueOf(in.ExpirationTime)
	// MISSING: LastModifiedTime
	// MISSING: Type
	out.View = ViewDefinition_ToProto(mapCtx, in.View)
	out.MaterializedView = MaterializedViewDefinition_ToProto(mapCtx, in.MaterializedView)
	// MISSING: MaterializedViewStatus
	out.ExternalDataConfiguration = ExternalDataConfiguration_ToProto(mapCtx, in.ExternalDataConfiguration)
	// MISSING: BiglakeConfiguration
	// MISSING: Location
	// MISSING: StreamingBuffer
	out.EncryptionConfiguration = TableEncryptionConfiguration_ToProto(mapCtx, in.EncryptionConfiguration)
	// MISSING: SnapshotDefinition
	// MISSING: DefaultCollation
	// MISSING: DefaultRoundingMode
	// MISSING: CloneDefinition
	// MISSING: NumTimeTravelPhysicalBytes
	// MISSING: NumTotalLogicalBytes
	// MISSING: NumActiveLogicalBytes
	// MISSING: NumLongTermLogicalBytes
	// MISSING: NumCurrentPhysicalBytes
	// MISSING: NumTotalPhysicalBytes
	// MISSING: NumActivePhysicalBytes
	// MISSING: NumLongTermPhysicalBytes
	// MISSING: NumPartitions
	out.MaxStaleness = direct.ValueOf(in.MaxStaleness)
	// MISSING: Restrictions
	out.TableConstraints = TableConstraints_ToProto(mapCtx, in.TableConstraints)
	// MISSING: ResourceTags
	// MISSING: TableReplicationInfo
	// MISSING: Replicas
	// MISSING: ExternalCatalogTableOptions
	return out
}
func ColumnReference_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ColumnReference) *pb.ColumnReference {
	if in == nil {
		return nil
	}
	out := &pb.ColumnReference{}
	out.ReferencingColumn = direct.ValueOf(in.ReferencingColumn)
	out.ReferencedColumn = direct.ValueOf(in.ReferencedColumn)
	return out
}

func CsvOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CsvOptions) *pb.CsvOptions {
	if in == nil {
		return nil
	}
	out := &pb.CsvOptions{}
	out.FieldDelimiter = direct.ValueOf(in.FieldDelimiter)
	out.SkipLeadingRows = direct.Int64Value_ToProto(mapCtx, in.SkipLeadingRows)
	out.Quote = direct.StringValue_ToProto(mapCtx, in.Quote)
	out.AllowQuotedNewlines = direct.BoolValue_ToProto(mapCtx, in.AllowQuotedNewlines)
	out.AllowJaggedRows = direct.BoolValue_ToProto(mapCtx, in.AllowJaggedRows)
	out.Encoding = direct.ValueOf(in.Encoding)
	// MISSING: PreserveAsciiControlCharacters
	// MISSING: NullMarker
	return out
}
func DatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.DatasetReference) *krmv1beta1.DatasetReference {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.DatasetReference{}
	// MISSING: DatasetID
	// (near miss): "DatasetID" vs "DatasetId"
	// MISSING: ProjectID
	// (near miss): "ProjectID" vs "ProjectId"
	return out
}
func DatasetReference_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.DatasetReference) *pb.DatasetReference {
	if in == nil {
		return nil
	}
	out := &pb.DatasetReference{}
	// MISSING: DatasetID
	// (near miss): "DatasetID" vs "DatasetId"
	// MISSING: ProjectID
	// (near miss): "ProjectID" vs "ProjectId"
	return out
}
func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfiguration) *krmv1beta1.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.EncryptionConfiguration{}
	if in.KmsKeyName != "" {
		out.KmsKeyRef = &refs.KMSCryptoKeyRef{External: in.KmsKeyName}
	}
	return out
}
func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.EncryptionConfiguration) *pb.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfiguration{}
	// MISSING: KMSKeyName
	return out
}
func ExternalCatalogDatasetOptions_FromProto(mapCtx *direct.MapContext, in *pb.ExternalCatalogDatasetOptions) *krmv1beta1.ExternalCatalogDatasetOptions {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ExternalCatalogDatasetOptions{}
	out.Parameters = in.Parameters
	// MISSING: DefaultStorageLocationURI
	// (near miss): "DefaultStorageLocationURI" vs "DefaultStorageLocationUri"
	return out
}
func ExternalCatalogDatasetOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ExternalCatalogDatasetOptions) *pb.ExternalCatalogDatasetOptions {
	if in == nil {
		return nil
	}
	out := &pb.ExternalCatalogDatasetOptions{}
	out.Parameters = in.Parameters
	// MISSING: DefaultStorageLocationURI
	// (near miss): "DefaultStorageLocationURI" vs "DefaultStorageLocationUri"
	return out
}
func ExternalDatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.ExternalDatasetReference) *krmv1beta1.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ExternalDatasetReference{}
	out.ExternalSource = direct.LazyPtr(in.GetExternalSource())
	out.Connection = direct.LazyPtr(in.GetConnection())
	return out
}
func ExternalDatasetReference_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ExternalDatasetReference) *pb.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &pb.ExternalDatasetReference{}
	out.ExternalSource = direct.ValueOf(in.ExternalSource)
	out.Connection = direct.ValueOf(in.Connection)
	return out
}
func ForeignKey_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ForeignKey) *pb.ForeignKey {
	if in == nil {
		return nil
	}
	out := &pb.ForeignKey{}
	out.Name = direct.ValueOf(in.Name)
	out.ReferencedTable = TableReference_ToProto(mapCtx, in.ReferencedTable)
	out.ColumnReferences = direct.Slice_ToProto(mapCtx, in.ColumnReferences, ColumnReference_ToProto)
	return out
}
func GcpTag_FromProto(mapCtx *direct.MapContext, in *pb.GcpTag) *krmv1beta1.GcpTag {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.GcpTag{}
	out.TagKey = direct.LazyPtr(in.GetTagKey())
	out.TagValue = direct.LazyPtr(in.GetTagValue())
	return out
}
func GcpTag_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.GcpTag) *pb.GcpTag {
	if in == nil {
		return nil
	}
	out := &pb.GcpTag{}
	out.TagKey = direct.ValueOf(in.TagKey)
	out.TagValue = direct.ValueOf(in.TagValue)
	return out
}
func GoogleSheetsOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.GoogleSheetsOptions) *pb.GoogleSheetsOptions {
	if in == nil {
		return nil
	}
	out := &pb.GoogleSheetsOptions{}
	out.SkipLeadingRows = direct.Int64Value_ToProto(mapCtx, in.SkipLeadingRows)
	out.Range = direct.ValueOf(in.Range)
	return out
}
func HivePartitioningOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.HivePartitioningOptions) *pb.HivePartitioningOptions {
	if in == nil {
		return nil
	}
	out := &pb.HivePartitioningOptions{}
	out.Mode = direct.ValueOf(in.Mode)
	// MISSING: SourceURIPrefix
	// (near miss): "SourceURIPrefix" vs "SourceUriPrefix"
	out.RequirePartitionFilter = direct.BoolValue_ToProto(mapCtx, in.RequirePartitionFilter)
	// MISSING: Fields
	return out
}

func JsonOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.JsonOptions) *pb.JsonOptions {
	if in == nil {
		return nil
	}
	out := &pb.JsonOptions{}
	out.Encoding = direct.ValueOf(in.Encoding)
	return out
}
func ParquetOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ParquetOptions) *pb.ParquetOptions {
	if in == nil {
		return nil
	}
	out := &pb.ParquetOptions{}
	out.EnumAsString = direct.BoolValue_ToProto(mapCtx, in.EnumAsString)
	out.EnableListInference = direct.BoolValue_ToProto(mapCtx, in.EnableListInference)
	// MISSING: MapTargetType
	return out
}
func PrimaryKey_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PrimaryKey) *pb.PrimaryKey {
	if in == nil {
		return nil
	}
	out := &pb.PrimaryKey{}
	out.Columns = in.Columns
	return out
}
func RestrictionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RestrictionConfig) *krmv1beta1.RestrictionConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RestrictionConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func RestrictionConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RestrictionConfig) *pb.RestrictionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RestrictionConfig{}
	out.Type = direct.Enum_ToProto[pb.RestrictionConfig_RestrictionType](mapCtx, in.Type)
	return out
}
func RoutineReference_FromProto(mapCtx *direct.MapContext, in *pb.RoutineReference) *krmv1beta1.RoutineReference {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RoutineReference{}
	out.DatasetId = direct.LazyPtr(in.DatasetId)
	out.ProjectId = direct.LazyPtr(in.ProjectId)
	out.RoutineId = direct.LazyPtr(in.RoutineId)
	return out
}
func RoutineReference_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RoutineReference) *pb.RoutineReference {
	if in == nil {
		return nil
	}
	out := &pb.RoutineReference{}
	// MISSING: ProjectID
	// (near miss): "ProjectID" vs "ProjectId"
	// MISSING: DatasetID
	// (near miss): "DatasetID" vs "DatasetId"
	// MISSING: RoutineID
	// (near miss): "RoutineID" vs "RoutineId"
	return out
}

func TableConstraints_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TableConstraints) *pb.TableConstraints {
	if in == nil {
		return nil
	}
	out := &pb.TableConstraints{}
	out.PrimaryKey = PrimaryKey_ToProto(mapCtx, in.PrimaryKey)
	out.ForeignKeys = direct.Slice_ToProto(mapCtx, in.ForeignKeys, ForeignKey_ToProto)
	return out
}
func TableReference_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TableReference) *pb.TableReference {
	if in == nil {
		return nil
	}
	out := &pb.TableReference{}
	// MISSING: ProjectID
	// (near miss): "ProjectID" vs "ProjectId"
	// MISSING: DatasetID
	// (near miss): "DatasetID" vs "DatasetId"
	// MISSING: TableID
	// (near miss): "TableID" vs "TableId"
	return out
}
func ViewDefinition_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ViewDefinition) *pb.ViewDefinition {
	if in == nil {
		return nil
	}
	out := &pb.ViewDefinition{}
	out.Query = direct.ValueOf(in.Query)
	// MISSING: UserDefinedFunctionResources
	// MISSING: UseLegacySQL
	// (near miss): "UseLegacySQL" vs "UseLegacySql"
	// MISSING: UseExplicitColumnNames
	// MISSING: PrivacyPolicy
	// MISSING: ForeignDefinitions
	return out
}
