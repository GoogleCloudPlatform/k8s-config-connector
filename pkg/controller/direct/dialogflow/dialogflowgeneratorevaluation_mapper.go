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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DialogflowGeneratorEvaluationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowGeneratorEvaluationSpec) *pb.GeneratorEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorEvaluation{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.GeneratorEvaluationConfig = GeneratorEvaluationConfig_ToProto(mapCtx, in.GeneratorEvaluationConfig)
	out.InitialGenerator = DialogflowGeneratorEvaluation_Generator_ToProto(mapCtx, in.InitialGenerator)
	return out
}

func DialogflowGeneratorEvaluationSpec_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorEvaluation) *krm.DialogflowGeneratorEvaluationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowGeneratorEvaluationSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GeneratorEvaluationConfig = GeneratorEvaluationConfig_FromProto(mapCtx, in.GetGeneratorEvaluationConfig())
	out.InitialGenerator = DialogflowGeneratorEvaluation_Generator_FromProto(mapCtx, in.GetInitialGenerator())
	return out
}

func DialogflowGeneratorEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowGeneratorEvaluationObservedState) *pb.GeneratorEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorEvaluation{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.InitialGenerator = DialogflowGeneratorObservedState_ToProto(mapCtx, in.InitialGenerator)
	if in.SummarizationMetrics != nil {
		out.Metrics = &pb.GeneratorEvaluation_SummarizationMetrics{
			SummarizationMetrics: SummarizationEvaluationMetricsObservedState_ToProto(mapCtx, in.SummarizationMetrics),
		}
	}
	out.EvaluationStatus = EvaluationStatusObservedState_ToProto(mapCtx, in.EvaluationStatus)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}

func DialogflowGeneratorEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorEvaluation) *krm.DialogflowGeneratorEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowGeneratorEvaluationObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.InitialGenerator = DialogflowGeneratorObservedState_FromProto(mapCtx, in.GetInitialGenerator())
	if summarizationMetrics := in.GetSummarizationMetrics(); summarizationMetrics != nil {
		out.SummarizationMetrics = SummarizationEvaluationMetricsObservedState_FromProto(mapCtx, summarizationMetrics)
	}
	out.EvaluationStatus = EvaluationStatusObservedState_FromProto(mapCtx, in.GetEvaluationStatus())
	if in.SatisfiesPzs != nil {
		out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	}
	if in.SatisfiesPzi != nil {
		out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	}
	return out
}

func DialogflowGeneratorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowGeneratorObservedState) *pb.Generator {
	if in == nil {
		return nil
	}
	out := &pb.Generator{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func DialogflowGeneratorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Generator) *krm.DialogflowGeneratorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowGeneratorObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func DialogflowGeneratorEvaluation_Generator_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowGeneratorEvaluation_Generator) *pb.Generator {
	if in == nil {
		return nil
	}
	out := &pb.Generator{}
	out.Description = direct.ValueOf(in.Description)
	if in.FreeFormContext != nil {
		out.Context = &pb.Generator_FreeFormContext{
			FreeFormContext: FreeFormContext_ToProto(mapCtx, in.FreeFormContext),
		}
	}
	if in.SummarizationContext != nil {
		out.Context = &pb.Generator_SummarizationContext{
			SummarizationContext: SummarizationContextWithRequired_ToProto(mapCtx, in.SummarizationContext),
		}
	}
	out.InferenceParameter = InferenceParameter_ToProto(mapCtx, in.InferenceParameter)
	out.TriggerEvent = direct.Enum_ToProto[pb.TriggerEvent](mapCtx, in.TriggerEvent)
	if in.PublishedModel != nil {
		out.FoundationModel = &pb.Generator_PublishedModel{
			PublishedModel: *in.PublishedModel,
		}
	}
	return out
}

