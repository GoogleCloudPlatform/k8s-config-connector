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
func Dataset_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.Dataset {
	if in == nil {
		return nil
	}
	out := &krm.Dataset{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SourceLanguageCode = direct.LazyPtr(in.GetSourceLanguageCode())
	out.TargetLanguageCode = direct.LazyPtr(in.GetTargetLanguageCode())
	// MISSING: ExampleCount
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Dataset_ToProto(mapCtx *direct.MapContext, in *krm.Dataset) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SourceLanguageCode = direct.ValueOf(in.SourceLanguageCode)
	out.TargetLanguageCode = direct.ValueOf(in.TargetLanguageCode)
	// MISSING: ExampleCount
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func DatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	out.ExampleCount = direct.LazyPtr(in.GetExampleCount())
	out.TrainExampleCount = direct.LazyPtr(in.GetTrainExampleCount())
	out.ValidateExampleCount = direct.LazyPtr(in.GetValidateExampleCount())
	out.TestExampleCount = direct.LazyPtr(in.GetTestExampleCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func DatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatasetObservedState) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	out.ExampleCount = direct.ValueOf(in.ExampleCount)
	out.TrainExampleCount = direct.ValueOf(in.TrainExampleCount)
	out.ValidateExampleCount = direct.ValueOf(in.ValidateExampleCount)
	out.TestExampleCount = direct.ValueOf(in.TestExampleCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func TranslationDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.TranslationDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationDatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationDatasetObservedState) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.TranslationDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationDatasetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationDatasetSpec) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: ExampleCount
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
