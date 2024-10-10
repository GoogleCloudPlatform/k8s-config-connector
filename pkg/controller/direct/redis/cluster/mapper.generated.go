// Copyright 2024 Google LLC
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

package cluster

import (
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ClusterPersistenceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ClusterPersistenceConfig) *krm.ClusterPersistenceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterPersistenceConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.RdbConfig = ClusterPersistenceConfig_RDBConfig_FromProto(mapCtx, in.GetRdbConfig())
	out.AofConfig = ClusterPersistenceConfig_AOFConfig_FromProto(mapCtx, in.GetAofConfig())
	return out
}
func ClusterPersistenceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ClusterPersistenceConfig) *pb.ClusterPersistenceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ClusterPersistenceConfig{}
	out.Mode = direct.Enum_ToProto[pb.ClusterPersistenceConfig_PersistenceMode](mapCtx, in.Mode)
	out.RdbConfig = ClusterPersistenceConfig_RDBConfig_ToProto(mapCtx, in.RdbConfig)
	out.AofConfig = ClusterPersistenceConfig_AOFConfig_ToProto(mapCtx, in.AofConfig)
	return out
}
func ClusterPersistenceConfig_AOFConfig_FromProto(mapCtx *direct.MapContext, in *pb.ClusterPersistenceConfig_AOFConfig) *krm.ClusterPersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterPersistenceConfig_AOFConfig{}
	out.AppendFsync = direct.Enum_FromProto(mapCtx, in.GetAppendFsync())
	return out
}
func ClusterPersistenceConfig_AOFConfig_ToProto(mapCtx *direct.MapContext, in *krm.ClusterPersistenceConfig_AOFConfig) *pb.ClusterPersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := &pb.ClusterPersistenceConfig_AOFConfig{}
	out.AppendFsync = direct.Enum_ToProto[pb.ClusterPersistenceConfig_AOFConfig_AppendFsync](mapCtx, in.AppendFsync)
	return out
}
func ClusterPersistenceConfig_RDBConfig_FromProto(mapCtx *direct.MapContext, in *pb.ClusterPersistenceConfig_RDBConfig) *krm.ClusterPersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterPersistenceConfig_RDBConfig{}
	out.RdbSnapshotPeriod = direct.Enum_FromProto(mapCtx, in.GetRdbSnapshotPeriod())
	out.RdbSnapshotStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRdbSnapshotStartTime())
	return out
}
func ClusterPersistenceConfig_RDBConfig_ToProto(mapCtx *direct.MapContext, in *krm.ClusterPersistenceConfig_RDBConfig) *pb.ClusterPersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := &pb.ClusterPersistenceConfig_RDBConfig{}
	out.RdbSnapshotPeriod = direct.Enum_ToProto[pb.ClusterPersistenceConfig_RDBConfig_SnapshotPeriod](mapCtx, in.RdbSnapshotPeriod)
	out.RdbSnapshotStartTime = direct.StringTimestamp_ToProto(mapCtx, in.RdbSnapshotStartTime)
	return out
}
func Cluster_StateInfo_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_StateInfo) *krm.Cluster_StateInfo {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_StateInfo{}
	out.UpdateInfo = Cluster_StateInfo_UpdateInfo_FromProto(mapCtx, in.GetUpdateInfo())
	return out
}
func Cluster_StateInfo_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_StateInfo) *pb.Cluster_StateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_StateInfo{}
	if oneof := Cluster_StateInfo_UpdateInfo_ToProto(mapCtx, in.UpdateInfo); oneof != nil {
		out.Info = &pb.Cluster_StateInfo_UpdateInfo_{UpdateInfo: oneof}
	}
	return out
}
func Cluster_StateInfo_UpdateInfo_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_StateInfo_UpdateInfo) *krm.Cluster_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_StateInfo_UpdateInfo{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func Cluster_StateInfo_UpdateInfo_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_StateInfo_UpdateInfo) *pb.Cluster_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_StateInfo_UpdateInfo{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func DiscoveryEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveryEndpoint) *krm.DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEndpoint{}
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	out.PscConfig = PscConfig_FromProto(mapCtx, in.GetPscConfig())
	return out
}
func DiscoveryEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEndpoint) *pb.DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveryEndpoint{}
	out.Address = direct.ValueOf(in.Address)
	out.Port = direct.ValueOf(in.Port)
	out.PscConfig = PscConfig_ToProto(mapCtx, in.PscConfig)
	return out
}
func PscConfig_FromProto(mapCtx *direct.MapContext, in *pb.PscConfig) *krm.PscConfig {
	if in == nil {
		return nil
	}
	out := &krm.PscConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func PscConfig_ToProto(mapCtx *direct.MapContext, in *krm.PscConfig) *pb.PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscConfig{}
	out.Network = direct.ValueOf(in.Network)
	return out
}
func PscConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krm.PscConnection {
	if in == nil {
		return nil
	}
	out := &krm.PscConnection{}
	out.PscConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	out.Address = direct.LazyPtr(in.GetAddress())
	out.ForwardingRule = direct.LazyPtr(in.GetForwardingRule())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func PscConnection_ToProto(mapCtx *direct.MapContext, in *krm.PscConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.PscConnectionId = direct.ValueOf(in.PscConnectionID)
	out.Address = direct.ValueOf(in.Address)
	out.ForwardingRule = direct.ValueOf(in.ForwardingRule)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Network = direct.ValueOf(in.Network)
	return out
}
func RedisClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.SizeGb = in.SizeGb
	out.DiscoveryEndpoints = direct.Slice_FromProto(mapCtx, in.DiscoveryEndpoints, DiscoveryEndpoint_FromProto)
	out.PscConnections = direct.Slice_FromProto(mapCtx, in.PscConnections, PscConnection_FromProto)
	out.StateInfo = Cluster_StateInfo_FromProto(mapCtx, in.GetStateInfo())
	out.PreciseSizeGb = in.PreciseSizeGb
	return out
}
func RedisClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RedisClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	out.SizeGb = in.SizeGb
	out.DiscoveryEndpoints = direct.Slice_ToProto(mapCtx, in.DiscoveryEndpoints, DiscoveryEndpoint_ToProto)
	out.PscConnections = direct.Slice_ToProto(mapCtx, in.PscConnections, PscConnection_ToProto)
	out.StateInfo = Cluster_StateInfo_ToProto(mapCtx, in.StateInfo)
	out.PreciseSizeGb = in.PreciseSizeGb
	return out
}
func RedisClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterSpec{}
	// MISSING: Name
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = in.ShardCount
	out.PscConfigs = direct.Slice_FromProto(mapCtx, in.PscConfigs, PscConfigSpec_FromProto)
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = ClusterPersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.RedisConfigs = in.RedisConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	return out
}
func RedisClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.RedisClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = in.ShardCount
	out.PscConfigs = direct.Slice_ToProto(mapCtx, in.PscConfigs, PscConfigSpec_ToProto)
	out.NodeType = direct.Enum_ToProto[pb.NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = ClusterPersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.RedisConfigs = in.RedisConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	return out
}
func ZoneDistributionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ZoneDistributionConfig) *krm.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ZoneDistributionConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.Zone = direct.LazyPtr(in.GetZone())
	return out
}
func ZoneDistributionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ZoneDistributionConfig) *pb.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ZoneDistributionConfig{}
	out.Mode = direct.Enum_ToProto[pb.ZoneDistributionConfig_ZoneDistributionMode](mapCtx, in.Mode)
	out.Zone = direct.ValueOf(in.Zone)
	return out
}
