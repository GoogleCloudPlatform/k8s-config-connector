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

package websecurityscanner

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/websecurityscanner/apiv1beta/websecurityscannerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/websecurityscanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ScanConfig_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig) *krm.ScanConfig {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.MaxQps = direct.LazyPtr(in.GetMaxQps())
	out.StartingUrls = in.StartingUrls
	out.Authentication = ScanConfig_Authentication_FromProto(mapCtx, in.GetAuthentication())
	out.UserAgent = direct.Enum_FromProto(mapCtx, in.GetUserAgent())
	out.BlacklistPatterns = in.BlacklistPatterns
	out.Schedule = ScanConfig_Schedule_FromProto(mapCtx, in.GetSchedule())
	out.TargetPlatforms = direct.EnumSlice_FromProto(mapCtx, in.TargetPlatforms)
	out.ExportToSecurityCommandCenter = direct.Enum_FromProto(mapCtx, in.GetExportToSecurityCommandCenter())
	out.LatestRun = ScanRun_FromProto(mapCtx, in.GetLatestRun())
	out.RiskLevel = direct.Enum_FromProto(mapCtx, in.GetRiskLevel())
	return out
}
func ScanConfig_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfig) *pb.ScanConfig {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.MaxQps = direct.ValueOf(in.MaxQps)
	out.StartingUrls = in.StartingUrls
	out.Authentication = ScanConfig_Authentication_ToProto(mapCtx, in.Authentication)
	out.UserAgent = direct.Enum_ToProto[pb.ScanConfig_UserAgent](mapCtx, in.UserAgent)
	out.BlacklistPatterns = in.BlacklistPatterns
	out.Schedule = ScanConfig_Schedule_ToProto(mapCtx, in.Schedule)
	out.TargetPlatforms = direct.EnumSlice_ToProto[pb.ScanConfig_TargetPlatform](mapCtx, in.TargetPlatforms)
	out.ExportToSecurityCommandCenter = direct.Enum_ToProto[pb.ScanConfig_ExportToSecurityCommandCenter](mapCtx, in.ExportToSecurityCommandCenter)
	out.LatestRun = ScanRun_ToProto(mapCtx, in.LatestRun)
	out.RiskLevel = direct.Enum_ToProto[pb.ScanConfig_RiskLevel](mapCtx, in.RiskLevel)
	return out
}
func ScanConfigError_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfigError) *krm.ScanConfigError {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfigError{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.FieldName = direct.LazyPtr(in.GetFieldName())
	return out
}
func ScanConfigError_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfigError) *pb.ScanConfigError {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfigError{}
	out.Code = direct.Enum_ToProto[pb.ScanConfigError_Code](mapCtx, in.Code)
	out.FieldName = direct.ValueOf(in.FieldName)
	return out
}
func ScanConfig_Authentication_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig_Authentication) *krm.ScanConfig_Authentication {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfig_Authentication{}
	out.GoogleAccount = ScanConfig_Authentication_GoogleAccount_FromProto(mapCtx, in.GetGoogleAccount())
	out.CustomAccount = ScanConfig_Authentication_CustomAccount_FromProto(mapCtx, in.GetCustomAccount())
	return out
}
func ScanConfig_Authentication_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfig_Authentication) *pb.ScanConfig_Authentication {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig_Authentication{}
	if oneof := ScanConfig_Authentication_GoogleAccount_ToProto(mapCtx, in.GoogleAccount); oneof != nil {
		out.Authentication = &pb.ScanConfig_Authentication_GoogleAccount_{GoogleAccount: oneof}
	}
	if oneof := ScanConfig_Authentication_CustomAccount_ToProto(mapCtx, in.CustomAccount); oneof != nil {
		out.Authentication = &pb.ScanConfig_Authentication_CustomAccount_{CustomAccount: oneof}
	}
	return out
}
func ScanConfig_Authentication_CustomAccount_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig_Authentication_CustomAccount) *krm.ScanConfig_Authentication_CustomAccount {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfig_Authentication_CustomAccount{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.LoginURL = direct.LazyPtr(in.GetLoginUrl())
	return out
}
func ScanConfig_Authentication_CustomAccount_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfig_Authentication_CustomAccount) *pb.ScanConfig_Authentication_CustomAccount {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig_Authentication_CustomAccount{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.LoginUrl = direct.ValueOf(in.LoginURL)
	return out
}
func ScanConfig_Authentication_GoogleAccount_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig_Authentication_GoogleAccount) *krm.ScanConfig_Authentication_GoogleAccount {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfig_Authentication_GoogleAccount{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	return out
}
func ScanConfig_Authentication_GoogleAccount_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfig_Authentication_GoogleAccount) *pb.ScanConfig_Authentication_GoogleAccount {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig_Authentication_GoogleAccount{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	return out
}
func ScanConfig_Schedule_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig_Schedule) *krm.ScanConfig_Schedule {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfig_Schedule{}
	out.ScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleTime())
	out.IntervalDurationDays = direct.LazyPtr(in.GetIntervalDurationDays())
	return out
}
func ScanConfig_Schedule_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfig_Schedule) *pb.ScanConfig_Schedule {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig_Schedule{}
	out.ScheduleTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleTime)
	out.IntervalDurationDays = direct.ValueOf(in.IntervalDurationDays)
	return out
}
func ScanRun_FromProto(mapCtx *direct.MapContext, in *pb.ScanRun) *krm.ScanRun {
	if in == nil {
		return nil
	}
	out := &krm.ScanRun{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExecutionState = direct.Enum_FromProto(mapCtx, in.GetExecutionState())
	out.ResultState = direct.Enum_FromProto(mapCtx, in.GetResultState())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UrlsCrawledCount = direct.LazyPtr(in.GetUrlsCrawledCount())
	out.UrlsTestedCount = direct.LazyPtr(in.GetUrlsTestedCount())
	out.HasVulnerabilities = direct.LazyPtr(in.GetHasVulnerabilities())
	out.ProgressPercent = direct.LazyPtr(in.GetProgressPercent())
	out.ErrorTrace = ScanRunErrorTrace_FromProto(mapCtx, in.GetErrorTrace())
	out.WarningTraces = direct.Slice_FromProto(mapCtx, in.WarningTraces, ScanRunWarningTrace_FromProto)
	return out
}
func ScanRun_ToProto(mapCtx *direct.MapContext, in *krm.ScanRun) *pb.ScanRun {
	if in == nil {
		return nil
	}
	out := &pb.ScanRun{}
	out.Name = direct.ValueOf(in.Name)
	out.ExecutionState = direct.Enum_ToProto[pb.ScanRun_ExecutionState](mapCtx, in.ExecutionState)
	out.ResultState = direct.Enum_ToProto[pb.ScanRun_ResultState](mapCtx, in.ResultState)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UrlsCrawledCount = direct.ValueOf(in.UrlsCrawledCount)
	out.UrlsTestedCount = direct.ValueOf(in.UrlsTestedCount)
	out.HasVulnerabilities = direct.ValueOf(in.HasVulnerabilities)
	out.ProgressPercent = direct.ValueOf(in.ProgressPercent)
	out.ErrorTrace = ScanRunErrorTrace_ToProto(mapCtx, in.ErrorTrace)
	out.WarningTraces = direct.Slice_ToProto(mapCtx, in.WarningTraces, ScanRunWarningTrace_ToProto)
	return out
}
func ScanRunErrorTrace_FromProto(mapCtx *direct.MapContext, in *pb.ScanRunErrorTrace) *krm.ScanRunErrorTrace {
	if in == nil {
		return nil
	}
	out := &krm.ScanRunErrorTrace{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.ScanConfigError = ScanConfigError_FromProto(mapCtx, in.GetScanConfigError())
	out.MostCommonHTTPErrorCode = direct.LazyPtr(in.GetMostCommonHttpErrorCode())
	return out
}
func ScanRunErrorTrace_ToProto(mapCtx *direct.MapContext, in *krm.ScanRunErrorTrace) *pb.ScanRunErrorTrace {
	if in == nil {
		return nil
	}
	out := &pb.ScanRunErrorTrace{}
	out.Code = direct.Enum_ToProto[pb.ScanRunErrorTrace_Code](mapCtx, in.Code)
	out.ScanConfigError = ScanConfigError_ToProto(mapCtx, in.ScanConfigError)
	out.MostCommonHttpErrorCode = direct.ValueOf(in.MostCommonHTTPErrorCode)
	return out
}
func ScanRunWarningTrace_FromProto(mapCtx *direct.MapContext, in *pb.ScanRunWarningTrace) *krm.ScanRunWarningTrace {
	if in == nil {
		return nil
	}
	out := &krm.ScanRunWarningTrace{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	return out
}
func ScanRunWarningTrace_ToProto(mapCtx *direct.MapContext, in *krm.ScanRunWarningTrace) *pb.ScanRunWarningTrace {
	if in == nil {
		return nil
	}
	out := &pb.ScanRunWarningTrace{}
	out.Code = direct.Enum_ToProto[pb.ScanRunWarningTrace_Code](mapCtx, in.Code)
	return out
}
func WebsecurityscannerScanConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig) *krm.WebsecurityscannerScanConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WebsecurityscannerScanConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MaxQps
	// MISSING: StartingUrls
	// MISSING: Authentication
	// MISSING: UserAgent
	// MISSING: BlacklistPatterns
	// MISSING: Schedule
	// MISSING: TargetPlatforms
	// MISSING: ExportToSecurityCommandCenter
	// MISSING: LatestRun
	// MISSING: RiskLevel
	return out
}
func WebsecurityscannerScanConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WebsecurityscannerScanConfigObservedState) *pb.ScanConfig {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MaxQps
	// MISSING: StartingUrls
	// MISSING: Authentication
	// MISSING: UserAgent
	// MISSING: BlacklistPatterns
	// MISSING: Schedule
	// MISSING: TargetPlatforms
	// MISSING: ExportToSecurityCommandCenter
	// MISSING: LatestRun
	// MISSING: RiskLevel
	return out
}
func WebsecurityscannerScanConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig) *krm.WebsecurityscannerScanConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.WebsecurityscannerScanConfigSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MaxQps
	// MISSING: StartingUrls
	// MISSING: Authentication
	// MISSING: UserAgent
	// MISSING: BlacklistPatterns
	// MISSING: Schedule
	// MISSING: TargetPlatforms
	// MISSING: ExportToSecurityCommandCenter
	// MISSING: LatestRun
	// MISSING: RiskLevel
	return out
}
func WebsecurityscannerScanConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.WebsecurityscannerScanConfigSpec) *pb.ScanConfig {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MaxQps
	// MISSING: StartingUrls
	// MISSING: Authentication
	// MISSING: UserAgent
	// MISSING: BlacklistPatterns
	// MISSING: Schedule
	// MISSING: TargetPlatforms
	// MISSING: ExportToSecurityCommandCenter
	// MISSING: LatestRun
	// MISSING: RiskLevel
	return out
}
