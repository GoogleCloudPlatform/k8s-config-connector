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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Deployment_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.Deployment {
	if in == nil {
		return nil
	}
	out := &krm.Deployment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.FlowVersion = direct.LazyPtr(in.GetFlowVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Result = Deployment_Result_FromProto(mapCtx, in.GetResult())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func Deployment_ToProto(mapCtx *direct.MapContext, in *krm.Deployment) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	out.Name = direct.ValueOf(in.Name)
	out.FlowVersion = direct.ValueOf(in.FlowVersion)
	out.State = direct.Enum_ToProto[pb.Deployment_State](mapCtx, in.State)
	out.Result = Deployment_Result_ToProto(mapCtx, in.Result)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func Deployment_Result_FromProto(mapCtx *direct.MapContext, in *pb.Deployment_Result) *krm.Deployment_Result {
	if in == nil {
		return nil
	}
	out := &krm.Deployment_Result{}
	out.DeploymentTestResults = in.DeploymentTestResults
	out.Experiment = direct.LazyPtr(in.GetExperiment())
	return out
}
func Deployment_Result_ToProto(mapCtx *direct.MapContext, in *krm.Deployment_Result) *pb.Deployment_Result {
	if in == nil {
		return nil
	}
	out := &pb.Deployment_Result{}
	out.DeploymentTestResults = in.DeploymentTestResults
	out.Experiment = direct.ValueOf(in.Experiment)
	return out
}
func DialogflowDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.DialogflowDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowDeploymentObservedState{}
	// MISSING: Name
	// MISSING: FlowVersion
	// MISSING: State
	// MISSING: Result
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DialogflowDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowDeploymentObservedState) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: FlowVersion
	// MISSING: State
	// MISSING: Result
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DialogflowDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.DialogflowDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowDeploymentSpec{}
	// MISSING: Name
	// MISSING: FlowVersion
	// MISSING: State
	// MISSING: Result
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DialogflowDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: FlowVersion
	// MISSING: State
	// MISSING: Result
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
