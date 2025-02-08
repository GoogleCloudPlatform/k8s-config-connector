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

package firestore

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
)
func BackupSchedule_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krm.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &krm.BackupSchedule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Retention = direct.StringDuration_FromProto(mapCtx, in.GetRetention())
	out.DailyRecurrence = DailyRecurrence_FromProto(mapCtx, in.GetDailyRecurrence())
	out.WeeklyRecurrence = WeeklyRecurrence_FromProto(mapCtx, in.GetWeeklyRecurrence())
	return out
}
func BackupSchedule_ToProto(mapCtx *direct.MapContext, in *krm.BackupSchedule) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Retention = direct.StringDuration_ToProto(mapCtx, in.Retention)
	if oneof := DailyRecurrence_ToProto(mapCtx, in.DailyRecurrence); oneof != nil {
		out.Recurrence = &pb.BackupSchedule_DailyRecurrence{DailyRecurrence: oneof}
	}
	if oneof := WeeklyRecurrence_ToProto(mapCtx, in.WeeklyRecurrence); oneof != nil {
		out.Recurrence = &pb.BackupSchedule_WeeklyRecurrence{WeeklyRecurrence: oneof}
	}
	return out
}
func BackupScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krm.BackupScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupScheduleObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Retention
	// MISSING: DailyRecurrence
	// MISSING: WeeklyRecurrence
	return out
}
func BackupScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupScheduleObservedState) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Retention
	// MISSING: DailyRecurrence
	// MISSING: WeeklyRecurrence
	return out
}
func DailyRecurrence_FromProto(mapCtx *direct.MapContext, in *pb.DailyRecurrence) *krm.DailyRecurrence {
	if in == nil {
		return nil
	}
	out := &krm.DailyRecurrence{}
	return out
}
func DailyRecurrence_ToProto(mapCtx *direct.MapContext, in *krm.DailyRecurrence) *pb.DailyRecurrence {
	if in == nil {
		return nil
	}
	out := &pb.DailyRecurrence{}
	return out
}
func FirestoreBackupScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krm.FirestoreBackupScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreBackupScheduleObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Retention
	// MISSING: DailyRecurrence
	// MISSING: WeeklyRecurrence
	return out
}
func FirestoreBackupScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreBackupScheduleObservedState) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Retention
	// MISSING: DailyRecurrence
	// MISSING: WeeklyRecurrence
	return out
}
func FirestoreBackupScheduleSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krm.FirestoreBackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreBackupScheduleSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Retention
	// MISSING: DailyRecurrence
	// MISSING: WeeklyRecurrence
	return out
}
func FirestoreBackupScheduleSpec_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreBackupScheduleSpec) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Retention
	// MISSING: DailyRecurrence
	// MISSING: WeeklyRecurrence
	return out
}
func WeeklyRecurrence_FromProto(mapCtx *direct.MapContext, in *pb.WeeklyRecurrence) *krm.WeeklyRecurrence {
	if in == nil {
		return nil
	}
	out := &krm.WeeklyRecurrence{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	return out
}
func WeeklyRecurrence_ToProto(mapCtx *direct.MapContext, in *krm.WeeklyRecurrence) *pb.WeeklyRecurrence {
	if in == nil {
		return nil
	}
	out := &pb.WeeklyRecurrence{}
	out.Day = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.Day)
	return out
}
