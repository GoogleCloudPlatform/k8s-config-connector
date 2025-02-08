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

package speech

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
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
func CustomClass_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.CustomClass {
	if in == nil {
		return nil
	}
	out := &krm.CustomClass{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, CustomClass_ClassItem_FromProto)
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Annotations = in.Annotations
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func CustomClass_ToProto(mapCtx *direct.MapContext, in *krm.CustomClass) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, CustomClass_ClassItem_ToProto)
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Annotations = in.Annotations
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func CustomClassObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.CustomClassObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomClassObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: DisplayName
	// MISSING: Items
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Annotations
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	return out
}
func CustomClassObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomClassObservedState) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: DisplayName
	// MISSING: Items
	out.State = direct.Enum_ToProto[pb.CustomClass_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Annotations
	out.Etag = direct.ValueOf(in.Etag)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
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
func ExplicitDecodingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExplicitDecodingConfig) *krm.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExplicitDecodingConfig{}
	out.Encoding = direct.Enum_FromProto(mapCtx, in.GetEncoding())
	out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz())
	out.AudioChannelCount = direct.LazyPtr(in.GetAudioChannelCount())
	return out
}
func ExplicitDecodingConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExplicitDecodingConfig) *pb.ExplicitDecodingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExplicitDecodingConfig{}
	out.Encoding = direct.Enum_ToProto[pb.ExplicitDecodingConfig_AudioEncoding](mapCtx, in.Encoding)
	out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz)
	out.AudioChannelCount = direct.ValueOf(in.AudioChannelCount)
	return out
}
func PhraseSet_FromProto(mapCtx *direct.MapContext, in *pb.PhraseSet) *krm.PhraseSet {
	if in == nil {
		return nil
	}
	out := &krm.PhraseSet{}
	// MISSING: Name
	// MISSING: Uid
	out.Phrases = direct.Slice_FromProto(mapCtx, in.Phrases, PhraseSet_Phrase_FromProto)
	out.Boost = direct.LazyPtr(in.GetBoost())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Annotations = in.Annotations
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func PhraseSet_ToProto(mapCtx *direct.MapContext, in *krm.PhraseSet) *pb.PhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.PhraseSet{}
	// MISSING: Name
	// MISSING: Uid
	out.Phrases = direct.Slice_ToProto(mapCtx, in.Phrases, PhraseSet_Phrase_ToProto)
	out.Boost = direct.ValueOf(in.Boost)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Annotations = in.Annotations
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func PhraseSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PhraseSet) *krm.PhraseSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PhraseSetObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Phrases
	// MISSING: Boost
	// MISSING: DisplayName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Annotations
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	return out
}
func PhraseSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PhraseSetObservedState) *pb.PhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.PhraseSet{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Phrases
	// MISSING: Boost
	// MISSING: DisplayName
	out.State = direct.Enum_ToProto[pb.PhraseSet_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Annotations
	out.Etag = direct.ValueOf(in.Etag)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
	return out
}
func PhraseSet_Phrase_FromProto(mapCtx *direct.MapContext, in *pb.PhraseSet_Phrase) *krm.PhraseSet_Phrase {
	if in == nil {
		return nil
	}
	out := &krm.PhraseSet_Phrase{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.Boost = direct.LazyPtr(in.GetBoost())
	return out
}
func PhraseSet_Phrase_ToProto(mapCtx *direct.MapContext, in *krm.PhraseSet_Phrase) *pb.PhraseSet_Phrase {
	if in == nil {
		return nil
	}
	out := &pb.PhraseSet_Phrase{}
	out.Value = direct.ValueOf(in.Value)
	out.Boost = direct.ValueOf(in.Boost)
	return out
}
func RecognitionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionConfig) *krm.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &krm.RecognitionConfig{}
	out.AutoDecodingConfig = AutoDetectDecodingConfig_FromProto(mapCtx, in.GetAutoDecodingConfig())
	out.ExplicitDecodingConfig = ExplicitDecodingConfig_FromProto(mapCtx, in.GetExplicitDecodingConfig())
	out.Model = direct.LazyPtr(in.GetModel())
	out.LanguageCodes = in.LanguageCodes
	out.Features = RecognitionFeatures_FromProto(mapCtx, in.GetFeatures())
	out.Adaptation = SpeechAdaptation_FromProto(mapCtx, in.GetAdaptation())
	out.TranscriptNormalization = TranscriptNormalization_FromProto(mapCtx, in.GetTranscriptNormalization())
	out.TranslationConfig = TranslationConfig_FromProto(mapCtx, in.GetTranslationConfig())
	return out
}
func RecognitionConfig_ToProto(mapCtx *direct.MapContext, in *krm.RecognitionConfig) *pb.RecognitionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecognitionConfig{}
	if oneof := AutoDetectDecodingConfig_ToProto(mapCtx, in.AutoDecodingConfig); oneof != nil {
		out.DecodingConfig = &pb.RecognitionConfig_AutoDecodingConfig{AutoDecodingConfig: oneof}
	}
	if oneof := ExplicitDecodingConfig_ToProto(mapCtx, in.ExplicitDecodingConfig); oneof != nil {
		out.DecodingConfig = &pb.RecognitionConfig_ExplicitDecodingConfig{ExplicitDecodingConfig: oneof}
	}
	out.Model = direct.ValueOf(in.Model)
	out.LanguageCodes = in.LanguageCodes
	out.Features = RecognitionFeatures_ToProto(mapCtx, in.Features)
	out.Adaptation = SpeechAdaptation_ToProto(mapCtx, in.Adaptation)
	out.TranscriptNormalization = TranscriptNormalization_ToProto(mapCtx, in.TranscriptNormalization)
	out.TranslationConfig = TranslationConfig_ToProto(mapCtx, in.TranslationConfig)
	return out
}
func RecognitionConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RecognitionConfig) *krm.RecognitionConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecognitionConfigObservedState{}
	// MISSING: AutoDecodingConfig
	// MISSING: ExplicitDecodingConfig
	// MISSING: Model
	// MISSING: LanguageCodes
	// MISSING: Features
	out.Adaptation = SpeechAdaptationObservedState_FromProto(mapCtx, in.GetAdaptation())
	// MISSING: TranscriptNormalization
	// MISSING: TranslationConfig
	return out
}
func RecognitionConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecognitionConfigObservedState) *pb.RecognitionConfig {
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
func Recognizer_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krm.Recognizer {
	if in == nil {
		return nil
	}
	out := &krm.Recognizer{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Model = direct.LazyPtr(in.GetModel())
	out.LanguageCodes = in.LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfig_FromProto(mapCtx, in.GetDefaultRecognitionConfig())
	out.Annotations = in.Annotations
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func Recognizer_ToProto(mapCtx *direct.MapContext, in *krm.Recognizer) *pb.Recognizer {
	if in == nil {
		return nil
	}
	out := &pb.Recognizer{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Model = direct.ValueOf(in.Model)
	out.LanguageCodes = in.LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfig_ToProto(mapCtx, in.DefaultRecognitionConfig)
	out.Annotations = in.Annotations
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func RecognizerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krm.RecognizerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecognizerObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfigObservedState_FromProto(mapCtx, in.GetDefaultRecognitionConfig())
	// MISSING: Annotations
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	return out
}
func RecognizerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecognizerObservedState) *pb.Recognizer {
	if in == nil {
		return nil
	}
	out := &pb.Recognizer{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: LanguageCodes
	out.DefaultRecognitionConfig = RecognitionConfigObservedState_ToProto(mapCtx, in.DefaultRecognitionConfig)
	// MISSING: Annotations
	out.State = direct.Enum_ToProto[pb.Recognizer_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
	return out
}
func SpeakerDiarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.SpeakerDiarizationConfig) *krm.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.SpeakerDiarizationConfig{}
	out.MinSpeakerCount = direct.LazyPtr(in.GetMinSpeakerCount())
	out.MaxSpeakerCount = direct.LazyPtr(in.GetMaxSpeakerCount())
	return out
}
func SpeakerDiarizationConfig_ToProto(mapCtx *direct.MapContext, in *krm.SpeakerDiarizationConfig) *pb.SpeakerDiarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.SpeakerDiarizationConfig{}
	out.MinSpeakerCount = direct.ValueOf(in.MinSpeakerCount)
	out.MaxSpeakerCount = direct.ValueOf(in.MaxSpeakerCount)
	return out
}
func SpeechAdaptation_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation) *krm.SpeechAdaptation {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptation{}
	out.PhraseSets = direct.Slice_FromProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSet_FromProto)
	out.CustomClasses = direct.Slice_FromProto(mapCtx, in.CustomClasses, CustomClass_FromProto)
	return out
}
func SpeechAdaptation_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptation) *pb.SpeechAdaptation {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation{}
	out.PhraseSets = direct.Slice_ToProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSet_ToProto)
	out.CustomClasses = direct.Slice_ToProto(mapCtx, in.CustomClasses, CustomClass_ToProto)
	return out
}
func SpeechAdaptationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation) *krm.SpeechAdaptationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptationObservedState{}
	out.PhraseSets = direct.Slice_FromProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSetObservedState_FromProto)
	out.CustomClasses = direct.Slice_FromProto(mapCtx, in.CustomClasses, CustomClassObservedState_FromProto)
	return out
}
func SpeechAdaptationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptationObservedState) *pb.SpeechAdaptation {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation{}
	out.PhraseSets = direct.Slice_ToProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSetObservedState_ToProto)
	out.CustomClasses = direct.Slice_ToProto(mapCtx, in.CustomClasses, CustomClassObservedState_ToProto)
	return out
}
func SpeechAdaptation_AdaptationPhraseSet_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation_AdaptationPhraseSet) *krm.SpeechAdaptation_AdaptationPhraseSet {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptation_AdaptationPhraseSet{}
	out.PhraseSet = direct.LazyPtr(in.GetPhraseSet())
	out.InlinePhraseSet = PhraseSet_FromProto(mapCtx, in.GetInlinePhraseSet())
	return out
}
func SpeechAdaptation_AdaptationPhraseSet_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptation_AdaptationPhraseSet) *pb.SpeechAdaptation_AdaptationPhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation_AdaptationPhraseSet{}
	if oneof := SpeechAdaptation_AdaptationPhraseSet_PhraseSet_ToProto(mapCtx, in.PhraseSet); oneof != nil {
		out.Value = oneof
	}
	if oneof := PhraseSet_ToProto(mapCtx, in.InlinePhraseSet); oneof != nil {
		out.Value = &pb.SpeechAdaptation_AdaptationPhraseSet_InlinePhraseSet{InlinePhraseSet: oneof}
	}
	return out
}
func SpeechAdaptation_AdaptationPhraseSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation_AdaptationPhraseSet) *krm.SpeechAdaptation_AdaptationPhraseSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptation_AdaptationPhraseSetObservedState{}
	// MISSING: PhraseSet
	out.InlinePhraseSet = PhraseSetObservedState_FromProto(mapCtx, in.GetInlinePhraseSet())
	return out
}
func SpeechAdaptation_AdaptationPhraseSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptation_AdaptationPhraseSetObservedState) *pb.SpeechAdaptation_AdaptationPhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation_AdaptationPhraseSet{}
	// MISSING: PhraseSet
	if oneof := PhraseSetObservedState_ToProto(mapCtx, in.InlinePhraseSet); oneof != nil {
		out.Value = &pb.SpeechAdaptation_AdaptationPhraseSet_InlinePhraseSet{InlinePhraseSet: oneof}
	}
	return out
}
func SpeechRecognizerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krm.SpeechRecognizerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechRecognizerObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: LanguageCodes
	// MISSING: DefaultRecognitionConfig
	// MISSING: Annotations
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func SpeechRecognizerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechRecognizerObservedState) *pb.Recognizer {
	if in == nil {
		return nil
	}
	out := &pb.Recognizer{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: LanguageCodes
	// MISSING: DefaultRecognitionConfig
	// MISSING: Annotations
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func SpeechRecognizerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krm.SpeechRecognizerSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpeechRecognizerSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: LanguageCodes
	// MISSING: DefaultRecognitionConfig
	// MISSING: Annotations
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	return out
}
func SpeechRecognizerSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpeechRecognizerSpec) *pb.Recognizer {
	if in == nil {
		return nil
	}
	out := &pb.Recognizer{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: LanguageCodes
	// MISSING: DefaultRecognitionConfig
	// MISSING: Annotations
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Etag
	// MISSING: Reconciling
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
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
func TranslationConfig_FromProto(mapCtx *direct.MapContext, in *pb.TranslationConfig) *krm.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &krm.TranslationConfig{}
	out.TargetLanguage = direct.LazyPtr(in.GetTargetLanguage())
	return out
}
func TranslationConfig_ToProto(mapCtx *direct.MapContext, in *krm.TranslationConfig) *pb.TranslationConfig {
	if in == nil {
		return nil
	}
	out := &pb.TranslationConfig{}
	out.TargetLanguage = direct.ValueOf(in.TargetLanguage)
	return out
}
