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

package enterpriseknowledgegraph

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/enterpriseknowledgegraph/apiv1/enterpriseknowledgegraphpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/enterpriseknowledgegraph/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AffinityClusteringConfig_FromProto(mapCtx *direct.MapContext, in *pb.AffinityClusteringConfig) *krm.AffinityClusteringConfig {
	if in == nil {
		return nil
	}
	out := &krm.AffinityClusteringConfig{}
	out.CompressionRoundCount = direct.LazyPtr(in.GetCompressionRoundCount())
	return out
}
func AffinityClusteringConfig_ToProto(mapCtx *direct.MapContext, in *krm.AffinityClusteringConfig) *pb.AffinityClusteringConfig {
	if in == nil {
		return nil
	}
	out := &pb.AffinityClusteringConfig{}
	out.CompressionRoundCount = direct.ValueOf(in.CompressionRoundCount)
	return out
}
func BigQueryInputConfig_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryInputConfig) *krm.BigQueryInputConfig {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryInputConfig{}
	out.BigqueryTable = direct.LazyPtr(in.GetBigqueryTable())
	out.GcsURI = direct.LazyPtr(in.GetGcsUri())
	return out
}
func BigQueryInputConfig_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryInputConfig) *pb.BigQueryInputConfig {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryInputConfig{}
	out.BigqueryTable = direct.ValueOf(in.BigqueryTable)
	out.GcsUri = direct.ValueOf(in.GcsURI)
	return out
}
func ConnectedComponentsConfig_FromProto(mapCtx *direct.MapContext, in *pb.ConnectedComponentsConfig) *krm.ConnectedComponentsConfig {
	if in == nil {
		return nil
	}
	out := &krm.ConnectedComponentsConfig{}
	out.WeightThreshold = direct.LazyPtr(in.GetWeightThreshold())
	return out
}
func ConnectedComponentsConfig_ToProto(mapCtx *direct.MapContext, in *krm.ConnectedComponentsConfig) *pb.ConnectedComponentsConfig {
	if in == nil {
		return nil
	}
	out := &pb.ConnectedComponentsConfig{}
	out.WeightThreshold = direct.ValueOf(in.WeightThreshold)
	return out
}
func EnterpriseknowledgegraphEntityReconciliationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntityReconciliationJob) *krm.EnterpriseknowledgegraphEntityReconciliationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnterpriseknowledgegraphEntityReconciliationJobObservedState{}
	// MISSING: Name
	// MISSING: InputConfig
	// MISSING: OutputConfig
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: ReconConfig
	return out
}
func EnterpriseknowledgegraphEntityReconciliationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EnterpriseknowledgegraphEntityReconciliationJobObservedState) *pb.EntityReconciliationJob {
	if in == nil {
		return nil
	}
	out := &pb.EntityReconciliationJob{}
	// MISSING: Name
	// MISSING: InputConfig
	// MISSING: OutputConfig
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: ReconConfig
	return out
}
func EnterpriseknowledgegraphEntityReconciliationJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntityReconciliationJob) *krm.EnterpriseknowledgegraphEntityReconciliationJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.EnterpriseknowledgegraphEntityReconciliationJobSpec{}
	// MISSING: Name
	// MISSING: InputConfig
	// MISSING: OutputConfig
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: ReconConfig
	return out
}
func EnterpriseknowledgegraphEntityReconciliationJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.EnterpriseknowledgegraphEntityReconciliationJobSpec) *pb.EntityReconciliationJob {
	if in == nil {
		return nil
	}
	out := &pb.EntityReconciliationJob{}
	// MISSING: Name
	// MISSING: InputConfig
	// MISSING: OutputConfig
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: ReconConfig
	return out
}
func EntityReconciliationJob_FromProto(mapCtx *direct.MapContext, in *pb.EntityReconciliationJob) *krm.EntityReconciliationJob {
	if in == nil {
		return nil
	}
	out := &krm.EntityReconciliationJob{}
	// MISSING: Name
	out.InputConfig = InputConfig_FromProto(mapCtx, in.GetInputConfig())
	out.OutputConfig = OutputConfig_FromProto(mapCtx, in.GetOutputConfig())
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	out.ReconConfig = ReconConfig_FromProto(mapCtx, in.GetReconConfig())
	return out
}
func EntityReconciliationJob_ToProto(mapCtx *direct.MapContext, in *krm.EntityReconciliationJob) *pb.EntityReconciliationJob {
	if in == nil {
		return nil
	}
	out := &pb.EntityReconciliationJob{}
	// MISSING: Name
	out.InputConfig = InputConfig_ToProto(mapCtx, in.InputConfig)
	out.OutputConfig = OutputConfig_ToProto(mapCtx, in.OutputConfig)
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	out.ReconConfig = ReconConfig_ToProto(mapCtx, in.ReconConfig)
	return out
}
func EntityReconciliationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntityReconciliationJob) *krm.EntityReconciliationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntityReconciliationJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: InputConfig
	// MISSING: OutputConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ReconConfig
	return out
}
func EntityReconciliationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntityReconciliationJobObservedState) *pb.EntityReconciliationJob {
	if in == nil {
		return nil
	}
	out := &pb.EntityReconciliationJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: InputConfig
	// MISSING: OutputConfig
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ReconConfig
	return out
}
func InputConfig_FromProto(mapCtx *direct.MapContext, in *pb.InputConfig) *krm.InputConfig {
	if in == nil {
		return nil
	}
	out := &krm.InputConfig{}
	out.BigqueryInputConfigs = direct.Slice_FromProto(mapCtx, in.BigqueryInputConfigs, BigQueryInputConfig_FromProto)
	out.EntityType = direct.Enum_FromProto(mapCtx, in.GetEntityType())
	out.PreviousResultBigqueryTable = direct.LazyPtr(in.GetPreviousResultBigqueryTable())
	return out
}
func InputConfig_ToProto(mapCtx *direct.MapContext, in *krm.InputConfig) *pb.InputConfig {
	if in == nil {
		return nil
	}
	out := &pb.InputConfig{}
	out.BigqueryInputConfigs = direct.Slice_ToProto(mapCtx, in.BigqueryInputConfigs, BigQueryInputConfig_ToProto)
	out.EntityType = direct.Enum_ToProto[pb.InputConfig_EntityType](mapCtx, in.EntityType)
	out.PreviousResultBigqueryTable = direct.ValueOf(in.PreviousResultBigqueryTable)
	return out
}
func OutputConfig_FromProto(mapCtx *direct.MapContext, in *pb.OutputConfig) *krm.OutputConfig {
	if in == nil {
		return nil
	}
	out := &krm.OutputConfig{}
	out.BigqueryDataset = direct.LazyPtr(in.GetBigqueryDataset())
	return out
}
func OutputConfig_ToProto(mapCtx *direct.MapContext, in *krm.OutputConfig) *pb.OutputConfig {
	if in == nil {
		return nil
	}
	out := &pb.OutputConfig{}
	out.BigqueryDataset = direct.ValueOf(in.BigqueryDataset)
	return out
}
func ReconConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReconConfig) *krm.ReconConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReconConfig{}
	out.ConnectedComponentsConfig = ConnectedComponentsConfig_FromProto(mapCtx, in.GetConnectedComponentsConfig())
	out.AffinityClusteringConfig = AffinityClusteringConfig_FromProto(mapCtx, in.GetAffinityClusteringConfig())
	out.Options = ReconConfig_Options_FromProto(mapCtx, in.GetOptions())
	out.ModelConfig = ReconConfig_ModelConfig_FromProto(mapCtx, in.GetModelConfig())
	return out
}
func ReconConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReconConfig) *pb.ReconConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReconConfig{}
	if oneof := ConnectedComponentsConfig_ToProto(mapCtx, in.ConnectedComponentsConfig); oneof != nil {
		out.ClusteringConfig = &pb.ReconConfig_ConnectedComponentsConfig{ConnectedComponentsConfig: oneof}
	}
	if oneof := AffinityClusteringConfig_ToProto(mapCtx, in.AffinityClusteringConfig); oneof != nil {
		out.ClusteringConfig = &pb.ReconConfig_AffinityClusteringConfig{AffinityClusteringConfig: oneof}
	}
	out.Options = ReconConfig_Options_ToProto(mapCtx, in.Options)
	out.ModelConfig = ReconConfig_ModelConfig_ToProto(mapCtx, in.ModelConfig)
	return out
}
func ReconConfig_ModelConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReconConfig_ModelConfig) *krm.ReconConfig_ModelConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReconConfig_ModelConfig{}
	out.ModelName = direct.LazyPtr(in.GetModelName())
	out.VersionTag = direct.LazyPtr(in.GetVersionTag())
	return out
}
func ReconConfig_ModelConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReconConfig_ModelConfig) *pb.ReconConfig_ModelConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReconConfig_ModelConfig{}
	out.ModelName = direct.ValueOf(in.ModelName)
	out.VersionTag = direct.ValueOf(in.VersionTag)
	return out
}
func ReconConfig_Options_FromProto(mapCtx *direct.MapContext, in *pb.ReconConfig_Options) *krm.ReconConfig_Options {
	if in == nil {
		return nil
	}
	out := &krm.ReconConfig_Options{}
	out.EnableGeocodingSeparation = direct.LazyPtr(in.GetEnableGeocodingSeparation())
	return out
}
func ReconConfig_Options_ToProto(mapCtx *direct.MapContext, in *krm.ReconConfig_Options) *pb.ReconConfig_Options {
	if in == nil {
		return nil
	}
	out := &pb.ReconConfig_Options{}
	out.EnableGeocodingSeparation = direct.ValueOf(in.EnableGeocodingSeparation)
	return out
}
