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

// +generated:mapper
// krm.group: speech.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.speech.v2

package speech

import (
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	krmspeechv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutoDetectDecodingConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutoDetectDecodingConfig) *krm.AutoDetectDecodingConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutoDetectDecodingConfig{}
	return out
}
func AutoDetectDecodingConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutoDetectDecodingConfig) *pb.AutoDetectDecodingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoDetectDecodingConfig{}
	return out
}
func AutoDetectDecodingConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutoDetectDecodingConfig) *krmspeechv1beta1.AutoDetectDecodingConfig {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.AutoDetectDecodingConfig{}
	return out
}
func AutoDetectDecodingConfig_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.AutoDetectDecodingConfig) *pb.AutoDetectDecodingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoDetectDecodingConfig{}
	return out
}
func CustomClass_ClassItem_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass_ClassItem) *krm.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &krm.CustomClass_ClassItem{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func CustomClass_ClassItem_ToProto(mapCtx *direct.MapContext, in *krm.CustomClass_ClassItem) *pb.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass_ClassItem{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
func CustomClass_ClassItem_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass_ClassItem) *krmspeechv1beta1.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.CustomClass_ClassItem{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func CustomClass_ClassItem_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.CustomClass_ClassItem) *pb.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass_ClassItem{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
func DenoiserConfig_FromProto(mapCtx *direct.MapContext, in *pb.DenoiserConfig) *krm.DenoiserConfig {
	if in == nil {
		return nil
	}
	out := &krm.DenoiserConfig{}
	out.DenoiseAudio = direct.LazyPtr(in.GetDenoiseAudio())
	out.SnrThreshold = direct.LazyPtr(in.GetSnrThreshold())
	return out
}
func DenoiserConfig_ToProto(mapCtx *direct.MapContext, in *krm.DenoiserConfig) *pb.DenoiserConfig {
	if in == nil {
		return nil
	}
	out := &pb.DenoiserConfig{}
	out.DenoiseAudio = direct.ValueOf(in.DenoiseAudio)
	out.SnrThreshold = direct.ValueOf(in.SnrThreshold)
	return out
}
func DenoiserConfig_FromProto(mapCtx *direct.MapContext, in *pb.DenoiserConfig) *krmspeechv1beta1.DenoiserConfig {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.DenoiserConfig{}
	out.DenoiseAudio = direct.LazyPtr(in.GetDenoiseAudio())
	out.SnrThreshold = direct.LazyPtr(in.GetSnrThreshold())
	return out
}
func DenoiserConfig_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.DenoiserConfig) *pb.DenoiserConfig {
	if in == nil {
		return nil
	}
	out := &pb.DenoiserConfig{}
	out.DenoiseAudio = direct.ValueOf(in.DenoiseAudio)
	out.SnrThreshold = direct.ValueOf(in.SnrThreshold)
	return out
}
func ExplicitDecodingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExplicitDecodingConfig) *krm.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExplicitDecodingConfig{}
	// MISSING: Encoding
	// MISSING: SampleRateHertz
	// MISSING: AudioChannelCount
	return out
}
func ExplicitDecodingConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExplicitDecodingConfig) *pb.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExplicitDecodingConfig{}
	// MISSING: Encoding
	// MISSING: SampleRateHertz
	// MISSING: AudioChannelCount
	return out
}
func ExplicitDecodingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExplicitDecodingConfig) *krmspeechv1beta1.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.ExplicitDecodingConfig{}
	// MISSING: Encoding
	// MISSING: SampleRateHertz
	// MISSING: AudioChannelCount
	return out
}
func ExplicitDecodingConfig_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.ExplicitDecodingConfig) *pb.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExplicitDecodingConfig{}
	// MISSING: Encoding
	// MISSING: SampleRateHertz
	// MISSING: AudioChannelCount
	return out
}
func RecognitionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionConfig) *krm.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &krm.RecognitionConfig{}
	// MISSING: AutoDecodingConfig
	// MISSING: ExplicitDecodingConfig
	out.Model = direct.LazyPtr(in.GetModel())
	out.LanguageCodes = in.LanguageCodes
	// MISSING: Features
	// MISSING: Adaptation
	// MISSING: TranscriptNormalization
	// MISSING: TranslationConfig
	// MISSING: DenoiserConfig
	return out
}
func RecognitionConfig_ToProto(mapCtx *direct.MapContext, in *krm.RecognitionConfig) *pb.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionConfig{}
	// MISSING: AutoDecodingConfig
	// MISSING: ExplicitDecodingConfig
	out.Model = direct.ValueOf(in.Model)
	out.LanguageCodes = in.LanguageCodes
	// MISSING: Features
	// MISSING: Adaptation
	// MISSING: TranscriptNormalization
	// MISSING: TranslationConfig
	// MISSING: DenoiserConfig
	return out
}
func RecognitionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionConfig) *krmspeechv1beta1.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.RecognitionConfig{}
	// MISSING: AutoDecodingConfig
	// MISSING: ExplicitDecodingConfig
	out.Model = direct.LazyPtr(in.GetModel())
	out.LanguageCodes = in.LanguageCodes
	// MISSING: Features
	// MISSING: Adaptation
	// MISSING: TranscriptNormalization
	// MISSING: TranslationConfig
	// MISSING: DenoiserConfig
	return out
}
func RecognitionConfig_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.RecognitionConfig) *pb.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionConfig{}
	// MISSING: AutoDecodingConfig
	// MISSING: ExplicitDecodingConfig
	out.Model = direct.ValueOf(in.Model)
	out.LanguageCodes = in.LanguageCodes
	// MISSING: Features
	// MISSING: Adaptation
	// MISSING: TranscriptNormalization
	// MISSING: TranslationConfig
	// MISSING: DenoiserConfig
	return out
}
func RecognitionFeatures_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionFeatures) *krm.RecognitionFeatures {
	if in == nil {
		return nil
	}
	out := &krm.RecognitionFeatures{}
	out.ProfanityFilter = direct.LazyPtr(in.GetProfanityFilter())
	out.EnableWordTimeOffsets = direct.LazyPtr(in.GetEnableWordTimeOffsets())
	out.EnableWordConfidence = direct.LazyPtr(in.GetEnableWordConfidence())
	out.EnableAutomaticPunctuation = direct.LazyPtr(in.GetEnableAutomaticPunctuation())
	out.EnableSpokenPunctuation = direct.LazyPtr(in.GetEnableSpokenPunctuation())
	out.EnableSpokenEmojis = direct.LazyPtr(in.GetEnableSpokenEmojis())
	out.MultiChannelMode = direct.Enum_FromProto(mapCtx, in.GetMultiChannelMode())
	out.DiarizationConfig = SpeakerDiarizationConfig_FromProto(mapCtx, in.GetDiarizationConfig())
	out.MaxAlternatives = direct.LazyPtr(in.GetMaxAlternatives())
	return out
}
func RecognitionFeatures_ToProto(mapCtx *direct.MapContext, in *krm.RecognitionFeatures) *pb.RecognitionFeatures {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionFeatures{}
	out.ProfanityFilter = direct.ValueOf(in.ProfanityFilter)
	out.EnableWordTimeOffsets = direct.ValueOf(in.EnableWordTimeOffsets)
	out.EnableWordConfidence = direct.ValueOf(in.EnableWordConfidence)
	out.EnableAutomaticPunctuation = direct.ValueOf(in.EnableAutomaticPunctuation)
	out.EnableSpokenPunctuation = direct.ValueOf(in.EnableSpokenPunctuation)
	out.EnableSpokenEmojis = direct.ValueOf(in.EnableSpokenEmojis)
	out.MultiChannelMode = direct.Enum_ToProto[pb.RecognitionFeatures_MultiChannelMode](mapCtx, in.MultiChannelMode)
	out.DiarizationConfig = SpeakerDiarizationConfig_ToProto(mapCtx, in.DiarizationConfig)
	out.MaxAlternatives = direct.ValueOf(in.MaxAlternatives)
	return out
}
func RecognitionFeatures_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionFeatures) *krmspeechv1beta1.RecognitionFeatures {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.RecognitionFeatures{}
	out.ProfanityFilter = direct.LazyPtr(in.GetProfanityFilter())
	out.EnableWordTimeOffsets = direct.LazyPtr(in.GetEnableWordTimeOffsets())
	out.EnableWordConfidence = direct.LazyPtr(in.GetEnableWordConfidence())
	out.EnableAutomaticPunctuation = direct.LazyPtr(in.GetEnableAutomaticPunctuation())
	out.EnableSpokenPunctuation = direct.LazyPtr(in.GetEnableSpokenPunctuation())
	out.EnableSpokenEmojis = direct.LazyPtr(in.GetEnableSpokenEmojis())
	out.MultiChannelMode = direct.Enum_FromProto(mapCtx, in.GetMultiChannelMode())
	out.DiarizationConfig = SpeakerDiarizationConfig_FromProto(mapCtx, in.GetDiarizationConfig())
	out.MaxAlternatives = direct.LazyPtr(in.GetMaxAlternatives())
	return out
}
func RecognitionFeatures_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.RecognitionFeatures) *pb.RecognitionFeatures {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionFeatures{}
	out.ProfanityFilter = direct.ValueOf(in.ProfanityFilter)
	out.EnableWordTimeOffsets = direct.ValueOf(in.EnableWordTimeOffsets)
	out.EnableWordConfidence = direct.ValueOf(in.EnableWordConfidence)
	out.EnableAutomaticPunctuation = direct.ValueOf(in.EnableAutomaticPunctuation)
	out.EnableSpokenPunctuation = direct.ValueOf(in.EnableSpokenPunctuation)
	out.EnableSpokenEmojis = direct.ValueOf(in.EnableSpokenEmojis)
	out.MultiChannelMode = direct.Enum_ToProto[pb.RecognitionFeatures_MultiChannelMode](mapCtx, in.MultiChannelMode)
	out.DiarizationConfig = SpeakerDiarizationConfig_ToProto(mapCtx, in.DiarizationConfig)
	out.MaxAlternatives = direct.ValueOf(in.MaxAlternatives)
	return out
}
func SpeakerDiarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.SpeakerDiarizationConfig) *krm.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.SpeakerDiarizationConfig{}
	// MISSING: MinSpeakerCount
	// MISSING: MaxSpeakerCount
	return out
}
func SpeakerDiarizationConfig_ToProto(mapCtx *direct.MapContext, in *krm.SpeakerDiarizationConfig) *pb.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.SpeakerDiarizationConfig{}
	// MISSING: MinSpeakerCount
	// MISSING: MaxSpeakerCount
	return out
}
func SpeakerDiarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.SpeakerDiarizationConfig) *krmspeechv1beta1.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.SpeakerDiarizationConfig{}
	// MISSING: MinSpeakerCount
	// MISSING: MaxSpeakerCount
	return out
}
func SpeakerDiarizationConfig_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.SpeakerDiarizationConfig) *pb.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.SpeakerDiarizationConfig{}
	// MISSING: MinSpeakerCount
	// MISSING: MaxSpeakerCount
	return out
}
func SpeechCustomClassSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.SpeechCustomClassSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpeechCustomClassSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, CustomClass_ClassItem_FromProto)
	out.Annotations = in.Annotations
	return out
}
func SpeechCustomClassSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpeechCustomClassSpec) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, CustomClass_ClassItem_ToProto)
	out.Annotations = in.Annotations
	return out
}
func SpeechCustomClassSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krmspeechv1beta1.SpeechCustomClassSpec {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.SpeechCustomClassSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, CustomClass_ClassItem_FromProto)
	out.Annotations = in.Annotations
	return out
}
func SpeechCustomClassSpec_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.SpeechCustomClassSpec) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, CustomClass_ClassItem_ToProto)
	out.Annotations = in.Annotations
	return out
}
func SpeechRecognizerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krm.SpeechRecognizerSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpeechRecognizerSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Model
	// MISSING: LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfig_FromProto(mapCtx, in.GetDefaultRecognitionConfig())
	out.Annotations = in.Annotations
	return out
}
func SpeechRecognizerSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpeechRecognizerSpec) *pb.Recognizer {
	if in == nil {
		return nil
	}
	out := &pb.Recognizer{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Model
	// MISSING: LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfig_ToProto(mapCtx, in.DefaultRecognitionConfig)
	out.Annotations = in.Annotations
	return out
}
func SpeechRecognizerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krmspeechv1beta1.SpeechRecognizerSpec {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.SpeechRecognizerSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Model
	// MISSING: LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfig_FromProto(mapCtx, in.GetDefaultRecognitionConfig())
	out.Annotations = in.Annotations
	return out
}
func SpeechRecognizerSpec_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.SpeechRecognizerSpec) *pb.Recognizer {
	if in == nil {
		return nil
	}
	out := &pb.Recognizer{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Model
	// MISSING: LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfig_ToProto(mapCtx, in.DefaultRecognitionConfig)
	out.Annotations = in.Annotations
	return out
}
func TranscriptNormalization_FromProto(mapCtx *direct.MapContext, in *pb.TranscriptNormalization) *krm.TranscriptNormalization {
	if in == nil {
		return nil
	}
	out := &krm.TranscriptNormalization{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, TranscriptNormalization_Entry_FromProto)
	return out
}
func TranscriptNormalization_ToProto(mapCtx *direct.MapContext, in *krm.TranscriptNormalization) *pb.TranscriptNormalization {
	if in == nil {
		return nil
	}
	out := &pb.TranscriptNormalization{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, TranscriptNormalization_Entry_ToProto)
	return out
}
func TranscriptNormalization_FromProto(mapCtx *direct.MapContext, in *pb.TranscriptNormalization) *krmspeechv1beta1.TranscriptNormalization {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.TranscriptNormalization{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, TranscriptNormalization_Entry_FromProto)
	return out
}
func TranscriptNormalization_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.TranscriptNormalization) *pb.TranscriptNormalization {
	if in == nil {
		return nil
	}
	out := &pb.TranscriptNormalization{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, TranscriptNormalization_Entry_ToProto)
	return out
}
func TranscriptNormalization_Entry_FromProto(mapCtx *direct.MapContext, in *pb.TranscriptNormalization_Entry) *krm.TranscriptNormalization_Entry {
	if in == nil {
		return nil
	}
	out := &krm.TranscriptNormalization_Entry{}
	out.Search = direct.LazyPtr(in.GetSearch())
	out.Replace = direct.LazyPtr(in.GetReplace())
	out.CaseSensitive = direct.LazyPtr(in.GetCaseSensitive())
	return out
}
func TranscriptNormalization_Entry_ToProto(mapCtx *direct.MapContext, in *krm.TranscriptNormalization_Entry) *pb.TranscriptNormalization_Entry {
	if in == nil {
		return nil
	}
	out := &pb.TranscriptNormalization_Entry{}
	out.Search = direct.ValueOf(in.Search)
	out.Replace = direct.ValueOf(in.Replace)
	out.CaseSensitive = direct.ValueOf(in.CaseSensitive)
	return out
}
func TranscriptNormalization_Entry_FromProto(mapCtx *direct.MapContext, in *pb.TranscriptNormalization_Entry) *krmspeechv1beta1.TranscriptNormalization_Entry {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.TranscriptNormalization_Entry{}
	out.Search = direct.LazyPtr(in.GetSearch())
	out.Replace = direct.LazyPtr(in.GetReplace())
	out.CaseSensitive = direct.LazyPtr(in.GetCaseSensitive())
	return out
}
func TranscriptNormalization_Entry_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.TranscriptNormalization_Entry) *pb.TranscriptNormalization_Entry {
	if in == nil {
		return nil
	}
	out := &pb.TranscriptNormalization_Entry{}
	out.Search = direct.ValueOf(in.Search)
	out.Replace = direct.ValueOf(in.Replace)
	out.CaseSensitive = direct.ValueOf(in.CaseSensitive)
	return out
}
func TranslationConfig_FromProto(mapCtx *direct.MapContext, in *pb.TranslationConfig) *krm.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &krm.TranslationConfig{}
	// MISSING: TargetLanguage
	return out
}
func TranslationConfig_ToProto(mapCtx *direct.MapContext, in *krm.TranslationConfig) *pb.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &pb.TranslationConfig{}
	// MISSING: TargetLanguage
	return out
}
func TranslationConfig_FromProto(mapCtx *direct.MapContext, in *pb.TranslationConfig) *krmspeechv1beta1.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &krmspeechv1beta1.TranslationConfig{}
	// MISSING: TargetLanguage
	return out
}
func TranslationConfig_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.TranslationConfig) *pb.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &pb.TranslationConfig{}
	// MISSING: TargetLanguage
	return out
}