func DialogflowGeneratorEvaluation_Generator_FromProto(mapCtx *direct.MapContext, in *pb.Generator) *krm.DialogflowGeneratorEvaluation_Generator {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowGeneratorEvaluation_Generator{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if oneof := in.GetFreeFormContext(); oneof != nil {
		out.FreeFormContext = FreeFormContext_FromProto(mapCtx, oneof)
	}
	if oneof := in.GetSummarizationContext(); oneof != nil {
		out.SummarizationContext = SummarizationContextWithRequired_FromProto(mapCtx, oneof)
	}
	out.InferenceParameter = InferenceParameter_FromProto(mapCtx, in.GetInferenceParameter())
	out.TriggerEvent = direct.Enum_FromProto(mapCtx, in.GetTriggerEvent())
	if publishedModel := in.GetPublishedModel(); publishedModel != "" {
		out.PublishedModel = direct.LazyPtr(publishedModel)
	}
	return out
}

func GeneratorEvaluationConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorEvaluationConfig) *pb.GeneratorEvaluationConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorEvaluationConfig{}
	out.InputDataConfig = GeneratorEvaluationConfig_InputDataConfig_ToProto(mapCtx, in.InputDataConfig)
	if in.SummarizationConfig != nil {
		out.EvaluationFeatureConfig = &pb.GeneratorEvaluationConfig_SummarizationConfig_{
			SummarizationConfig: GeneratorEvaluationConfig_SummarizationConfig_ToProto(mapCtx, in.SummarizationConfig),
		}
	}
	out.OutputGcsBucketPath = direct.ValueOf(in.OutputGCSBucketPath)
	return out
}

func GeneratorEvaluationConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorEvaluationConfig) *krm.GeneratorEvaluationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorEvaluationConfig{}
	out.InputDataConfig = GeneratorEvaluationConfig_InputDataConfig_FromProto(mapCtx, in.GetInputDataConfig())
	if oneof := in.GetSummarizationConfig(); oneof != nil {
		out.SummarizationConfig = GeneratorEvaluationConfig_SummarizationConfig_FromProto(mapCtx, oneof)
	}
	out.OutputGCSBucketPath = direct.LazyPtr(in.GetOutputGcsBucketPath())
	return out
}

func GeneratorEvaluationConfig_InputDataConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorEvaluationConfig_InputDataConfig) *pb.GeneratorEvaluationConfig_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorEvaluationConfig_InputDataConfig{}
	out.SampleSize = direct.ValueOf(in.SampleSize)
	out.SummaryGenerationOption = direct.Enum_ToProto[pb.GeneratorEvaluationConfig_InputDataConfig_SummaryGenerationOption](mapCtx, in.SummaryGenerationOption)
	if in.DatasetInputDataConfig != nil {
		out.SourceSpecificConfig = &pb.GeneratorEvaluationConfig_InputDataConfig_DatasetInputDataConfig{
			DatasetInputDataConfig: GeneratorEvaluationConfig_DatasetInputDataConfig_ToProto(mapCtx, in.DatasetInputDataConfig),
		}
	}
	if in.AgentAssistInputDataConfig != nil {
		out.SourceSpecificConfig = &pb.GeneratorEvaluationConfig_InputDataConfig_AgentAssistInputDataConfig{
			AgentAssistInputDataConfig: GeneratorEvaluationConfig_AgentAssistInputDataConfig_ToProto(mapCtx, in.AgentAssistInputDataConfig),
		}
	}
	return out
}

func GeneratorEvaluationConfig_InputDataConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorEvaluationConfig_InputDataConfig) *krm.GeneratorEvaluationConfig_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorEvaluationConfig_InputDataConfig{}
	out.SampleSize = direct.LazyPtr(in.GetSampleSize())
	out.SummaryGenerationOption = direct.Enum_FromProto(mapCtx, in.GetSummaryGenerationOption())
	if oneof := in.GetDatasetInputDataConfig(); oneof != nil {
		out.DatasetInputDataConfig = GeneratorEvaluationConfig_DatasetInputDataConfig_FromProto(mapCtx, oneof)
	}
	if oneof := in.GetAgentAssistInputDataConfig(); oneof != nil {
		out.AgentAssistInputDataConfig = GeneratorEvaluationConfig_AgentAssistInputDataConfig_FromProto(mapCtx, oneof)
	}
	return out
}

func GeneratorEvaluationConfig_DatasetInputDataConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorEvaluationConfig_DatasetInputDataConfig) *pb.GeneratorEvaluationConfig_DatasetInputDataConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorEvaluationConfig_DatasetInputDataConfig{}
	if in.DatasetRef != nil {
		out.Dataset = in.DatasetRef.External
	}
	return out
}

func GeneratorEvaluationConfig_DatasetInputDataConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorEvaluationConfig_DatasetInputDataConfig) *krm.GeneratorEvaluationConfig_DatasetInputDataConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorEvaluationConfig_DatasetInputDataConfig{}
	if in.GetDataset() != "" {
		out.DatasetRef = &vertexaiv1beta1.VertexAIDatasetRef{External: in.GetDataset()}
	}
	return out
}

func GeneratorEvaluationConfig_AgentAssistInputDataConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorEvaluationConfig_AgentAssistInputDataConfig) *pb.GeneratorEvaluationConfig_AgentAssistInputDataConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorEvaluationConfig_AgentAssistInputDataConfig{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}

