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
	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CodeBlock_FromProto(mapCtx *direct.MapContext, in *pb.CodeBlock) *krm.CodeBlock {
	if in == nil {
		return nil
	}
	out := &krm.CodeBlock{}
	out.Code = direct.LazyPtr(in.GetCode())
	return out
}

func CodeBlock_ToProto(mapCtx *direct.MapContext, in *krm.CodeBlock) *pb.CodeBlock {
	if in == nil {
		return nil
	}
	out := &pb.CodeBlock{}
	out.Code = direct.ValueOf(in.Code)
	return out
}

func DialogflowConversationModelObservedState_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.ConversationModel) *krm.DialogflowConversationModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}

func DialogflowConversationModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelObservedState) *dialogflowpb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.ConversationModel{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[dialogflowpb.ConversationModel_State](mapCtx, in.State)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}

func DialogflowConversationModelSpec_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.ConversationModel) *krm.DialogflowConversationModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Datasets = direct.Slice_FromProto(mapCtx, in.Datasets, InputDataset_FromProto)
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.ArticleSuggestionModelMetadata = ArticleSuggestionModelMetadata_FromProto(mapCtx, in.GetArticleSuggestionModelMetadata())
	out.SmartReplyModelMetadata = SmartReplyModelMetadata_FromProto(mapCtx, in.GetSmartReplyModelMetadata())
	return out
}

func DialogflowConversationModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelSpec) *dialogflowpb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.ConversationModel{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Datasets = direct.Slice_ToProto(mapCtx, in.Datasets, InputDataset_ToProto)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	if oneof := ArticleSuggestionModelMetadata_ToProto(mapCtx, in.ArticleSuggestionModelMetadata); oneof != nil {
		out.ModelMetadata = &dialogflowpb.ConversationModel_ArticleSuggestionModelMetadata{ArticleSuggestionModelMetadata: oneof}
	}
	if oneof := SmartReplyModelMetadata_ToProto(mapCtx, in.SmartReplyModelMetadata); oneof != nil {
		out.ModelMetadata = &dialogflowpb.ConversationModel_SmartReplyModelMetadata{SmartReplyModelMetadata: oneof}
	}
	return out
}

func ArticleSuggestionModelMetadata_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.ArticleSuggestionModelMetadata) *krm.ArticleSuggestionModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ArticleSuggestionModelMetadata{}
	out.TrainingModelType = direct.Enum_FromProto(mapCtx, in.GetTrainingModelType())
	return out
}

func ArticleSuggestionModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ArticleSuggestionModelMetadata) *dialogflowpb.ArticleSuggestionModelMetadata {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.ArticleSuggestionModelMetadata{}
	out.TrainingModelType = direct.Enum_ToProto[dialogflowpb.ConversationModel_ModelType](mapCtx, in.TrainingModelType)
	return out
}

func SmartReplyModelMetadata_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.SmartReplyModelMetadata) *krm.SmartReplyModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.SmartReplyModelMetadata{}
	out.TrainingModelType = direct.Enum_FromProto(mapCtx, in.GetTrainingModelType())
	return out
}

func SmartReplyModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.SmartReplyModelMetadata) *dialogflowpb.SmartReplyModelMetadata {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.SmartReplyModelMetadata{}
	out.TrainingModelType = direct.Enum_ToProto[dialogflowpb.ConversationModel_ModelType](mapCtx, in.TrainingModelType)
	return out
}

func DialogflowPlaybookObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Playbook) *krm.DialogflowPlaybookObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowPlaybookObservedState{}
	// MISSING: Name
	out.TokenCount = direct.LazyPtr(in.GetTokenCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ReferencedPlaybooks = in.ReferencedPlaybooks
	out.ReferencedFlows = in.ReferencedFlows
	out.InlineActions = in.InlineActions
	out.Handlers = direct.Slice_FromProto(mapCtx, in.Handlers, HandlerObservedState_FromProto)
	return out
}

func DialogflowPlaybookObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowPlaybookObservedState) *pb.Playbook {
	if in == nil {
		return nil
	}
	out := &pb.Playbook{}
	// MISSING: Name
	out.TokenCount = direct.ValueOf(in.TokenCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ReferencedPlaybooks = in.ReferencedPlaybooks
	out.ReferencedFlows = in.ReferencedFlows
	out.InlineActions = in.InlineActions
	out.Handlers = direct.Slice_ToProto(mapCtx, in.Handlers, HandlerObservedState_ToProto)
	return out
}

func DialogflowPlaybookSpec_FromProto(mapCtx *direct.MapContext, in *pb.Playbook) *krm.DialogflowPlaybookSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowPlaybookSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Goal = direct.LazyPtr(in.GetGoal())
	out.InputParameterDefinitions = direct.Slice_FromProto(mapCtx, in.InputParameterDefinitions, ParameterDefinition_FromProto)
	out.OutputParameterDefinitions = direct.Slice_FromProto(mapCtx, in.OutputParameterDefinitions, ParameterDefinition_FromProto)
	out.Instruction = Playbook_Instruction_FromProto(mapCtx, in.GetInstruction())
	for _, v := range in.GetReferencedTools() {
		out.ReferencedToolRefs = append(out.ReferencedToolRefs, krm.DialogflowToolRef{External: v})
	}
	out.CodeBlock = CodeBlock_FromProto(mapCtx, in.GetCodeBlock())
	out.LlmModelSettings = LlmModelSettings_FromProto(mapCtx, in.GetLlmModelSettings())
	out.Handlers = direct.Slice_FromProto(mapCtx, in.Handlers, Handler_FromProto)
	out.PlaybookType = direct.Enum_FromProto(mapCtx, in.GetPlaybookType())
	return out
}

func DialogflowPlaybookSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowPlaybookSpec) *pb.Playbook {
	if in == nil {
		return nil
	}
	out := &pb.Playbook{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Goal = direct.ValueOf(in.Goal)
	out.InputParameterDefinitions = direct.Slice_ToProto(mapCtx, in.InputParameterDefinitions, ParameterDefinition_ToProto)
	out.OutputParameterDefinitions = direct.Slice_ToProto(mapCtx, in.OutputParameterDefinitions, ParameterDefinition_ToProto)
	out.Instruction = Playbook_Instruction_ToProto(mapCtx, in.Instruction)
	for _, v := range in.ReferencedToolRefs {
		out.ReferencedTools = append(out.ReferencedTools, v.External)
	}
	out.CodeBlock = CodeBlock_ToProto(mapCtx, in.CodeBlock)
	out.LlmModelSettings = LlmModelSettings_ToProto(mapCtx, in.LlmModelSettings)
	out.Handlers = direct.Slice_ToProto(mapCtx, in.Handlers, Handler_ToProto)
	out.PlaybookType = direct.Enum_ToProto[pb.Playbook_PlaybookType](mapCtx, in.PlaybookType)
	return out
}

func HandlerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Handler) *krm.HandlerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HandlerObservedState{}
	out.EventHandler = Handler_EventHandlerObservedState_FromProto(mapCtx, in.GetEventHandler())
	out.LifecycleHandler = Handler_LifecycleHandlerObservedState_FromProto(mapCtx, in.GetLifecycleHandler())
	return out
}

func HandlerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HandlerObservedState) *pb.Handler {
	if in == nil {
		return nil
	}
	out := &pb.Handler{}
	if oneof := Handler_EventHandlerObservedState_ToProto(mapCtx, in.EventHandler); oneof != nil {
		out.Handler = &pb.Handler_EventHandler_{EventHandler: oneof}
	}
	if oneof := Handler_LifecycleHandlerObservedState_ToProto(mapCtx, in.LifecycleHandler); oneof != nil {
		out.Handler = &pb.Handler_LifecycleHandler_{LifecycleHandler: oneof}
	}
	return out
}

func Handler_EventHandlerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Handler_EventHandler) *krm.Handler_EventHandlerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Handler_EventHandlerObservedState{}
	// MISSING: Event
	// MISSING: Condition
	out.Fulfillment = FulfillmentObservedState_FromProto(mapCtx, in.GetFulfillment())
	return out
}

func Handler_EventHandlerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Handler_EventHandlerObservedState) *pb.Handler_EventHandler {
	if in == nil {
		return nil
	}
	out := &pb.Handler_EventHandler{}
	// MISSING: Event
	// MISSING: Condition
	out.Fulfillment = FulfillmentObservedState_ToProto(mapCtx, in.Fulfillment)
	return out
}

