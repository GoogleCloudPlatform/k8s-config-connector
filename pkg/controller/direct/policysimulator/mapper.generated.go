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

package policysimulator

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/policysimulator/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/policysimulator/apiv1/policysimulatorpb"
)
func PolicysimulatorReplayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Replay) *krm.PolicysimulatorReplayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicysimulatorReplayObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Config
	// MISSING: ResultsSummary
	return out
}
func PolicysimulatorReplayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicysimulatorReplayObservedState) *pb.Replay {
	if in == nil {
		return nil
	}
	out := &pb.Replay{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Config
	// MISSING: ResultsSummary
	return out
}
func PolicysimulatorReplaySpec_FromProto(mapCtx *direct.MapContext, in *pb.Replay) *krm.PolicysimulatorReplaySpec {
	if in == nil {
		return nil
	}
	out := &krm.PolicysimulatorReplaySpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Config
	// MISSING: ResultsSummary
	return out
}
func PolicysimulatorReplaySpec_ToProto(mapCtx *direct.MapContext, in *krm.PolicysimulatorReplaySpec) *pb.Replay {
	if in == nil {
		return nil
	}
	out := &pb.Replay{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Config
	// MISSING: ResultsSummary
	return out
}
func Replay_FromProto(mapCtx *direct.MapContext, in *pb.Replay) *krm.Replay {
	if in == nil {
		return nil
	}
	out := &krm.Replay{}
	// MISSING: Name
	// MISSING: State
	out.Config = ReplayConfig_FromProto(mapCtx, in.GetConfig())
	// MISSING: ResultsSummary
	return out
}
func Replay_ToProto(mapCtx *direct.MapContext, in *krm.Replay) *pb.Replay {
	if in == nil {
		return nil
	}
	out := &pb.Replay{}
	// MISSING: Name
	// MISSING: State
	out.Config = ReplayConfig_ToProto(mapCtx, in.Config)
	// MISSING: ResultsSummary
	return out
}
func ReplayConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReplayConfig) *krm.ReplayConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReplayConfig{}
	// MISSING: PolicyOverlay
	out.LogSource = direct.Enum_FromProto(mapCtx, in.GetLogSource())
	return out
}
func ReplayConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReplayConfig) *pb.ReplayConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReplayConfig{}
	// MISSING: PolicyOverlay
	out.LogSource = direct.Enum_ToProto[pb.ReplayConfig_LogSource](mapCtx, in.LogSource)
	return out
}
func ReplayObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Replay) *krm.ReplayObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReplayObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Config
	out.ResultsSummary = Replay_ResultsSummary_FromProto(mapCtx, in.GetResultsSummary())
	return out
}
func ReplayObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReplayObservedState) *pb.Replay {
	if in == nil {
		return nil
	}
	out := &pb.Replay{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Replay_State](mapCtx, in.State)
	// MISSING: Config
	out.ResultsSummary = Replay_ResultsSummary_ToProto(mapCtx, in.ResultsSummary)
	return out
}
func Replay_ResultsSummary_FromProto(mapCtx *direct.MapContext, in *pb.Replay_ResultsSummary) *krm.Replay_ResultsSummary {
	if in == nil {
		return nil
	}
	out := &krm.Replay_ResultsSummary{}
	out.LogCount = direct.LazyPtr(in.GetLogCount())
	out.UnchangedCount = direct.LazyPtr(in.GetUnchangedCount())
	out.DifferenceCount = direct.LazyPtr(in.GetDifferenceCount())
	out.ErrorCount = direct.LazyPtr(in.GetErrorCount())
	out.OldestDate = Date_FromProto(mapCtx, in.GetOldestDate())
	out.NewestDate = Date_FromProto(mapCtx, in.GetNewestDate())
	return out
}
func Replay_ResultsSummary_ToProto(mapCtx *direct.MapContext, in *krm.Replay_ResultsSummary) *pb.Replay_ResultsSummary {
	if in == nil {
		return nil
	}
	out := &pb.Replay_ResultsSummary{}
	out.LogCount = direct.ValueOf(in.LogCount)
	out.UnchangedCount = direct.ValueOf(in.UnchangedCount)
	out.DifferenceCount = direct.ValueOf(in.DifferenceCount)
	out.ErrorCount = direct.ValueOf(in.ErrorCount)
	out.OldestDate = Date_ToProto(mapCtx, in.OldestDate)
	out.NewestDate = Date_ToProto(mapCtx, in.NewestDate)
	return out
}
