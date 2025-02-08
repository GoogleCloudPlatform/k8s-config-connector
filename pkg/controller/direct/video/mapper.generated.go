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

package video

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Event_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.Event {
	if in == nil {
		return nil
	}
	out := &krm.Event{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.InputSwitch = Event_InputSwitchTask_FromProto(mapCtx, in.GetInputSwitch())
	out.AdBreak = Event_AdBreakTask_FromProto(mapCtx, in.GetAdBreak())
	out.ReturnToProgram = Event_ReturnToProgramTask_FromProto(mapCtx, in.GetReturnToProgram())
	out.Slate = Event_SlateTask_FromProto(mapCtx, in.GetSlate())
	out.Mute = Event_MuteTask_FromProto(mapCtx, in.GetMute())
	out.Unmute = Event_UnmuteTask_FromProto(mapCtx, in.GetUnmute())
	out.ExecuteNow = direct.LazyPtr(in.GetExecuteNow())
	out.ExecutionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExecutionTime())
	// MISSING: State
	// MISSING: Error
	return out
}
func Event_ToProto(mapCtx *direct.MapContext, in *krm.Event) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	if oneof := Event_InputSwitchTask_ToProto(mapCtx, in.InputSwitch); oneof != nil {
		out.Task = &pb.Event_InputSwitch{InputSwitch: oneof}
	}
	if oneof := Event_AdBreakTask_ToProto(mapCtx, in.AdBreak); oneof != nil {
		out.Task = &pb.Event_AdBreak{AdBreak: oneof}
	}
	if oneof := Event_ReturnToProgramTask_ToProto(mapCtx, in.ReturnToProgram); oneof != nil {
		out.Task = &pb.Event_ReturnToProgram{ReturnToProgram: oneof}
	}
	if oneof := Event_SlateTask_ToProto(mapCtx, in.Slate); oneof != nil {
		out.Task = &pb.Event_Slate{Slate: oneof}
	}
	if oneof := Event_MuteTask_ToProto(mapCtx, in.Mute); oneof != nil {
		out.Task = &pb.Event_Mute{Mute: oneof}
	}
	if oneof := Event_UnmuteTask_ToProto(mapCtx, in.Unmute); oneof != nil {
		out.Task = &pb.Event_Unmute{Unmute: oneof}
	}
	out.ExecuteNow = direct.ValueOf(in.ExecuteNow)
	out.ExecutionTime = direct.StringTimestamp_ToProto(mapCtx, in.ExecutionTime)
	// MISSING: State
	// MISSING: Error
	return out
}
func EventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.EventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: InputSwitch
	// MISSING: AdBreak
	// MISSING: ReturnToProgram
	// MISSING: Slate
	// MISSING: Mute
	// MISSING: Unmute
	// MISSING: ExecuteNow
	// MISSING: ExecutionTime
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func EventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventObservedState) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: InputSwitch
	// MISSING: AdBreak
	// MISSING: ReturnToProgram
	// MISSING: Slate
	// MISSING: Mute
	// MISSING: Unmute
	// MISSING: ExecuteNow
	// MISSING: ExecutionTime
	out.State = direct.Enum_ToProto[pb.Event_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func Event_AdBreakTask_FromProto(mapCtx *direct.MapContext, in *pb.Event_AdBreakTask) *krm.Event_AdBreakTask {
	if in == nil {
		return nil
	}
	out := &krm.Event_AdBreakTask{}
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}
func Event_AdBreakTask_ToProto(mapCtx *direct.MapContext, in *krm.Event_AdBreakTask) *pb.Event_AdBreakTask {
	if in == nil {
		return nil
	}
	out := &pb.Event_AdBreakTask{}
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}
func Event_InputSwitchTask_FromProto(mapCtx *direct.MapContext, in *pb.Event_InputSwitchTask) *krm.Event_InputSwitchTask {
	if in == nil {
		return nil
	}
	out := &krm.Event_InputSwitchTask{}
	out.InputKey = direct.LazyPtr(in.GetInputKey())
	return out
}
func Event_InputSwitchTask_ToProto(mapCtx *direct.MapContext, in *krm.Event_InputSwitchTask) *pb.Event_InputSwitchTask {
	if in == nil {
		return nil
	}
	out := &pb.Event_InputSwitchTask{}
	out.InputKey = direct.ValueOf(in.InputKey)
	return out
}
func Event_MuteTask_FromProto(mapCtx *direct.MapContext, in *pb.Event_MuteTask) *krm.Event_MuteTask {
	if in == nil {
		return nil
	}
	out := &krm.Event_MuteTask{}
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}
func Event_MuteTask_ToProto(mapCtx *direct.MapContext, in *krm.Event_MuteTask) *pb.Event_MuteTask {
	if in == nil {
		return nil
	}
	out := &pb.Event_MuteTask{}
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}
func Event_ReturnToProgramTask_FromProto(mapCtx *direct.MapContext, in *pb.Event_ReturnToProgramTask) *krm.Event_ReturnToProgramTask {
	if in == nil {
		return nil
	}
	out := &krm.Event_ReturnToProgramTask{}
	return out
}
func Event_ReturnToProgramTask_ToProto(mapCtx *direct.MapContext, in *krm.Event_ReturnToProgramTask) *pb.Event_ReturnToProgramTask {
	if in == nil {
		return nil
	}
	out := &pb.Event_ReturnToProgramTask{}
	return out
}
func Event_SlateTask_FromProto(mapCtx *direct.MapContext, in *pb.Event_SlateTask) *krm.Event_SlateTask {
	if in == nil {
		return nil
	}
	out := &krm.Event_SlateTask{}
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.Asset = direct.LazyPtr(in.GetAsset())
	return out
}
func Event_SlateTask_ToProto(mapCtx *direct.MapContext, in *krm.Event_SlateTask) *pb.Event_SlateTask {
	if in == nil {
		return nil
	}
	out := &pb.Event_SlateTask{}
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.Asset = direct.ValueOf(in.Asset)
	return out
}
func Event_UnmuteTask_FromProto(mapCtx *direct.MapContext, in *pb.Event_UnmuteTask) *krm.Event_UnmuteTask {
	if in == nil {
		return nil
	}
	out := &krm.Event_UnmuteTask{}
	return out
}
func Event_UnmuteTask_ToProto(mapCtx *direct.MapContext, in *krm.Event_UnmuteTask) *pb.Event_UnmuteTask {
	if in == nil {
		return nil
	}
	out := &pb.Event_UnmuteTask{}
	return out
}
func VideoEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.VideoEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoEventObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputSwitch
	// MISSING: AdBreak
	// MISSING: ReturnToProgram
	// MISSING: Slate
	// MISSING: Mute
	// MISSING: Unmute
	// MISSING: ExecuteNow
	// MISSING: ExecutionTime
	// MISSING: State
	// MISSING: Error
	return out
}
func VideoEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoEventObservedState) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputSwitch
	// MISSING: AdBreak
	// MISSING: ReturnToProgram
	// MISSING: Slate
	// MISSING: Mute
	// MISSING: Unmute
	// MISSING: ExecuteNow
	// MISSING: ExecutionTime
	// MISSING: State
	// MISSING: Error
	return out
}
func VideoEventSpec_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.VideoEventSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoEventSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputSwitch
	// MISSING: AdBreak
	// MISSING: ReturnToProgram
	// MISSING: Slate
	// MISSING: Mute
	// MISSING: Unmute
	// MISSING: ExecuteNow
	// MISSING: ExecutionTime
	// MISSING: State
	// MISSING: Error
	return out
}
func VideoEventSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoEventSpec) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputSwitch
	// MISSING: AdBreak
	// MISSING: ReturnToProgram
	// MISSING: Slate
	// MISSING: Mute
	// MISSING: Unmute
	// MISSING: ExecuteNow
	// MISSING: ExecutionTime
	// MISSING: State
	// MISSING: Error
	return out
}
