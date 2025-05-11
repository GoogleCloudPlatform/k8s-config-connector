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

package bigquery

import (
	"encoding/json"

	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/api/bigquery/v2"
)

func MaterializedViewDefinition_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedViewDefinition) *krmv1beta1.MaterializedViewDefinition {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MaterializedViewDefinition{}
	out.Query = direct.LazyPtr(in.Query)
	// MISSING: LastRefreshTime
	out.EnableRefresh = direct.LazyPtr(in.EnableRefresh)
	out.RefreshIntervalMs = direct.LazyPtr(in.RefreshIntervalMs)
	out.AllowNonIncrementalDefinition = direct.LazyPtr(in.AllowNonIncrementalDefinition)
	return out
}

func MaterializedViewDefinition_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MaterializedViewDefinition) *pb.MaterializedViewDefinition {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedViewDefinition{}
	out.Query = direct.ValueOf(in.Query)
	// MISSING: LastRefreshTime
	out.EnableRefresh = direct.ValueOf(in.EnableRefresh)
	out.RefreshIntervalMs = direct.ValueOf(in.RefreshIntervalMs)
	out.AllowNonIncrementalDefinition = direct.ValueOf(in.AllowNonIncrementalDefinition)
	return out
}

func RangePartitioning_Range_FromProto(mapCtx *direct.MapContext, in *pb.RangePartitioningRange) krmv1beta1.RangePartitioning_Range {
	if in == nil {
		return krmv1beta1.RangePartitioning_Range{}
	}
	out := krmv1beta1.RangePartitioning_Range{}
	out.Start = direct.LazyPtr(in.Start)
	out.End = direct.LazyPtr(in.End)
	out.Interval = direct.LazyPtr(in.Interval)
	return out
}

func RangePartitioning_Range_ToProto(mapCtx *direct.MapContext, in krmv1beta1.RangePartitioning_Range) *pb.RangePartitioningRange {
	out := &pb.RangePartitioningRange{}
	out.Start = direct.ValueOf(in.Start)
	out.End = direct.ValueOf(in.End)
	out.Interval = direct.ValueOf(in.Interval)
	return out
}

func RangePartitioning_FromProto(mapCtx *direct.MapContext, in *pb.RangePartitioning) *krmv1beta1.RangePartitioning {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RangePartitioning{}
	out.Field = direct.LazyPtr(in.Field)
	out.Range = RangePartitioning_Range_FromProto(mapCtx, in.Range)
	return out
}
func RangePartitioning_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RangePartitioning) *pb.RangePartitioning {
	if in == nil {
		return nil
	}
	out := &pb.RangePartitioning{}
	out.Field = direct.ValueOf(in.Field)
	out.Range = RangePartitioning_Range_ToProto(mapCtx, in.Range)
	return out
}

func TimePartitioning_FromProto(mapCtx *direct.MapContext, in *pb.TimePartitioning) *krmv1beta1.TimePartitioning {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TimePartitioning{}
	out.Type = in.Type
	out.ExpirationMs = direct.LazyPtr(in.ExpirationMs)
	out.Field = direct.LazyPtr(in.Field)
	return out
}

func TimePartitioning_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TimePartitioning) *pb.TimePartitioning {
	if in == nil {
		return nil
	}
	out := &pb.TimePartitioning{}
	out.Type = in.Type
	out.ExpirationMs = direct.ValueOf(in.ExpirationMs)
	out.Field = direct.ValueOf(in.Field)
	return out
}

func Table_Schema_FromProto(mapCtx *direct.MapContext, in *pb.TableSchema) *string {
	if in == nil {
		return nil
	}
	jsonDataBytes, err := json.Marshal(in)
	if err != nil {
		mapCtx.Errorf("failed to marshal json: %v", err)
		return nil
	}
	return direct.LazyPtr(string(jsonDataBytes))
}