func GeneratorEvaluationConfig_AgentAssistInputDataConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorEvaluationConfig_AgentAssistInputDataConfig) *krm.GeneratorEvaluationConfig_AgentAssistInputDataConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorEvaluationConfig_AgentAssistInputDataConfig{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}

func GeneratorEvaluationConfig_SummarizationConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorEvaluationConfig_SummarizationConfig) *pb.GeneratorEvaluationConfig_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorEvaluationConfig_SummarizationConfig{}
	out.EvaluatorVersion = in.EvaluatorVersion
	return out
}

func GeneratorEvaluationConfig_SummarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorEvaluationConfig_SummarizationConfig) *krm.GeneratorEvaluationConfig_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorEvaluationConfig_SummarizationConfig{}
	out.EvaluatorVersion = direct.LazyPtr(in.GetEvaluatorVersion())
	return out
}

func GeneratorSuggestionWithRequired_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorSuggestionWithRequired) *pb.GeneratorSuggestion {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorSuggestion{}
	if in.FreeFormSuggestion != nil {
		out.Suggestion = &pb.GeneratorSuggestion_FreeFormSuggestion{
			FreeFormSuggestion: FreeFormSuggestionWithRequired_ToProto(mapCtx, in.FreeFormSuggestion),
		}
	}
	if in.SummarySuggestion != nil {
		out.Suggestion = &pb.GeneratorSuggestion_SummarySuggestion{
			SummarySuggestion: SummarySuggestionWithRequired_ToProto(mapCtx, in.SummarySuggestion),
		}
	}
	return out
}

func GeneratorSuggestionWithRequired_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorSuggestion) *krm.GeneratorSuggestionWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorSuggestionWithRequired{}
	if oneof := in.GetFreeFormSuggestion(); oneof != nil {
		out.FreeFormSuggestion = FreeFormSuggestionWithRequired_FromProto(mapCtx, oneof)
	}
	if oneof := in.GetSummarySuggestion(); oneof != nil {
		out.SummarySuggestion = SummarySuggestionWithRequired_FromProto(mapCtx, oneof)
	}
	return out
}

func FreeFormSuggestionWithRequired_ToProto(mapCtx *direct.MapContext, in *krm.FreeFormSuggestionWithRequired) *pb.FreeFormSuggestion {
	if in == nil {
		return nil
	}
	out := &pb.FreeFormSuggestion{}
	out.Response = direct.ValueOf(in.Response)
	return out
}

func FreeFormSuggestionWithRequired_FromProto(mapCtx *direct.MapContext, in *pb.FreeFormSuggestion) *krm.FreeFormSuggestionWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.FreeFormSuggestionWithRequired{}
	out.Response = direct.LazyPtr(in.GetResponse())
	return out
}

func SummarySuggestionWithRequired_ToProto(mapCtx *direct.MapContext, in *krm.SummarySuggestionWithRequired) *pb.SummarySuggestion {
	if in == nil {
		return nil
	}
	out := &pb.SummarySuggestion{}
	out.SummarySections = direct.Slice_ToProto(mapCtx, in.SummarySections, SummarySuggestion_SummarySectionWithRequired_ToProto)
	return out
}

func SummarySuggestionWithRequired_FromProto(mapCtx *direct.MapContext, in *pb.SummarySuggestion) *krm.SummarySuggestionWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.SummarySuggestionWithRequired{}
	out.SummarySections = direct.Slice_FromProto(mapCtx, in.SummarySections, SummarySuggestion_SummarySectionWithRequired_FromProto)
	return out
}

func SummarySuggestion_SummarySectionWithRequired_ToProto(mapCtx *direct.MapContext, in *krm.SummarySuggestion_SummarySectionWithRequired) *pb.SummarySuggestion_SummarySection {
	if in == nil {
		return nil
	}
	out := &pb.SummarySuggestion_SummarySection{}
	out.Section = direct.ValueOf(in.Section)
	out.Summary = direct.ValueOf(in.Summary)
	return out
}

func SummarySuggestion_SummarySectionWithRequired_FromProto(mapCtx *direct.MapContext, in *pb.SummarySuggestion_SummarySection) *krm.SummarySuggestion_SummarySectionWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.SummarySuggestion_SummarySectionWithRequired{}
	out.Section = direct.LazyPtr(in.GetSection())
	out.Summary = direct.LazyPtr(in.GetSummary())
	return out
}

