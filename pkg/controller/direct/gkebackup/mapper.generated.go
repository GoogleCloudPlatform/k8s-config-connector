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
	container "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
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
func GKEBackupBackupPlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.GKEBackupBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupPlanSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetCluster() != "" {
		out.ClusterRef = &container.ContainerClusterRef{External: in.GetCluster()}
	}
	out.RetentionPolicy = BackupPlan_RetentionPolicy_FromProto(mapCtx, in.GetRetentionPolicy())
	out.Labels = in.Labels
	out.BackupSchedule = BackupPlan_Schedule_FromProto(mapCtx, in.GetBackupSchedule())
	out.Deactivated = direct.LazyPtr(in.GetDeactivated())
	out.BackupConfig = BackupPlan_BackupConfig_FromProto(mapCtx, in.GetBackupConfig())
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