func Handler_EventHandler_FromProto(mapCtx *direct.MapContext, in *pb.Handler_EventHandler) *krm.Handler_EventHandler {
	if in == nil {
		return nil
	}
	out := &krm.Handler_EventHandler{}
	out.Event = direct.LazyPtr(in.GetEvent())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Fulfillment = DialogflowCX_Fulfillment_FromProto(mapCtx, in.GetFulfillment())
	return out
}

func Handler_EventHandler_ToProto(mapCtx *direct.MapContext, in *krm.Handler_EventHandler) *pb.Handler_EventHandler {
	if in == nil {
		return nil
	}
	out := &pb.Handler_EventHandler{}
	out.Event = direct.ValueOf(in.Event)
	out.Condition = direct.ValueOf(in.Condition)
	out.Fulfillment = DialogflowCX_Fulfillment_ToProto(mapCtx, in.Fulfillment)
	return out
}

func Handler_FromProto(mapCtx *direct.MapContext, in *pb.Handler) *krm.Handler {
	if in == nil {
		return nil
	}
	out := &krm.Handler{}
	out.EventHandler = Handler_EventHandler_FromProto(mapCtx, in.GetEventHandler())
	out.LifecycleHandler = Handler_LifecycleHandler_FromProto(mapCtx, in.GetLifecycleHandler())
	return out
}

func Handler_LifecycleHandlerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Handler_LifecycleHandler) *krm.Handler_LifecycleHandlerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Handler_LifecycleHandlerObservedState{}
	// MISSING: LifecycleStage
	// MISSING: Condition
	out.Fulfillment = FulfillmentObservedState_FromProto(mapCtx, in.GetFulfillment())
	return out
}

func Handler_LifecycleHandlerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Handler_LifecycleHandlerObservedState) *pb.Handler_LifecycleHandler {
	if in == nil {
		return nil
	}
	out := &pb.Handler_LifecycleHandler{}
	// MISSING: LifecycleStage
	// MISSING: Condition
	out.Fulfillment = FulfillmentObservedState_ToProto(mapCtx, in.Fulfillment)
	return out
}

func Handler_LifecycleHandler_FromProto(mapCtx *direct.MapContext, in *pb.Handler_LifecycleHandler) *krm.Handler_LifecycleHandler {
	if in == nil {
		return nil
	}
	out := &krm.Handler_LifecycleHandler{}
	out.LifecycleStage = direct.LazyPtr(in.GetLifecycleStage())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Fulfillment = DialogflowCX_Fulfillment_FromProto(mapCtx, in.GetFulfillment())
	return out
}

func Handler_LifecycleHandler_ToProto(mapCtx *direct.MapContext, in *krm.Handler_LifecycleHandler) *pb.Handler_LifecycleHandler {
	if in == nil {
		return nil
	}
	out := &pb.Handler_LifecycleHandler{}
	out.LifecycleStage = direct.ValueOf(in.LifecycleStage)
	out.Condition = direct.ValueOf(in.Condition)
	out.Fulfillment = DialogflowCX_Fulfillment_ToProto(mapCtx, in.Fulfillment)
	return out
}

func Handler_ToProto(mapCtx *direct.MapContext, in *krm.Handler) *pb.Handler {
	if in == nil {
		return nil
	}
	out := &pb.Handler{}
	if oneof := Handler_EventHandler_ToProto(mapCtx, in.EventHandler); oneof != nil {
		out.Handler = &pb.Handler_EventHandler_{EventHandler: oneof}
	}
	if oneof := Handler_LifecycleHandler_ToProto(mapCtx, in.LifecycleHandler); oneof != nil {
		out.Handler = &pb.Handler_LifecycleHandler_{LifecycleHandler: oneof}
	}
	return out
}