func SummarizationContextWithRequired_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationContextWithRequired) *pb.SummarizationContext {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationContext{}
	out.SummarizationSections = direct.Slice_ToProto(mapCtx, in.SummarizationSections, SummarizationSection_ToProto)
	out.FewShotExamples = direct.Slice_ToProto(mapCtx, in.FewShotExamples, FewShotExampleWithRequired_ToProto)
	out.Version = direct.ValueOf(in.Version)
	out.OutputLanguageCode = direct.ValueOf(in.OutputLanguageCode)
	return out
}

func SummarizationContextWithRequired_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationContext) *krm.SummarizationContextWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationContextWithRequired{}
	out.SummarizationSections = direct.Slice_FromProto(mapCtx, in.SummarizationSections, SummarizationSection_FromProto)
	out.FewShotExamples = direct.Slice_FromProto(mapCtx, in.FewShotExamples, FewShotExampleWithRequired_FromProto)
	out.Version = direct.LazyPtr(in.GetVersion())
	out.OutputLanguageCode = direct.LazyPtr(in.GetOutputLanguageCode())
	return out
}

func FewShotExampleWithRequired_ToProto(mapCtx *direct.MapContext, in *krm.FewShotExampleWithRequired) *pb.FewShotExample {
	if in == nil {
		return nil
	}
	out := &pb.FewShotExample{}
	out.ConversationContext = ConversationContext_ToProto(mapCtx, in.ConversationContext)
	out.ExtraInfo = in.ExtraInfo
	out.Output = GeneratorSuggestionWithRequired_ToProto(mapCtx, in.Output)
	if in.SummarizationSectionList != nil {
		out.InstructionList = &pb.FewShotExample_SummarizationSectionList{
			SummarizationSectionList: SummarizationSectionList_ToProto(mapCtx, in.SummarizationSectionList),
		}
	}
	return out
}

func FewShotExampleWithRequired_FromProto(mapCtx *direct.MapContext, in *pb.FewShotExample) *krm.FewShotExampleWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.FewShotExampleWithRequired{}
	out.ConversationContext = ConversationContext_FromProto(mapCtx, in.GetConversationContext())
	out.ExtraInfo = in.ExtraInfo
	out.Output = GeneratorSuggestionWithRequired_FromProto(mapCtx, in.GetOutput())
	if oneof := in.GetSummarizationSectionList(); oneof != nil {
		out.SummarizationSectionList = SummarizationSectionList_FromProto(mapCtx, oneof)
	}
	return out
}

func FreeFormContext_ToProto(mapCtx *direct.MapContext, in *krm.FreeFormContext) *pb.FreeFormContext {
	if in == nil {
		return nil
	}
	out := &pb.FreeFormContext{}
	out.Text = direct.ValueOf(in.Text)
	return out
}

func FreeFormContext_FromProto(mapCtx *direct.MapContext, in *pb.FreeFormContext) *krm.FreeFormContext {
	if in == nil {
		return nil
	}
	out := &krm.FreeFormContext{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}

func InferenceParameter_ToProto(mapCtx *direct.MapContext, in *krm.InferenceParameter) *pb.InferenceParameter {
	if in == nil {
		return nil
	}
	out := &pb.InferenceParameter{}
	out.MaxOutputTokens = in.MaxOutputTokens
	out.Temperature = in.Temperature
	out.TopK = in.TopK
	out.TopP = in.TopP
	return out
}

func InferenceParameter_FromProto(mapCtx *direct.MapContext, in *pb.InferenceParameter) *krm.InferenceParameter {
	if in == nil {
		return nil
	}
	out := &krm.InferenceParameter{}
	out.MaxOutputTokens = in.MaxOutputTokens
	out.Temperature = in.Temperature
	out.TopK = in.TopK
	out.TopP = in.TopP
	return out
}

func SummarizationSection_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationSection) *pb.SummarizationSection {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationSection{}
	out.Key = direct.ValueOf(in.Key)
	out.Definition = direct.ValueOf(in.Definition)
	out.Type = direct.Enum_ToProto[pb.SummarizationSection_Type](mapCtx, in.Type)
	return out
}

