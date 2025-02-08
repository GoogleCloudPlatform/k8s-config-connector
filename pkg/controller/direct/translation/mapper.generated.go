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
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AdaptiveMtFile_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtFile) *krm.AdaptiveMtFile {
	if in == nil {
		return nil
	}
	out := &krm.AdaptiveMtFile{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.EntryCount = direct.LazyPtr(in.GetEntryCount())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AdaptiveMtFile_ToProto(mapCtx *direct.MapContext, in *krm.AdaptiveMtFile) *pb.AdaptiveMtFile {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtFile{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.EntryCount = direct.ValueOf(in.EntryCount)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AdaptiveMtFileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtFile) *krm.AdaptiveMtFileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdaptiveMtFileObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EntryCount
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func AdaptiveMtFileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AdaptiveMtFileObservedState) *pb.AdaptiveMtFile {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtFile{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EntryCount
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func TranslationAdaptiveMtFileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtFile) *krm.TranslationAdaptiveMtFileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationAdaptiveMtFileObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EntryCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtFileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationAdaptiveMtFileObservedState) *pb.AdaptiveMtFile {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtFile{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EntryCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtFileSpec_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtFile) *krm.TranslationAdaptiveMtFileSpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationAdaptiveMtFileSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EntryCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtFileSpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationAdaptiveMtFileSpec) *pb.AdaptiveMtFile {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtFile{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EntryCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
