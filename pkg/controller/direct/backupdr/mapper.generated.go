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

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func BackupPlan_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupPlan {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlan{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.BackupRules = direct.Slice_FromProto(mapCtx, in.BackupRules, BackupRule_FromProto)
	// MISSING: State
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.BackupVault = direct.LazyPtr(in.GetBackupVault())
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupPlan_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlan) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.BackupRules = direct.Slice_ToProto(mapCtx, in.BackupRules, BackupRule_ToProto)
	// MISSING: State
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Etag = direct.ValueOf(in.Etag)
	out.BackupVault = direct.ValueOf(in.BackupVault)
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupPlanObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: BackupRules
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	out.BackupVaultServiceAccount = direct.LazyPtr(in.GetBackupVaultServiceAccount())
	return out
}
func BackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: BackupRules
	out.State = direct.Enum_ToProto[pb.BackupPlan_State](mapCtx, in.State)
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	out.BackupVaultServiceAccount = direct.ValueOf(in.BackupVaultServiceAccount)
	return out
}
func BackupRule_FromProto(mapCtx *direct.MapContext, in *pb.BackupRule) *krm.BackupRule {
	if in == nil {
		return nil
	}
	out := &krm.BackupRule{}
	out.RuleID = direct.LazyPtr(in.GetRuleId())
	out.BackupRetentionDays = direct.LazyPtr(in.GetBackupRetentionDays())
	out.StandardSchedule = StandardSchedule_FromProto(mapCtx, in.GetStandardSchedule())
	return out
}
func BackupRule_ToProto(mapCtx *direct.MapContext, in *krm.BackupRule) *pb.BackupRule {
	if in == nil {
		return nil
	}
	out := &pb.BackupRule{}
	out.RuleId = direct.ValueOf(in.RuleID)
	out.BackupRetentionDays = direct.ValueOf(in.BackupRetentionDays)
	if oneof := StandardSchedule_ToProto(mapCtx, in.StandardSchedule); oneof != nil {
		out.BackupScheduleOneof = &pb.BackupRule_StandardSchedule{StandardSchedule: oneof}
	}
	return out
}
func BackupWindow_FromProto(mapCtx *direct.MapContext, in *pb.BackupWindow) *krm.BackupWindow {
	if in == nil {
		return nil
	}
	out := &krm.BackupWindow{}
	out.StartHourOfDay = direct.LazyPtr(in.GetStartHourOfDay())
	out.EndHourOfDay = direct.LazyPtr(in.GetEndHourOfDay())
	return out
}
func BackupWindow_ToProto(mapCtx *direct.MapContext, in *krm.BackupWindow) *pb.BackupWindow {
	if in == nil {
		return nil
	}
	out := &pb.BackupWindow{}
	out.StartHourOfDay = direct.ValueOf(in.StartHourOfDay)
	out.EndHourOfDay = direct.ValueOf(in.EndHourOfDay)
	return out
}
func BackupdrBackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrManagementServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerObservedState) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerSpec) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func StandardSchedule_FromProto(mapCtx *direct.MapContext, in *pb.StandardSchedule) *krm.StandardSchedule {
	if in == nil {
		return nil
	}
	out := &krm.StandardSchedule{}
	out.RecurrenceType = direct.Enum_FromProto(mapCtx, in.GetRecurrenceType())
	out.HourlyFrequency = direct.LazyPtr(in.GetHourlyFrequency())
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	out.DaysOfMonth = in.DaysOfMonth
	out.WeekDayOfMonth = WeekDayOfMonth_FromProto(mapCtx, in.GetWeekDayOfMonth())
	out.Months = direct.EnumSlice_FromProto(mapCtx, in.Months)
	out.BackupWindow = BackupWindow_FromProto(mapCtx, in.GetBackupWindow())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	return out
}
func StandardSchedule_ToProto(mapCtx *direct.MapContext, in *krm.StandardSchedule) *pb.StandardSchedule {
	if in == nil {
		return nil
	}
	out := &pb.StandardSchedule{}
	out.RecurrenceType = direct.Enum_ToProto[pb.StandardSchedule_RecurrenceType](mapCtx, in.RecurrenceType)
	out.HourlyFrequency = direct.ValueOf(in.HourlyFrequency)
	out.DaysOfWeek = direct.EnumSlice_ToProto[pb.DayOfWeek](mapCtx, in.DaysOfWeek)
	out.DaysOfMonth = in.DaysOfMonth
	out.WeekDayOfMonth = WeekDayOfMonth_ToProto(mapCtx, in.WeekDayOfMonth)
	out.Months = direct.EnumSlice_ToProto[pb.Month](mapCtx, in.Months)
	out.BackupWindow = BackupWindow_ToProto(mapCtx, in.BackupWindow)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	return out
}
func WeekDayOfMonth_FromProto(mapCtx *direct.MapContext, in *pb.WeekDayOfMonth) *krm.WeekDayOfMonth {
	if in == nil {
		return nil
	}
	out := &krm.WeekDayOfMonth{}
	out.WeekOfMonth = direct.Enum_FromProto(mapCtx, in.GetWeekOfMonth())
	out.DayOfWeek = direct.Enum_FromProto(mapCtx, in.GetDayOfWeek())
	return out
}
func WeekDayOfMonth_ToProto(mapCtx *direct.MapContext, in *krm.WeekDayOfMonth) *pb.WeekDayOfMonth {
	if in == nil {
		return nil
	}
	out := &pb.WeekDayOfMonth{}
	out.WeekOfMonth = direct.Enum_ToProto[pb.WeekDayOfMonth_WeekOfMonth](mapCtx, in.WeekOfMonth)
	out.DayOfWeek = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.DayOfWeek)
	return out
}
