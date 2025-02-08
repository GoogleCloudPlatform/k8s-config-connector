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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
)
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AgentVersion = direct.LazyPtr(in.GetAgentVersion())
	// MISSING: State
	// MISSING: UpdateTime
	out.TextToSpeechSettings = TextToSpeechSettings_FromProto(mapCtx, in.GetTextToSpeechSettings())
	out.Fulfillment = Fulfillment_FromProto(mapCtx, in.GetFulfillment())
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.AgentVersion = direct.ValueOf(in.AgentVersion)
	// MISSING: State
	// MISSING: UpdateTime
	out.TextToSpeechSettings = TextToSpeechSettings_ToProto(mapCtx, in.TextToSpeechSettings)
	out.Fulfillment = Fulfillment_ToProto(mapCtx, in.Fulfillment)
	return out
}
func EnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.EnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: AgentVersion
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: TextToSpeechSettings
	// MISSING: Fulfillment
	return out
}
func EnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: AgentVersion
	out.State = direct.Enum_ToProto[pb.Environment_State](mapCtx, in.State)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: TextToSpeechSettings
	// MISSING: Fulfillment
	return out
}
func Fulfillment_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment) *krm.Fulfillment {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GenericWebService = Fulfillment_GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.Features = direct.Slice_FromProto(mapCtx, in.Features, Fulfillment_Feature_FromProto)
	return out
}
func Fulfillment_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment) *pb.Fulfillment {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := Fulfillment_GenericWebService_ToProto(mapCtx, in.GenericWebService); oneof != nil {
		out.Fulfillment = &pb.Fulfillment_GenericWebService_{GenericWebService: oneof}
	}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.Features = direct.Slice_ToProto(mapCtx, in.Features, Fulfillment_Feature_ToProto)
	return out
}
func Fulfillment_Feature_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_Feature) *krm.Fulfillment_Feature {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_Feature{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Fulfillment_Feature_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_Feature) *pb.Fulfillment_Feature {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_Feature{}
	out.Type = direct.Enum_ToProto[pb.Fulfillment_Feature_Type](mapCtx, in.Type)
	return out
}
func Fulfillment_GenericWebService_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_GenericWebService) *krm.Fulfillment_GenericWebService {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_GenericWebService{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.RequestHeaders = in.RequestHeaders
	out.IsCloudFunction = direct.LazyPtr(in.GetIsCloudFunction())
	return out
}
func Fulfillment_GenericWebService_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_GenericWebService) *pb.Fulfillment_GenericWebService {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_GenericWebService{}
	out.Uri = direct.ValueOf(in.URI)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.RequestHeaders = in.RequestHeaders
	out.IsCloudFunction = direct.ValueOf(in.IsCloudFunction)
	return out
}
func SynthesizeSpeechConfig_FromProto(mapCtx *direct.MapContext, in *pb.SynthesizeSpeechConfig) *krm.SynthesizeSpeechConfig {
	if in == nil {
		return nil
	}
	out := &krm.SynthesizeSpeechConfig{}
	out.SpeakingRate = direct.LazyPtr(in.GetSpeakingRate())
	out.Pitch = direct.LazyPtr(in.GetPitch())
	out.VolumeGainDb = direct.LazyPtr(in.GetVolumeGainDb())
	out.EffectsProfileID = in.EffectsProfileId
	out.Voice = VoiceSelectionParams_FromProto(mapCtx, in.GetVoice())
	return out
}
func SynthesizeSpeechConfig_ToProto(mapCtx *direct.MapContext, in *krm.SynthesizeSpeechConfig) *pb.SynthesizeSpeechConfig {
	if in == nil {
		return nil
	}
	out := &pb.SynthesizeSpeechConfig{}
	out.SpeakingRate = direct.ValueOf(in.SpeakingRate)
	out.Pitch = direct.ValueOf(in.Pitch)
	out.VolumeGainDb = direct.ValueOf(in.VolumeGainDb)
	out.EffectsProfileId = in.EffectsProfileID
	out.Voice = VoiceSelectionParams_ToProto(mapCtx, in.Voice)
	return out
}
func TextToSpeechSettings_FromProto(mapCtx *direct.MapContext, in *pb.TextToSpeechSettings) *krm.TextToSpeechSettings {
	if in == nil {
		return nil
	}
	out := &krm.TextToSpeechSettings{}
	out.EnableTextToSpeech = direct.LazyPtr(in.GetEnableTextToSpeech())
	out.OutputAudioEncoding = direct.Enum_FromProto(mapCtx, in.GetOutputAudioEncoding())
	out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz())
	// MISSING: SynthesizeSpeechConfigs
	return out
}
func TextToSpeechSettings_ToProto(mapCtx *direct.MapContext, in *krm.TextToSpeechSettings) *pb.TextToSpeechSettings {
	if in == nil {
		return nil
	}
	out := &pb.TextToSpeechSettings{}
	out.EnableTextToSpeech = direct.ValueOf(in.EnableTextToSpeech)
	out.OutputAudioEncoding = direct.Enum_ToProto[pb.OutputAudioEncoding](mapCtx, in.OutputAudioEncoding)
	out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz)
	// MISSING: SynthesizeSpeechConfigs
	return out
}
func VoiceSelectionParams_FromProto(mapCtx *direct.MapContext, in *pb.VoiceSelectionParams) *krm.VoiceSelectionParams {
	if in == nil {
		return nil
	}
	out := &krm.VoiceSelectionParams{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SsmlGender = direct.Enum_FromProto(mapCtx, in.GetSsmlGender())
	return out
}
func VoiceSelectionParams_ToProto(mapCtx *direct.MapContext, in *krm.VoiceSelectionParams) *pb.VoiceSelectionParams {
	if in == nil {
		return nil
	}
	out := &pb.VoiceSelectionParams{}
	out.Name = direct.ValueOf(in.Name)
	out.SsmlGender = direct.Enum_ToProto[pb.SsmlVoiceGender](mapCtx, in.SsmlGender)
	return out
}
