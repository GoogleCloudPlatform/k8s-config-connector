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
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Action_FromProto(mapCtx *direct.MapContext, in *pb.Action) *krm.Action {
	if in == nil {
		return nil
	}
	out := &krm.Action{}
	out.UserUtterance = UserUtterance_FromProto(mapCtx, in.GetUserUtterance())
	out.AgentUtterance = AgentUtterance_FromProto(mapCtx, in.GetAgentUtterance())
	out.ToolUse = ToolUse_FromProto(mapCtx, in.GetToolUse())
	out.PlaybookInvocation = PlaybookInvocation_FromProto(mapCtx, in.GetPlaybookInvocation())
	out.FlowInvocation = FlowInvocation_FromProto(mapCtx, in.GetFlowInvocation())
	return out
}
func Action_ToProto(mapCtx *direct.MapContext, in *krm.Action) *pb.Action {
	if in == nil {
		return nil
	}
	out := &pb.Action{}
	if oneof := UserUtterance_ToProto(mapCtx, in.UserUtterance); oneof != nil {
		out.Action = &pb.Action_UserUtterance{UserUtterance: oneof}
	}
	if oneof := AgentUtterance_ToProto(mapCtx, in.AgentUtterance); oneof != nil {
		out.Action = &pb.Action_AgentUtterance{AgentUtterance: oneof}
	}
	if oneof := ToolUse_ToProto(mapCtx, in.ToolUse); oneof != nil {
		out.Action = &pb.Action_ToolUse{ToolUse: oneof}
	}
	if oneof := PlaybookInvocation_ToProto(mapCtx, in.PlaybookInvocation); oneof != nil {
		out.Action = &pb.Action_PlaybookInvocation{PlaybookInvocation: oneof}
	}
	if oneof := FlowInvocation_ToProto(mapCtx, in.FlowInvocation); oneof != nil {
		out.Action = &pb.Action_FlowInvocation{FlowInvocation: oneof}
	}
	return out
}
func AgentUtterance_FromProto(mapCtx *direct.MapContext, in *pb.AgentUtterance) *krm.AgentUtterance {
	if in == nil {
		return nil
	}
	out := &krm.AgentUtterance{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func AgentUtterance_ToProto(mapCtx *direct.MapContext, in *krm.AgentUtterance) *pb.AgentUtterance {
	if in == nil {
		return nil
	}
	out := &pb.AgentUtterance{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func DialogflowExampleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.DialogflowExampleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowExampleObservedState{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func DialogflowExampleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowExampleObservedState) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func DialogflowExampleSpec_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.DialogflowExampleSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowExampleSpec{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func DialogflowExampleSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowExampleSpec) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func Example_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.Example {
	if in == nil {
		return nil
	}
	out := &krm.Example{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PlaybookInput = PlaybookInput_FromProto(mapCtx, in.GetPlaybookInput())
	out.PlaybookOutput = PlaybookOutput_FromProto(mapCtx, in.GetPlaybookOutput())
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, Action_FromProto)
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ConversationState = direct.Enum_FromProto(mapCtx, in.GetConversationState())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func Example_ToProto(mapCtx *direct.MapContext, in *krm.Example) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	out.Name = direct.ValueOf(in.Name)
	out.PlaybookInput = PlaybookInput_ToProto(mapCtx, in.PlaybookInput)
	out.PlaybookOutput = PlaybookOutput_ToProto(mapCtx, in.PlaybookOutput)
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, Action_ToProto)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ConversationState = direct.Enum_ToProto[pb.OutputState](mapCtx, in.ConversationState)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
func ExampleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.ExampleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExampleObservedState{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	out.TokenCount = direct.LazyPtr(in.GetTokenCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func ExampleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExampleObservedState) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	out.TokenCount = direct.ValueOf(in.TokenCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func FlowInvocation_FromProto(mapCtx *direct.MapContext, in *pb.FlowInvocation) *krm.FlowInvocation {
	if in == nil {
		return nil
	}
	out := &krm.FlowInvocation{}
	out.Flow = direct.LazyPtr(in.GetFlow())
	out.InputActionParameters = InputActionParameters_FromProto(mapCtx, in.GetInputActionParameters())
	out.OutputActionParameters = OutputActionParameters_FromProto(mapCtx, in.GetOutputActionParameters())
	out.FlowState = direct.Enum_FromProto(mapCtx, in.GetFlowState())
	return out
}
func FlowInvocation_ToProto(mapCtx *direct.MapContext, in *krm.FlowInvocation) *pb.FlowInvocation {
	if in == nil {
		return nil
	}
	out := &pb.FlowInvocation{}
	out.Flow = direct.ValueOf(in.Flow)
	out.InputActionParameters = InputActionParameters_ToProto(mapCtx, in.InputActionParameters)
	out.OutputActionParameters = OutputActionParameters_ToProto(mapCtx, in.OutputActionParameters)
	out.FlowState = direct.Enum_ToProto[pb.OutputState](mapCtx, in.FlowState)
	return out
}
func PlaybookInput_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookInput) *krm.PlaybookInput {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookInput{}
	out.PrecedingConversationSummary = direct.LazyPtr(in.GetPrecedingConversationSummary())
	out.ActionParameters = ActionParameters_FromProto(mapCtx, in.GetActionParameters())
	return out
}
func PlaybookInput_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookInput) *pb.PlaybookInput {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookInput{}
	out.PrecedingConversationSummary = direct.ValueOf(in.PrecedingConversationSummary)
	out.ActionParameters = ActionParameters_ToProto(mapCtx, in.ActionParameters)
	return out
}
func PlaybookInvocation_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookInvocation) *krm.PlaybookInvocation {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookInvocation{}
	out.Playbook = direct.LazyPtr(in.GetPlaybook())
	out.PlaybookInput = PlaybookInput_FromProto(mapCtx, in.GetPlaybookInput())
	out.PlaybookOutput = PlaybookOutput_FromProto(mapCtx, in.GetPlaybookOutput())
	out.PlaybookState = direct.Enum_FromProto(mapCtx, in.GetPlaybookState())
	return out
}
func PlaybookInvocation_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookInvocation) *pb.PlaybookInvocation {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookInvocation{}
	out.Playbook = direct.ValueOf(in.Playbook)
	out.PlaybookInput = PlaybookInput_ToProto(mapCtx, in.PlaybookInput)
	out.PlaybookOutput = PlaybookOutput_ToProto(mapCtx, in.PlaybookOutput)
	out.PlaybookState = direct.Enum_ToProto[pb.OutputState](mapCtx, in.PlaybookState)
	return out
}
func PlaybookOutput_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookOutput) *krm.PlaybookOutput {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookOutput{}
	out.ExecutionSummary = direct.LazyPtr(in.GetExecutionSummary())
	out.ActionParameters = ActionParameters_FromProto(mapCtx, in.GetActionParameters())
	return out
}
func PlaybookOutput_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookOutput) *pb.PlaybookOutput {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookOutput{}
	out.ExecutionSummary = direct.ValueOf(in.ExecutionSummary)
	out.ActionParameters = ActionParameters_ToProto(mapCtx, in.ActionParameters)
	return out
}
func ToolUse_FromProto(mapCtx *direct.MapContext, in *pb.ToolUse) *krm.ToolUse {
	if in == nil {
		return nil
	}
	out := &krm.ToolUse{}
	out.Tool = direct.LazyPtr(in.GetTool())
	out.Action = direct.LazyPtr(in.GetAction())
	out.InputActionParameters = InputActionParameters_FromProto(mapCtx, in.GetInputActionParameters())
	out.OutputActionParameters = OutputActionParameters_FromProto(mapCtx, in.GetOutputActionParameters())
	return out
}
func ToolUse_ToProto(mapCtx *direct.MapContext, in *krm.ToolUse) *pb.ToolUse {
	if in == nil {
		return nil
	}
	out := &pb.ToolUse{}
	out.Tool = direct.ValueOf(in.Tool)
	out.Action = direct.ValueOf(in.Action)
	out.InputActionParameters = InputActionParameters_ToProto(mapCtx, in.InputActionParameters)
	out.OutputActionParameters = OutputActionParameters_ToProto(mapCtx, in.OutputActionParameters)
	return out
}
func UserUtterance_FromProto(mapCtx *direct.MapContext, in *pb.UserUtterance) *krm.UserUtterance {
	if in == nil {
		return nil
	}
	out := &krm.UserUtterance{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func UserUtterance_ToProto(mapCtx *direct.MapContext, in *krm.UserUtterance) *pb.UserUtterance {
	if in == nil {
		return nil
	}
	out := &pb.UserUtterance{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
