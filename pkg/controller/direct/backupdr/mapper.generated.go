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
// krm.group: backupdr.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.backupdr.v1

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupDRBackupPlanAssociationObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krmv1beta1.BackupDRBackupPlanAssociationObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BackupDRBackupPlanAssociationObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RulesConfigInfo = direct.Slice_FromProto(mapCtx, in.RulesConfigInfo, RuleConfigInfoObservedState_v1beta1_FromProto)
	out.DataSource = direct.LazyPtr(in.GetDataSource())
	return out
}
func BackupDRBackupPlanAssociationObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BackupDRBackupPlanAssociationObservedState) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.BackupPlanAssociation_State](mapCtx, in.State)
	out.RulesConfigInfo = direct.Slice_ToProto(mapCtx, in.RulesConfigInfo, RuleConfigInfoObservedState_v1beta1_ToProto)
	out.DataSource = direct.ValueOf(in.DataSource)
	return out
}
func BackupDRBackupPlanObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krmv1alpha1.BackupDRBackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BackupDRBackupPlanObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.BackupVaultServiceAccount = direct.LazyPtr(in.GetBackupVaultServiceAccount())
	return out
}
func BackupDRBackupPlanObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BackupDRBackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.BackupPlan_State](mapCtx, in.State)
	out.BackupVaultServiceAccount = direct.ValueOf(in.BackupVaultServiceAccount)
	return out
}
func BackupDRBackupVaultSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krmv1alpha1.BackupDRBackupVaultSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BackupDRBackupVaultSpec{}
	// MISSING: Name
	out.Description = in.Description
	out.Labels = in.Labels
	out.BackupMinimumEnforcedRetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetBackupMinimumEnforcedRetentionDuration())
	out.EffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEffectiveTime())
	// MISSING: Uid
	out.Annotations = in.Annotations
	out.AccessRestriction = direct.Enum_FromProto(mapCtx, in.GetAccessRestriction())
	return out
}
func BackupDRManagementServerSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krmv1alpha1.BackupDRManagementServerSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BackupDRManagementServerSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfig_v1alpha1_FromProto)
	// MISSING: OAUTH2ClientID
	// MISSING: WorkforceIdentityBasedOAUTH2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupDRManagementServerSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BackupDRManagementServerSpec) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.ManagementServer_InstanceType](mapCtx, in.Type)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfig_v1alpha1_ToProto)
	// MISSING: OAUTH2ClientID
	// MISSING: WorkforceIdentityBasedOAUTH2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupRule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupRule) *krmv1alpha1.BackupRule {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BackupRule{}
	out.RuleID = direct.LazyPtr(in.GetRuleId())
	out.BackupRetentionDays = direct.LazyPtr(in.GetBackupRetentionDays())
	out.StandardSchedule = StandardSchedule_v1alpha1_FromProto(mapCtx, in.GetStandardSchedule())
	return out
}
func BackupRule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BackupRule) *pb.BackupRule {
	if in == nil {
		return nil
	}
	out := &pb.BackupRule{}
	out.RuleId = direct.ValueOf(in.RuleID)
	out.BackupRetentionDays = direct.ValueOf(in.BackupRetentionDays)
	if oneof := StandardSchedule_v1alpha1_ToProto(mapCtx, in.StandardSchedule); oneof != nil {
		out.BackupScheduleOneof = &pb.BackupRule_StandardSchedule{StandardSchedule: oneof}
	}
	return out
}
func BackupWindow_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupWindow) *krmv1alpha1.BackupWindow {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BackupWindow{}
	out.StartHourOfDay = direct.LazyPtr(in.GetStartHourOfDay())
	out.EndHourOfDay = direct.LazyPtr(in.GetEndHourOfDay())
	return out
}
func BackupWindow_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BackupWindow) *pb.BackupWindow {
	if in == nil {
		return nil
	}
	out := &pb.BackupWindow{}
	out.StartHourOfDay = direct.ValueOf(in.StartHourOfDay)
	out.EndHourOfDay = direct.ValueOf(in.EndHourOfDay)
	return out
}
func NetworkConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.PeeringMode = direct.Enum_ToProto[pb.NetworkConfig_PeeringMode](mapCtx, in.PeeringMode)
	return out
}
func RuleConfigInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RuleConfigInfo) *krmv1beta1.RuleConfigInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RuleConfigInfo{}
	// MISSING: RuleID
	// MISSING: LastBackupState
	// MISSING: LastBackupError
	// MISSING: LastSuccessfulBackupConsistencyTime
	return out
}
func RuleConfigInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RuleConfigInfo) *pb.RuleConfigInfo {
	if in == nil {
		return nil
	}
	out := &pb.RuleConfigInfo{}
	// MISSING: RuleID
	// MISSING: LastBackupState
	// MISSING: LastBackupError
	// MISSING: LastSuccessfulBackupConsistencyTime
	return out
}
func RuleConfigInfoObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RuleConfigInfo) *krmv1beta1.RuleConfigInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RuleConfigInfoObservedState{}
	out.RuleID = direct.LazyPtr(in.GetRuleId())
	out.LastBackupState = direct.Enum_FromProto(mapCtx, in.GetLastBackupState())
	out.LastBackupError = Status_v1beta1_FromProto(mapCtx, in.GetLastBackupError())
	out.LastSuccessfulBackupConsistencyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastSuccessfulBackupConsistencyTime())
	return out
}
func RuleConfigInfoObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RuleConfigInfoObservedState) *pb.RuleConfigInfo {
	if in == nil {
		return nil
	}
	out := &pb.RuleConfigInfo{}
	out.RuleId = direct.ValueOf(in.RuleID)
	out.LastBackupState = direct.Enum_ToProto[pb.RuleConfigInfo_LastBackupState](mapCtx, in.LastBackupState)
	out.LastBackupError = Status_v1beta1_ToProto(mapCtx, in.LastBackupError)
	out.LastSuccessfulBackupConsistencyTime = direct.StringTimestamp_ToProto(mapCtx, in.LastSuccessfulBackupConsistencyTime)
	return out
}
func StandardSchedule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.StandardSchedule) *krmv1alpha1.StandardSchedule {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.StandardSchedule{}
	out.RecurrenceType = direct.Enum_FromProto(mapCtx, in.GetRecurrenceType())
	out.HourlyFrequency = direct.LazyPtr(in.GetHourlyFrequency())
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	out.DaysOfMonth = in.DaysOfMonth
	out.WeekDayOfMonth = WeekDayOfMonth_v1alpha1_FromProto(mapCtx, in.GetWeekDayOfMonth())
	out.Months = direct.EnumSlice_FromProto(mapCtx, in.Months)
	out.BackupWindow = BackupWindow_v1alpha1_FromProto(mapCtx, in.GetBackupWindow())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	return out
}
func WeekDayOfMonth_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.WeekDayOfMonth) *krmv1alpha1.WeekDayOfMonth {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.WeekDayOfMonth{}
	out.WeekOfMonth = direct.Enum_FromProto(mapCtx, in.GetWeekOfMonth())
	out.DayOfWeek = direct.Enum_FromProto(mapCtx, in.GetDayOfWeek())
	return out
}
func WorkforceIdentityBasedManagementURI_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedManagementURI) *krmv1alpha1.WorkforceIdentityBasedManagementURI {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.WorkforceIdentityBasedManagementURI{}
	// MISSING: FirstPartyManagementURI
	// MISSING: ThirdPartyManagementURI
	return out
}
func WorkforceIdentityBasedManagementURI_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.WorkforceIdentityBasedManagementURI) *pb.WorkforceIdentityBasedManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedManagementURI{}
	// MISSING: FirstPartyManagementURI
	// MISSING: ThirdPartyManagementURI
	return out
}
func WorkforceIdentityBasedManagementURIObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedManagementURI) *krmv1alpha1.WorkforceIdentityBasedManagementURIObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.WorkforceIdentityBasedManagementURIObservedState{}
	out.FirstPartyManagementURI = direct.LazyPtr(in.GetFirstPartyManagementUri())
	out.ThirdPartyManagementURI = direct.LazyPtr(in.GetThirdPartyManagementUri())
	return out
}
func WorkforceIdentityBasedManagementURIObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.WorkforceIdentityBasedManagementURIObservedState) *pb.WorkforceIdentityBasedManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedManagementURI{}
	out.FirstPartyManagementUri = direct.ValueOf(in.FirstPartyManagementURI)
	out.ThirdPartyManagementUri = direct.ValueOf(in.ThirdPartyManagementURI)
	return out
}
