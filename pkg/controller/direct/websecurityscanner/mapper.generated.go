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
func WebsecurityscannerScanRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ScanRun) *krm.WebsecurityscannerScanRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WebsecurityscannerScanRunObservedState{}
	// MISSING: Name
	// MISSING: ExecutionState
	// MISSING: ResultState
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UrlsCrawledCount
	// MISSING: UrlsTestedCount
	// MISSING: HasVulnerabilities
	// MISSING: ProgressPercent
	// MISSING: ErrorTrace
	// MISSING: WarningTraces
	return out
}
func WebsecurityscannerScanRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WebsecurityscannerScanRunObservedState) *pb.ScanRun {
	if in == nil {
		return nil
	}
	out := &pb.ScanRun{}
	// MISSING: Name
	// MISSING: ExecutionState
	// MISSING: ResultState
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UrlsCrawledCount
	// MISSING: UrlsTestedCount
	// MISSING: HasVulnerabilities
	// MISSING: ProgressPercent
	// MISSING: ErrorTrace
	// MISSING: WarningTraces
	return out
}
func WebsecurityscannerScanRunSpec_FromProto(mapCtx *direct.MapContext, in *pb.ScanRun) *krm.WebsecurityscannerScanRunSpec {
	if in == nil {
		return nil
	}
	out := &krm.WebsecurityscannerScanRunSpec{}
	// MISSING: Name
	// MISSING: ExecutionState
	// MISSING: ResultState
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UrlsCrawledCount
	// MISSING: UrlsTestedCount
	// MISSING: HasVulnerabilities
	// MISSING: ProgressPercent
	// MISSING: ErrorTrace
	// MISSING: WarningTraces
	return out
}
func WebsecurityscannerScanRunSpec_ToProto(mapCtx *direct.MapContext, in *krm.WebsecurityscannerScanRunSpec) *pb.ScanRun {
	if in == nil {
		return nil
	}
	out := &pb.ScanRun{}
	// MISSING: Name
	// MISSING: ExecutionState
	// MISSING: ResultState
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UrlsCrawledCount
	// MISSING: UrlsTestedCount
	// MISSING: HasVulnerabilities
	// MISSING: ProgressPercent
	// MISSING: ErrorTrace
	// MISSING: WarningTraces
	return out
}