func SummarizationSection_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationSection) *krm.SummarizationSection {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationSection{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Definition = direct.LazyPtr(in.GetDefinition())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}

func ConversationContext_ToProto(mapCtx *direct.MapContext, in *krm.ConversationContext) *pb.ConversationContext {
	if in == nil {
		return nil
	}
	out := &pb.ConversationContext{}
	out.MessageEntries = direct.Slice_ToProto(mapCtx, in.MessageEntries, MessageEntry_ToProto)
	return out
}

func ConversationContext_FromProto(mapCtx *direct.MapContext, in *pb.ConversationContext) *krm.ConversationContext {
	if in == nil {
		return nil
	}
	out := &krm.ConversationContext{}
	out.MessageEntries = direct.Slice_FromProto(mapCtx, in.MessageEntries, MessageEntry_FromProto)
	return out
}

func MessageEntry_ToProto(mapCtx *direct.MapContext, in *krm.MessageEntry) *pb.MessageEntry {
	if in == nil {
		return nil
	}
	out := &pb.MessageEntry{}
	out.Role = direct.Enum_ToProto[pb.MessageEntry_Role](mapCtx, in.Role)
	out.Text = direct.ValueOf(in.Text)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}

func MessageEntry_FromProto(mapCtx *direct.MapContext, in *pb.MessageEntry) *krm.MessageEntry {
	if in == nil {
		return nil
	}
	out := &krm.MessageEntry{}
	out.Role = direct.Enum_FromProto(mapCtx, in.GetRole())
	out.Text = direct.LazyPtr(in.GetText())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}

func SummarizationSectionList_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationSectionList) *pb.SummarizationSectionList {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationSectionList{}
	out.SummarizationSections = direct.Slice_ToProto(mapCtx, in.SummarizationSections, SummarizationSection_ToProto)
	return out
}

func SummarizationSectionList_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationSectionList) *krm.SummarizationSectionList {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationSectionList{}
	out.SummarizationSections = direct.Slice_FromProto(mapCtx, in.SummarizationSections, SummarizationSection_FromProto)
	return out
}

func EvaluationStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationStatusObservedState) *pb.EvaluationStatus {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationStatus{}
	out.Done = in.Done
	out.PipelineStatus = direct.Status_ToProto(mapCtx, in.PipelineStatus)
	return out
}

func EvaluationStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationStatus) *krm.EvaluationStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationStatusObservedState{}
	out.Done = in.Done
	out.PipelineStatus = direct.Status_FromProto(mapCtx, in.GetPipelineStatus())
	return out
}

func SummarizationEvaluationMetricsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetricsObservedState) *pb.SummarizationEvaluationMetrics {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics{}
	out.SummarizationEvaluationResults = direct.Slice_ToProto(mapCtx, in.SummarizationEvaluationResults, SummarizationEvaluationMetrics_SummarizationEvaluationResultObservedState_ToProto)
	out.SummarizationEvaluationMergedResultsUri = direct.ValueOf(in.SummarizationEvaluationMergedResultsURI)
	out.OverallMetrics = direct.Slice_ToProto(mapCtx, in.OverallMetrics, SummarizationEvaluationMetrics_OverallScoresByMetricObservedState_ToProto)
	out.OverallSectionTokens = direct.Slice_ToProto(mapCtx, in.OverallSectionTokens, SummarizationEvaluationMetrics_SectionTokenObservedState_ToProto)
	out.ConversationDetails = direct.Slice_ToProto(mapCtx, in.ConversationDetails, SummarizationEvaluationMetrics_ConversationDetailObservedState_ToProto)
	return out
}

func SummarizationEvaluationMetricsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics) *krm.SummarizationEvaluationMetricsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetricsObservedState{}
	out.SummarizationEvaluationResults = direct.Slice_FromProto(mapCtx, in.SummarizationEvaluationResults, SummarizationEvaluationMetrics_SummarizationEvaluationResultObservedState_FromProto)
	out.SummarizationEvaluationMergedResultsURI = direct.LazyPtr(in.GetSummarizationEvaluationMergedResultsUri())
	out.OverallMetrics = direct.Slice_FromProto(mapCtx, in.OverallMetrics, SummarizationEvaluationMetrics_OverallScoresByMetricObservedState_FromProto)
	out.OverallSectionTokens = direct.Slice_FromProto(mapCtx, in.OverallSectionTokens, SummarizationEvaluationMetrics_SectionTokenObservedState_FromProto)
	out.ConversationDetails = direct.Slice_FromProto(mapCtx, in.ConversationDetails, SummarizationEvaluationMetrics_ConversationDetailObservedState_FromProto)
	return out
}

