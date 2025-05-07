// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.backupdr.v1.ManagementServer
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	dayofweekpb "google.golang.org/genproto/googleapis/type/dayofweek"
	monthpb "google.golang.org/genproto/googleapis/type/month"
)

func StandardSchedule_ToProto(mapCtx *direct.MapContext, in *krm.StandardSchedule) *pb.StandardSchedule {
	if in == nil {
		return nil
	}
	out := &pb.StandardSchedule{}
	out.RecurrenceType = direct.Enum_ToProto[pb.StandardSchedule_RecurrenceType](mapCtx, in.RecurrenceType)
	out.HourlyFrequency = direct.ValueOf(in.HourlyFrequency)
	out.DaysOfWeek = direct.EnumSlice_ToProto[dayofweekpb.DayOfWeek](mapCtx, in.DaysOfWeek)
	out.DaysOfMonth = in.DaysOfMonth
	out.WeekDayOfMonth = WeekDayOfMonth_ToProto(mapCtx, in.WeekDayOfMonth)
	out.Months = direct.EnumSlice_ToProto[monthpb.Month](mapCtx, in.Months)
	out.BackupWindow = BackupWindow_ToProto(mapCtx, in.BackupWindow)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	return out
}
func WeekDayOfMonth_ToProto(mapCtx *direct.MapContext, in *krm.WeekDayOfMonth) *pb.WeekDayOfMonth {
	if in == nil {
		return nil
	}
	out := &pb.WeekDayOfMonth{}
	out.WeekOfMonth = direct.Enum_ToProto[pb.WeekDayOfMonth_WeekOfMonth](mapCtx, in.WeekOfMonth)
	out.DayOfWeek = direct.Enum_ToProto[dayofweekpb.DayOfWeek](mapCtx, in.DayOfWeek)
	return out
}
func BackupDRBackupPlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupDRBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRBackupPlanSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.BackupRules = direct.Slice_FromProto(mapCtx, in.BackupRules, BackupRule_FromProto)
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Etag = direct.LazyPtr(in.GetEtag())
	if in.GetBackupVault() != "" {
		out.BackupVaultRef = &krm.BackupVaultRef{
			External: in.GetBackupVault(),
		}
	}
	return out
}
func BackupDRBackupPlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.BackupRules = direct.Slice_ToProto(mapCtx, in.BackupRules, BackupRule_ToProto)
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Etag = direct.ValueOf(in.Etag)
	if in.BackupVaultRef != nil {
		out.BackupVault = in.BackupVaultRef.External
	}
	return out
}
