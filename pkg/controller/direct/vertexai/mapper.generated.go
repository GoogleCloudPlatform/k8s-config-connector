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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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

func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{
		KMSKeyRef: &v1beta1.KMSCryptoKeyRef{
			External: in.KmsKeyName,
		},
	}
	return out
}
func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
func MetadataStore_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krm.MetadataStore {
	if in == nil {
		return nil
	}
	out := &krm.MetadataStore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.DataplexConfig = MetadataStore_DataplexConfig_FromProto(mapCtx, in.GetDataplexConfig())
	return out
}
func MetadataStore_ToProto(mapCtx *direct.MapContext, in *krm.MetadataStore) *pb.MetadataStore {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.DataplexConfig = MetadataStore_DataplexConfig_ToProto(mapCtx, in.DataplexConfig)
	return out
}
func MetadataStore_DataplexConfig_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore_DataplexConfig) *krm.MetadataStore_DataplexConfig {
	if in == nil {
		return nil
	}
	out := &krm.MetadataStore_DataplexConfig{}
	out.EnabledPipelinesLineage = direct.LazyPtr(in.GetEnabledPipelinesLineage())
	return out
}
func MetadataStore_DataplexConfig_ToProto(mapCtx *direct.MapContext, in *krm.MetadataStore_DataplexConfig) *pb.MetadataStore_DataplexConfig {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore_DataplexConfig{}
	out.EnabledPipelinesLineage = direct.ValueOf(in.EnabledPipelinesLineage)
	return out
}
func MetadataStore_MetadataStoreState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore_MetadataStoreState) *krm.MetadataStore_MetadataStoreState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataStore_MetadataStoreState{}
	out.DiskUtilizationBytes = direct.LazyPtr(in.GetDiskUtilizationBytes())
	return out
}
func MetadataStore_MetadataStoreState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataStore_MetadataStoreState) *pb.MetadataStore_MetadataStoreState {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore_MetadataStoreState{}
	out.DiskUtilizationBytes = direct.ValueOf(in.DiskUtilizationBytes)
	return out
}
func VertexAIMetadataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krm.VertexAIMetadataStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIMetadataStoreObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = MetadataStore_MetadataStoreState_FromProto(mapCtx, in.GetState())
	return out
}
func VertexAIMetadataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIMetadataStoreObservedState) *pb.MetadataStore {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = MetadataStore_MetadataStoreState_ToProto(mapCtx, in.State)
	return out
}
func VertexAIMetadataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krm.VertexAIMetadataStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIMetadataStoreSpec{}
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DataplexConfig = MetadataStore_DataplexConfig_FromProto(mapCtx, in.GetDataplexConfig())
	return out
}
func VertexAIMetadataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIMetadataStoreSpec) *pb.MetadataStore {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore{}
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.Description = direct.ValueOf(in.Description)
	out.DataplexConfig = MetadataStore_DataplexConfig_ToProto(mapCtx, in.DataplexConfig)
	return out
}
