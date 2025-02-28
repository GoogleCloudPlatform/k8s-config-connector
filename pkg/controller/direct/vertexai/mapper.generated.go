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

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func Featurestore_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.Featurestore {
	if in == nil {
		return nil
	}
	out := &krm.Featurestore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_FromProto(mapCtx, in.GetOnlineServingConfig())
	// MISSING: State
	out.OnlineStorageTtlDays = direct.LazyPtr(in.GetOnlineStorageTtlDays())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func Featurestore_ToProto(mapCtx *direct.MapContext, in *krm.Featurestore) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_ToProto(mapCtx, in.OnlineServingConfig)
	// MISSING: State
	out.OnlineStorageTtlDays = direct.ValueOf(in.OnlineStorageTtlDays)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func FeaturestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.FeaturestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FeaturestoreObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func FeaturestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FeaturestoreObservedState) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	out.State = direct.Enum_ToProto[pb.Featurestore_State](mapCtx, in.State)
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func Featurestore_OnlineServingConfig_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig) *krm.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &krm.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.LazyPtr(in.GetFixedNodeCount())
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_FromProto(mapCtx, in.GetScaling())
	return out
}
func Featurestore_OnlineServingConfig_ToProto(mapCtx *direct.MapContext, in *krm.Featurestore_OnlineServingConfig) *pb.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.ValueOf(in.FixedNodeCount)
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_ToProto(mapCtx, in.Scaling)
	return out
}
func Featurestore_OnlineServingConfig_Scaling_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig_Scaling) *krm.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &krm.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	out.CPUUtilizationTarget = direct.LazyPtr(in.GetCpuUtilizationTarget())
	return out
}
func Featurestore_OnlineServingConfig_Scaling_ToProto(mapCtx *direct.MapContext, in *krm.Featurestore_OnlineServingConfig_Scaling) *pb.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	out.CpuUtilizationTarget = direct.ValueOf(in.CPUUtilizationTarget)
	return out
}
func VertexAIFeaturestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.VertexAIFeaturestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeaturestoreObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeaturestoreObservedState) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.VertexAIFeaturestoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeaturestoreSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeaturestoreSpec) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: OnlineServingConfig
	// MISSING: State
	// MISSING: OnlineStorageTtlDays
	// MISSING: EncryptionSpec
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
