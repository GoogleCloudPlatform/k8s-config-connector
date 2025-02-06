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

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"google.golang.org/genproto/googleapis/type/dayofweek"
	"google.golang.org/genproto/googleapis/type/timeofday"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutomatedBackupPolicy_WeeklySchedule_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_WeeklySchedule) *krm.AutomatedBackupPolicy_WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy_WeeklySchedule{}
	out.StartTimes = direct.Slice_FromProto(mapCtx, in.StartTimes, TimeOfDay_FromProto)
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	return out
}
func AutomatedBackupPolicy_WeeklySchedule_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy_WeeklySchedule) *pb.AutomatedBackupPolicy_WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_WeeklySchedule{}
	out.StartTimes = direct.Slice_ToProto(mapCtx, in.StartTimes, TimeOfDay_ToProto)
	out.DaysOfWeek = direct.EnumSlice_ToProto[dayofweek.DayOfWeek](mapCtx, in.DaysOfWeek)
	return out
}
func ContinuousBackupInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupInfo) *krm.ContinuousBackupInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContinuousBackupInfoObservedState{}
	encryptionInfo := EncryptionInfoObservedState_FromProto(mapCtx, in.GetEncryptionInfo())
	out.EncryptionInfo = []*krm.EncryptionInfoObservedState{encryptionInfo}
	out.EnabledTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEnabledTime())
	out.Schedule = direct.EnumSlice_FromProto(mapCtx, in.Schedule)
	out.EarliestRestorableTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEarliestRestorableTime())
	return out
}
func ContinuousBackupInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContinuousBackupInfoObservedState) *pb.ContinuousBackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupInfo{}
	out.EncryptionInfo = EncryptionInfoObservedState_ToProto(mapCtx, in.EncryptionInfo[0])
	out.EnabledTime = direct.StringTimestamp_ToProto(mapCtx, in.EnabledTime)
	out.Schedule = direct.EnumSlice_ToProto[dayofweek.DayOfWeek](mapCtx, in.Schedule)
	out.EarliestRestorableTime = direct.StringTimestamp_ToProto(mapCtx, in.EarliestRestorableTime)
	return out
}
func TimeOfDay_FromProto(mapCtx *direct.MapContext, in *timeofday.TimeOfDay) *krm.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &krm.TimeOfDay{}
	out.Hours = direct.PtrTo(in.GetHours())
	out.Minutes = direct.PtrTo(in.GetMinutes())
	out.Seconds = direct.PtrTo(in.GetSeconds())
	out.Nanos = direct.PtrTo(in.GetNanos())
	return out
}
func TimeOfDay_ToProto(mapCtx *direct.MapContext, in *krm.TimeOfDay) *timeofday.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &timeofday.TimeOfDay{}
	out.Hours = direct.ValueOf(in.Hours)
	out.Minutes = direct.ValueOf(in.Minutes)
	out.Seconds = direct.ValueOf(in.Seconds)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}
func MaintenanceUpdatePolicy_MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy_MaintenanceWindow) *krm.MaintenanceUpdatePolicy_MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceUpdatePolicy_MaintenanceWindow{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	return out
}
func MaintenanceUpdatePolicy_MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy_MaintenanceWindow) *pb.MaintenanceUpdatePolicy_MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy_MaintenanceWindow{}
	out.Day = direct.Enum_ToProto[dayofweek.DayOfWeek](mapCtx, in.Day)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	return out
}

func UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.UserPassword) *krm.UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.UserPassword{}
	out.User = direct.LazyPtr(in.GetUser())
	// TODO: Handle unreadable secret field.
	//out.Password = direct.LazyPtr(in.GetPassword())
	return out
}
func UserPassword_ToProto(mapCtx *direct.MapContext, in *krm.UserPassword) *pb.UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.UserPassword{}
	out.User = direct.ValueOf(in.User)
	out.Password = UserPassword_Password_ToProto(mapCtx, in.Password)
	return out
}

func UserPassword_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	// TODO: Process secret
	return ""
}
