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

package dataplex

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DataplexJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.DataplexJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexJobObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func DataplexJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func DataplexJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.DataplexJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexJobSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func DataplexJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: RetryCount
	// MISSING: Service
	// MISSING: ServiceJob
	// MISSING: Message
	// MISSING: Labels
	// MISSING: Trigger
	// MISSING: ExecutionSpec
	return out
}
func JobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.JobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RetryCount = direct.LazyPtr(in.GetRetryCount())
	out.Service = direct.Enum_FromProto(mapCtx, in.GetService())
	out.ServiceJob = direct.LazyPtr(in.GetServiceJob())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.Labels = in.Labels
	out.Trigger = direct.Enum_FromProto(mapCtx, in.GetTrigger())
	out.ExecutionSpec = Task_ExecutionSpec_FromProto(mapCtx, in.GetExecutionSpec())
	return out
}
func JobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Job_State](mapCtx, in.State)
	out.RetryCount = direct.ValueOf(in.RetryCount)
	out.Service = direct.Enum_ToProto[pb.Job_Service](mapCtx, in.Service)
	out.ServiceJob = direct.ValueOf(in.ServiceJob)
	out.Message = direct.ValueOf(in.Message)
	out.Labels = in.Labels
	out.Trigger = direct.Enum_ToProto[pb.Job_Trigger](mapCtx, in.Trigger)
	out.ExecutionSpec = Task_ExecutionSpec_ToProto(mapCtx, in.ExecutionSpec)
	return out
}
