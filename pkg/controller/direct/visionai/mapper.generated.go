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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Process_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.Process {
	if in == nil {
		return nil
	}
	out := &krm.Process{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Analysis = direct.LazyPtr(in.GetAnalysis())
	out.AttributeOverrides = in.AttributeOverrides
	out.RunStatus = RunStatus_FromProto(mapCtx, in.GetRunStatus())
	out.RunMode = direct.Enum_FromProto(mapCtx, in.GetRunMode())
	out.EventID = direct.LazyPtr(in.GetEventId())
	out.BatchID = direct.LazyPtr(in.GetBatchId())
	out.RetryCount = direct.LazyPtr(in.GetRetryCount())
	return out
}
func Process_ToProto(mapCtx *direct.MapContext, in *krm.Process) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Analysis = direct.ValueOf(in.Analysis)
	out.AttributeOverrides = in.AttributeOverrides
	out.RunStatus = RunStatus_ToProto(mapCtx, in.RunStatus)
	out.RunMode = direct.Enum_ToProto[pb.RunMode](mapCtx, in.RunMode)
	out.EventId = direct.ValueOf(in.EventID)
	out.BatchId = direct.ValueOf(in.BatchID)
	out.RetryCount = direct.ValueOf(in.RetryCount)
	return out
}
func ProcessObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.ProcessObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProcessObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Analysis
	// MISSING: AttributeOverrides
	// MISSING: RunStatus
	// MISSING: RunMode
	// MISSING: EventID
	// MISSING: BatchID
	// MISSING: RetryCount
	return out
}
func ProcessObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProcessObservedState) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Analysis
	// MISSING: AttributeOverrides
	// MISSING: RunStatus
	// MISSING: RunMode
	// MISSING: EventID
	// MISSING: BatchID
	// MISSING: RetryCount
	return out
}
func RunStatus_FromProto(mapCtx *direct.MapContext, in *pb.RunStatus) *krm.RunStatus {
	if in == nil {
		return nil
	}
	out := &krm.RunStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Reason = direct.LazyPtr(in.GetReason())
	return out
}
func RunStatus_ToProto(mapCtx *direct.MapContext, in *krm.RunStatus) *pb.RunStatus {
	if in == nil {
		return nil
	}
	out := &pb.RunStatus{}
	out.State = direct.Enum_ToProto[pb.RunStatus_State](mapCtx, in.State)
	out.Reason = direct.ValueOf(in.Reason)
	return out
}
func VisionaiProcessObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.VisionaiProcessObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiProcessObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Analysis
	// MISSING: AttributeOverrides
	// MISSING: RunStatus
	// MISSING: RunMode
	// MISSING: EventID
	// MISSING: BatchID
	// MISSING: RetryCount
	return out
}
func VisionaiProcessObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiProcessObservedState) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Analysis
	// MISSING: AttributeOverrides
	// MISSING: RunStatus
	// MISSING: RunMode
	// MISSING: EventID
	// MISSING: BatchID
	// MISSING: RetryCount
	return out
}
func VisionaiProcessSpec_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.VisionaiProcessSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiProcessSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Analysis
	// MISSING: AttributeOverrides
	// MISSING: RunStatus
	// MISSING: RunMode
	// MISSING: EventID
	// MISSING: BatchID
	// MISSING: RetryCount
	return out
}
func VisionaiProcessSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiProcessSpec) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Analysis
	// MISSING: AttributeOverrides
	// MISSING: RunStatus
	// MISSING: RunMode
	// MISSING: EventID
	// MISSING: BatchID
	// MISSING: RetryCount
	return out
}
