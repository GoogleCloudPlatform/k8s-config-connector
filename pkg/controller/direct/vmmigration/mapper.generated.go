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

package vmmigration

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmmigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CycleStep_FromProto(mapCtx *direct.MapContext, in *pb.CycleStep) *krm.CycleStep {
	if in == nil {
		return nil
	}
	out := &krm.CycleStep{}
	out.InitializingReplication = InitializingReplicationStep_FromProto(mapCtx, in.GetInitializingReplication())
	out.Replicating = ReplicatingStep_FromProto(mapCtx, in.GetReplicating())
	out.PostProcessing = PostProcessingStep_FromProto(mapCtx, in.GetPostProcessing())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func CycleStep_ToProto(mapCtx *direct.MapContext, in *krm.CycleStep) *pb.CycleStep {
	if in == nil {
		return nil
	}
	out := &pb.CycleStep{}
	if oneof := InitializingReplicationStep_ToProto(mapCtx, in.InitializingReplication); oneof != nil {
		out.Step = &pb.CycleStep_InitializingReplication{InitializingReplication: oneof}
	}
	if oneof := ReplicatingStep_ToProto(mapCtx, in.Replicating); oneof != nil {
		out.Step = &pb.CycleStep_Replicating{Replicating: oneof}
	}
	if oneof := PostProcessingStep_ToProto(mapCtx, in.PostProcessing); oneof != nil {
		out.Step = &pb.CycleStep_PostProcessing{PostProcessing: oneof}
	}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func InitializingReplicationStep_FromProto(mapCtx *direct.MapContext, in *pb.InitializingReplicationStep) *krm.InitializingReplicationStep {
	if in == nil {
		return nil
	}
	out := &krm.InitializingReplicationStep{}
	return out
}
func InitializingReplicationStep_ToProto(mapCtx *direct.MapContext, in *krm.InitializingReplicationStep) *pb.InitializingReplicationStep {
	if in == nil {
		return nil
	}
	out := &pb.InitializingReplicationStep{}
	return out
}
func PostProcessingStep_FromProto(mapCtx *direct.MapContext, in *pb.PostProcessingStep) *krm.PostProcessingStep {
	if in == nil {
		return nil
	}
	out := &krm.PostProcessingStep{}
	return out
}
func PostProcessingStep_ToProto(mapCtx *direct.MapContext, in *krm.PostProcessingStep) *pb.PostProcessingStep {
	if in == nil {
		return nil
	}
	out := &pb.PostProcessingStep{}
	return out
}
func ReplicatingStep_FromProto(mapCtx *direct.MapContext, in *pb.ReplicatingStep) *krm.ReplicatingStep {
	if in == nil {
		return nil
	}
	out := &krm.ReplicatingStep{}
	out.TotalBytes = direct.LazyPtr(in.GetTotalBytes())
	out.ReplicatedBytes = direct.LazyPtr(in.GetReplicatedBytes())
	out.LastTwoMinutesAverageBytesPerSecond = direct.LazyPtr(in.GetLastTwoMinutesAverageBytesPerSecond())
	out.LastThirtyMinutesAverageBytesPerSecond = direct.LazyPtr(in.GetLastThirtyMinutesAverageBytesPerSecond())
	return out
}
func ReplicatingStep_ToProto(mapCtx *direct.MapContext, in *krm.ReplicatingStep) *pb.ReplicatingStep {
	if in == nil {
		return nil
	}
	out := &pb.ReplicatingStep{}
	out.TotalBytes = direct.ValueOf(in.TotalBytes)
	out.ReplicatedBytes = direct.ValueOf(in.ReplicatedBytes)
	out.LastTwoMinutesAverageBytesPerSecond = direct.ValueOf(in.LastTwoMinutesAverageBytesPerSecond)
	out.LastThirtyMinutesAverageBytesPerSecond = direct.ValueOf(in.LastThirtyMinutesAverageBytesPerSecond)
	return out
}
func ReplicationCycle_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationCycle) *krm.ReplicationCycle {
	if in == nil {
		return nil
	}
	out := &krm.ReplicationCycle{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CycleNumber = direct.LazyPtr(in.GetCycleNumber())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.TotalPauseDuration = direct.StringDuration_FromProto(mapCtx, in.GetTotalPauseDuration())
	out.ProgressPercent = direct.LazyPtr(in.GetProgressPercent())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, CycleStep_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func ReplicationCycle_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationCycle) *pb.ReplicationCycle {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationCycle{}
	out.Name = direct.ValueOf(in.Name)
	out.CycleNumber = direct.ValueOf(in.CycleNumber)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.TotalPauseDuration = direct.StringDuration_ToProto(mapCtx, in.TotalPauseDuration)
	out.ProgressPercent = direct.ValueOf(in.ProgressPercent)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, CycleStep_ToProto)
	out.State = direct.Enum_ToProto[pb.ReplicationCycle_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func VmmigrationReplicationCycleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationCycle) *krm.VmmigrationReplicationCycleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationReplicationCycleObservedState{}
	// MISSING: Name
	// MISSING: CycleNumber
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TotalPauseDuration
	// MISSING: ProgressPercent
	// MISSING: Steps
	// MISSING: State
	// MISSING: Error
	return out
}
func VmmigrationReplicationCycleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationReplicationCycleObservedState) *pb.ReplicationCycle {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationCycle{}
	// MISSING: Name
	// MISSING: CycleNumber
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TotalPauseDuration
	// MISSING: ProgressPercent
	// MISSING: Steps
	// MISSING: State
	// MISSING: Error
	return out
}
func VmmigrationReplicationCycleSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationCycle) *krm.VmmigrationReplicationCycleSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationReplicationCycleSpec{}
	// MISSING: Name
	// MISSING: CycleNumber
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TotalPauseDuration
	// MISSING: ProgressPercent
	// MISSING: Steps
	// MISSING: State
	// MISSING: Error
	return out
}
func VmmigrationReplicationCycleSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationReplicationCycleSpec) *pb.ReplicationCycle {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationCycle{}
	// MISSING: Name
	// MISSING: CycleNumber
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TotalPauseDuration
	// MISSING: ProgressPercent
	// MISSING: Steps
	// MISSING: State
	// MISSING: Error
	return out
}
