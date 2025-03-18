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

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.BigtableClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableClusterObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func BigtableClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	return out
}
func BigtableClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.BigtableClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableClusterSpec{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ServeNodes = direct.LazyPtr(in.GetServeNodes())
	out.NodeScalingFactor = direct.Enum_FromProto(mapCtx, in.GetNodeScalingFactor())
	out.ClusterConfig = Cluster_ClusterConfig_FromProto(mapCtx, in.GetClusterConfig())
	out.DefaultStorageType = direct.Enum_FromProto(mapCtx, in.GetDefaultStorageType())
	out.EncryptionConfig = Cluster_EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func BigtableClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Location = direct.ValueOf(in.Location)
	out.ServeNodes = direct.ValueOf(in.ServeNodes)
	out.NodeScalingFactor = direct.Enum_ToProto[pb.Cluster_NodeScalingFactor](mapCtx, in.NodeScalingFactor)
	if oneof := Cluster_ClusterConfig_ToProto(mapCtx, in.ClusterConfig); oneof != nil {
		out.Config = &pb.Cluster_ClusterConfig_{ClusterConfig: oneof}
	}
	out.DefaultStorageType = direct.Enum_ToProto[pb.StorageType](mapCtx, in.DefaultStorageType)
	out.EncryptionConfig = Cluster_EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}

func Cluster_ClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ClusterConfig) *krm.Cluster_ClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ClusterConfig{}
	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_FromProto(mapCtx, in.GetClusterAutoscalingConfig())
	return out
}
func Cluster_ClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ClusterConfig) *pb.Cluster_ClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ClusterConfig{}
	out.ClusterAutoscalingConfig = Cluster_ClusterAutoscalingConfig_ToProto(mapCtx, in.ClusterAutoscalingConfig)
	return out
}

func Cluster_EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_EncryptionConfig) *krm.Cluster_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_EncryptionConfig{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func Cluster_EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_EncryptionConfig) *pb.Cluster_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_EncryptionConfig{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}

func Cluster_ClusterAutoscalingConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ClusterAutoscalingConfig) *krm.Cluster_ClusterAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ClusterAutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingLimits_FromProto(mapCtx, in.GetAutoscalingLimits())
	out.AutoscalingTargets = AutoscalingTargets_FromProto(mapCtx, in.GetAutoscalingTargets())
	return out
}
func Cluster_ClusterAutoscalingConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ClusterAutoscalingConfig) *pb.Cluster_ClusterAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ClusterAutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingLimits_ToProto(mapCtx, in.AutoscalingLimits)
	out.AutoscalingTargets = AutoscalingTargets_ToProto(mapCtx, in.AutoscalingTargets)
	return out
}

func AutoscalingLimits_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingLimits) *krm.AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingLimits{}
	out.MinServeNodes = direct.LazyPtr(in.GetMinServeNodes())
	out.MaxServeNodes = direct.LazyPtr(in.GetMaxServeNodes())
	return out
}
func AutoscalingLimits_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingLimits) *pb.AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingLimits{}
	out.MinServeNodes = direct.ValueOf(in.MinServeNodes)
	out.MaxServeNodes = direct.ValueOf(in.MaxServeNodes)
	return out
}
func AutoscalingTargets_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingTargets) *krm.AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingTargets{}
	out.CPUUtilizationPercent = direct.LazyPtr(in.GetCpuUtilizationPercent())
	out.StorageUtilizationGiBPerNode = direct.LazyPtr(in.GetStorageUtilizationGibPerNode())
	return out
}
func AutoscalingTargets_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingTargets) *pb.AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingTargets{}
	out.CpuUtilizationPercent = direct.ValueOf(in.CPUUtilizationPercent)
	out.StorageUtilizationGibPerNode = direct.ValueOf(in.StorageUtilizationGiBPerNode)
	return out
}
