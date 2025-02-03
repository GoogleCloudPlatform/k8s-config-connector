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

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessConfig_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfig{}
	out.NetworkConfigs = direct.Slice_FromProto(mapCtx, in.NetworkConfigs, NetworkConfig_FromProto)
	return out
}
func AccessConfig_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.NetworkConfigs = direct.Slice_ToProto(mapCtx, in.NetworkConfigs, NetworkConfig_ToProto)
	return out
}
func CapacityConfig_FromProto(mapCtx *direct.MapContext, in *pb.CapacityConfig) *krm.CapacityConfig {
	if in == nil {
		return nil
	}
	out := &krm.CapacityConfig{}
	out.VcpuCount = direct.LazyPtr(in.GetVcpuCount())
	out.MemoryBytes = direct.LazyPtr(in.GetMemoryBytes())
	return out
}
func CapacityConfig_ToProto(mapCtx *direct.MapContext, in *krm.CapacityConfig) *pb.CapacityConfig {
	if in == nil {
		return nil
	}
	out := &pb.CapacityConfig{}
	out.VcpuCount = direct.ValueOf(in.VcpuCount)
	out.MemoryBytes = direct.ValueOf(in.MemoryBytes)
	return out
}
func ManagedKafkaClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ManagedKafkaClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedKafkaClusterObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func ManagedKafkaClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedKafkaClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	return out
}
func ManagedKafkaClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ManagedKafkaClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.ManagedKafkaClusterSpec{}
	out.GcpConfig = GcpConfig_FromProto(mapCtx, in.GetGcpConfig())
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_FromProto(mapCtx, in.GetCapacityConfig())
	out.RebalanceConfig = RebalanceConfig_FromProto(mapCtx, in.GetRebalanceConfig())
	return out
}
func ManagedKafkaClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.ManagedKafkaClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	if oneof := GcpConfig_ToProto(mapCtx, in.GcpConfig); oneof != nil {
		out.PlatformConfig = &pb.Cluster_GcpConfig{GcpConfig: oneof}
	}
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_ToProto(mapCtx, in.CapacityConfig)
	out.RebalanceConfig = RebalanceConfig_ToProto(mapCtx, in.RebalanceConfig)
	return out
}
func RebalanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.RebalanceConfig) *krm.RebalanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.RebalanceConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func RebalanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.RebalanceConfig) *pb.RebalanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.RebalanceConfig{}
	out.Mode = direct.Enum_ToProto[pb.RebalanceConfig_Mode](mapCtx, in.Mode)
	return out
}
