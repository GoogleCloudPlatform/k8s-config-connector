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
// krm.version: v1beta1
// proto.service: google.cloud.speech.v2

package speech

import (
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutoDetectDecodingConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutoDetectDecodingConfig) *krmv1beta1.AutoDetectDecodingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AutoDetectDecodingConfig{}
	return out
}
func AutoDetectDecodingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AutoDetectDecodingConfig) *pb.AutoDetectDecodingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoDetectDecodingConfig{}
	return out
}
func CustomClass_ClassItem_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass_ClassItem) *krmv1beta1.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CustomClass_ClassItem{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func CustomClass_ClassItem_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CustomClass_ClassItem) *pb.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass_ClassItem{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
func ExplicitDecodingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExplicitDecodingConfig) *krmv1beta1.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ExplicitDecodingConfig{}
	/* NOTYET out.Encoding = direct.Enum_FromProto(mapCtx, in.GetEncoding()) */
	/* NOTYET out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz()) */
	/* NOTYET out.AudioChannelCount = direct.LazyPtr(in.GetAudioChannelCount()) */
	return out
}
func ExplicitDecodingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ExplicitDecodingConfig) *pb.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExplicitDecodingConfig{}
	/* NOTYET out.Encoding = direct.Enum_ToProto[pb.ExplicitDecodingConfig_AudioEncoding](mapCtx, in.Encoding) */
	/* NOTYET out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz) */
	/* NOTYET out.AudioChannelCount = direct.ValueOf(in.AudioChannelCount) */
	return out
}
func RecognitionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionConfig) *krmv1beta1.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RecognitionConfig{}
	/* NOTYET out.AutoDecodingConfig = AutoDetectDecodingConfig_FromProto(mapCtx, in.GetAutoDecodingConfig())*/
	/* NOTYET out.ExplicitDecodingConfig = ExplicitDecodingConfig_FromProto(mapCtx, in.GetExplicitDecodingConfig())*/
	out.Model = direct.LazyPtr(in.GetModel())
	out.LanguageCodes = in.LanguageCodes
	/* NOTYET out.Features = RecognitionFeatures_FromProto(mapCtx, in.GetFeatures())*/
	/* NOTYET out.Adaptation = SpeechAdaptation_FromProto(mapCtx, in.GetAdaptation())*/
	/* NOTYET out.TranscriptNormalization = TranscriptNormalization_FromProto(mapCtx, in.GetTranscriptNormalization())*/
	/* NOTYET out.TranslationConfig = TranslationConfig_FromProto(mapCtx, in.GetTranslationConfig())*/
	return out
}
func RecognitionConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RecognitionConfig) *pb.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionConfig{}
	/* NOTYET if oneof := AutoDetectDecodingConfig_ToProto(mapCtx, in.AutoDecodingConfig); oneof != nil {
		out.DecodingConfig = &pb.RecognitionConfig_AutoDecodingConfig{AutoDecodingConfig: oneof}
	}
	if oneof := ExplicitDecodingConfig_ToProto(mapCtx, in.ExplicitDecodingConfig); oneof != nil {
		out.DecodingConfig = &pb.RecognitionConfig_ExplicitDecodingConfig{ExplicitDecodingConfig: oneof}
	} */
	out.Model = direct.ValueOf(in.Model)
	out.LanguageCodes = in.LanguageCodes
	/* NOTYET 	out.Features = RecognitionFeatures_ToProto(mapCtx, in.Features) */
	/* NOTYET out.Adaptation = SpeechAdaptation_ToProto(mapCtx, in.Adaptation) */
	/* NOTYET out.TranscriptNormalization = TranscriptNormalization_ToProto(mapCtx, in.TranscriptNormalization) */
	/* NOTYET out.TranslationConfig = TranslationConfig_ToProto(mapCtx, in.TranslationConfig) */
	return out
}
func RecognitionConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionConfig) *krmv1beta1.RecognitionConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RecognitionConfigObservedState{}
	// MISSING: AutoDecodingConfig
	// MISSING: ExplicitDecodingConfig
	// MISSING: Model
	// MISSING: LanguageCodes
	// MISSING: Features
	// out.Adaptation = SpeechAdaptationObservedState_FromProto(mapCtx, in.GetAdaptation())
	// MISSING: TranscriptNormalization
	// MISSING: TranslationConfig
	return out
}
func RecognitionConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RecognitionConfigObservedState) *pb.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionConfig{}
	// MISSING: AutoDecodingConfig
	// MISSING: ExplicitDecodingConfig
	// MISSING: Model
	// MISSING: LanguageCodes
	// MISSING: Features
	out.Adaptation = SpeechAdaptationObservedState_ToProto(mapCtx, in.Adaptation)
	// MISSING: TranscriptNormalization
	// MISSING: TranslationConfig
	return out
}
func RecognitionFeatures_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionFeatures) *krmv1beta1.RecognitionFeatures {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.RecognitionFeatures{}
	/* NOTYET out.ProfanityFilter = direct.LazyPtr(in.GetProfanityFilter()) */
	/* NOTYET out.EnableWordTimeOffsets = direct.LazyPtr(in.GetEnableWordTimeOffsets()) */
	/* NOTYET out.EnableWordConfidence = direct.LazyPtr(in.GetEnableWordConfidence()) */
	/* NOTYET out.EnableAutomaticPunctuation = direct.LazyPtr(in.GetEnableAutomaticPunctuation()) */
	/* NOTYET out.EnableSpokenPunctuation = direct.LazyPtr(in.GetEnableSpokenPunctuation()) */
	/* NOTYET out.EnableSpokenEmojis = direct.LazyPtr(in.GetEnableSpokenEmojis()) */
	/* NOTYET out.MultiChannelMode = direct.Enum_FromProto(mapCtx, in.GetMultiChannelMode()) */
	/* NOTYET out.DiarizationConfig = SpeakerDiarizationConfig_FromProto(mapCtx, in.GetDiarizationConfig()) */
	/* NOTYET out.MaxAlternatives = direct.LazyPtr(in.GetMaxAlternatives()) */
	return out
}
func RecognitionFeatures_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.RecognitionFeatures) *pb.RecognitionFeatures {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionFeatures{}
	/* NOTYET out.ProfanityFilter = direct.ValueOf(in.ProfanityFilter) */
	/* NOTYET out.EnableWordTimeOffsets = direct.ValueOf(in.EnableWordTimeOffsets) */
	/* NOTYET out.EnableWordConfidence = direct.ValueOf(in.EnableWordConfidence) */
	/* NOTYET out.EnableAutomaticPunctuation = direct.ValueOf(in.EnableAutomaticPunctuation) */
	/* NOTYET out.EnableSpokenPunctuation = direct.ValueOf(in.EnableSpokenPunctuation) */
	/* NOTYET out.EnableSpokenEmojis = direct.ValueOf(in.EnableSpokenEmojis) */
	/* NOTYET out.MultiChannelMode = direct.Enum_ToProto[pb.RecognitionFeatures_MultiChannelMode](mapCtx, in.MultiChannelMode) */
	/* NOTYET out.DiarizationConfig = SpeakerDiarizationConfig_ToProto(mapCtx, in.DiarizationConfig) */
	/* NOTYET out.MaxAlternatives = direct.ValueOf(in.MaxAlternatives) */
	return out
}
func SpeakerDiarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.SpeakerDiarizationConfig) *krmv1beta1.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.SpeakerDiarizationConfig{}
	/* NOTYET out.MinSpeakerCount = direct.LazyPtr(in.GetMinSpeakerCount()) */
	/* NOTYET out.MaxSpeakerCount = direct.LazyPtr(in.GetMaxSpeakerCount()) */
	return out
}
func SpeakerDiarizationConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.SpeakerDiarizationConfig) *pb.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.SpeakerDiarizationConfig{}
	/* NOTYET out.MinSpeakerCount = direct.ValueOf(in.MinSpeakerCount) */
	/* NOTYET out.MaxSpeakerCount = direct.ValueOf(in.MaxSpeakerCount) */
	return out
}
func SpeechCustomClassSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krmv1beta1.SpeechCustomClassSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.SpeechCustomClassSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, CustomClass_ClassItem_FromProto)
	out.Annotations = in.Annotations
	return out
}
func SpeechCustomClassSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.SpeechCustomClassSpec) *pb.CustomClass {
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
func SpeechRecognizerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krmv1beta1.SpeechRecognizerSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.SpeechRecognizerSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Model
	// MISSING: LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfig_FromProto(mapCtx, in.GetDefaultRecognitionConfig())
	out.Annotations = in.Annotations
	return out
}
func SpeechRecognizerSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.SpeechRecognizerSpec) *pb.Recognizer {
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
func TranscriptNormalization_FromProto(mapCtx *direct.MapContext, in *pb.TranscriptNormalization) *krmv1beta1.TranscriptNormalization {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TranscriptNormalization{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, TranscriptNormalization_Entry_FromProto)
	return out
}
func TranscriptNormalization_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TranscriptNormalization) *pb.TranscriptNormalization {
	if in == nil {
		return nil
	}
	out := &pb.TranscriptNormalization{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, TranscriptNormalization_Entry_ToProto)
	return out
}
func TranscriptNormalization_Entry_FromProto(mapCtx *direct.MapContext, in *pb.TranscriptNormalization_Entry) *krmv1beta1.TranscriptNormalization_Entry {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TranscriptNormalization_Entry{}
	/* NOTYET out.Search = direct.LazyPtr(in.GetSearch()) */
	/* NOTYET out.Replace = direct.LazyPtr(in.GetReplace()) */
	/* NOTYET out.CaseSensitive = direct.LazyPtr(in.GetCaseSensitive()) */
	return out
}
func TranscriptNormalization_Entry_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TranscriptNormalization_Entry) *pb.TranscriptNormalization_Entry {
	if in == nil {
		return nil
	}
	out := &pb.TranscriptNormalization_Entry{}
	/* NOTYET out.Search = direct.ValueOf(in.Search) */
	/* NOTYET out.Replace = direct.ValueOf(in.Replace) */
	/* NOTYET out.CaseSensitive = direct.ValueOf(in.CaseSensitive) */
	return out
}
func TranslationConfig_FromProto(mapCtx *direct.MapContext, in *pb.TranslationConfig) *krmv1beta1.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TranslationConfig{}
	/* NOTYET out.TargetLanguage = direct.LazyPtr(in.GetTargetLanguage()) */
	return out
}
func TranslationConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TranslationConfig) *pb.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &pb.TranslationConfig{}
	/* NOTYET out.TargetLanguage = direct.ValueOf(in.TargetLanguage) */
	return out
}
