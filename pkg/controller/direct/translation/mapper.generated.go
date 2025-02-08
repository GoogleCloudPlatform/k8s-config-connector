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
func AdaptiveMtDataset_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtDataset) *krm.AdaptiveMtDataset {
	if in == nil {
		return nil
	}
	out := &krm.AdaptiveMtDataset{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SourceLanguageCode = direct.LazyPtr(in.GetSourceLanguageCode())
	out.TargetLanguageCode = direct.LazyPtr(in.GetTargetLanguageCode())
	out.ExampleCount = direct.LazyPtr(in.GetExampleCount())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AdaptiveMtDataset_ToProto(mapCtx *direct.MapContext, in *krm.AdaptiveMtDataset) *pb.AdaptiveMtDataset {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtDataset{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SourceLanguageCode = direct.ValueOf(in.SourceLanguageCode)
	out.TargetLanguageCode = direct.ValueOf(in.TargetLanguageCode)
	out.ExampleCount = direct.ValueOf(in.ExampleCount)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AdaptiveMtDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtDataset) *krm.AdaptiveMtDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdaptiveMtDatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func AdaptiveMtDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AdaptiveMtDatasetObservedState) *pb.AdaptiveMtDataset {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtDataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func TranslationAdaptiveMtDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtDataset) *krm.TranslationAdaptiveMtDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationAdaptiveMtDatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationAdaptiveMtDatasetObservedState) *pb.AdaptiveMtDataset {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtDataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.AdaptiveMtDataset) *krm.TranslationAdaptiveMtDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationAdaptiveMtDatasetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationAdaptiveMtDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationAdaptiveMtDatasetSpec) *pb.AdaptiveMtDataset {
	if in == nil {
		return nil
	}
	out := &pb.AdaptiveMtDataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
