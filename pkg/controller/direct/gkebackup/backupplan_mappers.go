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
	datepb "google.golang.org/genproto/googleapis/type/date"
	dayofweekpb "google.golang.org/genproto/googleapis/type/dayofweek"
	timeofdaypb "google.golang.org/genproto/googleapis/type/timeofday"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupPlan_Schedule_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan_Schedule) *krm.BackupPlan_Schedule {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan_Schedule{}
	out.CronSchedule = direct.LazyPtr(in.GetCronSchedule())
	out.Paused = direct.LazyPtr(in.GetPaused())
	out.RPOConfig = RPOConfig_FromProto(mapCtx, in.GetRpoConfig())
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
	out.RpoConfig = RPOConfig_ToProto(mapCtx, in.RPOConfig)
	// MISSING: NextScheduledBackupTime
	return out
}
func GKEBackupBackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.GKEBackupBackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupPlanObservedState{}
	// MISSING: Name
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.BackupSchedule = BackupPlan_ScheduleObservedState_FromProto(mapCtx, in.GetBackupSchedule())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ProtectedPodCount = direct.LazyPtr(in.GetProtectedPodCount())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	out.RPORiskLevel = direct.LazyPtr(in.GetRpoRiskLevel())
	out.RPORiskReason = direct.LazyPtr(in.GetRpoRiskReason())
	return out
}
func GKEBackupBackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.BackupSchedule = BackupPlan_ScheduleObservedState_ToProto(mapCtx, in.BackupSchedule)
	out.Etag = direct.ValueOf(in.Etag)
	out.ProtectedPodCount = direct.ValueOf(in.ProtectedPodCount)
	out.State = direct.Enum_ToProto[pb.BackupPlan_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	out.RpoRiskLevel = direct.ValueOf(in.RPORiskLevel)
	out.RpoRiskReason = direct.ValueOf(in.RPORiskReason)
	return out
}
func RPOConfig_FromProto(mapCtx *direct.MapContext, in *pb.RpoConfig) *krm.RPOConfig {
	if in == nil {
		return nil
	}
	out := &krm.RPOConfig{}
	out.TargetRPOMinutes = direct.LazyPtr(in.GetTargetRpoMinutes())
	out.ExclusionWindows = direct.Slice_FromProto(mapCtx, in.ExclusionWindows, ExclusionWindow_FromProto)
	return out
}
func RPOConfig_ToProto(mapCtx *direct.MapContext, in *krm.RPOConfig) *pb.RpoConfig {
	if in == nil {
		return nil
	}
	out := &pb.RpoConfig{}
	out.TargetRpoMinutes = direct.ValueOf(in.TargetRPOMinutes)
	out.ExclusionWindows = direct.Slice_ToProto(mapCtx, in.ExclusionWindows, ExclusionWindow_ToProto)
	return out
}
func ExclusionWindow_DayOfWeekList_ToProto(mapCtx *direct.MapContext, in *krm.ExclusionWindow_DayOfWeekList) *pb.ExclusionWindow_DayOfWeekList {
	if in == nil {
		return nil
	}
	out := &pb.ExclusionWindow_DayOfWeekList{}
	out.DaysOfWeek = direct.EnumSlice_ToProto[dayofweekpb.DayOfWeek](mapCtx, in.DaysOfWeek)
	return out
}
func TimeOfDay_FromProto(mapCtx *direct.MapContext, in *timeofdaypb.TimeOfDay) *krm.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &krm.TimeOfDay{}
	out.Hours = direct.LazyPtr(in.GetHours())
	out.Minutes = direct.LazyPtr(in.GetMinutes())
	out.Seconds = direct.LazyPtr(in.GetSeconds())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}
func TimeOfDay_ToProto(mapCtx *direct.MapContext, in *krm.TimeOfDay) *timeofdaypb.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &timeofdaypb.TimeOfDay{}
	out.Hours = direct.ValueOf(in.Hours)
	out.Minutes = direct.ValueOf(in.Minutes)
	out.Seconds = direct.ValueOf(in.Seconds)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}
func Date_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *krm.Date {
	if in == nil {
		return nil
	}
	out := &krm.Date{}
	out.Year = direct.LazyPtr(in.GetYear())
	out.Month = direct.LazyPtr(in.GetMonth())
	out.Day = direct.LazyPtr(in.GetDay())
	return out
}
func Date_ToProto(mapCtx *direct.MapContext, in *krm.Date) *datepb.Date {
	if in == nil {
		return nil
	}
	out := &datepb.Date{}
	out.Year = direct.ValueOf(in.Year)
	out.Month = direct.ValueOf(in.Month)
	out.Day = direct.ValueOf(in.Day)
	return out
}
func BackupPlan_BackupConfig_AllNamespaces_ToProto(mapCtx *direct.MapContext, allNamespaces *bool) *pb.BackupPlan_BackupConfig_AllNamespaces {
	if allNamespaces == nil {
		return nil
	}
	out := &pb.BackupPlan_BackupConfig_AllNamespaces{
		AllNamespaces: direct.ValueOf(allNamespaces),
	}
	return out
}
func ExclusionWindow_Daily_ToProto(mapCtx *direct.MapContext, daily *bool) *pb.ExclusionWindow_Daily {
	if daily == nil {
		return nil
	}
	out := &pb.ExclusionWindow_Daily{
		Daily: direct.ValueOf(daily),
	}
	return out
}
func EncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionKey) *krm.EncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionKey{}
	if in.GetGcpKmsEncryptionKey() != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{
			External: in.GetGcpKmsEncryptionKey(),
		}
	}
	return out
}
func EncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionKey) *pb.EncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionKey{}
	if in.KMSKeyRef != nil {
		out.GcpKmsEncryptionKey = in.KMSKeyRef.External
	}
	return out
}
func BackupPlan_BackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan_BackupConfig) *krm.BackupPlan_BackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan_BackupConfig{}
	// out.AllNamespaces = direct.LazyPtr(in.GetAllNamespaces())
	// out.AllNamespaces = direct.PtrTo(in.GetAllNamespaces())
	if _, ok := in.BackupScope.(*pb.BackupPlan_BackupConfig_AllNamespaces); ok {
		// special handling for oneof bool field to ensure it is round-trippable
		out.AllNamespaces = direct.PtrTo(in.GetAllNamespaces())
	}
	out.SelectedNamespaces = Namespaces_FromProto(mapCtx, in.GetSelectedNamespaces())
	out.SelectedApplications = NamespacedNames_FromProto(mapCtx, in.GetSelectedApplications())
	out.IncludeVolumeData = direct.LazyPtr(in.GetIncludeVolumeData())
	out.IncludeSecrets = direct.LazyPtr(in.GetIncludeSecrets())
	out.EncryptionKey = EncryptionKey_FromProto(mapCtx, in.GetEncryptionKey())
	out.PermissiveMode = direct.LazyPtr(in.GetPermissiveMode())
	return out
}