func SummarizationEvaluationMetrics_SummarizationEvaluationResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_SummarizationEvaluationResultObservedState) *pb.SummarizationEvaluationMetrics_SummarizationEvaluationResult {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_SummarizationEvaluationResult{}
	out.SessionId = direct.ValueOf(in.SessionID)
	out.Metric = direct.ValueOf(in.Metric)
	out.Section = direct.ValueOf(in.Section)
	out.Score = direct.ValueOf(in.Score)
	out.SectionSummary = direct.ValueOf(in.SectionSummary)
	out.Decompositions = direct.Slice_ToProto(mapCtx, in.Decompositions, SummarizationEvaluationMetrics_DecompositionObservedState_ToProto)
	out.EvaluationResults = direct.Slice_ToProto(mapCtx, in.EvaluationResults, SummarizationEvaluationMetrics_EvaluationResultObservedState_ToProto)
	return out
}

func SummarizationEvaluationMetrics_SummarizationEvaluationResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_SummarizationEvaluationResult) *krm.SummarizationEvaluationMetrics_SummarizationEvaluationResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_SummarizationEvaluationResultObservedState{}
	out.SessionID = direct.LazyPtr(in.GetSessionId())
	out.Metric = direct.LazyPtr(in.GetMetric())
	out.Section = direct.LazyPtr(in.GetSection())
	out.Score = direct.LazyPtr(in.GetScore())
	out.SectionSummary = direct.LazyPtr(in.GetSectionSummary())
	out.Decompositions = direct.Slice_FromProto(mapCtx, in.Decompositions, SummarizationEvaluationMetrics_DecompositionObservedState_FromProto)
	out.EvaluationResults = direct.Slice_FromProto(mapCtx, in.EvaluationResults, SummarizationEvaluationMetrics_EvaluationResultObservedState_FromProto)
	return out
}

func SummarizationEvaluationMetrics_DecompositionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_DecompositionObservedState) *pb.SummarizationEvaluationMetrics_Decomposition {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_Decomposition{}
	if in.AccuracyDecomposition != nil {
		out.Decomposition = &pb.SummarizationEvaluationMetrics_Decomposition_AccuracyDecomposition{
			AccuracyDecomposition: SummarizationEvaluationMetrics_AccuracyDecompositionObservedState_ToProto(mapCtx, in.AccuracyDecomposition),
		}
	}
	if in.AdherenceDecomposition != nil {
		out.Decomposition = &pb.SummarizationEvaluationMetrics_Decomposition_AdherenceDecomposition{
			AdherenceDecomposition: SummarizationEvaluationMetrics_AdherenceDecompositionObservedState_ToProto(mapCtx, in.AdherenceDecomposition),
		}
	}
	return out
}

func SummarizationEvaluationMetrics_DecompositionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_Decomposition) *krm.SummarizationEvaluationMetrics_DecompositionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_DecompositionObservedState{}
	if oneof := in.GetAccuracyDecomposition(); oneof != nil {
		out.AccuracyDecomposition = SummarizationEvaluationMetrics_AccuracyDecompositionObservedState_FromProto(mapCtx, oneof)
	}
	if oneof := in.GetAdherenceDecomposition(); oneof != nil {
		out.AdherenceDecomposition = SummarizationEvaluationMetrics_AdherenceDecompositionObservedState_FromProto(mapCtx, oneof)
	}
	return out
}

func SummarizationEvaluationMetrics_AccuracyDecompositionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_AccuracyDecompositionObservedState) *pb.SummarizationEvaluationMetrics_AccuracyDecomposition {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_AccuracyDecomposition{}
	out.Point = direct.ValueOf(in.Point)
	out.AccuracyReasoning = direct.ValueOf(in.AccuracyReasoning)
	out.IsAccurate = in.IsAccurate
	return out
}

func SummarizationEvaluationMetrics_AccuracyDecompositionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_AccuracyDecomposition) *krm.SummarizationEvaluationMetrics_AccuracyDecompositionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_AccuracyDecompositionObservedState{}
	out.Point = direct.LazyPtr(in.GetPoint())
	out.AccuracyReasoning = direct.LazyPtr(in.GetAccuracyReasoning())
	out.IsAccurate = in.IsAccurate
	return out
}

func SummarizationEvaluationMetrics_AdherenceDecompositionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_AdherenceDecompositionObservedState) *pb.SummarizationEvaluationMetrics_AdherenceDecomposition {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_AdherenceDecomposition{}
	out.Point = direct.ValueOf(in.Point)
	out.AdherenceReasoning = direct.ValueOf(in.AdherenceReasoning)
	out.IsAdherent = direct.ValueOf(in.IsAdherent)
	return out
}