func Table_Schema_ToProto(mapCtx *direct.MapContext, in *string) *pb.TableSchema {
	if in == nil {
		return nil
	}
	out := &pb.TableSchema{}
	err := json.Unmarshal([]byte(direct.ValueOf(in)), out)
	if err != nil {
		mapCtx.Errorf("failed to unmarshal json: %v", err)
		return nil
	}
	return out
}

func ViewDefinition_FromProto(mapCtx *direct.MapContext, in *pb.ViewDefinition) *krmv1beta1.ViewDefinition {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ViewDefinition{}
	out.Query = direct.LazyPtr(in.Query)
	// MISSING: UserDefinedFunctionResources
	// MISSING: UseLegacySQL
	// (near miss): "UseLegacySQL" vs "UseLegacySql"
	// MISSING: UseExplicitColumnNames
	// MISSING: PrivacyPolicy
	// MISSING: ForeignDefinitions
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

func BigQueryTableObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmv1beta1.BigQueryTableObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigQueryTableObservedState{}
	// MISSING: Kind
	// MISSING: Etag
	// MISSING: ID
	// MISSING: SelfLink
	// MISSING: TableReference
	// MISSING: PartitionDefinition
	// MISSING: NumBytes
	// MISSING: NumPhysicalBytes
	// MISSING: NumLongTermBytes
	// MISSING: NumRows
	// MISSING: CreationTime
	// MISSING: LastModifiedTime
	// MISSING: Type
	// MISSING: MaterializedViewStatus
	// MISSING: BiglakeConfiguration
	// MISSING: Location
	// MISSING: StreamingBuffer
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
	// MISSING: Restrictions
	// MISSING: ResourceTags
	// MISSING: TableReplicationInfo
	// MISSING: Replicas
	// MISSING: ExternalCatalogTableOptions
	return out
}

func BigQueryTableObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigQueryTableObservedState) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	// MISSING: Kind
	// MISSING: Etag
	// MISSING: ID
	// MISSING: SelfLink
	// MISSING: TableReference
	// MISSING: PartitionDefinition
	// MISSING: NumBytes
	// MISSING: NumPhysicalBytes
	// MISSING: NumLongTermBytes
	// MISSING: NumRows
	// MISSING: CreationTime
	// MISSING: LastModifiedTime
	// MISSING: Type
	// MISSING: MaterializedViewStatus
	// MISSING: BiglakeConfiguration
	// MISSING: Location
	// MISSING: StreamingBuffer
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
	// MISSING: Restrictions
	// MISSING: ResourceTags
	// MISSING: TableReplicationInfo
	// MISSING: Replicas
	// MISSING: ExternalCatalogTableOptions
	return out
}

func ExternalDataConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.ExternalDataConfiguration) *krmv1beta1.ExternalDataConfiguration {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ExternalDataConfiguration{}
	out.SourceUris = in.SourceUris
	out.FileSetSpecType = direct.LazyPtr(in.FileSetSpecType)
	out.Schema = Table_Schema_FromProto(mapCtx, in.Schema)
	out.SourceFormat = direct.LazyPtr(in.SourceFormat)
	out.MaxBadRecords = direct.LazyPtr(in.MaxBadRecords)
	out.Autodetect = direct.LazyPtr(in.Autodetect)
	out.IgnoreUnknownValues = direct.LazyPtr(in.IgnoreUnknownValues)
	out.Compression = direct.LazyPtr(in.Compression)
	out.CsvOptions = CsvOptions_FromProto(mapCtx, in.CsvOptions)
	out.JsonOptions = JsonOptions_FromProto(mapCtx, in.JsonOptions)
	// MISSING: BigtableOptions
	out.GoogleSheetsOptions = GoogleSheetsOptions_FromProto(mapCtx, in.GoogleSheetsOptions)
	out.HivePartitioningOptions = HivePartitioningOptions_FromProto(mapCtx, in.HivePartitioningOptions)
	// MISSING: ConnectionID
	// (near miss): "ConnectionID" vs "ConnectionId"
	// MISSING: DecimalTargetTypes
	out.AvroOptions = AvroOptions_FromProto(mapCtx, in.AvroOptions)
	// MISSING: JsonExtension
	out.ParquetOptions = ParquetOptions_FromProto(mapCtx, in.ParquetOptions)
	out.ObjectMetadata = direct.LazyPtr(in.ObjectMetadata)
	// MISSING: ReferenceFileSchemaURI
	// (near miss): "ReferenceFileSchemaURI" vs "ReferenceFileSchemaUri"
	out.MetadataCacheMode = direct.LazyPtr(in.MetadataCacheMode)
	return out
}

