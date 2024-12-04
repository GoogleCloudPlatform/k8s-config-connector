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

package memorystore

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
)
func DiscoveryEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveryEndpoint) *krm.DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEndpoint{}
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func DiscoveryEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEndpoint) *pb.DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveryEndpoint{}
	out.Address = direct.ValueOf(in.Address)
	out.Port = direct.ValueOf(in.Port)
	out.Network = direct.ValueOf(in.Network)
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Labels = in.Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateInfo = Instance_StateInfo_FromProto(mapCtx, in.GetStateInfo())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	out.DiscoveryEndpoints = direct.Slice_FromProto(mapCtx, in.DiscoveryEndpoints, DiscoveryEndpoint_FromProto)
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = PersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.EngineConfigs = in.EngineConfigs
	out.NodeConfig = NodeConfig_FromProto(mapCtx, in.GetNodeConfig())
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	out.PscAutoConnections = direct.Slice_FromProto(mapCtx, in.PscAutoConnections, PscAutoConnection_FromProto)
	// MISSING: Endpoints
	// MISSING: Mode
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Labels = in.Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateInfo = Instance_StateInfo_ToProto(mapCtx, in.StateInfo)
	out.Uid = direct.ValueOf(in.Uid)
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.Instance_AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	out.DiscoveryEndpoints = direct.Slice_ToProto(mapCtx, in.DiscoveryEndpoints, DiscoveryEndpoint_ToProto)
	out.NodeType = direct.Enum_ToProto[pb.Instance_NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = PersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.EngineConfigs = in.EngineConfigs
	out.NodeConfig = NodeConfig_ToProto(mapCtx, in.NodeConfig)
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	out.PscAutoConnections = direct.Slice_ToProto(mapCtx, in.PscAutoConnections, PscAutoConnection_ToProto)
	// MISSING: Endpoints
	// MISSING: Mode
	return out
}
func Instance_StateInfo_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo) *krm.Instance_StateInfo {
	if in == nil {
		return nil
	}
	out := &krm.Instance_StateInfo{}
	out.UpdateInfo = Instance_StateInfo_UpdateInfo_FromProto(mapCtx, in.GetUpdateInfo())
	return out
}
func Instance_StateInfo_ToProto(mapCtx *direct.MapContext, in *krm.Instance_StateInfo) *pb.Instance_StateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo{}
	if oneof := Instance_StateInfo_UpdateInfo_ToProto(mapCtx, in.UpdateInfo); oneof != nil {
		out.Info = &pb.Instance_StateInfo_UpdateInfo_{UpdateInfo: oneof}
	}
	return out
}
func Instance_StateInfo_UpdateInfo_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo_UpdateInfo) *krm.Instance_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &krm.Instance_StateInfo_UpdateInfo{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func Instance_StateInfo_UpdateInfo_ToProto(mapCtx *direct.MapContext, in *krm.Instance_StateInfo_UpdateInfo) *pb.Instance_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo_UpdateInfo{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func MemorystoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: StateInfo
	// MISSING: Uid
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	// MISSING: DiscoveryEndpoints
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = PersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.EngineConfigs = in.EngineConfigs
	// MISSING: NodeConfig
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	// MISSING: PscAutoConnections
	// MISSING: Endpoints
	// MISSING: Mode
	return out
}
func MemorystoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: StateInfo
	// MISSING: Uid
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.Instance_AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	// MISSING: DiscoveryEndpoints
	out.NodeType = direct.Enum_ToProto[pb.Instance_NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = PersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.EngineConfigs = in.EngineConfigs
	// MISSING: NodeConfig
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	// MISSING: PscAutoConnections
	// MISSING: Endpoints
	// MISSING: Mode
	return out
}
func NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krm.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig{}
	out.SizeGb = direct.LazyPtr(in.GetSizeGb())
	return out
}
func NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.SizeGb = direct.ValueOf(in.SizeGb)
	return out
}
func PersistenceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig) *krm.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &krm.PersistenceConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.RdbConfig = PersistenceConfig_RDBConfig_FromProto(mapCtx, in.GetRdbConfig())
	out.AofConfig = PersistenceConfig_AOFConfig_FromProto(mapCtx, in.GetAofConfig())
	return out
}
func PersistenceConfig_ToProto(mapCtx *direct.MapContext, in *krm.PersistenceConfig) *pb.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig{}
	out.Mode = direct.Enum_ToProto[pb.PersistenceConfig_PersistenceMode](mapCtx, in.Mode)
	out.RdbConfig = PersistenceConfig_RDBConfig_ToProto(mapCtx, in.RdbConfig)
	out.AofConfig = PersistenceConfig_AOFConfig_ToProto(mapCtx, in.AofConfig)
	return out
}
func PersistenceConfig_AOFConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig_AOFConfig) *krm.PersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := &krm.PersistenceConfig_AOFConfig{}
	out.AppendFsync = direct.Enum_FromProto(mapCtx, in.GetAppendFsync())
	return out
}
func PersistenceConfig_AOFConfig_ToProto(mapCtx *direct.MapContext, in *krm.PersistenceConfig_AOFConfig) *pb.PersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig_AOFConfig{}
	out.AppendFsync = direct.Enum_ToProto[pb.PersistenceConfig_AOFConfig_AppendFsync](mapCtx, in.AppendFsync)
	return out
}
func PersistenceConfig_RDBConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig_RDBConfig) *krm.PersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := &krm.PersistenceConfig_RDBConfig{}
	out.RdbSnapshotPeriod = direct.Enum_FromProto(mapCtx, in.GetRdbSnapshotPeriod())
	out.RdbSnapshotStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRdbSnapshotStartTime())
	return out
}
func PersistenceConfig_RDBConfig_ToProto(mapCtx *direct.MapContext, in *krm.PersistenceConfig_RDBConfig) *pb.PersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig_RDBConfig{}
	out.RdbSnapshotPeriod = direct.Enum_ToProto[pb.PersistenceConfig_RDBConfig_SnapshotPeriod](mapCtx, in.RdbSnapshotPeriod)
	out.RdbSnapshotStartTime = direct.StringTimestamp_ToProto(mapCtx, in.RdbSnapshotStartTime)
	return out
}
func PscAutoConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnection) *krm.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &krm.PscAutoConnection{}
	out.Port = direct.LazyPtr(in.GetPort())
	out.PscConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	out.IpAddress = direct.LazyPtr(in.GetIpAddress())
	out.ForwardingRule = direct.LazyPtr(in.GetForwardingRule())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	out.PscConnectionStatus = direct.Enum_FromProto(mapCtx, in.GetPscConnectionStatus())
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func PscAutoConnection_ToProto(mapCtx *direct.MapContext, in *krm.PscAutoConnection) *pb.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscAutoConnection{}
	if oneof := PscAutoConnection_Port_ToProto(mapCtx, in.Port); oneof != nil {
		out.Ports = oneof
	}
	out.PscConnectionId = direct.ValueOf(in.PscConnectionID)
	out.IpAddress = direct.ValueOf(in.IpAddress)
	out.ForwardingRule = direct.ValueOf(in.ForwardingRule)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Network = direct.ValueOf(in.Network)
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	out.PscConnectionStatus = direct.Enum_ToProto[pb.PscConnectionStatus](mapCtx, in.PscConnectionStatus)
	out.ConnectionType = direct.Enum_ToProto[pb.ConnectionType](mapCtx, in.ConnectionType)
	return out
}
func ZoneDistributionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ZoneDistributionConfig) *krm.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ZoneDistributionConfig{}
	out.Zone = direct.LazyPtr(in.GetZone())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func ZoneDistributionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ZoneDistributionConfig) *pb.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ZoneDistributionConfig{}
	out.Zone = direct.ValueOf(in.Zone)
	out.Mode = direct.Enum_ToProto[pb.ZoneDistributionConfig_ZoneDistributionMode](mapCtx, in.Mode)
	return out
}