func SummarizationEvaluationMetrics_AdherenceDecompositionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_AdherenceDecomposition) *krm.SummarizationEvaluationMetrics_AdherenceDecompositionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_AdherenceDecompositionObservedState{}
	out.Point = direct.LazyPtr(in.GetPoint())
	out.AdherenceReasoning = direct.LazyPtr(in.GetAdherenceReasoning())
	out.IsAdherent = direct.LazyPtr(in.GetIsAdherent())
	return out
}

func SummarizationEvaluationMetrics_EvaluationResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_EvaluationResultObservedState) *pb.SummarizationEvaluationMetrics_EvaluationResult {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_EvaluationResult{}
	if in.AccuracyDecomposition != nil {
		out.Result = &pb.SummarizationEvaluationMetrics_EvaluationResult_AccuracyDecomposition{
			AccuracyDecomposition: SummarizationEvaluationMetrics_AccuracyDecompositionObservedState_ToProto(mapCtx, in.AccuracyDecomposition),
		}
	}
	if in.AdherenceRubric != nil {
		out.Result = &pb.SummarizationEvaluationMetrics_EvaluationResult_AdherenceRubric{
			AdherenceRubric: SummarizationEvaluationMetrics_AdherenceRubricObservedState_ToProto(mapCtx, in.AdherenceRubric),
		}
	}
	if in.CompletenessRubric != nil {
		out.Result = &pb.SummarizationEvaluationMetrics_EvaluationResult_CompletenessRubric{
			CompletenessRubric: SummarizationEvaluationMetrics_CompletenessRubricObservedState_ToProto(mapCtx, in.CompletenessRubric),
		}
	}
	return out
}

func SummarizationEvaluationMetrics_EvaluationResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_EvaluationResult) *krm.SummarizationEvaluationMetrics_EvaluationResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_EvaluationResultObservedState{}
	if oneof := in.GetAccuracyDecomposition(); oneof != nil {
		out.AccuracyDecomposition = SummarizationEvaluationMetrics_AccuracyDecompositionObservedState_FromProto(mapCtx, oneof)
	}
	if oneof := in.GetAdherenceRubric(); oneof != nil {
		out.AdherenceRubric = SummarizationEvaluationMetrics_AdherenceRubricObservedState_FromProto(mapCtx, oneof)
	}
	if oneof := in.GetCompletenessRubric(); oneof != nil {
		out.CompletenessRubric = SummarizationEvaluationMetrics_CompletenessRubricObservedState_FromProto(mapCtx, oneof)
	}
	return out
}

func SummarizationEvaluationMetrics_AdherenceRubricObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_AdherenceRubricObservedState) *pb.SummarizationEvaluationMetrics_AdherenceRubric {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_AdherenceRubric{}
	out.Question = direct.ValueOf(in.Question)
	out.Reasoning = direct.ValueOf(in.Reasoning)
	out.IsAddressed = in.IsAddressed
	return out
}

func SummarizationEvaluationMetrics_AdherenceRubricObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_AdherenceRubric) *krm.SummarizationEvaluationMetrics_AdherenceRubricObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_AdherenceRubricObservedState{}
	out.Question = direct.LazyPtr(in.GetQuestion())
	out.Reasoning = direct.LazyPtr(in.GetReasoning())
	out.IsAddressed = in.IsAddressed
	return out
}

func SummarizationEvaluationMetrics_CompletenessRubricObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_CompletenessRubricObservedState) *pb.SummarizationEvaluationMetrics_CompletenessRubric {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_CompletenessRubric{}
	out.Question = direct.ValueOf(in.Question)
	out.IsAddressed = in.IsAddressed
	return out
}

func SummarizationEvaluationMetrics_CompletenessRubricObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_CompletenessRubric) *krm.SummarizationEvaluationMetrics_CompletenessRubricObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_CompletenessRubricObservedState{}
	out.Question = direct.LazyPtr(in.GetQuestion())
	out.IsAddressed = in.IsAddressed
	return out
}

func SummarizationEvaluationMetrics_OverallScoresByMetricObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_OverallScoresByMetricObservedState) *pb.SummarizationEvaluationMetrics_OverallScoresByMetric {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_OverallScoresByMetric{}
	out.Metric = direct.ValueOf(in.Metric)
	return out
}

func SummarizationEvaluationMetrics_OverallScoresByMetricObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_OverallScoresByMetric) *krm.SummarizationEvaluationMetrics_OverallScoresByMetricObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_OverallScoresByMetricObservedState{}
	out.Metric = direct.LazyPtr(in.GetMetric())
	return out
}

func SummarizationEvaluationMetrics_SectionTokenObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_SectionTokenObservedState) *pb.SummarizationEvaluationMetrics_SectionToken {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_SectionToken{}
	out.Section = direct.ValueOf(in.Section)
	out.TokenCount = in.TokenCount
	return out
}

