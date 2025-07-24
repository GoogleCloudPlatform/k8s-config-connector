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
// krm.group: memorystore.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.memorystore.v1beta

package memorystore

import (
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
)
func DiscoveryEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveryEndpoint) *krmv1alpha1.DiscoveryEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DiscoveryEndpointObservedState{}
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func DiscoveryEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DiscoveryEndpointObservedState) *pb.DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveryEndpoint{}
	out.Address = direct.ValueOf(in.Address)
	out.Port = direct.ValueOf(in.Port)
	out.Network = direct.ValueOf(in.Network)
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmv1alpha1.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.InstanceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateInfo = Instance_StateInfo_FromProto(mapCtx, in.GetStateInfo())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: ReplicaCount
	// MISSING: AuthorizationMode
	// MISSING: TransitEncryptionMode
	// MISSING: ShardCount
	// MISSING: DiscoveryEndpoints
	// MISSING: NodeType
	// MISSING: PersistenceConfig
	// MISSING: EngineVersion
	// MISSING: EngineConfigs
	out.NodeConfig = NodeConfig_FromProto(mapCtx, in.GetNodeConfig())
	// MISSING: ZoneDistributionConfig
	// MISSING: DeletionProtectionEnabled
	// MISSING: PSCAutoConnections
	// (near miss): "PSCAutoConnections" vs "PscAutoConnections"
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Instance_InstanceEndpointObservedState_FromProto)
	// MISSING: Mode
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateInfo = Instance_StateInfo_ToProto(mapCtx, in.StateInfo)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: ReplicaCount
	// MISSING: AuthorizationMode
	// MISSING: TransitEncryptionMode
	// MISSING: ShardCount
	// MISSING: DiscoveryEndpoints
	// MISSING: NodeType
	// MISSING: PersistenceConfig
	// MISSING: EngineVersion
	// MISSING: EngineConfigs
	out.NodeConfig = NodeConfig_ToProto(mapCtx, in.NodeConfig)
	// MISSING: ZoneDistributionConfig
	// MISSING: DeletionProtectionEnabled
	// MISSING: PSCAutoConnections
	// (near miss): "PSCAutoConnections" vs "PscAutoConnections"
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Instance_InstanceEndpointObservedState_ToProto)
	// MISSING: Mode
	return out
}
func Instance_ConnectionDetail_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionDetail) *krmv1alpha1.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_ConnectionDetail{}
	// MISSING: PSCAutoConnection
	// (near miss): "PSCAutoConnection" vs "PscAutoConnection"
	// MISSING: PSCConnection
	// (near miss): "PSCConnection" vs "PscConnection"
	return out
}
func Instance_ConnectionDetail_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_ConnectionDetail) *pb.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionDetail{}
	// MISSING: PSCAutoConnection
	// (near miss): "PSCAutoConnection" vs "PscAutoConnection"
	// MISSING: PSCConnection
	// (near miss): "PSCConnection" vs "PscConnection"
	return out
}
func Instance_ConnectionDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionDetail) *krmv1alpha1.Instance_ConnectionDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_ConnectionDetailObservedState{}
	// MISSING: PSCAutoConnection
	// MISSING: PSCConnection
	// (near miss): "PSCConnection" vs "PscConnection"
	return out
}
func Instance_ConnectionDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_ConnectionDetailObservedState) *pb.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionDetail{}
	// MISSING: PSCAutoConnection
	// MISSING: PSCConnection
	// (near miss): "PSCConnection" vs "PscConnection"
	return out
}
func Instance_InstanceEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceEndpoint) *krmv1alpha1.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, Instance_ConnectionDetail_FromProto)
	return out
}
func Instance_InstanceEndpoint_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_InstanceEndpoint) *pb.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Instance_ConnectionDetail_ToProto)
	return out
}
func Instance_InstanceEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceEndpoint) *krmv1alpha1.Instance_InstanceEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_InstanceEndpointObservedState{}
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, Instance_ConnectionDetailObservedState_FromProto)
	return out
}
func Instance_InstanceEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_InstanceEndpointObservedState) *pb.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Instance_ConnectionDetailObservedState_ToProto)
	return out
}
func Instance_StateInfo_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo) *krmv1alpha1.Instance_StateInfo {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_StateInfo{}
	// MISSING: UpdateInfo
	return out
}
func Instance_StateInfo_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_StateInfo) *pb.Instance_StateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo{}
	// MISSING: UpdateInfo
	return out
}
func Instance_StateInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo) *krmv1alpha1.Instance_StateInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_StateInfoObservedState{}
	out.UpdateInfo = Instance_StateInfo_UpdateInfo_FromProto(mapCtx, in.GetUpdateInfo())
	return out
}
func Instance_StateInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_StateInfoObservedState) *pb.Instance_StateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo{}
	if oneof := Instance_StateInfo_UpdateInfo_ToProto(mapCtx, in.UpdateInfo); oneof != nil {
		out.Info = &pb.Instance_StateInfo_UpdateInfo_{UpdateInfo: oneof}
	}
	return out
}
func Instance_StateInfo_UpdateInfo_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo_UpdateInfo) *krmv1alpha1.Instance_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_StateInfo_UpdateInfo{}
	// MISSING: TargetShardCount
	// MISSING: TargetReplicaCount
	return out
}
func Instance_StateInfo_UpdateInfo_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_StateInfo_UpdateInfo) *pb.Instance_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo_UpdateInfo{}
	// MISSING: TargetShardCount
	// MISSING: TargetReplicaCount
	return out
}
func Instance_StateInfo_UpdateInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo_UpdateInfo) *krmv1alpha1.Instance_StateInfo_UpdateInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Instance_StateInfo_UpdateInfoObservedState{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func Instance_StateInfo_UpdateInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Instance_StateInfo_UpdateInfoObservedState) *pb.Instance_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo_UpdateInfo{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func MemorystoreInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmv1alpha1.MemorystoreInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MemorystoreInstanceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateInfo = Instance_StateInfoObservedState_FromProto(mapCtx, in.GetStateInfo())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: DiscoveryEndpoints
	out.NodeConfig = NodeConfigObservedState_FromProto(mapCtx, in.GetNodeConfig())
	// MISSING: PSCAutoConnections
	// (near miss): "PSCAutoConnections" vs "PscAutoConnections"
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Instance_InstanceEndpointObservedState_FromProto)
	return out
}
func MemorystoreInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MemorystoreInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateInfo = Instance_StateInfoObservedState_ToProto(mapCtx, in.StateInfo)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: DiscoveryEndpoints
	out.NodeConfig = NodeConfigObservedState_ToProto(mapCtx, in.NodeConfig)
	// MISSING: PSCAutoConnections
	// (near miss): "PSCAutoConnections" vs "PscAutoConnections"
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Instance_InstanceEndpointObservedState_ToProto)
	return out
}
func MemorystoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmv1alpha1.MemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MemorystoreInstanceSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	// MISSING: DiscoveryEndpoints
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = PersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.EngineConfigs = in.EngineConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	// MISSING: PSCAutoConnections
	// (near miss): "PSCAutoConnections" vs "PscAutoConnections"
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Instance_InstanceEndpoint_FromProto)
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func MemorystoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MemorystoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.Labels = in.Labels
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.Instance_AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	// MISSING: DiscoveryEndpoints
	out.NodeType = direct.Enum_ToProto[pb.Instance_NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = PersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.EngineConfigs = in.EngineConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	// MISSING: PSCAutoConnections
	// (near miss): "PSCAutoConnections" vs "PscAutoConnections"
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Instance_InstanceEndpoint_ToProto)
	out.Mode = direct.Enum_ToProto[pb.Instance_Mode](mapCtx, in.Mode)
	return out
}
func NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krmv1alpha1.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NodeConfig{}
	// MISSING: SizeGB
	return out
}
func NodeConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	// MISSING: SizeGB
	return out
}
func NodeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krmv1alpha1.NodeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NodeConfigObservedState{}
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	return out
}
func NodeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NodeConfigObservedState) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.SizeGb = direct.ValueOf(in.SizeGB)
	return out
}
func PersistenceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig) *krmv1alpha1.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PersistenceConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.RdbConfig = PersistenceConfig_RdbConfig_FromProto(mapCtx, in.GetRdbConfig())
	out.AofConfig = PersistenceConfig_AofConfig_FromProto(mapCtx, in.GetAofConfig())
	return out
}
func PersistenceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PersistenceConfig) *pb.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig{}
	out.Mode = direct.Enum_ToProto[pb.PersistenceConfig_PersistenceMode](mapCtx, in.Mode)
	out.RdbConfig = PersistenceConfig_RdbConfig_ToProto(mapCtx, in.RdbConfig)
	out.AofConfig = PersistenceConfig_AofConfig_ToProto(mapCtx, in.AofConfig)
	return out
}
func PersistenceConfig_AofConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig_AOFConfig) *krmv1alpha1.PersistenceConfig_AofConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PersistenceConfig_AofConfig{}
	out.AppendFsync = direct.Enum_FromProto(mapCtx, in.GetAppendFsync())
	return out
}
func PersistenceConfig_AofConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PersistenceConfig_AofConfig) *pb.PersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig_AOFConfig{}
	out.AppendFsync = direct.Enum_ToProto[pb.PersistenceConfig_AOFConfig_AppendFsync](mapCtx, in.AppendFsync)
	return out
}
func PersistenceConfig_RdbConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig_RDBConfig) *krmv1alpha1.PersistenceConfig_RdbConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PersistenceConfig_RdbConfig{}
	out.RdbSnapshotPeriod = direct.Enum_FromProto(mapCtx, in.GetRdbSnapshotPeriod())
	out.RdbSnapshotStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRdbSnapshotStartTime())
	return out
}
func PersistenceConfig_RdbConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PersistenceConfig_RdbConfig) *pb.PersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig_RDBConfig{}
	out.RdbSnapshotPeriod = direct.Enum_ToProto[pb.PersistenceConfig_RDBConfig_SnapshotPeriod](mapCtx, in.RdbSnapshotPeriod)
	out.RdbSnapshotStartTime = direct.StringTimestamp_ToProto(mapCtx, in.RdbSnapshotStartTime)
	return out
}
func PscAutoConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnection) *krmv1alpha1.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PscAutoConnection{}
	// MISSING: Port
	// MISSING: PSCConnectionID
	// MISSING: IPAddress
	// MISSING: ForwardingRule
	// MISSING: ProjectID
	if in.GetNetwork() != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	// MISSING: ServiceAttachment
	// MISSING: PSCConnectionStatus
	// MISSING: ConnectionType
	return out
}
func PscAutoConnection_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PscAutoConnection) *pb.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscAutoConnection{}
	// MISSING: Port
	// MISSING: PSCConnectionID
	// MISSING: IPAddress
	// MISSING: ForwardingRule
	// MISSING: ProjectID
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	// MISSING: ServiceAttachment
	// MISSING: PSCConnectionStatus
	// MISSING: ConnectionType
	return out
}
func PscAutoConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnection) *krmv1alpha1.PscAutoConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PscAutoConnectionObservedState{}
	out.Port = direct.LazyPtr(in.GetPort())
	// MISSING: PSCConnectionID
	// (near miss): "PSCConnectionID" vs "PscConnectionID"
	// MISSING: IPAddress
	// (near miss): "IPAddress" vs "IpAddress"
	out.ForwardingRule = direct.LazyPtr(in.GetForwardingRule())
	// MISSING: ProjectID
	// MISSING: Network
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	// MISSING: PSCConnectionStatus
	// (near miss): "PSCConnectionStatus" vs "PscConnectionStatus"
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func PscAutoConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PscAutoConnectionObservedState) *pb.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscAutoConnection{}
	if oneof := PscAutoConnectionObservedState_Port_ToProto(mapCtx, in.Port); oneof != nil {
		out.Ports = oneof
	}
	// MISSING: PSCConnectionID
	// (near miss): "PSCConnectionID" vs "PscConnectionID"
	// MISSING: IPAddress
	// (near miss): "IPAddress" vs "IpAddress"
	out.ForwardingRule = direct.ValueOf(in.ForwardingRule)
	// MISSING: ProjectID
	// MISSING: Network
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	// MISSING: PSCConnectionStatus
	// (near miss): "PSCConnectionStatus" vs "PscConnectionStatus"
	out.ConnectionType = direct.Enum_ToProto[pb.ConnectionType](mapCtx, in.ConnectionType)
	return out
}
func PscConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krmv1alpha1.PscConnection {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PscConnection{}
	// MISSING: PSCConnectionID
	// MISSING: IPAddress
	// (near miss): "IPAddress" vs "IpAddress"
	// MISSING: ForwardingRule
	// MISSING: ProjectID
	if in.GetNetwork() != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetServiceAttachment() != "" {
		out.ServiceAttachmentRef = &refs.ComputeServiceAttachmentRef{External: in.GetServiceAttachment()}
	}
	// MISSING: PSCConnectionStatus
	// MISSING: ConnectionType
	return out
}
func PscConnection_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PscConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	// MISSING: PSCConnectionID
	// MISSING: IPAddress
	// (near miss): "IPAddress" vs "IpAddress"
	// MISSING: ForwardingRule
	// MISSING: ProjectID
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.ServiceAttachmentRef != nil {
		out.ServiceAttachment = in.ServiceAttachmentRef.External
	}
	// MISSING: PSCConnectionStatus
	// MISSING: ConnectionType
	return out
}
func PscConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krmv1alpha1.PscConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PscConnectionObservedState{}
	// MISSING: PSCConnectionID
	// (near miss): "PSCConnectionID" vs "PscConnectionID"
	// MISSING: IPAddress
	// MISSING: ForwardingRule
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	// MISSING: Network
	// MISSING: ServiceAttachment
	// MISSING: PSCConnectionStatus
	// (near miss): "PSCConnectionStatus" vs "PscConnectionStatus"
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func PscConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PscConnectionObservedState) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	// MISSING: PSCConnectionID
	// (near miss): "PSCConnectionID" vs "PscConnectionID"
	// MISSING: IPAddress
	// MISSING: ForwardingRule
	out.ProjectId = direct.ValueOf(in.ProjectID)
	// MISSING: Network
	// MISSING: ServiceAttachment
	// MISSING: PSCConnectionStatus
	// (near miss): "PSCConnectionStatus" vs "PscConnectionStatus"
	out.ConnectionType = direct.Enum_ToProto[pb.ConnectionType](mapCtx, in.ConnectionType)
	return out
}
func ZoneDistributionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ZoneDistributionConfig) *krmv1alpha1.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ZoneDistributionConfig{}
	out.Zone = direct.LazyPtr(in.GetZone())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func ZoneDistributionConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ZoneDistributionConfig) *pb.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ZoneDistributionConfig{}
	out.Zone = direct.ValueOf(in.Zone)
	out.Mode = direct.Enum_ToProto[pb.ZoneDistributionConfig_ZoneDistributionMode](mapCtx, in.Mode)
	return out
}
