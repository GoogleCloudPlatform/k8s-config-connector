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

package osconfig

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AptSettings_FromProto(mapCtx *direct.MapContext, in *pb.AptSettings) *krm.AptSettings {
	if in == nil {
		return nil
	}
	out := &krm.AptSettings{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Excludes = in.Excludes
	out.ExclusivePackages = in.ExclusivePackages
	return out
}
func AptSettings_ToProto(mapCtx *direct.MapContext, in *krm.AptSettings) *pb.AptSettings {
	if in == nil {
		return nil
	}
	out := &pb.AptSettings{}
	out.Type = direct.Enum_ToProto[pb.AptSettings_Type](mapCtx, in.Type)
	out.Excludes = in.Excludes
	out.ExclusivePackages = in.ExclusivePackages
	return out
}
func ExecStep_FromProto(mapCtx *direct.MapContext, in *pb.ExecStep) *krm.ExecStep {
	if in == nil {
		return nil
	}
	out := &krm.ExecStep{}
	out.LinuxExecStepConfig = ExecStepConfig_FromProto(mapCtx, in.GetLinuxExecStepConfig())
	out.WindowsExecStepConfig = ExecStepConfig_FromProto(mapCtx, in.GetWindowsExecStepConfig())
	return out
}
func ExecStep_ToProto(mapCtx *direct.MapContext, in *krm.ExecStep) *pb.ExecStep {
	if in == nil {
		return nil
	}
	out := &pb.ExecStep{}
	out.LinuxExecStepConfig = ExecStepConfig_ToProto(mapCtx, in.LinuxExecStepConfig)
	out.WindowsExecStepConfig = ExecStepConfig_ToProto(mapCtx, in.WindowsExecStepConfig)
	return out
}
func ExecStepConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExecStepConfig) *krm.ExecStepConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExecStepConfig{}
	out.LocalPath = direct.LazyPtr(in.GetLocalPath())
	out.GcsObject = GcsObject_FromProto(mapCtx, in.GetGcsObject())
	out.AllowedSuccessCodes = in.AllowedSuccessCodes
	out.Interpreter = direct.Enum_FromProto(mapCtx, in.GetInterpreter())
	return out
}
func ExecStepConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExecStepConfig) *pb.ExecStepConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExecStepConfig{}
	if oneof := ExecStepConfig_LocalPath_ToProto(mapCtx, in.LocalPath); oneof != nil {
		out.Executable = oneof
	}
	if oneof := GcsObject_ToProto(mapCtx, in.GcsObject); oneof != nil {
		out.Executable = &pb.ExecStepConfig_GcsObject{GcsObject: oneof}
	}
	out.AllowedSuccessCodes = in.AllowedSuccessCodes
	out.Interpreter = direct.Enum_ToProto[pb.ExecStepConfig_Interpreter](mapCtx, in.Interpreter)
	return out
}
func FixedOrPercent_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &krm.FixedOrPercent{}
	out.Fixed = direct.LazyPtr(in.GetFixed())
	out.Percent = direct.LazyPtr(in.GetPercent())
	return out
}
func FixedOrPercent_ToProto(mapCtx *direct.MapContext, in *krm.FixedOrPercent) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	if oneof := FixedOrPercent_Fixed_ToProto(mapCtx, in.Fixed); oneof != nil {
		out.Mode = oneof
	}
	if oneof := FixedOrPercent_Percent_ToProto(mapCtx, in.Percent); oneof != nil {
		out.Mode = oneof
	}
	return out
}
func GcsObject_FromProto(mapCtx *direct.MapContext, in *pb.GcsObject) *krm.GcsObject {
	if in == nil {
		return nil
	}
	out := &krm.GcsObject{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.Object = direct.LazyPtr(in.GetObject())
	out.GenerationNumber = direct.LazyPtr(in.GetGenerationNumber())
	return out
}
func GcsObject_ToProto(mapCtx *direct.MapContext, in *krm.GcsObject) *pb.GcsObject {
	if in == nil {
		return nil
	}
	out := &pb.GcsObject{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.Object = direct.ValueOf(in.Object)
	out.GenerationNumber = direct.ValueOf(in.GenerationNumber)
	return out
}
func GooSettings_FromProto(mapCtx *direct.MapContext, in *pb.GooSettings) *krm.GooSettings {
	if in == nil {
		return nil
	}
	out := &krm.GooSettings{}
	return out
}
func GooSettings_ToProto(mapCtx *direct.MapContext, in *krm.GooSettings) *pb.GooSettings {
	if in == nil {
		return nil
	}
	out := &pb.GooSettings{}
	return out
}
func MonthlySchedule_FromProto(mapCtx *direct.MapContext, in *pb.MonthlySchedule) *krm.MonthlySchedule {
	if in == nil {
		return nil
	}
	out := &krm.MonthlySchedule{}
	out.WeekDayOfMonth = WeekDayOfMonth_FromProto(mapCtx, in.GetWeekDayOfMonth())
	out.MonthDay = direct.LazyPtr(in.GetMonthDay())
	return out
}
func MonthlySchedule_ToProto(mapCtx *direct.MapContext, in *krm.MonthlySchedule) *pb.MonthlySchedule {
	if in == nil {
		return nil
	}
	out := &pb.MonthlySchedule{}
	if oneof := WeekDayOfMonth_ToProto(mapCtx, in.WeekDayOfMonth); oneof != nil {
		out.DayOfMonth = &pb.MonthlySchedule_WeekDayOfMonth{WeekDayOfMonth: oneof}
	}
	if oneof := MonthlySchedule_MonthDay_ToProto(mapCtx, in.MonthDay); oneof != nil {
		out.DayOfMonth = oneof
	}
	return out
}
func OneTimeSchedule_FromProto(mapCtx *direct.MapContext, in *pb.OneTimeSchedule) *krm.OneTimeSchedule {
	if in == nil {
		return nil
	}
	out := &krm.OneTimeSchedule{}
	out.ExecuteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExecuteTime())
	return out
}
func OneTimeSchedule_ToProto(mapCtx *direct.MapContext, in *krm.OneTimeSchedule) *pb.OneTimeSchedule {
	if in == nil {
		return nil
	}
	out := &pb.OneTimeSchedule{}
	out.ExecuteTime = direct.StringTimestamp_ToProto(mapCtx, in.ExecuteTime)
	return out
}
func OsconfigPatchDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PatchDeployment) *krm.OsconfigPatchDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigPatchDeploymentObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: InstanceFilter
	// MISSING: PatchConfig
	// MISSING: Duration
	// MISSING: OneTimeSchedule
	// MISSING: RecurringSchedule
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LastExecuteTime
	// MISSING: Rollout
	// MISSING: State
	return out
}
func OsconfigPatchDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigPatchDeploymentObservedState) *pb.PatchDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PatchDeployment{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: InstanceFilter
	// MISSING: PatchConfig
	// MISSING: Duration
	// MISSING: OneTimeSchedule
	// MISSING: RecurringSchedule
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LastExecuteTime
	// MISSING: Rollout
	// MISSING: State
	return out
}
func OsconfigPatchDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.PatchDeployment) *krm.OsconfigPatchDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigPatchDeploymentSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: InstanceFilter
	// MISSING: PatchConfig
	// MISSING: Duration
	// MISSING: OneTimeSchedule
	// MISSING: RecurringSchedule
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LastExecuteTime
	// MISSING: Rollout
	// MISSING: State
	return out
}
func OsconfigPatchDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigPatchDeploymentSpec) *pb.PatchDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PatchDeployment{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: InstanceFilter
	// MISSING: PatchConfig
	// MISSING: Duration
	// MISSING: OneTimeSchedule
	// MISSING: RecurringSchedule
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LastExecuteTime
	// MISSING: Rollout
	// MISSING: State
	return out
}
func PatchConfig_FromProto(mapCtx *direct.MapContext, in *pb.PatchConfig) *krm.PatchConfig {
	if in == nil {
		return nil
	}
	out := &krm.PatchConfig{}
	out.RebootConfig = direct.Enum_FromProto(mapCtx, in.GetRebootConfig())
	out.Apt = AptSettings_FromProto(mapCtx, in.GetApt())
	out.Yum = YumSettings_FromProto(mapCtx, in.GetYum())
	out.Goo = GooSettings_FromProto(mapCtx, in.GetGoo())
	out.Zypper = ZypperSettings_FromProto(mapCtx, in.GetZypper())
	out.WindowsUpdate = WindowsUpdateSettings_FromProto(mapCtx, in.GetWindowsUpdate())
	out.PreStep = ExecStep_FromProto(mapCtx, in.GetPreStep())
	out.PostStep = ExecStep_FromProto(mapCtx, in.GetPostStep())
	out.MigInstancesAllowed = direct.LazyPtr(in.GetMigInstancesAllowed())
	return out
}
func PatchConfig_ToProto(mapCtx *direct.MapContext, in *krm.PatchConfig) *pb.PatchConfig {
	if in == nil {
		return nil
	}
	out := &pb.PatchConfig{}
	out.RebootConfig = direct.Enum_ToProto[pb.PatchConfig_RebootConfig](mapCtx, in.RebootConfig)
	out.Apt = AptSettings_ToProto(mapCtx, in.Apt)
	out.Yum = YumSettings_ToProto(mapCtx, in.Yum)
	out.Goo = GooSettings_ToProto(mapCtx, in.Goo)
	out.Zypper = ZypperSettings_ToProto(mapCtx, in.Zypper)
	out.WindowsUpdate = WindowsUpdateSettings_ToProto(mapCtx, in.WindowsUpdate)
	out.PreStep = ExecStep_ToProto(mapCtx, in.PreStep)
	out.PostStep = ExecStep_ToProto(mapCtx, in.PostStep)
	out.MigInstancesAllowed = direct.ValueOf(in.MigInstancesAllowed)
	return out
}
func PatchDeployment_FromProto(mapCtx *direct.MapContext, in *pb.PatchDeployment) *krm.PatchDeployment {
	if in == nil {
		return nil
	}
	out := &krm.PatchDeployment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.InstanceFilter = PatchInstanceFilter_FromProto(mapCtx, in.GetInstanceFilter())
	out.PatchConfig = PatchConfig_FromProto(mapCtx, in.GetPatchConfig())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.OneTimeSchedule = OneTimeSchedule_FromProto(mapCtx, in.GetOneTimeSchedule())
	out.RecurringSchedule = RecurringSchedule_FromProto(mapCtx, in.GetRecurringSchedule())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LastExecuteTime
	out.Rollout = PatchRollout_FromProto(mapCtx, in.GetRollout())
	// MISSING: State
	return out
}
func PatchDeployment_ToProto(mapCtx *direct.MapContext, in *krm.PatchDeployment) *pb.PatchDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PatchDeployment{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.InstanceFilter = PatchInstanceFilter_ToProto(mapCtx, in.InstanceFilter)
	out.PatchConfig = PatchConfig_ToProto(mapCtx, in.PatchConfig)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	if oneof := OneTimeSchedule_ToProto(mapCtx, in.OneTimeSchedule); oneof != nil {
		out.Schedule = &pb.PatchDeployment_OneTimeSchedule{OneTimeSchedule: oneof}
	}
	if oneof := RecurringSchedule_ToProto(mapCtx, in.RecurringSchedule); oneof != nil {
		out.Schedule = &pb.PatchDeployment_RecurringSchedule{RecurringSchedule: oneof}
	}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LastExecuteTime
	out.Rollout = PatchRollout_ToProto(mapCtx, in.Rollout)
	// MISSING: State
	return out
}
func PatchDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PatchDeployment) *krm.PatchDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PatchDeploymentObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: InstanceFilter
	// MISSING: PatchConfig
	// MISSING: Duration
	// MISSING: OneTimeSchedule
	out.RecurringSchedule = RecurringScheduleObservedState_FromProto(mapCtx, in.GetRecurringSchedule())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LastExecuteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastExecuteTime())
	// MISSING: Rollout
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func PatchDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PatchDeploymentObservedState) *pb.PatchDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PatchDeployment{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: InstanceFilter
	// MISSING: PatchConfig
	// MISSING: Duration
	// MISSING: OneTimeSchedule
	if oneof := RecurringScheduleObservedState_ToProto(mapCtx, in.RecurringSchedule); oneof != nil {
		out.Schedule = &pb.PatchDeployment_RecurringSchedule{RecurringSchedule: oneof}
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LastExecuteTime = direct.StringTimestamp_ToProto(mapCtx, in.LastExecuteTime)
	// MISSING: Rollout
	out.State = direct.Enum_ToProto[pb.PatchDeployment_State](mapCtx, in.State)
	return out
}
func PatchInstanceFilter_FromProto(mapCtx *direct.MapContext, in *pb.PatchInstanceFilter) *krm.PatchInstanceFilter {
	if in == nil {
		return nil
	}
	out := &krm.PatchInstanceFilter{}
	out.All = direct.LazyPtr(in.GetAll())
	out.GroupLabels = direct.Slice_FromProto(mapCtx, in.GroupLabels, PatchInstanceFilter_GroupLabel_FromProto)
	out.Zones = in.Zones
	out.Instances = in.Instances
	out.InstanceNamePrefixes = in.InstanceNamePrefixes
	return out
}
func PatchInstanceFilter_ToProto(mapCtx *direct.MapContext, in *krm.PatchInstanceFilter) *pb.PatchInstanceFilter {
	if in == nil {
		return nil
	}
	out := &pb.PatchInstanceFilter{}
	out.All = direct.ValueOf(in.All)
	out.GroupLabels = direct.Slice_ToProto(mapCtx, in.GroupLabels, PatchInstanceFilter_GroupLabel_ToProto)
	out.Zones = in.Zones
	out.Instances = in.Instances
	out.InstanceNamePrefixes = in.InstanceNamePrefixes
	return out
}
func PatchInstanceFilter_GroupLabel_FromProto(mapCtx *direct.MapContext, in *pb.PatchInstanceFilter_GroupLabel) *krm.PatchInstanceFilter_GroupLabel {
	if in == nil {
		return nil
	}
	out := &krm.PatchInstanceFilter_GroupLabel{}
	out.Labels = in.Labels
	return out
}
func PatchInstanceFilter_GroupLabel_ToProto(mapCtx *direct.MapContext, in *krm.PatchInstanceFilter_GroupLabel) *pb.PatchInstanceFilter_GroupLabel {
	if in == nil {
		return nil
	}
	out := &pb.PatchInstanceFilter_GroupLabel{}
	out.Labels = in.Labels
	return out
}
func PatchRollout_FromProto(mapCtx *direct.MapContext, in *pb.PatchRollout) *krm.PatchRollout {
	if in == nil {
		return nil
	}
	out := &krm.PatchRollout{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.DisruptionBudget = FixedOrPercent_FromProto(mapCtx, in.GetDisruptionBudget())
	return out
}
func PatchRollout_ToProto(mapCtx *direct.MapContext, in *krm.PatchRollout) *pb.PatchRollout {
	if in == nil {
		return nil
	}
	out := &pb.PatchRollout{}
	out.Mode = direct.Enum_ToProto[pb.PatchRollout_Mode](mapCtx, in.Mode)
	out.DisruptionBudget = FixedOrPercent_ToProto(mapCtx, in.DisruptionBudget)
	return out
}
func RecurringSchedule_FromProto(mapCtx *direct.MapContext, in *pb.RecurringSchedule) *krm.RecurringSchedule {
	if in == nil {
		return nil
	}
	out := &krm.RecurringSchedule{}
	out.TimeZone = TimeZone_FromProto(mapCtx, in.GetTimeZone())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.TimeOfDay = TimeOfDay_FromProto(mapCtx, in.GetTimeOfDay())
	out.Frequency = direct.Enum_FromProto(mapCtx, in.GetFrequency())
	out.Weekly = WeeklySchedule_FromProto(mapCtx, in.GetWeekly())
	out.Monthly = MonthlySchedule_FromProto(mapCtx, in.GetMonthly())
	// MISSING: LastExecuteTime
	// MISSING: NextExecuteTime
	return out
}
func RecurringSchedule_ToProto(mapCtx *direct.MapContext, in *krm.RecurringSchedule) *pb.RecurringSchedule {
	if in == nil {
		return nil
	}
	out := &pb.RecurringSchedule{}
	out.TimeZone = TimeZone_ToProto(mapCtx, in.TimeZone)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.TimeOfDay = TimeOfDay_ToProto(mapCtx, in.TimeOfDay)
	out.Frequency = direct.Enum_ToProto[pb.RecurringSchedule_Frequency](mapCtx, in.Frequency)
	if oneof := WeeklySchedule_ToProto(mapCtx, in.Weekly); oneof != nil {
		out.ScheduleConfig = &pb.RecurringSchedule_Weekly{Weekly: oneof}
	}
	if oneof := MonthlySchedule_ToProto(mapCtx, in.Monthly); oneof != nil {
		out.ScheduleConfig = &pb.RecurringSchedule_Monthly{Monthly: oneof}
	}
	// MISSING: LastExecuteTime
	// MISSING: NextExecuteTime
	return out
}
func RecurringScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RecurringSchedule) *krm.RecurringScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecurringScheduleObservedState{}
	// MISSING: TimeZone
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TimeOfDay
	// MISSING: Frequency
	// MISSING: Weekly
	// MISSING: Monthly
	out.LastExecuteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastExecuteTime())
	out.NextExecuteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextExecuteTime())
	return out
}
func RecurringScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecurringScheduleObservedState) *pb.RecurringSchedule {
	if in == nil {
		return nil
	}
	out := &pb.RecurringSchedule{}
	// MISSING: TimeZone
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TimeOfDay
	// MISSING: Frequency
	// MISSING: Weekly
	// MISSING: Monthly
	out.LastExecuteTime = direct.StringTimestamp_ToProto(mapCtx, in.LastExecuteTime)
	out.NextExecuteTime = direct.StringTimestamp_ToProto(mapCtx, in.NextExecuteTime)
	return out
}
func WeekDayOfMonth_FromProto(mapCtx *direct.MapContext, in *pb.WeekDayOfMonth) *krm.WeekDayOfMonth {
	if in == nil {
		return nil
	}
	out := &krm.WeekDayOfMonth{}
	out.WeekOrdinal = direct.LazyPtr(in.GetWeekOrdinal())
	out.DayOfWeek = direct.Enum_FromProto(mapCtx, in.GetDayOfWeek())
	out.DayOffset = direct.LazyPtr(in.GetDayOffset())
	return out
}
func WeekDayOfMonth_ToProto(mapCtx *direct.MapContext, in *krm.WeekDayOfMonth) *pb.WeekDayOfMonth {
	if in == nil {
		return nil
	}
	out := &pb.WeekDayOfMonth{}
	out.WeekOrdinal = direct.ValueOf(in.WeekOrdinal)
	out.DayOfWeek = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.DayOfWeek)
	out.DayOffset = direct.ValueOf(in.DayOffset)
	return out
}
func WeeklySchedule_FromProto(mapCtx *direct.MapContext, in *pb.WeeklySchedule) *krm.WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &krm.WeeklySchedule{}
	out.DayOfWeek = direct.Enum_FromProto(mapCtx, in.GetDayOfWeek())
	return out
}
func WeeklySchedule_ToProto(mapCtx *direct.MapContext, in *krm.WeeklySchedule) *pb.WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &pb.WeeklySchedule{}
	out.DayOfWeek = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.DayOfWeek)
	return out
}
func WindowsUpdateSettings_FromProto(mapCtx *direct.MapContext, in *pb.WindowsUpdateSettings) *krm.WindowsUpdateSettings {
	if in == nil {
		return nil
	}
	out := &krm.WindowsUpdateSettings{}
	out.Classifications = direct.EnumSlice_FromProto(mapCtx, in.Classifications)
	out.Excludes = in.Excludes
	out.ExclusivePatches = in.ExclusivePatches
	return out
}
func WindowsUpdateSettings_ToProto(mapCtx *direct.MapContext, in *krm.WindowsUpdateSettings) *pb.WindowsUpdateSettings {
	if in == nil {
		return nil
	}
	out := &pb.WindowsUpdateSettings{}
	out.Classifications = direct.EnumSlice_ToProto[pb.WindowsUpdateSettings_Classification](mapCtx, in.Classifications)
	out.Excludes = in.Excludes
	out.ExclusivePatches = in.ExclusivePatches
	return out
}
func YumSettings_FromProto(mapCtx *direct.MapContext, in *pb.YumSettings) *krm.YumSettings {
	if in == nil {
		return nil
	}
	out := &krm.YumSettings{}
	out.Security = direct.LazyPtr(in.GetSecurity())
	out.Minimal = direct.LazyPtr(in.GetMinimal())
	out.Excludes = in.Excludes
	out.ExclusivePackages = in.ExclusivePackages
	return out
}
func YumSettings_ToProto(mapCtx *direct.MapContext, in *krm.YumSettings) *pb.YumSettings {
	if in == nil {
		return nil
	}
	out := &pb.YumSettings{}
	out.Security = direct.ValueOf(in.Security)
	out.Minimal = direct.ValueOf(in.Minimal)
	out.Excludes = in.Excludes
	out.ExclusivePackages = in.ExclusivePackages
	return out
}
func ZypperSettings_FromProto(mapCtx *direct.MapContext, in *pb.ZypperSettings) *krm.ZypperSettings {
	if in == nil {
		return nil
	}
	out := &krm.ZypperSettings{}
	out.WithOptional = direct.LazyPtr(in.GetWithOptional())
	out.WithUpdate = direct.LazyPtr(in.GetWithUpdate())
	out.Categories = in.Categories
	out.Severities = in.Severities
	out.Excludes = in.Excludes
	out.ExclusivePatches = in.ExclusivePatches
	return out
}
func ZypperSettings_ToProto(mapCtx *direct.MapContext, in *krm.ZypperSettings) *pb.ZypperSettings {
	if in == nil {
		return nil
	}
	out := &pb.ZypperSettings{}
	out.WithOptional = direct.ValueOf(in.WithOptional)
	out.WithUpdate = direct.ValueOf(in.WithUpdate)
	out.Categories = in.Categories
	out.Severities = in.Severities
	out.Excludes = in.Excludes
	out.ExclusivePatches = in.ExclusivePatches
	return out
}
