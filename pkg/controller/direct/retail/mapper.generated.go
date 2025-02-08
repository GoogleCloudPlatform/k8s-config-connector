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

package retail

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2beta/retailpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Model_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.Model {
	if in == nil {
		return nil
	}
	out := &krm.Model{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.TrainingState = direct.Enum_FromProto(mapCtx, in.GetTrainingState())
	// MISSING: ServingState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Type = direct.LazyPtr(in.GetType())
	out.OptimizationObjective = direct.LazyPtr(in.GetOptimizationObjective())
	out.PeriodicTuningState = direct.Enum_FromProto(mapCtx, in.GetPeriodicTuningState())
	// MISSING: LastTuneTime
	// MISSING: TuningOperation
	// MISSING: DataState
	out.FilteringOption = direct.Enum_FromProto(mapCtx, in.GetFilteringOption())
	// MISSING: ServingConfigLists
	out.ModelFeaturesConfig = Model_ModelFeaturesConfig_FromProto(mapCtx, in.GetModelFeaturesConfig())
	return out
}
func Model_ToProto(mapCtx *direct.MapContext, in *krm.Model) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.TrainingState = direct.Enum_ToProto[pb.Model_TrainingState](mapCtx, in.TrainingState)
	// MISSING: ServingState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Type = direct.ValueOf(in.Type)
	out.OptimizationObjective = direct.ValueOf(in.OptimizationObjective)
	out.PeriodicTuningState = direct.Enum_ToProto[pb.Model_PeriodicTuningState](mapCtx, in.PeriodicTuningState)
	// MISSING: LastTuneTime
	// MISSING: TuningOperation
	// MISSING: DataState
	out.FilteringOption = direct.Enum_ToProto[pb.RecommendationsFilteringOption](mapCtx, in.FilteringOption)
	// MISSING: ServingConfigLists
	out.ModelFeaturesConfig = Model_ModelFeaturesConfig_ToProto(mapCtx, in.ModelFeaturesConfig)
	return out
}
func ModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.ModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: TrainingState
	out.ServingState = direct.Enum_FromProto(mapCtx, in.GetServingState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Type
	// MISSING: OptimizationObjective
	// MISSING: PeriodicTuningState
	out.LastTuneTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastTuneTime())
	out.TuningOperation = direct.LazyPtr(in.GetTuningOperation())
	out.DataState = direct.Enum_FromProto(mapCtx, in.GetDataState())
	// MISSING: FilteringOption
	out.ServingConfigLists = direct.Slice_FromProto(mapCtx, in.ServingConfigLists, Model_ServingConfigList_FromProto)
	// MISSING: ModelFeaturesConfig
	return out
}
func ModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: TrainingState
	out.ServingState = direct.Enum_ToProto[pb.Model_ServingState](mapCtx, in.ServingState)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Type
	// MISSING: OptimizationObjective
	// MISSING: PeriodicTuningState
	out.LastTuneTime = direct.StringTimestamp_ToProto(mapCtx, in.LastTuneTime)
	out.TuningOperation = direct.ValueOf(in.TuningOperation)
	out.DataState = direct.Enum_ToProto[pb.Model_DataState](mapCtx, in.DataState)
	// MISSING: FilteringOption
	out.ServingConfigLists = direct.Slice_ToProto(mapCtx, in.ServingConfigLists, Model_ServingConfigList_ToProto)
	// MISSING: ModelFeaturesConfig
	return out
}
func Model_FrequentlyBoughtTogetherFeaturesConfig_FromProto(mapCtx *direct.MapContext, in *pb.Model_FrequentlyBoughtTogetherFeaturesConfig) *krm.Model_FrequentlyBoughtTogetherFeaturesConfig {
	if in == nil {
		return nil
	}
	out := &krm.Model_FrequentlyBoughtTogetherFeaturesConfig{}
	out.ContextProductsType = direct.Enum_FromProto(mapCtx, in.GetContextProductsType())
	return out
}
func Model_FrequentlyBoughtTogetherFeaturesConfig_ToProto(mapCtx *direct.MapContext, in *krm.Model_FrequentlyBoughtTogetherFeaturesConfig) *pb.Model_FrequentlyBoughtTogetherFeaturesConfig {
	if in == nil {
		return nil
	}
	out := &pb.Model_FrequentlyBoughtTogetherFeaturesConfig{}
	out.ContextProductsType = direct.Enum_ToProto[pb.Model_ContextProductsType](mapCtx, in.ContextProductsType)
	return out
}
func Model_ModelFeaturesConfig_FromProto(mapCtx *direct.MapContext, in *pb.Model_ModelFeaturesConfig) *krm.Model_ModelFeaturesConfig {
	if in == nil {
		return nil
	}
	out := &krm.Model_ModelFeaturesConfig{}
	out.FrequentlyBoughtTogetherConfig = Model_FrequentlyBoughtTogetherFeaturesConfig_FromProto(mapCtx, in.GetFrequentlyBoughtTogetherConfig())
	return out
}
func Model_ModelFeaturesConfig_ToProto(mapCtx *direct.MapContext, in *krm.Model_ModelFeaturesConfig) *pb.Model_ModelFeaturesConfig {
	if in == nil {
		return nil
	}
	out := &pb.Model_ModelFeaturesConfig{}
	if oneof := Model_FrequentlyBoughtTogetherFeaturesConfig_ToProto(mapCtx, in.FrequentlyBoughtTogetherConfig); oneof != nil {
		out.TypeDedicatedConfig = &pb.Model_ModelFeaturesConfig_FrequentlyBoughtTogetherConfig{FrequentlyBoughtTogetherConfig: oneof}
	}
	return out
}
func Model_ServingConfigList_FromProto(mapCtx *direct.MapContext, in *pb.Model_ServingConfigList) *krm.Model_ServingConfigList {
	if in == nil {
		return nil
	}
	out := &krm.Model_ServingConfigList{}
	out.ServingConfigIds = in.ServingConfigIds
	return out
}
func Model_ServingConfigList_ToProto(mapCtx *direct.MapContext, in *krm.Model_ServingConfigList) *pb.Model_ServingConfigList {
	if in == nil {
		return nil
	}
	out := &pb.Model_ServingConfigList{}
	out.ServingConfigIds = in.ServingConfigIds
	return out
}
