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
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ConversationContext_FromProto(mapCtx *direct.MapContext, in *pb.ConversationContext) *krm.ConversationContext {
	if in == nil {
		return nil
	}
	out := &krm.ConversationContext{}
	out.MessageEntries = direct.Slice_FromProto(mapCtx, in.MessageEntries, MessageEntry_FromProto)
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
func FewShotExample_FromProto(mapCtx *direct.MapContext, in *pb.FewShotExample) *krm.FewShotExample {
	if in == nil {
		return nil
	}
	out := &krm.FewShotExample{}
	out.ConversationContext = ConversationContext_FromProto(mapCtx, in.GetConversationContext())
	out.ExtraInfo = in.ExtraInfo
	out.SummarizationSectionList = SummarizationSectionList_FromProto(mapCtx, in.GetSummarizationSectionList())
	out.Output = GeneratorSuggestion_FromProto(mapCtx, in.GetOutput())
	return out
}
func FewShotExample_ToProto(mapCtx *direct.MapContext, in *krm.FewShotExample) *pb.FewShotExample {
	if in == nil {
		return nil
	}
	out := &pb.FewShotExample{}
	out.ConversationContext = ConversationContext_ToProto(mapCtx, in.ConversationContext)
	out.ExtraInfo = in.ExtraInfo
	if oneof := SummarizationSectionList_ToProto(mapCtx, in.SummarizationSectionList); oneof != nil {
		out.InstructionList = &pb.FewShotExample_SummarizationSectionList{SummarizationSectionList: oneof}
	}
	out.Output = GeneratorSuggestion_ToProto(mapCtx, in.Output)
	return out
}
func Generator_FromProto(mapCtx *direct.MapContext, in *pb.Generator) *krm.Generator {
	if in == nil {
		return nil
	}
	out := &krm.Generator{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.SummarizationContext = SummarizationContext_FromProto(mapCtx, in.GetSummarizationContext())
	out.InferenceParameter = InferenceParameter_FromProto(mapCtx, in.GetInferenceParameter())
	out.TriggerEvent = direct.Enum_FromProto(mapCtx, in.GetTriggerEvent())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Generator_ToProto(mapCtx *direct.MapContext, in *krm.Generator) *pb.Generator {
	if in == nil {
		return nil
	}
	out := &pb.Generator{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	if oneof := SummarizationContext_ToProto(mapCtx, in.SummarizationContext); oneof != nil {
		out.Context = &pb.Generator_SummarizationContext{SummarizationContext: oneof}
	}
	out.InferenceParameter = InferenceParameter_ToProto(mapCtx, in.InferenceParameter)
	out.TriggerEvent = direct.Enum_ToProto[pb.TriggerEvent](mapCtx, in.TriggerEvent)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func GeneratorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Generator) *krm.GeneratorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: SummarizationContext
	// MISSING: InferenceParameter
	// MISSING: TriggerEvent
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func GeneratorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorObservedState) *pb.Generator {
	if in == nil {
		return nil
	}
	out := &pb.Generator{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: SummarizationContext
	// MISSING: InferenceParameter
	// MISSING: TriggerEvent
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func GeneratorSuggestion_FromProto(mapCtx *direct.MapContext, in *pb.GeneratorSuggestion) *krm.GeneratorSuggestion {
	if in == nil {
		return nil
	}
	out := &krm.GeneratorSuggestion{}
	out.SummarySuggestion = SummarySuggestion_FromProto(mapCtx, in.GetSummarySuggestion())
	return out
}
func GeneratorSuggestion_ToProto(mapCtx *direct.MapContext, in *krm.GeneratorSuggestion) *pb.GeneratorSuggestion {
	if in == nil {
		return nil
	}
	out := &pb.GeneratorSuggestion{}
	if oneof := SummarySuggestion_ToProto(mapCtx, in.SummarySuggestion); oneof != nil {
		out.Suggestion = &pb.GeneratorSuggestion_SummarySuggestion{SummarySuggestion: oneof}
	}
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
func SummarizationContext_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationContext) *krm.SummarizationContext {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationContext{}
	out.SummarizationSections = direct.Slice_FromProto(mapCtx, in.SummarizationSections, SummarizationSection_FromProto)
	out.FewShotExamples = direct.Slice_FromProto(mapCtx, in.FewShotExamples, FewShotExample_FromProto)
	out.Version = direct.LazyPtr(in.GetVersion())
	out.OutputLanguageCode = direct.LazyPtr(in.GetOutputLanguageCode())
	return out
}
func SummarizationContext_ToProto(mapCtx *direct.MapContext, in *krm.SummarizationContext) *pb.SummarizationContext {
	if in == nil {
		return nil
	}
	out := &pb.SummarizationContext{}
	out.SummarizationSections = direct.Slice_ToProto(mapCtx, in.SummarizationSections, SummarizationSection_ToProto)
	out.FewShotExamples = direct.Slice_ToProto(mapCtx, in.FewShotExamples, FewShotExample_ToProto)
	out.Version = direct.ValueOf(in.Version)
	out.OutputLanguageCode = direct.ValueOf(in.OutputLanguageCode)
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
func SummarizationSectionList_FromProto(mapCtx *direct.MapContext, in *pb.SummarizationSectionList) *krm.SummarizationSectionList {
	if in == nil {
		return nil
	}
	out := &krm.SummarizationSectionList{}
	out.SummarizationSections = direct.Slice_FromProto(mapCtx, in.SummarizationSections, SummarizationSection_FromProto)
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
func SummarySuggestion_FromProto(mapCtx *direct.MapContext, in *pb.SummarySuggestion) *krm.SummarySuggestion {
	if in == nil {
		return nil
	}
	out := &krm.SummarySuggestion{}
	out.SummarySections = direct.Slice_FromProto(mapCtx, in.SummarySections, SummarySuggestion_SummarySection_FromProto)
	return out
}
func SummarySuggestion_ToProto(mapCtx *direct.MapContext, in *krm.SummarySuggestion) *pb.SummarySuggestion {
	if in == nil {
		return nil
	}
	out := &pb.SummarySuggestion{}
	out.SummarySections = direct.Slice_ToProto(mapCtx, in.SummarySections, SummarySuggestion_SummarySection_ToProto)
	return out
}
func SummarySuggestion_SummarySection_FromProto(mapCtx *direct.MapContext, in *pb.SummarySuggestion_SummarySection) *krm.SummarySuggestion_SummarySection {
	if in == nil {
		return nil
	}
	out := &krm.SummarySuggestion_SummarySection{}
	out.Section = direct.LazyPtr(in.GetSection())
	out.Summary = direct.LazyPtr(in.GetSummary())
	return out
}
func SummarySuggestion_SummarySection_ToProto(mapCtx *direct.MapContext, in *krm.SummarySuggestion_SummarySection) *pb.SummarySuggestion_SummarySection {
	if in == nil {
		return nil
	}
	out := &pb.SummarySuggestion_SummarySection{}
	out.Section = direct.ValueOf(in.Section)
	out.Summary = direct.ValueOf(in.Summary)
	return out
}
