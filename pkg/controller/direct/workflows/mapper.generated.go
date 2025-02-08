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

package workflows

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/workflows/executions/apiv1/executionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Execution_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.Execution {
	if in == nil {
		return nil
	}
	out := &krm.Execution{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: Duration
	// MISSING: State
	out.Argument = direct.LazyPtr(in.GetArgument())
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
	out.CallLogLevel = direct.Enum_FromProto(mapCtx, in.GetCallLogLevel())
	// MISSING: Status
	out.Labels = in.Labels
	// MISSING: StateError
	return out
}
func Execution_ToProto(mapCtx *direct.MapContext, in *krm.Execution) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: Duration
	// MISSING: State
	out.Argument = direct.ValueOf(in.Argument)
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
	out.CallLogLevel = direct.Enum_ToProto[pb.Execution_CallLogLevel](mapCtx, in.CallLogLevel)
	// MISSING: Status
	out.Labels = in.Labels
	// MISSING: StateError
	return out
}
func ExecutionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.ExecutionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Argument
	out.Result = direct.LazyPtr(in.GetResult())
	out.Error = Execution_Error_FromProto(mapCtx, in.GetError())
	out.WorkflowRevisionID = direct.LazyPtr(in.GetWorkflowRevisionId())
	// MISSING: CallLogLevel
	out.Status = Execution_Status_FromProto(mapCtx, in.GetStatus())
	// MISSING: Labels
	out.StateError = Execution_StateError_FromProto(mapCtx, in.GetStateError())
	return out
}
func ExecutionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionObservedState) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	out.Name = direct.ValueOf(in.Name)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.State = direct.Enum_ToProto[pb.Execution_State](mapCtx, in.State)
	// MISSING: Argument
	out.Result = direct.ValueOf(in.Result)
	out.Error = Execution_Error_ToProto(mapCtx, in.Error)
	out.WorkflowRevisionId = direct.ValueOf(in.WorkflowRevisionID)
	// MISSING: CallLogLevel
	out.Status = Execution_Status_ToProto(mapCtx, in.Status)
	// MISSING: Labels
	out.StateError = Execution_StateError_ToProto(mapCtx, in.StateError)
	return out
}
func Execution_Error_FromProto(mapCtx *direct.MapContext, in *pb.Execution_Error) *krm.Execution_Error {
	if in == nil {
		return nil
	}
	out := &krm.Execution_Error{}
	out.Payload = direct.LazyPtr(in.GetPayload())
	out.Context = direct.LazyPtr(in.GetContext())
	out.StackTrace = Execution_StackTrace_FromProto(mapCtx, in.GetStackTrace())
	return out
}
func Execution_Error_ToProto(mapCtx *direct.MapContext, in *krm.Execution_Error) *pb.Execution_Error {
	if in == nil {
		return nil
	}
	out := &pb.Execution_Error{}
	out.Payload = direct.ValueOf(in.Payload)
	out.Context = direct.ValueOf(in.Context)
	out.StackTrace = Execution_StackTrace_ToProto(mapCtx, in.StackTrace)
	return out
}
func Execution_StackTrace_FromProto(mapCtx *direct.MapContext, in *pb.Execution_StackTrace) *krm.Execution_StackTrace {
	if in == nil {
		return nil
	}
	out := &krm.Execution_StackTrace{}
	out.Elements = direct.Slice_FromProto(mapCtx, in.Elements, Execution_StackTraceElement_FromProto)
	return out
}
func Execution_StackTrace_ToProto(mapCtx *direct.MapContext, in *krm.Execution_StackTrace) *pb.Execution_StackTrace {
	if in == nil {
		return nil
	}
	out := &pb.Execution_StackTrace{}
	out.Elements = direct.Slice_ToProto(mapCtx, in.Elements, Execution_StackTraceElement_ToProto)
	return out
}
func Execution_StackTraceElement_FromProto(mapCtx *direct.MapContext, in *pb.Execution_StackTraceElement) *krm.Execution_StackTraceElement {
	if in == nil {
		return nil
	}
	out := &krm.Execution_StackTraceElement{}
	out.Step = direct.LazyPtr(in.GetStep())
	out.Routine = direct.LazyPtr(in.GetRoutine())
	out.Position = Execution_StackTraceElement_Position_FromProto(mapCtx, in.GetPosition())
	return out
}
func Execution_StackTraceElement_ToProto(mapCtx *direct.MapContext, in *krm.Execution_StackTraceElement) *pb.Execution_StackTraceElement {
	if in == nil {
		return nil
	}
	out := &pb.Execution_StackTraceElement{}
	out.Step = direct.ValueOf(in.Step)
	out.Routine = direct.ValueOf(in.Routine)
	out.Position = Execution_StackTraceElement_Position_ToProto(mapCtx, in.Position)
	return out
}
func Execution_StackTraceElement_Position_FromProto(mapCtx *direct.MapContext, in *pb.Execution_StackTraceElement_Position) *krm.Execution_StackTraceElement_Position {
	if in == nil {
		return nil
	}
	out := &krm.Execution_StackTraceElement_Position{}
	out.Line = direct.LazyPtr(in.GetLine())
	out.Column = direct.LazyPtr(in.GetColumn())
	out.Length = direct.LazyPtr(in.GetLength())
	return out
}
func Execution_StackTraceElement_Position_ToProto(mapCtx *direct.MapContext, in *krm.Execution_StackTraceElement_Position) *pb.Execution_StackTraceElement_Position {
	if in == nil {
		return nil
	}
	out := &pb.Execution_StackTraceElement_Position{}
	out.Line = direct.ValueOf(in.Line)
	out.Column = direct.ValueOf(in.Column)
	out.Length = direct.ValueOf(in.Length)
	return out
}
func Execution_StateError_FromProto(mapCtx *direct.MapContext, in *pb.Execution_StateError) *krm.Execution_StateError {
	if in == nil {
		return nil
	}
	out := &krm.Execution_StateError{}
	out.Details = direct.LazyPtr(in.GetDetails())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Execution_StateError_ToProto(mapCtx *direct.MapContext, in *krm.Execution_StateError) *pb.Execution_StateError {
	if in == nil {
		return nil
	}
	out := &pb.Execution_StateError{}
	out.Details = direct.ValueOf(in.Details)
	out.Type = direct.Enum_ToProto[pb.Execution_StateError_Type](mapCtx, in.Type)
	return out
}
func Execution_Status_FromProto(mapCtx *direct.MapContext, in *pb.Execution_Status) *krm.Execution_Status {
	if in == nil {
		return nil
	}
	out := &krm.Execution_Status{}
	out.CurrentSteps = direct.Slice_FromProto(mapCtx, in.CurrentSteps, Execution_Status_Step_FromProto)
	return out
}
func Execution_Status_ToProto(mapCtx *direct.MapContext, in *krm.Execution_Status) *pb.Execution_Status {
	if in == nil {
		return nil
	}
	out := &pb.Execution_Status{}
	out.CurrentSteps = direct.Slice_ToProto(mapCtx, in.CurrentSteps, Execution_Status_Step_ToProto)
	return out
}
func Execution_Status_Step_FromProto(mapCtx *direct.MapContext, in *pb.Execution_Status_Step) *krm.Execution_Status_Step {
	if in == nil {
		return nil
	}
	out := &krm.Execution_Status_Step{}
	out.Routine = direct.LazyPtr(in.GetRoutine())
	out.Step = direct.LazyPtr(in.GetStep())
	return out
}
func Execution_Status_Step_ToProto(mapCtx *direct.MapContext, in *krm.Execution_Status_Step) *pb.Execution_Status_Step {
	if in == nil {
		return nil
	}
	out := &pb.Execution_Status_Step{}
	out.Routine = direct.ValueOf(in.Routine)
	out.Step = direct.ValueOf(in.Step)
	return out
}
func WorkflowsExecutionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.WorkflowsExecutionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsExecutionObservedState{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: Duration
	// MISSING: State
	// MISSING: Argument
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
	// MISSING: CallLogLevel
	// MISSING: Status
	// MISSING: Labels
	// MISSING: StateError
	return out
}
func WorkflowsExecutionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsExecutionObservedState) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: Duration
	// MISSING: State
	// MISSING: Argument
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
	// MISSING: CallLogLevel
	// MISSING: Status
	// MISSING: Labels
	// MISSING: StateError
	return out
}
func WorkflowsExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.WorkflowsExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsExecutionSpec{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: Duration
	// MISSING: State
	// MISSING: Argument
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
	// MISSING: CallLogLevel
	// MISSING: Status
	// MISSING: Labels
	// MISSING: StateError
	return out
}
func WorkflowsExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsExecutionSpec) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: Duration
	// MISSING: State
	// MISSING: Argument
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
	// MISSING: CallLogLevel
	// MISSING: Status
	// MISSING: Labels
	// MISSING: StateError
	return out
}
