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

package migrationcenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ExecutionReport_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionReport) *krm.ExecutionReport {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionReport{}
	out.FramesReported = direct.LazyPtr(in.GetFramesReported())
	out.ExecutionErrors = ValidationReport_FromProto(mapCtx, in.GetExecutionErrors())
	// MISSING: TotalRowsCount
	return out
}
func ExecutionReport_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionReport) *pb.ExecutionReport {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionReport{}
	out.FramesReported = direct.ValueOf(in.FramesReported)
	out.ExecutionErrors = ValidationReport_ToProto(mapCtx, in.ExecutionErrors)
	// MISSING: TotalRowsCount
	return out
}
func ExecutionReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionReport) *krm.ExecutionReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionReportObservedState{}
	// MISSING: FramesReported
	// MISSING: ExecutionErrors
	out.TotalRowsCount = direct.LazyPtr(in.GetTotalRowsCount())
	return out
}
func ExecutionReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionReportObservedState) *pb.ExecutionReport {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionReport{}
	// MISSING: FramesReported
	// MISSING: ExecutionErrors
	out.TotalRowsCount = direct.ValueOf(in.TotalRowsCount)
	return out
}
func FileValidationReport_FromProto(mapCtx *direct.MapContext, in *pb.FileValidationReport) *krm.FileValidationReport {
	if in == nil {
		return nil
	}
	out := &krm.FileValidationReport{}
	out.FileName = direct.LazyPtr(in.GetFileName())
	out.RowErrors = direct.Slice_FromProto(mapCtx, in.RowErrors, ImportRowError_FromProto)
	out.PartialReport = direct.LazyPtr(in.GetPartialReport())
	out.FileErrors = direct.Slice_FromProto(mapCtx, in.FileErrors, ImportError_FromProto)
	return out
}
func FileValidationReport_ToProto(mapCtx *direct.MapContext, in *krm.FileValidationReport) *pb.FileValidationReport {
	if in == nil {
		return nil
	}
	out := &pb.FileValidationReport{}
	out.FileName = direct.ValueOf(in.FileName)
	out.RowErrors = direct.Slice_ToProto(mapCtx, in.RowErrors, ImportRowError_ToProto)
	out.PartialReport = direct.ValueOf(in.PartialReport)
	out.FileErrors = direct.Slice_ToProto(mapCtx, in.FileErrors, ImportError_ToProto)
	return out
}
func ImportError_FromProto(mapCtx *direct.MapContext, in *pb.ImportError) *krm.ImportError {
	if in == nil {
		return nil
	}
	out := &krm.ImportError{}
	out.ErrorDetails = direct.LazyPtr(in.GetErrorDetails())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	return out
}
func ImportError_ToProto(mapCtx *direct.MapContext, in *krm.ImportError) *pb.ImportError {
	if in == nil {
		return nil
	}
	out := &pb.ImportError{}
	out.ErrorDetails = direct.ValueOf(in.ErrorDetails)
	out.Severity = direct.Enum_ToProto[pb.ImportError_Severity](mapCtx, in.Severity)
	return out
}
func ImportJob_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.ImportJob {
	if in == nil {
		return nil
	}
	out := &krm.ImportJob{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CompleteTime
	// MISSING: State
	out.Labels = in.Labels
	out.AssetSource = direct.LazyPtr(in.GetAssetSource())
	// MISSING: ValidationReport
	// MISSING: ExecutionReport
	return out
}
func ImportJob_ToProto(mapCtx *direct.MapContext, in *krm.ImportJob) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CompleteTime
	// MISSING: State
	out.Labels = in.Labels
	out.AssetSource = direct.ValueOf(in.AssetSource)
	// MISSING: ValidationReport
	// MISSING: ExecutionReport
	return out
}
func ImportJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.ImportJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ImportJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Labels
	// MISSING: AssetSource
	out.ValidationReport = ValidationReport_FromProto(mapCtx, in.GetValidationReport())
	out.ExecutionReport = ExecutionReport_FromProto(mapCtx, in.GetExecutionReport())
	return out
}
func ImportJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ImportJobObservedState) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.State = direct.Enum_ToProto[pb.ImportJob_ImportJobState](mapCtx, in.State)
	// MISSING: Labels
	// MISSING: AssetSource
	if oneof := ValidationReport_ToProto(mapCtx, in.ValidationReport); oneof != nil {
		out.Report = &pb.ImportJob_ValidationReport{ValidationReport: oneof}
	}
	if oneof := ExecutionReport_ToProto(mapCtx, in.ExecutionReport); oneof != nil {
		out.Report = &pb.ImportJob_ExecutionReport{ExecutionReport: oneof}
	}
	return out
}
func ImportRowError_FromProto(mapCtx *direct.MapContext, in *pb.ImportRowError) *krm.ImportRowError {
	if in == nil {
		return nil
	}
	out := &krm.ImportRowError{}
	out.RowNumber = direct.LazyPtr(in.GetRowNumber())
	out.VmName = direct.LazyPtr(in.GetVmName())
	out.VmUuid = direct.LazyPtr(in.GetVmUuid())
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, ImportError_FromProto)
	return out
}
func ImportRowError_ToProto(mapCtx *direct.MapContext, in *krm.ImportRowError) *pb.ImportRowError {
	if in == nil {
		return nil
	}
	out := &pb.ImportRowError{}
	out.RowNumber = direct.ValueOf(in.RowNumber)
	out.VmName = direct.ValueOf(in.VmName)
	out.VmUuid = direct.ValueOf(in.VmUuid)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, ImportError_ToProto)
	return out
}
func MigrationcenterImportJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.MigrationcenterImportJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterImportJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: AssetSource
	// MISSING: ValidationReport
	// MISSING: ExecutionReport
	return out
}
func MigrationcenterImportJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterImportJobObservedState) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: AssetSource
	// MISSING: ValidationReport
	// MISSING: ExecutionReport
	return out
}
func MigrationcenterImportJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.MigrationcenterImportJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterImportJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: AssetSource
	// MISSING: ValidationReport
	// MISSING: ExecutionReport
	return out
}
func MigrationcenterImportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterImportJobSpec) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: AssetSource
	// MISSING: ValidationReport
	// MISSING: ExecutionReport
	return out
}
func ValidationReport_FromProto(mapCtx *direct.MapContext, in *pb.ValidationReport) *krm.ValidationReport {
	if in == nil {
		return nil
	}
	out := &krm.ValidationReport{}
	out.FileValidations = direct.Slice_FromProto(mapCtx, in.FileValidations, FileValidationReport_FromProto)
	out.JobErrors = direct.Slice_FromProto(mapCtx, in.JobErrors, ImportError_FromProto)
	return out
}
func ValidationReport_ToProto(mapCtx *direct.MapContext, in *krm.ValidationReport) *pb.ValidationReport {
	if in == nil {
		return nil
	}
	out := &pb.ValidationReport{}
	out.FileValidations = direct.Slice_ToProto(mapCtx, in.FileValidations, FileValidationReport_ToProto)
	out.JobErrors = direct.Slice_ToProto(mapCtx, in.JobErrors, ImportError_ToProto)
	return out
}
