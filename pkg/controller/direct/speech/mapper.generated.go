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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
)
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
