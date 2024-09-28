// Copyright 2024 Google LLC
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
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Access_FromProto(mapCtx *direct.MapContext, in *pb.DatasetAccess) *krm.Access {
	if in == nil {
		return nil
	}
	out := &krm.Access{}
	out.Role = direct.LazyPtr(in.GetRole())
	out.UserByEmail = direct.LazyPtr(in.GetUserByEmail())
	out.GroupByEmail = direct.LazyPtr(in.GetGroupByEmail())
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.SpecialGroup = direct.LazyPtr(in.GetSpecialGroup())
	out.IamMember = direct.LazyPtr(in.GetIamMember())
	out.View = TableReference_FromProto(mapCtx, in.GetView())
	out.Routine = RoutineReference_FromProto(mapCtx, in.GetRoutine())
	out.Dataset = DatasetAccessEntry_FromProto(mapCtx, in.GetDataset())
	return out
}
func Access_ToProto(mapCtx *direct.MapContext, in *krm.Access) *pb.DatasetAccess {
	if in == nil {
		return nil
	}
	out := &pb.DatasetAccess{}
	out.Role = in.Role
	out.UserByEmail = in.UserByEmail
	out.GroupByEmail = in.GroupByEmail
	out.Domain = in.Domain
	out.SpecialGroup = in.SpecialGroup
	out.IamMember = in.IamMember
	out.View = TableReference_ToProto(mapCtx, in.View)
	out.Routine = RoutineReference_ToProto(mapCtx, in.Routine)
	out.Dataset = DatasetAccessEntry_ToProto(mapCtx, in.Dataset)
	return out
}
func AggregationThresholdPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AggregationThresholdPolicy) *krm.AggregationThresholdPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AggregationThresholdPolicy{}
	out.Threshold = in.Threshold
	out.PrivacyUnitColumns = in.PrivacyUnitColumns
	return out
}
func AggregationThresholdPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AggregationThresholdPolicy) *pb.AggregationThresholdPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AggregationThresholdPolicy{}
	out.Threshold = in.Threshold
	out.PrivacyUnitColumns = in.PrivacyUnitColumns
	return out
}
func AvroOptions_FromProto(mapCtx *direct.MapContext, in *pb.AvroOptions) *krm.AvroOptions {
	if in == nil {
		return nil
	}
	out := &krm.AvroOptions{}
	out.UseAvroLogicalTypes = direct.LazyPtr(in.GetUseAvroLogicalTypes())
	return out
}
func AvroOptions_ToProto(mapCtx *direct.MapContext, in *krm.AvroOptions) *pb.AvroOptions {
	if in == nil {
		return nil
	}
	out := &pb.AvroOptions{}
	out.UseAvroLogicalTypes = in.UseAvroLogicalTypes
	return out
}
func BiEngineReason_FromProto(mapCtx *direct.MapContext, in *pb.BiEngineReason) *krm.BiEngineReason {
	if in == nil {
		return nil
	}
	out := &krm.BiEngineReason{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func BiEngineReason_ToProto(mapCtx *direct.MapContext, in *krm.BiEngineReason) *pb.BiEngineReason {
	if in == nil {
		return nil
	}
	out := &pb.BiEngineReason{}
	out.Code = in.Code
	out.Message = in.Message
	return out
}
func BiEngineStatistics_FromProto(mapCtx *direct.MapContext, in *pb.BiEngineStatistics) *krm.BiEngineStatistics {
	if in == nil {
		return nil
	}
	out := &krm.BiEngineStatistics{}
	out.BiEngineMode = direct.LazyPtr(in.GetBiEngineMode())
	out.AccelerationMode = direct.LazyPtr(in.GetAccelerationMode())
	out.BiEngineReasons = direct.Slice_FromProto(mapCtx, in.BiEngineReasons, BiEngineReason_FromProto)
	return out
}
func BiEngineStatistics_ToProto(mapCtx *direct.MapContext, in *krm.BiEngineStatistics) *pb.BiEngineStatistics {
	if in == nil {
		return nil
	}
	out := &pb.BiEngineStatistics{}
	out.BiEngineMode = in.BiEngineMode
	out.AccelerationMode = in.AccelerationMode
	out.BiEngineReasons = direct.Slice_ToProto(mapCtx, in.BiEngineReasons, BiEngineReason_ToProto)
	return out
}
func BigLakeConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.BigLakeConfiguration) *krm.BigLakeConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.BigLakeConfiguration{}
	out.ConnectionID = direct.LazyPtr(in.GetConnectionId())
	out.StorageUri = direct.LazyPtr(in.GetStorageUri())
	out.FileFormat = direct.LazyPtr(in.GetFileFormat())
	out.TableFormat = direct.LazyPtr(in.GetTableFormat())
	return out
}
func BigLakeConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.BigLakeConfiguration) *pb.BigLakeConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.BigLakeConfiguration{}
	out.ConnectionId = in.ConnectionID
	out.StorageUri = in.StorageUri
	out.FileFormat = in.FileFormat
	out.TableFormat = in.TableFormat
	return out
}
func BigtableColumn_FromProto(mapCtx *direct.MapContext, in *pb.BigtableColumn) *krm.BigtableColumn {
	if in == nil {
		return nil
	}
	out := &krm.BigtableColumn{}
	out.QualifierEncoded = &in.GetQualifierEncoded()[0]
	out.QualifierString = direct.LazyPtr(in.GetQualifierString())
	out.FieldName = direct.LazyPtr(in.GetFieldName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.OnlyReadLatest = direct.LazyPtr(in.GetOnlyReadLatest())
	return out
}
func BigtableColumn_ToProto(mapCtx *direct.MapContext, in *krm.BigtableColumn) *pb.BigtableColumn {
	if in == nil {
		return nil
	}
	out := &pb.BigtableColumn{}
	out.QualifierEncoded = []byte{*in.QualifierEncoded}
	out.QualifierString = in.QualifierString
	out.FieldName = in.FieldName
	out.Type = in.Type
	out.Encoding = in.Encoding
	out.OnlyReadLatest = in.OnlyReadLatest
	return out
}
func BigtableColumnFamily_FromProto(mapCtx *direct.MapContext, in *pb.BigtableColumnFamily) *krm.BigtableColumnFamily {
	if in == nil {
		return nil
	}
	out := &krm.BigtableColumnFamily{}
	out.FamilyID = direct.LazyPtr(in.GetFamilyId())
	out.Type = direct.LazyPtr(in.GetType())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, BigtableColumn_FromProto)
	out.OnlyReadLatest = direct.LazyPtr(in.GetOnlyReadLatest())
	return out
}
func BigtableColumnFamily_ToProto(mapCtx *direct.MapContext, in *krm.BigtableColumnFamily) *pb.BigtableColumnFamily {
	if in == nil {
		return nil
	}
	out := &pb.BigtableColumnFamily{}
	out.FamilyId = in.FamilyID
	out.Type = in.Type
	out.Encoding = in.Encoding
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, BigtableColumn_ToProto)
	out.OnlyReadLatest = in.OnlyReadLatest
	return out
}
func BigtableOptions_FromProto(mapCtx *direct.MapContext, in *pb.BigtableOptions) *krm.BigtableOptions {
	if in == nil {
		return nil
	}
	out := &krm.BigtableOptions{}
	out.ColumnFamilies = direct.Slice_FromProto(mapCtx, in.ColumnFamilies, BigtableColumnFamily_FromProto)
	out.IgnoreUnspecifiedColumnFamilies = direct.LazyPtr(in.GetIgnoreUnspecifiedColumnFamilies())
	out.ReadRowkeyAsString = direct.LazyPtr(in.GetReadRowkeyAsString())
	out.OutputColumnFamiliesAsJson = direct.LazyPtr(in.GetOutputColumnFamiliesAsJson())
	return out
}
func BigtableOptions_ToProto(mapCtx *direct.MapContext, in *krm.BigtableOptions) *pb.BigtableOptions {
	if in == nil {
		return nil
	}
	out := &pb.BigtableOptions{}
	out.ColumnFamilies = direct.Slice_ToProto(mapCtx, in.ColumnFamilies, BigtableColumnFamily_ToProto)
	out.IgnoreUnspecifiedColumnFamilies = in.IgnoreUnspecifiedColumnFamilies
	out.ReadRowkeyAsString = in.ReadRowkeyAsString
	out.OutputColumnFamiliesAsJson = in.OutputColumnFamiliesAsJson
	return out
}
func CloneDefinition_FromProto(mapCtx *direct.MapContext, in *pb.CloneDefinition) *krm.CloneDefinition {
	if in == nil {
		return nil
	}
	out := &krm.CloneDefinition{}
	out.BaseTableReference = TableReference_FromProto(mapCtx, in.GetBaseTableReference())
	out.CloneTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCloneTime())
	return out
}
func CloneDefinition_ToProto(mapCtx *direct.MapContext, in *krm.CloneDefinition) *pb.CloneDefinition {
	if in == nil {
		return nil
	}
	out := &pb.CloneDefinition{}
	out.BaseTableReference = TableReference_ToProto(mapCtx, in.BaseTableReference)
	out.CloneTime = direct.StringTimestamp_ToProto(mapCtx, in.CloneTime)
	return out
}
func Clustering_FromProto(mapCtx *direct.MapContext, in *pb.Clustering) *krm.Clustering {
	if in == nil {
		return nil
	}
	out := &krm.Clustering{}
	out.Fields = in.Fields
	return out
}
func Clustering_ToProto(mapCtx *direct.MapContext, in *krm.Clustering) *pb.Clustering {
	if in == nil {
		return nil
	}
	out := &pb.Clustering{}
	out.Fields = in.Fields
	return out
}
func ColumnReference_FromProto(mapCtx *direct.MapContext, in *pb.ColumnReferences) *krm.ColumnReference {
	if in == nil {
		return nil
	}
	out := &krm.ColumnReference{}
	out.ReferencingColumn = direct.LazyPtr(in.GetReferencingColumn())
	out.ReferencedColumn = direct.LazyPtr(in.GetReferencedColumn())
	return out
}
func ColumnReference_ToProto(mapCtx *direct.MapContext, in *krm.ColumnReference) *pb.ColumnReferences {
	if in == nil {
		return nil
	}
	out := &pb.ColumnReferences{}
	out.ReferencingColumn = in.ReferencingColumn
	out.ReferencedColumn = in.ReferencedColumn
	return out
}
func ConnectionProperty_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProperty) *krm.ConnectionProperty {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionProperty{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func ConnectionProperty_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionProperty) *pb.ConnectionProperty {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProperty{}
	out.Key = in.Key
	out.Value = in.Value
	return out
}
func CopyJobStatistics_FromProto(mapCtx *direct.MapContext, in *pb.JobStatistics5) *krm.CopyJobStatistics {
	if in == nil {
		return nil
	}
	out := &krm.CopyJobStatistics{}
	out.CopiedRows = direct.LazyPtr(in.GetCopiedRows())
	out.CopiedLogicalBytes = direct.LazyPtr(in.GetCopiedLogicalBytes())
	return out
}
func CopyJobStatistics_ToProto(mapCtx *direct.MapContext, in *krm.CopyJobStatistics) *pb.JobStatistics5 {
	if in == nil {
		return nil
	}
	out := &pb.JobStatistics5{}
	out.CopiedRows = in.CopiedRows
	out.CopiedLogicalBytes = in.CopiedLogicalBytes
	return out
}
func CsvOptions_FromProto(mapCtx *direct.MapContext, in *pb.CsvOptions) *krm.CsvOptions {
	if in == nil {
		return nil
	}
	out := &krm.CsvOptions{}
	out.FieldDelimiter = direct.LazyPtr(in.GetFieldDelimiter())
	out.SkipLeadingRows = direct.LazyPtr(in.GetSkipLeadingRows())
	out.Quote = direct.LazyPtr(in.GetQuote())
	out.AllowQuotedNewlines = direct.LazyPtr(in.GetAllowQuotedNewlines())
	out.AllowJaggedRows = direct.LazyPtr(in.GetAllowJaggedRows())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.PreserveAsciiControlCharacters = direct.LazyPtr(in.GetPreserveAsciiControlCharacters())
	out.NullMarker = direct.LazyPtr(in.GetNullMarker())
	return out
}
func CsvOptions_ToProto(mapCtx *direct.MapContext, in *krm.CsvOptions) *pb.CsvOptions {
	if in == nil {
		return nil
	}
	out := &pb.CsvOptions{}
	out.FieldDelimiter = in.FieldDelimiter
	out.SkipLeadingRows = in.SkipLeadingRows
	out.Quote = in.Quote
	out.AllowQuotedNewlines = in.AllowQuotedNewlines
	out.AllowJaggedRows = in.AllowJaggedRows
	out.Encoding = in.Encoding
	out.PreserveAsciiControlCharacters = in.PreserveAsciiControlCharacters
	out.NullMarker = in.NullMarker
	return out
}
func DataFormatOptions_FromProto(mapCtx *direct.MapContext, in *pb.DataFormatOptions) *krm.DataFormatOptions {
	if in == nil {
		return nil
	}
	out := &krm.DataFormatOptions{}
	out.UseInt64Timestamp = direct.LazyPtr(in.GetUseInt64Timestamp())
	return out
}
func DataFormatOptions_ToProto(mapCtx *direct.MapContext, in *krm.DataFormatOptions) *pb.DataFormatOptions {
	if in == nil {
		return nil
	}
	out := &pb.DataFormatOptions{}
	out.UseInt64Timestamp = in.UseInt64Timestamp
	return out
}
func DataMaskingStatistics_FromProto(mapCtx *direct.MapContext, in *pb.DataMaskingStatistics) *krm.DataMaskingStatistics {
	if in == nil {
		return nil
	}
	out := &krm.DataMaskingStatistics{}
	out.DataMaskingApplied = direct.LazyPtr(in.GetDataMaskingApplied())
	return out
}
func DataMaskingStatistics_ToProto(mapCtx *direct.MapContext, in *krm.DataMaskingStatistics) *pb.DataMaskingStatistics {
	if in == nil {
		return nil
	}
	out := &pb.DataMaskingStatistics{}
	out.DataMaskingApplied = in.DataMaskingApplied
	return out
}
func Dataset_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.Dataset {
	if in == nil {
		return nil
	}
	out := &krm.Dataset{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ID = direct.LazyPtr(in.GetId())
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.DatasetReference = DatasetReference_FromProto(mapCtx, in.GetDatasetReference())
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DefaultTableExpirationMs = direct.LazyPtr(in.GetDefaultTableExpirationMs())
	out.DefaultPartitionExpirationMs = direct.LazyPtr(in.GetDefaultPartitionExpirationMs())
	out.Labels = in.Labels
	out.Access = direct.Slice_FromProto(mapCtx, in.Access, Access_FromProto)
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.LastModifiedTime = direct.LazyPtr(in.GetLastModifiedTime())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.DefaultEncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetDefaultEncryptionConfiguration())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	out.Type = direct.LazyPtr(in.GetType())
	out.LinkedDatasetSource = LinkedDatasetSource_FromProto(mapCtx, in.GetLinkedDatasetSource())
	out.LinkedDatasetMetadata = LinkedDatasetMetadata_FromProto(mapCtx, in.GetLinkedDatasetMetadata())
	out.ExternalDatasetReference = ExternalDatasetReference_FromProto(mapCtx, in.GetExternalDatasetReference())
	out.ExternalCatalogDatasetOptions = ExternalCatalogDatasetOptions_FromProto(mapCtx, in.GetExternalCatalogDatasetOptions())
	out.IsCaseInsensitive = direct.LazyPtr(in.GetIsCaseInsensitive())
	out.DefaultCollation = direct.LazyPtr(in.GetDefaultCollation())
	out.DefaultRoundingMode = direct.LazyPtr(in.GetDefaultRoundingMode())
	out.MaxTimeTravelHours = direct.LazyPtr(in.GetMaxTimeTravelHours())
	out.Tags = direct.Slice_FromProto(mapCtx, in.Tags, GcpTag_FromProto)
	out.StorageBillingModel = direct.LazyPtr(in.GetStorageBillingModel())
	out.Restrictions = RestrictionConfig_FromProto(mapCtx, in.GetRestrictions())
	return out
}
func Dataset_ToProto(mapCtx *direct.MapContext, in *krm.Dataset) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	out.Kind = in.Kind
	out.Etag = in.Etag
	out.Id = in.ID
	out.SelfLink = in.SelfLink
	out.DatasetReference = DatasetReference_ToProto(mapCtx, in.DatasetReference)
	out.FriendlyName = in.FriendlyName
	out.Description = in.Description
	out.DefaultTableExpirationMs = in.DefaultTableExpirationMs
	out.DefaultPartitionExpirationMs = in.DefaultPartitionExpirationMs
	out.Labels = in.Labels
	out.Access = direct.Slice_ToProto(mapCtx, in.Access, Access_ToProto)
	out.CreationTime = in.CreationTime
	out.LastModifiedTime = in.LastModifiedTime
	out.Location = in.Location
	out.DefaultEncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.DefaultEncryptionConfiguration)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	out.Type = in.Type
	out.LinkedDatasetSource = LinkedDatasetSource_ToProto(mapCtx, in.LinkedDatasetSource)
	out.LinkedDatasetMetadata = LinkedDatasetMetadata_ToProto(mapCtx, in.LinkedDatasetMetadata)
	out.ExternalDatasetReference = ExternalDatasetReference_ToProto(mapCtx, in.ExternalDatasetReference)
	out.ExternalCatalogDatasetOptions = ExternalCatalogDatasetOptions_ToProto(mapCtx, in.ExternalCatalogDatasetOptions)
	out.IsCaseInsensitive = in.IsCaseInsensitive
	out.DefaultCollation = in.DefaultCollation
	out.DefaultRoundingMode = in.DefaultRoundingMode
	out.MaxTimeTravelHours = in.MaxTimeTravelHours
	out.Tags = direct.Slice_ToProto(mapCtx, in.Tags, GcpTag_ToProto)
	out.StorageBillingModel = in.StorageBillingModel
	out.Restrictions = RestrictionConfig_ToProto(mapCtx, in.Restrictions)
	return out
}
func DatasetAccessEntry_FromProto(mapCtx *direct.MapContext, in *pb.DatasetAccessEntry) *krm.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &krm.DatasetAccessEntry{}
	out.Dataset = DatasetReference_FromProto(mapCtx, in.GetDataset())
	out.TargetTypes = in.TargetTypes
	return out
}
func DatasetAccessEntry_ToProto(mapCtx *direct.MapContext, in *krm.DatasetAccessEntry) *pb.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &pb.DatasetAccessEntry{}
	out.Dataset = DatasetReference_ToProto(mapCtx, in.Dataset)
	out.TargetTypes = in.TargetTypes
	return out
}
func DatasetList_FromProto(mapCtx *direct.MapContext, in *pb.DatasetList) *krm.DatasetList {
	if in == nil {
		return nil
	}
	out := &krm.DatasetList{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.NextPageToken = direct.LazyPtr(in.GetNextPageToken())
	out.Datasets = direct.Slice_FromProto(mapCtx, in.Datasets, ListFormatDataset_FromProto)
	out.Unreachable = in.Unreachable
	return out
}
func DatasetList_ToProto(mapCtx *direct.MapContext, in *krm.DatasetList) *pb.DatasetList {
	if in == nil {
		return nil
	}
	out := &pb.DatasetList{}
	out.Kind = in.Kind
	out.Etag = in.Etag
	out.NextPageToken = in.NextPageToken
	out.Datasets = direct.Slice_ToProto(mapCtx, in.Datasets, ListFormatDataset_ToProto)
	out.Unreachable = in.Unreachable
	return out
}
func DatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.DatasetReference) *krm.DatasetReference {
	if in == nil {
		return nil
	}
	out := &krm.DatasetReference{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	return out
}
func DatasetReference_ToProto(mapCtx *direct.MapContext, in *krm.DatasetReference) *pb.DatasetReference {
	if in == nil {
		return nil
	}
	out := &pb.DatasetReference{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	return out
}
func DestinationTableProperties_FromProto(mapCtx *direct.MapContext, in *pb.DestinationTableProperties) *krm.DestinationTableProperties {
	if in == nil {
		return nil
	}
	out := &krm.DestinationTableProperties{}
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	return out
}
func DestinationTableProperties_ToProto(mapCtx *direct.MapContext, in *krm.DestinationTableProperties) *pb.DestinationTableProperties {
	if in == nil {
		return nil
	}
	out := &pb.DestinationTableProperties{}
	out.FriendlyName = in.FriendlyName
	out.Description = in.Description
	out.Labels = in.Labels
	return out
}
func DifferentialPrivacyPolicy_FromProto(mapCtx *direct.MapContext, in *pb.DifferentialPrivacyPolicy) *krm.DifferentialPrivacyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.DifferentialPrivacyPolicy{}
	out.MaxEpsilonPerQuery = in.MaxEpsilonPerQuery
	out.DeltaPerQuery = in.DeltaPerQuery
	out.MaxGroupsContributed = in.MaxGroupsContributed
	out.PrivacyUnitColumn = in.PrivacyUnitColumn
	out.EpsilonBudget = in.EpsilonBudget
	out.DeltaBudget = in.DeltaBudget
	out.EpsilonBudgetRemaining = in.EpsilonBudgetRemaining
	out.DeltaBudgetRemaining = in.DeltaBudgetRemaining
	return out
}
func DifferentialPrivacyPolicy_ToProto(mapCtx *direct.MapContext, in *krm.DifferentialPrivacyPolicy) *pb.DifferentialPrivacyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DifferentialPrivacyPolicy{}
	out.MaxEpsilonPerQuery = in.MaxEpsilonPerQuery
	out.DeltaPerQuery = in.DeltaPerQuery
	out.MaxGroupsContributed = in.MaxGroupsContributed
	out.PrivacyUnitColumn = in.PrivacyUnitColumn
	out.EpsilonBudget = in.EpsilonBudget
	out.DeltaBudget = in.DeltaBudget
	out.EpsilonBudgetRemaining = in.EpsilonBudgetRemaining
	out.DeltaBudgetRemaining = in.DeltaBudgetRemaining
	return out
}
func DmlStats_FromProto(mapCtx *direct.MapContext, in *pb.DmlStatistics) *krm.DmlStats {
	if in == nil {
		return nil
	}
	out := &krm.DmlStats{}
	out.InsertedRowCount = in.InsertedRowCount
	out.DeletedRowCount = in.DeletedRowCount
	out.UpdatedRowCount = in.UpdatedRowCount
	return out
}
func DmlStats_ToProto(mapCtx *direct.MapContext, in *krm.DmlStats) *pb.DmlStatistics {
	if in == nil {
		return nil
	}
	out := &pb.DmlStatistics{}
	out.InsertedRowCount = in.InsertedRowCount
	out.DeletedRowCount = in.DeletedRowCount
	out.UpdatedRowCount = in.UpdatedRowCount
	return out
}
func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfiguration) *krm.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfiguration{}
	out.KmsKeyRef.Name = *in.KmsKeyName
	return out
}
func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *pb.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfiguration{}
	out.KmsKeyName = &in.KmsKeyRef.Name
	return out
}
func ErrorProto_FromProto(mapCtx *direct.MapContext, in *pb.ErrorProto) *krm.ErrorProto {
	if in == nil {
		return nil
	}
	out := &krm.ErrorProto{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.DebugInfo = direct.LazyPtr(in.GetDebugInfo())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func ErrorProto_ToProto(mapCtx *direct.MapContext, in *krm.ErrorProto) *pb.ErrorProto {
	if in == nil {
		return nil
	}
	out := &pb.ErrorProto{}
	out.Reason = in.Reason
	out.Location = in.Location
	out.DebugInfo = in.DebugInfo
	out.Message = in.Message
	return out
}
func ExplainQueryStage_FromProto(mapCtx *direct.MapContext, in *pb.ExplainQueryStage) *krm.ExplainQueryStage {
	if in == nil {
		return nil
	}
	out := &krm.ExplainQueryStage{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.StartMs = direct.LazyPtr(in.GetStartMs())
	out.EndMs = direct.LazyPtr(in.GetEndMs())
	out.InputStages = in.InputStages
	out.WaitRatioAvg = direct.LazyPtr(in.GetWaitRatioAvg())
	out.WaitMsAvg = direct.LazyPtr(in.GetWaitMsAvg())
	out.WaitRatioMax = direct.LazyPtr(in.GetWaitRatioMax())
	out.WaitMsMax = direct.LazyPtr(in.GetWaitMsMax())
	out.ReadRatioAvg = direct.LazyPtr(in.GetReadRatioAvg())
	out.ReadMsAvg = direct.LazyPtr(in.GetReadMsAvg())
	out.ReadRatioMax = direct.LazyPtr(in.GetReadRatioMax())
	out.ReadMsMax = direct.LazyPtr(in.GetReadMsMax())
	out.ComputeRatioAvg = direct.LazyPtr(in.GetComputeRatioAvg())
	out.ComputeMsAvg = direct.LazyPtr(in.GetComputeMsAvg())
	out.ComputeRatioMax = direct.LazyPtr(in.GetComputeRatioMax())
	out.ComputeMsMax = direct.LazyPtr(in.GetComputeMsMax())
	out.WriteRatioAvg = direct.LazyPtr(in.GetWriteRatioAvg())
	out.WriteMsAvg = direct.LazyPtr(in.GetWriteMsAvg())
	out.WriteRatioMax = direct.LazyPtr(in.GetWriteRatioMax())
	out.WriteMsMax = direct.LazyPtr(in.GetWriteMsMax())
	out.ShuffleOutputBytes = direct.LazyPtr(in.GetShuffleOutputBytes())
	out.ShuffleOutputBytesSpilled = direct.LazyPtr(in.GetShuffleOutputBytesSpilled())
	out.RecordsRead = direct.LazyPtr(in.GetRecordsRead())
	out.RecordsWritten = direct.LazyPtr(in.GetRecordsWritten())
	out.ParallelInputs = direct.LazyPtr(in.GetParallelInputs())
	out.CompletedParallelInputs = direct.LazyPtr(in.GetCompletedParallelInputs())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, ExplainQueryStep_FromProto)
	out.SlotMs = direct.LazyPtr(in.GetSlotMs())
	out.ComputeMode = direct.LazyPtr(in.GetComputeMode())
	return out
}
func ExplainQueryStage_ToProto(mapCtx *direct.MapContext, in *krm.ExplainQueryStage) *pb.ExplainQueryStage {
	if in == nil {
		return nil
	}
	out := &pb.ExplainQueryStage{}
	out.Name = in.Name
	out.Id = in.ID
	out.StartMs = in.StartMs
	out.EndMs = in.EndMs
	out.InputStages = in.InputStages
	out.WaitRatioAvg = in.WaitRatioAvg
	out.WaitMsAvg = in.WaitMsAvg
	out.WaitRatioMax = in.WaitRatioMax
	out.WaitMsMax = in.WaitMsMax
	out.ReadRatioAvg = in.ReadRatioAvg
	out.ReadMsAvg = in.ReadMsAvg
	out.ReadRatioMax = in.ReadRatioMax
	out.ReadMsMax = in.ReadMsMax
	out.ComputeMsAvg = in.ComputeMsAvg
	out.ComputeRatioMax = in.ComputeRatioMax
	out.ComputeMsMax = in.ComputeMsMax
	out.WriteRatioAvg = in.WriteRatioAvg
	out.WriteMsAvg = in.WriteMsAvg
	out.WriteRatioMax = in.WriteRatioMax
	out.WriteMsMax = in.WriteMsMax
	out.ShuffleOutputBytes = in.ShuffleOutputBytes
	out.ShuffleOutputBytesSpilled = in.ShuffleOutputBytesSpilled
	out.RecordsRead = in.RecordsRead
	out.RecordsWritten = in.RecordsWritten
	out.ParallelInputs = in.ParallelInputs
	out.CompletedParallelInputs = in.CompletedParallelInputs
	out.Status = in.Status
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, ExplainQueryStep_ToProto)
	out.SlotMs = in.SlotMs
	out.ComputeMode = in.ComputeMode
	return out
}
func ExplainQueryStep_FromProto(mapCtx *direct.MapContext, in *pb.ExplainQueryStep) *krm.ExplainQueryStep {
	if in == nil {
		return nil
	}
	out := &krm.ExplainQueryStep{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Substeps = in.Substeps
	return out
}
func ExplainQueryStep_ToProto(mapCtx *direct.MapContext, in *krm.ExplainQueryStep) *pb.ExplainQueryStep {
	if in == nil {
		return nil
	}
	out := &pb.ExplainQueryStep{}
	out.Kind = in.Kind
	out.Substeps = in.Substeps
	return out
}
func ExportDataStatistics_FromProto(mapCtx *direct.MapContext, in *pb.ExportDataStatistics) *krm.ExportDataStatistics {
	if in == nil {
		return nil
	}
	out := &krm.ExportDataStatistics{}
	out.FileCount = direct.LazyPtr(in.GetFileCount())
	out.RowCount = direct.LazyPtr(in.GetRowCount())
	return out
}
func ExportDataStatistics_ToProto(mapCtx *direct.MapContext, in *krm.ExportDataStatistics) *pb.ExportDataStatistics {
	if in == nil {
		return nil
	}
	out := &pb.ExportDataStatistics{}
	out.FileCount = in.FileCount
	out.RowCount = in.RowCount
	return out
}
func ExternalCatalogDatasetOptions_FromProto(mapCtx *direct.MapContext, in *pb.ExternalCatalogDatasetOptions) *krm.ExternalCatalogDatasetOptions {
	if in == nil {
		return nil
	}
	out := &krm.ExternalCatalogDatasetOptions{}
	out.Parameters = in.Parameters
	out.DefaultStorageLocationUri = direct.LazyPtr(in.GetDefaultStorageLocationUri())
	return out
}
func ExternalCatalogDatasetOptions_ToProto(mapCtx *direct.MapContext, in *krm.ExternalCatalogDatasetOptions) *pb.ExternalCatalogDatasetOptions {
	if in == nil {
		return nil
	}
	out := &pb.ExternalCatalogDatasetOptions{}
	out.Parameters = in.Parameters
	out.DefaultStorageLocationUri = in.DefaultStorageLocationUri
	return out
}
func ExternalCatalogTableOptions_FromProto(mapCtx *direct.MapContext, in *pb.ExternalCatalogTableOptions) *krm.ExternalCatalogTableOptions {
	if in == nil {
		return nil
	}
	out := &krm.ExternalCatalogTableOptions{}
	out.Parameters = in.Parameters
	out.StorageDescriptor = StorageDescriptor_FromProto(mapCtx, in.GetStorageDescriptor())
	out.ConnectionID = direct.LazyPtr(in.GetConnectionId())
	return out
}
func ExternalCatalogTableOptions_ToProto(mapCtx *direct.MapContext, in *krm.ExternalCatalogTableOptions) *pb.ExternalCatalogTableOptions {
	if in == nil {
		return nil
	}
	out := &pb.ExternalCatalogTableOptions{}
	out.Parameters = in.Parameters
	out.StorageDescriptor = StorageDescriptor_ToProto(mapCtx, in.StorageDescriptor)
	out.ConnectionId = in.ConnectionID
	return out
}
func ExternalDataConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.ExternalDataConfiguration) *krm.ExternalDataConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.ExternalDataConfiguration{}
	out.SourceUris = in.SourceUris
	out.FileSetSpecType = direct.LazyPtr(in.GetFileSetSpecType())
	out.Schema = TableSchema_FromProto(mapCtx, in.GetSchema())
	out.SourceFormat = direct.LazyPtr(in.GetSourceFormat())
	out.MaxBadRecords = direct.LazyPtr(in.GetMaxBadRecords())
	out.Autodetect = direct.LazyPtr(in.GetAutodetect())
	out.IgnoreUnknownValues = direct.LazyPtr(in.GetIgnoreUnknownValues())
	out.Compression = direct.LazyPtr(in.GetCompression())
	out.CsvOptions = CsvOptions_FromProto(mapCtx, in.GetCsvOptions())
	out.JsonOptions = JsonOptions_FromProto(mapCtx, in.GetJsonOptions())
	out.BigtableOptions = BigtableOptions_FromProto(mapCtx, in.GetBigtableOptions())
	out.GoogleSheetsOptions = GoogleSheetsOptions_FromProto(mapCtx, in.GetGoogleSheetsOptions())
	out.HivePartitioningOptions = HivePartitioningOptions_FromProto(mapCtx, in.GetHivePartitioningOptions())
	out.ConnectionID = direct.LazyPtr(in.GetConnectionId())
	out.DecimalTargetTypes = in.DecimalTargetTypes
	out.AvroOptions = AvroOptions_FromProto(mapCtx, in.GetAvroOptions())
	out.JsonExtension = direct.LazyPtr(in.GetJsonExtension())
	out.ParquetOptions = ParquetOptions_FromProto(mapCtx, in.GetParquetOptions())
	out.ObjectMetadata = direct.LazyPtr(in.GetObjectMetadata())
	out.ReferenceFileSchemaUri = direct.LazyPtr(in.GetReferenceFileSchemaUri())
	out.MetadataCacheMode = direct.LazyPtr(in.GetMetadataCacheMode())
	return out
}
func ExternalDataConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.ExternalDataConfiguration) *pb.ExternalDataConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.ExternalDataConfiguration{}
	out.SourceUris = in.SourceUris
	out.FileSetSpecType = in.FileSetSpecType
	out.Schema = TableSchema_ToProto(mapCtx, in.Schema)
	out.SourceFormat = in.SourceFormat
	out.MaxBadRecords = in.MaxBadRecords
	out.Autodetect = in.Autodetect
	out.IgnoreUnknownValues = in.IgnoreUnknownValues
	out.Compression = in.Compression
	out.CsvOptions = CsvOptions_ToProto(mapCtx, in.CsvOptions)
	out.JsonOptions = JsonOptions_ToProto(mapCtx, in.JsonOptions)
	out.BigtableOptions = BigtableOptions_ToProto(mapCtx, in.BigtableOptions)
	out.GoogleSheetsOptions = GoogleSheetsOptions_ToProto(mapCtx, in.GoogleSheetsOptions)
	out.HivePartitioningOptions = HivePartitioningOptions_ToProto(mapCtx, in.HivePartitioningOptions)
	out.ConnectionId = in.ConnectionID
	out.DecimalTargetTypes = in.DecimalTargetTypes
	out.AvroOptions = AvroOptions_ToProto(mapCtx, in.AvroOptions)
	out.JsonExtension = in.JsonExtension
	out.ParquetOptions = ParquetOptions_ToProto(mapCtx, in.ParquetOptions)
	out.ObjectMetadata = in.ObjectMetadata
	out.ReferenceFileSchemaUri = in.ReferenceFileSchemaUri
	out.MetadataCacheMode = in.MetadataCacheMode
	return out
}
func ExternalDatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.ExternalDatasetReference) *krm.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &krm.ExternalDatasetReference{}
	out.ExternalSource = direct.LazyPtr(in.GetExternalSource())
	out.Connection = direct.LazyPtr(in.GetConnection())
	return out
}
func ExternalDatasetReference_ToProto(mapCtx *direct.MapContext, in *krm.ExternalDatasetReference) *pb.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &pb.ExternalDatasetReference{}
	out.ExternalSource = in.ExternalSource
	out.Connection = in.Connection
	return out
}
func ExternalServiceCost_FromProto(mapCtx *direct.MapContext, in *pb.ExternalServiceCost) *krm.ExternalServiceCost {
	if in == nil {
		return nil
	}
	out := &krm.ExternalServiceCost{}
	out.ExternalService = direct.LazyPtr(in.GetExternalService())
	out.BytesProcessed = direct.LazyPtr(in.GetBytesProcessed())
	out.BytesBilled = direct.LazyPtr(in.GetBytesBilled())
	out.SlotMs = direct.LazyPtr(in.GetSlotMs())
	out.ReservedSlotCount = direct.LazyPtr(in.GetReservedSlotCount())
	return out
}
func ExternalServiceCost_ToProto(mapCtx *direct.MapContext, in *krm.ExternalServiceCost) *pb.ExternalServiceCost {
	if in == nil {
		return nil
	}
	out := &pb.ExternalServiceCost{}
	out.ExternalService = in.ExternalService
	out.BytesProcessed = in.BytesProcessed
	out.BytesBilled = in.BytesBilled
	out.SlotMs = in.SlotMs
	out.ReservedSlotCount = in.ReservedSlotCount
	return out
}
func ForeignKey_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraintsForeignKeys) *krm.ForeignKey {
	if in == nil {
		return nil
	}
	out := &krm.ForeignKey{}
	out.Name = direct.LazyPtr(in.GetName())

	out.ReferencedTable = ReferencedTable_FromProto(mapCtx, in.GetReferencedTable())
	out.ColumnReferences = direct.Slice_FromProto(mapCtx, in.ColumnReferences, ColumnReference_FromProto)
	return out
}
func ForeignKey_ToProto(mapCtx *direct.MapContext, in *krm.ForeignKey) *pb.TableConstraintsForeignKeys {
	if in == nil {
		return nil
	}
	out := &pb.TableConstraintsForeignKeys{}
	out.Name = in.Name
	out.ReferencedTable = ReferencedTable_ToProto(mapCtx, in.ReferencedTable)
	out.ColumnReferences = direct.Slice_ToProto(mapCtx, in.ColumnReferences, ColumnReference_ToProto)
	return out
}
func ForeignTypeInfo_FromProto(mapCtx *direct.MapContext, in *pb.ForeignTypeInfo) *krm.ForeignTypeInfo {
	if in == nil {
		return nil
	}
	out := &krm.ForeignTypeInfo{}
	out.TypeSystem = direct.LazyPtr(in.GetTypeSystem())
	return out
}
func ForeignTypeInfo_ToProto(mapCtx *direct.MapContext, in *krm.ForeignTypeInfo) *pb.ForeignTypeInfo {
	if in == nil {
		return nil
	}
	out := &pb.ForeignTypeInfo{}
	out.TypeSystem = in.TypeSystem
	return out
}
func GcpTag_FromProto(mapCtx *direct.MapContext, in *pb.DatasetTags) *krm.GcpTag {
	if in == nil {
		return nil
	}
	out := &krm.GcpTag{}
	out.TagKey = direct.LazyPtr(in.GetTagKey())
	out.TagValue = direct.LazyPtr(in.GetTagValue())
	return out
}
func GcpTag_ToProto(mapCtx *direct.MapContext, in *krm.GcpTag) *pb.DatasetTags {
	if in == nil {
		return nil
	}
	out := &pb.DatasetTags{}
	out.TagKey = in.TagKey
	out.TagValue = in.TagValue
	return out
}
func GoogleSheetsOptions_FromProto(mapCtx *direct.MapContext, in *pb.GoogleSheetsOptions) *krm.GoogleSheetsOptions {
	if in == nil {
		return nil
	}
	out := &krm.GoogleSheetsOptions{}
	out.SkipLeadingRows = direct.LazyPtr(in.GetSkipLeadingRows())
	out.Range = direct.LazyPtr(in.GetRange())
	return out
}
func GoogleSheetsOptions_ToProto(mapCtx *direct.MapContext, in *krm.GoogleSheetsOptions) *pb.GoogleSheetsOptions {
	if in == nil {
		return nil
	}
	out := &pb.GoogleSheetsOptions{}
	out.SkipLeadingRows = in.SkipLeadingRows
	out.Range = in.Range
	return out
}
func HighCardinalityJoin_FromProto(mapCtx *direct.MapContext, in *pb.HighCardinalityJoin) *krm.HighCardinalityJoin {
	if in == nil {
		return nil
	}
	out := &krm.HighCardinalityJoin{}
	out.LeftRows = direct.LazyPtr(in.GetLeftRows())
	out.RightRows = direct.LazyPtr(in.GetRightRows())
	out.OutputRows = direct.LazyPtr(in.GetOutputRows())
	out.StepIndex = direct.LazyPtr(in.GetStepIndex())
	return out
}
func HighCardinalityJoin_ToProto(mapCtx *direct.MapContext, in *krm.HighCardinalityJoin) *pb.HighCardinalityJoin {
	if in == nil {
		return nil
	}
	out := &pb.HighCardinalityJoin{}
	out.LeftRows = in.LeftRows
	out.RightRows = in.RightRows
	out.OutputRows = in.OutputRows
	out.StepIndex = in.StepIndex
	return out
}
func HivePartitioningOptions_FromProto(mapCtx *direct.MapContext, in *pb.HivePartitioningOptions) *krm.HivePartitioningOptions {
	if in == nil {
		return nil
	}
	out := &krm.HivePartitioningOptions{}
	out.Mode = direct.LazyPtr(in.GetMode())
	out.SourceUriPrefix = direct.LazyPtr(in.GetSourceUriPrefix())
	out.RequirePartitionFilter = direct.LazyPtr(in.GetRequirePartitionFilter())
	out.Fields = in.Fields
	return out
}
func HivePartitioningOptions_ToProto(mapCtx *direct.MapContext, in *krm.HivePartitioningOptions) *pb.HivePartitioningOptions {
	if in == nil {
		return nil
	}
	out := &pb.HivePartitioningOptions{}
	out.Mode = in.Mode
	out.SourceUriPrefix = in.SourceUriPrefix
	out.RequirePartitionFilter = in.RequirePartitionFilter
	out.Fields = in.Fields
	return out
}
func IndexUnusedReason_FromProto(mapCtx *direct.MapContext, in *pb.IndexUnusedReason) *krm.IndexUnusedReason {
	if in == nil {
		return nil
	}
	out := &krm.IndexUnusedReason{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = in.Message
	out.BaseTable = TableReference_FromProto(mapCtx, in.GetBaseTable())
	out.IndexName = in.IndexName
	return out
}
func IndexUnusedReason_ToProto(mapCtx *direct.MapContext, in *krm.IndexUnusedReason) *pb.IndexUnusedReason {
	if in == nil {
		return nil
	}
	out := &pb.IndexUnusedReason{}
	out.Code = in.Code
	out.Message = in.Message
	out.BaseTable = TableReference_ToProto(mapCtx, in.BaseTable)
	out.IndexName = in.IndexName
	return out
}
func InputDataChange_FromProto(mapCtx *direct.MapContext, in *pb.InputDataChange) *krm.InputDataChange {
	if in == nil {
		return nil
	}
	out := &krm.InputDataChange{}
	out.RecordsReadDiffPercentage = direct.LazyPtr(float64(in.GetRecordsReadDiffPercentage()))
	return out
}
func InputDataChange_ToProto(mapCtx *direct.MapContext, in *krm.InputDataChange) *pb.InputDataChange {
	if in == nil {
		return nil
	}
	out := &pb.InputDataChange{}
	out.RecordsReadDiffPercentage = direct.LazyPtr(float32(*in.RecordsReadDiffPercentage))
	return out
}
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ID = direct.LazyPtr(in.GetId())
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.UserEmail = direct.LazyPtr(in.GetUserEmail())
	out.Configuration = JobConfiguration_FromProto(mapCtx, in.GetConfiguration())
	out.JobReference = JobReference_FromProto(mapCtx, in.GetJobReference())
	out.Statistics = JobStatistics_FromProto(mapCtx, in.GetStatistics())
	out.Status = JobStatus_FromProto(mapCtx, in.GetStatus())
	out.PrincipalSubject = direct.LazyPtr(in.GetPrincipalSubject())
	out.JobCreationReason = JobCreationReason_FromProto(mapCtx, in.GetJobCreationReason())
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Kind = in.Kind
	out.Etag = in.Etag
	out.Id = in.ID
	out.SelfLink = in.SelfLink
	out.UserEmail = in.UserEmail
	out.Configuration = JobConfiguration_ToProto(mapCtx, in.Configuration)
	out.JobReference = JobReference_ToProto(mapCtx, in.JobReference)
	out.Statistics = JobStatistics_ToProto(mapCtx, in.Statistics)
	out.Status = JobStatus_ToProto(mapCtx, in.Status)
	out.PrincipalSubject = in.PrincipalSubject
	out.JobCreationReason = JobCreationReason_ToProto(mapCtx, in.JobCreationReason)
	return out
}
func JobConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.JobConfiguration) *krm.JobConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.JobConfiguration{}
	out.JobType = direct.LazyPtr(in.GetJobType())
	out.Query = JobConfigurationQuery_FromProto(mapCtx, in.GetQuery())
	out.Load = JobConfigurationLoad_FromProto(mapCtx, in.GetLoad())
	out.Copy = JobConfigurationTableCopy_FromProto(mapCtx, in.GetCopy())
	out.Extract = JobConfigurationExtract_FromProto(mapCtx, in.GetExtract())
	out.DryRun = direct.LazyPtr(in.GetDryRun())
	out.JobTimeoutMs = direct.LazyPtr(in.GetJobTimeoutMs())
	out.Labels = in.Labels
	return out
}
func JobConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.JobConfiguration) *pb.JobConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.JobConfiguration{}
	out.JobType = in.JobType
	out.Query = JobConfigurationQuery_ToProto(mapCtx, in.Query)
	out.Load = JobConfigurationLoad_ToProto(mapCtx, in.Load)
	out.Copy = JobConfigurationTableCopy_ToProto(mapCtx, in.Copy)
	out.Extract = JobConfigurationExtract_ToProto(mapCtx, in.Extract)
	out.DryRun = in.DryRun
	out.JobTimeoutMs = in.JobTimeoutMs
	out.Labels = in.Labels
	return out
}
func JobConfigurationExtract_FromProto(mapCtx *direct.MapContext, in *pb.JobConfigurationExtract) *krm.JobConfigurationExtract {
	if in == nil {
		return nil
	}
	out := &krm.JobConfigurationExtract{}
	out.SourceTable = TableReference_FromProto(mapCtx, in.GetSourceTable())
	out.SourceModel = ModelReference_FromProto(mapCtx, in.GetSourceModel())
	out.DestinationUris = in.DestinationUris
	out.PrintHeader = direct.LazyPtr(in.GetPrintHeader())
	out.FieldDelimiter = direct.LazyPtr(in.GetFieldDelimiter())
	out.DestinationFormat = direct.LazyPtr(in.GetDestinationFormat())
	out.Compression = direct.LazyPtr(in.GetCompression())
	out.UseAvroLogicalTypes = direct.LazyPtr(in.GetUseAvroLogicalTypes())
	out.ModelExtractOptions = JobConfigurationExtract_ModelExtractOptions_FromProto(mapCtx, in.GetModelExtractOptions())
	return out
}
func JobConfigurationExtract_ToProto(mapCtx *direct.MapContext, in *krm.JobConfigurationExtract) *pb.JobConfigurationExtract {
	if in == nil {
		return nil
	}
	out := &pb.JobConfigurationExtract{}
	out.SourceTable = TableReference_ToProto(mapCtx, in.SourceTable)
	out.SourceModel = ModelReference_ToProto(mapCtx, in.SourceModel)
	out.DestinationUris = in.DestinationUris
	out.PrintHeader = in.PrintHeader
	out.FieldDelimiter = in.FieldDelimiter
	out.DestinationFormat = in.DestinationFormat
	out.Compression = in.Compression
	out.UseAvroLogicalTypes = in.UseAvroLogicalTypes
	out.ModelExtractOptions = JobConfigurationExtract_ModelExtractOptions_ToProto(mapCtx, in.ModelExtractOptions)
	return out
}
func JobConfigurationExtract_ModelExtractOptions_FromProto(mapCtx *direct.MapContext, in *pb.ModelExtractOptions) *krm.JobConfigurationExtract_ModelExtractOptions {
	if in == nil {
		return nil
	}
	out := &krm.JobConfigurationExtract_ModelExtractOptions{}
	out.TrialID = direct.LazyPtr(in.GetTrialId())
	return out
}
func JobConfigurationExtract_ModelExtractOptions_ToProto(mapCtx *direct.MapContext, in *krm.JobConfigurationExtract_ModelExtractOptions) *pb.ModelExtractOptions {
	if in == nil {
		return nil
	}
	out := &pb.ModelExtractOptions{}
	out.TrialId = in.TrialID
	return out
}
func JobConfigurationLoad_FromProto(mapCtx *direct.MapContext, in *pb.JobConfigurationLoad) *krm.JobConfigurationLoad {
	if in == nil {
		return nil
	}
	out := &krm.JobConfigurationLoad{}
	out.SourceUris = in.SourceUris
	out.FileSetSpecType = direct.LazyPtr(in.GetFileSetSpecType())
	out.Schema = TableSchema_FromProto(mapCtx, in.GetSchema())
	out.DestinationTable = TableReference_FromProto(mapCtx, in.GetDestinationTable())
	out.DestinationTableProperties = DestinationTableProperties_FromProto(mapCtx, in.GetDestinationTableProperties())
	out.CreateDisposition = direct.LazyPtr(in.GetCreateDisposition())
	out.WriteDisposition = direct.LazyPtr(in.GetWriteDisposition())
	out.NullMarker = direct.LazyPtr(in.GetNullMarker())
	out.FieldDelimiter = direct.LazyPtr(in.GetFieldDelimiter())
	out.SkipLeadingRows = direct.LazyPtr(in.GetSkipLeadingRows())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.Quote = direct.LazyPtr(in.GetQuote())
	out.MaxBadRecords = direct.LazyPtr(in.GetMaxBadRecords())
	out.AllowQuotedNewlines = direct.LazyPtr(in.GetAllowQuotedNewlines())
	out.SourceFormat = direct.LazyPtr(in.GetSourceFormat())
	out.AllowJaggedRows = direct.LazyPtr(in.GetAllowJaggedRows())
	out.IgnoreUnknownValues = direct.LazyPtr(in.GetIgnoreUnknownValues())
	out.ProjectionFields = in.ProjectionFields
	out.Autodetect = direct.LazyPtr(in.GetAutodetect())
	out.SchemaUpdateOptions = in.SchemaUpdateOptions
	out.TimePartitioning = TimePartitioning_FromProto(mapCtx, in.GetTimePartitioning())
	out.RangePartitioning = RangePartitioning_FromProto(mapCtx, in.GetRangePartitioning())
	out.Clustering = Clustering_FromProto(mapCtx, in.GetClustering())
	out.DestinationEncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetDestinationEncryptionConfiguration())
	out.UseAvroLogicalTypes = direct.LazyPtr(in.GetUseAvroLogicalTypes())
	out.ReferenceFileSchemaUri = direct.LazyPtr(in.GetReferenceFileSchemaUri())
	out.HivePartitioningOptions = HivePartitioningOptions_FromProto(mapCtx, in.GetHivePartitioningOptions())
	out.DecimalTargetTypes = in.DecimalTargetTypes
	out.JsonExtension = direct.LazyPtr(in.GetJsonExtension())
	out.ParquetOptions = ParquetOptions_FromProto(mapCtx, in.GetParquetOptions())
	out.PreserveAsciiControlCharacters = direct.LazyPtr(in.GetPreserveAsciiControlCharacters())
	out.ConnectionProperties = direct.Slice_FromProto(mapCtx, in.ConnectionProperties, ConnectionProperty_FromProto)
	out.CreateSession = direct.LazyPtr(in.GetCreateSession())
	out.CopyFilesOnly = direct.LazyPtr(in.GetCopyFilesOnly())
	return out
}
func JobConfigurationLoad_ToProto(mapCtx *direct.MapContext, in *krm.JobConfigurationLoad) *pb.JobConfigurationLoad {
	if in == nil {
		return nil
	}
	out := &pb.JobConfigurationLoad{}
	out.SourceUris = in.SourceUris
	out.FileSetSpecType = in.FileSetSpecType
	out.Schema = TableSchema_ToProto(mapCtx, in.Schema)
	out.DestinationTable = TableReference_ToProto(mapCtx, in.DestinationTable)
	out.DestinationTableProperties = DestinationTableProperties_ToProto(mapCtx, in.DestinationTableProperties)
	out.CreateDisposition = in.CreateDisposition
	out.WriteDisposition = in.WriteDisposition
	out.NullMarker = in.NullMarker
	out.FieldDelimiter = in.FieldDelimiter
	out.SkipLeadingRows = in.SkipLeadingRows
	out.Encoding = in.Encoding
	out.Quote = in.Quote
	out.MaxBadRecords = in.MaxBadRecords
	out.AllowQuotedNewlines = in.AllowQuotedNewlines
	out.SourceFormat = in.SourceFormat
	out.AllowJaggedRows = in.AllowJaggedRows
	out.IgnoreUnknownValues = in.IgnoreUnknownValues
	out.ProjectionFields = in.ProjectionFields
	out.Autodetect = in.Autodetect
	out.SchemaUpdateOptions = in.SchemaUpdateOptions
	out.TimePartitioning = TimePartitioning_ToProto(mapCtx, in.TimePartitioning)
	out.RangePartitioning = RangePartitioning_ToProto(mapCtx, in.RangePartitioning)
	out.Clustering = Clustering_ToProto(mapCtx, in.Clustering)
	out.DestinationEncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.DestinationEncryptionConfiguration)
	out.UseAvroLogicalTypes = in.UseAvroLogicalTypes
	out.ReferenceFileSchemaUri = in.ReferenceFileSchemaUri
	out.HivePartitioningOptions = HivePartitioningOptions_ToProto(mapCtx, in.HivePartitioningOptions)
	out.DecimalTargetTypes = in.DecimalTargetTypes
	out.JsonExtension = in.JsonExtension
	out.ParquetOptions = ParquetOptions_ToProto(mapCtx, in.ParquetOptions)
	out.PreserveAsciiControlCharacters = in.PreserveAsciiControlCharacters
	out.ConnectionProperties = direct.Slice_ToProto(mapCtx, in.ConnectionProperties, ConnectionProperty_ToProto)
	out.CreateSession = in.CreateSession
	out.CopyFilesOnly = in.CopyFilesOnly
	return out
}
func JobConfigurationQuery_FromProto(mapCtx *direct.MapContext, in *pb.JobConfigurationQuery) *krm.JobConfigurationQuery {
	if in == nil {
		return nil
	}
	out := &krm.JobConfigurationQuery{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.DestinationTable = TableReference_FromProto(mapCtx, in.GetDestinationTable())
	// 	/* NOTYET */: ExternalTableDefinitions
	out.UserDefinedFunctionResources = direct.Slice_FromProto(mapCtx, in.UserDefinedFunctionResources, UserDefinedFunctionResource_FromProto)
	out.CreateDisposition = direct.LazyPtr(in.GetCreateDisposition())
	out.WriteDisposition = direct.LazyPtr(in.GetWriteDisposition())
	out.DefaultDataset = DatasetReference_FromProto(mapCtx, in.GetDefaultDataset())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.AllowLargeResults = direct.LazyPtr(in.GetAllowLargeResults())
	out.UseQueryCache = direct.LazyPtr(in.GetUseQueryCache())
	out.FlattenResults = direct.LazyPtr(in.GetFlattenResults())
	out.MaximumBytesBilled = direct.LazyPtr(in.GetMaximumBytesBilled())
	out.UseLegacySql = direct.LazyPtr(in.GetUseLegacySql())
	out.ParameterMode = direct.LazyPtr(in.GetParameterMode())
	out.QueryParameters = direct.Slice_FromProto(mapCtx, in.QueryParameters, QueryParameter_FromProto)
	out.SystemVariables = SystemVariables_FromProto(mapCtx, in.GetSystemVariables())
	out.SchemaUpdateOptions = in.SchemaUpdateOptions
	out.TimePartitioning = TimePartitioning_FromProto(mapCtx, in.GetTimePartitioning())
	out.RangePartitioning = RangePartitioning_FromProto(mapCtx, in.GetRangePartitioning())
	out.Clustering = Clustering_FromProto(mapCtx, in.GetClustering())
	out.DestinationEncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetDestinationEncryptionConfiguration())
	out.ScriptOptions = ScriptOptions_FromProto(mapCtx, in.GetScriptOptions())
	out.ConnectionProperties = direct.Slice_FromProto(mapCtx, in.ConnectionProperties, ConnectionProperty_FromProto)
	out.CreateSession = direct.LazyPtr(in.GetCreateSession())
	out.Continuous = direct.LazyPtr(in.GetContinuous())
	return out
}
func JobConfigurationQuery_ToProto(mapCtx *direct.MapContext, in *krm.JobConfigurationQuery) *pb.JobConfigurationQuery {
	if in == nil {
		return nil
	}
	out := &pb.JobConfigurationQuery{}
	out.Query = in.Query
	out.DestinationTable = TableReference_ToProto(mapCtx, in.DestinationTable)
	// 	/* NOTYET */: ExternalTableDefinitions
	out.UserDefinedFunctionResources = direct.Slice_ToProto(mapCtx, in.UserDefinedFunctionResources, UserDefinedFunctionResource_ToProto)
	out.CreateDisposition = in.CreateDisposition
	out.WriteDisposition = in.WriteDisposition
	out.DefaultDataset = DatasetReference_ToProto(mapCtx, in.DefaultDataset)
	out.Priority = in.Priority
	out.AllowLargeResults = in.AllowLargeResults
	out.UseQueryCache = in.UseQueryCache
	out.FlattenResults = in.FlattenResults
	out.MaximumBytesBilled = in.MaximumBytesBilled
	out.UseLegacySql = in.UseLegacySql
	out.ParameterMode = in.ParameterMode
	out.QueryParameters = direct.Slice_ToProto(mapCtx, in.QueryParameters, QueryParameter_ToProto)
	if oneof := SystemVariables_ToProto(mapCtx, in.SystemVariables); oneof != nil {
		out.SystemVariables = oneof
	}
	out.SchemaUpdateOptions = in.SchemaUpdateOptions
	out.TimePartitioning = TimePartitioning_ToProto(mapCtx, in.TimePartitioning)
	out.RangePartitioning = RangePartitioning_ToProto(mapCtx, in.RangePartitioning)
	out.Clustering = Clustering_ToProto(mapCtx, in.Clustering)
	out.DestinationEncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.DestinationEncryptionConfiguration)
	out.ScriptOptions = ScriptOptions_ToProto(mapCtx, in.ScriptOptions)
	out.ConnectionProperties = direct.Slice_ToProto(mapCtx, in.ConnectionProperties, ConnectionProperty_ToProto)
	out.CreateSession = in.CreateSession
	out.Continuous = in.Continuous
	return out
}
func JobConfigurationTableCopy_FromProto(mapCtx *direct.MapContext, in *pb.JobConfigurationTableCopy) *krm.JobConfigurationTableCopy {
	if in == nil {
		return nil
	}
	out := &krm.JobConfigurationTableCopy{}
	out.SourceTable = TableReference_FromProto(mapCtx, in.GetSourceTable())
	out.SourceTables = direct.Slice_FromProto(mapCtx, in.SourceTables, TableReference_FromProto)
	out.DestinationTable = TableReference_FromProto(mapCtx, in.GetDestinationTable())
	out.CreateDisposition = direct.LazyPtr(in.GetCreateDisposition())
	out.WriteDisposition = direct.LazyPtr(in.GetWriteDisposition())
	out.DestinationEncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetDestinationEncryptionConfiguration())
	out.OperationType = direct.LazyPtr(in.GetOperationType())
	out.DestinationExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDestinationExpirationTime())
	return out
}
func JobConfigurationTableCopy_ToProto(mapCtx *direct.MapContext, in *krm.JobConfigurationTableCopy) *pb.JobConfigurationTableCopy {
	if in == nil {
		return nil
	}
	out := &pb.JobConfigurationTableCopy{}
	out.SourceTable = TableReference_ToProto(mapCtx, in.SourceTable)
	out.SourceTables = direct.Slice_ToProto(mapCtx, in.SourceTables, TableReference_ToProto)
	out.DestinationTable = TableReference_ToProto(mapCtx, in.DestinationTable)
	out.CreateDisposition = in.CreateDisposition
	out.WriteDisposition = in.WriteDisposition
	out.DestinationEncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.DestinationEncryptionConfiguration)
	out.OperationType = in.OperationType
	out.DestinationExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.DestinationExpirationTime)
	return out
}
func JobCreationReason_FromProto(mapCtx *direct.MapContext, in *pb.JobCreationReason) *krm.JobCreationReason {
	if in == nil {
		return nil
	}
	out := &krm.JobCreationReason{}
	out.Code = direct.LazyPtr(in.GetCode())
	return out
}
func JobCreationReason_ToProto(mapCtx *direct.MapContext, in *krm.JobCreationReason) *pb.JobCreationReason {
	if in == nil {
		return nil
	}
	out := &pb.JobCreationReason{}
	out.Code = in.Code
	return out
}
func JobList_FromProto(mapCtx *direct.MapContext, in *pb.JobList) *krm.JobList {
	if in == nil {
		return nil
	}
	out := &krm.JobList{}
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Kind = direct.LazyPtr(in.GetKind())
	out.NextPageToken = direct.LazyPtr(in.GetNextPageToken())
	out.Jobs = direct.Slice_FromProto(mapCtx, in.Jobs, ListFormatJob_FromProto)
	out.Unreachable = in.Unreachable
	return out
}
func JobList_ToProto(mapCtx *direct.MapContext, in *krm.JobList) *pb.JobList {
	if in == nil {
		return nil
	}
	out := &pb.JobList{}
	out.Etag = in.Etag
	out.Kind = in.Kind
	out.NextPageToken = in.NextPageToken
	out.Jobs = direct.Slice_ToProto(mapCtx, in.Jobs, ListFormatJob_ToProto)
	out.Unreachable = in.Unreachable
	return out
}
func JobReference_FromProto(mapCtx *direct.MapContext, in *pb.JobReference) *krm.JobReference {
	if in == nil {
		return nil
	}
	out := &krm.JobReference{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.JobID = direct.LazyPtr(in.GetJobId())
	out.Location = direct.LazyPtr((in.GetLocation()))
	return out
}
func JobReference_ToProto(mapCtx *direct.MapContext, in *krm.JobReference) *pb.JobReference {
	if in == nil {
		return nil
	}
	out := &pb.JobReference{}
	out.ProjectId = in.ProjectID
	out.JobId = in.JobID
	out.Location = in.Location
	return out
}
func JobStatistics_FromProto(mapCtx *direct.MapContext, in *pb.JobStatistics) *krm.JobStatistics {
	if in == nil {
		return nil
	}
	out := &krm.JobStatistics{}
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.StartTime = direct.LazyPtr(in.GetStartTime())
	out.EndTime = direct.LazyPtr(in.GetEndTime())
	out.TotalBytesProcessed = direct.LazyPtr(in.GetTotalBytesProcessed())
	out.CompletionRatio = direct.LazyPtr(in.GetCompletionRatio())
	out.QuotaDeferments = in.QuotaDeferments
	out.Query = JobStatistics2_FromProto(mapCtx, in.GetQuery())
	out.Load = JobStatistics3_FromProto(mapCtx, in.GetLoad())
	out.Extract = JobStatistics4_FromProto(mapCtx, in.GetExtract())
	out.Copy = CopyJobStatistics_FromProto(mapCtx, in.GetCopy())
	out.TotalSlotMs = direct.LazyPtr(in.GetTotalSlotMs())
	out.ReservationID = direct.LazyPtr(in.GetReservationId())
	out.NumChildJobs = direct.LazyPtr(in.GetNumChildJobs())
	out.ParentJobID = direct.LazyPtr(in.GetParentJobId())
	out.ScriptStatistics = ScriptStatistics_FromProto(mapCtx, in.GetScriptStatistics())
	out.RowLevelSecurityStatistics = RowLevelSecurityStatistics_FromProto(mapCtx, in.GetRowLevelSecurityStatistics())
	out.DataMaskingStatistics = DataMaskingStatistics_FromProto(mapCtx, in.GetDataMaskingStatistics())
	out.TransactionInfo = JobStatistics_TransactionInfo_FromProto(mapCtx, in.GetTransactionInfo())
	out.SessionInfo = SessionInfo_FromProto(mapCtx, in.GetSessionInfo())
	out.FinalExecutionDurationMs = direct.LazyPtr(in.GetFinalExecutionDurationMs())
	return out
}
func JobStatistics_ToProto(mapCtx *direct.MapContext, in *krm.JobStatistics) *pb.JobStatistics {
	if in == nil {
		return nil
	}
	out := &pb.JobStatistics{}
	out.CreationTime = in.CreationTime
	out.StartTime = in.StartTime
	out.EndTime = in.EndTime
	out.TotalBytesProcessed = in.TotalBytesProcessed
	out.CompletionRatio = in.CompletionRatio
	out.QuotaDeferments = in.QuotaDeferments
	out.Query = JobStatistics2_ToProto(mapCtx, in.Query)
	out.Load = JobStatistics3_ToProto(mapCtx, in.Load)
	out.Extract = JobStatistics4_ToProto(mapCtx, in.Extract)
	out.Copy = CopyJobStatistics_ToProto(mapCtx, in.Copy)
	out.TotalSlotMs = in.TotalSlotMs
	out.ReservationId = in.ReservationID
	out.NumChildJobs = in.NumChildJobs
	out.ParentJobId = in.ParentJobID
	out.ScriptStatistics = ScriptStatistics_ToProto(mapCtx, in.ScriptStatistics)
	out.RowLevelSecurityStatistics = RowLevelSecurityStatistics_ToProto(mapCtx, in.RowLevelSecurityStatistics)
	out.DataMaskingStatistics = DataMaskingStatistics_ToProto(mapCtx, in.DataMaskingStatistics)
	out.TransactionInfo = JobStatistics_TransactionInfo_ToProto(mapCtx, in.TransactionInfo)
	out.SessionInfo = SessionInfo_ToProto(mapCtx, in.SessionInfo)
	out.FinalExecutionDurationMs = in.FinalExecutionDurationMs
	return out
}
func JobStatistics2_FromProto(mapCtx *direct.MapContext, in *pb.JobStatistics2) *krm.JobStatistics2 {
	if in == nil {
		return nil
	}
	out := &krm.JobStatistics2{}
	out.QueryPlan = direct.Slice_FromProto(mapCtx, in.QueryPlan, ExplainQueryStage_FromProto)
	out.EstimatedBytesProcessed = direct.LazyPtr(in.GetEstimatedBytesProcessed())
	out.Timeline = direct.Slice_FromProto(mapCtx, in.Timeline, QueryTimelineSample_FromProto)
	out.TotalPartitionsProcessed = direct.LazyPtr(in.GetTotalPartitionsProcessed())
	out.TotalBytesProcessed = direct.LazyPtr(in.GetTotalBytesProcessed())
	out.TotalBytesProcessedAccuracy = direct.LazyPtr(in.GetTotalBytesProcessedAccuracy())
	out.TotalBytesBilled = direct.LazyPtr(in.GetTotalBytesBilled())
	out.BillingTier = direct.LazyPtr(in.GetBillingTier())
	out.TotalSlotMs = direct.LazyPtr(in.GetTotalSlotMs())
	out.CacheHit = direct.LazyPtr(in.GetCacheHit())
	out.ReferencedTables = direct.Slice_FromProto(mapCtx, in.ReferencedTables, TableReference_FromProto)
	out.ReferencedRoutines = direct.Slice_FromProto(mapCtx, in.ReferencedRoutines, RoutineReference_FromProto)
	out.Schema = TableSchema_FromProto(mapCtx, in.GetSchema())
	out.NumDmlAffectedRows = direct.LazyPtr(in.GetNumDmlAffectedRows())
	out.DmlStats = DmlStats_FromProto(mapCtx, in.GetDmlStats())
	out.UndeclaredQueryParameters = direct.Slice_FromProto(mapCtx, in.UndeclaredQueryParameters, QueryParameter_FromProto)
	out.StatementType = direct.LazyPtr(in.GetStatementType())
	out.DdlOperationPerformed = direct.LazyPtr(in.GetDdlOperationPerformed())
	out.DdlTargetTable = TableReference_FromProto(mapCtx, in.GetDdlTargetTable())
	out.DdlDestinationTable = TableReference_FromProto(mapCtx, in.GetDdlDestinationTable())
	out.DdlTargetRowAccessPolicy = RowAccessPolicyReference_FromProto(mapCtx, in.GetDdlTargetRowAccessPolicy())
	out.DdlAffectedRowAccessPolicyCount = direct.LazyPtr(in.GetDdlAffectedRowAccessPolicyCount())
	out.DdlTargetRoutine = RoutineReference_FromProto(mapCtx, in.GetDdlTargetRoutine())
	out.DdlTargetDataset = DatasetReference_FromProto(mapCtx, in.GetDdlTargetDataset())
	out.MlStatistics = MlStatistics_FromProto(mapCtx, in.GetMlStatistics())
	out.ExportDataStatistics = ExportDataStatistics_FromProto(mapCtx, in.GetExportDataStatistics())
	out.ExternalServiceCosts = direct.Slice_FromProto(mapCtx, in.ExternalServiceCosts, ExternalServiceCost_FromProto)
	out.BiEngineStatistics = BiEngineStatistics_FromProto(mapCtx, in.GetBiEngineStatistics())
	out.LoadQueryStatistics = LoadQueryStatistics_FromProto(mapCtx, in.GetLoadQueryStatistics())
	out.DclTargetTable = TableReference_FromProto(mapCtx, in.GetDclTargetTable())
	out.DclTargetView = TableReference_FromProto(mapCtx, in.GetDclTargetView())
	out.DclTargetDataset = DatasetReference_FromProto(mapCtx, in.GetDclTargetDataset())
	out.SearchStatistics = SearchStatistics_FromProto(mapCtx, in.GetSearchStatistics())
	out.VectorSearchStatistics = VectorSearchStatistics_FromProto(mapCtx, in.GetVectorSearchStatistics())
	out.PerformanceInsights = PerformanceInsights_FromProto(mapCtx, in.GetPerformanceInsights())
	out.QueryInfo = QueryInfo_FromProto(mapCtx, in.GetQueryInfo())
	out.SparkStatistics = SparkStatistics_FromProto(mapCtx, in.GetSparkStatistics())
	out.TransferredBytes = direct.LazyPtr(in.GetTransferredBytes())
	out.MaterializedViewStatistics = MaterializedViewStatistics_FromProto(mapCtx, in.GetMaterializedViewStatistics())
	out.MetadataCacheStatistics = MetadataCacheStatistics_FromProto(mapCtx, in.GetMetadataCacheStatistics())
	return out
}
func JobStatistics2_ToProto(mapCtx *direct.MapContext, in *krm.JobStatistics2) *pb.JobStatistics2 {
	if in == nil {
		return nil
	}
	out := &pb.JobStatistics2{}
	out.QueryPlan = direct.Slice_ToProto(mapCtx, in.QueryPlan, ExplainQueryStage_ToProto)
	out.EstimatedBytesProcessed = in.EstimatedBytesProcessed
	out.Timeline = direct.Slice_ToProto(mapCtx, in.Timeline, QueryTimelineSample_ToProto)
	out.TotalPartitionsProcessed = in.TotalPartitionsProcessed
	out.TotalBytesProcessed = in.TotalBytesProcessed
	out.TotalBytesProcessedAccuracy = in.TotalBytesProcessedAccuracy
	out.TotalBytesBilled = in.TotalBytesBilled
	out.BillingTier = in.BillingTier
	out.TotalSlotMs = in.TotalSlotMs
	out.CacheHit = in.CacheHit
	out.ReferencedTables = direct.Slice_ToProto(mapCtx, in.ReferencedTables, TableReference_ToProto)
	out.ReferencedRoutines = direct.Slice_ToProto(mapCtx, in.ReferencedRoutines, RoutineReference_ToProto)
	out.Schema = TableSchema_ToProto(mapCtx, in.Schema)
	out.NumDmlAffectedRows = in.NumDmlAffectedRows
	out.DmlStats = DmlStats_ToProto(mapCtx, in.DmlStats)
	out.UndeclaredQueryParameters = direct.Slice_ToProto(mapCtx, in.UndeclaredQueryParameters, QueryParameter_ToProto)
	out.StatementType = in.StatementType
	out.DdlOperationPerformed = in.DdlOperationPerformed
	out.DdlTargetTable = TableReference_ToProto(mapCtx, in.DdlTargetTable)
	out.DdlDestinationTable = TableReference_ToProto(mapCtx, in.DdlDestinationTable)
	out.DdlTargetRowAccessPolicy = RowAccessPolicyReference_ToProto(mapCtx, in.DdlTargetRowAccessPolicy)
	out.DdlAffectedRowAccessPolicyCount = in.DdlAffectedRowAccessPolicyCount
	out.DdlTargetRoutine = RoutineReference_ToProto(mapCtx, in.DdlTargetRoutine)
	out.DdlTargetDataset = DatasetReference_ToProto(mapCtx, in.DdlTargetDataset)
	out.MlStatistics = MlStatistics_ToProto(mapCtx, in.MlStatistics)
	out.ExportDataStatistics = ExportDataStatistics_ToProto(mapCtx, in.ExportDataStatistics)
	out.ExternalServiceCosts = direct.Slice_ToProto(mapCtx, in.ExternalServiceCosts, ExternalServiceCost_ToProto)
	out.BiEngineStatistics = BiEngineStatistics_ToProto(mapCtx, in.BiEngineStatistics)
	out.LoadQueryStatistics = LoadQueryStatistics_ToProto(mapCtx, in.LoadQueryStatistics)
	out.DclTargetTable = TableReference_ToProto(mapCtx, in.DclTargetTable)
	out.DclTargetView = TableReference_ToProto(mapCtx, in.DclTargetView)
	out.DclTargetDataset = DatasetReference_ToProto(mapCtx, in.DclTargetDataset)
	out.SearchStatistics = SearchStatistics_ToProto(mapCtx, in.SearchStatistics)
	out.VectorSearchStatistics = VectorSearchStatistics_ToProto(mapCtx, in.VectorSearchStatistics)
	out.PerformanceInsights = PerformanceInsights_ToProto(mapCtx, in.PerformanceInsights)
	out.QueryInfo = QueryInfo_ToProto(mapCtx, in.QueryInfo)
	out.SparkStatistics = SparkStatistics_ToProto(mapCtx, in.SparkStatistics)
	out.TransferredBytes = in.TransferredBytes
	out.MaterializedViewStatistics = MaterializedViewStatistics_ToProto(mapCtx, in.MaterializedViewStatistics)
	out.MetadataCacheStatistics = MetadataCacheStatistics_ToProto(mapCtx, in.MetadataCacheStatistics)
	return out
}
func JobStatistics3_FromProto(mapCtx *direct.MapContext, in *pb.JobStatistics3) *krm.JobStatistics3 {
	if in == nil {
		return nil
	}
	out := &krm.JobStatistics3{}
	out.InputFiles = direct.LazyPtr((in.GetInputFiles()))
	out.InputFileBytes = direct.LazyPtr(in.GetInputFileBytes())
	out.OutputRows = direct.LazyPtr(in.GetOutputRows())
	out.OutputBytes = direct.LazyPtr(in.GetOutputBytes())
	out.BadRecords = direct.LazyPtr(in.GetBadRecords())
	out.Timeline = direct.Slice_FromProto(mapCtx, in.Timeline, QueryTimelineSample_FromProto)
	return out
}
func JobStatistics3_ToProto(mapCtx *direct.MapContext, in *krm.JobStatistics3) *pb.JobStatistics3 {
	if in == nil {
		return nil
	}
	out := &pb.JobStatistics3{}
	out.InputFiles = in.InputFiles
	out.InputFileBytes = in.InputFileBytes
	out.OutputRows = in.OutputRows
	out.OutputBytes = in.OutputBytes
	out.BadRecords = in.BadRecords
	out.Timeline = direct.Slice_ToProto(mapCtx, in.Timeline, QueryTimelineSample_ToProto)
	return out
}
func JobStatistics4_FromProto(mapCtx *direct.MapContext, in *pb.JobStatistics4) *krm.JobStatistics4 {
	if in == nil {
		return nil
	}
	out := &krm.JobStatistics4{}
	out.DestinationUriFileCounts = in.DestinationUriFileCounts
	out.InputBytes = direct.LazyPtr(in.GetInputBytes())
	out.Timeline = direct.Slice_FromProto(mapCtx, in.Timeline, QueryTimelineSample_FromProto)
	return out
}
func JobStatistics4_ToProto(mapCtx *direct.MapContext, in *krm.JobStatistics4) *pb.JobStatistics4 {
	if in == nil {
		return nil
	}
	out := &pb.JobStatistics4{}
	out.DestinationUriFileCounts = in.DestinationUriFileCounts
	out.InputBytes = in.InputBytes
	out.Timeline = direct.Slice_ToProto(mapCtx, in.Timeline, QueryTimelineSample_ToProto)
	return out
}
func JobStatistics_TransactionInfo_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo) *krm.JobStatistics_TransactionInfo {
	if in == nil {
		return nil
	}
	out := &krm.JobStatistics_TransactionInfo{}
	out.TransactionID = direct.LazyPtr(in.GetTransactionId())
	return out
}
func JobStatistics_TransactionInfo_ToProto(mapCtx *direct.MapContext, in *krm.JobStatistics_TransactionInfo) *pb.TransactionInfo {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo{}
	out.TransactionId = in.TransactionID
	return out
}
func JobStatus_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus) *krm.JobStatus {
	if in == nil {
		return nil
	}
	out := &krm.JobStatus{}
	out.ErrorResult = ErrorProto_FromProto(mapCtx, in.GetErrorResult())
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, ErrorProto_FromProto)
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func JobStatus_ToProto(mapCtx *direct.MapContext, in *krm.JobStatus) *pb.JobStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus{}
	out.ErrorResult = ErrorProto_ToProto(mapCtx, in.ErrorResult)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, ErrorProto_ToProto)
	out.State = in.State
	return out
}
func JoinRestrictionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.JoinRestrictionPolicy) *krm.JoinRestrictionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.JoinRestrictionPolicy{}
	out.JoinCondition = direct.LazyPtr(in.GetJoinCondition())
	out.JoinAllowedColumns = in.JoinAllowedColumns
	return out
}
func JoinRestrictionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.JoinRestrictionPolicy) *pb.JoinRestrictionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.JoinRestrictionPolicy{}
	out.JoinCondition = in.JoinCondition
	out.JoinAllowedColumns = in.JoinAllowedColumns
	return out
}
func JsonOptions_FromProto(mapCtx *direct.MapContext, in *pb.JsonOptions) *krm.JsonOptions {
	if in == nil {
		return nil
	}
	out := &krm.JsonOptions{}
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	return out
}
func JsonOptions_ToProto(mapCtx *direct.MapContext, in *krm.JsonOptions) *pb.JsonOptions {
	if in == nil {
		return nil
	}
	out := &pb.JsonOptions{}
	out.Encoding = in.Encoding
	return out
}
func LinkedDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.LinkedDatasetMetadata) *krm.LinkedDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.LinkedDatasetMetadata{}
	return out
}
func LinkedDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.LinkedDatasetMetadata) *pb.LinkedDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.LinkedDatasetMetadata{}
	return out
}
func LinkedDatasetSource_FromProto(mapCtx *direct.MapContext, in *pb.LinkedDatasetSource) *krm.LinkedDatasetSource {
	if in == nil {
		return nil
	}
	out := &krm.LinkedDatasetSource{}
	out.SourceDataset = DatasetReference_FromProto(mapCtx, in.GetSourceDataset())
	return out
}
func LinkedDatasetSource_ToProto(mapCtx *direct.MapContext, in *krm.LinkedDatasetSource) *pb.LinkedDatasetSource {
	if in == nil {
		return nil
	}
	out := &pb.LinkedDatasetSource{}
	out.SourceDataset = DatasetReference_ToProto(mapCtx, in.SourceDataset)
	return out
}
func ListFormatDataset_FromProto(mapCtx *direct.MapContext, in *pb.DatasetListDatasets) *krm.ListFormatDataset {
	if in == nil {
		return nil
	}
	out := &krm.ListFormatDataset{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.ID = direct.LazyPtr(in.GetId())
	out.DatasetReference = DatasetReference_FromProto(mapCtx, in.GetDatasetReference())
	out.Labels = in.Labels
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func ListFormatDataset_ToProto(mapCtx *direct.MapContext, in *krm.ListFormatDataset) *pb.DatasetListDatasets {
	if in == nil {
		return nil
	}
	out := &pb.DatasetListDatasets{}
	out.Kind = in.Kind
	out.Id = in.ID
	out.DatasetReference = DatasetReference_ToProto(mapCtx, in.DatasetReference)
	out.Labels = in.Labels
	out.FriendlyName = in.FriendlyName
	out.Location = in.Location
	return out
}
func ListFormatJob_FromProto(mapCtx *direct.MapContext, in *pb.JobListJobs) *krm.ListFormatJob {
	if in == nil {
		return nil
	}
	out := &krm.ListFormatJob{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Kind = direct.LazyPtr(in.GetKind())
	out.JobReference = JobReference_FromProto(mapCtx, in.GetJobReference())
	out.State = direct.LazyPtr(in.GetState())
	out.ErrorResult = ErrorProto_FromProto(mapCtx, in.GetErrorResult())
	out.Statistics = JobStatistics_FromProto(mapCtx, in.GetStatistics())
	out.Configuration = JobConfiguration_FromProto(mapCtx, in.GetConfiguration())
	out.Status = JobStatus_FromProto(mapCtx, in.GetStatus())
	out.UserEmail = direct.LazyPtr(in.GetUserEmail())
	out.PrincipalSubject = direct.LazyPtr(in.GetPrincipalSubject())
	return out
}
func ListFormatJob_ToProto(mapCtx *direct.MapContext, in *krm.ListFormatJob) *pb.JobListJobs {
	if in == nil {
		return nil
	}
	out := &pb.JobListJobs{}
	out.Id = in.ID
	out.Kind = in.Kind
	out.JobReference = JobReference_ToProto(mapCtx, in.JobReference)
	out.State = in.State
	out.ErrorResult = ErrorProto_ToProto(mapCtx, in.ErrorResult)
	out.Statistics = JobStatistics_ToProto(mapCtx, in.Statistics)
	out.Configuration = JobConfiguration_ToProto(mapCtx, in.Configuration)
	out.Status = JobStatus_ToProto(mapCtx, in.Status)
	out.UserEmail = in.UserEmail
	out.PrincipalSubject = in.PrincipalSubject
	return out
}
func ListFormatTable_FromProto(mapCtx *direct.MapContext, in *pb.TableListTables) *krm.ListFormatTable {
	if in == nil {
		return nil
	}
	out := &krm.ListFormatTable{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.ID = direct.LazyPtr(in.GetId())
	out.TableReference = TableReference_FromProto(mapCtx, in.GetTableReference())
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Type = direct.LazyPtr(in.GetType())
	out.TimePartitioning = TimePartitioning_FromProto(mapCtx, in.GetTimePartitioning())
	out.RangePartitioning = RangePartitioning_FromProto(mapCtx, in.GetRangePartitioning())
	out.Clustering = Clustering_FromProto(mapCtx, in.GetClustering())
	out.Labels = in.Labels
	out.View = ListFormatView_FromProto(mapCtx, in.GetView())
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.ExpirationTime = direct.LazyPtr(in.GetExpirationTime())
	out.RequirePartitionFilter = direct.LazyPtr(in.GetRequirePartitionFilter())
	return out
}
func ListFormatTable_ToProto(mapCtx *direct.MapContext, in *krm.ListFormatTable) *pb.TableListTables {
	if in == nil {
		return nil
	}
	out := &pb.TableListTables{}
	out.Kind = in.Kind
	out.Id = in.ID
	out.TableReference = TableReference_ToProto(mapCtx, in.TableReference)
	out.FriendlyName = in.FriendlyName
	out.Type = in.Type
	out.TimePartitioning = TimePartitioning_ToProto(mapCtx, in.TimePartitioning)
	out.RangePartitioning = RangePartitioning_ToProto(mapCtx, in.RangePartitioning)
	out.Clustering = Clustering_ToProto(mapCtx, in.Clustering)
	out.Labels = in.Labels
	out.View = ListFormatView_ToProto(mapCtx, in.View)
	out.CreationTime = in.CreationTime
	out.ExpirationTime = in.ExpirationTime
	out.RequirePartitionFilter = in.RequirePartitionFilter
	return out
}
func ListFormatView_FromProto(mapCtx *direct.MapContext, in *pb.View) *krm.ListFormatView {
	if in == nil {
		return nil
	}
	out := &krm.ListFormatView{}
	out.UseLegacySql = direct.LazyPtr(in.GetUseLegacySql())
	out.PrivacyPolicy = PrivacyPolicy_FromProto(mapCtx, in.GetPrivacyPolicy())
	return out
}
func ListFormatView_ToProto(mapCtx *direct.MapContext, in *krm.ListFormatView) *pb.View {
	if in == nil {
		return nil
	}
	out := &pb.View{}
	out.UseLegacySql = in.UseLegacySql
	out.PrivacyPolicy = PrivacyPolicy_ToProto(mapCtx, in.PrivacyPolicy)
	return out
}
func LoadQueryStatistics_FromProto(mapCtx *direct.MapContext, in *pb.LoadQueryStatistics) *krm.LoadQueryStatistics {
	if in == nil {
		return nil
	}
	out := &krm.LoadQueryStatistics{}
	out.InputFiles = direct.LazyPtr(in.GetInputFiles())
	out.InputFileBytes = direct.LazyPtr(in.GetInputFileBytes())
	out.OutputRows = direct.LazyPtr(in.GetOutputRows())
	out.OutputBytes = direct.LazyPtr(in.GetOutputBytes())
	out.BadRecords = direct.LazyPtr(in.GetBadRecords())
	return out
}
func LoadQueryStatistics_ToProto(mapCtx *direct.MapContext, in *krm.LoadQueryStatistics) *pb.LoadQueryStatistics {
	if in == nil {
		return nil
	}
	out := &pb.LoadQueryStatistics{}
	out.InputFiles = in.InputFiles
	out.InputFileBytes = in.InputFileBytes
	out.OutputRows = in.OutputRows
	out.OutputBytes = in.OutputBytes
	out.BadRecords = in.BadRecords
	return out
}
func MaterializedView_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krm.MaterializedView {
	if in == nil {
		return nil
	}
	out := &krm.MaterializedView{}
	out.TableReference = TableReference_FromProto(mapCtx, in.GetTableReference())
	out.Chosen = in.Chosen
	out.EstimatedBytesSaved = in.EstimatedBytesSaved
	out.RejectedReason = direct.LazyPtr(in.GetRejectedReason())
	return out
}
func MaterializedView_ToProto(mapCtx *direct.MapContext, in *krm.MaterializedView) *pb.MaterializedView {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedView{}
	if oneof := TableReference_ToProto(mapCtx, in.TableReference); oneof != nil {
		out.TableReference = oneof
	}
	out.Chosen = in.Chosen
	out.EstimatedBytesSaved = in.EstimatedBytesSaved
	out.RejectedReason = in.RejectedReason
	return out
}
func MaterializedViewDefinition_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedViewDefinition) *krm.MaterializedViewDefinition {
	if in == nil {
		return nil
	}
	out := &krm.MaterializedViewDefinition{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.LastRefreshTime = direct.LazyPtr(in.GetLastRefreshTime())
	out.EnableRefresh = direct.LazyPtr(in.GetEnableRefresh())
	out.RefreshIntervalMs = direct.LazyPtr(uint64(in.GetRefreshIntervalMs()))
	out.AllowNonIncrementalDefinition = direct.LazyPtr(in.GetAllowNonIncrementalDefinition())
	return out
}
func MaterializedViewDefinition_ToProto(mapCtx *direct.MapContext, in *krm.MaterializedViewDefinition) *pb.MaterializedViewDefinition {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedViewDefinition{}
	out.Query = in.Query
	out.LastRefreshTime = in.LastRefreshTime
	out.EnableRefresh = in.EnableRefresh
	out.RefreshIntervalMs = direct.LazyPtr(int64(*in.RefreshIntervalMs))
	out.AllowNonIncrementalDefinition = in.AllowNonIncrementalDefinition
	return out
}
func MaterializedViewStatistics_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedViewStatistics) *krm.MaterializedViewStatistics {
	if in == nil {
		return nil
	}
	out := &krm.MaterializedViewStatistics{}
	out.MaterializedView = direct.Slice_FromProto(mapCtx, in.MaterializedView, MaterializedView_FromProto)
	return out
}
func MaterializedViewStatistics_ToProto(mapCtx *direct.MapContext, in *krm.MaterializedViewStatistics) *pb.MaterializedViewStatistics {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedViewStatistics{}
	out.MaterializedView = direct.Slice_ToProto(mapCtx, in.MaterializedView, MaterializedView_ToProto)
	return out
}
func MaterializedViewStatus_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedViewStatus) *krm.MaterializedViewStatus {
	if in == nil {
		return nil
	}
	out := &krm.MaterializedViewStatus{}
	out.RefreshWatermark = direct.StringTimestamp_FromProto(mapCtx, in.GetRefreshWatermark())
	out.LastRefreshStatus = ErrorProto_FromProto(mapCtx, in.GetLastRefreshStatus())
	return out
}
func MaterializedViewStatus_ToProto(mapCtx *direct.MapContext, in *krm.MaterializedViewStatus) *pb.MaterializedViewStatus {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedViewStatus{}
	out.RefreshWatermark = direct.StringTimestamp_ToProto(mapCtx, in.RefreshWatermark)
	out.LastRefreshStatus = ErrorProto_ToProto(mapCtx, in.LastRefreshStatus)
	return out
}
func MetadataCacheStatistics_FromProto(mapCtx *direct.MapContext, in *pb.MetadataCacheStatistics) *krm.MetadataCacheStatistics {
	if in == nil {
		return nil
	}
	out := &krm.MetadataCacheStatistics{}
	out.TableMetadataCacheUsage = direct.Slice_FromProto(mapCtx, in.TableMetadataCacheUsage, TableMetadataCacheUsage_FromProto)
	return out
}
func MetadataCacheStatistics_ToProto(mapCtx *direct.MapContext, in *krm.MetadataCacheStatistics) *pb.MetadataCacheStatistics {
	if in == nil {
		return nil
	}
	out := &pb.MetadataCacheStatistics{}
	out.TableMetadataCacheUsage = direct.Slice_ToProto(mapCtx, in.TableMetadataCacheUsage, TableMetadataCacheUsage_ToProto)
	return out
}
func MlStatistics_FromProto(mapCtx *direct.MapContext, in *pb.MlStatistics) *krm.MlStatistics {
	if in == nil {
		return nil
	}
	out := &krm.MlStatistics{}
	out.MaxIterations = direct.LazyPtr(in.GetMaxIterations())
	out.IterationResults = direct.Slice_FromProto(mapCtx, in.IterationResults, Model_TrainingRun_IterationResult_FromProto)
	out.ModelType = direct.LazyPtr(in.GetModelType())
	out.TrainingType = direct.LazyPtr(in.GetTrainingType())
	out.HparamTrials = direct.Slice_FromProto(mapCtx, in.HparamTrials, Model_HparamTuningTrial_FromProto)
	return out
}
func MlStatistics_ToProto(mapCtx *direct.MapContext, in *krm.MlStatistics) *pb.MlStatistics {
	if in == nil {
		return nil
	}
	out := &pb.MlStatistics{}
	out.MaxIterations = in.MaxIterations
	out.IterationResults = direct.Slice_ToProto(mapCtx, in.IterationResults, Model_TrainingRun_IterationResult_ToProto)
	out.ModelType = in.ModelType
	out.TrainingType = in.TrainingType
	out.HparamTrials = direct.Slice_ToProto(mapCtx, in.HparamTrials, Model_HparamTuningTrial_ToProto)
	return out
}
func Model_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.Model {
	if in == nil {
		return nil
	}
	out := &krm.Model{}
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ModelReference = ModelReference_FromProto(mapCtx, in.GetModelReference())
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.LastModifiedTime = direct.LazyPtr(in.GetLastModifiedTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Labels = in.Labels
	out.ExpirationTime = direct.LazyPtr(in.GetExpirationTime())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.EncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetEncryptionConfiguration())
	out.ModelType = direct.LazyPtr(in.GetModelType())
	out.TrainingRuns = direct.Slice_FromProto(mapCtx, in.TrainingRuns, Model_TrainingRun_FromProto)
	out.FeatureColumns = direct.Slice_FromProto(mapCtx, in.FeatureColumns, StandardSqlField_FromProto)
	out.LabelColumns = direct.Slice_FromProto(mapCtx, in.LabelColumns, StandardSqlField_FromProto)
	out.TransformColumns = direct.Slice_FromProto(mapCtx, in.TransformColumns, TransformColumn_FromProto)
	out.HparamSearchSpaces = Model_HparamSearchSpaces_FromProto(mapCtx, in.GetHparamSearchSpaces())
	out.DefaultTrialID = direct.LazyPtr(in.GetDefaultTrialId())
	out.HparamTrials = direct.Slice_FromProto(mapCtx, in.HparamTrials, Model_HparamTuningTrial_FromProto)
	out.OptimalTrialIds = in.OptimalTrialIds
	out.RemoteModelInfo = RemoteModelInfo_FromProto(mapCtx, in.GetRemoteModelInfo())
	return out
}
func Model_ToProto(mapCtx *direct.MapContext, in *krm.Model) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.Etag = in.Etag
	out.ModelReference = ModelReference_ToProto(mapCtx, in.ModelReference)
	out.CreationTime = in.CreationTime
	out.LastModifiedTime = in.LastModifiedTime
	out.Description = in.Description
	out.FriendlyName = in.FriendlyName
	out.Labels = in.Labels
	out.ExpirationTime = in.ExpirationTime
	out.Location = in.Location
	out.EncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.EncryptionConfiguration)
	out.ModelType = in.ModelType
	out.TrainingRuns = direct.Slice_ToProto(mapCtx, in.TrainingRuns, Model_TrainingRun_ToProto)
	out.FeatureColumns = direct.Slice_ToProto(mapCtx, in.FeatureColumns, StandardSqlField_ToProto)
	out.LabelColumns = direct.Slice_ToProto(mapCtx, in.LabelColumns, StandardSqlField_ToProto)
	out.TransformColumns = direct.Slice_ToProto(mapCtx, in.TransformColumns, TransformColumn_ToProto)
	out.HparamSearchSpaces = Model_HparamSearchSpaces_ToProto(mapCtx, in.HparamSearchSpaces)
	out.DefaultTrialId = in.DefaultTrialID
	out.HparamTrials = direct.Slice_ToProto(mapCtx, in.HparamTrials, Model_HparamTuningTrial_ToProto)
	out.OptimalTrialIds = in.OptimalTrialIds
	out.RemoteModelInfo = RemoteModelInfo_ToProto(mapCtx, in.RemoteModelInfo)
	return out
}
func ModelReference_FromProto(mapCtx *direct.MapContext, in *pb.ModelReference) *krm.ModelReference {
	if in == nil {
		return nil
	}
	out := &krm.ModelReference{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	out.ModelID = direct.LazyPtr(in.GetModelId())
	return out
}
func ModelReference_ToProto(mapCtx *direct.MapContext, in *krm.ModelReference) *pb.ModelReference {
	if in == nil {
		return nil
	}
	out := &pb.ModelReference{}
	out.ProjectId = in.ProjectID
	out.DatasetId = in.DatasetID
	out.ModelId = in.ModelID
	return out
}
func Model_AggregateClassificationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.AggregateClassificationMetrics) *krm.Model_AggregateClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_AggregateClassificationMetrics{}
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.Accuracy = direct.LazyPtr(in.GetAccuracy())
	out.Threshold = direct.LazyPtr(in.GetThreshold())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	out.LogLoss = direct.LazyPtr(in.GetLogLoss())
	out.RocAuc = direct.LazyPtr(in.GetRocAuc())
	return out
}
func Model_AggregateClassificationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_AggregateClassificationMetrics) *pb.AggregateClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.AggregateClassificationMetrics{}
	out.Precision = in.Precision
	out.Recall = in.Recall
	out.Accuracy = in.Accuracy
	out.Threshold = in.Threshold
	out.F1Score = in.F1Score
	out.LogLoss = in.LogLoss
	out.RocAuc = in.RocAuc
	return out
}
func Model_ArimaFittingMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ArimaFittingMetrics) *krm.Model_ArimaFittingMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_ArimaFittingMetrics{}
	out.LogLikelihood = direct.LazyPtr(in.GetLogLikelihood())
	out.Aic = direct.LazyPtr(in.GetAic())
	out.Variance = direct.LazyPtr(in.GetVariance())
	return out
}
func Model_ArimaFittingMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_ArimaFittingMetrics) *pb.ArimaFittingMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ArimaFittingMetrics{}
	out.LogLikelihood = in.LogLikelihood
	out.Aic = in.Aic
	out.Variance = in.Variance
	return out
}
func Model_ArimaForecastingMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ArimaForecastingMetrics) *krm.Model_ArimaForecastingMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_ArimaForecastingMetrics{}
	out.ArimaSingleModelForecastingMetrics = direct.Slice_FromProto(mapCtx, in.ArimaSingleModelForecastingMetrics, Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics_FromProto)
	return out
}
func Model_ArimaForecastingMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_ArimaForecastingMetrics) *pb.ArimaForecastingMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ArimaForecastingMetrics{}
	out.ArimaSingleModelForecastingMetrics = direct.Slice_ToProto(mapCtx, in.ArimaSingleModelForecastingMetrics, Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics_ToProto)
	return out
}
func Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ArimaSingleModelForecastingMetrics) *krm.Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics{}
	out.NonSeasonalOrder = Model_ArimaOrder_FromProto(mapCtx, in.GetNonSeasonalOrder())
	out.ArimaFittingMetrics = Model_ArimaFittingMetrics_FromProto(mapCtx, in.GetArimaFittingMetrics())
	out.HasDrift = direct.LazyPtr(in.GetHasDrift())
	out.TimeSeriesID = direct.LazyPtr(in.GetTimeSeriesId())
	out.TimeSeriesIds = in.TimeSeriesIds
	out.SeasonalPeriods = in.SeasonalPeriods
	out.HasHolidayEffect = direct.LazyPtr(in.GetHasHolidayEffect())
	out.HasSpikesAndDips = direct.LazyPtr(in.GetHasSpikesAndDips())
	out.HasStepChanges = direct.LazyPtr(in.GetHasStepChanges())
	return out
}
func Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_ArimaForecastingMetrics_ArimaSingleModelForecastingMetrics) *pb.ArimaSingleModelForecastingMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ArimaSingleModelForecastingMetrics{}
	out.NonSeasonalOrder = Model_ArimaOrder_ToProto(mapCtx, in.NonSeasonalOrder)
	out.ArimaFittingMetrics = Model_ArimaFittingMetrics_ToProto(mapCtx, in.ArimaFittingMetrics)
	out.HasDrift = in.HasDrift
	out.TimeSeriesId = in.TimeSeriesID
	out.TimeSeriesIds = in.TimeSeriesIds
	out.SeasonalPeriods = in.SeasonalPeriods
	out.HasHolidayEffect = in.HasHolidayEffect
	out.HasSpikesAndDips = in.HasSpikesAndDips
	out.HasStepChanges = in.HasStepChanges
	return out
}
func Model_ArimaOrder_FromProto(mapCtx *direct.MapContext, in *pb.ArimaOrder) *krm.Model_ArimaOrder {
	if in == nil {
		return nil
	}
	out := &krm.Model_ArimaOrder{}
	out.P = direct.LazyPtr(in.GetP())
	out.D = direct.LazyPtr(in.GetD())
	out.Q = direct.LazyPtr(in.GetQ())
	return out
}
func Model_ArimaOrder_ToProto(mapCtx *direct.MapContext, in *krm.Model_ArimaOrder) *pb.ArimaOrder {
	if in == nil {
		return nil
	}
	out := &pb.ArimaOrder{}
	out.P = in.P
	out.D = in.D
	out.Q = in.Q
	return out
}
func Model_BinaryClassificationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.BinaryClassificationMetrics) *krm.Model_BinaryClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_BinaryClassificationMetrics{}
	out.AggregateClassificationMetrics = Model_AggregateClassificationMetrics_FromProto(mapCtx, in.GetAggregateClassificationMetrics())
	out.BinaryConfusionMatrixList = direct.Slice_FromProto(mapCtx, in.BinaryConfusionMatrixList, Model_BinaryClassificationMetrics_BinaryConfusionMatrix_FromProto)
	out.PositiveLabel = direct.LazyPtr(in.GetPositiveLabel())
	out.NegativeLabel = direct.LazyPtr(in.GetNegativeLabel())
	return out
}
func Model_BinaryClassificationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_BinaryClassificationMetrics) *pb.BinaryClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.BinaryClassificationMetrics{}
	out.AggregateClassificationMetrics = Model_AggregateClassificationMetrics_ToProto(mapCtx, in.AggregateClassificationMetrics)
	out.BinaryConfusionMatrixList = direct.Slice_ToProto(mapCtx, in.BinaryConfusionMatrixList, Model_BinaryClassificationMetrics_BinaryConfusionMatrix_ToProto)
	out.PositiveLabel = in.PositiveLabel
	out.NegativeLabel = in.NegativeLabel
	return out
}
func Model_BinaryClassificationMetrics_BinaryConfusionMatrix_FromProto(mapCtx *direct.MapContext, in *pb.BinaryConfusionMatrix) *krm.Model_BinaryClassificationMetrics_BinaryConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &krm.Model_BinaryClassificationMetrics_BinaryConfusionMatrix{}
	out.PositiveClassThreshold = direct.LazyPtr(in.GetPositiveClassThreshold())
	out.TruePositives = direct.LazyPtr(in.GetTruePositives())
	out.FalsePositives = direct.LazyPtr(in.GetFalsePositives())
	out.TrueNegatives = direct.LazyPtr(in.GetTrueNegatives())
	out.FalseNegatives = direct.LazyPtr(in.GetFalseNegatives())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Recall = direct.LazyPtr(in.GetRecall())
	out.F1Score = direct.LazyPtr(in.GetF1Score())
	out.Accuracy = direct.LazyPtr(in.GetAccuracy())
	return out
}
func Model_BinaryClassificationMetrics_BinaryConfusionMatrix_ToProto(mapCtx *direct.MapContext, in *krm.Model_BinaryClassificationMetrics_BinaryConfusionMatrix) *pb.BinaryConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &pb.BinaryConfusionMatrix{}
	out.PositiveClassThreshold = in.PositiveClassThreshold
	out.TruePositives = in.TruePositives
	out.FalsePositives = in.FalsePositives
	out.TrueNegatives = in.TrueNegatives
	out.FalseNegatives = in.FalseNegatives
	out.Precision = in.Precision
	out.Recall = in.Recall
	out.F1Score = in.F1Score
	out.Accuracy = in.Accuracy
	return out
}
func Model_ClusteringMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ClusteringMetrics) *krm.Model_ClusteringMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_ClusteringMetrics{}
	out.DaviesBouldinIndex = direct.LazyPtr(in.GetDaviesBouldinIndex())
	out.MeanSquaredDistance = direct.LazyPtr(in.GetMeanSquaredDistance())
	out.Clusters = direct.Slice_FromProto(mapCtx, in.Clusters, Model_ClusteringMetrics_Cluster_FromProto)
	return out
}
func Model_ClusteringMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_ClusteringMetrics) *pb.ClusteringMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ClusteringMetrics{}
	out.DaviesBouldinIndex = in.DaviesBouldinIndex
	out.MeanSquaredDistance = in.MeanSquaredDistance
	out.Clusters = direct.Slice_ToProto(mapCtx, in.Clusters, Model_ClusteringMetrics_Cluster_ToProto)
	return out
}
func Model_ClusteringMetrics_Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.Model_ClusteringMetrics_Cluster {
	if in == nil {
		return nil
	}
	out := &krm.Model_ClusteringMetrics_Cluster{}
	out.CentroidID = direct.LazyPtr(in.GetCentroidId())
	out.FeatureValues = direct.Slice_FromProto(mapCtx, in.FeatureValues, Model_ClusteringMetrics_Cluster_FeatureValue_FromProto)
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func Model_ClusteringMetrics_Cluster_ToProto(mapCtx *direct.MapContext, in *krm.Model_ClusteringMetrics_Cluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.CentroidId = in.CentroidID
	out.FeatureValues = direct.Slice_ToProto(mapCtx, in.FeatureValues, Model_ClusteringMetrics_Cluster_FeatureValue_ToProto)
	out.Count = in.Count
	return out
}
func Model_ClusteringMetrics_Cluster_FeatureValue_FromProto(mapCtx *direct.MapContext, in *pb.FeatureValue) *krm.Model_ClusteringMetrics_Cluster_FeatureValue {
	if in == nil {
		return nil
	}
	out := &krm.Model_ClusteringMetrics_Cluster_FeatureValue{}
	out.FeatureColumn = direct.LazyPtr(in.GetFeatureColumn())
	out.NumericalValue = direct.LazyPtr(in.GetNumericalValue())
	out.CategoricalValue = Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_FromProto(mapCtx, in.GetCategoricalValue())
	return out
}
func Model_ClusteringMetrics_Cluster_FeatureValue_ToProto(mapCtx *direct.MapContext, in *krm.Model_ClusteringMetrics_Cluster_FeatureValue) *pb.FeatureValue {
	if in == nil {
		return nil
	}
	out := &pb.FeatureValue{}
	out.FeatureColumn = in.FeatureColumn
	out.NumericalValue = in.NumericalValue
	out.CategoricalValue = Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_ToProto(mapCtx, in.CategoricalValue)
	return out
}
func Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_FromProto(mapCtx *direct.MapContext, in *pb.CategoricalValue) *krm.Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue {
	if in == nil {
		return nil
	}
	out := &krm.Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue{}
	out.CategoryCounts = direct.Slice_FromProto(mapCtx, in.CategoryCounts, Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount_FromProto)
	return out
}
func Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_ToProto(mapCtx *direct.MapContext, in *krm.Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue) *pb.CategoricalValue {
	if in == nil {
		return nil
	}
	out := &pb.CategoricalValue{}
	out.CategoryCounts = direct.Slice_ToProto(mapCtx, in.CategoryCounts, Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount_ToProto)
	return out
}
func Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount_FromProto(mapCtx *direct.MapContext, in *pb.CategoryCount) *krm.Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount {
	if in == nil {
		return nil
	}
	out := &krm.Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount{}
	out.Category = direct.LazyPtr(in.GetCategory())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount_ToProto(mapCtx *direct.MapContext, in *krm.Model_ClusteringMetrics_Cluster_FeatureValue_CategoricalValue_CategoryCount) *pb.CategoryCount {
	if in == nil {
		return nil
	}
	out := &pb.CategoryCount{}
	out.Category = in.Category
	out.Count = in.Count
	return out
}
func Model_DataSplitResult_FromProto(mapCtx *direct.MapContext, in *pb.DataSplitResult) *krm.Model_DataSplitResult {
	if in == nil {
		return nil
	}
	out := &krm.Model_DataSplitResult{}
	out.TrainingTable = TableReference_FromProto(mapCtx, in.GetTrainingTable())
	out.EvaluationTable = TableReference_FromProto(mapCtx, in.GetEvaluationTable())
	out.TestTable = TableReference_FromProto(mapCtx, in.GetTestTable())
	return out
}
func Model_DataSplitResult_ToProto(mapCtx *direct.MapContext, in *krm.Model_DataSplitResult) *pb.DataSplitResult {
	if in == nil {
		return nil
	}
	out := &pb.DataSplitResult{}
	out.TrainingTable = TableReference_ToProto(mapCtx, in.TrainingTable)
	out.EvaluationTable = TableReference_ToProto(mapCtx, in.EvaluationTable)
	out.TestTable = TableReference_ToProto(mapCtx, in.TestTable)
	return out
}
func Model_DimensionalityReductionMetrics_FromProto(mapCtx *direct.MapContext, in *pb.DimensionalityReductionMetrics) *krm.Model_DimensionalityReductionMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_DimensionalityReductionMetrics{}
	out.TotalExplainedVarianceRatio = direct.LazyPtr(in.GetTotalExplainedVarianceRatio())
	return out
}
func Model_DimensionalityReductionMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_DimensionalityReductionMetrics) *pb.DimensionalityReductionMetrics {
	if in == nil {
		return nil
	}
	out := &pb.DimensionalityReductionMetrics{}
	out.TotalExplainedVarianceRatio = in.TotalExplainedVarianceRatio
	return out
}
func Model_EvaluationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationMetrics) *krm.Model_EvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_EvaluationMetrics{}
	out.RegressionMetrics = Model_RegressionMetrics_FromProto(mapCtx, in.GetRegressionMetrics())
	out.BinaryClassificationMetrics = Model_BinaryClassificationMetrics_FromProto(mapCtx, in.GetBinaryClassificationMetrics())
	out.MultiClassClassificationMetrics = Model_MultiClassClassificationMetrics_FromProto(mapCtx, in.GetMultiClassClassificationMetrics())
	out.ClusteringMetrics = Model_ClusteringMetrics_FromProto(mapCtx, in.GetClusteringMetrics())
	out.RankingMetrics = Model_RankingMetrics_FromProto(mapCtx, in.GetRankingMetrics())
	out.ArimaForecastingMetrics = Model_ArimaForecastingMetrics_FromProto(mapCtx, in.GetArimaForecastingMetrics())
	out.DimensionalityReductionMetrics = Model_DimensionalityReductionMetrics_FromProto(mapCtx, in.GetDimensionalityReductionMetrics())
	return out
}
func Model_EvaluationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_EvaluationMetrics) *pb.EvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationMetrics{}
	if oneof := Model_RegressionMetrics_ToProto(mapCtx, in.RegressionMetrics); oneof != nil {
		out.RegressionMetrics = oneof
	}
	if oneof := Model_BinaryClassificationMetrics_ToProto(mapCtx, in.BinaryClassificationMetrics); oneof != nil {
		out.BinaryClassificationMetrics = oneof
	}
	if oneof := Model_MultiClassClassificationMetrics_ToProto(mapCtx, in.MultiClassClassificationMetrics); oneof != nil {
		out.MultiClassClassificationMetrics = oneof
	}
	if oneof := Model_ClusteringMetrics_ToProto(mapCtx, in.ClusteringMetrics); oneof != nil {
		out.ClusteringMetrics = oneof
	}
	if oneof := Model_RankingMetrics_ToProto(mapCtx, in.RankingMetrics); oneof != nil {
		out.RankingMetrics = oneof
	}
	if oneof := Model_ArimaForecastingMetrics_ToProto(mapCtx, in.ArimaForecastingMetrics); oneof != nil {
		out.ArimaForecastingMetrics = oneof
	}
	if oneof := Model_DimensionalityReductionMetrics_ToProto(mapCtx, in.DimensionalityReductionMetrics); oneof != nil {
		out.DimensionalityReductionMetrics = oneof
	}
	return out
}
func Model_float64HparamSearchSpace_float64Range_FromProto(mapCtx *direct.MapContext, in *pb.DoubleRange) *krm.Model_float64HparamSearchSpace_float64Range {
	if in == nil {
		return nil
	}
	out := &krm.Model_float64HparamSearchSpace_float64Range{}
	out.Max = in.Max
	out.Min = in.Min
	return out
}
func Model_float64HparamSearchSpace_float64Range_ToProto(mapCtx *direct.MapContext, in *krm.Model_float64HparamSearchSpace_float64Range) *pb.DoubleRange {
	if in == nil {
		return nil
	}
	out := &pb.DoubleRange{}
	out.Max = in.Max
	out.Min = in.Min
	return out
}
func Model_float64HparamSearchSpace_float64Candidates_FromProto(mapCtx *direct.MapContext, in *pb.DoubleCandidates) *krm.Model_float64HparamSearchSpace_float64Candidates {
	if in == nil {
		return nil
	}
	out := &krm.Model_float64HparamSearchSpace_float64Candidates{}
	out.Candidates = in.Candidates
	return out
}
func Model_float64HparamSearchSpace_float64Candidates_ToProto(mapCtx *direct.MapContext, in *krm.Model_float64HparamSearchSpace_float64Candidates) *pb.DoubleCandidates {
	if in == nil {
		return nil
	}
	out := &pb.DoubleCandidates{}
	out.Candidates = in.Candidates
	return out
}
func Model_float64HparamSearchSpace_FromProto(mapCtx *direct.MapContext, in *pb.DoubleHparamSearchSpace) *krm.Model_float64HparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &krm.Model_float64HparamSearchSpace{}
	out.Range = Model_float64HparamSearchSpace_float64Range_FromProto(mapCtx, in.GetRange())
	out.Candidates = Model_float64HparamSearchSpace_float64Candidates_FromProto(mapCtx, in.GetCandidates())
	return out
}
func Model_float64HparamSearchSpace_ToProto(mapCtx *direct.MapContext, in *krm.Model_float64HparamSearchSpace) *pb.DoubleHparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &pb.DoubleHparamSearchSpace{}
	out.Range = Model_float64HparamSearchSpace_float64Range_ToProto(mapCtx, in.Range)
	out.Candidates = Model_float64HparamSearchSpace_float64Candidates_ToProto(mapCtx, in.Candidates)
	return out
}
func Model_GlobalExplanation_FromProto(mapCtx *direct.MapContext, in *pb.GlobalExplanation) *krm.Model_GlobalExplanation {
	if in == nil {
		return nil
	}
	out := &krm.Model_GlobalExplanation{}
	out.Explanations = direct.Slice_FromProto(mapCtx, in.Explanations, Model_GlobalExplanation_Explanation_FromProto)
	out.ClassLabel = direct.LazyPtr(in.GetClassLabel())
	return out
}
func Model_GlobalExplanation_ToProto(mapCtx *direct.MapContext, in *krm.Model_GlobalExplanation) *pb.GlobalExplanation {
	if in == nil {
		return nil
	}
	out := &pb.GlobalExplanation{}
	out.Explanations = direct.Slice_ToProto(mapCtx, in.Explanations, Model_GlobalExplanation_Explanation_ToProto)
	out.ClassLabel = in.ClassLabel
	return out
}
func Model_GlobalExplanation_Explanation_FromProto(mapCtx *direct.MapContext, in *pb.Explanation) *krm.Model_GlobalExplanation_Explanation {
	if in == nil {
		return nil
	}
	out := &krm.Model_GlobalExplanation_Explanation{}
	out.FeatureName = direct.LazyPtr(in.GetFeatureName())
	out.Attribution = direct.LazyPtr(in.GetAttribution())
	return out
}
func Model_GlobalExplanation_Explanation_ToProto(mapCtx *direct.MapContext, in *krm.Model_GlobalExplanation_Explanation) *pb.Explanation {
	if in == nil {
		return nil
	}
	out := &pb.Explanation{}
	out.FeatureName = in.FeatureName
	out.Attribution = in.Attribution
	return out
}
func Model_HparamSearchSpaces_FromProto(mapCtx *direct.MapContext, in *pb.HparamSearchSpaces) *krm.Model_HparamSearchSpaces {
	if in == nil {
		return nil
	}
	out := &krm.Model_HparamSearchSpaces{}
	out.LearnRate = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetLearnRate())
	out.L1Reg = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetL1Reg())
	out.L2Reg = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetL2Reg())
	out.NumClusters = Model_IntHparamSearchSpace_FromProto(mapCtx, in.GetNumClusters())
	out.NumFactors = Model_IntHparamSearchSpace_FromProto(mapCtx, in.GetNumFactors())
	out.HiddenUnits = Model_IntArrayHparamSearchSpace_FromProto(mapCtx, in.GetHiddenUnits())
	out.BatchSize = Model_IntHparamSearchSpace_FromProto(mapCtx, in.GetBatchSize())
	out.Dropout = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetDropout())
	out.MaxTreeDepth = Model_IntHparamSearchSpace_FromProto(mapCtx, in.GetMaxTreeDepth())
	out.Subsample = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetSubsample())
	out.MinSplitLoss = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetMinSplitLoss())
	out.WalsAlpha = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetWalsAlpha())
	out.BoosterType = Model_StringHparamSearchSpace_FromProto(mapCtx, in.GetBoosterType())
	out.NumParallelTree = Model_IntHparamSearchSpace_FromProto(mapCtx, in.GetNumParallelTree())
	out.DartNormalizeType = Model_StringHparamSearchSpace_FromProto(mapCtx, in.GetDartNormalizeType())
	out.TreeMethod = Model_StringHparamSearchSpace_FromProto(mapCtx, in.GetTreeMethod())
	out.MinTreeChildWeight = Model_IntHparamSearchSpace_FromProto(mapCtx, in.GetMinTreeChildWeight())
	out.ColsampleBytree = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetColsampleBytree())
	out.ColsampleBylevel = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetColsampleBylevel())
	out.ColsampleBynode = Model_float64HparamSearchSpace_FromProto(mapCtx, in.GetColsampleBynode())
	out.ActivationFn = Model_StringHparamSearchSpace_FromProto(mapCtx, in.GetActivationFn())
	out.Optimizer = Model_StringHparamSearchSpace_FromProto(mapCtx, in.GetOptimizer())
	return out
}
func Model_HparamSearchSpaces_ToProto(mapCtx *direct.MapContext, in *krm.Model_HparamSearchSpaces) *pb.HparamSearchSpaces {
	if in == nil {
		return nil
	}
	out := &pb.HparamSearchSpaces{}
	out.LearnRate = Model_float64HparamSearchSpace_ToProto(mapCtx, in.LearnRate)
	out.L1Reg = Model_float64HparamSearchSpace_ToProto(mapCtx, in.L1Reg)
	out.L2Reg = Model_float64HparamSearchSpace_ToProto(mapCtx, in.L2Reg)
	out.NumClusters = Model_IntHparamSearchSpace_ToProto(mapCtx, in.NumClusters)
	out.NumFactors = Model_IntHparamSearchSpace_ToProto(mapCtx, in.NumFactors)
	out.HiddenUnits = Model_IntArrayHparamSearchSpace_ToProto(mapCtx, in.HiddenUnits)
	out.BatchSize = Model_IntHparamSearchSpace_ToProto(mapCtx, in.BatchSize)
	out.Dropout = Model_float64HparamSearchSpace_ToProto(mapCtx, in.Dropout)
	out.MaxTreeDepth = Model_IntHparamSearchSpace_ToProto(mapCtx, in.MaxTreeDepth)
	out.Subsample = Model_float64HparamSearchSpace_ToProto(mapCtx, in.Subsample)
	out.MinSplitLoss = Model_float64HparamSearchSpace_ToProto(mapCtx, in.MinSplitLoss)
	out.WalsAlpha = Model_float64HparamSearchSpace_ToProto(mapCtx, in.WalsAlpha)
	out.BoosterType = Model_StringHparamSearchSpace_ToProto(mapCtx, in.BoosterType)
	out.NumParallelTree = Model_IntHparamSearchSpace_ToProto(mapCtx, in.NumParallelTree)
	out.DartNormalizeType = Model_StringHparamSearchSpace_ToProto(mapCtx, in.DartNormalizeType)
	out.TreeMethod = Model_StringHparamSearchSpace_ToProto(mapCtx, in.TreeMethod)
	out.MinTreeChildWeight = Model_IntHparamSearchSpace_ToProto(mapCtx, in.MinTreeChildWeight)
	out.ColsampleBytree = Model_float64HparamSearchSpace_ToProto(mapCtx, in.ColsampleBytree)
	out.ColsampleBylevel = Model_float64HparamSearchSpace_ToProto(mapCtx, in.ColsampleBylevel)
	out.ColsampleBynode = Model_float64HparamSearchSpace_ToProto(mapCtx, in.ColsampleBynode)
	out.ActivationFn = Model_StringHparamSearchSpace_ToProto(mapCtx, in.ActivationFn)
	out.Optimizer = Model_StringHparamSearchSpace_ToProto(mapCtx, in.Optimizer)
	return out
}
func Model_HparamTuningTrial_FromProto(mapCtx *direct.MapContext, in *pb.HparamTuningTrial) *krm.Model_HparamTuningTrial {
	if in == nil {
		return nil
	}
	out := &krm.Model_HparamTuningTrial{}
	out.TrialID = direct.LazyPtr(in.GetTrialId())
	out.StartTimeMs = direct.LazyPtr(in.GetStartTimeMs())
	out.EndTimeMs = direct.LazyPtr(in.GetEndTimeMs())
	out.Hparams = Model_TrainingRun_TrainingOptions_FromProto(mapCtx, in.GetHparams())
	out.EvaluationMetrics = Model_EvaluationMetrics_FromProto(mapCtx, in.GetEvaluationMetrics())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	out.TrainingLoss = direct.LazyPtr(in.GetTrainingLoss())
	out.EvalLoss = direct.LazyPtr(in.GetEvalLoss())
	out.HparamTuningEvaluationMetrics = Model_EvaluationMetrics_FromProto(mapCtx, in.GetHparamTuningEvaluationMetrics())
	return out
}
func Model_HparamTuningTrial_ToProto(mapCtx *direct.MapContext, in *krm.Model_HparamTuningTrial) *pb.HparamTuningTrial {
	if in == nil {
		return nil
	}
	out := &pb.HparamTuningTrial{}
	out.TrialId = in.TrialID
	out.StartTimeMs = in.StartTimeMs
	out.EndTimeMs = in.EndTimeMs
	out.Hparams = Model_TrainingRun_TrainingOptions_ToProto(mapCtx, in.Hparams)
	out.EvaluationMetrics = Model_EvaluationMetrics_ToProto(mapCtx, in.EvaluationMetrics)
	out.Status = in.Status
	out.ErrorMessage = in.ErrorMessage
	out.TrainingLoss = in.TrainingLoss
	out.EvalLoss = in.EvalLoss
	out.HparamTuningEvaluationMetrics = Model_EvaluationMetrics_ToProto(mapCtx, in.HparamTuningEvaluationMetrics)
	return out
}
func Model_IntArrayHparamSearchSpace_FromProto(mapCtx *direct.MapContext, in *pb.IntArrayHparamSearchSpace) *krm.Model_IntArrayHparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &krm.Model_IntArrayHparamSearchSpace{}
	out.Candidates = direct.Slice_FromProto(mapCtx, in.Candidates, Model_IntArrayHparamSearchSpace_IntArray_FromProto)
	return out
}
func Model_IntArrayHparamSearchSpace_ToProto(mapCtx *direct.MapContext, in *krm.Model_IntArrayHparamSearchSpace) *pb.IntArrayHparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &pb.IntArrayHparamSearchSpace{}
	out.Candidates = direct.Slice_ToProto(mapCtx, in.Candidates, Model_IntArrayHparamSearchSpace_IntArray_ToProto)
	return out
}
func Model_IntArrayHparamSearchSpace_IntArray_FromProto(mapCtx *direct.MapContext, in *pb.IntArray) *krm.Model_IntArrayHparamSearchSpace_IntArray {
	if in == nil {
		return nil
	}
	out := &krm.Model_IntArrayHparamSearchSpace_IntArray{}
	out.Elements = in.Elements
	return out
}
func Model_IntArrayHparamSearchSpace_IntArray_ToProto(mapCtx *direct.MapContext, in *krm.Model_IntArrayHparamSearchSpace_IntArray) *pb.IntArray {
	if in == nil {
		return nil
	}
	out := &pb.IntArray{}
	out.Elements = in.Elements
	return out
}
func Model_IntHparamSearchSpace_FromProto(mapCtx *direct.MapContext, in *pb.IntHparamSearchSpace) *krm.Model_IntHparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &krm.Model_IntHparamSearchSpace{}
	out.Range = Model_IntHparamSearchSpace_IntRange_FromProto(mapCtx, in.GetRange())
	out.Candidates = Model_IntHparamSearchSpace_IntCandidates_FromProto(mapCtx, in.GetCandidates())
	return out
}
func Model_IntHparamSearchSpace_ToProto(mapCtx *direct.MapContext, in *krm.Model_IntHparamSearchSpace) *pb.IntHparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &pb.IntHparamSearchSpace{}
	if oneof := Model_IntHparamSearchSpace_IntRange_ToProto(mapCtx, in.Range); oneof != nil {
		out.Range = oneof
	}
	if oneof := Model_IntHparamSearchSpace_IntCandidates_ToProto(mapCtx, in.Candidates); oneof != nil {
		out.Candidates = oneof
	}
	return out
}
func Model_IntHparamSearchSpace_IntCandidates_FromProto(mapCtx *direct.MapContext, in *pb.IntCandidates) *krm.Model_IntHparamSearchSpace_IntCandidates {
	if in == nil {
		return nil
	}
	out := &krm.Model_IntHparamSearchSpace_IntCandidates{}
	out.Candidates = in.Candidates
	return out
}
func Model_IntHparamSearchSpace_IntCandidates_ToProto(mapCtx *direct.MapContext, in *krm.Model_IntHparamSearchSpace_IntCandidates) *pb.IntCandidates {
	if in == nil {
		return nil
	}
	out := &pb.IntCandidates{}
	out.Candidates = in.Candidates
	return out
}
func Model_IntHparamSearchSpace_IntRange_FromProto(mapCtx *direct.MapContext, in *pb.IntRange) *krm.Model_IntHparamSearchSpace_IntRange {
	if in == nil {
		return nil
	}
	out := &krm.Model_IntHparamSearchSpace_IntRange{}
	out.Min = direct.LazyPtr(in.GetMin())
	out.Max = direct.LazyPtr(in.GetMax())
	return out
}
func Model_IntHparamSearchSpace_IntRange_ToProto(mapCtx *direct.MapContext, in *krm.Model_IntHparamSearchSpace_IntRange) *pb.IntRange {
	if in == nil {
		return nil
	}
	out := &pb.IntRange{}
	out.Min = in.Min
	out.Max = in.Max
	return out
}
func Model_MultiClassClassificationMetrics_FromProto(mapCtx *direct.MapContext, in *pb.MultiClassClassificationMetrics) *krm.Model_MultiClassClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_MultiClassClassificationMetrics{}
	out.AggregateClassificationMetrics = Model_AggregateClassificationMetrics_FromProto(mapCtx, in.GetAggregateClassificationMetrics())
	out.ConfusionMatrixList = direct.Slice_FromProto(mapCtx, in.ConfusionMatrixList, Model_MultiClassClassificationMetrics_ConfusionMatrix_FromProto)
	return out
}
func Model_MultiClassClassificationMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_MultiClassClassificationMetrics) *pb.MultiClassClassificationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.MultiClassClassificationMetrics{}
	out.AggregateClassificationMetrics = Model_AggregateClassificationMetrics_ToProto(mapCtx, in.AggregateClassificationMetrics)
	out.ConfusionMatrixList = direct.Slice_ToProto(mapCtx, in.ConfusionMatrixList, Model_MultiClassClassificationMetrics_ConfusionMatrix_ToProto)
	return out
}
func Model_MultiClassClassificationMetrics_ConfusionMatrix_FromProto(mapCtx *direct.MapContext, in *pb.ConfusionMatrix) *krm.Model_MultiClassClassificationMetrics_ConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &krm.Model_MultiClassClassificationMetrics_ConfusionMatrix{}
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.Rows = direct.Slice_FromProto(mapCtx, in.Rows, Model_MultiClassClassificationMetrics_ConfusionMatrix_Row_FromProto)
	return out
}
func Model_MultiClassClassificationMetrics_ConfusionMatrix_ToProto(mapCtx *direct.MapContext, in *krm.Model_MultiClassClassificationMetrics_ConfusionMatrix) *pb.ConfusionMatrix {
	if in == nil {
		return nil
	}
	out := &pb.ConfusionMatrix{}
	out.ConfidenceThreshold = in.ConfidenceThreshold
	out.Rows = direct.Slice_ToProto(mapCtx, in.Rows, Model_MultiClassClassificationMetrics_ConfusionMatrix_Row_ToProto)
	return out
}
func Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry {
	if in == nil {
		return nil
	}
	out := &krm.Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry{}
	out.PredictedLabel = direct.LazyPtr(in.GetPredictedLabel())
	out.ItemCount = direct.LazyPtr(in.GetItemCount())
	return out
}
func Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry_ToProto(mapCtx *direct.MapContext, in *krm.Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	out.PredictedLabel = in.PredictedLabel
	out.ItemCount = in.ItemCount
	return out
}
func Model_MultiClassClassificationMetrics_ConfusionMatrix_Row_FromProto(mapCtx *direct.MapContext, in *pb.Row) *krm.Model_MultiClassClassificationMetrics_ConfusionMatrix_Row {
	if in == nil {
		return nil
	}
	out := &krm.Model_MultiClassClassificationMetrics_ConfusionMatrix_Row{}
	out.ActualLabel = direct.LazyPtr(in.GetActualLabel())
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry_FromProto)
	return out
}
func Model_MultiClassClassificationMetrics_ConfusionMatrix_Row_ToProto(mapCtx *direct.MapContext, in *krm.Model_MultiClassClassificationMetrics_ConfusionMatrix_Row) *pb.Row {
	if in == nil {
		return nil
	}
	out := &pb.Row{}
	out.ActualLabel = in.ActualLabel
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, Model_MultiClassClassificationMetrics_ConfusionMatrix_Entry_ToProto)
	return out
}
func Model_RankingMetrics_FromProto(mapCtx *direct.MapContext, in *pb.RankingMetrics) *krm.Model_RankingMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_RankingMetrics{}
	out.MeanAveragePrecision = direct.LazyPtr(in.GetMeanAveragePrecision())
	out.MeanSquaredError = direct.LazyPtr(in.GetMeanSquaredError())
	out.NormalizedDiscountedCumulativeGain = direct.LazyPtr(in.GetNormalizedDiscountedCumulativeGain())
	out.AverageRank = direct.LazyPtr(in.GetAverageRank())
	return out
}
func Model_RankingMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_RankingMetrics) *pb.RankingMetrics {
	if in == nil {
		return nil
	}
	out := &pb.RankingMetrics{}
	out.MeanAveragePrecision = in.MeanAveragePrecision
	out.MeanSquaredError = in.MeanSquaredError
	out.NormalizedDiscountedCumulativeGain = in.NormalizedDiscountedCumulativeGain
	out.AverageRank = in.AverageRank
	return out
}
func Model_RegressionMetrics_FromProto(mapCtx *direct.MapContext, in *pb.RegressionMetrics) *krm.Model_RegressionMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Model_RegressionMetrics{}
	out.MeanAbsoluteError = direct.LazyPtr(in.GetMeanAbsoluteError())
	out.MeanSquaredError = direct.LazyPtr(in.GetMeanSquaredError())
	out.MeanSquaredLogError = direct.LazyPtr(in.GetMeanSquaredLogError())
	out.MedianAbsoluteError = direct.LazyPtr(in.GetMedianAbsoluteError())
	out.RSquared = direct.LazyPtr(in.GetRSquared())
	return out
}
func Model_RegressionMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Model_RegressionMetrics) *pb.RegressionMetrics {
	if in == nil {
		return nil
	}
	out := &pb.RegressionMetrics{}
	out.MeanAbsoluteError = in.MeanAbsoluteError
	out.MeanSquaredError = in.MeanSquaredError
	out.MeanSquaredLogError = in.MeanSquaredLogError
	out.MedianAbsoluteError = in.MedianAbsoluteError
	out.RSquared = in.RSquared
	return out
}