func InlineSchema_FromProto(mapCtx *direct.MapContext, in *pb.InlineSchema) *krm.InlineSchema {
	if in == nil {
		return nil
	}
	out := &krm.InlineSchema{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Items = TypeSchema_FromProto(mapCtx, in.GetItems())
	return out
}

func InlineSchema_ToProto(mapCtx *direct.MapContext, in *krm.InlineSchema) *pb.InlineSchema {
	if in == nil {
		return nil
	}
	out := &pb.InlineSchema{}
	out.Type = direct.Enum_ToProto[pb.DataType](mapCtx, in.Type)
	out.Items = TypeSchema_ToProto(mapCtx, in.Items)
	return out
}

func InputDataset_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.InputDataset) *krm.InputDataset {
	if in == nil {
		return nil
	}
	out := &krm.InputDataset{}
	if in.GetDataset() != "" {
		out.DatasetRef = &krm.DialogflowConversationDatasetRef{External: in.GetDataset()}
	}
	return out
}

func InputDataset_ToProto(mapCtx *direct.MapContext, in *krm.InputDataset) *dialogflowpb.InputDataset {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.InputDataset{}
	if in.DatasetRef != nil {
		out.Dataset = in.DatasetRef.External
	}
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
	out.TypeSchema = TypeSchema_FromProto(mapCtx, in.GetTypeSchema())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func ParameterDefinition_ToProto(mapCtx *direct.MapContext, in *krm.ParameterDefinition) *pb.ParameterDefinition {
	if in == nil {
		return nil
	}
	out := &pb.ParameterDefinition{}
	out.Name = direct.ValueOf(in.Name)
	out.TypeSchema = TypeSchema_ToProto(mapCtx, in.TypeSchema)
	out.Description = direct.ValueOf(in.Description)
	return out
}

func Playbook_Instruction_FromProto(mapCtx *direct.MapContext, in *pb.Playbook_Instruction) *krm.Playbook_Instruction {
	if in == nil {
		return nil
	}
	out := &krm.Playbook_Instruction{}
	out.Guidelines = direct.LazyPtr(in.GetGuidelines())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, Playbook_Step_FromProto)
	return out
}

func Playbook_Instruction_ToProto(mapCtx *direct.MapContext, in *krm.Playbook_Instruction) *pb.Playbook_Instruction {
	if in == nil {
		return nil
	}
	out := &pb.Playbook_Instruction{}
	out.Guidelines = direct.ValueOf(in.Guidelines)
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

func Playbook_Step_Text_ToProto(mapCtx *direct.MapContext, in *string) *pb.Playbook_Step_Text {
	if in == nil {
		return nil
	}
	return &pb.Playbook_Step_Text{Text: *in}
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

func TypeSchema_FromProto(mapCtx *direct.MapContext, in *pb.TypeSchema) *krm.TypeSchema {
	if in == nil {
		return nil
	}
	out := &krm.TypeSchema{}
	out.InlineSchema = InlineSchema_FromProto(mapCtx, in.GetInlineSchema())
	out.SchemaReference = TypeSchema_SchemaReference_FromProto(mapCtx, in.GetSchemaReference())
	return out
}

func TypeSchema_SchemaReference_FromProto(mapCtx *direct.MapContext, in *pb.TypeSchema_SchemaReference) *krm.TypeSchema_SchemaReference {
	if in == nil {
		return nil
	}
	out := &krm.TypeSchema_SchemaReference{}
	if in.GetTool() != "" {
		out.ToolRef = &krm.DialogflowToolRef{External: in.GetTool()}
	}
	out.Schema = direct.LazyPtr(in.GetSchema())
	return out
}

func TypeSchema_SchemaReference_ToProto(mapCtx *direct.MapContext, in *krm.TypeSchema_SchemaReference) *pb.TypeSchema_SchemaReference {
	if in == nil {
		return nil
	}
	out := &pb.TypeSchema_SchemaReference{}
	if in.ToolRef != nil {
		out.Tool = in.ToolRef.External
	}
	out.Schema = direct.ValueOf(in.Schema)
	return out
}

func TypeSchema_ToProto(mapCtx *direct.MapContext, in *krm.TypeSchema) *pb.TypeSchema {
	if in == nil {
		return nil
	}
	out := &pb.TypeSchema{}
	if oneof := InlineSchema_ToProto(mapCtx, in.InlineSchema); oneof != nil {
		out.Schema = &pb.TypeSchema_InlineSchema{InlineSchema: oneof}
	}
	if oneof := TypeSchema_SchemaReference_ToProto(mapCtx, in.SchemaReference); oneof != nil {
		out.Schema = &pb.TypeSchema_SchemaReference_{SchemaReference: oneof}
	}
	return out
}
