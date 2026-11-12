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

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DialogflowCXAgentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Agent) *krm.DialogflowCXAgentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXAgentObservedState{}
	return out
}

func DialogflowCXAgentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXAgentObservedState) *pb.Agent {
	if in == nil {
		return nil
	}
	out := &pb.Agent{}
	return out
}

func DialogflowCXAgentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Agent) *krm.DialogflowCXAgentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXAgentSpec{}
	out.DisplayName = in.GetDisplayName()
	out.DefaultLanguageCode = in.GetDefaultLanguageCode()
	out.SupportedLanguageCodes = in.SupportedLanguageCodes
	out.TimeZone = in.GetTimeZone()
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AvatarURI = direct.LazyPtr(in.GetAvatarUri())
	out.SpeechToTextSettings = AgentSpeechToTextSettings_FromProto(mapCtx, in.GetSpeechToTextSettings())
	out.SecuritySettings = direct.LazyPtr(in.GetSecuritySettings())
	out.EnableStackdriverLogging = direct.LazyPtr(in.GetEnableStackdriverLogging())
	out.EnableSpellCorrection = direct.LazyPtr(in.GetEnableSpellCorrection())
	return out
}

func DialogflowCXAgentSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXAgentSpec) *pb.Agent {
	if in == nil {
		return nil
	}
	out := &pb.Agent{}
	out.DisplayName = in.DisplayName
	out.DefaultLanguageCode = in.DefaultLanguageCode
	out.SupportedLanguageCodes = in.SupportedLanguageCodes
	out.TimeZone = in.TimeZone
	out.Description = direct.ValueOf(in.Description)
	out.AvatarUri = direct.ValueOf(in.AvatarURI)
	out.SpeechToTextSettings = AgentSpeechToTextSettings_ToProto(mapCtx, in.SpeechToTextSettings)
	out.SecuritySettings = direct.ValueOf(in.SecuritySettings)
	out.EnableStackdriverLogging = direct.ValueOf(in.EnableStackdriverLogging)
	out.EnableSpellCorrection = direct.ValueOf(in.EnableSpellCorrection)
	return out
}

func AgentSpeechToTextSettings_FromProto(mapCtx *direct.MapContext, in *pb.SpeechToTextSettings) *krm.AgentSpeechToTextSettings {
	if in == nil {
		return nil
	}
	out := &krm.AgentSpeechToTextSettings{}
	out.EnableSpeechAdaptation = direct.LazyPtr(in.GetEnableSpeechAdaptation())
	return out
}

func AgentSpeechToTextSettings_ToProto(mapCtx *direct.MapContext, in *krm.AgentSpeechToTextSettings) *pb.SpeechToTextSettings {
	if in == nil {
		return nil
	}
	out := &pb.SpeechToTextSettings{}
	out.EnableSpeechAdaptation = direct.ValueOf(in.EnableSpeechAdaptation)
	return out
}
