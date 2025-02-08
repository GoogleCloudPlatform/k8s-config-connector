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
	pb "cloud.google.com/go/workflows/executions/apiv1beta/executionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Execution_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.Execution {
	if in == nil {
		return nil
	}
	out := &krm.Execution{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	out.Argument = direct.LazyPtr(in.GetArgument())
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
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
	// MISSING: State
	out.Argument = direct.ValueOf(in.Argument)
	// MISSING: Result
	// MISSING: Error
	// MISSING: WorkflowRevisionID
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
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Argument
	out.Result = direct.LazyPtr(in.GetResult())
	out.Error = Execution_Error_FromProto(mapCtx, in.GetError())
	out.WorkflowRevisionID = direct.LazyPtr(in.GetWorkflowRevisionId())
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
	out.State = direct.Enum_ToProto[pb.Execution_State](mapCtx, in.State)
	// MISSING: Argument
	out.Result = direct.ValueOf(in.Result)
	out.Error = Execution_Error_ToProto(mapCtx, in.Error)
	out.WorkflowRevisionId = direct.ValueOf(in.WorkflowRevisionID)
	return out
}
func Execution_Error_FromProto(mapCtx *direct.MapContext, in *pb.Execution_Error) *krm.Execution_Error {
	if in == nil {
		return nil
	}
	out := &krm.Execution_Error{}
	out.Payload = direct.LazyPtr(in.GetPayload())
	out.Context = direct.LazyPtr(in.GetContext())
	return out
}
func Execution_Error_ToProto(mapCtx *direct.MapContext, in *krm.Execution_Error) *pb.Execution_Error {
	if in == nil {
		return nil
	}
	out := &pb.Execution_Error{}
	out.Payload = direct.ValueOf(in.Payload)
	out.Context = direct.ValueOf(in.Context)
	return out
}
