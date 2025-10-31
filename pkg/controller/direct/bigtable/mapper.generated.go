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
	krmbigtablev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppProfile_DataBoostIsolationReadOnly_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_DataBoostIsolationReadOnly) *krm.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_DataBoostIsolationReadOnly{}
	out.ComputeBillingOwner = direct.Enum_FromProto(mapCtx, in.GetComputeBillingOwner())
	return out
}
func AppProfile_DataBoostIsolationReadOnly_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_DataBoostIsolationReadOnly) *pb.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_DataBoostIsolationReadOnly{}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx, in.ComputeBillingOwner); oneof != nil {
		out.ComputeBillingOwner = oneof
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny) *krm.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	out.RowAffinity = AppProfile_MultiClusterRoutingUseAny_RowAffinity_v1beta1_FromProto(mapCtx, in.GetRowAffinity())
	return out
}
func AppProfile_MultiClusterRoutingUseAny_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_MultiClusterRoutingUseAny) *pb.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	if oneof := AppProfile_MultiClusterRoutingUseAny_RowAffinity_v1beta1_ToProto(mapCtx, in.RowAffinity); oneof != nil {
		out.Affinity = &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity_{RowAffinity: oneof}
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *krm.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_SingleClusterRouting_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_SingleClusterRouting) *krm.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_SingleClusterRouting{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.AllowTransactionalWrites = direct.LazyPtr(in.GetAllowTransactionalWrites())
	return out
}
func AppProfile_SingleClusterRouting_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_SingleClusterRouting) *pb.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_SingleClusterRouting{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.AllowTransactionalWrites = direct.ValueOf(in.AllowTransactionalWrites)
	return out
}
func AppProfile_StandardIsolation_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_StandardIsolation) *krm.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func AppProfile_StandardIsolation_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_StandardIsolation) *pb.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_ToProto[pb.AppProfile_Priority](mapCtx, in.Priority)
	return out
}
func AuthorizedView_FamilySubsets_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView_FamilySubsets) *krmbigtablev1alpha1.AuthorizedView_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.AuthorizedView_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func AuthorizedView_FamilySubsets_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.AuthorizedView_FamilySubsets) *pb.AuthorizedView_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func AuthorizedView_SubsetView_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView_SubsetView) *krmbigtablev1alpha1.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}
func AuthorizedView_SubsetView_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.AuthorizedView_SubsetView) *pb.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}
func AutoscalingLimits_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingLimits) *krmbigtablev1alpha1.AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.AutoscalingLimits{}
	out.MinServeNodes = direct.LazyPtr(in.GetMinServeNodes())
	out.MaxServeNodes = direct.LazyPtr(in.GetMaxServeNodes())
	return out
}
func AutoscalingLimits_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.AutoscalingLimits) *pb.AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingLimits{}
	out.MinServeNodes = direct.ValueOf(in.MinServeNodes)
	out.MaxServeNodes = direct.ValueOf(in.MaxServeNodes)
	return out
}
func AutoscalingTargets_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingTargets) *krmbigtablev1alpha1.AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.AutoscalingTargets{}
	out.CPUUtilizationPercent = direct.LazyPtr(in.GetCpuUtilizationPercent())
	// MISSING: StorageUtilizationGibPerNode
	// (near miss): "StorageUtilizationGibPerNode" vs "StorageUtilizationGiBPerNode"
	return out
}
func AutoscalingTargets_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.AutoscalingTargets) *pb.AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingTargets{}
	out.CpuUtilizationPercent = direct.ValueOf(in.CPUUtilizationPercent)
	// MISSING: StorageUtilizationGibPerNode
	// (near miss): "StorageUtilizationGibPerNode" vs "StorageUtilizationGiBPerNode"
	return out
}
func BackupInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krm.BackupInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupInfo{}
	// MISSING: Backup
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SourceTable
	// MISSING: SourceBackup
	return out
}
func BackupInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackupInfo) *pb.BackupInfo {
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
func BackupInfoObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krm.BackupInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupInfoObservedState{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	return out
}
func BackupInfoObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackupInfoObservedState) *pb.BackupInfo {
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
func BigtableAppProfileObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krm.BigtableAppProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableAppProfileObservedState{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAppProfileObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableAppProfileObservedState) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAuthorizedViewObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krmbigtablev1alpha1.BigtableAuthorizedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableAuthorizedViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableAuthorizedViewObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableAuthorizedViewObservedState) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableAuthorizedViewSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krmbigtablev1alpha1.BigtableAuthorizedViewSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableAuthorizedViewSpec{}
	// MISSING: Name
	out.SubsetView = AuthorizedView_SubsetView_v1alpha1_FromProto(mapCtx, in.GetSubsetView())
	// MISSING: Etag
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func BigtableAuthorizedViewSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableAuthorizedViewSpec) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	if oneof := AuthorizedView_SubsetView_v1alpha1_ToProto(mapCtx, in.SubsetView); oneof != nil {
		out.AuthorizedView = &pb.AuthorizedView_SubsetView_{SubsetView: oneof}
	}
	// MISSING: Etag
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}
func BigtableBackupObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krmbigtablev1alpha1.BigtableBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableBackupObservedState{}
	// MISSING: Name
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.EncryptionInfo = EncryptionInfoObservedState_v1alpha1_FromProto(mapCtx, in.GetEncryptionInfo())
	return out
}
func BigtableBackupObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.EncryptionInfo = EncryptionInfoObservedState_v1alpha1_ToProto(mapCtx, in.EncryptionInfo)
	return out
}
func BigtableBackupSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krmbigtablev1alpha1.BigtableBackupSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableBackupSpec{}
	// MISSING: Name
	if in.GetSourceTable() != "" {
		out.SourceTableRef = &krm.TableRef{External: in.GetSourceTable()}
	}
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.HotToStandardTime = direct.StringTimestamp_FromProto(mapCtx, in.GetHotToStandardTime())
	return out
}
func BigtableBackupSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	if in.SourceTableRef != nil {
		out.SourceTable = in.SourceTableRef.External
	}
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	out.HotToStandardTime = direct.StringTimestamp_ToProto(mapCtx, in.HotToStandardTime)
	return out
}
func BigtableClusterObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmbigtablev1alpha1.BigtableClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableClusterObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func BigtableClusterObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	return out
}
func BigtableClusterSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmbigtablev1alpha1.BigtableClusterSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableClusterSpec{}
	// MISSING: Name
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ServeNodes = direct.LazyPtr(in.GetServeNodes())
	out.NodeScalingFactor = direct.Enum_FromProto(mapCtx, in.GetNodeScalingFactor())
	out.ClusterConfig = Cluster_ClusterConfig_v1alpha1_FromProto(mapCtx, in.GetClusterConfig())
	out.DefaultStorageType = direct.Enum_FromProto(mapCtx, in.GetDefaultStorageType())
	out.EncryptionConfig = Cluster_EncryptionConfig_v1alpha1_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func BigtableClusterSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.Location = direct.ValueOf(in.Location)
	out.ServeNodes = direct.ValueOf(in.ServeNodes)
	out.NodeScalingFactor = direct.Enum_ToProto[pb.Cluster_NodeScalingFactor](mapCtx, in.NodeScalingFactor)
	if oneof := Cluster_ClusterConfig_v1alpha1_ToProto(mapCtx, in.ClusterConfig); oneof != nil {
		out.Config = &pb.Cluster_ClusterConfig_{ClusterConfig: oneof}
	}
	out.DefaultStorageType = direct.Enum_ToProto[pb.StorageType](mapCtx, in.DefaultStorageType)
	out.EncryptionConfig = Cluster_EncryptionConfig_v1alpha1_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func BigtableInstanceSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.BigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableInstanceSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: Type
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: Tags
	return out
}
func BigtableInstanceSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: Type
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: Tags
	return out
}
func BigtableLogicalViewObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.LogicalView) *krmbigtablev1alpha1.BigtableLogicalViewObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableLogicalViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableLogicalViewObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableLogicalViewObservedState) *pb.LogicalView {
	if in == nil {
		return nil
	}
	out := &pb.LogicalView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableLogicalViewSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.LogicalView) *krmbigtablev1alpha1.BigtableLogicalViewSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableLogicalViewSpec{}
	// MISSING: Name
	out.Query = direct.LazyPtr(in.GetQuery())
	// MISSING: Etag
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func BigtableLogicalViewSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableLogicalViewSpec) *pb.LogicalView {
	if in == nil {
		return nil
	}
	out := &pb.LogicalView{}
	// MISSING: Name
	out.Query = direct.ValueOf(in.Query)
	// MISSING: Etag
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}

func BigtableMaterializedViewSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krmbigtablev1alpha1.BigtableMaterializedViewSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.BigtableMaterializedViewSpec{}
	// MISSING: Name
	out.Query = direct.LazyPtr(in.GetQuery())
	// MISSING: Etag
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func BigtableMaterializedViewSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.BigtableMaterializedViewSpec) *pb.MaterializedView {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedView{}
	// MISSING: Name
	out.Query = direct.ValueOf(in.Query)
	// MISSING: Etag
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}
func ChangeStreamConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ChangeStreamConfig) *krm.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &krm.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func ChangeStreamConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ChangeStreamConfig) *pb.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &pb.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	return out
}
func Cluster_ClusterAutoscalingConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ClusterAutoscalingConfig) *krmbigtablev1alpha1.Cluster_ClusterAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.Cluster_ClusterAutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingLimits_v1alpha1_FromProto(mapCtx, in.GetAutoscalingLimits())
	out.AutoscalingTargets = AutoscalingTargets_v1alpha1_FromProto(mapCtx, in.GetAutoscalingTargets())
	return out
}
func Cluster_ClusterAutoscalingConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.Cluster_ClusterAutoscalingConfig) *pb.Cluster_ClusterAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ClusterAutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingLimits_v1alpha1_ToProto(mapCtx, in.AutoscalingLimits)
	out.AutoscalingTargets = AutoscalingTargets_v1alpha1_ToProto(mapCtx, in.AutoscalingTargets)
	return out
}
func Cluster_ClusterConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ClusterConfig) *krmbigtablev1alpha1.Cluster_ClusterConfig {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.Cluster_ClusterConfig{}
	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_v1alpha1_FromProto(mapCtx, in.GetClusterAutoscalingConfig())
	return out
}
func Cluster_ClusterConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.Cluster_ClusterConfig) *pb.Cluster_ClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ClusterConfig{}
	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_v1alpha1_ToProto(mapCtx, in.ClusterAutoscalingConfig)
	return out
}
func Cluster_EncryptionConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_EncryptionConfig) *krmbigtablev1alpha1.Cluster_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.Cluster_EncryptionConfig{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func Cluster_EncryptionConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.Cluster_EncryptionConfig) *pb.Cluster_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_EncryptionConfig{}
	if in.KMSKeyNameRef != nil {
		out.KmsKeyName = in.KMSKeyNameRef.External
	}
	return out
}
func EncryptionInfo_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmbigtablev1alpha1.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	return out
}
func EncryptionInfo_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	return out
}
func EncryptionInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
func EncryptionInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
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
func EncryptionInfoObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmbigtablev1alpha1.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.EncryptionStatus = Status_v1alpha1_FromProto(mapCtx, in.GetEncryptionStatus())
	out.KMSKeyVersion = direct.LazyPtr(in.GetKmsKeyVersion())
	return out
}
func EncryptionInfoObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	out.EncryptionStatus = Status_v1alpha1_ToProto(mapCtx, in.EncryptionStatus)
	out.KmsKeyVersion = direct.ValueOf(in.KMSKeyVersion)
	return out
}
func GcRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GcRule) *krm.GcRule {
	if in == nil {
		return nil
	}
	out := &krm.GcRule{}
	out.MaxNumVersions = direct.LazyPtr(in.GetMaxNumVersions())
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.Intersection = GcRule_Intersection_v1beta1_FromProto(mapCtx, in.GetIntersection())
	out.Union = GcRule_Union_v1beta1_FromProto(mapCtx, in.GetUnion())
	return out
}
func GcRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GcRule) *pb.GcRule {
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
	if oneof := GcRule_Intersection_v1beta1_ToProto(mapCtx, in.Intersection); oneof != nil {
		out.Rule = &pb.GcRule_Intersection_{Intersection: oneof}
	}
	if oneof := GcRule_Union_v1beta1_ToProto(mapCtx, in.Union); oneof != nil {
		out.Rule = &pb.GcRule_Union_{Union: oneof}
	}
	return out
}
func GcRule_MaxNumVersions_ToProto(mapCtx *direct.MapContext, in *int32) *pb.GcRule_MaxNumVersions {
	if in == nil {
		return nil
	}
	return &pb.GcRule_MaxNumVersions{MaxNumVersions: *in}
}
func GcRule_Intersection_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Intersection) *krm.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &krm.GcRule_Intersection{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_v1beta1_FromProto)
	return out
}
func GcRule_Intersection_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GcRule_Intersection) *pb.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Intersection{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_v1beta1_ToProto)
	return out
}
func GcRule_Union_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Union) *krm.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &krm.GcRule_Union{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_v1beta1_FromProto)
	return out
}
func GcRule_Union_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GcRule_Union) *pb.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Union{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_v1beta1_ToProto)
	return out
}
func RestoreInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krm.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &krm.RestoreInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.BackupInfo = BackupInfo_v1beta1_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RestoreInfo) *pb.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &pb.RestoreInfo{}
	out.SourceType = direct.Enum_ToProto[pb.RestoreSourceType](mapCtx, in.SourceType)
	if oneof := BackupInfo_v1beta1_ToProto(mapCtx, in.BackupInfo); oneof != nil {
		out.SourceInfo = &pb.RestoreInfo_BackupInfo{BackupInfo: oneof}
	}
	return out
}
func RestoreInfoObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krm.RestoreInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RestoreInfoObservedState{}
	// MISSING: SourceType
	out.BackupInfo = BackupInfoObservedState_v1beta1_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfoObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RestoreInfoObservedState) *pb.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &pb.RestoreInfo{}
	// MISSING: SourceType
	if oneof := BackupInfoObservedState_v1beta1_ToProto(mapCtx, in.BackupInfo); oneof != nil {
		out.SourceInfo = &pb.RestoreInfo_BackupInfo{BackupInfo: oneof}
	}
	return out
}
func TableColumnFamily_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ColumnFamily) *krm.TableColumnFamily {
	if in == nil {
		return nil
	}
	out := &krm.TableColumnFamily{}
	// MISSING: GcRule
	// MISSING: ValueType
	return out
}
func TableColumnFamily_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.TableColumnFamily) *pb.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &pb.ColumnFamily{}
	// MISSING: GcRule
	// MISSING: ValueType
	return out
}
func Table_AutomatedBackupPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Table_AutomatedBackupPolicy) *krm.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	out.Frequency = direct.StringDuration_FromProto(mapCtx, in.GetFrequency())
	return out
}
func Table_AutomatedBackupPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Table_AutomatedBackupPolicy) *pb.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	out.Frequency = direct.StringDuration_ToProto(mapCtx, in.Frequency)
	return out
}
func Table_ClusterState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Table_ClusterState) *krm.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &krm.Table_ClusterState{}
	// MISSING: ReplicationState
	// MISSING: EncryptionInfo
	return out
}
func Table_ClusterState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Table_ClusterState) *pb.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &pb.Table_ClusterState{}
	// MISSING: ReplicationState
	// MISSING: EncryptionInfo
	return out
}
func Type_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type) *krm.Type {
	if in == nil {
		return nil
	}
	out := &krm.Type{}
	out.BytesType = Type_Bytes_v1beta1_FromProto(mapCtx, in.GetBytesType())
	out.StringType = Type_String_v1beta1_FromProto(mapCtx, in.GetStringType())
	out.Int64Type = Type_Int64_v1beta1_FromProto(mapCtx, in.GetInt64Type())
	out.Float32Type = Type_Float32_v1beta1_FromProto(mapCtx, in.GetFloat32Type())
	out.Float64Type = Type_Float64_v1beta1_FromProto(mapCtx, in.GetFloat64Type())
	out.BoolType = Type_Bool_v1beta1_FromProto(mapCtx, in.GetBoolType())
	out.TimestampType = Type_Timestamp_v1beta1_FromProto(mapCtx, in.GetTimestampType())
	out.DateType = Type_Date_v1beta1_FromProto(mapCtx, in.GetDateType())
	out.AggregateType = Type_Aggregate_v1beta1_FromProto(mapCtx, in.GetAggregateType())
	out.StructType = Type_Struct_v1beta1_FromProto(mapCtx, in.GetStructType())
	out.ArrayType = Type_Array_v1beta1_FromProto(mapCtx, in.GetArrayType())
	out.MapType = Type_Map_v1beta1_FromProto(mapCtx, in.GetMapType())
	out.ProtoType = Type_Proto_v1beta1_FromProto(mapCtx, in.GetProtoType())
	out.EnumType = Type_Enum_v1beta1_FromProto(mapCtx, in.GetEnumType())
	return out
}
func Type_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type) *pb.Type {
	if in == nil {
		return nil
	}
	out := &pb.Type{}
	if oneof := Type_Bytes_v1beta1_ToProto(mapCtx, in.BytesType); oneof != nil {
		out.Kind = &pb.Type_BytesType{BytesType: oneof}
	}
	if oneof := Type_String_v1beta1_ToProto(mapCtx, in.StringType); oneof != nil {
		out.Kind = &pb.Type_StringType{StringType: oneof}
	}
	if oneof := Type_Int64_v1beta1_ToProto(mapCtx, in.Int64Type); oneof != nil {
		out.Kind = &pb.Type_Int64Type{Int64Type: oneof}
	}
	if oneof := Type_Float32_v1beta1_ToProto(mapCtx, in.Float32Type); oneof != nil {
		out.Kind = &pb.Type_Float32Type{Float32Type: oneof}
	}
	if oneof := Type_Float64_v1beta1_ToProto(mapCtx, in.Float64Type); oneof != nil {
		out.Kind = &pb.Type_Float64Type{Float64Type: oneof}
	}
	if oneof := Type_Bool_v1beta1_ToProto(mapCtx, in.BoolType); oneof != nil {
		out.Kind = &pb.Type_BoolType{BoolType: oneof}
	}
	if oneof := Type_Timestamp_v1beta1_ToProto(mapCtx, in.TimestampType); oneof != nil {
		out.Kind = &pb.Type_TimestampType{TimestampType: oneof}
	}
	if oneof := Type_Date_v1beta1_ToProto(mapCtx, in.DateType); oneof != nil {
		out.Kind = &pb.Type_DateType{DateType: oneof}
	}
	if oneof := Type_Aggregate_v1beta1_ToProto(mapCtx, in.AggregateType); oneof != nil {
		out.Kind = &pb.Type_AggregateType{AggregateType: oneof}
	}
	if oneof := Type_Struct_v1beta1_ToProto(mapCtx, in.StructType); oneof != nil {
		out.Kind = &pb.Type_StructType{StructType: oneof}
	}
	if oneof := Type_Array_v1beta1_ToProto(mapCtx, in.ArrayType); oneof != nil {
		out.Kind = &pb.Type_ArrayType{ArrayType: oneof}
	}
	if oneof := Type_Map_v1beta1_ToProto(mapCtx, in.MapType); oneof != nil {
		out.Kind = &pb.Type_MapType{MapType: oneof}
	}
	if oneof := Type_Proto_v1beta1_ToProto(mapCtx, in.ProtoType); oneof != nil {
		out.Kind = &pb.Type_ProtoType{ProtoType: oneof}
	}
	if oneof := Type_Enum_v1beta1_ToProto(mapCtx, in.EnumType); oneof != nil {
		out.Kind = &pb.Type_EnumType{EnumType: oneof}
	}
	return out
}
func TypeObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type) *krm.TypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TypeObservedState{}
	// MISSING: BytesType
	// MISSING: StringType
	// MISSING: Int64Type
	// MISSING: Float32Type
	// MISSING: Float64Type
	// MISSING: BoolType
	// MISSING: TimestampType
	// MISSING: DateType
	out.AggregateType = Type_AggregateObservedState_v1beta1_FromProto(mapCtx, in.GetAggregateType())
	// MISSING: StructType
	// MISSING: ArrayType
	// MISSING: MapType
	// MISSING: ProtoType
	// MISSING: EnumType
	return out
}
func TypeObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.TypeObservedState) *pb.Type {
	if in == nil {
		return nil
	}
	out := &pb.Type{}
	// MISSING: BytesType
	// MISSING: StringType
	// MISSING: Int64Type
	// MISSING: Float32Type
	// MISSING: Float64Type
	// MISSING: BoolType
	// MISSING: TimestampType
	// MISSING: DateType
	if oneof := Type_AggregateObservedState_v1beta1_ToProto(mapCtx, in.AggregateType); oneof != nil {
		out.Kind = &pb.Type_AggregateType{AggregateType: oneof}
	}
	// MISSING: StructType
	// MISSING: ArrayType
	// MISSING: MapType
	// MISSING: ProtoType
	// MISSING: EnumType
	return out
}
func Type_Aggregate_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate) *krm.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate{}
	out.InputType = Type_v1beta1_FromProto(mapCtx, in.GetInputType())
	// MISSING: StateType
	out.Sum = Type_Aggregate_Sum_v1beta1_FromProto(mapCtx, in.GetSum())
	out.HllppUniqueCount = Type_Aggregate_HyperLogLogPlusPlusUniqueCount_v1beta1_FromProto(mapCtx, in.GetHllppUniqueCount())
	out.Max = Type_Aggregate_Max_v1beta1_FromProto(mapCtx, in.GetMax())
	out.Min = Type_Aggregate_Min_v1beta1_FromProto(mapCtx, in.GetMin())
	return out
}
func Type_Aggregate_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate) *pb.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate{}
	out.InputType = Type_v1beta1_ToProto(mapCtx, in.InputType)
	// MISSING: StateType
	if oneof := Type_Aggregate_Sum_v1beta1_ToProto(mapCtx, in.Sum); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Sum_{Sum: oneof}
	}
	if oneof := Type_Aggregate_HyperLogLogPlusPlusUniqueCount_v1beta1_ToProto(mapCtx, in.HllppUniqueCount); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_HllppUniqueCount{HllppUniqueCount: oneof}
	}
	if oneof := Type_Aggregate_Max_v1beta1_ToProto(mapCtx, in.Max); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Max_{Max: oneof}
	}
	if oneof := Type_Aggregate_Min_v1beta1_ToProto(mapCtx, in.Min); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Min_{Min: oneof}
	}
	return out
}
func Type_AggregateObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate) *krm.Type_AggregateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Type_AggregateObservedState{}
	// MISSING: InputType
	out.StateType = Type_v1beta1_FromProto(mapCtx, in.GetStateType())
	// MISSING: Sum
	// MISSING: HllppUniqueCount
	// MISSING: Max
	// MISSING: Min
	return out
}
func Type_AggregateObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_AggregateObservedState) *pb.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate{}
	// MISSING: InputType
	out.StateType = Type_v1beta1_ToProto(mapCtx, in.StateType)
	// MISSING: Sum
	// MISSING: HllppUniqueCount
	// MISSING: Max
	// MISSING: Min
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *krm.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_Max_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Max) *krm.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Max_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_Max) *pb.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Min_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Min) *krm.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Min_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_Min) *pb.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Sum_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Sum) *krm.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_Sum{}
	return out
}
func Type_Aggregate_Sum_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_Sum) *pb.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Sum{}
	return out
}
func Type_Array_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Array) *krm.Type_Array {
	if in == nil {
		return nil
	}
	out := &krm.Type_Array{}
	out.ElementType = Type_v1beta1_FromProto(mapCtx, in.GetElementType())
	return out
}
func Type_Array_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Array) *pb.Type_Array {
	if in == nil {
		return nil
	}
	out := &pb.Type_Array{}
	out.ElementType = Type_v1beta1_ToProto(mapCtx, in.ElementType)
	return out
}
func Type_Bool_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bool) *krm.Type_Bool {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bool{}
	return out
}
func Type_Bool_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bool) *pb.Type_Bool {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bool{}
	return out
}
func Type_Bytes_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes) *krm.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_v1beta1_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Bytes_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes) *pb.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_v1beta1_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Bytes_Encoding_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding) *krm.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes_Encoding{}
	out.Raw = Type_Bytes_Encoding_Raw_v1beta1_FromProto(mapCtx, in.GetRaw())
	return out
}
func Type_Bytes_Encoding_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes_Encoding) *pb.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding{}
	if oneof := Type_Bytes_Encoding_Raw_v1beta1_ToProto(mapCtx, in.Raw); oneof != nil {
		out.Encoding = &pb.Type_Bytes_Encoding_Raw_{Raw: oneof}
	}
	return out
}
func Type_Bytes_Encoding_Raw_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding_Raw) *krm.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Bytes_Encoding_Raw_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes_Encoding_Raw) *pb.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Date_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Date) *krm.Type_Date {
	if in == nil {
		return nil
	}
	out := &krm.Type_Date{}
	return out
}
func Type_Date_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Date) *pb.Type_Date {
	if in == nil {
		return nil
	}
	out := &pb.Type_Date{}
	return out
}
func Type_Enum_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Enum) *krm.Type_Enum {
	if in == nil {
		return nil
	}
	out := &krm.Type_Enum{}
	out.SchemaBundleID = direct.LazyPtr(in.GetSchemaBundleId())
	out.EnumName = direct.LazyPtr(in.GetEnumName())
	return out
}
func Type_Enum_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Enum) *pb.Type_Enum {
	if in == nil {
		return nil
	}
	out := &pb.Type_Enum{}
	out.SchemaBundleId = direct.ValueOf(in.SchemaBundleID)
	out.EnumName = direct.ValueOf(in.EnumName)
	return out
}
func Type_Float32_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float32) *krm.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &krm.Type_Float32{}
	return out
}
func Type_Float32_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Float32) *pb.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float32{}
	return out
}
func Type_Float64_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float64) *krm.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &krm.Type_Float64{}
	return out
}
func Type_Float64_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Float64) *pb.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float64{}
	return out
}
func Type_Int64_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64) *krm.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_v1beta1_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Int64_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64) *pb.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_v1beta1_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Int64_Encoding_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding) *krm.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64_Encoding{}
	out.BigEndianBytes = Type_Int64_Encoding_BigEndianBytes_v1beta1_FromProto(mapCtx, in.GetBigEndianBytes())
	out.OrderedCodeBytes = Type_Int64_Encoding_OrderedCodeBytes_v1beta1_FromProto(mapCtx, in.GetOrderedCodeBytes())
	return out
}
func Type_Int64_Encoding_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64_Encoding) *pb.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding{}
	if oneof := Type_Int64_Encoding_BigEndianBytes_v1beta1_ToProto(mapCtx, in.BigEndianBytes); oneof != nil {
		out.Encoding = &pb.Type_Int64_Encoding_BigEndianBytes_{BigEndianBytes: oneof}
	}
	if oneof := Type_Int64_Encoding_OrderedCodeBytes_v1beta1_ToProto(mapCtx, in.OrderedCodeBytes); oneof != nil {
		out.Encoding = &pb.Type_Int64_Encoding_OrderedCodeBytes_{OrderedCodeBytes: oneof}
	}
	return out
}
func Type_Int64_Encoding_BigEndianBytes_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding_BigEndianBytes) *krm.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_v1beta1_FromProto(mapCtx, in.GetBytesType())
	return out
}
func Type_Int64_Encoding_BigEndianBytes_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64_Encoding_BigEndianBytes) *pb.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_v1beta1_ToProto(mapCtx, in.BytesType)
	return out
}
func Type_Int64_Encoding_OrderedCodeBytes_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding_OrderedCodeBytes) *krm.Type_Int64_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Int64_Encoding_OrderedCodeBytes_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64_Encoding_OrderedCodeBytes) *pb.Type_Int64_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Map_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Map) *krm.Type_Map {
	if in == nil {
		return nil
	}
	out := &krm.Type_Map{}
	out.KeyType = Type_v1beta1_FromProto(mapCtx, in.GetKeyType())
	out.ValueType = Type_v1beta1_FromProto(mapCtx, in.GetValueType())
	return out
}
func Type_Map_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Map) *pb.Type_Map {
	if in == nil {
		return nil
	}
	out := &pb.Type_Map{}
	out.KeyType = Type_v1beta1_ToProto(mapCtx, in.KeyType)
	out.ValueType = Type_v1beta1_ToProto(mapCtx, in.ValueType)
	return out
}
func Type_Proto_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Proto) *krm.Type_Proto {
	if in == nil {
		return nil
	}
	out := &krm.Type_Proto{}
	out.SchemaBundleID = direct.LazyPtr(in.GetSchemaBundleId())
	out.MessageName = direct.LazyPtr(in.GetMessageName())
	return out
}
func Type_Proto_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Proto) *pb.Type_Proto {
	if in == nil {
		return nil
	}
	out := &pb.Type_Proto{}
	out.SchemaBundleId = direct.ValueOf(in.SchemaBundleID)
	out.MessageName = direct.ValueOf(in.MessageName)
	return out
}
func Type_String_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_String) *krm.Type_String {
	if in == nil {
		return nil
	}
	out := &krm.Type_String{}
	out.Encoding = Type_String_Encoding_v1beta1_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_String_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_String) *pb.Type_String {
	if in == nil {
		return nil
	}
	out := &pb.Type_String{}
	out.Encoding = Type_String_Encoding_v1beta1_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_String_Encoding_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding) *krm.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding{}
	out.Utf8Raw = Type_String_Encoding_Utf8Raw_v1beta1_FromProto(mapCtx, in.GetUtf8Raw())
	out.Utf8Bytes = Type_String_Encoding_Utf8Bytes_v1beta1_FromProto(mapCtx, in.GetUtf8Bytes())
	return out
}
func Type_String_Encoding_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding) *pb.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding{}
	if oneof := Type_String_Encoding_Utf8Raw_v1beta1_ToProto(mapCtx, in.Utf8Raw); oneof != nil {
		out.Encoding = &pb.Type_String_Encoding_Utf8Raw_{Utf8Raw: oneof}
	}
	if oneof := Type_String_Encoding_Utf8Bytes_v1beta1_ToProto(mapCtx, in.Utf8Bytes); oneof != nil {
		out.Encoding = &pb.Type_String_Encoding_Utf8Bytes_{Utf8Bytes: oneof}
	}
	return out
}
func Type_String_Encoding_Utf8Bytes_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Bytes) *krm.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Bytes_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding_Utf8Bytes) *pb.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Raw_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Raw) *krm.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_String_Encoding_Utf8Raw_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding_Utf8Raw) *pb.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_Struct_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct) *krm.Type_Struct {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Type_Struct_Field_v1beta1_FromProto)
	out.Encoding = Type_Struct_Encoding_v1beta1_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Struct_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct) *pb.Type_Struct {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Type_Struct_Field_v1beta1_ToProto)
	out.Encoding = Type_Struct_Encoding_v1beta1_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_StructObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct) *krm.Type_StructObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Type_StructObservedState{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Type_Struct_FieldObservedState_v1beta1_FromProto)
	// MISSING: Encoding
	return out
}
func Type_StructObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_StructObservedState) *pb.Type_Struct {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Type_Struct_FieldObservedState_v1beta1_ToProto)
	// MISSING: Encoding
	return out
}
func Type_Struct_Encoding_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding) *krm.Type_Struct_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct_Encoding{}
	out.Singleton = Type_Struct_Encoding_Singleton_v1beta1_FromProto(mapCtx, in.GetSingleton())
	out.DelimitedBytes = Type_Struct_Encoding_DelimitedBytes_v1beta1_FromProto(mapCtx, in.GetDelimitedBytes())
	out.OrderedCodeBytes = Type_Struct_Encoding_OrderedCodeBytes_v1beta1_FromProto(mapCtx, in.GetOrderedCodeBytes())
	return out
}
func Type_Struct_Encoding_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct_Encoding) *pb.Type_Struct_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding{}
	if oneof := Type_Struct_Encoding_Singleton_v1beta1_ToProto(mapCtx, in.Singleton); oneof != nil {
		out.Encoding = &pb.Type_Struct_Encoding_Singleton_{Singleton: oneof}
	}
	if oneof := Type_Struct_Encoding_DelimitedBytes_v1beta1_ToProto(mapCtx, in.DelimitedBytes); oneof != nil {
		out.Encoding = &pb.Type_Struct_Encoding_DelimitedBytes_{DelimitedBytes: oneof}
	}
	if oneof := Type_Struct_Encoding_OrderedCodeBytes_v1beta1_ToProto(mapCtx, in.OrderedCodeBytes); oneof != nil {
		out.Encoding = &pb.Type_Struct_Encoding_OrderedCodeBytes_{OrderedCodeBytes: oneof}
	}
	return out
}
func Type_Struct_Encoding_OrderedCodeBytes_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding_OrderedCodeBytes) *krm.Type_Struct_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Struct_Encoding_OrderedCodeBytes_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct_Encoding_OrderedCodeBytes) *pb.Type_Struct_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Struct_Encoding_Singleton_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding_Singleton) *krm.Type_Struct_Encoding_Singleton {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct_Encoding_Singleton{}
	return out
}
func Type_Struct_Encoding_Singleton_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct_Encoding_Singleton) *pb.Type_Struct_Encoding_Singleton {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding_Singleton{}
	return out
}
func Type_Struct_Field_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Field) *krm.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct_Field{}
	out.FieldName = direct.LazyPtr(in.GetFieldName())
	out.Type = Type_v1beta1_FromProto(mapCtx, in.GetType())
	return out
}
func Type_Struct_Field_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct_Field) *pb.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Field{}
	out.FieldName = direct.ValueOf(in.FieldName)
	out.Type = Type_v1beta1_ToProto(mapCtx, in.Type)
	return out
}
func Type_Struct_FieldObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Field) *krm.Type_Struct_FieldObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct_FieldObservedState{}
	// MISSING: FieldName
	out.Type = TypeObservedState_v1beta1_FromProto(mapCtx, in.GetType())
	return out
}
func Type_Struct_FieldObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct_FieldObservedState) *pb.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Field{}
	// MISSING: FieldName
	out.Type = TypeObservedState_v1beta1_ToProto(mapCtx, in.Type)
	return out
}
func Type_Timestamp_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Timestamp) *krm.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &krm.Type_Timestamp{}
	out.Encoding = Type_Timestamp_Encoding_v1beta1_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Timestamp_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Timestamp) *pb.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &pb.Type_Timestamp{}
	out.Encoding = Type_Timestamp_Encoding_v1beta1_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Timestamp_Encoding_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Timestamp_Encoding) *krm.Type_Timestamp_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Timestamp_Encoding{}
	out.UnixMicrosInt64 = Type_Int64_Encoding_v1beta1_FromProto(mapCtx, in.GetUnixMicrosInt64())
	return out
}
func Type_Timestamp_Encoding_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Timestamp_Encoding) *pb.Type_Timestamp_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Timestamp_Encoding{}
	if oneof := Type_Int64_Encoding_v1beta1_ToProto(mapCtx, in.UnixMicrosInt64); oneof != nil {
		out.Encoding = &pb.Type_Timestamp_Encoding_UnixMicrosInt64{UnixMicrosInt64: oneof}
	}
	return out
}
