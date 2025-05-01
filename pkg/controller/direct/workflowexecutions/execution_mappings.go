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

// krm.version: v1alpha1
// proto.service: google.cloud.workflows.executions.v1
// krm.group: workflowexecutions.cnrm.cloud.google.com

package workflowexecutions

import (
	pb "cloud.google.com/go/workflows/executions/apiv1/executionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflowexecutions/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkflowsExecutionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.WorkflowsExecutionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsExecutionObservedState{}
	out.Name = direct.StringValue_FromProto(mapCtx, in.GetName())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Result = direct.LazyPtr(in.GetResult())
	out.Error = Execution_Error_FromProto(mapCtx, in.GetError())
	out.WorkflowRevisionID = direct.LazyPtr(in.GetWorkflowRevisionId())
	out.Status = Execution_Status_FromProto(mapCtx, in.GetStatus())
	out.StateError = Execution_StateError_FromProto(mapCtx, in.GetStateError())
	return out
}
func WorkflowsExecutionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsExecutionObservedState) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	out.Name = direct.StringValue_FromProto(mapCtx, in.GetName())
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.State = direct.Enum_ToProto[pb.Execution_State](mapCtx, in.State)
	out.Result = direct.ValueOf(in.Result)
	out.Error = Execution_Error_ToProto(mapCtx, in.Error)
	out.WorkflowRevisionId = direct.ValueOf(in.WorkflowRevisionID)
	out.Status = Execution_Status_ToProto(mapCtx, in.Status)
	out.StateError = Execution_StateError_ToProto(mapCtx, in.StateError)
	return out
}
func WorkflowsExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.WorkflowsExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsExecutionSpec{}
	out.Argument = direct.LazyPtr(in.GetArgument())
	out.CallLogLevel = direct.Enum_FromProto(mapCtx, in.GetCallLogLevel())
	out.Labels = in.Labels
	return out
}
func WorkflowsExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsExecutionSpec) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	// in.Parent
	// TODO Use in.Parent to get Name
	out.Argument = direct.ValueOf(in.Argument)
	out.CallLogLevel = direct.Enum_ToProto[pb.Execution_CallLogLevel](mapCtx, in.CallLogLevel)
	out.Labels = in.Labels
	return out
}