func SummarizationEvaluationMetrics_SectionTokenObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_SectionToken) *krm.SummarizationEvaluationMetrics_SectionTokenObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_SectionTokenObservedState{}
	out.Section = direct.LazyPtr(in.GetSection())
	out.TokenCount = in.TokenCount
	return out
}

func SummarizationEvaluationMetrics_ConversationDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_ConversationDetailObservedState) *pb.SummarizationEvaluationMetrics_ConversationDetail {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_ConversationDetail{}
	out.MessageEntries = direct.Slice_ToProto(mapCtx, in.MessageEntries, MessageEntry_ToProto)
	out.SummarySections = direct.Slice_ToProto(mapCtx, in.SummarySections, SummarySuggestion_SummarySectionObservedState_ToProto)
	out.MetricDetails = direct.Slice_ToProto(mapCtx, in.MetricDetails, SummarizationEvaluationMetrics_ConversationDetail_MetricDetailObservedState_ToProto)
	out.SectionTokens = direct.Slice_ToProto(mapCtx, in.SectionTokens, SummarizationEvaluationMetrics_SectionTokenObservedState_ToProto)
	return out
}

func SummarizationEvaluationMetrics_ConversationDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_ConversationDetail) *krm.SummarizationEvaluationMetrics_ConversationDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_ConversationDetailObservedState{}
	out.MessageEntries = direct.Slice_FromProto(mapCtx, in.MessageEntries, MessageEntry_FromProto)
	out.SummarySections = direct.Slice_FromProto(mapCtx, in.SummarySections, SummarySuggestion_SummarySectionObservedState_FromProto)
	out.MetricDetails = direct.Slice_FromProto(mapCtx, in.MetricDetails, SummarizationEvaluationMetrics_ConversationDetail_MetricDetailObservedState_FromProto)
	out.SectionTokens = direct.Slice_FromProto(mapCtx, in.SectionTokens, SummarizationEvaluationMetrics_SectionTokenObservedState_FromProto)
	return out
}

func SummarizationEvaluationMetrics_ConversationDetail_MetricDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_ConversationDetail_MetricDetailObservedState) *pb.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail{}
	out.Metric = direct.ValueOf(in.Metric)
	out.Score = in.Score
	out.SectionDetails = direct.Slice_ToProto(mapCtx, in.SectionDetails, SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetailObservedState_ToProto)
	return out
}

func SummarizationEvaluationMetrics_ConversationDetail_MetricDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail) *krm.SummarizationEvaluationMetrics_ConversationDetail_MetricDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_ConversationDetail_MetricDetailObservedState{}
	out.Metric = direct.LazyPtr(in.GetMetric())
	out.Score = in.Score
	out.SectionDetails = direct.Slice_FromProto(mapCtx, in.SectionDetails, SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetailObservedState_FromProto)
	return out
}

func SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetailObservedState) *pb.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetail{}
	out.Section = direct.ValueOf(in.Section)
	out.Score = in.Score
	out.SectionSummary = direct.ValueOf(in.SectionSummary)
	out.EvaluationResults = direct.Slice_ToProto(mapCtx, in.EvaluationResults, SummarizationEvaluationMetrics_EvaluationResultObservedState_ToProto)
	return out
}

func SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetail) *krm.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationEvaluationMetrics_ConversationDetail_MetricDetail_SectionDetailObservedState{}
	out.Section = direct.LazyPtr(in.GetSection())
	out.Score = in.Score
	out.SectionSummary = direct.LazyPtr(in.GetSectionSummary())
	out.EvaluationResults = direct.Slice_FromProto(mapCtx, in.EvaluationResults, SummarizationEvaluationMetrics_EvaluationResultObservedState_FromProto)
	return out
}

func SummarySuggestion_SummarySectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SummarySuggestion_SummarySectionObservedState) *pb.SummarySuggestion_SummarySection {
	if in == nil {
		return nil
	}
	out := &pb.SummarySuggestion_SummarySection{}
	out.Section = direct.ValueOf(in.Section)
	out.Summary = direct.ValueOf(in.Summary)
	return out
}

func SummarySuggestion_SummarySectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SummarySuggestion_SummarySection) *krm.SummarySuggestion_SummarySectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SummarySuggestion_SummarySectionObservedState{}
	out.Section = direct.LazyPtr(in.GetSection())
	out.Summary = direct.LazyPtr(in.GetSummary())
	return out
}
