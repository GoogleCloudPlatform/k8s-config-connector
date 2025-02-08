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
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func DialogflowPlaybookVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookVersion) *krm.DialogflowPlaybookVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowPlaybookVersionObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Playbook
	// MISSING: Examples
	// MISSING: UpdateTime
	return out
}
func DialogflowPlaybookVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowPlaybookVersionObservedState) *pb.PlaybookVersion {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookVersion{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Playbook
	// MISSING: Examples
	// MISSING: UpdateTime
	return out
}
func DialogflowPlaybookVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookVersion) *krm.DialogflowPlaybookVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowPlaybookVersionSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Playbook
	// MISSING: Examples
	// MISSING: UpdateTime
	return out
}
func DialogflowPlaybookVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowPlaybookVersionSpec) *pb.PlaybookVersion {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookVersion{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Playbook
	// MISSING: Examples
	// MISSING: UpdateTime
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
func LlmModelSettings_FromProto(mapCtx *direct.MapContext, in *pb.LlmModelSettings) *krm.LlmModelSettings {
	if in == nil {
		return nil
	}
	out := &krm.LlmModelSettings{}
	out.Model = direct.LazyPtr(in.GetModel())
	out.PromptText = direct.LazyPtr(in.GetPromptText())
	return out
}
func LlmModelSettings_ToProto(mapCtx *direct.MapContext, in *krm.LlmModelSettings) *pb.LlmModelSettings {
	if in == nil {
		return nil
	}
	out := &pb.LlmModelSettings{}
	out.Model = direct.ValueOf(in.Model)
	out.PromptText = direct.ValueOf(in.PromptText)
	return out
}
func ParameterDefinition_FromProto(mapCtx *direct.MapContext, in *pb.ParameterDefinition) *krm.ParameterDefinition {
	if in == nil {
		return nil
	}
	out := &krm.ParameterDefinition{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func ParameterDefinition_ToProto(mapCtx *direct.MapContext, in *krm.ParameterDefinition) *pb.ParameterDefinition {
	if in == nil {
		return nil
	}
	out := &pb.ParameterDefinition{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.Enum_ToProto[pb.ParameterDefinition_ParameterType](mapCtx, in.Type)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func Playbook_FromProto(mapCtx *direct.MapContext, in *pb.Playbook) *krm.Playbook {
	if in == nil {
		return nil
	}
	out := &krm.Playbook{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Goal = direct.LazyPtr(in.GetGoal())
	out.InputParameterDefinitions = direct.Slice_FromProto(mapCtx, in.InputParameterDefinitions, ParameterDefinition_FromProto)
	out.OutputParameterDefinitions = direct.Slice_FromProto(mapCtx, in.OutputParameterDefinitions, ParameterDefinition_FromProto)
	out.Instruction = Playbook_Instruction_FromProto(mapCtx, in.GetInstruction())
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReferencedPlaybooks
	// MISSING: ReferencedFlows
	out.ReferencedTools = in.ReferencedTools
	out.LlmModelSettings = LlmModelSettings_FromProto(mapCtx, in.GetLlmModelSettings())
	return out
}
func Playbook_ToProto(mapCtx *direct.MapContext, in *krm.Playbook) *pb.Playbook {
	if in == nil {
		return nil
	}
	out := &pb.Playbook{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Goal = direct.ValueOf(in.Goal)
	out.InputParameterDefinitions = direct.Slice_ToProto(mapCtx, in.InputParameterDefinitions, ParameterDefinition_ToProto)
	out.OutputParameterDefinitions = direct.Slice_ToProto(mapCtx, in.OutputParameterDefinitions, ParameterDefinition_ToProto)
	out.Instruction = Playbook_Instruction_ToProto(mapCtx, in.Instruction)
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReferencedPlaybooks
	// MISSING: ReferencedFlows
	out.ReferencedTools = in.ReferencedTools
	out.LlmModelSettings = LlmModelSettings_ToProto(mapCtx, in.LlmModelSettings)
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
func PlaybookObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Playbook) *krm.PlaybookObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Goal
	// MISSING: InputParameterDefinitions
	// MISSING: OutputParameterDefinitions
	// MISSING: Instruction
	out.TokenCount = direct.LazyPtr(in.GetTokenCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ReferencedPlaybooks = in.ReferencedPlaybooks
	out.ReferencedFlows = in.ReferencedFlows
	// MISSING: ReferencedTools
	// MISSING: LlmModelSettings
	return out
}
func PlaybookObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookObservedState) *pb.Playbook {
	if in == nil {
		return nil
	}
	out := &pb.Playbook{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Goal
	// MISSING: InputParameterDefinitions
	// MISSING: OutputParameterDefinitions
	// MISSING: Instruction
	out.TokenCount = direct.ValueOf(in.TokenCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ReferencedPlaybooks = in.ReferencedPlaybooks
	out.ReferencedFlows = in.ReferencedFlows
	// MISSING: ReferencedTools
	// MISSING: LlmModelSettings
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
func PlaybookVersion_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookVersion) *krm.PlaybookVersion {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Playbook
	// MISSING: Examples
	// MISSING: UpdateTime
	return out
}
func PlaybookVersion_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookVersion) *pb.PlaybookVersion {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Playbook
	// MISSING: Examples
	// MISSING: UpdateTime
	return out
}
func PlaybookVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookVersion) *krm.PlaybookVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookVersionObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.Playbook = Playbook_FromProto(mapCtx, in.GetPlaybook())
	out.Examples = direct.Slice_FromProto(mapCtx, in.Examples, Example_FromProto)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func PlaybookVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookVersionObservedState) *pb.PlaybookVersion {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookVersion{}
	// MISSING: Name
	// MISSING: Description
	out.Playbook = Playbook_ToProto(mapCtx, in.Playbook)
	out.Examples = direct.Slice_ToProto(mapCtx, in.Examples, Example_ToProto)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Playbook_Instruction_FromProto(mapCtx *direct.MapContext, in *pb.Playbook_Instruction) *krm.Playbook_Instruction {
	if in == nil {
		return nil
	}
	out := &krm.Playbook_Instruction{}
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, Playbook_Step_FromProto)
	return out
}
func Playbook_Instruction_ToProto(mapCtx *direct.MapContext, in *krm.Playbook_Instruction) *pb.Playbook_Instruction {
	if in == nil {
		return nil
	}
	out := &pb.Playbook_Instruction{}
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, Playbook_Step_ToProto)
	return out
}
func Playbook_Step_FromProto(mapCtx *direct.MapContext, in *pb.Playbook_Step) *krm.Playbook_Step {
	if in == nil {
		return nil
	}
	out := &krm.Playbook_Step{}
	out.Text = direct.LazyPtr(in.GetText())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, Playbook_Step_FromProto)
	return out
}
func Playbook_Step_ToProto(mapCtx *direct.MapContext, in *krm.Playbook_Step) *pb.Playbook_Step {
	if in == nil {
		return nil
	}
	out := &pb.Playbook_Step{}
	if oneof := Playbook_Step_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Instruction = oneof
	}
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, Playbook_Step_ToProto)
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
