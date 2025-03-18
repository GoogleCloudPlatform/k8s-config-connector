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

package bigtable

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func AuthorizedView_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krm.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizedView{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SubsetView = AuthorizedView_SubsetView_FromProto(mapCtx, in.GetSubsetView())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func AuthorizedView_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizedView) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := AuthorizedView_SubsetView_ToProto(mapCtx, in.SubsetView); oneof != nil {
		out.AuthorizedView = &pb.AuthorizedView_SubsetView_{SubsetView: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}
func AuthorizedView_FamilySubsets_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView_FamilySubsets) *krm.AuthorizedView_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizedView_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func AuthorizedView_FamilySubsets_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizedView_FamilySubsets) *pb.AuthorizedView_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func AuthorizedView_SubsetView_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView_SubsetView) *krm.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}
func AuthorizedView_SubsetView_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizedView_SubsetView) *pb.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}
func AutoscalingLimits_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingLimits) *krm.AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingLimits{}
	out.MinServeNodes = direct.LazyPtr(in.GetMinServeNodes())
	out.MaxServeNodes = direct.LazyPtr(in.GetMaxServeNodes())
	return out
}
func AutoscalingLimits_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingLimits) *pb.AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingLimits{}
	out.MinServeNodes = direct.ValueOf(in.MinServeNodes)
	out.MaxServeNodes = direct.ValueOf(in.MaxServeNodes)
	return out
}
func AutoscalingTargets_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingTargets) *krm.AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingTargets{}
	out.CpuUtilizationPercent = direct.LazyPtr(in.GetCpuUtilizationPercent())
	out.StorageUtilizationGibPerNode = direct.LazyPtr(in.GetStorageUtilizationGibPerNode())
	return out
}
func AutoscalingTargets_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingTargets) *pb.AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingTargets{}
	out.CpuUtilizationPercent = direct.ValueOf(in.CpuUtilizationPercent)
	out.StorageUtilizationGibPerNode = direct.ValueOf(in.StorageUtilizationGibPerNode)
	return out
}
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	out.ExpireTime = Backup_ExpireTime_FromProto(mapCtx, in.GetExpireTime())
	out.StartTime = Backup_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = Backup_EndTime_FromProto(mapCtx, in.GetEndTime())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.EncryptionInfo = EncryptionInfo_FromProto(mapCtx, in.GetEncryptionInfo())
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	out.SourceTable = direct.ValueOf(in.SourceTable)
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	out.ExpireTime = Backup_ExpireTime_ToProto(mapCtx, in.ExpireTime)
	out.StartTime = Backup_StartTime_ToProto(mapCtx, in.StartTime)
	out.EndTime = Backup_EndTime_ToProto(mapCtx, in.EndTime)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.EncryptionInfo = EncryptionInfo_ToProto(mapCtx, in.EncryptionInfo)
	return out
}
func BackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krm.BackupInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupInfo{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.StartTime = BackupInfo_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = BackupInfo_EndTime_FromProto(mapCtx, in.GetEndTime())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	return out
}
func BackupInfo_ToProto(mapCtx *direct.MapContext, in *krm.BackupInfo) *pb.BackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupInfo{}
	out.Backup = direct.ValueOf(in.Backup)
	out.StartTime = BackupInfo_StartTime_ToProto(mapCtx, in.StartTime)
	out.EndTime = BackupInfo_EndTime_ToProto(mapCtx, in.EndTime)
	out.SourceTable = direct.ValueOf(in.SourceTable)
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	return out
}
func ChangeStreamConfig_FromProto(mapCtx *direct.MapContext, in *pb.ChangeStreamConfig) *krm.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &krm.ChangeStreamConfig{}
	out.RetentionPeriod = ChangeStreamConfig_RetentionPeriod_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func ChangeStreamConfig_ToProto(mapCtx *direct.MapContext, in *krm.ChangeStreamConfig) *pb.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &pb.ChangeStreamConfig{}
	out.RetentionPeriod = ChangeStreamConfig_RetentionPeriod_ToProto(mapCtx, in.RetentionPeriod)
	return out
}
func Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.Cluster {
	if in == nil {
		return nil
	}
	out := &krm.Cluster{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServeNodes = direct.LazyPtr(in.GetServeNodes())
	out.ClusterConfig = Cluster_ClusterConfig_FromProto(mapCtx, in.GetClusterConfig())
	out.DefaultStorageType = direct.Enum_FromProto(mapCtx, in.GetDefaultStorageType())
	out.EncryptionConfig = Cluster_EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func Cluster_ToProto(mapCtx *direct.MapContext, in *krm.Cluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Name = direct.ValueOf(in.Name)
	out.Location = direct.ValueOf(in.Location)
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	out.ServeNodes = direct.ValueOf(in.ServeNodes)
	if oneof := Cluster_ClusterConfig_ToProto(mapCtx, in.ClusterConfig); oneof != nil {
		out.Config = &pb.Cluster_ClusterConfig_{ClusterConfig: oneof}
	}
	out.DefaultStorageType = direct.Enum_ToProto[pb.StorageType](mapCtx, in.DefaultStorageType)
	out.EncryptionConfig = Cluster_EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func Cluster_ClusterAutoscalingConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ClusterAutoscalingConfig) *krm.Cluster_ClusterAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ClusterAutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingLimits_FromProto(mapCtx, in.GetAutoscalingLimits())
	out.AutoscalingTargets = AutoscalingTargets_FromProto(mapCtx, in.GetAutoscalingTargets())
	return out
}
func Cluster_ClusterAutoscalingConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ClusterAutoscalingConfig) *pb.Cluster_ClusterAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ClusterAutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingLimits_ToProto(mapCtx, in.AutoscalingLimits)
	out.AutoscalingTargets = AutoscalingTargets_ToProto(mapCtx, in.AutoscalingTargets)
	return out
}
func Cluster_ClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ClusterConfig) *krm.Cluster_ClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ClusterConfig{}
	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_FromProto(mapCtx, in.GetClusterAutoscalingConfig())
	return out
}
func Cluster_ClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ClusterConfig) *pb.Cluster_ClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ClusterConfig{}
	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_ToProto(mapCtx, in.ClusterAutoscalingConfig)
	return out
}
func Cluster_EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_EncryptionConfig) *krm.Cluster_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_EncryptionConfig{}
	out.KmsKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func Cluster_EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_EncryptionConfig) *pb.Cluster_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_EncryptionConfig{}
	out.KmsKeyName = direct.ValueOf(in.KmsKeyName)
	return out
}
func ColumnFamily_FromProto(mapCtx *direct.MapContext, in *pb.ColumnFamily) *krm.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &krm.ColumnFamily{}
	out.GcRule = GcRule_FromProto(mapCtx, in.GetGcRule())
	out.ValueType = Type_FromProto(mapCtx, in.GetValueType())
	return out
}
func ColumnFamily_ToProto(mapCtx *direct.MapContext, in *krm.ColumnFamily) *pb.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &pb.ColumnFamily{}
	out.GcRule = GcRule_ToProto(mapCtx, in.GcRule)
	out.ValueType = Type_ToProto(mapCtx, in.ValueType)
	return out
}
func DataBoostReadLocalWrites_FromProto(mapCtx *direct.MapContext, in *pb.DataBoostReadLocalWrites) *krm.DataBoostReadLocalWrites {
	if in == nil {
		return nil
	}
	out := &krm.DataBoostReadLocalWrites{}
	return out
}
func DataBoostReadLocalWrites_ToProto(mapCtx *direct.MapContext, in *krm.DataBoostReadLocalWrites) *pb.DataBoostReadLocalWrites {
	if in == nil {
		return nil
	}
	out := &pb.DataBoostReadLocalWrites{}
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	out.KmsKeyVersion = direct.LazyPtr(in.GetKmsKeyVersion())
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	// MISSING: EncryptionStatus
	out.KmsKeyVersion = direct.ValueOf(in.KmsKeyVersion)
	return out
}
func GcRule_FromProto(mapCtx *direct.MapContext, in *pb.GcRule) *krm.GcRule {
	if in == nil {
		return nil
	}
	out := &krm.GcRule{}
	out.MaxNumVersions = direct.LazyPtr(in.GetMaxNumVersions())
	out.MaxAge = GcRule_MaxAge_FromProto(mapCtx, in.GetMaxAge())
	out.Intersection = GcRule_Intersection_FromProto(mapCtx, in.GetIntersection())
	out.Union = GcRule_Union_FromProto(mapCtx, in.GetUnion())
	return out
}
func GcRule_ToProto(mapCtx *direct.MapContext, in *krm.GcRule) *pb.GcRule {
	if in == nil {
		return nil
	}
	out := &pb.GcRule{}
	if oneof := GcRule_MaxNumVersions_ToProto(mapCtx, in.MaxNumVersions); oneof != nil {
		out.Rule = oneof
	}
	if oneof := GcRule_MaxAge_ToProto(mapCtx, in.MaxAge); oneof != nil {
		out.Rule = &pb.GcRule_MaxAge{MaxAge: oneof}
	}
	if oneof := GcRule_Intersection_ToProto(mapCtx, in.Intersection); oneof != nil {
		out.Rule = &pb.GcRule_Intersection_{Intersection: oneof}
	}
	if oneof := GcRule_Union_ToProto(mapCtx, in.Union); oneof != nil {
		out.Rule = &pb.GcRule_Union_{Union: oneof}
	}
	return out
}
func GcRule_Intersection_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Intersection) *krm.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &krm.GcRule_Intersection{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Intersection_ToProto(mapCtx *direct.MapContext, in *krm.GcRule_Intersection) *pb.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Intersection{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func GcRule_Union_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Union) *krm.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &krm.GcRule_Union{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Union_ToProto(mapCtx *direct.MapContext, in *krm.GcRule_Union) *pb.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Union{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func HotTablet_FromProto(mapCtx *direct.MapContext, in *pb.HotTablet) *krm.HotTablet {
	if in == nil {
		return nil
	}
	out := &krm.HotTablet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.TableName = direct.LazyPtr(in.GetTableName())
	out.StartTime = HotTablet_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = HotTablet_EndTime_FromProto(mapCtx, in.GetEndTime())
	out.StartKey = direct.LazyPtr(in.GetStartKey())
	out.EndKey = direct.LazyPtr(in.GetEndKey())
	out.NodeCpuUsagePercent = direct.LazyPtr(in.GetNodeCpuUsagePercent())
	return out
}
func HotTablet_ToProto(mapCtx *direct.MapContext, in *krm.HotTablet) *pb.HotTablet {
	if in == nil {
		return nil
	}
	out := &pb.HotTablet{}
	out.Name = direct.ValueOf(in.Name)
	out.TableName = direct.ValueOf(in.TableName)
	out.StartTime = HotTablet_StartTime_ToProto(mapCtx, in.StartTime)
	out.EndTime = HotTablet_EndTime_ToProto(mapCtx, in.EndTime)
	out.StartKey = direct.ValueOf(in.StartKey)
	out.EndKey = direct.ValueOf(in.EndKey)
	out.NodeCpuUsagePercent = direct.ValueOf(in.NodeCpuUsagePercent)
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Labels = in.Labels
	out.CreateTime = Instance_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.SatisfiesPzs = in.SatisfiesPzs
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.Type = direct.Enum_ToProto[pb.Instance_Type](mapCtx, in.Type)
	out.Labels = in.Labels
	out.CreateTime = Instance_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.SatisfiesPzs = in.SatisfiesPzs
	return out
}
func OperationProgress_FromProto(mapCtx *direct.MapContext, in *pb.OperationProgress) *krm.OperationProgress {
	if in == nil {
		return nil
	}
	out := &krm.OperationProgress{}
	out.ProgressPercent = direct.LazyPtr(in.GetProgressPercent())
	out.StartTime = OperationProgress_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = OperationProgress_EndTime_FromProto(mapCtx, in.GetEndTime())
	return out
}
func OperationProgress_ToProto(mapCtx *direct.MapContext, in *krm.OperationProgress) *pb.OperationProgress {
	if in == nil {
		return nil
	}
	out := &pb.OperationProgress{}
	out.ProgressPercent = direct.ValueOf(in.ProgressPercent)
	out.StartTime = OperationProgress_StartTime_ToProto(mapCtx, in.StartTime)
	out.EndTime = OperationProgress_EndTime_ToProto(mapCtx, in.EndTime)
	return out
}
func RestoreInfo_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krm.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &krm.RestoreInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.BackupInfo = BackupInfo_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfo_ToProto(mapCtx *direct.MapContext, in *krm.RestoreInfo) *pb.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &pb.RestoreInfo{}
	out.SourceType = direct.Enum_ToProto[pb.RestoreSourceType](mapCtx, in.SourceType)
	if oneof := BackupInfo_ToProto(mapCtx, in.BackupInfo); oneof != nil {
		out.SourceInfo = &pb.RestoreInfo_BackupInfo{BackupInfo: oneof}
	}
	return out
}
func Snapshot_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.Snapshot {
	if in == nil {
		return nil
	}
	out := &krm.Snapshot{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SourceTable = Table_FromProto(mapCtx, in.GetSourceTable())
	out.DataSizeBytes = direct.LazyPtr(in.GetDataSizeBytes())
	out.CreateTime = Snapshot_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.DeleteTime = Snapshot_DeleteTime_FromProto(mapCtx, in.GetDeleteTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Snapshot_ToProto(mapCtx *direct.MapContext, in *krm.Snapshot) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	out.Name = direct.ValueOf(in.Name)
	out.SourceTable = Table_ToProto(mapCtx, in.SourceTable)
	out.DataSizeBytes = direct.ValueOf(in.DataSizeBytes)
	out.CreateTime = Snapshot_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.DeleteTime = Snapshot_DeleteTime_ToProto(mapCtx, in.DeleteTime)
	out.State = direct.Enum_ToProto[pb.Snapshot_State](mapCtx, in.State)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func StandardReadRemoteWrites_FromProto(mapCtx *direct.MapContext, in *pb.StandardReadRemoteWrites) *krm.StandardReadRemoteWrites {
	if in == nil {
		return nil
	}
	out := &krm.StandardReadRemoteWrites{}
	return out
}
func StandardReadRemoteWrites_ToProto(mapCtx *direct.MapContext, in *krm.StandardReadRemoteWrites) *pb.StandardReadRemoteWrites {
	if in == nil {
		return nil
	}
	out := &pb.StandardReadRemoteWrites{}
	return out
}
func Table_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.Table {
	if in == nil {
		return nil
	}
	out := &krm.Table{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	out.Granularity = direct.Enum_FromProto(mapCtx, in.GetGranularity())
	out.RestoreInfo = RestoreInfo_FromProto(mapCtx, in.GetRestoreInfo())
	out.ChangeStreamConfig = ChangeStreamConfig_FromProto(mapCtx, in.GetChangeStreamConfig())
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	out.AutomatedBackupPolicy = Table_AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())
	return out
}
func Table_ToProto(mapCtx *direct.MapContext, in *krm.Table) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	out.Granularity = direct.Enum_ToProto[pb.Table_TimestampGranularity](mapCtx, in.Granularity)
	out.RestoreInfo = RestoreInfo_ToProto(mapCtx, in.RestoreInfo)
	out.ChangeStreamConfig = ChangeStreamConfig_ToProto(mapCtx, in.ChangeStreamConfig)
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	if oneof := Table_AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy); oneof != nil {
		out.AutomatedBackupConfig = &pb.Table_AutomatedBackupPolicy_{AutomatedBackupPolicy: oneof}
	}
	return out
}
func Table_AutomatedBackupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Table_AutomatedBackupPolicy) *krm.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = AutomatedBackupPolicy_RetentionPeriod_FromProto(mapCtx, in.GetRetentionPeriod())
	out.Frequency = AutomatedBackupPolicy_Frequency_FromProto(mapCtx, in.GetFrequency())
	return out
}
func Table_AutomatedBackupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.Table_AutomatedBackupPolicy) *pb.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = AutomatedBackupPolicy_RetentionPeriod_ToProto(mapCtx, in.RetentionPeriod)
	out.Frequency = AutomatedBackupPolicy_Frequency_ToProto(mapCtx, in.Frequency)
	return out
}
func Table_ClusterState_FromProto(mapCtx *direct.MapContext, in *pb.Table_ClusterState) *krm.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &krm.Table_ClusterState{}
	out.ReplicationState = direct.Enum_FromProto(mapCtx, in.GetReplicationState())
	out.EncryptionInfo = direct.Slice_FromProto(mapCtx, in.EncryptionInfo, EncryptionInfo_FromProto)
	return out
}
func Table_ClusterState_ToProto(mapCtx *direct.MapContext, in *krm.Table_ClusterState) *pb.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &pb.Table_ClusterState{}
	out.ReplicationState = direct.Enum_ToProto[pb.Table_ClusterState_ReplicationState](mapCtx, in.ReplicationState)
	out.EncryptionInfo = direct.Slice_ToProto(mapCtx, in.EncryptionInfo, EncryptionInfo_ToProto)
	return out
}
func Type_FromProto(mapCtx *direct.MapContext, in *pb.Type) *krm.Type {
	if in == nil {
		return nil
	}
	out := &krm.Type{}
	out.BytesType = Type_Bytes_FromProto(mapCtx, in.GetBytesType())
	out.StringType = Type_String_FromProto(mapCtx, in.GetStringType())
	out.Int64Type = Type_Int64_FromProto(mapCtx, in.GetInt64Type())
	out.AggregateType = Type_Aggregate_FromProto(mapCtx, in.GetAggregateType())
	return out
}
func Type_ToProto(mapCtx *direct.MapContext, in *krm.Type) *pb.Type {
	if in == nil {
		return nil
	}
	out := &pb.Type{}
	if oneof := Type_Bytes_ToProto(mapCtx, in.BytesType); oneof != nil {
		out.Kind = &pb.Type_BytesType{BytesType: oneof}
	}
	if oneof := Type_String_ToProto(mapCtx, in.StringType); oneof != nil {
		out.Kind = &pb.Type_StringType{StringType: oneof}
	}
	if oneof := Type_Int64_ToProto(mapCtx, in.Int64Type); oneof != nil {
		out.Kind = &pb.Type_Int64Type{Int64Type: oneof}
	}
	if oneof := Type_Aggregate_ToProto(mapCtx, in.AggregateType); oneof != nil {
		out.Kind = &pb.Type_AggregateType{AggregateType: oneof}
	}
	return out
}
func Type_Aggregate_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate) *krm.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate{}
	out.InputType = Type_FromProto(mapCtx, in.GetInputType())
	out.StateType = Type_FromProto(mapCtx, in.GetStateType())
	out.Sum = Type_Aggregate_Sum_FromProto(mapCtx, in.GetSum())
	return out
}
func Type_Aggregate_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate) *pb.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate{}
	out.InputType = Type_ToProto(mapCtx, in.InputType)
	out.StateType = Type_ToProto(mapCtx, in.StateType)
	if oneof := Type_Aggregate_Sum_ToProto(mapCtx, in.Sum); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Sum_{Sum: oneof}
	}
	return out
}
func Type_Aggregate_Sum_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Sum) *krm.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_Sum{}
	return out
}
func Type_Aggregate_Sum_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_Sum) *pb.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Sum{}
	return out
}
func Type_Bytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes) *krm.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Bytes_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes) *pb.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Bytes_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding) *krm.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes_Encoding{}
	out.Raw = Type_Bytes_Encoding_Raw_FromProto(mapCtx, in.GetRaw())
	return out
}
func Type_Bytes_Encoding_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes_Encoding) *pb.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding{}
	if oneof := Type_Bytes_Encoding_Raw_ToProto(mapCtx, in.Raw); oneof != nil {
		out.Encoding = &pb.Type_Bytes_Encoding_Raw_{Raw: oneof}
	}
	return out
}
func Type_Bytes_Encoding_Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding_Raw) *krm.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Bytes_Encoding_Raw_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes_Encoding_Raw) *pb.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Int64_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64) *krm.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Int64_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64) *pb.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Int64_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding) *krm.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64_Encoding{}
	out.BigEndianBytes = Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx, in.GetBigEndianBytes())
	return out
}
func Type_Int64_Encoding_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64_Encoding) *pb.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding{}
	if oneof := Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx, in.BigEndianBytes); oneof != nil {
		out.Encoding = &pb.Type_Int64_Encoding_BigEndianBytes_{BigEndianBytes: oneof}
	}
	return out
}
func Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding_BigEndianBytes) *krm.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_FromProto(mapCtx, in.GetBytesType())
	return out
}
func Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64_Encoding_BigEndianBytes) *pb.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_ToProto(mapCtx, in.BytesType)
	return out
}
func Type_String_FromProto(mapCtx *direct.MapContext, in *pb.Type_String) *krm.Type_String {
	if in == nil {
		return nil
	}
	out := &krm.Type_String{}
	out.Encoding = Type_String_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_String_ToProto(mapCtx *direct.MapContext, in *krm.Type_String) *pb.Type_String {
	if in == nil {
		return nil
	}
	out := &pb.Type_String{}
	out.Encoding = Type_String_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_String_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding) *krm.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding{}
	out.Utf8Raw = Type_String_Encoding_Utf8Raw_FromProto(mapCtx, in.GetUtf8Raw())
	return out
}
func Type_String_Encoding_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding) *pb.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding{}
	if oneof := Type_String_Encoding_Utf8Raw_ToProto(mapCtx, in.Utf8Raw); oneof != nil {
		out.Encoding = &pb.Type_String_Encoding_Utf8Raw_{Utf8Raw: oneof}
	}
	return out
}
func Type_String_Encoding_Utf8Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Raw) *krm.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_String_Encoding_Utf8Raw_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding_Utf8Raw) *pb.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Raw{}
	return out
}
