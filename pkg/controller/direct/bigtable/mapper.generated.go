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
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.bigtable.admin.v2

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_DataBoostIsolationReadOnly) *krmv1beta1.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_DataBoostIsolationReadOnly{}
	out.ComputeBillingOwner = direct.Enum_FromProto(mapCtx, in.GetComputeBillingOwner())
	return out
}
func AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_DataBoostIsolationReadOnly) *pb.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_DataBoostIsolationReadOnly{}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx, in.ComputeBillingOwner); oneof != nil {
		out.ComputeBillingOwner = oneof
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny) *krmv1beta1.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	out.RowAffinity = AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx, in.GetRowAffinity())
	return out
}
func AppProfile_MultiClusterRoutingUseAny_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_MultiClusterRoutingUseAny) *pb.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	if oneof := AppProfile_MultiClusterRoutingUseAny_RowAffinity_ToProto(mapCtx, in.RowAffinity); oneof != nil {
		out.Affinity = &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity_{RowAffinity: oneof}
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *krmv1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_SingleClusterRouting_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_SingleClusterRouting) *krmv1beta1.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_SingleClusterRouting{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.AllowTransactionalWrites = direct.LazyPtr(in.GetAllowTransactionalWrites())
	return out
}
func AppProfile_SingleClusterRouting_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_SingleClusterRouting) *pb.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_SingleClusterRouting{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.AllowTransactionalWrites = direct.ValueOf(in.AllowTransactionalWrites)
	return out
}
func AppProfile_StandardIsolation_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_StandardIsolation) *krmv1beta1.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func AppProfile_StandardIsolation_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_StandardIsolation) *pb.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_ToProto[pb.AppProfile_Priority](mapCtx, in.Priority)
	return out
}
func AuthorizedView_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krmv1alpha1.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.AuthorizedView{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SubsetView = AuthorizedView_SubsetView_FromProto(mapCtx, in.GetSubsetView())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func AuthorizedView_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.AuthorizedView) *pb.AuthorizedView {
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
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krmv1alpha1.Backup {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Backup{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	// MISSING: SourceBackup
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SizeBytes
	// MISSING: State
	// MISSING: EncryptionInfo
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.HotToStandardTime = direct.StringTimestamp_FromProto(mapCtx, in.GetHotToStandardTime())
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	out.SourceTable = direct.ValueOf(in.SourceTable)
	// MISSING: SourceBackup
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SizeBytes
	// MISSING: State
	// MISSING: EncryptionInfo
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	out.HotToStandardTime = direct.StringTimestamp_ToProto(mapCtx, in.HotToStandardTime)
	return out
}
func BackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krmv1beta1.BackupInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BackupInfo{}
	// MISSING: Backup
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SourceTable
	// MISSING: SourceBackup
	return out
}
func BackupInfo_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BackupInfo) *pb.BackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupInfo{}
	// MISSING: Backup
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SourceTable
	// MISSING: SourceBackup
	return out
}
func BackupInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krmv1beta1.BackupInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BackupInfoObservedState{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	return out
}
func BackupInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BackupInfoObservedState) *pb.BackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupInfo{}
	out.Backup = direct.ValueOf(in.Backup)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.SourceTable = direct.ValueOf(in.SourceTable)
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	return out
}
func BigtableAppProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krmv1beta1.BigtableAppProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigtableAppProfileObservedState{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	// MISSING: DataBoostIsolationReadOnly
	return out
}
func BigtableAppProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigtableAppProfileObservedState) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	// MISSING: DataBoostIsolationReadOnly
	return out
}
func BigtableAuthorizedViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krmv1alpha1.BigtableAuthorizedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigtableAuthorizedViewObservedState{}
	// MISSING: Name
	return out
}
func BigtableAuthorizedViewObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableAuthorizedViewObservedState) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	return out
}
func ChangeStreamConfig_FromProto(mapCtx *direct.MapContext, in *pb.ChangeStreamConfig) *krmv1beta1.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func ChangeStreamConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ChangeStreamConfig) *pb.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &pb.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	return out
}
func ColumnFamily_FromProto(mapCtx *direct.MapContext, in *pb.ColumnFamily) *krmv1beta1.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ColumnFamily{}
	out.GcRule = GcRule_FromProto(mapCtx, in.GetGcRule())
	out.ValueType = Type_FromProto(mapCtx, in.GetValueType())
	return out
}
func ColumnFamily_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ColumnFamily) *pb.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &pb.ColumnFamily{}
	out.GcRule = GcRule_ToProto(mapCtx, in.GcRule)
	out.ValueType = Type_ToProto(mapCtx, in.ValueType)
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmv1beta1.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.EncryptionInfo{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
func GcRule_FromProto(mapCtx *direct.MapContext, in *pb.GcRule) *krmv1beta1.GcRule {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.GcRule{}
	out.MaxNumVersions = direct.LazyPtr(in.GetMaxNumVersions())
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.Intersection = GcRule_Intersection_FromProto(mapCtx, in.GetIntersection())
	out.Union = GcRule_Union_FromProto(mapCtx, in.GetUnion())
	return out
}
func GcRule_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.GcRule) *pb.GcRule {
	if in == nil {
		return nil
	}
	out := &pb.GcRule{}
	if oneof := GcRule_MaxNumVersions_ToProto(mapCtx, in.MaxNumVersions); oneof != nil {
		out.Rule = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.MaxAge); oneof != nil {
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
func GcRule_Intersection_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Intersection) *krmv1beta1.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.GcRule_Intersection{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Intersection_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.GcRule_Intersection) *pb.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Intersection{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func GcRule_Union_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Union) *krmv1beta1.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.GcRule_Union{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Union_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.GcRule_Union) *pb.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Union{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func RestoreInfo_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krmv1beta1.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RestoreInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.BackupInfo = BackupInfo_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfo_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RestoreInfo) *pb.RestoreInfo {
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
func RestoreInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krmv1beta1.RestoreInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RestoreInfoObservedState{}
	// MISSING: SourceType
	out.BackupInfo = BackupInfoObservedState_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RestoreInfoObservedState) *pb.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &pb.RestoreInfo{}
	// MISSING: SourceType
	if oneof := BackupInfoObservedState_ToProto(mapCtx, in.BackupInfo); oneof != nil {
		out.SourceInfo = &pb.RestoreInfo_BackupInfo{BackupInfo: oneof}
	}
	return out
}
func Table_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmv1beta1.Table {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Table{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	out.Granularity = direct.Enum_FromProto(mapCtx, in.GetGranularity())
	// MISSING: RestoreInfo
	out.ChangeStreamConfig = ChangeStreamConfig_FromProto(mapCtx, in.GetChangeStreamConfig())
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	out.AutomatedBackupPolicy = Table_AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())
	return out
}
func Table_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Table) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	out.Granularity = direct.Enum_ToProto[pb.Table_TimestampGranularity](mapCtx, in.Granularity)
	// MISSING: RestoreInfo
	out.ChangeStreamConfig = ChangeStreamConfig_ToProto(mapCtx, in.ChangeStreamConfig)
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	if oneof := Table_AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy); oneof != nil {
		out.AutomatedBackupConfig = &pb.Table_AutomatedBackupPolicy_{AutomatedBackupPolicy: oneof}
	}
	return out
}
func TableObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmv1beta1.TableObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TableObservedState{}
	// MISSING: Name
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	out.RestoreInfo = RestoreInfo_FromProto(mapCtx, in.GetRestoreInfo())
	// MISSING: ChangeStreamConfig
	// MISSING: DeletionProtection
	// MISSING: AutomatedBackupPolicy
	return out
}
func TableObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TableObservedState) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	// MISSING: Name
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	out.RestoreInfo = RestoreInfo_ToProto(mapCtx, in.RestoreInfo)
	// MISSING: ChangeStreamConfig
	// MISSING: DeletionProtection
	// MISSING: AutomatedBackupPolicy
	return out
}
func Table_AutomatedBackupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Table_AutomatedBackupPolicy) *krmv1beta1.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	out.Frequency = direct.StringDuration_FromProto(mapCtx, in.GetFrequency())
	return out
}
func Table_AutomatedBackupPolicy_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Table_AutomatedBackupPolicy) *pb.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	out.Frequency = direct.StringDuration_ToProto(mapCtx, in.Frequency)
	return out
}
func Table_ClusterState_FromProto(mapCtx *direct.MapContext, in *pb.Table_ClusterState) *krmv1beta1.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Table_ClusterState{}
	// MISSING: ReplicationState
	// MISSING: EncryptionInfo
	return out
}
func Table_ClusterState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Table_ClusterState) *pb.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &pb.Table_ClusterState{}
	// MISSING: ReplicationState
	// MISSING: EncryptionInfo
	return out
}
func Type_FromProto(mapCtx *direct.MapContext, in *pb.Type) *krmv1beta1.Type {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type{}
	out.BytesType = Type_Bytes_FromProto(mapCtx, in.GetBytesType())
	out.StringType = Type_String_FromProto(mapCtx, in.GetStringType())
	out.Int64Type = Type_Int64_FromProto(mapCtx, in.GetInt64Type())
	out.Float32Type = Type_Float32_FromProto(mapCtx, in.GetFloat32Type())
	out.Float64Type = Type_Float64_FromProto(mapCtx, in.GetFloat64Type())
	out.BoolType = Type_Bool_FromProto(mapCtx, in.GetBoolType())
	out.TimestampType = Type_Timestamp_FromProto(mapCtx, in.GetTimestampType())
	out.DateType = Type_Date_FromProto(mapCtx, in.GetDateType())
	out.AggregateType = Type_Aggregate_FromProto(mapCtx, in.GetAggregateType())
	out.StructType = Type_Struct_FromProto(mapCtx, in.GetStructType())
	out.ArrayType = Type_Array_FromProto(mapCtx, in.GetArrayType())
	out.MapType = Type_Map_FromProto(mapCtx, in.GetMapType())
	return out
}
func Type_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type) *pb.Type {
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
	if oneof := Type_Float32_ToProto(mapCtx, in.Float32Type); oneof != nil {
		out.Kind = &pb.Type_Float32Type{Float32Type: oneof}
	}
	if oneof := Type_Float64_ToProto(mapCtx, in.Float64Type); oneof != nil {
		out.Kind = &pb.Type_Float64Type{Float64Type: oneof}
	}
	if oneof := Type_Bool_ToProto(mapCtx, in.BoolType); oneof != nil {
		out.Kind = &pb.Type_BoolType{BoolType: oneof}
	}
	if oneof := Type_Timestamp_ToProto(mapCtx, in.TimestampType); oneof != nil {
		out.Kind = &pb.Type_TimestampType{TimestampType: oneof}
	}
	if oneof := Type_Date_ToProto(mapCtx, in.DateType); oneof != nil {
		out.Kind = &pb.Type_DateType{DateType: oneof}
	}
	if oneof := Type_Aggregate_ToProto(mapCtx, in.AggregateType); oneof != nil {
		out.Kind = &pb.Type_AggregateType{AggregateType: oneof}
	}
	if oneof := Type_Struct_ToProto(mapCtx, in.StructType); oneof != nil {
		out.Kind = &pb.Type_StructType{StructType: oneof}
	}
	if oneof := Type_Array_ToProto(mapCtx, in.ArrayType); oneof != nil {
		out.Kind = &pb.Type_ArrayType{ArrayType: oneof}
	}
	if oneof := Type_Map_ToProto(mapCtx, in.MapType); oneof != nil {
		out.Kind = &pb.Type_MapType{MapType: oneof}
	}
	return out
}
func Type_Aggregate_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate) *krmv1beta1.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Aggregate{}
	out.InputType = Type_FromProto(mapCtx, in.GetInputType())
	// MISSING: StateType
	out.Sum = Type_Aggregate_Sum_FromProto(mapCtx, in.GetSum())
	out.HllppUniqueCount = Type_Aggregate_HyperLogLogPlusPlusUniqueCount_FromProto(mapCtx, in.GetHllppUniqueCount())
	out.Max = Type_Aggregate_Max_FromProto(mapCtx, in.GetMax())
	out.Min = Type_Aggregate_Min_FromProto(mapCtx, in.GetMin())
	return out
}
func Type_Aggregate_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Aggregate) *pb.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate{}
	out.InputType = Type_ToProto(mapCtx, in.InputType)
	// MISSING: StateType
	if oneof := Type_Aggregate_Sum_ToProto(mapCtx, in.Sum); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Sum_{Sum: oneof}
	}
	if oneof := Type_Aggregate_HyperLogLogPlusPlusUniqueCount_ToProto(mapCtx, in.HllppUniqueCount); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_HllppUniqueCount{HllppUniqueCount: oneof}
	}
	if oneof := Type_Aggregate_Max_ToProto(mapCtx, in.Max); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Max_{Max: oneof}
	}
	if oneof := Type_Aggregate_Min_ToProto(mapCtx, in.Min); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Min_{Min: oneof}
	}
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *krmv1beta1.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_Max_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Max) *krmv1beta1.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Max_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Aggregate_Max) *pb.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Min_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Min) *krmv1beta1.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Min_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Aggregate_Min) *pb.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Sum_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Sum) *krmv1beta1.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Aggregate_Sum{}
	return out
}
func Type_Aggregate_Sum_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Aggregate_Sum) *pb.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Sum{}
	return out
}
func Type_Array_FromProto(mapCtx *direct.MapContext, in *pb.Type_Array) *krmv1beta1.Type_Array {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Array{}
	out.ElementType = Type_FromProto(mapCtx, in.GetElementType())
	return out
}
func Type_Array_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Array) *pb.Type_Array {
	if in == nil {
		return nil
	}
	out := &pb.Type_Array{}
	out.ElementType = Type_ToProto(mapCtx, in.ElementType)
	return out
}
func Type_Bool_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bool) *krmv1beta1.Type_Bool {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Bool{}
	return out
}
func Type_Bool_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Bool) *pb.Type_Bool {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bool{}
	return out
}
func Type_Bytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes) *krmv1beta1.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Bytes_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Bytes) *pb.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Bytes_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding) *krmv1beta1.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Bytes_Encoding{}
	out.Raw = Type_Bytes_Encoding_Raw_FromProto(mapCtx, in.GetRaw())
	return out
}
func Type_Bytes_Encoding_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Bytes_Encoding) *pb.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding{}
	if oneof := Type_Bytes_Encoding_Raw_ToProto(mapCtx, in.Raw); oneof != nil {
		out.Encoding = &pb.Type_Bytes_Encoding_Raw_{Raw: oneof}
	}
	return out
}
func Type_Bytes_Encoding_Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding_Raw) *krmv1beta1.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Bytes_Encoding_Raw_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Bytes_Encoding_Raw) *pb.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Date_FromProto(mapCtx *direct.MapContext, in *pb.Type_Date) *krmv1beta1.Type_Date {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Date{}
	return out
}
func Type_Date_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Date) *pb.Type_Date {
	if in == nil {
		return nil
	}
	out := &pb.Type_Date{}
	return out
}
func Type_Float32_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float32) *krmv1beta1.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Float32{}
	return out
}
func Type_Float32_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Float32) *pb.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float32{}
	return out
}
func Type_Float64_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float64) *krmv1beta1.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Float64{}
	return out
}
func Type_Float64_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Float64) *pb.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float64{}
	return out
}
func Type_Int64_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64) *krmv1beta1.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Int64_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Int64) *pb.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Int64_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding) *krmv1beta1.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Int64_Encoding{}
	out.BigEndianBytes = Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx, in.GetBigEndianBytes())
	return out
}
func Type_Int64_Encoding_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Int64_Encoding) *pb.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding{}
	if oneof := Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx, in.BigEndianBytes); oneof != nil {
		out.Encoding = &pb.Type_Int64_Encoding_BigEndianBytes_{BigEndianBytes: oneof}
	}
	return out
}
func Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding_BigEndianBytes) *krmv1beta1.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_FromProto(mapCtx, in.GetBytesType())
	return out
}
func Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Int64_Encoding_BigEndianBytes) *pb.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_ToProto(mapCtx, in.BytesType)
	return out
}
func Type_Map_FromProto(mapCtx *direct.MapContext, in *pb.Type_Map) *krmv1beta1.Type_Map {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Map{}
	out.KeyType = Type_FromProto(mapCtx, in.GetKeyType())
	out.ValueType = Type_FromProto(mapCtx, in.GetValueType())
	return out
}
func Type_Map_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Map) *pb.Type_Map {
	if in == nil {
		return nil
	}
	out := &pb.Type_Map{}
	out.KeyType = Type_ToProto(mapCtx, in.KeyType)
	out.ValueType = Type_ToProto(mapCtx, in.ValueType)
	return out
}
func Type_String_FromProto(mapCtx *direct.MapContext, in *pb.Type_String) *krmv1beta1.Type_String {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_String{}
	out.Encoding = Type_String_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_String_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_String) *pb.Type_String {
	if in == nil {
		return nil
	}
	out := &pb.Type_String{}
	out.Encoding = Type_String_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_String_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding) *krmv1beta1.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_String_Encoding{}
	out.Utf8Raw = Type_String_Encoding_Utf8Raw_FromProto(mapCtx, in.GetUtf8Raw())
	out.Utf8Bytes = Type_String_Encoding_Utf8Bytes_FromProto(mapCtx, in.GetUtf8Bytes())
	return out
}
func Type_String_Encoding_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_String_Encoding) *pb.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding{}
	if oneof := Type_String_Encoding_Utf8Raw_ToProto(mapCtx, in.Utf8Raw); oneof != nil {
		out.Encoding = &pb.Type_String_Encoding_Utf8Raw_{Utf8Raw: oneof}
	}
	if oneof := Type_String_Encoding_Utf8Bytes_ToProto(mapCtx, in.Utf8Bytes); oneof != nil {
		out.Encoding = &pb.Type_String_Encoding_Utf8Bytes_{Utf8Bytes: oneof}
	}
	return out
}
func Type_String_Encoding_Utf8Bytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Bytes) *krmv1beta1.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Bytes_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_String_Encoding_Utf8Bytes) *pb.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Raw) *krmv1beta1.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_String_Encoding_Utf8Raw_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_String_Encoding_Utf8Raw) *pb.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_Struct_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct) *krmv1beta1.Type_Struct {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Struct{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Type_Struct_Field_FromProto)
	return out
}
func Type_Struct_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Struct) *pb.Type_Struct {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Type_Struct_Field_ToProto)
	return out
}
func Type_Struct_Field_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Field) *krmv1beta1.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Struct_Field{}
	out.FieldName = direct.LazyPtr(in.GetFieldName())
	out.Type = Type_FromProto(mapCtx, in.GetType())
	return out
}
func Type_Struct_Field_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Struct_Field) *pb.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Field{}
	out.FieldName = direct.ValueOf(in.FieldName)
	out.Type = Type_ToProto(mapCtx, in.Type)
	return out
}
func Type_Timestamp_FromProto(mapCtx *direct.MapContext, in *pb.Type_Timestamp) *krmv1beta1.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Type_Timestamp{}
	return out
}
func Type_Timestamp_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Type_Timestamp) *pb.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &pb.Type_Timestamp{}
	return out
}
