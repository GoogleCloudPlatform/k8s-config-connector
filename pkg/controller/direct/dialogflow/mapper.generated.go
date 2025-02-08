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
func DialogflowPlaybookObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Playbook) *krm.DialogflowPlaybookObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowPlaybookObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Goal
	// MISSING: InputParameterDefinitions
	// MISSING: OutputParameterDefinitions
	// MISSING: Instruction
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReferencedPlaybooks
	// MISSING: ReferencedFlows
	// MISSING: ReferencedTools
	// MISSING: LlmModelSettings
	return out
}
func DialogflowPlaybookObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowPlaybookObservedState) *pb.Playbook {
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
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReferencedPlaybooks
	// MISSING: ReferencedFlows
	// MISSING: ReferencedTools
	// MISSING: LlmModelSettings
	return out
}
func DialogflowPlaybookSpec_FromProto(mapCtx *direct.MapContext, in *pb.Playbook) *krm.DialogflowPlaybookSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowPlaybookSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Goal
	// MISSING: InputParameterDefinitions
	// MISSING: OutputParameterDefinitions
	// MISSING: Instruction
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReferencedPlaybooks
	// MISSING: ReferencedFlows
	// MISSING: ReferencedTools
	// MISSING: LlmModelSettings
	return out
}
func DialogflowPlaybookSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowPlaybookSpec) *pb.Playbook {
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
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReferencedPlaybooks
	// MISSING: ReferencedFlows
	// MISSING: ReferencedTools
	// MISSING: LlmModelSettings
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
