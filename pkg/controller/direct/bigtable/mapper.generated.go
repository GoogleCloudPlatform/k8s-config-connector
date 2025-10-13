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
// krm.version: v1alpha1
// proto.service: google.bigtable.admin.v2

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krmbigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_DataBoostIsolationReadOnly) *krmbigtablev1beta1.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_DataBoostIsolationReadOnly{}
	out.ComputeBillingOwner = direct.Enum_FromProto(mapCtx, in.GetComputeBillingOwner())
	return out
}
func AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_DataBoostIsolationReadOnly) *pb.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_DataBoostIsolationReadOnly{}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx, in.ComputeBillingOwner); oneof != nil {
		out.ComputeBillingOwner = oneof
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny) *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	out.RowAffinity = AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx, in.GetRowAffinity())
	return out
}
func AppProfile_MultiClusterRoutingUseAny_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny) *pb.AppProfile_MultiClusterRoutingUseAny {
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
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_SingleClusterRouting_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_SingleClusterRouting) *krmbigtablev1beta1.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_SingleClusterRouting{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.AllowTransactionalWrites = direct.LazyPtr(in.GetAllowTransactionalWrites())
	return out
}
func AppProfile_SingleClusterRouting_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_SingleClusterRouting) *pb.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_SingleClusterRouting{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.AllowTransactionalWrites = direct.ValueOf(in.AllowTransactionalWrites)
	return out
}
func AppProfile_StandardIsolation_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_StandardIsolation) *krmbigtablev1beta1.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func AppProfile_StandardIsolation_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_StandardIsolation) *pb.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_ToProto[pb.AppProfile_Priority](mapCtx, in.Priority)
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
	out.CPUUtilizationPercent = direct.LazyPtr(in.GetCpuUtilizationPercent())
	// MISSING: StorageUtilizationGibPerNode
	// (near miss): "StorageUtilizationGibPerNode" vs "StorageUtilizationGiBPerNode"
	return out
}
func AutoscalingTargets_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingTargets) *pb.AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingTargets{}
	out.CpuUtilizationPercent = direct.ValueOf(in.CPUUtilizationPercent)
	// MISSING: StorageUtilizationGibPerNode
	// (near miss): "StorageUtilizationGibPerNode" vs "StorageUtilizationGiBPerNode"
	return out
}
func BackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krmbigtablev1beta1.BackupInfo {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BackupInfo{}
	// MISSING: Backup
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SourceTable
	// MISSING: SourceBackup
	return out
}
func BackupInfo_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BackupInfo) *pb.BackupInfo {
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
func BackupInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krmbigtablev1beta1.BackupInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BackupInfoObservedState{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	return out
}
func BackupInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BackupInfoObservedState) *pb.BackupInfo {
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
func BigtableAppProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krmbigtablev1beta1.BigtableAppProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BigtableAppProfileObservedState{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAppProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BigtableAppProfileObservedState) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAppProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krmbigtablev1beta1.BigtableAppProfileSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BigtableAppProfileSpec{}
	// MISSING: Name
	// MISSING: Etag
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MultiClusterRoutingUseAny = bool_FromProto(mapCtx, in.GetMultiClusterRoutingUseAny())
	out.SingleClusterRouting = AppProfile_SingleClusterRouting_FromProto(mapCtx, in.GetSingleClusterRouting())
	// MISSING: Priority
	out.StandardIsolation = AppProfile_StandardIsolation_FromProto(mapCtx, in.GetStandardIsolation())
	out.DataBoostIsolationReadOnly = AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx, in.GetDataBoostIsolationReadOnly())
	return out
}
func BigtableAppProfileSpec_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BigtableAppProfileSpec) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	// MISSING: Name
	// MISSING: Etag
	out.Description = direct.ValueOf(in.Description)
	if oneof := bool_ToProto(mapCtx, in.MultiClusterRoutingUseAny); oneof != nil {
		out.RoutingPolicy = &pb.AppProfile_MultiClusterRoutingUseAny_{MultiClusterRoutingUseAny: oneof}
	}
	if oneof := AppProfile_SingleClusterRouting_ToProto(mapCtx, in.SingleClusterRouting); oneof != nil {
		out.RoutingPolicy = &pb.AppProfile_SingleClusterRouting_{SingleClusterRouting: oneof}
	}
	// MISSING: Priority
	if oneof := AppProfile_StandardIsolation_ToProto(mapCtx, in.StandardIsolation); oneof != nil {
		out.Isolation = &pb.AppProfile_StandardIsolation_{StandardIsolation: oneof}
	}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx, in.DataBoostIsolationReadOnly); oneof != nil {
		out.Isolation = &pb.AppProfile_DataBoostIsolationReadOnly_{DataBoostIsolationReadOnly: oneof}
	}
	return out
}
func BigtableAuthorizedViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krm.BigtableAuthorizedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableAuthorizedViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableAuthorizedViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableAuthorizedViewObservedState) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableAuthorizedViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krm.BigtableAuthorizedViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableAuthorizedViewSpec{}
	// MISSING: Name
	out.SubsetView = AuthorizedView_SubsetView_FromProto(mapCtx, in.GetSubsetView())
	// MISSING: Etag
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func BigtableAuthorizedViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableAuthorizedViewSpec) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	if oneof := AuthorizedView_SubsetView_ToProto(mapCtx, in.SubsetView); oneof != nil {
		out.AuthorizedView = &pb.AuthorizedView_SubsetView_{SubsetView: oneof}
	}
	// MISSING: Etag
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}
func BigtableBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BigtableBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableBackupObservedState{}
	// MISSING: Name
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.EncryptionInfo = EncryptionInfoObservedState_FromProto(mapCtx, in.GetEncryptionInfo())
	return out
}
func BigtableBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableBackupObservedState) *pb.Backup {
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
	out.EncryptionInfo = EncryptionInfoObservedState_ToProto(mapCtx, in.EncryptionInfo)
	return out
}
func BigtableBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BigtableBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableBackupSpec{}
	// MISSING: Name
	if in.GetSourceTable() != "" {
		out.SourceTableRef = &krmbigtablev1beta1.TableRef{External: in.GetSourceTable()}
	}
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.HotToStandardTime = direct.StringTimestamp_FromProto(mapCtx, in.GetHotToStandardTime())
	return out
}
func BigtableBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableBackupSpec) *pb.Backup {
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
func BigtableClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.BigtableClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableClusterObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func BigtableClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	return out
}
func BigtableClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.BigtableClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableClusterSpec{}
	// MISSING: Name
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ServeNodes = direct.LazyPtr(in.GetServeNodes())
	out.NodeScalingFactor = direct.Enum_FromProto(mapCtx, in.GetNodeScalingFactor())
	out.ClusterConfig = Cluster_ClusterConfig_FromProto(mapCtx, in.GetClusterConfig())
	out.DefaultStorageType = direct.Enum_FromProto(mapCtx, in.GetDefaultStorageType())
	out.EncryptionConfig = Cluster_EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func BigtableClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.Location = direct.ValueOf(in.Location)
	out.ServeNodes = direct.ValueOf(in.ServeNodes)
	out.NodeScalingFactor = direct.Enum_ToProto[pb.Cluster_NodeScalingFactor](mapCtx, in.NodeScalingFactor)
	if oneof := Cluster_ClusterConfig_ToProto(mapCtx, in.ClusterConfig); oneof != nil {
		out.Config = &pb.Cluster_ClusterConfig_{ClusterConfig: oneof}
	}
	out.DefaultStorageType = direct.Enum_ToProto[pb.StorageType](mapCtx, in.DefaultStorageType)
	out.EncryptionConfig = Cluster_EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func BigtableInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmbigtablev1beta1.BigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BigtableInstanceSpec{}
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
func BigtableInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BigtableInstanceSpec) *pb.Instance {
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
func BigtableLogicalViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogicalView) *krm.BigtableLogicalViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableLogicalViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: DeletionProtection
	return out
}
func BigtableLogicalViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableLogicalViewObservedState) *pb.LogicalView {
	if in == nil {
		return nil
	}
	out := &pb.LogicalView{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: DeletionProtection
	return out
}
func BigtableLogicalViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogicalView) *krm.BigtableLogicalViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableLogicalViewSpec{}
	// MISSING: Name
	out.Query = direct.LazyPtr(in.GetQuery())
	// MISSING: Etag
	// MISSING: DeletionProtection
	return out
}
func BigtableLogicalViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableLogicalViewSpec) *pb.LogicalView {
	if in == nil {
		return nil
	}
	out := &pb.LogicalView{}
	// MISSING: Name
	out.Query = direct.ValueOf(in.Query)
	// MISSING: Etag
	// MISSING: DeletionProtection
	return out
}
func BigtableMaterializedViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krm.BigtableMaterializedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableMaterializedViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableMaterializedViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMaterializedViewObservedState) *pb.MaterializedView {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableMaterializedViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krm.BigtableMaterializedViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableMaterializedViewSpec{}
	// MISSING: Name
	out.Query = direct.LazyPtr(in.GetQuery())
	// MISSING: Etag
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func BigtableMaterializedViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMaterializedViewSpec) *pb.MaterializedView {
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
func BigtableTableObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmbigtablev1beta1.BigtableTableObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BigtableTableObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// TODO: map type string message for field ClusterStates
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	out.RestoreInfo = RestoreInfo_FromProto(mapCtx, in.GetRestoreInfo())
	// MISSING: ChangeStreamConfig
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}
func BigtableTableObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BigtableTableObservedState) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.Name = direct.ValueOf(in.Name)
	// TODO: map type string message for field ClusterStates
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	out.RestoreInfo = RestoreInfo_ToProto(mapCtx, in.RestoreInfo)
	// MISSING: ChangeStreamConfig
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}
func BigtableTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmbigtablev1beta1.BigtableTableSpec {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BigtableTableSpec{}
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	// MISSING: ChangeStreamConfig
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}
func BigtableTableSpec_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BigtableTableSpec) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	// MISSING: ChangeStreamConfig
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}
func ChangeStreamConfig_FromProto(mapCtx *direct.MapContext, in *pb.ChangeStreamConfig) *krmbigtablev1beta1.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func ChangeStreamConfig_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.ChangeStreamConfig) *pb.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &pb.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
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
	if in.GetKmsKeyName() != "" {
		out.KMSKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func Cluster_EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_EncryptionConfig) *pb.Cluster_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_EncryptionConfig{}
	if in.KMSKeyNameRef != nil {
		out.KmsKeyName = in.KMSKeyNameRef.External
	}
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmbigtablev1beta1.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.EncryptionInfo{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.EncryptionInfo) *pb.EncryptionInfo {
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
func EncryptionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.EncryptionStatus = Status_FromProto(mapCtx, in.GetEncryptionStatus())
	out.KMSKeyVersion = direct.LazyPtr(in.GetKmsKeyVersion())
	return out
}
func EncryptionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	out.EncryptionStatus = Status_ToProto(mapCtx, in.EncryptionStatus)
	out.KmsKeyVersion = direct.ValueOf(in.KMSKeyVersion)
	return out
}
func GcRule_FromProto(mapCtx *direct.MapContext, in *pb.GcRule) *krmbigtablev1beta1.GcRule {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.GcRule{}
	out.MaxNumVersions = direct.LazyPtr(in.GetMaxNumVersions())
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.Intersection = GcRule_Intersection_FromProto(mapCtx, in.GetIntersection())
	out.Union = GcRule_Union_FromProto(mapCtx, in.GetUnion())
	return out
}
func GcRule_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.GcRule) *pb.GcRule {
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
func GcRule_MaxNumVersions_ToProto(mapCtx *direct.MapContext, in *int32) *pb.GcRule_MaxNumVersions {
	if in == nil {
		return nil
	}
	return &pb.GcRule_MaxNumVersions{MaxNumVersions: *in}
}
func GcRule_Intersection_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Intersection) *krmbigtablev1beta1.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.GcRule_Intersection{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Intersection_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.GcRule_Intersection) *pb.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Intersection{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func GcRule_Union_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Union) *krmbigtablev1beta1.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.GcRule_Union{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Union_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.GcRule_Union) *pb.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Union{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func RestoreInfo_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krmbigtablev1beta1.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.RestoreInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.BackupInfo = BackupInfo_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfo_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.RestoreInfo) *pb.RestoreInfo {
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
func RestoreInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krmbigtablev1beta1.RestoreInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.RestoreInfoObservedState{}
	// MISSING: SourceType
	out.BackupInfo = BackupInfoObservedState_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.RestoreInfoObservedState) *pb.RestoreInfo {
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
func TableColumnFamily_FromProto(mapCtx *direct.MapContext, in *pb.ColumnFamily) *krmbigtablev1beta1.TableColumnFamily {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.TableColumnFamily{}
	// MISSING: GcRule
	// MISSING: ValueType
	return out
}
func TableColumnFamily_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.TableColumnFamily) *pb.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &pb.ColumnFamily{}
	// MISSING: GcRule
	// MISSING: ValueType
	return out
}
func Table_AutomatedBackupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Table_AutomatedBackupPolicy) *krmbigtablev1beta1.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	out.Frequency = direct.StringDuration_FromProto(mapCtx, in.GetFrequency())
	return out
}
func Table_AutomatedBackupPolicy_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Table_AutomatedBackupPolicy) *pb.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	out.Frequency = direct.StringDuration_ToProto(mapCtx, in.Frequency)
	return out
}
func Table_ClusterState_FromProto(mapCtx *direct.MapContext, in *pb.Table_ClusterState) *krmbigtablev1beta1.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Table_ClusterState{}
	// MISSING: ReplicationState
	// MISSING: EncryptionInfo
	return out
}
func Table_ClusterState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Table_ClusterState) *pb.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &pb.Table_ClusterState{}
	// MISSING: ReplicationState
	// MISSING: EncryptionInfo
	return out
}
func Type_FromProto(mapCtx *direct.MapContext, in *pb.Type) *krmbigtablev1beta1.Type {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type{}
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
	out.ProtoType = Type_Proto_FromProto(mapCtx, in.GetProtoType())
	out.EnumType = Type_Enum_FromProto(mapCtx, in.GetEnumType())
	return out
}
func Type_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type) *pb.Type {
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
	if oneof := Type_Proto_ToProto(mapCtx, in.ProtoType); oneof != nil {
		out.Kind = &pb.Type_ProtoType{ProtoType: oneof}
	}
	if oneof := Type_Enum_ToProto(mapCtx, in.EnumType); oneof != nil {
		out.Kind = &pb.Type_EnumType{EnumType: oneof}
	}
	return out
}
func TypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Type) *krmbigtablev1beta1.TypeObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.TypeObservedState{}
	// MISSING: BytesType
	// MISSING: StringType
	// MISSING: Int64Type
	// MISSING: Float32Type
	// MISSING: Float64Type
	// MISSING: BoolType
	// MISSING: TimestampType
	// MISSING: DateType
	out.AggregateType = Type_AggregateObservedState_FromProto(mapCtx, in.GetAggregateType())
	// MISSING: StructType
	// MISSING: ArrayType
	// MISSING: MapType
	// MISSING: ProtoType
	// MISSING: EnumType
	return out
}
func TypeObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.TypeObservedState) *pb.Type {
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
	if oneof := Type_AggregateObservedState_ToProto(mapCtx, in.AggregateType); oneof != nil {
		out.Kind = &pb.Type_AggregateType{AggregateType: oneof}
	}
	// MISSING: StructType
	// MISSING: ArrayType
	// MISSING: MapType
	// MISSING: ProtoType
	// MISSING: EnumType
	return out
}
func Type_Aggregate_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate) *krmbigtablev1beta1.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Aggregate{}
	out.InputType = Type_FromProto(mapCtx, in.GetInputType())
	// MISSING: StateType
	out.Sum = Type_Aggregate_Sum_FromProto(mapCtx, in.GetSum())
	out.HllppUniqueCount = Type_Aggregate_HyperLogLogPlusPlusUniqueCount_FromProto(mapCtx, in.GetHllppUniqueCount())
	out.Max = Type_Aggregate_Max_FromProto(mapCtx, in.GetMax())
	out.Min = Type_Aggregate_Min_FromProto(mapCtx, in.GetMin())
	return out
}
func Type_Aggregate_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Aggregate) *pb.Type_Aggregate {
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
func Type_AggregateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate) *krmbigtablev1beta1.Type_AggregateObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_AggregateObservedState{}
	// MISSING: InputType
	out.StateType = Type_FromProto(mapCtx, in.GetStateType())
	// MISSING: Sum
	// MISSING: HllppUniqueCount
	// MISSING: Max
	// MISSING: Min
	return out
}
func Type_AggregateObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_AggregateObservedState) *pb.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate{}
	// MISSING: InputType
	out.StateType = Type_ToProto(mapCtx, in.StateType)
	// MISSING: Sum
	// MISSING: HllppUniqueCount
	// MISSING: Max
	// MISSING: Min
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *krmbigtablev1beta1.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_Max_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Max) *krmbigtablev1beta1.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Max_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Aggregate_Max) *pb.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Min_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Min) *krmbigtablev1beta1.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Min_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Aggregate_Min) *pb.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Sum_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Sum) *krmbigtablev1beta1.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Aggregate_Sum{}
	return out
}
func Type_Aggregate_Sum_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Aggregate_Sum) *pb.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Sum{}
	return out
}
func Type_Array_FromProto(mapCtx *direct.MapContext, in *pb.Type_Array) *krmbigtablev1beta1.Type_Array {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Array{}
	out.ElementType = Type_FromProto(mapCtx, in.GetElementType())
	return out
}
func Type_Array_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Array) *pb.Type_Array {
	if in == nil {
		return nil
	}
	out := &pb.Type_Array{}
	out.ElementType = Type_ToProto(mapCtx, in.ElementType)
	return out
}
func Type_Bool_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bool) *krmbigtablev1beta1.Type_Bool {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Bool{}
	return out
}
func Type_Bool_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Bool) *pb.Type_Bool {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bool{}
	return out
}
func Type_Bytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes) *krmbigtablev1beta1.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Bytes_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Bytes) *pb.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Bytes_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding) *krmbigtablev1beta1.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Bytes_Encoding{}
	out.Raw = Type_Bytes_Encoding_Raw_FromProto(mapCtx, in.GetRaw())
	return out
}
func Type_Bytes_Encoding_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Bytes_Encoding) *pb.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding{}
	if oneof := Type_Bytes_Encoding_Raw_ToProto(mapCtx, in.Raw); oneof != nil {
		out.Encoding = &pb.Type_Bytes_Encoding_Raw_{Raw: oneof}
	}
	return out
}
func Type_Bytes_Encoding_Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding_Raw) *krmbigtablev1beta1.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Bytes_Encoding_Raw_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Bytes_Encoding_Raw) *pb.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Date_FromProto(mapCtx *direct.MapContext, in *pb.Type_Date) *krmbigtablev1beta1.Type_Date {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Date{}
	return out
}
func Type_Date_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Date) *pb.Type_Date {
	if in == nil {
		return nil
	}
	out := &pb.Type_Date{}
	return out
}
func Type_Enum_FromProto(mapCtx *direct.MapContext, in *pb.Type_Enum) *krmbigtablev1beta1.Type_Enum {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Enum{}
	out.SchemaBundleID = direct.LazyPtr(in.GetSchemaBundleId())
	out.EnumName = direct.LazyPtr(in.GetEnumName())
	return out
}
func Type_Enum_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Enum) *pb.Type_Enum {
	if in == nil {
		return nil
	}
	out := &pb.Type_Enum{}
	out.SchemaBundleId = direct.ValueOf(in.SchemaBundleID)
	out.EnumName = direct.ValueOf(in.EnumName)
	return out
}
func Type_Float32_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float32) *krmbigtablev1beta1.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Float32{}
	return out
}
func Type_Float32_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Float32) *pb.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float32{}
	return out
}
func Type_Float64_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float64) *krmbigtablev1beta1.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Float64{}
	return out
}
func Type_Float64_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Float64) *pb.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float64{}
	return out
}
func Type_Int64_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64) *krmbigtablev1beta1.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Int64_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Int64) *pb.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Int64_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding) *krmbigtablev1beta1.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Int64_Encoding{}
	out.BigEndianBytes = Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx, in.GetBigEndianBytes())
	out.OrderedCodeBytes = Type_Int64_Encoding_OrderedCodeBytes_FromProto(mapCtx, in.GetOrderedCodeBytes())
	return out
}
func Type_Int64_Encoding_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Int64_Encoding) *pb.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding{}
	if oneof := Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx, in.BigEndianBytes); oneof != nil {
		out.Encoding = &pb.Type_Int64_Encoding_BigEndianBytes_{BigEndianBytes: oneof}
	}
	if oneof := Type_Int64_Encoding_OrderedCodeBytes_ToProto(mapCtx, in.OrderedCodeBytes); oneof != nil {
		out.Encoding = &pb.Type_Int64_Encoding_OrderedCodeBytes_{OrderedCodeBytes: oneof}
	}
	return out
}
func Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding_BigEndianBytes) *krmbigtablev1beta1.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_FromProto(mapCtx, in.GetBytesType())
	return out
}
func Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Int64_Encoding_BigEndianBytes) *pb.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_ToProto(mapCtx, in.BytesType)
	return out
}
func Type_Int64_Encoding_OrderedCodeBytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding_OrderedCodeBytes) *krmbigtablev1beta1.Type_Int64_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Int64_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Int64_Encoding_OrderedCodeBytes_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Int64_Encoding_OrderedCodeBytes) *pb.Type_Int64_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Map_FromProto(mapCtx *direct.MapContext, in *pb.Type_Map) *krmbigtablev1beta1.Type_Map {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Map{}
	out.KeyType = Type_FromProto(mapCtx, in.GetKeyType())
	out.ValueType = Type_FromProto(mapCtx, in.GetValueType())
	return out
}
func Type_Map_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Map) *pb.Type_Map {
	if in == nil {
		return nil
	}
	out := &pb.Type_Map{}
	out.KeyType = Type_ToProto(mapCtx, in.KeyType)
	out.ValueType = Type_ToProto(mapCtx, in.ValueType)
	return out
}
func Type_Proto_FromProto(mapCtx *direct.MapContext, in *pb.Type_Proto) *krmbigtablev1beta1.Type_Proto {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Proto{}
	out.SchemaBundleID = direct.LazyPtr(in.GetSchemaBundleId())
	out.MessageName = direct.LazyPtr(in.GetMessageName())
	return out
}
func Type_Proto_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Proto) *pb.Type_Proto {
	if in == nil {
		return nil
	}
	out := &pb.Type_Proto{}
	out.SchemaBundleId = direct.ValueOf(in.SchemaBundleID)
	out.MessageName = direct.ValueOf(in.MessageName)
	return out
}
func Type_String_FromProto(mapCtx *direct.MapContext, in *pb.Type_String) *krmbigtablev1beta1.Type_String {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_String{}
	out.Encoding = Type_String_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_String_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_String) *pb.Type_String {
	if in == nil {
		return nil
	}
	out := &pb.Type_String{}
	out.Encoding = Type_String_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_String_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding) *krmbigtablev1beta1.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_String_Encoding{}
	out.Utf8Raw = Type_String_Encoding_Utf8Raw_FromProto(mapCtx, in.GetUtf8Raw())
	out.Utf8Bytes = Type_String_Encoding_Utf8Bytes_FromProto(mapCtx, in.GetUtf8Bytes())
	return out
}
func Type_String_Encoding_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_String_Encoding) *pb.Type_String_Encoding {
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
func Type_String_Encoding_Utf8Bytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Bytes) *krmbigtablev1beta1.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Bytes_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_String_Encoding_Utf8Bytes) *pb.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Raw) *krmbigtablev1beta1.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_String_Encoding_Utf8Raw_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_String_Encoding_Utf8Raw) *pb.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_Struct_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct) *krmbigtablev1beta1.Type_Struct {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Struct{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Type_Struct_Field_FromProto)
	out.Encoding = Type_Struct_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Struct_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Struct) *pb.Type_Struct {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Type_Struct_Field_ToProto)
	out.Encoding = Type_Struct_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_StructObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct) *krmbigtablev1beta1.Type_StructObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_StructObservedState{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Type_Struct_FieldObservedState_FromProto)
	// MISSING: Encoding
	return out
}
func Type_StructObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_StructObservedState) *pb.Type_Struct {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Type_Struct_FieldObservedState_ToProto)
	// MISSING: Encoding
	return out
}
func Type_Struct_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding) *krmbigtablev1beta1.Type_Struct_Encoding {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Struct_Encoding{}
	out.Singleton = Type_Struct_Encoding_Singleton_FromProto(mapCtx, in.GetSingleton())
	out.DelimitedBytes = Type_Struct_Encoding_DelimitedBytes_FromProto(mapCtx, in.GetDelimitedBytes())
	out.OrderedCodeBytes = Type_Struct_Encoding_OrderedCodeBytes_FromProto(mapCtx, in.GetOrderedCodeBytes())
	return out
}
func Type_Struct_Encoding_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Struct_Encoding) *pb.Type_Struct_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding{}
	if oneof := Type_Struct_Encoding_Singleton_ToProto(mapCtx, in.Singleton); oneof != nil {
		out.Encoding = &pb.Type_Struct_Encoding_Singleton_{Singleton: oneof}
	}
	if oneof := Type_Struct_Encoding_DelimitedBytes_ToProto(mapCtx, in.DelimitedBytes); oneof != nil {
		out.Encoding = &pb.Type_Struct_Encoding_DelimitedBytes_{DelimitedBytes: oneof}
	}
	if oneof := Type_Struct_Encoding_OrderedCodeBytes_ToProto(mapCtx, in.OrderedCodeBytes); oneof != nil {
		out.Encoding = &pb.Type_Struct_Encoding_OrderedCodeBytes_{OrderedCodeBytes: oneof}
	}
	return out
}
func Type_Struct_Encoding_DelimitedBytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding_DelimitedBytes) *krmbigtablev1beta1.Type_Struct_Encoding_DelimitedBytes {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Struct_Encoding_DelimitedBytes{}
	out.Delimiter = []krmbigtablev1beta1.byte{direct.LazyPtr(in.GetDelimiter())}
	return out
}
func Type_Struct_Encoding_DelimitedBytes_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Struct_Encoding_DelimitedBytes) *pb.Type_Struct_Encoding_DelimitedBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding_DelimitedBytes{}
	if len(in.Delimiter) > 0 && in.Delimiter[0] != nil {
		out.Delimiter = direct.ValueOf(in.Delimiter[0])
	}
	return out
}
func Type_Struct_Encoding_OrderedCodeBytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding_OrderedCodeBytes) *krmbigtablev1beta1.Type_Struct_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Struct_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Struct_Encoding_OrderedCodeBytes_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Struct_Encoding_OrderedCodeBytes) *pb.Type_Struct_Encoding_OrderedCodeBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding_OrderedCodeBytes{}
	return out
}
func Type_Struct_Encoding_Singleton_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding_Singleton) *krmbigtablev1beta1.Type_Struct_Encoding_Singleton {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Struct_Encoding_Singleton{}
	return out
}
func Type_Struct_Encoding_Singleton_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Struct_Encoding_Singleton) *pb.Type_Struct_Encoding_Singleton {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding_Singleton{}
	return out
}
func Type_Struct_Field_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Field) *krmbigtablev1beta1.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Struct_Field{}
	out.FieldName = direct.LazyPtr(in.GetFieldName())
	out.Type = Type_FromProto(mapCtx, in.GetType())
	return out
}
func Type_Struct_Field_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Struct_Field) *pb.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Field{}
	out.FieldName = direct.ValueOf(in.FieldName)
	out.Type = Type_ToProto(mapCtx, in.Type)
	return out
}
func Type_Struct_FieldObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Field) *krmbigtablev1beta1.Type_Struct_FieldObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Struct_FieldObservedState{}
	// MISSING: FieldName
	out.Type = TypeObservedState_FromProto(mapCtx, in.GetType())
	return out
}
func Type_Struct_FieldObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Struct_FieldObservedState) *pb.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Field{}
	// MISSING: FieldName
	out.Type = TypeObservedState_ToProto(mapCtx, in.Type)
	return out
}
func Type_Timestamp_FromProto(mapCtx *direct.MapContext, in *pb.Type_Timestamp) *krmbigtablev1beta1.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Timestamp{}
	out.Encoding = Type_Timestamp_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Timestamp_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Timestamp) *pb.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &pb.Type_Timestamp{}
	out.Encoding = Type_Timestamp_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Timestamp_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Timestamp_Encoding) *krmbigtablev1beta1.Type_Timestamp_Encoding {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.Type_Timestamp_Encoding{}
	out.UnixMicrosInt64 = Type_Int64_Encoding_FromProto(mapCtx, in.GetUnixMicrosInt64())
	return out
}
func Type_Timestamp_Encoding_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.Type_Timestamp_Encoding) *pb.Type_Timestamp_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Timestamp_Encoding{}
	if oneof := Type_Int64_Encoding_ToProto(mapCtx, in.UnixMicrosInt64); oneof != nil {
		out.Encoding = &pb.Type_Timestamp_Encoding_UnixMicrosInt64{UnixMicrosInt64: oneof}
	}
	return out
}
