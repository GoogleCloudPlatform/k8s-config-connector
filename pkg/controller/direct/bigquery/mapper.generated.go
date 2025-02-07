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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/migration/apiv2/migrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BigqueryMigrationSubtaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSubtask) *krm.BigqueryMigrationSubtaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryMigrationSubtaskObservedState{}
	// MISSING: Name
	// MISSING: TaskID
	// MISSING: Type
	// MISSING: State
	// MISSING: ProcessingError
	// MISSING: ResourceErrorDetails
	// MISSING: ResourceErrorCount
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	// MISSING: Metrics
	return out
}
func BigqueryMigrationSubtaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryMigrationSubtaskObservedState) *pb.MigrationSubtask {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSubtask{}
	// MISSING: Name
	// MISSING: TaskID
	// MISSING: Type
	// MISSING: State
	// MISSING: ProcessingError
	// MISSING: ResourceErrorDetails
	// MISSING: ResourceErrorCount
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	// MISSING: Metrics
	return out
}
func BigqueryMigrationSubtaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSubtask) *krm.BigqueryMigrationSubtaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryMigrationSubtaskSpec{}
	// MISSING: Name
	// MISSING: TaskID
	// MISSING: Type
	// MISSING: State
	// MISSING: ProcessingError
	// MISSING: ResourceErrorDetails
	// MISSING: ResourceErrorCount
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	// MISSING: Metrics
	return out
}
func BigqueryMigrationSubtaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryMigrationSubtaskSpec) *pb.MigrationSubtask {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSubtask{}
	// MISSING: Name
	// MISSING: TaskID
	// MISSING: Type
	// MISSING: State
	// MISSING: ProcessingError
	// MISSING: ResourceErrorDetails
	// MISSING: ResourceErrorCount
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	// MISSING: Metrics
	return out
}
func BigqueryMigrationWorkflowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationWorkflow) *krm.BigqueryMigrationWorkflowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryMigrationWorkflowObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func BigqueryMigrationWorkflowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryMigrationWorkflowObservedState) *pb.MigrationWorkflow {
	if in == nil {
		return nil
	}
	out := &pb.MigrationWorkflow{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func BigqueryMigrationWorkflowSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationWorkflow) *krm.BigqueryMigrationWorkflowSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryMigrationWorkflowSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func BigqueryMigrationWorkflowSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryMigrationWorkflowSpec) *pb.MigrationWorkflow {
	if in == nil {
		return nil
	}
	out := &pb.MigrationWorkflow{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func ErrorDetail_FromProto(mapCtx *direct.MapContext, in *pb.ErrorDetail) *krm.ErrorDetail {
	if in == nil {
		return nil
	}
	out := &krm.ErrorDetail{}
	out.Location = ErrorLocation_FromProto(mapCtx, in.GetLocation())
	out.ErrorInfo = ErrorInfo_FromProto(mapCtx, in.GetErrorInfo())
	return out
}
func ErrorDetail_ToProto(mapCtx *direct.MapContext, in *krm.ErrorDetail) *pb.ErrorDetail {
	if in == nil {
		return nil
	}
	out := &pb.ErrorDetail{}
	out.Location = ErrorLocation_ToProto(mapCtx, in.Location)
	out.ErrorInfo = ErrorInfo_ToProto(mapCtx, in.ErrorInfo)
	return out
}
func ErrorLocation_FromProto(mapCtx *direct.MapContext, in *pb.ErrorLocation) *krm.ErrorLocation {
	if in == nil {
		return nil
	}
	out := &krm.ErrorLocation{}
	out.Line = direct.LazyPtr(in.GetLine())
	out.Column = direct.LazyPtr(in.GetColumn())
	return out
}
func ErrorLocation_ToProto(mapCtx *direct.MapContext, in *krm.ErrorLocation) *pb.ErrorLocation {
	if in == nil {
		return nil
	}
	out := &pb.ErrorLocation{}
	out.Line = direct.ValueOf(in.Line)
	out.Column = direct.ValueOf(in.Column)
	return out
}
func MigrationSubtask_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSubtask) *krm.MigrationSubtask {
	if in == nil {
		return nil
	}
	out := &krm.MigrationSubtask{}
	// MISSING: Name
	out.TaskID = direct.LazyPtr(in.GetTaskId())
	out.Type = direct.LazyPtr(in.GetType())
	// MISSING: State
	// MISSING: ProcessingError
	// MISSING: ResourceErrorDetails
	out.ResourceErrorCount = direct.LazyPtr(in.GetResourceErrorCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, TimeSeries_FromProto)
	return out
}
func MigrationSubtask_ToProto(mapCtx *direct.MapContext, in *krm.MigrationSubtask) *pb.MigrationSubtask {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSubtask{}
	// MISSING: Name
	out.TaskId = direct.ValueOf(in.TaskID)
	out.Type = direct.ValueOf(in.Type)
	// MISSING: State
	// MISSING: ProcessingError
	// MISSING: ResourceErrorDetails
	out.ResourceErrorCount = direct.ValueOf(in.ResourceErrorCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, TimeSeries_ToProto)
	return out
}
func MigrationSubtaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSubtask) *krm.MigrationSubtaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationSubtaskObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: TaskID
	// MISSING: Type
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ProcessingError = ErrorInfo_FromProto(mapCtx, in.GetProcessingError())
	out.ResourceErrorDetails = direct.Slice_FromProto(mapCtx, in.ResourceErrorDetails, ResourceErrorDetail_FromProto)
	// MISSING: ResourceErrorCount
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	// MISSING: Metrics
	return out
}
func MigrationSubtaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationSubtaskObservedState) *pb.MigrationSubtask {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSubtask{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: TaskID
	// MISSING: Type
	out.State = direct.Enum_ToProto[pb.MigrationSubtask_State](mapCtx, in.State)
	out.ProcessingError = ErrorInfo_ToProto(mapCtx, in.ProcessingError)
	out.ResourceErrorDetails = direct.Slice_ToProto(mapCtx, in.ResourceErrorDetails, ResourceErrorDetail_ToProto)
	// MISSING: ResourceErrorCount
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	// MISSING: Metrics
	return out
}
func Point_FromProto(mapCtx *direct.MapContext, in *pb.Point) *krm.Point {
	if in == nil {
		return nil
	}
	out := &krm.Point{}
	out.Interval = TimeInterval_FromProto(mapCtx, in.GetInterval())
	out.Value = TypedValue_FromProto(mapCtx, in.GetValue())
	return out
}
func Point_ToProto(mapCtx *direct.MapContext, in *krm.Point) *pb.Point {
	if in == nil {
		return nil
	}
	out := &pb.Point{}
	out.Interval = TimeInterval_ToProto(mapCtx, in.Interval)
	out.Value = TypedValue_ToProto(mapCtx, in.Value)
	return out
}
func ResourceErrorDetail_FromProto(mapCtx *direct.MapContext, in *pb.ResourceErrorDetail) *krm.ResourceErrorDetail {
	if in == nil {
		return nil
	}
	out := &krm.ResourceErrorDetail{}
	out.ResourceInfo = ResourceInfo_FromProto(mapCtx, in.GetResourceInfo())
	out.ErrorDetails = direct.Slice_FromProto(mapCtx, in.ErrorDetails, ErrorDetail_FromProto)
	out.ErrorCount = direct.LazyPtr(in.GetErrorCount())
	return out
}
func ResourceErrorDetail_ToProto(mapCtx *direct.MapContext, in *krm.ResourceErrorDetail) *pb.ResourceErrorDetail {
	if in == nil {
		return nil
	}
	out := &pb.ResourceErrorDetail{}
	out.ResourceInfo = ResourceInfo_ToProto(mapCtx, in.ResourceInfo)
	out.ErrorDetails = direct.Slice_ToProto(mapCtx, in.ErrorDetails, ErrorDetail_ToProto)
	out.ErrorCount = direct.ValueOf(in.ErrorCount)
	return out
}
func TimeInterval_FromProto(mapCtx *direct.MapContext, in *pb.TimeInterval) *krm.TimeInterval {
	if in == nil {
		return nil
	}
	out := &krm.TimeInterval{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func TimeInterval_ToProto(mapCtx *direct.MapContext, in *krm.TimeInterval) *pb.TimeInterval {
	if in == nil {
		return nil
	}
	out := &pb.TimeInterval{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func TimeSeries_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeries) *krm.TimeSeries {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeries{}
	out.Metric = direct.LazyPtr(in.GetMetric())
	out.ValueType = direct.Enum_FromProto(mapCtx, in.GetValueType())
	out.MetricKind = direct.Enum_FromProto(mapCtx, in.GetMetricKind())
	out.Points = direct.Slice_FromProto(mapCtx, in.Points, Point_FromProto)
	return out
}
func TimeSeries_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeries) *pb.TimeSeries {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeries{}
	out.Metric = direct.ValueOf(in.Metric)
	out.ValueType = direct.Enum_ToProto[pb.MetricDescriptor_ValueType](mapCtx, in.ValueType)
	out.MetricKind = direct.Enum_ToProto[pb.MetricDescriptor_MetricKind](mapCtx, in.MetricKind)
	out.Points = direct.Slice_ToProto(mapCtx, in.Points, Point_ToProto)
	return out
}
func TypedValue_FromProto(mapCtx *direct.MapContext, in *pb.TypedValue) *krm.TypedValue {
	if in == nil {
		return nil
	}
	out := &krm.TypedValue{}
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.Int64Value = direct.LazyPtr(in.GetInt64Value())
	out.DoubleValue = direct.LazyPtr(in.GetDoubleValue())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.DistributionValue = Distribution_FromProto(mapCtx, in.GetDistributionValue())
	return out
}
func TypedValue_ToProto(mapCtx *direct.MapContext, in *krm.TypedValue) *pb.TypedValue {
	if in == nil {
		return nil
	}
	out := &pb.TypedValue{}
	if oneof := TypedValue_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := TypedValue_Int64Value_ToProto(mapCtx, in.Int64Value); oneof != nil {
		out.Value = oneof
	}
	if oneof := TypedValue_DoubleValue_ToProto(mapCtx, in.DoubleValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := TypedValue_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := Distribution_ToProto(mapCtx, in.DistributionValue); oneof != nil {
		out.Value = &pb.TypedValue_DistributionValue{DistributionValue: oneof}
	}
	return out
}
