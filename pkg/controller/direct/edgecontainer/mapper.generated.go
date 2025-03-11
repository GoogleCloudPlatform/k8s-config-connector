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

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Authorization_FromProto(mapCtx *direct.MapContext, in *pb.Authorization) *krm.Authorization {
	if in == nil {
		return nil
	}
	out := &krm.Authorization{}
	out.AdminUsers = ClusterUser_FromProto(mapCtx, in.GetAdminUsers())
	return out
}
func Authorization_ToProto(mapCtx *direct.MapContext, in *krm.Authorization) *pb.Authorization {
	if in == nil {
		return nil
	}
	out := &pb.Authorization{}
	out.AdminUsers = ClusterUser_ToProto(mapCtx, in.AdminUsers)
	return out
}
func Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.Cluster {
	if in == nil {
		return nil
	}
	out := &krm.Cluster{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Fleet = Fleet_FromProto(mapCtx, in.GetFleet())
	out.Networking = ClusterNetworking_FromProto(mapCtx, in.GetNetworking())
	out.Authorization = Authorization_FromProto(mapCtx, in.GetAuthorization())
	out.DefaultMaxPodsPerNode = direct.LazyPtr(in.GetDefaultMaxPodsPerNode())
	// MISSING: Endpoint
	// MISSING: Port
	// MISSING: ClusterCACertificate
	out.MaintenancePolicy = MaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	// MISSING: ControlPlaneVersion
	// MISSING: NodeVersion
	out.ControlPlane = Cluster_ControlPlane_FromProto(mapCtx, in.GetControlPlane())
	out.SystemAddonsConfig = Cluster_SystemAddonsConfig_FromProto(mapCtx, in.GetSystemAddonsConfig())
	out.ExternalLoadBalancerIPV4AddressPools = in.ExternalLoadBalancerIpv4AddressPools
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryption_FromProto(mapCtx, in.GetControlPlaneEncryption())
	// MISSING: Status
	// MISSING: MaintenanceEvents
	out.TargetVersion = direct.LazyPtr(in.GetTargetVersion())
	out.ReleaseChannel = direct.Enum_FromProto(mapCtx, in.GetReleaseChannel())
	out.SurvivabilityConfig = Cluster_SurvivabilityConfig_FromProto(mapCtx, in.GetSurvivabilityConfig())
	out.ExternalLoadBalancerIPV6AddressPools = in.ExternalLoadBalancerIpv6AddressPools
	// MISSING: ConnectionState
	return out
}
func Cluster_ToProto(mapCtx *direct.MapContext, in *krm.Cluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Fleet = Fleet_ToProto(mapCtx, in.Fleet)
	out.Networking = ClusterNetworking_ToProto(mapCtx, in.Networking)
	out.Authorization = Authorization_ToProto(mapCtx, in.Authorization)
	out.DefaultMaxPodsPerNode = direct.ValueOf(in.DefaultMaxPodsPerNode)
	// MISSING: Endpoint
	// MISSING: Port
	// MISSING: ClusterCACertificate
	out.MaintenancePolicy = MaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	// MISSING: ControlPlaneVersion
	// MISSING: NodeVersion
	out.ControlPlane = Cluster_ControlPlane_ToProto(mapCtx, in.ControlPlane)
	out.SystemAddonsConfig = Cluster_SystemAddonsConfig_ToProto(mapCtx, in.SystemAddonsConfig)
	out.ExternalLoadBalancerIpv4AddressPools = in.ExternalLoadBalancerIPV4AddressPools
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryption_ToProto(mapCtx, in.ControlPlaneEncryption)
	// MISSING: Status
	// MISSING: MaintenanceEvents
	out.TargetVersion = direct.ValueOf(in.TargetVersion)
	out.ReleaseChannel = direct.Enum_ToProto[pb.Cluster_ReleaseChannel](mapCtx, in.ReleaseChannel)
	out.SurvivabilityConfig = Cluster_SurvivabilityConfig_ToProto(mapCtx, in.SurvivabilityConfig)
	out.ExternalLoadBalancerIpv6AddressPools = in.ExternalLoadBalancerIPV6AddressPools
	// MISSING: ConnectionState
	return out
}
func ClusterNetworking_FromProto(mapCtx *direct.MapContext, in *pb.ClusterNetworking) *krm.ClusterNetworking {
	if in == nil {
		return nil
	}
	out := &krm.ClusterNetworking{}
	out.ClusterIPV4CIDRBlocks = in.ClusterIpv4CidrBlocks
	out.ServicesIPV4CIDRBlocks = in.ServicesIpv4CidrBlocks
	return out
}
func ClusterNetworking_ToProto(mapCtx *direct.MapContext, in *krm.ClusterNetworking) *pb.ClusterNetworking {
	if in == nil {
		return nil
	}
	out := &pb.ClusterNetworking{}
	out.ClusterIpv4CidrBlocks = in.ClusterIPV4CIDRBlocks
	out.ServicesIpv4CidrBlocks = in.ServicesIPV4CIDRBlocks
	return out
}
func ClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClusterObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.Fleet = FleetObservedState_FromProto(mapCtx, in.GetFleet())
	// MISSING: Networking
	// MISSING: Authorization
	// MISSING: DefaultMaxPodsPerNode
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	out.Port = direct.LazyPtr(in.GetPort())
	out.ClusterCACertificate = direct.LazyPtr(in.GetClusterCaCertificate())
	// MISSING: MaintenancePolicy
	out.ControlPlaneVersion = direct.LazyPtr(in.GetControlPlaneVersion())
	out.NodeVersion = direct.LazyPtr(in.GetNodeVersion())
	// MISSING: ControlPlane
	// MISSING: SystemAddonsConfig
	// MISSING: ExternalLoadBalancerIPV4AddressPools
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryptionObservedState_FromProto(mapCtx, in.GetControlPlaneEncryption())
	out.Status = direct.Enum_FromProto(mapCtx, in.GetStatus())
	out.MaintenanceEvents = direct.Slice_FromProto(mapCtx, in.MaintenanceEvents, Cluster_MaintenanceEvent_FromProto)
	// MISSING: TargetVersion
	// MISSING: ReleaseChannel
	// MISSING: SurvivabilityConfig
	// MISSING: ExternalLoadBalancerIPV6AddressPools
	out.ConnectionState = Cluster_ConnectionState_FromProto(mapCtx, in.GetConnectionState())
	return out
}
func ClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.Fleet = FleetObservedState_ToProto(mapCtx, in.Fleet)
	// MISSING: Networking
	// MISSING: Authorization
	// MISSING: DefaultMaxPodsPerNode
	out.Endpoint = direct.ValueOf(in.Endpoint)
	out.Port = direct.ValueOf(in.Port)
	out.ClusterCaCertificate = direct.ValueOf(in.ClusterCACertificate)
	// MISSING: MaintenancePolicy
	out.ControlPlaneVersion = direct.ValueOf(in.ControlPlaneVersion)
	out.NodeVersion = direct.ValueOf(in.NodeVersion)
	// MISSING: ControlPlane
	// MISSING: SystemAddonsConfig
	// MISSING: ExternalLoadBalancerIPV4AddressPools
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryptionObservedState_ToProto(mapCtx, in.ControlPlaneEncryption)
	out.Status = direct.Enum_ToProto[pb.Cluster_Status](mapCtx, in.Status)
	out.MaintenanceEvents = direct.Slice_ToProto(mapCtx, in.MaintenanceEvents, Cluster_MaintenanceEvent_ToProto)
	// MISSING: TargetVersion
	// MISSING: ReleaseChannel
	// MISSING: SurvivabilityConfig
	// MISSING: ExternalLoadBalancerIPV6AddressPools
	out.ConnectionState = Cluster_ConnectionState_ToProto(mapCtx, in.ConnectionState)
	return out
}
func ClusterUser_FromProto(mapCtx *direct.MapContext, in *pb.ClusterUser) *krm.ClusterUser {
	if in == nil {
		return nil
	}
	out := &krm.ClusterUser{}
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func ClusterUser_ToProto(mapCtx *direct.MapContext, in *krm.ClusterUser) *pb.ClusterUser {
	if in == nil {
		return nil
	}
	out := &pb.ClusterUser{}
	out.Username = direct.ValueOf(in.Username)
	return out
}
func Cluster_ConnectionState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ConnectionState) *krm.Cluster_ConnectionState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ConnectionState{}
	// MISSING: State
	// MISSING: UpdateTime
	return out
}
func Cluster_ConnectionState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ConnectionState) *pb.Cluster_ConnectionState {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ConnectionState{}
	// MISSING: State
	// MISSING: UpdateTime
	return out
}
func Cluster_ConnectionStateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ConnectionState) *krm.Cluster_ConnectionStateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ConnectionStateObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Cluster_ConnectionStateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ConnectionStateObservedState) *pb.Cluster_ConnectionState {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ConnectionState{}
	out.State = direct.Enum_ToProto[pb.Cluster_ConnectionState_State](mapCtx, in.State)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Cluster_ControlPlane_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlane) *krm.Cluster_ControlPlane {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ControlPlane{}
	out.Remote = Cluster_ControlPlane_Remote_FromProto(mapCtx, in.GetRemote())
	out.Local = Cluster_ControlPlane_Local_FromProto(mapCtx, in.GetLocal())
	return out
}
func Cluster_ControlPlane_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ControlPlane) *pb.Cluster_ControlPlane {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlane{}
	if oneof := Cluster_ControlPlane_Remote_ToProto(mapCtx, in.Remote); oneof != nil {
		out.Config = &pb.Cluster_ControlPlane_Remote_{Remote: oneof}
	}
	if oneof := Cluster_ControlPlane_Local_ToProto(mapCtx, in.Local); oneof != nil {
		out.Config = &pb.Cluster_ControlPlane_Local_{Local: oneof}
	}
	return out
}
func Cluster_ControlPlaneEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlaneEncryption) *krm.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ControlPlaneEncryption{}
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	// MISSING: KMSKeyActiveVersion
	// MISSING: KMSKeyState
	// MISSING: KMSStatus
	// MISSING: ResourceState
	return out
}
func Cluster_ControlPlaneEncryption_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ControlPlaneEncryption) *pb.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlaneEncryption{}
	out.KmsKey = direct.ValueOf(in.KMSKey)
	// MISSING: KMSKeyActiveVersion
	// MISSING: KMSKeyState
	// MISSING: KMSStatus
	// MISSING: ResourceState
	return out
}
func Cluster_ControlPlaneEncryptionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlaneEncryption) *krm.Cluster_ControlPlaneEncryptionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ControlPlaneEncryptionObservedState{}
	// MISSING: KMSKey
	out.KMSKeyActiveVersion = direct.LazyPtr(in.GetKmsKeyActiveVersion())
	out.KMSKeyState = direct.Enum_FromProto(mapCtx, in.GetKmsKeyState())
	out.KMSStatus = Status_FromProto(mapCtx, in.GetKmsStatus())
	out.ResourceState = direct.Enum_FromProto(mapCtx, in.GetResourceState())
	return out
}
func Cluster_ControlPlaneEncryptionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ControlPlaneEncryptionObservedState) *pb.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlaneEncryption{}
	// MISSING: KMSKey
	out.KmsKeyActiveVersion = direct.ValueOf(in.KMSKeyActiveVersion)
	out.KmsKeyState = direct.Enum_ToProto[pb.KmsKeyState](mapCtx, in.KMSKeyState)
	out.KmsStatus = Status_ToProto(mapCtx, in.KMSStatus)
	out.ResourceState = direct.Enum_ToProto[pb.ResourceState](mapCtx, in.ResourceState)
	return out
}
func Cluster_ControlPlane_Local_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlane_Local) *krm.Cluster_ControlPlane_Local {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ControlPlane_Local{}
	out.NodeLocation = direct.LazyPtr(in.GetNodeLocation())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.MachineFilter = direct.LazyPtr(in.GetMachineFilter())
	out.SharedDeploymentPolicy = direct.Enum_FromProto(mapCtx, in.GetSharedDeploymentPolicy())
	out.ControlPlaneNodeStorageSchema = direct.LazyPtr(in.GetControlPlaneNodeStorageSchema())
	return out
}
func Cluster_ControlPlane_Local_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ControlPlane_Local) *pb.Cluster_ControlPlane_Local {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlane_Local{}
	out.NodeLocation = direct.ValueOf(in.NodeLocation)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.MachineFilter = direct.ValueOf(in.MachineFilter)
	out.SharedDeploymentPolicy = direct.Enum_ToProto[pb.Cluster_ControlPlane_SharedDeploymentPolicy](mapCtx, in.SharedDeploymentPolicy)
	out.ControlPlaneNodeStorageSchema = direct.ValueOf(in.ControlPlaneNodeStorageSchema)
	return out
}
func Cluster_ControlPlane_Remote_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlane_Remote) *krm.Cluster_ControlPlane_Remote {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ControlPlane_Remote{}
	return out
}
func Cluster_ControlPlane_Remote_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ControlPlane_Remote) *pb.Cluster_ControlPlane_Remote {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlane_Remote{}
	return out
}
func Cluster_MaintenanceEvent_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_MaintenanceEvent) *krm.Cluster_MaintenanceEvent {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_MaintenanceEvent{}
	// MISSING: Uuid
	// MISSING: TargetVersion
	// MISSING: Operation
	// MISSING: Type
	// MISSING: Schedule
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	return out
}
func Cluster_MaintenanceEvent_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_MaintenanceEvent) *pb.Cluster_MaintenanceEvent {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_MaintenanceEvent{}
	// MISSING: Uuid
	// MISSING: TargetVersion
	// MISSING: Operation
	// MISSING: Type
	// MISSING: Schedule
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	return out
}
func Cluster_MaintenanceEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_MaintenanceEvent) *krm.Cluster_MaintenanceEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_MaintenanceEventObservedState{}
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.TargetVersion = direct.LazyPtr(in.GetTargetVersion())
	out.Operation = direct.LazyPtr(in.GetOperation())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Schedule = direct.Enum_FromProto(mapCtx, in.GetSchedule())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Cluster_MaintenanceEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_MaintenanceEventObservedState) *pb.Cluster_MaintenanceEvent {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_MaintenanceEvent{}
	out.Uuid = direct.ValueOf(in.Uuid)
	out.TargetVersion = direct.ValueOf(in.TargetVersion)
	out.Operation = direct.ValueOf(in.Operation)
	out.Type = direct.Enum_ToProto[pb.Cluster_MaintenanceEvent_Type](mapCtx, in.Type)
	out.Schedule = direct.Enum_ToProto[pb.Cluster_MaintenanceEvent_Schedule](mapCtx, in.Schedule)
	out.State = direct.Enum_ToProto[pb.Cluster_MaintenanceEvent_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Cluster_SurvivabilityConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_SurvivabilityConfig) *krm.Cluster_SurvivabilityConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_SurvivabilityConfig{}
	out.OfflineRebootTtl = direct.StringDuration_FromProto(mapCtx, in.GetOfflineRebootTtl())
	return out
}
func Cluster_SurvivabilityConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_SurvivabilityConfig) *pb.Cluster_SurvivabilityConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_SurvivabilityConfig{}
	out.OfflineRebootTtl = direct.StringDuration_ToProto(mapCtx, in.OfflineRebootTtl)
	return out
}
func Cluster_SystemAddonsConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_SystemAddonsConfig) *krm.Cluster_SystemAddonsConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_SystemAddonsConfig{}
	out.Ingress = Cluster_SystemAddonsConfig_Ingress_FromProto(mapCtx, in.GetIngress())
	out.VmServiceConfig = Cluster_SystemAddonsConfig_VmServiceConfig_FromProto(mapCtx, in.GetVmServiceConfig())
	return out
}
func Cluster_SystemAddonsConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_SystemAddonsConfig) *pb.Cluster_SystemAddonsConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_SystemAddonsConfig{}
	out.Ingress = Cluster_SystemAddonsConfig_Ingress_ToProto(mapCtx, in.Ingress)
	out.VmServiceConfig = Cluster_SystemAddonsConfig_VmServiceConfig_ToProto(mapCtx, in.VmServiceConfig)
	return out
}
func Cluster_SystemAddonsConfig_Ingress_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_SystemAddonsConfig_Ingress) *krm.Cluster_SystemAddonsConfig_Ingress {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_SystemAddonsConfig_Ingress{}
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.IPV4Vip = direct.LazyPtr(in.GetIpv4Vip())
	return out
}
func Cluster_SystemAddonsConfig_Ingress_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_SystemAddonsConfig_Ingress) *pb.Cluster_SystemAddonsConfig_Ingress {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_SystemAddonsConfig_Ingress{}
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Ipv4Vip = direct.ValueOf(in.IPV4Vip)
	return out
}
func Cluster_SystemAddonsConfig_VmServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_SystemAddonsConfig_VMServiceConfig) *krm.Cluster_SystemAddonsConfig_VmServiceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_SystemAddonsConfig_VmServiceConfig{}
	out.VmmEnabled = direct.LazyPtr(in.GetVmmEnabled())
	return out
}
func Cluster_SystemAddonsConfig_VmServiceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_SystemAddonsConfig_VmServiceConfig) *pb.Cluster_SystemAddonsConfig_VMServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_SystemAddonsConfig_VMServiceConfig{}
	out.VmmEnabled = direct.ValueOf(in.VmmEnabled)
	return out
}
func EdgeContainerClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.EdgeContainerClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Fleet
	// MISSING: Networking
	// MISSING: Authorization
	// MISSING: DefaultMaxPodsPerNode
	// MISSING: Endpoint
	// MISSING: Port
	// MISSING: ClusterCACertificate
	// MISSING: MaintenancePolicy
	// MISSING: ControlPlaneVersion
	// MISSING: NodeVersion
	// MISSING: ControlPlane
	// MISSING: SystemAddonsConfig
	// MISSING: ExternalLoadBalancerIPV4AddressPools
	// MISSING: ControlPlaneEncryption
	// MISSING: Status
	// MISSING: MaintenanceEvents
	// MISSING: TargetVersion
	// MISSING: ReleaseChannel
	// MISSING: SurvivabilityConfig
	// MISSING: ExternalLoadBalancerIPV6AddressPools
	// MISSING: ConnectionState
	return out
}
func EdgeContainerClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Fleet
	// MISSING: Networking
	// MISSING: Authorization
	// MISSING: DefaultMaxPodsPerNode
	// MISSING: Endpoint
	// MISSING: Port
	// MISSING: ClusterCACertificate
	// MISSING: MaintenancePolicy
	// MISSING: ControlPlaneVersion
	// MISSING: NodeVersion
	// MISSING: ControlPlane
	// MISSING: SystemAddonsConfig
	// MISSING: ExternalLoadBalancerIPV4AddressPools
	// MISSING: ControlPlaneEncryption
	// MISSING: Status
	// MISSING: MaintenanceEvents
	// MISSING: TargetVersion
	// MISSING: ReleaseChannel
	// MISSING: SurvivabilityConfig
	// MISSING: ExternalLoadBalancerIPV6AddressPools
	// MISSING: ConnectionState
	return out
}
func EdgeContainerClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.EdgeContainerClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Fleet
	// MISSING: Networking
	// MISSING: Authorization
	// MISSING: DefaultMaxPodsPerNode
	// MISSING: Endpoint
	// MISSING: Port
	// MISSING: ClusterCACertificate
	// MISSING: MaintenancePolicy
	// MISSING: ControlPlaneVersion
	// MISSING: NodeVersion
	// MISSING: ControlPlane
	// MISSING: SystemAddonsConfig
	// MISSING: ExternalLoadBalancerIPV4AddressPools
	// MISSING: ControlPlaneEncryption
	// MISSING: Status
	// MISSING: MaintenanceEvents
	// MISSING: TargetVersion
	// MISSING: ReleaseChannel
	// MISSING: SurvivabilityConfig
	// MISSING: ExternalLoadBalancerIPV6AddressPools
	// MISSING: ConnectionState
	return out
}
func EdgeContainerClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Fleet
	// MISSING: Networking
	// MISSING: Authorization
	// MISSING: DefaultMaxPodsPerNode
	// MISSING: Endpoint
	// MISSING: Port
	// MISSING: ClusterCACertificate
	// MISSING: MaintenancePolicy
	// MISSING: ControlPlaneVersion
	// MISSING: NodeVersion
	// MISSING: ControlPlane
	// MISSING: SystemAddonsConfig
	// MISSING: ExternalLoadBalancerIPV4AddressPools
	// MISSING: ControlPlaneEncryption
	// MISSING: Status
	// MISSING: MaintenanceEvents
	// MISSING: TargetVersion
	// MISSING: ReleaseChannel
	// MISSING: SurvivabilityConfig
	// MISSING: ExternalLoadBalancerIPV6AddressPools
	// MISSING: ConnectionState
	return out
}
func Fleet_FromProto(mapCtx *direct.MapContext, in *pb.Fleet) *krm.Fleet {
	if in == nil {
		return nil
	}
	out := &krm.Fleet{}
	out.Project = direct.LazyPtr(in.GetProject())
	// MISSING: Membership
	return out
}
func Fleet_ToProto(mapCtx *direct.MapContext, in *krm.Fleet) *pb.Fleet {
	if in == nil {
		return nil
	}
	out := &pb.Fleet{}
	out.Project = direct.ValueOf(in.Project)
	// MISSING: Membership
	return out
}
func FleetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Fleet) *krm.FleetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FleetObservedState{}
	// MISSING: Project
	out.Membership = direct.LazyPtr(in.GetMembership())
	return out
}
func FleetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FleetObservedState) *pb.Fleet {
	if in == nil {
		return nil
	}
	out := &pb.Fleet{}
	// MISSING: Project
	out.Membership = direct.ValueOf(in.Membership)
	return out
}
func MaintenanceExclusionWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceExclusionWindow) *krm.MaintenanceExclusionWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceExclusionWindow{}
	out.Window = TimeWindow_FromProto(mapCtx, in.GetWindow())
	out.ID = direct.LazyPtr(in.GetId())
	return out
}
func MaintenanceExclusionWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceExclusionWindow) *pb.MaintenanceExclusionWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceExclusionWindow{}
	out.Window = TimeWindow_ToProto(mapCtx, in.Window)
	out.Id = direct.ValueOf(in.ID)
	return out
}
func MaintenancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krm.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &krm.MaintenancePolicy{}
	out.Window = MaintenanceWindow_FromProto(mapCtx, in.GetWindow())
	out.MaintenanceExclusions = direct.Slice_FromProto(mapCtx, in.MaintenanceExclusions, MaintenanceExclusionWindow_FromProto)
	return out
}
func MaintenancePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MaintenancePolicy) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	out.Window = MaintenanceWindow_ToProto(mapCtx, in.Window)
	out.MaintenanceExclusions = direct.Slice_ToProto(mapCtx, in.MaintenanceExclusions, MaintenanceExclusionWindow_ToProto)
	return out
}
func MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceWindow) *krm.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceWindow{}
	out.RecurringWindow = RecurringTimeWindow_FromProto(mapCtx, in.GetRecurringWindow())
	return out
}
func MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceWindow) *pb.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceWindow{}
	out.RecurringWindow = RecurringTimeWindow_ToProto(mapCtx, in.RecurringWindow)
	return out
}
func RecurringTimeWindow_FromProto(mapCtx *direct.MapContext, in *pb.RecurringTimeWindow) *krm.RecurringTimeWindow {
	if in == nil {
		return nil
	}
	out := &krm.RecurringTimeWindow{}
	out.Window = TimeWindow_FromProto(mapCtx, in.GetWindow())
	out.Recurrence = direct.LazyPtr(in.GetRecurrence())
	return out
}
func RecurringTimeWindow_ToProto(mapCtx *direct.MapContext, in *krm.RecurringTimeWindow) *pb.RecurringTimeWindow {
	if in == nil {
		return nil
	}
	out := &pb.RecurringTimeWindow{}
	out.Window = TimeWindow_ToProto(mapCtx, in.Window)
	out.Recurrence = direct.ValueOf(in.Recurrence)
	return out
}
func TimeWindow_FromProto(mapCtx *direct.MapContext, in *pb.TimeWindow) *krm.TimeWindow {
	if in == nil {
		return nil
	}
	out := &krm.TimeWindow{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func TimeWindow_ToProto(mapCtx *direct.MapContext, in *krm.TimeWindow) *pb.TimeWindow {
	if in == nil {
		return nil
	}
	out := &pb.TimeWindow{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
