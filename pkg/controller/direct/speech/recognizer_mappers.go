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
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func InlineCustomClassObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.InlineCustomClassObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InlineCustomClassObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UID = direct.LazyPtr(in.GetUid())
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
func InlineCustomClassObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InlineCustomClassObservedState) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.UID)
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
func InlinePhraseSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PhraseSet) *krm.InlinePhraseSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InlinePhraseSetObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UID = direct.LazyPtr(in.GetUid())
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
func InlinePhraseSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InlinePhraseSetObservedState) *pb.PhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.PhraseSet{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.UID)
	out.State = direct.Enum_ToProto[pb.PhraseSet_State](mapCtx, in.State)
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
func SpeechRecognizerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krm.SpeechRecognizerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechRecognizerObservedState{}
	// MISSING: Name
	out.UID = direct.LazyPtr(in.GetUid())
	out.DefaultRecognitionConfig = RecognitionConfigObservedState_FromProto(mapCtx, in.GetDefaultRecognitionConfig())
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
func SpeechRecognizerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechRecognizerObservedState) *pb.Recognizer {
	if in == nil {
		return nil
	}
	out := &pb.Recognizer{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.UID)
	out.DefaultRecognitionConfig = RecognitionConfigObservedState_ToProto(mapCtx, in.DefaultRecognitionConfig)
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
func SpeechAdaptation_AdaptationPhraseSet_PhraseSet_ToProto(mapCtx *direct.MapContext, in *krm.PhraseSetRef) *pb.SpeechAdaptation_AdaptationPhraseSet_PhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation_AdaptationPhraseSet_PhraseSet{}
	out.PhraseSet = in.External
	return out
}
func InlineCustomClass_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.InlineCustomClass {
	if in == nil {
		return nil
	}
	out := &krm.InlineCustomClass{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, CustomClass_ClassItem_FromProto)
	out.Annotations = in.Annotations
	return out
}
func InlineCustomClass_ToProto(mapCtx *direct.MapContext, in *krm.InlineCustomClass) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, CustomClass_ClassItem_ToProto)
	out.Annotations = in.Annotations
	return out
}
func SpeechAdaptation_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation) *krm.SpeechAdaptation {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptation{}
	out.PhraseSets = direct.Slice_FromProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSet_FromProto)
	out.CustomClasses = direct.Slice_FromProto(mapCtx, in.CustomClasses, InlineCustomClass_FromProto)
	return out
}
func SpeechAdaptation_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptation) *pb.SpeechAdaptation {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation{}
	out.PhraseSets = direct.Slice_ToProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSet_ToProto)
	out.CustomClasses = direct.Slice_ToProto(mapCtx, in.CustomClasses, InlineCustomClass_ToProto)
	return out
}
func SpeechAdaptationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation) *krm.SpeechAdaptationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptationObservedState{}
	out.PhraseSets = direct.Slice_FromProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSetObservedState_FromProto)
	out.CustomClasses = direct.Slice_FromProto(mapCtx, in.CustomClasses, InlineCustomClassObservedState_FromProto)
	return out
}
func SpeechAdaptationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptationObservedState) *pb.SpeechAdaptation {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation{}
	out.PhraseSets = direct.Slice_ToProto(mapCtx, in.PhraseSets, SpeechAdaptation_AdaptationPhraseSetObservedState_ToProto)
	out.CustomClasses = direct.Slice_ToProto(mapCtx, in.CustomClasses, InlineCustomClassObservedState_ToProto)
	return out
}
func InlinePhraseSet_FromProto(mapCtx *direct.MapContext, in *pb.PhraseSet) *krm.InlinePhraseSet {
	if in == nil {
		return nil
	}
	out := &krm.InlinePhraseSet{}
	out.Phrases = direct.Slice_FromProto(mapCtx, in.Phrases, PhraseSet_Phrase_FromProto)
	out.Boost = direct.LazyPtr(direct.Float32ToString(mapCtx, in.GetBoost()))
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Annotations = in.Annotations
	return out
}
func InlinePhraseSet_ToProto(mapCtx *direct.MapContext, in *krm.InlinePhraseSet) *pb.PhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.PhraseSet{}

	out.Phrases = direct.Slice_ToProto(mapCtx, in.Phrases, PhraseSet_Phrase_ToProto)
	out.Boost = direct.StringToFloat32(mapCtx, direct.ValueOf(in.Boost))
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Annotations = in.Annotations
	return out
}
func SpeechAdaptation_AdaptationPhraseSet_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation_AdaptationPhraseSet) *krm.SpeechAdaptation_AdaptationPhraseSet {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptation_AdaptationPhraseSet{}
	if in.GetPhraseSet() != "" {
		out.PhraseSetRef = &krm.PhraseSetRef{
			External: in.GetPhraseSet(),
		}
	}
	out.InlinePhraseSet = InlinePhraseSet_FromProto(mapCtx, in.GetInlinePhraseSet())
	return out
}
func SpeechAdaptation_AdaptationPhraseSet_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptation_AdaptationPhraseSet) *pb.SpeechAdaptation_AdaptationPhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation_AdaptationPhraseSet{}
	if oneof := SpeechAdaptation_AdaptationPhraseSet_PhraseSet_ToProto(mapCtx, in.PhraseSetRef); oneof != nil {
		out.Value = oneof
	}
	if oneof := InlinePhraseSet_ToProto(mapCtx, in.InlinePhraseSet); oneof != nil {
		out.Value = &pb.SpeechAdaptation_AdaptationPhraseSet_InlinePhraseSet{InlinePhraseSet: oneof}
	}
	return out
}
func SpeechAdaptation_AdaptationPhraseSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpeechAdaptation_AdaptationPhraseSet) *krm.SpeechAdaptation_AdaptationPhraseSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechAdaptation_AdaptationPhraseSetObservedState{}
	out.InlinePhraseSet = InlinePhraseSetObservedState_FromProto(mapCtx, in.GetInlinePhraseSet())
	return out
}
func SpeechAdaptation_AdaptationPhraseSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechAdaptation_AdaptationPhraseSetObservedState) *pb.SpeechAdaptation_AdaptationPhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation_AdaptationPhraseSet{}
	if oneof := InlinePhraseSetObservedState_ToProto(mapCtx, in.InlinePhraseSet); oneof != nil {
		out.Value = &pb.SpeechAdaptation_AdaptationPhraseSet_InlinePhraseSet{InlinePhraseSet: oneof}
	}
	return out
}