func Model_StringHparamSearchSpace_FromProto(mapCtx *direct.MapContext, in *pb.StringHparamSearchSpace) *krm.Model_StringHparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &krm.Model_StringHparamSearchSpace{}
	out.Candidates = in.Candidates
	return out
}
func Model_StringHparamSearchSpace_ToProto(mapCtx *direct.MapContext, in *krm.Model_StringHparamSearchSpace) *pb.StringHparamSearchSpace {
	if in == nil {
		return nil
	}
	out := &pb.StringHparamSearchSpace{}
	out.Candidates = in.Candidates
	return out
}
func Model_TrainingRun_FromProto(mapCtx *direct.MapContext, in *pb.TrainingRun) *krm.Model_TrainingRun {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun{}
	out.TrainingOptions = Model_TrainingRun_TrainingOptions_FromProto(mapCtx, in.GetTrainingOptions())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.Results = direct.Slice_FromProto(mapCtx, in.Results, Model_TrainingRun_IterationResult_FromProto)
	out.EvaluationMetrics = Model_EvaluationMetrics_FromProto(mapCtx, in.GetEvaluationMetrics())
	out.DataSplitResult = Model_DataSplitResult_FromProto(mapCtx, in.GetDataSplitResult())
	out.ModelLevelGlobalExplanation = Model_GlobalExplanation_FromProto(mapCtx, in.GetModelLevelGlobalExplanation())
	out.ClassLevelGlobalExplanations = direct.Slice_FromProto(mapCtx, in.ClassLevelGlobalExplanations, Model_GlobalExplanation_FromProto)
	out.VertexAiModelID = direct.LazyPtr(in.GetVertexAiModelId())
	out.VertexAiModelVersion = direct.LazyPtr(in.GetVertexAiModelVersion())
	return out
}
func Model_TrainingRun_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun) *pb.TrainingRun {
	if in == nil {
		return nil
	}
	out := &pb.TrainingRun{}
	out.TrainingOptions = Model_TrainingRun_TrainingOptions_ToProto(mapCtx, in.TrainingOptions)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.Results = direct.Slice_ToProto(mapCtx, in.Results, Model_TrainingRun_IterationResult_ToProto)
	out.EvaluationMetrics = Model_EvaluationMetrics_ToProto(mapCtx, in.EvaluationMetrics)
	out.DataSplitResult = Model_DataSplitResult_ToProto(mapCtx, in.DataSplitResult)
	out.ModelLevelGlobalExplanation = Model_GlobalExplanation_ToProto(mapCtx, in.ModelLevelGlobalExplanation)
	out.ClassLevelGlobalExplanations = direct.Slice_ToProto(mapCtx, in.ClassLevelGlobalExplanations, Model_GlobalExplanation_ToProto)
	out.VertexAiModelId = in.VertexAiModelID
	out.VertexAiModelVersion = in.VertexAiModelVersion
	return out
}
func Model_TrainingRun_IterationResult_FromProto(mapCtx *direct.MapContext, in *pb.IterationResult) *krm.Model_TrainingRun_IterationResult {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun_IterationResult{}
	out.Index = direct.LazyPtr(in.GetIndex())
	out.DurationMs = direct.LazyPtr(in.GetDurationMs())
	out.TrainingLoss = direct.LazyPtr(in.GetTrainingLoss())
	out.EvalLoss = direct.LazyPtr(in.GetEvalLoss())
	out.LearnRate = direct.LazyPtr(in.GetLearnRate())
	out.ClusterInfos = direct.Slice_FromProto(mapCtx, in.ClusterInfos, Model_TrainingRun_IterationResult_ClusterInfo_FromProto)
	out.ArimaResult = Model_TrainingRun_IterationResult_ArimaResult_FromProto(mapCtx, in.GetArimaResult())
	out.PrincipalComponentInfos = direct.Slice_FromProto(mapCtx, in.PrincipalComponentInfos, Model_TrainingRun_IterationResult_PrincipalComponentInfo_FromProto)
	return out
}
func Model_TrainingRun_IterationResult_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun_IterationResult) *pb.IterationResult {
	if in == nil {
		return nil
	}
	out := &pb.IterationResult{}
	out.Index = in.Index
	out.DurationMs = in.DurationMs
	out.TrainingLoss = in.TrainingLoss
	out.EvalLoss = in.EvalLoss
	out.LearnRate = in.LearnRate
	out.ClusterInfos = direct.Slice_ToProto(mapCtx, in.ClusterInfos, Model_TrainingRun_IterationResult_ClusterInfo_ToProto)
	out.ArimaResult = Model_TrainingRun_IterationResult_ArimaResult_ToProto(mapCtx, in.ArimaResult)
	out.PrincipalComponentInfos = direct.Slice_ToProto(mapCtx, in.PrincipalComponentInfos, Model_TrainingRun_IterationResult_PrincipalComponentInfo_ToProto)
	return out
}
func Model_TrainingRun_IterationResult_ArimaResult_FromProto(mapCtx *direct.MapContext, in *pb.ArimaResult) *krm.Model_TrainingRun_IterationResult_ArimaResult {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun_IterationResult_ArimaResult{}
	out.ArimaModelInfo = direct.Slice_FromProto(mapCtx, in.ArimaModelInfo, Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo_FromProto)
	out.SeasonalPeriods = in.SeasonalPeriods
	return out
}
func Model_TrainingRun_IterationResult_ArimaResult_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun_IterationResult_ArimaResult) *pb.ArimaResult {
	if in == nil {
		return nil
	}
	out := &pb.ArimaResult{}
	out.ArimaModelInfo = direct.Slice_ToProto(mapCtx, in.ArimaModelInfo, Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo_ToProto)
	out.SeasonalPeriods = in.SeasonalPeriods
	return out
}
func Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients_FromProto(mapCtx *direct.MapContext, in *pb.ArimaCoefficients) *krm.Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients{}
	out.AutoRegressiveCoefficients = in.AutoRegressiveCoefficients
	out.MovingAverageCoefficients = in.MovingAverageCoefficients
	out.InterceptCoefficient = direct.LazyPtr(in.GetInterceptCoefficient())
	return out
}
func Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients) *pb.ArimaCoefficients {
	if in == nil {
		return nil
	}
	out := &pb.ArimaCoefficients{}
	out.AutoRegressiveCoefficients = in.AutoRegressiveCoefficients
	out.MovingAverageCoefficients = in.MovingAverageCoefficients
	out.InterceptCoefficient = in.InterceptCoefficient
	return out
}
func Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo_FromProto(mapCtx *direct.MapContext, in *pb.ArimaModelInfo) *krm.Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo{}
	out.NonSeasonalOrder = Model_ArimaOrder_FromProto(mapCtx, in.GetNonSeasonalOrder())
	out.ArimaCoefficients = Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients_FromProto(mapCtx, in.GetArimaCoefficients())
	out.ArimaFittingMetrics = Model_ArimaFittingMetrics_FromProto(mapCtx, in.GetArimaFittingMetrics())
	out.HasDrift = direct.LazyPtr(in.GetHasDrift())
	out.TimeSeriesID = direct.LazyPtr(in.GetTimeSeriesId())
	out.TimeSeriesIds = in.TimeSeriesIds
	out.SeasonalPeriods = in.SeasonalPeriods
	out.HasHolidayEffect = direct.LazyPtr(in.GetHasHolidayEffect())
	out.HasSpikesAndDips = direct.LazyPtr(in.GetHasSpikesAndDips())
	out.HasStepChanges = direct.LazyPtr(in.GetHasStepChanges())
	return out
}
func Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun_IterationResult_ArimaResult_ArimaModelInfo) *pb.ArimaModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.ArimaModelInfo{}
	out.NonSeasonalOrder = Model_ArimaOrder_ToProto(mapCtx, in.NonSeasonalOrder)
	out.ArimaCoefficients = Model_TrainingRun_IterationResult_ArimaResult_ArimaCoefficients_ToProto(mapCtx, in.ArimaCoefficients)
	out.ArimaFittingMetrics = Model_ArimaFittingMetrics_ToProto(mapCtx, in.ArimaFittingMetrics)
	out.HasDrift = in.HasDrift
	out.TimeSeriesId = in.TimeSeriesID
	out.TimeSeriesIds = in.TimeSeriesIds
	out.SeasonalPeriods = in.SeasonalPeriods
	out.HasHolidayEffect = in.HasHolidayEffect
	out.HasSpikesAndDips = in.HasSpikesAndDips
	out.HasStepChanges = in.HasStepChanges
	return out
}
func Model_TrainingRun_IterationResult_ClusterInfo_FromProto(mapCtx *direct.MapContext, in *pb.ClusterInfo) *krm.Model_TrainingRun_IterationResult_ClusterInfo {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun_IterationResult_ClusterInfo{}
	out.CentroidID = direct.LazyPtr(in.GetCentroidId())
	out.ClusterRadius = direct.LazyPtr(in.GetClusterRadius())
	out.ClusterSize = direct.LazyPtr(in.GetClusterSize())
	return out
}
func Model_TrainingRun_IterationResult_ClusterInfo_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun_IterationResult_ClusterInfo) *pb.ClusterInfo {
	if in == nil {
		return nil
	}
	out := &pb.ClusterInfo{}
	out.CentroidId = in.CentroidID
	out.ClusterRadius = in.ClusterRadius
	out.ClusterSize = in.ClusterSize
	return out
}
func Model_TrainingRun_IterationResult_PrincipalComponentInfo_FromProto(mapCtx *direct.MapContext, in *pb.PrincipalComponentInfo) *krm.Model_TrainingRun_IterationResult_PrincipalComponentInfo {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun_IterationResult_PrincipalComponentInfo{}
	out.PrincipalComponentID = direct.LazyPtr(in.GetPrincipalComponentId())
	out.ExplainedVariance = direct.LazyPtr(in.GetExplainedVariance())
	out.ExplainedVarianceRatio = direct.LazyPtr(in.GetExplainedVarianceRatio())
	out.CumulativeExplainedVarianceRatio = direct.LazyPtr(in.GetCumulativeExplainedVarianceRatio())
	return out
}
func Model_TrainingRun_IterationResult_PrincipalComponentInfo_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun_IterationResult_PrincipalComponentInfo) *pb.PrincipalComponentInfo {
	if in == nil {
		return nil
	}
	out := &pb.PrincipalComponentInfo{}
	out.PrincipalComponentId = in.PrincipalComponentID
	out.ExplainedVariance = in.ExplainedVariance
	out.ExplainedVarianceRatio = in.ExplainedVarianceRatio
	out.CumulativeExplainedVarianceRatio = in.CumulativeExplainedVarianceRatio
	return out
}
func Model_TrainingRun_TrainingOptions_FromProto(mapCtx *direct.MapContext, in *pb.TrainingOptions) *krm.Model_TrainingRun_TrainingOptions {
	if in == nil {
		return nil
	}
	out := &krm.Model_TrainingRun_TrainingOptions{}
	out.MaxIterations = direct.LazyPtr(in.GetMaxIterations())
	out.LossType = direct.LazyPtr(in.GetLossType())
	out.LearnRate = direct.LazyPtr(in.GetLearnRate())
	out.L1Regularization = direct.LazyPtr(in.GetL1Regularization())
	out.L2Regularization = direct.LazyPtr(in.GetL2Regularization())
	out.MinRelativeProgress = direct.LazyPtr(in.GetMinRelativeProgress())
	out.WarmStart = direct.LazyPtr(in.GetWarmStart())
	out.EarlyStop = direct.LazyPtr(in.GetEarlyStop())
	out.InputLabelColumns = in.InputLabelColumns
	out.DataSplitMethod = direct.LazyPtr(in.GetDataSplitMethod())
	out.DataSplitEvalFraction = direct.LazyPtr(in.GetDataSplitEvalFraction())
	out.DataSplitColumn = direct.LazyPtr(in.GetDataSplitColumn())
	out.LearnRateStrategy = direct.LazyPtr(in.GetLearnRateStrategy())
	out.InitialLearnRate = direct.LazyPtr(in.GetInitialLearnRate())
	// 	/* NOTYET */: LabelClassWeights
	out.UserColumn = direct.LazyPtr(in.GetUserColumn())
	out.ItemColumn = direct.LazyPtr(in.GetItemColumn())
	out.DistanceType = direct.LazyPtr(in.GetDistanceType())
	out.NumClusters = direct.LazyPtr(in.GetNumClusters())
	out.ModelUri = direct.LazyPtr(in.GetModelUri())
	out.OptimizationStrategy = direct.LazyPtr(in.GetOptimizationStrategy())
	out.HiddenUnits = in.HiddenUnits
	out.BatchSize = direct.LazyPtr(in.GetBatchSize())
	out.Dropout = direct.LazyPtr(in.GetDropout())
	out.MaxTreeDepth = direct.LazyPtr(in.GetMaxTreeDepth())
	out.Subsample = direct.LazyPtr(in.GetSubsample())
	out.MinSplitLoss = direct.LazyPtr(in.GetMinSplitLoss())
	out.BoosterType = direct.LazyPtr(in.GetBoosterType())
	out.NumParallelTree = direct.LazyPtr(in.GetNumParallelTree())
	out.DartNormalizeType = direct.LazyPtr(in.GetDartNormalizeType())
	out.TreeMethod = direct.LazyPtr(in.GetTreeMethod())
	out.MinTreeChildWeight = direct.LazyPtr(in.GetMinTreeChildWeight())
	out.ColsampleBytree = direct.LazyPtr(in.GetColsampleBytree())
	out.ColsampleBylevel = direct.LazyPtr(in.GetColsampleBylevel())
	out.ColsampleBynode = direct.LazyPtr(in.GetColsampleBynode())
	out.NumFactors = direct.LazyPtr(in.GetNumFactors())
	out.FeedbackType = direct.LazyPtr(in.GetFeedbackType())
	out.WalsAlpha = direct.LazyPtr(in.GetWalsAlpha())
	out.KmeansInitializationMethod = direct.LazyPtr(in.GetKmeansInitializationMethod())
	out.KmeansInitializationColumn = direct.LazyPtr(in.GetKmeansInitializationColumn())
	out.TimeSeriesTimestampColumn = direct.LazyPtr(in.GetTimeSeriesTimestampColumn())
	out.TimeSeriesDataColumn = direct.LazyPtr(in.GetTimeSeriesDataColumn())
	out.AutoArima = direct.LazyPtr(in.GetAutoArima())
	out.NonSeasonalOrder = Model_ArimaOrder_FromProto(mapCtx, in.GetNonSeasonalOrder())
	out.DataFrequency = direct.LazyPtr(in.GetDataFrequency())
	out.CalculatePValues = direct.LazyPtr(in.GetCalculatePValues())
	out.IncludeDrift = direct.LazyPtr(in.GetIncludeDrift())
	out.HolidayRegion = direct.LazyPtr(in.GetHolidayRegion())
	out.HolidayRegions = in.HolidayRegions
	out.TimeSeriesIDColumn = direct.LazyPtr(in.GetTimeSeriesIdColumn())
	out.TimeSeriesIDColumns = in.TimeSeriesIdColumns
	out.Horizon = direct.LazyPtr(in.GetHorizon())
	out.AutoArimaMaxOrder = direct.LazyPtr(in.GetAutoArimaMaxOrder())
	out.AutoArimaMinOrder = direct.LazyPtr(in.GetAutoArimaMinOrder())
	out.NumTrials = direct.LazyPtr(in.GetNumTrials())
	out.MaxParallelTrials = direct.LazyPtr(in.GetMaxParallelTrials())
	out.HparamTuningObjectives = in.HparamTuningObjectives
	out.DecomposeTimeSeries = direct.LazyPtr(in.GetDecomposeTimeSeries())
	out.CleanSpikesAndDips = direct.LazyPtr(in.GetCleanSpikesAndDips())
	out.AdjustStepChanges = direct.LazyPtr(in.GetAdjustStepChanges())
	out.EnableGlobalExplain = direct.LazyPtr(in.GetEnableGlobalExplain())
	out.SampledShapleyNumPaths = direct.LazyPtr(in.GetSampledShapleyNumPaths())
	out.IntegratedGradientsNumSteps = direct.LazyPtr(in.GetIntegratedGradientsNumSteps())
	out.CategoryEncodingMethod = direct.LazyPtr(in.GetCategoryEncodingMethod())
	out.TfVersion = direct.LazyPtr(in.GetTfVersion())
	out.ColorSpace = direct.LazyPtr(in.GetColorSpace())
	out.InstanceWeightColumn = direct.LazyPtr(in.GetInstanceWeightColumn())
	out.TrendSmoothingWindowSize = direct.LazyPtr(in.GetTrendSmoothingWindowSize())
	out.TimeSeriesLengthFraction = direct.LazyPtr(in.GetTimeSeriesLengthFraction())
	out.MinTimeSeriesLength = direct.LazyPtr(in.GetMinTimeSeriesLength())
	out.MaxTimeSeriesLength = direct.LazyPtr(in.GetMaxTimeSeriesLength())
	out.XgboostVersion = direct.LazyPtr(in.GetXgboostVersion())
	out.ApproxGlobalFeatureContrib = direct.LazyPtr(in.GetApproxGlobalFeatureContrib())
	out.FitIntercept = direct.LazyPtr(in.GetFitIntercept())
	out.NumPrincipalComponents = direct.LazyPtr(in.GetNumPrincipalComponents())
	out.PcaExplainedVarianceRatio = direct.LazyPtr(in.GetPcaExplainedVarianceRatio())
	out.ScaleFeatures = direct.LazyPtr(in.GetScaleFeatures())
	out.PcaSolver = direct.LazyPtr(in.GetPcaSolver())
	out.AutoClassWeights = direct.LazyPtr(in.GetAutoClassWeights())
	out.ActivationFn = direct.LazyPtr(in.GetActivationFn())
	out.Optimizer = direct.LazyPtr(in.GetOptimizer())
	out.BudgetHours = direct.LazyPtr(in.GetBudgetHours())
	out.StandardizeFeatures = direct.LazyPtr(in.GetStandardizeFeatures())
	out.L1RegActivation = direct.LazyPtr(in.GetL1RegActivation())
	out.ModelRegistry = direct.LazyPtr(in.GetModelRegistry())
	out.VertexAiModelVersionAliases = in.VertexAiModelVersionAliases
	return out
}
func Model_TrainingRun_TrainingOptions_ToProto(mapCtx *direct.MapContext, in *krm.Model_TrainingRun_TrainingOptions) *pb.TrainingOptions {
	if in == nil {
		return nil
	}
	out := &pb.TrainingOptions{}
	out.MaxIterations = in.MaxIterations
	out.LossType = in.LossType
	out.LearnRate = in.LearnRate
	out.L1Regularization = in.L1Regularization
	out.L2Regularization = in.L2Regularization
	out.MinRelativeProgress = in.MinRelativeProgress
	out.WarmStart = in.WarmStart
	out.EarlyStop = in.EarlyStop
	out.InputLabelColumns = in.InputLabelColumns
	out.DataSplitMethod = in.DataSplitMethod
	out.DataSplitEvalFraction = in.DataSplitEvalFraction
	out.DataSplitColumn = in.DataSplitColumn
	out.LearnRateStrategy = in.LearnRateStrategy
	out.InitialLearnRate = in.InitialLearnRate
	// 	/* NOTYET */: LabelClassWeights
	out.UserColumn = in.UserColumn
	out.ItemColumn = in.ItemColumn
	out.DistanceType = in.DistanceType
	out.NumClusters = in.NumClusters
	out.ModelUri = in.ModelUri
	out.OptimizationStrategy = in.OptimizationStrategy
	out.HiddenUnits = in.HiddenUnits
	out.BatchSize = in.BatchSize
	out.Dropout = in.Dropout
	out.MaxTreeDepth = in.MaxTreeDepth
	out.Subsample = in.Subsample
	out.MinSplitLoss = in.MinSplitLoss
	out.BoosterType = in.BoosterType
	out.NumParallelTree = in.NumParallelTree
	out.DartNormalizeType = in.DartNormalizeType
	out.TreeMethod = in.TreeMethod
	out.MinTreeChildWeight = in.MinTreeChildWeight
	out.ColsampleBytree = in.ColsampleBytree
	out.ColsampleBylevel = in.ColsampleBylevel
	out.ColsampleBynode = in.ColsampleBynode
	out.NumFactors = in.NumFactors
	out.FeedbackType = in.FeedbackType
	out.WalsAlpha = in.WalsAlpha
	out.KmeansInitializationMethod = in.KmeansInitializationMethod
	out.KmeansInitializationColumn = in.KmeansInitializationColumn
	out.TimeSeriesTimestampColumn = in.TimeSeriesTimestampColumn
	out.TimeSeriesDataColumn = in.TimeSeriesDataColumn
	out.AutoArima = in.AutoArima
	out.NonSeasonalOrder = Model_ArimaOrder_ToProto(mapCtx, in.NonSeasonalOrder)
	out.DataFrequency = in.DataFrequency
	out.CalculatePValues = in.CalculatePValues
	out.IncludeDrift = in.IncludeDrift
	out.HolidayRegion = in.HolidayRegion
	out.HolidayRegions = in.HolidayRegions
	out.TimeSeriesIdColumn = in.TimeSeriesIDColumn
	out.TimeSeriesIdColumns = in.TimeSeriesIDColumns
	out.Horizon = in.Horizon
	out.AutoArimaMaxOrder = in.AutoArimaMaxOrder
	out.AutoArimaMinOrder = in.AutoArimaMinOrder
	out.NumTrials = in.NumTrials
	out.MaxParallelTrials = in.MaxParallelTrials
	out.HparamTuningObjectives = in.HparamTuningObjectives
	out.DecomposeTimeSeries = in.DecomposeTimeSeries
	out.CleanSpikesAndDips = in.CleanSpikesAndDips
	out.AdjustStepChanges = in.AdjustStepChanges
	out.EnableGlobalExplain = in.EnableGlobalExplain
	out.SampledShapleyNumPaths = in.SampledShapleyNumPaths
	out.IntegratedGradientsNumSteps = in.IntegratedGradientsNumSteps
	out.CategoryEncodingMethod = in.CategoryEncodingMethod
	out.TfVersion = in.TfVersion
	out.ColorSpace = in.ColorSpace
	out.InstanceWeightColumn = in.InstanceWeightColumn
	out.TrendSmoothingWindowSize = in.TrendSmoothingWindowSize
	out.TimeSeriesLengthFraction = in.TimeSeriesLengthFraction
	out.MinTimeSeriesLength = in.MinTimeSeriesLength
	out.MaxTimeSeriesLength = in.MaxTimeSeriesLength
	out.XgboostVersion = in.XgboostVersion
	out.ApproxGlobalFeatureContrib = in.ApproxGlobalFeatureContrib
	out.FitIntercept = in.FitIntercept
	out.NumPrincipalComponents = in.NumPrincipalComponents
	out.PcaExplainedVarianceRatio = in.PcaExplainedVarianceRatio
	out.ScaleFeatures = in.ScaleFeatures
	out.PcaSolver = in.PcaSolver
	out.AutoClassWeights = in.AutoClassWeights
	out.ActivationFn = in.ActivationFn
	out.Optimizer = in.Optimizer
	out.BudgetHours = in.BudgetHours
	out.StandardizeFeatures = in.StandardizeFeatures
	out.L1RegActivation = in.L1RegActivation
	out.ModelRegistry = in.ModelRegistry
	out.VertexAiModelVersionAliases = in.VertexAiModelVersionAliases
	return out
}
func ParquetOptions_FromProto(mapCtx *direct.MapContext, in *pb.ParquetOptions) *krm.ParquetOptions {
	if in == nil {
		return nil
	}
	out := &krm.ParquetOptions{}
	out.EnumAsString = direct.LazyPtr(in.GetEnumAsString())
	out.EnableListInference = direct.LazyPtr(in.GetEnableListInference())
	out.MapTargetType = direct.LazyPtr(in.GetMapTargetType())
	return out
}
func ParquetOptions_ToProto(mapCtx *direct.MapContext, in *krm.ParquetOptions) *pb.ParquetOptions {
	if in == nil {
		return nil
	}
	out := &pb.ParquetOptions{}
	out.EnumAsString = in.EnumAsString
	out.EnableListInference = in.EnableListInference
	out.MapTargetType = in.MapTargetType
	return out
}
func PartitionSkew_FromProto(mapCtx *direct.MapContext, in *pb.PartitionSkew) *krm.PartitionSkew {
	if in == nil {
		return nil
	}
	out := &krm.PartitionSkew{}
	out.SkewSources = direct.Slice_FromProto(mapCtx, in.SkewSources, PartitionSkew_SkewSource_FromProto)
	return out
}
func PartitionSkew_ToProto(mapCtx *direct.MapContext, in *krm.PartitionSkew) *pb.PartitionSkew {
	if in == nil {
		return nil
	}
	out := &pb.PartitionSkew{}
	out.SkewSources = direct.Slice_ToProto(mapCtx, in.SkewSources, PartitionSkew_SkewSource_ToProto)
	return out
}
func PartitionSkew_SkewSource_FromProto(mapCtx *direct.MapContext, in *pb.SkewSource) *krm.PartitionSkew_SkewSource {
	if in == nil {
		return nil
	}
	out := &krm.PartitionSkew_SkewSource{}
	out.StageID = direct.LazyPtr(in.GetStageId())
	return out
}
func PartitionSkew_SkewSource_ToProto(mapCtx *direct.MapContext, in *krm.PartitionSkew_SkewSource) *pb.SkewSource {
	if in == nil {
		return nil
	}
	out := &pb.SkewSource{}
	out.StageId = in.StageID
	return out
}
func PartitionedColumn_FromProto(mapCtx *direct.MapContext, in *pb.PartitionedColumn) *krm.PartitionedColumn {
	if in == nil {
		return nil
	}
	out := &krm.PartitionedColumn{}
	out.Field = in.Field
	return out
}
func PartitionedColumn_ToProto(mapCtx *direct.MapContext, in *krm.PartitionedColumn) *pb.PartitionedColumn {
	if in == nil {
		return nil
	}
	out := &pb.PartitionedColumn{}
	out.Field = in.Field
	return out
}
func PartitioningDefinition_FromProto(mapCtx *direct.MapContext, in *pb.PartitioningDefinition) *krm.PartitioningDefinition {
	if in == nil {
		return nil
	}
	out := &krm.PartitioningDefinition{}
	out.PartitionedColumn = direct.Slice_FromProto(mapCtx, in.PartitionedColumn, PartitionedColumn_FromProto)
	return out
}
func PartitioningDefinition_ToProto(mapCtx *direct.MapContext, in *krm.PartitioningDefinition) *pb.PartitioningDefinition {
	if in == nil {
		return nil
	}
	out := &pb.PartitioningDefinition{}
	out.PartitionedColumn = direct.Slice_ToProto(mapCtx, in.PartitionedColumn, PartitionedColumn_ToProto)
	return out
}
func PerformanceInsights_FromProto(mapCtx *direct.MapContext, in *pb.PerformanceInsights) *krm.PerformanceInsights {
	if in == nil {
		return nil
	}
	out := &krm.PerformanceInsights{}
	out.AvgPreviousExecutionMs = direct.LazyPtr(in.GetAvgPreviousExecutionMs())
	out.StagePerformanceStandaloneInsights = direct.Slice_FromProto(mapCtx, in.StagePerformanceStandaloneInsights, StagePerformanceStandaloneInsight_FromProto)
	out.StagePerformanceChangeInsights = direct.Slice_FromProto(mapCtx, in.StagePerformanceChangeInsights, StagePerformanceChangeInsight_FromProto)
	return out
}
func PerformanceInsights_ToProto(mapCtx *direct.MapContext, in *krm.PerformanceInsights) *pb.PerformanceInsights {
	if in == nil {
		return nil
	}
	out := &pb.PerformanceInsights{}
	out.AvgPreviousExecutionMs = in.AvgPreviousExecutionMs
	out.StagePerformanceStandaloneInsights = direct.Slice_ToProto(mapCtx, in.StagePerformanceStandaloneInsights, StagePerformanceStandaloneInsight_ToProto)
	out.StagePerformanceChangeInsights = direct.Slice_ToProto(mapCtx, in.StagePerformanceChangeInsights, StagePerformanceChangeInsight_ToProto)
	return out
}
func PrimaryKey_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraintsPrimaryKey) *krm.PrimaryKey {
	if in == nil {
		return nil
	}
	out := &krm.PrimaryKey{}
	out.Columns = in.Columns
	return out
}
func PrimaryKey_ToProto(mapCtx *direct.MapContext, in *krm.PrimaryKey) *pb.TableConstraintsPrimaryKey {
	if in == nil {
		return nil
	}
	out := &pb.TableConstraintsPrimaryKey{}
	out.Columns = in.Columns
	return out
}
func PrivacyPolicy_FromProto(mapCtx *direct.MapContext, in *pb.PrivacyPolicy) *krm.PrivacyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.PrivacyPolicy{}
	out.AggregationThresholdPolicy = AggregationThresholdPolicy_FromProto(mapCtx, in.GetAggregationThresholdPolicy())
	out.DifferentialPrivacyPolicy = DifferentialPrivacyPolicy_FromProto(mapCtx, in.GetDifferentialPrivacyPolicy())
	out.JoinRestrictionPolicy = JoinRestrictionPolicy_FromProto(mapCtx, in.GetJoinRestrictionPolicy())
	return out
}
func PrivacyPolicy_ToProto(mapCtx *direct.MapContext, in *krm.PrivacyPolicy) *pb.PrivacyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.PrivacyPolicy{}
	if oneof := AggregationThresholdPolicy_ToProto(mapCtx, in.AggregationThresholdPolicy); oneof != nil {
		out.AggregationThresholdPolicy = oneof
	}
	if oneof := DifferentialPrivacyPolicy_ToProto(mapCtx, in.DifferentialPrivacyPolicy); oneof != nil {
		out.DifferentialPrivacyPolicy = oneof
	}
	if oneof := JoinRestrictionPolicy_ToProto(mapCtx, in.JoinRestrictionPolicy); oneof != nil {
		out.JoinRestrictionPolicy = oneof
	}
	return out
}
func QueryInfo_FromProto(mapCtx *direct.MapContext, in *pb.QueryInfo) *krm.QueryInfo {
	if in == nil {
		return nil
	}
	out := &krm.QueryInfo{}
	/* NOTYET: OptimizationDetails */
	return out
}
func QueryInfo_ToProto(mapCtx *direct.MapContext, in *krm.QueryInfo) *pb.QueryInfo {
	if in == nil {
		return nil
	}
	out := &pb.QueryInfo{}
	/* NOTYET: OptimizationDetails */
	return out
}
func QueryParameter_FromProto(mapCtx *direct.MapContext, in *pb.QueryParameter) *krm.QueryParameter {
	if in == nil {
		return nil
	}
	out := &krm.QueryParameter{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ParameterType = QueryParameterType_FromProto(mapCtx, in.GetParameterType())
	out.ParameterValue = QueryParameterValue_FromProto(mapCtx, in.GetParameterValue())
	return out
}
func QueryParameter_ToProto(mapCtx *direct.MapContext, in *krm.QueryParameter) *pb.QueryParameter {
	if in == nil {
		return nil
	}
	out := &pb.QueryParameter{}
	out.Name = in.Name
	out.ParameterType = QueryParameterType_ToProto(mapCtx, in.ParameterType)
	out.ParameterValue = QueryParameterValue_ToProto(mapCtx, in.ParameterValue)
	return out
}
func QueryParameterStructType_FromProto(mapCtx *direct.MapContext, in *pb.QueryParameterTypeStructTypes) *krm.QueryParameterStructType {
	if in == nil {
		return nil
	}
	out := &krm.QueryParameterStructType{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = QueryParameterType_FromProto(mapCtx, in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func QueryParameterStructType_ToProto(mapCtx *direct.MapContext, in *krm.QueryParameterStructType) *pb.QueryParameterTypeStructTypes {
	if in == nil {
		return nil
	}
	out := &pb.QueryParameterTypeStructTypes{}
	out.Name = in.Name
	out.Type = QueryParameterType_ToProto(mapCtx, in.Type)
	out.Description = in.Description
	return out
}
func QueryParameterType_FromProto(mapCtx *direct.MapContext, in *pb.QueryParameterType) *krm.QueryParameterType {
	if in == nil {
		return nil
	}
	out := &krm.QueryParameterType{}
	out.Type = direct.LazyPtr(in.GetType())
	out.ArrayType = QueryParameterType_FromProto(mapCtx, in.GetArrayType())
	out.StructTypes = direct.Slice_FromProto(mapCtx, in.StructTypes, QueryParameterStructType_FromProto)
	out.RangeElementType = QueryParameterType_FromProto(mapCtx, in.GetRangeElementType())
	return out
}
func QueryParameterType_ToProto(mapCtx *direct.MapContext, in *krm.QueryParameterType) *pb.QueryParameterType {
	if in == nil {
		return nil
	}
	out := &pb.QueryParameterType{}
	out.Type = in.Type
	out.ArrayType = QueryParameterType_ToProto(mapCtx, in.ArrayType)
	out.StructTypes = direct.Slice_ToProto(mapCtx, in.StructTypes, QueryParameterStructType_ToProto)
	out.RangeElementType = QueryParameterType_ToProto(mapCtx, in.RangeElementType)
	return out
}
func QueryParameterValue_FromProto(mapCtx *direct.MapContext, in *pb.QueryParameterValue) *krm.QueryParameterValue {
	if in == nil {
		return nil
	}
	out := &krm.QueryParameterValue{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.ArrayValues = direct.Slice_FromProto(mapCtx, in.ArrayValues, QueryParameterValue_FromProto)
	/* NOTYET: StructValues*/
	out.RangeValue = RangeValue_FromProto(mapCtx, in.GetRangeValue())
	return out
}
func QueryParameterValue_ToProto(mapCtx *direct.MapContext, in *krm.QueryParameterValue) *pb.QueryParameterValue {
	if in == nil {
		return nil
	}
	out := &pb.QueryParameterValue{}
	out.Value = in.Value
	out.ArrayValues = direct.Slice_ToProto(mapCtx, in.ArrayValues, QueryParameterValue_ToProto)
	/* NOTYET: StructValues*/
	out.RangeValue = RangeValue_ToProto(mapCtx, in.RangeValue)
	return out
}
func QueryTimelineSample_FromProto(mapCtx *direct.MapContext, in *pb.QueryTimelineSample) *krm.QueryTimelineSample {
	if in == nil {
		return nil
	}
	out := &krm.QueryTimelineSample{}
	out.ElapsedMs = direct.LazyPtr(in.GetElapsedMs())
	out.TotalSlotMs = direct.LazyPtr(in.GetTotalSlotMs())
	out.PendingUnits = direct.LazyPtr(in.GetPendingUnits())
	out.CompletedUnits = direct.LazyPtr(in.GetCompletedUnits())
	out.ActiveUnits = direct.LazyPtr(in.GetActiveUnits())
	out.EstimatedRunnableUnits = direct.LazyPtr(in.GetEstimatedRunnableUnits())
	return out
}
func QueryTimelineSample_ToProto(mapCtx *direct.MapContext, in *krm.QueryTimelineSample) *pb.QueryTimelineSample {
	if in == nil {
		return nil
	}
	out := &pb.QueryTimelineSample{}
	out.ElapsedMs = in.ElapsedMs
	out.TotalSlotMs = in.TotalSlotMs
	out.PendingUnits = in.PendingUnits
	out.CompletedUnits = in.CompletedUnits
	out.ActiveUnits = in.ActiveUnits
	out.EstimatedRunnableUnits = in.EstimatedRunnableUnits
	return out
}
func RangePartitioning_FromProto(mapCtx *direct.MapContext, in *pb.RangePartitioning) *krm.RangePartitioning {
	if in == nil {
		return nil
	}
	out := &krm.RangePartitioning{}
	out.Field = direct.LazyPtr(in.GetField())
	out.Range = RangePartitioning_Range_FromProto(mapCtx, in.GetRange())
	return out
}
func RangePartitioning_ToProto(mapCtx *direct.MapContext, in *krm.RangePartitioning) *pb.RangePartitioning {
	if in == nil {
		return nil
	}
	out := &pb.RangePartitioning{}
	out.Field = in.Field
	out.Range = RangePartitioning_Range_ToProto(mapCtx, in.Range)
	return out
}
func RangePartitioning_Range_FromProto(mapCtx *direct.MapContext, in *pb.RangePartitioningRange) *krm.RangePartitioning_Range {
	if in == nil {
		return nil
	}
	out := &krm.RangePartitioning_Range{}
	start := strconv.FormatInt(*in.Start, 10)
	end := strconv.FormatInt(*in.End, 10)
	interval := strconv.FormatInt(*in.Interval, 10)
	out.Start = direct.LazyPtr(start)
	out.End = direct.LazyPtr(end)
	out.Interval = direct.LazyPtr(interval)
	return out
}
func RangePartitioning_Range_ToProto(mapCtx *direct.MapContext, in *krm.RangePartitioning_Range) *pb.RangePartitioningRange {
	if in == nil {
		return nil
	}
	out := &pb.RangePartitioningRange{}
	start, _ := strconv.ParseInt(*in.Start, 10, 64)
	end, _ := strconv.ParseInt(*in.End, 10, 64)
	interval, _ := strconv.ParseInt(*in.Interval, 10, 64)
	out.Start = direct.LazyPtr(start)
	out.End = direct.LazyPtr(end)
	out.Interval = direct.LazyPtr(interval)
	return out
}
func RangeValue_FromProto(mapCtx *direct.MapContext, in *pb.RangeValue) *krm.RangeValue {
	if in == nil {
		return nil
	}
	out := &krm.RangeValue{}
	out.Start = QueryParameterValue_FromProto(mapCtx, in.GetStart())
	out.End = QueryParameterValue_FromProto(mapCtx, in.GetEnd())
	return out
}
func RangeValue_ToProto(mapCtx *direct.MapContext, in *krm.RangeValue) *pb.RangeValue {
	if in == nil {
		return nil
	}
	out := &pb.RangeValue{}
	out.Start = QueryParameterValue_ToProto(mapCtx, in.Start)
	out.End = QueryParameterValue_ToProto(mapCtx, in.End)
	return out
}
func ReferencedTable_FromProto(mapCtx *direct.MapContext, in *pb.ReferencedTable) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func ReferencedTable_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.ReferencedTable {
	if in == nil {
		return nil
	}
	out := &pb.ReferencedTable{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func RemoteModelInfo_FromProto(mapCtx *direct.MapContext, in *pb.RemoteModelInfo) *krm.RemoteModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.RemoteModelInfo{}
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	out.RemoteServiceType = direct.LazyPtr(in.GetRemoteServiceType())
	out.Connection = direct.LazyPtr(in.GetConnection())
	out.MaxBatchingRows = direct.LazyPtr(in.GetMaxBatchingRows())
	out.RemoteModelVersion = direct.LazyPtr(in.GetRemoteModelVersion())
	out.SpeechRecognizer = direct.LazyPtr(in.GetSpeechRecognizer())
	return out
}
func RemoteModelInfo_ToProto(mapCtx *direct.MapContext, in *krm.RemoteModelInfo) *pb.RemoteModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.RemoteModelInfo{}
	out.Endpoint = in.Endpoint
	out.RemoteServiceType = in.RemoteServiceType
	out.Connection = in.Connection
	out.MaxBatchingRows = in.MaxBatchingRows
	out.RemoteModelVersion = in.RemoteModelVersion
	out.SpeechRecognizer = in.SpeechRecognizer
	return out
}
func RestrictionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RestrictionConfig) *krm.RestrictionConfig {
	if in == nil {
		return nil
	}
	out := &krm.RestrictionConfig{}
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func RestrictionConfig_ToProto(mapCtx *direct.MapContext, in *krm.RestrictionConfig) *pb.RestrictionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RestrictionConfig{}
	out.Type = in.Type
	return out
}
func Routine_FromProto(mapCtx *direct.MapContext, in *pb.Routine) *krm.Routine {
	if in == nil {
		return nil
	}
	out := &krm.Routine{}
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.RoutineReference = RoutineReference_FromProto(mapCtx, in.GetRoutineReference())
	out.RoutineType = direct.LazyPtr(in.GetRoutineType())
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.LastModifiedTime = direct.LazyPtr(in.GetLastModifiedTime())
	out.Language = direct.LazyPtr(in.GetLanguage())
	out.Arguments = direct.Slice_FromProto(mapCtx, in.Arguments, Routine_Argument_FromProto)
	out.ReturnType = StandardSqlDataType_FromProto(mapCtx, in.GetReturnType())
	out.ReturnTableType = StandardSqlTableType_FromProto(mapCtx, in.GetReturnTableType())
	out.ImportedLibraries = in.ImportedLibraries
	out.DefinitionBody = direct.LazyPtr(in.GetDefinitionBody())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DeterminismLevel = direct.LazyPtr(in.GetDeterminismLevel())
	out.SecurityMode = direct.LazyPtr(in.GetSecurityMode())
	out.StrictMode = direct.LazyPtr(in.GetStrictMode())
	out.RemoteFunctionOptions = Routine_RemoteFunctionOptions_FromProto(mapCtx, in.GetRemoteFunctionOptions())
	out.SparkOptions = SparkOptions_FromProto(mapCtx, in.GetSparkOptions())
	out.DataGovernanceType = direct.LazyPtr(in.GetDataGovernanceType())
	return out
}
func Routine_ToProto(mapCtx *direct.MapContext, in *krm.Routine) *pb.Routine {
	if in == nil {
		return nil
	}
	out := &pb.Routine{}
	out.Etag = in.Etag
	out.RoutineReference = RoutineReference_ToProto(mapCtx, in.RoutineReference)
	out.RoutineType = in.RoutineType
	out.CreationTime = in.CreationTime
	out.LastModifiedTime = in.LastModifiedTime
	out.Language = in.Language
	out.Arguments = direct.Slice_ToProto(mapCtx, in.Arguments, Routine_Argument_ToProto)
	out.ReturnType = StandardSqlDataType_ToProto(mapCtx, in.ReturnType)
	out.ReturnTableType = StandardSqlTableType_ToProto(mapCtx, in.ReturnTableType)
	out.ImportedLibraries = in.ImportedLibraries
	out.DefinitionBody = in.DefinitionBody
	out.Description = in.Description
	out.DeterminismLevel = in.DeterminismLevel
	out.SecurityMode = in.SecurityMode
	out.StrictMode = in.StrictMode
	out.RemoteFunctionOptions = Routine_RemoteFunctionOptions_ToProto(mapCtx, in.RemoteFunctionOptions)
	out.SparkOptions = SparkOptions_ToProto(mapCtx, in.SparkOptions)
	out.DataGovernanceType = in.DataGovernanceType
	return out
}
func RoutineReference_FromProto(mapCtx *direct.MapContext, in *pb.RoutineReference) *krm.RoutineReference {
	if in == nil {
		return nil
	}
	out := &krm.RoutineReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.RoutineId = in.RoutineId
	return out
}
func RoutineReference_ToProto(mapCtx *direct.MapContext, in *krm.RoutineReference) *pb.RoutineReference {
	if in == nil {
		return nil
	}
	out := &pb.RoutineReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.RoutineId = in.RoutineId
	return out
}
func Routine_Argument_FromProto(mapCtx *direct.MapContext, in *pb.Argument) *krm.Routine_Argument {
	if in == nil {
		return nil
	}
	out := &krm.Routine_Argument{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ArgumentKind = direct.LazyPtr(in.GetArgumentKind())
	out.Mode = direct.LazyPtr(in.GetMode())
	out.DataType = StandardSqlDataType_FromProto(mapCtx, in.GetDataType())
	out.IsAggregate = direct.LazyPtr(in.GetIsAggregate())
	return out
}
func Routine_Argument_ToProto(mapCtx *direct.MapContext, in *krm.Routine_Argument) *pb.Argument {
	if in == nil {
		return nil
	}
	out := &pb.Argument{}
	out.Name = in.Name
	out.ArgumentKind = in.ArgumentKind
	out.Mode = in.Mode
	out.DataType = StandardSqlDataType_ToProto(mapCtx, in.DataType)
	out.IsAggregate = in.IsAggregate
	return out
}
func Routine_RemoteFunctionOptions_FromProto(mapCtx *direct.MapContext, in *pb.RemoteFunctionOptions) *krm.Routine_RemoteFunctionOptions {
	if in == nil {
		return nil
	}
	out := &krm.Routine_RemoteFunctionOptions{}
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	out.Connection = direct.LazyPtr(in.GetConnection())
	out.UserDefinedContext = in.UserDefinedContext
	out.MaxBatchingRows = direct.LazyPtr(in.GetMaxBatchingRows())
	return out
}
func Routine_RemoteFunctionOptions_ToProto(mapCtx *direct.MapContext, in *krm.Routine_RemoteFunctionOptions) *pb.RemoteFunctionOptions {
	if in == nil {
		return nil
	}
	out := &pb.RemoteFunctionOptions{}
	out.Endpoint = in.Endpoint
	out.Connection = in.Connection
	out.UserDefinedContext = in.UserDefinedContext
	out.MaxBatchingRows = in.MaxBatchingRows
	return out
}
func RowAccessPolicy_FromProto(mapCtx *direct.MapContext, in *pb.RowAccessPolicy) *krm.RowAccessPolicy {
	if in == nil {
		return nil
	}
	out := &krm.RowAccessPolicy{}
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.RowAccessPolicyReference = RowAccessPolicyReference_FromProto(mapCtx, in.GetRowAccessPolicyReference())
	out.FilterPredicate = direct.LazyPtr(in.GetFilterPredicate())
	out.CreationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreationTime())
	out.LastModifiedTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastModifiedTime())
	return out
}
func RowAccessPolicy_ToProto(mapCtx *direct.MapContext, in *krm.RowAccessPolicy) *pb.RowAccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.RowAccessPolicy{}
	out.Etag = in.Etag
	out.RowAccessPolicyReference = RowAccessPolicyReference_ToProto(mapCtx, in.RowAccessPolicyReference)
	out.FilterPredicate = in.FilterPredicate
	out.CreationTime = direct.StringTimestamp_ToProto(mapCtx, in.CreationTime)
	out.LastModifiedTime = direct.StringTimestamp_ToProto(mapCtx, in.LastModifiedTime)
	return out
}
func RowAccessPolicyReference_FromProto(mapCtx *direct.MapContext, in *pb.RowAccessPolicyReference) *krm.RowAccessPolicyReference {
	if in == nil {
		return nil
	}
	out := &krm.RowAccessPolicyReference{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	out.TableID = direct.LazyPtr(in.GetTableId())
	out.PolicyID = direct.LazyPtr(in.GetPolicyId())
	return out
}
func RowAccessPolicyReference_ToProto(mapCtx *direct.MapContext, in *krm.RowAccessPolicyReference) *pb.RowAccessPolicyReference {
	if in == nil {
		return nil
	}
	out := &pb.RowAccessPolicyReference{}
	out.ProjectId = in.ProjectID
	out.DatasetId = in.DatasetID
	out.TableId = in.TableID
	out.PolicyId = in.PolicyID
	return out
}
func RowLevelSecurityStatistics_FromProto(mapCtx *direct.MapContext, in *pb.RowLevelSecurityStatistics) *krm.RowLevelSecurityStatistics {
	if in == nil {
		return nil
	}
	out := &krm.RowLevelSecurityStatistics{}
	out.RowLevelSecurityApplied = direct.LazyPtr(in.GetRowLevelSecurityApplied())
	return out
}
func RowLevelSecurityStatistics_ToProto(mapCtx *direct.MapContext, in *krm.RowLevelSecurityStatistics) *pb.RowLevelSecurityStatistics {
	if in == nil {
		return nil
	}
	out := &pb.RowLevelSecurityStatistics{}
	out.RowLevelSecurityApplied = in.RowLevelSecurityApplied
	return out
}
func ScriptOptions_FromProto(mapCtx *direct.MapContext, in *pb.ScriptOptions) *krm.ScriptOptions {
	if in == nil {
		return nil
	}
	out := &krm.ScriptOptions{}
	out.StatementTimeoutMs = direct.LazyPtr(in.GetStatementTimeoutMs())
	out.StatementByteBudget = direct.LazyPtr(in.GetStatementByteBudget())
	out.KeyResultStatement = direct.LazyPtr(in.GetKeyResultStatement())
	return out
}
func ScriptOptions_ToProto(mapCtx *direct.MapContext, in *krm.ScriptOptions) *pb.ScriptOptions {
	if in == nil {
		return nil
	}
	out := &pb.ScriptOptions{}
	out.StatementTimeoutMs = in.StatementTimeoutMs
	out.StatementByteBudget = in.StatementByteBudget
	out.KeyResultStatement = in.KeyResultStatement
	return out
}
func ScriptStatistics_FromProto(mapCtx *direct.MapContext, in *pb.ScriptStatistics) *krm.ScriptStatistics {
	if in == nil {
		return nil
	}
	out := &krm.ScriptStatistics{}
	out.EvaluationKind = direct.LazyPtr(in.GetEvaluationKind())
	out.StackFrames = direct.Slice_FromProto(mapCtx, in.StackFrames, ScriptStatistics_ScriptStackFrame_FromProto)
	return out
}
func ScriptStatistics_ToProto(mapCtx *direct.MapContext, in *krm.ScriptStatistics) *pb.ScriptStatistics {
	if in == nil {
		return nil
	}
	out := &pb.ScriptStatistics{}
	out.EvaluationKind = in.EvaluationKind
	out.StackFrames = direct.Slice_ToProto(mapCtx, in.StackFrames, ScriptStatistics_ScriptStackFrame_ToProto)
	return out
}
func ScriptStatistics_ScriptStackFrame_FromProto(mapCtx *direct.MapContext, in *pb.ScriptStackFrame) *krm.ScriptStatistics_ScriptStackFrame {
	if in == nil {
		return nil
	}
	out := &krm.ScriptStatistics_ScriptStackFrame{}
	out.StartLine = direct.LazyPtr(in.GetStartLine())
	out.StartColumn = direct.LazyPtr(in.GetStartColumn())
	out.EndLine = direct.LazyPtr(in.GetEndLine())
	out.EndColumn = direct.LazyPtr(in.GetEndColumn())
	out.ProcedureID = direct.LazyPtr(in.GetProcedureId())
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func ScriptStatistics_ScriptStackFrame_ToProto(mapCtx *direct.MapContext, in *krm.ScriptStatistics_ScriptStackFrame) *pb.ScriptStackFrame {
	if in == nil {
		return nil
	}
	out := &pb.ScriptStackFrame{}
	out.StartLine = in.StartLine
	out.StartColumn = in.StartColumn
	out.EndLine = in.EndLine
	out.EndColumn = in.EndColumn
	out.ProcedureId = in.ProcedureID
	out.Text = in.Text
	return out
}
func SearchStatistics_FromProto(mapCtx *direct.MapContext, in *pb.SearchStatistics) *krm.SearchStatistics {
	if in == nil {
		return nil
	}
	out := &krm.SearchStatistics{}
	out.IndexUsageMode = direct.LazyPtr(in.GetIndexUsageMode())
	out.IndexUnusedReasons = direct.Slice_FromProto(mapCtx, in.IndexUnusedReasons, IndexUnusedReason_FromProto)
	return out
}
func SearchStatistics_ToProto(mapCtx *direct.MapContext, in *krm.SearchStatistics) *pb.SearchStatistics {
	if in == nil {
		return nil
	}
	out := &pb.SearchStatistics{}
	out.IndexUsageMode = in.IndexUsageMode
	out.IndexUnusedReasons = direct.Slice_ToProto(mapCtx, in.IndexUnusedReasons, IndexUnusedReason_ToProto)
	return out
}
func SerDeInfo_FromProto(mapCtx *direct.MapContext, in *pb.SerDeInfo) *krm.SerDeInfo {
	if in == nil {
		return nil
	}
	out := &krm.SerDeInfo{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SerializationLibrary = direct.LazyPtr(in.GetSerializationLibrary())
	out.Parameters = in.Parameters
	return out
}
func SerDeInfo_ToProto(mapCtx *direct.MapContext, in *krm.SerDeInfo) *pb.SerDeInfo {
	if in == nil {
		return nil
	}
	out := &pb.SerDeInfo{}
	out.Name = in.Name
	out.SerializationLibrary = in.SerializationLibrary
	out.Parameters = in.Parameters
	return out
}
func SessionInfo_FromProto(mapCtx *direct.MapContext, in *pb.SessionInfo) *krm.SessionInfo {
	if in == nil {
		return nil
	}
	out := &krm.SessionInfo{}
	out.SessionID = direct.LazyPtr(in.GetSessionId())
	return out
}
func SessionInfo_ToProto(mapCtx *direct.MapContext, in *krm.SessionInfo) *pb.SessionInfo {
	if in == nil {
		return nil
	}
	out := &pb.SessionInfo{}
	out.SessionId = in.SessionID
	return out
}
func SnapshotDefinition_FromProto(mapCtx *direct.MapContext, in *pb.SnapshotDefinition) *krm.SnapshotDefinition {
	if in == nil {
		return nil
	}
	out := &krm.SnapshotDefinition{}
	out.BaseTableReference = TableReference_FromProto(mapCtx, in.GetBaseTableReference())
	out.SnapshotTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSnapshotTime())
	return out
}
func SnapshotDefinition_ToProto(mapCtx *direct.MapContext, in *krm.SnapshotDefinition) *pb.SnapshotDefinition {
	if in == nil {
		return nil
	}
	out := &pb.SnapshotDefinition{}
	out.BaseTableReference = TableReference_ToProto(mapCtx, in.BaseTableReference)
	out.SnapshotTime = direct.StringTimestamp_ToProto(mapCtx, in.SnapshotTime)
	return out
}
func SparkOptions_FromProto(mapCtx *direct.MapContext, in *pb.SparkOptions) *krm.SparkOptions {
	if in == nil {
		return nil
	}
	out := &krm.SparkOptions{}
	out.Connection = direct.LazyPtr(in.GetConnection())
	out.RuntimeVersion = direct.LazyPtr(in.GetRuntimeVersion())
	out.ContainerImage = direct.LazyPtr(in.GetContainerImage())
	out.Properties = in.Properties
	out.MainFileUri = direct.LazyPtr(in.GetMainFileUri())
	out.PyFileUris = in.PyFileUris
	out.JarUris = in.JarUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	return out
}
func SparkOptions_ToProto(mapCtx *direct.MapContext, in *krm.SparkOptions) *pb.SparkOptions {
	if in == nil {
		return nil
	}
	out := &pb.SparkOptions{}
	out.Connection = in.Connection
	out.RuntimeVersion = in.RuntimeVersion
	out.ContainerImage = in.ContainerImage
	out.Properties = in.Properties
	out.MainFileUri = in.MainFileUri
	out.PyFileUris = in.PyFileUris
	out.JarUris = in.JarUris
	out.FileUris = in.FileUris
	out.ArchiveUris = in.ArchiveUris
	out.MainClass = in.MainClass
	return out
}
func SparkStatistics_FromProto(mapCtx *direct.MapContext, in *pb.SparkStatistics) *krm.SparkStatistics {
	if in == nil {
		return nil
	}
	out := &krm.SparkStatistics{}
	out.SparkJobID = in.SparkJobId
	out.SparkJobLocation = in.SparkJobLocation
	out.Endpoints = in.Endpoints
	out.LoggingInfo = SparkStatistics_LoggingInfo_FromProto(mapCtx, in.GetLoggingInfo())
	out.KmsKeyName = in.KmsKeyName
	out.GcsStagingBucket = in.GcsStagingBucket
	return out
}
func SparkStatistics_ToProto(mapCtx *direct.MapContext, in *krm.SparkStatistics) *pb.SparkStatistics {
	if in == nil {
		return nil
	}
	out := &pb.SparkStatistics{}
	out.SparkJobId = in.SparkJobID
	out.SparkJobLocation = in.SparkJobLocation
	out.Endpoints = in.Endpoints
	if oneof := SparkStatistics_LoggingInfo_ToProto(mapCtx, in.LoggingInfo); oneof != nil {
		out.LoggingInfo = oneof
	}
	out.KmsKeyName = in.KmsKeyName
	out.GcsStagingBucket = in.GcsStagingBucket
	return out
}
func SparkStatistics_LoggingInfo_FromProto(mapCtx *direct.MapContext, in *pb.SparkLoggingInfo) *krm.SparkStatistics_LoggingInfo {
	if in == nil {
		return nil
	}
	out := &krm.SparkStatistics_LoggingInfo{}
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	return out
}
func SparkStatistics_LoggingInfo_ToProto(mapCtx *direct.MapContext, in *krm.SparkStatistics_LoggingInfo) *pb.SparkLoggingInfo {
	if in == nil {
		return nil
	}
	out := &pb.SparkLoggingInfo{}
	out.ResourceType = in.ResourceType
	out.ProjectId = in.ProjectID
	return out
}
func StagePerformanceChangeInsight_FromProto(mapCtx *direct.MapContext, in *pb.StagePerformanceChangeInsight) *krm.StagePerformanceChangeInsight {
	if in == nil {
		return nil
	}
	out := &krm.StagePerformanceChangeInsight{}
	out.StageID = direct.LazyPtr(in.GetStageId())
	out.InputDataChange = InputDataChange_FromProto(mapCtx, in.GetInputDataChange())
	return out
}
func StagePerformanceChangeInsight_ToProto(mapCtx *direct.MapContext, in *krm.StagePerformanceChangeInsight) *pb.StagePerformanceChangeInsight {
	if in == nil {
		return nil
	}
	out := &pb.StagePerformanceChangeInsight{}
	out.StageId = in.StageID
	if oneof := InputDataChange_ToProto(mapCtx, in.InputDataChange); oneof != nil {
		out.InputDataChange = oneof
	}
	return out
}
func StagePerformanceStandaloneInsight_FromProto(mapCtx *direct.MapContext, in *pb.StagePerformanceStandaloneInsight) *krm.StagePerformanceStandaloneInsight {
	if in == nil {
		return nil
	}
	out := &krm.StagePerformanceStandaloneInsight{}
	out.StageID = direct.LazyPtr(in.GetStageId())
	out.SlotContention = in.SlotContention
	out.InsufficientShuffleQuota = in.InsufficientShuffleQuota
	out.BiEngineReasons = direct.Slice_FromProto(mapCtx, in.BiEngineReasons, BiEngineReason_FromProto)
	out.HighCardinalityJoins = direct.Slice_FromProto(mapCtx, in.HighCardinalityJoins, HighCardinalityJoin_FromProto)
	out.PartitionSkew = PartitionSkew_FromProto(mapCtx, in.GetPartitionSkew())
	return out
}
func StagePerformanceStandaloneInsight_ToProto(mapCtx *direct.MapContext, in *krm.StagePerformanceStandaloneInsight) *pb.StagePerformanceStandaloneInsight {
	if in == nil {
		return nil
	}
	out := &pb.StagePerformanceStandaloneInsight{}
	out.StageId = in.StageID
	out.SlotContention = in.SlotContention
	out.InsufficientShuffleQuota = in.InsufficientShuffleQuota
	out.BiEngineReasons = direct.Slice_ToProto(mapCtx, in.BiEngineReasons, BiEngineReason_ToProto)
	out.HighCardinalityJoins = direct.Slice_ToProto(mapCtx, in.HighCardinalityJoins, HighCardinalityJoin_ToProto)
	if oneof := PartitionSkew_ToProto(mapCtx, in.PartitionSkew); oneof != nil {
		out.PartitionSkew = oneof
	}
	return out
}
func StandardSqlDataType_FromProto(mapCtx *direct.MapContext, in *pb.StandardSqlDataType) *krm.StandardSqlDataType {
	if in == nil {
		return nil
	}
	out := &krm.StandardSqlDataType{}
	out.TypeKind = direct.LazyPtr(in.GetTypeKind())
	out.ArrayElementType = StandardSqlDataType_FromProto(mapCtx, in.GetArrayElementType())
	out.StructType = StandardSqlStructType_FromProto(mapCtx, in.GetStructType())
	out.RangeElementType = StandardSqlDataType_FromProto(mapCtx, in.GetRangeElementType())
	return out
}
func StandardSqlDataType_ToProto(mapCtx *direct.MapContext, in *krm.StandardSqlDataType) *pb.StandardSqlDataType {
	if in == nil {
		return nil
	}
	out := &pb.StandardSqlDataType{}
	out.TypeKind = in.TypeKind
	if oneof := StandardSqlDataType_ToProto(mapCtx, in.ArrayElementType); oneof != nil {
		out.ArrayElementType = oneof
	}
	if oneof := StandardSqlStructType_ToProto(mapCtx, in.StructType); oneof != nil {
		out.StructType = oneof
	}
	if oneof := StandardSqlDataType_ToProto(mapCtx, in.RangeElementType); oneof != nil {
		out.RangeElementType = oneof
	}
	return out
}
func StandardSqlField_FromProto(mapCtx *direct.MapContext, in *pb.StandardSqlField) *krm.StandardSqlField {
	if in == nil {
		return nil
	}
	out := &krm.StandardSqlField{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = StandardSqlDataType_FromProto(mapCtx, in.GetType())
	return out
}
func StandardSqlField_ToProto(mapCtx *direct.MapContext, in *krm.StandardSqlField) *pb.StandardSqlField {
	if in == nil {
		return nil
	}
	out := &pb.StandardSqlField{}
	out.Name = in.Name
	out.Type = StandardSqlDataType_ToProto(mapCtx, in.Type)
	return out
}
func StandardSqlStructType_FromProto(mapCtx *direct.MapContext, in *pb.StandardSqlStructType) *krm.StandardSqlStructType {
	if in == nil {
		return nil
	}
	out := &krm.StandardSqlStructType{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, StandardSqlField_FromProto)
	return out
}
func StandardSqlStructType_ToProto(mapCtx *direct.MapContext, in *krm.StandardSqlStructType) *pb.StandardSqlStructType {
	if in == nil {
		return nil
	}
	out := &pb.StandardSqlStructType{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, StandardSqlField_ToProto)
	return out
}
func StandardSqlTableType_FromProto(mapCtx *direct.MapContext, in *pb.StandardSqlTableType) *krm.StandardSqlTableType {
	if in == nil {
		return nil
	}
	out := &krm.StandardSqlTableType{}
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, StandardSqlField_FromProto)
	return out
}
func StandardSqlTableType_ToProto(mapCtx *direct.MapContext, in *krm.StandardSqlTableType) *pb.StandardSqlTableType {
	if in == nil {
		return nil
	}
	out := &pb.StandardSqlTableType{}
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, StandardSqlField_ToProto)
	return out
}
func StorageDescriptor_FromProto(mapCtx *direct.MapContext, in *pb.StorageDescriptor) *krm.StorageDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.StorageDescriptor{}
	out.LocationUri = direct.LazyPtr(in.GetLocationUri())
	out.InputFormat = direct.LazyPtr(in.GetInputFormat())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.SerdeInfo = SerDeInfo_FromProto(mapCtx, in.GetSerdeInfo())
	return out
}
func StorageDescriptor_ToProto(mapCtx *direct.MapContext, in *krm.StorageDescriptor) *pb.StorageDescriptor {
	if in == nil {
		return nil
	}
	out := &pb.StorageDescriptor{}
	out.LocationUri = in.LocationUri
	out.InputFormat = in.InputFormat
	out.OutputFormat = in.OutputFormat
	out.SerdeInfo = SerDeInfo_ToProto(mapCtx, in.SerdeInfo)
	return out
}
func Streamingbuffer_FromProto(mapCtx *direct.MapContext, in *pb.Streamingbuffer) *krm.Streamingbuffer {
	if in == nil {
		return nil
	}
	out := &krm.Streamingbuffer{}
	out.EstimatedBytes = direct.LazyPtr(in.GetEstimatedBytes())
	out.EstimatedRows = direct.LazyPtr(in.GetEstimatedRows())
	out.OldestEntryTime = direct.LazyPtr(in.GetOldestEntryTime())
	return out
}
func Streamingbuffer_ToProto(mapCtx *direct.MapContext, in *krm.Streamingbuffer) *pb.Streamingbuffer {
	if in == nil {
		return nil
	}
	out := &pb.Streamingbuffer{}
	out.EstimatedBytes = in.EstimatedBytes
	out.EstimatedRows = in.EstimatedRows
	out.OldestEntryTime = in.OldestEntryTime
	return out
}
func SystemVariables_FromProto(mapCtx *direct.MapContext, in *pb.SystemVariables) *krm.SystemVariables {
	if in == nil {
		return nil
	}
	out := &krm.SystemVariables{}
	/* NOTYET */
	return out
}
func SystemVariables_ToProto(mapCtx *direct.MapContext, in *krm.SystemVariables) *pb.SystemVariables {
	if in == nil {
		return nil
	}
	out := &pb.SystemVariables{}
	/* NOTYET */
	return out
}
func Table_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.Table {
	if in == nil {
		return nil
	}
	out := &krm.Table{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ID = direct.LazyPtr(in.GetId())
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.TableReference = TableReference_FromProto(mapCtx, in.GetTableReference())
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Schema = TableSchema_FromProto(mapCtx, in.GetSchema())
	out.TimePartitioning = TimePartitioning_FromProto(mapCtx, in.GetTimePartitioning())
	out.RangePartitioning = RangePartitioning_FromProto(mapCtx, in.GetRangePartitioning())
	out.Clustering = Clustering_FromProto(mapCtx, in.GetClustering())
	out.RequirePartitionFilter = direct.LazyPtr(in.GetRequirePartitionFilter())
	out.PartitionDefinition = PartitioningDefinition_FromProto(mapCtx, in.GetPartitionDefinition())
	out.NumBytes = direct.LazyPtr(in.GetNumBytes())
	out.NumPhysicalBytes = direct.LazyPtr(in.GetNumPhysicalBytes())
	out.NumLongTermBytes = direct.LazyPtr(in.GetNumLongTermBytes())
	out.NumRows = direct.LazyPtr(in.GetNumRows())
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.ExpirationTime = direct.LazyPtr(in.GetExpirationTime())
	out.LastModifiedTime = direct.LazyPtr(in.GetLastModifiedTime())
	out.Type = direct.LazyPtr(in.GetType())
	out.View = ViewDefinition_FromProto(mapCtx, in.GetView())
	out.MaterializedView = MaterializedViewDefinition_FromProto(mapCtx, in.GetMaterializedView())
	out.MaterializedViewStatus = MaterializedViewStatus_FromProto(mapCtx, in.GetMaterializedViewStatus())
	out.ExternalDataConfiguration = ExternalDataConfiguration_FromProto(mapCtx, in.GetExternalDataConfiguration())
	out.BiglakeConfiguration = BigLakeConfiguration_FromProto(mapCtx, in.GetBiglakeConfiguration())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.StreamingBuffer = Streamingbuffer_FromProto(mapCtx, in.GetStreamingBuffer())
	out.EncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetEncryptionConfiguration())
	out.SnapshotDefinition = SnapshotDefinition_FromProto(mapCtx, in.GetSnapshotDefinition())
	out.DefaultCollation = direct.LazyPtr(in.GetDefaultCollation())
	out.DefaultRoundingMode = direct.LazyPtr(in.GetDefaultRoundingMode())
	out.CloneDefinition = CloneDefinition_FromProto(mapCtx, in.GetCloneDefinition())
	out.NumTimeTravelPhysicalBytes = direct.LazyPtr(in.GetNumTimeTravelPhysicalBytes())
	out.NumTotalLogicalBytes = direct.LazyPtr(in.GetNumTotalLogicalBytes())
	out.NumActiveLogicalBytes = direct.LazyPtr(in.GetNumActiveLogicalBytes())
	out.NumLongTermLogicalBytes = direct.LazyPtr(in.GetNumLongTermLogicalBytes())
	out.NumCurrentPhysicalBytes = direct.LazyPtr(in.GetNumPhysicalBytes())
	out.NumTotalPhysicalBytes = direct.LazyPtr(in.GetNumTotalPhysicalBytes())
	out.NumActivePhysicalBytes = direct.LazyPtr(in.GetNumActivePhysicalBytes())
	out.NumLongTermPhysicalBytes = direct.LazyPtr(in.GetNumLongTermPhysicalBytes())
	out.NumPartitions = direct.LazyPtr(in.GetNumPartitions())
	out.MaxStaleness = direct.LazyPtr(in.GetMaxStaleness())
	out.Restrictions = RestrictionConfig_FromProto(mapCtx, in.GetRestrictions())
	out.TableConstraints = TableConstraints_FromProto(mapCtx, in.GetTableConstraints())
	out.ResourceTags = in.ResourceTags
	out.TableReplicationInfo = TableReplicationInfo_FromProto(mapCtx, in.GetTableReplicationInfo())
	out.Replicas = direct.Slice_FromProto(mapCtx, in.Replicas, TableReference_FromProto)
	out.ExternalCatalogTableOptions = ExternalCatalogTableOptions_FromProto(mapCtx, in.GetExternalCatalogTableOptions())
	return out
}
func Table_ToProto(mapCtx *direct.MapContext, in *krm.Table) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.Kind = in.Kind
	out.Etag = in.Etag
	out.Id = in.ID
	out.SelfLink = in.SelfLink
	out.TableReference = TableReference_ToProto(mapCtx, in.TableReference)
	out.FriendlyName = in.FriendlyName
	out.Description = in.Description
	out.Labels = in.Labels
	out.Schema = TableSchema_ToProto(mapCtx, in.Schema)
	out.TimePartitioning = TimePartitioning_ToProto(mapCtx, in.TimePartitioning)
	out.RangePartitioning = RangePartitioning_ToProto(mapCtx, in.RangePartitioning)
	out.Clustering = Clustering_ToProto(mapCtx, in.Clustering)
	out.RequirePartitionFilter = in.RequirePartitionFilter
	if oneof := PartitioningDefinition_ToProto(mapCtx, in.PartitionDefinition); oneof != nil {
		out.PartitionDefinition = oneof
	}
	out.NumBytes = in.NumBytes
	out.NumPhysicalBytes = in.NumPhysicalBytes
	out.NumLongTermBytes = in.NumLongTermBytes
	out.NumRows = in.NumRows
	out.CreationTime = in.CreationTime
	out.ExpirationTime = in.ExpirationTime
	out.LastModifiedTime = in.LastModifiedTime
	out.Type = in.Type
	out.View = ViewDefinition_ToProto(mapCtx, in.View)
	out.MaterializedView = MaterializedViewDefinition_ToProto(mapCtx, in.MaterializedView)
	out.MaterializedViewStatus = MaterializedViewStatus_ToProto(mapCtx, in.MaterializedViewStatus)
	out.ExternalDataConfiguration = ExternalDataConfiguration_ToProto(mapCtx, in.ExternalDataConfiguration)
	out.BiglakeConfiguration = BigLakeConfiguration_ToProto(mapCtx, in.BiglakeConfiguration)
	out.Location = in.Location
	out.StreamingBuffer = Streamingbuffer_ToProto(mapCtx, in.StreamingBuffer)
	out.EncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.EncryptionConfiguration)
	out.SnapshotDefinition = SnapshotDefinition_ToProto(mapCtx, in.SnapshotDefinition)
	out.DefaultCollation = in.DefaultCollation
	out.DefaultRoundingMode = in.DefaultRoundingMode
	out.CloneDefinition = CloneDefinition_ToProto(mapCtx, in.CloneDefinition)
	out.NumTimeTravelPhysicalBytes = in.NumTimeTravelPhysicalBytes
	out.NumTotalLogicalBytes = in.NumTotalLogicalBytes
	out.NumActiveLogicalBytes = in.NumActiveLogicalBytes
	out.NumLongTermLogicalBytes = in.NumLongTermLogicalBytes
	out.NumPhysicalBytes = in.NumCurrentPhysicalBytes
	out.NumTotalPhysicalBytes = in.NumTotalPhysicalBytes
	out.NumActivePhysicalBytes = in.NumActivePhysicalBytes
	out.NumLongTermPhysicalBytes = in.NumLongTermPhysicalBytes
	out.NumPartitions = in.NumPartitions
	out.MaxStaleness = in.MaxStaleness
	out.Restrictions = RestrictionConfig_ToProto(mapCtx, in.Restrictions)
	out.TableConstraints = TableConstraints_ToProto(mapCtx, in.TableConstraints)
	out.ResourceTags = in.ResourceTags
	out.TableReplicationInfo = TableReplicationInfo_ToProto(mapCtx, in.TableReplicationInfo)
	out.Replicas = direct.Slice_ToProto(mapCtx, in.Replicas, TableReference_ToProto)
	out.ExternalCatalogTableOptions = ExternalCatalogTableOptions_ToProto(mapCtx, in.ExternalCatalogTableOptions)
	return out
}
func TableConstraints_FromProto(mapCtx *direct.MapContext, in *pb.TableConstraints) *krm.TableConstraints {
	if in == nil {
		return nil
	}
	out := &krm.TableConstraints{}
	out.PrimaryKey = PrimaryKey_FromProto(mapCtx, in.GetPrimaryKey())
	out.ForeignKeys = direct.Slice_FromProto(mapCtx, in.ForeignKeys, ForeignKey_FromProto)
	return out
}
func TableConstraints_ToProto(mapCtx *direct.MapContext, in *krm.TableConstraints) *pb.TableConstraints {
	if in == nil {
		return nil
	}
	out := &pb.TableConstraints{}
	out.PrimaryKey = PrimaryKey_ToProto(mapCtx, in.PrimaryKey)
	out.ForeignKeys = direct.Slice_ToProto(mapCtx, in.ForeignKeys, ForeignKey_ToProto)
	return out
}
func TableFieldSchema_FromProto(mapCtx *direct.MapContext, in *pb.TableFieldSchema) *krm.TableFieldSchema {
	if in == nil {
		return nil
	}
	out := &krm.TableFieldSchema{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Mode = direct.LazyPtr(in.GetMode())
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, TableFieldSchema_FromProto)
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PolicyTags = TableFieldSchema_PolicyTagList_FromProto(mapCtx, in.GetPolicyTags())
	out.MaxLength = direct.LazyPtr(in.GetMaxLength())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Scale = direct.LazyPtr(in.GetScale())
	out.RoundingMode = direct.LazyPtr(in.GetRoundingMode())
	out.Collation = direct.LazyPtr(in.GetCollation())
	out.DefaultValueExpression = direct.LazyPtr(in.GetDefaultValueExpression())
	out.RangeElementType = TableFieldSchema_FieldElementType_FromProto(mapCtx, in.GetRangeElementType())
	out.ForeignTypeDefinition = direct.LazyPtr(in.GetForeignTypeDefinition())
	return out
}
func TableFieldSchema_ToProto(mapCtx *direct.MapContext, in *krm.TableFieldSchema) *pb.TableFieldSchema {
	if in == nil {
		return nil
	}
	out := &pb.TableFieldSchema{}
	out.Name = in.Name
	out.Type = in.Type
	out.Mode = in.Mode
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, TableFieldSchema_ToProto)
	out.Description = in.Description
	out.PolicyTags = TableFieldSchema_PolicyTagList_ToProto(mapCtx, in.PolicyTags)
	out.MaxLength = in.MaxLength
	out.Precision = in.Precision
	out.Scale = in.Scale
	out.RoundingMode = in.RoundingMode
	out.Collation = in.Collation
	out.DefaultValueExpression = in.DefaultValueExpression
	out.RangeElementType = TableFieldSchema_FieldElementType_ToProto(mapCtx, in.RangeElementType)
	out.ForeignTypeDefinition = in.ForeignTypeDefinition
	return out
}
func TableFieldSchema_FieldElementType_FromProto(mapCtx *direct.MapContext, in *pb.TableFieldSchemaRangeElementType) *krm.TableFieldSchema_FieldElementType {
	if in == nil {
		return nil
	}
	out := &krm.TableFieldSchema_FieldElementType{}
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func TableFieldSchema_FieldElementType_ToProto(mapCtx *direct.MapContext, in *krm.TableFieldSchema_FieldElementType) *pb.TableFieldSchemaRangeElementType {
	if in == nil {
		return nil
	}
	out := &pb.TableFieldSchemaRangeElementType{}
	out.Type = in.Type
	return out
}
func TableFieldSchema_PolicyTagList_FromProto(mapCtx *direct.MapContext, in *pb.TableFieldSchemaPolicyTags) *krm.TableFieldSchema_PolicyTagList {
	if in == nil {
		return nil
	}
	out := &krm.TableFieldSchema_PolicyTagList{}
	out.Names = in.Names
	return out
}
func TableFieldSchema_PolicyTagList_ToProto(mapCtx *direct.MapContext, in *krm.TableFieldSchema_PolicyTagList) *pb.TableFieldSchemaPolicyTags {
	if in == nil {
		return nil
	}
	out := &pb.TableFieldSchemaPolicyTags{}
	out.Names = in.Names
	return out
}
func TableList_FromProto(mapCtx *direct.MapContext, in *pb.TableList) *krm.TableList {
	if in == nil {
		return nil
	}
	out := &krm.TableList{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.NextPageToken = direct.LazyPtr(in.GetNextPageToken())
	out.Tables = direct.Slice_FromProto(mapCtx, in.Tables, ListFormatTable_FromProto)
	out.TotalItems = direct.LazyPtr(in.GetTotalItems())
	return out
}
func TableList_ToProto(mapCtx *direct.MapContext, in *krm.TableList) *pb.TableList {
	if in == nil {
		return nil
	}
	out := &pb.TableList{}
	out.Kind = in.Kind
	out.Etag = in.Etag
	out.NextPageToken = in.NextPageToken
	out.Tables = direct.Slice_ToProto(mapCtx, in.Tables, ListFormatTable_ToProto)
	out.TotalItems = in.TotalItems
	return out
}
func TableMetadataCacheUsage_FromProto(mapCtx *direct.MapContext, in *pb.TableMetadataCacheUsage) *krm.TableMetadataCacheUsage {
	if in == nil {
		return nil
	}
	out := &krm.TableMetadataCacheUsage{}
	out.TableReference = TableReference_FromProto(mapCtx, in.GetTableReference())
	out.UnusedReason = direct.LazyPtr(in.GetUnusedReason())
	out.Explanation = in.Explanation
	out.TableType = direct.LazyPtr(in.GetTableType())
	return out
}
func TableMetadataCacheUsage_ToProto(mapCtx *direct.MapContext, in *krm.TableMetadataCacheUsage) *pb.TableMetadataCacheUsage {
	if in == nil {
		return nil
	}
	out := &pb.TableMetadataCacheUsage{}
	if oneof := TableReference_ToProto(mapCtx, in.TableReference); oneof != nil {
		out.TableReference = oneof
	}
	out.UnusedReason = in.UnusedReason
	out.Explanation = in.Explanation
	out.TableType = in.TableType
	return out
}
func TableReference_FromProto(mapCtx *direct.MapContext, in *pb.TableReference) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func TableReference_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.TableReference {
	if in == nil {
		return nil
	}
	out := &pb.TableReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func TableReplicationInfo_FromProto(mapCtx *direct.MapContext, in *pb.TableReplicationInfo) *krm.TableReplicationInfo {
	if in == nil {
		return nil
	}
	out := &krm.TableReplicationInfo{}
	out.SourceTable = TableReference_FromProto(mapCtx, in.GetSourceTable())
	out.ReplicationIntervalMs = direct.LazyPtr(in.GetReplicationIntervalMs())
	out.ReplicatedSourceLastRefreshTime = direct.LazyPtr(in.GetReplicatedSourceLastRefreshTime())
	out.ReplicationStatus = direct.LazyPtr(in.GetReplicationStatus())
	out.ReplicationError = ErrorProto_FromProto(mapCtx, in.GetReplicationError())
	return out
}
func TableReplicationInfo_ToProto(mapCtx *direct.MapContext, in *krm.TableReplicationInfo) *pb.TableReplicationInfo {
	if in == nil {
		return nil
	}
	out := &pb.TableReplicationInfo{}
	out.SourceTable = TableReference_ToProto(mapCtx, in.SourceTable)
	out.ReplicationIntervalMs = in.ReplicationIntervalMs
	out.ReplicatedSourceLastRefreshTime = in.ReplicatedSourceLastRefreshTime
	out.ReplicationStatus = in.ReplicationStatus
	out.ReplicationError = ErrorProto_ToProto(mapCtx, in.ReplicationError)
	return out
}
func TableSchema_FromProto(mapCtx *direct.MapContext, in *pb.TableSchema) *krm.TableSchema {
	if in == nil {
		return nil
	}
	out := &krm.TableSchema{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, TableFieldSchema_FromProto)
	out.ForeignTypeInfo = ForeignTypeInfo_FromProto(mapCtx, in.GetForeignTypeInfo())
	return out
}
func TableSchema_ToProto(mapCtx *direct.MapContext, in *krm.TableSchema) *pb.TableSchema {
	if in == nil {
		return nil
	}
	out := &pb.TableSchema{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, TableFieldSchema_ToProto)
	out.ForeignTypeInfo = ForeignTypeInfo_ToProto(mapCtx, in.ForeignTypeInfo)
	return out
}
func TimePartitioning_FromProto(mapCtx *direct.MapContext, in *pb.TimePartitioning) *krm.TimePartitioning {
	if in == nil {
		return nil
	}
	out := &krm.TimePartitioning{}
	out.Type = direct.LazyPtr(in.GetType())
	out.ExpirationMs = direct.LazyPtr(in.GetExpirationMs())
	out.Field = direct.LazyPtr(in.GetField())
	return out
}
func TimePartitioning_ToProto(mapCtx *direct.MapContext, in *krm.TimePartitioning) *pb.TimePartitioning {
	if in == nil {
		return nil
	}
	out := &pb.TimePartitioning{}
	out.Type = in.Type
	out.ExpirationMs = in.ExpirationMs
	out.Field = in.Field
	return out
}
func TransformColumn_FromProto(mapCtx *direct.MapContext, in *pb.TransformColumn) *krm.TransformColumn {
	if in == nil {
		return nil
	}
	out := &krm.TransformColumn{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = StandardSqlDataType_FromProto(mapCtx, in.GetType())
	out.TransformSql = direct.LazyPtr(in.GetTransformSql())
	return out
}
func TransformColumn_ToProto(mapCtx *direct.MapContext, in *krm.TransformColumn) *pb.TransformColumn {
	if in == nil {
		return nil
	}
	out := &pb.TransformColumn{}
	out.Name = in.Name
	out.Type = StandardSqlDataType_ToProto(mapCtx, in.Type)
	out.TransformSql = in.TransformSql
	return out
}
func UserDefinedFunctionResource_FromProto(mapCtx *direct.MapContext, in *pb.UserDefinedFunctionResource) *krm.UserDefinedFunctionResource {
	if in == nil {
		return nil
	}
	out := &krm.UserDefinedFunctionResource{}
	out.ResourceUri = direct.LazyPtr(in.GetResourceUri())
	out.InlineCode = direct.LazyPtr(in.GetInlineCode())
	return out
}
func UserDefinedFunctionResource_ToProto(mapCtx *direct.MapContext, in *krm.UserDefinedFunctionResource) *pb.UserDefinedFunctionResource {
	if in == nil {
		return nil
	}
	out := &pb.UserDefinedFunctionResource{}
	out.ResourceUri = in.ResourceUri
	out.InlineCode = in.InlineCode
	return out
}
func VectorSearchStatistics_FromProto(mapCtx *direct.MapContext, in *pb.VectorSearchStatistics) *krm.VectorSearchStatistics {
	if in == nil {
		return nil
	}
	out := &krm.VectorSearchStatistics{}
	out.IndexUsageMode = direct.LazyPtr(in.GetIndexUsageMode())
	out.IndexUnusedReasons = direct.Slice_FromProto(mapCtx, in.IndexUnusedReasons, IndexUnusedReason_FromProto)
	return out
}
func VectorSearchStatistics_ToProto(mapCtx *direct.MapContext, in *krm.VectorSearchStatistics) *pb.VectorSearchStatistics {
	if in == nil {
		return nil
	}
	out := &pb.VectorSearchStatistics{}
	out.IndexUsageMode = in.IndexUsageMode
	out.IndexUnusedReasons = direct.Slice_ToProto(mapCtx, in.IndexUnusedReasons, IndexUnusedReason_ToProto)
	return out
}
func ViewDefinition_FromProto(mapCtx *direct.MapContext, in *pb.ViewDefinition) *krm.ViewDefinition {
	if in == nil {
		return nil
	}
	out := &krm.ViewDefinition{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.UserDefinedFunctionResources = direct.Slice_FromProto(mapCtx, in.UserDefinedFunctionResources, UserDefinedFunctionResource_FromProto)
	out.UseLegacySql = direct.LazyPtr(in.GetUseLegacySql())
	out.UseExplicitColumnNames = direct.LazyPtr(in.GetUseExplicitColumnNames())
	out.PrivacyPolicy = PrivacyPolicy_FromProto(mapCtx, in.GetPrivacyPolicy())
	return out
}
func ViewDefinition_ToProto(mapCtx *direct.MapContext, in *krm.ViewDefinition) *pb.ViewDefinition {
	if in == nil {
		return nil
	}
	out := &pb.ViewDefinition{}
	out.Query = in.Query
	out.UserDefinedFunctionResources = direct.Slice_ToProto(mapCtx, in.UserDefinedFunctionResources, UserDefinedFunctionResource_ToProto)
	out.UseLegacySql = in.UseLegacySql
	out.UseExplicitColumnNames = in.UseExplicitColumnNames
	out.PrivacyPolicy = PrivacyPolicy_ToProto(mapCtx, in.PrivacyPolicy)
	return out
}
