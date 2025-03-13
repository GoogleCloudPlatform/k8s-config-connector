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

func CustomClassObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.CustomClassObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomClassObservedState{}
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
func CustomClassObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomClassObservedState) *pb.CustomClass {
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
func PhraseSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PhraseSet) *krm.PhraseSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PhraseSetObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UID = direct.LazyPtr(in.GetUid())
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
	out.Uid = direct.ValueOf(in.UID)
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
func SpeechAdaptation_AdaptationPhraseSet_PhraseSet_ToProto(mapCtx *direct.MapContext, in *string) *pb.SpeechAdaptation_AdaptationPhraseSet_PhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.SpeechAdaptation_AdaptationPhraseSet_PhraseSet{}
	out.PhraseSet = direct.ValueOf(in)
	return out
}
