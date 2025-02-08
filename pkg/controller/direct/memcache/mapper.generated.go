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

package memcache

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/memcache/apiv1/memcachepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memcache/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.AuthorizedNetwork = direct.LazyPtr(in.GetAuthorizedNetwork())
	out.Zones = in.Zones
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.NodeConfig = Instance_NodeConfig_FromProto(mapCtx, in.GetNodeConfig())
	out.MemcacheVersion = direct.Enum_FromProto(mapCtx, in.GetMemcacheVersion())
	out.Parameters = MemcacheParameters_FromProto(mapCtx, in.GetParameters())
	// MISSING: MemcacheNodes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: MemcacheFullVersion
	out.InstanceMessages = direct.Slice_FromProto(mapCtx, in.InstanceMessages, Instance_InstanceMessage_FromProto)
	// MISSING: DiscoveryEndpoint
	out.MaintenancePolicy = MaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	// MISSING: MaintenanceSchedule
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.AuthorizedNetwork = direct.ValueOf(in.AuthorizedNetwork)
	out.Zones = in.Zones
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.NodeConfig = Instance_NodeConfig_ToProto(mapCtx, in.NodeConfig)
	out.MemcacheVersion = direct.Enum_ToProto[pb.MemcacheVersion](mapCtx, in.MemcacheVersion)
	out.Parameters = MemcacheParameters_ToProto(mapCtx, in.Parameters)
	// MISSING: MemcacheNodes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: MemcacheFullVersion
	out.InstanceMessages = direct.Slice_ToProto(mapCtx, in.InstanceMessages, Instance_InstanceMessage_ToProto)
	// MISSING: DiscoveryEndpoint
	out.MaintenancePolicy = MaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	// MISSING: MaintenanceSchedule
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AuthorizedNetwork
	// MISSING: Zones
	// MISSING: NodeCount
	// MISSING: NodeConfig
	// MISSING: MemcacheVersion
	out.Parameters = MemcacheParametersObservedState_FromProto(mapCtx, in.GetParameters())
	out.MemcacheNodes = direct.Slice_FromProto(mapCtx, in.MemcacheNodes, Instance_Node_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.MemcacheFullVersion = direct.LazyPtr(in.GetMemcacheFullVersion())
	// MISSING: InstanceMessages
	out.DiscoveryEndpoint = direct.LazyPtr(in.GetDiscoveryEndpoint())
	out.MaintenancePolicy = MaintenancePolicyObservedState_FromProto(mapCtx, in.GetMaintenancePolicy())
	out.MaintenanceSchedule = MaintenanceSchedule_FromProto(mapCtx, in.GetMaintenanceSchedule())
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AuthorizedNetwork
	// MISSING: Zones
	// MISSING: NodeCount
	// MISSING: NodeConfig
	// MISSING: MemcacheVersion
	out.Parameters = MemcacheParametersObservedState_ToProto(mapCtx, in.Parameters)
	out.MemcacheNodes = direct.Slice_ToProto(mapCtx, in.MemcacheNodes, Instance_Node_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.MemcacheFullVersion = direct.ValueOf(in.MemcacheFullVersion)
	// MISSING: InstanceMessages
	out.DiscoveryEndpoint = direct.ValueOf(in.DiscoveryEndpoint)
	out.MaintenancePolicy = MaintenancePolicyObservedState_ToProto(mapCtx, in.MaintenancePolicy)
	out.MaintenanceSchedule = MaintenanceSchedule_ToProto(mapCtx, in.MaintenanceSchedule)
	return out
}
func Instance_InstanceMessage_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceMessage) *krm.Instance_InstanceMessage {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InstanceMessage{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func Instance_InstanceMessage_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InstanceMessage) *pb.Instance_InstanceMessage {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceMessage{}
	out.Code = direct.Enum_ToProto[pb.Instance_InstanceMessage_Code](mapCtx, in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func Instance_Node_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krm.Instance_Node {
	if in == nil {
		return nil
	}
	out := &krm.Instance_Node{}
	// MISSING: NodeID
	// MISSING: Zone
	// MISSING: State
	// MISSING: Host
	// MISSING: Port
	out.Parameters = MemcacheParameters_FromProto(mapCtx, in.GetParameters())
	return out
}
func Instance_Node_ToProto(mapCtx *direct.MapContext, in *krm.Instance_Node) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	// MISSING: NodeID
	// MISSING: Zone
	// MISSING: State
	// MISSING: Host
	// MISSING: Port
	out.Parameters = MemcacheParameters_ToProto(mapCtx, in.Parameters)
	return out
}
func Instance_NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_NodeConfig) *krm.Instance_NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_NodeConfig{}
	out.CpuCount = direct.LazyPtr(in.GetCpuCount())
	out.MemorySizeMb = direct.LazyPtr(in.GetMemorySizeMb())
	return out
}
func Instance_NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_NodeConfig) *pb.Instance_NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_NodeConfig{}
	out.CpuCount = direct.ValueOf(in.CpuCount)
	out.MemorySizeMb = direct.ValueOf(in.MemorySizeMb)
	return out
}
func Instance_NodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krm.Instance_NodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_NodeObservedState{}
	out.NodeID = direct.LazyPtr(in.GetNodeId())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	// MISSING: Parameters
	return out
}
func Instance_NodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_NodeObservedState) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	out.NodeId = direct.ValueOf(in.NodeID)
	out.Zone = direct.ValueOf(in.Zone)
	out.State = direct.Enum_ToProto[pb.Instance_Node_State](mapCtx, in.State)
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	// MISSING: Parameters
	return out
}
func MaintenancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krm.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &krm.MaintenancePolicy{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.WeeklyMaintenanceWindow = direct.Slice_FromProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindow_FromProto)
	return out
}
func MaintenancePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MaintenancePolicy) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.WeeklyMaintenanceWindow = direct.Slice_ToProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindow_ToProto)
	return out
}
func MaintenancePolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krm.MaintenancePolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MaintenancePolicyObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: WeeklyMaintenanceWindow
	return out
}
func MaintenancePolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MaintenancePolicyObservedState) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: WeeklyMaintenanceWindow
	return out
}
func MaintenanceSchedule_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceSchedule{}
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ScheduleDeadlineTime
	return out
}
func MaintenanceSchedule_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceSchedule) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: ScheduleDeadlineTime
	return out
}
func MaintenanceScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceScheduleObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.ScheduleDeadlineTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleDeadlineTime())
	return out
}
func MaintenanceScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceScheduleObservedState) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.ScheduleDeadlineTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleDeadlineTime)
	return out
}
func MemcacheInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemcacheInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MemcacheInstanceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AuthorizedNetwork
	// MISSING: Zones
	// MISSING: NodeCount
	// MISSING: NodeConfig
	// MISSING: MemcacheVersion
	// MISSING: Parameters
	// MISSING: MemcacheNodes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: MemcacheFullVersion
	// MISSING: InstanceMessages
	// MISSING: DiscoveryEndpoint
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	return out
}
func MemcacheInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MemcacheInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AuthorizedNetwork
	// MISSING: Zones
	// MISSING: NodeCount
	// MISSING: NodeConfig
	// MISSING: MemcacheVersion
	// MISSING: Parameters
	// MISSING: MemcacheNodes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: MemcacheFullVersion
	// MISSING: InstanceMessages
	// MISSING: DiscoveryEndpoint
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	return out
}
func MemcacheInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemcacheInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MemcacheInstanceSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AuthorizedNetwork
	// MISSING: Zones
	// MISSING: NodeCount
	// MISSING: NodeConfig
	// MISSING: MemcacheVersion
	// MISSING: Parameters
	// MISSING: MemcacheNodes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: MemcacheFullVersion
	// MISSING: InstanceMessages
	// MISSING: DiscoveryEndpoint
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	return out
}
func MemcacheInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MemcacheInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AuthorizedNetwork
	// MISSING: Zones
	// MISSING: NodeCount
	// MISSING: NodeConfig
	// MISSING: MemcacheVersion
	// MISSING: Parameters
	// MISSING: MemcacheNodes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: MemcacheFullVersion
	// MISSING: InstanceMessages
	// MISSING: DiscoveryEndpoint
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	return out
}
func MemcacheParameters_FromProto(mapCtx *direct.MapContext, in *pb.MemcacheParameters) *krm.MemcacheParameters {
	if in == nil {
		return nil
	}
	out := &krm.MemcacheParameters{}
	// MISSING: ID
	out.Params = in.Params
	return out
}
func MemcacheParameters_ToProto(mapCtx *direct.MapContext, in *krm.MemcacheParameters) *pb.MemcacheParameters {
	if in == nil {
		return nil
	}
	out := &pb.MemcacheParameters{}
	// MISSING: ID
	out.Params = in.Params
	return out
}
func MemcacheParametersObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MemcacheParameters) *krm.MemcacheParametersObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MemcacheParametersObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	// MISSING: Params
	return out
}
func MemcacheParametersObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MemcacheParametersObservedState) *pb.MemcacheParameters {
	if in == nil {
		return nil
	}
	out := &pb.MemcacheParameters{}
	out.Id = direct.ValueOf(in.ID)
	// MISSING: Params
	return out
}
func WeeklyMaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.WeeklyMaintenanceWindow) *krm.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.WeeklyMaintenanceWindow{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}
func WeeklyMaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.WeeklyMaintenanceWindow) *pb.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.WeeklyMaintenanceWindow{}
	out.Day = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.Day)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}