func ExternalDataConfiguration_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ExternalDataConfiguration) *pb.ExternalDataConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.ExternalDataConfiguration{}
	out.SourceUris = in.SourceUris
	out.FileSetSpecType = direct.ValueOf(in.FileSetSpecType)
	out.Schema = Table_Schema_ToProto(mapCtx, in.Schema)
	out.SourceFormat = direct.ValueOf(in.SourceFormat)
	out.MaxBadRecords = direct.ValueOf(in.MaxBadRecords)
	out.Autodetect = direct.ValueOf(in.Autodetect)
	out.IgnoreUnknownValues = direct.ValueOf(in.IgnoreUnknownValues)
	out.Compression = direct.ValueOf(in.Compression)
	out.CsvOptions = CsvOptions_ToProto(mapCtx, in.CsvOptions)
	out.JsonOptions = JsonOptions_ToProto(mapCtx, in.JsonOptions)
	// MISSING: BigtableOptions
	out.GoogleSheetsOptions = GoogleSheetsOptions_ToProto(mapCtx, in.GoogleSheetsOptions)
	out.HivePartitioningOptions = HivePartitioningOptions_ToProto(mapCtx, in.HivePartitioningOptions)
	// MISSING: ConnectionID
	// (near miss): "ConnectionID" vs "ConnectionId"
	// MISSING: DecimalTargetTypes
	out.AvroOptions = AvroOptions_ToProto(mapCtx, in.AvroOptions)
	// MISSING: JsonExtension
	out.ParquetOptions = ParquetOptions_ToProto(mapCtx, in.ParquetOptions)
	out.ObjectMetadata = direct.ValueOf(in.ObjectMetadata)
	// MISSING: ReferenceFileSchemaURI
	// (near miss): "ReferenceFileSchemaURI" vs "ReferenceFileSchemaUri"
	out.MetadataCacheMode = direct.ValueOf(in.MetadataCacheMode)
	return out
}

func CsvOptions_FromProto(mapCtx *direct.MapContext, in *pb.CsvOptions) *krmv1beta1.CsvOptions {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CsvOptions{}
	out.FieldDelimiter = direct.LazyPtr(in.FieldDelimiter)
	out.SkipLeadingRows = direct.LazyPtr(in.SkipLeadingRows)
	out.Quote = in.Quote
	out.AllowQuotedNewlines = direct.LazyPtr(in.AllowQuotedNewlines)
	out.AllowJaggedRows = direct.LazyPtr(in.AllowJaggedRows)
	out.Encoding = direct.LazyPtr(in.Encoding)
	// MISSING: PreserveAsciiControlCharacters
	// MISSING: NullMarker
	return out
}

func CsvOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CsvOptions) *pb.CsvOptions {
	if in == nil {
		return nil
	}
	out := &pb.CsvOptions{}
	out.FieldDelimiter = direct.ValueOf(in.FieldDelimiter)
	out.SkipLeadingRows = direct.ValueOf(in.SkipLeadingRows)
	out.Quote = in.Quote
	out.AllowQuotedNewlines = direct.ValueOf(in.AllowQuotedNewlines)
	out.AllowJaggedRows = direct.ValueOf(in.AllowJaggedRows)
	out.Encoding = direct.ValueOf(in.Encoding)
	// MISSING: PreserveAsciiControlCharacters
	// MISSING: NullMarker
	return out
}

