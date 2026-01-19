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

// +generated:mapper
// krm.group: managedkafka.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.managedkafka.v1

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	krmmanagedkafkav1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1beta1"
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
func AccessConfig_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krmmanagedkafkav1beta1.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.AccessConfig{}
	out.NetworkConfigs = direct.Slice_FromProto(mapCtx, in.NetworkConfigs, NetworkConfig_FromProto)
	return out
}
func AccessConfig_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.AccessConfig) *pb.AccessConfig {
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
func CapacityConfig_FromProto(mapCtx *direct.MapContext, in *pb.CapacityConfig) *krmmanagedkafkav1beta1.CapacityConfig {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.CapacityConfig{}
	out.VcpuCount = direct.LazyPtr(in.GetVcpuCount())
	out.MemoryBytes = direct.LazyPtr(in.GetMemoryBytes())
	return out
}
func CapacityConfig_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.CapacityConfig) *pb.CapacityConfig {
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
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
	return out
}
func ManagedKafkaClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedKafkaClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
	return out
}
func ManagedKafkaClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmmanagedkafkav1beta1.ManagedKafkaClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.ManagedKafkaClusterObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
	return out
}
func ManagedKafkaClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.ManagedKafkaClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
	return out
}
func ManagedKafkaClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ManagedKafkaClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.ManagedKafkaClusterSpec{}
	out.GcpConfig = GcpConfig_FromProto(mapCtx, in.GetGcpConfig())
	// MISSING: Name
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_FromProto(mapCtx, in.GetCapacityConfig())
	out.RebalanceConfig = RebalanceConfig_FromProto(mapCtx, in.GetRebalanceConfig())
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
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
	// MISSING: Name
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_ToProto(mapCtx, in.CapacityConfig)
	out.RebalanceConfig = RebalanceConfig_ToProto(mapCtx, in.RebalanceConfig)
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
	return out
}
func ManagedKafkaClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmmanagedkafkav1beta1.ManagedKafkaClusterSpec {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.ManagedKafkaClusterSpec{}
	out.GcpConfig = GcpConfig_FromProto(mapCtx, in.GetGcpConfig())
	// MISSING: Name
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_FromProto(mapCtx, in.GetCapacityConfig())
	out.RebalanceConfig = RebalanceConfig_FromProto(mapCtx, in.GetRebalanceConfig())
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
	return out
}
func ManagedKafkaClusterSpec_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.ManagedKafkaClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	if oneof := GcpConfig_ToProto(mapCtx, in.GcpConfig); oneof != nil {
		out.PlatformConfig = &pb.Cluster_GcpConfig{GcpConfig: oneof}
	}
	// MISSING: Name
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_ToProto(mapCtx, in.CapacityConfig)
	out.RebalanceConfig = RebalanceConfig_ToProto(mapCtx, in.RebalanceConfig)
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: TLSConfig
	return out
}
func ManagedKafkaTopicSpec_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.ManagedKafkaTopicSpec {
	if in == nil {
		return nil
	}
	out := &krm.ManagedKafkaTopicSpec{}
	// MISSING: Name
	out.PartitionCount = direct.LazyPtr(in.GetPartitionCount())
	out.ReplicationFactor = direct.LazyPtr(in.GetReplicationFactor())
	out.Configs = in.Configs
	return out
}
func ManagedKafkaTopicSpec_ToProto(mapCtx *direct.MapContext, in *krm.ManagedKafkaTopicSpec) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	out.PartitionCount = direct.ValueOf(in.PartitionCount)
	out.ReplicationFactor = direct.ValueOf(in.ReplicationFactor)
	out.Configs = in.Configs
	return out
}
func ManagedKafkaTopicSpec_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krmmanagedkafkav1beta1.ManagedKafkaTopicSpec {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.ManagedKafkaTopicSpec{}
	// MISSING: Name
	out.PartitionCount = direct.LazyPtr(in.GetPartitionCount())
	out.ReplicationFactor = direct.LazyPtr(in.GetReplicationFactor())
	out.Configs = in.Configs
	return out
}
func ManagedKafkaTopicSpec_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.ManagedKafkaTopicSpec) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	out.PartitionCount = direct.ValueOf(in.PartitionCount)
	out.ReplicationFactor = direct.ValueOf(in.ReplicationFactor)
	out.Configs = in.Configs
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
func RebalanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.RebalanceConfig) *krmmanagedkafkav1beta1.RebalanceConfig {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.RebalanceConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func RebalanceConfig_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.RebalanceConfig) *pb.RebalanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.RebalanceConfig{}
	out.Mode = direct.Enum_ToProto[pb.RebalanceConfig_Mode](mapCtx, in.Mode)
	return out
}
func TLSConfig_FromProto(mapCtx *direct.MapContext, in *pb.TlsConfig) *krm.TLSConfig {
	if in == nil {
		return nil
	}
	out := &krm.TLSConfig{}
	out.TrustConfig = TrustConfig_FromProto(mapCtx, in.GetTrustConfig())
	out.SSLPrincipalMappingRules = direct.LazyPtr(in.GetSslPrincipalMappingRules())
	return out
}
func TLSConfig_ToProto(mapCtx *direct.MapContext, in *krm.TLSConfig) *pb.TlsConfig {
	if in == nil {
		return nil
	}
	out := &pb.TlsConfig{}
	out.TrustConfig = TrustConfig_ToProto(mapCtx, in.TrustConfig)
	out.SslPrincipalMappingRules = direct.ValueOf(in.SSLPrincipalMappingRules)
	return out
}
func TLSConfig_FromProto(mapCtx *direct.MapContext, in *pb.TlsConfig) *krmmanagedkafkav1beta1.TLSConfig {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.TLSConfig{}
	out.TrustConfig = TrustConfig_FromProto(mapCtx, in.GetTrustConfig())
	out.SSLPrincipalMappingRules = direct.LazyPtr(in.GetSslPrincipalMappingRules())
	return out
}
func TLSConfig_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.TLSConfig) *pb.TlsConfig {
	if in == nil {
		return nil
	}
	out := &pb.TlsConfig{}
	out.TrustConfig = TrustConfig_ToProto(mapCtx, in.TrustConfig)
	out.SslPrincipalMappingRules = direct.ValueOf(in.SSLPrincipalMappingRules)
	return out
}
func TrustConfig_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krm.TrustConfig {
	if in == nil {
		return nil
	}
	out := &krm.TrustConfig{}
	out.CasConfigs = direct.Slice_FromProto(mapCtx, in.CasConfigs, TrustConfig_CertificateAuthorityServiceConfig_FromProto)
	return out
}
func TrustConfig_ToProto(mapCtx *direct.MapContext, in *krm.TrustConfig) *pb.TrustConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig{}
	out.CasConfigs = direct.Slice_ToProto(mapCtx, in.CasConfigs, TrustConfig_CertificateAuthorityServiceConfig_ToProto)
	return out
}
func TrustConfig_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig) *krmmanagedkafkav1beta1.TrustConfig {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.TrustConfig{}
	out.CasConfigs = direct.Slice_FromProto(mapCtx, in.CasConfigs, TrustConfig_CertificateAuthorityServiceConfig_FromProto)
	return out
}
func TrustConfig_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.TrustConfig) *pb.TrustConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig{}
	out.CasConfigs = direct.Slice_ToProto(mapCtx, in.CasConfigs, TrustConfig_CertificateAuthorityServiceConfig_ToProto)
	return out
}
func TrustConfig_CertificateAuthorityServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_CertificateAuthorityServiceConfig) *krm.TrustConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &krm.TrustConfig_CertificateAuthorityServiceConfig{}
	out.CAPool = direct.LazyPtr(in.GetCaPool())
	return out
}
func TrustConfig_CertificateAuthorityServiceConfig_ToProto(mapCtx *direct.MapContext, in *krm.TrustConfig_CertificateAuthorityServiceConfig) *pb.TrustConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_CertificateAuthorityServiceConfig{}
	out.CaPool = direct.ValueOf(in.CAPool)
	return out
}
func TrustConfig_CertificateAuthorityServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.TrustConfig_CertificateAuthorityServiceConfig) *krmmanagedkafkav1beta1.TrustConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1beta1.TrustConfig_CertificateAuthorityServiceConfig{}
	out.CAPool = direct.LazyPtr(in.GetCaPool())
	return out
}
func TrustConfig_CertificateAuthorityServiceConfig_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1beta1.TrustConfig_CertificateAuthorityServiceConfig) *pb.TrustConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrustConfig_CertificateAuthorityServiceConfig{}
	out.CaPool = direct.ValueOf(in.CAPool)
	return out
}
