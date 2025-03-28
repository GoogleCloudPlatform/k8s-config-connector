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
// krm.group: gkebackup.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.gkebackup.v1

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupPlan_BackupConfig_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlan_BackupConfig) *pb.BackupPlan_BackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan_BackupConfig{}
	if oneof := BackupPlan_BackupConfig_AllNamespaces_ToProto(mapCtx, in.AllNamespaces); oneof != nil {
		out.BackupScope = oneof
	}
	if oneof := Namespaces_ToProto(mapCtx, in.SelectedNamespaces); oneof != nil {
		out.BackupScope = &pb.BackupPlan_BackupConfig_SelectedNamespaces{SelectedNamespaces: oneof}
	}
	if oneof := NamespacedNames_ToProto(mapCtx, in.SelectedApplications); oneof != nil {
		out.BackupScope = &pb.BackupPlan_BackupConfig_SelectedApplications{SelectedApplications: oneof}
	}
	out.IncludeVolumeData = direct.ValueOf(in.IncludeVolumeData)
	out.IncludeSecrets = direct.ValueOf(in.IncludeSecrets)
	out.EncryptionKey = EncryptionKey_ToProto(mapCtx, in.EncryptionKey)
	out.PermissiveMode = direct.ValueOf(in.PermissiveMode)
	return out
}
func BackupPlan_RetentionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan_RetentionPolicy) *krm.BackupPlan_RetentionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan_RetentionPolicy{}
	out.BackupDeleteLockDays = direct.LazyPtr(in.GetBackupDeleteLockDays())
	out.BackupRetainDays = direct.LazyPtr(in.GetBackupRetainDays())
	out.Locked = direct.LazyPtr(in.GetLocked())
	return out
}
func BackupPlan_RetentionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlan_RetentionPolicy) *pb.BackupPlan_RetentionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan_RetentionPolicy{}
	out.BackupDeleteLockDays = direct.ValueOf(in.BackupDeleteLockDays)
	out.BackupRetainDays = direct.ValueOf(in.BackupRetainDays)
	out.Locked = direct.ValueOf(in.Locked)
	return out
}
func BackupPlan_ScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan_Schedule) *krm.BackupPlan_ScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan_ScheduleObservedState{}
	// MISSING: CronSchedule
	// MISSING: Paused
	// MISSING: RpoConfig
	out.NextScheduledBackupTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextScheduledBackupTime())
	return out
}
func BackupPlan_ScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlan_ScheduleObservedState) *pb.BackupPlan_Schedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan_Schedule{}
	// MISSING: CronSchedule
	// MISSING: Paused
	// MISSING: RpoConfig
	out.NextScheduledBackupTime = direct.StringTimestamp_ToProto(mapCtx, in.NextScheduledBackupTime)
	return out
}
func ExclusionWindow_FromProto(mapCtx *direct.MapContext, in *pb.ExclusionWindow) *krm.ExclusionWindow {
	if in == nil {
		return nil
	}
	out := &krm.ExclusionWindow{}
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.SingleOccurrenceDate = Date_FromProto(mapCtx, in.GetSingleOccurrenceDate())
	out.Daily = direct.LazyPtr(in.GetDaily())
	out.DaysOfWeek = ExclusionWindow_DayOfWeekList_FromProto(mapCtx, in.GetDaysOfWeek())
	return out
}
func ExclusionWindow_ToProto(mapCtx *direct.MapContext, in *krm.ExclusionWindow) *pb.ExclusionWindow {
	if in == nil {
		return nil
	}
	out := &pb.ExclusionWindow{}
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	if oneof := Date_ToProto(mapCtx, in.SingleOccurrenceDate); oneof != nil {
		out.Recurrence = &pb.ExclusionWindow_SingleOccurrenceDate{SingleOccurrenceDate: oneof}
	}
	if oneof := ExclusionWindow_Daily_ToProto(mapCtx, in.Daily); oneof != nil {
		out.Recurrence = oneof
	}
	if oneof := ExclusionWindow_DayOfWeekList_ToProto(mapCtx, in.DaysOfWeek); oneof != nil {
		out.Recurrence = &pb.ExclusionWindow_DaysOfWeek{DaysOfWeek: oneof}
	}
	return out
}
func ExclusionWindow_DayOfWeekList_FromProto(mapCtx *direct.MapContext, in *pb.ExclusionWindow_DayOfWeekList) *krm.ExclusionWindow_DayOfWeekList {
	if in == nil {
		return nil
	}
	out := &krm.ExclusionWindow_DayOfWeekList{}
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	return out
}
func GKEBackupBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.GKEBackupBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Manual = direct.LazyPtr(in.GetManual())
	out.DeleteLockExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteLockExpireTime())
	out.RetainExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRetainExpireTime())
	out.EncryptionKey = EncryptionKey_FromProto(mapCtx, in.GetEncryptionKey())
	out.AllNamespaces = direct.LazyPtr(in.GetAllNamespaces())
	out.SelectedNamespaces = Namespaces_FromProto(mapCtx, in.GetSelectedNamespaces())
	out.SelectedApplications = NamespacedNames_FromProto(mapCtx, in.GetSelectedApplications())
	out.ContainsVolumeData = direct.LazyPtr(in.GetContainsVolumeData())
	out.ContainsSecrets = direct.LazyPtr(in.GetContainsSecrets())
	out.ClusterMetadata = Backup_ClusterMetadataObservedState_FromProto(mapCtx, in.GetClusterMetadata())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.ResourceCount = direct.LazyPtr(in.GetResourceCount())
	out.VolumeCount = direct.LazyPtr(in.GetVolumeCount())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.PodCount = direct.LazyPtr(in.GetPodCount())
	out.ConfigBackupSizeBytes = direct.LazyPtr(in.GetConfigBackupSizeBytes())
	out.PermissiveMode = direct.LazyPtr(in.GetPermissiveMode())
	return out
}
func GKEBackupBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Manual = direct.ValueOf(in.Manual)
	out.DeleteLockExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteLockExpireTime)
	out.RetainExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.RetainExpireTime)
	out.EncryptionKey = EncryptionKey_ToProto(mapCtx, in.EncryptionKey)
	if oneof := GKEBackupBackupObservedState_AllNamespaces_ToProto(mapCtx, in.AllNamespaces); oneof != nil {
		out.BackupScope = oneof
	}
	if oneof := Namespaces_ToProto(mapCtx, in.SelectedNamespaces); oneof != nil {
		out.BackupScope = &pb.Backup_SelectedNamespaces{SelectedNamespaces: oneof}
	}
	if oneof := NamespacedNames_ToProto(mapCtx, in.SelectedApplications); oneof != nil {
		out.BackupScope = &pb.Backup_SelectedApplications{SelectedApplications: oneof}
	}
	out.ContainsVolumeData = direct.ValueOf(in.ContainsVolumeData)
	out.ContainsSecrets = direct.ValueOf(in.ContainsSecrets)
	out.ClusterMetadata = Backup_ClusterMetadataObservedState_ToProto(mapCtx, in.ClusterMetadata)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.ResourceCount = direct.ValueOf(in.ResourceCount)
	out.VolumeCount = direct.ValueOf(in.VolumeCount)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.Etag = direct.ValueOf(in.Etag)
	out.PodCount = direct.ValueOf(in.PodCount)
	out.ConfigBackupSizeBytes = direct.ValueOf(in.ConfigBackupSizeBytes)
	out.PermissiveMode = direct.ValueOf(in.PermissiveMode)
	return out
}
func GKEBackupBackupPlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	if in.ClusterRef != nil {
		out.Cluster = in.ClusterRef.External
	}
	out.RetentionPolicy = BackupPlan_RetentionPolicy_ToProto(mapCtx, in.RetentionPolicy)
	out.Labels = in.Labels
	out.BackupSchedule = BackupPlan_Schedule_ToProto(mapCtx, in.BackupSchedule)
	out.Deactivated = direct.ValueOf(in.Deactivated)
	out.BackupConfig = BackupPlan_BackupConfig_ToProto(mapCtx, in.BackupConfig)
	// MISSING: RpoRiskLevel
	// MISSING: RpoRiskReason
	return out
}
func GKEBackupBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.GKEBackupBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.Labels = in.Labels
	out.DeleteLockDays = direct.LazyPtr(in.GetDeleteLockDays())
	out.RetainDays = direct.LazyPtr(in.GetRetainDays())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func GKEBackupBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Uid
	out.Labels = in.Labels
	out.DeleteLockDays = direct.ValueOf(in.DeleteLockDays)
	out.RetainDays = direct.ValueOf(in.RetainDays)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func GKEBackupRestorePlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupRestorePlanSpec) *pb.RestorePlan {
	if in == nil {
		return nil
	}
	out := &pb.RestorePlan{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	if in.BackupPlanRef != nil {
		out.BackupPlan = in.BackupPlanRef.External
	}
	if in.ClusterRef != nil {
		out.Cluster = in.ClusterRef.External
	}
	out.RestoreConfig = RestoreConfig_ToProto(mapCtx, in.RestoreConfig)
	out.Labels = in.Labels
	return out
}
func NamespacedName_FromProto(mapCtx *direct.MapContext, in *pb.NamespacedName) *krm.NamespacedName {
	if in == nil {
		return nil
	}
	out := &krm.NamespacedName{}
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func NamespacedName_ToProto(mapCtx *direct.MapContext, in *krm.NamespacedName) *pb.NamespacedName {
	if in == nil {
		return nil
	}
	out := &pb.NamespacedName{}
	out.Namespace = direct.ValueOf(in.Namespace)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func NamespacedNames_FromProto(mapCtx *direct.MapContext, in *pb.NamespacedNames) *krm.NamespacedNames {
	if in == nil {
		return nil
	}
	out := &krm.NamespacedNames{}
	out.NamespacedNames = direct.Slice_FromProto(mapCtx, in.NamespacedNames, NamespacedName_FromProto)
	return out
}
func NamespacedNames_ToProto(mapCtx *direct.MapContext, in *krm.NamespacedNames) *pb.NamespacedNames {
	if in == nil {
		return nil
	}
	out := &pb.NamespacedNames{}
	out.NamespacedNames = direct.Slice_ToProto(mapCtx, in.NamespacedNames, NamespacedName_ToProto)
	return out
}
func Namespaces_FromProto(mapCtx *direct.MapContext, in *pb.Namespaces) *krm.Namespaces {
	if in == nil {
		return nil
	}
	out := &krm.Namespaces{}
	out.Namespaces = in.Namespaces
	return out
}
func Namespaces_ToProto(mapCtx *direct.MapContext, in *krm.Namespaces) *pb.Namespaces {
	if in == nil {
		return nil
	}
	out := &pb.Namespaces{}
	out.Namespaces = in.Namespaces
	return out
}
func RestoreConfig_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig) *pb.RestoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig{}
	out.VolumeDataRestorePolicy = direct.Enum_ToProto[pb.RestoreConfig_VolumeDataRestorePolicy](mapCtx, in.VolumeDataRestorePolicy)
	out.ClusterResourceConflictPolicy = direct.Enum_ToProto[pb.RestoreConfig_ClusterResourceConflictPolicy](mapCtx, in.ClusterResourceConflictPolicy)
	out.NamespacedResourceRestoreMode = direct.Enum_ToProto[pb.RestoreConfig_NamespacedResourceRestoreMode](mapCtx, in.NamespacedResourceRestoreMode)
	out.ClusterResourceRestoreScope = RestoreConfig_ClusterResourceRestoreScope_ToProto(mapCtx, in.ClusterResourceRestoreScope)
	if oneof := RestoreConfig_AllNamespaces_ToProto(mapCtx, in.AllNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = oneof
	}
	if oneof := Namespaces_ToProto(mapCtx, in.SelectedNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = &pb.RestoreConfig_SelectedNamespaces{SelectedNamespaces: oneof}
	}
	if oneof := NamespacedNames_ToProto(mapCtx, in.SelectedApplications); oneof != nil {
		out.NamespacedResourceRestoreScope = &pb.RestoreConfig_SelectedApplications{SelectedApplications: oneof}
	}
	if oneof := RestoreConfig_NoNamespaces_ToProto(mapCtx, in.NoNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = oneof
	}
	if oneof := Namespaces_ToProto(mapCtx, in.ExcludedNamespaces); oneof != nil {
		out.NamespacedResourceRestoreScope = &pb.RestoreConfig_ExcludedNamespaces{ExcludedNamespaces: oneof}
	}
	out.SubstitutionRules = direct.Slice_ToProto(mapCtx, in.SubstitutionRules, RestoreConfig_SubstitutionRule_ToProto)
	out.TransformationRules = direct.Slice_ToProto(mapCtx, in.TransformationRules, RestoreConfig_TransformationRule_ToProto)
	out.VolumeDataRestorePolicyBindings = direct.Slice_ToProto(mapCtx, in.VolumeDataRestorePolicyBindings, RestoreConfig_VolumeDataRestorePolicyBinding_ToProto)
	out.RestoreOrder = RestoreConfig_RestoreOrder_ToProto(mapCtx, in.RestoreOrder)
	return out
}
func RestoreConfig_ClusterResourceRestoreScope_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_ClusterResourceRestoreScope) *krm.RestoreConfig_ClusterResourceRestoreScope {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_ClusterResourceRestoreScope{}
	out.SelectedGroupKinds = direct.Slice_FromProto(mapCtx, in.SelectedGroupKinds, RestoreConfig_GroupKind_FromProto)
	out.ExcludedGroupKinds = direct.Slice_FromProto(mapCtx, in.ExcludedGroupKinds, RestoreConfig_GroupKind_FromProto)
	out.AllGroupKinds = direct.LazyPtr(in.GetAllGroupKinds())
	out.NoGroupKinds = direct.LazyPtr(in.GetNoGroupKinds())
	return out
}
func RestoreConfig_ClusterResourceRestoreScope_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_ClusterResourceRestoreScope) *pb.RestoreConfig_ClusterResourceRestoreScope {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_ClusterResourceRestoreScope{}
	out.SelectedGroupKinds = direct.Slice_ToProto(mapCtx, in.SelectedGroupKinds, RestoreConfig_GroupKind_ToProto)
	out.ExcludedGroupKinds = direct.Slice_ToProto(mapCtx, in.ExcludedGroupKinds, RestoreConfig_GroupKind_ToProto)
	out.AllGroupKinds = direct.ValueOf(in.AllGroupKinds)
	out.NoGroupKinds = direct.ValueOf(in.NoGroupKinds)
	return out
}
func RestoreConfig_GroupKind_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_GroupKind) *krm.RestoreConfig_GroupKind {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_GroupKind{}
	out.ResourceGroup = direct.LazyPtr(in.GetResourceGroup())
	out.ResourceKind = direct.LazyPtr(in.GetResourceKind())
	return out
}
func RestoreConfig_GroupKind_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_GroupKind) *pb.RestoreConfig_GroupKind {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_GroupKind{}
	out.ResourceGroup = direct.ValueOf(in.ResourceGroup)
	out.ResourceKind = direct.ValueOf(in.ResourceKind)
	return out
}
func RestoreConfig_RestoreOrder_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_RestoreOrder) *krm.RestoreConfig_RestoreOrder {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_RestoreOrder{}
	out.GroupKindDependencies = direct.Slice_FromProto(mapCtx, in.GroupKindDependencies, RestoreConfig_RestoreOrder_GroupKindDependency_FromProto)
	return out
}
func RestoreConfig_RestoreOrder_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_RestoreOrder) *pb.RestoreConfig_RestoreOrder {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_RestoreOrder{}
	out.GroupKindDependencies = direct.Slice_ToProto(mapCtx, in.GroupKindDependencies, RestoreConfig_RestoreOrder_GroupKindDependency_ToProto)
	return out
}
func RestoreConfig_RestoreOrder_GroupKindDependency_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_RestoreOrder_GroupKindDependency) *krm.RestoreConfig_RestoreOrder_GroupKindDependency {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_RestoreOrder_GroupKindDependency{}
	out.Satisfying = RestoreConfig_GroupKind_FromProto(mapCtx, in.GetSatisfying())
	out.Requiring = RestoreConfig_GroupKind_FromProto(mapCtx, in.GetRequiring())
	return out
}
func RestoreConfig_RestoreOrder_GroupKindDependency_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_RestoreOrder_GroupKindDependency) *pb.RestoreConfig_RestoreOrder_GroupKindDependency {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_RestoreOrder_GroupKindDependency{}
	out.Satisfying = RestoreConfig_GroupKind_ToProto(mapCtx, in.Satisfying)
	out.Requiring = RestoreConfig_GroupKind_ToProto(mapCtx, in.Requiring)
	return out
}
func RestoreConfig_TransformationRule_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_TransformationRule) *krm.RestoreConfig_TransformationRule {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_TransformationRule{}
	out.FieldActions = direct.Slice_FromProto(mapCtx, in.FieldActions, RestoreConfig_TransformationRuleAction_FromProto)
	out.ResourceFilter = RestoreConfig_ResourceFilter_FromProto(mapCtx, in.GetResourceFilter())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func RestoreConfig_TransformationRule_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_TransformationRule) *pb.RestoreConfig_TransformationRule {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_TransformationRule{}
	out.FieldActions = direct.Slice_ToProto(mapCtx, in.FieldActions, RestoreConfig_TransformationRuleAction_ToProto)
	out.ResourceFilter = RestoreConfig_ResourceFilter_ToProto(mapCtx, in.ResourceFilter)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func RestoreConfig_TransformationRuleAction_FromProto(mapCtx *direct.MapContext, in *pb.RestoreConfig_TransformationRuleAction) *krm.RestoreConfig_TransformationRuleAction {
	if in == nil {
		return nil
	}
	out := &krm.RestoreConfig_TransformationRuleAction{}
	out.Op = direct.Enum_FromProto(mapCtx, in.GetOp())
	out.FromPath = direct.LazyPtr(in.GetFromPath())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func RestoreConfig_TransformationRuleAction_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_TransformationRuleAction) *pb.RestoreConfig_TransformationRuleAction {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_TransformationRuleAction{}
	out.Op = direct.Enum_ToProto[pb.RestoreConfig_TransformationRuleAction_Op](mapCtx, in.Op)
	out.FromPath = direct.ValueOf(in.FromPath)
	out.Path = direct.ValueOf(in.Path)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func RestoreConfig_VolumeDataRestorePolicyBinding_ToProto(mapCtx *direct.MapContext, in *krm.RestoreConfig_VolumeDataRestorePolicyBinding) *pb.RestoreConfig_VolumeDataRestorePolicyBinding {
	if in == nil {
		return nil
	}
	out := &pb.RestoreConfig_VolumeDataRestorePolicyBinding{}
	out.Policy = direct.Enum_ToProto[pb.RestoreConfig_VolumeDataRestorePolicy](mapCtx, in.Policy)
	if oneof := RestoreConfig_VolumeDataRestorePolicyBinding_VolumeType_ToProto(mapCtx, in.VolumeType); oneof != nil {
		out.Scope = oneof
	}
	return out
}
