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

package translation

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
)
func AdaptiveMtSentence_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtSentence) *krm.AdaptiveMtSentence {
	if in == nil {
		return nil
	}
	out := &krm.AdaptiveMtSentence{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SourceSentence = direct.LazyPtr(in.GetSourceSentence())
	out.TargetSentence = direct.LazyPtr(in.GetTargetSentence())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AdaptiveMtSentence_ToProto(mapCtx *direct.MapContext, in *krm.AdaptiveMtSentence) *pb.AdaptiveMtSentence {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtSentence{}
	out.Name = direct.ValueOf(in.Name)
	out.SourceSentence = direct.ValueOf(in.SourceSentence)
	out.TargetSentence = direct.ValueOf(in.TargetSentence)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AdaptiveMtSentenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtSentence) *krm.AdaptiveMtSentenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdaptiveMtSentenceObservedState{}
	// MISSING: Name
	// MISSING: SourceSentence
	// MISSING: TargetSentence
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func AdaptiveMtSentenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AdaptiveMtSentenceObservedState) *pb.AdaptiveMtSentence {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtSentence{}
	// MISSING: Name
	// MISSING: SourceSentence
	// MISSING: TargetSentence
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func TranslationAdaptiveMtSentenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtSentence) *krm.TranslationAdaptiveMtSentenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationAdaptiveMtSentenceObservedState{}
	// MISSING: Name
	// MISSING: SourceSentence
	// MISSING: TargetSentence
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtSentenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationAdaptiveMtSentenceObservedState) *pb.AdaptiveMtSentence {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtSentence{}
	// MISSING: Name
	// MISSING: SourceSentence
	// MISSING: TargetSentence
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtSentenceSpec_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtSentence) *krm.TranslationAdaptiveMtSentenceSpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationAdaptiveMtSentenceSpec{}
	// MISSING: Name
	// MISSING: SourceSentence
	// MISSING: TargetSentence
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtSentenceSpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationAdaptiveMtSentenceSpec) *pb.AdaptiveMtSentence {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtSentence{}
	// MISSING: Name
	// MISSING: SourceSentence
	// MISSING: TargetSentence
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
