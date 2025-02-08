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

package tpu

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/tpu/apiv2/tpupb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Topology = direct.LazyPtr(in.GetTopology())
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.Type = direct.Enum_ToProto[pb.AcceleratorConfig_Type](mapCtx, in.Type)
	out.Topology = direct.ValueOf(in.Topology)
	return out
}
func AccessConfig_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfig{}
	// MISSING: ExternalIP
	return out
}
func AccessConfig_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	// MISSING: ExternalIP
	return out
}
func AccessConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfigObservedState{}
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	return out
}
func AccessConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfigObservedState) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	return out
}
func AttachedDisk_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &krm.AttachedDisk{}
	out.SourceDisk = direct.LazyPtr(in.GetSourceDisk())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func AttachedDisk_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDisk) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	out.SourceDisk = direct.ValueOf(in.SourceDisk)
	out.Mode = direct.Enum_ToProto[pb.AttachedDisk_DiskMode](mapCtx, in.Mode)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	out.EnableExternalIps = direct.LazyPtr(in.GetEnableExternalIps())
	out.CanIPForward = direct.LazyPtr(in.GetCanIpForward())
	out.QueueCount = direct.LazyPtr(in.GetQueueCount())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	out.EnableExternalIps = direct.ValueOf(in.EnableExternalIps)
	out.CanIpForward = direct.ValueOf(in.CanIPForward)
	out.QueueCount = direct.ValueOf(in.QueueCount)
	return out
}
func NetworkEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krm.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.NetworkEndpoint{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	out.AccessConfig = AccessConfig_FromProto(mapCtx, in.GetAccessConfig())
	return out
}
func NetworkEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.NetworkEndpoint) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Port = direct.ValueOf(in.Port)
	out.AccessConfig = AccessConfig_ToProto(mapCtx, in.AccessConfig)
	return out
}
func NetworkEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krm.NetworkEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkEndpointObservedState{}
	// MISSING: IPAddress
	// MISSING: Port
	out.AccessConfig = AccessConfigObservedState_FromProto(mapCtx, in.GetAccessConfig())
	return out
}
func NetworkEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkEndpointObservedState) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	// MISSING: IPAddress
	// MISSING: Port
	out.AccessConfig = AccessConfigObservedState_ToProto(mapCtx, in.AccessConfig)
	return out
}
func Node_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.Node {
	if in == nil {
		return nil
	}
	out := &krm.Node{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorType())
	// MISSING: State
	// MISSING: HealthDescription
	out.RuntimeVersion = direct.LazyPtr(in.GetRuntimeVersion())
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.NetworkConfigs = direct.Slice_FromProto(mapCtx, in.NetworkConfigs, NetworkConfig_FromProto)
	out.CidrBlock = direct.LazyPtr(in.GetCidrBlock())
	out.ServiceAccount = ServiceAccount_FromProto(mapCtx, in.GetServiceAccount())
	// MISSING: CreateTime
	out.SchedulingConfig = SchedulingConfig_FromProto(mapCtx, in.GetSchedulingConfig())
	// MISSING: NetworkEndpoints
	out.Health = direct.Enum_FromProto(mapCtx, in.GetHealth())
	out.Labels = in.Labels
	out.Metadata = in.Metadata
	out.Tags = in.Tags
	// MISSING: ID
	out.DataDisks = direct.Slice_FromProto(mapCtx, in.DataDisks, AttachedDisk_FromProto)
	// MISSING: ApiVersion
	// MISSING: Symptoms
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.AcceleratorConfig = AcceleratorConfig_FromProto(mapCtx, in.GetAcceleratorConfig())
	// MISSING: QueuedResource
	// MISSING: MultisliceNode
	return out
}
func Node_ToProto(mapCtx *direct.MapContext, in *krm.Node) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.AcceleratorType = direct.ValueOf(in.AcceleratorType)
	// MISSING: State
	// MISSING: HealthDescription
	out.RuntimeVersion = direct.ValueOf(in.RuntimeVersion)
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.NetworkConfigs = direct.Slice_ToProto(mapCtx, in.NetworkConfigs, NetworkConfig_ToProto)
	out.CidrBlock = direct.ValueOf(in.CidrBlock)
	out.ServiceAccount = ServiceAccount_ToProto(mapCtx, in.ServiceAccount)
	// MISSING: CreateTime
	out.SchedulingConfig = SchedulingConfig_ToProto(mapCtx, in.SchedulingConfig)
	// MISSING: NetworkEndpoints
	out.Health = direct.Enum_ToProto[pb.Node_Health](mapCtx, in.Health)
	out.Labels = in.Labels
	out.Metadata = in.Metadata
	out.Tags = in.Tags
	// MISSING: ID
	out.DataDisks = direct.Slice_ToProto(mapCtx, in.DataDisks, AttachedDisk_ToProto)
	// MISSING: ApiVersion
	// MISSING: Symptoms
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.AcceleratorConfig = AcceleratorConfig_ToProto(mapCtx, in.AcceleratorConfig)
	// MISSING: QueuedResource
	// MISSING: MultisliceNode
	return out
}
func NodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.NodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NodeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: AcceleratorType
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.HealthDescription = direct.LazyPtr(in.GetHealthDescription())
	// MISSING: RuntimeVersion
	// MISSING: NetworkConfig
	// MISSING: NetworkConfigs
	// MISSING: CidrBlock
	// MISSING: ServiceAccount
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: SchedulingConfig
	out.NetworkEndpoints = direct.Slice_FromProto(mapCtx, in.NetworkEndpoints, NetworkEndpoint_FromProto)
	// MISSING: Health
	// MISSING: Labels
	// MISSING: Metadata
	// MISSING: Tags
	out.ID = direct.LazyPtr(in.GetId())
	// MISSING: DataDisks
	out.ApiVersion = direct.Enum_FromProto(mapCtx, in.GetApiVersion())
	out.Symptoms = direct.Slice_FromProto(mapCtx, in.Symptoms, Symptom_FromProto)
	// MISSING: ShieldedInstanceConfig
	// MISSING: AcceleratorConfig
	out.QueuedResource = direct.LazyPtr(in.GetQueuedResource())
	out.MultisliceNode = direct.LazyPtr(in.GetMultisliceNode())
	return out
}
func NodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NodeObservedState) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: AcceleratorType
	out.State = direct.Enum_ToProto[pb.Node_State](mapCtx, in.State)
	out.HealthDescription = direct.ValueOf(in.HealthDescription)
	// MISSING: RuntimeVersion
	// MISSING: NetworkConfig
	// MISSING: NetworkConfigs
	// MISSING: CidrBlock
	// MISSING: ServiceAccount
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: SchedulingConfig
	out.NetworkEndpoints = direct.Slice_ToProto(mapCtx, in.NetworkEndpoints, NetworkEndpoint_ToProto)
	// MISSING: Health
	// MISSING: Labels
	// MISSING: Metadata
	// MISSING: Tags
	out.Id = direct.ValueOf(in.ID)
	// MISSING: DataDisks
	out.ApiVersion = direct.Enum_ToProto[pb.Node_ApiVersion](mapCtx, in.ApiVersion)
	out.Symptoms = direct.Slice_ToProto(mapCtx, in.Symptoms, Symptom_ToProto)
	// MISSING: ShieldedInstanceConfig
	// MISSING: AcceleratorConfig
	out.QueuedResource = direct.ValueOf(in.QueuedResource)
	out.MultisliceNode = direct.ValueOf(in.MultisliceNode)
	return out
}
func QueuedResource_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource) *krm.QueuedResource {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource{}
	// MISSING: Name
	// MISSING: CreateTime
	out.Tpu = QueuedResource_Tpu_FromProto(mapCtx, in.GetTpu())
	out.Spot = QueuedResource_Spot_FromProto(mapCtx, in.GetSpot())
	out.Guaranteed = QueuedResource_Guaranteed_FromProto(mapCtx, in.GetGuaranteed())
	out.QueueingPolicy = QueuedResource_QueueingPolicy_FromProto(mapCtx, in.GetQueueingPolicy())
	// MISSING: State
	out.ReservationName = direct.LazyPtr(in.GetReservationName())
	return out
}
func QueuedResource_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource) *pb.QueuedResource {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource{}
	// MISSING: Name
	// MISSING: CreateTime
	if oneof := QueuedResource_Tpu_ToProto(mapCtx, in.Tpu); oneof != nil {
		out.Resource = &pb.QueuedResource_Tpu_{Tpu: oneof}
	}
	if oneof := QueuedResource_Spot_ToProto(mapCtx, in.Spot); oneof != nil {
		out.Tier = &pb.QueuedResource_Spot_{Spot: oneof}
	}
	if oneof := QueuedResource_Guaranteed_ToProto(mapCtx, in.Guaranteed); oneof != nil {
		out.Tier = &pb.QueuedResource_Guaranteed_{Guaranteed: oneof}
	}
	out.QueueingPolicy = QueuedResource_QueueingPolicy_ToProto(mapCtx, in.QueueingPolicy)
	// MISSING: State
	out.ReservationName = direct.ValueOf(in.ReservationName)
	return out
}
func QueuedResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource) *krm.QueuedResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Tpu = QueuedResource_TpuObservedState_FromProto(mapCtx, in.GetTpu())
	// MISSING: Spot
	// MISSING: Guaranteed
	// MISSING: QueueingPolicy
	out.State = QueuedResourceState_FromProto(mapCtx, in.GetState())
	// MISSING: ReservationName
	return out
}
func QueuedResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceObservedState) *pb.QueuedResource {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	if oneof := QueuedResource_TpuObservedState_ToProto(mapCtx, in.Tpu); oneof != nil {
		out.Resource = &pb.QueuedResource_Tpu_{Tpu: oneof}
	}
	// MISSING: Spot
	// MISSING: Guaranteed
	// MISSING: QueueingPolicy
	out.State = QueuedResourceState_ToProto(mapCtx, in.State)
	// MISSING: ReservationName
	return out
}
func QueuedResourceState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState) *krm.QueuedResourceState {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState{}
	// MISSING: State
	// MISSING: CreatingData
	// MISSING: AcceptedData
	// MISSING: ProvisioningData
	// MISSING: FailedData
	// MISSING: DeletingData
	// MISSING: ActiveData
	// MISSING: SuspendingData
	// MISSING: SuspendedData
	// MISSING: StateInitiator
	return out
}
func QueuedResourceState_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState) *pb.QueuedResourceState {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState{}
	// MISSING: State
	// MISSING: CreatingData
	// MISSING: AcceptedData
	// MISSING: ProvisioningData
	// MISSING: FailedData
	// MISSING: DeletingData
	// MISSING: ActiveData
	// MISSING: SuspendingData
	// MISSING: SuspendedData
	// MISSING: StateInitiator
	return out
}
func QueuedResourceStateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState) *krm.QueuedResourceStateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceStateObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreatingData = QueuedResourceState_CreatingData_FromProto(mapCtx, in.GetCreatingData())
	out.AcceptedData = QueuedResourceState_AcceptedData_FromProto(mapCtx, in.GetAcceptedData())
	out.ProvisioningData = QueuedResourceState_ProvisioningData_FromProto(mapCtx, in.GetProvisioningData())
	out.FailedData = QueuedResourceState_FailedData_FromProto(mapCtx, in.GetFailedData())
	out.DeletingData = QueuedResourceState_DeletingData_FromProto(mapCtx, in.GetDeletingData())
	out.ActiveData = QueuedResourceState_ActiveData_FromProto(mapCtx, in.GetActiveData())
	out.SuspendingData = QueuedResourceState_SuspendingData_FromProto(mapCtx, in.GetSuspendingData())
	out.SuspendedData = QueuedResourceState_SuspendedData_FromProto(mapCtx, in.GetSuspendedData())
	out.StateInitiator = direct.Enum_FromProto(mapCtx, in.GetStateInitiator())
	return out
}
func QueuedResourceStateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceStateObservedState) *pb.QueuedResourceState {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState{}
	out.State = direct.Enum_ToProto[pb.QueuedResourceState_State](mapCtx, in.State)
	if oneof := QueuedResourceState_CreatingData_ToProto(mapCtx, in.CreatingData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_CreatingData_{CreatingData: oneof}
	}
	if oneof := QueuedResourceState_AcceptedData_ToProto(mapCtx, in.AcceptedData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_AcceptedData_{AcceptedData: oneof}
	}
	if oneof := QueuedResourceState_ProvisioningData_ToProto(mapCtx, in.ProvisioningData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_ProvisioningData_{ProvisioningData: oneof}
	}
	if oneof := QueuedResourceState_FailedData_ToProto(mapCtx, in.FailedData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_FailedData_{FailedData: oneof}
	}
	if oneof := QueuedResourceState_DeletingData_ToProto(mapCtx, in.DeletingData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_DeletingData_{DeletingData: oneof}
	}
	if oneof := QueuedResourceState_ActiveData_ToProto(mapCtx, in.ActiveData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_ActiveData_{ActiveData: oneof}
	}
	if oneof := QueuedResourceState_SuspendingData_ToProto(mapCtx, in.SuspendingData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_SuspendingData_{SuspendingData: oneof}
	}
	if oneof := QueuedResourceState_SuspendedData_ToProto(mapCtx, in.SuspendedData); oneof != nil {
		out.StateData = &pb.QueuedResourceState_SuspendedData_{SuspendedData: oneof}
	}
	out.StateInitiator = direct.Enum_ToProto[pb.QueuedResourceState_StateInitiator](mapCtx, in.StateInitiator)
	return out
}
func QueuedResourceState_AcceptedData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_AcceptedData) *krm.QueuedResourceState_AcceptedData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_AcceptedData{}
	return out
}
func QueuedResourceState_AcceptedData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_AcceptedData) *pb.QueuedResourceState_AcceptedData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_AcceptedData{}
	return out
}
func QueuedResourceState_ActiveData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_ActiveData) *krm.QueuedResourceState_ActiveData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_ActiveData{}
	return out
}
func QueuedResourceState_ActiveData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_ActiveData) *pb.QueuedResourceState_ActiveData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_ActiveData{}
	return out
}
func QueuedResourceState_CreatingData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_CreatingData) *krm.QueuedResourceState_CreatingData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_CreatingData{}
	return out
}
func QueuedResourceState_CreatingData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_CreatingData) *pb.QueuedResourceState_CreatingData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_CreatingData{}
	return out
}
func QueuedResourceState_DeletingData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_DeletingData) *krm.QueuedResourceState_DeletingData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_DeletingData{}
	return out
}
func QueuedResourceState_DeletingData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_DeletingData) *pb.QueuedResourceState_DeletingData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_DeletingData{}
	return out
}
func QueuedResourceState_FailedData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_FailedData) *krm.QueuedResourceState_FailedData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_FailedData{}
	// MISSING: Error
	return out
}
func QueuedResourceState_FailedData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_FailedData) *pb.QueuedResourceState_FailedData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_FailedData{}
	// MISSING: Error
	return out
}
func QueuedResourceState_FailedDataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_FailedData) *krm.QueuedResourceState_FailedDataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_FailedDataObservedState{}
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func QueuedResourceState_FailedDataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_FailedDataObservedState) *pb.QueuedResourceState_FailedData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_FailedData{}
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func QueuedResourceState_ProvisioningData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_ProvisioningData) *krm.QueuedResourceState_ProvisioningData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_ProvisioningData{}
	return out
}
func QueuedResourceState_ProvisioningData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_ProvisioningData) *pb.QueuedResourceState_ProvisioningData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_ProvisioningData{}
	return out
}
func QueuedResourceState_SuspendedData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_SuspendedData) *krm.QueuedResourceState_SuspendedData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_SuspendedData{}
	return out
}
func QueuedResourceState_SuspendedData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_SuspendedData) *pb.QueuedResourceState_SuspendedData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_SuspendedData{}
	return out
}
func QueuedResourceState_SuspendingData_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResourceState_SuspendingData) *krm.QueuedResourceState_SuspendingData {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResourceState_SuspendingData{}
	return out
}
func QueuedResourceState_SuspendingData_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResourceState_SuspendingData) *pb.QueuedResourceState_SuspendingData {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResourceState_SuspendingData{}
	return out
}
func QueuedResource_Guaranteed_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_Guaranteed) *krm.QueuedResource_Guaranteed {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_Guaranteed{}
	out.MinDuration = direct.StringDuration_FromProto(mapCtx, in.GetMinDuration())
	return out
}
func QueuedResource_Guaranteed_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_Guaranteed) *pb.QueuedResource_Guaranteed {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_Guaranteed{}
	out.MinDuration = direct.StringDuration_ToProto(mapCtx, in.MinDuration)
	return out
}
func QueuedResource_QueueingPolicy_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_QueueingPolicy) *krm.QueuedResource_QueueingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_QueueingPolicy{}
	out.ValidUntilDuration = direct.StringDuration_FromProto(mapCtx, in.GetValidUntilDuration())
	out.ValidUntilTime = direct.StringTimestamp_FromProto(mapCtx, in.GetValidUntilTime())
	out.ValidAfterDuration = direct.StringDuration_FromProto(mapCtx, in.GetValidAfterDuration())
	out.ValidAfterTime = direct.StringTimestamp_FromProto(mapCtx, in.GetValidAfterTime())
	out.ValidInterval = Interval_FromProto(mapCtx, in.GetValidInterval())
	return out
}
func QueuedResource_QueueingPolicy_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_QueueingPolicy) *pb.QueuedResource_QueueingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_QueueingPolicy{}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.ValidUntilDuration); oneof != nil {
		out.StartTimingConstraints = &pb.QueuedResource_QueueingPolicy_ValidUntilDuration{ValidUntilDuration: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ValidUntilTime); oneof != nil {
		out.StartTimingConstraints = &pb.QueuedResource_QueueingPolicy_ValidUntilTime{ValidUntilTime: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.ValidAfterDuration); oneof != nil {
		out.StartTimingConstraints = &pb.QueuedResource_QueueingPolicy_ValidAfterDuration{ValidAfterDuration: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ValidAfterTime); oneof != nil {
		out.StartTimingConstraints = &pb.QueuedResource_QueueingPolicy_ValidAfterTime{ValidAfterTime: oneof}
	}
	if oneof := Interval_ToProto(mapCtx, in.ValidInterval); oneof != nil {
		out.StartTimingConstraints = &pb.QueuedResource_QueueingPolicy_ValidInterval{ValidInterval: oneof}
	}
	return out
}
func QueuedResource_Spot_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_Spot) *krm.QueuedResource_Spot {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_Spot{}
	return out
}
func QueuedResource_Spot_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_Spot) *pb.QueuedResource_Spot {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_Spot{}
	return out
}
func QueuedResource_Tpu_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_Tpu) *krm.QueuedResource_Tpu {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_Tpu{}
	out.NodeSpec = direct.Slice_FromProto(mapCtx, in.NodeSpec, QueuedResource_Tpu_NodeSpec_FromProto)
	return out
}
func QueuedResource_Tpu_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_Tpu) *pb.QueuedResource_Tpu {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_Tpu{}
	out.NodeSpec = direct.Slice_ToProto(mapCtx, in.NodeSpec, QueuedResource_Tpu_NodeSpec_ToProto)
	return out
}
func QueuedResource_TpuObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_Tpu) *krm.QueuedResource_TpuObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_TpuObservedState{}
	out.NodeSpec = direct.Slice_FromProto(mapCtx, in.NodeSpec, QueuedResource_Tpu_NodeSpecObservedState_FromProto)
	return out
}
func QueuedResource_TpuObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_TpuObservedState) *pb.QueuedResource_Tpu {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_Tpu{}
	out.NodeSpec = direct.Slice_ToProto(mapCtx, in.NodeSpec, QueuedResource_Tpu_NodeSpecObservedState_ToProto)
	return out
}
func QueuedResource_Tpu_NodeSpec_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_Tpu_NodeSpec) *krm.QueuedResource_Tpu_NodeSpec {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_Tpu_NodeSpec{}
	out.Parent = direct.LazyPtr(in.GetParent())
	out.NodeID = direct.LazyPtr(in.GetNodeId())
	out.MultisliceParams = QueuedResource_Tpu_NodeSpec_MultisliceParams_FromProto(mapCtx, in.GetMultisliceParams())
	out.Node = Node_FromProto(mapCtx, in.GetNode())
	return out
}
func QueuedResource_Tpu_NodeSpec_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_Tpu_NodeSpec) *pb.QueuedResource_Tpu_NodeSpec {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_Tpu_NodeSpec{}
	out.Parent = direct.ValueOf(in.Parent)
	if oneof := QueuedResource_Tpu_NodeSpec_NodeId_ToProto(mapCtx, in.NodeID); oneof != nil {
		out.NameStrategy = oneof
	}
	if oneof := QueuedResource_Tpu_NodeSpec_MultisliceParams_ToProto(mapCtx, in.MultisliceParams); oneof != nil {
		out.NameStrategy = &pb.QueuedResource_Tpu_NodeSpec_MultisliceParams_{MultisliceParams: oneof}
	}
	out.Node = Node_ToProto(mapCtx, in.Node)
	return out
}
func QueuedResource_Tpu_NodeSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_Tpu_NodeSpec) *krm.QueuedResource_Tpu_NodeSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_Tpu_NodeSpecObservedState{}
	// MISSING: Parent
	// MISSING: NodeID
	// MISSING: MultisliceParams
	out.Node = NodeObservedState_FromProto(mapCtx, in.GetNode())
	return out
}
func QueuedResource_Tpu_NodeSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_Tpu_NodeSpecObservedState) *pb.QueuedResource_Tpu_NodeSpec {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_Tpu_NodeSpec{}
	// MISSING: Parent
	// MISSING: NodeID
	// MISSING: MultisliceParams
	out.Node = NodeObservedState_ToProto(mapCtx, in.Node)
	return out
}
func QueuedResource_Tpu_NodeSpec_MultisliceParams_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource_Tpu_NodeSpec_MultisliceParams) *krm.QueuedResource_Tpu_NodeSpec_MultisliceParams {
	if in == nil {
		return nil
	}
	out := &krm.QueuedResource_Tpu_NodeSpec_MultisliceParams{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.NodeIDPrefix = direct.LazyPtr(in.GetNodeIdPrefix())
	return out
}
func QueuedResource_Tpu_NodeSpec_MultisliceParams_ToProto(mapCtx *direct.MapContext, in *krm.QueuedResource_Tpu_NodeSpec_MultisliceParams) *pb.QueuedResource_Tpu_NodeSpec_MultisliceParams {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource_Tpu_NodeSpec_MultisliceParams{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.NodeIdPrefix = direct.ValueOf(in.NodeIDPrefix)
	return out
}
func SchedulingConfig_FromProto(mapCtx *direct.MapContext, in *pb.SchedulingConfig) *krm.SchedulingConfig {
	if in == nil {
		return nil
	}
	out := &krm.SchedulingConfig{}
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	out.Reserved = direct.LazyPtr(in.GetReserved())
	out.Spot = direct.LazyPtr(in.GetSpot())
	return out
}
func SchedulingConfig_ToProto(mapCtx *direct.MapContext, in *krm.SchedulingConfig) *pb.SchedulingConfig {
	if in == nil {
		return nil
	}
	out := &pb.SchedulingConfig{}
	out.Preemptible = direct.ValueOf(in.Preemptible)
	out.Reserved = direct.ValueOf(in.Reserved)
	out.Spot = direct.ValueOf(in.Spot)
	return out
}
func ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccount{}
	out.Email = direct.LazyPtr(in.GetEmail())
	out.Scope = in.Scope
	return out
}
func ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	out.Email = direct.ValueOf(in.Email)
	out.Scope = in.Scope
	return out
}
func ShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	return out
}
func ShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	return out
}
func Symptom_FromProto(mapCtx *direct.MapContext, in *pb.Symptom) *krm.Symptom {
	if in == nil {
		return nil
	}
	out := &krm.Symptom{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.SymptomType = direct.Enum_FromProto(mapCtx, in.GetSymptomType())
	out.Details = direct.LazyPtr(in.GetDetails())
	out.WorkerID = direct.LazyPtr(in.GetWorkerId())
	return out
}
func Symptom_ToProto(mapCtx *direct.MapContext, in *krm.Symptom) *pb.Symptom {
	if in == nil {
		return nil
	}
	out := &pb.Symptom{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.SymptomType = direct.Enum_ToProto[pb.Symptom_SymptomType](mapCtx, in.SymptomType)
	out.Details = direct.ValueOf(in.Details)
	out.WorkerId = direct.ValueOf(in.WorkerID)
	return out
}
func TpuQueuedResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource) *krm.TpuQueuedResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TpuQueuedResourceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Tpu
	// MISSING: Spot
	// MISSING: Guaranteed
	// MISSING: QueueingPolicy
	// MISSING: State
	// MISSING: ReservationName
	return out
}
func TpuQueuedResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TpuQueuedResourceObservedState) *pb.QueuedResource {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Tpu
	// MISSING: Spot
	// MISSING: Guaranteed
	// MISSING: QueueingPolicy
	// MISSING: State
	// MISSING: ReservationName
	return out
}
func TpuQueuedResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource) *krm.TpuQueuedResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.TpuQueuedResourceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Tpu
	// MISSING: Spot
	// MISSING: Guaranteed
	// MISSING: QueueingPolicy
	// MISSING: State
	// MISSING: ReservationName
	return out
}
func TpuQueuedResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.TpuQueuedResourceSpec) *pb.QueuedResource {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Tpu
	// MISSING: Spot
	// MISSING: Guaranteed
	// MISSING: QueueingPolicy
	// MISSING: State
	// MISSING: ReservationName
	return out
}