func JsonOptions_FromProto(mapCtx *direct.MapContext, in *pb.JsonOptions) *krmv1beta1.JsonOptions {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.JsonOptions{}
	out.Encoding = direct.LazyPtr(in.Encoding)
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

func HivePartitioningOptions_FromProto(mapCtx *direct.MapContext, in *pb.HivePartitioningOptions) *krmv1beta1.HivePartitioningOptions {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.HivePartitioningOptions{}
	out.Mode = direct.LazyPtr(in.Mode)
	// MISSING: SourceURIPrefix
	// (near miss): "SourceURIPrefix" vs "SourceUriPrefix"
	out.RequirePartitionFilter = direct.LazyPtr(in.RequirePartitionFilter)
	// MISSING: Fields
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
	out.RequirePartitionFilter = direct.ValueOf(in.RequirePartitionFilter)
	// MISSING: Fields
	return out
}

func AvroOptions_FromProto(mapCtx *direct.MapContext, in *pb.AvroOptions) *krmv1beta1.AvroOptions {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AvroOptions{}
	out.UseAvroLogicalTypes = direct.LazyPtr(in.UseAvroLogicalTypes)
	return out
}

func AvroOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AvroOptions) *pb.AvroOptions {
	if in == nil {
		return nil
	}
	out := &pb.AvroOptions{}
	out.UseAvroLogicalTypes = direct.ValueOf(in.UseAvroLogicalTypes)
	return out
}
func ParquetOptions_FromProto(mapCtx *direct.MapContext, in *pb.ParquetOptions) *krmv1beta1.ParquetOptions {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ParquetOptions{}
	out.EnumAsString = direct.LazyPtr(in.EnumAsString)
	out.EnableListInference = direct.LazyPtr(in.EnableListInference)
	// MISSING: MapTargetType
	return out
}
func ParquetOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ParquetOptions) *pb.ParquetOptions {
	if in == nil {
		return nil
	}
	out := &pb.ParquetOptions{}
	out.EnumAsString = direct.ValueOf(in.EnumAsString)
	out.EnableListInference = direct.ValueOf(in.EnableListInference)
	// MISSING: MapTargetType
	return out
}
func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfiguration) *krmv1beta1.TableEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TableEncryptionConfiguration{}
	if in.KmsKeyName != "" {
		out.KmsKeyRef = &refs.KMSCryptoKeyRef{External: in.KmsKeyName}
	}
	return out
}

func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TableEncryptionConfiguration) *pb.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfiguration{}
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	return out
}

func TableConstraints_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraints) *krmv1beta1.TableConstraints {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TableConstraints{}
	out.PrimaryKey = PrimaryKey_FromProto(mapCtx, in.PrimaryKey)
	out.ForeignKeys = direct.Slice_FromProto(mapCtx, in.ForeignKeys, ForeignKey_FromProto)
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
func PrimaryKey_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraintsPrimaryKey) *krmv1beta1.PrimaryKey {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PrimaryKey{}
	out.Columns = in.Columns
	return out
}

func PrimaryKey_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PrimaryKey) *pb.TableConstraintsPrimaryKey {
	if in == nil {
		return nil
	}
	out := &pb.TableConstraintsPrimaryKey{}
	out.Columns = in.Columns
	return out
}
func ForeignKey_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraintsForeignKeys) *krmv1beta1.ForeignKey {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ForeignKey{}
	out.Name = direct.LazyPtr(in.Name)
	out.ReferencedTable = TableReference_FromProto(mapCtx, in.ReferencedTable)
	if len(in.ColumnReferences) > 0 {
		out.ColumnReferences = ColumnReference_FromProto(mapCtx, in.ColumnReferences[0])
	}
	return out
}
func ForeignKey_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ForeignKey) *pb.TableConstraintsForeignKeys {
	if in == nil {
		return nil
	}
	out := &pb.TableConstraintsForeignKeys{}
	out.Name = direct.ValueOf(in.Name)
	out.ReferencedTable = TableReference_ToProto(mapCtx, in.ReferencedTable)
	out.ColumnReferences = []*pb.TableConstraintsForeignKeysColumnReferences{ColumnReference_ToProto(mapCtx, in.ColumnReferences)}
	return out
}

