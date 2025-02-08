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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Model_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.Model {
	if in == nil {
		return nil
	}
	out := &krm.Model{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Dataset = direct.LazyPtr(in.GetDataset())
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Model_ToProto(mapCtx *direct.MapContext, in *krm.Model) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Dataset = direct.ValueOf(in.Dataset)
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.ModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Dataset
	out.SourceLanguageCode = direct.LazyPtr(in.GetSourceLanguageCode())
	out.TargetLanguageCode = direct.LazyPtr(in.GetTargetLanguageCode())
	out.TrainExampleCount = direct.LazyPtr(in.GetTrainExampleCount())
	out.ValidateExampleCount = direct.LazyPtr(in.GetValidateExampleCount())
	out.TestExampleCount = direct.LazyPtr(in.GetTestExampleCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Dataset
	out.SourceLanguageCode = direct.ValueOf(in.SourceLanguageCode)
	out.TargetLanguageCode = direct.ValueOf(in.TargetLanguageCode)
	out.TrainExampleCount = direct.ValueOf(in.TrainExampleCount)
	out.ValidateExampleCount = direct.ValueOf(in.ValidateExampleCount)
	out.TestExampleCount = direct.ValueOf(in.TestExampleCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func TranslationModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.TranslationModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationModelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Dataset
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Dataset
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.TranslationModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationModelSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Dataset
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TranslationModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationModelSpec) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Dataset
	// MISSING: SourceLanguageCode
	// MISSING: TargetLanguageCode
	// MISSING: TrainExampleCount
	// MISSING: ValidateExampleCount
	// MISSING: TestExampleCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
