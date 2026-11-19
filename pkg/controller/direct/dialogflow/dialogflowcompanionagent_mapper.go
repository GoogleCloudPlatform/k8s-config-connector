// Copyright 2026 Google LLC
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
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krmcesv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ces/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CesToolSpec_FromProto(mapCtx *direct.MapContext, in *pb.CesToolSpec) *krm.CesToolSpec {
	if in == nil {
		return nil
	}
	out := &krm.CesToolSpec{}
	if in.GetCesTool() != "" {
		out.CesToolRef = &krmcesv1alpha1.CesToolRef{External: in.GetCesTool()}
	}
	out.ConfirmationRequirement = direct.Enum_FromProto(mapCtx, in.GetConfirmationRequirement())
	return out
}

func CesToolSpec_ToProto(mapCtx *direct.MapContext, in *krm.CesToolSpec) *pb.CesToolSpec {
	if in == nil {
		return nil
	}
	out := &pb.CesToolSpec{}
	if in.CesToolRef != nil {
		out.CesTool = in.CesToolRef.External
	}
	out.ConfirmationRequirement = direct.Enum_ToProto[pb.Tool_ConfirmationRequirement](mapCtx, in.ConfirmationRequirement)
	return out
}

func CompanionAgent_GuidanceSkillConfig_FromProto(mapCtx *direct.MapContext, in *pb.CompanionAgent_GuidanceSkillConfig) *krm.CompanionAgent_GuidanceSkillConfig {
	if in == nil {
		return nil
	}
	out := &krm.CompanionAgent_GuidanceSkillConfig{}
	out.GuidanceInstructions = direct.Slice_FromProto(mapCtx, in.GuidanceInstructions, GuidanceInstruction_FromProto)
	out.OverarchingGuidance = direct.LazyPtr(in.GetOverarchingGuidance())
	out.KnowledgeSource = CompanionAgent_KnowledgeSource_FromProto(mapCtx, in.GetKnowledgeSource())
	return out
}

func CompanionAgent_GuidanceSkillConfig_ToProto(mapCtx *direct.MapContext, in *krm.CompanionAgent_GuidanceSkillConfig) *pb.CompanionAgent_GuidanceSkillConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompanionAgent_GuidanceSkillConfig{}
	out.GuidanceInstructions = direct.Slice_ToProto(mapCtx, in.GuidanceInstructions, GuidanceInstruction_ToProto)
	out.OverarchingGuidance = direct.ValueOf(in.OverarchingGuidance)
	out.KnowledgeSource = CompanionAgent_KnowledgeSource_ToProto(mapCtx, in.KnowledgeSource)
	return out
}

func CompanionAgent_KnowledgeSource_FromProto(mapCtx *direct.MapContext, in *pb.CompanionAgent_KnowledgeSource) *krm.CompanionAgent_KnowledgeSource {
	if in == nil {
		return nil
	}
	out := &krm.CompanionAgent_KnowledgeSource{}
	return out
}

func CompanionAgent_KnowledgeSource_ToProto(mapCtx *direct.MapContext, in *krm.CompanionAgent_KnowledgeSource) *pb.CompanionAgent_KnowledgeSource {
	if in == nil {
		return nil
	}
	out := &pb.CompanionAgent_KnowledgeSource{}
	return out
}

func CompanionAgent_SkillConfig_FromProto(mapCtx *direct.MapContext, in *pb.CompanionAgent_SkillConfig) *krm.CompanionAgent_SkillConfig {
	if in == nil {
		return nil
	}
	out := &krm.CompanionAgent_SkillConfig{}
	out.GuidanceSkillConfig = CompanionAgent_GuidanceSkillConfig_FromProto(mapCtx, in.GetGuidanceSkillConfig())
	out.TranslationSkillConfig = CompanionAgent_TranslationSkillConfig_FromProto(mapCtx, in.GetTranslationSkillConfig())
	out.SkillTriggeringEvent = direct.Enum_FromProto(mapCtx, in.GetSkillTriggeringEvent())
	return out
}

func CompanionAgent_SkillConfig_ToProto(mapCtx *direct.MapContext, in *krm.CompanionAgent_SkillConfig) *pb.CompanionAgent_SkillConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompanionAgent_SkillConfig{}
	if oneof := CompanionAgent_GuidanceSkillConfig_ToProto(mapCtx, in.GuidanceSkillConfig); oneof != nil {
		out.Config = &pb.CompanionAgent_SkillConfig_GuidanceSkillConfig{GuidanceSkillConfig: oneof}
	}
	if oneof := CompanionAgent_TranslationSkillConfig_ToProto(mapCtx, in.TranslationSkillConfig); oneof != nil {
		out.Config = &pb.CompanionAgent_SkillConfig_TranslationSkillConfig{TranslationSkillConfig: oneof}
	}
	out.SkillTriggeringEvent = direct.Enum_ToProto[pb.CompanionAgent_SkillConfig_SkillTriggerEvent](mapCtx, in.SkillTriggeringEvent)
	return out
}

func CompanionAgent_TranslationSkillConfig_FromProto(mapCtx *direct.MapContext, in *pb.CompanionAgent_TranslationSkillConfig) *krm.CompanionAgent_TranslationSkillConfig {
	if in == nil {
		return nil
	}
	out := &krm.CompanionAgent_TranslationSkillConfig{}
	out.AgentLanguageCode = direct.LazyPtr(in.GetAgentLanguageCode())
	out.CustomerLanguageCode = direct.LazyPtr(in.GetCustomerLanguageCode())
	return out
}

