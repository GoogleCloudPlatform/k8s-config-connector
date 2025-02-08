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
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DialogflowGenerativeSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeSettings) *krm.DialogflowGenerativeSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowGenerativeSettingsObservedState{}
	// MISSING: Name
	// MISSING: FallbackSettings
	// MISSING: GenerativeSafetySettings
	// MISSING: KnowledgeConnectorSettings
	// MISSING: LanguageCode
	return out
}
func DialogflowGenerativeSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowGenerativeSettingsObservedState) *pb.GenerativeSettings {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeSettings{}
	// MISSING: Name
	// MISSING: FallbackSettings
	// MISSING: GenerativeSafetySettings
	// MISSING: KnowledgeConnectorSettings
	// MISSING: LanguageCode
	return out
}
func DialogflowGenerativeSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeSettings) *krm.DialogflowGenerativeSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowGenerativeSettingsSpec{}
	// MISSING: Name
	// MISSING: FallbackSettings
	// MISSING: GenerativeSafetySettings
	// MISSING: KnowledgeConnectorSettings
	// MISSING: LanguageCode
	return out
}
func DialogflowGenerativeSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowGenerativeSettingsSpec) *pb.GenerativeSettings {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeSettings{}
	// MISSING: Name
	// MISSING: FallbackSettings
	// MISSING: GenerativeSafetySettings
	// MISSING: KnowledgeConnectorSettings
	// MISSING: LanguageCode
	return out
}
func GenerativeSettings_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeSettings) *krm.GenerativeSettings {
	if in == nil {
		return nil
	}
	out := &krm.GenerativeSettings{}
	out.Name = direct.LazyPtr(in.GetName())
	out.FallbackSettings = GenerativeSettings_FallbackSettings_FromProto(mapCtx, in.GetFallbackSettings())
	out.GenerativeSafetySettings = SafetySettings_FromProto(mapCtx, in.GetGenerativeSafetySettings())
	out.KnowledgeConnectorSettings = GenerativeSettings_KnowledgeConnectorSettings_FromProto(mapCtx, in.GetKnowledgeConnectorSettings())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func GenerativeSettings_ToProto(mapCtx *direct.MapContext, in *krm.GenerativeSettings) *pb.GenerativeSettings {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeSettings{}
	out.Name = direct.ValueOf(in.Name)
	out.FallbackSettings = GenerativeSettings_FallbackSettings_ToProto(mapCtx, in.FallbackSettings)
	out.GenerativeSafetySettings = SafetySettings_ToProto(mapCtx, in.GenerativeSafetySettings)
	out.KnowledgeConnectorSettings = GenerativeSettings_KnowledgeConnectorSettings_ToProto(mapCtx, in.KnowledgeConnectorSettings)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
func GenerativeSettings_FallbackSettings_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeSettings_FallbackSettings) *krm.GenerativeSettings_FallbackSettings {
	if in == nil {
		return nil
	}
	out := &krm.GenerativeSettings_FallbackSettings{}
	out.SelectedPrompt = direct.LazyPtr(in.GetSelectedPrompt())
	out.PromptTemplates = direct.Slice_FromProto(mapCtx, in.PromptTemplates, GenerativeSettings_FallbackSettings_PromptTemplate_FromProto)
	return out
}
func GenerativeSettings_FallbackSettings_ToProto(mapCtx *direct.MapContext, in *krm.GenerativeSettings_FallbackSettings) *pb.GenerativeSettings_FallbackSettings {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeSettings_FallbackSettings{}
	out.SelectedPrompt = direct.ValueOf(in.SelectedPrompt)
	out.PromptTemplates = direct.Slice_ToProto(mapCtx, in.PromptTemplates, GenerativeSettings_FallbackSettings_PromptTemplate_ToProto)
	return out
}
func GenerativeSettings_FallbackSettings_PromptTemplate_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeSettings_FallbackSettings_PromptTemplate) *krm.GenerativeSettings_FallbackSettings_PromptTemplate {
	if in == nil {
		return nil
	}
	out := &krm.GenerativeSettings_FallbackSettings_PromptTemplate{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.PromptText = direct.LazyPtr(in.GetPromptText())
	out.Frozen = direct.LazyPtr(in.GetFrozen())
	return out
}
func GenerativeSettings_FallbackSettings_PromptTemplate_ToProto(mapCtx *direct.MapContext, in *krm.GenerativeSettings_FallbackSettings_PromptTemplate) *pb.GenerativeSettings_FallbackSettings_PromptTemplate {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeSettings_FallbackSettings_PromptTemplate{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.PromptText = direct.ValueOf(in.PromptText)
	out.Frozen = direct.ValueOf(in.Frozen)
	return out
}
func GenerativeSettings_KnowledgeConnectorSettings_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeSettings_KnowledgeConnectorSettings) *krm.GenerativeSettings_KnowledgeConnectorSettings {
	if in == nil {
		return nil
	}
	out := &krm.GenerativeSettings_KnowledgeConnectorSettings{}
	out.Business = direct.LazyPtr(in.GetBusiness())
	out.Agent = direct.LazyPtr(in.GetAgent())
	out.AgentIdentity = direct.LazyPtr(in.GetAgentIdentity())
	out.BusinessDescription = direct.LazyPtr(in.GetBusinessDescription())
	out.AgentScope = direct.LazyPtr(in.GetAgentScope())
	out.DisableDataStoreFallback = direct.LazyPtr(in.GetDisableDataStoreFallback())
	return out
}
func GenerativeSettings_KnowledgeConnectorSettings_ToProto(mapCtx *direct.MapContext, in *krm.GenerativeSettings_KnowledgeConnectorSettings) *pb.GenerativeSettings_KnowledgeConnectorSettings {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeSettings_KnowledgeConnectorSettings{}
	out.Business = direct.ValueOf(in.Business)
	out.Agent = direct.ValueOf(in.Agent)
	out.AgentIdentity = direct.ValueOf(in.AgentIdentity)
	out.BusinessDescription = direct.ValueOf(in.BusinessDescription)
	out.AgentScope = direct.ValueOf(in.AgentScope)
	out.DisableDataStoreFallback = direct.ValueOf(in.DisableDataStoreFallback)
	return out
}
func SafetySettings_FromProto(mapCtx *direct.MapContext, in *pb.SafetySettings) *krm.SafetySettings {
	if in == nil {
		return nil
	}
	out := &krm.SafetySettings{}
	out.BannedPhrases = direct.Slice_FromProto(mapCtx, in.BannedPhrases, SafetySettings_Phrase_FromProto)
	return out
}
func SafetySettings_ToProto(mapCtx *direct.MapContext, in *krm.SafetySettings) *pb.SafetySettings {
	if in == nil {
		return nil
	}
	out := &pb.SafetySettings{}
	out.BannedPhrases = direct.Slice_ToProto(mapCtx, in.BannedPhrases, SafetySettings_Phrase_ToProto)
	return out
}
func SafetySettings_Phrase_FromProto(mapCtx *direct.MapContext, in *pb.SafetySettings_Phrase) *krm.SafetySettings_Phrase {
	if in == nil {
		return nil
	}
	out := &krm.SafetySettings_Phrase{}
	out.Text = direct.LazyPtr(in.GetText())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func SafetySettings_Phrase_ToProto(mapCtx *direct.MapContext, in *krm.SafetySettings_Phrase) *pb.SafetySettings_Phrase {
	if in == nil {
		return nil
	}
	out := &pb.SafetySettings_Phrase{}
	out.Text = direct.ValueOf(in.Text)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