func TableReference_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraintsForeignKeysReferencedTable) krmv1beta1.TableReference {
	if in == nil {
		return krmv1beta1.TableReference{}
	}
	out := krmv1beta1.TableReference{}
	out.DatasetId = direct.LazyPtr(in.DatasetId)
	out.ProjectId = direct.LazyPtr(in.ProjectId)
	out.TableId = direct.LazyPtr(in.TableId)
	return out
}

func ColumnReference_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraintsForeignKeysColumnReferences) krmv1beta1.ColumnReference {
	out := krmv1beta1.ColumnReference{}
	out.ReferencingColumn = direct.LazyPtr(in.ReferencingColumn)
	out.ReferencedColumn = direct.LazyPtr(in.ReferencedColumn)
	return out
}

func ColumnReference_ToProto(mapCtx *direct.MapContext, in krmv1beta1.ColumnReference) *pb.TableConstraintsForeignKeysColumnReferences {
	out := &pb.TableConstraintsForeignKeysColumnReferences{}
	out.ReferencingColumn = direct.ValueOf(in.ReferencingColumn)
	out.ReferencedColumn = direct.ValueOf(in.ReferencedColumn)
	return out
}

func TableReference_ToProto(mapCtx *direct.MapContext, in krmv1beta1.TableReference) *pb.TableConstraintsForeignKeysReferencedTable {
	out := &pb.TableConstraintsForeignKeysReferencedTable{}
	out.DatasetId = direct.ValueOf(in.DatasetId)
	out.ProjectId = direct.ValueOf(in.ProjectId)
	out.TableId = direct.ValueOf(in.TableId)
	return out
}

func GoogleSheetsOptions_FromProto(mapCtx *direct.MapContext, in *pb.GoogleSheetsOptions) *krmv1beta1.GoogleSheetsOptions {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.GoogleSheetsOptions{}
	out.SkipLeadingRows = direct.LazyPtr(in.SkipLeadingRows)
	out.Range = direct.LazyPtr(in.Range)
	return out
}

func GoogleSheetsOptions_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.GoogleSheetsOptions) *pb.GoogleSheetsOptions {
	if in == nil {
		return nil
	}
	out := &pb.GoogleSheetsOptions{}
	out.SkipLeadingRows = direct.ValueOf(in.SkipLeadingRows)
	out.Range = direct.ValueOf(in.Range)
	return out
}

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
	out.EncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.EncryptionConfiguration)
	// MISSING: SnapshotDefinition
	// MISSING: DefaultCollation
	// MISSING: DefaultRoundingMode
	// MISSING: CloneDefinition
	// MISSING: NumTimeTravelPhysicalBytesEncryptionConfiguration_FromProto
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

func BigQueryTableStatus_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmv1beta1.BigQueryTableStatus {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigQueryTableStatus{}
	out.CreationTime = direct.LazyPtr(in.CreationTime)
	out.Etag = direct.LazyPtr(in.Etag)
	out.LastModifiedTime = direct.LazyPtr(int64(in.LastModifiedTime))
	out.Location = direct.LazyPtr(in.Location)
	out.NumBytes = direct.LazyPtr(in.NumBytes)
	out.NumRows = direct.LazyPtr(int64(in.NumRows))
	out.SelfLink = direct.LazyPtr(in.SelfLink)
	out.Type = direct.LazyPtr(in.Type)
	return out
}

func BigQueryTableStatus_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigQueryTableStatus) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.CreationTime = direct.ValueOf(in.CreationTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.LastModifiedTime = uint64(direct.ValueOf(in.LastModifiedTime))
	out.Location = direct.ValueOf(in.Location)
	out.NumBytes = direct.ValueOf(in.NumBytes)
	out.NumRows = uint64(direct.ValueOf(in.NumRows))
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.Type = direct.ValueOf(in.Type)
	return out
}
