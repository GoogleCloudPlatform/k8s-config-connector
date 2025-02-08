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
	pb "cloud.google.com/go/speech/apiv1p1beta1/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func PhraseSet_FromProto(mapCtx *direct.MapContext, in *pb.PhraseSet) *krm.PhraseSet {
	if in == nil {
		return nil
	}
	out := &krm.PhraseSet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Phrases = direct.Slice_FromProto(mapCtx, in.Phrases, PhraseSet_Phrase_FromProto)
	out.Boost = direct.LazyPtr(in.GetBoost())
	return out
}
func PhraseSet_ToProto(mapCtx *direct.MapContext, in *krm.PhraseSet) *pb.PhraseSet {
	if in == nil {
		return nil
	}
	out := &pb.PhraseSet{}
	out.Name = direct.ValueOf(in.Name)
	out.Phrases = direct.Slice_ToProto(mapCtx, in.Phrases, PhraseSet_Phrase_ToProto)
	out.Boost = direct.ValueOf(in.Boost)
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