func CompanionAgent_TranslationSkillConfig_ToProto(mapCtx *direct.MapContext, in *krm.CompanionAgent_TranslationSkillConfig) *pb.CompanionAgent_TranslationSkillConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompanionAgent_TranslationSkillConfig{}
	out.AgentLanguageCode = direct.ValueOf(in.AgentLanguageCode)
	out.CustomerLanguageCode = direct.ValueOf(in.CustomerLanguageCode)
	return out
}

func DialogflowCompanionAgentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompanionAgent) *krm.DialogflowCompanionAgentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCompanionAgentObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func DialogflowCompanionAgentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCompanionAgentObservedState) *pb.CompanionAgent {
	if in == nil {
		return nil
	}
	out := &pb.CompanionAgent{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func DialogflowCompanionAgentSpec_FromProto(mapCtx *direct.MapContext, in *pb.CompanionAgent) *krm.DialogflowCompanionAgentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCompanionAgentSpec{}
	// MISSING: Name
	out.ToolsetTools = direct.Slice_FromProto(mapCtx, in.ToolsetTools, ToolsetTool_FromProto)
	out.CesToolSpecs = direct.Slice_FromProto(mapCtx, in.CesToolSpecs, CesToolSpec_FromProto)
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.SkillConfigs = direct.Slice_FromProto(mapCtx, in.SkillConfigs, CompanionAgent_SkillConfig_FromProto)
	return out
}

func DialogflowCompanionAgentSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCompanionAgentSpec) *pb.CompanionAgent {
	if in == nil {
		return nil
	}
	out := &pb.CompanionAgent{}
	// MISSING: Name
	out.ToolsetTools = direct.Slice_ToProto(mapCtx, in.ToolsetTools, ToolsetTool_ToProto)
	out.CesToolSpecs = direct.Slice_ToProto(mapCtx, in.CesToolSpecs, CesToolSpec_ToProto)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.SkillConfigs = direct.Slice_ToProto(mapCtx, in.SkillConfigs, CompanionAgent_SkillConfig_ToProto)
	return out
}

func GuidanceInstruction_FromProto(mapCtx *direct.MapContext, in *pb.GuidanceInstruction) *krm.GuidanceInstruction {
	if in == nil {
		return nil
	}
	out := &krm.GuidanceInstruction{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DisplayDetails = direct.LazyPtr(in.GetDisplayDetails())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, GuidanceInstruction_Action_FromProto)
	out.TriggerEvent = direct.Enum_FromProto(mapCtx, in.GetTriggerEvent())
	out.DisableSuggestedReply = direct.LazyPtr(in.GetDisableSuggestedReply())
	return out
}

func GuidanceInstruction_ToProto(mapCtx *direct.MapContext, in *krm.GuidanceInstruction) *pb.GuidanceInstruction {
	if in == nil {
		return nil
	}
	out := &pb.GuidanceInstruction{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DisplayDetails = direct.ValueOf(in.DisplayDetails)
	out.Condition = direct.ValueOf(in.Condition)
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, GuidanceInstruction_Action_ToProto)
	out.TriggerEvent = direct.Enum_ToProto[pb.CompanionAgent_TriggerEvent](mapCtx, in.TriggerEvent)
	out.DisableSuggestedReply = direct.ValueOf(in.DisableSuggestedReply)
	return out
}

func GuidanceInstruction_Action_FromProto(mapCtx *direct.MapContext, in *pb.GuidanceInstruction_Action) *krm.GuidanceInstruction_Action {
	if in == nil {
		return nil
	}
	out := &krm.GuidanceInstruction_Action{}
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func GuidanceInstruction_Action_ToProto(mapCtx *direct.MapContext, in *krm.GuidanceInstruction_Action) *pb.GuidanceInstruction_Action {
	if in == nil {
		return nil
	}
	out := &pb.GuidanceInstruction_Action{}
	out.Description = direct.ValueOf(in.Description)
	return out
}

func ToolsetTool_FromProto(mapCtx *direct.MapContext, in *pb.ToolsetTool) *krm.ToolsetTool {
	if in == nil {
		return nil
	}
	out := &krm.ToolsetTool{}
	if in.GetToolset() != "" {
		out.ToolsetRef = &krmcesv1alpha1.CesToolsetRef{External: in.GetToolset()}
	}
	out.OperationID = direct.LazyPtr(in.GetOperationId())
	out.ConfirmationRequirement = direct.Enum_FromProto(mapCtx, in.GetConfirmationRequirement())
	return out
}

func ToolsetTool_ToProto(mapCtx *direct.MapContext, in *krm.ToolsetTool) *pb.ToolsetTool {
	if in == nil {
		return nil
	}
	out := &pb.ToolsetTool{}
	if in.ToolsetRef != nil {
		out.Toolset = in.ToolsetRef.External
	}
	out.OperationId = direct.ValueOf(in.OperationID)
	out.ConfirmationRequirement = direct.Enum_ToProto[pb.Tool_ConfirmationRequirement](mapCtx, in.ConfirmationRequirement)
	return out
}
