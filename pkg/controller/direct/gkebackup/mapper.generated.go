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

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupPlan_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupPlan {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.RetentionPolicy = BackupPlan_RetentionPolicy_FromProto(mapCtx, in.GetRetentionPolicy())
	out.Labels = in.Labels
	out.BackupSchedule = BackupPlan_Schedule_FromProto(mapCtx, in.GetBackupSchedule())
	// MISSING: Etag
	out.Deactivated = direct.LazyPtr(in.GetDeactivated())
	out.BackupConfig = BackupPlan_BackupConfig_FromProto(mapCtx, in.GetBackupConfig())
	// MISSING: ProtectedPodCount
	// MISSING: State
	// MISSING: StateReason
	// MISSING: RpoRiskLevel
	// MISSING: RpoRiskReason
	return out
}
func BackupPlan_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlan) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.Cluster = direct.ValueOf(in.Cluster)
	out.RetentionPolicy = BackupPlan_RetentionPolicy_ToProto(mapCtx, in.RetentionPolicy)
	out.Labels = in.Labels
	out.BackupSchedule = BackupPlan_Schedule_ToProto(mapCtx, in.BackupSchedule)
	// MISSING: Etag
	out.Deactivated = direct.ValueOf(in.Deactivated)
	out.BackupConfig = BackupPlan_BackupConfig_ToProto(mapCtx, in.BackupConfig)
	// MISSING: ProtectedPodCount
	// MISSING: State
	// MISSING: StateReason
	// MISSING: RpoRiskLevel
	// MISSING: RpoRiskReason
	return out
}
func BackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlanObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: Cluster
	// MISSING: RetentionPolicy
	// MISSING: Labels
	out.BackupSchedule = BackupPlan_ScheduleObservedState_FromProto(mapCtx, in.GetBackupSchedule())
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Deactivated
	// MISSING: BackupConfig
	out.ProtectedPodCount = direct.LazyPtr(in.GetProtectedPodCount())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	out.RpoRiskLevel = direct.LazyPtr(in.GetRpoRiskLevel())
	out.RpoRiskReason = direct.LazyPtr(in.GetRpoRiskReason())
	return out
}
func BackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: Cluster
	// MISSING: RetentionPolicy
	// MISSING: Labels
	out.BackupSchedule = BackupPlan_ScheduleObservedState_ToProto(mapCtx, in.BackupSchedule)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Deactivated
	// MISSING: BackupConfig
	out.ProtectedPodCount = direct.ValueOf(in.ProtectedPodCount)
	out.State = direct.Enum_ToProto[pb.BackupPlan_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	out.RpoRiskLevel = direct.ValueOf(in.RpoRiskLevel)
	out.RpoRiskReason = direct.ValueOf(in.RpoRiskReason)
	return out
}
func BackupPlan_BackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan_BackupConfig) *krm.BackupPlan_BackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan_BackupConfig{}
	out.AllNamespaces = direct.LazyPtr(in.GetAllNamespaces())
	out.SelectedNamespaces = Namespaces_FromProto(mapCtx, in.GetSelectedNamespaces())
	out.SelectedApplications = NamespacedNames_FromProto(mapCtx, in.GetSelectedApplications())
	out.IncludeVolumeData = direct.LazyPtr(in.GetIncludeVolumeData())
	out.IncludeSecrets = direct.LazyPtr(in.GetIncludeSecrets())
	out.EncryptionKey = EncryptionKey_FromProto(mapCtx, in.GetEncryptionKey())
	out.PermissiveMode = direct.LazyPtr(in.GetPermissiveMode())
	return out
}
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
func BackupPlan_Schedule_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan_Schedule) *krm.BackupPlan_Schedule {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan_Schedule{}
	out.CronSchedule = direct.LazyPtr(in.GetCronSchedule())
	out.Paused = direct.LazyPtr(in.GetPaused())
	out.RpoConfig = RpoConfig_FromProto(mapCtx, in.GetRpoConfig())
	// MISSING: NextScheduledBackupTime
	return out
}
func BackupPlan_Schedule_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlan_Schedule) *pb.BackupPlan_Schedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan_Schedule{}
	out.CronSchedule = direct.ValueOf(in.CronSchedule)
	out.Paused = direct.ValueOf(in.Paused)
	out.RpoConfig = RpoConfig_ToProto(mapCtx, in.RpoConfig)
	// MISSING: NextScheduledBackupTime
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
func EncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionKey) *krm.EncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionKey{}
	out.GcpKMSEncryptionKey = direct.LazyPtr(in.GetGcpKmsEncryptionKey())
	return out
}
func EncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionKey) *pb.EncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionKey{}
	out.GcpKmsEncryptionKey = direct.ValueOf(in.GcpKMSEncryptionKey)
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
func ExclusionWindow_DayOfWeekList_ToProto(mapCtx *direct.MapContext, in *krm.ExclusionWindow_DayOfWeekList) *pb.ExclusionWindow_DayOfWeekList {
	if in == nil {
		return nil
	}
	out := &pb.ExclusionWindow_DayOfWeekList{}
	out.DaysOfWeek = direct.EnumSlice_ToProto[pb.DayOfWeek](mapCtx, in.DaysOfWeek)
	return out
}
func GKEBackupBackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.GKEBackupBackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupPlanObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Cluster
	// MISSING: RetentionPolicy
	// MISSING: Labels
	// MISSING: BackupSchedule
	// MISSING: Etag
	// MISSING: Deactivated
	// MISSING: BackupConfig
	// MISSING: ProtectedPodCount
	// MISSING: State
	// MISSING: StateReason
	// MISSING: RpoRiskLevel
	// MISSING: RpoRiskReason
	return out
}
func GKEBackupBackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Cluster
	// MISSING: RetentionPolicy
	// MISSING: Labels
	// MISSING: BackupSchedule
	// MISSING: Etag
	// MISSING: Deactivated
	// MISSING: BackupConfig
	// MISSING: ProtectedPodCount
	// MISSING: State
	// MISSING: StateReason
	// MISSING: RpoRiskLevel
	// MISSING: RpoRiskReason
	return out
}
func GKEBackupBackupPlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.GKEBackupBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupPlanSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Cluster
	// MISSING: RetentionPolicy
	// MISSING: Labels
	// MISSING: BackupSchedule
	// MISSING: Etag
	// MISSING: Deactivated
	// MISSING: BackupConfig
	// MISSING: ProtectedPodCount
	// MISSING: State
	// MISSING: StateReason
	// MISSING: RpoRiskLevel
	// MISSING: RpoRiskReason
	return out
}
func GKEBackupBackupPlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Cluster
	// MISSING: RetentionPolicy
	// MISSING: Labels
	// MISSING: BackupSchedule
	// MISSING: Etag
	// MISSING: Deactivated
	// MISSING: BackupConfig
	// MISSING: ProtectedPodCount
	// MISSING: State
	// MISSING: StateReason
	// MISSING: RpoRiskLevel
	// MISSING: RpoRiskReason
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
func RpoConfig_FromProto(mapCtx *direct.MapContext, in *pb.RpoConfig) *krm.RpoConfig {
	if in == nil {
		return nil
	}
	out := &krm.RpoConfig{}
	out.TargetRpoMinutes = direct.LazyPtr(in.GetTargetRpoMinutes())
	out.ExclusionWindows = direct.Slice_FromProto(mapCtx, in.ExclusionWindows, ExclusionWindow_FromProto)
	return out
}
func RpoConfig_ToProto(mapCtx *direct.MapContext, in *krm.RpoConfig) *pb.RpoConfig {
	if in == nil {
		return nil
	}
	out := &pb.RpoConfig{}
	out.TargetRpoMinutes = direct.ValueOf(in.TargetRpoMinutes)
	out.ExclusionWindows = direct.Slice_ToProto(mapCtx, in.ExclusionWindows, ExclusionWindow_ToProto)
	return out
}
