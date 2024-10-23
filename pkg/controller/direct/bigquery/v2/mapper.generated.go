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
	out.View = ReferencedTable_FromProto(mapCtx, in.GetView())
	out.Routine = ReferencedRoutine_FromProto(mapCtx, in.GetRoutine())
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
	out.View = ReferencedTable_ToProto(mapCtx, in.View)
	out.Routine = ReferencedRoutine_ToProto(mapCtx, in.Routine)
	out.Dataset = DatasetAccessEntry_ToProto(mapCtx, in.Dataset)
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
	out.SatisfiesPzs.Value = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi.Value = direct.LazyPtr(in.GetSatisfiesPzi())
	out.Type = direct.LazyPtr(in.GetType())
	out.LinkedDatasetSource = LinkedDatasetSource_FromProto(mapCtx, in.GetLinkedDatasetSource())
	out.LinkedDatasetMetadata = LinkedDatasetMetadata_FromProto(mapCtx, in.GetLinkedDatasetMetadata())
	out.ExternalDatasetReference = ExternalDatasetReference_FromProto(mapCtx, in.GetExternalDatasetReference())
	out.ExternalCatalogDatasetOptions = ExternalCatalogDatasetOptions_FromProto(mapCtx, in.GetExternalCatalogDatasetOptions())
	out.IsCaseInsensitive.Value = direct.LazyPtr(in.GetIsCaseInsensitive())
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
	out.SatisfiesPzs = in.SatisfiesPzs.Value
	out.SatisfiesPzi = in.SatisfiesPzi.Value
	out.Type = in.Type
	out.LinkedDatasetSource = LinkedDatasetSource_ToProto(mapCtx, in.LinkedDatasetSource)
	out.LinkedDatasetMetadata = LinkedDatasetMetadata_ToProto(mapCtx, in.LinkedDatasetMetadata)
	out.ExternalDatasetReference = ExternalDatasetReference_ToProto(mapCtx, in.ExternalDatasetReference)
	out.ExternalCatalogDatasetOptions = ExternalCatalogDatasetOptions_ToProto(mapCtx, in.ExternalCatalogDatasetOptions)
	out.IsCaseInsensitive = in.IsCaseInsensitive.Value
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
func ReferencedTable_FromProto(mapCtx *direct.MapContext, in *pb.TableReference) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func ReferencedTable_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.TableReference {
	if in == nil {
		return nil
	}
	out := &pb.TableReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func ReferencedRoutine_FromProto(mapCtx *direct.MapContext, in *pb.RoutineReference) *krm.RoutineReference {
	if in == nil {
		return nil
	}
	out := &krm.RoutineReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.RoutineId = in.RoutineId
	return out
}
func ReferencedRoutine_ToProto(mapCtx *direct.MapContext, in *krm.RoutineReference) *pb.RoutineReference {
	if in == nil {
		return nil
	}
	out := &pb.RoutineReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.RoutineId = in.RoutineId
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
